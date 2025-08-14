// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateProgramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateProgramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateProgramResponse
	GetStatusCode() *int32
	SetBody(v *CreateProgramResponseBody) *CreateProgramResponse
	GetBody() *CreateProgramResponseBody
}

type CreateProgramResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateProgramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateProgramResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateProgramResponse) GoString() string {
	return s.String()
}

func (s *CreateProgramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateProgramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateProgramResponse) GetBody() *CreateProgramResponseBody {
	return s.Body
}

func (s *CreateProgramResponse) SetHeaders(v map[string]*string) *CreateProgramResponse {
	s.Headers = v
	return s
}

func (s *CreateProgramResponse) SetStatusCode(v int32) *CreateProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateProgramResponse) SetBody(v *CreateProgramResponseBody) *CreateProgramResponse {
	s.Body = v
	return s
}

func (s *CreateProgramResponse) Validate() error {
	return dara.Validate(s)
}
