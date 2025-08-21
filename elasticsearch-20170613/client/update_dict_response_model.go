// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDictResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateDictResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateDictResponse
	GetStatusCode() *int32
	SetBody(v *UpdateDictResponseBody) *UpdateDictResponse
	GetBody() *UpdateDictResponseBody
}

type UpdateDictResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateDictResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDictResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateDictResponse) GoString() string {
	return s.String()
}

func (s *UpdateDictResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateDictResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateDictResponse) GetBody() *UpdateDictResponseBody {
	return s.Body
}

func (s *UpdateDictResponse) SetHeaders(v map[string]*string) *UpdateDictResponse {
	s.Headers = v
	return s
}

func (s *UpdateDictResponse) SetStatusCode(v int32) *UpdateDictResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDictResponse) SetBody(v *UpdateDictResponseBody) *UpdateDictResponse {
	s.Body = v
	return s
}

func (s *UpdateDictResponse) Validate() error {
	return dara.Validate(s)
}
