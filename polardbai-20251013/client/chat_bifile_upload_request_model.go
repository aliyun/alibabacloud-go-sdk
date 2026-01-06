// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIFileUploadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIFileUploadRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIFileUploadRequest
	GetAuthType() *string
	SetFileName(v string) *ChatBIFileUploadRequest
	GetFileName() *string
	SetInstanceName(v string) *ChatBIFileUploadRequest
	GetInstanceName() *string
}

type ChatBIFileUploadRequest struct {
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	AuthType    *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// example:
	//
	// pattern_test.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
}

func (s ChatBIFileUploadRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIFileUploadRequest) GoString() string {
	return s.String()
}

func (s *ChatBIFileUploadRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIFileUploadRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIFileUploadRequest) GetFileName() *string {
	return s.FileName
}

func (s *ChatBIFileUploadRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIFileUploadRequest) SetAuthMessage(v string) *ChatBIFileUploadRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIFileUploadRequest) SetAuthType(v string) *ChatBIFileUploadRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIFileUploadRequest) SetFileName(v string) *ChatBIFileUploadRequest {
	s.FileName = &v
	return s
}

func (s *ChatBIFileUploadRequest) SetInstanceName(v string) *ChatBIFileUploadRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIFileUploadRequest) Validate() error {
	return dara.Validate(s)
}
