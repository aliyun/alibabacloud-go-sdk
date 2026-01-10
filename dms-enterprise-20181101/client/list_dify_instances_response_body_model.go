// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDifyInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListDifyInstancesResponseBody
	GetCode() *string
	SetErrorCode(v string) *ListDifyInstancesResponseBody
	GetErrorCode() *string
	SetHttpStatusCode(v int32) *ListDifyInstancesResponseBody
	GetHttpStatusCode() *int32
	SetMaxResults(v int32) *ListDifyInstancesResponseBody
	GetMaxResults() *int32
	SetMessage(v string) *ListDifyInstancesResponseBody
	GetMessage() *string
	SetNextToken(v string) *ListDifyInstancesResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListDifyInstancesResponseBody
	GetRequestId() *string
	SetRoot(v *ListDifyInstancesResponseBodyRoot) *ListDifyInstancesResponseBody
	GetRoot() *ListDifyInstancesResponseBodyRoot
	SetSuccess(v bool) *ListDifyInstancesResponseBody
	GetSuccess() *bool
}

type ListDifyInstancesResponseBody struct {
	Code           *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	ErrorCode      *string                            `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	HttpStatusCode *int32                             `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	MaxResults     *int32                             `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	Message        *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	NextToken      *string                            `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	RequestId      *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Root           *ListDifyInstancesResponseBodyRoot `json:"Root,omitempty" xml:"Root,omitempty" type:"Struct"`
	Success        *bool                              `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListDifyInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDifyInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDifyInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListDifyInstancesResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *ListDifyInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ListDifyInstancesResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDifyInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListDifyInstancesResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDifyInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDifyInstancesResponseBody) GetRoot() *ListDifyInstancesResponseBodyRoot {
	return s.Root
}

func (s *ListDifyInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListDifyInstancesResponseBody) SetCode(v string) *ListDifyInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *ListDifyInstancesResponseBody) SetErrorCode(v string) *ListDifyInstancesResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *ListDifyInstancesResponseBody) SetHttpStatusCode(v int32) *ListDifyInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ListDifyInstancesResponseBody) SetMaxResults(v int32) *ListDifyInstancesResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListDifyInstancesResponseBody) SetMessage(v string) *ListDifyInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *ListDifyInstancesResponseBody) SetNextToken(v string) *ListDifyInstancesResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListDifyInstancesResponseBody) SetRequestId(v string) *ListDifyInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDifyInstancesResponseBody) SetRoot(v *ListDifyInstancesResponseBodyRoot) *ListDifyInstancesResponseBody {
	s.Root = v
	return s
}

func (s *ListDifyInstancesResponseBody) SetSuccess(v bool) *ListDifyInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *ListDifyInstancesResponseBody) Validate() error {
	if s.Root != nil {
		if err := s.Root.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDifyInstancesResponseBodyRoot struct {
	Data []*ListDifyInstancesResponseBodyRootData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
}

func (s ListDifyInstancesResponseBodyRoot) String() string {
	return dara.Prettify(s)
}

func (s ListDifyInstancesResponseBodyRoot) GoString() string {
	return s.String()
}

func (s *ListDifyInstancesResponseBodyRoot) GetData() []*ListDifyInstancesResponseBodyRootData {
	return s.Data
}

func (s *ListDifyInstancesResponseBodyRoot) SetData(v []*ListDifyInstancesResponseBodyRootData) *ListDifyInstancesResponseBodyRoot {
	s.Data = v
	return s
}

func (s *ListDifyInstancesResponseBodyRoot) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDifyInstancesResponseBodyRootData struct {
	AppUuid     *string `json:"AppUuid,omitempty" xml:"AppUuid,omitempty"`
	CreatedTime *string `json:"CreatedTime,omitempty" xml:"CreatedTime,omitempty"`
	// example:
	//
	// 实例描述
	Description           *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Edition               *string `json:"Edition,omitempty" xml:"Edition,omitempty"`
	EnterpriseInternetUrl *string `json:"EnterpriseInternetUrl,omitempty" xml:"EnterpriseInternetUrl,omitempty"`
	EnterpriseIntranetUrl *string `json:"EnterpriseIntranetUrl,omitempty" xml:"EnterpriseIntranetUrl,omitempty"`
	InstanceId            *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	InstanceName          *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	InternetUrl           *string `json:"InternetUrl,omitempty" xml:"InternetUrl,omitempty"`
	IntranetUrl           *string `json:"IntranetUrl,omitempty" xml:"IntranetUrl,omitempty"`
	MajorVersion          *string `json:"MajorVersion,omitempty" xml:"MajorVersion,omitempty"`
	// example:
	//
	// 地域信息
	RegionCode      *string `json:"RegionCode,omitempty" xml:"RegionCode,omitempty"`
	SecurityGroupId *string `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	// example:
	//
	// running
	Status      *string `json:"Status,omitempty" xml:"Status,omitempty"`
	VSwitchId   *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	VpcId       *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
	ZoneId      *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s ListDifyInstancesResponseBodyRootData) String() string {
	return dara.Prettify(s)
}

