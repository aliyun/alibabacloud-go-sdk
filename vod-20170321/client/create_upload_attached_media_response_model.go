// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateUploadAttachedMediaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateUploadAttachedMediaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateUploadAttachedMediaResponse
	GetStatusCode() *int32
	SetBody(v *CreateUploadAttachedMediaResponseBody) *CreateUploadAttachedMediaResponse
	GetBody() *CreateUploadAttachedMediaResponseBody
}

type CreateUploadAttachedMediaResponse struct {
	Headers    map[string]*string                     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUploadAttachedMediaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUploadAttachedMediaResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateUploadAttachedMediaResponse) GoString() string {
	return s.String()
}

func (s *CreateUploadAttachedMediaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateUploadAttachedMediaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateUploadAttachedMediaResponse) GetBody() *CreateUploadAttachedMediaResponseBody {
	return s.Body
}

func (s *CreateUploadAttachedMediaResponse) SetHeaders(v map[string]*string) *CreateUploadAttachedMediaResponse {
	s.Headers = v
	return s
}

func (s *CreateUploadAttachedMediaResponse) SetStatusCode(v int32) *CreateUploadAttachedMediaResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUploadAttachedMediaResponse) SetBody(v *CreateUploadAttachedMediaResponseBody) *CreateUploadAttachedMediaResponse {
	s.Body = v
	return s
}

func (s *CreateUploadAttachedMediaResponse) Validate() error {
	return dara.Validate(s)
}
