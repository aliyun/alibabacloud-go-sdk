// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeNsasSuspEventTypeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContainerFieldName(v string) *DescribeNsasSuspEventTypeRequest
	GetContainerFieldName() *string
	SetContainerFieldValue(v string) *DescribeNsasSuspEventTypeRequest
	GetContainerFieldValue() *string
	SetFrom(v string) *DescribeNsasSuspEventTypeRequest
	GetFrom() *string
	SetLang(v string) *DescribeNsasSuspEventTypeRequest
	GetLang() *string
	SetMultiAccountActionType(v int32) *DescribeNsasSuspEventTypeRequest
	GetMultiAccountActionType() *int32
	SetName(v string) *DescribeNsasSuspEventTypeRequest
	GetName() *string
	SetRemark(v string) *DescribeNsasSuspEventTypeRequest
	GetRemark() *string
	SetSourceIp(v string) *DescribeNsasSuspEventTypeRequest
	GetSourceIp() *string
	SetSupportOperateCodeList(v []*string) *DescribeNsasSuspEventTypeRequest
	GetSupportOperateCodeList() []*string
	SetUuids(v string) *DescribeNsasSuspEventTypeRequest
	GetUuids() *string
}

type DescribeNsasSuspEventTypeRequest struct {
	// The name of the container field. Valid values:
	//
	// 	- **clusterId**: the ID of the cluster
	//
	// example:
	//
	// clusterId
	ContainerFieldName *string `json:"ContainerFieldName,omitempty" xml:"ContainerFieldName,omitempty"`
	// The value of the container field.
	//
	// example:
	//
	// ca3108551c83c4d949106e1ab9e1e****
	ContainerFieldValue *string `json:"ContainerFieldValue,omitempty" xml:"ContainerFieldValue,omitempty"`
	// The ID of the request source. Set the value to **sas**, which indicates that the request is sent from Security Center.
	//
	// example:
	//
	// sas
	From *string `json:"From,omitempty" xml:"From,omitempty"`
	// The language of the content within the request and response. Default value: **zh**. Valid values:
	//
	// 	- **zh**: Chinese
	//
	// 	- **en**: English
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The type of the accounts that you want to query. Default value: **0**. Valid values:
	//
	// 	- **0**: the current account.
	//
	// 	- **1**: all accounts.
	//
	// example:
	//
	// 0
	MultiAccountActionType *int32 `json:"MultiAccountActionType,omitempty" xml:"MultiAccountActionType,omitempty"`
	// The name of the alert type.
	//
	// example:
	//
	// Unusual Logon
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The name of the server.
	//
	// example:
	//
	// testECS
	Remark *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	// The source IP address of the request.
	//
	// example:
	//
	// 180.212.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
	// List of supported alarm operation types.
	SupportOperateCodeList []*string `json:"SupportOperateCodeList,omitempty" xml:"SupportOperateCodeList,omitempty" type:"Repeated"`
	// The UUIDs of servers. Separate multiple UUIDs with commas (,).
	//
	// > You can call the [DescribeCloudCenterInstances](https://help.aliyun.com/document_detail/141932.html) operation to query the UUIDs of servers.
	//
	// example:
	//
	// f56406cb-916d-42db-b6f7-2ff79e34****
	Uuids *string `json:"Uuids,omitempty" xml:"Uuids,omitempty"`
}

func (s DescribeNsasSuspEventTypeRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeNsasSuspEventTypeRequest) GoString() string {
	return s.String()
}

func (s *DescribeNsasSuspEventTypeRequest) GetContainerFieldName() *string {
	return s.ContainerFieldName
}

func (s *DescribeNsasSuspEventTypeRequest) GetContainerFieldValue() *string {
	return s.ContainerFieldValue
}

func (s *DescribeNsasSuspEventTypeRequest) GetFrom() *string {
	return s.From
}

func (s *DescribeNsasSuspEventTypeRequest) GetLang() *string {
	return s.Lang
}

func (s *DescribeNsasSuspEventTypeRequest) GetMultiAccountActionType() *int32 {
	return s.MultiAccountActionType
}

func (s *DescribeNsasSuspEventTypeRequest) GetName() *string {
	return s.Name
}

func (s *DescribeNsasSuspEventTypeRequest) GetRemark() *string {
	return s.Remark
}

func (s *DescribeNsasSuspEventTypeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *DescribeNsasSuspEventTypeRequest) GetSupportOperateCodeList() []*string {
	return s.SupportOperateCodeList
}

func (s *DescribeNsasSuspEventTypeRequest) GetUuids() *string {
	return s.Uuids
}

func (s *DescribeNsasSuspEventTypeRequest) SetContainerFieldName(v string) *DescribeNsasSuspEventTypeRequest {
	s.ContainerFieldName = &v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) SetContainerFieldValue(v string) *DescribeNsasSuspEventTypeRequest {
	s.ContainerFieldValue = &v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) SetFrom(v string) *DescribeNsasSuspEventTypeRequest {
	s.From = &v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) SetLang(v string) *DescribeNsasSuspEventTypeRequest {
	s.Lang = &v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) SetMultiAccountActionType(v int32) *DescribeNsasSuspEventTypeRequest {
	s.MultiAccountActionType = &v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) SetName(v string) *DescribeNsasSuspEventTypeRequest {
	s.Name = &v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) SetRemark(v string) *DescribeNsasSuspEventTypeRequest {
	s.Remark = &v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) SetSourceIp(v string) *DescribeNsasSuspEventTypeRequest {
	s.SourceIp = &v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) SetSupportOperateCodeList(v []*string) *DescribeNsasSuspEventTypeRequest {
	s.SupportOperateCodeList = v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) SetUuids(v string) *DescribeNsasSuspEventTypeRequest {
	s.Uuids = &v
	return s
}

func (s *DescribeNsasSuspEventTypeRequest) Validate() error {
	return dara.Validate(s)
}
