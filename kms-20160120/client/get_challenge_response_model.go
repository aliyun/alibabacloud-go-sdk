// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetChallengeResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetChallengeResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetChallengeResponse
	GetStatusCode() *int32
	SetBody(v *GetChallengeResponseBody) *GetChallengeResponse
	GetBody() *GetChallengeResponseBody
}

type GetChallengeResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetChallengeResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetChallengeResponse) String() string {
	return dara.Prettify(s)
}

func (s GetChallengeResponse) GoString() string {
	return s.String()
}

func (s *GetChallengeResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetChallengeResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetChallengeResponse) GetBody() *GetChallengeResponseBody {
	return s.Body
}

func (s *GetChallengeResponse) SetHeaders(v map[string]*string) *GetChallengeResponse {
	s.Headers = v
	return s
}

func (s *GetChallengeResponse) SetStatusCode(v int32) *GetChallengeResponse {
	s.StatusCode = &v
	return s
}

func (s *GetChallengeResponse) SetBody(v *GetChallengeResponseBody) *GetChallengeResponse {
	s.Body = v
	return s
}

func (s *GetChallengeResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
