// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataArchiveOrderShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetComment(v string) *CreateDataArchiveOrderShrinkRequest
	GetComment() *string
	SetParamShrink(v string) *CreateDataArchiveOrderShrinkRequest
	GetParamShrink() *string
	SetParentId(v int64) *CreateDataArchiveOrderShrinkRequest
	GetParentId() *int64
	SetPluginType(v string) *CreateDataArchiveOrderShrinkRequest
	GetPluginType() *string
	SetRelatedUserListShrink(v string) *CreateDataArchiveOrderShrinkRequest
	GetRelatedUserListShrink() *string
	SetTid(v int64) *CreateDataArchiveOrderShrinkRequest
	GetTid() *int64
}

type CreateDataArchiveOrderShrinkRequest struct {
	// The description of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The parameters for archiving data.
	//
	// This parameter is required.
	ParamShrink *string `json:"Param,omitempty" xml:"Param,omitempty"`
	// The ID of the parent ticket. A parent ticket is generated only when a child ticket is created.
	//
	// example:
	//
	// 123****
	ParentId *int64 `json:"ParentId,omitempty" xml:"ParentId,omitempty"`
	// The type of the plug-in. Default value: DATA_ARCHIVE.
	//
	// example:
	//
	// DATA_ARCHIVE
	PluginType *string `json:"PluginType,omitempty" xml:"PluginType,omitempty"`
	// The list of the related users.
	RelatedUserListShrink *string `json:"RelatedUserList,omitempty" xml:"RelatedUserList,omitempty"`
	// The tenant ID. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s CreateDataArchiveOrderShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateDataArchiveOrderShrinkRequest) GoString() string {
	return s.String()
}

func (s *CreateDataArchiveOrderShrinkRequest) GetComment() *string {
	return s.Comment
}

func (s *CreateDataArchiveOrderShrinkRequest) GetParamShrink() *string {
	return s.ParamShrink
}

func (s *CreateDataArchiveOrderShrinkRequest) GetParentId() *int64 {
	return s.ParentId
}

func (s *CreateDataArchiveOrderShrinkRequest) GetPluginType() *string {
	return s.PluginType
}

func (s *CreateDataArchiveOrderShrinkRequest) GetRelatedUserListShrink() *string {
	return s.RelatedUserListShrink
}

func (s *CreateDataArchiveOrderShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *CreateDataArchiveOrderShrinkRequest) SetComment(v string) *CreateDataArchiveOrderShrinkRequest {
	s.Comment = &v
	return s
}

func (s *CreateDataArchiveOrderShrinkRequest) SetParamShrink(v string) *CreateDataArchiveOrderShrinkRequest {
	s.ParamShrink = &v
	return s
}

func (s *CreateDataArchiveOrderShrinkRequest) SetParentId(v int64) *CreateDataArchiveOrderShrinkRequest {
	s.ParentId = &v
	return s
}

func (s *CreateDataArchiveOrderShrinkRequest) SetPluginType(v string) *CreateDataArchiveOrderShrinkRequest {
	s.PluginType = &v
	return s
}

func (s *CreateDataArchiveOrderShrinkRequest) SetRelatedUserListShrink(v string) *CreateDataArchiveOrderShrinkRequest {
	s.RelatedUserListShrink = &v
	return s
}

func (s *CreateDataArchiveOrderShrinkRequest) SetTid(v int64) *CreateDataArchiveOrderShrinkRequest {
	s.Tid = &v
	return s
}

func (s *CreateDataArchiveOrderShrinkRequest) Validate() error {
	return dara.Validate(s)
}
