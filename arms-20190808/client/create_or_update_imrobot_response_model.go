// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOrUpdateIMRobotResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateOrUpdateIMRobotResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateOrUpdateIMRobotResponse
	GetStatusCode() *int32
	SetBody(v *CreateOrUpdateIMRobotResponseBody) *CreateOrUpdateIMRobotResponse
	GetBody() *CreateOrUpdateIMRobotResponseBody
}

type CreateOrUpdateIMRobotResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrUpdateIMRobotResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrUpdateIMRobotResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateOrUpdateIMRobotResponse) GoString() string {
	return s.String()
}

func (s *CreateOrUpdateIMRobotResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateOrUpdateIMRobotResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateOrUpdateIMRobotResponse) GetBody() *CreateOrUpdateIMRobotResponseBody {
	return s.Body
}

func (s *CreateOrUpdateIMRobotResponse) SetHeaders(v map[string]*string) *CreateOrUpdateIMRobotResponse {
	s.Headers = v
	return s
}

func (s *CreateOrUpdateIMRobotResponse) SetStatusCode(v int32) *CreateOrUpdateIMRobotResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrUpdateIMRobotResponse) SetBody(v *CreateOrUpdateIMRobotResponseBody) *CreateOrUpdateIMRobotResponse {
	s.Body = v
	return s
}

func (s *CreateOrUpdateIMRobotResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
