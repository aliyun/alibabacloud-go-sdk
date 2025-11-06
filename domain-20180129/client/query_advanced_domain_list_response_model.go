// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAdvancedDomainListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryAdvancedDomainListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryAdvancedDomainListResponse
	GetStatusCode() *int32
	SetBody(v *QueryAdvancedDomainListResponseBody) *QueryAdvancedDomainListResponse
	GetBody() *QueryAdvancedDomainListResponseBody
}

type QueryAdvancedDomainListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryAdvancedDomainListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryAdvancedDomainListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryAdvancedDomainListResponse) GoString() string {
	return s.String()
}

func (s *QueryAdvancedDomainListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryAdvancedDomainListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryAdvancedDomainListResponse) GetBody() *QueryAdvancedDomainListResponseBody {
	return s.Body
}

func (s *QueryAdvancedDomainListResponse) SetHeaders(v map[string]*string) *QueryAdvancedDomainListResponse {
	s.Headers = v
	return s
}

func (s *QueryAdvancedDomainListResponse) SetStatusCode(v int32) *QueryAdvancedDomainListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryAdvancedDomainListResponse) SetBody(v *QueryAdvancedDomainListResponseBody) *QueryAdvancedDomainListResponse {
	s.Body = v
	return s
}

func (s *QueryAdvancedDomainListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
