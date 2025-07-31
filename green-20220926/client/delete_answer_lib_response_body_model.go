// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAnswerLibResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteAnswerLibResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteAnswerLibResponseBody
	GetRequestId() *string
}

type DeleteAnswerLibResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAnswerLibResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAnswerLibResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAnswerLibResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAnswerLibResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAnswerLibResponseBody) SetData(v bool) *DeleteAnswerLibResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAnswerLibResponseBody) SetRequestId(v string) *DeleteAnswerLibResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAnswerLibResponseBody) Validate() error {
	return dara.Validate(s)
}
