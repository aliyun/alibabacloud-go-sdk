// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddTranscodeTemplateGroupRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *AddTranscodeTemplateGroupRequest
	GetAppId() *string
	SetName(v string) *AddTranscodeTemplateGroupRequest
	GetName() *string
	SetTranscodeTemplateGroupId(v string) *AddTranscodeTemplateGroupRequest
	GetTranscodeTemplateGroupId() *string
	SetTranscodeTemplateList(v string) *AddTranscodeTemplateGroupRequest
	GetTranscodeTemplateList() *string
}

type AddTranscodeTemplateGroupRequest struct {
	// The application ID. Default value: **app-1000000**. For more information, see [Multi-application](https://help.aliyun.com/document_detail/113600.html).
	//
	// example:
	//
	// app-****
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The name of the transcoding template group.
	//
	// - The name can be up to 128 bytes in length.
	//
	// - The value is encoded in UTF-8.
	//
	// > You must specify either TranscodeTemplateGroupId or Name.
	//
	// example:
	//
	// transcodetemplate
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the transcoding template group. If you specify the ID of a transcoding template group, new transcoding templates are added to the specified template group.
	//
	// > You must specify either TranscodeTemplateGroupId or Name.
	//
	// example:
	//
	// 4c71a339fe52b4fa6f4527****
	TranscodeTemplateGroupId *string `json:"TranscodeTemplateGroupId,omitempty" xml:"TranscodeTemplateGroupId,omitempty"`
	// The transcoding template configurations in the JSON format. For more information about the parameter structure, see [TranscodeTemplate](https://help.aliyun.com/document_detail/52839.html).
	//
	// >- If you do not specify this parameter, the transcoding process is not created and video uploads do not trigger transcoding.
	//
	// >- If you do not need to set the Width or Height property, do not specify the corresponding property. Do not set it to an empty string, such as "Height":"".
	//
	// example:
	//
	// [{"Video":{"Bitrate":"400","Codec":"H.264","Fps":"30","Height":360,"Width":640},"Definition":"SD","Container":{"Format":"mp4"},"TemplateName":"testName","MuxConfig":{},"Audio":{"Codec":"AAC","Bitrate":"64","Samplerate":"44100"}}]
	TranscodeTemplateList *string `json:"TranscodeTemplateList,omitempty" xml:"TranscodeTemplateList,omitempty"`
}

func (s AddTranscodeTemplateGroupRequest) String() string {
	return dara.Prettify(s)
}

func (s AddTranscodeTemplateGroupRequest) GoString() string {
	return s.String()
}

func (s *AddTranscodeTemplateGroupRequest) GetAppId() *string {
	return s.AppId
}

func (s *AddTranscodeTemplateGroupRequest) GetName() *string {
	return s.Name
}

func (s *AddTranscodeTemplateGroupRequest) GetTranscodeTemplateGroupId() *string {
	return s.TranscodeTemplateGroupId
}

func (s *AddTranscodeTemplateGroupRequest) GetTranscodeTemplateList() *string {
	return s.TranscodeTemplateList
}

func (s *AddTranscodeTemplateGroupRequest) SetAppId(v string) *AddTranscodeTemplateGroupRequest {
	s.AppId = &v
	return s
}

func (s *AddTranscodeTemplateGroupRequest) SetName(v string) *AddTranscodeTemplateGroupRequest {
	s.Name = &v
	return s
}

func (s *AddTranscodeTemplateGroupRequest) SetTranscodeTemplateGroupId(v string) *AddTranscodeTemplateGroupRequest {
	s.TranscodeTemplateGroupId = &v
	return s
}

func (s *AddTranscodeTemplateGroupRequest) SetTranscodeTemplateList(v string) *AddTranscodeTemplateGroupRequest {
	s.TranscodeTemplateList = &v
	return s
}

func (s *AddTranscodeTemplateGroupRequest) Validate() error {
	return dara.Validate(s)
}
