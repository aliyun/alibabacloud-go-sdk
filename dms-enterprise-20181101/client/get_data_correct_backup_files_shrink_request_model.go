// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectBackupFilesShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionDetailShrink(v string) *GetDataCorrectBackupFilesShrinkRequest
	GetActionDetailShrink() *string
	SetOrderId(v int64) *GetDataCorrectBackupFilesShrinkRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataCorrectBackupFilesShrinkRequest
	GetTid() *int64
}

type GetDataCorrectBackupFilesShrinkRequest struct {
	// The parameters that are required to perform the operation. You do not need to specify this parameter.
	//
	// example:
	//
	// {}
	ActionDetailShrink *string `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
	// The ID of the ticket. You can call the [ListOrders](https://help.aliyun.com/document_detail/144643.html) operation to obtain the ticket ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4200000
	OrderId *int64 `json:"OrderId,omitempty" xml:"OrderId,omitempty"`
	// The ID of the tenant. You can call the [GetUserActiveTenant](https://help.aliyun.com/document_detail/198073.html) or [ListUserTenants](https://help.aliyun.com/document_detail/198074.html) operation to obtain the tenant ID.
	//
	// example:
	//
	// 3***
	Tid *int64 `json:"Tid,omitempty" xml:"Tid,omitempty"`
}

func (s GetDataCorrectBackupFilesShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectBackupFilesShrinkRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesShrinkRequest) GetActionDetailShrink() *string {
	return s.ActionDetailShrink
}

func (s *GetDataCorrectBackupFilesShrinkRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataCorrectBackupFilesShrinkRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataCorrectBackupFilesShrinkRequest) SetActionDetailShrink(v string) *GetDataCorrectBackupFilesShrinkRequest {
	s.ActionDetailShrink = &v
	return s
}

func (s *GetDataCorrectBackupFilesShrinkRequest) SetOrderId(v int64) *GetDataCorrectBackupFilesShrinkRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCorrectBackupFilesShrinkRequest) SetTid(v int64) *GetDataCorrectBackupFilesShrinkRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCorrectBackupFilesShrinkRequest) Validate() error {
	return dara.Validate(s)
}
