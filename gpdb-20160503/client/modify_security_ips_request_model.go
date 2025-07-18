// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifySecurityIpsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest
	GetDBInstanceIPArrayAttribute() *string
	SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest
	GetDBInstanceIPArrayName() *string
	SetDBInstanceId(v string) *ModifySecurityIpsRequest
	GetDBInstanceId() *string
	SetModifyMode(v string) *ModifySecurityIpsRequest
	GetModifyMode() *string
	SetResourceGroupId(v string) *ModifySecurityIpsRequest
	GetResourceGroupId() *string
	SetSecurityIPList(v string) *ModifySecurityIpsRequest
	GetSecurityIPList() *string
}

type ModifySecurityIpsRequest struct {
	// The attribute of the IP address whitelist. By default, this parameter is empty. A whitelist with the `hidden` attribute does not appear in the console.
	//
	// example:
	//
	// hidden
	DBInstanceIPArrayAttribute *string `json:"DBInstanceIPArrayAttribute,omitempty" xml:"DBInstanceIPArrayAttribute,omitempty"`
	// The name of the whitelist. If you do not enter a name, IP addresses are added to the default whitelist.
	//
	// >  You can create up to 50 whitelists for an instance.
	//
	// example:
	//
	// default
	DBInstanceIPArrayName *string `json:"DBInstanceIPArrayName,omitempty" xml:"DBInstanceIPArrayName,omitempty"`
	// The ID of the instance.
	//
	// >  You can call the [DescribeDBInstances](https://help.aliyun.com/document_detail/86911.html) operation to query the instance IDs of all AnalyticDB for PostgreSQL instances in a specific region.
	//
	// This parameter is required.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The method of modification. Valid values:
	//
	// 	- **Cover**: overwrites the whitelist.
	//
	// 	- **Append**: appends data to the whitelist.
	//
	// 	- **Delete**: deletes the whitelist.
	//
	// example:
	//
	// 0
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The ID of the resource group to which the instance belongs. For more information about how to obtain the ID of a resource group, see [View basic information of a resource group](https://help.aliyun.com/document_detail/151181.html).
	//
	// example:
	//
	// rg-bp67acfmxazb4p****
	ResourceGroupId *string `json:"ResourceGroupId,omitempty" xml:"ResourceGroupId,omitempty"`
	// The IP addresses listed in the whitelist. You can add up to 1,000 IP addresses to the whitelist. Separate multiple IP addresses with commas (,). The IP addresses must use one of the following formats:
	//
	// 	- 0.0.0.0/0
	//
	// 	- 10.23.12.24. This is a standard IP address.
	//
	// 	- 10.23.12.24/24. This is a CIDR block. The value `/24` indicates that the prefix of the CIDR block is 24-bit long. You can replace 24 with a value in the range of `1 to 32`.
	//
	// This parameter is required.
	//
	// example:
	//
	// ``10.10.**.**``
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s ModifySecurityIpsRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifySecurityIpsRequest) GoString() string {
	return s.String()
}

func (s *ModifySecurityIpsRequest) GetDBInstanceIPArrayAttribute() *string {
	return s.DBInstanceIPArrayAttribute
}

func (s *ModifySecurityIpsRequest) GetDBInstanceIPArrayName() *string {
	return s.DBInstanceIPArrayName
}

func (s *ModifySecurityIpsRequest) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *ModifySecurityIpsRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifySecurityIpsRequest) GetResourceGroupId() *string {
	return s.ResourceGroupId
}

func (s *ModifySecurityIpsRequest) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayAttribute(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayAttribute = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceIPArrayName(v string) *ModifySecurityIpsRequest {
	s.DBInstanceIPArrayName = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetDBInstanceId(v string) *ModifySecurityIpsRequest {
	s.DBInstanceId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetModifyMode(v string) *ModifySecurityIpsRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetResourceGroupId(v string) *ModifySecurityIpsRequest {
	s.ResourceGroupId = &v
	return s
}

func (s *ModifySecurityIpsRequest) SetSecurityIPList(v string) *ModifySecurityIpsRequest {
	s.SecurityIPList = &v
	return s
}

func (s *ModifySecurityIpsRequest) Validate() error {
	return dara.Validate(s)
}
