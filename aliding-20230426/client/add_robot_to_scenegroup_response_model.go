// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddRobotToScenegroupResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddRobotToScenegroupResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddRobotToScenegroupResponse
	GetStatusCode() *int32
	SetBody(v *AddRobotToScenegroupResponseBody) *AddRobotToScenegroupResponse
	GetBody() *AddRobotToScenegroupResponseBody
}

type AddRobotToScenegroupResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddRobotToScenegroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddRobotToScenegroupResponse) String() string {
	return dara.Prettify(s)
}

func (s AddRobotToScenegroupResponse) GoString() string {
	return s.String()
}

func (s *AddRobotToScenegroupResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddRobotToScenegroupResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddRobotToScenegroupResponse) GetBody() *AddRobotToScenegroupResponseBody {
	return s.Body
}

func (s *AddRobotToScenegroupResponse) SetHeaders(v map[string]*string) *AddRobotToScenegroupResponse {
	s.Headers = v
	return s
}

func (s *AddRobotToScenegroupResponse) SetStatusCode(v int32) *AddRobotToScenegroupResponse {
	s.StatusCode = &v
	return s
}

func (s *AddRobotToScenegroupResponse) SetBody(v *AddRobotToScenegroupResponseBody) *AddRobotToScenegroupResponse {
	s.Body = v
	return s
}

func (s *AddRobotToScenegroupResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
