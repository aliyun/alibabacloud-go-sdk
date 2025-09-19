// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetSensitiveDefineRuleConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetSensitiveDefineRuleConfigResponseBody
	GetCode() *string
	SetData(v *SetSensitiveDefineRuleConfigResponseBodyData) *SetSensitiveDefineRuleConfigResponseBody
	GetData() *SetSensitiveDefineRuleConfigResponseBodyData
	SetMessage(v string) *SetSensitiveDefineRuleConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetSensitiveDefineRuleConfigResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetSensitiveDefineRuleConfigResponseBody
	GetSuccess() *bool
}

type SetSensitiveDefineRuleConfigResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data *SetSensitiveDefineRuleConfigResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 0B48AB3C-84FC-424D-A01D-B9270EF4****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetSensitiveDefineRuleConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetSensitiveDefineRuleConfigResponseBody) GoString() string {
	return s.String()
}

func (s *SetSensitiveDefineRuleConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetSensitiveDefineRuleConfigResponseBody) GetData() *SetSensitiveDefineRuleConfigResponseBodyData {
	return s.Data
}

func (s *SetSensitiveDefineRuleConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetSensitiveDefineRuleConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetSensitiveDefineRuleConfigResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetSensitiveDefineRuleConfigResponseBody) SetCode(v string) *SetSensitiveDefineRuleConfigResponseBody {
	s.Code = &v
	return s
}

func (s *SetSensitiveDefineRuleConfigResponseBody) SetData(v *SetSensitiveDefineRuleConfigResponseBodyData) *SetSensitiveDefineRuleConfigResponseBody {
	s.Data = v
	return s
}

func (s *SetSensitiveDefineRuleConfigResponseBody) SetMessage(v string) *SetSensitiveDefineRuleConfigResponseBody {
	s.Message = &v
	return s
}

func (s *SetSensitiveDefineRuleConfigResponseBody) SetRequestId(v string) *SetSensitiveDefineRuleConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetSensitiveDefineRuleConfigResponseBody) SetSuccess(v bool) *SetSensitiveDefineRuleConfigResponseBody {
	s.Success = &v
	return s
}

func (s *SetSensitiveDefineRuleConfigResponseBody) Validate() error {
	return dara.Validate(s)
}

type SetSensitiveDefineRuleConfigResponseBodyData struct {
	// The custom primary key.
	//
	// example:
	//
	// 44616
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s SetSensitiveDefineRuleConfigResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s SetSensitiveDefineRuleConfigResponseBodyData) GoString() string {
	return s.String()
}

func (s *SetSensitiveDefineRuleConfigResponseBodyData) GetId() *int64 {
	return s.Id
}

func (s *SetSensitiveDefineRuleConfigResponseBodyData) SetId(v int64) *SetSensitiveDefineRuleConfigResponseBodyData {
	s.Id = &v
	return s
}

func (s *SetSensitiveDefineRuleConfigResponseBodyData) Validate() error {
	return dara.Validate(s)
}
