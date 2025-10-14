// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateDataImportTaskResponseBodyData) *CreateDataImportTaskResponseBody
	GetData() *CreateDataImportTaskResponseBodyData
	SetMessage(v string) *CreateDataImportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateDataImportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataImportTaskResponseBody
	GetSuccess() *bool
}

type CreateDataImportTaskResponseBody struct {
	Data *CreateDataImportTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 173CA69A-3513-591D-8A09-C1EA37CBE2D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateDataImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataImportTaskResponseBody) GetData() *CreateDataImportTaskResponseBodyData {
	return s.Data
}

func (s *CreateDataImportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateDataImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataImportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataImportTaskResponseBody) SetData(v *CreateDataImportTaskResponseBodyData) *CreateDataImportTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateDataImportTaskResponseBody) SetMessage(v string) *CreateDataImportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateDataImportTaskResponseBody) SetRequestId(v string) *CreateDataImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataImportTaskResponseBody) SetSuccess(v bool) *CreateDataImportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataImportTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateDataImportTaskResponseBodyData struct {
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s CreateDataImportTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateDataImportTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateDataImportTaskResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *CreateDataImportTaskResponseBodyData) SetSlinkTaskId(v string) *CreateDataImportTaskResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *CreateDataImportTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