func (s ListDifyInstancesResponseBodyRootData) GoString() string {
	return s.String()
}

func (s *ListDifyInstancesResponseBodyRootData) GetAppUuid() *string {
	return s.AppUuid
}

func (s *ListDifyInstancesResponseBodyRootData) GetCreatedTime() *string {
	return s.CreatedTime
}

func (s *ListDifyInstancesResponseBodyRootData) GetDescription() *string {
	return s.Description
}

func (s *ListDifyInstancesResponseBodyRootData) GetEdition() *string {
	return s.Edition
}

func (s *ListDifyInstancesResponseBodyRootData) GetEnterpriseInternetUrl() *string {
	return s.EnterpriseInternetUrl
}

func (s *ListDifyInstancesResponseBodyRootData) GetEnterpriseIntranetUrl() *string {
	return s.EnterpriseIntranetUrl
}

func (s *ListDifyInstancesResponseBodyRootData) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListDifyInstancesResponseBodyRootData) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ListDifyInstancesResponseBodyRootData) GetInternetUrl() *string {
	return s.InternetUrl
}

func (s *ListDifyInstancesResponseBodyRootData) GetIntranetUrl() *string {
	return s.IntranetUrl
}

func (s *ListDifyInstancesResponseBodyRootData) GetMajorVersion() *string {
	return s.MajorVersion
}

func (s *ListDifyInstancesResponseBodyRootData) GetRegionCode() *string {
	return s.RegionCode
}

func (s *ListDifyInstancesResponseBodyRootData) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *ListDifyInstancesResponseBodyRootData) GetStatus() *string {
	return s.Status
}

func (s *ListDifyInstancesResponseBodyRootData) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *ListDifyInstancesResponseBodyRootData) GetVpcId() *string {
	return s.VpcId
}

func (s *ListDifyInstancesResponseBodyRootData) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *ListDifyInstancesResponseBodyRootData) GetZoneId() *string {
	return s.ZoneId
}

func (s *ListDifyInstancesResponseBodyRootData) SetAppUuid(v string) *ListDifyInstancesResponseBodyRootData {
	s.AppUuid = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetCreatedTime(v string) *ListDifyInstancesResponseBodyRootData {
	s.CreatedTime = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetDescription(v string) *ListDifyInstancesResponseBodyRootData {
	s.Description = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetEdition(v string) *ListDifyInstancesResponseBodyRootData {
	s.Edition = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetEnterpriseInternetUrl(v string) *ListDifyInstancesResponseBodyRootData {
	s.EnterpriseInternetUrl = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetEnterpriseIntranetUrl(v string) *ListDifyInstancesResponseBodyRootData {
	s.EnterpriseIntranetUrl = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetInstanceId(v string) *ListDifyInstancesResponseBodyRootData {
	s.InstanceId = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetInstanceName(v string) *ListDifyInstancesResponseBodyRootData {
	s.InstanceName = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetInternetUrl(v string) *ListDifyInstancesResponseBodyRootData {
	s.InternetUrl = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetIntranetUrl(v string) *ListDifyInstancesResponseBodyRootData {
	s.IntranetUrl = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetMajorVersion(v string) *ListDifyInstancesResponseBodyRootData {
	s.MajorVersion = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetRegionCode(v string) *ListDifyInstancesResponseBodyRootData {
	s.RegionCode = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetSecurityGroupId(v string) *ListDifyInstancesResponseBodyRootData {
	s.SecurityGroupId = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetStatus(v string) *ListDifyInstancesResponseBodyRootData {
	s.Status = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetVSwitchId(v string) *ListDifyInstancesResponseBodyRootData {
	s.VSwitchId = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetVpcId(v string) *ListDifyInstancesResponseBodyRootData {
	s.VpcId = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetWorkspaceId(v string) *ListDifyInstancesResponseBodyRootData {
	s.WorkspaceId = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) SetZoneId(v string) *ListDifyInstancesResponseBodyRootData {
	s.ZoneId = &v
	return s
}

func (s *ListDifyInstancesResponseBodyRootData) Validate() error {
	return dara.Validate(s)
}
