// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetImageDiagnoseResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v int32) *GetImageDiagnoseResponseBody
	GetCode() *int32
	SetData(v *GetImageDiagnoseResponseBodyData) *GetImageDiagnoseResponseBody
	GetData() *GetImageDiagnoseResponseBodyData
	SetMessage(v string) *GetImageDiagnoseResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetImageDiagnoseResponseBody
	GetRequestId() *string
}

type GetImageDiagnoseResponseBody struct {
	// example:
	//
	// 200
	Code *int32                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *GetImageDiagnoseResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// ok
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// DC2DCCC9-C3DF-4F59-8D8E-78185729F16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetImageDiagnoseResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetImageDiagnoseResponseBody) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseResponseBody) GetCode() *int32 {
	return s.Code
}

func (s *GetImageDiagnoseResponseBody) GetData() *GetImageDiagnoseResponseBodyData {
	return s.Data
}

func (s *GetImageDiagnoseResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetImageDiagnoseResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetImageDiagnoseResponseBody) SetCode(v int32) *GetImageDiagnoseResponseBody {
	s.Code = &v
	return s
}

func (s *GetImageDiagnoseResponseBody) SetData(v *GetImageDiagnoseResponseBodyData) *GetImageDiagnoseResponseBody {
	s.Data = v
	return s
}

func (s *GetImageDiagnoseResponseBody) SetMessage(v string) *GetImageDiagnoseResponseBody {
	s.Message = &v
	return s
}

func (s *GetImageDiagnoseResponseBody) SetRequestId(v string) *GetImageDiagnoseResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetImageDiagnoseResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetImageDiagnoseResponseBodyData struct {
	// example:
	//
	// zh
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
}

func (s GetImageDiagnoseResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetImageDiagnoseResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetImageDiagnoseResponseBodyData) GetLanguage() *string {
	return s.Language
}

func (s *GetImageDiagnoseResponseBodyData) SetLanguage(v string) *GetImageDiagnoseResponseBodyData {
	s.Language = &v
	return s
}

func (s *GetImageDiagnoseResponseBodyData) Validate() error {
	return dara.Validate(s)
}
