// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNisInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *UpdateNisInspectionTaskResponseBody
	GetData() *bool
	SetRequestId(v string) *UpdateNisInspectionTaskResponseBody
	GetRequestId() *string
}

type UpdateNisInspectionTaskResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateNisInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNisInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNisInspectionTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *UpdateNisInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNisInspectionTaskResponseBody) SetData(v bool) *UpdateNisInspectionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateNisInspectionTaskResponseBody) SetRequestId(v string) *UpdateNisInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNisInspectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
