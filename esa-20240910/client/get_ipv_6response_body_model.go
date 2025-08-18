// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIPv6ResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEnable(v string) *GetIPv6ResponseBody
	GetEnable() *string
	SetRegion(v string) *GetIPv6ResponseBody
	GetRegion() *string
	SetRequestId(v string) *GetIPv6ResponseBody
	GetRequestId() *string
}

type GetIPv6ResponseBody struct {
	// Indicates whether IPv6 is enabled. Valid values:
	//
	// 	- **on**
	//
	// 	- **off**
	//
	// example:
	//
	// on
	Enable *string `json:"Enable,omitempty" xml:"Enable,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 156A6B-677B1A-4297B7-9187B7-2B44792
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetIPv6ResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetIPv6ResponseBody) GoString() string {
	return s.String()
}

func (s *GetIPv6ResponseBody) GetEnable() *string {
	return s.Enable
}

func (s *GetIPv6ResponseBody) GetRegion() *string {
	return s.Region
}

func (s *GetIPv6ResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetIPv6ResponseBody) SetEnable(v string) *GetIPv6ResponseBody {
	s.Enable = &v
	return s
}

func (s *GetIPv6ResponseBody) SetRegion(v string) *GetIPv6ResponseBody {
	s.Region = &v
	return s
}

func (s *GetIPv6ResponseBody) SetRequestId(v string) *GetIPv6ResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetIPv6ResponseBody) Validate() error {
	return dara.Validate(s)
}
