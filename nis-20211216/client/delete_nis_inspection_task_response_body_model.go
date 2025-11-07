// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteNisInspectionTaskResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteNisInspectionTaskResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteNisInspectionTaskResponseBody
	GetRequestId() *string
}

type DeleteNisInspectionTaskResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// A7F0D6EC-E19E-58AC-AC9F-08036763960F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteNisInspectionTaskResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteNisInspectionTaskResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteNisInspectionTaskResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteNisInspectionTaskResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteNisInspectionTaskResponseBody) SetData(v bool) *DeleteNisInspectionTaskResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteNisInspectionTaskResponseBody) SetRequestId(v string) *DeleteNisInspectionTaskResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteNisInspectionTaskResponseBody) Validate() error {
	return dara.Validate(s)
}
