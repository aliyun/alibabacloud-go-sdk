// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListInstancesWithEcsInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListInstancesWithEcsInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListInstancesWithEcsInfoResponse
	GetStatusCode() *int32
	SetBody(v *ListInstancesWithEcsInfoResponseBody) *ListInstancesWithEcsInfoResponse
	GetBody() *ListInstancesWithEcsInfoResponseBody
}

type ListInstancesWithEcsInfoResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListInstancesWithEcsInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListInstancesWithEcsInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s ListInstancesWithEcsInfoResponse) GoString() string {
	return s.String()
}

func (s *ListInstancesWithEcsInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListInstancesWithEcsInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListInstancesWithEcsInfoResponse) GetBody() *ListInstancesWithEcsInfoResponseBody {
	return s.Body
}

func (s *ListInstancesWithEcsInfoResponse) SetHeaders(v map[string]*string) *ListInstancesWithEcsInfoResponse {
	s.Headers = v
	return s
}

func (s *ListInstancesWithEcsInfoResponse) SetStatusCode(v int32) *ListInstancesWithEcsInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *ListInstancesWithEcsInfoResponse) SetBody(v *ListInstancesWithEcsInfoResponseBody) *ListInstancesWithEcsInfoResponse {
	s.Body = v
	return s
}

func (s *ListInstancesWithEcsInfoResponse) Validate() error {
	return dara.Validate(s)
}
