// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMessageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteMessageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteMessageResponse
	GetStatusCode() *int32
}

type DeleteMessageResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteMessageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteMessageResponse) GoString() string {
	return s.String()
}

func (s *DeleteMessageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteMessageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteMessageResponse) SetHeaders(v map[string]*string) *DeleteMessageResponse {
	s.Headers = v
	return s
}

func (s *DeleteMessageResponse) SetStatusCode(v int32) *DeleteMessageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteMessageResponse) Validate() error {
	return dara.Validate(s)
}
