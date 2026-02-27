// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysForExpressConnectResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeTagKeysForExpressConnectResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTagKeysForExpressConnectResponseBody
	GetRequestId() *string
	SetTagKeys(v *DescribeTagKeysForExpressConnectResponseBodyTagKeys) *DescribeTagKeysForExpressConnectResponseBody
	GetTagKeys() *DescribeTagKeysForExpressConnectResponseBodyTagKeys
}

type DescribeTagKeysForExpressConnectResponseBody struct {
	// A pagination token. It can be used in the next request to retrieve a new page of results.
	//
	// 	- If **NextToken*	- is empty, no next page exists.
	//
	// 	- If a value is returned for **NextToken**, the value can be used in the next request to retrieve a new page of results.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 54B48E3D-DF70-471B-AA93-08E683A1B45
	RequestId *string                                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	TagKeys   *DescribeTagKeysForExpressConnectResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Struct"`
}

func (s DescribeTagKeysForExpressConnectResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysForExpressConnectResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysForExpressConnectResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagKeysForExpressConnectResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagKeysForExpressConnectResponseBody) GetTagKeys() *DescribeTagKeysForExpressConnectResponseBodyTagKeys {
	return s.TagKeys
}

func (s *DescribeTagKeysForExpressConnectResponseBody) SetNextToken(v string) *DescribeTagKeysForExpressConnectResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectResponseBody) SetRequestId(v string) *DescribeTagKeysForExpressConnectResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectResponseBody) SetTagKeys(v *DescribeTagKeysForExpressConnectResponseBodyTagKeys) *DescribeTagKeysForExpressConnectResponseBody {
	s.TagKeys = v
	return s
}

func (s *DescribeTagKeysForExpressConnectResponseBody) Validate() error {
	if s.TagKeys != nil {
		if err := s.TagKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeTagKeysForExpressConnectResponseBodyTagKeys struct {
	TagKey []*DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s DescribeTagKeysForExpressConnectResponseBodyTagKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysForExpressConnectResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysForExpressConnectResponseBodyTagKeys) GetTagKey() []*DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey {
	return s.TagKey
}

func (s *DescribeTagKeysForExpressConnectResponseBodyTagKeys) SetTagKey(v []*DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey) *DescribeTagKeysForExpressConnectResponseBodyTagKeys {
	s.TagKey = v
	return s
}

func (s *DescribeTagKeysForExpressConnectResponseBodyTagKeys) Validate() error {
	if s.TagKey != nil {
		for _, item := range s.TagKey {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey struct {
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey) GetType() *string {
	return s.Type
}

func (s *DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey) SetTagKey(v string) *DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey {
	s.TagKey = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey) SetType(v string) *DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey {
	s.Type = &v
	return s
}

func (s *DescribeTagKeysForExpressConnectResponseBodyTagKeysTagKey) Validate() error {
	return dara.Validate(s)
}
