// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSecurityIpsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeSecurityIpsResponseBodyData) *DescribeSecurityIpsResponseBody
	GetData() *DescribeSecurityIpsResponseBodyData
	SetMessage(v string) *DescribeSecurityIpsResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeSecurityIpsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeSecurityIpsResponseBody
	GetSuccess() *bool
}

type DescribeSecurityIpsResponseBody struct {
	Data *DescribeSecurityIpsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 14036EBE-CF2E-44DB-ACE9-AC6157D3A6FC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeSecurityIpsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIpsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBody) GetData() *DescribeSecurityIpsResponseBodyData {
	return s.Data
}

func (s *DescribeSecurityIpsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeSecurityIpsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSecurityIpsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeSecurityIpsResponseBody) SetData(v *DescribeSecurityIpsResponseBodyData) *DescribeSecurityIpsResponseBody {
	s.Data = v
	return s
}

func (s *DescribeSecurityIpsResponseBody) SetMessage(v string) *DescribeSecurityIpsResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeSecurityIpsResponseBody) SetRequestId(v string) *DescribeSecurityIpsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSecurityIpsResponseBody) SetSuccess(v bool) *DescribeSecurityIpsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeSecurityIpsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeSecurityIpsResponseBodyData struct {
	// example:
	//
	// pxc-hzjasd****
	DBInstanceName *string                                          `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	GroupItems     []*DescribeSecurityIpsResponseBodyDataGroupItems `json:"GroupItems,omitempty" xml:"GroupItems,omitempty" type:"Repeated"`
}

func (s DescribeSecurityIpsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIpsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBodyData) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeSecurityIpsResponseBodyData) GetGroupItems() []*DescribeSecurityIpsResponseBodyDataGroupItems {
	return s.GroupItems
}

func (s *DescribeSecurityIpsResponseBodyData) SetDBInstanceName(v string) *DescribeSecurityIpsResponseBodyData {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodyData) SetGroupItems(v []*DescribeSecurityIpsResponseBodyDataGroupItems) *DescribeSecurityIpsResponseBodyData {
	s.GroupItems = v
	return s
}

func (s *DescribeSecurityIpsResponseBodyData) Validate() error {
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

type DescribeSecurityIpsResponseBodyDataGroupItems struct {
	// example:
	//
	// defaultGroup
	GroupName *string `json:"GroupName,omitempty" xml:"GroupName,omitempty"`
	// example:
	//
	// 127.0.0.1,172.168.0.0
	SecurityIPList *string `json:"SecurityIPList,omitempty" xml:"SecurityIPList,omitempty"`
}

func (s DescribeSecurityIpsResponseBodyDataGroupItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeSecurityIpsResponseBodyDataGroupItems) GoString() string {
	return s.String()
}

func (s *DescribeSecurityIpsResponseBodyDataGroupItems) GetGroupName() *string {
	return s.GroupName
}

func (s *DescribeSecurityIpsResponseBodyDataGroupItems) GetSecurityIPList() *string {
	return s.SecurityIPList
}

func (s *DescribeSecurityIpsResponseBodyDataGroupItems) SetGroupName(v string) *DescribeSecurityIpsResponseBodyDataGroupItems {
	s.GroupName = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodyDataGroupItems) SetSecurityIPList(v string) *DescribeSecurityIpsResponseBodyDataGroupItems {
	s.SecurityIPList = &v
	return s
}

func (s *DescribeSecurityIpsResponseBodyDataGroupItems) Validate() error {
	return dara.Validate(s)
}
