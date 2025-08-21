// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopMoveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *StopMoveResponseBody
	GetId() *string
	SetRequestId(v string) *StopMoveResponseBody
	GetRequestId() *string
}

type StopMoveResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StopMoveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopMoveResponseBody) GoString() string {
	return s.String()
}

func (s *StopMoveResponseBody) GetId() *string {
	return s.Id
}

func (s *StopMoveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopMoveResponseBody) SetId(v string) *StopMoveResponseBody {
	s.Id = &v
	return s
}

func (s *StopMoveResponseBody) SetRequestId(v string) *StopMoveResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopMoveResponseBody) Validate() error {
	return dara.Validate(s)
}
