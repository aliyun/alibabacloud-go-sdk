// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iShareRCDeploymentSetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ShareRCDeploymentSetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ShareRCDeploymentSetResponse
	GetStatusCode() *int32
	SetBody(v *ShareRCDeploymentSetResponseBody) *ShareRCDeploymentSetResponse
	GetBody() *ShareRCDeploymentSetResponseBody
}

type ShareRCDeploymentSetResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareRCDeploymentSetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ShareRCDeploymentSetResponse) String() string {
	return dara.Prettify(s)
}

func (s ShareRCDeploymentSetResponse) GoString() string {
	return s.String()
}

func (s *ShareRCDeploymentSetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ShareRCDeploymentSetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ShareRCDeploymentSetResponse) GetBody() *ShareRCDeploymentSetResponseBody {
	return s.Body
}

func (s *ShareRCDeploymentSetResponse) SetHeaders(v map[string]*string) *ShareRCDeploymentSetResponse {
	s.Headers = v
	return s
}

func (s *ShareRCDeploymentSetResponse) SetStatusCode(v int32) *ShareRCDeploymentSetResponse {
	s.StatusCode = &v
	return s
}

func (s *ShareRCDeploymentSetResponse) SetBody(v *ShareRCDeploymentSetResponseBody) *ShareRCDeploymentSetResponse {
	s.Body = v
	return s
}

func (s *ShareRCDeploymentSetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
