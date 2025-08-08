// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadUserAppToMsaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UploadUserAppToMsaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UploadUserAppToMsaResponse
	GetStatusCode() *int32
	SetBody(v *UploadUserAppToMsaResponseBody) *UploadUserAppToMsaResponse
	GetBody() *UploadUserAppToMsaResponseBody
}

type UploadUserAppToMsaResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UploadUserAppToMsaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UploadUserAppToMsaResponse) String() string {
	return dara.Prettify(s)
}

func (s UploadUserAppToMsaResponse) GoString() string {
	return s.String()
}

func (s *UploadUserAppToMsaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UploadUserAppToMsaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UploadUserAppToMsaResponse) GetBody() *UploadUserAppToMsaResponseBody {
	return s.Body
}

func (s *UploadUserAppToMsaResponse) SetHeaders(v map[string]*string) *UploadUserAppToMsaResponse {
	s.Headers = v
	return s
}

func (s *UploadUserAppToMsaResponse) SetStatusCode(v int32) *UploadUserAppToMsaResponse {
	s.StatusCode = &v
	return s
}

func (s *UploadUserAppToMsaResponse) SetBody(v *UploadUserAppToMsaResponseBody) *UploadUserAppToMsaResponse {
	s.Body = v
	return s
}

func (s *UploadUserAppToMsaResponse) Validate() error {
	return dara.Validate(s)
}
