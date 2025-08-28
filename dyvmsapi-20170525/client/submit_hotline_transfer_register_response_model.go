// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitHotlineTransferRegisterResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SubmitHotlineTransferRegisterResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SubmitHotlineTransferRegisterResponse
	GetStatusCode() *int32
	SetBody(v *SubmitHotlineTransferRegisterResponseBody) *SubmitHotlineTransferRegisterResponse
	GetBody() *SubmitHotlineTransferRegisterResponseBody
}

type SubmitHotlineTransferRegisterResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SubmitHotlineTransferRegisterResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SubmitHotlineTransferRegisterResponse) String() string {
	return dara.Prettify(s)
}

func (s SubmitHotlineTransferRegisterResponse) GoString() string {
	return s.String()
}

func (s *SubmitHotlineTransferRegisterResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SubmitHotlineTransferRegisterResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SubmitHotlineTransferRegisterResponse) GetBody() *SubmitHotlineTransferRegisterResponseBody {
	return s.Body
}

func (s *SubmitHotlineTransferRegisterResponse) SetHeaders(v map[string]*string) *SubmitHotlineTransferRegisterResponse {
	s.Headers = v
	return s
}

func (s *SubmitHotlineTransferRegisterResponse) SetStatusCode(v int32) *SubmitHotlineTransferRegisterResponse {
	s.StatusCode = &v
	return s
}

func (s *SubmitHotlineTransferRegisterResponse) SetBody(v *SubmitHotlineTransferRegisterResponseBody) *SubmitHotlineTransferRegisterResponse {
	s.Body = v
	return s
}

func (s *SubmitHotlineTransferRegisterResponse) Validate() error {
	return dara.Validate(s)
}
