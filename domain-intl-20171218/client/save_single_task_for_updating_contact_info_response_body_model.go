// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveSingleTaskForUpdatingContactInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveSingleTaskForUpdatingContactInfoResponseBody
	GetRequestId() *string
	SetTaskNo(v string) *SaveSingleTaskForUpdatingContactInfoResponseBody
	GetTaskNo() *string
}

type SaveSingleTaskForUpdatingContactInfoResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveSingleTaskForUpdatingContactInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveSingleTaskForUpdatingContactInfoResponseBody) GoString() string {
	return s.String()
}

func (s *SaveSingleTaskForUpdatingContactInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveSingleTaskForUpdatingContactInfoResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveSingleTaskForUpdatingContactInfoResponseBody) SetRequestId(v string) *SaveSingleTaskForUpdatingContactInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoResponseBody) SetTaskNo(v string) *SaveSingleTaskForUpdatingContactInfoResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveSingleTaskForUpdatingContactInfoResponseBody) Validate() error {
	return dara.Validate(s)
}
