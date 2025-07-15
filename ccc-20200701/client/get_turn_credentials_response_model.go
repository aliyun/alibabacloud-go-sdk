// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTurnCredentialsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTurnCredentialsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTurnCredentialsResponse
	GetStatusCode() *int32
	SetBody(v *GetTurnCredentialsResponseBody) *GetTurnCredentialsResponse
	GetBody() *GetTurnCredentialsResponseBody
}

type GetTurnCredentialsResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTurnCredentialsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTurnCredentialsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTurnCredentialsResponse) GoString() string {
	return s.String()
}

func (s *GetTurnCredentialsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTurnCredentialsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTurnCredentialsResponse) GetBody() *GetTurnCredentialsResponseBody {
	return s.Body
}

func (s *GetTurnCredentialsResponse) SetHeaders(v map[string]*string) *GetTurnCredentialsResponse {
	s.Headers = v
	return s
}

func (s *GetTurnCredentialsResponse) SetStatusCode(v int32) *GetTurnCredentialsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTurnCredentialsResponse) SetBody(v *GetTurnCredentialsResponseBody) *GetTurnCredentialsResponse {
	s.Body = v
	return s
}

func (s *GetTurnCredentialsResponse) Validate() error {
	return dara.Validate(s)
}
