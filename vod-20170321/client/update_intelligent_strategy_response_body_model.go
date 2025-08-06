// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateIntelligentStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateIntelligentStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateIntelligentStrategyResponseBody
	GetSuccess() *bool
}

type UpdateIntelligentStrategyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateIntelligentStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateIntelligentStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateIntelligentStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateIntelligentStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateIntelligentStrategyResponseBody) SetRequestId(v string) *UpdateIntelligentStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateIntelligentStrategyResponseBody) SetSuccess(v bool) *UpdateIntelligentStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateIntelligentStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
