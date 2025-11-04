// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListHotwordLibrariesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListHotwordLibrariesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListHotwordLibrariesResponse
	GetStatusCode() *int32
	SetBody(v *ListHotwordLibrariesResponseBody) *ListHotwordLibrariesResponse
	GetBody() *ListHotwordLibrariesResponseBody
}

type ListHotwordLibrariesResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListHotwordLibrariesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListHotwordLibrariesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListHotwordLibrariesResponse) GoString() string {
	return s.String()
}

func (s *ListHotwordLibrariesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListHotwordLibrariesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListHotwordLibrariesResponse) GetBody() *ListHotwordLibrariesResponseBody {
	return s.Body
}

func (s *ListHotwordLibrariesResponse) SetHeaders(v map[string]*string) *ListHotwordLibrariesResponse {
	s.Headers = v
	return s
}

func (s *ListHotwordLibrariesResponse) SetStatusCode(v int32) *ListHotwordLibrariesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListHotwordLibrariesResponse) SetBody(v *ListHotwordLibrariesResponseBody) *ListHotwordLibrariesResponse {
	s.Body = v
	return s
}

func (s *ListHotwordLibrariesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
