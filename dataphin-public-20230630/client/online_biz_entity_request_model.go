// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOnlineBizEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOnlineCommand(v *OnlineBizEntityRequestOnlineCommand) *OnlineBizEntityRequest
	GetOnlineCommand() *OnlineBizEntityRequestOnlineCommand
	SetOpTenantId(v int64) *OnlineBizEntityRequest
	GetOpTenantId() *int64
}

type OnlineBizEntityRequest struct {
	// The online request.
	//
	// This parameter is required.
	OnlineCommand *OnlineBizEntityRequestOnlineCommand `json:"OnlineCommand,omitempty" xml:"OnlineCommand,omitempty" type:"Struct"`
	// The tenant ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s OnlineBizEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s OnlineBizEntityRequest) GoString() string {
	return s.String()
}

func (s *OnlineBizEntityRequest) GetOnlineCommand() *OnlineBizEntityRequestOnlineCommand {
	return s.OnlineCommand
}

func (s *OnlineBizEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OnlineBizEntityRequest) SetOnlineCommand(v *OnlineBizEntityRequestOnlineCommand) *OnlineBizEntityRequest {
	s.OnlineCommand = v
	return s
}

func (s *OnlineBizEntityRequest) SetOpTenantId(v int64) *OnlineBizEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *OnlineBizEntityRequest) Validate() error {
	if s.OnlineCommand != nil {
		if err := s.OnlineCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OnlineBizEntityRequestOnlineCommand struct {
	// The ID of the business unit to which the entity belongs.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6798087749072704
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// The remarks for the offline operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// The ID of the business entity.
	//
	// This parameter is required.
	//
	// example:
	//
	// 101001201
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The business type. Valid values:
	//
	// - BIZ_OBJECT
	//
	// - BIZ_PROCESS.
	//
	// This parameter is required.
	//
	// example:
	//
	// BIZ_OBJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OnlineBizEntityRequestOnlineCommand) String() string {
	return dara.Prettify(s)
}

func (s OnlineBizEntityRequestOnlineCommand) GoString() string {
	return s.String()
}

func (s *OnlineBizEntityRequestOnlineCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *OnlineBizEntityRequestOnlineCommand) GetComment() *string {
	return s.Comment
}

func (s *OnlineBizEntityRequestOnlineCommand) GetId() *int64 {
	return s.Id
}

func (s *OnlineBizEntityRequestOnlineCommand) GetType() *string {
	return s.Type
}

func (s *OnlineBizEntityRequestOnlineCommand) SetBizUnitId(v int64) *OnlineBizEntityRequestOnlineCommand {
	s.BizUnitId = &v
	return s
}

func (s *OnlineBizEntityRequestOnlineCommand) SetComment(v string) *OnlineBizEntityRequestOnlineCommand {
	s.Comment = &v
	return s
}

func (s *OnlineBizEntityRequestOnlineCommand) SetId(v int64) *OnlineBizEntityRequestOnlineCommand {
	s.Id = &v
	return s
}

func (s *OnlineBizEntityRequestOnlineCommand) SetType(v string) *OnlineBizEntityRequestOnlineCommand {
	s.Type = &v
	return s
}

func (s *OnlineBizEntityRequestOnlineCommand) Validate() error {
	return dara.Validate(s)
}
