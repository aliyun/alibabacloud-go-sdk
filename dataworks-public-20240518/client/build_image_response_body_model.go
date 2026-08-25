// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iBuildImageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *BuildImageResponseBodyData) *BuildImageResponseBody
	GetData() *BuildImageResponseBodyData
	SetRequestId(v string) *BuildImageResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *BuildImageResponseBody
	GetSuccess() *bool
}

type BuildImageResponseBody struct {
	// The result of the API request.
	//
	// example:
	//
	// true
	Data *BuildImageResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID, which is used for locating logs and troubleshooting issues.
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

func (s BuildImageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s BuildImageResponseBody) GoString() string {
	return s.String()
}

func (s *BuildImageResponseBody) GetData() *BuildImageResponseBodyData {
	return s.Data
}

func (s *BuildImageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *BuildImageResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *BuildImageResponseBody) SetData(v *BuildImageResponseBodyData) *BuildImageResponseBody {
	s.Data = v
	return s
}

func (s *BuildImageResponseBody) SetRequestId(v string) *BuildImageResponseBody {
	s.RequestId = &v
	return s
}

func (s *BuildImageResponseBody) SetSuccess(v bool) *BuildImageResponseBody {
	s.Success = &v
	return s
}

func (s *BuildImageResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type BuildImageResponseBodyData struct {
	// The image build execution ID.
	//
	// example:
	//
	// 582d4896-d224-413b-b883-239eeebe0bc5
	ProcessId *string `json:"ProcessId,omitempty" xml:"ProcessId,omitempty"`
	// Indicates whether the build was triggered successfully.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s BuildImageResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s BuildImageResponseBodyData) GoString() string {
	return s.String()
}

func (s *BuildImageResponseBodyData) GetProcessId() *string {
	return s.ProcessId
}

func (s *BuildImageResponseBodyData) GetSuccess() *bool {
	return s.Success
}

func (s *BuildImageResponseBodyData) SetProcessId(v string) *BuildImageResponseBodyData {
	s.ProcessId = &v
	return s
}

func (s *BuildImageResponseBodyData) SetSuccess(v bool) *BuildImageResponseBodyData {
	s.Success = &v
	return s
}

func (s *BuildImageResponseBodyData) Validate() error {
	return dara.Validate(s)
}
