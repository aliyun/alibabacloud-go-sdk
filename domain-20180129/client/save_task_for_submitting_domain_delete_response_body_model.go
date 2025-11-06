// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForSubmittingDomainDeleteResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForSubmittingDomainDeleteResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveTaskForSubmittingDomainDeleteResponseBody
	GetTaskNo() *string
}

type SaveTaskForSubmittingDomainDeleteResponseBody struct {
	// example:
	//
	// 23C9B3C4-9E2C-4405-A88D-BD33E459D140
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 3cb1adc3-20e8-44ae-9e76-e812fa6fc9d8
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForSubmittingDomainDeleteResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForSubmittingDomainDeleteResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForSubmittingDomainDeleteResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForSubmittingDomainDeleteResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForSubmittingDomainDeleteResponseBody) SetRequestId(v string) *SaveTaskForSubmittingDomainDeleteResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteResponseBody) SetTaskNo(v string) *SaveTaskForSubmittingDomainDeleteResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForSubmittingDomainDeleteResponseBody) Validate() error {
	return dara.Validate(s)
}
