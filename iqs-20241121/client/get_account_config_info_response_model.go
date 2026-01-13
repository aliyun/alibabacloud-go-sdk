// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAccountConfigInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAccountConfigInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAccountConfigInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetAccountConfigInfoResponseBody) *GetAccountConfigInfoResponse
	GetBody() *GetAccountConfigInfoResponseBody
}

type GetAccountConfigInfoResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAccountConfigInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAccountConfigInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAccountConfigInfoResponse) GoString() string {
	return s.String()
}

func (s *GetAccountConfigInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAccountConfigInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAccountConfigInfoResponse) GetBody() *GetAccountConfigInfoResponseBody {
	return s.Body
}

func (s *GetAccountConfigInfoResponse) SetHeaders(v map[string]*string) *GetAccountConfigInfoResponse {
	s.Headers = v
	return s
}

func (s *GetAccountConfigInfoResponse) SetStatusCode(v int32) *GetAccountConfigInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAccountConfigInfoResponse) SetBody(v *GetAccountConfigInfoResponseBody) *GetAccountConfigInfoResponse {
	s.Body = v
	return s
}

func (s *GetAccountConfigInfoResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
