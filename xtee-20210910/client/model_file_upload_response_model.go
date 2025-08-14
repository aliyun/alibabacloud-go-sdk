// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelFileUploadResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModelFileUploadResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModelFileUploadResponse
	GetStatusCode() *int32
	SetBody(v *ModelFileUploadResponseBody) *ModelFileUploadResponse
	GetBody() *ModelFileUploadResponseBody
}

type ModelFileUploadResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModelFileUploadResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModelFileUploadResponse) String() string {
	return dara.Prettify(s)
}

func (s ModelFileUploadResponse) GoString() string {
	return s.String()
}

func (s *ModelFileUploadResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModelFileUploadResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModelFileUploadResponse) GetBody() *ModelFileUploadResponseBody {
	return s.Body
}

func (s *ModelFileUploadResponse) SetHeaders(v map[string]*string) *ModelFileUploadResponse {
	s.Headers = v
	return s
}

func (s *ModelFileUploadResponse) SetStatusCode(v int32) *ModelFileUploadResponse {
	s.StatusCode = &v
	return s
}

func (s *ModelFileUploadResponse) SetBody(v *ModelFileUploadResponseBody) *ModelFileUploadResponse {
	s.Body = v
	return s
}

func (s *ModelFileUploadResponse) Validate() error {
	return dara.Validate(s)
}
