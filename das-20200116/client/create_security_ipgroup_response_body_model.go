// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSecurityIPGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSecurityIPGroupResponseBody
	GetCode() *string
	SetData(v *CreateSecurityIPGroupResponseBodyData) *CreateSecurityIPGroupResponseBody
	GetData() *CreateSecurityIPGroupResponseBodyData
	SetMessage(v string) *CreateSecurityIPGroupResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateSecurityIPGroupResponseBody
	GetRequestId() *string
	SetSuccess(v string) *CreateSecurityIPGroupResponseBody
	GetSuccess() *string
}

type CreateSecurityIPGroupResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// ListResult<InstanceSSL>
	Data *CreateSecurityIPGroupResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// D578DB3C-06BF-54F2-A78F-C6C25Exxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSecurityIPGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIPGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSecurityIPGroupResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSecurityIPGroupResponseBody) GetData() *CreateSecurityIPGroupResponseBodyData {
	return s.Data
}

func (s *CreateSecurityIPGroupResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSecurityIPGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSecurityIPGroupResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *CreateSecurityIPGroupResponseBody) SetCode(v string) *CreateSecurityIPGroupResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBody) SetData(v *CreateSecurityIPGroupResponseBodyData) *CreateSecurityIPGroupResponseBody {
	s.Data = v
	return s
}

func (s *CreateSecurityIPGroupResponseBody) SetMessage(v string) *CreateSecurityIPGroupResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBody) SetRequestId(v string) *CreateSecurityIPGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBody) SetSuccess(v string) *CreateSecurityIPGroupResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSecurityIPGroupResponseBodyData struct {
	GlobalSecurityIPGroup []*CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup `json:"GlobalSecurityIPGroup,omitempty" xml:"GlobalSecurityIPGroup,omitempty" type:"Repeated"`
}

func (s CreateSecurityIPGroupResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIPGroupResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateSecurityIPGroupResponseBodyData) GetGlobalSecurityIPGroup() []*CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	return s.GlobalSecurityIPGroup
}

func (s *CreateSecurityIPGroupResponseBodyData) SetGlobalSecurityIPGroup(v []*CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) *CreateSecurityIPGroupResponseBodyData {
	s.GlobalSecurityIPGroup = v
	return s
}

func (s *CreateSecurityIPGroupResponseBodyData) Validate() error {
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

type CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup struct {
	// example:
	//
	// 192.168.0.0/24
	GIpList *string `json:"GIpList,omitempty" xml:"GIpList,omitempty"`
	// example:
	//
	// test
	GlobalIgName *string `json:"GlobalIgName,omitempty" xml:"GlobalIgName,omitempty"`
	// example:
	//
	// g-2uztsd6yvhmsqyjXXX
	GlobalSecurityGroupId *string `json:"GlobalSecurityGroupId,omitempty" xml:"GlobalSecurityGroupId,omitempty"`
	// example:
	//
	// cn-shenzhen
	RegionId *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	// example:
	//
	// ipv4
	SecurityIpType *string `json:"SecurityIpType,omitempty" xml:"SecurityIpType,omitempty"`
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) String() string {
	return dara.Prettify(s)
}

func (s CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GoString() string {
	return s.String()
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetGIpList() *string {
	return s.GIpList
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetGlobalIgName() *string {
	return s.GlobalIgName
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetGlobalSecurityGroupId() *string {
	return s.GlobalSecurityGroupId
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetRegionId() *string {
	return s.RegionId
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetSecurityIpType() *string {
	return s.SecurityIpType
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) GetWhitelistNetType() *string {
	return s.WhitelistNetType
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetGIpList(v string) *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.GIpList = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetGlobalIgName(v string) *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.GlobalIgName = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetGlobalSecurityGroupId(v string) *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.GlobalSecurityGroupId = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetRegionId(v string) *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.RegionId = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetSecurityIpType(v string) *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.SecurityIpType = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) SetWhitelistNetType(v string) *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup {
	s.WhitelistNetType = &v
	return s
}

func (s *CreateSecurityIPGroupResponseBodyDataGlobalSecurityIPGroup) Validate() error {
	return dara.Validate(s)
}
