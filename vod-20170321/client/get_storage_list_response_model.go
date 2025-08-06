// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageListResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageListResponseBody) *GetStorageListResponse
	GetBody() *GetStorageListResponseBody
}

type GetStorageListResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageListResponse) GoString() string {
	return s.String()
}

func (s *GetStorageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageListResponse) GetBody() *GetStorageListResponseBody {
	return s.Body
}

func (s *GetStorageListResponse) SetHeaders(v map[string]*string) *GetStorageListResponse {
	s.Headers = v
	return s
}

func (s *GetStorageListResponse) SetStatusCode(v int32) *GetStorageListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageListResponse) SetBody(v *GetStorageListResponseBody) *GetStorageListResponse {
	s.Body = v
	return s
}

func (s *GetStorageListResponse) Validate() error {
	return dara.Validate(s)
}
