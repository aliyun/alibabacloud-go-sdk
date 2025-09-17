// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSellerInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeSellerInstancesResponseBody
	GetCode() *string
	SetCount(v int64) *DescribeSellerInstancesResponseBody
	GetCount() *int64
	SetFatal(v bool) *DescribeSellerInstancesResponseBody
	GetFatal() *bool
	SetMessage(v string) *DescribeSellerInstancesResponseBody
	GetMessage() *string
	SetPageNumber(v int64) *DescribeSellerInstancesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *DescribeSellerInstancesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *DescribeSellerInstancesResponseBody
	GetRequestId() *string
	SetResult(v []*DescribeSellerInstancesResponseBodyResult) *DescribeSellerInstancesResponseBody
	GetResult() []*DescribeSellerInstancesResponseBodyResult
	SetSuccess(v bool) *DescribeSellerInstancesResponseBody
	GetSuccess() *bool
	SetVersion(v string) *DescribeSellerInstancesResponseBody
	GetVersion() *string
}

type DescribeSellerInstancesResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 10
	Count *int64 `json:"Count,omitempty" xml:"Count,omitempty"`
	// fatal
	//
	// example:
	//
	// False
	Fatal *bool `json:"Fatal,omitempty" xml:"Fatal,omitempty"`
	// example:
	//
	// Instance 5723f7ee-952d-411f-94f4-b942a550d9b8 does not exist.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// example:
	//
	// 10
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// example:
	//
	// A6A33748-D573-593C-A3BC-593E33D68311
	RequestId *string                                      `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    []*DescribeSellerInstancesResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// example:
	//
	// True
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// 103
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DescribeSellerInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSellerInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSellerInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeSellerInstancesResponseBody) GetCount() *int64 {
	return s.Count
}

func (s *DescribeSellerInstancesResponseBody) GetFatal() *bool {
	return s.Fatal
}

func (s *DescribeSellerInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSellerInstancesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *DescribeSellerInstancesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *DescribeSellerInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSellerInstancesResponseBody) GetResult() []*DescribeSellerInstancesResponseBodyResult {
	return s.Result
}

func (s *DescribeSellerInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSellerInstancesResponseBody) GetVersion() *string {
	return s.Version
}

func (s *DescribeSellerInstancesResponseBody) SetCode(v string) *DescribeSellerInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetCount(v int64) *DescribeSellerInstancesResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetFatal(v bool) *DescribeSellerInstancesResponseBody {
	s.Fatal = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetMessage(v string) *DescribeSellerInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetPageNumber(v int64) *DescribeSellerInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetPageSize(v int64) *DescribeSellerInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetRequestId(v string) *DescribeSellerInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetResult(v []*DescribeSellerInstancesResponseBodyResult) *DescribeSellerInstancesResponseBody {
	s.Result = v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetSuccess(v bool) *DescribeSellerInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) SetVersion(v string) *DescribeSellerInstancesResponseBody {
	s.Version = &v
	return s
}

func (s *DescribeSellerInstancesResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeSellerInstancesResponseBodyResult struct {
	// example:
	//
	// {\\"authUrl\\":\\"https://marketplace.alibabacloud.com/\\"}
	AppInfo *string `json:"AppInfo,omitempty" xml:"AppInfo,omitempty"`
	// example:
	//
	// 1
	ChargeType *int32 `json:"ChargeType,omitempty" xml:"ChargeType,omitempty"`
	// example:
	//
	// sgcmgj000356
	CommodityCode *string `json:"CommodityCode,omitempty" xml:"CommodityCode,omitempty"`
	// example:
	//
	// 1741752000000
	CreatedOn *int64 `json:"CreatedOn,omitempty" xml:"CreatedOn,omitempty"`
	// example:
	//
	// {\\"userName\\":\\"marketplace\\"}
	HostInfo *string `json:"HostInfo,omitempty" xml:"HostInfo,omitempty"`
	// example:
	//
	// {\\"userName\\":\\"marketplace\\"}
	Info *string `json:"Info,omitempty" xml:"Info,omitempty"`
	// example:
	//
	// 5000002763123
	InstanceId *int64 `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// OPENED
	InstanceStatus *string `json:"InstanceStatus,omitempty" xml:"InstanceStatus,omitempty"`
	// example:
	//
	// 5322460655123456
	UserId *int64 `json:"UserId,omitempty" xml:"UserId,omitempty"`
}

func (s DescribeSellerInstancesResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeSellerInstancesResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeSellerInstancesResponseBodyResult) GetAppInfo() *string {
	return s.AppInfo
}

func (s *DescribeSellerInstancesResponseBodyResult) GetChargeType() *int32 {
	return s.ChargeType
}

func (s *DescribeSellerInstancesResponseBodyResult) GetCommodityCode() *string {
	return s.CommodityCode
}

func (s *DescribeSellerInstancesResponseBodyResult) GetCreatedOn() *int64 {
	return s.CreatedOn
}

func (s *DescribeSellerInstancesResponseBodyResult) GetHostInfo() *string {
	return s.HostInfo
}

func (s *DescribeSellerInstancesResponseBodyResult) GetInfo() *string {
	return s.Info
}

func (s *DescribeSellerInstancesResponseBodyResult) GetInstanceId() *int64 {
	return s.InstanceId
}

func (s *DescribeSellerInstancesResponseBodyResult) GetInstanceStatus() *string {
	return s.InstanceStatus
}

func (s *DescribeSellerInstancesResponseBodyResult) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeSellerInstancesResponseBodyResult) SetAppInfo(v string) *DescribeSellerInstancesResponseBodyResult {
	s.AppInfo = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetChargeType(v int32) *DescribeSellerInstancesResponseBodyResult {
	s.ChargeType = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetCommodityCode(v string) *DescribeSellerInstancesResponseBodyResult {
	s.CommodityCode = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetCreatedOn(v int64) *DescribeSellerInstancesResponseBodyResult {
	s.CreatedOn = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetHostInfo(v string) *DescribeSellerInstancesResponseBodyResult {
	s.HostInfo = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetInfo(v string) *DescribeSellerInstancesResponseBodyResult {
	s.Info = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetInstanceId(v int64) *DescribeSellerInstancesResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetInstanceStatus(v string) *DescribeSellerInstancesResponseBodyResult {
	s.InstanceStatus = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) SetUserId(v int64) *DescribeSellerInstancesResponseBodyResult {
	s.UserId = &v
	return s
}

func (s *DescribeSellerInstancesResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
