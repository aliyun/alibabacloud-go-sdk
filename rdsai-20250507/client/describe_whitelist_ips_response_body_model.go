// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeWhitelistIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeWhitelistIpsResponseBodyData) *DescribeWhitelistIpsResponseBody
	GetData() *DescribeWhitelistIpsResponseBodyData
	SetMessage(v string) *DescribeWhitelistIpsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeWhitelistIpsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeWhitelistIpsResponseBody
	GetSuccess() *bool
}

type DescribeWhitelistIpsResponseBody struct {
	Data *DescribeWhitelistIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeWhitelistIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistIpsResponseBody) GetData() *DescribeWhitelistIpsResponseBodyData {
	return s.Data
}

func (s *DescribeWhitelistIpsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeWhitelistIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeWhitelistIpsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeWhitelistIpsResponseBody) SetData(v *DescribeWhitelistIpsResponseBodyData) *DescribeWhitelistIpsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeWhitelistIpsResponseBody) SetMessage(v string) *DescribeWhitelistIpsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeWhitelistIpsResponseBody) SetRequestId(v string) *DescribeWhitelistIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeWhitelistIpsResponseBody) SetSuccess(v bool) *DescribeWhitelistIpsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeWhitelistIpsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeWhitelistIpsResponseBodyData struct {
	// example:
	//
	// www.test123.com
	Domain *string `json:"Domain,omitempty" xml:"Domain,omitempty"`
	// example:
	//
	// 127.0.0.1,192.168.1.0/24
	IpWhitelist *string `json:"IpWhitelist,omitempty" xml:"IpWhitelist,omitempty"`
}

func (s DescribeWhitelistIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeWhitelistIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeWhitelistIpsResponseBodyData) GetDomain() *string {
	return s.Domain
}

func (s *DescribeWhitelistIpsResponseBodyData) GetIpWhitelist() *string {
	return s.IpWhitelist
}

func (s *DescribeWhitelistIpsResponseBodyData) SetDomain(v string) *DescribeWhitelistIpsResponseBodyData {
	s.Domain = &v
	return s
}

func (s *DescribeWhitelistIpsResponseBodyData) SetIpWhitelist(v string) *DescribeWhitelistIpsResponseBodyData {
	s.IpWhitelist = &v
	return s
}

func (s *DescribeWhitelistIpsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
