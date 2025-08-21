// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateConnQuestionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOutlineId(v int64) *CreateConnQuestionResponseBody
	GetOutlineId() *int64
	SetRequestId(v string) *CreateConnQuestionResponseBody
	GetRequestId() *string
}

type CreateConnQuestionResponseBody struct {
	// example:
	//
	// 1000002123
	OutlineId *int64 `json:"OutlineId,omitempty" xml:"OutlineId,omitempty"`
	// example:
	//
	// C191B48B-9268-4FB1-A3C2-5143B4A91D0C
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateConnQuestionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateConnQuestionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateConnQuestionResponseBody) GetOutlineId() *int64 {
	return s.OutlineId
}

func (s *CreateConnQuestionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateConnQuestionResponseBody) SetOutlineId(v int64) *CreateConnQuestionResponseBody {
	s.OutlineId = &v
	return s
}

func (s *CreateConnQuestionResponseBody) SetRequestId(v string) *CreateConnQuestionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateConnQuestionResponseBody) Validate() error {
	return dara.Validate(s)
}
