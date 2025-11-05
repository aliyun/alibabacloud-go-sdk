// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountInfoResponseBody) *GetAccountInfoResponse
	GetBody() *GetAccountInfoResponseBody
}

type GetAccountInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccountInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountInfoResponse) GetBody() *GetAccountInfoResponseBody {
	return s.Body
}

func (s *GetAccountInfoResponse) SetHeaders(v map[string]*string) *GetAccountInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccountInfoResponse) SetStatusCode(v int32) *GetAccountInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountInfoResponse) SetBody(v *GetAccountInfoResponseBody) *GetAccountInfoResponse {
	s.Body = v
	return s
}

func (s *GetAccountInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
