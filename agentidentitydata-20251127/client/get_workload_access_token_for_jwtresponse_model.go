// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadAccessTokenForJWTResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkloadAccessTokenForJWTResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkloadAccessTokenForJWTResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkloadAccessTokenForJWTResponseBody) *GetWorkloadAccessTokenForJWTResponse
	GetBody() *GetWorkloadAccessTokenForJWTResponseBody
}

type GetWorkloadAccessTokenForJWTResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkloadAccessTokenForJWTResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkloadAccessTokenForJWTResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenForJWTResponse) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenForJWTResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkloadAccessTokenForJWTResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkloadAccessTokenForJWTResponse) GetBody() *GetWorkloadAccessTokenForJWTResponseBody {
	return s.Body
}

func (s *GetWorkloadAccessTokenForJWTResponse) SetHeaders(v map[string]*string) *GetWorkloadAccessTokenForJWTResponse {
	s.Headers = v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponse) SetStatusCode(v int32) *GetWorkloadAccessTokenForJWTResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponse) SetBody(v *GetWorkloadAccessTokenForJWTResponseBody) *GetWorkloadAccessTokenForJWTResponse {
	s.Body = v
	return s
}

func (s *GetWorkloadAccessTokenForJWTResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
