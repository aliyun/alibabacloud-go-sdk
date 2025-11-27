// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRCDeploymentSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateRCDeploymentSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateRCDeploymentSetResponse
	GetStatusCode() *int32
	SetBody(v *CreateRCDeploymentSetResponseBody) *CreateRCDeploymentSetResponse
	GetBody() *CreateRCDeploymentSetResponseBody
}

type CreateRCDeploymentSetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateRCDeploymentSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateRCDeploymentSetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateRCDeploymentSetResponse) GoString() string {
	return s.String()
}

func (s *CreateRCDeploymentSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateRCDeploymentSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateRCDeploymentSetResponse) GetBody() *CreateRCDeploymentSetResponseBody {
	return s.Body
}

func (s *CreateRCDeploymentSetResponse) SetHeaders(v map[string]*string) *CreateRCDeploymentSetResponse {
	s.Headers = v
	return s
}

func (s *CreateRCDeploymentSetResponse) SetStatusCode(v int32) *CreateRCDeploymentSetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateRCDeploymentSetResponse) SetBody(v *CreateRCDeploymentSetResponseBody) *CreateRCDeploymentSetResponse {
	s.Body = v
	return s
}

func (s *CreateRCDeploymentSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
