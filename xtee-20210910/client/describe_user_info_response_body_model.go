// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserInfoResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeUserInfoResponseBody
	GetCode() *string
	SetHttpStatusCode(v string) *DescribeUserInfoResponseBody
	GetHttpStatusCode() *string
	SetMessage(v string) *DescribeUserInfoResponseBody
	GetMessage() *string
	SetResultObject(v *DescribeUserInfoResponseBodyResultObject) *DescribeUserInfoResponseBody
	GetResultObject() *DescribeUserInfoResponseBodyResultObject
	SetSuccess(v bool) *DescribeUserInfoResponseBody
	GetSuccess() *bool
}

type DescribeUserInfoResponseBody struct {
	// Status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// HTTP status code
	//
	// example:
	//
	// 200
	HttpStatusCode *string `json:"httpStatusCode,omitempty" xml:"httpStatusCode,omitempty"`
	// Error message
	//
	// example:
	//
	// The input parameter data is not valid. order_storage_company_num component not found
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Returned object
	ResultObject *DescribeUserInfoResponseBodyResultObject `json:"resultObject,omitempty" xml:"resultObject,omitempty" type:"Struct"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"success,omitempty" xml:"success,omitempty"`
}

func (s DescribeUserInfoResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserInfoResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeUserInfoResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeUserInfoResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeUserInfoResponseBody) GetResultObject() *DescribeUserInfoResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeUserInfoResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeUserInfoResponseBody) SetCode(v string) *DescribeUserInfoResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeUserInfoResponseBody) SetHttpStatusCode(v string) *DescribeUserInfoResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeUserInfoResponseBody) SetMessage(v string) *DescribeUserInfoResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeUserInfoResponseBody) SetResultObject(v *DescribeUserInfoResponseBodyResultObject) *DescribeUserInfoResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeUserInfoResponseBody) SetSuccess(v bool) *DescribeUserInfoResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeUserInfoResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUserInfoResponseBodyResultObject struct {
	// Client IP.
	//
	// example:
	//
	// 192.168.0.1
	ClientIp *string `json:"clientIp,omitempty" xml:"clientIp,omitempty"`
	// Sub-account ID
	//
	// example:
	//
	// 555666
	SubId *string `json:"subId,omitempty" xml:"subId,omitempty"`
	// User UID
	//
	// example:
	//
	// 15633333331111
	UserId *int64 `json:"userId,omitempty" xml:"userId,omitempty"`
	// User name
	//
	// example:
	//
	// root
	UserName *string `json:"userName,omitempty" xml:"userName,omitempty"`
}

func (s DescribeUserInfoResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserInfoResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeUserInfoResponseBodyResultObject) GetClientIp() *string {
	return s.ClientIp
}

func (s *DescribeUserInfoResponseBodyResultObject) GetSubId() *string {
	return s.SubId
}

func (s *DescribeUserInfoResponseBodyResultObject) GetUserId() *int64 {
	return s.UserId
}

func (s *DescribeUserInfoResponseBodyResultObject) GetUserName() *string {
	return s.UserName
}

func (s *DescribeUserInfoResponseBodyResultObject) SetClientIp(v string) *DescribeUserInfoResponseBodyResultObject {
	s.ClientIp = &v
	return s
}

func (s *DescribeUserInfoResponseBodyResultObject) SetSubId(v string) *DescribeUserInfoResponseBodyResultObject {
	s.SubId = &v
	return s
}

func (s *DescribeUserInfoResponseBodyResultObject) SetUserId(v int64) *DescribeUserInfoResponseBodyResultObject {
	s.UserId = &v
	return s
}

func (s *DescribeUserInfoResponseBodyResultObject) SetUserName(v string) *DescribeUserInfoResponseBodyResultObject {
	s.UserName = &v
	return s
}

func (s *DescribeUserInfoResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
