// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateCacheRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateCacheRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateCacheRuleResponseBody
	GetRequestId() *string
}

type CreateCacheRuleResponseBody struct {
	ConfigId  *int64  `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateCacheRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateCacheRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCacheRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateCacheRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateCacheRuleResponseBody) SetConfigId(v int64) *CreateCacheRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateCacheRuleResponseBody) SetRequestId(v string) *CreateCacheRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateCacheRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
