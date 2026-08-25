// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelImageTestResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *CancelImageTestResponseBodyData) *CancelImageTestResponseBody
	GetData() *CancelImageTestResponseBodyData
	SetRequestId(v string) *CancelImageTestResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelImageTestResponseBody
	GetSuccess() *bool
}

type CancelImageTestResponseBody struct {
	// The result of the API request.
	//
	// example:
	//
	// true
	Data *CancelImageTestResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used to locate logs and troubleshoot issues.
	//
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelImageTestResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelImageTestResponseBody) GoString() string {
	return s.String()
}

func (s *CancelImageTestResponseBody) GetData() *CancelImageTestResponseBodyData {
	return s.Data
}

func (s *CancelImageTestResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelImageTestResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelImageTestResponseBody) SetData(v *CancelImageTestResponseBodyData) *CancelImageTestResponseBody {
	s.Data = v
	return s
}

func (s *CancelImageTestResponseBody) SetRequestId(v string) *CancelImageTestResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelImageTestResponseBody) SetSuccess(v bool) *CancelImageTestResponseBody {
	s.Success = &v
	return s
}

func (s *CancelImageTestResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelImageTestResponseBodyData struct {
	// The ID of the canceled image test execution.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// Indicates whether the cancellation was triggered successfully.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelImageTestResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelImageTestResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelImageTestResponseBodyData) GetProcessId() *string {
	return s.ProcessId
}

func (s *CancelImageTestResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *CancelImageTestResponseBodyData) SetProcessId(v string) *CancelImageTestResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *CancelImageTestResponseBodyData) SetSuccess(v bool) *CancelImageTestResponseBodyData {
	s.Success = &v
	return s
}

func (s *CancelImageTestResponseBodyData) Validate() error {
	return dara.Validate(s)
}
