// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForDeletingDnsHostResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForDeletingDnsHostResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForDeletingDnsHostResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForDeletingDnsHostResponseBody struct {
	// example:
	//
	// 8fc97e44-837a-447d-ac61-ea28d2fe8a38
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 8fc97e44-837a-447d-ac61-ea28d2fexxxx
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForDeletingDnsHostResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForDeletingDnsHostResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForDeletingDnsHostResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForDeletingDnsHostResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForDeletingDnsHostResponseBody) SetRequestId(v string) *SaveSingleTaskForDeletingDnsHostResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostResponseBody) SetTaskNo(v string) *SaveSingleTaskForDeletingDnsHostResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForDeletingDnsHostResponseBody) Validate() error {
	return dara.Validate(s)
}
