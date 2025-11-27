// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSmsOcrOssInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOwnerId(v int64) *GetSmsOcrOssInfoRequest
	GetOwnerId() *int64
	SetResourceOwnerAccount(v string) *GetSmsOcrOssInfoRequest
	GetResourceOwnerAccount() *string
	SetResourceOwnerId(v int64) *GetSmsOcrOssInfoRequest
	GetResourceOwnerId() *int64
	SetTaskType(v string) *GetSmsOcrOssInfoRequest
	GetTaskType() *string
}

type GetSmsOcrOssInfoRequest struct {
	OwnerId              *int64  `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	ResourceOwnerAccount *string `json:"ResourceOwnerAccount,omitempty" xml:"ResourceOwnerAccount,omitempty"`
	ResourceOwnerId      *int64  `json:"ResourceOwnerId,omitempty" xml:"ResourceOwnerId,omitempty"`
	// OCR任务类型
	//
	// example:
	//
	// 示例值
	TaskType *string `json:"TaskType,omitempty" xml:"TaskType,omitempty"`
}

func (s GetSmsOcrOssInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSmsOcrOssInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSmsOcrOssInfoRequest) GetOwnerId() *int64 {
	return s.OwnerId
}

func (s *GetSmsOcrOssInfoRequest) GetResourceOwnerAccount() *string {
	return s.ResourceOwnerAccount
}

func (s *GetSmsOcrOssInfoRequest) GetResourceOwnerId() *int64 {
	return s.ResourceOwnerId
}

func (s *GetSmsOcrOssInfoRequest) GetTaskType() *string {
	return s.TaskType
}

func (s *GetSmsOcrOssInfoRequest) SetOwnerId(v int64) *GetSmsOcrOssInfoRequest {
	s.OwnerId = &v
	return s
}

func (s *GetSmsOcrOssInfoRequest) SetResourceOwnerAccount(v string) *GetSmsOcrOssInfoRequest {
	s.ResourceOwnerAccount = &v
	return s
}

func (s *GetSmsOcrOssInfoRequest) SetResourceOwnerId(v int64) *GetSmsOcrOssInfoRequest {
	s.ResourceOwnerId = &v
	return s
}

func (s *GetSmsOcrOssInfoRequest) SetTaskType(v string) *GetSmsOcrOssInfoRequest {
	s.TaskType = &v
	return s
}

func (s *GetSmsOcrOssInfoRequest) Validate() error {
	return dara.Validate(s)
}
