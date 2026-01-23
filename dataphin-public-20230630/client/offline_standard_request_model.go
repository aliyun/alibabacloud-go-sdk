// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineStandardRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOfflineCommand(v *OfflineStandardRequestOfflineCommand) *OfflineStandardRequest
	GetOfflineCommand() *OfflineStandardRequestOfflineCommand
	SetOpTenantId(v int64) *OfflineStandardRequest
	GetOpTenantId() *int64
}

type OfflineStandardRequest struct {
	// This parameter is required.
	OfflineCommand *OfflineStandardRequestOfflineCommand `json:"OfflineCommand,omitempty" xml:"OfflineCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s OfflineStandardRequest) String() string {
	return dara.Prettify(s)
}

func (s OfflineStandardRequest) GoString() string {
	return s.String()
}

func (s *OfflineStandardRequest) GetOfflineCommand() *OfflineStandardRequestOfflineCommand {
	return s.OfflineCommand
}

func (s *OfflineStandardRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *OfflineStandardRequest) SetOfflineCommand(v *OfflineStandardRequestOfflineCommand) *OfflineStandardRequest {
	s.OfflineCommand = v
	return s
}

func (s *OfflineStandardRequest) SetOpTenantId(v int64) *OfflineStandardRequest {
	s.OpTenantId = &v
	return s
}

func (s *OfflineStandardRequest) Validate() error {
	if s.OfflineCommand != nil {
		if err := s.OfflineCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type OfflineStandardRequestOfflineCommand struct {
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
	// 1234
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s OfflineStandardRequestOfflineCommand) String() string {
	return dara.Prettify(s)
}

func (s OfflineStandardRequestOfflineCommand) GoString() string {
	return s.String()
}

func (s *OfflineStandardRequestOfflineCommand) GetComment() *string {
	return s.Comment
}

func (s *OfflineStandardRequestOfflineCommand) GetId() *int64 {
	return s.Id
}

func (s *OfflineStandardRequestOfflineCommand) SetComment(v string) *OfflineStandardRequestOfflineCommand {
	s.Comment = &v
	return s
}

func (s *OfflineStandardRequestOfflineCommand) SetId(v int64) *OfflineStandardRequestOfflineCommand {
	s.Id = &v
	return s
}

func (s *OfflineStandardRequestOfflineCommand) Validate() error {
	return dara.Validate(s)
}
