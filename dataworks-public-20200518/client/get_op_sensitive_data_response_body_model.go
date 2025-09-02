// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetOpSensitiveDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOpSensitiveData(v string) *GetOpSensitiveDataResponseBody
	GetOpSensitiveData() *string
	SetRequestId(v string) *GetOpSensitiveDataResponseBody
	GetRequestId() *string
}

type GetOpSensitiveDataResponseBody struct {
	// The information about the access records of the sensitive data. The information includes totalCount and opRiskDatas. opRiskDatas includes the following parameters:
	//
	// 	- sensType: the type of the sensitive data.
	//
	// 	- sensLevel: the sensitivity level of the sensitive data. A larger value indicates a higher sensitivity level.
	//
	// 	- opType: the type of the operation.
	//
	// 	- sql: the SQL statement that is executed.
	//
	// 	- opAccount: the account that is used to perform the operation.
	//
	// 	- opTime: the time when the operation was performed.
	//
	// example:
	//
	// "opSensDatas": [       {         "sensLevel": "L4",         "opTime": "2021-02-07 00:14:51",         "opAccount": "ALIYUN$dsg_test",         "sensType": "Mobile phone number",         "sql": "select 	- from dsg_demo.tbl_phonebook where phone_no = &#39;1331111****&#39;;"       }     ],     "totalCount": 6
	OpSensitiveData *string `json:"OpSensitiveData,omitempty" xml:"OpSensitiveData,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0000-ABCD-EFG****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetOpSensitiveDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetOpSensitiveDataResponseBody) GoString() string {
	return s.String()
}

func (s *GetOpSensitiveDataResponseBody) GetOpSensitiveData() *string {
	return s.OpSensitiveData
}

func (s *GetOpSensitiveDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetOpSensitiveDataResponseBody) SetOpSensitiveData(v string) *GetOpSensitiveDataResponseBody {
	s.OpSensitiveData = &v
	return s
}

func (s *GetOpSensitiveDataResponseBody) SetRequestId(v string) *GetOpSensitiveDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetOpSensitiveDataResponseBody) Validate() error {
	return dara.Validate(s)
}
