// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileUploadLimitResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateFileUploadLimitResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateFileUploadLimitResponse
	GetStatusCode() *int32
	SetBody(v *CreateFileUploadLimitResponseBody) *CreateFileUploadLimitResponse
	GetBody() *CreateFileUploadLimitResponseBody
}

type CreateFileUploadLimitResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileUploadLimitResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileUploadLimitResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateFileUploadLimitResponse) GoString() string {
	return s.String()
}

func (s *CreateFileUploadLimitResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateFileUploadLimitResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateFileUploadLimitResponse) GetBody() *CreateFileUploadLimitResponseBody {
	return s.Body
}

func (s *CreateFileUploadLimitResponse) SetHeaders(v map[string]*string) *CreateFileUploadLimitResponse {
	s.Headers = v
	return s
}

func (s *CreateFileUploadLimitResponse) SetStatusCode(v int32) *CreateFileUploadLimitResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileUploadLimitResponse) SetBody(v *CreateFileUploadLimitResponseBody) *CreateFileUploadLimitResponse {
	s.Body = v
	return s
}

func (s *CreateFileUploadLimitResponse) Validate() error {
	return dara.Validate(s)
}
