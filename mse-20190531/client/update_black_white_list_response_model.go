// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateBlackWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *UpdateBlackWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *UpdateBlackWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *UpdateBlackWhiteListResponseBody) *UpdateBlackWhiteListResponse
	GetBody() *UpdateBlackWhiteListResponseBody
}

type UpdateBlackWhiteListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateBlackWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateBlackWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s UpdateBlackWhiteListResponse) GoString() string {
	return s.String()
}

func (s *UpdateBlackWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *UpdateBlackWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *UpdateBlackWhiteListResponse) GetBody() *UpdateBlackWhiteListResponseBody {
	return s.Body
}

func (s *UpdateBlackWhiteListResponse) SetHeaders(v map[string]*string) *UpdateBlackWhiteListResponse {
	s.Headers = v
	return s
}

func (s *UpdateBlackWhiteListResponse) SetStatusCode(v int32) *UpdateBlackWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateBlackWhiteListResponse) SetBody(v *UpdateBlackWhiteListResponseBody) *UpdateBlackWhiteListResponse {
	s.Body = v
	return s
}

func (s *UpdateBlackWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
