// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListMediaBasicInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListMediaBasicInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListMediaBasicInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListMediaBasicInfosResponseBody) *ListMediaBasicInfosResponse
	GetBody() *ListMediaBasicInfosResponseBody
}

type ListMediaBasicInfosResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMediaBasicInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMediaBasicInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListMediaBasicInfosResponse) GoString() string {
	return s.String()
}

func (s *ListMediaBasicInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListMediaBasicInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListMediaBasicInfosResponse) GetBody() *ListMediaBasicInfosResponseBody {
	return s.Body
}

func (s *ListMediaBasicInfosResponse) SetHeaders(v map[string]*string) *ListMediaBasicInfosResponse {
	s.Headers = v
	return s
}

func (s *ListMediaBasicInfosResponse) SetStatusCode(v int32) *ListMediaBasicInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMediaBasicInfosResponse) SetBody(v *ListMediaBasicInfosResponseBody) *ListMediaBasicInfosResponse {
	s.Body = v
	return s
}

func (s *ListMediaBasicInfosResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
