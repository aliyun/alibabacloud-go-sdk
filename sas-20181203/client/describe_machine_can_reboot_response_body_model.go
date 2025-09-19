// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeMachineCanRebootResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCanReboot(v bool) *DescribeMachineCanRebootResponseBody
	GetCanReboot() *bool
	SetRequestId(v string) *DescribeMachineCanRebootResponseBody
	GetRequestId() *string
}

type DescribeMachineCanRebootResponseBody struct {
	// Indicates whether the server can be restarted. Valid values:
	//
	// 	- **true**: yes
	//
	// 	- **false**: no
	//
	// example:
	//
	// true
	CanReboot *bool `json:"CanReboot,omitempty" xml:"CanReboot,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 79CFF74D-E967-5407-8A78-EE03B925FDAA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeMachineCanRebootResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeMachineCanRebootResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeMachineCanRebootResponseBody) GetCanReboot() *bool {
	return s.CanReboot
}

func (s *DescribeMachineCanRebootResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeMachineCanRebootResponseBody) SetCanReboot(v bool) *DescribeMachineCanRebootResponseBody {
	s.CanReboot = &v
	return s
}

func (s *DescribeMachineCanRebootResponseBody) SetRequestId(v string) *DescribeMachineCanRebootResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeMachineCanRebootResponseBody) Validate() error {
	return dara.Validate(s)
}
