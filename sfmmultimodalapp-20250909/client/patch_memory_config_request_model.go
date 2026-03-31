// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPatchMemoryConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *PatchMemoryConfigRequest
	GetAppId() *string
	SetAutoUpdate(v bool) *PatchMemoryConfigRequest
	GetAutoUpdate() *bool
	SetExpirationTime(v int32) *PatchMemoryConfigRequest
	GetExpirationTime() *int32
	SetPrompt(v string) *PatchMemoryConfigRequest
	GetPrompt() *string
	SetThreshold(v float64) *PatchMemoryConfigRequest
	GetThreshold() *float64
	SetTopK(v int32) *PatchMemoryConfigRequest
	GetTopK() *int32
	SetUserDefinedId(v string) *PatchMemoryConfigRequest
	GetUserDefinedId() *string
	SetWorkspaceId(v string) *PatchMemoryConfigRequest
	GetWorkspaceId() *string
}

type PatchMemoryConfigRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// mm_bfaf7e110b6d4359977d1686a3f8
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// example:
	//
	// true
	AutoUpdate *bool `json:"AutoUpdate,omitempty" xml:"AutoUpdate,omitempty"`
	// example:
	//
	// 30
	ExpirationTime *int32  `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	Prompt         *string `json:"Prompt,omitempty" xml:"Prompt,omitempty"`
	// example:
	//
	// 0.03
	Threshold *float64 `json:"Threshold,omitempty" xml:"Threshold,omitempty"`
	// example:
	//
	// 3
	TopK *int32 `json:"TopK,omitempty" xml:"TopK,omitempty"`
	// example:
	//
	// 110b6d4359977d1
	UserDefinedId *string `json:"UserDefinedId,omitempty" xml:"UserDefinedId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// llm-jb5sabg80b4ts71g
	WorkspaceId *string `json:"WorkspaceId,omitempty" xml:"WorkspaceId,omitempty"`
}

func (s PatchMemoryConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s PatchMemoryConfigRequest) GoString() string {
	return s.String()
}

func (s *PatchMemoryConfigRequest) GetAppId() *string {
	return s.AppId
}

func (s *PatchMemoryConfigRequest) GetAutoUpdate() *bool {
	return s.AutoUpdate
}

func (s *PatchMemoryConfigRequest) GetExpirationTime() *int32 {
	return s.ExpirationTime
}

func (s *PatchMemoryConfigRequest) GetPrompt() *string {
	return s.Prompt
}

func (s *PatchMemoryConfigRequest) GetThreshold() *float64 {
	return s.Threshold
}

func (s *PatchMemoryConfigRequest) GetTopK() *int32 {
	return s.TopK
}

func (s *PatchMemoryConfigRequest) GetUserDefinedId() *string {
	return s.UserDefinedId
}

func (s *PatchMemoryConfigRequest) GetWorkspaceId() *string {
	return s.WorkspaceId
}

func (s *PatchMemoryConfigRequest) SetAppId(v string) *PatchMemoryConfigRequest {
	s.AppId = &v
	return s
}

func (s *PatchMemoryConfigRequest) SetAutoUpdate(v bool) *PatchMemoryConfigRequest {
	s.AutoUpdate = &v
	return s
}

func (s *PatchMemoryConfigRequest) SetExpirationTime(v int32) *PatchMemoryConfigRequest {
	s.ExpirationTime = &v
	return s
}

func (s *PatchMemoryConfigRequest) SetPrompt(v string) *PatchMemoryConfigRequest {
	s.Prompt = &v
	return s
}

func (s *PatchMemoryConfigRequest) SetThreshold(v float64) *PatchMemoryConfigRequest {
	s.Threshold = &v
	return s
}

func (s *PatchMemoryConfigRequest) SetTopK(v int32) *PatchMemoryConfigRequest {
	s.TopK = &v
	return s
}

func (s *PatchMemoryConfigRequest) SetUserDefinedId(v string) *PatchMemoryConfigRequest {
	s.UserDefinedId = &v
	return s
}

func (s *PatchMemoryConfigRequest) SetWorkspaceId(v string) *PatchMemoryConfigRequest {
	s.WorkspaceId = &v
	return s
}

func (s *PatchMemoryConfigRequest) Validate() error {
	return dara.Validate(s)
}
