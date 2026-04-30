// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyJVSInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyJVSInstanceResponseBody
	GetRequestId() *string
}

type ModifyJVSInstanceResponseBody struct {
	// example:
	//
	// 6C8439B9-7DBF-57F4-92AE-55A9B9D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyJVSInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyJVSInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyJVSInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyJVSInstanceResponseBody) SetRequestId(v string) *ModifyJVSInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyJVSInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
