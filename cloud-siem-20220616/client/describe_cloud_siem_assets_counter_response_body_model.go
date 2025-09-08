// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudSiemAssetsCounterResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *DescribeCloudSiemAssetsCounterResponseBody
	GetCode() *int32
	SetData(v []*DescribeCloudSiemAssetsCounterResponseBodyData) *DescribeCloudSiemAssetsCounterResponseBody
	GetData() []*DescribeCloudSiemAssetsCounterResponseBodyData
	SetMessage(v string) *DescribeCloudSiemAssetsCounterResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCloudSiemAssetsCounterResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeCloudSiemAssetsCounterResponseBody
	GetSuccess() *bool
}

type DescribeCloudSiemAssetsCounterResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	//
	// example:
	//
	// 123456
	Data []*DescribeCloudSiemAssetsCounterResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9AAA9ED9-78F4-5021-86DC-D51C7511****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) GetData() []*DescribeCloudSiemAssetsCounterResponseBodyData {
	return s.Data
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetCode(v int32) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetData(v []*DescribeCloudSiemAssetsCounterResponseBodyData) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetMessage(v string) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetRequestId(v string) *DescribeCloudSiemAssetsCounterResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) SetSuccess(v bool) *DescribeCloudSiemAssetsCounterResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeCloudSiemAssetsCounterResponseBodyData struct {
	// The number of assets.
	//
	// example:
	//
	// 1
	AssetNum *int32 `json:"AssetNum,omitempty" xml:"AssetNum,omitempty"`
	// The type of the asset. Valid values:
	//
	// 	- ip
	//
	// 	- domain
	//
	// 	- url
	//
	// 	- process
	//
	// 	- file
	//
	// 	- host
	//
	// example:
	//
	// domain
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
}

func (s DescribeCloudSiemAssetsCounterResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudSiemAssetsCounterResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCloudSiemAssetsCounterResponseBodyData) GetAssetNum() *int32 {
	return s.AssetNum
}

func (s *DescribeCloudSiemAssetsCounterResponseBodyData) GetAssetType() *string {
	return s.AssetType
}

func (s *DescribeCloudSiemAssetsCounterResponseBodyData) SetAssetNum(v int32) *DescribeCloudSiemAssetsCounterResponseBodyData {
	s.AssetNum = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBodyData) SetAssetType(v string) *DescribeCloudSiemAssetsCounterResponseBodyData {
	s.AssetType = &v
	return s
}

func (s *DescribeCloudSiemAssetsCounterResponseBodyData) Validate() error {
	return dara.Validate(s)
}
