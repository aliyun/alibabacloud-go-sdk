// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenPlanAccountDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTokenPlanAccountDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTokenPlanAccountDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetTokenPlanAccountDetailResponseBody) *GetTokenPlanAccountDetailResponse
	GetBody() *GetTokenPlanAccountDetailResponseBody
}

type GetTokenPlanAccountDetailResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTokenPlanAccountDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTokenPlanAccountDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTokenPlanAccountDetailResponse) GoString() string {
	return s.String()
}

func (s *GetTokenPlanAccountDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTokenPlanAccountDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTokenPlanAccountDetailResponse) GetBody() *GetTokenPlanAccountDetailResponseBody {
	return s.Body
}

func (s *GetTokenPlanAccountDetailResponse) SetHeaders(v map[string]*string) *GetTokenPlanAccountDetailResponse {
	s.Headers = v
	return s
}

func (s *GetTokenPlanAccountDetailResponse) SetStatusCode(v int32) *GetTokenPlanAccountDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTokenPlanAccountDetailResponse) SetBody(v *GetTokenPlanAccountDetailResponseBody) *GetTokenPlanAccountDetailResponse {
	s.Body = v
	return s
}

func (s *GetTokenPlanAccountDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
