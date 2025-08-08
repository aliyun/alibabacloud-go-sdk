// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserAppUploadProcessInMsaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetUserAppUploadProcessInMsaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetUserAppUploadProcessInMsaResponse
	GetStatusCode() *int32
	SetBody(v *GetUserAppUploadProcessInMsaResponseBody) *GetUserAppUploadProcessInMsaResponse
	GetBody() *GetUserAppUploadProcessInMsaResponseBody
}

type GetUserAppUploadProcessInMsaResponse struct {
	Headers    map[string]*string                        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUserAppUploadProcessInMsaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserAppUploadProcessInMsaResponse) String() string {
	return dara.Prettify(s)
}

func (s GetUserAppUploadProcessInMsaResponse) GoString() string {
	return s.String()
}

func (s *GetUserAppUploadProcessInMsaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetUserAppUploadProcessInMsaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetUserAppUploadProcessInMsaResponse) GetBody() *GetUserAppUploadProcessInMsaResponseBody {
	return s.Body
}

func (s *GetUserAppUploadProcessInMsaResponse) SetHeaders(v map[string]*string) *GetUserAppUploadProcessInMsaResponse {
	s.Headers = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponse) SetStatusCode(v int32) *GetUserAppUploadProcessInMsaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponse) SetBody(v *GetUserAppUploadProcessInMsaResponseBody) *GetUserAppUploadProcessInMsaResponse {
	s.Body = v
	return s
}

func (s *GetUserAppUploadProcessInMsaResponse) Validate() error {
	return dara.Validate(s)
}
