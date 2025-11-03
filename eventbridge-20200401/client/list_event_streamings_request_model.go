// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEventStreamingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLimit(v int32) *ListEventStreamingsRequest
	GetLimit() *int32
	SetNamePrefix(v string) *ListEventStreamingsRequest
	GetNamePrefix() *string
	SetNextToken(v string) *ListEventStreamingsRequest
	GetNextToken() *string
	SetSinkArn(v string) *ListEventStreamingsRequest
	GetSinkArn() *string
	SetSourceArn(v string) *ListEventStreamingsRequest
	GetSourceArn() *string
	SetTags(v []*ListEventStreamingsRequestTags) *ListEventStreamingsRequest
	GetTags() []*ListEventStreamingsRequestTags
}

type ListEventStreamingsRequest struct {
	// The maximum number of entries to be returned in a call. You can use this parameter and NextToken to implement paging. A maximum of 100 entries can be returned in a call.
	//
	// example:
	//
	// 10
	Limit *int32 `json:"Limit,omitempty" xml:"Limit,omitempty"`
	// The name of the event stream that you want to query.
	//
	// example:
	//
	// name
	NamePrefix *string `json:"NamePrefix,omitempty" xml:"NamePrefix,omitempty"`
	// If you configure Limit and excess return values exist, this parameter is returned.
	//
	// example:
	//
	// 10
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ARN of the event target.
	//
	// example:
	//
	// acs:fc:cn-hangzhou:118609547428****:services/fw1.LATEST/functions/log1
	SinkArn *string `json:"SinkArn,omitempty" xml:"SinkArn,omitempty"`
	// The Alibaba Cloud Resource Name (ARN) of the event source.
	SourceArn *string                           `json:"SourceArn,omitempty" xml:"SourceArn,omitempty"`
	Tags      []*ListEventStreamingsRequestTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListEventStreamingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsRequest) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsRequest) GetLimit() *int32 {
	return s.Limit
}

func (s *ListEventStreamingsRequest) GetNamePrefix() *string {
	return s.NamePrefix
}

func (s *ListEventStreamingsRequest) GetNextToken() *string {
	return s.NextToken
}

func (s *ListEventStreamingsRequest) GetSinkArn() *string {
	return s.SinkArn
}

func (s *ListEventStreamingsRequest) GetSourceArn() *string {
	return s.SourceArn
}

func (s *ListEventStreamingsRequest) GetTags() []*ListEventStreamingsRequestTags {
	return s.Tags
}

func (s *ListEventStreamingsRequest) SetLimit(v int32) *ListEventStreamingsRequest {
	s.Limit = &v
	return s
}

func (s *ListEventStreamingsRequest) SetNamePrefix(v string) *ListEventStreamingsRequest {
	s.NamePrefix = &v
	return s
}

func (s *ListEventStreamingsRequest) SetNextToken(v string) *ListEventStreamingsRequest {
	s.NextToken = &v
	return s
}

func (s *ListEventStreamingsRequest) SetSinkArn(v string) *ListEventStreamingsRequest {
	s.SinkArn = &v
	return s
}

func (s *ListEventStreamingsRequest) SetSourceArn(v string) *ListEventStreamingsRequest {
	s.SourceArn = &v
	return s
}

func (s *ListEventStreamingsRequest) SetTags(v []*ListEventStreamingsRequestTags) *ListEventStreamingsRequest {
	s.Tags = v
	return s
}

func (s *ListEventStreamingsRequest) Validate() error {
	if s.Tags != nil {
		for _, item := range s.Tags {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEventStreamingsRequestTags struct {
	// example:
	//
	// mns
	Key *string `json:"Key,omitempty" xml:"Key,omitempty"`
	// example:
	//
	// mnstest
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s ListEventStreamingsRequestTags) String() string {
	return dara.Prettify(s)
}

func (s ListEventStreamingsRequestTags) GoString() string {
	return s.String()
}

func (s *ListEventStreamingsRequestTags) GetKey() *string {
	return s.Key
}

func (s *ListEventStreamingsRequestTags) GetValue() *string {
	return s.Value
}

func (s *ListEventStreamingsRequestTags) SetKey(v string) *ListEventStreamingsRequestTags {
	s.Key = &v
	return s
}

func (s *ListEventStreamingsRequestTags) SetValue(v string) *ListEventStreamingsRequestTags {
	s.Value = &v
	return s
}

func (s *ListEventStreamingsRequestTags) Validate() error {
	return dara.Validate(s)
}
