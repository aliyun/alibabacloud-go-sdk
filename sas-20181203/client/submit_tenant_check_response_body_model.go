// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitTenantCheckResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *SubmitTenantCheckResponseBodyData) *SubmitTenantCheckResponseBody
	GetData() *SubmitTenantCheckResponseBodyData
	SetRequestId(v string) *SubmitTenantCheckResponseBody
	GetRequestId() *string
}

type SubmitTenantCheckResponseBody struct {
	// The data returned.
	Data *SubmitTenantCheckResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 1162D670-E633-5676-AE87-8359B066****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitTenantCheckResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitTenantCheckResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitTenantCheckResponseBody) GetData() *SubmitTenantCheckResponseBodyData {
	return s.Data
}

func (s *SubmitTenantCheckResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitTenantCheckResponseBody) SetData(v *SubmitTenantCheckResponseBodyData) *SubmitTenantCheckResponseBody {
	s.Data = v
	return s
}

func (s *SubmitTenantCheckResponseBody) SetRequestId(v string) *SubmitTenantCheckResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitTenantCheckResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SubmitTenantCheckResponseBodyData struct {
	// The ID of the scan task.
	//
	// example:
	//
	// fc98d58eb56f699d49bf7ebbd6d7****
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s SubmitTenantCheckResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SubmitTenantCheckResponseBodyData) GoString() string {
	return s.String()
}

func (s *SubmitTenantCheckResponseBodyData) GetTaskId() *string {
	return s.TaskId
}

func (s *SubmitTenantCheckResponseBodyData) SetTaskId(v string) *SubmitTenantCheckResponseBodyData {
	s.TaskId = &v
	return s
}

func (s *SubmitTenantCheckResponseBodyData) Validate() error {
	return dara.Validate(s)
}
