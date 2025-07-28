// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeResponseCodeTrendGraphResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeResponseCodeTrendGraphResponseBody
	GetRequestId() *string
	SetResponseCodes(v []*DescribeResponseCodeTrendGraphResponseBodyResponseCodes) *DescribeResponseCodeTrendGraphResponseBody
	GetResponseCodes() []*DescribeResponseCodeTrendGraphResponseBodyResponseCodes
}

type DescribeResponseCodeTrendGraphResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 7D46493E-84DD-58CE-80A7-8643****9ECC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The array of the statistics of the error codes.
	ResponseCodes []*DescribeResponseCodeTrendGraphResponseBodyResponseCodes `json:"ResponseCodes,omitempty" xml:"ResponseCodes,omitempty" type:"Repeated"`
}

func (s DescribeResponseCodeTrendGraphResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeResponseCodeTrendGraphResponseBody) GetResponseCodes() []*DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	return s.ResponseCodes
}

func (s *DescribeResponseCodeTrendGraphResponseBody) SetRequestId(v string) *DescribeResponseCodeTrendGraphResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBody) SetResponseCodes(v []*DescribeResponseCodeTrendGraphResponseBodyResponseCodes) *DescribeResponseCodeTrendGraphResponseBody {
	s.ResponseCodes = v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeResponseCodeTrendGraphResponseBodyResponseCodes struct {
	// The number of 302 error codes that are returned.
	//
	// example:
	//
	// 0
	Code302Pv *int64 `json:"302Pv,omitempty" xml:"302Pv,omitempty"`
	// The number of 405 error codes that are returned.
	//
	// example:
	//
	// 121645464
	Code405Pv *int64 `json:"405Pv,omitempty" xml:"405Pv,omitempty"`
	// The number of 444 error codes that are returned.
	//
	// example:
	//
	// 0
	Code444Pv *int64 `json:"444Pv,omitempty" xml:"444Pv,omitempty"`
	// The number of 499 error codes that are returned.
	//
	// example:
	//
	// 0
	Code499Pv *int64 `json:"499Pv,omitempty" xml:"499Pv,omitempty"`
	// The number of 5xx error codes that are returned.
	//
	// example:
	//
	// 2932
	Code5xxPv *int64 `json:"5xxPv,omitempty" xml:"5xxPv,omitempty"`
	// The serial number of the time interval. The serial numbers are arranged in chronological order.
	//
	// example:
	//
	// 10
	Index *int64 `json:"Index,omitempty" xml:"Index,omitempty"`
}

func (s DescribeResponseCodeTrendGraphResponseBodyResponseCodes) String() string {
	return dara.Prettify(s)
}

func (s DescribeResponseCodeTrendGraphResponseBodyResponseCodes) GoString() string {
	return s.String()
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) GetCode302Pv() *int64 {
	return s.Code302Pv
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) GetCode405Pv() *int64 {
	return s.Code405Pv
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) GetCode444Pv() *int64 {
	return s.Code444Pv
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) GetCode499Pv() *int64 {
	return s.Code499Pv
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) GetCode5xxPv() *int64 {
	return s.Code5xxPv
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) GetIndex() *int64 {
	return s.Index
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode302Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code302Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode405Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code405Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode444Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code444Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode499Pv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code499Pv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetCode5xxPv(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Code5xxPv = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) SetIndex(v int64) *DescribeResponseCodeTrendGraphResponseBodyResponseCodes {
	s.Index = &v
	return s
}

func (s *DescribeResponseCodeTrendGraphResponseBodyResponseCodes) Validate() error {
	return dara.Validate(s)
}
