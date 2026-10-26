// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyDBInstanceDeletionProtectionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyDBInstanceDeletionProtectionResponseBody
	GetRequestId() *string
}

type ModifyDBInstanceDeletionProtectionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 3C5CFDEE-F774-4DED-89A2-1D76EC63C575
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyDBInstanceDeletionProtectionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyDBInstanceDeletionProtectionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyDBInstanceDeletionProtectionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyDBInstanceDeletionProtectionResponseBody) SetRequestId(v string) *ModifyDBInstanceDeletionProtectionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyDBInstanceDeletionProtectionResponseBody) Validate() error {
	return dara.Validate(s)
}
