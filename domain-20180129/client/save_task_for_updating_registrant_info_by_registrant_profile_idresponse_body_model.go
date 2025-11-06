// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody
	GetTaskNo() *string
}

type SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody struct {
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3cb1adc3-20e8-44ae-9e76-e812fa6fc9d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) SetRequestId(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) SetTaskNo(v string) *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForUpdatingRegistrantInfoByRegistrantProfileIDResponseBody) Validate() error {
	return dara.Validate(s)
}
