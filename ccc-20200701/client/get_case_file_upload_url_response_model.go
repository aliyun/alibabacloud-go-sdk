// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaseFileUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetCaseFileUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetCaseFileUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *GetCaseFileUploadUrlResponseBody) *GetCaseFileUploadUrlResponse
	GetBody() *GetCaseFileUploadUrlResponseBody
}

type GetCaseFileUploadUrlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetCaseFileUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetCaseFileUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s GetCaseFileUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetCaseFileUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetCaseFileUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetCaseFileUploadUrlResponse) GetBody() *GetCaseFileUploadUrlResponseBody {
	return s.Body
}

func (s *GetCaseFileUploadUrlResponse) SetHeaders(v map[string]*string) *GetCaseFileUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetCaseFileUploadUrlResponse) SetStatusCode(v int32) *GetCaseFileUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetCaseFileUploadUrlResponse) SetBody(v *GetCaseFileUploadUrlResponseBody) *GetCaseFileUploadUrlResponse {
	s.Body = v
	return s
}

func (s *GetCaseFileUploadUrlResponse) Validate() error {
	return dara.Validate(s)
}
