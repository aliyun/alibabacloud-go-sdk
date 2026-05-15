// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexCurrentValueRequest interface {
	dara.Model
	String() string
	GoString() string
	SetDepIds(v string) *GetIndexCurrentValueRequest
	GetDepIds() *string
	SetGroupIds(v string) *GetIndexCurrentValueRequest
	GetGroupIds() *string
	SetInstanceId(v string) *GetIndexCurrentValueRequest
	GetInstanceId() *string
}

type GetIndexCurrentValueRequest struct {
	// example:
	//
	// 2332****,2334****
	DepIds *string `json:"DepIds,omitempty" xml:"DepIds,omitempty"`
	// example:
	//
	// 2323****,2324****
	GroupIds *string `json:"GroupIds,omitempty" xml:"GroupIds,omitempty"`
	// example:
	//
	// ccc_xp_pre-cn-***
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetIndexCurrentValueRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIndexCurrentValueRequest) GoString() string {
	return s.String()
}

func (s *GetIndexCurrentValueRequest) GetDepIds() *string {
	return s.DepIds
}

func (s *GetIndexCurrentValueRequest) GetGroupIds() *string {
	return s.GroupIds
}

func (s *GetIndexCurrentValueRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetIndexCurrentValueRequest) SetDepIds(v string) *GetIndexCurrentValueRequest {
	s.DepIds = &v
	return s
}

func (s *GetIndexCurrentValueRequest) SetGroupIds(v string) *GetIndexCurrentValueRequest {
	s.GroupIds = &v
	return s
}

func (s *GetIndexCurrentValueRequest) SetInstanceId(v string) *GetIndexCurrentValueRequest {
	s.InstanceId = &v
	return s
}

func (s *GetIndexCurrentValueRequest) Validate() error {
	return dara.Validate(s)
}
