// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySharesToUserListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySharesToUserListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySharesToUserListResponse
	GetStatusCode() *int32
	SetBody(v *QuerySharesToUserListResponseBody) *QuerySharesToUserListResponse
	GetBody() *QuerySharesToUserListResponseBody
}

type QuerySharesToUserListResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySharesToUserListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySharesToUserListResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySharesToUserListResponse) GoString() string {
	return s.String()
}

func (s *QuerySharesToUserListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySharesToUserListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySharesToUserListResponse) GetBody() *QuerySharesToUserListResponseBody {
	return s.Body
}

func (s *QuerySharesToUserListResponse) SetHeaders(v map[string]*string) *QuerySharesToUserListResponse {
	s.Headers = v
	return s
}

func (s *QuerySharesToUserListResponse) SetStatusCode(v int32) *QuerySharesToUserListResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySharesToUserListResponse) SetBody(v *QuerySharesToUserListResponseBody) *QuerySharesToUserListResponse {
	s.Body = v
	return s
}

func (s *QuerySharesToUserListResponse) Validate() error {
	return dara.Validate(s)
}
