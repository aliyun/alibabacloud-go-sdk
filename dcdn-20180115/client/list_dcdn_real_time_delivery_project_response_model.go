// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDcdnRealTimeDeliveryProjectResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListDcdnRealTimeDeliveryProjectResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListDcdnRealTimeDeliveryProjectResponse
	GetStatusCode() *int32
	SetBody(v *ListDcdnRealTimeDeliveryProjectResponseBody) *ListDcdnRealTimeDeliveryProjectResponse
	GetBody() *ListDcdnRealTimeDeliveryProjectResponseBody
}

type ListDcdnRealTimeDeliveryProjectResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDcdnRealTimeDeliveryProjectResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDcdnRealTimeDeliveryProjectResponse) String() string {
	return dara.Prettify(s)
}

func (s ListDcdnRealTimeDeliveryProjectResponse) GoString() string {
	return s.String()
}

func (s *ListDcdnRealTimeDeliveryProjectResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListDcdnRealTimeDeliveryProjectResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListDcdnRealTimeDeliveryProjectResponse) GetBody() *ListDcdnRealTimeDeliveryProjectResponseBody {
	return s.Body
}

func (s *ListDcdnRealTimeDeliveryProjectResponse) SetHeaders(v map[string]*string) *ListDcdnRealTimeDeliveryProjectResponse {
	s.Headers = v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponse) SetStatusCode(v int32) *ListDcdnRealTimeDeliveryProjectResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponse) SetBody(v *ListDcdnRealTimeDeliveryProjectResponseBody) *ListDcdnRealTimeDeliveryProjectResponse {
	s.Body = v
	return s
}

func (s *ListDcdnRealTimeDeliveryProjectResponse) Validate() error {
	return dara.Validate(s)
}
