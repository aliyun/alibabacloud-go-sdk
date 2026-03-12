// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingContactByTempateIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForUpdatingContactByTempateIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveTaskForUpdatingContactByTempateIdResponseBody
	GetSuccess() *bool
	SetTaskNo(v string) *SaveTaskForUpdatingContactByTempateIdResponseBody
	GetTaskNo() *string
}

type SaveTaskForUpdatingContactByTempateIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForUpdatingContactByTempateIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingContactByTempateIdResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingContactByTempateIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForUpdatingContactByTempateIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveTaskForUpdatingContactByTempateIdResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForUpdatingContactByTempateIdResponseBody) SetRequestId(v string) *SaveTaskForUpdatingContactByTempateIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdResponseBody) SetSuccess(v bool) *SaveTaskForUpdatingContactByTempateIdResponseBody {
	s.Success = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdResponseBody) SetTaskNo(v string) *SaveTaskForUpdatingContactByTempateIdResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTempateIdResponseBody) Validate() error {
	return dara.Validate(s)
}
