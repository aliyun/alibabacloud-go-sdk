// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveRCInstancesFromDeploymentSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *RemoveRCInstancesFromDeploymentSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *RemoveRCInstancesFromDeploymentSetResponse
	GetStatusCode() *int32
	SetBody(v *RemoveRCInstancesFromDeploymentSetResponseBody) *RemoveRCInstancesFromDeploymentSetResponse
	GetBody() *RemoveRCInstancesFromDeploymentSetResponseBody
}

type RemoveRCInstancesFromDeploymentSetResponse struct {
	Headers    map[string]*string                              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveRCInstancesFromDeploymentSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveRCInstancesFromDeploymentSetResponse) String() string {
	return dara.Prettify(s)
}

func (s RemoveRCInstancesFromDeploymentSetResponse) GoString() string {
	return s.String()
}

func (s *RemoveRCInstancesFromDeploymentSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *RemoveRCInstancesFromDeploymentSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *RemoveRCInstancesFromDeploymentSetResponse) GetBody() *RemoveRCInstancesFromDeploymentSetResponseBody {
	return s.Body
}

func (s *RemoveRCInstancesFromDeploymentSetResponse) SetHeaders(v map[string]*string) *RemoveRCInstancesFromDeploymentSetResponse {
	s.Headers = v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetResponse) SetStatusCode(v int32) *RemoveRCInstancesFromDeploymentSetResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetResponse) SetBody(v *RemoveRCInstancesFromDeploymentSetResponseBody) *RemoveRCInstancesFromDeploymentSetResponse {
	s.Body = v
	return s
}

func (s *RemoveRCInstancesFromDeploymentSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
