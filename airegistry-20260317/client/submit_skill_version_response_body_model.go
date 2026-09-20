// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSubmitSkillVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *SubmitSkillVersionResponseBody
	GetData() *string
	SetRequestId(v string) *SubmitSkillVersionResponseBody
	GetRequestId() *string
}

type SubmitSkillVersionResponseBody struct {
	// The skill version.
	//
	// example:
	//
	// ba9b5c2466dc408c9fcd9df72bcd762a
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// BE66410A-37F8-55C5-8471-589CA195760C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s SubmitSkillVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SubmitSkillVersionResponseBody) GoString() string {
	return s.String()
}

func (s *SubmitSkillVersionResponseBody) GetData() *string {
	return s.Data
}

func (s *SubmitSkillVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SubmitSkillVersionResponseBody) SetData(v string) *SubmitSkillVersionResponseBody {
	s.Data = &v
	return s
}

func (s *SubmitSkillVersionResponseBody) SetRequestId(v string) *SubmitSkillVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *SubmitSkillVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
