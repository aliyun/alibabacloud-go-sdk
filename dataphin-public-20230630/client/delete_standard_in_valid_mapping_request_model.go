// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardInValidMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteStandardInValidMappingRequestDeleteCommand) *DeleteStandardInValidMappingRequest
	GetDeleteCommand() *DeleteStandardInValidMappingRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteStandardInValidMappingRequest
	GetOpTenantId() *int64
}

type DeleteStandardInValidMappingRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteStandardInValidMappingRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardInValidMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardInValidMappingRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardInValidMappingRequest) GetDeleteCommand() *DeleteStandardInValidMappingRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteStandardInValidMappingRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardInValidMappingRequest) SetDeleteCommand(v *DeleteStandardInValidMappingRequestDeleteCommand) *DeleteStandardInValidMappingRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteStandardInValidMappingRequest) SetOpTenantId(v int64) *DeleteStandardInValidMappingRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardInValidMappingRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteStandardInValidMappingRequestDeleteCommand struct {
	BelongGuidList []*string `json:"BelongGuidList,omitempty" xml:"BelongGuidList,omitempty" type:"Repeated"`
	GuidList       []*string `json:"GuidList,omitempty" xml:"GuidList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
}

func (s DeleteStandardInValidMappingRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardInValidMappingRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteStandardInValidMappingRequestDeleteCommand) GetBelongGuidList() []*string {
	return s.BelongGuidList
}

func (s *DeleteStandardInValidMappingRequestDeleteCommand) GetGuidList() []*string {
	return s.GuidList
}

func (s *DeleteStandardInValidMappingRequestDeleteCommand) GetStandardId() *int64 {
	return s.StandardId
}

func (s *DeleteStandardInValidMappingRequestDeleteCommand) SetBelongGuidList(v []*string) *DeleteStandardInValidMappingRequestDeleteCommand {
	s.BelongGuidList = v
	return s
}

func (s *DeleteStandardInValidMappingRequestDeleteCommand) SetGuidList(v []*string) *DeleteStandardInValidMappingRequestDeleteCommand {
	s.GuidList = v
	return s
}

func (s *DeleteStandardInValidMappingRequestDeleteCommand) SetStandardId(v int64) *DeleteStandardInValidMappingRequestDeleteCommand {
	s.StandardId = &v
	return s
}

func (s *DeleteStandardInValidMappingRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
