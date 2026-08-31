// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitAssetsOnShelveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *SubmitAssetsOnShelveRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *SubmitAssetsOnShelveRequest
	GetOpUserId() *string
	SetSubmitCommand(v *SubmitAssetsOnShelveRequestSubmitCommand) *SubmitAssetsOnShelveRequest
	GetSubmitCommand() *SubmitAssetsOnShelveRequestSubmitCommand
}

type SubmitAssetsOnShelveRequest struct {
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
	// The submit listing instruction.
	//
	// This parameter is required.
	SubmitCommand *SubmitAssetsOnShelveRequestSubmitCommand `json:"SubmitCommand,omitempty" xml:"SubmitCommand,omitempty" type:"Struct"`
}

func (s SubmitAssetsOnShelveRequest) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOnShelveRequest) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOnShelveRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *SubmitAssetsOnShelveRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *SubmitAssetsOnShelveRequest) GetSubmitCommand() *SubmitAssetsOnShelveRequestSubmitCommand {
	return s.SubmitCommand
}

func (s *SubmitAssetsOnShelveRequest) SetOpTenantId(v int64) *SubmitAssetsOnShelveRequest {
	s.OpTenantId = &v
	return s
}

func (s *SubmitAssetsOnShelveRequest) SetOpUserId(v string) *SubmitAssetsOnShelveRequest {
	s.OpUserId = &v
	return s
}

func (s *SubmitAssetsOnShelveRequest) SetSubmitCommand(v *SubmitAssetsOnShelveRequestSubmitCommand) *SubmitAssetsOnShelveRequest {
	s.SubmitCommand = v
	return s
}

func (s *SubmitAssetsOnShelveRequest) Validate() error {
	if s.SubmitCommand != nil {
		if err := s.SubmitCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitAssetsOnShelveRequestSubmitCommand struct {
	// The list of asset GUIDs to be listed. A maximum of 50 GUIDs can be specified per request.
	//
	// This parameter is required.
	GuidList []*string `json:"GuidList,omitempty" xml:"GuidList,omitempty" type:"Repeated"`
}

func (s SubmitAssetsOnShelveRequestSubmitCommand) String() string {
	return dara.Prettify(s)
}

func (s SubmitAssetsOnShelveRequestSubmitCommand) GoString() string {
	return s.String()
}

func (s *SubmitAssetsOnShelveRequestSubmitCommand) GetGuidList() []*string {
	return s.GuidList
}

func (s *SubmitAssetsOnShelveRequestSubmitCommand) SetGuidList(v []*string) *SubmitAssetsOnShelveRequestSubmitCommand {
	s.GuidList = v
	return s
}

func (s *SubmitAssetsOnShelveRequestSubmitCommand) Validate() error {
	return dara.Validate(s)
}
