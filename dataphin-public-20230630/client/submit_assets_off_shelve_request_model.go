// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAssetsOffShelveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitAssetsOffShelveRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitAssetsOffShelveRequest
	GetOpUserId() *string
	SetSubmitCommand(v *SubmitAssetsOffShelveRequestSubmitCommand) *SubmitAssetsOffShelveRequest
	GetSubmitCommand() *SubmitAssetsOffShelveRequestSubmitCommand
}

type SubmitAssetsOffShelveRequest struct {
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operator user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
	// The delisting submit command.
	//
	// This parameter is required.
	SubmitCommand *SubmitAssetsOffShelveRequestSubmitCommand `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty" type:"Struct"`
}

func (s SubmitAssetsOffShelveRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOffShelveRequest) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOffShelveRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitAssetsOffShelveRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitAssetsOffShelveRequest) GetSubmitCommand() *SubmitAssetsOffShelveRequestSubmitCommand {
	return s.SubmitCommand
}

func (s *SubmitAssetsOffShelveRequest) SetOpTenantId(v int64) *SubmitAssetsOffShelveRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitAssetsOffShelveRequest) SetOpUserId(v string) *SubmitAssetsOffShelveRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitAssetsOffShelveRequest) SetSubmitCommand(v *SubmitAssetsOffShelveRequestSubmitCommand) *SubmitAssetsOffShelveRequest {
	s.SubmitCommand = v
	return s
}

func (s *SubmitAssetsOffShelveRequest) Validate() error {
	if s.SubmitCommand != nil {
		if err := s.SubmitCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAssetsOffShelveRequestSubmitCommand struct {
	// The list of asset GUIDs to be delisted. A maximum of 50 GUIDs can be specified per request.
	//
	// This parameter is required.
	GuidList []*string `json:"GuidList,omitempty" xml:"GuidList,omitempty" type:"Repeated"`
	// The delisting description. The value must be 1 to 100 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// Business adjustment, no longer available externally
	OffShelveDescription *string `json:"OffShelveDescription,omitempty" xml:"OffShelveDescription,omitempty"`
}

func (s SubmitAssetsOffShelveRequestSubmitCommand) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOffShelveRequestSubmitCommand) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOffShelveRequestSubmitCommand) GetGuidList() []*string {
	return s.GuidList
}

func (s *SubmitAssetsOffShelveRequestSubmitCommand) GetOffShelveDescription() *string {
	return s.OffShelveDescription
}

func (s *SubmitAssetsOffShelveRequestSubmitCommand) SetGuidList(v []*string) *SubmitAssetsOffShelveRequestSubmitCommand {
	s.GuidList = v
	return s
}

func (s *SubmitAssetsOffShelveRequestSubmitCommand) SetOffShelveDescription(v string) *SubmitAssetsOffShelveRequestSubmitCommand {
	s.OffShelveDescription = &v
	return s
}

func (s *SubmitAssetsOffShelveRequestSubmitCommand) Validate() error {
	return dara.Validate(s)
}
