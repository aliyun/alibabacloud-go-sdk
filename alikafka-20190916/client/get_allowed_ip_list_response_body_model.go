// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAllowedIpListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAllowedList(v *GetAllowedIpListResponseBodyAllowedList) *GetAllowedIpListResponseBody
	GetAllowedList() *GetAllowedIpListResponseBodyAllowedList
	SetCode(v int32) *GetAllowedIpListResponseBody
	GetCode() *int32
	SetMessage(v string) *GetAllowedIpListResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAllowedIpListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAllowedIpListResponseBody
	GetSuccess() *bool
}

type GetAllowedIpListResponseBody struct {
	// The IP address whitelist.
	AllowedList *GetAllowedIpListResponseBodyAllowedList `json:"AllowedList,omitempty" xml:"AllowedList,omitempty" type:"Struct"`
	// The HTTP status code returned. The HTTP status code 200 indicates that the request is successful.
	//
	// example:
	//
	// 200
	Code *int32 `json:"Code,omitempty" xml:"Code,omitempty"`
	// The message returned.
	//
	// example:
	//
	// operation success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A421CCD7-5BC5-4B32-8DD8-64668A8FCB56
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAllowedIpListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAllowedIpListResponseBody) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponseBody) GetAllowedList() *GetAllowedIpListResponseBodyAllowedList {
	return s.AllowedList
}

func (s *GetAllowedIpListResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetAllowedIpListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAllowedIpListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAllowedIpListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAllowedIpListResponseBody) SetAllowedList(v *GetAllowedIpListResponseBodyAllowedList) *GetAllowedIpListResponseBody {
	s.AllowedList = v
	return s
}

func (s *GetAllowedIpListResponseBody) SetCode(v int32) *GetAllowedIpListResponseBody {
	s.Code = &v
	return s
}

func (s *GetAllowedIpListResponseBody) SetMessage(v string) *GetAllowedIpListResponseBody {
	s.Message = &v
	return s
}

func (s *GetAllowedIpListResponseBody) SetRequestId(v string) *GetAllowedIpListResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAllowedIpListResponseBody) SetSuccess(v bool) *GetAllowedIpListResponseBody {
	s.Success = &v
	return s
}

