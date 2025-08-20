// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetLakeStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetLakeStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetLakeStorageResponse
	GetStatusCode() *int32
	SetBody(v *GetLakeStorageResponseBody) *GetLakeStorageResponse
	GetBody() *GetLakeStorageResponseBody
}

type GetLakeStorageResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLakeStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLakeStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetLakeStorageResponse) GoString() string {
	return s.String()
}

func (s *GetLakeStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetLakeStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetLakeStorageResponse) GetBody() *GetLakeStorageResponseBody {
	return s.Body
}

func (s *GetLakeStorageResponse) SetHeaders(v map[string]*string) *GetLakeStorageResponse {
	s.Headers = v
	return s
}

func (s *GetLakeStorageResponse) SetStatusCode(v int32) *GetLakeStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLakeStorageResponse) SetBody(v *GetLakeStorageResponseBody) *GetLakeStorageResponse {
	s.Body = v
	return s
}

func (s *GetLakeStorageResponse) Validate() error {
	return dara.Validate(s)
}
