// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpaClusterImageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *GetOpaClusterImageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *GetOpaClusterImageListResponse
	GetStatusCode() *int32
	SetBody(v *GetOpaClusterImageListResponseBody) *GetOpaClusterImageListResponse
	GetBody() *GetOpaClusterImageListResponseBody
}

type GetOpaClusterImageListResponse struct {
	Headers    map[string]*string                  `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                              `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetOpaClusterImageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetOpaClusterImageListResponse) String() string {
	return dara.Prettify(s)
}

func (s GetOpaClusterImageListResponse) GoString() string {
	return s.String()
}

func (s *GetOpaClusterImageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *GetOpaClusterImageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *GetOpaClusterImageListResponse) GetBody() *GetOpaClusterImageListResponseBody {
	return s.Body
}

func (s *GetOpaClusterImageListResponse) SetHeaders(v map[string]*string) *GetOpaClusterImageListResponse {
	s.Headers = v
	return s
}

func (s *GetOpaClusterImageListResponse) SetStatusCode(v int32) *GetOpaClusterImageListResponse {
	s.StatusCode = &v
	return s
}

func (s *GetOpaClusterImageListResponse) SetBody(v *GetOpaClusterImageListResponseBody) *GetOpaClusterImageListResponse {
	s.Body = v
	return s
}

func (s *GetOpaClusterImageListResponse) Validate() error {
	return dara.Validate(s)
}
