// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateFileMetaResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateFileMetaResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateFileMetaResponse
	GetStatusCode() *int32
	SetBody(v *UpdateFileMetaResponseBody) *UpdateFileMetaResponse
	GetBody() *UpdateFileMetaResponseBody
}

type UpdateFileMetaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFileMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileMetaResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateFileMetaResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileMetaResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateFileMetaResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateFileMetaResponse) GetBody() *UpdateFileMetaResponseBody {
	return s.Body
}

func (s *UpdateFileMetaResponse) SetHeaders(v map[string]*string) *UpdateFileMetaResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileMetaResponse) SetStatusCode(v int32) *UpdateFileMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileMetaResponse) SetBody(v *UpdateFileMetaResponseBody) *UpdateFileMetaResponse {
	s.Body = v
	return s
}

func (s *UpdateFileMetaResponse) Validate() error {
	return dara.Validate(s)
}
