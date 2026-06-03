// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPnsPackageListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPnsPackageListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPnsPackageListResponse
	GetStatusCode() *int32
	SetBody(v *QueryPnsPackageListResponseBody) *QueryPnsPackageListResponse
	GetBody() *QueryPnsPackageListResponseBody
}

type QueryPnsPackageListResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPnsPackageListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPnsPackageListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPnsPackageListResponse) GoString() string {
	return s.String()
}

func (s *QueryPnsPackageListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPnsPackageListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPnsPackageListResponse) GetBody() *QueryPnsPackageListResponseBody {
	return s.Body
}

func (s *QueryPnsPackageListResponse) SetHeaders(v map[string]*string) *QueryPnsPackageListResponse {
	s.Headers = v
	return s
}

func (s *QueryPnsPackageListResponse) SetStatusCode(v int32) *QueryPnsPackageListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPnsPackageListResponse) SetBody(v *QueryPnsPackageListResponseBody) *QueryPnsPackageListResponse {
	s.Body = v
	return s
}

func (s *QueryPnsPackageListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
