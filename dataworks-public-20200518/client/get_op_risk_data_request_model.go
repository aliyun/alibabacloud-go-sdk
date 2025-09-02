// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpRiskDataRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDate(v string) *GetOpRiskDataRequest
	GetDate() *string
	SetName(v string) *GetOpRiskDataRequest
	GetName() *string
	SetPageNo(v int32) *GetOpRiskDataRequest
	GetPageNo() *int32
	SetPageSize(v int32) *GetOpRiskDataRequest
	GetPageSize() *int32
	SetRiskType(v string) *GetOpRiskDataRequest
	GetRiskType() *string
}

type GetOpRiskDataRequest struct {
	// The date on which access records were generated. Specify the value in the yyyyMMdd format.
	//
	// This parameter is required.
	//
	// example:
	//
	// 20210221
	Date *string `json:"Date,omitempty" xml:"Date,omitempty"`
	// The parameters that you can configure to query the access records. Valid values:
	//
	// 	- dbType
	//
	// 	- instanceName
	//
	// 	- databaseName
	//
	// 	- projectName
	//
	// 	- clusterName
	//
	// The following example shows the parameters configured to query the access records of the sensitive data in the abc database of the Hologres instance ABC: [ {"dbType":"hologres","instanceName":"ABC","databaseName":"abc"} ]
	//
	// You must configure the parameters based on the compute engine that you use in your business.
	//
	// example:
	//
	// [ {"dbType":"hologres","instanceName":"ABC","databaseName":"abc"}, {"dbType":"ODPS.ODPS","projectName":"adbc"} ]
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The page number. Pages start from 1.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The number of entries per page. Maximum value: 1000.
	//
	// This parameter is required.
	//
	// example:
	//
	// 100
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The method that you use to identify risks.
	//
	// 	- You can manually identify risks.
	//
	// 	- You can also use a sensitive data identification rule to identify risks. You can log on to the DataWorks console and go to the Risk Identification Rules page in Data Security Guard to obtain the name of the rule.
	//
	// example:
	//
	// Manual identification
	RiskType *string `json:"RiskType,omitempty" xml:"RiskType,omitempty"`
}

func (s GetOpRiskDataRequest) String() string {
	return dara.Prettify(s)
}

func (s GetOpRiskDataRequest) GoString() string {
	return s.String()
}

func (s *GetOpRiskDataRequest) GetDate() *string {
	return s.Date
}

func (s *GetOpRiskDataRequest) GetName() *string {
	return s.Name
}

func (s *GetOpRiskDataRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *GetOpRiskDataRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *GetOpRiskDataRequest) GetRiskType() *string {
	return s.RiskType
}

func (s *GetOpRiskDataRequest) SetDate(v string) *GetOpRiskDataRequest {
	s.Date = &v
	return s
}

func (s *GetOpRiskDataRequest) SetName(v string) *GetOpRiskDataRequest {
	s.Name = &v
	return s
}

func (s *GetOpRiskDataRequest) SetPageNo(v int32) *GetOpRiskDataRequest {
	s.PageNo = &v
	return s
}

func (s *GetOpRiskDataRequest) SetPageSize(v int32) *GetOpRiskDataRequest {
	s.PageSize = &v
	return s
}

func (s *GetOpRiskDataRequest) SetRiskType(v string) *GetOpRiskDataRequest {
	s.RiskType = &v
	return s
}

func (s *GetOpRiskDataRequest) Validate() error {
	return dara.Validate(s)
}
