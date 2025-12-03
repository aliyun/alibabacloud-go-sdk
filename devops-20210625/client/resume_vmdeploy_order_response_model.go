// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iResumeVMDeployOrderResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ResumeVMDeployOrderResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ResumeVMDeployOrderResponse
	GetStatusCode() *int32
	SetBody(v *ResumeVMDeployOrderResponseBody) *ResumeVMDeployOrderResponse
	GetBody() *ResumeVMDeployOrderResponseBody
}

type ResumeVMDeployOrderResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ResumeVMDeployOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ResumeVMDeployOrderResponse) String() string {
	return dara.Prettify(s)
}

func (s ResumeVMDeployOrderResponse) GoString() string {
	return s.String()
}

func (s *ResumeVMDeployOrderResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ResumeVMDeployOrderResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ResumeVMDeployOrderResponse) GetBody() *ResumeVMDeployOrderResponseBody {
	return s.Body
}

func (s *ResumeVMDeployOrderResponse) SetHeaders(v map[string]*string) *ResumeVMDeployOrderResponse {
	s.Headers = v
	return s
}

func (s *ResumeVMDeployOrderResponse) SetStatusCode(v int32) *ResumeVMDeployOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *ResumeVMDeployOrderResponse) SetBody(v *ResumeVMDeployOrderResponseBody) *ResumeVMDeployOrderResponse {
	s.Body = v
	return s
}

func (s *ResumeVMDeployOrderResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
