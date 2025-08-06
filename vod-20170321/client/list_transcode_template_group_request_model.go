// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListTranscodeTemplateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *ListTranscodeTemplateGroupRequest
	GetAppId() *string
}

type ListTranscodeTemplateGroupRequest struct {
	// The ID of the application. Default value: **app-1000000**. For more information, see [Overview](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
}

func (s ListTranscodeTemplateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s ListTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *ListTranscodeTemplateGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *ListTranscodeTemplateGroupRequest) SetAppId(v string) *ListTranscodeTemplateGroupRequest {
	s.AppId = &v
	return s
}

func (s *ListTranscodeTemplateGroupRequest) Validate() error {
	return dara.Validate(s)
}
