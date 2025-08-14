// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetProgramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetProgramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetProgramResponse
	GetStatusCode() *int32
	SetBody(v *GetProgramResponseBody) *GetProgramResponse
	GetBody() *GetProgramResponseBody
}

type GetProgramResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetProgramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetProgramResponse) String() string {
	return dara.Prettify(s)
}

func (s GetProgramResponse) GoString() string {
	return s.String()
}

func (s *GetProgramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetProgramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetProgramResponse) GetBody() *GetProgramResponseBody {
	return s.Body
}

func (s *GetProgramResponse) SetHeaders(v map[string]*string) *GetProgramResponse {
	s.Headers = v
	return s
}

func (s *GetProgramResponse) SetStatusCode(v int32) *GetProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *GetProgramResponse) SetBody(v *GetProgramResponseBody) *GetProgramResponse {
	s.Body = v
	return s
}

func (s *GetProgramResponse) Validate() error {
	return dara.Validate(s)
}
