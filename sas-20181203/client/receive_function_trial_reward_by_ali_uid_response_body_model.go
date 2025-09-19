// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iReceiveFunctionTrialRewardByAliUidResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ReceiveFunctionTrialRewardByAliUidResponseBody
	GetRequestId() *string
}

type ReceiveFunctionTrialRewardByAliUidResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 94004FDB-27EC-5666-83D4-D0C5C624****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ReceiveFunctionTrialRewardByAliUidResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ReceiveFunctionTrialRewardByAliUidResponseBody) GoString() string {
	return s.String()
}

func (s *ReceiveFunctionTrialRewardByAliUidResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ReceiveFunctionTrialRewardByAliUidResponseBody) SetRequestId(v string) *ReceiveFunctionTrialRewardByAliUidResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReceiveFunctionTrialRewardByAliUidResponseBody) Validate() error {
	return dara.Validate(s)
}
