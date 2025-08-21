// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyInstanceRemarkResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyInstanceRemarkResponseBody
	GetRequestId() *string
}

type ModifyInstanceRemarkResponseBody struct {
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 7EFA2BA6-9C0A-4410-B735-FC337EB634A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyInstanceRemarkResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyInstanceRemarkResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyInstanceRemarkResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyInstanceRemarkResponseBody) SetRequestId(v string) *ModifyInstanceRemarkResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyInstanceRemarkResponseBody) Validate() error {
	return dara.Validate(s)
}
