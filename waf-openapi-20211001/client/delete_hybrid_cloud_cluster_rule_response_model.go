// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteHybridCloudClusterRuleResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteHybridCloudClusterRuleResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteHybridCloudClusterRuleResponse
	GetStatusCode() *int32
	SetBody(v *DeleteHybridCloudClusterRuleResponseBody) *DeleteHybridCloudClusterRuleResponse
	GetBody() *DeleteHybridCloudClusterRuleResponseBody
}

type DeleteHybridCloudClusterRuleResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteHybridCloudClusterRuleResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteHybridCloudClusterRuleResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteHybridCloudClusterRuleResponse) GoString() string {
	return s.String()
}

func (s *DeleteHybridCloudClusterRuleResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteHybridCloudClusterRuleResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteHybridCloudClusterRuleResponse) GetBody() *DeleteHybridCloudClusterRuleResponseBody {
	return s.Body
}

func (s *DeleteHybridCloudClusterRuleResponse) SetHeaders(v map[string]*string) *DeleteHybridCloudClusterRuleResponse {
	s.Headers = v
	return s
}

func (s *DeleteHybridCloudClusterRuleResponse) SetStatusCode(v int32) *DeleteHybridCloudClusterRuleResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteHybridCloudClusterRuleResponse) SetBody(v *DeleteHybridCloudClusterRuleResponseBody) *DeleteHybridCloudClusterRuleResponse {
	s.Body = v
	return s
}

func (s *DeleteHybridCloudClusterRuleResponse) Validate() error {
	return dara.Validate(s)
}
