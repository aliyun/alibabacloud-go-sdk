// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddBlackWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *AddBlackWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *AddBlackWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *AddBlackWhiteListResponseBody) *AddBlackWhiteListResponse
	GetBody() *AddBlackWhiteListResponseBody
}

type AddBlackWhiteListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddBlackWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddBlackWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s AddBlackWhiteListResponse) GoString() string {
	return s.String()
}

func (s *AddBlackWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *AddBlackWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *AddBlackWhiteListResponse) GetBody() *AddBlackWhiteListResponseBody {
	return s.Body
}

func (s *AddBlackWhiteListResponse) SetHeaders(v map[string]*string) *AddBlackWhiteListResponse {
	s.Headers = v
	return s
}

func (s *AddBlackWhiteListResponse) SetStatusCode(v int32) *AddBlackWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *AddBlackWhiteListResponse) SetBody(v *AddBlackWhiteListResponseBody) *AddBlackWhiteListResponse {
	s.Body = v
	return s
}

func (s *AddBlackWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
