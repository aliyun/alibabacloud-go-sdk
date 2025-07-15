// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUnbindConfigGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupIds(v []*string) *UnbindConfigGroupResponseBody
	GetGroupIds() []*string
	SetRequestId(v string) *UnbindConfigGroupResponseBody
	GetRequestId() *string
}

type UnbindConfigGroupResponseBody struct {
	// The IDs of the configuration groups.
	GroupIds []*string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// AD0FF13D-FC7D-56AD-934F-91C8487*****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UnbindConfigGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UnbindConfigGroupResponseBody) GoString() string {
	return s.String()
}

func (s *UnbindConfigGroupResponseBody) GetGroupIds() []*string {
	return s.GroupIds
}

func (s *UnbindConfigGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UnbindConfigGroupResponseBody) SetGroupIds(v []*string) *UnbindConfigGroupResponseBody {
	s.GroupIds = v
	return s
}

func (s *UnbindConfigGroupResponseBody) SetRequestId(v string) *UnbindConfigGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *UnbindConfigGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
