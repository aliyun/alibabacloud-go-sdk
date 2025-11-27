// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRedeployRCInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RedeployRCInstanceResponseBody
	GetRequestId() *string
	SetTaskId(v string) *RedeployRCInstanceResponseBody
	GetTaskId() *string
}

type RedeployRCInstanceResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 866F5EB8-4650-4061-87F0-379F6F968BCE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The task ID.
	//
	// example:
	//
	// t-bp10e8orkp8x****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s RedeployRCInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RedeployRCInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *RedeployRCInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RedeployRCInstanceResponseBody) GetTaskId() *string {
	return s.TaskId
}

func (s *RedeployRCInstanceResponseBody) SetRequestId(v string) *RedeployRCInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RedeployRCInstanceResponseBody) SetTaskId(v string) *RedeployRCInstanceResponseBody {
	s.TaskId = &v
	return s
}

func (s *RedeployRCInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
