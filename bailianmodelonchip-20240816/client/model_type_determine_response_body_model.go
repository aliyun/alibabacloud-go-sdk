// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModelTypeDetermineResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ModelTypeDetermineResponseBody
	GetCode() *string
	SetData(v *ModelTypeDetermineResponseBodyData) *ModelTypeDetermineResponseBody
	GetData() *ModelTypeDetermineResponseBodyData
	SetHttpStatusCode(v int32) *ModelTypeDetermineResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *ModelTypeDetermineResponseBody
	GetMessage() *string
	SetRequestId(v string) *ModelTypeDetermineResponseBody
	GetRequestId() *string
	SetSuccess(v string) *ModelTypeDetermineResponseBody
	GetSuccess() *string
}

type ModelTypeDetermineResponseBody struct {
	// example:
	//
	// success
	Code *string                             `json:"code,omitempty" xml:"code,omitempty"`
	Data *ModelTypeDetermineResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 7B0FC4BC-9E4B-5AD7-9D35-6559BDC0788E
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// example:
	//
	// True
	Success *string `json:"success,omitempty" xml:"success,omitempty"`
}

func (s ModelTypeDetermineResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModelTypeDetermineResponseBody) GoString() string {
	return s.String()
}

func (s *ModelTypeDetermineResponseBody) GetCode() *string {
	return s.Code
}

func (s *ModelTypeDetermineResponseBody) GetData() *ModelTypeDetermineResponseBodyData {
	return s.Data
}

func (s *ModelTypeDetermineResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *ModelTypeDetermineResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ModelTypeDetermineResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModelTypeDetermineResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *ModelTypeDetermineResponseBody) SetCode(v string) *ModelTypeDetermineResponseBody {
	s.Code = &v
	return s
}

func (s *ModelTypeDetermineResponseBody) SetData(v *ModelTypeDetermineResponseBodyData) *ModelTypeDetermineResponseBody {
	s.Data = v
	return s
}

func (s *ModelTypeDetermineResponseBody) SetHttpStatusCode(v int32) *ModelTypeDetermineResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *ModelTypeDetermineResponseBody) SetMessage(v string) *ModelTypeDetermineResponseBody {
	s.Message = &v
	return s
}

func (s *ModelTypeDetermineResponseBody) SetRequestId(v string) *ModelTypeDetermineResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModelTypeDetermineResponseBody) SetSuccess(v string) *ModelTypeDetermineResponseBody {
	s.Success = &v
	return s
}

func (s *ModelTypeDetermineResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ModelTypeDetermineResponseBodyData struct {
	// example:
	//
	// false
	FollowUp    *bool   `json:"followUp,omitempty" xml:"followUp,omitempty"`
	RewriteText *string `json:"rewriteText,omitempty" xml:"rewriteText,omitempty"`
	// example:
	//
	// true
	Vl *bool `json:"vl,omitempty" xml:"vl,omitempty"`
}

func (s ModelTypeDetermineResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ModelTypeDetermineResponseBodyData) GoString() string {
	return s.String()
}

func (s *ModelTypeDetermineResponseBodyData) GetFollowUp() *bool {
	return s.FollowUp
}

func (s *ModelTypeDetermineResponseBodyData) GetRewriteText() *string {
	return s.RewriteText
}

func (s *ModelTypeDetermineResponseBodyData) GetVl() *bool {
	return s.Vl
}

func (s *ModelTypeDetermineResponseBodyData) SetFollowUp(v bool) *ModelTypeDetermineResponseBodyData {
	s.FollowUp = &v
	return s
}

func (s *ModelTypeDetermineResponseBodyData) SetRewriteText(v string) *ModelTypeDetermineResponseBodyData {
	s.RewriteText = &v
	return s
}

func (s *ModelTypeDetermineResponseBodyData) SetVl(v bool) *ModelTypeDetermineResponseBodyData {
	s.Vl = &v
	return s
}

func (s *ModelTypeDetermineResponseBodyData) Validate() error {
	return dara.Validate(s)
}
