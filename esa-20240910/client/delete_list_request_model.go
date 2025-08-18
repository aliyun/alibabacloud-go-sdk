// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteListRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeleteListRequest
	GetId() *int64
}

type DeleteListRequest struct {
	// The ID of the custom list, which can be obtained by calling the [ListLists](https://help.aliyun.com/document_detail/2850217.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteListRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteListRequest) GoString() string {
	return s.String()
}

func (s *DeleteListRequest) GetId() *int64 {
	return s.Id
}

func (s *DeleteListRequest) SetId(v int64) *DeleteListRequest {
	s.Id = &v
	return s
}

func (s *DeleteListRequest) Validate() error {
	return dara.Validate(s)
}
