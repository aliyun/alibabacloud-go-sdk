// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetWorkloadAccessTokenResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetWorkloadAccessTokenResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetWorkloadAccessTokenResponse
	GetStatusCode() *int32
	SetBody(v *GetWorkloadAccessTokenResponseBody) *GetWorkloadAccessTokenResponse
	GetBody() *GetWorkloadAccessTokenResponseBody
}

type GetWorkloadAccessTokenResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetWorkloadAccessTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetWorkloadAccessTokenResponse) String() string {
	return dara.Prettify(s)
}

func (s GetWorkloadAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *GetWorkloadAccessTokenResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetWorkloadAccessTokenResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetWorkloadAccessTokenResponse) GetBody() *GetWorkloadAccessTokenResponseBody {
	return s.Body
}

func (s *GetWorkloadAccessTokenResponse) SetHeaders(v map[string]*string) *GetWorkloadAccessTokenResponse {
	s.Headers = v
	return s
}

func (s *GetWorkloadAccessTokenResponse) SetStatusCode(v int32) *GetWorkloadAccessTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetWorkloadAccessTokenResponse) SetBody(v *GetWorkloadAccessTokenResponseBody) *GetWorkloadAccessTokenResponse {
	s.Body = v
	return s
}

func (s *GetWorkloadAccessTokenResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
