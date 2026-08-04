// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryEncryptedAccountProfileInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryEncryptedAccountProfileInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryEncryptedAccountProfileInfoResponse
	GetStatusCode() *int32
	SetBody(v *QueryEncryptedAccountProfileInfoResponseBody) *QueryEncryptedAccountProfileInfoResponse
	GetBody() *QueryEncryptedAccountProfileInfoResponseBody
}

type QueryEncryptedAccountProfileInfoResponse struct {
	Headers    map[string]*string                            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryEncryptedAccountProfileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryEncryptedAccountProfileInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryEncryptedAccountProfileInfoResponse) GoString() string {
	return s.String()
}

func (s *QueryEncryptedAccountProfileInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryEncryptedAccountProfileInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryEncryptedAccountProfileInfoResponse) GetBody() *QueryEncryptedAccountProfileInfoResponseBody {
	return s.Body
}

func (s *QueryEncryptedAccountProfileInfoResponse) SetHeaders(v map[string]*string) *QueryEncryptedAccountProfileInfoResponse {
	s.Headers = v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponse) SetStatusCode(v int32) *QueryEncryptedAccountProfileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponse) SetBody(v *QueryEncryptedAccountProfileInfoResponseBody) *QueryEncryptedAccountProfileInfoResponse {
	s.Body = v
	return s
}

func (s *QueryEncryptedAccountProfileInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
