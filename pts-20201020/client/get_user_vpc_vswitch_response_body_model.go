// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetUserVpcVSwitchResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetUserVpcVSwitchResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *GetUserVpcVSwitchResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetUserVpcVSwitchResponseBody
	GetMessage() *string
	SetPageNumber(v int32) *GetUserVpcVSwitchResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *GetUserVpcVSwitchResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *GetUserVpcVSwitchResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetUserVpcVSwitchResponseBody
	GetSuccess() *bool
	SetVSwitchCount(v int32) *GetUserVpcVSwitchResponseBody
	GetVSwitchCount() *int32
	SetVSwitchList(v []*GetUserVpcVSwitchResponseBodyVSwitchList) *GetUserVpcVSwitchResponseBody
	GetVSwitchList() []*GetUserVpcVSwitchResponseBodyVSwitchList
}

type GetUserVpcVSwitchResponseBody struct {
	// The system status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The returned message. If the request was successful, this parameter is left empty.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of returned entries per page.
	//
	// example:
	//
	// 5
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0235E5FC-4C7C-5F0C-843C-FC674F15F947
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
	// The number of vSwitches.
	//
	// example:
	//
	// 6
	VSwitchCount *int32 `json:"VSwitchCount,omitempty" xml:"VSwitchCount,omitempty"`
	// The vSwitches.
	VSwitchList []*GetUserVpcVSwitchResponseBodyVSwitchList `json:"VSwitchList,omitempty" xml:"VSwitchList,omitempty" type:"Repeated"`
}

func (s GetUserVpcVSwitchResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcVSwitchResponseBody) GoString() string {
	return s.String()
}

func (s *GetUserVpcVSwitchResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetUserVpcVSwitchResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetUserVpcVSwitchResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetUserVpcVSwitchResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *GetUserVpcVSwitchResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetUserVpcVSwitchResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetUserVpcVSwitchResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetUserVpcVSwitchResponseBody) GetVSwitchCount() *int32 {
	return s.VSwitchCount
}

func (s *GetUserVpcVSwitchResponseBody) GetVSwitchList() []*GetUserVpcVSwitchResponseBodyVSwitchList {
	return s.VSwitchList
}

func (s *GetUserVpcVSwitchResponseBody) SetCode(v string) *GetUserVpcVSwitchResponseBody {
	s.Code = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBody) SetHttpStatusCode(v int32) *GetUserVpcVSwitchResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBody) SetMessage(v string) *GetUserVpcVSwitchResponseBody {
	s.Message = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBody) SetPageNumber(v int32) *GetUserVpcVSwitchResponseBody {
	s.PageNumber = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBody) SetPageSize(v int32) *GetUserVpcVSwitchResponseBody {
	s.PageSize = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBody) SetRequestId(v string) *GetUserVpcVSwitchResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBody) SetSuccess(v bool) *GetUserVpcVSwitchResponseBody {
	s.Success = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBody) SetVSwitchCount(v int32) *GetUserVpcVSwitchResponseBody {
	s.VSwitchCount = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBody) SetVSwitchList(v []*GetUserVpcVSwitchResponseBodyVSwitchList) *GetUserVpcVSwitchResponseBody {
	s.VSwitchList = v
	return s
}

func (s *GetUserVpcVSwitchResponseBody) Validate() error {
	if s.VSwitchList != nil {
		for _, item := range s.VSwitchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetUserVpcVSwitchResponseBodyVSwitchList struct {
	// The number of available IP addresses in the vSwitch.
	//
	// example:
	//
	// 1000
	AvailableIpAddressCount *int64 `json:"AvailableIpAddressCount,omitempty" xml:"AvailableIpAddressCount,omitempty"`
	// The maximum number of stress testers to be added.
	//
	// example:
	//
	// 1000
	MaxAgentCount *int32 `json:"MaxAgentCount,omitempty" xml:"MaxAgentCount,omitempty"`
	// The vSwitch ID.
	//
	// example:
	//
	// vsw-bp1eil9df23rsd8l1sevebiszooj
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// The vSwitch name.
	//
	// example:
	//
	// my-vswitch
	VSwitchName *string `json:"VSwitchName,omitempty" xml:"VSwitchName,omitempty"`
	// The VPC ID.
	//
	// example:
	//
	// vpc-wz9bpdaebft6j23fesdf84v2f1um3a
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s GetUserVpcVSwitchResponseBodyVSwitchList) String() string {
	return dara.Prettify(s)
}

func (s GetUserVpcVSwitchResponseBodyVSwitchList) GoString() string {
	return s.String()
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) GetAvailableIpAddressCount() *int64 {
	return s.AvailableIpAddressCount
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) GetMaxAgentCount() *int32 {
	return s.MaxAgentCount
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) GetVSwitchName() *string {
	return s.VSwitchName
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) GetVpcId() *string {
	return s.VpcId
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) SetAvailableIpAddressCount(v int64) *GetUserVpcVSwitchResponseBodyVSwitchList {
	s.AvailableIpAddressCount = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) SetMaxAgentCount(v int32) *GetUserVpcVSwitchResponseBodyVSwitchList {
	s.MaxAgentCount = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) SetVSwitchId(v string) *GetUserVpcVSwitchResponseBodyVSwitchList {
	s.VSwitchId = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) SetVSwitchName(v string) *GetUserVpcVSwitchResponseBodyVSwitchList {
	s.VSwitchName = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) SetVpcId(v string) *GetUserVpcVSwitchResponseBodyVSwitchList {
	s.VpcId = &v
	return s
}

func (s *GetUserVpcVSwitchResponseBodyVSwitchList) Validate() error {
	return dara.Validate(s)
}
