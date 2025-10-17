// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMPULayoutResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateMPULayoutResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateMPULayoutResponse
	GetStatusCode() *int32
	SetBody(v *CreateMPULayoutResponseBody) *CreateMPULayoutResponse
	GetBody() *CreateMPULayoutResponseBody
}

type CreateMPULayoutResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateMPULayoutResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateMPULayoutResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateMPULayoutResponse) GoString() string {
	return s.String()
}

func (s *CreateMPULayoutResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateMPULayoutResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateMPULayoutResponse) GetBody() *CreateMPULayoutResponseBody {
	return s.Body
}

func (s *CreateMPULayoutResponse) SetHeaders(v map[string]*string) *CreateMPULayoutResponse {
	s.Headers = v
	return s
}

func (s *CreateMPULayoutResponse) SetStatusCode(v int32) *CreateMPULayoutResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateMPULayoutResponse) SetBody(v *CreateMPULayoutResponseBody) *CreateMPULayoutResponse {
	s.Body = v
	return s
}

func (s *CreateMPULayoutResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
