// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityModeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetModule(v string) *DescribeSecurityModeResponseBody
	GetModule() *string
	SetRequestId(v string) *DescribeSecurityModeResponseBody
	GetRequestId() *string
	SetSecurityMode(v int32) *DescribeSecurityModeResponseBody
	GetSecurityMode() *int32
}

type DescribeSecurityModeResponseBody struct {
	// Deprecated
	//
	// example:
	//
	// sg_server
	Module *string `json:"Module,omitempty" xml:"Module,omitempty"`
	// example:
	//
	// C6D68A02-54D5-5F5C-A8AA-6D6C2874****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	SecurityMode *int32 `json:"SecurityMode,omitempty" xml:"SecurityMode,omitempty"`
}

func (s DescribeSecurityModeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityModeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityModeResponseBody) GetModule() *string {
	return s.Module
}

func (s *DescribeSecurityModeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityModeResponseBody) GetSecurityMode() *int32 {
	return s.SecurityMode
}

func (s *DescribeSecurityModeResponseBody) SetModule(v string) *DescribeSecurityModeResponseBody {
	s.Module = &v
	return s
}

func (s *DescribeSecurityModeResponseBody) SetRequestId(v string) *DescribeSecurityModeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityModeResponseBody) SetSecurityMode(v int32) *DescribeSecurityModeResponseBody {
	s.SecurityMode = &v
	return s
}

func (s *DescribeSecurityModeResponseBody) Validate() error {
	return dara.Validate(s)
}
