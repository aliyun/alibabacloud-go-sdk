// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAccessKeyResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAccessKey(v *CreateAccessKeyResponseBodyAccessKey) *CreateAccessKeyResponseBody
	GetAccessKey() *CreateAccessKeyResponseBodyAccessKey
	SetRequestId(v string) *CreateAccessKeyResponseBody
	GetRequestId() *string
}

type CreateAccessKeyResponseBody struct {
	// The information about the AccessKey pair.
	AccessKey *CreateAccessKeyResponseBodyAccessKey `json:"AccessKey,omitempty" xml:"AccessKey,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 19DDD9F7-AFCC-4D72-8CBA-CCE5A142E7AB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAccessKeyResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessKeyResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponseBody) GetAccessKey() *CreateAccessKeyResponseBodyAccessKey {
	return s.AccessKey
}

func (s *CreateAccessKeyResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAccessKeyResponseBody) SetAccessKey(v *CreateAccessKeyResponseBodyAccessKey) *CreateAccessKeyResponseBody {
	s.AccessKey = v
	return s
}

func (s *CreateAccessKeyResponseBody) SetRequestId(v string) *CreateAccessKeyResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAccessKeyResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateAccessKeyResponseBodyAccessKey struct {
	// The AccessKey ID.
	//
	// example:
	//
	// yourAccessKeyID
	AccessKeyId *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	// The AccessKey secret.
	//
	// example:
	//
	// yourAccessKeySecret
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	// The time when the AccessKey pair was created.
	//
	// example:
	//
	// 2020-10-15T08:08:54Z
	CreateDate *string `json:"CreateDate,omitempty" xml:"CreateDate,omitempty"`
	// The status of the AccessKey pair. Valid values:
	//
	// 	- Active
	//
	// 	- Inactive
	//
	// example:
	//
	// Active
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateAccessKeyResponseBodyAccessKey) String() string {
	return dara.Prettify(s)
}

func (s CreateAccessKeyResponseBodyAccessKey) GoString() string {
	return s.String()
}

func (s *CreateAccessKeyResponseBodyAccessKey) GetAccessKeyId() *string {
	return s.AccessKeyId
}

func (s *CreateAccessKeyResponseBodyAccessKey) GetAccessKeySecret() *string {
	return s.AccessKeySecret
}

func (s *CreateAccessKeyResponseBodyAccessKey) GetCreateDate() *string {
	return s.CreateDate
}

func (s *CreateAccessKeyResponseBodyAccessKey) GetStatus() *string {
	return s.Status
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetAccessKeyId(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.AccessKeyId = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetAccessKeySecret(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.AccessKeySecret = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetCreateDate(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.CreateDate = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) SetStatus(v string) *CreateAccessKeyResponseBodyAccessKey {
	s.Status = &v
	return s
}

func (s *CreateAccessKeyResponseBodyAccessKey) Validate() error {
	return dara.Validate(s)
}
