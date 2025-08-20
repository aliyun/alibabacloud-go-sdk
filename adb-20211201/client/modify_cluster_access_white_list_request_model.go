// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterAccessWhiteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterIPArrayAttribute(v string) *ModifyClusterAccessWhiteListRequest
	GetDBClusterIPArrayAttribute() *string
	SetDBClusterIPArrayName(v string) *ModifyClusterAccessWhiteListRequest
	GetDBClusterIPArrayName() *string
	SetDBClusterId(v string) *ModifyClusterAccessWhiteListRequest
	GetDBClusterId() *string
	SetModifyMode(v string) *ModifyClusterAccessWhiteListRequest
	GetModifyMode() *string
	SetSecurityIps(v string) *ModifyClusterAccessWhiteListRequest
	GetSecurityIps() *string
}

type ModifyClusterAccessWhiteListRequest struct {
	// The attribute of the IP address whitelist. By default, this parameter is empty.
	//
	// >  IP address whitelists with the hidden attribute are not displayed in the console. Those whitelists are used to access Data Transmission Service (DTS) and PolarDB.
	//
	// example:
	//
	// hidden
	DBClusterIPArrayAttribute *string `json:"DBClusterIPArrayAttribute,omitempty" xml:"DBClusterIPArrayAttribute,omitempty"`
	// The name of the IP address whitelist. If you do not specify this parameter, the Default whitelist is modified.
	//
	// 	- The whitelist name must be 2 to 32 characters in length. The name can contain lowercase letters, digits, and underscores (_). The name must start with a lowercase letter and end with a lowercase letter or a digit.
	//
	// 	- Each cluster supports up to 50 IP address whitelists.
	//
	// example:
	//
	// test
	DBClusterIPArrayName *string `json:"DBClusterIPArrayName,omitempty" xml:"DBClusterIPArrayName,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// This parameter is required.
	//
	// example:
	//
	// amv-bp11q28kvl688****
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The method used to modify the IP address whitelist. Valid values:
	//
	// 	- **Cover*	- (default)
	//
	// 	- **Append**
	//
	// 	- **Delete**
	//
	// example:
	//
	// Cover
	ModifyMode *string `json:"ModifyMode,omitempty" xml:"ModifyMode,omitempty"`
	// The IP addresses in an IP address whitelist of a cluster. Separate multiple IP addresses with commas (,). You can add a maximum of 500 different IP addresses to a whitelist. The entries in the IP address whitelist must be in one of the following formats:
	//
	// 	- IP addresses, such as 10.23.XX.XX.
	//
	// 	- CIDR blocks, such as 10.23.xx.xx/24. In this example, 24 indicates that the prefix of each IP address in the IP whitelist is 24 bits in length. You can replace 24 with a value within the range of 1 to 32.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10.23.xx.xx
	SecurityIps *string `json:"SecurityIps,omitempty" xml:"SecurityIps,omitempty"`
}

func (s ModifyClusterAccessWhiteListRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterAccessWhiteListRequest) GoString() string {
	return s.String()
}

func (s *ModifyClusterAccessWhiteListRequest) GetDBClusterIPArrayAttribute() *string {
	return s.DBClusterIPArrayAttribute
}

func (s *ModifyClusterAccessWhiteListRequest) GetDBClusterIPArrayName() *string {
	return s.DBClusterIPArrayName
}

func (s *ModifyClusterAccessWhiteListRequest) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *ModifyClusterAccessWhiteListRequest) GetModifyMode() *string {
	return s.ModifyMode
}

func (s *ModifyClusterAccessWhiteListRequest) GetSecurityIps() *string {
	return s.SecurityIps
}

func (s *ModifyClusterAccessWhiteListRequest) SetDBClusterIPArrayAttribute(v string) *ModifyClusterAccessWhiteListRequest {
	s.DBClusterIPArrayAttribute = &v
	return s
}

func (s *ModifyClusterAccessWhiteListRequest) SetDBClusterIPArrayName(v string) *ModifyClusterAccessWhiteListRequest {
	s.DBClusterIPArrayName = &v
	return s
}

func (s *ModifyClusterAccessWhiteListRequest) SetDBClusterId(v string) *ModifyClusterAccessWhiteListRequest {
	s.DBClusterId = &v
	return s
}

func (s *ModifyClusterAccessWhiteListRequest) SetModifyMode(v string) *ModifyClusterAccessWhiteListRequest {
	s.ModifyMode = &v
	return s
}

func (s *ModifyClusterAccessWhiteListRequest) SetSecurityIps(v string) *ModifyClusterAccessWhiteListRequest {
	s.SecurityIps = &v
	return s
}

func (s *ModifyClusterAccessWhiteListRequest) Validate() error {
	return dara.Validate(s)
}
