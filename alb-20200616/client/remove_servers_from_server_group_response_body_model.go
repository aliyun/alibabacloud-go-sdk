// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRemoveServersFromServerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *RemoveServersFromServerGroupResponseBody
	GetJobId() *string
	SetRequestId(v string) *RemoveServersFromServerGroupResponseBody
	GetRequestId() *string
}

type RemoveServersFromServerGroupResponseBody struct {
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

func (s RemoveServersFromServerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RemoveServersFromServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveServersFromServerGroupResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *RemoveServersFromServerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RemoveServersFromServerGroupResponseBody) SetJobId(v string) *RemoveServersFromServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) SetRequestId(v string) *RemoveServersFromServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveServersFromServerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
