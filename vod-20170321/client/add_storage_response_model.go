// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddStorageResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddStorageResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddStorageResponse
	GetStatusCode() *int32
	SetBody(v *AddStorageResponseBody) *AddStorageResponse
	GetBody() *AddStorageResponseBody
}

type AddStorageResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddStorageResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddStorageResponse) String() string {
	return dara.Prettify(s)
}

func (s AddStorageResponse) GoString() string {
	return s.String()
}

func (s *AddStorageResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddStorageResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddStorageResponse) GetBody() *AddStorageResponseBody {
	return s.Body
}

func (s *AddStorageResponse) SetHeaders(v map[string]*string) *AddStorageResponse {
	s.Headers = v
	return s
}

func (s *AddStorageResponse) SetStatusCode(v int32) *AddStorageResponse {
	s.StatusCode = &v
	return s
}

func (s *AddStorageResponse) SetBody(v *AddStorageResponseBody) *AddStorageResponse {
	s.Body = v
	return s
}

func (s *AddStorageResponse) Validate() error {
	return dara.Validate(s)
}
