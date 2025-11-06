// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLocalityRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLocalityRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLocalityRuleResponse
	GetStatusCode() *int32
	SetBody(v *GetLocalityRuleResponseBody) *GetLocalityRuleResponse
	GetBody() *GetLocalityRuleResponseBody
}

type GetLocalityRuleResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLocalityRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLocalityRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLocalityRuleResponse) GoString() string {
	return s.String()
}

func (s *GetLocalityRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLocalityRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLocalityRuleResponse) GetBody() *GetLocalityRuleResponseBody {
	return s.Body
}

func (s *GetLocalityRuleResponse) SetHeaders(v map[string]*string) *GetLocalityRuleResponse {
	s.Headers = v
	return s
}

func (s *GetLocalityRuleResponse) SetStatusCode(v int32) *GetLocalityRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLocalityRuleResponse) SetBody(v *GetLocalityRuleResponseBody) *GetLocalityRuleResponse {
	s.Body = v
	return s
}

func (s *GetLocalityRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
