// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryPhoneNoAByTrackNoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *QueryPhoneNoAByTrackNoResponseBody
	GetCode() *string
	SetMessage(v string) *QueryPhoneNoAByTrackNoResponseBody
	GetMessage() *string
	SetModule(v []*QueryPhoneNoAByTrackNoResponseBodyModule) *QueryPhoneNoAByTrackNoResponseBody
	GetModule() []*QueryPhoneNoAByTrackNoResponseBodyModule
	SetRequestId(v string) *QueryPhoneNoAByTrackNoResponseBody
	GetRequestId() *string
}

type QueryPhoneNoAByTrackNoResponseBody struct {
	// The response code.
	//
	// 	- The value OK indicates that the request was successful.
	//
	// 	- Other status codes indicate that the request failed. For more information, see [Error codes](https://help.aliyun.com/document_detail/109196.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The returned message.
	//
	// example:
	//
	// OK
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The information returned after the phone numbers were bound.
	Module []*QueryPhoneNoAByTrackNoResponseBodyModule `json:"Module,omitempty" xml:"Module,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s QueryPhoneNoAByTrackNoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoResponseBody) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoResponseBody) GetCode() *string {
	return s.Code
}

func (s *QueryPhoneNoAByTrackNoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryPhoneNoAByTrackNoResponseBody) GetModule() []*QueryPhoneNoAByTrackNoResponseBodyModule {
	return s.Module
}

func (s *QueryPhoneNoAByTrackNoResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetCode(v string) *QueryPhoneNoAByTrackNoResponseBody {
	s.Code = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetMessage(v string) *QueryPhoneNoAByTrackNoResponseBody {
	s.Message = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetModule(v []*QueryPhoneNoAByTrackNoResponseBodyModule) *QueryPhoneNoAByTrackNoResponseBody {
	s.Module = v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) SetRequestId(v string) *QueryPhoneNoAByTrackNoResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBody) Validate() error {
	if s.Module != nil {
		for _, item := range s.Module {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryPhoneNoAByTrackNoResponseBodyModule struct {
	// The extension of phone number X.
	//
	// example:
	//
	// 130
	Extension *string `json:"Extension,omitempty" xml:"Extension,omitempty"`
	// Phone number A.
	//
	// example:
	//
	// 1310000****
	PhoneNoA *string `json:"PhoneNoA,omitempty" xml:"PhoneNoA,omitempty"`
	// The private number, that is, phone number X.
	//
	// example:
	//
	// 1710000****
	PhoneNoX *string `json:"PhoneNoX,omitempty" xml:"PhoneNoX,omitempty"`
}

func (s QueryPhoneNoAByTrackNoResponseBodyModule) String() string {
	return dara.Prettify(s)
}

func (s QueryPhoneNoAByTrackNoResponseBodyModule) GoString() string {
	return s.String()
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) GetExtension() *string {
	return s.Extension
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) GetPhoneNoA() *string {
	return s.PhoneNoA
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) GetPhoneNoX() *string {
	return s.PhoneNoX
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) SetExtension(v string) *QueryPhoneNoAByTrackNoResponseBodyModule {
	s.Extension = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) SetPhoneNoA(v string) *QueryPhoneNoAByTrackNoResponseBodyModule {
	s.PhoneNoA = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) SetPhoneNoX(v string) *QueryPhoneNoAByTrackNoResponseBodyModule {
	s.PhoneNoX = &v
	return s
}

func (s *QueryPhoneNoAByTrackNoResponseBodyModule) Validate() error {
	return dara.Validate(s)
}
