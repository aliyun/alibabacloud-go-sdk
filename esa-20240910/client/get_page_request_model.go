// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetPageRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetPageRequest
	GetId() *int64
}

type GetPageRequest struct {
	// The ID of the custom error page, which can be obtained by calling the [ListPages](https://help.aliyun.com/document_detail/2850223.html) operation.
	//
	// example:
	//
	// 50000001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetPageRequest) String() string {
	return dara.Prettify(s)
}

func (s GetPageRequest) GoString() string {
	return s.String()
}

func (s *GetPageRequest) GetId() *int64 {
	return s.Id
}

func (s *GetPageRequest) SetId(v int64) *GetPageRequest {
	s.Id = &v
	return s
}

func (s *GetPageRequest) Validate() error {
	return dara.Validate(s)
}
