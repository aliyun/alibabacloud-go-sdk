// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServerlessAppAuthDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetServerlessAppAuthDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetServerlessAppAuthDetailResponse
	GetStatusCode() *int32
	SetBody(v *GetServerlessAppAuthDetailResponseBody) *GetServerlessAppAuthDetailResponse
	GetBody() *GetServerlessAppAuthDetailResponseBody
}

type GetServerlessAppAuthDetailResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetServerlessAppAuthDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetServerlessAppAuthDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s GetServerlessAppAuthDetailResponse) GoString() string {
	return s.String()
}

func (s *GetServerlessAppAuthDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetServerlessAppAuthDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetServerlessAppAuthDetailResponse) GetBody() *GetServerlessAppAuthDetailResponseBody {
	return s.Body
}

func (s *GetServerlessAppAuthDetailResponse) SetHeaders(v map[string]*string) *GetServerlessAppAuthDetailResponse {
	s.Headers = v
	return s
}

func (s *GetServerlessAppAuthDetailResponse) SetStatusCode(v int32) *GetServerlessAppAuthDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *GetServerlessAppAuthDetailResponse) SetBody(v *GetServerlessAppAuthDetailResponseBody) *GetServerlessAppAuthDetailResponse {
	s.Body = v
	return s
}

func (s *GetServerlessAppAuthDetailResponse) Validate() error {
	return dara.Validate(s)
}
