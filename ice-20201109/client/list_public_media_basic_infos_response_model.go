// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPublicMediaBasicInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListPublicMediaBasicInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListPublicMediaBasicInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListPublicMediaBasicInfosResponseBody) *ListPublicMediaBasicInfosResponse
	GetBody() *ListPublicMediaBasicInfosResponseBody
}

type ListPublicMediaBasicInfosResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListPublicMediaBasicInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListPublicMediaBasicInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListPublicMediaBasicInfosResponse) GoString() string {
	return s.String()
}

func (s *ListPublicMediaBasicInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListPublicMediaBasicInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListPublicMediaBasicInfosResponse) GetBody() *ListPublicMediaBasicInfosResponseBody {
	return s.Body
}

func (s *ListPublicMediaBasicInfosResponse) SetHeaders(v map[string]*string) *ListPublicMediaBasicInfosResponse {
	s.Headers = v
	return s
}

func (s *ListPublicMediaBasicInfosResponse) SetStatusCode(v int32) *ListPublicMediaBasicInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListPublicMediaBasicInfosResponse) SetBody(v *ListPublicMediaBasicInfosResponseBody) *ListPublicMediaBasicInfosResponse {
	s.Body = v
	return s
}

func (s *ListPublicMediaBasicInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
