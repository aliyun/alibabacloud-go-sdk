// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateNotifyStrategyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNotifyStrategyId(v string) *UpdateNotifyStrategyResponseBody
	GetNotifyStrategyId() *string
	SetRequestId(v string) *UpdateNotifyStrategyResponseBody
	GetRequestId() *string
}

type UpdateNotifyStrategyResponseBody struct {
	// example:
	//
	// 12312-31-23-123-1-23123
	NotifyStrategyId *string `json:"notifyStrategyId,omitempty" xml:"notifyStrategyId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateNotifyStrategyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateNotifyStrategyResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateNotifyStrategyResponseBody) GetNotifyStrategyId() *string {
	return s.NotifyStrategyId
}

func (s *UpdateNotifyStrategyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateNotifyStrategyResponseBody) SetNotifyStrategyId(v string) *UpdateNotifyStrategyResponseBody {
	s.NotifyStrategyId = &v
	return s
}

func (s *UpdateNotifyStrategyResponseBody) SetRequestId(v string) *UpdateNotifyStrategyResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateNotifyStrategyResponseBody) Validate() error {
	return dara.Validate(s)
}
