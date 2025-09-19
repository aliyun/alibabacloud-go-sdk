// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTagWithUuidResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddTagWithUuidResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddTagWithUuidResponse
	GetStatusCode() *int32
	SetBody(v *AddTagWithUuidResponseBody) *AddTagWithUuidResponse
	GetBody() *AddTagWithUuidResponseBody
}

type AddTagWithUuidResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddTagWithUuidResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddTagWithUuidResponse) String() string {
	return dara.Prettify(s)
}

func (s AddTagWithUuidResponse) GoString() string {
	return s.String()
}

func (s *AddTagWithUuidResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddTagWithUuidResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddTagWithUuidResponse) GetBody() *AddTagWithUuidResponseBody {
	return s.Body
}

func (s *AddTagWithUuidResponse) SetHeaders(v map[string]*string) *AddTagWithUuidResponse {
	s.Headers = v
	return s
}

func (s *AddTagWithUuidResponse) SetStatusCode(v int32) *AddTagWithUuidResponse {
	s.StatusCode = &v
	return s
}

func (s *AddTagWithUuidResponse) SetBody(v *AddTagWithUuidResponseBody) *AddTagWithUuidResponse {
	s.Body = v
	return s
}

func (s *AddTagWithUuidResponse) Validate() error {
	return dara.Validate(s)
}
