// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDIJobResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDIJobId(v int64) *CreateDIJobResponseBody
	GetDIJobId() *int64
	SetId(v int64) *CreateDIJobResponseBody
	GetId() *int64
	SetRequestId(v string) *CreateDIJobResponseBody
	GetRequestId() *string
}

type CreateDIJobResponseBody struct {
	// Deprecated
	//
	// This parameter is deprecated and is replaced by the Id parameter.
	//
	// example:
	//
	// 11792
	DIJobId *int64 `json:"DIJobId,omitempty" xml:"DIJobId,omitempty"`
	// The ID of the synchronization task.
	//
	// example:
	//
	// 11792
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The request ID. You can use the ID to query logs and troubleshoot issues.
	//
	// example:
	//
	// 4F6AB6B3-41FB-5EBB-AFB2-0C98D49DA2BB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDIJobResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDIJobResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDIJobResponseBody) GetDIJobId() *int64 {
	return s.DIJobId
}

func (s *CreateDIJobResponseBody) GetId() *int64 {
	return s.Id
}

func (s *CreateDIJobResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDIJobResponseBody) SetDIJobId(v int64) *CreateDIJobResponseBody {
	s.DIJobId = &v
	return s
}

func (s *CreateDIJobResponseBody) SetId(v int64) *CreateDIJobResponseBody {
	s.Id = &v
	return s
}

func (s *CreateDIJobResponseBody) SetRequestId(v string) *CreateDIJobResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDIJobResponseBody) Validate() error {
	return dara.Validate(s)
}
