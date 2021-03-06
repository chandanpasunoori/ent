// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/json/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL *url.URL `json:"url,omitempty"`
	// Raw holds the value of the "raw" field.
	Raw json.RawMessage `json:"raw,omitempty"`
	// Dirs holds the value of the "dirs" field.
	Dirs []http.Dir `json:"dirs,omitempty"`
	// Ints holds the value of the "ints" field.
	Ints []int `json:"ints,omitempty"`
	// Floats holds the value of the "floats" field.
	Floats []float64 `json:"floats,omitempty"`
	// Strings holds the value of the "strings" field.
	Strings []string `json:"strings,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},
		&[]byte{},
		&[]byte{},
		&[]byte{},
		&[]byte{},
		&[]byte{},
		&[]byte{},
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m != n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]

	if value, ok := values[0].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field url", values[0])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.URL); err != nil {
			return fmt.Errorf("unmarshal field url: %v", err)
		}
	}

	if value, ok := values[1].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field raw", values[1])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.Raw); err != nil {
			return fmt.Errorf("unmarshal field raw: %v", err)
		}
	}

	if value, ok := values[2].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field dirs", values[2])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.Dirs); err != nil {
			return fmt.Errorf("unmarshal field dirs: %v", err)
		}
	}

	if value, ok := values[3].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field ints", values[3])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.Ints); err != nil {
			return fmt.Errorf("unmarshal field ints: %v", err)
		}
	}

	if value, ok := values[4].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field floats", values[4])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.Floats); err != nil {
			return fmt.Errorf("unmarshal field floats: %v", err)
		}
	}

	if value, ok := values[5].(*[]byte); !ok {
		return fmt.Errorf("unexpected type %T for field strings", values[5])
	} else if value != nil && len(*value) > 0 {
		if err := json.Unmarshal(*value, &u.Strings); err != nil {
			return fmt.Errorf("unmarshal field strings: %v", err)
		}
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", url=")
	builder.WriteString(fmt.Sprintf("%v", u.URL))
	builder.WriteString(", raw=")
	builder.WriteString(fmt.Sprintf("%v", u.Raw))
	builder.WriteString(", dirs=")
	builder.WriteString(fmt.Sprintf("%v", u.Dirs))
	builder.WriteString(", ints=")
	builder.WriteString(fmt.Sprintf("%v", u.Ints))
	builder.WriteString(", floats=")
	builder.WriteString(fmt.Sprintf("%v", u.Floats))
	builder.WriteString(", strings=")
	builder.WriteString(fmt.Sprintf("%v", u.Strings))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
