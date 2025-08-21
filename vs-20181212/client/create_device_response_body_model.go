// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *CreateDeviceResponseBody
	GetId() *string
	SetRequestId(v string) *CreateDeviceResponseBody
	GetRequestId() *string
}

type CreateDeviceResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDeviceResponseBody) GetId() *string {
	return s.Id
}

func (s *CreateDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDeviceResponseBody) SetId(v string) *CreateDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDeviceResponseBody) SetRequestId(v string) *CreateDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
