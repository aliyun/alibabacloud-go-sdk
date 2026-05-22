// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetListRequest
	GetId() *int64
}

type GetListRequest struct {
	// The ID of the custom list, which can be obtained by calling the [ListLists](https://help.aliyun.com/document_detail/2850217.html) operation.
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetListRequest) String() string {
	return dara.Prettify(s)
}

func (s GetListRequest) GoString() string {
	return s.String()
}

func (s *GetListRequest) GetId() *int64 {
	return s.Id
}

func (s *GetListRequest) SetId(v int64) *GetListRequest {
	s.Id = &v
	return s
}

func (s *GetListRequest) Validate() error {
	return dara.Validate(s)
}
