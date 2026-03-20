// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRCDeploymentSetAttributeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyRCDeploymentSetAttributeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyRCDeploymentSetAttributeResponse
	GetStatusCode() *int32
	SetBody(v *ModifyRCDeploymentSetAttributeResponseBody) *ModifyRCDeploymentSetAttributeResponse
	GetBody() *ModifyRCDeploymentSetAttributeResponseBody
}

type ModifyRCDeploymentSetAttributeResponse struct {
	Headers    map[string]*string                          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyRCDeploymentSetAttributeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyRCDeploymentSetAttributeResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyRCDeploymentSetAttributeResponse) GoString() string {
	return s.String()
}

func (s *ModifyRCDeploymentSetAttributeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyRCDeploymentSetAttributeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyRCDeploymentSetAttributeResponse) GetBody() *ModifyRCDeploymentSetAttributeResponseBody {
	return s.Body
}

func (s *ModifyRCDeploymentSetAttributeResponse) SetHeaders(v map[string]*string) *ModifyRCDeploymentSetAttributeResponse {
	s.Headers = v
	return s
}

func (s *ModifyRCDeploymentSetAttributeResponse) SetStatusCode(v int32) *ModifyRCDeploymentSetAttributeResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyRCDeploymentSetAttributeResponse) SetBody(v *ModifyRCDeploymentSetAttributeResponseBody) *ModifyRCDeploymentSetAttributeResponse {
	s.Body = v
	return s
}

func (s *ModifyRCDeploymentSetAttributeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
