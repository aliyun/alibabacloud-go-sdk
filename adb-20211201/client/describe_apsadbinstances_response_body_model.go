// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeAPSADBInstancesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeAPSADBInstancesResponseBody
	GetCode() *string
	SetHttpStatusCode(v int32) *DescribeAPSADBInstancesResponseBody
	GetHttpStatusCode() *int32
	SetItems(v []*DescribeAPSADBInstancesResponseBodyItems) *DescribeAPSADBInstancesResponseBody
	GetItems() []*DescribeAPSADBInstancesResponseBodyItems
	SetMessage(v string) *DescribeAPSADBInstancesResponseBody
	GetMessage() *string
	SetPageNumber(v string) *DescribeAPSADBInstancesResponseBody
	GetPageNumber() *string
	SetPageSize(v string) *DescribeAPSADBInstancesResponseBody
	GetPageSize() *string
	SetRequestId(v string) *DescribeAPSADBInstancesResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeAPSADBInstancesResponseBody
	GetSuccess() *bool
	SetTotalCount(v string) *DescribeAPSADBInstancesResponseBody
	GetTotalCount() *string
}

type DescribeAPSADBInstancesResponseBody struct {
	// The HTTP status code or the error code.
	//
	// example:
	//
	// 0
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The HTTP status code.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The queried clusters.
	//
	// example:
	//
	// -
	Items []*DescribeAPSADBInstancesResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The returned message. Valid values:
	//
	// 	- If the request was successful, a success message is returned.****
	//
	// 	- If the request failed, an error message is returned.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The page number.
	//
	// example:
	//
	// 1
	PageNumber *string `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 30
	PageSize *string `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// ******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 100
	TotalCount *string `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s DescribeAPSADBInstancesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeAPSADBInstancesResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeAPSADBInstancesResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeAPSADBInstancesResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *DescribeAPSADBInstancesResponseBody) GetItems() []*DescribeAPSADBInstancesResponseBodyItems {
	return s.Items
}

func (s *DescribeAPSADBInstancesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeAPSADBInstancesResponseBody) GetPageNumber() *string {
	return s.PageNumber
}

func (s *DescribeAPSADBInstancesResponseBody) GetPageSize() *string {
	return s.PageSize
}

func (s *DescribeAPSADBInstancesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeAPSADBInstancesResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeAPSADBInstancesResponseBody) GetTotalCount() *string {
	return s.TotalCount
}

func (s *DescribeAPSADBInstancesResponseBody) SetCode(v string) *DescribeAPSADBInstancesResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBody) SetHttpStatusCode(v int32) *DescribeAPSADBInstancesResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBody) SetItems(v []*DescribeAPSADBInstancesResponseBodyItems) *DescribeAPSADBInstancesResponseBody {
	s.Items = v
	return s
}

func (s *DescribeAPSADBInstancesResponseBody) SetMessage(v string) *DescribeAPSADBInstancesResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBody) SetPageNumber(v string) *DescribeAPSADBInstancesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBody) SetPageSize(v string) *DescribeAPSADBInstancesResponseBody {
	s.PageSize = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBody) SetRequestId(v string) *DescribeAPSADBInstancesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBody) SetSuccess(v bool) *DescribeAPSADBInstancesResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBody) SetTotalCount(v string) *DescribeAPSADBInstancesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBody) Validate() error {
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

type DescribeAPSADBInstancesResponseBodyItems struct {
	// The specifications of the reserved computing resources.
	//
	// example:
	//
	// 16ACU
	ComputeResource *string `json:"ComputeResource,omitempty" xml:"ComputeResource,omitempty"`
	// The description of the cluster.
	//
	// example:
	//
	// adb_test
	DBClusterDescription *string `json:"DBClusterDescription,omitempty" xml:"DBClusterDescription,omitempty"`
	// The ID of the AnalyticDB for MySQL Data Lakehouse Edition cluster.
	//
	// example:
	//
	// am-bp1********
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// The status of the cluster.
	//
	// example:
	//
	// Running
	DBClusterStatus *string `json:"DBClusterStatus,omitempty" xml:"DBClusterStatus,omitempty"`
	// The amount of remaining reserved computing resources that are available in the cluster.
	//
	// example:
	//
	// 24ACU
	ReservedACU *string `json:"ReservedACU,omitempty" xml:"ReservedACU,omitempty"`
	// The specifications of the reserved storage resources.
	//
	// example:
	//
	// 24ACU
	StorageResource *int64 `json:"StorageResource,omitempty" xml:"StorageResource,omitempty"`
	// The zone ID of the cluster.
	//
	// example:
	//
	// cn-hangzhou-h
	ZoneId *string `json:"ZoneId,omitempty" xml:"ZoneId,omitempty"`
}

func (s DescribeAPSADBInstancesResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s DescribeAPSADBInstancesResponseBodyItems) GoString() string {
	return s.String()
}

func (s *DescribeAPSADBInstancesResponseBodyItems) GetComputeResource() *string {
	return s.ComputeResource
}

func (s *DescribeAPSADBInstancesResponseBodyItems) GetDBClusterDescription() *string {
	return s.DBClusterDescription
}

func (s *DescribeAPSADBInstancesResponseBodyItems) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeAPSADBInstancesResponseBodyItems) GetDBClusterStatus() *string {
	return s.DBClusterStatus
}

func (s *DescribeAPSADBInstancesResponseBodyItems) GetReservedACU() *string {
	return s.ReservedACU
}

func (s *DescribeAPSADBInstancesResponseBodyItems) GetStorageResource() *int64 {
	return s.StorageResource
}

func (s *DescribeAPSADBInstancesResponseBodyItems) GetZoneId() *string {
	return s.ZoneId
}

func (s *DescribeAPSADBInstancesResponseBodyItems) SetComputeResource(v string) *DescribeAPSADBInstancesResponseBodyItems {
	s.ComputeResource = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBodyItems) SetDBClusterDescription(v string) *DescribeAPSADBInstancesResponseBodyItems {
	s.DBClusterDescription = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBodyItems) SetDBClusterId(v string) *DescribeAPSADBInstancesResponseBodyItems {
	s.DBClusterId = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBodyItems) SetDBClusterStatus(v string) *DescribeAPSADBInstancesResponseBodyItems {
	s.DBClusterStatus = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBodyItems) SetReservedACU(v string) *DescribeAPSADBInstancesResponseBodyItems {
	s.ReservedACU = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBodyItems) SetStorageResource(v int64) *DescribeAPSADBInstancesResponseBodyItems {
	s.StorageResource = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBodyItems) SetZoneId(v string) *DescribeAPSADBInstancesResponseBodyItems {
	s.ZoneId = &v
	return s
}

func (s *DescribeAPSADBInstancesResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
