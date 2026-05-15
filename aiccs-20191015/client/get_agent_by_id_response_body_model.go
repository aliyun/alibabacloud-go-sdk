// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAgentByIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetAgentByIdResponseBody
	GetCode() *string
	SetData(v *GetAgentByIdResponseBodyData) *GetAgentByIdResponseBody
	GetData() *GetAgentByIdResponseBodyData
	SetMessage(v string) *GetAgentByIdResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetAgentByIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetAgentByIdResponseBody
	GetSuccess() *bool
}

type GetAgentByIdResponseBody struct {
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetAgentByIdResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool                         `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetAgentByIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAgentByIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetAgentByIdResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetAgentByIdResponseBody) GetData() *GetAgentByIdResponseBodyData {
	return s.Data
}

func (s *GetAgentByIdResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetAgentByIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAgentByIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetAgentByIdResponseBody) SetCode(v string) *GetAgentByIdResponseBody {
	s.Code = &v
	return s
}

func (s *GetAgentByIdResponseBody) SetData(v *GetAgentByIdResponseBodyData) *GetAgentByIdResponseBody {
	s.Data = v
	return s
}

func (s *GetAgentByIdResponseBody) SetMessage(v string) *GetAgentByIdResponseBody {
	s.Message = &v
	return s
}

func (s *GetAgentByIdResponseBody) SetRequestId(v string) *GetAgentByIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAgentByIdResponseBody) SetSuccess(v bool) *GetAgentByIdResponseBody {
	s.Success = &v
	return s
}

func (s *GetAgentByIdResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetAgentByIdResponseBodyData struct {
	AgentId        *int32  `json:"AgentId,omitempty" xml:"AgentId,omitempty"`
	CreateUserName *string `json:"CreateUserName,omitempty" xml:"CreateUserName,omitempty"`
	ForeignKey     *string `json:"ForeignKey,omitempty" xml:"ForeignKey,omitempty"`
	ForeignNick    *string `json:"ForeignNick,omitempty" xml:"ForeignNick,omitempty"`
	RealName       *string `json:"RealName,omitempty" xml:"RealName,omitempty"`
	ServicerType   *int32  `json:"ServicerType,omitempty" xml:"ServicerType,omitempty"`
	ShowName       *string `json:"ShowName,omitempty" xml:"ShowName,omitempty"`
}

func (s GetAgentByIdResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetAgentByIdResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetAgentByIdResponseBodyData) GetAgentId() *int32 {
	return s.AgentId
}

func (s *GetAgentByIdResponseBodyData) GetCreateUserName() *string {
	return s.CreateUserName
}

func (s *GetAgentByIdResponseBodyData) GetForeignKey() *string {
	return s.ForeignKey
}

func (s *GetAgentByIdResponseBodyData) GetForeignNick() *string {
	return s.ForeignNick
}

func (s *GetAgentByIdResponseBodyData) GetRealName() *string {
	return s.RealName
}

func (s *GetAgentByIdResponseBodyData) GetServicerType() *int32 {
	return s.ServicerType
}

func (s *GetAgentByIdResponseBodyData) GetShowName() *string {
	return s.ShowName
}

func (s *GetAgentByIdResponseBodyData) SetAgentId(v int32) *GetAgentByIdResponseBodyData {
	s.AgentId = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetCreateUserName(v string) *GetAgentByIdResponseBodyData {
	s.CreateUserName = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetForeignKey(v string) *GetAgentByIdResponseBodyData {
	s.ForeignKey = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetForeignNick(v string) *GetAgentByIdResponseBodyData {
	s.ForeignNick = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetRealName(v string) *GetAgentByIdResponseBodyData {
	s.RealName = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetServicerType(v int32) *GetAgentByIdResponseBodyData {
	s.ServicerType = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) SetShowName(v string) *GetAgentByIdResponseBodyData {
	s.ShowName = &v
	return s
}

func (s *GetAgentByIdResponseBodyData) Validate() error {
	return dara.Validate(s)
}
