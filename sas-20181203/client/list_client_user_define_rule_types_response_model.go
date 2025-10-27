// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientUserDefineRuleTypesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientUserDefineRuleTypesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientUserDefineRuleTypesResponse
	GetStatusCode() *int32
	SetBody(v *ListClientUserDefineRuleTypesResponseBody) *ListClientUserDefineRuleTypesResponse
	GetBody() *ListClientUserDefineRuleTypesResponseBody
}

type ListClientUserDefineRuleTypesResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientUserDefineRuleTypesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientUserDefineRuleTypesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientUserDefineRuleTypesResponse) GoString() string {
	return s.String()
}

func (s *ListClientUserDefineRuleTypesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientUserDefineRuleTypesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientUserDefineRuleTypesResponse) GetBody() *ListClientUserDefineRuleTypesResponseBody {
	return s.Body
}

func (s *ListClientUserDefineRuleTypesResponse) SetHeaders(v map[string]*string) *ListClientUserDefineRuleTypesResponse {
	s.Headers = v
	return s
}

func (s *ListClientUserDefineRuleTypesResponse) SetStatusCode(v int32) *ListClientUserDefineRuleTypesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientUserDefineRuleTypesResponse) SetBody(v *ListClientUserDefineRuleTypesResponseBody) *ListClientUserDefineRuleTypesResponse {
	s.Body = v
	return s
}

func (s *ListClientUserDefineRuleTypesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
