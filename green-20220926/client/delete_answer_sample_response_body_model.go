// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAnswerSampleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DeleteAnswerSampleResponseBody
	GetData() *bool
	SetRequestId(v string) *DeleteAnswerSampleResponseBody
	GetRequestId() *string
}

type DeleteAnswerSampleResponseBody struct {
	// example:
	//
	// True
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// AAAAAA-BBBB-CCCCC-DDDD-EEEEEEEE****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAnswerSampleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAnswerSampleResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAnswerSampleResponseBody) GetData() *bool {
	return s.Data
}

func (s *DeleteAnswerSampleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAnswerSampleResponseBody) SetData(v bool) *DeleteAnswerSampleResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAnswerSampleResponseBody) SetRequestId(v string) *DeleteAnswerSampleResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAnswerSampleResponseBody) Validate() error {
	return dara.Validate(s)
}
