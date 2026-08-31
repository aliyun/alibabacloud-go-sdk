// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAssetAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetOpTenantId(v int64) *GetAssetAttributesRequest
	GetOpTenantId() *int64
	SetOpUserId(v string) *GetAssetAttributesRequest
	GetOpUserId() *string
	SetQueryCommand(v *GetAssetAttributesRequestQueryCommand) *GetAssetAttributesRequest
	GetQueryCommand() *GetAssetAttributesRequestQueryCommand
}

type GetAssetAttributesRequest struct {
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
	// The query instruction.
	//
	// This parameter is required.
	QueryCommand *GetAssetAttributesRequestQueryCommand `json:"QueryCommand,omitempty" xml:"QueryCommand,omitempty" type:"Struct"`
}

func (s GetAssetAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAssetAttributesRequest) GoString() string {
	return s.String()
}

func (s *GetAssetAttributesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *GetAssetAttributesRequest) GetOpUserId() *string {
	return s.OpUserId
}

func (s *GetAssetAttributesRequest) GetQueryCommand() *GetAssetAttributesRequestQueryCommand {
	return s.QueryCommand
}

func (s *GetAssetAttributesRequest) SetOpTenantId(v int64) *GetAssetAttributesRequest {
	s.OpTenantId = &v
	return s
}

func (s *GetAssetAttributesRequest) SetOpUserId(v string) *GetAssetAttributesRequest {
	s.OpUserId = &v
	return s
}

func (s *GetAssetAttributesRequest) SetQueryCommand(v *GetAssetAttributesRequestQueryCommand) *GetAssetAttributesRequest {
	s.QueryCommand = v
	return s
}

func (s *GetAssetAttributesRequest) Validate() error {
	if s.QueryCommand != nil {
		if err := s.QueryCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAssetAttributesRequestQueryCommand struct {
	// The list of property codes to return. If this parameter is not specified, all custom properties of the asset are returned.
	AttributeCodeList []*string `json:"AttributeCodeList,omitempty" xml:"AttributeCodeList,omitempty" type:"Repeated"`
	// The list of asset GUIDs. A maximum of 50 GUIDs are supported.
	//
	// This parameter is required.
	GuidList []*string `json:"GuidList,omitempty" xml:"GuidList,omitempty" type:"Repeated"`
}

func (s GetAssetAttributesRequestQueryCommand) String() string {
	return dara.Prettify(s)
}

func (s GetAssetAttributesRequestQueryCommand) GoString() string {
	return s.String()
}

func (s *GetAssetAttributesRequestQueryCommand) GetAttributeCodeList() []*string {
	return s.AttributeCodeList
}

func (s *GetAssetAttributesRequestQueryCommand) GetGuidList() []*string {
	return s.GuidList
}

func (s *GetAssetAttributesRequestQueryCommand) SetAttributeCodeList(v []*string) *GetAssetAttributesRequestQueryCommand {
	s.AttributeCodeList = v
	return s
}

func (s *GetAssetAttributesRequestQueryCommand) SetGuidList(v []*string) *GetAssetAttributesRequestQueryCommand {
	s.GuidList = v
	return s
}

func (s *GetAssetAttributesRequestQueryCommand) Validate() error {
	return dara.Validate(s)
}
