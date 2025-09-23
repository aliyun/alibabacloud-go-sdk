// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTokenResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetTokenResponseBody
	GetCode() *string
	SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody
	GetData() *GetTokenResponseBodyData
	SetMessage(v string) *GetTokenResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetTokenResponseBody
	GetRequestId() *string
}

type GetTokenResponseBody struct {
	// The HTTP status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the token.
	Data *GetTokenResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The returned message.
	//
	// example:
	//
	// Success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// A07FFDF2-78FA-1B48-9E38-88E833A93187
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTokenResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetTokenResponseBody) GetData() *GetTokenResponseBodyData {
	return s.Data
}

func (s *GetTokenResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetTokenResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTokenResponseBody) SetCode(v string) *GetTokenResponseBody {
	s.Code = &v
	return s
}

func (s *GetTokenResponseBody) SetData(v *GetTokenResponseBodyData) *GetTokenResponseBody {
	s.Data = v
	return s
}

func (s *GetTokenResponseBody) SetMessage(v string) *GetTokenResponseBody {
	s.Message = &v
	return s
}

func (s *GetTokenResponseBody) SetRequestId(v string) *GetTokenResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTokenResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetTokenResponseBodyData struct {
	// The AccessKey ID that is used to access OSS.
	//
	// example:
	//
	// STS.NTm*****8tu
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret used to access OSS.
	//
	// example:
	//
	// 9NG*****K4X
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The OSS bucket that is used to store the architecture image.
	//
	// example:
	//
	// bucket-1
	Bucket *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	// The OSS endpoint.
	//
	// example:
	//
	// https://oss-cn-beijing.aliyuncs.com
	Endpoint *string `json:"Endpoint,omitempty" xml:"Endpoint,omitempty"`
	// The token that is used to access the Object Storage Service (OSS) bucket that stores the architecture image.
	//
	// example:
	//
	// ABCD
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The OSS bucket that is used to save data snapshots.
	//
	// example:
	//
	// bucket-2
	SnapshotBucket *string `json:"SnapshotBucket,omitempty" xml:"SnapshotBucket,omitempty"`
}

func (s GetTokenResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetTokenResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetTokenResponseBodyData) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *GetTokenResponseBodyData) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *GetTokenResponseBodyData) GetBucket() *string {
	return s.Bucket
}

func (s *GetTokenResponseBodyData) GetEndpoint() *string {
	return s.Endpoint
}

func (s *GetTokenResponseBodyData) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *GetTokenResponseBodyData) GetSnapshotBucket() *string {
	return s.SnapshotBucket
}

func (s *GetTokenResponseBodyData) SetAccessKeyId(v string) *GetTokenResponseBodyData {
	s.AccessKeyId = &v
	return s
}

func (s *GetTokenResponseBodyData) SetAccessKeySecret(v string) *GetTokenResponseBodyData {
	s.AccessKeySecret = &v
	return s
}

func (s *GetTokenResponseBodyData) SetBucket(v string) *GetTokenResponseBodyData {
	s.Bucket = &v
	return s
}

func (s *GetTokenResponseBodyData) SetEndpoint(v string) *GetTokenResponseBodyData {
	s.Endpoint = &v
	return s
}

func (s *GetTokenResponseBodyData) SetSecurityToken(v string) *GetTokenResponseBodyData {
	s.SecurityToken = &v
	return s
}

func (s *GetTokenResponseBodyData) SetSnapshotBucket(v string) *GetTokenResponseBodyData {
	s.SnapshotBucket = &v
	return s
}

func (s *GetTokenResponseBodyData) Validate() error {
	return dara.Validate(s)
}
