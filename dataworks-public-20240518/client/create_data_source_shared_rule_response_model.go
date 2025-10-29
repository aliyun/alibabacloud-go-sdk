// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataSourceSharedRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDataSourceSharedRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDataSourceSharedRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateDataSourceSharedRuleResponseBody) *CreateDataSourceSharedRuleResponse
	GetBody() *CreateDataSourceSharedRuleResponseBody
}

type CreateDataSourceSharedRuleResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDataSourceSharedRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDataSourceSharedRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateDataSourceSharedRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateDataSourceSharedRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDataSourceSharedRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDataSourceSharedRuleResponse) GetBody() *CreateDataSourceSharedRuleResponseBody {
	return s.Body
}

func (s *CreateDataSourceSharedRuleResponse) SetHeaders(v map[string]*string) *CreateDataSourceSharedRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateDataSourceSharedRuleResponse) SetStatusCode(v int32) *CreateDataSourceSharedRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDataSourceSharedRuleResponse) SetBody(v *CreateDataSourceSharedRuleResponseBody) *CreateDataSourceSharedRuleResponse {
	s.Body = v
	return s
}

func (s *CreateDataSourceSharedRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
