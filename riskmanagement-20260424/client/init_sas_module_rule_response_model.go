// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iInitSasModuleRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *InitSasModuleRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *InitSasModuleRuleResponse
	GetStatusCode() *int32
	SetBody(v *InitSasModuleRuleResponseBody) *InitSasModuleRuleResponse
	GetBody() *InitSasModuleRuleResponseBody
}

type InitSasModuleRuleResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *InitSasModuleRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s InitSasModuleRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s InitSasModuleRuleResponse) GoString() string {
	return s.String()
}

func (s *InitSasModuleRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *InitSasModuleRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *InitSasModuleRuleResponse) GetBody() *InitSasModuleRuleResponseBody {
	return s.Body
}

func (s *InitSasModuleRuleResponse) SetHeaders(v map[string]*string) *InitSasModuleRuleResponse {
	s.Headers = v
	return s
}

func (s *InitSasModuleRuleResponse) SetStatusCode(v int32) *InitSasModuleRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *InitSasModuleRuleResponse) SetBody(v *InitSasModuleRuleResponseBody) *InitSasModuleRuleResponse {
	s.Body = v
	return s
}

func (s *InitSasModuleRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
