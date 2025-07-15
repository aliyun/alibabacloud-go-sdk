// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAttachPluginRequest interface {
	dara.Model
	String() string
	GoString() string
	SetApiId(v string) *AttachPluginRequest
	GetApiId() *string
	SetApiIds(v string) *AttachPluginRequest
	GetApiIds() *string
	SetGroupId(v string) *AttachPluginRequest
	GetGroupId() *string
	SetPluginId(v string) *AttachPluginRequest
	GetPluginId() *string
	SetSecurityToken(v string) *AttachPluginRequest
	GetSecurityToken() *string
	SetStageName(v string) *AttachPluginRequest
	GetStageName() *string
}

type AttachPluginRequest struct {
	// The number of the API to be bound.
	//
	// example:
	//
	// 8afff6c8c4c6447abb035812e4d66b65
	ApiId *string `json:"ApiId,omitempty" xml:"ApiId,omitempty"`
	// The number of the API to be operated. Separate multiple numbers with commas (,). A maximum of 100 numbers can be entered.
	//
	// example:
	//
	// xxx
	ApiIds *string `json:"ApiIds,omitempty" xml:"ApiIds,omitempty"`
	// The ID of the API group that contains the API to which the plug-in is to be bound.
	//
	// example:
	//
	// 285bb759342649a1b70c2093a772e087
	GroupId *string `json:"GroupId,omitempty" xml:"GroupId,omitempty"`
	// The ID of the plug-in to be bound.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9a3f1a5279434f2ba74ccd91c295af9f
	PluginId      *string `json:"PluginId,omitempty" xml:"PluginId,omitempty"`
	SecurityToken *string `json:"SecurityToken,omitempty" xml:"SecurityToken,omitempty"`
	// The name of the runtime environment. Valid values:
	//
	// 	- **RELEASE**
	//
	// 	- **PRE: the pre-release environment**
	//
	// 	- **TEST**
	//
	// This parameter is required.
	//
	// example:
	//
	// TEST
	StageName *string `json:"StageName,omitempty" xml:"StageName,omitempty"`
}

func (s AttachPluginRequest) String() string {
	return dara.Prettify(s)
}

func (s AttachPluginRequest) GoString() string {
	return s.String()
}

func (s *AttachPluginRequest) GetApiId() *string {
	return s.ApiId
}

func (s *AttachPluginRequest) GetApiIds() *string {
	return s.ApiIds
}

func (s *AttachPluginRequest) GetGroupId() *string {
	return s.GroupId
}

func (s *AttachPluginRequest) GetPluginId() *string {
	return s.PluginId
}

func (s *AttachPluginRequest) GetSecurityToken() *string {
	return s.SecurityToken
}

func (s *AttachPluginRequest) GetStageName() *string {
	return s.StageName
}

func (s *AttachPluginRequest) SetApiId(v string) *AttachPluginRequest {
	s.ApiId = &v
	return s
}

func (s *AttachPluginRequest) SetApiIds(v string) *AttachPluginRequest {
	s.ApiIds = &v
	return s
}

func (s *AttachPluginRequest) SetGroupId(v string) *AttachPluginRequest {
	s.GroupId = &v
	return s
}

func (s *AttachPluginRequest) SetPluginId(v string) *AttachPluginRequest {
	s.PluginId = &v
	return s
}

func (s *AttachPluginRequest) SetSecurityToken(v string) *AttachPluginRequest {
	s.SecurityToken = &v
	return s
}

func (s *AttachPluginRequest) SetStageName(v string) *AttachPluginRequest {
	s.StageName = &v
	return s
}

func (s *AttachPluginRequest) Validate() error {
	return dara.Validate(s)
}
