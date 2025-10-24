// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateHybridCloudClusterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateHybridCloudClusterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateHybridCloudClusterRuleResponse
	GetStatusCode() *int32
	SetBody(v *CreateHybridCloudClusterRuleResponseBody) *CreateHybridCloudClusterRuleResponse
	GetBody() *CreateHybridCloudClusterRuleResponseBody
}

type CreateHybridCloudClusterRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateHybridCloudClusterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateHybridCloudClusterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateHybridCloudClusterRuleResponse) GoString() string {
	return s.String()
}

func (s *CreateHybridCloudClusterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateHybridCloudClusterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateHybridCloudClusterRuleResponse) GetBody() *CreateHybridCloudClusterRuleResponseBody {
	return s.Body
}

func (s *CreateHybridCloudClusterRuleResponse) SetHeaders(v map[string]*string) *CreateHybridCloudClusterRuleResponse {
	s.Headers = v
	return s
}

func (s *CreateHybridCloudClusterRuleResponse) SetStatusCode(v int32) *CreateHybridCloudClusterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateHybridCloudClusterRuleResponse) SetBody(v *CreateHybridCloudClusterRuleResponseBody) *CreateHybridCloudClusterRuleResponse {
	s.Body = v
	return s
}

func (s *CreateHybridCloudClusterRuleResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
