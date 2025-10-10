// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddServersToServerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetJobId(v string) *AddServersToServerGroupResponseBody
	GetJobId() *string
	SetRequestId(v string) *AddServersToServerGroupResponseBody
	GetRequestId() *string
}

type AddServersToServerGroupResponseBody struct {
	// The ID of the asynchronous job.
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

func (s AddServersToServerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddServersToServerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *AddServersToServerGroupResponseBody) GetJobId() *string {
	return s.JobId
}

func (s *AddServersToServerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddServersToServerGroupResponseBody) SetJobId(v string) *AddServersToServerGroupResponseBody {
	s.JobId = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) SetRequestId(v string) *AddServersToServerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddServersToServerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
