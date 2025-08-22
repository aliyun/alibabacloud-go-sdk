// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetRoutineSubdomainResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetRoutineSubdomainResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetRoutineSubdomainResponse
	GetStatusCode() *int32
	SetBody(v *SetRoutineSubdomainResponseBody) *SetRoutineSubdomainResponse
	GetBody() *SetRoutineSubdomainResponseBody
}

type SetRoutineSubdomainResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetRoutineSubdomainResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetRoutineSubdomainResponse) String() string {
	return dara.Prettify(s)
}

func (s SetRoutineSubdomainResponse) GoString() string {
	return s.String()
}

func (s *SetRoutineSubdomainResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetRoutineSubdomainResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetRoutineSubdomainResponse) GetBody() *SetRoutineSubdomainResponseBody {
	return s.Body
}

func (s *SetRoutineSubdomainResponse) SetHeaders(v map[string]*string) *SetRoutineSubdomainResponse {
	s.Headers = v
	return s
}

func (s *SetRoutineSubdomainResponse) SetStatusCode(v int32) *SetRoutineSubdomainResponse {
	s.StatusCode = &v
	return s
}

func (s *SetRoutineSubdomainResponse) SetBody(v *SetRoutineSubdomainResponseBody) *SetRoutineSubdomainResponse {
	s.Body = v
	return s
}

func (s *SetRoutineSubdomainResponse) Validate() error {
	return dara.Validate(s)
}
