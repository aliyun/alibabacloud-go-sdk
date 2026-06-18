// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDBInstanceConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *DescribeDBInstanceConfigResponseBodyData) *DescribeDBInstanceConfigResponseBody
	GetData() *DescribeDBInstanceConfigResponseBodyData
	SetRequestId(v string) *DescribeDBInstanceConfigResponseBody
	GetRequestId() *string
}

type DescribeDBInstanceConfigResponseBody struct {
	// The data struct.
	Data *DescribeDBInstanceConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// C458B1E8-1683-3645-B154-6BA32080EEA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeDBInstanceConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponseBody) GetData() *DescribeDBInstanceConfigResponseBodyData {
	return s.Data
}

func (s *DescribeDBInstanceConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDBInstanceConfigResponseBody) SetData(v *DescribeDBInstanceConfigResponseBodyData) *DescribeDBInstanceConfigResponseBody {
	s.Data = v
	return s
}

func (s *DescribeDBInstanceConfigResponseBody) SetRequestId(v string) *DescribeDBInstanceConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDBInstanceConfigResponseBodyData struct {
	// The configuration key.
	//
	// example:
	//
	// htap
	ConfigName *string `json:"ConfigName,omitempty" xml:"ConfigName,omitempty"`
	// The configuration item. The following parameters are included:
	//
	// - attendHtapList: the list of instances for which HTAP is enabled.
	//
	// - autoAttendHtap: specifies whether to automatically add newly created read-only instances to the HTAP list.
	//
	// - delayExecutionStrategy: when the read-only instance lag reaches the value specified by storageDelayThreshold, read-only traffic is routed back to the primary instance. Default value: 1. Valid values: 0 and 1.
	//
	// - enableConsistentReplicaRead: specifies whether to enable consistent reads.
	//
	// - storageDelayThreshold: the latency threshold for read-only instances. Default value: 3s. Valid values: 0 to 86400.
	//
	// - enableHtap: specifies whether to enable HTAP.
	//
	// - masterReadWeight: the read weight of the primary node. A value of 100 indicates that 100% of traffic is routed to the primary node. Valid values: 0 to 100.
	//
	// example:
	//
	// {\\"attendHtapList\\":[\\"pxc-hzruyzes08****\\"],\\"autoAttendHtap\\":\\"true\\",\\"delayExecutionStrategy\\":1,\\"enableConsistentReplicaRead\\":true,\\"enableHtap\\":\\"true\\",\\"masterReadWeight\\":100,\\"storageDelayThreshold\\":3}
	ConfigValue *string `json:"ConfigValue,omitempty" xml:"ConfigValue,omitempty"`
	// The instance ID.
	//
	// example:
	//
	// pxc-sprcym7g7w****
	DbInstanceName *string `json:"DbInstanceName,omitempty" xml:"DbInstanceName,omitempty"`
}

func (s DescribeDBInstanceConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribeDBInstanceConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribeDBInstanceConfigResponseBodyData) GetConfigName() *string {
	return s.ConfigName
}

func (s *DescribeDBInstanceConfigResponseBodyData) GetConfigValue() *string {
	return s.ConfigValue
}

func (s *DescribeDBInstanceConfigResponseBodyData) GetDbInstanceName() *string {
	return s.DbInstanceName
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetConfigName(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.ConfigName = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetConfigValue(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.ConfigValue = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) SetDbInstanceName(v string) *DescribeDBInstanceConfigResponseBodyData {
	s.DbInstanceName = &v
	return s
}

func (s *DescribeDBInstanceConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
