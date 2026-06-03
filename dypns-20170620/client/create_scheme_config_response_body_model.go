// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSchemeConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CreateSchemeConfigResponseBody
	GetCode() *string
	SetMessage(v string) *CreateSchemeConfigResponseBody
	GetMessage() *string
	SetModel(v *CreateSchemeConfigResponseBodyModel) *CreateSchemeConfigResponseBody
	GetModel() *CreateSchemeConfigResponseBodyModel
	SetRequestId(v string) *CreateSchemeConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateSchemeConfigResponseBody
	GetSuccess() *bool
}

type CreateSchemeConfigResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// 示例值示例值
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// "Model": {
	//
	//         "SchemeCode": "FA000000003961804087"
	//
	//     }
	Model *CreateSchemeConfigResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// example:
	//
	// 2E7CA885-8D97-C497-A02C-2D31214D3285
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CreateSchemeConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemeConfigResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSchemeConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *CreateSchemeConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateSchemeConfigResponseBody) GetModel() *CreateSchemeConfigResponseBodyModel {
	return s.Model
}

func (s *CreateSchemeConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateSchemeConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateSchemeConfigResponseBody) SetCode(v string) *CreateSchemeConfigResponseBody {
	s.Code = &v
	return s
}

func (s *CreateSchemeConfigResponseBody) SetMessage(v string) *CreateSchemeConfigResponseBody {
	s.Message = &v
	return s
}

func (s *CreateSchemeConfigResponseBody) SetModel(v *CreateSchemeConfigResponseBodyModel) *CreateSchemeConfigResponseBody {
	s.Model = v
	return s
}

func (s *CreateSchemeConfigResponseBody) SetRequestId(v string) *CreateSchemeConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateSchemeConfigResponseBody) SetSuccess(v bool) *CreateSchemeConfigResponseBody {
	s.Success = &v
	return s
}

func (s *CreateSchemeConfigResponseBody) Validate() error {
	if s.Model != nil {
		if err := s.Model.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CreateSchemeConfigResponseBodyModel struct {
	// example:
	//
	// FA000000003540014055
	SchemeCode *string `json:"SchemeCode,omitempty" xml:"SchemeCode,omitempty"`
}

func (s CreateSchemeConfigResponseBodyModel) String() string {
	return dara.Prettify(s)
}

func (s CreateSchemeConfigResponseBodyModel) GoString() string {
	return s.String()
}

func (s *CreateSchemeConfigResponseBodyModel) GetSchemeCode() *string {
	return s.SchemeCode
}

func (s *CreateSchemeConfigResponseBodyModel) SetSchemeCode(v string) *CreateSchemeConfigResponseBodyModel {
	s.SchemeCode = &v
	return s
}

func (s *CreateSchemeConfigResponseBodyModel) Validate() error {
	return dara.Validate(s)
}
