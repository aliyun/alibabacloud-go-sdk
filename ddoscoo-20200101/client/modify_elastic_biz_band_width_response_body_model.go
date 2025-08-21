// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyElasticBizBandWidthResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyElasticBizBandWidthResponseBody
	GetRequestId() *string
}

type ModifyElasticBizBandWidthResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// C566BA3A-192F-5D32-8A33-21422F975145
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyElasticBizBandWidthResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyElasticBizBandWidthResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyElasticBizBandWidthResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyElasticBizBandWidthResponseBody) SetRequestId(v string) *ModifyElasticBizBandWidthResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyElasticBizBandWidthResponseBody) Validate() error {
	return dara.Validate(s)
}
