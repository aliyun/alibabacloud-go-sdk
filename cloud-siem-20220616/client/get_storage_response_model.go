// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetStorageResponse
	GetStatusCode() *int32
	SetBody(v *GetStorageResponseBody) *GetStorageResponse
	GetBody() *GetStorageResponseBody
}

type GetStorageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s GetStorageResponse) GoString() string {
	return s.String()
}

func (s *GetStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetStorageResponse) GetBody() *GetStorageResponseBody {
	return s.Body
}

func (s *GetStorageResponse) SetHeaders(v map[string]*string) *GetStorageResponse {
	s.Headers = v
	return s
}

func (s *GetStorageResponse) SetStatusCode(v int32) *GetStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStorageResponse) SetBody(v *GetStorageResponseBody) *GetStorageResponse {
	s.Body = v
	return s
}

func (s *GetStorageResponse) Validate() error {
	return dara.Validate(s)
}
