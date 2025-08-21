// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDirectoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModifyDirectoryResponseBody
	GetId() *string
	SetRequestId(v string) *ModifyDirectoryResponseBody
	GetRequestId() *string
}

type ModifyDirectoryResponseBody struct {
	// example:
	//
	// 399*****488-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDirectoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDirectoryResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDirectoryResponseBody) GetId() *string {
	return s.Id
}

func (s *ModifyDirectoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDirectoryResponseBody) SetId(v string) *ModifyDirectoryResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyDirectoryResponseBody) SetRequestId(v string) *ModifyDirectoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDirectoryResponseBody) Validate() error {
	return dara.Validate(s)
}
