// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesEcsInfoListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancesEcsInfoListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancesEcsInfoListResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancesEcsInfoListResponseBody) *ListInstancesEcsInfoListResponse
	GetBody() *ListInstancesEcsInfoListResponseBody
}

type ListInstancesEcsInfoListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesEcsInfoListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesEcsInfoListResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesEcsInfoListResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesEcsInfoListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancesEcsInfoListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancesEcsInfoListResponse) GetBody() *ListInstancesEcsInfoListResponseBody {
	return s.Body
}

func (s *ListInstancesEcsInfoListResponse) SetHeaders(v map[string]*string) *ListInstancesEcsInfoListResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesEcsInfoListResponse) SetStatusCode(v int32) *ListInstancesEcsInfoListResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesEcsInfoListResponse) SetBody(v *ListInstancesEcsInfoListResponseBody) *ListInstancesEcsInfoListResponse {
	s.Body = v
	return s
}

func (s *ListInstancesEcsInfoListResponse) Validate() error {
	return dara.Validate(s)
}
