// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGroupUpdateNameResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GroupUpdateNameResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GroupUpdateNameResponse
	GetStatusCode() *int32
}

type GroupUpdateNameResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GroupUpdateNameResponse) String() string {
	return dara.Prettify(s)
}

func (s GroupUpdateNameResponse) GoString() string {
	return s.String()
}

func (s *GroupUpdateNameResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GroupUpdateNameResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GroupUpdateNameResponse) SetHeaders(v map[string]*string) *GroupUpdateNameResponse {
	s.Headers = v
	return s
}

func (s *GroupUpdateNameResponse) SetStatusCode(v int32) *GroupUpdateNameResponse {
	s.StatusCode = &v
	return s
}

func (s *GroupUpdateNameResponse) Validate() error {
	return dara.Validate(s)
}
