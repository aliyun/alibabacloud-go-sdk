// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateNodeGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeGroupId(v string) *CreateNodeGroupResponseBody
	GetNodeGroupId() *string
	SetRequestId(v string) *CreateNodeGroupResponseBody
	GetRequestId() *string
}

type CreateNodeGroupResponseBody struct {
	// The ID of the machine group.
	//
	// example:
	//
	// G-21E39B11837E****
	NodeGroupId *string `json:"NodeGroupId,omitempty" xml:"NodeGroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 9E3A7161-EB7B-172B-8D18-FFB06BA3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateNodeGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateNodeGroupResponseBody) GoString() string {
	return s.String()
}

func (s *CreateNodeGroupResponseBody) GetNodeGroupId() *string {
	return s.NodeGroupId
}

func (s *CreateNodeGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateNodeGroupResponseBody) SetNodeGroupId(v string) *CreateNodeGroupResponseBody {
	s.NodeGroupId = &v
	return s
}

func (s *CreateNodeGroupResponseBody) SetRequestId(v string) *CreateNodeGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateNodeGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
