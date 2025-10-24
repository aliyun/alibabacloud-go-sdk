// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHybridCloudClusterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyHybridCloudClusterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyHybridCloudClusterRuleResponse
	GetStatusCode() *int32
	SetBody(v *ModifyHybridCloudClusterRuleResponseBody) *ModifyHybridCloudClusterRuleResponse
	GetBody() *ModifyHybridCloudClusterRuleResponseBody
}

type ModifyHybridCloudClusterRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyHybridCloudClusterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyHybridCloudClusterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyHybridCloudClusterRuleResponse) GoString() string {
	return s.String()
}

func (s *ModifyHybridCloudClusterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyHybridCloudClusterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyHybridCloudClusterRuleResponse) GetBody() *ModifyHybridCloudClusterRuleResponseBody {
	return s.Body
}

func (s *ModifyHybridCloudClusterRuleResponse) SetHeaders(v map[string]*string) *ModifyHybridCloudClusterRuleResponse {
	s.Headers = v
	return s
}

func (s *ModifyHybridCloudClusterRuleResponse) SetStatusCode(v int32) *ModifyHybridCloudClusterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyHybridCloudClusterRuleResponse) SetBody(v *ModifyHybridCloudClusterRuleResponseBody) *ModifyHybridCloudClusterRuleResponse {
	s.Body = v
	return s
}

func (s *ModifyHybridCloudClusterRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
