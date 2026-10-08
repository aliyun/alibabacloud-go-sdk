// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBatchHandoverAssetRequest interface {
	dara.Model
	String() string
	GoString() string
	SetHandoverCommand(v *BatchHandoverAssetRequestHandoverCommand) *BatchHandoverAssetRequest
	GetHandoverCommand() *BatchHandoverAssetRequestHandoverCommand
	SetOpTenantId(v int64) *BatchHandoverAssetRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *BatchHandoverAssetRequest
	GetOpUserId() *string
}

type BatchHandoverAssetRequest struct {
	// This parameter is required.
	HandoverCommand *BatchHandoverAssetRequestHandoverCommand `json:"HandoverCommand,omitempty" xml:"HandoverCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s BatchHandoverAssetRequest) String() string {
	return dara.Prettify(s)
}

func (s BatchHandoverAssetRequest) GoString() string {
	return s.String()
}

func (s *BatchHandoverAssetRequest) GetHandoverCommand() *BatchHandoverAssetRequestHandoverCommand {
	return s.HandoverCommand
}

func (s *BatchHandoverAssetRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *BatchHandoverAssetRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *BatchHandoverAssetRequest) SetHandoverCommand(v *BatchHandoverAssetRequestHandoverCommand) *BatchHandoverAssetRequest {
	s.HandoverCommand = v
	return s
}

func (s *BatchHandoverAssetRequest) SetOpTenantId(v int64) *BatchHandoverAssetRequest {
	s.OpTenantId = &v
	return s
}

func (s *BatchHandoverAssetRequest) SetOpUserId(v string) *BatchHandoverAssetRequest {
	s.OpUserId = &v
	return s
}

func (s *BatchHandoverAssetRequest) Validate() error {
	if s.HandoverCommand != nil {
		if err := s.HandoverCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BatchHandoverAssetRequestHandoverCommand struct {
	// This parameter is required.
	GuidList []*string `json:"GuidList,omitempty" xml:"GuidList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 300004567
	TargetUserId *string `json:"TargetUserId,omitempty" xml:"TargetUserId,omitempty"`
}

func (s BatchHandoverAssetRequestHandoverCommand) String() string {
	return dara.Prettify(s)
}

func (s BatchHandoverAssetRequestHandoverCommand) GoString() string {
	return s.String()
}

func (s *BatchHandoverAssetRequestHandoverCommand) GetGuidList() []*string {
	return s.GuidList
}

func (s *BatchHandoverAssetRequestHandoverCommand) GetTargetUserId() *string {
	return s.TargetUserId
}

func (s *BatchHandoverAssetRequestHandoverCommand) SetGuidList(v []*string) *BatchHandoverAssetRequestHandoverCommand {
	s.GuidList = v
	return s
}

func (s *BatchHandoverAssetRequestHandoverCommand) SetTargetUserId(v string) *BatchHandoverAssetRequestHandoverCommand {
	s.TargetUserId = &v
	return s
}

func (s *BatchHandoverAssetRequestHandoverCommand) Validate() error {
	return dara.Validate(s)
}
