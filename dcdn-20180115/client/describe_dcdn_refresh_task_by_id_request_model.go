// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDcdnRefreshTaskByIdRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTaskId(v string) *DescribeDcdnRefreshTaskByIdRequest
	GetTaskId() *string
}

type DescribeDcdnRefreshTaskByIdRequest struct {
	// The ID of the task that you want to query. The following signature algorithms require different message digest algorithms:
	//
	// 	- Perform the [RefreshDcdnObjectCaches](https://help.aliyun.com/document_detail/130620.html) operation to query refresh task IDs.
	//
	// 	- Perform the [PreloadDcdnObjectCaches](https://help.aliyun.com/document_detail/130636.html) operation to query prefetch task IDs.
	//
	// > You can specify at most 10 task IDs in each call. Separate IDs with commas (,).
	//
	// This parameter is required.
	//
	// example:
	//
	// 113681**
	TaskId *string `json:"TaskId,omitempty" xml:"TaskId,omitempty"`
}

func (s DescribeDcdnRefreshTaskByIdRequest) String() string {
	return dara.Prettify(s)
}

func (s DescribeDcdnRefreshTaskByIdRequest) GoString() string {
	return s.String()
}

func (s *DescribeDcdnRefreshTaskByIdRequest) GetTaskId() *string {
	return s.TaskId
}

func (s *DescribeDcdnRefreshTaskByIdRequest) SetTaskId(v string) *DescribeDcdnRefreshTaskByIdRequest {
	s.TaskId = &v
	return s
}

func (s *DescribeDcdnRefreshTaskByIdRequest) Validate() error {
	return dara.Validate(s)
}
