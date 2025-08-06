// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSdkListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetSdkListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetSdkListResponse
	GetStatusCode() *int32
	SetBody(v *GetSdkListResponseBody) *GetSdkListResponse
	GetBody() *GetSdkListResponseBody
}

type GetSdkListResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetSdkListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetSdkListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetSdkListResponse) GoString() string {
	return s.String()
}

func (s *GetSdkListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetSdkListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetSdkListResponse) GetBody() *GetSdkListResponseBody {
	return s.Body
}

func (s *GetSdkListResponse) SetHeaders(v map[string]*string) *GetSdkListResponse {
	s.Headers = v
	return s
}

func (s *GetSdkListResponse) SetStatusCode(v int32) *GetSdkListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetSdkListResponse) SetBody(v *GetSdkListResponseBody) *GetSdkListResponse {
	s.Body = v
	return s
}

func (s *GetSdkListResponse) Validate() error {
	return dara.Validate(s)
}
