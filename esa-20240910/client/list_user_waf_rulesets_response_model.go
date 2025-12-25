// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListUserWafRulesetsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListUserWafRulesetsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListUserWafRulesetsResponse
	GetStatusCode() *int32
	SetBody(v *ListUserWafRulesetsResponseBody) *ListUserWafRulesetsResponse
	GetBody() *ListUserWafRulesetsResponseBody
}

type ListUserWafRulesetsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserWafRulesetsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserWafRulesetsResponse) String() string {
	return dara.Prettify(s)
}

func (s ListUserWafRulesetsResponse) GoString() string {
	return s.String()
}

func (s *ListUserWafRulesetsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListUserWafRulesetsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListUserWafRulesetsResponse) GetBody() *ListUserWafRulesetsResponseBody {
	return s.Body
}

func (s *ListUserWafRulesetsResponse) SetHeaders(v map[string]*string) *ListUserWafRulesetsResponse {
	s.Headers = v
	return s
}

func (s *ListUserWafRulesetsResponse) SetStatusCode(v int32) *ListUserWafRulesetsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserWafRulesetsResponse) SetBody(v *ListUserWafRulesetsResponseBody) *ListUserWafRulesetsResponse {
	s.Body = v
	return s
}

func (s *ListUserWafRulesetsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
