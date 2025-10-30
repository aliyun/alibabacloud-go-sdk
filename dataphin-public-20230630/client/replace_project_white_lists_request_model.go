// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceProjectWhiteListsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *ReplaceProjectWhiteListsRequest
	GetId() *int64
	SetOpTenantId(v int64) *ReplaceProjectWhiteListsRequest
	GetOpTenantId() *int64
	SetReplaceCommand(v *ReplaceProjectWhiteListsRequestReplaceCommand) *ReplaceProjectWhiteListsRequest
	GetReplaceCommand() *ReplaceProjectWhiteListsRequestReplaceCommand
}

type ReplaceProjectWhiteListsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1030111021
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
	// This parameter is required.
	ReplaceCommand *ReplaceProjectWhiteListsRequestReplaceCommand `json:"ReplaceCommand,omitempty" xml:"ReplaceCommand,omitempty" type:"Struct"`
}

func (s ReplaceProjectWhiteListsRequest) String() string {
	return dara.Prettify(s)
}

func (s ReplaceProjectWhiteListsRequest) GoString() string {
	return s.String()
}

func (s *ReplaceProjectWhiteListsRequest) GetId() *int64 {
	return s.Id
}

func (s *ReplaceProjectWhiteListsRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *ReplaceProjectWhiteListsRequest) GetReplaceCommand() *ReplaceProjectWhiteListsRequestReplaceCommand {
	return s.ReplaceCommand
}

func (s *ReplaceProjectWhiteListsRequest) SetId(v int64) *ReplaceProjectWhiteListsRequest {
	s.Id = &v
	return s
}

func (s *ReplaceProjectWhiteListsRequest) SetOpTenantId(v int64) *ReplaceProjectWhiteListsRequest {
	s.OpTenantId = &v
	return s
}

func (s *ReplaceProjectWhiteListsRequest) SetReplaceCommand(v *ReplaceProjectWhiteListsRequestReplaceCommand) *ReplaceProjectWhiteListsRequest {
	s.ReplaceCommand = v
	return s
}

func (s *ReplaceProjectWhiteListsRequest) Validate() error {
	if s.ReplaceCommand != nil {
		if err := s.ReplaceCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ReplaceProjectWhiteListsRequestReplaceCommand struct {
	// This parameter is required.
	WhiteLists []*ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists `json:"WhiteLists,omitempty" xml:"WhiteLists,omitempty" type:"Repeated"`
}

func (s ReplaceProjectWhiteListsRequestReplaceCommand) String() string {
	return dara.Prettify(s)
}

func (s ReplaceProjectWhiteListsRequestReplaceCommand) GoString() string {
	return s.String()
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommand) GetWhiteLists() []*ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists {
	return s.WhiteLists
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommand) SetWhiteLists(v []*ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) *ReplaceProjectWhiteListsRequestReplaceCommand {
	s.WhiteLists = v
	return s
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommand) Validate() error {
	if s.WhiteLists != nil {
		for _, item := range s.WhiteLists {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists struct {
	// example:
	//
	// 测试
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// ip
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.1.0.2
	Ip *string `json:"Ip,omitempty" xml:"Ip,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 5432
	Port *string `json:"Port,omitempty" xml:"Port,omitempty"`
}

func (s ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) String() string {
	return dara.Prettify(s)
}

func (s ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) GoString() string {
	return s.String()
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) GetDescription() *string {
	return s.Description
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) GetIp() *string {
	return s.Ip
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) GetPort() *string {
	return s.Port
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) SetDescription(v string) *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists {
	s.Description = &v
	return s
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) SetIp(v string) *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists {
	s.Ip = &v
	return s
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) SetPort(v string) *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists {
	s.Port = &v
	return s
}

func (s *ReplaceProjectWhiteListsRequestReplaceCommandWhiteLists) Validate() error {
	return dara.Validate(s)
}
