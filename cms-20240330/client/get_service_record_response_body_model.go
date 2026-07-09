// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetServiceRecordResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecord(v *GetServiceRecordResponseBodyRecord) *GetServiceRecordResponseBody
	GetRecord() *GetServiceRecordResponseBodyRecord
	SetRequestId(v string) *GetServiceRecordResponseBody
	GetRequestId() *string
}

type GetServiceRecordResponseBody struct {
	// The record.
	Record *GetServiceRecordResponseBodyRecord `json:"record,omitempty" xml:"record,omitempty" type:"Struct"`
	// Id of the request
	//
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetServiceRecordResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRecordResponseBody) GoString() string {
	return s.String()
}

func (s *GetServiceRecordResponseBody) GetRecord() *GetServiceRecordResponseBodyRecord {
	return s.Record
}

func (s *GetServiceRecordResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetServiceRecordResponseBody) SetRecord(v *GetServiceRecordResponseBodyRecord) *GetServiceRecordResponseBody {
	s.Record = v
	return s
}

func (s *GetServiceRecordResponseBody) SetRequestId(v string) *GetServiceRecordResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetServiceRecordResponseBody) Validate() error {
	if s.Record != nil {
		if err := s.Record.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetServiceRecordResponseBodyRecord struct {
	// The entry content in JSON string format. The returned content varies depending on the recordType.
	//
	// example:
	//
	// {"project":"proj-xtrace-16c988dcfe21fcb73c5e6f234927d998-cn-hangzhou","storeName":"app-biz-log","regionId":"cn-hangzhou","bindType":"logstore","traceIdRelateField":""}
	RecordContent *string `json:"recordContent,omitempty" xml:"recordContent,omitempty"`
	// The type of the linked entry. Currently supported values:
	//
	// logCorrelation: indicates application log association.
	//
	// example:
	//
	// logCorrelation
	RecordType *string `json:"recordType,omitempty" xml:"recordType,omitempty"`
	// The unique identifier of the service.
	//
	// example:
	//
	// ckj0xn6ma3@b96491402f8e1f15a8c79
	ServiceId *string `json:"serviceId,omitempty" xml:"serviceId,omitempty"`
	// The workspace.
	//
	// example:
	//
	// default-cms-1610600919225911-cn-beijing
	Workspace *string `json:"workspace,omitempty" xml:"workspace,omitempty"`
}

func (s GetServiceRecordResponseBodyRecord) String() string {
	return dara.Prettify(s)
}

func (s GetServiceRecordResponseBodyRecord) GoString() string {
	return s.String()
}

func (s *GetServiceRecordResponseBodyRecord) GetRecordContent() *string {
	return s.RecordContent
}

func (s *GetServiceRecordResponseBodyRecord) GetRecordType() *string {
	return s.RecordType
}

func (s *GetServiceRecordResponseBodyRecord) GetServiceId() *string {
	return s.ServiceId
}

func (s *GetServiceRecordResponseBodyRecord) GetWorkspace() *string {
	return s.Workspace
}

func (s *GetServiceRecordResponseBodyRecord) SetRecordContent(v string) *GetServiceRecordResponseBodyRecord {
	s.RecordContent = &v
	return s
}

func (s *GetServiceRecordResponseBodyRecord) SetRecordType(v string) *GetServiceRecordResponseBodyRecord {
	s.RecordType = &v
	return s
}

func (s *GetServiceRecordResponseBodyRecord) SetServiceId(v string) *GetServiceRecordResponseBodyRecord {
	s.ServiceId = &v
	return s
}

func (s *GetServiceRecordResponseBodyRecord) SetWorkspace(v string) *GetServiceRecordResponseBodyRecord {
	s.Workspace = &v
	return s
}

func (s *GetServiceRecordResponseBodyRecord) Validate() error {
	return dara.Validate(s)
}
