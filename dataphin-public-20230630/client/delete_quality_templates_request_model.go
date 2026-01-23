// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteQualityTemplatesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDeleteCommand(v *DeleteQualityTemplatesRequestDeleteCommand) *DeleteQualityTemplatesRequest
	GetDeleteCommand() *DeleteQualityTemplatesRequestDeleteCommand
	SetOpTenantId(v int64) *DeleteQualityTemplatesRequest
	GetOpTenantId() *int64
}

type DeleteQualityTemplatesRequest struct {
	// This parameter is required.
	DeleteCommand *DeleteQualityTemplatesRequestDeleteCommand `json:"DeleteCommand,omitempty" xml:"DeleteCommand,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 30001011
	OpTenantId *int64 `json:"OpTenantId,omitempty" xml:"OpTenantId,omitempty"`
}

func (s DeleteQualityTemplatesRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityTemplatesRequest) GoString() string {
	return s.String()
}

func (s *DeleteQualityTemplatesRequest) GetDeleteCommand() *DeleteQualityTemplatesRequestDeleteCommand {
	return s.DeleteCommand
}

func (s *DeleteQualityTemplatesRequest) GetOpTenantId() *int64 {
	return s.OpTenantId
}

func (s *DeleteQualityTemplatesRequest) SetDeleteCommand(v *DeleteQualityTemplatesRequestDeleteCommand) *DeleteQualityTemplatesRequest {
	s.DeleteCommand = v
	return s
}

func (s *DeleteQualityTemplatesRequest) SetOpTenantId(v int64) *DeleteQualityTemplatesRequest {
	s.OpTenantId = &v
	return s
}

func (s *DeleteQualityTemplatesRequest) Validate() error {
	if s.DeleteCommand != nil {
		if err := s.DeleteCommand.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DeleteQualityTemplatesRequestDeleteCommand struct {
	// This parameter is required.
	TemplateIdList []*int64 `json:"TemplateIdList,omitempty" xml:"TemplateIdList,omitempty" type:"Repeated"`
}

func (s DeleteQualityTemplatesRequestDeleteCommand) String() string {
	return dara.Prettify(s)
}

func (s DeleteQualityTemplatesRequestDeleteCommand) GoString() string {
	return s.String()
}

func (s *DeleteQualityTemplatesRequestDeleteCommand) GetTemplateIdList() []*int64 {
	return s.TemplateIdList
}

func (s *DeleteQualityTemplatesRequestDeleteCommand) SetTemplateIdList(v []*int64) *DeleteQualityTemplatesRequestDeleteCommand {
	s.TemplateIdList = v
	return s
}

func (s *DeleteQualityTemplatesRequestDeleteCommand) Validate() error {
	return dara.Validate(s)
}
