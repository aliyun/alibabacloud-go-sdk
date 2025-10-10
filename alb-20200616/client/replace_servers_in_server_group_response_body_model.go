// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReplaceServersInServerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *ReplaceServersInServerGroupResponseBody
	GetJobId() *string
	SetRequestId(v string) *ReplaceServersInServerGroupResponseBody
	GetRequestId() *string
}

type ReplaceServersInServerGroupResponseBody struct {
	// The asynchronous task ID.
	//
	// example:
	//
	// 72dcd26b-f12d-4c27-b3af-18f6aed5****
	JobId *string `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 365F4154-92F6-4AE4-92F8-7FF34B540710
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReplaceServersInServerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReplaceServersInServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ReplaceServersInServerGroupResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *ReplaceServersInServerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReplaceServersInServerGroupResponseBody) SetJobId(v string) *ReplaceServersInServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *ReplaceServersInServerGroupResponseBody) SetRequestId(v string) *ReplaceServersInServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReplaceServersInServerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
