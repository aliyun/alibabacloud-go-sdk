// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyRuleStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedIds(v string) *ModifyRuleStatusResponseBody
	GetFailedIds() *string
	SetRequestId(v string) *ModifyRuleStatusResponseBody
	GetRequestId() *string
}

type ModifyRuleStatusResponseBody struct {
	// The IDs of sensitive data detection rules whose status failed to be changed. Multiple IDs are separated with commas (,).
	//
	// example:
	//
	// 1,2,3,4
	FailedIds *string `json:"FailedIds,omitempty" xml:"FailedIds,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 7C3AC882-E5A8-4855-BE77-B6837B695EF1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyRuleStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyRuleStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyRuleStatusResponseBody) GetFailedIds() *string {
	return s.FailedIds
}

func (s *ModifyRuleStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyRuleStatusResponseBody) SetFailedIds(v string) *ModifyRuleStatusResponseBody {
	s.FailedIds = &v
	return s
}

func (s *ModifyRuleStatusResponseBody) SetRequestId(v string) *ModifyRuleStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyRuleStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
