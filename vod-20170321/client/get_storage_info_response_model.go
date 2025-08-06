// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageInfoResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageInfoResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageInfoResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageInfoResponseBody) *GetStorageInfoResponse
	GetBody() *GetStorageInfoResponseBody
}

type GetStorageInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageInfoResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageInfoResponse) GoString() string {
	return s.String()
}

func (s *GetStorageInfoResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageInfoResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageInfoResponse) GetBody() *GetStorageInfoResponseBody {
	return s.Body
}

func (s *GetStorageInfoResponse) SetHeaders(v map[string]*string) *GetStorageInfoResponse {
	s.Headers = v
	return s
}

func (s *GetStorageInfoResponse) SetStatusCode(v int32) *GetStorageInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageInfoResponse) SetBody(v *GetStorageInfoResponseBody) *GetStorageInfoResponse {
	s.Body = v
	return s
}

func (s *GetStorageInfoResponse) Validate() error {
	return dara.Validate(s)
}
