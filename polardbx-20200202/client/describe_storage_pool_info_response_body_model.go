// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeStoragePoolInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeStoragePoolInfoResponseBody
	GetCode() *string
	SetData(v *DescribeStoragePoolInfoResponseBodyData) *DescribeStoragePoolInfoResponseBody
	GetData() *DescribeStoragePoolInfoResponseBodyData
	SetRequestId(v string) *DescribeStoragePoolInfoResponseBody
	GetRequestId() *string
}

type DescribeStoragePoolInfoResponseBody struct {
	// example:
	//
	// 200
	Code *string                                  `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *DescribeStoragePoolInfoResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeStoragePoolInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePoolInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeStoragePoolInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeStoragePoolInfoResponseBody) GetData() *DescribeStoragePoolInfoResponseBodyData {
	return s.Data
}

func (s *DescribeStoragePoolInfoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeStoragePoolInfoResponseBody) SetCode(v string) *DescribeStoragePoolInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeStoragePoolInfoResponseBody) SetData(v *DescribeStoragePoolInfoResponseBodyData) *DescribeStoragePoolInfoResponseBody {
	s.Data = v
	return s
}

func (s *DescribeStoragePoolInfoResponseBody) SetRequestId(v string) *DescribeStoragePoolInfoResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeStoragePoolInfoResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeStoragePoolInfoResponseBodyData struct {
	StoragePools []*DescribeStoragePoolInfoResponseBodyDataStoragePools `json:"StoragePools,omitempty" xml:"StoragePools,omitempty" type:"Repeated"`
}

func (s DescribeStoragePoolInfoResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePoolInfoResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeStoragePoolInfoResponseBodyData) GetStoragePools() []*DescribeStoragePoolInfoResponseBodyDataStoragePools {
	return s.StoragePools
}

func (s *DescribeStoragePoolInfoResponseBodyData) SetStoragePools(v []*DescribeStoragePoolInfoResponseBodyDataStoragePools) *DescribeStoragePoolInfoResponseBodyData {
	s.StoragePools = v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyData) Validate() error {
	if s.StoragePools != nil {
		for _, item := range s.StoragePools {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeStoragePoolInfoResponseBodyDataStoragePools struct {
	// example:
	//
	// polar.mysql.x4.large
	Class *string `json:"Class,omitempty" xml:"Class,omitempty"`
	// example:
	//
	// [pxc-xdb-s-pxcbj****xxnwp0wac2c4\\",\\"pxc-xdb-s-pxcbjr3b8****p0wa1589\\"]
	DnIdList []*string `json:"DnIdList,omitempty" xml:"DnIdList,omitempty" type:"Repeated"`
	// DN id
	//
	// example:
	//
	// pxc-xdb-s-pxcshr****rh7fn2654fc
	DnIdString *string `json:"DnIdString,omitempty" xml:"DnIdString,omitempty"`
	// example:
	//
	// {\\"appId\\":\\"APP_ETLUP1E5FMI5ND5IFO8W\\"}
	Extra *string `json:"Extra,omitempty" xml:"Extra,omitempty"`
	// example:
	//
	// 2025-08-27 23:19:52.0
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// example:
	//
	// 2025-09-02 16:52:47.0
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// example:
	//
	// none
	IdleDNIdList []*string `json:"IdleDNIdList,omitempty" xml:"IdleDNIdList,omitempty" type:"Repeated"`
	// example:
	//
	// test-cacheapi
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// pxc-xdb-s-hzr*****6j4se284344
	UnDeletableDNId *string `json:"UnDeletableDNId,omitempty" xml:"UnDeletableDNId,omitempty"`
}

func (s DescribeStoragePoolInfoResponseBodyDataStoragePools) String() string {
	return dara.Prettify(s)
}

func (s DescribeStoragePoolInfoResponseBodyDataStoragePools) GoString() string {
	return s.String()
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) GetClass() *string {
	return s.Class
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) GetDnIdList() []*string {
	return s.DnIdList
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) GetDnIdString() *string {
	return s.DnIdString
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) GetExtra() *string {
	return s.Extra
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) GetGmtModified() *string {
	return s.GmtModified
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) GetIdleDNIdList() []*string {
	return s.IdleDNIdList
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) GetName() *string {
	return s.Name
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) GetUnDeletableDNId() *string {
	return s.UnDeletableDNId
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) SetClass(v string) *DescribeStoragePoolInfoResponseBodyDataStoragePools {
	s.Class = &v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) SetDnIdList(v []*string) *DescribeStoragePoolInfoResponseBodyDataStoragePools {
	s.DnIdList = v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) SetDnIdString(v string) *DescribeStoragePoolInfoResponseBodyDataStoragePools {
	s.DnIdString = &v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) SetExtra(v string) *DescribeStoragePoolInfoResponseBodyDataStoragePools {
	s.Extra = &v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) SetGmtCreated(v string) *DescribeStoragePoolInfoResponseBodyDataStoragePools {
	s.GmtCreated = &v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) SetGmtModified(v string) *DescribeStoragePoolInfoResponseBodyDataStoragePools {
	s.GmtModified = &v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) SetIdleDNIdList(v []*string) *DescribeStoragePoolInfoResponseBodyDataStoragePools {
	s.IdleDNIdList = v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) SetName(v string) *DescribeStoragePoolInfoResponseBodyDataStoragePools {
	s.Name = &v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) SetUnDeletableDNId(v string) *DescribeStoragePoolInfoResponseBodyDataStoragePools {
	s.UnDeletableDNId = &v
	return s
}

func (s *DescribeStoragePoolInfoResponseBodyDataStoragePools) Validate() error {
	return dara.Validate(s)
}
