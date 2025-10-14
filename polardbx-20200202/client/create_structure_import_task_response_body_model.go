// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateStructureImportTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CreateStructureImportTaskResponseBodyData) *CreateStructureImportTaskResponseBody
	GetData() *CreateStructureImportTaskResponseBodyData
	SetMessage(v string) *CreateStructureImportTaskResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateStructureImportTaskResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateStructureImportTaskResponseBody
	GetSuccess() *bool
}

type CreateStructureImportTaskResponseBody struct {
	Data *CreateStructureImportTaskResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 73559800-3c8c-****-bd40-99cfcff3fe1e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateStructureImportTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateStructureImportTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStructureImportTaskResponseBody) GetData() *CreateStructureImportTaskResponseBodyData {
	return s.Data
}

func (s *CreateStructureImportTaskResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateStructureImportTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateStructureImportTaskResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateStructureImportTaskResponseBody) SetData(v *CreateStructureImportTaskResponseBodyData) *CreateStructureImportTaskResponseBody {
	s.Data = v
	return s
}

func (s *CreateStructureImportTaskResponseBody) SetMessage(v string) *CreateStructureImportTaskResponseBody {
	s.Message = &v
	return s
}

func (s *CreateStructureImportTaskResponseBody) SetRequestId(v string) *CreateStructureImportTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateStructureImportTaskResponseBody) SetSuccess(v bool) *CreateStructureImportTaskResponseBody {
	s.Success = &v
	return s
}

func (s *CreateStructureImportTaskResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateStructureImportTaskResponseBodyData struct {
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s CreateStructureImportTaskResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CreateStructureImportTaskResponseBodyData) GoString() string {
	return s.String()
}

func (s *CreateStructureImportTaskResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *CreateStructureImportTaskResponseBodyData) SetSlinkTaskId(v string) *CreateStructureImportTaskResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *CreateStructureImportTaskResponseBodyData) Validate() error {
	return dara.Validate(s)
}
