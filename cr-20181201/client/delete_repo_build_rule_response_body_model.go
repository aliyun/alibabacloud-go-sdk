// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteRepoBuildRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteRepoBuildRuleResponseBody
	GetCode() *string
	SetIsSuccess(v bool) *DeleteRepoBuildRuleResponseBody
	GetIsSuccess() *bool
	SetRequestId(v string) *DeleteRepoBuildRuleResponseBody
	GetRequestId() *string
}

type DeleteRepoBuildRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	IsSuccess *bool `json:"IsSuccess,omitempty" xml:"IsSuccess,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 2E3F55BF-FA7B-454E-B2C6-85265E243ADC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteRepoBuildRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteRepoBuildRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteRepoBuildRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteRepoBuildRuleResponseBody) GetIsSuccess() *bool {
	return s.IsSuccess
}

func (s *DeleteRepoBuildRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteRepoBuildRuleResponseBody) SetCode(v string) *DeleteRepoBuildRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteRepoBuildRuleResponseBody) SetIsSuccess(v bool) *DeleteRepoBuildRuleResponseBody {
	s.IsSuccess = &v
	return s
}

func (s *DeleteRepoBuildRuleResponseBody) SetRequestId(v string) *DeleteRepoBuildRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteRepoBuildRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
