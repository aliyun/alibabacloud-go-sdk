// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSplitItemBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *DescribeSplitItemBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *DescribeSplitItemBillResponse
	GetStatusCode() *int32
	SetBody(v *DescribeSplitItemBillResponseBody) *DescribeSplitItemBillResponse
	GetBody() *DescribeSplitItemBillResponseBody
}

type DescribeSplitItemBillResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DescribeSplitItemBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DescribeSplitItemBillResponse) String() string {
	return dara.Prettify(s)
}

func (s DescribeSplitItemBillResponse) GoString() string {
	return s.String()
}

func (s *DescribeSplitItemBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *DescribeSplitItemBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *DescribeSplitItemBillResponse) GetBody() *DescribeSplitItemBillResponseBody {
	return s.Body
}

func (s *DescribeSplitItemBillResponse) SetHeaders(v map[string]*string) *DescribeSplitItemBillResponse {
	s.Headers = v
	return s
}

func (s *DescribeSplitItemBillResponse) SetStatusCode(v int32) *DescribeSplitItemBillResponse {
	s.StatusCode = &v
	return s
}

func (s *DescribeSplitItemBillResponse) SetBody(v *DescribeSplitItemBillResponseBody) *DescribeSplitItemBillResponse {
	s.Body = v
	return s
}

func (s *DescribeSplitItemBillResponse) Validate() error {
	return dara.Validate(s)
}
