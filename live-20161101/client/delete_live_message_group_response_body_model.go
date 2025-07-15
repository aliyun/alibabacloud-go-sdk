// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteLiveMessageGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *DeleteLiveMessageGroupResponseBody
	GetGroupId() *string
	SetRequestId(v string) *DeleteLiveMessageGroupResponseBody
	GetRequestId() *string
}

type DeleteLiveMessageGroupResponseBody struct {
	// The ID of the deleted or non-existent group.
	//
	// example:
	//
	// grouptest
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// B95BE680-5A6A-1CAD-8AB1-09DFF5D6****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteLiveMessageGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteLiveMessageGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteLiveMessageGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *DeleteLiveMessageGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteLiveMessageGroupResponseBody) SetGroupId(v string) *DeleteLiveMessageGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *DeleteLiveMessageGroupResponseBody) SetRequestId(v string) *DeleteLiveMessageGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteLiveMessageGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
