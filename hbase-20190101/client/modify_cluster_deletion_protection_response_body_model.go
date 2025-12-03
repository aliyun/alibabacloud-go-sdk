// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyClusterDeletionProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyClusterDeletionProtectionResponseBody
	GetRequestId() *string
}

type ModifyClusterDeletionProtectionResponseBody struct {
	// example:
	//
	// 24C80BD8-C710-4138-893A-D2AFED4FC13D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyClusterDeletionProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyClusterDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyClusterDeletionProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyClusterDeletionProtectionResponseBody) SetRequestId(v string) *ModifyClusterDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyClusterDeletionProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
