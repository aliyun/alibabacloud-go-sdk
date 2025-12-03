// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTitleGenerateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetTitleGenerateResponseBody
	GetCode() *int32
	SetData(v *GetTitleGenerateResponseBodyData) *GetTitleGenerateResponseBody
	GetData() *GetTitleGenerateResponseBodyData
	SetMessage(v string) *GetTitleGenerateResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTitleGenerateResponseBody
	GetRequestId() *string
}

type GetTitleGenerateResponseBody struct {
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetTitleGenerateResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTitleGenerateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTitleGenerateResponseBody) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetTitleGenerateResponseBody) GetData() *GetTitleGenerateResponseBodyData {
	return s.Data
}

func (s *GetTitleGenerateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTitleGenerateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTitleGenerateResponseBody) SetCode(v int32) *GetTitleGenerateResponseBody {
	s.Code = &v
	return s
}

func (s *GetTitleGenerateResponseBody) SetData(v *GetTitleGenerateResponseBodyData) *GetTitleGenerateResponseBody {
	s.Data = v
	return s
}

func (s *GetTitleGenerateResponseBody) SetMessage(v string) *GetTitleGenerateResponseBody {
	s.Message = &v
	return s
}

func (s *GetTitleGenerateResponseBody) SetRequestId(v string) *GetTitleGenerateResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTitleGenerateResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetTitleGenerateResponseBodyData struct {
	// example:
	//
	// 10pcs 80ml Kitchen Disposable Plastic Sauce Cup Pot Chutney Container
	Titles *string `json:"Titles,omitempty" xml:"Titles,omitempty"`
}

func (s GetTitleGenerateResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTitleGenerateResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTitleGenerateResponseBodyData) GetTitles() *string {
	return s.Titles
}

func (s *GetTitleGenerateResponseBodyData) SetTitles(v string) *GetTitleGenerateResponseBodyData {
	s.Titles = &v
	return s
}

func (s *GetTitleGenerateResponseBodyData) Validate() error {
	return dara.Validate(s)
}
