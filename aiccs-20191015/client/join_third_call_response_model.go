// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iJoinThirdCallResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *JoinThirdCallResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *JoinThirdCallResponse
	GetStatusCode() *int32
	SetBody(v *JoinThirdCallResponseBody) *JoinThirdCallResponse
	GetBody() *JoinThirdCallResponseBody
}

type JoinThirdCallResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *JoinThirdCallResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s JoinThirdCallResponse) String() string {
	return dara.Prettify(s)
}

func (s JoinThirdCallResponse) GoString() string {
	return s.String()
}

func (s *JoinThirdCallResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *JoinThirdCallResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *JoinThirdCallResponse) GetBody() *JoinThirdCallResponseBody {
	return s.Body
}

func (s *JoinThirdCallResponse) SetHeaders(v map[string]*string) *JoinThirdCallResponse {
	s.Headers = v
	return s
}

func (s *JoinThirdCallResponse) SetStatusCode(v int32) *JoinThirdCallResponse {
	s.StatusCode = &v
	return s
}

func (s *JoinThirdCallResponse) SetBody(v *JoinThirdCallResponseBody) *JoinThirdCallResponse {
	s.Body = v
	return s
}

func (s *JoinThirdCallResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
