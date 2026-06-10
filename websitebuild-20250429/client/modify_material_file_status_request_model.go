// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterialFileStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBizId(v string) *ModifyMaterialFileStatusRequest
	GetBizId() *string
	SetFileIds(v []*string) *ModifyMaterialFileStatusRequest
	GetFileIds() []*string
	SetStatus(v string) *ModifyMaterialFileStatusRequest
	GetStatus() *string
}

type ModifyMaterialFileStatusRequest struct {
	// Business instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// WS12345678
	BizId *string `json:"BizId,omitempty" xml:"BizId,omitempty"`
	// File ID
	//
	// This parameter is required.
	FileIds []*string `json:"FileIds,omitempty" xml:"FileIds,omitempty" type:"Repeated"`
	// Modified status
	//
	// This parameter is required.
	//
	// example:
	//
	// NORMAL
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyMaterialFileStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterialFileStatusRequest) GoString() string {
	return s.String()
}

func (s *ModifyMaterialFileStatusRequest) GetBizId() *string {
	return s.BizId
}

func (s *ModifyMaterialFileStatusRequest) GetFileIds() []*string {
	return s.FileIds
}

func (s *ModifyMaterialFileStatusRequest) GetStatus() *string {
	return s.Status
}

func (s *ModifyMaterialFileStatusRequest) SetBizId(v string) *ModifyMaterialFileStatusRequest {
	s.BizId = &v
	return s
}

func (s *ModifyMaterialFileStatusRequest) SetFileIds(v []*string) *ModifyMaterialFileStatusRequest {
	s.FileIds = v
	return s
}

func (s *ModifyMaterialFileStatusRequest) SetStatus(v string) *ModifyMaterialFileStatusRequest {
	s.Status = &v
	return s
}

func (s *ModifyMaterialFileStatusRequest) Validate() error {
	return dara.Validate(s)
}
