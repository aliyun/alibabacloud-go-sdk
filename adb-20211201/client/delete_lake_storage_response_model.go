// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLakeStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteLakeStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteLakeStorageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteLakeStorageResponseBody) *DeleteLakeStorageResponse
	GetBody() *DeleteLakeStorageResponseBody
}

type DeleteLakeStorageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteLakeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteLakeStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteLakeStorageResponse) GoString() string {
	return s.String()
}

func (s *DeleteLakeStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteLakeStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteLakeStorageResponse) GetBody() *DeleteLakeStorageResponseBody {
	return s.Body
}

func (s *DeleteLakeStorageResponse) SetHeaders(v map[string]*string) *DeleteLakeStorageResponse {
	s.Headers = v
	return s
}

func (s *DeleteLakeStorageResponse) SetStatusCode(v int32) *DeleteLakeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteLakeStorageResponse) SetBody(v *DeleteLakeStorageResponseBody) *DeleteLakeStorageResponse {
	s.Body = v
	return s
}

func (s *DeleteLakeStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
