// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryBillOverviewResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QueryBillOverviewResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QueryBillOverviewResponse
	GetStatusCode() *int32
	SetBody(v *QueryBillOverviewResponseBody) *QueryBillOverviewResponse
	GetBody() *QueryBillOverviewResponseBody
}

type QueryBillOverviewResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryBillOverviewResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryBillOverviewResponse) String() string {
	return dara.Prettify(s)
}

func (s QueryBillOverviewResponse) GoString() string {
	return s.String()
}

func (s *QueryBillOverviewResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QueryBillOverviewResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QueryBillOverviewResponse) GetBody() *QueryBillOverviewResponseBody {
	return s.Body
}

func (s *QueryBillOverviewResponse) SetHeaders(v map[string]*string) *QueryBillOverviewResponse {
	s.Headers = v
	return s
}

func (s *QueryBillOverviewResponse) SetStatusCode(v int32) *QueryBillOverviewResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryBillOverviewResponse) SetBody(v *QueryBillOverviewResponseBody) *QueryBillOverviewResponse {
	s.Body = v
	return s
}

func (s *QueryBillOverviewResponse) Validate() error {
	return dara.Validate(s)
}
