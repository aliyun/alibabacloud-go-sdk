// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iTriggerNetworkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *TriggerNetworkResponseBody
	GetRequestId() *string
	SetResult(v bool) *TriggerNetworkResponseBody
	GetResult() *bool
}

type TriggerNetworkResponseBody struct {
	// example:
	//
	// 5A5D8E74-565C-43DC-B031-29289FA****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Result *bool `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s TriggerNetworkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s TriggerNetworkResponseBody) GoString() string {
	return s.String()
}

func (s *TriggerNetworkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *TriggerNetworkResponseBody) GetResult() *bool {
	return s.Result
}

func (s *TriggerNetworkResponseBody) SetRequestId(v string) *TriggerNetworkResponseBody {
	s.RequestId = &v
	return s
}

func (s *TriggerNetworkResponseBody) SetResult(v bool) *TriggerNetworkResponseBody {
	s.Result = &v
	return s
}

func (s *TriggerNetworkResponseBody) Validate() error {
	return dara.Validate(s)
}
