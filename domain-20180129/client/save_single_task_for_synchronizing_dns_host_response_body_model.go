// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForSynchronizingDnsHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForSynchronizingDnsHostResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForSynchronizingDnsHostResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForSynchronizingDnsHostResponseBody struct {
	// example:
	//
	// 0F1B3547-BE50-4206-8F78-9540FFB85BC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// e9b8e8b4-7334-4548-9cec-c30b6891f292
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForSynchronizingDnsHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForSynchronizingDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponseBody) SetRequestId(v string) *SaveSingleTaskForSynchronizingDnsHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponseBody) SetTaskNo(v string) *SaveSingleTaskForSynchronizingDnsHostResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForSynchronizingDnsHostResponseBody) Validate() error {
	return dara.Validate(s)
}
