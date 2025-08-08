// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEnvironmentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteEnvironmentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteEnvironmentResponse
	GetStatusCode() *int32
}

type DeleteEnvironmentResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteEnvironmentResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteEnvironmentResponse) GoString() string {
	return s.String()
}

func (s *DeleteEnvironmentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteEnvironmentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteEnvironmentResponse) SetHeaders(v map[string]*string) *DeleteEnvironmentResponse {
	s.Headers = v
	return s
}

func (s *DeleteEnvironmentResponse) SetStatusCode(v int32) *DeleteEnvironmentResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteEnvironmentResponse) Validate() error {
	return dara.Validate(s)
}
