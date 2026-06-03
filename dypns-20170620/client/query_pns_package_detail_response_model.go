// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsPackageDetailResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPnsPackageDetailResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPnsPackageDetailResponse
	GetStatusCode() *int32
	SetBody(v *QueryPnsPackageDetailResponseBody) *QueryPnsPackageDetailResponse
	GetBody() *QueryPnsPackageDetailResponseBody
}

type QueryPnsPackageDetailResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPnsPackageDetailResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPnsPackageDetailResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsPackageDetailResponse) GoString() string {
	return s.String()
}

func (s *QueryPnsPackageDetailResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPnsPackageDetailResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPnsPackageDetailResponse) GetBody() *QueryPnsPackageDetailResponseBody {
	return s.Body
}

func (s *QueryPnsPackageDetailResponse) SetHeaders(v map[string]*string) *QueryPnsPackageDetailResponse {
	s.Headers = v
	return s
}

func (s *QueryPnsPackageDetailResponse) SetStatusCode(v int32) *QueryPnsPackageDetailResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPnsPackageDetailResponse) SetBody(v *QueryPnsPackageDetailResponseBody) *QueryPnsPackageDetailResponse {
	s.Body = v
	return s
}

func (s *QueryPnsPackageDetailResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
