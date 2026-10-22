// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitPenetrationTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SubmitPenetrationTestResponseBody
	GetCode() *string
	SetData(v *SubmitPenetrationTestResponseBodyData) *SubmitPenetrationTestResponseBody
	GetData() *SubmitPenetrationTestResponseBodyData
	SetMessage(v string) *SubmitPenetrationTestResponseBody
	GetMessage() *string
	SetRequestId(v string) *SubmitPenetrationTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SubmitPenetrationTestResponseBody
	GetSuccess() *bool
}

type SubmitPenetrationTestResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response data.
	Data *SubmitPenetrationTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 67BD8435-6624-5484-A75D-170231B51615
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The call status. Valid values:
	//
	// - **true**: Successful.
	//
	// - **false**: Failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SubmitPenetrationTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitPenetrationTestResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitPenetrationTestResponseBody) GetCode() *string {
	return s.Code
}

func (s *SubmitPenetrationTestResponseBody) GetData() *SubmitPenetrationTestResponseBodyData {
	return s.Data
}

func (s *SubmitPenetrationTestResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SubmitPenetrationTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitPenetrationTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SubmitPenetrationTestResponseBody) SetCode(v string) *SubmitPenetrationTestResponseBody {
	s.Code = &v
	return s
}

func (s *SubmitPenetrationTestResponseBody) SetData(v *SubmitPenetrationTestResponseBodyData) *SubmitPenetrationTestResponseBody {
	s.Data = v
	return s
}

func (s *SubmitPenetrationTestResponseBody) SetMessage(v string) *SubmitPenetrationTestResponseBody {
	s.Message = &v
	return s
}

func (s *SubmitPenetrationTestResponseBody) SetRequestId(v string) *SubmitPenetrationTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitPenetrationTestResponseBody) SetSuccess(v bool) *SubmitPenetrationTestResponseBody {
	s.Success = &v
	return s
}

func (s *SubmitPenetrationTestResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitPenetrationTestResponseBodyData struct {
	// The response data.
	List []*SubmitPenetrationTestResponseBodyDataList `json:"List,omitempty" xml:"List,omitempty" type:"Repeated"`
}

func (s SubmitPenetrationTestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitPenetrationTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitPenetrationTestResponseBodyData) GetList() []*SubmitPenetrationTestResponseBodyDataList {
	return s.List
}

func (s *SubmitPenetrationTestResponseBodyData) SetList(v []*SubmitPenetrationTestResponseBodyDataList) *SubmitPenetrationTestResponseBodyData {
	s.List = v
	return s
}

func (s *SubmitPenetrationTestResponseBodyData) Validate() error {
	if s.List != nil {
		for _, item := range s.List {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type SubmitPenetrationTestResponseBodyDataList struct {
	// The type of the asset. Valid values:
	//
	// - **EIP**: The public IP address of a virtual private cloud (VPC).
	//
	// - **SLB**: load balancing.
	//
	// - **ECS**: The public IP address of an Elastic Compute Service (ECS) instance in the classic network.
	//
	// example:
	//
	// EIP
	AssetType *string `json:"AssetType,omitempty" xml:"AssetType,omitempty"`
	// The IP address of the asset.
	//
	// example:
	//
	// 10.0.xxx.58
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
}

func (s SubmitPenetrationTestResponseBodyDataList) String() string {
	return dara.Prettify(s)
}

func (s SubmitPenetrationTestResponseBodyDataList) GoString() string {
	return s.String()
}

func (s *SubmitPenetrationTestResponseBodyDataList) GetAssetType() *string {
	return s.AssetType
}

func (s *SubmitPenetrationTestResponseBodyDataList) GetIp() *string {
	return s.Ip
}

func (s *SubmitPenetrationTestResponseBodyDataList) SetAssetType(v string) *SubmitPenetrationTestResponseBodyDataList {
	s.AssetType = &v
	return s
}

func (s *SubmitPenetrationTestResponseBodyDataList) SetIp(v string) *SubmitPenetrationTestResponseBodyDataList {
	s.Ip = &v
	return s
}

func (s *SubmitPenetrationTestResponseBodyDataList) Validate() error {
	return dara.Validate(s)
}
