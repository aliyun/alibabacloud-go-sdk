// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppInfosResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetAppInfosResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetAppInfosResponse
	GetStatusCode() *int32
	SetBody(v *GetAppInfosResponseBody) *GetAppInfosResponse
	GetBody() *GetAppInfosResponseBody
}

type GetAppInfosResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAppInfosResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAppInfosResponse) String() string {
	return dara.Prettify(s)
}

func (s GetAppInfosResponse) GoString() string {
	return s.String()
}

func (s *GetAppInfosResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetAppInfosResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetAppInfosResponse) GetBody() *GetAppInfosResponseBody {
	return s.Body
}

func (s *GetAppInfosResponse) SetHeaders(v map[string]*string) *GetAppInfosResponse {
	s.Headers = v
	return s
}

func (s *GetAppInfosResponse) SetStatusCode(v int32) *GetAppInfosResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAppInfosResponse) SetBody(v *GetAppInfosResponseBody) *GetAppInfosResponse {
	s.Body = v
	return s
}

func (s *GetAppInfosResponse) Validate() error {
	return dara.Validate(s)
}
