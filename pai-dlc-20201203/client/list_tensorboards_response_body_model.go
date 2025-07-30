// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTensorboardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListTensorboardsResponseBody
	GetRequestId() *string
	SetTensorboards(v []*Tensorboard) *ListTensorboardsResponseBody
	GetTensorboards() []*Tensorboard
	SetTotalCount(v int64) *ListTensorboardsResponseBody
	GetTotalCount() *int64
}

type ListTensorboardsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The TensorBoard instances.
	Tensorboards []*Tensorboard `json:"Tensorboards,omitempty" xml:"Tensorboards,omitempty" type:"Repeated"`
	// The total number of data sources that meet the conditions.
	//
	// example:
	//
	// 100
	TotalCount *int64 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTensorboardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListTensorboardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTensorboardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListTensorboardsResponseBody) GetTensorboards() []*Tensorboard {
	return s.Tensorboards
}

func (s *ListTensorboardsResponseBody) GetTotalCount() *int64 {
	return s.TotalCount
}

func (s *ListTensorboardsResponseBody) SetRequestId(v string) *ListTensorboardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTensorboardsResponseBody) SetTensorboards(v []*Tensorboard) *ListTensorboardsResponseBody {
	s.Tensorboards = v
	return s
}

func (s *ListTensorboardsResponseBody) SetTotalCount(v int64) *ListTensorboardsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTensorboardsResponseBody) Validate() error {
	return dara.Validate(s)
}
