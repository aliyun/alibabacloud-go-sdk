// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListKeyPairsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetKeyPairs(v []*ListKeyPairsResponseBodyKeyPairs) *ListKeyPairsResponseBody
	GetKeyPairs() []*ListKeyPairsResponseBodyKeyPairs
	SetPageNumber(v int32) *ListKeyPairsResponseBody
	GetPageNumber() *int32
	SetPageSize(v int32) *ListKeyPairsResponseBody
	GetPageSize() *int32
	SetRequestId(v string) *ListKeyPairsResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListKeyPairsResponseBody
	GetTotalCount() *int32
}

type ListKeyPairsResponseBody struct {
	// The collection of key pair information.
	KeyPairs []*ListKeyPairsResponseBodyKeyPairs `json:"KeyPairs,omitempty" xml:"KeyPairs,omitempty" type:"Repeated"`
	// The current page number.
	//
	// example:
	//
	// 1
	PageNumber *int32 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries to return on each page.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 20758A-585D-4A41-A9B2-28DA8F4F534F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of key pairs.
	//
	// example:
	//
	// 3
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListKeyPairsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListKeyPairsResponseBody) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponseBody) GetKeyPairs() []*ListKeyPairsResponseBodyKeyPairs {
	return s.KeyPairs
}

func (s *ListKeyPairsResponseBody) GetPageNumber() *int32 {
	return s.PageNumber
}

func (s *ListKeyPairsResponseBody) GetPageSize() *int32 {
	return s.PageSize
}

func (s *ListKeyPairsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListKeyPairsResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListKeyPairsResponseBody) SetKeyPairs(v []*ListKeyPairsResponseBodyKeyPairs) *ListKeyPairsResponseBody {
	s.KeyPairs = v
	return s
}

func (s *ListKeyPairsResponseBody) SetPageNumber(v int32) *ListKeyPairsResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListKeyPairsResponseBody) SetPageSize(v int32) *ListKeyPairsResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListKeyPairsResponseBody) SetRequestId(v string) *ListKeyPairsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListKeyPairsResponseBody) SetTotalCount(v int32) *ListKeyPairsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListKeyPairsResponseBody) Validate() error {
	if s.KeyPairs != nil {
		for _, item := range s.KeyPairs {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListKeyPairsResponseBodyKeyPairs struct {
	// The time when the key pair was created. The time follows the UTC (Z) time zone and is in the YYYY-MM-DDThh:mmZ format.
	//
	// example:
	//
	// 2024-05-06T02:28Z
	CreationTime *string `json:"CreationTime,omitempty" xml:"CreationTime,omitempty"`
	// The IDs of the simple application server instances. A maximum of 50 instance IDs are supported.
	InstanceIds []*string `json:"InstanceIds,omitempty" xml:"InstanceIds,omitempty" type:"Repeated"`
	// The name of the key pair. The name must be 2 to 64 characters in length and can contain letters and Chinese characters. It must start with an uppercase or lowercase letter or a Chinese character, and cannot start with http:// or https://. It can contain digits, colons (:), underscores (_), or hyphens (-).
	//
	// example:
	//
	// testKeyPairName
	KeyPairName *string `json:"KeyPairName,omitempty" xml:"KeyPairName,omitempty"`
	// The content of the public key.
	//
	// example:
	//
	// ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCbO5Govwhb0iHzoMYKkIQxjlHyHH8nxFsW6KF5saxgYhOwdeIpWngpi+/NDWQKvuOnXFFDh/o3eJJkh3rqP+RlMggt4HLQWOd9TS0f4/cgbAzud1caW9PnankCr****
	PublicKey *string `json:"PublicKey,omitempty" xml:"PublicKey,omitempty"`
}

func (s ListKeyPairsResponseBodyKeyPairs) String() string {
	return dara.Prettify(s)
}

func (s ListKeyPairsResponseBodyKeyPairs) GoString() string {
	return s.String()
}

func (s *ListKeyPairsResponseBodyKeyPairs) GetCreationTime() *string {
	return s.CreationTime
}

func (s *ListKeyPairsResponseBodyKeyPairs) GetInstanceIds() []*string {
	return s.InstanceIds
}

func (s *ListKeyPairsResponseBodyKeyPairs) GetKeyPairName() *string {
	return s.KeyPairName
}

func (s *ListKeyPairsResponseBodyKeyPairs) GetPublicKey() *string {
	return s.PublicKey
}

func (s *ListKeyPairsResponseBodyKeyPairs) SetCreationTime(v string) *ListKeyPairsResponseBodyKeyPairs {
	s.CreationTime = &v
	return s
}

func (s *ListKeyPairsResponseBodyKeyPairs) SetInstanceIds(v []*string) *ListKeyPairsResponseBodyKeyPairs {
	s.InstanceIds = v
	return s
}

func (s *ListKeyPairsResponseBodyKeyPairs) SetKeyPairName(v string) *ListKeyPairsResponseBodyKeyPairs {
	s.KeyPairName = &v
	return s
}

func (s *ListKeyPairsResponseBodyKeyPairs) SetPublicKey(v string) *ListKeyPairsResponseBodyKeyPairs {
	s.PublicKey = &v
	return s
}

func (s *ListKeyPairsResponseBodyKeyPairs) Validate() error {
	return dara.Validate(s)
}
