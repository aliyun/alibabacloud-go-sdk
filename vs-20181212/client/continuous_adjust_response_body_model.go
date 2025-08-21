// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iContinuousAdjustResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *ContinuousAdjustResponseBody
	GetId() *string
	SetRequestId(v string) *ContinuousAdjustResponseBody
	GetRequestId() *string
}

type ContinuousAdjustResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ContinuousAdjustResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ContinuousAdjustResponseBody) GoString() string {
	return s.String()
}

func (s *ContinuousAdjustResponseBody) GetId() *string {
	return s.Id
}

func (s *ContinuousAdjustResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ContinuousAdjustResponseBody) SetId(v string) *ContinuousAdjustResponseBody {
	s.Id = &v
	return s
}

func (s *ContinuousAdjustResponseBody) SetRequestId(v string) *ContinuousAdjustResponseBody {
	s.RequestId = &v
	return s
}

func (s *ContinuousAdjustResponseBody) Validate() error {
	return dara.Validate(s)
}
