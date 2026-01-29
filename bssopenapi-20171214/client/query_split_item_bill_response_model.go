// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQuerySplitItemBillResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *QuerySplitItemBillResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *QuerySplitItemBillResponse
	GetStatusCode() *int32
	SetBody(v *QuerySplitItemBillResponseBody) *QuerySplitItemBillResponse
	GetBody() *QuerySplitItemBillResponseBody
}

type QuerySplitItemBillResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QuerySplitItemBillResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QuerySplitItemBillResponse) String() string {
	return dara.Prettify(s)
}

func (s QuerySplitItemBillResponse) GoString() string {
	return s.String()
}

func (s *QuerySplitItemBillResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *QuerySplitItemBillResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *QuerySplitItemBillResponse) GetBody() *QuerySplitItemBillResponseBody {
	return s.Body
}

func (s *QuerySplitItemBillResponse) SetHeaders(v map[string]*string) *QuerySplitItemBillResponse {
	s.Headers = v
	return s
}

func (s *QuerySplitItemBillResponse) SetStatusCode(v int32) *QuerySplitItemBillResponse {
	s.StatusCode = &v
	return s
}

func (s *QuerySplitItemBillResponse) SetBody(v *QuerySplitItemBillResponseBody) *QuerySplitItemBillResponse {
	s.Body = v
	return s
}

func (s *QuerySplitItemBillResponse) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	return nil
}
