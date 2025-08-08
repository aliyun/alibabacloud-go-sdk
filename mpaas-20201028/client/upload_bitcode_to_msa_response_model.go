// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadBitcodeToMsaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadBitcodeToMsaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadBitcodeToMsaResponse
	GetStatusCode() *int32
	SetBody(v *UploadBitcodeToMsaResponseBody) *UploadBitcodeToMsaResponse
	GetBody() *UploadBitcodeToMsaResponseBody
}

type UploadBitcodeToMsaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadBitcodeToMsaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadBitcodeToMsaResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadBitcodeToMsaResponse) GoString() string {
	return s.String()
}

func (s *UploadBitcodeToMsaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadBitcodeToMsaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadBitcodeToMsaResponse) GetBody() *UploadBitcodeToMsaResponseBody {
	return s.Body
}

func (s *UploadBitcodeToMsaResponse) SetHeaders(v map[string]*string) *UploadBitcodeToMsaResponse {
	s.Headers = v
	return s
}

func (s *UploadBitcodeToMsaResponse) SetStatusCode(v int32) *UploadBitcodeToMsaResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadBitcodeToMsaResponse) SetBody(v *UploadBitcodeToMsaResponseBody) *UploadBitcodeToMsaResponse {
	s.Body = v
	return s
}

func (s *UploadBitcodeToMsaResponse) Validate() error {
	return dara.Validate(s)
}
