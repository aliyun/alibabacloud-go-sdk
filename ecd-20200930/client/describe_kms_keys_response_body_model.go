// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeKmsKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAuthorizeStatus(v string) *DescribeKmsKeysResponseBody
	GetAuthorizeStatus() *string
	SetKeys(v []*DescribeKmsKeysResponseBodyKeys) *DescribeKmsKeysResponseBody
	GetKeys() []*DescribeKmsKeysResponseBodyKeys
	SetKmsServiceStatus(v string) *DescribeKmsKeysResponseBody
	GetKmsServiceStatus() *string
	SetRequestId(v string) *DescribeKmsKeysResponseBody
	GetRequestId() *string
}

type DescribeKmsKeysResponseBody struct {
	// The authorization status.
	//
	// Valid value:
	//
	// 	- not_authorized
	//
	// 	- authorized
	//
	// example:
	//
	// authorized
	AuthorizeStatus *string `json:"AuthorizeStatus,omitempty" xml:"AuthorizeStatus,omitempty"`
	// Customer master key (CMK)
	Keys []*DescribeKmsKeysResponseBodyKeys `json:"Keys,omitempty" xml:"Keys,omitempty" type:"Repeated"`
	// Indicates whether KMS is activated.
	//
	// Valid value:
	//
	// 	- disabled
	//
	// 	- enabled
	//
	// example:
	//
	// enabled
	KmsServiceStatus *string `json:"KmsServiceStatus,omitempty" xml:"KmsServiceStatus,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// 1CBAFFAB-B697-4049-A9B1-67E1FC5F****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeKmsKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysResponseBody) GetAuthorizeStatus() *string {
	return s.AuthorizeStatus
}

func (s *DescribeKmsKeysResponseBody) GetKeys() []*DescribeKmsKeysResponseBodyKeys {
	return s.Keys
}

func (s *DescribeKmsKeysResponseBody) GetKmsServiceStatus() *string {
	return s.KmsServiceStatus
}

func (s *DescribeKmsKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeKmsKeysResponseBody) SetAuthorizeStatus(v string) *DescribeKmsKeysResponseBody {
	s.AuthorizeStatus = &v
	return s
}

func (s *DescribeKmsKeysResponseBody) SetKeys(v []*DescribeKmsKeysResponseBodyKeys) *DescribeKmsKeysResponseBody {
	s.Keys = v
	return s
}

func (s *DescribeKmsKeysResponseBody) SetKmsServiceStatus(v string) *DescribeKmsKeysResponseBody {
	s.KmsServiceStatus = &v
	return s
}

func (s *DescribeKmsKeysResponseBody) SetRequestId(v string) *DescribeKmsKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeKmsKeysResponseBody) Validate() error {
	if s.Keys != nil {
		for _, item := range s.Keys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeKmsKeysResponseBodyKeys struct {
	// The alias of the key.
	//
	// example:
	//
	// TestAlias
	Alias *string `json:"Alias,omitempty" xml:"Alias,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the key in KMS.
	//
	// example:
	//
	// acs:kms:cn-hangzhou:154035569884****:key/05754286-3ba2-4fa6-8d41-4323aca6****
	Arn *string `json:"Arn,omitempty" xml:"Arn,omitempty"`
	// The ID of the key.
	//
	// example:
	//
	// 05754286-3ba2-4fa6-8d41-4323aca6****
	KeyId *string `json:"KeyId,omitempty" xml:"KeyId,omitempty"`
	// The type of the key.
	//
	// example:
	//
	// ServiceKey
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeKmsKeysResponseBodyKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeKmsKeysResponseBodyKeys) GoString() string {
	return s.String()
}

func (s *DescribeKmsKeysResponseBodyKeys) GetAlias() *string {
	return s.Alias
}

func (s *DescribeKmsKeysResponseBodyKeys) GetArn() *string {
	return s.Arn
}

func (s *DescribeKmsKeysResponseBodyKeys) GetKeyId() *string {
	return s.KeyId
}

func (s *DescribeKmsKeysResponseBodyKeys) GetType() *string {
	return s.Type
}

func (s *DescribeKmsKeysResponseBodyKeys) SetAlias(v string) *DescribeKmsKeysResponseBodyKeys {
	s.Alias = &v
	return s
}

func (s *DescribeKmsKeysResponseBodyKeys) SetArn(v string) *DescribeKmsKeysResponseBodyKeys {
	s.Arn = &v
	return s
}

func (s *DescribeKmsKeysResponseBodyKeys) SetKeyId(v string) *DescribeKmsKeysResponseBodyKeys {
	s.KeyId = &v
	return s
}

func (s *DescribeKmsKeysResponseBodyKeys) SetType(v string) *DescribeKmsKeysResponseBodyKeys {
	s.Type = &v
	return s
}

func (s *DescribeKmsKeysResponseBodyKeys) Validate() error {
	return dara.Validate(s)
}
