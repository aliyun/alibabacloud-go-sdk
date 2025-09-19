// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRulesCountResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRulesCountResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRulesCountResponse
	GetStatusCode() *int32
	SetBody(v *GetRulesCountResponseBody) *GetRulesCountResponse
	GetBody() *GetRulesCountResponseBody
}

type GetRulesCountResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRulesCountResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRulesCountResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRulesCountResponse) GoString() string {
	return s.String()
}

func (s *GetRulesCountResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRulesCountResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRulesCountResponse) GetBody() *GetRulesCountResponseBody {
	return s.Body
}

func (s *GetRulesCountResponse) SetHeaders(v map[string]*string) *GetRulesCountResponse {
	s.Headers = v
	return s
}

func (s *GetRulesCountResponse) SetStatusCode(v int32) *GetRulesCountResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRulesCountResponse) SetBody(v *GetRulesCountResponseBody) *GetRulesCountResponse {
	s.Body = v
	return s
}

func (s *GetRulesCountResponse) Validate() error {
	return dara.Validate(s)
}