func (s *GetAllowedIpListResponseBody) Validate() error {
	if s.AllowedList != nil {
		if err := s.AllowedList.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAllowedIpListResponseBodyAllowedList struct {
	// The deployment mode of the instance. Valid values:
	//
	// 	- **4**: allows access from the Internet and a virtual private cloud (VPC).
	//
	// 	- **5**: allows access from a VPC.
	//
	// >  Only integrators need to concern themselves with the value of this parameter.
	//
	// example:
	//
	// 4
	DeployType *int32 `json:"DeployType,omitempty" xml:"DeployType,omitempty"`
	// The whitelist for access from the Internet.
	InternetList []*GetAllowedIpListResponseBodyAllowedListInternetList `json:"InternetList,omitempty" xml:"InternetList,omitempty" type:"Repeated"`
	// The whitelist for access from a virtual private cloud (VPC).
	VpcList []*GetAllowedIpListResponseBodyAllowedListVpcList `json:"VpcList,omitempty" xml:"VpcList,omitempty" type:"Repeated"`
}

func (s GetAllowedIpListResponseBodyAllowedList) String() string {
	return dara.Prettify(s)
}

func (s GetAllowedIpListResponseBodyAllowedList) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponseBodyAllowedList) GetDeployType() *int32 {
	return s.DeployType
}

func (s *GetAllowedIpListResponseBodyAllowedList) GetInternetList() []*GetAllowedIpListResponseBodyAllowedListInternetList {
	return s.InternetList
}

func (s *GetAllowedIpListResponseBodyAllowedList) GetVpcList() []*GetAllowedIpListResponseBodyAllowedListVpcList {
	return s.VpcList
}

func (s *GetAllowedIpListResponseBodyAllowedList) SetDeployType(v int32) *GetAllowedIpListResponseBodyAllowedList {
	s.DeployType = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedList) SetInternetList(v []*GetAllowedIpListResponseBodyAllowedListInternetList) *GetAllowedIpListResponseBodyAllowedList {
	s.InternetList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedList) SetVpcList(v []*GetAllowedIpListResponseBodyAllowedListVpcList) *GetAllowedIpListResponseBodyAllowedList {
	s.VpcList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedList) Validate() error {
	if s.InternetList != nil {
		for _, item := range s.InternetList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VpcList != nil {
		for _, item := range s.VpcList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetAllowedIpListResponseBodyAllowedListInternetList struct {
	// The group to which the IP address whitelist belongs.
	AllowedIpGroup map[string]*string `json:"AllowedIpGroup,omitempty" xml:"AllowedIpGroup,omitempty"`
	// The information about the IP address whitelist.
	AllowedIpList []*string          `json:"AllowedIpList,omitempty" xml:"AllowedIpList,omitempty" type:"Repeated"`
	BlackIPList   []*string          `json:"BlackIPList,omitempty" xml:"BlackIPList,omitempty" type:"Repeated"`
	BlackIPMap    map[string]*string `json:"BlackIPMap,omitempty" xml:"BlackIPMap,omitempty"`
	// The port range. Valid value:
	//
	// **9093/9093**.
	//
	// example:
	//
	// 9093/9093
	PortRange                      *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SecurityGroupId                *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	UserDefinedSharedSecurityGroup *bool   `json:"UserDefinedSharedSecurityGroup,omitempty" xml:"UserDefinedSharedSecurityGroup,omitempty"`
}

func (s GetAllowedIpListResponseBodyAllowedListInternetList) String() string {
	return dara.Prettify(s)
}

func (s GetAllowedIpListResponseBodyAllowedListInternetList) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) GetAllowedIpGroup() map[string]*string {
	return s.AllowedIpGroup
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) GetAllowedIpList() []*string {
	return s.AllowedIpList
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) GetBlackIPList() []*string {
	return s.BlackIPList
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) GetBlackIPMap() map[string]*string {
	return s.BlackIPMap
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) GetPortRange() *string {
	return s.PortRange
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) GetUserDefinedSharedSecurityGroup() *bool {
	return s.UserDefinedSharedSecurityGroup
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetAllowedIpGroup(v map[string]*string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.AllowedIpGroup = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetAllowedIpList(v []*string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.AllowedIpList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetBlackIPList(v []*string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.BlackIPList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetBlackIPMap(v map[string]*string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.BlackIPMap = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetPortRange(v string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.PortRange = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetSecurityGroupId(v string) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.SecurityGroupId = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) SetUserDefinedSharedSecurityGroup(v bool) *GetAllowedIpListResponseBodyAllowedListInternetList {
	s.UserDefinedSharedSecurityGroup = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListInternetList) Validate() error {
	return dara.Validate(s)
}

type GetAllowedIpListResponseBodyAllowedListVpcList struct {
	// The group to which the IP address whitelist belongs.
	AllowedIpGroup map[string]*string `json:"AllowedIpGroup,omitempty" xml:"AllowedIpGroup,omitempty"`
	// The information about the IP address whitelist.
	AllowedIpList []*string          `json:"AllowedIpList,omitempty" xml:"AllowedIpList,omitempty" type:"Repeated"`
	BlackIPList   []*string          `json:"BlackIPList,omitempty" xml:"BlackIPList,omitempty" type:"Repeated"`
	BlackIPMap    map[string]*string `json:"BlackIPMap,omitempty" xml:"BlackIPMap,omitempty"`
	// The port range. Valid value:
	//
	// **9092/9092**.
	//
	// example:
	//
	// 9092/9092
	PortRange                      *string `json:"PortRange,omitempty" xml:"PortRange,omitempty"`
	SecurityGroupId                *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	UserDefinedSharedSecurityGroup *bool   `json:"UserDefinedSharedSecurityGroup,omitempty" xml:"UserDefinedSharedSecurityGroup,omitempty"`
}

func (s GetAllowedIpListResponseBodyAllowedListVpcList) String() string {
	return dara.Prettify(s)
}

func (s GetAllowedIpListResponseBodyAllowedListVpcList) GoString() string {
	return s.String()
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) GetAllowedIpGroup() map[string]*string {
	return s.AllowedIpGroup
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) GetAllowedIpList() []*string {
	return s.AllowedIpList
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) GetBlackIPList() []*string {
	return s.BlackIPList
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) GetBlackIPMap() map[string]*string {
	return s.BlackIPMap
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) GetPortRange() *string {
	return s.PortRange
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) GetUserDefinedSharedSecurityGroup() *bool {
	return s.UserDefinedSharedSecurityGroup
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetAllowedIpGroup(v map[string]*string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.AllowedIpGroup = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetAllowedIpList(v []*string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.AllowedIpList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetBlackIPList(v []*string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.BlackIPList = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetBlackIPMap(v map[string]*string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.BlackIPMap = v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetPortRange(v string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.PortRange = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetSecurityGroupId(v string) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.SecurityGroupId = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) SetUserDefinedSharedSecurityGroup(v bool) *GetAllowedIpListResponseBodyAllowedListVpcList {
	s.UserDefinedSharedSecurityGroup = &v
	return s
}

func (s *GetAllowedIpListResponseBodyAllowedListVpcList) Validate() error {
	return dara.Validate(s)
}
