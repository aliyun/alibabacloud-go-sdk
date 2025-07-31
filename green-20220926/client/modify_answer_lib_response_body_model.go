// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAnswerLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *ModifyAnswerLibResponseBody
	GetData() *bool
	SetRequestId(v string) *ModifyAnswerLibResponseBody
	GetRequestId() *string
}

type ModifyAnswerLibResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAnswerLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAnswerLibResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAnswerLibResponseBody) GetData() *bool {
	return s.Data
}

func (s *ModifyAnswerLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAnswerLibResponseBody) SetData(v bool) *ModifyAnswerLibResponseBody {
	s.Data = &v
	return s
}

func (s *ModifyAnswerLibResponseBody) SetRequestId(v string) *ModifyAnswerLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAnswerLibResponseBody) Validate() error {
	return dara.Validate(s)
}
