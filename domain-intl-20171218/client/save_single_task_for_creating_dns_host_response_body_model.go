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
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
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
