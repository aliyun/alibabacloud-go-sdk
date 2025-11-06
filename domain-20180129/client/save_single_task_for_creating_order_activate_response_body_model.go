// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingOrderActivateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForCreatingOrderActivateResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForCreatingOrderActivateResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForCreatingOrderActivateResponseBody struct {
	// example:
	//
	// 40F46D3D-F4F3-4CCB-AC30-2DD20E32E528
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3cb1adc3-20e8-44ae-9e76-e812fa6fc9d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCreatingOrderActivateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingOrderActivateResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingOrderActivateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForCreatingOrderActivateResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForCreatingOrderActivateResponseBody) SetRequestId(v string) *SaveSingleTaskForCreatingOrderActivateResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateResponseBody) SetTaskNo(v string) *SaveSingleTaskForCreatingOrderActivateResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingOrderActivateResponseBody) Validate() error {
	return dara.Validate(s)
}
