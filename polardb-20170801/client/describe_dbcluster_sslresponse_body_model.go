// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBClusterSSLResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*DescribeDBClusterSSLResponseBodyItems) *DescribeDBClusterSSLResponseBody
	GetItems() []*DescribeDBClusterSSLResponseBodyItems
	SetRequestId(v string) *DescribeDBClusterSSLResponseBody
	GetRequestId() *string
	SetSSLAutoRotate(v string) *DescribeDBClusterSSLResponseBody
	GetSSLAutoRotate() *string
}

type DescribeDBClusterSSLResponseBody struct {
	// The list of SSL connections.
	Items []*DescribeDBClusterSSLResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// C890995A-CF06-4F4D-8DB8-DD26C2******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether automatic rotation of SSL certificates is enabled. Valid values:
	//
	// 	- **Enable**: The feature is enabled.
	//
	// 	- **Disable**: The feature is disabled.
	//
	// > This parameter is valid only for a PolarDB for MySQL cluster.
	//
	// example:
	//
	// Enable
	SSLAutoRotate *string `json:"SSLAutoRotate,omitempty" xml:"SSLAutoRotate,omitempty"`
}

func (s DescribeDBClusterSSLResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSSLResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLResponseBody) GetItems() []*DescribeDBClusterSSLResponseBodyItems {
	return s.Items
}

func (s *DescribeDBClusterSSLResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBClusterSSLResponseBody) GetSSLAutoRotate() *string {
	return s.SSLAutoRotate
}

func (s *DescribeDBClusterSSLResponseBody) SetItems(v []*DescribeDBClusterSSLResponseBodyItems) *DescribeDBClusterSSLResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDBClusterSSLResponseBody) SetRequestId(v string) *DescribeDBClusterSSLResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBody) SetSSLAutoRotate(v string) *DescribeDBClusterSSLResponseBody {
	s.SSLAutoRotate = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBody) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDBClusterSSLResponseBodyItems struct {
	// The ID of the endpoint.
	//
	// example:
	//
	// pe-************
	DBEndpointId *string `json:"DBEndpointId,omitempty" xml:"DBEndpointId,omitempty"`
	// example:
	//
	// Enable
	SSLAutoRotate *string `json:"SSLAutoRotate,omitempty" xml:"SSLAutoRotate,omitempty"`
	// The SSL connection string.
	//
	// example:
	//
	// pc-************.mysql.polardb.rds.aliyuncs.com
	SSLConnectionString *string `json:"SSLConnectionString,omitempty" xml:"SSLConnectionString,omitempty"`
	// Indicates whether SSL encryption is enabled. Valid values:
	//
	// 	- **Enabled**: SSL is enabled.
	//
	// 	- **Disable**: SSL is disabled.
	//
	// example:
	//
	// Enabled
	SSLEnabled *string `json:"SSLEnabled,omitempty" xml:"SSLEnabled,omitempty"`
	// The time when the server certificate expires. The time is in the `yyyy-MM-ddTHH:mm:ssZ` format. The time is displayed in UTC.
	//
	// example:
	//
	// 2021-11-13T07:14:22Z
	SSLExpireTime *string `json:"SSLExpireTime,omitempty" xml:"SSLExpireTime,omitempty"`
}

func (s DescribeDBClusterSSLResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBClusterSSLResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDBClusterSSLResponseBodyItems) GetDBEndpointId() *string {
	return s.DBEndpointId
}

func (s *DescribeDBClusterSSLResponseBodyItems) GetSSLAutoRotate() *string {
	return s.SSLAutoRotate
}

func (s *DescribeDBClusterSSLResponseBodyItems) GetSSLConnectionString() *string {
	return s.SSLConnectionString
}

func (s *DescribeDBClusterSSLResponseBodyItems) GetSSLEnabled() *string {
	return s.SSLEnabled
}

func (s *DescribeDBClusterSSLResponseBodyItems) GetSSLExpireTime() *string {
	return s.SSLExpireTime
}

func (s *DescribeDBClusterSSLResponseBodyItems) SetDBEndpointId(v string) *DescribeDBClusterSSLResponseBodyItems {
	s.DBEndpointId = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBodyItems) SetSSLAutoRotate(v string) *DescribeDBClusterSSLResponseBodyItems {
	s.SSLAutoRotate = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBodyItems) SetSSLConnectionString(v string) *DescribeDBClusterSSLResponseBodyItems {
	s.SSLConnectionString = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBodyItems) SetSSLEnabled(v string) *DescribeDBClusterSSLResponseBodyItems {
	s.SSLEnabled = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBodyItems) SetSSLExpireTime(v string) *DescribeDBClusterSSLResponseBodyItems {
	s.SSLExpireTime = &v
	return s
}

func (s *DescribeDBClusterSSLResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
