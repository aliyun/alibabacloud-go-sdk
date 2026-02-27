// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTerminalsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTerminalsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTerminalsResponse
	GetStatusCode() *int32
	SetBody(v *AddTerminalsResponseBody) *AddTerminalsResponse
	GetBody() *AddTerminalsResponseBody
}

type AddTerminalsResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTerminalsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTerminalsResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTerminalsResponse) GoString() string {
	return s.String()
}

func (s *AddTerminalsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTerminalsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTerminalsResponse) GetBody() *AddTerminalsResponseBody {
	return s.Body
}

func (s *AddTerminalsResponse) SetHeaders(v map[string]*string) *AddTerminalsResponse {
	s.Headers = v
	return s
}

func (s *AddTerminalsResponse) SetStatusCode(v int32) *AddTerminalsResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTerminalsResponse) SetBody(v *AddTerminalsResponseBody) *AddTerminalsResponse {
	s.Body = v
	return s
}

func (s *AddTerminalsResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
