// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSavingPlanDeductableCommodityResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSavingPlanDeductableCommodityResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSavingPlanDeductableCommodityResponse
	GetStatusCode() *int32
	SetBody(v *GetSavingPlanDeductableCommodityResponseBody) *GetSavingPlanDeductableCommodityResponse
	GetBody() *GetSavingPlanDeductableCommodityResponseBody
}

type GetSavingPlanDeductableCommodityResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSavingPlanDeductableCommodityResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSavingPlanDeductableCommodityResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSavingPlanDeductableCommodityResponse) GoString() string {
	return s.String()
}

func (s *GetSavingPlanDeductableCommodityResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSavingPlanDeductableCommodityResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSavingPlanDeductableCommodityResponse) GetBody() *GetSavingPlanDeductableCommodityResponseBody {
	return s.Body
}

func (s *GetSavingPlanDeductableCommodityResponse) SetHeaders(v map[string]*string) *GetSavingPlanDeductableCommodityResponse {
	s.Headers = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponse) SetStatusCode(v int32) *GetSavingPlanDeductableCommodityResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponse) SetBody(v *GetSavingPlanDeductableCommodityResponseBody) *GetSavingPlanDeductableCommodityResponse {
	s.Body = v
	return s
}

func (s *GetSavingPlanDeductableCommodityResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
