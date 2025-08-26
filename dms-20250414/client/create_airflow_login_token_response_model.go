// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAirflowLoginTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateAirflowLoginTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateAirflowLoginTokenResponse
	GetStatusCode() *int32
	SetBody(v *CreateAirflowLoginTokenResponseBody) *CreateAirflowLoginTokenResponse
	GetBody() *CreateAirflowLoginTokenResponseBody
}

type CreateAirflowLoginTokenResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateAirflowLoginTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateAirflowLoginTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateAirflowLoginTokenResponse) GoString() string {
	return s.String()
}

func (s *CreateAirflowLoginTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateAirflowLoginTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateAirflowLoginTokenResponse) GetBody() *CreateAirflowLoginTokenResponseBody {
	return s.Body
}

func (s *CreateAirflowLoginTokenResponse) SetHeaders(v map[string]*string) *CreateAirflowLoginTokenResponse {
	s.Headers = v
	return s
}

func (s *CreateAirflowLoginTokenResponse) SetStatusCode(v int32) *CreateAirflowLoginTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateAirflowLoginTokenResponse) SetBody(v *CreateAirflowLoginTokenResponseBody) *CreateAirflowLoginTokenResponse {
	s.Body = v
	return s
}

func (s *CreateAirflowLoginTokenResponse) Validate() error {
	return dara.Validate(s)
}
