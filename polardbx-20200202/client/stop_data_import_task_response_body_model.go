// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDataImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *StopDataImportTaskResponseBody
	GetData() *bool
	SetMessage(v string) *StopDataImportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *StopDataImportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StopDataImportTaskResponseBody
	GetSuccess() *bool
}

type StopDataImportTaskResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 14036EBE-***-44DB-ACE9-AC6157D3A6FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StopDataImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDataImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *StopDataImportTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *StopDataImportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StopDataImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDataImportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StopDataImportTaskResponseBody) SetData(v bool) *StopDataImportTaskResponseBody {
	s.Data = &v
	return s
}

func (s *StopDataImportTaskResponseBody) SetMessage(v string) *StopDataImportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *StopDataImportTaskResponseBody) SetRequestId(v string) *StopDataImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDataImportTaskResponseBody) SetSuccess(v bool) *StopDataImportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *StopDataImportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
