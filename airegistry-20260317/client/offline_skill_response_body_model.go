// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iOfflineSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *OfflineSkillResponseBody
	GetRequestId() *string
}

type OfflineSkillResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s OfflineSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s OfflineSkillResponseBody) GoString() string {
	return s.String()
}

func (s *OfflineSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *OfflineSkillResponseBody) SetRequestId(v string) *OfflineSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *OfflineSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
