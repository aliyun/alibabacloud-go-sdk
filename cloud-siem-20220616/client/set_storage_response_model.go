// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *SetStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *SetStorageResponse
	GetStatusCode() *int32
	SetBody(v *SetStorageResponseBody) *SetStorageResponse
	GetBody() *SetStorageResponseBody
}

type SetStorageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SetStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SetStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s SetStorageResponse) GoString() string {
	return s.String()
}

func (s *SetStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *SetStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *SetStorageResponse) GetBody() *SetStorageResponseBody {
	return s.Body
}

func (s *SetStorageResponse) SetHeaders(v map[string]*string) *SetStorageResponse {
	s.Headers = v
	return s
}

func (s *SetStorageResponse) SetStatusCode(v int32) *SetStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *SetStorageResponse) SetBody(v *SetStorageResponseBody) *SetStorageResponse {
	s.Body = v
	return s
}

func (s *SetStorageResponse) Validate() error {
	return dara.Validate(s)
}
