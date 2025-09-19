// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSasContainerWebDefenseRuleApplicationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContainerWebDefenseAppList(v []*GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) *GetSasContainerWebDefenseRuleApplicationResponseBody
	GetContainerWebDefenseAppList() []*GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList
	SetRequestId(v string) *GetSasContainerWebDefenseRuleApplicationResponseBody
	GetRequestId() *string
}

type GetSasContainerWebDefenseRuleApplicationResponseBody struct {
	// The applications.
	ContainerWebDefenseAppList []*GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList `json:"ContainerWebDefenseAppList,omitempty" xml:"ContainerWebDefenseAppList,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 09969D2C-4FAD-429E-BFBF-9A60DEF8****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetSasContainerWebDefenseRuleApplicationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleApplicationResponseBody) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBody) GetContainerWebDefenseAppList() []*GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList {
	return s.ContainerWebDefenseAppList
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBody) SetContainerWebDefenseAppList(v []*GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) *GetSasContainerWebDefenseRuleApplicationResponseBody {
	s.ContainerWebDefenseAppList = v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBody) SetRequestId(v string) *GetSasContainerWebDefenseRuleApplicationResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList struct {
	// The user ID.
	//
	// example:
	//
	// 5944922169365****
	AliUid *int64 `json:"AliUid,omitempty" xml:"AliUid,omitempty"`
	// The ID of the container cluster.
	//
	// >  The IDs of clusters can be obtained by using the [DescribeGroupedContainerInstances](https://help.aliyun.com/document_detail/182997.html) operation.
	//
	// example:
	//
	// cfb41a869c71e4678a97021582dd8****
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// The time when the application was created. Unit: milliseconds.
	//
	// example:
	//
	// 1677839038000
	GmtCreate *int64 `json:"GmtCreate,omitempty" xml:"GmtCreate,omitempty"`
	// The last modification time. Unit: milliseconds.
	//
	// example:
	//
	// 1667891185000
	GmtModified *int64 `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the node.
	//
	// example:
	//
	// 143761
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the rule.
	//
	// example:
	//
	// 403327
	RuleId *int64 `json:"RuleId,omitempty" xml:"RuleId,omitempty"`
	// The value of the application label.
	//
	// example:
	//
	// app:test
	Tag *string `json:"Tag,omitempty" xml:"Tag,omitempty"`
}

func (s GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) String() string {
	return dara.Prettify(s)
}

func (s GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) GoString() string {
	return s.String()
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) GetAliUid() *int64 {
	return s.AliUid
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) GetGmtCreate() *int64 {
	return s.GmtCreate
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) GetGmtModified() *int64 {
	return s.GmtModified
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) GetId() *int64 {
	return s.Id
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) GetRuleId() *int64 {
	return s.RuleId
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) GetTag() *string {
	return s.Tag
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) SetAliUid(v int64) *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList {
	s.AliUid = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) SetClusterId(v string) *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList {
	s.ClusterId = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) SetGmtCreate(v int64) *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList {
	s.GmtCreate = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) SetGmtModified(v int64) *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList {
	s.GmtModified = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) SetId(v int64) *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList {
	s.Id = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) SetRuleId(v int64) *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList {
	s.RuleId = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) SetTag(v string) *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList {
	s.Tag = &v
	return s
}

func (s *GetSasContainerWebDefenseRuleApplicationResponseBodyContainerWebDefenseAppList) Validate() error {
	return dara.Validate(s)
}
