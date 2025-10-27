// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLogStoreKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLogStoreKeys(v *DescribeLogStoreKeysResponseBodyLogStoreKeys) *DescribeLogStoreKeysResponseBody
	GetLogStoreKeys() *DescribeLogStoreKeysResponseBodyLogStoreKeys
	SetRequestId(v string) *DescribeLogStoreKeysResponseBody
	GetRequestId() *string
}

type DescribeLogStoreKeysResponseBody struct {
	// The queried log keywords.
	LogStoreKeys *DescribeLogStoreKeysResponseBodyLogStoreKeys `json:"LogStoreKeys,omitempty" xml:"LogStoreKeys,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// 3BB185E9-BB54-1727-B876-13243E4C0EB5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeLogStoreKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreKeysResponseBody) GetLogStoreKeys() *DescribeLogStoreKeysResponseBodyLogStoreKeys {
	return s.LogStoreKeys
}

func (s *DescribeLogStoreKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLogStoreKeysResponseBody) SetLogStoreKeys(v *DescribeLogStoreKeysResponseBodyLogStoreKeys) *DescribeLogStoreKeysResponseBody {
	s.LogStoreKeys = v
	return s
}

func (s *DescribeLogStoreKeysResponseBody) SetRequestId(v string) *DescribeLogStoreKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLogStoreKeysResponseBody) Validate() error {
	if s.LogStoreKeys != nil {
		if err := s.LogStoreKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeLogStoreKeysResponseBodyLogStoreKeys struct {
	LogStoreKey []*string `json:"LogStoreKey,omitempty" xml:"LogStoreKey,omitempty" type:"Repeated"`
}

func (s DescribeLogStoreKeysResponseBodyLogStoreKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeLogStoreKeysResponseBodyLogStoreKeys) GoString() string {
	return s.String()
}

func (s *DescribeLogStoreKeysResponseBodyLogStoreKeys) GetLogStoreKey() []*string {
	return s.LogStoreKey
}

func (s *DescribeLogStoreKeysResponseBodyLogStoreKeys) SetLogStoreKey(v []*string) *DescribeLogStoreKeysResponseBodyLogStoreKeys {
	s.LogStoreKey = v
	return s
}

func (s *DescribeLogStoreKeysResponseBodyLogStoreKeys) Validate() error {
	return dara.Validate(s)
}
