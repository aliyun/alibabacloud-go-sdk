// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUserEncryptionKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDBClusterId(v string) *DescribeUserEncryptionKeyListResponseBody
	GetDBClusterId() *string
	SetKeyList(v []*string) *DescribeUserEncryptionKeyListResponseBody
	GetKeyList() []*string
	SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody
	GetRequestId() *string
}

type DescribeUserEncryptionKeyListResponseBody struct {
	// The ID of the cluster.
	//
	// example:
	//
	// pc-************
	DBClusterId *string `json:"DBClusterId,omitempty" xml:"DBClusterId,omitempty"`
	// Cluster key list.
	KeyList []*string `json:"KeyList,omitempty" xml:"KeyList,omitempty" type:"Repeated"`
	// The ID of the request.
	//
	// example:
	//
	// A7E6A8FD-C50B-46B2-BA85-D8B8D3******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeUserEncryptionKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUserEncryptionKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetDBClusterId() *string {
	return s.DBClusterId
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetKeyList() []*string {
	return s.KeyList
}

func (s *DescribeUserEncryptionKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetDBClusterId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.DBClusterId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetKeyList(v []*string) *DescribeUserEncryptionKeyListResponseBody {
	s.KeyList = v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) SetRequestId(v string) *DescribeUserEncryptionKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUserEncryptionKeyListResponseBody) Validate() error {
	return dara.Validate(s)
}
