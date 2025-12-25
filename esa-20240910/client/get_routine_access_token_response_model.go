// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetRoutineAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetRoutineAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetRoutineAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetRoutineAccessTokenResponseBody) *GetRoutineAccessTokenResponse
	GetBody() *GetRoutineAccessTokenResponseBody
}

type GetRoutineAccessTokenResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetRoutineAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRoutineAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetRoutineAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GetRoutineAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetRoutineAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetRoutineAccessTokenResponse) GetBody() *GetRoutineAccessTokenResponseBody {
	return s.Body
}

func (s *GetRoutineAccessTokenResponse) SetHeaders(v map[string]*string) *GetRoutineAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GetRoutineAccessTokenResponse) SetStatusCode(v int32) *GetRoutineAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRoutineAccessTokenResponse) SetBody(v *GetRoutineAccessTokenResponseBody) *GetRoutineAccessTokenResponse {
	s.Body = v
	return s
}

func (s *GetRoutineAccessTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
