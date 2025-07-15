// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagKeysResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeTagKeysResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeTagKeysResponseBody
	GetRequestId() *string
	SetTagKeys(v *DescribeTagKeysResponseBodyTagKeys) *DescribeTagKeysResponseBody
	GetTagKeys() *DescribeTagKeysResponseBodyTagKeys
}

type DescribeTagKeysResponseBody struct {
	// The token that is used for the next query. Valid values:
	//
	// 	- If the value of **NextToken*	- is not returned, it indicates that no next query is to be sent.
	//
	// 	- If a value of **NextToken*	- is returned, the value is the token that is used for the subsequent query.
	//
	// example:
	//
	// FFmyTO70tTpLG6I3FmYAXGKPd****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// DE65F6B7-7566-4802-9007-96F2494AC512
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of tag keys.
	TagKeys *DescribeTagKeysResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Struct"`
}

func (s DescribeTagKeysResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeTagKeysResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagKeysResponseBody) GetTagKeys() *DescribeTagKeysResponseBodyTagKeys {
	return s.TagKeys
}

func (s *DescribeTagKeysResponseBody) SetNextToken(v string) *DescribeTagKeysResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetRequestId(v string) *DescribeTagKeysResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagKeysResponseBody) SetTagKeys(v *DescribeTagKeysResponseBodyTagKeys) *DescribeTagKeysResponseBody {
	s.TagKeys = v
	return s
}

func (s *DescribeTagKeysResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeTagKeysResponseBodyTagKeys struct {
	TagKey []*DescribeTagKeysResponseBodyTagKeysTagKey `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s DescribeTagKeysResponseBodyTagKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponseBodyTagKeys) GetTagKey() []*DescribeTagKeysResponseBodyTagKeysTagKey {
	return s.TagKey
}

func (s *DescribeTagKeysResponseBodyTagKeys) SetTagKey(v []*DescribeTagKeysResponseBodyTagKeysTagKey) *DescribeTagKeysResponseBodyTagKeys {
	s.TagKey = v
	return s
}

func (s *DescribeTagKeysResponseBodyTagKeys) Validate() error {
	return dara.Validate(s)
}

type DescribeTagKeysResponseBodyTagKeysTagKey struct {
	// The tag key.
	//
	// example:
	//
	// FinanceDept
	TagKey *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	// The type of the tag key. Valid values:
	//
	// 	- **Custom**: custom
	//
	// 	- **System**: system
	//
	// example:
	//
	// Custom
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s DescribeTagKeysResponseBodyTagKeysTagKey) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagKeysResponseBodyTagKeysTagKey) GoString() string {
	return s.String()
}

func (s *DescribeTagKeysResponseBodyTagKeysTagKey) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagKeysResponseBodyTagKeysTagKey) GetType() *string {
	return s.Type
}

func (s *DescribeTagKeysResponseBodyTagKeysTagKey) SetTagKey(v string) *DescribeTagKeysResponseBodyTagKeysTagKey {
	s.TagKey = &v
	return s
}

func (s *DescribeTagKeysResponseBodyTagKeysTagKey) SetType(v string) *DescribeTagKeysResponseBodyTagKeysTagKey {
	s.Type = &v
	return s
}

func (s *DescribeTagKeysResponseBodyTagKeysTagKey) Validate() error {
	return dara.Validate(s)
}
