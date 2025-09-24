// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetResellerUserQuotaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *SetResellerUserQuotaResponseBody
	GetCode() *string
	SetData(v bool) *SetResellerUserQuotaResponseBody
	GetData() *bool
	SetMessage(v string) *SetResellerUserQuotaResponseBody
	GetMessage() *string
	SetRequestId(v string) *SetResellerUserQuotaResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetResellerUserQuotaResponseBody
	GetSuccess() *bool
}

type SetResellerUserQuotaResponseBody struct {
	// The error code returned if the call failed. For more information, see the "Error codes" section of this topic.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// Indicates whether the request is successful.
	//
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// The message returned.
	//
	// example:
	//
	// Successful!
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// F5B803CF-94D8-43AF-ADB3-D819AAD30E27
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call is successful. A value of true indicates that the call is successful. A value of false indicates that the call failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetResellerUserQuotaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetResellerUserQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetResellerUserQuotaResponseBody) GetCode() *string {
	return s.Code
}

func (s *SetResellerUserQuotaResponseBody) GetData() *bool {
	return s.Data
}

func (s *SetResellerUserQuotaResponseBody) GetMessage() *string {
	return s.Message
}

func (s *SetResellerUserQuotaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetResellerUserQuotaResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetResellerUserQuotaResponseBody) SetCode(v string) *SetResellerUserQuotaResponseBody {
	s.Code = &v
	return s
}

func (s *SetResellerUserQuotaResponseBody) SetData(v bool) *SetResellerUserQuotaResponseBody {
	s.Data = &v
	return s
}

func (s *SetResellerUserQuotaResponseBody) SetMessage(v string) *SetResellerUserQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *SetResellerUserQuotaResponseBody) SetRequestId(v string) *SetResellerUserQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetResellerUserQuotaResponseBody) SetSuccess(v bool) *SetResellerUserQuotaResponseBody {
	s.Success = &v
	return s
}

func (s *SetResellerUserQuotaResponseBody) Validate() error {
	return dara.Validate(s)
}
