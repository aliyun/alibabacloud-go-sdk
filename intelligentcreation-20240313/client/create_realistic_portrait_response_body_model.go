// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRealisticPortraitResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateRealisticPortraitResponseBody
	GetRequestId() *string
	SetTaskId(v string) *CreateRealisticPortraitResponseBody
	GetTaskId() *string
}

type CreateRealisticPortraitResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	TaskId    *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s CreateRealisticPortraitResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRealisticPortraitResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRealisticPortraitResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRealisticPortraitResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *CreateRealisticPortraitResponseBody) SetRequestId(v string) *CreateRealisticPortraitResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRealisticPortraitResponseBody) SetTaskId(v string) *CreateRealisticPortraitResponseBody {
	s.TaskId = &v
	return s
}

func (s *CreateRealisticPortraitResponseBody) Validate() error {
	return dara.Validate(s)
}
