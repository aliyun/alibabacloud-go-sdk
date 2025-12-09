// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCloudVendorProductTemplateConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DescribeCloudVendorProductTemplateConfigResponseBody
	GetData() *string
	SetRequestId(v string) *DescribeCloudVendorProductTemplateConfigResponseBody
	GetRequestId() *string
}

type DescribeCloudVendorProductTemplateConfigResponseBody struct {
	// example:
	//
	// {\\"vendors\\":[{\\"vendorType\\":100,\\"vendor\\":\\"CHAITIN\\",\\"displayName\\":\\"Chaitin\\",\\"products\\":[{\\"product\\":\\"webFirewall\\",\\"displayName\\":\\"WAF\\",\\"backendConfig\\":{\\"apiParams\\":{\\"SecretKey\\":{\\"field\\":\\"apiToken\\",\\"format\\":\\"text\\"},\\"SecretId\\":{\\"field\\":\\"endpoint\\",\\"format\\":\\"text\\"},\\"Vendor\\":{\\"field\\":\\"vendor\\",\\"format\\":\\"text\\"},\\"CtdrCloudUserId\\":{\\"field\\":\\"ctdrCloudUserId\\",\\"format\\":\\"text\\"},\\"ExtendInfo\\":{\\"format\\":\\"json\\",\\"fields\\":[\\"product\\",\\"remark\\"]}}},\\"description\\":\\"https://docs.waf-ce.chaitin.cn/%E6%9B%B4%E5%A4%9A%E6%8A%80%E6%9C%AF%E6%96%87%E6%A1%A3/OPENAPI\\",\\"fields\\":[{\\"displayName\\":\\"Endpoint\\",\\"fieldType\\":\\"text\\",\\"prompt\\":\\"Enter an endpoint that is in the IP address:Port number format.\\",\\"required\\":true,\\"fieldId\\":\\"endpoint\\"},{\\"displayName\\":\\"API Token\\",\\"fieldType\\":\\"password\\",\\"required\\":true,\\"fieldId\\":\\"apiToken\\"},{\\"displayName\\":\\"‌Device Name‌ (The cloud_user_id field in the device connection logs will uniquely identify the device)\\",\\"fieldType\\":\\"text\\",\\"required\\":true,\\"fieldId\\":\\"ctdrCloudUserId\\"},{\\"displayName\\":\\"Remark\\",\\"fieldType\\":\\"text\\",\\"required\\":false,\\"fieldId\\":\\"remark\\"}]}]}]}
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// CD380235-A0B8-540D-A0D5-D6288446****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCloudVendorProductTemplateConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCloudVendorProductTemplateConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCloudVendorProductTemplateConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *DescribeCloudVendorProductTemplateConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCloudVendorProductTemplateConfigResponseBody) SetData(v string) *DescribeCloudVendorProductTemplateConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DescribeCloudVendorProductTemplateConfigResponseBody) SetRequestId(v string) *DescribeCloudVendorProductTemplateConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCloudVendorProductTemplateConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
