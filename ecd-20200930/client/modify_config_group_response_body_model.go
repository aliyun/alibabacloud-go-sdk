// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyConfigGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetGroupId(v string) *ModifyConfigGroupResponseBody
	GetGroupId() *string
	SetRequestId(v string) *ModifyConfigGroupResponseBody
	GetRequestId() *string
}

type ModifyConfigGroupResponseBody struct {
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

func (s ModifyConfigGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyConfigGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyConfigGroupResponseBody) GetGroupId() *string {
	return s.GroupId
}

func (s *ModifyConfigGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyConfigGroupResponseBody) SetGroupId(v string) *ModifyConfigGroupResponseBody {
	s.GroupId = &v
	return s
}

func (s *ModifyConfigGroupResponseBody) SetRequestId(v string) *ModifyConfigGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyConfigGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
