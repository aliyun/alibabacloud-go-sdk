// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBatchUploadUrlResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateBatchUploadUrlResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateBatchUploadUrlResponse
	GetStatusCode() *int32
	SetBody(v *CreateBatchUploadUrlResponseBody) *CreateBatchUploadUrlResponse
	GetBody() *CreateBatchUploadUrlResponseBody
}

type CreateBatchUploadUrlResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateBatchUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateBatchUploadUrlResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateBatchUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *CreateBatchUploadUrlResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateBatchUploadUrlResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateBatchUploadUrlResponse) GetBody() *CreateBatchUploadUrlResponseBody {
	return s.Body
}

func (s *CreateBatchUploadUrlResponse) SetHeaders(v map[string]*string) *CreateBatchUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *CreateBatchUploadUrlResponse) SetStatusCode(v int32) *CreateBatchUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateBatchUploadUrlResponse) SetBody(v *CreateBatchUploadUrlResponseBody) *CreateBatchUploadUrlResponse {
	s.Body = v
	return s
}

func (s *CreateBatchUploadUrlResponse) Validate() error {
	return dara.Validate(s)
}
