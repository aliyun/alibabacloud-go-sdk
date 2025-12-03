// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyBackupPlanConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyBackupPlanConfigResponseBody
	GetRequestId() *string
}

type ModifyBackupPlanConfigResponseBody struct {
	// example:
	//
	// 50F4A8C2-076F-4703-9813-2FCD7FBB91C2
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyBackupPlanConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyBackupPlanConfigResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyBackupPlanConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyBackupPlanConfigResponseBody) SetRequestId(v string) *ModifyBackupPlanConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyBackupPlanConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
