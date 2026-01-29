// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryCommodityListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryCommodityListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryCommodityListResponse
	GetStatusCode() *int32
	SetBody(v *QueryCommodityListResponseBody) *QueryCommodityListResponse
	GetBody() *QueryCommodityListResponseBody
}

type QueryCommodityListResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryCommodityListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryCommodityListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryCommodityListResponse) GoString() string {
	return s.String()
}

func (s *QueryCommodityListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryCommodityListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryCommodityListResponse) GetBody() *QueryCommodityListResponseBody {
	return s.Body
}

func (s *QueryCommodityListResponse) SetHeaders(v map[string]*string) *QueryCommodityListResponse {
	s.Headers = v
	return s
}

func (s *QueryCommodityListResponse) SetStatusCode(v int32) *QueryCommodityListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryCommodityListResponse) SetBody(v *QueryCommodityListResponseBody) *QueryCommodityListResponse {
	s.Body = v
	return s
}

func (s *QueryCommodityListResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
