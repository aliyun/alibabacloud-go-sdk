// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTagsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ListTagsRequest
	GetInstanceId() *string
	SetScriptId(v string) *ListTagsRequest
	GetScriptId() *string
}

type ListTagsRequest struct {
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 8fa1953f-4a84-46d8-b80c-8ce9cf684fb3
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Script ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 8a4c6d3d-5ed6-44ca-b779-16c20f8862be
	ScriptId *string `json:"ScriptId,omitempty" xml:"ScriptId,omitempty"`
}

func (s ListTagsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ListTagsRequest) GetScriptId() *string {
	return s.ScriptId
}

func (s *ListTagsRequest) SetInstanceId(v string) *ListTagsRequest {
	s.InstanceId = &v
	return s
}

func (s *ListTagsRequest) SetScriptId(v string) *ListTagsRequest {
	s.ScriptId = &v
	return s
}

func (s *ListTagsRequest) Validate() error {
	return dara.Validate(s)
}
