// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDataSetBloodResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DataSetBloodResponseBody
	GetRequestId() *string
	SetResult(v []*DataSetBloodResponseBodyResult) *DataSetBloodResponseBody
	GetResult() []*DataSetBloodResponseBodyResult
	SetSuccess(v bool) *DataSetBloodResponseBody
	GetSuccess() *bool
}

type DataSetBloodResponseBody struct {
	// Request ID.
	//
	// example:
	//
	// 46e537a5****,3dadsu****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Array of works.
	Result []*DataSetBloodResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
	// Indicates whether the request was successful. Possible values:
	//
	// - true: Request succeeded
	//
	// - false: Request failed
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DataSetBloodResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DataSetBloodResponseBody) GoString() string {
	return s.String()
}

func (s *DataSetBloodResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DataSetBloodResponseBody) GetResult() []*DataSetBloodResponseBodyResult {
	return s.Result
}

func (s *DataSetBloodResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DataSetBloodResponseBody) SetRequestId(v string) *DataSetBloodResponseBody {
	s.RequestId = &v
	return s
}

func (s *DataSetBloodResponseBody) SetResult(v []*DataSetBloodResponseBodyResult) *DataSetBloodResponseBody {
	s.Result = v
	return s
}

func (s *DataSetBloodResponseBody) SetSuccess(v bool) *DataSetBloodResponseBody {
	s.Success = &v
	return s
}

func (s *DataSetBloodResponseBody) Validate() error {
	return dara.Validate(s)
}

type DataSetBloodResponseBodyResult struct {
	// Work ID.
	//
	// example:
	//
	// ccd3428c-****-****-a608-26bae29dffee
	WorksId *string `json:"WorksId,omitempty" xml:"WorksId,omitempty"`
	// Work types: - REPORT:
	//
	// - REPORT: Workbooks
	//
	// - dashboardOfflineQuery: Downloads
	//
	// - DASHBOARD: Dashboard
	//
	// - ANALYSIS: Ad Hoc Analysis
	//
	// - SCREEN: Visualization Screen
	//
	// - PAGE: Old dashboard
	//
	// example:
	//
	// PAGE
	WorksType *string `json:"WorksType,omitempty" xml:"WorksType,omitempty"`
}

func (s DataSetBloodResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DataSetBloodResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DataSetBloodResponseBodyResult) GetWorksId() *string {
	return s.WorksId
}

func (s *DataSetBloodResponseBodyResult) GetWorksType() *string {
	return s.WorksType
}

func (s *DataSetBloodResponseBodyResult) SetWorksId(v string) *DataSetBloodResponseBodyResult {
	s.WorksId = &v
	return s
}

func (s *DataSetBloodResponseBodyResult) SetWorksType(v string) *DataSetBloodResponseBodyResult {
	s.WorksType = &v
	return s
}

func (s *DataSetBloodResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
