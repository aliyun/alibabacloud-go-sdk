// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeModelOssTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DescribeModelOssTokenResponseBody
	GetAccessId() *string
	SetCode(v int64) *DescribeModelOssTokenResponseBody
	GetCode() *int64
	SetHost(v string) *DescribeModelOssTokenResponseBody
	GetHost() *string
	SetHttpStatusCode(v int64) *DescribeModelOssTokenResponseBody
	GetHttpStatusCode() *int64
	SetKey(v string) *DescribeModelOssTokenResponseBody
	GetKey() *string
	SetPolicy(v string) *DescribeModelOssTokenResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribeModelOssTokenResponseBody
	GetRequestId() *string
	SetResultObject(v *DescribeModelOssTokenResponseBodyResultObject) *DescribeModelOssTokenResponseBody
	GetResultObject() *DescribeModelOssTokenResponseBodyResultObject
	SetSignature(v string) *DescribeModelOssTokenResponseBody
	GetSignature() *string
	SetSuccess(v bool) *DescribeModelOssTokenResponseBody
	GetSuccess() *bool
	SetXOssSecurityToken(v string) *DescribeModelOssTokenResponseBody
	GetXOssSecurityToken() *string
}

type DescribeModelOssTokenResponseBody struct {
	// example:
	//
	// Lxxxxxxxxxxxxxxxxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// 200
	Code *int64 `json:"Code,omitempty" xml:"Code,omitempty"`
	// example:
	//
	// https://safxxxxxxxxx.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *int64 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// example:
	//
	// saf/xxxxxx/xxxxxx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMxxxxxxxxxxxxxxxxxxxxxxxxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// 4A91D2D1-AEC9-1658-ABCE-5A39DE66A5C2
	RequestId    *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	ResultObject *DescribeModelOssTokenResponseBodyResultObject `json:"ResultObject,omitempty" xml:"ResultObject,omitempty" type:"Struct"`
	// example:
	//
	// lUxxxxxxxxxxxxxxxxxxxx
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
	// example:
	//
	// smxxxxxxxxxxx
	XOssSecurityToken *string `json:"XOssSecurityToken,omitempty" xml:"XOssSecurityToken,omitempty"`
}

func (s DescribeModelOssTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOssTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeModelOssTokenResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribeModelOssTokenResponseBody) GetCode() *int64 {
	return s.Code
}

func (s *DescribeModelOssTokenResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribeModelOssTokenResponseBody) GetHttpStatusCode() *int64 {
	return s.HttpStatusCode
}

func (s *DescribeModelOssTokenResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribeModelOssTokenResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeModelOssTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeModelOssTokenResponseBody) GetResultObject() *DescribeModelOssTokenResponseBodyResultObject {
	return s.ResultObject
}

func (s *DescribeModelOssTokenResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribeModelOssTokenResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeModelOssTokenResponseBody) GetXOssSecurityToken() *string {
	return s.XOssSecurityToken
}

func (s *DescribeModelOssTokenResponseBody) SetAccessId(v string) *DescribeModelOssTokenResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetCode(v int64) *DescribeModelOssTokenResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetHost(v string) *DescribeModelOssTokenResponseBody {
	s.Host = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetHttpStatusCode(v int64) *DescribeModelOssTokenResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetKey(v string) *DescribeModelOssTokenResponseBody {
	s.Key = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetPolicy(v string) *DescribeModelOssTokenResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetRequestId(v string) *DescribeModelOssTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetResultObject(v *DescribeModelOssTokenResponseBodyResultObject) *DescribeModelOssTokenResponseBody {
	s.ResultObject = v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetSignature(v string) *DescribeModelOssTokenResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetSuccess(v bool) *DescribeModelOssTokenResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) SetXOssSecurityToken(v string) *DescribeModelOssTokenResponseBody {
	s.XOssSecurityToken = &v
	return s
}

func (s *DescribeModelOssTokenResponseBody) Validate() error {
	if s.ResultObject != nil {
		if err := s.ResultObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeModelOssTokenResponseBodyResultObject struct {
	// example:
	//
	// Lxxxxxxxxxxxxxxxxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// example:
	//
	// https://safxxxxxxxxx.aliyuncs.com
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// example:
	//
	// saf/xxxxxx/xxxxxx
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// eyJleHBpcmF0aW9uIjoiMxxxxxxxxxxxxxxxxxxxxxxxxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// example:
	//
	// lUxxxxxxxxxxxxxxxxxxxx
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// example:
	//
	// smxxxxxxxxxxx
	XOssSecurityToken *string `json:"XOssSecurityToken,omitempty" xml:"XOssSecurityToken,omitempty"`
}

func (s DescribeModelOssTokenResponseBodyResultObject) String() string {
	return dara.Prettify(s)
}

func (s DescribeModelOssTokenResponseBodyResultObject) GoString() string {
	return s.String()
}

func (s *DescribeModelOssTokenResponseBodyResultObject) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribeModelOssTokenResponseBodyResultObject) GetHost() *string {
	return s.Host
}

func (s *DescribeModelOssTokenResponseBodyResultObject) GetKey() *string {
	return s.Key
}

func (s *DescribeModelOssTokenResponseBodyResultObject) GetPolicy() *string {
	return s.Policy
}

func (s *DescribeModelOssTokenResponseBodyResultObject) GetSignature() *string {
	return s.Signature
}

func (s *DescribeModelOssTokenResponseBodyResultObject) GetXOssSecurityToken() *string {
	return s.XOssSecurityToken
}

func (s *DescribeModelOssTokenResponseBodyResultObject) SetAccessId(v string) *DescribeModelOssTokenResponseBodyResultObject {
	s.AccessId = &v
	return s
}

func (s *DescribeModelOssTokenResponseBodyResultObject) SetHost(v string) *DescribeModelOssTokenResponseBodyResultObject {
	s.Host = &v
	return s
}

func (s *DescribeModelOssTokenResponseBodyResultObject) SetKey(v string) *DescribeModelOssTokenResponseBodyResultObject {
	s.Key = &v
	return s
}

func (s *DescribeModelOssTokenResponseBodyResultObject) SetPolicy(v string) *DescribeModelOssTokenResponseBodyResultObject {
	s.Policy = &v
	return s
}

func (s *DescribeModelOssTokenResponseBodyResultObject) SetSignature(v string) *DescribeModelOssTokenResponseBodyResultObject {
	s.Signature = &v
	return s
}

func (s *DescribeModelOssTokenResponseBodyResultObject) SetXOssSecurityToken(v string) *DescribeModelOssTokenResponseBodyResultObject {
	s.XOssSecurityToken = &v
	return s
}

func (s *DescribeModelOssTokenResponseBodyResultObject) Validate() error {
	return dara.Validate(s)
}
