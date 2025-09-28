// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iVerifyMobileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *VerifyMobileResponseBody
	GetCode() *string
	SetGateVerifyResultDTO(v *VerifyMobileResponseBodyGateVerifyResultDTO) *VerifyMobileResponseBody
	GetGateVerifyResultDTO() *VerifyMobileResponseBodyGateVerifyResultDTO
	SetMessage(v string) *VerifyMobileResponseBody
	GetMessage() *string
	SetRequestId(v string) *VerifyMobileResponseBody
	GetRequestId() *string
}

type VerifyMobileResponseBody struct {
	// The response code.
	//
	// 	- If OK is returned, the request is successful.
	//
	// 	- For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
	//
	// example:
	//
	// OK
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The response parameters.
	GateVerifyResultDTO *VerifyMobileResponseBodyGateVerifyResultDTO `json:"GateVerifyResultDTO,omitempty" xml:"GateVerifyResultDTO,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 8906582E-6722
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s VerifyMobileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s VerifyMobileResponseBody) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponseBody) GetCode() *string {
	return s.Code
}

func (s *VerifyMobileResponseBody) GetGateVerifyResultDTO() *VerifyMobileResponseBodyGateVerifyResultDTO {
	return s.GateVerifyResultDTO
}

func (s *VerifyMobileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *VerifyMobileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *VerifyMobileResponseBody) SetCode(v string) *VerifyMobileResponseBody {
	s.Code = &v
	return s
}

func (s *VerifyMobileResponseBody) SetGateVerifyResultDTO(v *VerifyMobileResponseBodyGateVerifyResultDTO) *VerifyMobileResponseBody {
	s.GateVerifyResultDTO = v
	return s
}

func (s *VerifyMobileResponseBody) SetMessage(v string) *VerifyMobileResponseBody {
	s.Message = &v
	return s
}

func (s *VerifyMobileResponseBody) SetRequestId(v string) *VerifyMobileResponseBody {
	s.RequestId = &v
	return s
}

func (s *VerifyMobileResponseBody) Validate() error {
	return dara.Validate(s)
}

type VerifyMobileResponseBodyGateVerifyResultDTO struct {
	// The verification ID.
	//
	// example:
	//
	// 121343241
	VerifyId *string `json:"VerifyId,omitempty" xml:"VerifyId,omitempty"`
	// The verification results. Valid values:
	//
	// 	- **PASS: The input phone number is consistent with the phone number that you use.**
	//
	// 	- **REJECT: The input phone number is different from the phone number that you use.**
	//
	// 	- **UNKNOWN: The system cannot judge whether the input phone number is consistent with the phone number that you use.
	//
	// example:
	//
	// PASS
	VerifyResult *string `json:"VerifyResult,omitempty" xml:"VerifyResult,omitempty"`
}

func (s VerifyMobileResponseBodyGateVerifyResultDTO) String() string {
	return dara.Prettify(s)
}

func (s VerifyMobileResponseBodyGateVerifyResultDTO) GoString() string {
	return s.String()
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) GetVerifyId() *string {
	return s.VerifyId
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) GetVerifyResult() *string {
	return s.VerifyResult
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) SetVerifyId(v string) *VerifyMobileResponseBodyGateVerifyResultDTO {
	s.VerifyId = &v
	return s
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) SetVerifyResult(v string) *VerifyMobileResponseBodyGateVerifyResultDTO {
	s.VerifyResult = &v
	return s
}

func (s *VerifyMobileResponseBodyGateVerifyResultDTO) Validate() error {
	return dara.Validate(s)
}
