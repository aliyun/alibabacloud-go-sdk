// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadRuleResponse
	GetStatusCode() *int32
	SetBody(v *UploadRuleResponseBody) *UploadRuleResponse
	GetBody() *UploadRuleResponseBody
}

type UploadRuleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadRuleResponse) GoString() string {
	return s.String()
}

func (s *UploadRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadRuleResponse) GetBody() *UploadRuleResponseBody {
	return s.Body
}

func (s *UploadRuleResponse) SetHeaders(v map[string]*string) *UploadRuleResponse {
	s.Headers = v
	return s
}

func (s *UploadRuleResponse) SetStatusCode(v int32) *UploadRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadRuleResponse) SetBody(v *UploadRuleResponseBody) *UploadRuleResponse {
	s.Body = v
	return s
}

func (s *UploadRuleResponse) Validate() error {
	return dara.Validate(s)
}
