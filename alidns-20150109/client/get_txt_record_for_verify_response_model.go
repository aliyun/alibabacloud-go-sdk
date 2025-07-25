// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTxtRecordForVerifyResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetTxtRecordForVerifyResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetTxtRecordForVerifyResponse
	GetStatusCode() *int32
	SetBody(v *GetTxtRecordForVerifyResponseBody) *GetTxtRecordForVerifyResponse
	GetBody() *GetTxtRecordForVerifyResponseBody
}

type GetTxtRecordForVerifyResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTxtRecordForVerifyResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTxtRecordForVerifyResponse) String() string {
	return dara.Prettify(s)
}

func (s GetTxtRecordForVerifyResponse) GoString() string {
	return s.String()
}

func (s *GetTxtRecordForVerifyResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetTxtRecordForVerifyResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetTxtRecordForVerifyResponse) GetBody() *GetTxtRecordForVerifyResponseBody {
	return s.Body
}

func (s *GetTxtRecordForVerifyResponse) SetHeaders(v map[string]*string) *GetTxtRecordForVerifyResponse {
	s.Headers = v
	return s
}

func (s *GetTxtRecordForVerifyResponse) SetStatusCode(v int32) *GetTxtRecordForVerifyResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTxtRecordForVerifyResponse) SetBody(v *GetTxtRecordForVerifyResponseBody) *GetTxtRecordForVerifyResponse {
	s.Body = v
	return s
}

func (s *GetTxtRecordForVerifyResponse) Validate() error {
	return dara.Validate(s)
}
