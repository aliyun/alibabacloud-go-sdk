// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteWebCacheCustomRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteWebCacheCustomRuleResponseBody
	GetRequestId() *string
}

type DeleteWebCacheCustomRuleResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 6623EA1F-30FB-5BC8-BEC9-74D55F6F08F1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteWebCacheCustomRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteWebCacheCustomRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteWebCacheCustomRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteWebCacheCustomRuleResponseBody) SetRequestId(v string) *DeleteWebCacheCustomRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteWebCacheCustomRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
