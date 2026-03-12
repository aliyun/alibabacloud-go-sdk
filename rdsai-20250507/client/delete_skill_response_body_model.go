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
	SetResult(v string) *DeleteSkillResponseBody
	GetResult() *string
}

type DeleteSkillResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// FE9C65D7-930F-57A5-A207-8C396329****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The returned results.
	//
	// example:
	//
	// success
	Result *string `json:"Result,omitempty" xml:"Result,omitempty"`
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

func (s *DeleteSkillResponseBody) GetResult() *string {
	return s.Result
}

func (s *DeleteSkillResponseBody) SetRequestId(v string) *DeleteSkillResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteSkillResponseBody) SetResult(v string) *DeleteSkillResponseBody {
	s.Result = &v
	return s
}

func (s *DeleteSkillResponseBody) Validate() error {
	return dara.Validate(s)
}
