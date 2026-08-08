// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateCrossAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *UpdateCrossAccountsResponseBodyData) *UpdateCrossAccountsResponseBody
	GetData() *UpdateCrossAccountsResponseBodyData
	SetRequestId(v string) *UpdateCrossAccountsResponseBody
	GetRequestId() *string
}

type UpdateCrossAccountsResponseBody struct {
	Data *UpdateCrossAccountsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 8724BC18-904D-5A0D-BFF4-F0554F0037E7
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateCrossAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateCrossAccountsResponseBody) GetData() *UpdateCrossAccountsResponseBodyData {
	return s.Data
}

func (s *UpdateCrossAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateCrossAccountsResponseBody) SetData(v *UpdateCrossAccountsResponseBodyData) *UpdateCrossAccountsResponseBody {
	s.Data = v
	return s
}

func (s *UpdateCrossAccountsResponseBody) SetRequestId(v string) *UpdateCrossAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateCrossAccountsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateCrossAccountsResponseBodyData struct {
	// example:
	//
	// t-0000e4w0u1v592zdf6s7
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s UpdateCrossAccountsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s UpdateCrossAccountsResponseBodyData) GoString() string {
	return s.String()
}

func (s *UpdateCrossAccountsResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *UpdateCrossAccountsResponseBodyData) SetTaskId(v string) *UpdateCrossAccountsResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *UpdateCrossAccountsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
