// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountMFAInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountMFAInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountMFAInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountMFAInfoResponseBody) *GetAccountMFAInfoResponse
	GetBody() *GetAccountMFAInfoResponseBody
}

type GetAccountMFAInfoResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountMFAInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountMFAInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountMFAInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccountMFAInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountMFAInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountMFAInfoResponse) GetBody() *GetAccountMFAInfoResponseBody {
	return s.Body
}

func (s *GetAccountMFAInfoResponse) SetHeaders(v map[string]*string) *GetAccountMFAInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccountMFAInfoResponse) SetStatusCode(v int32) *GetAccountMFAInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountMFAInfoResponse) SetBody(v *GetAccountMFAInfoResponseBody) *GetAccountMFAInfoResponse {
	s.Body = v
	return s
}

func (s *GetAccountMFAInfoResponse) Validate() error {
	return dara.Validate(s)
}
