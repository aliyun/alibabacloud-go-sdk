// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListImageAssociatedProjectsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*int64) *ListImageAssociatedProjectsResponseBody
	GetData() []*int64
	SetRequestId(v string) *ListImageAssociatedProjectsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListImageAssociatedProjectsResponseBody
	GetSuccess() *bool
}

type ListImageAssociatedProjectsResponseBody struct {
	Data []*int64 `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// 6A6CBE87-9F91-1323-B680-E7A7065XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListImageAssociatedProjectsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListImageAssociatedProjectsResponseBody) GoString() string {
	return s.String()
}

func (s *ListImageAssociatedProjectsResponseBody) GetData() []*int64 {
	return s.Data
}

func (s *ListImageAssociatedProjectsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListImageAssociatedProjectsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListImageAssociatedProjectsResponseBody) SetData(v []*int64) *ListImageAssociatedProjectsResponseBody {
	s.Data = v
	return s
}

func (s *ListImageAssociatedProjectsResponseBody) SetRequestId(v string) *ListImageAssociatedProjectsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListImageAssociatedProjectsResponseBody) SetSuccess(v bool) *ListImageAssociatedProjectsResponseBody {
	s.Success = &v
	return s
}

func (s *ListImageAssociatedProjectsResponseBody) Validate() error {
	return dara.Validate(s)
}
