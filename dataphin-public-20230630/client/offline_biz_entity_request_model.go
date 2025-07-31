// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineBizEntityRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfflineCommand(v *OfflineBizEntityRequestOfflineCommand) *OfflineBizEntityRequest
	GetOfflineCommand() *OfflineBizEntityRequestOfflineCommand
	SetOpTenantId(v int64) *OfflineBizEntityRequest
	GetOpTenantId() *int64
}

type OfflineBizEntityRequest struct {
	// This parameter is required.
	OfflineCommand *OfflineBizEntityRequestOfflineCommand `json:"OfflineCommand,omitempty" xml:"OfflineCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s OfflineBizEntityRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineBizEntityRequest) GoString() string {
	return s.String()
}

func (s *OfflineBizEntityRequest) GetOfflineCommand() *OfflineBizEntityRequestOfflineCommand {
	return s.OfflineCommand
}

func (s *OfflineBizEntityRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OfflineBizEntityRequest) SetOfflineCommand(v *OfflineBizEntityRequestOfflineCommand) *OfflineBizEntityRequest {
	s.OfflineCommand = v
	return s
}

func (s *OfflineBizEntityRequest) SetOpTenantId(v int64) *OfflineBizEntityRequest {
	s.OpTenantId = &v
	return s
}

func (s *OfflineBizEntityRequest) Validate() error {
	return dara.Validate(s)
}

type OfflineBizEntityRequestOfflineCommand struct {
	// This parameter is required.
	//
	// example:
	//
	// 6798087749072704
	BizUnitId *int64 `json:"BizUnitId,omitempty" xml:"BizUnitId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// test
	Comment *string `json:"Comment,omitempty" xml:"Comment,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 101001201
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// BIZ_OBJECT
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s OfflineBizEntityRequestOfflineCommand) String() string {
	return dara.Prettify(s)
}

func (s OfflineBizEntityRequestOfflineCommand) GoString() string {
	return s.String()
}

func (s *OfflineBizEntityRequestOfflineCommand) GetBizUnitId() *int64 {
	return s.BizUnitId
}

func (s *OfflineBizEntityRequestOfflineCommand) GetComment() *string {
	return s.Comment
}

func (s *OfflineBizEntityRequestOfflineCommand) GetId() *int64 {
	return s.Id
}

func (s *OfflineBizEntityRequestOfflineCommand) GetType() *string {
	return s.Type
}

func (s *OfflineBizEntityRequestOfflineCommand) SetBizUnitId(v int64) *OfflineBizEntityRequestOfflineCommand {
	s.BizUnitId = &v
	return s
}

func (s *OfflineBizEntityRequestOfflineCommand) SetComment(v string) *OfflineBizEntityRequestOfflineCommand {
	s.Comment = &v
	return s
}

func (s *OfflineBizEntityRequestOfflineCommand) SetId(v int64) *OfflineBizEntityRequestOfflineCommand {
	s.Id = &v
	return s
}

func (s *OfflineBizEntityRequestOfflineCommand) SetType(v string) *OfflineBizEntityRequestOfflineCommand {
	s.Type = &v
	return s
}

func (s *OfflineBizEntityRequestOfflineCommand) Validate() error {
	return dara.Validate(s)
}
