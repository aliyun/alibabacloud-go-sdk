// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSampleDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBInstanceId(v string) *DescribeSampleDataResponseBody
	GetDBInstanceId() *string
	SetErrorMessage(v string) *DescribeSampleDataResponseBody
	GetErrorMessage() *string
	SetHasSampleData(v bool) *DescribeSampleDataResponseBody
	GetHasSampleData() *bool
	SetRequestId(v string) *DescribeSampleDataResponseBody
	GetRequestId() *string
	SetSampleDataStatus(v string) *DescribeSampleDataResponseBody
	GetSampleDataStatus() *string
}

type DescribeSampleDataResponseBody struct {
	// The ID of the instance.
	//
	// example:
	//
	// gp-bp12ga6v69h86****
	DBInstanceId *string `json:"DBInstanceId,omitempty" xml:"DBInstanceId,omitempty"`
	// The error message returned if an error occurs. This message does not affect the execution of the operation.
	//
	// example:
	//
	// ******
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// Indicates whether a sample dataset is loaded to the instance. Valid values:
	//
	// 	- **true**: A sample dataset is loaded.
	//
	// 	- **false**: No sample dataset is loaded.
	//
	// example:
	//
	// true
	HasSampleData *bool `json:"HasSampleData,omitempty" xml:"HasSampleData,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 84CD7CAF-FA7B-5178-B19F-D8CDE307D5FA_8111
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The loading status of the sample dataset. Valid values:
	//
	// 	- **loaded**
	//
	// 	- **loading**
	//
	// 	- **unload**
	//
	// example:
	//
	// loaded
	SampleDataStatus *string `json:"SampleDataStatus,omitempty" xml:"SampleDataStatus,omitempty"`
}

func (s DescribeSampleDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSampleDataResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSampleDataResponseBody) GetDBInstanceId() *string {
	return s.DBInstanceId
}

func (s *DescribeSampleDataResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeSampleDataResponseBody) GetHasSampleData() *bool {
	return s.HasSampleData
}

func (s *DescribeSampleDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSampleDataResponseBody) GetSampleDataStatus() *string {
	return s.SampleDataStatus
}

func (s *DescribeSampleDataResponseBody) SetDBInstanceId(v string) *DescribeSampleDataResponseBody {
	s.DBInstanceId = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetErrorMessage(v string) *DescribeSampleDataResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetHasSampleData(v bool) *DescribeSampleDataResponseBody {
	s.HasSampleData = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetRequestId(v string) *DescribeSampleDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSampleDataResponseBody) SetSampleDataStatus(v string) *DescribeSampleDataResponseBody {
	s.SampleDataStatus = &v
	return s
}

func (s *DescribeSampleDataResponseBody) Validate() error {
	return dara.Validate(s)
}
