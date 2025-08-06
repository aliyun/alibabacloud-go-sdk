// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DeleteStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DeleteStorageResponse
	GetStatusCode() *int32
	SetBody(v *DeleteStorageResponseBody) *DeleteStorageResponse
	GetBody() *DeleteStorageResponseBody
}

type DeleteStorageResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s DeleteStorageResponse) GoString() string {
	return s.String()
}

func (s *DeleteStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DeleteStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DeleteStorageResponse) GetBody() *DeleteStorageResponseBody {
	return s.Body
}

func (s *DeleteStorageResponse) SetHeaders(v map[string]*string) *DeleteStorageResponse {
	s.Headers = v
	return s
}

func (s *DeleteStorageResponse) SetStatusCode(v int32) *DeleteStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStorageResponse) SetBody(v *DeleteStorageResponseBody) *DeleteStorageResponse {
	s.Body = v
	return s
}

func (s *DeleteStorageResponse) Validate() error {
	return dara.Validate(s)
}
