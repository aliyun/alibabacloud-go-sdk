// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMobileAgentPackageResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteMobileAgentPackageResponseBody
	GetCode() *string
	SetMessage(v string) *DeleteMobileAgentPackageResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteMobileAgentPackageResponseBody
	GetRequestId() *string
}

type DeleteMobileAgentPackageResponseBody struct {
	// The status code. A value of 200 indicates success.
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
	// The request ID.
	//
	// example:
	//
	// F07A1DA1-E1EB-5CCA-8EED-12F85D32****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMobileAgentPackageResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMobileAgentPackageResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMobileAgentPackageResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteMobileAgentPackageResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteMobileAgentPackageResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMobileAgentPackageResponseBody) SetCode(v string) *DeleteMobileAgentPackageResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteMobileAgentPackageResponseBody) SetMessage(v string) *DeleteMobileAgentPackageResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteMobileAgentPackageResponseBody) SetRequestId(v string) *DeleteMobileAgentPackageResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMobileAgentPackageResponseBody) Validate() error {
	return dara.Validate(s)
}
