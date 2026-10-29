// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsVerifyRuleCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SdnsVerifyRuleCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SdnsVerifyRuleCategoryResponse
	GetStatusCode() *int32
	SetBody(v *SdnsVerifyRuleCategoryResponseBody) *SdnsVerifyRuleCategoryResponse
	GetBody() *SdnsVerifyRuleCategoryResponseBody
}

type SdnsVerifyRuleCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SdnsVerifyRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SdnsVerifyRuleCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s SdnsVerifyRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *SdnsVerifyRuleCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SdnsVerifyRuleCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SdnsVerifyRuleCategoryResponse) GetBody() *SdnsVerifyRuleCategoryResponseBody {
	return s.Body
}

func (s *SdnsVerifyRuleCategoryResponse) SetHeaders(v map[string]*string) *SdnsVerifyRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *SdnsVerifyRuleCategoryResponse) SetStatusCode(v int32) *SdnsVerifyRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SdnsVerifyRuleCategoryResponse) SetBody(v *SdnsVerifyRuleCategoryResponseBody) *SdnsVerifyRuleCategoryResponse {
	s.Body = v
	return s
}

func (s *SdnsVerifyRuleCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
