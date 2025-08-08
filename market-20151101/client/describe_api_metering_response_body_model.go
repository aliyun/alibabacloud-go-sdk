// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeApiMeteringResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeApiMeteringResponseBody
	GetCode() *string
	SetCount(v int64) *DescribeApiMeteringResponseBody
	GetCount() *int64
	SetFatal(v bool) *DescribeApiMeteringResponseBody
	GetFatal() *bool
	SetMessage(v string) *DescribeApiMeteringResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *DescribeApiMeteringResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeApiMeteringResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeApiMeteringResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeApiMeteringResponseBodyResult) *DescribeApiMeteringResponseBody
	GetResult() []*DescribeApiMeteringResponseBodyResult
	SetSuccess(v bool) *DescribeApiMeteringResponseBody
	GetSuccess() *bool
	SetVersion(v string) *DescribeApiMeteringResponseBody
	GetVersion() *string
}

type DescribeApiMeteringResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 100
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// fatal
	//
	// example:
	//
	// false
	Fatal   *bool   `json:"Fatal,omitempty" xml:"Fatal,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// 62FC432C55A1A63534A6CB34
	RequestId *string                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeApiMeteringResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeApiMeteringResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiMeteringResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeApiMeteringResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeApiMeteringResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *DescribeApiMeteringResponseBody) GetFatal() *bool {
	return s.Fatal
}

func (s *DescribeApiMeteringResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeApiMeteringResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeApiMeteringResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeApiMeteringResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeApiMeteringResponseBody) GetResult() []*DescribeApiMeteringResponseBodyResult {
	return s.Result
}

func (s *DescribeApiMeteringResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeApiMeteringResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeApiMeteringResponseBody) SetCode(v string) *DescribeApiMeteringResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetCount(v int64) *DescribeApiMeteringResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetFatal(v bool) *DescribeApiMeteringResponseBody {
	s.Fatal = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetMessage(v string) *DescribeApiMeteringResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetPageNumber(v int64) *DescribeApiMeteringResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetPageSize(v int64) *DescribeApiMeteringResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetRequestId(v string) *DescribeApiMeteringResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetResult(v []*DescribeApiMeteringResponseBodyResult) *DescribeApiMeteringResponseBody {
	s.Result = v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetSuccess(v bool) *DescribeApiMeteringResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) SetVersion(v string) *DescribeApiMeteringResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeApiMeteringResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeApiMeteringResponseBodyResult struct {
	// example:
	//
	// 102277855749****
	AliyunPk *int64 `json:"AliyunPk,omitempty" xml:"AliyunPk,omitempty"`
	// example:
	//
	// cmapi0004****
	ProductCode   *string `json:"ProductCode,omitempty" xml:"ProductCode,omitempty"`
	ProductName   *string `json:"ProductName,omitempty" xml:"ProductName,omitempty"`
	TotalCapacity *int64  `json:"TotalCapacity,omitempty" xml:"TotalCapacity,omitempty"`
	// example:
	//
	// 98
	TotalQuota *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	// example:
	//
	// 220
	TotalUsage *int64  `json:"TotalUsage,omitempty" xml:"TotalUsage,omitempty"`
	Unit       *string `json:"Unit,omitempty" xml:"Unit,omitempty"`
}

func (s DescribeApiMeteringResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeApiMeteringResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeApiMeteringResponseBodyResult) GetAliyunPk() *int64 {
	return s.AliyunPk
}

func (s *DescribeApiMeteringResponseBodyResult) GetProductCode() *string {
	return s.ProductCode
}

func (s *DescribeApiMeteringResponseBodyResult) GetProductName() *string {
	return s.ProductName
}

func (s *DescribeApiMeteringResponseBodyResult) GetTotalCapacity() *int64 {
	return s.TotalCapacity
}

func (s *DescribeApiMeteringResponseBodyResult) GetTotalQuota() *int64 {
	return s.TotalQuota
}

func (s *DescribeApiMeteringResponseBodyResult) GetTotalUsage() *int64 {
	return s.TotalUsage
}

func (s *DescribeApiMeteringResponseBodyResult) GetUnit() *string {
	return s.Unit
}

func (s *DescribeApiMeteringResponseBodyResult) SetAliyunPk(v int64) *DescribeApiMeteringResponseBodyResult {
	s.AliyunPk = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetProductCode(v string) *DescribeApiMeteringResponseBodyResult {
	s.ProductCode = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetProductName(v string) *DescribeApiMeteringResponseBodyResult {
	s.ProductName = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetTotalCapacity(v int64) *DescribeApiMeteringResponseBodyResult {
	s.TotalCapacity = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetTotalQuota(v int64) *DescribeApiMeteringResponseBodyResult {
	s.TotalQuota = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetTotalUsage(v int64) *DescribeApiMeteringResponseBodyResult {
	s.TotalUsage = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) SetUnit(v string) *DescribeApiMeteringResponseBodyResult {
	s.Unit = &v
	return s
}

func (s *DescribeApiMeteringResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
