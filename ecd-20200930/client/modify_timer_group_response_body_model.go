// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyTimerGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ModifyTimerGroupResponseBody
	GetGroupId() *string
	SetRequestId(v string) *ModifyTimerGroupResponseBody
	GetRequestId() *string
}

type ModifyTimerGroupResponseBody struct {
	// The ID of the configuration group.
	//
	// example:
	//
	// cg-i1ruuudp92qpj****
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyTimerGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyTimerGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyTimerGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyTimerGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyTimerGroupResponseBody) SetGroupId(v string) *ModifyTimerGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *ModifyTimerGroupResponseBody) SetRequestId(v string) *ModifyTimerGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyTimerGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
