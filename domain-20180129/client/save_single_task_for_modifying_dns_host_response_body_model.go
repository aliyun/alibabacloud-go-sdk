// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForModifyingDnsHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForModifyingDnsHostResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForModifyingDnsHostResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForModifyingDnsHostResponseBody struct {
	// example:
	//
	// 0F1B3547-BE50-4206-8F78-9540FFB85BC1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// e9b8e8b4-7334-4548-9cec-c30b6891f292
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForModifyingDnsHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForModifyingDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForModifyingDnsHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForModifyingDnsHostResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForModifyingDnsHostResponseBody) SetRequestId(v string) *SaveSingleTaskForModifyingDnsHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostResponseBody) SetTaskNo(v string) *SaveSingleTaskForModifyingDnsHostResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForModifyingDnsHostResponseBody) Validate() error {
	return dara.Validate(s)
}
