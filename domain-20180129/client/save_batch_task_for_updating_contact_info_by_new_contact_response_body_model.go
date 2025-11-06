// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdatingContactInfoByNewContactResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody struct {
	// example:
	//
	// 464AF466-CA8E-43A8-B61D-test
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 65de2165-ca09-491f-9fe0-test
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) SetRequestId(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) SetTaskNo(v string) *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByNewContactResponseBody) Validate() error {
	return dara.Validate(s)
}
