// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSdnsListRuleCategoriesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SdnsListRuleCategoriesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SdnsListRuleCategoriesResponse
	GetStatusCode() *int32
	SetBody(v *SdnsListRuleCategoriesResponseBody) *SdnsListRuleCategoriesResponse
	GetBody() *SdnsListRuleCategoriesResponseBody
}

type SdnsListRuleCategoriesResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SdnsListRuleCategoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SdnsListRuleCategoriesResponse) String() string {
	return dara.Prettify(s)
}

func (s SdnsListRuleCategoriesResponse) GoString() string {
	return s.String()
}

func (s *SdnsListRuleCategoriesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SdnsListRuleCategoriesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SdnsListRuleCategoriesResponse) GetBody() *SdnsListRuleCategoriesResponseBody {
	return s.Body
}

func (s *SdnsListRuleCategoriesResponse) SetHeaders(v map[string]*string) *SdnsListRuleCategoriesResponse {
	s.Headers = v
	return s
}

func (s *SdnsListRuleCategoriesResponse) SetStatusCode(v int32) *SdnsListRuleCategoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SdnsListRuleCategoriesResponse) SetBody(v *SdnsListRuleCategoriesResponseBody) *SdnsListRuleCategoriesResponse {
	s.Body = v
	return s
}

func (s *SdnsListRuleCategoriesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
