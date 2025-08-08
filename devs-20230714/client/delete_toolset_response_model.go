// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteToolsetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteToolsetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteToolsetResponse
	GetStatusCode() *int32
}

type DeleteToolsetResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteToolsetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteToolsetResponse) GoString() string {
	return s.String()
}

func (s *DeleteToolsetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteToolsetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteToolsetResponse) SetHeaders(v map[string]*string) *DeleteToolsetResponse {
	s.Headers = v
	return s
}

func (s *DeleteToolsetResponse) SetStatusCode(v int32) *DeleteToolsetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteToolsetResponse) Validate() error {
	return dara.Validate(s)
}
