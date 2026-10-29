// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsGetRuleCategoryResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SdnsGetRuleCategoryResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SdnsGetRuleCategoryResponse
	GetStatusCode() *int32
	SetBody(v *SdnsGetRuleCategoryResponseBody) *SdnsGetRuleCategoryResponse
	GetBody() *SdnsGetRuleCategoryResponseBody
}

type SdnsGetRuleCategoryResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SdnsGetRuleCategoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SdnsGetRuleCategoryResponse) String() string {
	return dara.Prettify(s)
}

func (s SdnsGetRuleCategoryResponse) GoString() string {
	return s.String()
}

func (s *SdnsGetRuleCategoryResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SdnsGetRuleCategoryResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SdnsGetRuleCategoryResponse) GetBody() *SdnsGetRuleCategoryResponseBody {
	return s.Body
}

func (s *SdnsGetRuleCategoryResponse) SetHeaders(v map[string]*string) *SdnsGetRuleCategoryResponse {
	s.Headers = v
	return s
}

func (s *SdnsGetRuleCategoryResponse) SetStatusCode(v int32) *SdnsGetRuleCategoryResponse {
	s.StatusCode = &v
	return s
}

func (s *SdnsGetRuleCategoryResponse) SetBody(v *SdnsGetRuleCategoryResponseBody) *SdnsGetRuleCategoryResponse {
	s.Body = v
	return s
}

func (s *SdnsGetRuleCategoryResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
