// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVerifySchemeResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeVerifySchemeResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeVerifySchemeResponseBody
	GetMessage() *string
	SetRequestId(v string) *DescribeVerifySchemeResponseBody
	GetRequestId() *string
	SetSchemeQueryResultDTO(v *DescribeVerifySchemeResponseBodySchemeQueryResultDTO) *DescribeVerifySchemeResponseBody
	GetSchemeQueryResultDTO() *DescribeVerifySchemeResponseBodySchemeQueryResultDTO
}

type DescribeVerifySchemeResponseBody struct {
	// The response code. OK indicates that the request is successful. For more information about other error codes, see [API response codes](https://help.aliyun.com/document_detail/85198.html).
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
	// The request ID.
	//
	// example:
	//
	// 0C5380A7-2032-5F7D-9614-1BF8B54D16CB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The response parameters.
	SchemeQueryResultDTO *DescribeVerifySchemeResponseBodySchemeQueryResultDTO `json:"SchemeQueryResultDTO,omitempty" xml:"SchemeQueryResultDTO,omitempty" type:"Struct"`
}

func (s DescribeVerifySchemeResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySchemeResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeVerifySchemeResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeVerifySchemeResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVerifySchemeResponseBody) GetSchemeQueryResultDTO() *DescribeVerifySchemeResponseBodySchemeQueryResultDTO {
	return s.SchemeQueryResultDTO
}

func (s *DescribeVerifySchemeResponseBody) SetCode(v string) *DescribeVerifySchemeResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetMessage(v string) *DescribeVerifySchemeResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetRequestId(v string) *DescribeVerifySchemeResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVerifySchemeResponseBody) SetSchemeQueryResultDTO(v *DescribeVerifySchemeResponseBodySchemeQueryResultDTO) *DescribeVerifySchemeResponseBody {
	s.SchemeQueryResultDTO = v
	return s
}

func (s *DescribeVerifySchemeResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVerifySchemeResponseBodySchemeQueryResultDTO struct {
	// The key generated when you create a service in the console.
	//
	// example:
	//
	// ZDMARqPkyQzWVJjB/sB/+fCp5TA4lNsRnY7rEC+HfGsOIOk1Brj8UyXFW2RBYIWqLieCSo8ZypEaEj+h9rLd3FgpXAjGYDfmOperod6jPUUwFHhBObxK+HuKVoi2jOqN7aDOlyPyGcATyq3BDdlf922JmnFLT8Hvnu4qgzzCZk0LXWTb0XVPnm5/fHUGHEA2Q+aTrGkaWcHjmTDqQ7BtvrAIIcJJkCJu4i1aeU++/0EzGWap4mcb2VhKROBs****
	AppEncryptInfo *string `json:"AppEncryptInfo,omitempty" xml:"AppEncryptInfo,omitempty"`
}

func (s DescribeVerifySchemeResponseBodySchemeQueryResultDTO) String() string {
	return dara.Prettify(s)
}

func (s DescribeVerifySchemeResponseBodySchemeQueryResultDTO) GoString() string {
	return s.String()
}

func (s *DescribeVerifySchemeResponseBodySchemeQueryResultDTO) GetAppEncryptInfo() *string {
	return s.AppEncryptInfo
}

func (s *DescribeVerifySchemeResponseBodySchemeQueryResultDTO) SetAppEncryptInfo(v string) *DescribeVerifySchemeResponseBodySchemeQueryResultDTO {
	s.AppEncryptInfo = &v
	return s
}

func (s *DescribeVerifySchemeResponseBodySchemeQueryResultDTO) Validate() error {
	return dara.Validate(s)
}
