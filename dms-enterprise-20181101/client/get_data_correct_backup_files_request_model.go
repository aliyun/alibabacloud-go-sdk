// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataCorrectBackupFilesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetActionDetail(v map[string]interface{}) *GetDataCorrectBackupFilesRequest
	GetActionDetail() map[string]interface{}
	SetOrderId(v int64) *GetDataCorrectBackupFilesRequest
	GetOrderId() *int64
	SetTid(v int64) *GetDataCorrectBackupFilesRequest
	GetTid() *int64
}

type GetDataCorrectBackupFilesRequest struct {
	// The parameters that are required to perform the operation. You do not need to specify this parameter.
	//
	// example:
	//
	// {}
	ActionDetail map[string]interface{} `json:"ActionDetail,omitempty" xml:"ActionDetail,omitempty"`
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

func (s GetDataCorrectBackupFilesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataCorrectBackupFilesRequest) GoString() string {
	return s.String()
}

func (s *GetDataCorrectBackupFilesRequest) GetActionDetail() map[string]interface{} {
	return s.ActionDetail
}

func (s *GetDataCorrectBackupFilesRequest) GetOrderId() *int64 {
	return s.OrderId
}

func (s *GetDataCorrectBackupFilesRequest) GetTid() *int64 {
	return s.Tid
}

func (s *GetDataCorrectBackupFilesRequest) SetActionDetail(v map[string]interface{}) *GetDataCorrectBackupFilesRequest {
	s.ActionDetail = v
	return s
}

func (s *GetDataCorrectBackupFilesRequest) SetOrderId(v int64) *GetDataCorrectBackupFilesRequest {
	s.OrderId = &v
	return s
}

func (s *GetDataCorrectBackupFilesRequest) SetTid(v int64) *GetDataCorrectBackupFilesRequest {
	s.Tid = &v
	return s
}

func (s *GetDataCorrectBackupFilesRequest) Validate() error {
	return dara.Validate(s)
}
