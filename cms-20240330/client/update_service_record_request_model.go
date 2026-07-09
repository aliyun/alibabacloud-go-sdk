// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateServiceRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordContent(v string) *UpdateServiceRecordRequest
	GetRecordContent() *string
	SetRecordType(v string) *UpdateServiceRecordRequest
	GetRecordType() *string
}

type UpdateServiceRecordRequest struct {
	// The entry content in JSON string format. The format varies depending on the recordType value.
	//
	// This parameter is required.
	//
	// example:
	//
	// {
	//
	//   "project": "proj-xtrace-16c988dcfe21fcb73c5e6f234927d998-cn-hangzhou",
	//
	//   "storeName": "app-biz-log",
	//
	//   "regionId": "cn-hangzhou",
	//
	//   "bindType": "logstore",
	//
	//   "traceIdRelateField": ""
	//
	// }
	RecordContent *string `json:"recordContent,omitempty" xml:"recordContent,omitempty"`
	// The type of the linked entry. Currently supported value:
	//
	// logCorrelation, which indicates application log association.
	//
	// This parameter is required.
	//
	// example:
	//
	// logCorrelation
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
}

func (s UpdateServiceRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateServiceRecordRequest) GoString() string {
	return s.String()
}

func (s *UpdateServiceRecordRequest) GetRecordContent() *string {
	return s.RecordContent
}

func (s *UpdateServiceRecordRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *UpdateServiceRecordRequest) SetRecordContent(v string) *UpdateServiceRecordRequest {
	s.RecordContent = &v
	return s
}

func (s *UpdateServiceRecordRequest) SetRecordType(v string) *UpdateServiceRecordRequest {
	s.RecordType = &v
	return s
}

func (s *UpdateServiceRecordRequest) Validate() error {
	return dara.Validate(s)
}
