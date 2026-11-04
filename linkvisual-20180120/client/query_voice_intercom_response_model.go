// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryVoiceIntercomResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryVoiceIntercomResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryVoiceIntercomResponse
	GetStatusCode() *int32
	SetBody(v *QueryVoiceIntercomResponseBody) *QueryVoiceIntercomResponse
	GetBody() *QueryVoiceIntercomResponseBody
}

type QueryVoiceIntercomResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryVoiceIntercomResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryVoiceIntercomResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryVoiceIntercomResponse) GoString() string {
	return s.String()
}

func (s *QueryVoiceIntercomResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryVoiceIntercomResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryVoiceIntercomResponse) GetBody() *QueryVoiceIntercomResponseBody {
	return s.Body
}

func (s *QueryVoiceIntercomResponse) SetHeaders(v map[string]*string) *QueryVoiceIntercomResponse {
	s.Headers = v
	return s
}

func (s *QueryVoiceIntercomResponse) SetStatusCode(v int32) *QueryVoiceIntercomResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryVoiceIntercomResponse) SetBody(v *QueryVoiceIntercomResponseBody) *QueryVoiceIntercomResponse {
	s.Body = v
	return s
}

func (s *QueryVoiceIntercomResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
