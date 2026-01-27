// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBTablesRecoveryBackupSetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDBTablesRecoveryBackupSetResponseBody
	GetCode() *string
	SetData(v string) *DescribeDBTablesRecoveryBackupSetResponseBody
	GetData() *string
	SetErrCode(v string) *DescribeDBTablesRecoveryBackupSetResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDBTablesRecoveryBackupSetResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeDBTablesRecoveryBackupSetResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDBTablesRecoveryBackupSetResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDBTablesRecoveryBackupSetResponseBody
	GetSuccess() *string
}

type DescribeDBTablesRecoveryBackupSetResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBTablesRecoveryBackupSetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBTablesRecoveryBackupSetResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetCode(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetData(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetErrCode(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetErrMessage(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetMessage(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetRequestId(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) SetSuccess(v string) *DescribeDBTablesRecoveryBackupSetResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDBTablesRecoveryBackupSetResponseBody) Validate() error {
	return dara.Validate(s)
}
