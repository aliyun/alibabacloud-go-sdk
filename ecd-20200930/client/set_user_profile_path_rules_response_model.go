// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetUserProfilePathRulesResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetUserProfilePathRulesResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetUserProfilePathRulesResponse
	GetStatusCode() *int32
	SetBody(v *SetUserProfilePathRulesResponseBody) *SetUserProfilePathRulesResponse
	GetBody() *SetUserProfilePathRulesResponseBody
}

type SetUserProfilePathRulesResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetUserProfilePathRulesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetUserProfilePathRulesResponse) String() string {
	return dara.Prettify(s)
}

func (s SetUserProfilePathRulesResponse) GoString() string {
	return s.String()
}

func (s *SetUserProfilePathRulesResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetUserProfilePathRulesResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetUserProfilePathRulesResponse) GetBody() *SetUserProfilePathRulesResponseBody {
	return s.Body
}

func (s *SetUserProfilePathRulesResponse) SetHeaders(v map[string]*string) *SetUserProfilePathRulesResponse {
	s.Headers = v
	return s
}

func (s *SetUserProfilePathRulesResponse) SetStatusCode(v int32) *SetUserProfilePathRulesResponse {
	s.StatusCode = &v
	return s
}

func (s *SetUserProfilePathRulesResponse) SetBody(v *SetUserProfilePathRulesResponseBody) *SetUserProfilePathRulesResponse {
	s.Body = v
	return s
}

func (s *SetUserProfilePathRulesResponse) Validate() error {
	return dara.Validate(s)
}
