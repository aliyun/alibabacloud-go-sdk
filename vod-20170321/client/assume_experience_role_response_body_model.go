// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAssumeExperienceRoleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AssumeExperienceRoleResponseBody
	GetRequestId() *string
	SetResult(v string) *AssumeExperienceRoleResponseBody
	GetResult() *string
}

type AssumeExperienceRoleResponseBody struct {
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Result    *string `json:"Result,omitempty" xml:"Result,omitempty"`
}

func (s AssumeExperienceRoleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AssumeExperienceRoleResponseBody) GoString() string {
	return s.String()
}

func (s *AssumeExperienceRoleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AssumeExperienceRoleResponseBody) GetResult() *string {
	return s.Result
}

func (s *AssumeExperienceRoleResponseBody) SetRequestId(v string) *AssumeExperienceRoleResponseBody {
	s.RequestId = &v
	return s
}

func (s *AssumeExperienceRoleResponseBody) SetResult(v string) *AssumeExperienceRoleResponseBody {
	s.Result = &v
	return s
}

func (s *AssumeExperienceRoleResponseBody) Validate() error {
	return dara.Validate(s)
}
