// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAdjustResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StopAdjustResponseBody
	GetId() *string
	SetRequestId(v string) *StopAdjustResponseBody
	GetRequestId() *string
}

type StopAdjustResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopAdjustResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAdjustResponseBody) GoString() string {
	return s.String()
}

func (s *StopAdjustResponseBody) GetId() *string {
	return s.Id
}

func (s *StopAdjustResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAdjustResponseBody) SetId(v string) *StopAdjustResponseBody {
	s.Id = &v
	return s
}

func (s *StopAdjustResponseBody) SetRequestId(v string) *StopAdjustResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAdjustResponseBody) Validate() error {
	return dara.Validate(s)
}
