// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartSwitchDatabaseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *StartSwitchDatabaseResponseBodyData) *StartSwitchDatabaseResponseBody
	GetData() *StartSwitchDatabaseResponseBodyData
	SetMessage(v string) *StartSwitchDatabaseResponseBody
	GetMessage() *string
	SetRequestId(v string) *StartSwitchDatabaseResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *StartSwitchDatabaseResponseBody
	GetSuccess() *bool
}

type StartSwitchDatabaseResponseBody struct {
	Data *StartSwitchDatabaseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// *****
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 9B2F3840-5C98-475C-B269-2D5C3A31797C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s StartSwitchDatabaseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartSwitchDatabaseResponseBody) GoString() string {
	return s.String()
}

func (s *StartSwitchDatabaseResponseBody) GetData() *StartSwitchDatabaseResponseBodyData {
	return s.Data
}

func (s *StartSwitchDatabaseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *StartSwitchDatabaseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartSwitchDatabaseResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *StartSwitchDatabaseResponseBody) SetData(v *StartSwitchDatabaseResponseBodyData) *StartSwitchDatabaseResponseBody {
	s.Data = v
	return s
}

func (s *StartSwitchDatabaseResponseBody) SetMessage(v string) *StartSwitchDatabaseResponseBody {
	s.Message = &v
	return s
}

func (s *StartSwitchDatabaseResponseBody) SetRequestId(v string) *StartSwitchDatabaseResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartSwitchDatabaseResponseBody) SetSuccess(v bool) *StartSwitchDatabaseResponseBody {
	s.Success = &v
	return s
}

func (s *StartSwitchDatabaseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type StartSwitchDatabaseResponseBodyData struct {
	// example:
	//
	// etx-szr2rr6i*****
	SlinkTaskId *string `json:"SlinkTaskId,omitempty" xml:"SlinkTaskId,omitempty"`
}

func (s StartSwitchDatabaseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s StartSwitchDatabaseResponseBodyData) GoString() string {
	return s.String()
}

func (s *StartSwitchDatabaseResponseBodyData) GetSlinkTaskId() *string {
	return s.SlinkTaskId
}

func (s *StartSwitchDatabaseResponseBodyData) SetSlinkTaskId(v string) *StartSwitchDatabaseResponseBodyData {
	s.SlinkTaskId = &v
	return s
}

func (s *StartSwitchDatabaseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
