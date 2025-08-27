// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBtripBillInfoAdjustResponse interface {
	dara.Model
	String() string
	GoString() string
	SetHeaders(v map[string]*string) *BtripBillInfoAdjustResponse
	GetHeaders() map[string]*string
	SetStatusCode(v int32) *BtripBillInfoAdjustResponse
	GetStatusCode() *int32
	SetBody(v *BtripBillInfoAdjustResponseBody) *BtripBillInfoAdjustResponse
	GetBody() *BtripBillInfoAdjustResponseBody
}

type BtripBillInfoAdjustResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BtripBillInfoAdjustResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BtripBillInfoAdjustResponse) String() string {
	return dara.Prettify(s)
}

func (s BtripBillInfoAdjustResponse) GoString() string {
	return s.String()
}

func (s *BtripBillInfoAdjustResponse) GetHeaders() map[string]*string {
	return s.Headers
}

func (s *BtripBillInfoAdjustResponse) GetStatusCode() *int32 {
	return s.StatusCode
}

func (s *BtripBillInfoAdjustResponse) GetBody() *BtripBillInfoAdjustResponseBody {
	return s.Body
}

func (s *BtripBillInfoAdjustResponse) SetHeaders(v map[string]*string) *BtripBillInfoAdjustResponse {
	s.Headers = v
	return s
}

func (s *BtripBillInfoAdjustResponse) SetStatusCode(v int32) *BtripBillInfoAdjustResponse {
	s.StatusCode = &v
	return s
}

func (s *BtripBillInfoAdjustResponse) SetBody(v *BtripBillInfoAdjustResponseBody) *BtripBillInfoAdjustResponse {
	s.Body = v
	return s
}

func (s *BtripBillInfoAdjustResponse) Validate() error {
	return dara.Validate(s)
}
