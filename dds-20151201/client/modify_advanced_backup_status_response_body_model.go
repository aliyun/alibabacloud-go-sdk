// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAdvancedBackupStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAdvancedBackupStatusResponseBody
	GetRequestId() *string
}

type ModifyAdvancedBackupStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 45D2B592-DEBA-4347-BBF3-47FF6C97DBBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAdvancedBackupStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAdvancedBackupStatusResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAdvancedBackupStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAdvancedBackupStatusResponseBody) SetRequestId(v string) *ModifyAdvancedBackupStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAdvancedBackupStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
