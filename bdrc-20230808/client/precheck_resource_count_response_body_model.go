// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPrecheckResourceCountResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *PrecheckResourceCountResponseBodyData) *PrecheckResourceCountResponseBody
	GetData() *PrecheckResourceCountResponseBodyData
	SetRequestId(v string) *PrecheckResourceCountResponseBody
	GetRequestId() *string
}

type PrecheckResourceCountResponseBody struct {
	Data *PrecheckResourceCountResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 86DEBAC9-AB6A-59AB-9E5C-A540E579ECC9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PrecheckResourceCountResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PrecheckResourceCountResponseBody) GoString() string {
	return s.String()
}

func (s *PrecheckResourceCountResponseBody) GetData() *PrecheckResourceCountResponseBodyData {
	return s.Data
}

func (s *PrecheckResourceCountResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PrecheckResourceCountResponseBody) SetData(v *PrecheckResourceCountResponseBodyData) *PrecheckResourceCountResponseBody {
	s.Data = v
	return s
}

func (s *PrecheckResourceCountResponseBody) SetRequestId(v string) *PrecheckResourceCountResponseBody {
	s.RequestId = &v
	return s
}

func (s *PrecheckResourceCountResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PrecheckResourceCountResponseBodyData struct {
	// example:
	//
	// t-bp1ewftyzmeg3bl4dtd2
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s PrecheckResourceCountResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s PrecheckResourceCountResponseBodyData) GoString() string {
	return s.String()
}

func (s *PrecheckResourceCountResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *PrecheckResourceCountResponseBodyData) SetTaskId(v string) *PrecheckResourceCountResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *PrecheckResourceCountResponseBodyData) Validate() error {
	return dara.Validate(s)
}
