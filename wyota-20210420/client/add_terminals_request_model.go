// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTerminalsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAddTerminalParams(v []*AddTerminalsRequestAddTerminalParams) *AddTerminalsRequest
	GetAddTerminalParams() []*AddTerminalsRequestAddTerminalParams
	SetMainBizType(v string) *AddTerminalsRequest
	GetMainBizType() *string
}

type AddTerminalsRequest struct {
	AddTerminalParams []*AddTerminalsRequestAddTerminalParams `json:"AddTerminalParams,omitempty" xml:"AddTerminalParams,omitempty" type:"Repeated"`
	MainBizType       *string                                 `json:"MainBizType,omitempty" xml:"MainBizType,omitempty"`
}

func (s AddTerminalsRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTerminalsRequest) GoString() string {
	return s.String()
}

func (s *AddTerminalsRequest) GetAddTerminalParams() []*AddTerminalsRequestAddTerminalParams {
	return s.AddTerminalParams
}

func (s *AddTerminalsRequest) GetMainBizType() *string {
	return s.MainBizType
}

func (s *AddTerminalsRequest) SetAddTerminalParams(v []*AddTerminalsRequestAddTerminalParams) *AddTerminalsRequest {
	s.AddTerminalParams = v
	return s
}

func (s *AddTerminalsRequest) SetMainBizType(v string) *AddTerminalsRequest {
	s.MainBizType = &v
	return s
}

func (s *AddTerminalsRequest) Validate() error {
	if s.AddTerminalParams != nil {
		for _, item := range s.AddTerminalParams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type AddTerminalsRequestAddTerminalParams struct {
	Alias           *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	ClientType      *int32  `json:"ClientType,omitempty" xml:"ClientType,omitempty"`
	SerialNumber    *string `json:"SerialNumber,omitempty" xml:"SerialNumber,omitempty"`
	TerminalGroupId *string `json:"TerminalGroupId,omitempty" xml:"TerminalGroupId,omitempty"`
	Uuid            *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s AddTerminalsRequestAddTerminalParams) String() string {
	return dara.Prettify(s)
}

func (s AddTerminalsRequestAddTerminalParams) GoString() string {
	return s.String()
}

func (s *AddTerminalsRequestAddTerminalParams) GetAlias() *string {
	return s.Alias
}

func (s *AddTerminalsRequestAddTerminalParams) GetClientType() *int32 {
	return s.ClientType
}

func (s *AddTerminalsRequestAddTerminalParams) GetSerialNumber() *string {
	return s.SerialNumber
}

func (s *AddTerminalsRequestAddTerminalParams) GetTerminalGroupId() *string {
	return s.TerminalGroupId
}

func (s *AddTerminalsRequestAddTerminalParams) GetUuid() *string {
	return s.Uuid
}

func (s *AddTerminalsRequestAddTerminalParams) SetAlias(v string) *AddTerminalsRequestAddTerminalParams {
	s.Alias = &v
	return s
}

func (s *AddTerminalsRequestAddTerminalParams) SetClientType(v int32) *AddTerminalsRequestAddTerminalParams {
	s.ClientType = &v
	return s
}

func (s *AddTerminalsRequestAddTerminalParams) SetSerialNumber(v string) *AddTerminalsRequestAddTerminalParams {
	s.SerialNumber = &v
	return s
}

func (s *AddTerminalsRequestAddTerminalParams) SetTerminalGroupId(v string) *AddTerminalsRequestAddTerminalParams {
	s.TerminalGroupId = &v
	return s
}

func (s *AddTerminalsRequestAddTerminalParams) SetUuid(v string) *AddTerminalsRequestAddTerminalParams {
	s.Uuid = &v
	return s
}

func (s *AddTerminalsRequestAddTerminalParams) Validate() error {
	return dara.Validate(s)
}
