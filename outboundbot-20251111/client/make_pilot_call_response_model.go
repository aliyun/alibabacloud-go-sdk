// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMakePilotCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *MakePilotCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *MakePilotCallResponse
	GetStatusCode() *int32
	SetBody(v *MakePilotCallResponseBody) *MakePilotCallResponse
	GetBody() *MakePilotCallResponseBody
}

type MakePilotCallResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MakePilotCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MakePilotCallResponse) String() string {
	return dara.Prettify(s)
}

func (s MakePilotCallResponse) GoString() string {
	return s.String()
}

func (s *MakePilotCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *MakePilotCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *MakePilotCallResponse) GetBody() *MakePilotCallResponseBody {
	return s.Body
}

func (s *MakePilotCallResponse) SetHeaders(v map[string]*string) *MakePilotCallResponse {
	s.Headers = v
	return s
}

func (s *MakePilotCallResponse) SetStatusCode(v int32) *MakePilotCallResponse {
	s.StatusCode = &v
	return s
}

func (s *MakePilotCallResponse) SetBody(v *MakePilotCallResponseBody) *MakePilotCallResponse {
	s.Body = v
	return s
}

func (s *MakePilotCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
