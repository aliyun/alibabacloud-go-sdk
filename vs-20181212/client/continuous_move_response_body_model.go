// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuousMoveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ContinuousMoveResponseBody
	GetId() *string
	SetRequestId(v string) *ContinuousMoveResponseBody
	GetRequestId() *string
}

type ContinuousMoveResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinuousMoveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContinuousMoveResponseBody) GoString() string {
	return s.String()
}

func (s *ContinuousMoveResponseBody) GetId() *string {
	return s.Id
}

func (s *ContinuousMoveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContinuousMoveResponseBody) SetId(v string) *ContinuousMoveResponseBody {
	s.Id = &v
	return s
}

func (s *ContinuousMoveResponseBody) SetRequestId(v string) *ContinuousMoveResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinuousMoveResponseBody) Validate() error {
	return dara.Validate(s)
}
