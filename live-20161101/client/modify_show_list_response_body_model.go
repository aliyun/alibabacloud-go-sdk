// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyShowListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyShowListResponseBody
	GetRequestId() *string
}

type ModifyShowListResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 16A96B9A-F203-4EC5-8E43-CB92E68F4CD8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyShowListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyShowListResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyShowListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyShowListResponseBody) SetRequestId(v string) *ModifyShowListResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyShowListResponseBody) Validate() error {
	return dara.Validate(s)
}
