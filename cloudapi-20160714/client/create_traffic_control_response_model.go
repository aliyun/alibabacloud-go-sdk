// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTrafficControlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateTrafficControlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateTrafficControlResponse
	GetStatusCode() *int32
	SetBody(v *CreateTrafficControlResponseBody) *CreateTrafficControlResponse
	GetBody() *CreateTrafficControlResponseBody
}

type CreateTrafficControlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateTrafficControlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateTrafficControlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateTrafficControlResponse) GoString() string {
	return s.String()
}

func (s *CreateTrafficControlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateTrafficControlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateTrafficControlResponse) GetBody() *CreateTrafficControlResponseBody {
	return s.Body
}

func (s *CreateTrafficControlResponse) SetHeaders(v map[string]*string) *CreateTrafficControlResponse {
	s.Headers = v
	return s
}

func (s *CreateTrafficControlResponse) SetStatusCode(v int32) *CreateTrafficControlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateTrafficControlResponse) SetBody(v *CreateTrafficControlResponseBody) *CreateTrafficControlResponse {
	s.Body = v
	return s
}

func (s *CreateTrafficControlResponse) Validate() error {
	return dara.Validate(s)
}
