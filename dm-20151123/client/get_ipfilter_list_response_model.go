// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIpfilterListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetIpfilterListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetIpfilterListResponse
	GetStatusCode() *int32
	SetBody(v *GetIpfilterListResponseBody) *GetIpfilterListResponse
	GetBody() *GetIpfilterListResponseBody
}

type GetIpfilterListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetIpfilterListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIpfilterListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetIpfilterListResponse) GoString() string {
	return s.String()
}

func (s *GetIpfilterListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetIpfilterListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetIpfilterListResponse) GetBody() *GetIpfilterListResponseBody {
	return s.Body
}

func (s *GetIpfilterListResponse) SetHeaders(v map[string]*string) *GetIpfilterListResponse {
	s.Headers = v
	return s
}

func (s *GetIpfilterListResponse) SetStatusCode(v int32) *GetIpfilterListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIpfilterListResponse) SetBody(v *GetIpfilterListResponseBody) *GetIpfilterListResponse {
	s.Body = v
	return s
}

func (s *GetIpfilterListResponse) Validate() error {
	return dara.Validate(s)
}
