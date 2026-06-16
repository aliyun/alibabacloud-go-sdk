// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePocOssTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessId(v string) *DescribePocOssTokenResponseBody
	GetAccessId() *string
	SetHost(v string) *DescribePocOssTokenResponseBody
	GetHost() *string
	SetKey(v string) *DescribePocOssTokenResponseBody
	GetKey() *string
	SetPolicy(v string) *DescribePocOssTokenResponseBody
	GetPolicy() *string
	SetRequestId(v string) *DescribePocOssTokenResponseBody
	GetRequestId() *string
	SetSignature(v string) *DescribePocOssTokenResponseBody
	GetSignature() *string
	SetResultObject(v bool) *DescribePocOssTokenResponseBody
	GetResultObject() *bool
}

type DescribePocOssTokenResponseBody struct {
	// The AccessKey ID used to upload files to OSS.
	//
	// example:
	//
	// LTAxxxxxxxxx
	AccessId *string `json:"AccessId,omitempty" xml:"AccessId,omitempty"`
	// The host address.
	//
	// example:
	//
	// 192.168.34.191
	Host *string `json:"Host,omitempty" xml:"Host,omitempty"`
	// The key required for file upload.
	//
	// example:
	//
	// saf/de/namelist/e924/ufzgsedX9bd3a7
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// The OSS security policy.
	//
	// example:
	//
	// eyJleHBpxxxcmF0aW9xxxIjoiMjAyNS0wNy0zMFQxxxxxwNjowNTo0OSxxxxx45NTRaIiwiY29uxxxZGl0a5Ymxxxxxxxxxxxxx
	Policy *string `json:"Policy,omitempty" xml:"Policy,omitempty"`
	// The request ID.
	//
	// example:
	//
	// A32FE941-35F2-5378-B37C-4B8FDB16F094
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The signature information for the upload.
	//
	// example:
	//
	// 0lxQEWM0BqHd476JJE0fNXdS3UA=
	Signature *string `json:"Signature,omitempty" xml:"Signature,omitempty"`
	// The returned result.
	//
	// example:
	//
	// true
	ResultObject *bool `json:"resultObject,omitempty" xml:"resultObject,omitempty"`
}

func (s DescribePocOssTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePocOssTokenResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePocOssTokenResponseBody) GetAccessId() *string {
	return s.AccessId
}

func (s *DescribePocOssTokenResponseBody) GetHost() *string {
	return s.Host
}

func (s *DescribePocOssTokenResponseBody) GetKey() *string {
	return s.Key
}

func (s *DescribePocOssTokenResponseBody) GetPolicy() *string {
	return s.Policy
}

func (s *DescribePocOssTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePocOssTokenResponseBody) GetSignature() *string {
	return s.Signature
}

func (s *DescribePocOssTokenResponseBody) GetResultObject() *bool {
	return s.ResultObject
}

func (s *DescribePocOssTokenResponseBody) SetAccessId(v string) *DescribePocOssTokenResponseBody {
	s.AccessId = &v
	return s
}

func (s *DescribePocOssTokenResponseBody) SetHost(v string) *DescribePocOssTokenResponseBody {
	s.Host = &v
	return s
}

func (s *DescribePocOssTokenResponseBody) SetKey(v string) *DescribePocOssTokenResponseBody {
	s.Key = &v
	return s
}

func (s *DescribePocOssTokenResponseBody) SetPolicy(v string) *DescribePocOssTokenResponseBody {
	s.Policy = &v
	return s
}

func (s *DescribePocOssTokenResponseBody) SetRequestId(v string) *DescribePocOssTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePocOssTokenResponseBody) SetSignature(v string) *DescribePocOssTokenResponseBody {
	s.Signature = &v
	return s
}

func (s *DescribePocOssTokenResponseBody) SetResultObject(v bool) *DescribePocOssTokenResponseBody {
	s.ResultObject = &v
	return s
}

func (s *DescribePocOssTokenResponseBody) Validate() error {
	return dara.Validate(s)
}
