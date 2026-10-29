// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsDeleteRuleCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SdnsDeleteRuleCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SdnsDeleteRuleCategoryResponse
	GetStatusCode() *int32
	SetBody(v *SdnsDeleteRuleCategoryResponseBody) *SdnsDeleteRuleCategoryResponse
	GetBody() *SdnsDeleteRuleCategoryResponseBody
}

type SdnsDeleteRuleCategoryResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SdnsDeleteRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SdnsDeleteRuleCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s SdnsDeleteRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *SdnsDeleteRuleCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SdnsDeleteRuleCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SdnsDeleteRuleCategoryResponse) GetBody() *SdnsDeleteRuleCategoryResponseBody {
	return s.Body
}

func (s *SdnsDeleteRuleCategoryResponse) SetHeaders(v map[string]*string) *SdnsDeleteRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *SdnsDeleteRuleCategoryResponse) SetStatusCode(v int32) *SdnsDeleteRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SdnsDeleteRuleCategoryResponse) SetBody(v *SdnsDeleteRuleCategoryResponseBody) *SdnsDeleteRuleCategoryResponse {
	s.Body = v
	return s
}

func (s *SdnsDeleteRuleCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
