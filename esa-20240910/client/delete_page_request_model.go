// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeletePageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *DeletePageRequest
	GetId() *int64
}

type DeletePageRequest struct {
	// The ID of the custom error page, which can be obtained by calling the [ListPages](https://help.aliyun.com/document_detail/2850223.html) operation.
	//
	// This parameter is required.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeletePageRequest) String() string {
	return dara.Prettify(s)
}

func (s DeletePageRequest) GoString() string {
	return s.String()
}

func (s *DeletePageRequest) GetId() *int64 {
	return s.Id
}

func (s *DeletePageRequest) SetId(v int64) *DeletePageRequest {
	s.Id = &v
	return s
}

func (s *DeletePageRequest) Validate() error {
	return dara.Validate(s)
}
