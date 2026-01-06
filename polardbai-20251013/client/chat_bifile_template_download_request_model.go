// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iChatBIFileTemplateDownloadRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAuthMessage(v string) *ChatBIFileTemplateDownloadRequest
	GetAuthMessage() *string
	SetAuthType(v string) *ChatBIFileTemplateDownloadRequest
	GetAuthType() *string
	SetInstanceName(v string) *ChatBIFileTemplateDownloadRequest
	GetInstanceName() *string
	SetTableType(v string) *ChatBIFileTemplateDownloadRequest
	GetTableType() *string
}

type ChatBIFileTemplateDownloadRequest struct {
	AuthMessage *string `json:"AuthMessage,omitempty" xml:"AuthMessage,omitempty"`
	AuthType    *string `json:"AuthType,omitempty" xml:"AuthType,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pc-2ze454l20me07****
	InstanceName *string `json:"InstanceName,omitempty" xml:"InstanceName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// pattern/config
	TableType *string `json:"TableType,omitempty" xml:"TableType,omitempty"`
}

func (s ChatBIFileTemplateDownloadRequest) String() string {
	return dara.Prettify(s)
}

func (s ChatBIFileTemplateDownloadRequest) GoString() string {
	return s.String()
}

func (s *ChatBIFileTemplateDownloadRequest) GetAuthMessage() *string {
	return s.AuthMessage
}

func (s *ChatBIFileTemplateDownloadRequest) GetAuthType() *string {
	return s.AuthType
}

func (s *ChatBIFileTemplateDownloadRequest) GetInstanceName() *string {
	return s.InstanceName
}

func (s *ChatBIFileTemplateDownloadRequest) GetTableType() *string {
	return s.TableType
}

func (s *ChatBIFileTemplateDownloadRequest) SetAuthMessage(v string) *ChatBIFileTemplateDownloadRequest {
	s.AuthMessage = &v
	return s
}

func (s *ChatBIFileTemplateDownloadRequest) SetAuthType(v string) *ChatBIFileTemplateDownloadRequest {
	s.AuthType = &v
	return s
}

func (s *ChatBIFileTemplateDownloadRequest) SetInstanceName(v string) *ChatBIFileTemplateDownloadRequest {
	s.InstanceName = &v
	return s
}

func (s *ChatBIFileTemplateDownloadRequest) SetTableType(v string) *ChatBIFileTemplateDownloadRequest {
	s.TableType = &v
	return s
}

func (s *ChatBIFileTemplateDownloadRequest) Validate() error {
	return dara.Validate(s)
}
