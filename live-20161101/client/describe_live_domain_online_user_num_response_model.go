// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveDomainOnlineUserNumResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeLiveDomainOnlineUserNumResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeLiveDomainOnlineUserNumResponse
	GetStatusCode() *int32
	SetBody(v *DescribeLiveDomainOnlineUserNumResponseBody) *DescribeLiveDomainOnlineUserNumResponse
	GetBody() *DescribeLiveDomainOnlineUserNumResponseBody
}

type DescribeLiveDomainOnlineUserNumResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeLiveDomainOnlineUserNumResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeLiveDomainOnlineUserNumResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveDomainOnlineUserNumResponse) GoString() string {
	return s.String()
}

func (s *DescribeLiveDomainOnlineUserNumResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeLiveDomainOnlineUserNumResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeLiveDomainOnlineUserNumResponse) GetBody() *DescribeLiveDomainOnlineUserNumResponseBody {
	return s.Body
}

func (s *DescribeLiveDomainOnlineUserNumResponse) SetHeaders(v map[string]*string) *DescribeLiveDomainOnlineUserNumResponse {
	s.Headers = v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponse) SetStatusCode(v int32) *DescribeLiveDomainOnlineUserNumResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponse) SetBody(v *DescribeLiveDomainOnlineUserNumResponseBody) *DescribeLiveDomainOnlineUserNumResponse {
	s.Body = v
	return s
}

func (s *DescribeLiveDomainOnlineUserNumResponse) Validate() error {
	return dara.Validate(s)
}
