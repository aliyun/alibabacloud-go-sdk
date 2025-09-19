// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSoarPlaybookTaskDetailRequest interface {
	dara.Model
	String() string
	GoString() string
	SetPlaybookId(v int64) *DescribeSoarPlaybookTaskDetailRequest
	GetPlaybookId() *int64
	SetRecordId(v int64) *DescribeSoarPlaybookTaskDetailRequest
	GetRecordId() *int64
	SetRequestUuid(v string) *DescribeSoarPlaybookTaskDetailRequest
	GetRequestUuid() *string
}

type DescribeSoarPlaybookTaskDetailRequest struct {
	// Playbook ID.
	//
	// > You can obtain this parameter by calling the [DescribePlaybooks](https://help.aliyun.com/document_detail/2627461.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 123
	PlaybookId *int64 `json:"PlaybookId,omitempty" xml:"PlaybookId,omitempty"`
	// The vulnerability ID passed when creating the policy task.
	//
	// > You can obtain this parameter by calling the [DescribeVulList](~~DescribeVulList~~) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// 14
	RecordId *int64 `json:"RecordId,omitempty" xml:"RecordId,omitempty"`
	// UUID of the playbook task execution.
	//
	// > You can obtain this parameter by calling the [DescribeSoarRecords](https://help.aliyun.com/document_detail/2627455.html) interface.
	//
	// This parameter is required.
	//
	// example:
	//
	// a15e37da-abe0-4d87-acd2-024e875a****
	RequestUuid *string `json:"RequestUuid,omitempty" xml:"RequestUuid,omitempty"`
}

func (s DescribeSoarPlaybookTaskDetailRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeSoarPlaybookTaskDetailRequest) GoString() string {
	return s.String()
}

func (s *DescribeSoarPlaybookTaskDetailRequest) GetPlaybookId() *int64 {
	return s.PlaybookId
}

func (s *DescribeSoarPlaybookTaskDetailRequest) GetRecordId() *int64 {
	return s.RecordId
}

func (s *DescribeSoarPlaybookTaskDetailRequest) GetRequestUuid() *string {
	return s.RequestUuid
}

func (s *DescribeSoarPlaybookTaskDetailRequest) SetPlaybookId(v int64) *DescribeSoarPlaybookTaskDetailRequest {
	s.PlaybookId = &v
	return s
}

func (s *DescribeSoarPlaybookTaskDetailRequest) SetRecordId(v int64) *DescribeSoarPlaybookTaskDetailRequest {
	s.RecordId = &v
	return s
}

func (s *DescribeSoarPlaybookTaskDetailRequest) SetRequestUuid(v string) *DescribeSoarPlaybookTaskDetailRequest {
	s.RequestUuid = &v
	return s
}

func (s *DescribeSoarPlaybookTaskDetailRequest) Validate() error {
	return dara.Validate(s)
}
