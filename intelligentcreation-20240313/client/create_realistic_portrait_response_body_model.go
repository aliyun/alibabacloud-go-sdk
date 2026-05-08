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
	// Id of the request
	//
	// example:
	//
	// D5798660-1531-5D12-9C20-16FEE9D22351
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
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
