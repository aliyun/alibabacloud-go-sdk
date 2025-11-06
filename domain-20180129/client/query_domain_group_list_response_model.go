// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDomainGroupListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDomainGroupListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDomainGroupListResponse
	GetStatusCode() *int32
	SetBody(v *QueryDomainGroupListResponseBody) *QueryDomainGroupListResponse
	GetBody() *QueryDomainGroupListResponseBody
}

type QueryDomainGroupListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDomainGroupListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDomainGroupListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDomainGroupListResponse) GoString() string {
	return s.String()
}

func (s *QueryDomainGroupListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDomainGroupListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDomainGroupListResponse) GetBody() *QueryDomainGroupListResponseBody {
	return s.Body
}

func (s *QueryDomainGroupListResponse) SetHeaders(v map[string]*string) *QueryDomainGroupListResponse {
	s.Headers = v
	return s
}

func (s *QueryDomainGroupListResponse) SetStatusCode(v int32) *QueryDomainGroupListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDomainGroupListResponse) SetBody(v *QueryDomainGroupListResponseBody) *QueryDomainGroupListResponse {
	s.Body = v
	return s
}

func (s *QueryDomainGroupListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
