// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpRiskDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *GetOpRiskDataResponseBody
	GetRequestId() *string
	SetRiskData(v string) *GetOpRiskDataResponseBody
	GetRiskData() *string
}

type GetOpRiskDataResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The information about the high-risk sensitive data. The information includes totalCount and opRiskDatas. opRiskDatas includes the following parameters:
	//
	// 	- sensType: the type of the sensitive data
	//
	// 	- sensLevel: the sensitivity level of the sensitive data
	//
	// 	- opType: the type of the operation
	//
	// 	- sql: the SQL statement that is executed
	//
	// 	- opAccount: the account that is used to perform the operation
	//
	// 	- opTime: the time when the operation was performed
	//
	// example:
	//
	// {
	//
	//       "opRiskDatas": [
	//
	//             {
	//
	//                   "riskType": "Hierarchical dimension, EMR engine dimension and project dimension, EMR engine project dimension operation data, export method dimension, EMR engine",
	//
	//                   "opTime": "2021-01-04 23:39:13",
	//
	//                   "opType": "SQL_SELECT",
	//
	//                   "opAccount": "user",
	//
	//                   "sensType": "Email/name/mobile phone number",
	//
	//                   "sql": "SELECT 	- FROM default.jiade_1219_test_create LIMIT 20"
	//
	//             }
	//
	//       ],
	//
	//       "totalCount": 499
	//
	// }
	RiskData *string `json:"RiskData,omitempty" xml:"RiskData,omitempty"`
}

func (s GetOpRiskDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpRiskDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpRiskDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpRiskDataResponseBody) GetRiskData() *string {
	return s.RiskData
}

func (s *GetOpRiskDataResponseBody) SetRequestId(v string) *GetOpRiskDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpRiskDataResponseBody) SetRiskData(v string) *GetOpRiskDataResponseBody {
	s.RiskData = &v
	return s
}

func (s *GetOpRiskDataResponseBody) Validate() error {
	return dara.Validate(s)
}
