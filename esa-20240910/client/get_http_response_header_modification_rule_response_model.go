// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHttpResponseHeaderModificationRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetHttpResponseHeaderModificationRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetHttpResponseHeaderModificationRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetHttpResponseHeaderModificationRuleResponseBody) *GetHttpResponseHeaderModificationRuleResponse
	GetBody() *GetHttpResponseHeaderModificationRuleResponseBody
}

type GetHttpResponseHeaderModificationRuleResponse struct {
	Headers    map[string]*string                                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetHttpResponseHeaderModificationRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetHttpResponseHeaderModificationRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetHttpResponseHeaderModificationRuleResponse) GoString() string {
	return s.String()
}

func (s *GetHttpResponseHeaderModificationRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetHttpResponseHeaderModificationRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetHttpResponseHeaderModificationRuleResponse) GetBody() *GetHttpResponseHeaderModificationRuleResponseBody {
	return s.Body
}

func (s *GetHttpResponseHeaderModificationRuleResponse) SetHeaders(v map[string]*string) *GetHttpResponseHeaderModificationRuleResponse {
	s.Headers = v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponse) SetStatusCode(v int32) *GetHttpResponseHeaderModificationRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponse) SetBody(v *GetHttpResponseHeaderModificationRuleResponseBody) *GetHttpResponseHeaderModificationRuleResponse {
	s.Body = v
	return s
}

func (s *GetHttpResponseHeaderModificationRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
