// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUploadedHoneyPotFileRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileKey(v string) *UploadedHoneyPotFileRequest
	GetFileKey() *string
	SetFileName(v string) *UploadedHoneyPotFileRequest
	GetFileName() *string
	SetFileType(v string) *UploadedHoneyPotFileRequest
	GetFileType() *string
	SetHoneypotImageName(v string) *UploadedHoneyPotFileRequest
	GetHoneypotImageName() *string
	SetLang(v string) *UploadedHoneyPotFileRequest
	GetLang() *string
	SetNodeId(v string) *UploadedHoneyPotFileRequest
	GetNodeId() *string
	SetTemplateExtra(v string) *UploadedHoneyPotFileRequest
	GetTemplateExtra() *string
}

type UploadedHoneyPotFileRequest struct {
	// The FileKey used to upload the file.
	//
	// > Format: HONEYPOT_FILE/{timestamp}_{custom_file_name}.
	//
	// This parameter is required.
	//
	// example:
	//
	// HONEYPOT_FILE/1601097845544644_********
	FileKey *string `json:"FileKey,omitempty" xml:"FileKey,omitempty"`
	// The name of the uploaded file.
	//
	// This parameter is required.
	//
	// example:
	//
	// trojan.zip
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The file type.
	//
	// This parameter is required.
	//
	// example:
	//
	// application/zip
	FileType *string `json:"FileType,omitempty" xml:"FileType,omitempty"`
	// The name of the honeypot image.
	//
	// This parameter is required.
	//
	// example:
	//
	// ruoyi
	HoneypotImageName *string `json:"HoneypotImageName,omitempty" xml:"HoneypotImageName,omitempty"`
	// The language type of the request and response. Default value: **zh**. Valid values:
	//
	// - **zh**: Chinese
	//
	// - **en**: English.
	//
	// example:
	//
	// zh
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The ID of the honeypot management node.
	//
	// > Call the [ListHoneypotNode](~~ListHoneypotNode~~) operation to obtain this value.
	//
	// example:
	//
	// cc427e14-f257-4670-9d2b-d83bbbe*****
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The template prompt corresponding to the uploaded file.
	//
	// This parameter is required.
	//
	// example:
	//
	// {\\"help\\":\\".zip\\",\\"label\\":\\"file\\",\\"type\\":\\"file\\",\\"key\\":\\"ftpfiles.zip\\"}
	TemplateExtra *string `json:"TemplateExtra,omitempty" xml:"TemplateExtra,omitempty"`
}

func (s UploadedHoneyPotFileRequest) String() string {
	return dara.Prettify(s)
}

func (s UploadedHoneyPotFileRequest) GoString() string {
	return s.String()
}

func (s *UploadedHoneyPotFileRequest) GetFileKey() *string {
	return s.FileKey
}

func (s *UploadedHoneyPotFileRequest) GetFileName() *string {
	return s.FileName
}

func (s *UploadedHoneyPotFileRequest) GetFileType() *string {
	return s.FileType
}

func (s *UploadedHoneyPotFileRequest) GetHoneypotImageName() *string {
	return s.HoneypotImageName
}

func (s *UploadedHoneyPotFileRequest) GetLang() *string {
	return s.Lang
}

func (s *UploadedHoneyPotFileRequest) GetNodeId() *string {
	return s.NodeId
}

func (s *UploadedHoneyPotFileRequest) GetTemplateExtra() *string {
	return s.TemplateExtra
}

func (s *UploadedHoneyPotFileRequest) SetFileKey(v string) *UploadedHoneyPotFileRequest {
	s.FileKey = &v
	return s
}

func (s *UploadedHoneyPotFileRequest) SetFileName(v string) *UploadedHoneyPotFileRequest {
	s.FileName = &v
	return s
}

func (s *UploadedHoneyPotFileRequest) SetFileType(v string) *UploadedHoneyPotFileRequest {
	s.FileType = &v
	return s
}

func (s *UploadedHoneyPotFileRequest) SetHoneypotImageName(v string) *UploadedHoneyPotFileRequest {
	s.HoneypotImageName = &v
	return s
}

func (s *UploadedHoneyPotFileRequest) SetLang(v string) *UploadedHoneyPotFileRequest {
	s.Lang = &v
	return s
}

func (s *UploadedHoneyPotFileRequest) SetNodeId(v string) *UploadedHoneyPotFileRequest {
	s.NodeId = &v
	return s
}

func (s *UploadedHoneyPotFileRequest) SetTemplateExtra(v string) *UploadedHoneyPotFileRequest {
	s.TemplateExtra = &v
	return s
}

func (s *UploadedHoneyPotFileRequest) Validate() error {
	return dara.Validate(s)
}
