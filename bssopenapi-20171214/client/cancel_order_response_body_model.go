// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCancelOrderResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *CancelOrderResponseBody
	GetCode() *string
	SetData(v *CancelOrderResponseBodyData) *CancelOrderResponseBody
	GetData() *CancelOrderResponseBodyData
	SetMessage(v string) *CancelOrderResponseBody
	GetMessage() *string
	SetRequestId(v string) *CancelOrderResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CancelOrderResponseBody
	GetSuccess() *bool
}

type CancelOrderResponseBody struct {
	// The status code.
	//
	// example:
	//
	// Success
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The data returned.
	Data *CancelOrderResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The error message.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 30D2285E-A848-4ECB-AA74-4954C60858A1
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s CancelOrderResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CancelOrderResponseBody) GetCode() *string {
	return s.Code
}

func (s *CancelOrderResponseBody) GetData() *CancelOrderResponseBodyData {
	return s.Data
}

func (s *CancelOrderResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CancelOrderResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CancelOrderResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CancelOrderResponseBody) SetCode(v string) *CancelOrderResponseBody {
	s.Code = &v
	return s
}

func (s *CancelOrderResponseBody) SetData(v *CancelOrderResponseBodyData) *CancelOrderResponseBody {
	s.Data = v
	return s
}

func (s *CancelOrderResponseBody) SetMessage(v string) *CancelOrderResponseBody {
	s.Message = &v
	return s
}

func (s *CancelOrderResponseBody) SetRequestId(v string) *CancelOrderResponseBody {
	s.RequestId = &v
	return s
}

func (s *CancelOrderResponseBody) SetSuccess(v bool) *CancelOrderResponseBody {
	s.Success = &v
	return s
}

func (s *CancelOrderResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type CancelOrderResponseBodyData struct {
	// The ID of the host.
	//
	// example:
	//
	// cn
	HostId *string `json:"HostId,omitempty" xml:"HostId,omitempty"`
}

func (s CancelOrderResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s CancelOrderResponseBodyData) GoString() string {
	return s.String()
}

func (s *CancelOrderResponseBodyData) GetHostId() *string {
	return s.HostId
}

func (s *CancelOrderResponseBodyData) SetHostId(v string) *CancelOrderResponseBodyData {
	s.HostId = &v
	return s
}

func (s *CancelOrderResponseBodyData) Validate() error {
	return dara.Validate(s)
}
