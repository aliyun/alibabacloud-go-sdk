// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartBackToBackCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartBackToBackCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartBackToBackCallResponse
	GetStatusCode() *int32
	SetBody(v *StartBackToBackCallResponseBody) *StartBackToBackCallResponse
	GetBody() *StartBackToBackCallResponseBody
}

type StartBackToBackCallResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartBackToBackCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartBackToBackCallResponse) String() string {
	return dara.Prettify(s)
}

func (s StartBackToBackCallResponse) GoString() string {
	return s.String()
}

func (s *StartBackToBackCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartBackToBackCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartBackToBackCallResponse) GetBody() *StartBackToBackCallResponseBody {
	return s.Body
}

func (s *StartBackToBackCallResponse) SetHeaders(v map[string]*string) *StartBackToBackCallResponse {
	s.Headers = v
	return s
}

func (s *StartBackToBackCallResponse) SetStatusCode(v int32) *StartBackToBackCallResponse {
	s.StatusCode = &v
	return s
}

func (s *StartBackToBackCallResponse) SetBody(v *StartBackToBackCallResponseBody) *StartBackToBackCallResponse {
	s.Body = v
	return s
}

func (s *StartBackToBackCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
