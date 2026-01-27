// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBTablesRecoveryStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyDBTablesRecoveryStateResponseBody
	GetCode() *string
	SetData(v string) *ModifyDBTablesRecoveryStateResponseBody
	GetData() *string
	SetErrCode(v string) *ModifyDBTablesRecoveryStateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ModifyDBTablesRecoveryStateResponseBody
	GetErrMessage() *string
	SetMessage(v string) *ModifyDBTablesRecoveryStateResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyDBTablesRecoveryStateResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyDBTablesRecoveryStateResponseBody
	GetSuccess() *string
}

type ModifyDBTablesRecoveryStateResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyDBTablesRecoveryStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBTablesRecoveryStateResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBTablesRecoveryStateResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyDBTablesRecoveryStateResponseBody) GetData() *string {
	return s.Data
}

func (s *ModifyDBTablesRecoveryStateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ModifyDBTablesRecoveryStateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ModifyDBTablesRecoveryStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyDBTablesRecoveryStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBTablesRecoveryStateResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetCode(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetData(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetErrCode(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetErrMessage(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetMessage(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetRequestId(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) SetSuccess(v string) *ModifyDBTablesRecoveryStateResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyDBTablesRecoveryStateResponseBody) Validate() error {
	return dara.Validate(s)
}
