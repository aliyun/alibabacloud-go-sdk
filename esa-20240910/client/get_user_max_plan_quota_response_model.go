// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMaxPlanQuotaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserMaxPlanQuotaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserMaxPlanQuotaResponse
	GetStatusCode() *int32
	SetBody(v *GetUserMaxPlanQuotaResponseBody) *GetUserMaxPlanQuotaResponse
	GetBody() *GetUserMaxPlanQuotaResponseBody
}

type GetUserMaxPlanQuotaResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserMaxPlanQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserMaxPlanQuotaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserMaxPlanQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetUserMaxPlanQuotaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserMaxPlanQuotaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserMaxPlanQuotaResponse) GetBody() *GetUserMaxPlanQuotaResponseBody {
	return s.Body
}

func (s *GetUserMaxPlanQuotaResponse) SetHeaders(v map[string]*string) *GetUserMaxPlanQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetUserMaxPlanQuotaResponse) SetStatusCode(v int32) *GetUserMaxPlanQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserMaxPlanQuotaResponse) SetBody(v *GetUserMaxPlanQuotaResponseBody) *GetUserMaxPlanQuotaResponse {
	s.Body = v
	return s
}

func (s *GetUserMaxPlanQuotaResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
