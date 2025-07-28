// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPageDemandPlanWithDemandInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PageDemandPlanWithDemandInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PageDemandPlanWithDemandInfoResponse
	GetStatusCode() *int32
	SetBody(v *PageDemandPlanWithDemandInfoResponseBody) *PageDemandPlanWithDemandInfoResponse
	GetBody() *PageDemandPlanWithDemandInfoResponseBody
}

type PageDemandPlanWithDemandInfoResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PageDemandPlanWithDemandInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PageDemandPlanWithDemandInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s PageDemandPlanWithDemandInfoResponse) GoString() string {
	return s.String()
}

func (s *PageDemandPlanWithDemandInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PageDemandPlanWithDemandInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PageDemandPlanWithDemandInfoResponse) GetBody() *PageDemandPlanWithDemandInfoResponseBody {
	return s.Body
}

func (s *PageDemandPlanWithDemandInfoResponse) SetHeaders(v map[string]*string) *PageDemandPlanWithDemandInfoResponse {
	s.Headers = v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponse) SetStatusCode(v int32) *PageDemandPlanWithDemandInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponse) SetBody(v *PageDemandPlanWithDemandInfoResponseBody) *PageDemandPlanWithDemandInfoResponse {
	s.Body = v
	return s
}

func (s *PageDemandPlanWithDemandInfoResponse) Validate() error {
	return dara.Validate(s)
}
