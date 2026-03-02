// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMetadataInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMetadataInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMetadataInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListMetadataInfosResponseBody) *ListMetadataInfosResponse
	GetBody() *ListMetadataInfosResponseBody
}

type ListMetadataInfosResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMetadataInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMetadataInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMetadataInfosResponse) GoString() string {
	return s.String()
}

func (s *ListMetadataInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMetadataInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMetadataInfosResponse) GetBody() *ListMetadataInfosResponseBody {
	return s.Body
}

func (s *ListMetadataInfosResponse) SetHeaders(v map[string]*string) *ListMetadataInfosResponse {
	s.Headers = v
	return s
}

func (s *ListMetadataInfosResponse) SetStatusCode(v int32) *ListMetadataInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMetadataInfosResponse) SetBody(v *ListMetadataInfosResponseBody) *ListMetadataInfosResponse {
	s.Body = v
	return s
}

func (s *ListMetadataInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
