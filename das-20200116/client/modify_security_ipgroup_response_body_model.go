// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifySecurityIPGroupResponseBody
	GetCode() *string
	SetData(v *ModifySecurityIPGroupResponseBodyData) *ModifySecurityIPGroupResponseBody
	GetData() *ModifySecurityIPGroupResponseBodyData
	SetMessage(v string) *ModifySecurityIPGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifySecurityIPGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifySecurityIPGroupResponseBody
	GetSuccess() *string
}

type ModifySecurityIPGroupResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *ModifySecurityIPGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// > If the request is successful, **Successful*	- is returned. If the request fails, an error message is returned, such as an error code.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B6D17591-B48B-4D31-9CD6-9B9796B2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful:
	//
	// - **true**: The request was successful.
	//
	// - **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifySecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifySecurityIPGroupResponseBody) GetData() *ModifySecurityIPGroupResponseBodyData {
	return s.Data
}

func (s *ModifySecurityIPGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifySecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifySecurityIPGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifySecurityIPGroupResponseBody) SetCode(v string) *ModifySecurityIPGroupResponseBody {
	s.Code = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBody) SetData(v *ModifySecurityIPGroupResponseBodyData) *ModifySecurityIPGroupResponseBody {
	s.Data = v
	return s
}

func (s *ModifySecurityIPGroupResponseBody) SetMessage(v string) *ModifySecurityIPGroupResponseBody {
	s.Message = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBody) SetRequestId(v string) *ModifySecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBody) SetSuccess(v string) *ModifySecurityIPGroupResponseBody {
	s.Success = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifySecurityIPGroupResponseBodyData struct {
	// The information about the cross-product whitelist template.
	GlobalSecurityIPGroup []*ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
}

func (s ModifySecurityIPGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupResponseBodyData) GetGlobalSecurityIPGroup() []*ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	return s.GlobalSecurityIPGroup
}

func (s *ModifySecurityIPGroupResponseBodyData) SetGlobalSecurityIPGroup(v []*ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) *ModifySecurityIPGroupResponseBodyData {
	s.GlobalSecurityIPGroup = v
	return s
}

func (s *ModifySecurityIPGroupResponseBodyData) Validate() error {
	if s.GlobalSecurityIPGroup != nil {
		for _, item := range s.GlobalSecurityIPGroup {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup struct {
	// The IP addresses in the whitelist template.
	//
	// > Separate multiple IP addresses with commas. All IP address whitelists support a combined total of 1,000 IP addresses or address segments.
	//
	// example:
	//
	// 192.168.0.1
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// The name of the IP address whitelist template.
	//
	// example:
	//
	// test
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// The ID of the IP address whitelist template.
	//
	// example:
	//
	// g-b1asblm5ae****
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// The region ID.
	//
	// Example: cn-hangzhou
	//
	// example:
	//
	// cn-beijing
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// The IP type.
	//
	// example:
	//
	// ipv4
	SecurityIpType *string `json:"SecurityIpType,omitempty" xml:"SecurityIpType,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetGIpList() *string {
	return s.GIpList
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetSecurityIpType() *string {
	return s.SecurityIpType
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetWhitelistNetType() *string {
	return s.WhitelistNetType
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetGIpList(v string) *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.GIpList = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetGlobalIgName(v string) *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.GlobalIgName = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetGlobalSecurityGroupId(v string) *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetRegionId(v string) *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.RegionId = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetSecurityIpType(v string) *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.SecurityIpType = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetWhitelistNetType(v string) *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.WhitelistNetType = &v
	return s
}

func (s *ModifySecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
