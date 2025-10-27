// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListClientUserDefineRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListClientUserDefineRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListClientUserDefineRulesResponse
	GetStatusCode() *int32
	SetBody(v *ListClientUserDefineRulesResponseBody) *ListClientUserDefineRulesResponse
	GetBody() *ListClientUserDefineRulesResponseBody
}

type ListClientUserDefineRulesResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListClientUserDefineRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListClientUserDefineRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s ListClientUserDefineRulesResponse) GoString() string {
	return s.String()
}

func (s *ListClientUserDefineRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListClientUserDefineRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListClientUserDefineRulesResponse) GetBody() *ListClientUserDefineRulesResponseBody {
	return s.Body
}

func (s *ListClientUserDefineRulesResponse) SetHeaders(v map[string]*string) *ListClientUserDefineRulesResponse {
	s.Headers = v
	return s
}

func (s *ListClientUserDefineRulesResponse) SetStatusCode(v int32) *ListClientUserDefineRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListClientUserDefineRulesResponse) SetBody(v *ListClientUserDefineRulesResponseBody) *ListClientUserDefineRulesResponse {
	s.Body = v
	return s
}

func (s *ListClientUserDefineRulesResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
