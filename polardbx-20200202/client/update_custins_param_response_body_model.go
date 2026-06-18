// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCustinsParamResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateCustinsParamResponseBodyData) *UpdateCustinsParamResponseBody
	GetData() *UpdateCustinsParamResponseBodyData
	SetRequestId(v string) *UpdateCustinsParamResponseBody
	GetRequestId() *string
}

type UpdateCustinsParamResponseBody struct {
	Data *UpdateCustinsParamResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 73559800-3c8c-****-bd40-99cfcff3fe1e
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCustinsParamResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustinsParamResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCustinsParamResponseBody) GetData() *UpdateCustinsParamResponseBodyData {
	return s.Data
}

func (s *UpdateCustinsParamResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCustinsParamResponseBody) SetData(v *UpdateCustinsParamResponseBodyData) *UpdateCustinsParamResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCustinsParamResponseBody) SetRequestId(v string) *UpdateCustinsParamResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCustinsParamResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCustinsParamResponseBodyData struct {
	// example:
	//
	// ******
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateCustinsParamResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateCustinsParamResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCustinsParamResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *UpdateCustinsParamResponseBodyData) SetTaskId(v int32) *UpdateCustinsParamResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateCustinsParamResponseBodyData) Validate() error {
	return dara.Validate(s)
}
