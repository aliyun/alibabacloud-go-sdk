// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartEaisEiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartEaisEiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartEaisEiResponse
	GetStatusCode() *int32
	SetBody(v *StartEaisEiResponseBody) *StartEaisEiResponse
	GetBody() *StartEaisEiResponseBody
}

type StartEaisEiResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartEaisEiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartEaisEiResponse) String() string {
	return dara.Prettify(s)
}

func (s StartEaisEiResponse) GoString() string {
	return s.String()
}

func (s *StartEaisEiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartEaisEiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartEaisEiResponse) GetBody() *StartEaisEiResponseBody {
	return s.Body
}

func (s *StartEaisEiResponse) SetHeaders(v map[string]*string) *StartEaisEiResponse {
	s.Headers = v
	return s
}

func (s *StartEaisEiResponse) SetStatusCode(v int32) *StartEaisEiResponse {
	s.StatusCode = &v
	return s
}

func (s *StartEaisEiResponse) SetBody(v *StartEaisEiResponseBody) *StartEaisEiResponse {
	s.Body = v
	return s
}

func (s *StartEaisEiResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
