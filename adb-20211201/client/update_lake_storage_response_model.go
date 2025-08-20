// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateLakeStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateLakeStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateLakeStorageResponse
	GetStatusCode() *int32
	SetBody(v *UpdateLakeStorageResponseBody) *UpdateLakeStorageResponse
	GetBody() *UpdateLakeStorageResponseBody
}

type UpdateLakeStorageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateLakeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateLakeStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateLakeStorageResponse) GoString() string {
	return s.String()
}

func (s *UpdateLakeStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateLakeStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateLakeStorageResponse) GetBody() *UpdateLakeStorageResponseBody {
	return s.Body
}

func (s *UpdateLakeStorageResponse) SetHeaders(v map[string]*string) *UpdateLakeStorageResponse {
	s.Headers = v
	return s
}

func (s *UpdateLakeStorageResponse) SetStatusCode(v int32) *UpdateLakeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateLakeStorageResponse) SetBody(v *UpdateLakeStorageResponseBody) *UpdateLakeStorageResponse {
	s.Body = v
	return s
}

func (s *UpdateLakeStorageResponse) Validate() error {
	return dara.Validate(s)
}
