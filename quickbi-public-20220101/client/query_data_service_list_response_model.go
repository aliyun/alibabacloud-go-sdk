// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryDataServiceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryDataServiceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryDataServiceListResponse
	GetStatusCode() *int32
	SetBody(v *QueryDataServiceListResponseBody) *QueryDataServiceListResponse
	GetBody() *QueryDataServiceListResponseBody
}

type QueryDataServiceListResponse struct {
	Headers    map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                            `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryDataServiceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryDataServiceListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryDataServiceListResponse) GoString() string {
	return s.String()
}

func (s *QueryDataServiceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryDataServiceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryDataServiceListResponse) GetBody() *QueryDataServiceListResponseBody {
	return s.Body
}

func (s *QueryDataServiceListResponse) SetHeaders(v map[string]*string) *QueryDataServiceListResponse {
	s.Headers = v
	return s
}

func (s *QueryDataServiceListResponse) SetStatusCode(v int32) *QueryDataServiceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryDataServiceListResponse) SetBody(v *QueryDataServiceListResponseBody) *QueryDataServiceListResponse {
	s.Body = v
	return s
}

func (s *QueryDataServiceListResponse) Validate() error {
	return dara.Validate(s)
}
