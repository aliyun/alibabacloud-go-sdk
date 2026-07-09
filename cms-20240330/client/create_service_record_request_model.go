// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetRecordContent(v string) *CreateServiceRecordRequest
	GetRecordContent() *string
	SetRecordType(v string) *CreateServiceRecordRequest
	GetRecordType() *string
}

type CreateServiceRecordRequest struct {
	// The entry content in JSON string format. The format may vary depending on the value of recordType.
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
	// The type of the linked entry. Valid values:
	//
	// - logCorrelation: application log association.
	//
	// This parameter is required.
	//
	// example:
	//
	// logCorrelation
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
}

func (s CreateServiceRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRecordRequest) GoString() string {
	return s.String()
}

func (s *CreateServiceRecordRequest) GetRecordContent() *string {
	return s.RecordContent
}

func (s *CreateServiceRecordRequest) GetRecordType() *string {
	return s.RecordType
}

func (s *CreateServiceRecordRequest) SetRecordContent(v string) *CreateServiceRecordRequest {
	s.RecordContent = &v
	return s
}

func (s *CreateServiceRecordRequest) SetRecordType(v string) *CreateServiceRecordRequest {
	s.RecordType = &v
	return s
}

func (s *CreateServiceRecordRequest) Validate() error {
	return dara.Validate(s)
}
