// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHTTPHeader interface {
	dara.Model
	String() string
	GoString() string
	SetName(v string) *HTTPHeader
	GetName() *string
	SetValue(v string) *HTTPHeader
	GetValue() *string
}

type HTTPHeader struct {
	Name  *string `json:"name,omitempty" xml:"name,omitempty"`
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s HTTPHeader) String() string {
	return dara.Prettify(s)
}

func (s HTTPHeader) GoString() string {
	return s.String()
}

func (s *HTTPHeader) GetName() *string {
	return s.Name
}

func (s *HTTPHeader) GetValue() *string {
	return s.Value
}

func (s *HTTPHeader) SetName(v string) *HTTPHeader {
	s.Name = &v
	return s
}

func (s *HTTPHeader) SetValue(v string) *HTTPHeader {
	s.Value = &v
	return s
}

func (s *HTTPHeader) Validate() error {
	return dara.Validate(s)
}
