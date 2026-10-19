// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetAccountAttributesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *SetAccountAttributesRequest
	GetClientToken() *string
	SetLogEnabled(v bool) *SetAccountAttributesRequest
	GetLogEnabled() *bool
	SetLogStoreName(v string) *SetAccountAttributesRequest
	GetLogStoreName() *string
	SetMessageTraceEnabled(v bool) *SetAccountAttributesRequest
	GetMessageTraceEnabled() *bool
	SetProjectName(v string) *SetAccountAttributesRequest
	GetProjectName() *string
}

type SetAccountAttributesRequest struct {
	// Ensures the idempotence of the request. Generate a parameter value from your client to ensure that the value is unique across different requests. ClientToken supports only ASCII characters and cannot exceed 64 characters in length. For more information, see How to ensure idempotence.
	//
	// example:
	//
	// 06273500-249
	ClientToken *string `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	// Specifies whether to enable the log management feature.
	//
	// example:
	//
	// true
	LogEnabled *bool `json:"LogEnabled,omitempty" xml:"LogEnabled,omitempty"`
	// The name of the Simple Log Service Logstore.
	//
	// example:
	//
	// test-logstore
	LogStoreName *string `json:"LogStoreName,omitempty" xml:"LogStoreName,omitempty"`
	// Specifies whether to enable the message trace feature.
	//
	// example:
	//
	// true
	MessageTraceEnabled *bool `json:"MessageTraceEnabled,omitempty" xml:"MessageTraceEnabled,omitempty"`
	// The name of the Simple Log Service project.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
}

func (s SetAccountAttributesRequest) String() string {
	return dara.Prettify(s)
}

func (s SetAccountAttributesRequest) GoString() string {
	return s.String()
}

func (s *SetAccountAttributesRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *SetAccountAttributesRequest) GetLogEnabled() *bool {
	return s.LogEnabled
}

func (s *SetAccountAttributesRequest) GetLogStoreName() *string {
	return s.LogStoreName
}

func (s *SetAccountAttributesRequest) GetMessageTraceEnabled() *bool {
	return s.MessageTraceEnabled
}

func (s *SetAccountAttributesRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *SetAccountAttributesRequest) SetClientToken(v string) *SetAccountAttributesRequest {
	s.ClientToken = &v
	return s
}

func (s *SetAccountAttributesRequest) SetLogEnabled(v bool) *SetAccountAttributesRequest {
	s.LogEnabled = &v
	return s
}

func (s *SetAccountAttributesRequest) SetLogStoreName(v string) *SetAccountAttributesRequest {
	s.LogStoreName = &v
	return s
}

func (s *SetAccountAttributesRequest) SetMessageTraceEnabled(v bool) *SetAccountAttributesRequest {
	s.MessageTraceEnabled = &v
	return s
}

func (s *SetAccountAttributesRequest) SetProjectName(v string) *SetAccountAttributesRequest {
	s.ProjectName = &v
	return s
}

func (s *SetAccountAttributesRequest) Validate() error {
	return dara.Validate(s)
}
