// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *StartCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *StartCallResponse
	GetStatusCode() *int32
	SetBody(v *StartCallResponseBody) *StartCallResponse
	GetBody() *StartCallResponseBody
}

type StartCallResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *StartCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s StartCallResponse) String() string {
	return dara.Prettify(s)
}

func (s StartCallResponse) GoString() string {
	return s.String()
}

func (s *StartCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *StartCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *StartCallResponse) GetBody() *StartCallResponseBody {
	return s.Body
}

func (s *StartCallResponse) SetHeaders(v map[string]*string) *StartCallResponse {
	s.Headers = v
	return s
}

func (s *StartCallResponse) SetStatusCode(v int32) *StartCallResponse {
	s.StatusCode = &v
	return s
}

func (s *StartCallResponse) SetBody(v *StartCallResponseBody) *StartCallResponse {
	s.Body = v
	return s
}

func (s *StartCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
