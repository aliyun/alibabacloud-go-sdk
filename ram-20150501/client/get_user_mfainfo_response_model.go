// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserMFAInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserMFAInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserMFAInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetUserMFAInfoResponseBody) *GetUserMFAInfoResponse
	GetBody() *GetUserMFAInfoResponseBody
}

type GetUserMFAInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserMFAInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserMFAInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserMFAInfoResponse) GoString() string {
	return s.String()
}

func (s *GetUserMFAInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserMFAInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserMFAInfoResponse) GetBody() *GetUserMFAInfoResponseBody {
	return s.Body
}

func (s *GetUserMFAInfoResponse) SetHeaders(v map[string]*string) *GetUserMFAInfoResponse {
	s.Headers = v
	return s
}

func (s *GetUserMFAInfoResponse) SetStatusCode(v int32) *GetUserMFAInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserMFAInfoResponse) SetBody(v *GetUserMFAInfoResponseBody) *GetUserMFAInfoResponse {
	s.Body = v
	return s
}

func (s *GetUserMFAInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
