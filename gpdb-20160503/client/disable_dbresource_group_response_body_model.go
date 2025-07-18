// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDBResourceGroupResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DisableDBResourceGroupResponseBody
	GetRequestId() *string
}

type DisableDBResourceGroupResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 53EA07B7-FC2A-521B-AB7C-27**********
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableDBResourceGroupResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDBResourceGroupResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDBResourceGroupResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDBResourceGroupResponseBody) SetRequestId(v string) *DisableDBResourceGroupResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDBResourceGroupResponseBody) Validate() error {
	return dara.Validate(s)
}
