// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCostInfoByDbsInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeCostInfoByDbsInstanceResponseBody
	GetCode() *string
	SetData(v string) *DescribeCostInfoByDbsInstanceResponseBody
	GetData() *string
	SetErrCode(v string) *DescribeCostInfoByDbsInstanceResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeCostInfoByDbsInstanceResponseBody
	GetErrMessage() *string
	SetMessage(v string) *DescribeCostInfoByDbsInstanceResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeCostInfoByDbsInstanceResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeCostInfoByDbsInstanceResponseBody
	GetSuccess() *string
}

type DescribeCostInfoByDbsInstanceResponseBody struct {
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// {
	//
	//       "backupPlanComment": "",
	//
	//       "commodity": "cbs_post",
	//
	//       "product": "cbs",
	//
	//       "moduleCode": "BackupStorageSize",
	//
	//       "instanceName": "d-2zefd6337d766294",
	//
	//       "backupPlanId": "dbs:d-2zefd6337d766294",
	//
	//       "moduleName": "mongodb"
	//
	//     }
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Param.NotFound
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// example:
	//
	// The specified parameter %s value is not valid.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// example:
	//
	// The specified parameter %s value is not valid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 9C397502-B4F2-4E22-AD97-C81F0049F3F3
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeCostInfoByDbsInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCostInfoByDbsInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) SetCode(v string) *DescribeCostInfoByDbsInstanceResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) SetData(v string) *DescribeCostInfoByDbsInstanceResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) SetErrCode(v string) *DescribeCostInfoByDbsInstanceResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) SetErrMessage(v string) *DescribeCostInfoByDbsInstanceResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) SetMessage(v string) *DescribeCostInfoByDbsInstanceResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) SetRequestId(v string) *DescribeCostInfoByDbsInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) SetSuccess(v string) *DescribeCostInfoByDbsInstanceResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeCostInfoByDbsInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
