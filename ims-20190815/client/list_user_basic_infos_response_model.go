// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserBasicInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserBasicInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserBasicInfosResponse
	GetStatusCode() *int32
	SetBody(v *ListUserBasicInfosResponseBody) *ListUserBasicInfosResponse
	GetBody() *ListUserBasicInfosResponseBody
}

type ListUserBasicInfosResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserBasicInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserBasicInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserBasicInfosResponse) GoString() string {
	return s.String()
}

func (s *ListUserBasicInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserBasicInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserBasicInfosResponse) GetBody() *ListUserBasicInfosResponseBody {
	return s.Body
}

func (s *ListUserBasicInfosResponse) SetHeaders(v map[string]*string) *ListUserBasicInfosResponse {
	s.Headers = v
	return s
}

func (s *ListUserBasicInfosResponse) SetStatusCode(v int32) *ListUserBasicInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserBasicInfosResponse) SetBody(v *ListUserBasicInfosResponseBody) *ListUserBasicInfosResponse {
	s.Body = v
	return s
}

func (s *ListUserBasicInfosResponse) Validate() error {
	return dara.Validate(s)
}
