// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDNAJobListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDNAJobListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDNAJobListResponse
	GetStatusCode() *int32
	SetBody(v *QueryDNAJobListResponseBody) *QueryDNAJobListResponse
	GetBody() *QueryDNAJobListResponseBody
}

type QueryDNAJobListResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDNAJobListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDNAJobListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDNAJobListResponse) GoString() string {
	return s.String()
}

func (s *QueryDNAJobListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDNAJobListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDNAJobListResponse) GetBody() *QueryDNAJobListResponseBody {
	return s.Body
}

func (s *QueryDNAJobListResponse) SetHeaders(v map[string]*string) *QueryDNAJobListResponse {
	s.Headers = v
	return s
}

func (s *QueryDNAJobListResponse) SetStatusCode(v int32) *QueryDNAJobListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDNAJobListResponse) SetBody(v *QueryDNAJobListResponseBody) *QueryDNAJobListResponse {
	s.Body = v
	return s
}

func (s *QueryDNAJobListResponse) Validate() error {
	return dara.Validate(s)
}
