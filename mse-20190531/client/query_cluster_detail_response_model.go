// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryClusterDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryClusterDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryClusterDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryClusterDetailResponseBody) *QueryClusterDetailResponse
	GetBody() *QueryClusterDetailResponseBody
}

type QueryClusterDetailResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryClusterDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryClusterDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryClusterDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryClusterDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryClusterDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryClusterDetailResponse) GetBody() *QueryClusterDetailResponseBody {
	return s.Body
}

func (s *QueryClusterDetailResponse) SetHeaders(v map[string]*string) *QueryClusterDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryClusterDetailResponse) SetStatusCode(v int32) *QueryClusterDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryClusterDetailResponse) SetBody(v *QueryClusterDetailResponseBody) *QueryClusterDetailResponse {
	s.Body = v
	return s
}

func (s *QueryClusterDetailResponse) Validate() error {
	return dara.Validate(s)
}
