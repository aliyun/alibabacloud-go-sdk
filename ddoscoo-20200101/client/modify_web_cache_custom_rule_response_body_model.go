// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyWebCacheCustomRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyWebCacheCustomRuleResponseBody
	GetRequestId() *string
}

type ModifyWebCacheCustomRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 0bcf28g5-d57c-11e7-9bs0-d89d6717dxbc
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyWebCacheCustomRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyWebCacheCustomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyWebCacheCustomRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyWebCacheCustomRuleResponseBody) SetRequestId(v string) *ModifyWebCacheCustomRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyWebCacheCustomRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
