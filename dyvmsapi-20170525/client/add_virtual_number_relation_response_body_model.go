// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddVirtualNumberRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *AddVirtualNumberRelationResponseBody
	GetCode() *string
	SetData(v string) *AddVirtualNumberRelationResponseBody
	GetData() *string
	SetMessage(v string) *AddVirtualNumberRelationResponseBody
	GetMessage() *string
	SetRequestId(v string) *AddVirtualNumberRelationResponseBody
	GetRequestId() *string
}

type AddVirtualNumberRelationResponseBody struct {
	// The response code.
	//
	// 	- The value 200 indicates that the request was successful.
	//
	// 	- For more information about other response codes, see [API error codes](https://help.aliyun.com/document_detail/112502.html).
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The numbers that failed to be associated.
	//
	// > If all numbers are associated, no value is returned for this parameter.
	//
	// example:
	//
	// 1321111****
	Data *string `json:"Data,omitempty" xml:"Data,omitempty"`
	// The returned message.
	//
	// example:
	//
	// successful
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F655A8D5-B967-440B-8683-DAD6FF8DE990
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddVirtualNumberRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddVirtualNumberRelationResponseBody) GoString() string {
	return s.String()
}

func (s *AddVirtualNumberRelationResponseBody) GetCode() *string {
	return s.Code
}

func (s *AddVirtualNumberRelationResponseBody) GetData() *string {
	return s.Data
}

func (s *AddVirtualNumberRelationResponseBody) GetMessage() *string {
	return s.Message
}

func (s *AddVirtualNumberRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddVirtualNumberRelationResponseBody) SetCode(v string) *AddVirtualNumberRelationResponseBody {
	s.Code = &v
	return s
}

func (s *AddVirtualNumberRelationResponseBody) SetData(v string) *AddVirtualNumberRelationResponseBody {
	s.Data = &v
	return s
}

func (s *AddVirtualNumberRelationResponseBody) SetMessage(v string) *AddVirtualNumberRelationResponseBody {
	s.Message = &v
	return s
}

func (s *AddVirtualNumberRelationResponseBody) SetRequestId(v string) *AddVirtualNumberRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddVirtualNumberRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
