// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModifyGroupResponseBody
	GetId() *string
	SetRequestId(v string) *ModifyGroupResponseBody
	GetRequestId() *string
}

type ModifyGroupResponseBody struct {
	// example:
	//
	// 32388487739092994-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyGroupResponseBody) GetId() *string {
	return s.Id
}

func (s *ModifyGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyGroupResponseBody) SetId(v string) *ModifyGroupResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyGroupResponseBody) SetRequestId(v string) *ModifyGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
