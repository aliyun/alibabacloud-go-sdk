// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStartInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *StartInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *StartInstanceResponseBody
	GetRequestId() *string
}

type StartInstanceResponseBody struct {
	// The ID of the bastion host that you enable.
	//
	// example:
	//
	// bastionhost-cn-78v1gh****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 0F85AE0A-6A87-48F3-A0C7-90B998788A36
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s StartInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StartInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *StartInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *StartInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StartInstanceResponseBody) SetInstanceId(v string) *StartInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *StartInstanceResponseBody) SetRequestId(v string) *StartInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *StartInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
