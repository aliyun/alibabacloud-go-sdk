// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteIntelligentStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteIntelligentStrategyResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteIntelligentStrategyResponseBody
	GetSuccess() *bool
}

type DeleteIntelligentStrategyResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Success   *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteIntelligentStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteIntelligentStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteIntelligentStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteIntelligentStrategyResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteIntelligentStrategyResponseBody) SetRequestId(v string) *DeleteIntelligentStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteIntelligentStrategyResponseBody) SetSuccess(v bool) *DeleteIntelligentStrategyResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteIntelligentStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
