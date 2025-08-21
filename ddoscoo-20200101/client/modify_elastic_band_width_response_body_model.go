// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBandWidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyElasticBandWidthResponseBody
	GetRequestId() *string
}

type ModifyElasticBandWidthResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticBandWidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBandWidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticBandWidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElasticBandWidthResponseBody) SetRequestId(v string) *ModifyElasticBandWidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElasticBandWidthResponseBody) Validate() error {
	return dara.Validate(s)
}
