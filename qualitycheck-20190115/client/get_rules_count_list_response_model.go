// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRulesCountListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRulesCountListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRulesCountListResponse
	GetStatusCode() *int32
	SetBody(v *GetRulesCountListResponseBody) *GetRulesCountListResponse
	GetBody() *GetRulesCountListResponseBody
}

type GetRulesCountListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRulesCountListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRulesCountListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRulesCountListResponse) GoString() string {
	return s.String()
}

func (s *GetRulesCountListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRulesCountListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRulesCountListResponse) GetBody() *GetRulesCountListResponseBody {
	return s.Body
}

func (s *GetRulesCountListResponse) SetHeaders(v map[string]*string) *GetRulesCountListResponse {
	s.Headers = v
	return s
}

func (s *GetRulesCountListResponse) SetStatusCode(v int32) *GetRulesCountListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRulesCountListResponse) SetBody(v *GetRulesCountListResponseBody) *GetRulesCountListResponse {
	s.Body = v
	return s
}

func (s *GetRulesCountListResponse) Validate() error {
	return dara.Validate(s)
}
