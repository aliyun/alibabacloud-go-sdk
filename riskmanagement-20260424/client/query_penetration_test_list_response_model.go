// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPenetrationTestListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPenetrationTestListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPenetrationTestListResponse
	GetStatusCode() *int32
	SetBody(v *QueryPenetrationTestListResponseBody) *QueryPenetrationTestListResponse
	GetBody() *QueryPenetrationTestListResponseBody
}

type QueryPenetrationTestListResponse struct {
	Headers    map[string]*string                    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPenetrationTestListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPenetrationTestListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPenetrationTestListResponse) GoString() string {
	return s.String()
}

func (s *QueryPenetrationTestListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPenetrationTestListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPenetrationTestListResponse) GetBody() *QueryPenetrationTestListResponseBody {
	return s.Body
}

func (s *QueryPenetrationTestListResponse) SetHeaders(v map[string]*string) *QueryPenetrationTestListResponse {
	s.Headers = v
	return s
}

func (s *QueryPenetrationTestListResponse) SetStatusCode(v int32) *QueryPenetrationTestListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPenetrationTestListResponse) SetBody(v *QueryPenetrationTestListResponseBody) *QueryPenetrationTestListResponse {
	s.Body = v
	return s
}

func (s *QueryPenetrationTestListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
