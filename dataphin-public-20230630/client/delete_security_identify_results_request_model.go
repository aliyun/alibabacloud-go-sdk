// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSecurityIdentifyResultsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteSecurityIdentifyResultsRequestDeleteCommand) *DeleteSecurityIdentifyResultsRequest
	GetDeleteCommand() *DeleteSecurityIdentifyResultsRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteSecurityIdentifyResultsRequest
	GetOpTenantId() *int64
}

type DeleteSecurityIdentifyResultsRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteSecurityIdentifyResultsRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteSecurityIdentifyResultsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIdentifyResultsRequest) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIdentifyResultsRequest) GetDeleteCommand() *DeleteSecurityIdentifyResultsRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteSecurityIdentifyResultsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteSecurityIdentifyResultsRequest) SetDeleteCommand(v *DeleteSecurityIdentifyResultsRequestDeleteCommand) *DeleteSecurityIdentifyResultsRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteSecurityIdentifyResultsRequest) SetOpTenantId(v int64) *DeleteSecurityIdentifyResultsRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteSecurityIdentifyResultsRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteSecurityIdentifyResultsRequestDeleteCommand struct {
	IdentifyResultIdList []*int64 `json:"IdentifyResultIdList,omitempty" xml:"IdentifyResultIdList,omitempty" type:"Repeated"`
}

func (s DeleteSecurityIdentifyResultsRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteSecurityIdentifyResultsRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteSecurityIdentifyResultsRequestDeleteCommand) GetIdentifyResultIdList() []*int64 {
	return s.IdentifyResultIdList
}

func (s *DeleteSecurityIdentifyResultsRequestDeleteCommand) SetIdentifyResultIdList(v []*int64) *DeleteSecurityIdentifyResultsRequestDeleteCommand {
	s.IdentifyResultIdList = v
	return s
}

func (s *DeleteSecurityIdentifyResultsRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
