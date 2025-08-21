// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartDeviceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StartDeviceResponseBody
	GetId() *string
	SetRequestId(v string) *StartDeviceResponseBody
	GetRequestId() *string
}

type StartDeviceResponseBody struct {
	// example:
	//
	// 323884****9092996
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartDeviceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartDeviceResponseBody) GoString() string {
	return s.String()
}

func (s *StartDeviceResponseBody) GetId() *string {
	return s.Id
}

func (s *StartDeviceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartDeviceResponseBody) SetId(v string) *StartDeviceResponseBody {
	s.Id = &v
	return s
}

func (s *StartDeviceResponseBody) SetRequestId(v string) *StartDeviceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartDeviceResponseBody) Validate() error {
	return dara.Validate(s)
}
