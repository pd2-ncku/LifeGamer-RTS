package protocal

var generator (func() int)
var chans map[string]chan string

func init() {
    generator = func() (func() int) {
        var i int = -1
        return func() int {
            i++
            return i
        }
    }()

    chans = make(map[string]chan string)
}

type Client struct {
    cid int
    name string
}

func NewClient(name string) Client {
    new_id := generator()

    chans[name] = make(chan string, 100)
    return Client { new_id, name }
}

func (c Client) Get() string {
    return <- chans[c.name]
}

func (c Client) Put(dst string, msg string) (ok bool) {
    dstchan, ok := chans[dst]
    if ok {
        dstchan <- msg
    }

    return ok
}
