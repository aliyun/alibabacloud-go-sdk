// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSaveTaskForUpdatingContactByTemplateIdResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SaveTaskForUpdatingContactByTemplateIdResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SaveTaskForUpdatingContactByTemplateIdResponseBody
	GetSuccess() *bool
	SetTaskNo(v string) *SaveTaskForUpdatingContactByTemplateIdResponseBody
	GetTaskNo() *string
}

type SaveTaskForUpdatingContactByTemplateIdResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
	TaskNo    *string `json:"TaskNo,omitempty" xml:"TaskNo,omitempty"`
}

func (s SaveTaskForUpdatingContactByTemplateIdResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SaveTaskForUpdatingContactByTemplateIdResponseBody) GoString() string {
	return s.String()
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponseBody) GetTaskNo() *string {
	return s.TaskNo
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponseBody) SetRequestId(v string) *SaveTaskForUpdatingContactByTemplateIdResponseBody {
	s.RequestId = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponseBody) SetSuccess(v bool) *SaveTaskForUpdatingContactByTemplateIdResponseBody {
	s.Success = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponseBody) SetTaskNo(v string) *SaveTaskForUpdatingContactByTemplateIdResponseBody {
	s.TaskNo = &v
	return s
}

func (s *SaveTaskForUpdatingContactByTemplateIdResponseBody) Validate() error {
	return dara.Validate(s)
}
