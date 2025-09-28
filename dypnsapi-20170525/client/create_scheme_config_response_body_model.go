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
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/zh/pnvs/developer-reference/api-return-code?spm=a2c4g.11186623.0.0.5c3a662fbgeAuk).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The returned results.
	Model *CreateSchemeConfigResponseBodyModel `json:"Model,omitempty" xml:"Model,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// B95B36EC-8108-4479-D3AA-2BB27F9B155A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
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
	return dara.Validate(s)
}

type CreateSchemeConfigResponseBodyModel struct {
	// The service code.
	//
	// example:
	//
	// FA100000168468035
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
