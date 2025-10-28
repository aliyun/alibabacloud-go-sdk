// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetEndpointSwitchTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetEndpointSwitchTaskResponseBody
	GetCode() *string
	SetData(v *GetEndpointSwitchTaskResponseBodyData) *GetEndpointSwitchTaskResponseBody
	GetData() *GetEndpointSwitchTaskResponseBodyData
	SetMessage(v string) *GetEndpointSwitchTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetEndpointSwitchTaskResponseBody
	GetRequestId() *string
	SetSuccess(v string) *GetEndpointSwitchTaskResponseBody
	GetSuccess() *string
	SetSynchro(v string) *GetEndpointSwitchTaskResponseBody
	GetSynchro() *string
}

type GetEndpointSwitchTaskResponseBody struct {
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
	Data      *GetEndpointSwitchTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *string                                `json:"Success,omitempty" xml:"Success,omitempty"`
	Synchro   *string                                `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s GetEndpointSwitchTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetEndpointSwitchTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetEndpointSwitchTaskResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetEndpointSwitchTaskResponseBody) GetData() *GetEndpointSwitchTaskResponseBodyData {
	return s.Data
}

func (s *GetEndpointSwitchTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetEndpointSwitchTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetEndpointSwitchTaskResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *GetEndpointSwitchTaskResponseBody) GetSynchro() *string {
	return s.Synchro
}

func (s *GetEndpointSwitchTaskResponseBody) SetCode(v string) *GetEndpointSwitchTaskResponseBody {
	s.Code = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetData(v *GetEndpointSwitchTaskResponseBodyData) *GetEndpointSwitchTaskResponseBody {
	s.Data = v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetMessage(v string) *GetEndpointSwitchTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetRequestId(v string) *GetEndpointSwitchTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetSuccess(v string) *GetEndpointSwitchTaskResponseBody {
	s.Success = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) SetSynchro(v string) *GetEndpointSwitchTaskResponseBody {
	s.Synchro = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetEndpointSwitchTaskResponseBodyData struct {
	AccountId *string `json:"AccountId,omitempty" xml:"AccountId,omitempty"`
	DbLinkId  *int64  `json:"DbLinkId,omitempty" xml:"DbLinkId,omitempty"`
	ErrMsg    *string `json:"ErrMsg,omitempty" xml:"ErrMsg,omitempty"`
	OriUuid   *string `json:"OriUuid,omitempty" xml:"OriUuid,omitempty"`
	Status    *string `json:"Status,omitempty" xml:"Status,omitempty"`
	TaskId    *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
	Uuid      *string `json:"Uuid,omitempty" xml:"Uuid,omitempty"`
}

func (s GetEndpointSwitchTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetEndpointSwitchTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetEndpointSwitchTaskResponseBodyData) GetAccountId() *string {
	return s.AccountId
}

func (s *GetEndpointSwitchTaskResponseBodyData) GetDbLinkId() *int64 {
	return s.DbLinkId
}

func (s *GetEndpointSwitchTaskResponseBodyData) GetErrMsg() *string {
	return s.ErrMsg
}

func (s *GetEndpointSwitchTaskResponseBodyData) GetOriUuid() *string {
	return s.OriUuid
}

func (s *GetEndpointSwitchTaskResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *GetEndpointSwitchTaskResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *GetEndpointSwitchTaskResponseBodyData) GetUuid() *string {
	return s.Uuid
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetAccountId(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.AccountId = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetDbLinkId(v int64) *GetEndpointSwitchTaskResponseBodyData {
	s.DbLinkId = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetErrMsg(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.ErrMsg = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetOriUuid(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.OriUuid = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetStatus(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.Status = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetTaskId(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) SetUuid(v string) *GetEndpointSwitchTaskResponseBodyData {
	s.Uuid = &v
	return s
}

func (s *GetEndpointSwitchTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
