// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryWildcardByTopicRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFirstTopic(v string) *QueryWildcardByTopicRequest
	GetFirstTopic() *string
	SetInstanceId(v string) *QueryWildcardByTopicRequest
	GetInstanceId() *string
	SetPageNo(v int32) *QueryWildcardByTopicRequest
	GetPageNo() *int32
	SetPageSize(v int32) *QueryWildcardByTopicRequest
	GetPageSize() *int32
}

type QueryWildcardByTopicRequest struct {
	// The parent topic.
	//
	// This parameter is required.
	//
	// example:
	//
	// test
	FirstTopic *string `json:"FirstTopic,omitempty" xml:"FirstTopic,omitempty"`
	// The ID of the ApsaraMQ for MQTT instance.
	//
	// This parameter is required.
	//
	// example:
	//
	// mqtt-cn-0pp1ldu****
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// The page number of the results to return. Pages start from page 1.
	//
	// example:
	//
	// 1
	PageNo *int32 `json:"PageNo,omitempty" xml:"PageNo,omitempty"`
	// The maximum number of records to display per page.
	//
	// Valid values: 1 to 100.
	//
	// example:
	//
	// 10
	PageSize *int32 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
}

func (s QueryWildcardByTopicRequest) String() string {
	return dara.Prettify(s)
}

func (s QueryWildcardByTopicRequest) GoString() string {
	return s.String()
}

func (s *QueryWildcardByTopicRequest) GetFirstTopic() *string {
	return s.FirstTopic
}

func (s *QueryWildcardByTopicRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *QueryWildcardByTopicRequest) GetPageNo() *int32 {
	return s.PageNo
}

func (s *QueryWildcardByTopicRequest) GetPageSize() *int32 {
	return s.PageSize
}

func (s *QueryWildcardByTopicRequest) SetFirstTopic(v string) *QueryWildcardByTopicRequest {
	s.FirstTopic = &v
	return s
}

func (s *QueryWildcardByTopicRequest) SetInstanceId(v string) *QueryWildcardByTopicRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryWildcardByTopicRequest) SetPageNo(v int32) *QueryWildcardByTopicRequest {
	s.PageNo = &v
	return s
}

func (s *QueryWildcardByTopicRequest) SetPageSize(v int32) *QueryWildcardByTopicRequest {
	s.PageSize = &v
	return s
}

func (s *QueryWildcardByTopicRequest) Validate() error {
	return dara.Validate(s)
}
