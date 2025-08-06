// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPartnerProduceListResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryPartnerProduceListResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryPartnerProduceListResponse
	GetStatusCode() *int32
	SetBody(v *QueryPartnerProduceListResponseBody) *QueryPartnerProduceListResponse
	GetBody() *QueryPartnerProduceListResponseBody
}

type QueryPartnerProduceListResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryPartnerProduceListResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryPartnerProduceListResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryPartnerProduceListResponse) GoString() string {
	return s.String()
}

func (s *QueryPartnerProduceListResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryPartnerProduceListResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryPartnerProduceListResponse) GetBody() *QueryPartnerProduceListResponseBody {
	return s.Body
}

func (s *QueryPartnerProduceListResponse) SetHeaders(v map[string]*string) *QueryPartnerProduceListResponse {
	s.Headers = v
	return s
}

func (s *QueryPartnerProduceListResponse) SetStatusCode(v int32) *QueryPartnerProduceListResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryPartnerProduceListResponse) SetBody(v *QueryPartnerProduceListResponseBody) *QueryPartnerProduceListResponse {
	s.Body = v
	return s
}

func (s *QueryPartnerProduceListResponse) Validate() error {
	return dara.Validate(s)
}
