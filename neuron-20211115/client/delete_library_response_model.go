// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLibraryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLibraryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLibraryResponse
	GetStatusCode() *int32
}

type DeleteLibraryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteLibraryResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLibraryResponse) GoString() string {
	return s.String()
}

func (s *DeleteLibraryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLibraryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLibraryResponse) SetHeaders(v map[string]*string) *DeleteLibraryResponse {
	s.Headers = v
	return s
}

func (s *DeleteLibraryResponse) SetStatusCode(v int32) *DeleteLibraryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLibraryResponse) Validate() error {
	return dara.Validate(s)
}
