// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRestartDataImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *RestartDataImportTaskResponseBody
	GetData() *bool
	SetMessage(v string) *RestartDataImportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *RestartDataImportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RestartDataImportTaskResponseBody
	GetSuccess() *bool
}

type RestartDataImportTaskResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// FEA5DC20-6D8A-5979-97AA-FC57546ADC20
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RestartDataImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RestartDataImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *RestartDataImportTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *RestartDataImportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RestartDataImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RestartDataImportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RestartDataImportTaskResponseBody) SetData(v bool) *RestartDataImportTaskResponseBody {
	s.Data = &v
	return s
}

func (s *RestartDataImportTaskResponseBody) SetMessage(v string) *RestartDataImportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *RestartDataImportTaskResponseBody) SetRequestId(v string) *RestartDataImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *RestartDataImportTaskResponseBody) SetSuccess(v bool) *RestartDataImportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *RestartDataImportTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
