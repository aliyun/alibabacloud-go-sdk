// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDeploymentTargetV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDeploymentTargetV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDeploymentTargetV2Response
	GetStatusCode() *int32
	SetBody(v *UpdateDeploymentTargetV2ResponseBody) *UpdateDeploymentTargetV2Response
	GetBody() *UpdateDeploymentTargetV2ResponseBody
}

type UpdateDeploymentTargetV2Response struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDeploymentTargetV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDeploymentTargetV2Response) String() string {
	return dara.Prettify(s)
}

func (s UpdateDeploymentTargetV2Response) GoString() string {
	return s.String()
}

func (s *UpdateDeploymentTargetV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDeploymentTargetV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDeploymentTargetV2Response) GetBody() *UpdateDeploymentTargetV2ResponseBody {
	return s.Body
}

func (s *UpdateDeploymentTargetV2Response) SetHeaders(v map[string]*string) *UpdateDeploymentTargetV2Response {
	s.Headers = v
	return s
}

func (s *UpdateDeploymentTargetV2Response) SetStatusCode(v int32) *UpdateDeploymentTargetV2Response {
	s.StatusCode = &v
	return s
}

func (s *UpdateDeploymentTargetV2Response) SetBody(v *UpdateDeploymentTargetV2ResponseBody) *UpdateDeploymentTargetV2Response {
	s.Body = v
	return s
}

func (s *UpdateDeploymentTargetV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
