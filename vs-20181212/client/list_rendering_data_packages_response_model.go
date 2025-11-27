// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListRenderingDataPackagesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListRenderingDataPackagesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListRenderingDataPackagesResponse
	GetStatusCode() *int32
	SetBody(v *ListRenderingDataPackagesResponseBody) *ListRenderingDataPackagesResponse
	GetBody() *ListRenderingDataPackagesResponseBody
}

type ListRenderingDataPackagesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRenderingDataPackagesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRenderingDataPackagesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListRenderingDataPackagesResponse) GoString() string {
	return s.String()
}

func (s *ListRenderingDataPackagesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListRenderingDataPackagesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListRenderingDataPackagesResponse) GetBody() *ListRenderingDataPackagesResponseBody {
	return s.Body
}

func (s *ListRenderingDataPackagesResponse) SetHeaders(v map[string]*string) *ListRenderingDataPackagesResponse {
	s.Headers = v
	return s
}

func (s *ListRenderingDataPackagesResponse) SetStatusCode(v int32) *ListRenderingDataPackagesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRenderingDataPackagesResponse) SetBody(v *ListRenderingDataPackagesResponseBody) *ListRenderingDataPackagesResponse {
	s.Body = v
	return s
}

func (s *ListRenderingDataPackagesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
