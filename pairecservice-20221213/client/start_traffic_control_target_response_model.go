// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartTrafficControlTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartTrafficControlTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartTrafficControlTargetResponse
	GetStatusCode() *int32
	SetBody(v *StartTrafficControlTargetResponseBody) *StartTrafficControlTargetResponse
	GetBody() *StartTrafficControlTargetResponseBody
}

type StartTrafficControlTargetResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartTrafficControlTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s StartTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *StartTrafficControlTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartTrafficControlTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartTrafficControlTargetResponse) GetBody() *StartTrafficControlTargetResponseBody {
	return s.Body
}

func (s *StartTrafficControlTargetResponse) SetHeaders(v map[string]*string) *StartTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *StartTrafficControlTargetResponse) SetStatusCode(v int32) *StartTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *StartTrafficControlTargetResponse) SetBody(v *StartTrafficControlTargetResponseBody) *StartTrafficControlTargetResponse {
	s.Body = v
	return s
}

func (s *StartTrafficControlTargetResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
