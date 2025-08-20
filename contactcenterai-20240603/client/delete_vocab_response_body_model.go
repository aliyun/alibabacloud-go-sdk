// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *DeleteVocabResponseBody
	GetData() *string
	SetRequestId(v string) *DeleteVocabResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DeleteVocabResponseBody
	GetSuccess() *string
}

type DeleteVocabResponseBody struct {
	// example:
	//
	// true
	Data *string `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// 968A8634-FA2C-5381-9B3E-*******F
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// true
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DeleteVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVocabResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVocabResponseBody) GetData() *string {
	return s.Data
}

func (s *DeleteVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVocabResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DeleteVocabResponseBody) SetData(v string) *DeleteVocabResponseBody {
	s.Data = &v
	return s
}

func (s *DeleteVocabResponseBody) SetRequestId(v string) *DeleteVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVocabResponseBody) SetSuccess(v string) *DeleteVocabResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteVocabResponseBody) Validate() error {
	return dara.Validate(s)
}
