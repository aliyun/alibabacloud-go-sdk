// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceInstanceBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListServiceInstanceBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListServiceInstanceBillResponse
	GetStatusCode() *int32
	SetBody(v *ListServiceInstanceBillResponseBody) *ListServiceInstanceBillResponse
	GetBody() *ListServiceInstanceBillResponseBody
}

type ListServiceInstanceBillResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListServiceInstanceBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListServiceInstanceBillResponse) String() string {
	return dara.Prettify(s)
}

func (s ListServiceInstanceBillResponse) GoString() string {
	return s.String()
}

func (s *ListServiceInstanceBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListServiceInstanceBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListServiceInstanceBillResponse) GetBody() *ListServiceInstanceBillResponseBody {
	return s.Body
}

func (s *ListServiceInstanceBillResponse) SetHeaders(v map[string]*string) *ListServiceInstanceBillResponse {
	s.Headers = v
	return s
}

func (s *ListServiceInstanceBillResponse) SetStatusCode(v int32) *ListServiceInstanceBillResponse {
	s.StatusCode = &v
	return s
}

func (s *ListServiceInstanceBillResponse) SetBody(v *ListServiceInstanceBillResponseBody) *ListServiceInstanceBillResponse {
	s.Body = v
	return s
}

func (s *ListServiceInstanceBillResponse) Validate() error {
	return dara.Validate(s)
}
