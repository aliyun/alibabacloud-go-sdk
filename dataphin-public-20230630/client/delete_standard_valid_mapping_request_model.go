// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteStandardValidMappingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteStandardValidMappingRequestDeleteCommand) *DeleteStandardValidMappingRequest
	GetDeleteCommand() *DeleteStandardValidMappingRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteStandardValidMappingRequest
	GetOpTenantId() *int64
}

type DeleteStandardValidMappingRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteStandardValidMappingRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteStandardValidMappingRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardValidMappingRequest) GoString() string {
	return s.String()
}

func (s *DeleteStandardValidMappingRequest) GetDeleteCommand() *DeleteStandardValidMappingRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteStandardValidMappingRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteStandardValidMappingRequest) SetDeleteCommand(v *DeleteStandardValidMappingRequestDeleteCommand) *DeleteStandardValidMappingRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteStandardValidMappingRequest) SetOpTenantId(v int64) *DeleteStandardValidMappingRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteStandardValidMappingRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteStandardValidMappingRequestDeleteCommand struct {
	BelongGuidList []*string `json:"BelongGuidList,omitempty" xml:"BelongGuidList,omitempty" type:"Repeated"`
	GuidList       []*string `json:"GuidList,omitempty" xml:"GuidList,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	StandardId *int64 `json:"StandardId,omitempty" xml:"StandardId,omitempty"`
}

func (s DeleteStandardValidMappingRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteStandardValidMappingRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteStandardValidMappingRequestDeleteCommand) GetBelongGuidList() []*string {
	return s.BelongGuidList
}

func (s *DeleteStandardValidMappingRequestDeleteCommand) GetGuidList() []*string {
	return s.GuidList
}

func (s *DeleteStandardValidMappingRequestDeleteCommand) GetStandardId() *int64 {
	return s.StandardId
}

func (s *DeleteStandardValidMappingRequestDeleteCommand) SetBelongGuidList(v []*string) *DeleteStandardValidMappingRequestDeleteCommand {
	s.BelongGuidList = v
	return s
}

func (s *DeleteStandardValidMappingRequestDeleteCommand) SetGuidList(v []*string) *DeleteStandardValidMappingRequestDeleteCommand {
	s.GuidList = v
	return s
}

func (s *DeleteStandardValidMappingRequestDeleteCommand) SetStandardId(v int64) *DeleteStandardValidMappingRequestDeleteCommand {
	s.StandardId = &v
	return s
}

func (s *DeleteStandardValidMappingRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
