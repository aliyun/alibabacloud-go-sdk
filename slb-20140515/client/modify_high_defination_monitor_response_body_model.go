// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyHighDefinationMonitorResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyHighDefinationMonitorResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModifyHighDefinationMonitorResponseBody
	GetSuccess() *string
}

type ModifyHighDefinationMonitorResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8B9DB03B-ED39-5DB8-9C9F-1ED5F548D61E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ModifyHighDefinationMonitorResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyHighDefinationMonitorResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyHighDefinationMonitorResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyHighDefinationMonitorResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModifyHighDefinationMonitorResponseBody) SetRequestId(v string) *ModifyHighDefinationMonitorResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyHighDefinationMonitorResponseBody) SetSuccess(v string) *ModifyHighDefinationMonitorResponseBody {
	s.Success = &v
	return s
}

func (s *ModifyHighDefinationMonitorResponseBody) Validate() error {
	return dara.Validate(s)
}
