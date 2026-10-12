// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMobileAgentPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModifyMobileAgentPackageResponseBody
	GetCode() *string
	SetMessage(v string) *ModifyMobileAgentPackageResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModifyMobileAgentPackageResponseBody
	GetRequestId() *string
}

type ModifyMobileAgentPackageResponseBody struct {
	// The status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response message.
	//
	// example:
	//
	// Success.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 69BCBBE4-FCF2-59B8-AD9D-531EB422****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMobileAgentPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMobileAgentPackageResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMobileAgentPackageResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModifyMobileAgentPackageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModifyMobileAgentPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMobileAgentPackageResponseBody) SetCode(v string) *ModifyMobileAgentPackageResponseBody {
	s.Code = &v
	return s
}

func (s *ModifyMobileAgentPackageResponseBody) SetMessage(v string) *ModifyMobileAgentPackageResponseBody {
	s.Message = &v
	return s
}

func (s *ModifyMobileAgentPackageResponseBody) SetRequestId(v string) *ModifyMobileAgentPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMobileAgentPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
