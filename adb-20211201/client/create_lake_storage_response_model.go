// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateLakeStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *CreateLakeStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *CreateLakeStorageResponse
	GetStatusCode() *int32
	SetBody(v *CreateLakeStorageResponseBody) *CreateLakeStorageResponse
	GetBody() *CreateLakeStorageResponseBody
}

type CreateLakeStorageResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateLakeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateLakeStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s CreateLakeStorageResponse) GoString() string {
	return s.String()
}

func (s *CreateLakeStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *CreateLakeStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *CreateLakeStorageResponse) GetBody() *CreateLakeStorageResponseBody {
	return s.Body
}

func (s *CreateLakeStorageResponse) SetHeaders(v map[string]*string) *CreateLakeStorageResponse {
	s.Headers = v
	return s
}

func (s *CreateLakeStorageResponse) SetStatusCode(v int32) *CreateLakeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateLakeStorageResponse) SetBody(v *CreateLakeStorageResponseBody) *CreateLakeStorageResponse {
	s.Body = v
	return s
}

func (s *CreateLakeStorageResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
