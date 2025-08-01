// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetBlackWhiteListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetBlackWhiteListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetBlackWhiteListResponse
	GetStatusCode() *int32
	SetBody(v *GetBlackWhiteListResponseBody) *GetBlackWhiteListResponse
	GetBody() *GetBlackWhiteListResponseBody
}

type GetBlackWhiteListResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetBlackWhiteListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetBlackWhiteListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetBlackWhiteListResponse) GoString() string {
	return s.String()
}

func (s *GetBlackWhiteListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetBlackWhiteListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetBlackWhiteListResponse) GetBody() *GetBlackWhiteListResponseBody {
	return s.Body
}

func (s *GetBlackWhiteListResponse) SetHeaders(v map[string]*string) *GetBlackWhiteListResponse {
	s.Headers = v
	return s
}

func (s *GetBlackWhiteListResponse) SetStatusCode(v int32) *GetBlackWhiteListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetBlackWhiteListResponse) SetBody(v *GetBlackWhiteListResponseBody) *GetBlackWhiteListResponse {
	s.Body = v
	return s
}

func (s *GetBlackWhiteListResponse) Validate() error {
	return dara.Validate(s)
}
