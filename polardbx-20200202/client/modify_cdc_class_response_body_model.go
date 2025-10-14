// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyCdcClassResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ModifyCdcClassResponseBodyData) *ModifyCdcClassResponseBody
	GetData() *ModifyCdcClassResponseBodyData
	SetRequestId(v string) *ModifyCdcClassResponseBody
	GetRequestId() *string
}

type ModifyCdcClassResponseBody struct {
	Data *ModifyCdcClassResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyCdcClassResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdcClassResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyCdcClassResponseBody) GetData() *ModifyCdcClassResponseBodyData {
	return s.Data
}

func (s *ModifyCdcClassResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyCdcClassResponseBody) SetData(v *ModifyCdcClassResponseBodyData) *ModifyCdcClassResponseBody {
	s.Data = v
	return s
}

func (s *ModifyCdcClassResponseBody) SetRequestId(v string) *ModifyCdcClassResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyCdcClassResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModifyCdcClassResponseBodyData struct {
	// example:
	//
	// 2209883
	TaskId *int32 `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s ModifyCdcClassResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModifyCdcClassResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModifyCdcClassResponseBodyData) GetTaskId() *int32 {
	return s.TaskId
}

func (s *ModifyCdcClassResponseBodyData) SetTaskId(v int32) *ModifyCdcClassResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *ModifyCdcClassResponseBodyData) Validate() error {
	return dara.Validate(s)
}
