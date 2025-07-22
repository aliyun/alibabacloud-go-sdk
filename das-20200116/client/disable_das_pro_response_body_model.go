// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableDasProResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableDasProResponseBody
	GetCode() *string
	SetData(v string) *DisableDasProResponseBody
	GetData() *string
	SetMessage(v string) *DisableDasProResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableDasProResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DisableDasProResponseBody
	GetSuccess() *string
	SetSynchro(v string) *DisableDasProResponseBody
	GetSynchro() *string
}

type DisableDasProResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// >  If the request was successful, **Successful*	- is returned. If the request failed, an error message such as an error code is returned.
	//
	// example:
	//
	// Successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 7172BECE-588A-5961-8126-C216E16B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
	// The reserved parameter.
	//
	// example:
	//
	// None
	Synchro *string `json:"Synchro,omitempty" xml:"Synchro,omitempty"`
}

func (s DisableDasProResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableDasProResponseBody) GoString() string {
	return s.String()
}

func (s *DisableDasProResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableDasProResponseBody) GetData() *string {
	return s.Data
}

func (s *DisableDasProResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableDasProResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableDasProResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DisableDasProResponseBody) GetSynchro() *string {
	return s.Synchro
}

func (s *DisableDasProResponseBody) SetCode(v string) *DisableDasProResponseBody {
	s.Code = &v
	return s
}

func (s *DisableDasProResponseBody) SetData(v string) *DisableDasProResponseBody {
	s.Data = &v
	return s
}

func (s *DisableDasProResponseBody) SetMessage(v string) *DisableDasProResponseBody {
	s.Message = &v
	return s
}

func (s *DisableDasProResponseBody) SetRequestId(v string) *DisableDasProResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableDasProResponseBody) SetSuccess(v string) *DisableDasProResponseBody {
	s.Success = &v
	return s
}

func (s *DisableDasProResponseBody) SetSynchro(v string) *DisableDasProResponseBody {
	s.Synchro = &v
	return s
}

func (s *DisableDasProResponseBody) Validate() error {
	return dara.Validate(s)
}
