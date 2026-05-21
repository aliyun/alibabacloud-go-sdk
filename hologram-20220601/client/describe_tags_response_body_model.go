// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeTagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetErrorCode(v string) *DescribeTagsResponseBody
	GetErrorCode() *string
	SetErrorMessage(v string) *DescribeTagsResponseBody
	GetErrorMessage() *string
	SetHttpStatusCode(v string) *DescribeTagsResponseBody
	GetHttpStatusCode() *string
	SetRequestId(v string) *DescribeTagsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DescribeTagsResponseBody
	GetSuccess() *bool
	SetTagResources(v []*DescribeTagsResponseBodyTagResources) *DescribeTagsResponseBody
	GetTagResources() []*DescribeTagsResponseBodyTagResources
}

type DescribeTagsResponseBody struct {
	// example:
	//
	// 404
	ErrorCode *string `json:"ErrorCode,omitempty" xml:"ErrorCode,omitempty"`
	// example:
	//
	// Internal server error.
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	// example:
	//
	// 200
	HttpStatusCode *string `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 14E2358E-0499-509E-8E22-CA30EC474A9A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success      *bool                                   `json:"Success,omitempty" xml:"Success,omitempty"`
	TagResources []*DescribeTagsResponseBodyTagResources `json:"TagResources,omitempty" xml:"TagResources,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBody) GetErrorCode() *string {
	return s.ErrorCode
}

func (s *DescribeTagsResponseBody) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *DescribeTagsResponseBody) GetHttpStatusCode() *string {
	return s.HttpStatusCode
}

func (s *DescribeTagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeTagsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DescribeTagsResponseBody) GetTagResources() []*DescribeTagsResponseBodyTagResources {
	return s.TagResources
}

func (s *DescribeTagsResponseBody) SetErrorCode(v string) *DescribeTagsResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *DescribeTagsResponseBody) SetErrorMessage(v string) *DescribeTagsResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *DescribeTagsResponseBody) SetHttpStatusCode(v string) *DescribeTagsResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *DescribeTagsResponseBody) SetRequestId(v string) *DescribeTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeTagsResponseBody) SetSuccess(v bool) *DescribeTagsResponseBody {
	s.Success = &v
	return s
}

func (s *DescribeTagsResponseBody) SetTagResources(v []*DescribeTagsResponseBodyTagResources) *DescribeTagsResponseBody {
	s.TagResources = v
	return s
}

func (s *DescribeTagsResponseBody) Validate() error {
	if s.TagResources != nil {
		for _, item := range s.TagResources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeTagsResponseBodyTagResources struct {
	// example:
	//
	// hgprecn-cn-0ju2uaanj001
	ResourceId *string `json:"ResourceId,omitempty" xml:"ResourceId,omitempty"`
	// example:
	//
	// Instance
	ResourceType *string `json:"ResourceType,omitempty" xml:"ResourceType,omitempty"`
	// example:
	//
	// acs:rm:rgId
	TagKey   *string   `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	TagValue []*string `json:"TagValue,omitempty" xml:"TagValue,omitempty" type:"Repeated"`
}

func (s DescribeTagsResponseBodyTagResources) String() string {
	return dara.Prettify(s)
}

func (s DescribeTagsResponseBodyTagResources) GoString() string {
	return s.String()
}

func (s *DescribeTagsResponseBodyTagResources) GetResourceId() *string {
	return s.ResourceId
}

func (s *DescribeTagsResponseBodyTagResources) GetResourceType() *string {
	return s.ResourceType
}

func (s *DescribeTagsResponseBodyTagResources) GetTagKey() *string {
	return s.TagKey
}

func (s *DescribeTagsResponseBodyTagResources) GetTagValue() []*string {
	return s.TagValue
}

func (s *DescribeTagsResponseBodyTagResources) SetResourceId(v string) *DescribeTagsResponseBodyTagResources {
	s.ResourceId = &v
	return s
}

func (s *DescribeTagsResponseBodyTagResources) SetResourceType(v string) *DescribeTagsResponseBodyTagResources {
	s.ResourceType = &v
	return s
}

func (s *DescribeTagsResponseBodyTagResources) SetTagKey(v string) *DescribeTagsResponseBodyTagResources {
	s.TagKey = &v
	return s
}

func (s *DescribeTagsResponseBodyTagResources) SetTagValue(v []*string) *DescribeTagsResponseBodyTagResources {
	s.TagValue = v
	return s
}

func (s *DescribeTagsResponseBodyTagResources) Validate() error {
	return dara.Validate(s)
}
