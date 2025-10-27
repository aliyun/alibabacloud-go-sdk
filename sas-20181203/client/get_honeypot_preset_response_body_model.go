// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetHoneypotPresetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetHoneypotPresetResponseBody
	GetCode() *string
	SetData(v *GetHoneypotPresetResponseBodyData) *GetHoneypotPresetResponseBody
	GetData() *GetHoneypotPresetResponseBodyData
	SetHttpStatusCode(v int32) *GetHoneypotPresetResponseBody
	GetHttpStatusCode() *int32
	SetMessage(v string) *GetHoneypotPresetResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetHoneypotPresetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetHoneypotPresetResponseBody
	GetSuccess() *bool
}

type GetHoneypotPresetResponseBody struct {
	// The status code returned. The status code **200*	- indicates that the request was successful. Other status codes indicate that the request failed. You can identify the cause of the failure based on the status code.
	//
	// example:
	//
	// 200
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The information about the honeypot template.
	Data *GetHoneypotPresetResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The HTTP status code returned.
	//
	// example:
	//
	// 200
	HttpStatusCode *int32 `json:"HttpStatusCode,omitempty" xml:"HttpStatusCode,omitempty"`
	// The message returned.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The ID of the request, which is used to locate and troubleshoot issues.
	//
	// example:
	//
	// 38AFE393-88E8-5642-B3E2-D57C6E76025D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful. Valid values:
	//
	// 	- **true**: The request was successful.
	//
	// 	- **false**: The request failed.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetHoneypotPresetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotPresetResponseBody) GoString() string {
	return s.String()
}

func (s *GetHoneypotPresetResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetHoneypotPresetResponseBody) GetData() *GetHoneypotPresetResponseBodyData {
	return s.Data
}

func (s *GetHoneypotPresetResponseBody) GetHttpStatusCode() *int32 {
	return s.HttpStatusCode
}

func (s *GetHoneypotPresetResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetHoneypotPresetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetHoneypotPresetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetHoneypotPresetResponseBody) SetCode(v string) *GetHoneypotPresetResponseBody {
	s.Code = &v
	return s
}

func (s *GetHoneypotPresetResponseBody) SetData(v *GetHoneypotPresetResponseBodyData) *GetHoneypotPresetResponseBody {
	s.Data = v
	return s
}

func (s *GetHoneypotPresetResponseBody) SetHttpStatusCode(v int32) *GetHoneypotPresetResponseBody {
	s.HttpStatusCode = &v
	return s
}

func (s *GetHoneypotPresetResponseBody) SetMessage(v string) *GetHoneypotPresetResponseBody {
	s.Message = &v
	return s
}

func (s *GetHoneypotPresetResponseBody) SetRequestId(v string) *GetHoneypotPresetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetHoneypotPresetResponseBody) SetSuccess(v bool) *GetHoneypotPresetResponseBody {
	s.Success = &v
	return s
}

func (s *GetHoneypotPresetResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetHoneypotPresetResponseBodyData struct {
	// The name of the management node.
	//
	// example:
	//
	// managerNodename
	ControlNodeName *string `json:"ControlNodeName,omitempty" xml:"ControlNodeName,omitempty"`
	// An array that consists of the configurations of the uploaded file.
	FileInfoList []*GetHoneypotPresetResponseBodyDataFileInfoList `json:"FileInfoList,omitempty" xml:"FileInfoList,omitempty" type:"Repeated"`
	// The display name of the honeypot image.
	//
	// example:
	//
	// RuoYi
	HoneypotImageDisplayName *string `json:"HoneypotImageDisplayName,omitempty" xml:"HoneypotImageDisplayName,omitempty"`
	// The name of the honeypot image.
	//
	// example:
	//
	// ruoyi
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The ID of the honeypot template.
	//
	// example:
	//
	// 94fd8805-d178-4361-84d3-de47fb4e****
	HoneypotPresetId *string `json:"HoneypotPresetId,omitempty" xml:"HoneypotPresetId,omitempty"`
	// The custom configuration of the honeypot template.
	//
	// example:
	//
	// {"trojan_git":"zip","burp":"open","portrait_option":"true"}
	Meta *string `json:"Meta,omitempty" xml:"Meta,omitempty"`
	// The ID of the management node.
	//
	// example:
	//
	// a882e590-b87b-45a6-87b9-d0a3e5a0****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The custom name of the honeypot template.
	//
	// example:
	//
	// ssh
	PresetName *string `json:"PresetName,omitempty" xml:"PresetName,omitempty"`
	// The type of the honeypot template. Valid values:
	//
	// 	- **TEMP**: automatically generated template
	//
	// 	- **CUSTOM**: custom template
	//
	// 	- **DEFAULT**: default template
	//
	// example:
	//
	// CUSTOM
	PresetType *string `json:"PresetType,omitempty" xml:"PresetType,omitempty"`
}

func (s GetHoneypotPresetResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotPresetResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetHoneypotPresetResponseBodyData) GetControlNodeName() *string {
	return s.ControlNodeName
}

func (s *GetHoneypotPresetResponseBodyData) GetFileInfoList() []*GetHoneypotPresetResponseBodyDataFileInfoList {
	return s.FileInfoList
}

func (s *GetHoneypotPresetResponseBodyData) GetHoneypotImageDisplayName() *string {
	return s.HoneypotImageDisplayName
}

func (s *GetHoneypotPresetResponseBodyData) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *GetHoneypotPresetResponseBodyData) GetHoneypotPresetId() *string {
	return s.HoneypotPresetId
}

