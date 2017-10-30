package entity
import (
    "testing"
    "reflect"
)
var t_users = []User {
        {"a", "a", "a", "a"},
        {"b", "b", "b", "b"},
        {"c", "c", "c", "c"},
      }
func init() {
  userlist = nil
}
func TestCreateUser(t *testing.T) {
   
    cases := []struct {
        in User
        want []User
    }{
        {t_users[0], t_users[:1]},
        {t_users[1], t_users[:2]},
        {t_users[2], t_users[:3]},
    }
    for _, c := range cases {
        createUser(c.in)
        if got := userlist; !reflect.DeepEqual(got, c.want)  {
            t.Errorf("CreateUser(%q) == %q, want %q", c.in, got, c.want)
        }
    }
}

