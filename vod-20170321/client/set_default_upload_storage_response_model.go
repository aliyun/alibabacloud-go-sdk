// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetDefaultUploadStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetDefaultUploadStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetDefaultUploadStorageResponse
	GetStatusCode() *int32
	SetBody(v *SetDefaultUploadStorageResponseBody) *SetDefaultUploadStorageResponse
	GetBody() *SetDefaultUploadStorageResponseBody
}

type SetDefaultUploadStorageResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetDefaultUploadStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetDefaultUploadStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s SetDefaultUploadStorageResponse) GoString() string {
	return s.String()
}

func (s *SetDefaultUploadStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetDefaultUploadStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetDefaultUploadStorageResponse) GetBody() *SetDefaultUploadStorageResponseBody {
	return s.Body
}

func (s *SetDefaultUploadStorageResponse) SetHeaders(v map[string]*string) *SetDefaultUploadStorageResponse {
	s.Headers = v
	return s
}

func (s *SetDefaultUploadStorageResponse) SetStatusCode(v int32) *SetDefaultUploadStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *SetDefaultUploadStorageResponse) SetBody(v *SetDefaultUploadStorageResponseBody) *SetDefaultUploadStorageResponse {
	s.Body = v
	return s
}

func (s *SetDefaultUploadStorageResponse) Validate() error {
	return dara.Validate(s)
}
