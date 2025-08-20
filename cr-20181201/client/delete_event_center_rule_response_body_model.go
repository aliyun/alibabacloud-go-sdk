// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteEventCenterRuleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteEventCenterRuleResponseBody
	GetCode() *string
	SetRequestId(v string) *DeleteEventCenterRuleResponseBody
	GetRequestId() *string
}

type DeleteEventCenterRuleResponseBody struct {
	// The return value.
	//
	// example:
	//
	// success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 031572FA-7D8F-4C05-B790-1071E0E05DE6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteEventCenterRuleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteEventCenterRuleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventCenterRuleResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteEventCenterRuleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteEventCenterRuleResponseBody) SetCode(v string) *DeleteEventCenterRuleResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteEventCenterRuleResponseBody) SetRequestId(v string) *DeleteEventCenterRuleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventCenterRuleResponseBody) Validate() error {
	return dara.Validate(s)
}
