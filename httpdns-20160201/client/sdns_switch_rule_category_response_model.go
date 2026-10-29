// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsSwitchRuleCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SdnsSwitchRuleCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SdnsSwitchRuleCategoryResponse
	GetStatusCode() *int32
	SetBody(v *SdnsSwitchRuleCategoryResponseBody) *SdnsSwitchRuleCategoryResponse
	GetBody() *SdnsSwitchRuleCategoryResponseBody
}

type SdnsSwitchRuleCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SdnsSwitchRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SdnsSwitchRuleCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s SdnsSwitchRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *SdnsSwitchRuleCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SdnsSwitchRuleCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SdnsSwitchRuleCategoryResponse) GetBody() *SdnsSwitchRuleCategoryResponseBody {
	return s.Body
}

func (s *SdnsSwitchRuleCategoryResponse) SetHeaders(v map[string]*string) *SdnsSwitchRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *SdnsSwitchRuleCategoryResponse) SetStatusCode(v int32) *SdnsSwitchRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SdnsSwitchRuleCategoryResponse) SetBody(v *SdnsSwitchRuleCategoryResponseBody) *SdnsSwitchRuleCategoryResponse {
	s.Body = v
	return s
}

func (s *SdnsSwitchRuleCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
