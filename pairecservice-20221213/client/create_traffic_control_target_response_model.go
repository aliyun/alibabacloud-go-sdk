// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficControlTargetResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrafficControlTargetResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrafficControlTargetResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrafficControlTargetResponseBody) *CreateTrafficControlTargetResponse
	GetBody() *CreateTrafficControlTargetResponseBody
}

type CreateTrafficControlTargetResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficControlTargetResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficControlTargetResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlTargetResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTargetResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrafficControlTargetResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrafficControlTargetResponse) GetBody() *CreateTrafficControlTargetResponseBody {
	return s.Body
}

func (s *CreateTrafficControlTargetResponse) SetHeaders(v map[string]*string) *CreateTrafficControlTargetResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficControlTargetResponse) SetStatusCode(v int32) *CreateTrafficControlTargetResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficControlTargetResponse) SetBody(v *CreateTrafficControlTargetResponseBody) *CreateTrafficControlTargetResponse {
	s.Body = v
	return s
}

func (s *CreateTrafficControlTargetResponse) Validate() error {
	return dara.Validate(s)
}
