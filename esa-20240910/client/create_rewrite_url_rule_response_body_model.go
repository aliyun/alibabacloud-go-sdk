// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRewriteUrlRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateRewriteUrlRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateRewriteUrlRuleResponseBody
	GetRequestId() *string
}

type CreateRewriteUrlRuleResponseBody struct {
	// Configuration ID.
	//
	// example:
	//
	// 39237781679****
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// Request ID.
	//
	// example:
	//
	// CB1A380B-09F0-41BB-280B-72F8FD6DA2FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRewriteUrlRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRewriteUrlRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRewriteUrlRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateRewriteUrlRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRewriteUrlRuleResponseBody) SetConfigId(v int64) *CreateRewriteUrlRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateRewriteUrlRuleResponseBody) SetRequestId(v string) *CreateRewriteUrlRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRewriteUrlRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
