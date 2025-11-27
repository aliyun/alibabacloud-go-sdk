// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRCDeploymentSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteRCDeploymentSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteRCDeploymentSetResponse
	GetStatusCode() *int32
	SetBody(v *DeleteRCDeploymentSetResponseBody) *DeleteRCDeploymentSetResponse
	GetBody() *DeleteRCDeploymentSetResponseBody
}

type DeleteRCDeploymentSetResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteRCDeploymentSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteRCDeploymentSetResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteRCDeploymentSetResponse) GoString() string {
	return s.String()
}

func (s *DeleteRCDeploymentSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteRCDeploymentSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteRCDeploymentSetResponse) GetBody() *DeleteRCDeploymentSetResponseBody {
	return s.Body
}

func (s *DeleteRCDeploymentSetResponse) SetHeaders(v map[string]*string) *DeleteRCDeploymentSetResponse {
	s.Headers = v
	return s
}

func (s *DeleteRCDeploymentSetResponse) SetStatusCode(v int32) *DeleteRCDeploymentSetResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteRCDeploymentSetResponse) SetBody(v *DeleteRCDeploymentSetResponseBody) *DeleteRCDeploymentSetResponse {
	s.Body = v
	return s
}

func (s *DeleteRCDeploymentSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
