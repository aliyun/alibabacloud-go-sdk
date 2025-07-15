// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppVersionsResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppVersionsResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppVersionsResponse
	GetStatusCode() *int32
	SetBody(v *GetAppVersionsResponseBody) *GetAppVersionsResponse
	GetBody() *GetAppVersionsResponseBody
}

type GetAppVersionsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppVersionsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppVersionsResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppVersionsResponse) GoString() string {
	return s.String()
}

func (s *GetAppVersionsResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppVersionsResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppVersionsResponse) GetBody() *GetAppVersionsResponseBody {
	return s.Body
}

func (s *GetAppVersionsResponse) SetHeaders(v map[string]*string) *GetAppVersionsResponse {
	s.Headers = v
	return s
}

func (s *GetAppVersionsResponse) SetStatusCode(v int32) *GetAppVersionsResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppVersionsResponse) SetBody(v *GetAppVersionsResponseBody) *GetAppVersionsResponse {
	s.Body = v
	return s
}

func (s *GetAppVersionsResponse) Validate() error {
	return dara.Validate(s)
}
