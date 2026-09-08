// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataMaskingRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataMaskingRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataMaskingRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataMaskingRuleResponseBody) *CreateDataMaskingRuleResponse
	GetBody() *CreateDataMaskingRuleResponseBody
}

type CreateDataMaskingRuleResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataMaskingRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataMaskingRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataMaskingRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDataMaskingRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataMaskingRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataMaskingRuleResponse) GetBody() *CreateDataMaskingRuleResponseBody {
	return s.Body
}

func (s *CreateDataMaskingRuleResponse) SetHeaders(v map[string]*string) *CreateDataMaskingRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDataMaskingRuleResponse) SetStatusCode(v int32) *CreateDataMaskingRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataMaskingRuleResponse) SetBody(v *CreateDataMaskingRuleResponseBody) *CreateDataMaskingRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDataMaskingRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
