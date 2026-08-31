// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataAssetsGovernObjectRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCommand(v *GetDataAssetsGovernObjectRequestCommand) *GetDataAssetsGovernObjectRequest
	GetCommand() *GetDataAssetsGovernObjectRequestCommand
	SetOpTenantId(v int64) *GetDataAssetsGovernObjectRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetDataAssetsGovernObjectRequest
	GetOpUserId() *string
}

type GetDataAssetsGovernObjectRequest struct {
	// The query instruction.
	//
	// This parameter is required.
	Command *GetDataAssetsGovernObjectRequestCommand `json:"Command,omitempty" xml:"Command,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// The ID of the operation user.
	//
	// example:
	//
	// 30001011
	OpUserId *string `json:"OpUserId,omitempty" xml:"OpUserId,omitempty"`
}

func (s GetDataAssetsGovernObjectRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataAssetsGovernObjectRequest) GoString() string {
	return s.String()
}

func (s *GetDataAssetsGovernObjectRequest) GetCommand() *GetDataAssetsGovernObjectRequestCommand {
	return s.Command
}

func (s *GetDataAssetsGovernObjectRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetDataAssetsGovernObjectRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetDataAssetsGovernObjectRequest) SetCommand(v *GetDataAssetsGovernObjectRequestCommand) *GetDataAssetsGovernObjectRequest {
	s.Command = v
	return s
}

func (s *GetDataAssetsGovernObjectRequest) SetOpTenantId(v int64) *GetDataAssetsGovernObjectRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetDataAssetsGovernObjectRequest) SetOpUserId(v string) *GetDataAssetsGovernObjectRequest {
	s.OpUserId = &v
	return s
}

func (s *GetDataAssetsGovernObjectRequest) Validate() error {
	if s.Command != nil {
		if err := s.Command.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetDataAssetsGovernObjectRequestCommand struct {
	// The governance object ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 22004
	GovernObjectId *int64 `json:"GovernObjectId,omitempty" xml:"GovernObjectId,omitempty"`
}

func (s GetDataAssetsGovernObjectRequestCommand) String() string {
	return dara.Prettify(s)
}

func (s GetDataAssetsGovernObjectRequestCommand) GoString() string {
	return s.String()
}

func (s *GetDataAssetsGovernObjectRequestCommand) GetGovernObjectId() *int64 {
	return s.GovernObjectId
}

func (s *GetDataAssetsGovernObjectRequestCommand) SetGovernObjectId(v int64) *GetDataAssetsGovernObjectRequestCommand {
	s.GovernObjectId = &v
	return s
}

func (s *GetDataAssetsGovernObjectRequestCommand) Validate() error {
	return dara.Validate(s)
}
