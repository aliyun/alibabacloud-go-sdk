// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteSkillResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteSkillResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteSkillResponseBody
	GetSuccess() *bool
}

type DeleteSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Returns `true` when both the logical deletion of the Skill and the optional logical archiving of the Artifact are complete.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteSkillResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteSkillResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteSkillResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteSkillResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteSkillResponseBody) SetRequestId(v string) *DeleteSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillResponseBody) SetSuccess(v bool) *DeleteSkillResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
