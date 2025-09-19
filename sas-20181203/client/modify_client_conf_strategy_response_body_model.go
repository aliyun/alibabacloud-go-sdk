// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClientConfStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClientConfStrategyResponseBody
	GetRequestId() *string
}

type ModifyClientConfStrategyResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 9089D0AB-835F-5663-AB5E-4FF646BB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClientConfStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClientConfStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClientConfStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClientConfStrategyResponseBody) SetRequestId(v string) *ModifyClientConfStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClientConfStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
