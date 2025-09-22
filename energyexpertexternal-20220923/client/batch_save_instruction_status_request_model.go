// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchSaveInstructionStatusRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFactoryId(v string) *BatchSaveInstructionStatusRequest
	GetFactoryId() *string
	SetPKey(v string) *BatchSaveInstructionStatusRequest
	GetPKey() *string
	SetStatusList(v string) *BatchSaveInstructionStatusRequest
	GetStatusList() *string
}

type BatchSaveInstructionStatusRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ***
	FactoryId *string `json:"factoryId,omitempty" xml:"factoryId,omitempty"`
	// example:
	//
	// ib
	PKey       *string `json:"pKey,omitempty" xml:"pKey,omitempty"`
	StatusList *string `json:"statusList,omitempty" xml:"statusList,omitempty"`
}

func (s BatchSaveInstructionStatusRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchSaveInstructionStatusRequest) GoString() string {
	return s.String()
}

func (s *BatchSaveInstructionStatusRequest) GetFactoryId() *string {
	return s.FactoryId
}

func (s *BatchSaveInstructionStatusRequest) GetPKey() *string {
	return s.PKey
}

func (s *BatchSaveInstructionStatusRequest) GetStatusList() *string {
	return s.StatusList
}

func (s *BatchSaveInstructionStatusRequest) SetFactoryId(v string) *BatchSaveInstructionStatusRequest {
	s.FactoryId = &v
	return s
}

func (s *BatchSaveInstructionStatusRequest) SetPKey(v string) *BatchSaveInstructionStatusRequest {
	s.PKey = &v
	return s
}

func (s *BatchSaveInstructionStatusRequest) SetStatusList(v string) *BatchSaveInstructionStatusRequest {
	s.StatusList = &v
	return s
}

func (s *BatchSaveInstructionStatusRequest) Validate() error {
	return dara.Validate(s)
}
