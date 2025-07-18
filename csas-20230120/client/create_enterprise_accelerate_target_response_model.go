// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateEnterpriseAccelerateTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateEnterpriseAccelerateTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateEnterpriseAccelerateTargetResponse
	GetStatusCode() *int32
	SetBody(v *CreateEnterpriseAccelerateTargetResponseBody) *CreateEnterpriseAccelerateTargetResponse
	GetBody() *CreateEnterpriseAccelerateTargetResponseBody
}

type CreateEnterpriseAccelerateTargetResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateEnterpriseAccelerateTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateEnterpriseAccelerateTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateEnterpriseAccelerateTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateEnterpriseAccelerateTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateEnterpriseAccelerateTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateEnterpriseAccelerateTargetResponse) GetBody() *CreateEnterpriseAccelerateTargetResponseBody {
	return s.Body
}

func (s *CreateEnterpriseAccelerateTargetResponse) SetHeaders(v map[string]*string) *CreateEnterpriseAccelerateTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateEnterpriseAccelerateTargetResponse) SetStatusCode(v int32) *CreateEnterpriseAccelerateTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateEnterpriseAccelerateTargetResponse) SetBody(v *CreateEnterpriseAccelerateTargetResponseBody) *CreateEnterpriseAccelerateTargetResponse {
	s.Body = v
	return s
}

func (s *CreateEnterpriseAccelerateTargetResponse) Validate() error {
	return dara.Validate(s)
}
