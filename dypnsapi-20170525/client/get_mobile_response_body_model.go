// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMobileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetMobileResponseBody
	GetCode() *string
	SetGetMobileResultDTO(v *GetMobileResponseBodyGetMobileResultDTO) *GetMobileResponseBody
	GetGetMobileResultDTO() *GetMobileResponseBodyGetMobileResultDTO
	SetMessage(v string) *GetMobileResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetMobileResponseBody
	GetRequestId() *string
}

type GetMobileResponseBody struct {
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
	GetMobileResultDTO *GetMobileResponseBodyGetMobileResultDTO `json:"GetMobileResultDTO,omitempty" xml:"GetMobileResultDTO,omitempty" type:"Struct"`
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

func (s GetMobileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMobileResponseBody) GoString() string {
	return s.String()
}

func (s *GetMobileResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetMobileResponseBody) GetGetMobileResultDTO() *GetMobileResponseBodyGetMobileResultDTO {
	return s.GetMobileResultDTO
}

func (s *GetMobileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetMobileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMobileResponseBody) SetCode(v string) *GetMobileResponseBody {
	s.Code = &v
	return s
}

func (s *GetMobileResponseBody) SetGetMobileResultDTO(v *GetMobileResponseBodyGetMobileResultDTO) *GetMobileResponseBody {
	s.GetMobileResultDTO = v
	return s
}

func (s *GetMobileResponseBody) SetMessage(v string) *GetMobileResponseBody {
	s.Message = &v
	return s
}

func (s *GetMobileResponseBody) SetRequestId(v string) *GetMobileResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMobileResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMobileResponseBodyGetMobileResultDTO struct {
	// The phone number,
	//
	// example:
	//
	// 13900001234
	Mobile *string `json:"Mobile,omitempty" xml:"Mobile,omitempty"`
}

func (s GetMobileResponseBodyGetMobileResultDTO) String() string {
	return dara.Prettify(s)
}

func (s GetMobileResponseBodyGetMobileResultDTO) GoString() string {
	return s.String()
}

func (s *GetMobileResponseBodyGetMobileResultDTO) GetMobile() *string {
	return s.Mobile
}

func (s *GetMobileResponseBodyGetMobileResultDTO) SetMobile(v string) *GetMobileResponseBodyGetMobileResultDTO {
	s.Mobile = &v
	return s
}

func (s *GetMobileResponseBodyGetMobileResultDTO) Validate() error {
	return dara.Validate(s)
}
