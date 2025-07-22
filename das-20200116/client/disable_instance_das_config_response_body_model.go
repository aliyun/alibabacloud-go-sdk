// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableInstanceDasConfigResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DisableInstanceDasConfigResponseBody
	GetCode() *string
	SetData(v string) *DisableInstanceDasConfigResponseBody
	GetData() *string
	SetMessage(v string) *DisableInstanceDasConfigResponseBody
	GetMessage() *string
	SetRequestId(v string) *DisableInstanceDasConfigResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DisableInstanceDasConfigResponseBody
	GetSuccess() *string
}

type DisableInstanceDasConfigResponseBody struct {
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The result of disabling the auto scaling feature for the database instance.
	//
	// example:
	//
	// success
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
	// 	- **true**
	//
	// 	- **false**
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DisableInstanceDasConfigResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableInstanceDasConfigResponseBody) GoString() string {
	return s.String()
}

func (s *DisableInstanceDasConfigResponseBody) GetCode() *string {
	return s.Code
}

func (s *DisableInstanceDasConfigResponseBody) GetData() *string {
	return s.Data
}

func (s *DisableInstanceDasConfigResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DisableInstanceDasConfigResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableInstanceDasConfigResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DisableInstanceDasConfigResponseBody) SetCode(v string) *DisableInstanceDasConfigResponseBody {
	s.Code = &v
	return s
}

func (s *DisableInstanceDasConfigResponseBody) SetData(v string) *DisableInstanceDasConfigResponseBody {
	s.Data = &v
	return s
}

func (s *DisableInstanceDasConfigResponseBody) SetMessage(v string) *DisableInstanceDasConfigResponseBody {
	s.Message = &v
	return s
}

func (s *DisableInstanceDasConfigResponseBody) SetRequestId(v string) *DisableInstanceDasConfigResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableInstanceDasConfigResponseBody) SetSuccess(v string) *DisableInstanceDasConfigResponseBody {
	s.Success = &v
	return s
}

func (s *DisableInstanceDasConfigResponseBody) Validate() error {
	return dara.Validate(s)
}
