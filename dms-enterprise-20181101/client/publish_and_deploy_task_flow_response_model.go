// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPublishAndDeployTaskFlowResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *PublishAndDeployTaskFlowResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *PublishAndDeployTaskFlowResponse
	GetStatusCode() *int32
	SetBody(v *PublishAndDeployTaskFlowResponseBody) *PublishAndDeployTaskFlowResponse
	GetBody() *PublishAndDeployTaskFlowResponseBody
}

type PublishAndDeployTaskFlowResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *PublishAndDeployTaskFlowResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s PublishAndDeployTaskFlowResponse) String() string {
	return dara.Prettify(s)
}

func (s PublishAndDeployTaskFlowResponse) GoString() string {
	return s.String()
}

func (s *PublishAndDeployTaskFlowResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *PublishAndDeployTaskFlowResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *PublishAndDeployTaskFlowResponse) GetBody() *PublishAndDeployTaskFlowResponseBody {
	return s.Body
}

func (s *PublishAndDeployTaskFlowResponse) SetHeaders(v map[string]*string) *PublishAndDeployTaskFlowResponse {
	s.Headers = v
	return s
}

func (s *PublishAndDeployTaskFlowResponse) SetStatusCode(v int32) *PublishAndDeployTaskFlowResponse {
	s.StatusCode = &v
	return s
}

func (s *PublishAndDeployTaskFlowResponse) SetBody(v *PublishAndDeployTaskFlowResponseBody) *PublishAndDeployTaskFlowResponse {
	s.Body = v
	return s
}

func (s *PublishAndDeployTaskFlowResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
