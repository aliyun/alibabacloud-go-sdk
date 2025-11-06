// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForCreatingDnsHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForCreatingDnsHostResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForCreatingDnsHostResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForCreatingDnsHostResponseBody struct {
	// example:
	//
	// 0F1B3547-BE50-4206-8F78-9540FFB85BC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// e9b8e8b4-7334-4548-9cec-c30b6891f292
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForCreatingDnsHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForCreatingDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForCreatingDnsHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForCreatingDnsHostResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForCreatingDnsHostResponseBody) SetRequestId(v string) *SaveSingleTaskForCreatingDnsHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostResponseBody) SetTaskNo(v string) *SaveSingleTaskForCreatingDnsHostResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForCreatingDnsHostResponseBody) Validate() error {
	return dara.Validate(s)
}
