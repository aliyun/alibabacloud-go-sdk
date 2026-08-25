// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunImageTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *RunImageTestResponseBodyData) *RunImageTestResponseBody
	GetData() *RunImageTestResponseBodyData
	SetRequestId(v string) *RunImageTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RunImageTestResponseBody
	GetSuccess() *bool
}

type RunImageTestResponseBody struct {
	// The result of the API request.
	//
	// example:
	//
	// true
	Data *RunImageTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunImageTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunImageTestResponseBody) GoString() string {
	return s.String()
}

func (s *RunImageTestResponseBody) GetData() *RunImageTestResponseBodyData {
	return s.Data
}

func (s *RunImageTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunImageTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RunImageTestResponseBody) SetData(v *RunImageTestResponseBodyData) *RunImageTestResponseBody {
	s.Data = v
	return s
}

func (s *RunImageTestResponseBody) SetRequestId(v string) *RunImageTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunImageTestResponseBody) SetSuccess(v bool) *RunImageTestResponseBody {
	s.Success = &v
	return s
}

func (s *RunImageTestResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RunImageTestResponseBodyData struct {
	// The image test execution ID.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// Indicates whether the trigger is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RunImageTestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RunImageTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *RunImageTestResponseBodyData) GetProcessId() *string {
	return s.ProcessId
}

func (s *RunImageTestResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *RunImageTestResponseBodyData) SetProcessId(v string) *RunImageTestResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *RunImageTestResponseBodyData) SetSuccess(v bool) *RunImageTestResponseBodyData {
	s.Success = &v
	return s
}

func (s *RunImageTestResponseBodyData) Validate() error {
	return dara.Validate(s)
}
