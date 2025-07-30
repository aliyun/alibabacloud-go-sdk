// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDTSIPResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDynamicCode(v string) *DescribeDTSIPResponseBody
	GetDynamicCode() *string
	SetDynamicMessage(v string) *DescribeDTSIPResponseBody
	GetDynamicMessage() *string
	SetErrCode(v string) *DescribeDTSIPResponseBody
	GetErrCode() *string
	SetErrMessage(v string) *DescribeDTSIPResponseBody
	GetErrMessage() *string
	SetRequestId(v string) *DescribeDTSIPResponseBody
	GetRequestId() *string
	SetSuccess(v string) *DescribeDTSIPResponseBody
	GetSuccess() *string
}

type DescribeDTSIPResponseBody struct {
	// The internal error code. This parameter will be removed in the future.
	//
	// example:
	//
	// 500
	DynamicCode *string `json:"DynamicCode,omitempty" xml:"DynamicCode,omitempty"`
	// The CIDR blocks of DTS servers.
	//
	// example:
	//
	// 10.151.12.0/24,47.102.181.0/24,47.101.109.0/24,120.55.129.0/24,11.115.103.0/24,47.102.234.0/24
	DynamicMessage *string `json:"DynamicMessage,omitempty" xml:"DynamicMessage,omitempty"`
	// The error code returned if the call failed.
	//
	// example:
	//
	// InternalError
	ErrCode *string `json:"ErrCode,omitempty" xml:"ErrCode,omitempty"`
	// The error message returned if the call failed.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	ErrMessage *string `json:"ErrMessage,omitempty" xml:"ErrMessage,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 0C99C0BE-F312-40FA-ADFA-4DC1166B****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the call was successful.
	//
	// example:
	//
	// true
	Success *string `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DescribeDTSIPResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDTSIPResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDTSIPResponseBody) GetDynamicCode() *string {
	return s.DynamicCode
}

func (s *DescribeDTSIPResponseBody) GetDynamicMessage() *string {
	return s.DynamicMessage
}

func (s *DescribeDTSIPResponseBody) GetErrCode() *string {
	return s.ErrCode
}

func (s *DescribeDTSIPResponseBody) GetErrMessage() *string {
	return s.ErrMessage
}

func (s *DescribeDTSIPResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDTSIPResponseBody) GetSuccess() *string {
	return s.Success
}

func (s *DescribeDTSIPResponseBody) SetDynamicCode(v string) *DescribeDTSIPResponseBody {
	s.DynamicCode = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetDynamicMessage(v string) *DescribeDTSIPResponseBody {
	s.DynamicMessage = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetErrCode(v string) *DescribeDTSIPResponseBody {
	s.ErrCode = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetErrMessage(v string) *DescribeDTSIPResponseBody {
	s.ErrMessage = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetRequestId(v string) *DescribeDTSIPResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDTSIPResponseBody) SetSuccess(v string) *DescribeDTSIPResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeDTSIPResponseBody) Validate() error {
	return dara.Validate(s)
}
