// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficControlTaskResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrafficControlTaskResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrafficControlTaskResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrafficControlTaskResponseBody) *CreateTrafficControlTaskResponse
	GetBody() *CreateTrafficControlTaskResponseBody
}

type CreateTrafficControlTaskResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficControlTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficControlTaskResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlTaskResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrafficControlTaskResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrafficControlTaskResponse) GetBody() *CreateTrafficControlTaskResponseBody {
	return s.Body
}

func (s *CreateTrafficControlTaskResponse) SetHeaders(v map[string]*string) *CreateTrafficControlTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficControlTaskResponse) SetStatusCode(v int32) *CreateTrafficControlTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficControlTaskResponse) SetBody(v *CreateTrafficControlTaskResponseBody) *CreateTrafficControlTaskResponse {
	s.Body = v
	return s
}

func (s *CreateTrafficControlTaskResponse) Validate() error {
	return dara.Validate(s)
}
