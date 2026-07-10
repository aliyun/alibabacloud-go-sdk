// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLangfuseSecurityIPListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeLangfuseSecurityIPListResponseBodyData) *DescribeLangfuseSecurityIPListResponseBody
	GetData() *DescribeLangfuseSecurityIPListResponseBodyData
	SetRequestId(v string) *DescribeLangfuseSecurityIPListResponseBody
	GetRequestId() *string
}

type DescribeLangfuseSecurityIPListResponseBody struct {
	// The returned result.
	Data *DescribeLangfuseSecurityIPListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 2FED790E-FB61-4721-8C1C-07C627FA5A19
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLangfuseSecurityIPListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseSecurityIPListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseSecurityIPListResponseBody) GetData() *DescribeLangfuseSecurityIPListResponseBodyData {
	return s.Data
}

func (s *DescribeLangfuseSecurityIPListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLangfuseSecurityIPListResponseBody) SetData(v *DescribeLangfuseSecurityIPListResponseBodyData) *DescribeLangfuseSecurityIPListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBody) SetRequestId(v string) *DescribeLangfuseSecurityIPListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLangfuseSecurityIPListResponseBodyData struct {
	// The instance ID.
	//
	// example:
	//
	// lfs-****
	DBInstanceID *int32 `json:"DBInstanceID,omitempty" xml:"DBInstanceID,omitempty"`
	// The instance name.
	//
	// example:
	//
	// lfs-****
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// The whitelist group list.
	GroupItems []*DescribeLangfuseSecurityIPListResponseBodyDataGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
}

func (s DescribeLangfuseSecurityIPListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseSecurityIPListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseSecurityIPListResponseBodyData) GetDBInstanceID() *int32 {
	return s.DBInstanceID
}

func (s *DescribeLangfuseSecurityIPListResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeLangfuseSecurityIPListResponseBodyData) GetGroupItems() []*DescribeLangfuseSecurityIPListResponseBodyDataGroupItems {
	return s.GroupItems
}

func (s *DescribeLangfuseSecurityIPListResponseBodyData) SetDBInstanceID(v int32) *DescribeLangfuseSecurityIPListResponseBodyData {
	s.DBInstanceID = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBodyData) SetDBInstanceName(v string) *DescribeLangfuseSecurityIPListResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBodyData) SetGroupItems(v []*DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) *DescribeLangfuseSecurityIPListResponseBodyData {
	s.GroupItems = v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBodyData) Validate() error {
	if s.GroupItems != nil {
		for _, item := range s.GroupItems {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLangfuseSecurityIPListResponseBodyDataGroupItems struct {
	// The name of the whitelist group.
	//
	// example:
	//
	// default
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// The group tag.
	//
	// example:
	//
	// test
	GroupTag *string `json:"GroupTag,omitempty" xml:"GroupTag,omitempty"`
	// The list of IP addresses in the whitelist group.
	//
	// example:
	//
	// 127.0.XX.XX
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
	// The IP address type. The value is fixed to IPv4. IPv6 is not supported.
	//
	// example:
	//
	// ipv4
	SecurityIPType *string `json:"SecurityIPType,omitempty" xml:"SecurityIPType,omitempty"`
	// The network type of the whitelist.
	//
	// example:
	//
	// mix
	WhitelistNetType *string `json:"WhitelistNetType,omitempty" xml:"WhitelistNetType,omitempty"`
}

func (s DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) GetGroupTag() *string {
	return s.GroupTag
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) GetSecurityIPType() *string {
	return s.SecurityIPType
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) GetWhitelistNetType() *string {
	return s.WhitelistNetType
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) SetGroupName(v string) *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) SetGroupTag(v string) *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems {
	s.GroupTag = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) SetSecurityIPList(v string) *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) SetSecurityIPType(v string) *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems {
	s.SecurityIPType = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) SetWhitelistNetType(v string) *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems {
	s.WhitelistNetType = &v
	return s
}

func (s *DescribeLangfuseSecurityIPListResponseBodyDataGroupItems) Validate() error {
	return dara.Validate(s)
}
