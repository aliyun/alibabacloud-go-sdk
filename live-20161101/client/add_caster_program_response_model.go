// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddCasterProgramResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddCasterProgramResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddCasterProgramResponse
	GetStatusCode() *int32
	SetBody(v *AddCasterProgramResponseBody) *AddCasterProgramResponse
	GetBody() *AddCasterProgramResponseBody
}

type AddCasterProgramResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddCasterProgramResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddCasterProgramResponse) String() string {
	return dara.Prettify(s)
}

func (s AddCasterProgramResponse) GoString() string {
	return s.String()
}

func (s *AddCasterProgramResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddCasterProgramResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddCasterProgramResponse) GetBody() *AddCasterProgramResponseBody {
	return s.Body
}

func (s *AddCasterProgramResponse) SetHeaders(v map[string]*string) *AddCasterProgramResponse {
	s.Headers = v
	return s
}

func (s *AddCasterProgramResponse) SetStatusCode(v int32) *AddCasterProgramResponse {
	s.StatusCode = &v
	return s
}

func (s *AddCasterProgramResponse) SetBody(v *AddCasterProgramResponseBody) *AddCasterProgramResponse {
	s.Body = v
	return s
}

func (s *AddCasterProgramResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
