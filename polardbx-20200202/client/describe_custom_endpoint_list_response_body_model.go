// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeCustomEndpointListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeCustomEndpointListResponseBodyData) *DescribeCustomEndpointListResponseBody
	GetData() *DescribeCustomEndpointListResponseBodyData
	SetRequestId(v string) *DescribeCustomEndpointListResponseBody
	GetRequestId() *string
}

type DescribeCustomEndpointListResponseBody struct {
	Data *DescribeCustomEndpointListResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 6352AC16-76BF-5135-B1EA-ED49293526E6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeCustomEndpointListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEndpointListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeCustomEndpointListResponseBody) GetData() *DescribeCustomEndpointListResponseBodyData {
	return s.Data
}

func (s *DescribeCustomEndpointListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeCustomEndpointListResponseBody) SetData(v *DescribeCustomEndpointListResponseBodyData) *DescribeCustomEndpointListResponseBody {
	s.Data = v
	return s
}

func (s *DescribeCustomEndpointListResponseBody) SetRequestId(v string) *DescribeCustomEndpointListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeCustomEndpointListResponseBodyData struct {
	// example:
	//
	// 0
	CanDeleteCount *int32                                                 `json:"CanDeleteCount,omitempty" xml:"CanDeleteCount,omitempty"`
	Endpoints      []*DescribeCustomEndpointListResponseBodyDataEndpoints `json:"Endpoints,omitempty" xml:"Endpoints,omitempty" type:"Repeated"`
}

func (s DescribeCustomEndpointListResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEndpointListResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeCustomEndpointListResponseBodyData) GetCanDeleteCount() *int32 {
	return s.CanDeleteCount
}

func (s *DescribeCustomEndpointListResponseBodyData) GetEndpoints() []*DescribeCustomEndpointListResponseBodyDataEndpoints {
	return s.Endpoints
}

func (s *DescribeCustomEndpointListResponseBodyData) SetCanDeleteCount(v int32) *DescribeCustomEndpointListResponseBodyData {
	s.CanDeleteCount = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyData) SetEndpoints(v []*DescribeCustomEndpointListResponseBodyDataEndpoints) *DescribeCustomEndpointListResponseBodyData {
	s.Endpoints = v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyData) Validate() error {
	if s.Endpoints != nil {
		for _, item := range s.Endpoints {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeCustomEndpointListResponseBodyDataEndpoints struct {
	// [\\"pxc-i-vb1sqa7llp\\",\\"pxc-i-bemprx50ad\\"]
	CnNames []*string `json:"CnNames,omitempty" xml:"CnNames,omitempty" type:"Repeated"`
	// example:
	//
	// pxc-shra****zq0j01.polarx.rds.aliyuncs.com
	ConnectionString *string `json:"ConnectionString,omitempty" xml:"ConnectionString,omitempty"`
	// example:
	//
	// pxe-b6e****o4pfap1s
	CustomEndpointId *string `json:"CustomEndpointId,omitempty" xml:"CustomEndpointId,omitempty"`
	// example:
	//
	// pxc-hz****zoxherr7
	DBInstanceName *string `json:"DBInstanceName,omitempty" xml:"DBInstanceName,omitempty"`
	// example:
	//
	// Name
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// true
	NodeAutoEnter *string `json:"NodeAutoEnter,omitempty" xml:"NodeAutoEnter,omitempty"`
	// example:
	//
	// same_azone_as_last
	NodeRole *string `json:"NodeRole,omitempty" xml:"NodeRole,omitempty"`
	// example:
	//
	// 3306
	Port *int32 `json:"Port,omitempty" xml:"Port,omitempty"`
	// example:
	//
	// created
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// example:
	//
	// vsw-8vbkw****5yh4nrd639ih
	VSwitchId *string `json:"VSwitchId,omitempty" xml:"VSwitchId,omitempty"`
	// example:
	//
	// vpc-uf61h****dj1zg5fqp5x7
	VpcId *string `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s DescribeCustomEndpointListResponseBodyDataEndpoints) String() string {
	return dara.Prettify(s)
}

func (s DescribeCustomEndpointListResponseBodyDataEndpoints) GoString() string {
	return s.String()
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetCnNames() []*string {
	return s.CnNames
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetConnectionString() *string {
	return s.ConnectionString
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetCustomEndpointId() *string {
	return s.CustomEndpointId
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetDBInstanceName() *string {
	return s.DBInstanceName
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetName() *string {
	return s.Name
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetNodeAutoEnter() *string {
	return s.NodeAutoEnter
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetNodeRole() *string {
	return s.NodeRole
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetPort() *int32 {
	return s.Port
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetStatus() *string {
	return s.Status
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetVSwitchId() *string {
	return s.VSwitchId
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) GetVpcId() *string {
	return s.VpcId
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetCnNames(v []*string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.CnNames = v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetConnectionString(v string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.ConnectionString = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetCustomEndpointId(v string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.CustomEndpointId = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetDBInstanceName(v string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.DBInstanceName = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetName(v string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.Name = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetNodeAutoEnter(v string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.NodeAutoEnter = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetNodeRole(v string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.NodeRole = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetPort(v int32) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.Port = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetStatus(v string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.Status = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetVSwitchId(v string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.VSwitchId = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) SetVpcId(v string) *DescribeCustomEndpointListResponseBodyDataEndpoints {
	s.VpcId = &v
	return s
}

func (s *DescribeCustomEndpointListResponseBodyDataEndpoints) Validate() error {
	return dara.Validate(s)
}
