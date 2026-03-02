// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLoginUserInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLoginUserInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLoginUserInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetLoginUserInfoResponseBody) *GetLoginUserInfoResponse
	GetBody() *GetLoginUserInfoResponseBody
}

type GetLoginUserInfoResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLoginUserInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLoginUserInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLoginUserInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLoginUserInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLoginUserInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLoginUserInfoResponse) GetBody() *GetLoginUserInfoResponseBody {
	return s.Body
}

func (s *GetLoginUserInfoResponse) SetHeaders(v map[string]*string) *GetLoginUserInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLoginUserInfoResponse) SetStatusCode(v int32) *GetLoginUserInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLoginUserInfoResponse) SetBody(v *GetLoginUserInfoResponseBody) *GetLoginUserInfoResponse {
	s.Body = v
	return s
}

func (s *GetLoginUserInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
