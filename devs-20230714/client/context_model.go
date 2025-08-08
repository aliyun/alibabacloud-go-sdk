// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContext interface {
	dara.Model
	String() string
	GoString() string
	SetData(v map[string]interface{}) *Context
	GetData() map[string]interface{}
}

type Context struct {
	Data map[string]interface{} `json:"data,omitempty" xml:"data,omitempty"`
}

func (s Context) String() string {
	return dara.Prettify(s)
}

func (s Context) GoString() string {
	return s.String()
}

func (s *Context) GetData() map[string]interface{} {
	return s.Data
}

func (s *Context) SetData(v map[string]interface{}) *Context {
	s.Data = v
	return s
}

func (s *Context) Validate() error {
	return dara.Validate(s)
}
