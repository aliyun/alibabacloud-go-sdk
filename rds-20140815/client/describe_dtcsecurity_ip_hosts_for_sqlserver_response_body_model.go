// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDTCSecurityIpHostsForSQLServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeDTCSecurityIpHostsForSQLServerResponseBody
	GetDBInstanceId() *string
	SetIpHostPairNum(v string) *DescribeDTCSecurityIpHostsForSQLServerResponseBody
	GetIpHostPairNum() *string
	SetItems(v *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems) *DescribeDTCSecurityIpHostsForSQLServerResponseBody
	GetItems() *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems
	SetRequestId(v string) *DescribeDTCSecurityIpHostsForSQLServerResponseBody
	GetRequestId() *string
}

type DescribeDTCSecurityIpHostsForSQLServerResponseBody struct {
	// The instance ID.
	//
	// example:
	//
	// rm-uf6wjk5xxxxxxx
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The number of distributed transaction whitelists.
	//
	// example:
	//
	// 1
	IpHostPairNum *string                                                  `json:"IpHostPairNum,omitempty" xml:"IpHostPairNum,omitempty"`
	Items         *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 2CA62A70-2203-45C6-8E90-8971D5ACC0C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDTCSecurityIpHostsForSQLServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDTCSecurityIpHostsForSQLServerResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBody) GetIpHostPairNum() *string {
	return s.IpHostPairNum
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBody) GetItems() *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems {
	return s.Items
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBody) SetDBInstanceId(v string) *DescribeDTCSecurityIpHostsForSQLServerResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBody) SetIpHostPairNum(v string) *DescribeDTCSecurityIpHostsForSQLServerResponseBody {
	s.IpHostPairNum = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBody) SetItems(v *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems) *DescribeDTCSecurityIpHostsForSQLServerResponseBody {
	s.Items = v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBody) SetRequestId(v string) *DescribeDTCSecurityIpHostsForSQLServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBody) Validate() error {
	if s.Items != nil {
		if err := s.Items.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems struct {
	WhiteListGroups []*DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups `json:"WhiteListGroups,omitempty" xml:"WhiteListGroups,omitempty" type:"Repeated"`
}

func (s DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems) GetWhiteListGroups() []*DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups {
	return s.WhiteListGroups
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems) SetWhiteListGroups(v []*DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups) *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems {
	s.WhiteListGroups = v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItems) Validate() error {
	if s.WhiteListGroups != nil {
		for _, item := range s.WhiteListGroups {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups struct {
	SecurityIpHosts    *string `json:"SecurityIpHosts,omitempty" xml:"SecurityIpHosts,omitempty"`
	WhitelistGroupName *string `json:"WhitelistGroupName,omitempty" xml:"WhitelistGroupName,omitempty"`
}

func (s DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups) String() string {
	return dara.Prettify(s)
}

func (s DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups) GoString() string {
	return s.String()
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups) GetSecurityIpHosts() *string {
	return s.SecurityIpHosts
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups) GetWhitelistGroupName() *string {
	return s.WhitelistGroupName
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups) SetSecurityIpHosts(v string) *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups {
	s.SecurityIpHosts = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups) SetWhitelistGroupName(v string) *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups {
	s.WhitelistGroupName = &v
	return s
}

func (s *DescribeDTCSecurityIpHostsForSQLServerResponseBodyItemsWhiteListGroups) Validate() error {
	return dara.Validate(s)
}
