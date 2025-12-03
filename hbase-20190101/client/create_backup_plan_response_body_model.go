// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateBackupPlanResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateBackupPlanResponseBody
	GetRequestId() *string
}

type CreateBackupPlanResponseBody struct {
	// example:
	//
	// 50373857-C47B-4B64-9332-D0B5280B59EA
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateBackupPlanResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateBackupPlanResponseBody) GoString() string {
	return s.String()
}

func (s *CreateBackupPlanResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateBackupPlanResponseBody) SetRequestId(v string) *CreateBackupPlanResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateBackupPlanResponseBody) Validate() error {
	return dara.Validate(s)
}
