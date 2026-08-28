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
	// The response data.
	//
	// example:
	//
	// 1.0.0
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A1B2C3D4-E5F6-47A8-90AB-CDEF12345678
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
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
