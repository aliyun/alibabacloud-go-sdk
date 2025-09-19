// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaStrategyTemplateSummaryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpaStrategyTemplateSummaryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpaStrategyTemplateSummaryResponse
	GetStatusCode() *int32
	SetBody(v *GetOpaStrategyTemplateSummaryResponseBody) *GetOpaStrategyTemplateSummaryResponse
	GetBody() *GetOpaStrategyTemplateSummaryResponseBody
}

type GetOpaStrategyTemplateSummaryResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpaStrategyTemplateSummaryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpaStrategyTemplateSummaryResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpaStrategyTemplateSummaryResponse) GoString() string {
	return s.String()
}

func (s *GetOpaStrategyTemplateSummaryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpaStrategyTemplateSummaryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpaStrategyTemplateSummaryResponse) GetBody() *GetOpaStrategyTemplateSummaryResponseBody {
	return s.Body
}

func (s *GetOpaStrategyTemplateSummaryResponse) SetHeaders(v map[string]*string) *GetOpaStrategyTemplateSummaryResponse {
	s.Headers = v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponse) SetStatusCode(v int32) *GetOpaStrategyTemplateSummaryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponse) SetBody(v *GetOpaStrategyTemplateSummaryResponseBody) *GetOpaStrategyTemplateSummaryResponse {
	s.Body = v
	return s
}

func (s *GetOpaStrategyTemplateSummaryResponse) Validate() error {
	return dara.Validate(s)
}
