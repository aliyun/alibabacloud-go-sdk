// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListBlacklistCallTaggingsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListBlacklistCallTaggingsRequest
	GetInstanceId() *string
	SetNumberList(v string) *ListBlacklistCallTaggingsRequest
	GetNumberList() *string
}

type ListBlacklistCallTaggingsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// [{"number":"1521083xxxx","jobId":"job-481841171213393920"}]
	NumberList *string `json:"NumberList,omitempty" xml:"NumberList,omitempty"`
}

func (s ListBlacklistCallTaggingsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListBlacklistCallTaggingsRequest) GoString() string {
	return s.String()
}

func (s *ListBlacklistCallTaggingsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListBlacklistCallTaggingsRequest) GetNumberList() *string {
	return s.NumberList
}

func (s *ListBlacklistCallTaggingsRequest) SetInstanceId(v string) *ListBlacklistCallTaggingsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListBlacklistCallTaggingsRequest) SetNumberList(v string) *ListBlacklistCallTaggingsRequest {
	s.NumberList = &v
	return s
}

func (s *ListBlacklistCallTaggingsRequest) Validate() error {
	return dara.Validate(s)
}
