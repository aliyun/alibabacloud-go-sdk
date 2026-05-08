// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTransferPortraitStyleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TransferPortraitStyleResponseBody
	GetRequestId() *string
	SetTaskId(v string) *TransferPortraitStyleResponseBody
	GetTaskId() *string
}

type TransferPortraitStyleResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 725E87CD-F2DE-5FC4-8A09-2EBDFBF26DAA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// 313123123
	TaskId *string `json:"taskId,omitempty" xml:"taskId,omitempty"`
}

func (s TransferPortraitStyleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TransferPortraitStyleResponseBody) GoString() string {
	return s.String()
}

func (s *TransferPortraitStyleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TransferPortraitStyleResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *TransferPortraitStyleResponseBody) SetRequestId(v string) *TransferPortraitStyleResponseBody {
	s.RequestId = &v
	return s
}

func (s *TransferPortraitStyleResponseBody) SetTaskId(v string) *TransferPortraitStyleResponseBody {
	s.TaskId = &v
	return s
}

func (s *TransferPortraitStyleResponseBody) Validate() error {
	return dara.Validate(s)
}
