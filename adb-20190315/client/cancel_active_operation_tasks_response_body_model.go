// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelActiveOperationTasksResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetIds(v string) *CancelActiveOperationTasksResponseBody
	GetIds() *string
	SetRequestId(v string) *CancelActiveOperationTasksResponseBody
	GetRequestId() *string
}

type CancelActiveOperationTasksResponseBody struct {
	// The cluster ID.
	//
	// example:
	//
	// am-xxxxxxxx
	Ids *string `json:"Ids,omitempty" xml:"Ids,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 642F3512-C628-5D0C-8815-F6670CEA00D4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CancelActiveOperationTasksResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelActiveOperationTasksResponseBody) GoString() string {
	return s.String()
}

func (s *CancelActiveOperationTasksResponseBody) GetIds() *string {
	return s.Ids
}

func (s *CancelActiveOperationTasksResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelActiveOperationTasksResponseBody) SetIds(v string) *CancelActiveOperationTasksResponseBody {
	s.Ids = &v
	return s
}

func (s *CancelActiveOperationTasksResponseBody) SetRequestId(v string) *CancelActiveOperationTasksResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelActiveOperationTasksResponseBody) Validate() error {
	return dara.Validate(s)
}
