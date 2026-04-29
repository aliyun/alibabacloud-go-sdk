// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryRecallManagementTableRecordsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecords(v []map[string]interface{}) *QueryRecallManagementTableRecordsResponseBody
	GetRecords() []map[string]interface{}
	SetRequestId(v string) *QueryRecallManagementTableRecordsResponseBody
	GetRequestId() *string
}

type QueryRecallManagementTableRecordsResponseBody struct {
	Records []map[string]interface{} `json:"Records,omitempty" xml:"Records,omitempty" type:"Repeated"`
	// example:
	//
	// 7D59453C-48AA-5FC5-8848-2D373BD1A17F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryRecallManagementTableRecordsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryRecallManagementTableRecordsResponseBody) GoString() string {
	return s.String()
}

func (s *QueryRecallManagementTableRecordsResponseBody) GetRecords() []map[string]interface{} {
	return s.Records
}

func (s *QueryRecallManagementTableRecordsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryRecallManagementTableRecordsResponseBody) SetRecords(v []map[string]interface{}) *QueryRecallManagementTableRecordsResponseBody {
	s.Records = v
	return s
}

func (s *QueryRecallManagementTableRecordsResponseBody) SetRequestId(v string) *QueryRecallManagementTableRecordsResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryRecallManagementTableRecordsResponseBody) Validate() error {
	return dara.Validate(s)
}
