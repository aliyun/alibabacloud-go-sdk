// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyApiResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *ModifyApiResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *ModifyApiResponse
	GetStatusCode() *int32
	SetBody(v *ModifyApiResponseBody) *ModifyApiResponse
	GetBody() *ModifyApiResponseBody
}

type ModifyApiResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ModifyApiResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ModifyApiResponse) String() string {
	return dara.Prettify(s)
}

func (s ModifyApiResponse) GoString() string {
	return s.String()
}

func (s *ModifyApiResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *ModifyApiResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *ModifyApiResponse) GetBody() *ModifyApiResponseBody {
	return s.Body
}

func (s *ModifyApiResponse) SetHeaders(v map[string]*string) *ModifyApiResponse {
	s.Headers = v
	return s
}

func (s *ModifyApiResponse) SetStatusCode(v int32) *ModifyApiResponse {
	s.StatusCode = &v
	return s
}

func (s *ModifyApiResponse) SetBody(v *ModifyApiResponseBody) *ModifyApiResponse {
	s.Body = v
	return s
}

func (s *ModifyApiResponse) Validate() error {
	return dara.Validate(s)
}
