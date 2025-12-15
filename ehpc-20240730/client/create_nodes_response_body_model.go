// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceIds(v []*string) *CreateNodesResponseBody
	GetInstanceIds() []*string
	SetRequestId(v string) *CreateNodesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateNodesResponseBody
	GetSuccess() *bool
}

type CreateNodesResponseBody struct {
	// The IDs of the compute nodes to be created.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE73368
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateNodesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodesResponseBody) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *CreateNodesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateNodesResponseBody) SetInstanceIds(v []*string) *CreateNodesResponseBody {
	s.InstanceIds = v
	return s
}

func (s *CreateNodesResponseBody) SetRequestId(v string) *CreateNodesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodesResponseBody) SetSuccess(v bool) *CreateNodesResponseBody {
	s.Success = &v
	return s
}

func (s *CreateNodesResponseBody) Validate() error {
	return dara.Validate(s)
}
