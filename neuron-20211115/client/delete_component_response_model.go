// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteComponentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteComponentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteComponentResponse
	GetStatusCode() *int32
}

type DeleteComponentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteComponentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteComponentResponse) GoString() string {
	return s.String()
}

func (s *DeleteComponentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteComponentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteComponentResponse) SetHeaders(v map[string]*string) *DeleteComponentResponse {
	s.Headers = v
	return s
}

func (s *DeleteComponentResponse) SetStatusCode(v int32) *DeleteComponentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteComponentResponse) Validate() error {
	return dara.Validate(s)
}
