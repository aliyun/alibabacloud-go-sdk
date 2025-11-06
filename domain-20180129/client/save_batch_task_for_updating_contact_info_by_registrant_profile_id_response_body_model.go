// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody
	GetTaskNo() *string
}

type SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody struct {
	// example:
	//
	// EDC28FEC-6BE0-4583-95BC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 880f1579-be51-4dd3-a69d
	TaskNo *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) GoString() string {
	return s.String()
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) SetRequestId(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) SetTaskNo(v string) *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveBatchTaskForUpdatingContactInfoByRegistrantProfileIdResponseBody) Validate() error {
	return dara.Validate(s)
}
