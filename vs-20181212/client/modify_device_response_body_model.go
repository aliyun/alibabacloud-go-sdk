// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ModifyDeviceResponseBody
	GetId() *string
	SetRequestId(v string) *ModifyDeviceResponseBody
	GetRequestId() *string
}

type ModifyDeviceResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDeviceResponseBody) GetId() *string {
	return s.Id
}

func (s *ModifyDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDeviceResponseBody) SetId(v string) *ModifyDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *ModifyDeviceResponseBody) SetRequestId(v string) *ModifyDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
