// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeploymentTargetV2Response interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateDeploymentTargetV2Response
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateDeploymentTargetV2Response
	GetStatusCode() *int32
	SetBody(v *CreateDeploymentTargetV2ResponseBody) *CreateDeploymentTargetV2Response
	GetBody() *CreateDeploymentTargetV2ResponseBody
}

type CreateDeploymentTargetV2Response struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDeploymentTargetV2ResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDeploymentTargetV2Response) String() string {
	return dara.Prettify(s)
}

func (s CreateDeploymentTargetV2Response) GoString() string {
	return s.String()
}

func (s *CreateDeploymentTargetV2Response) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateDeploymentTargetV2Response) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateDeploymentTargetV2Response) GetBody() *CreateDeploymentTargetV2ResponseBody {
	return s.Body
}

func (s *CreateDeploymentTargetV2Response) SetHeaders(v map[string]*string) *CreateDeploymentTargetV2Response {
	s.Headers = v
	return s
}

func (s *CreateDeploymentTargetV2Response) SetStatusCode(v int32) *CreateDeploymentTargetV2Response {
	s.StatusCode = &v
	return s
}

func (s *CreateDeploymentTargetV2Response) SetBody(v *CreateDeploymentTargetV2ResponseBody) *CreateDeploymentTargetV2Response {
	s.Body = v
	return s
}

func (s *CreateDeploymentTargetV2Response) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
