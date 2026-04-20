// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMmAppBindingMcpRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *MmAppBindingMcpRequest
	GetAppId() *string
	SetMcps(v []*MmAppBindingMcpRequestMcps) *MmAppBindingMcpRequest
	GetMcps() []*MmAppBindingMcpRequestMcps
	SetWorkspaceId(v string) *MmAppBindingMcpRequest
	GetWorkspaceId() *string
}

type MmAppBindingMcpRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_a2eb4e04b48041108edb1f6de815
	AppId *string                       `json:"AppId,omitempty" xml:"AppId,omitempty"`
	Mcps  []*MmAppBindingMcpRequestMcps `json:"Mcps,omitempty" xml:"Mcps,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// llm-o8ixktz41iyd2b6p
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s MmAppBindingMcpRequest) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingMcpRequest) GoString() string {
	return s.String()
}

func (s *MmAppBindingMcpRequest) GetAppId() *string {
	return s.AppId
}

func (s *MmAppBindingMcpRequest) GetMcps() []*MmAppBindingMcpRequestMcps {
	return s.Mcps
}

func (s *MmAppBindingMcpRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *MmAppBindingMcpRequest) SetAppId(v string) *MmAppBindingMcpRequest {
	s.AppId = &v
	return s
}

func (s *MmAppBindingMcpRequest) SetMcps(v []*MmAppBindingMcpRequestMcps) *MmAppBindingMcpRequest {
	s.Mcps = v
	return s
}

func (s *MmAppBindingMcpRequest) SetWorkspaceId(v string) *MmAppBindingMcpRequest {
	s.WorkspaceId = &v
	return s
}

func (s *MmAppBindingMcpRequest) Validate() error {
	if s.Mcps != nil {
		for _, item := range s.Mcps {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type MmAppBindingMcpRequestMcps struct {
	// example:
	//
	// mcp-ZDI1MDU2ZTExZGZh
	Code     *string   `json:"Code,omitempty" xml:"Code,omitempty"`
	ToolList []*string `json:"ToolList,omitempty" xml:"ToolList,omitempty" type:"Repeated"`
	// example:
	//
	// CUSTOM
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s MmAppBindingMcpRequestMcps) String() string {
	return dara.Prettify(s)
}

func (s MmAppBindingMcpRequestMcps) GoString() string {
	return s.String()
}

func (s *MmAppBindingMcpRequestMcps) GetCode() *string {
	return s.Code
}

func (s *MmAppBindingMcpRequestMcps) GetToolList() []*string {
	return s.ToolList
}

func (s *MmAppBindingMcpRequestMcps) GetType() *string {
	return s.Type
}

func (s *MmAppBindingMcpRequestMcps) SetCode(v string) *MmAppBindingMcpRequestMcps {
	s.Code = &v
	return s
}

func (s *MmAppBindingMcpRequestMcps) SetToolList(v []*string) *MmAppBindingMcpRequestMcps {
	s.ToolList = v
	return s
}

func (s *MmAppBindingMcpRequestMcps) SetType(v string) *MmAppBindingMcpRequestMcps {
	s.Type = &v
	return s
}

func (s *MmAppBindingMcpRequestMcps) Validate() error {
	return dara.Validate(s)
}
