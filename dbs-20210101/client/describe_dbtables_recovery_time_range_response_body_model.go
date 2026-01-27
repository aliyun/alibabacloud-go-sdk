// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBTablesRecoveryTimeRangeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody
	GetCode() *string
	SetData(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody
	GetData() *string
	SetErrCode(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody
	GetSuccess() *string
}

type DescribeDBTablesRecoveryTimeRangeResponseBody struct {
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Data       *string `json:"Data,omitempty" xml:"Data,omitempty"`
	ErrCode    *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success    *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDBTablesRecoveryTimeRangeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBTablesRecoveryTimeRangeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetCode(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetData(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetErrCode(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetErrMessage(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetMessage(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetRequestId(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) SetSuccess(v string) *DescribeDBTablesRecoveryTimeRangeResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDBTablesRecoveryTimeRangeResponseBody) Validate() error {
	return dara.Validate(s)
}
