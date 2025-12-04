// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCheckScaleOutBalancedResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCheckCode(v string) *CheckScaleOutBalancedResponseBody
	GetCheckCode() *string
	SetPageNumber(v int32) *CheckScaleOutBalancedResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *CheckScaleOutBalancedResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *CheckScaleOutBalancedResponseBody
	GetRequestId() *string
	SetTableDetails(v *CheckScaleOutBalancedResponseBodyTableDetails) *CheckScaleOutBalancedResponseBody
	GetTableDetails() *CheckScaleOutBalancedResponseBodyTableDetails
	SetTimeDuration(v string) *CheckScaleOutBalancedResponseBody
	GetTimeDuration() *string
	SetTotalCount(v int32) *CheckScaleOutBalancedResponseBody
	GetTotalCount() *int32
}

type CheckScaleOutBalancedResponseBody struct {
	// The check result. Valid values:
	//
	// 	- **400**: The cluster failed the check.
	//
	// 	- **200**: The cluster passed the check.
	//
	// example:
	//
	// 400
	CheckCode *string `json:"CheckCode,omitempty" xml:"CheckCode,omitempty"`
	// The total number of returned pages.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries returned per page. Valid values:
	//
	// 	- **30*	- (default)
	//
	// 	- **50**
	//
	// 	- **100**
	//
	// example:
	//
	// 30
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F5178C10-1407-4987-9133-DE4DC9119F75
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The error information returned for a check failure.
	TableDetails *CheckScaleOutBalancedResponseBodyTableDetails `json:"TableDetails,omitempty" xml:"TableDetails,omitempty" type:"Struct"`
	// The amount of time that is required for the migration and scale-out. Unit: minutes.
	//
	// example:
	//
	// 21
	TimeDuration *string `json:"TimeDuration,omitempty" xml:"TimeDuration,omitempty"`
	// The total number of entries that are returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s CheckScaleOutBalancedResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CheckScaleOutBalancedResponseBody) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedResponseBody) GetCheckCode() *string {
	return s.CheckCode
}

func (s *CheckScaleOutBalancedResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *CheckScaleOutBalancedResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *CheckScaleOutBalancedResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CheckScaleOutBalancedResponseBody) GetTableDetails() *CheckScaleOutBalancedResponseBodyTableDetails {
	return s.TableDetails
}

func (s *CheckScaleOutBalancedResponseBody) GetTimeDuration() *string {
	return s.TimeDuration
}

func (s *CheckScaleOutBalancedResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *CheckScaleOutBalancedResponseBody) SetCheckCode(v string) *CheckScaleOutBalancedResponseBody {
	s.CheckCode = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetPageNumber(v int32) *CheckScaleOutBalancedResponseBody {
	s.PageNumber = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetPageSize(v int32) *CheckScaleOutBalancedResponseBody {
	s.PageSize = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetRequestId(v string) *CheckScaleOutBalancedResponseBody {
	s.RequestId = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetTableDetails(v *CheckScaleOutBalancedResponseBodyTableDetails) *CheckScaleOutBalancedResponseBody {
	s.TableDetails = v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetTimeDuration(v string) *CheckScaleOutBalancedResponseBody {
	s.TimeDuration = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) SetTotalCount(v int32) *CheckScaleOutBalancedResponseBody {
	s.TotalCount = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBody) Validate() error {
	if s.TableDetails != nil {
		if err := s.TableDetails.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CheckScaleOutBalancedResponseBodyTableDetails struct {
	TableDetail []*CheckScaleOutBalancedResponseBodyTableDetailsTableDetail `json:"TableDetail,omitempty" xml:"TableDetail,omitempty" type:"Repeated"`
}

func (s CheckScaleOutBalancedResponseBodyTableDetails) String() string {
	return dara.Prettify(s)
}

func (s CheckScaleOutBalancedResponseBodyTableDetails) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedResponseBodyTableDetails) GetTableDetail() []*CheckScaleOutBalancedResponseBodyTableDetailsTableDetail {
	return s.TableDetail
}

func (s *CheckScaleOutBalancedResponseBodyTableDetails) SetTableDetail(v []*CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) *CheckScaleOutBalancedResponseBodyTableDetails {
	s.TableDetail = v
	return s
}

func (s *CheckScaleOutBalancedResponseBodyTableDetails) Validate() error {
	if s.TableDetail != nil {
		for _, item := range s.TableDetail {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CheckScaleOutBalancedResponseBodyTableDetailsTableDetail struct {
	// The cluster. The value is fixed as **default**.
	//
	// example:
	//
	// default
	Cluster *string `json:"Cluster,omitempty" xml:"Cluster,omitempty"`
	// The database name.
	//
	// example:
	//
	// db_name
	Database *string `json:"Database,omitempty" xml:"Database,omitempty"`
	// The error details. Valid values:
	//
	// 	- **1**: The unique distributed table is missing.
	//
	// 	- **2**: More than one distributed table exists for the local table.
	//
	// example:
	//
	// 1
	Detail *int32 `json:"Detail,omitempty" xml:"Detail,omitempty"`
	// The name of the local table.
	//
	// example:
	//
	// test
	TableName *string `json:"TableName,omitempty" xml:"TableName,omitempty"`
}

func (s CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) String() string {
	return dara.Prettify(s)
}

func (s CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) GoString() string {
	return s.String()
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) GetCluster() *string {
	return s.Cluster
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) GetDatabase() *string {
	return s.Database
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) GetDetail() *int32 {
	return s.Detail
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) GetTableName() *string {
	return s.TableName
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) SetCluster(v string) *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail {
	s.Cluster = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) SetDatabase(v string) *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail {
	s.Database = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) SetDetail(v int32) *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail {
	s.Detail = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) SetTableName(v string) *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail {
	s.TableName = &v
	return s
}

func (s *CheckScaleOutBalancedResponseBodyTableDetailsTableDetail) Validate() error {
	return dara.Validate(s)
}
