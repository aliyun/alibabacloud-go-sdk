// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iConfigureBackupSetInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ConfigureBackupSetInfoResponseBody
	GetCode() *string
	SetData(v *ConfigureBackupSetInfoResponseBodyData) *ConfigureBackupSetInfoResponseBody
	GetData() *ConfigureBackupSetInfoResponseBodyData
	SetErrCode(v string) *ConfigureBackupSetInfoResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *ConfigureBackupSetInfoResponseBody
	GetErrMessage() *string
	SetMessage(v string) *ConfigureBackupSetInfoResponseBody
	GetMessage() *string
	SetRequestId(v string) *ConfigureBackupSetInfoResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ConfigureBackupSetInfoResponseBody
	GetSuccess() *string
}

type ConfigureBackupSetInfoResponseBody struct {
	// example:
	//
	// Success
	Code *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ConfigureBackupSetInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Request.Forbidden
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 54A63B3B-AA10-1CC3-A6BB-6CCE98D1****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ConfigureBackupSetInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ConfigureBackupSetInfoResponseBody) GoString() string {
	return s.String()
}

func (s *ConfigureBackupSetInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *ConfigureBackupSetInfoResponseBody) GetData() *ConfigureBackupSetInfoResponseBodyData {
	return s.Data
}

func (s *ConfigureBackupSetInfoResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *ConfigureBackupSetInfoResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *ConfigureBackupSetInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ConfigureBackupSetInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ConfigureBackupSetInfoResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ConfigureBackupSetInfoResponseBody) SetCode(v string) *ConfigureBackupSetInfoResponseBody {
	s.Code = &v
	return s
}

func (s *ConfigureBackupSetInfoResponseBody) SetData(v *ConfigureBackupSetInfoResponseBodyData) *ConfigureBackupSetInfoResponseBody {
	s.Data = v
	return s
}

func (s *ConfigureBackupSetInfoResponseBody) SetErrCode(v string) *ConfigureBackupSetInfoResponseBody {
	s.ErrCode = &v
	return s
}

func (s *ConfigureBackupSetInfoResponseBody) SetErrMessage(v string) *ConfigureBackupSetInfoResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *ConfigureBackupSetInfoResponseBody) SetMessage(v string) *ConfigureBackupSetInfoResponseBody {
	s.Message = &v
	return s
}

func (s *ConfigureBackupSetInfoResponseBody) SetRequestId(v string) *ConfigureBackupSetInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *ConfigureBackupSetInfoResponseBody) SetSuccess(v string) *ConfigureBackupSetInfoResponseBody {
	s.Success = &v
	return s
}

func (s *ConfigureBackupSetInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ConfigureBackupSetInfoResponseBodyData struct {
	// example:
	//
	// ee-d84wwm3ca****
	BackupSetId *string `json:"BackupSetId,omitempty" xml:"BackupSetId,omitempty"`
}

func (s ConfigureBackupSetInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ConfigureBackupSetInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *ConfigureBackupSetInfoResponseBodyData) GetBackupSetId() *string {
	return s.BackupSetId
}

func (s *ConfigureBackupSetInfoResponseBodyData) SetBackupSetId(v string) *ConfigureBackupSetInfoResponseBodyData {
	s.BackupSetId = &v
	return s
}

func (s *ConfigureBackupSetInfoResponseBodyData) Validate() error {
	return dara.Validate(s)
}
