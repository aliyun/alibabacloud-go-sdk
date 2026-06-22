// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebLockProcessStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDealAll(v int32) *ModifyWebLockProcessStatusRequest
	GetDealAll() *int32
	SetOperateInfo(v string) *ModifyWebLockProcessStatusRequest
	GetOperateInfo() *string
	SetProcessPath(v []*string) *ModifyWebLockProcessStatusRequest
	GetProcessPath() []*string
	SetStatus(v int32) *ModifyWebLockProcessStatusRequest
	GetStatus() *int32
	SetUuid(v string) *ModifyWebLockProcessStatusRequest
	GetUuid() *string
}

type ModifyWebLockProcessStatusRequest struct {
	// The identity is simultaneous processing of servers that have the same process. Valid values:
	//
	// - **0**: Do not use simultaneous processing.
	//
	// - **1**: Use simultaneous processing.
	//
	// example:
	//
	// 1
	DealAll *int32 `json:"DealAll,omitempty" xml:"DealAll,omitempty"`
	// The operation parameters for batch setting the tamper-proofing process status. The value is in JSON format.
	//
	// example:
	//
	// [{"processPath":"/etc/test1","uuid":"0c1714dc-f7a3-4265-8364-7aa3fce8****"},{"processPath":"/etc/test2","uuid":"1cc45e7d-7698-4b2c-89d8-e8cba407****"}]
	OperateInfo *string `json:"OperateInfo,omitempty" xml:"OperateInfo,omitempty"`
	// The list of process paths.
	ProcessPath []*string `json:"ProcessPath,omitempty" xml:"ProcessPath,omitempty" type:"Repeated"`
	// The status of the tamper-proofing process. Valid values:
	//
	// - **0**: Remove from the whitelist.
	//
	// - **1**: Add to the whitelist.
	//
	// example:
	//
	// 1
	Status *int32 `json:"Status,omitempty" xml:"Status,omitempty"`
	// The UUID of the server for which you want to set the tamper-proofing process status.
	//
	// example:
	//
	// bc8510e7-7327-4030-b75c-956e434d****
	Uuid *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s ModifyWebLockProcessStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebLockProcessStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyWebLockProcessStatusRequest) GetDealAll() *int32 {
	return s.DealAll
}

func (s *ModifyWebLockProcessStatusRequest) GetOperateInfo() *string {
	return s.OperateInfo
}

func (s *ModifyWebLockProcessStatusRequest) GetProcessPath() []*string {
	return s.ProcessPath
}

func (s *ModifyWebLockProcessStatusRequest) GetStatus() *int32 {
	return s.Status
}

func (s *ModifyWebLockProcessStatusRequest) GetUuid() *string {
	return s.Uuid
}

func (s *ModifyWebLockProcessStatusRequest) SetDealAll(v int32) *ModifyWebLockProcessStatusRequest {
	s.DealAll = &v
	return s
}

func (s *ModifyWebLockProcessStatusRequest) SetOperateInfo(v string) *ModifyWebLockProcessStatusRequest {
	s.OperateInfo = &v
	return s
}

func (s *ModifyWebLockProcessStatusRequest) SetProcessPath(v []*string) *ModifyWebLockProcessStatusRequest {
	s.ProcessPath = v
	return s
}

func (s *ModifyWebLockProcessStatusRequest) SetStatus(v int32) *ModifyWebLockProcessStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyWebLockProcessStatusRequest) SetUuid(v string) *ModifyWebLockProcessStatusRequest {
	s.Uuid = &v
	return s
}

func (s *ModifyWebLockProcessStatusRequest) Validate() error {
	return dara.Validate(s)
}
