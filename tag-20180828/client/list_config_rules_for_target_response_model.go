// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConfigRulesForTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ListConfigRulesForTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ListConfigRulesForTargetResponse
	GetStatusCode() *int32
	SetBody(v *ListConfigRulesForTargetResponseBody) *ListConfigRulesForTargetResponse
	GetBody() *ListConfigRulesForTargetResponseBody
}

type ListConfigRulesForTargetResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListConfigRulesForTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListConfigRulesForTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s ListConfigRulesForTargetResponse) GoString() string {
	return s.String()
}

func (s *ListConfigRulesForTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ListConfigRulesForTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ListConfigRulesForTargetResponse) GetBody() *ListConfigRulesForTargetResponseBody {
	return s.Body
}

func (s *ListConfigRulesForTargetResponse) SetHeaders(v map[string]*string) *ListConfigRulesForTargetResponse {
	s.Headers = v
	return s
}

func (s *ListConfigRulesForTargetResponse) SetStatusCode(v int32) *ListConfigRulesForTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *ListConfigRulesForTargetResponse) SetBody(v *ListConfigRulesForTargetResponseBody) *ListConfigRulesForTargetResponse {
	s.Body = v
	return s
}

func (s *ListConfigRulesForTargetResponse) Validate() error {
	return dara.Validate(s)
}