func (s *GetHoneypotPresetResponseBodyData) GetMeta() *string {
	return s.Meta
}

func (s *GetHoneypotPresetResponseBodyData) GetNodeId() *string {
	return s.NodeId
}

func (s *GetHoneypotPresetResponseBodyData) GetPresetName() *string {
	return s.PresetName
}

func (s *GetHoneypotPresetResponseBodyData) GetPresetType() *string {
	return s.PresetType
}

func (s *GetHoneypotPresetResponseBodyData) SetControlNodeName(v string) *GetHoneypotPresetResponseBodyData {
	s.ControlNodeName = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyData) SetFileInfoList(v []*GetHoneypotPresetResponseBodyDataFileInfoList) *GetHoneypotPresetResponseBodyData {
	s.FileInfoList = v
	return s
}

func (s *GetHoneypotPresetResponseBodyData) SetHoneypotImageDisplayName(v string) *GetHoneypotPresetResponseBodyData {
	s.HoneypotImageDisplayName = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyData) SetHoneypotImageName(v string) *GetHoneypotPresetResponseBodyData {
	s.HoneypotImageName = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyData) SetHoneypotPresetId(v string) *GetHoneypotPresetResponseBodyData {
	s.HoneypotPresetId = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyData) SetMeta(v string) *GetHoneypotPresetResponseBodyData {
	s.Meta = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyData) SetNodeId(v string) *GetHoneypotPresetResponseBodyData {
	s.NodeId = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyData) SetPresetName(v string) *GetHoneypotPresetResponseBodyData {
	s.PresetName = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyData) SetPresetType(v string) *GetHoneypotPresetResponseBodyData {
	s.PresetType = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyData) Validate() error {
	if s.FileInfoList != nil {
		for _, item := range s.FileInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetHoneypotPresetResponseBodyDataFileInfoList struct {
	// The ID of the uploaded file.
	//
	// example:
	//
	// HONEYPOT_FILE/1765_167040128****
	FileId *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	// The name of the uploaded file.
	//
	// example:
	//
	// HONEYPOT_FILE****
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The download URL.
	//
	// example:
	//
	// http://aegis****
	OssUrl *string `json:"OssUrl,omitempty" xml:"OssUrl,omitempty"`
}

func (s GetHoneypotPresetResponseBodyDataFileInfoList) String() string {
	return dara.Prettify(s)
}

func (s GetHoneypotPresetResponseBodyDataFileInfoList) GoString() string {
	return s.String()
}

func (s *GetHoneypotPresetResponseBodyDataFileInfoList) GetFileId() *string {
	return s.FileId
}

func (s *GetHoneypotPresetResponseBodyDataFileInfoList) GetFileName() *string {
	return s.FileName
}

func (s *GetHoneypotPresetResponseBodyDataFileInfoList) GetOssUrl() *string {
	return s.OssUrl
}

func (s *GetHoneypotPresetResponseBodyDataFileInfoList) SetFileId(v string) *GetHoneypotPresetResponseBodyDataFileInfoList {
	s.FileId = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyDataFileInfoList) SetFileName(v string) *GetHoneypotPresetResponseBodyDataFileInfoList {
	s.FileName = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyDataFileInfoList) SetOssUrl(v string) *GetHoneypotPresetResponseBodyDataFileInfoList {
	s.OssUrl = &v
	return s
}

func (s *GetHoneypotPresetResponseBodyDataFileInfoList) Validate() error {
	return dara.Validate(s)
}
