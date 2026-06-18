// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAiccsSmartCallResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AiccsSmartCallResponseBody
	GetCode() *string
	SetData(v string) *AiccsSmartCallResponseBody
	GetData() *string
	SetMessage(v string) *AiccsSmartCallResponseBody
	GetMessage() *string
	SetRequestId(v string) *AiccsSmartCallResponseBody
	GetRequestId() *string
}

type AiccsSmartCallResponseBody struct {
	// Unique receipt ID for this call.
	//
	// example:
	//
	// 116012854210^10281427****
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Request status code. A return value of "OK" indicates that the request succeeded.
	//
	// example:
	//
	// OK
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// Description of the status code.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// Request ID.
	//
	// example:
	//
	// A90E4451-FED7-49D2-87C8-00700A8C4D0D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AiccsSmartCallResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AiccsSmartCallResponseBody) GoString() string {
	return s.String()
}

func (s *AiccsSmartCallResponseBody) GetCode() *string {
	return s.Code
}

func (s *AiccsSmartCallResponseBody) GetData() *string {
	return s.Data
}

func (s *AiccsSmartCallResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AiccsSmartCallResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AiccsSmartCallResponseBody) SetCode(v string) *AiccsSmartCallResponseBody {
	s.Code = &v
	return s
}

func (s *AiccsSmartCallResponseBody) SetData(v string) *AiccsSmartCallResponseBody {
	s.Data = &v
	return s
}

func (s *AiccsSmartCallResponseBody) SetMessage(v string) *AiccsSmartCallResponseBody {
	s.Message = &v
	return s
}

func (s *AiccsSmartCallResponseBody) SetRequestId(v string) *AiccsSmartCallResponseBody {
	s.RequestId = &v
	return s
}

func (s *AiccsSmartCallResponseBody) Validate() error {
	return dara.Validate(s)
}
