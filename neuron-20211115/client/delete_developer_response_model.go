// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDeveloperResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteDeveloperResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteDeveloperResponse
	GetStatusCode() *int32
}

type DeleteDeveloperResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteDeveloperResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteDeveloperResponse) GoString() string {
	return s.String()
}

func (s *DeleteDeveloperResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteDeveloperResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteDeveloperResponse) SetHeaders(v map[string]*string) *DeleteDeveloperResponse {
	s.Headers = v
	return s
}

func (s *DeleteDeveloperResponse) SetStatusCode(v int32) *DeleteDeveloperResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteDeveloperResponse) Validate() error {
	return dara.Validate(s)
}
