// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateErrorPagesRedirectRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetConfigId(v int64) *CreateErrorPagesRedirectRuleResponseBody
	GetConfigId() *int64
	SetRequestId(v string) *CreateErrorPagesRedirectRuleResponseBody
	GetRequestId() *string
}

type CreateErrorPagesRedirectRuleResponseBody struct {
	// The configuration ID.
	//
	// example:
	//
	// 427428371703808
	ConfigId *int64 `json:"ConfigId,omitempty" xml:"ConfigId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// D1D7BBB5-9B5B-5A29-8848-398F3CA18A8A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateErrorPagesRedirectRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateErrorPagesRedirectRuleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateErrorPagesRedirectRuleResponseBody) GetConfigId() *int64 {
	return s.ConfigId
}

func (s *CreateErrorPagesRedirectRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateErrorPagesRedirectRuleResponseBody) SetConfigId(v int64) *CreateErrorPagesRedirectRuleResponseBody {
	s.ConfigId = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleResponseBody) SetRequestId(v string) *CreateErrorPagesRedirectRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateErrorPagesRedirectRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
