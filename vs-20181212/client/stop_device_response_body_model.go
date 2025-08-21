// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StopDeviceResponseBody
	GetId() *string
	SetRequestId(v string) *StopDeviceResponseBody
	GetRequestId() *string
}

type StopDeviceResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *StopDeviceResponseBody) GetId() *string {
	return s.Id
}

func (s *StopDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopDeviceResponseBody) SetId(v string) *StopDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *StopDeviceResponseBody) SetRequestId(v string) *StopDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
