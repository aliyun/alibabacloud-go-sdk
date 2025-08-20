// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateVocabResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v string) *UpdateVocabResponseBody
	GetData() *string
	SetRequestId(v string) *UpdateVocabResponseBody
	GetRequestId() *string
	SetSuccess(v string) *UpdateVocabResponseBody
	GetSuccess() *string
}

type UpdateVocabResponseBody struct {
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

func (s UpdateVocabResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateVocabResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateVocabResponseBody) GetData() *string {
	return s.Data
}

func (s *UpdateVocabResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateVocabResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *UpdateVocabResponseBody) SetData(v string) *UpdateVocabResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateVocabResponseBody) SetRequestId(v string) *UpdateVocabResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateVocabResponseBody) SetSuccess(v string) *UpdateVocabResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateVocabResponseBody) Validate() error {
	return dara.Validate(s)
}
