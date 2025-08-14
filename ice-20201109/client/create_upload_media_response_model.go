// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUploadMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUploadMediaResponse
	GetStatusCode() *int32
	SetBody(v *CreateUploadMediaResponseBody) *CreateUploadMediaResponse
	GetBody() *CreateUploadMediaResponseBody
}

type CreateUploadMediaResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUploadMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUploadMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadMediaResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUploadMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUploadMediaResponse) GetBody() *CreateUploadMediaResponseBody {
	return s.Body
}

func (s *CreateUploadMediaResponse) SetHeaders(v map[string]*string) *CreateUploadMediaResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadMediaResponse) SetStatusCode(v int32) *CreateUploadMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUploadMediaResponse) SetBody(v *CreateUploadMediaResponseBody) *CreateUploadMediaResponse {
	s.Body = v
	return s
}

func (s *CreateUploadMediaResponse) Validate() error {
	return dara.Validate(s)
}
