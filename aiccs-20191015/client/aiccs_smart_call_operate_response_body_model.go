// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiccsSmartCallOperateResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AiccsSmartCallOperateResponseBody
	GetCode() *string
	SetData(v string) *AiccsSmartCallOperateResponseBody
	GetData() *string
	SetMessage(v string) *AiccsSmartCallOperateResponseBody
	GetMessage() *string
	SetRequestId(v string) *AiccsSmartCallOperateResponseBody
	GetRequestId() *string
}

type AiccsSmartCallOperateResponseBody struct {
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// true
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AiccsSmartCallOperateResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AiccsSmartCallOperateResponseBody) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallOperateResponseBody) GetCode() *string {
	return s.Code
}

func (s *AiccsSmartCallOperateResponseBody) GetData() *string {
	return s.Data
}

func (s *AiccsSmartCallOperateResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AiccsSmartCallOperateResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AiccsSmartCallOperateResponseBody) SetCode(v string) *AiccsSmartCallOperateResponseBody {
	s.Code = &v
	return s
}

func (s *AiccsSmartCallOperateResponseBody) SetData(v string) *AiccsSmartCallOperateResponseBody {
	s.Data = &v
	return s
}

func (s *AiccsSmartCallOperateResponseBody) SetMessage(v string) *AiccsSmartCallOperateResponseBody {
	s.Message = &v
	return s
}

func (s *AiccsSmartCallOperateResponseBody) SetRequestId(v string) *AiccsSmartCallOperateResponseBody {
	s.RequestId = &v
	return s
}

func (s *AiccsSmartCallOperateResponseBody) Validate() error {
	return dara.Validate(s)
}
