// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAsrVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DeleteAsrVocabResponseBody
	GetCode() *string
	SetData(v string) *DeleteAsrVocabResponseBody
	GetData() *string
	SetMessage(v string) *DeleteAsrVocabResponseBody
	GetMessage() *string
	SetRequestId(v string) *DeleteAsrVocabResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteAsrVocabResponseBody
	GetSuccess() *bool
}

type DeleteAsrVocabResponseBody struct {
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 71b1795ac8634bd8bdf4d3878480c7c2
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 4987D326-83D9-4A42-B9A5-0B27F9B40539
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteAsrVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAsrVocabResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAsrVocabResponseBody) GetCode() *string {
	return s.Code
}

func (s *DeleteAsrVocabResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteAsrVocabResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DeleteAsrVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAsrVocabResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteAsrVocabResponseBody) SetCode(v string) *DeleteAsrVocabResponseBody {
	s.Code = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetData(v string) *DeleteAsrVocabResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetMessage(v string) *DeleteAsrVocabResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetRequestId(v string) *DeleteAsrVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) SetSuccess(v bool) *DeleteAsrVocabResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteAsrVocabResponseBody) Validate() error {
	return dara.Validate(s)
}
