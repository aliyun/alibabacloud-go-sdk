// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPurchaseRatePlanResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PurchaseRatePlanResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PurchaseRatePlanResponse
	GetStatusCode() *int32
	SetBody(v *PurchaseRatePlanResponseBody) *PurchaseRatePlanResponse
	GetBody() *PurchaseRatePlanResponseBody
}

type PurchaseRatePlanResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PurchaseRatePlanResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PurchaseRatePlanResponse) String() string {
	return dara.Prettify(s)
}

func (s PurchaseRatePlanResponse) GoString() string {
	return s.String()
}

func (s *PurchaseRatePlanResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PurchaseRatePlanResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PurchaseRatePlanResponse) GetBody() *PurchaseRatePlanResponseBody {
	return s.Body
}

func (s *PurchaseRatePlanResponse) SetHeaders(v map[string]*string) *PurchaseRatePlanResponse {
	s.Headers = v
	return s
}

func (s *PurchaseRatePlanResponse) SetStatusCode(v int32) *PurchaseRatePlanResponse {
	s.StatusCode = &v
	return s
}

func (s *PurchaseRatePlanResponse) SetBody(v *PurchaseRatePlanResponseBody) *PurchaseRatePlanResponse {
	s.Body = v
	return s
}

func (s *PurchaseRatePlanResponse) Validate() error {
	return dara.Validate(s)
}
