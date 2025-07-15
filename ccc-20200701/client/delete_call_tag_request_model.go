// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteCallTagRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *DeleteCallTagRequest
	GetInstanceId() *string
	SetTagName(v string) *DeleteCallTagRequest
	GetTagName() *string
}

type DeleteCallTagRequest struct {
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
	// TagA
	TagName *string `json:"TagName,omitempty" xml:"TagName,omitempty"`
}

func (s DeleteCallTagRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteCallTagRequest) GoString() string {
	return s.String()
}

func (s *DeleteCallTagRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *DeleteCallTagRequest) GetTagName() *string {
	return s.TagName
}

func (s *DeleteCallTagRequest) SetInstanceId(v string) *DeleteCallTagRequest {
	s.InstanceId = &v
	return s
}

func (s *DeleteCallTagRequest) SetTagName(v string) *DeleteCallTagRequest {
	s.TagName = &v
	return s
}

func (s *DeleteCallTagRequest) Validate() error {
	return dara.Validate(s)
}
