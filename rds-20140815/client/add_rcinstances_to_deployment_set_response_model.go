// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRCInstancesToDeploymentSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRCInstancesToDeploymentSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRCInstancesToDeploymentSetResponse
	GetStatusCode() *int32
	SetBody(v *AddRCInstancesToDeploymentSetResponseBody) *AddRCInstancesToDeploymentSetResponse
	GetBody() *AddRCInstancesToDeploymentSetResponseBody
}

type AddRCInstancesToDeploymentSetResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRCInstancesToDeploymentSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRCInstancesToDeploymentSetResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRCInstancesToDeploymentSetResponse) GoString() string {
	return s.String()
}

func (s *AddRCInstancesToDeploymentSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRCInstancesToDeploymentSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRCInstancesToDeploymentSetResponse) GetBody() *AddRCInstancesToDeploymentSetResponseBody {
	return s.Body
}

func (s *AddRCInstancesToDeploymentSetResponse) SetHeaders(v map[string]*string) *AddRCInstancesToDeploymentSetResponse {
	s.Headers = v
	return s
}

func (s *AddRCInstancesToDeploymentSetResponse) SetStatusCode(v int32) *AddRCInstancesToDeploymentSetResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRCInstancesToDeploymentSetResponse) SetBody(v *AddRCInstancesToDeploymentSetResponseBody) *AddRCInstancesToDeploymentSetResponse {
	s.Body = v
	return s
}

func (s *AddRCInstancesToDeploymentSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
