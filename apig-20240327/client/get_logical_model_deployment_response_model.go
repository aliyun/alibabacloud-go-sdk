// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLogicalModelDeploymentResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLogicalModelDeploymentResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLogicalModelDeploymentResponse
	GetStatusCode() *int32
	SetBody(v *GetLogicalModelDeploymentResponseBody) *GetLogicalModelDeploymentResponse
	GetBody() *GetLogicalModelDeploymentResponseBody
}

type GetLogicalModelDeploymentResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLogicalModelDeploymentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLogicalModelDeploymentResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLogicalModelDeploymentResponse) GoString() string {
	return s.String()
}

func (s *GetLogicalModelDeploymentResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLogicalModelDeploymentResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLogicalModelDeploymentResponse) GetBody() *GetLogicalModelDeploymentResponseBody {
	return s.Body
}

func (s *GetLogicalModelDeploymentResponse) SetHeaders(v map[string]*string) *GetLogicalModelDeploymentResponse {
	s.Headers = v
	return s
}

func (s *GetLogicalModelDeploymentResponse) SetStatusCode(v int32) *GetLogicalModelDeploymentResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLogicalModelDeploymentResponse) SetBody(v *GetLogicalModelDeploymentResponseBody) *GetLogicalModelDeploymentResponse {
	s.Body = v
	return s
}

func (s *GetLogicalModelDeploymentResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
