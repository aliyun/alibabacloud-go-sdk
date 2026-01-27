// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBTablesRecoveryStateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDBTablesRecoveryStateResponseBody
	GetCode() *string
	SetData(v string) *DescribeDBTablesRecoveryStateResponseBody
	GetData() *string
	SetErrCode(v string) *DescribeDBTablesRecoveryStateResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDBTablesRecoveryStateResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeDBTablesRecoveryStateResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDBTablesRecoveryStateResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDBTablesRecoveryStateResponseBody
	GetSuccess() *string
}

type DescribeDBTablesRecoveryStateResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBTablesRecoveryStateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBTablesRecoveryStateResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryStateResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDBTablesRecoveryStateResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeDBTablesRecoveryStateResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDBTablesRecoveryStateResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDBTablesRecoveryStateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDBTablesRecoveryStateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBTablesRecoveryStateResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetCode(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetData(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetErrCode(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetErrMessage(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetMessage(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetRequestId(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) SetSuccess(v string) *DescribeDBTablesRecoveryStateResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDBTablesRecoveryStateResponseBody) Validate() error {
	return dara.Validate(s)
}
