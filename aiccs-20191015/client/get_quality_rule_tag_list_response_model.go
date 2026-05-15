// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetQualityRuleTagListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetQualityRuleTagListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetQualityRuleTagListResponse
	GetStatusCode() *int32
	SetBody(v *GetQualityRuleTagListResponseBody) *GetQualityRuleTagListResponse
	GetBody() *GetQualityRuleTagListResponseBody
}

type GetQualityRuleTagListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetQualityRuleTagListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetQualityRuleTagListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetQualityRuleTagListResponse) GoString() string {
	return s.String()
}

func (s *GetQualityRuleTagListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetQualityRuleTagListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetQualityRuleTagListResponse) GetBody() *GetQualityRuleTagListResponseBody {
	return s.Body
}

func (s *GetQualityRuleTagListResponse) SetHeaders(v map[string]*string) *GetQualityRuleTagListResponse {
	s.Headers = v
	return s
}

func (s *GetQualityRuleTagListResponse) SetStatusCode(v int32) *GetQualityRuleTagListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetQualityRuleTagListResponse) SetBody(v *GetQualityRuleTagListResponseBody) *GetQualityRuleTagListResponse {
	s.Body = v
	return s
}

func (s *GetQualityRuleTagListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
