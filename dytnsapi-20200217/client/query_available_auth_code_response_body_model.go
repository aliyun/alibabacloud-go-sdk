// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryAvailableAuthCodeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryAvailableAuthCodeResponseBody
	GetCode() *string
	SetData(v []*string) *QueryAvailableAuthCodeResponseBody
	GetData() []*string
	SetMessage(v string) *QueryAvailableAuthCodeResponseBody
	GetMessage() *string
	SetRequestId(v string) *QueryAvailableAuthCodeResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *QueryAvailableAuthCodeResponseBody
	GetSuccess() *bool
}

type QueryAvailableAuthCodeResponseBody struct {
	// The response code. **OK*	- indicates that the request is successful.
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned data.
	Data []*string `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6ADDCD31-6BC7-5913-A47F-E29A07E37FEE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful. Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// false
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s QueryAvailableAuthCodeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryAvailableAuthCodeResponseBody) GoString() string {
	return s.String()
}

func (s *QueryAvailableAuthCodeResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryAvailableAuthCodeResponseBody) GetData() []*string {
	return s.Data
}

func (s *QueryAvailableAuthCodeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryAvailableAuthCodeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryAvailableAuthCodeResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *QueryAvailableAuthCodeResponseBody) SetCode(v string) *QueryAvailableAuthCodeResponseBody {
	s.Code = &v
	return s
}

func (s *QueryAvailableAuthCodeResponseBody) SetData(v []*string) *QueryAvailableAuthCodeResponseBody {
	s.Data = v
	return s
}

func (s *QueryAvailableAuthCodeResponseBody) SetMessage(v string) *QueryAvailableAuthCodeResponseBody {
	s.Message = &v
	return s
}

func (s *QueryAvailableAuthCodeResponseBody) SetRequestId(v string) *QueryAvailableAuthCodeResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryAvailableAuthCodeResponseBody) SetSuccess(v bool) *QueryAvailableAuthCodeResponseBody {
	s.Success = &v
	return s
}

func (s *QueryAvailableAuthCodeResponseBody) Validate() error {
	return dara.Validate(s)
}
