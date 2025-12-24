// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarRecordInOutputResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInOutputInfo(v string) *DescribeSoarRecordInOutputResponseBody
	GetInOutputInfo() *string
	SetRequestId(v string) *DescribeSoarRecordInOutputResponseBody
	GetRequestId() *string
}

type DescribeSoarRecordInOutputResponseBody struct {
	// The execution result of the component action.
	//
	// example:
	//
	// {
	//
	//     "actionUuid": "3896a25d-4967-493c-942e-4e60f27da1f7-xxxxx",
	//
	//     "outputSummary": {
	//
	//         "datalist": [
	//
	//             {
	//
	//                 "a": "a"
	//
	//             }
	//
	//         ],
	//
	//         "total_data_successful": 1,
	//
	//         "total_data": 1,
	//
	//         "total_exe_successful": 1,
	//
	//         "total_exe": 1,
	//
	//         "total_data_with_dup": 1,
	//
	//         "status": true
	//
	//     },
	//
	//     "outputSchema": {
	//
	//         "a": "String",
	//
	//         "startTime": "DateTime"
	//
	//     },
	//
	//     "inputParams": {
	//
	//         "inputData": [
	//
	//             {
	//
	//                 "outputFields": {
	//
	//                     "a": "a"
	//
	//                 }
	//
	//             }
	//
	//         ],
	//
	//         "totalSize": 1
	//
	//     },
	//
	//     "startTime": "2023-11-13 17:47:28.645",
	//
	//     "taskName": "92af3c79-1754-4646-9366-9ddbxxxxx"
	//
	// }
	InOutputInfo *string `json:"InOutputInfo,omitempty" xml:"InOutputInfo,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 372D8B41-AF8D-573A-9B3F-0924950F241F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeSoarRecordInOutputResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarRecordInOutputResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSoarRecordInOutputResponseBody) GetInOutputInfo() *string {
	return s.InOutputInfo
}

func (s *DescribeSoarRecordInOutputResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSoarRecordInOutputResponseBody) SetInOutputInfo(v string) *DescribeSoarRecordInOutputResponseBody {
	s.InOutputInfo = &v
	return s
}

func (s *DescribeSoarRecordInOutputResponseBody) SetRequestId(v string) *DescribeSoarRecordInOutputResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSoarRecordInOutputResponseBody) Validate() error {
	return dara.Validate(s)
}
