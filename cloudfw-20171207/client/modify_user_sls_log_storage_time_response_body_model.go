// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyUserSlsLogStorageTimeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyUserSlsLogStorageTimeResponseBody
	GetRequestId() *string
}

type ModifyUserSlsLogStorageTimeResponseBody struct {
	// example:
	//
	// 337A4DBA-8A01-5E9C-99CA-84293E13****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyUserSlsLogStorageTimeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyUserSlsLogStorageTimeResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyUserSlsLogStorageTimeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyUserSlsLogStorageTimeResponseBody) SetRequestId(v string) *ModifyUserSlsLogStorageTimeResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyUserSlsLogStorageTimeResponseBody) Validate() error {
	return dara.Validate(s)
}
