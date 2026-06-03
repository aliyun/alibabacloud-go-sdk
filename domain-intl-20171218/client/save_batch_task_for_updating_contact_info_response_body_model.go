// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdatingContactInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForUpdatingContactInfoResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForUpdatingContactInfoResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForUpdatingContactInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForUpdatingContactInfoResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForUpdatingContactInfoResponseBody) SetRequestId(v string) *SaveBatchTaskForUpdatingContactInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoResponseBody) SetTaskNo(v string) *SaveBatchTaskForUpdatingContactInfoResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
