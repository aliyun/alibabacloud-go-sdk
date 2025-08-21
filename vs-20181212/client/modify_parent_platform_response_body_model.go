// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyParentPlatformResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModifyParentPlatformResponseBody
	GetId() *string
	SetRequestId(v string) *ModifyParentPlatformResponseBody
	GetRequestId() *string
}

type ModifyParentPlatformResponseBody struct {
	// example:
	//
	// 359*****374-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyParentPlatformResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyParentPlatformResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyParentPlatformResponseBody) GetId() *string {
	return s.Id
}

func (s *ModifyParentPlatformResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyParentPlatformResponseBody) SetId(v string) *ModifyParentPlatformResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyParentPlatformResponseBody) SetRequestId(v string) *ModifyParentPlatformResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyParentPlatformResponseBody) Validate() error {
	return dara.Validate(s)
}
