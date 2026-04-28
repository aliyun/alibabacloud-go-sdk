// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRevisionResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRevisionResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRevisionResponse
	GetStatusCode() *int32
}

type DeleteRevisionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteRevisionResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRevisionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRevisionResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRevisionResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRevisionResponse) SetHeaders(v map[string]*string) *DeleteRevisionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRevisionResponse) SetStatusCode(v int32) *DeleteRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRevisionResponse) Validate() error {
	return dara.Validate(s)
}
