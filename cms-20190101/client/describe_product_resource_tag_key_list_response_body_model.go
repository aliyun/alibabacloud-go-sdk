// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeProductResourceTagKeyListResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *DescribeProductResourceTagKeyListResponseBody
	GetCode() *string
	SetMessage(v string) *DescribeProductResourceTagKeyListResponseBody
	GetMessage() *string
	SetNextToken(v string) *DescribeProductResourceTagKeyListResponseBody
	GetNextToken() *string
	SetRequestId(v string) *DescribeProductResourceTagKeyListResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeProductResourceTagKeyListResponseBody
	GetSuccess() *bool
	SetTagKeys(v *DescribeProductResourceTagKeyListResponseBodyTagKeys) *DescribeProductResourceTagKeyListResponseBody
	GetTagKeys() *DescribeProductResourceTagKeyListResponseBodyTagKeys
}

type DescribeProductResourceTagKeyListResponseBody struct {
	// The status code.
	//
	// >  The status code 200 indicates that the request was successful.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The error message.
	//
	// example:
	//
	// The request processing has failed due to some unknown error.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The pagination token. If more entries are to be returned on the next page, a pagination token is returned.
	//
	// >  If the value of this parameter is not null, more entries are to be returned on the next page. You can use the returned pagination token as a request parameter to retrieve a new page of results. If the value of this parameter is null, all the entries have been returned.
	//
	// example:
	//
	// dbc2826f237e****
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 945ACAA9-89F2-4A62-8913-076FDEDAA8DB
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values: true and false.
	//
	// example:
	//
	// true
	Success *bool                                                 `json:"Success,omitempty" xml:"Success,omitempty"`
	TagKeys *DescribeProductResourceTagKeyListResponseBodyTagKeys `json:"TagKeys,omitempty" xml:"TagKeys,omitempty" type:"Struct"`
}

func (s DescribeProductResourceTagKeyListResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResourceTagKeyListResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeProductResourceTagKeyListResponseBody) GetCode() *string {
	return s.Code
}

func (s *DescribeProductResourceTagKeyListResponseBody) GetMessage() *string {
	return s.Message
}

func (s *DescribeProductResourceTagKeyListResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeProductResourceTagKeyListResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeProductResourceTagKeyListResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeProductResourceTagKeyListResponseBody) GetTagKeys() *DescribeProductResourceTagKeyListResponseBodyTagKeys {
	return s.TagKeys
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetCode(v string) *DescribeProductResourceTagKeyListResponseBody {
	s.Code = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetMessage(v string) *DescribeProductResourceTagKeyListResponseBody {
	s.Message = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetNextToken(v string) *DescribeProductResourceTagKeyListResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetRequestId(v string) *DescribeProductResourceTagKeyListResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetSuccess(v bool) *DescribeProductResourceTagKeyListResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) SetTagKeys(v *DescribeProductResourceTagKeyListResponseBodyTagKeys) *DescribeProductResourceTagKeyListResponseBody {
	s.TagKeys = v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBody) Validate() error {
	if s.TagKeys != nil {
		if err := s.TagKeys.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeProductResourceTagKeyListResponseBodyTagKeys struct {
	TagKey []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
}

func (s DescribeProductResourceTagKeyListResponseBodyTagKeys) String() string {
	return dara.Prettify(s)
}

func (s DescribeProductResourceTagKeyListResponseBodyTagKeys) GoString() string {
	return s.String()
}

func (s *DescribeProductResourceTagKeyListResponseBodyTagKeys) GetTagKey() []*string {
	return s.TagKey
}

func (s *DescribeProductResourceTagKeyListResponseBodyTagKeys) SetTagKey(v []*string) *DescribeProductResourceTagKeyListResponseBodyTagKeys {
	s.TagKey = v
	return s
}

func (s *DescribeProductResourceTagKeyListResponseBodyTagKeys) Validate() error {
	return dara.Validate(s)
}
