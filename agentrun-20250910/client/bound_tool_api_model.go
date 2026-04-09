// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBoundToolApi interface {
	dara.Model
	String() string
	GoString() string
	SetMethod(v string) *BoundToolApi
	GetMethod() *string
	SetPath(v string) *BoundToolApi
	GetPath() *string
}

type BoundToolApi struct {
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	Path   *string `json:"path,omitempty" xml:"path,omitempty"`
}

func (s BoundToolApi) String() string {
	return dara.Prettify(s)
}

func (s BoundToolApi) GoString() string {
	return s.String()
}

func (s *BoundToolApi) GetMethod() *string {
	return s.Method
}

func (s *BoundToolApi) GetPath() *string {
	return s.Path
}

func (s *BoundToolApi) SetMethod(v string) *BoundToolApi {
	s.Method = &v
	return s
}

func (s *BoundToolApi) SetPath(v string) *BoundToolApi {
	s.Path = &v
	return s
}

func (s *BoundToolApi) Validate() error {
	return dara.Validate(s)
}
