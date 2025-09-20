// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebofficeTokenRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCachePreview(v bool) *GenerateWebofficeTokenRequest
	GetCachePreview() *bool
	SetCredentialConfig(v *CredentialConfig) *GenerateWebofficeTokenRequest
	GetCredentialConfig() *CredentialConfig
	SetExternalUploaded(v bool) *GenerateWebofficeTokenRequest
	GetExternalUploaded() *bool
	SetFilename(v string) *GenerateWebofficeTokenRequest
	GetFilename() *string
	SetHidecmb(v bool) *GenerateWebofficeTokenRequest
	GetHidecmb() *bool
	SetNotification(v *Notification) *GenerateWebofficeTokenRequest
	GetNotification() *Notification
	SetNotifyTopicName(v string) *GenerateWebofficeTokenRequest
	GetNotifyTopicName() *string
	SetPassword(v string) *GenerateWebofficeTokenRequest
	GetPassword() *string
	SetPermission(v *WebofficePermission) *GenerateWebofficeTokenRequest
	GetPermission() *WebofficePermission
	SetPreviewPages(v int64) *GenerateWebofficeTokenRequest
	GetPreviewPages() *int64
	SetProjectName(v string) *GenerateWebofficeTokenRequest
	GetProjectName() *string
	SetReferer(v string) *GenerateWebofficeTokenRequest
	GetReferer() *string
	SetSourceURI(v string) *GenerateWebofficeTokenRequest
	GetSourceURI() *string
	SetUser(v *WebofficeUser) *GenerateWebofficeTokenRequest
	GetUser() *WebofficeUser
	SetUserData(v string) *GenerateWebofficeTokenRequest
	GetUserData() *string
	SetWatermark(v *WebofficeWatermark) *GenerateWebofficeTokenRequest
	GetWatermark() *WebofficeWatermark
}

type GenerateWebofficeTokenRequest struct {
	// Specifies whether to enable cache preview.
	//
	// 	- true: enables cache preview. The document can be previewed only and cannot be collaboratively edited.
	//
	// 	- false: does not enable cache preview. The document can be collaboratively edited when it is being previewed.
	//
	// >  The pricing for document previews varies based on whether cache preview is enabled or disabled.
	//
	// >  During a cache preview, document content search and printing are not supported.
	//
	// example:
	//
	// false
	CachePreview *bool `json:"CachePreview,omitempty" xml:"CachePreview,omitempty"`
	// **If you have no special requirements, leave this parameter empty.**
	//
	// The configurations of authorization chains. For more information, see [Use authorization chains to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// Specifies whether to allow an upload of a document to the Object Storage Service (OSS) bucket. Valid values:
	//
	// 	- true: Documents can be directly uploaded to OSS. The uploaded document overwrites the existing document and a new version is generated for the document. Before you upload a new document, close the existing document if it is being edited. After the document is uploaded, wait for approximately 5 minutes before you open the document again so that the new version can successfully load. Upload a new document only when the existing is closed. Otherwise, the uploaded document is overwritten when the existing document is saved.
	//
	// 	- false: Documents cannot be directly uploaded to OSS. If you try to upload a document, an error is returned. This is the default value.
	//
	// example:
	//
	// false
	ExternalUploaded *bool `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	// The name of the file. The extension must be included in the file name. By default, this parameter is set to the last depth level of the **SourceURI*	- parameter value.
	//
	// Supported extensions (only preview supported for .pdf):
	//
	// 	- Word documents: .doc, .docx, .txt, .dot, .wps, .wpt, .dotx, .docm, .dotm, and .rtf
	//
	// 	- Presentation documents: .ppt, .pptx, .pptm, .ppsx, .ppsm, .pps, .potx, .potm, .dpt, and .dps
	//
	// 	- Table documents: .et, .xls, .xlt, .xlsx, .xlsm, .xltx, .xltm, and .csv
	//
	// 	- PDF documents: .pdf
	//
	// example:
	//
	// test.pptx
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// example:
	//
	// false
	Hidecmb *bool `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	// The notification settings. Only SMQ messages are supported. For more information, see [WebOffice message example](https://help.aliyun.com/document_detail/2743999.html).
	//
	// >  A notification is sent after the document is saved or renamed.
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// example:
	//
	// topic1
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The user permission settings in the JSON format.
	//
	// The parameter supports the following permission fields:
	//
	// Each field is of type Boolean and can have a value of true and false (the default value):
	//
	// 	- Readonly: grants the permission to preview the document. This field is optional.
	//
	// 	- Rename: grants the permission to rename the document. Notification messages of a rename event can be sent only by using SMQ. This field is optional.
	//
	// 	- History: grants the permission to view historical versions. This field is optional.
	//
	// 	- Copy: grants the permission to copy the document. This field is optional.
	//
	// 	- Export: grants the permission to export the document as a PDF file. This field is optional.
	//
	// 	- Print: grants the permission to print the document. This field is optional.
	//
	// >  Only online preview is supported for PDF documents. When you call the operation on a PDF document, you can set Readonly only to true.
	//
	// >  To manage multiple versions of the document, you must enable versioning for the bucket that stores the document and set the History parameter to true.
	//
	// >  Printing is not supported during cache preview.
	Permission *WebofficePermission `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// example:
	//
	// 5
	PreviewPages *int64 `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// immtest
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// example:
	//
	// *
	Referer *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// oss://imm-test/test.pptx
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The user information. The user information that you want to display on the WebOffice page. If you do not specify this parameter, the user name displayed is Unknown.
	User *WebofficeUser `json:"User,omitempty" xml:"User,omitempty"`
	// The user-defined data that you want to return in asynchronous messages. This parameter takes effect only when you specify the MNS settings in the Notification parameter. The maximum length of the value is 2,048 bytes.
	//
	// example:
	//
	// {"file_id": "abc"}
	UserData  *string             `json:"UserData,omitempty" xml:"UserData,omitempty"`
	Watermark *WebofficeWatermark `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s GenerateWebofficeTokenRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebofficeTokenRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebofficeTokenRequest) GetCachePreview() *bool {
	return s.CachePreview
}

func (s *GenerateWebofficeTokenRequest) GetCredentialConfig() *CredentialConfig {
	return s.CredentialConfig
}

func (s *GenerateWebofficeTokenRequest) GetExternalUploaded() *bool {
	return s.ExternalUploaded
}

func (s *GenerateWebofficeTokenRequest) GetFilename() *string {
	return s.Filename
}

func (s *GenerateWebofficeTokenRequest) GetHidecmb() *bool {
	return s.Hidecmb
}

func (s *GenerateWebofficeTokenRequest) GetNotification() *Notification {
	return s.Notification
}

func (s *GenerateWebofficeTokenRequest) GetNotifyTopicName() *string {
	return s.NotifyTopicName
}

func (s *GenerateWebofficeTokenRequest) GetPassword() *string {
	return s.Password
}

func (s *GenerateWebofficeTokenRequest) GetPermission() *WebofficePermission {
	return s.Permission
}

func (s *GenerateWebofficeTokenRequest) GetPreviewPages() *int64 {
	return s.PreviewPages
}

func (s *GenerateWebofficeTokenRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GenerateWebofficeTokenRequest) GetReferer() *string {
	return s.Referer
}

func (s *GenerateWebofficeTokenRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *GenerateWebofficeTokenRequest) GetUser() *WebofficeUser {
	return s.User
}

func (s *GenerateWebofficeTokenRequest) GetUserData() *string {
	return s.UserData
}

func (s *GenerateWebofficeTokenRequest) GetWatermark() *WebofficeWatermark {
	return s.Watermark
}

func (s *GenerateWebofficeTokenRequest) SetCachePreview(v bool) *GenerateWebofficeTokenRequest {
	s.CachePreview = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetCredentialConfig(v *CredentialConfig) *GenerateWebofficeTokenRequest {
	s.CredentialConfig = v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetExternalUploaded(v bool) *GenerateWebofficeTokenRequest {
	s.ExternalUploaded = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetFilename(v string) *GenerateWebofficeTokenRequest {
	s.Filename = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetHidecmb(v bool) *GenerateWebofficeTokenRequest {
	s.Hidecmb = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetNotification(v *Notification) *GenerateWebofficeTokenRequest {
	s.Notification = v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetNotifyTopicName(v string) *GenerateWebofficeTokenRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetPassword(v string) *GenerateWebofficeTokenRequest {
	s.Password = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetPermission(v *WebofficePermission) *GenerateWebofficeTokenRequest {
	s.Permission = v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetPreviewPages(v int64) *GenerateWebofficeTokenRequest {
	s.PreviewPages = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetProjectName(v string) *GenerateWebofficeTokenRequest {
	s.ProjectName = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetReferer(v string) *GenerateWebofficeTokenRequest {
	s.Referer = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetSourceURI(v string) *GenerateWebofficeTokenRequest {
	s.SourceURI = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetUser(v *WebofficeUser) *GenerateWebofficeTokenRequest {
	s.User = v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetUserData(v string) *GenerateWebofficeTokenRequest {
	s.UserData = &v
	return s
}

func (s *GenerateWebofficeTokenRequest) SetWatermark(v *WebofficeWatermark) *GenerateWebofficeTokenRequest {
	s.Watermark = v
	return s
}

func (s *GenerateWebofficeTokenRequest) Validate() error {
	return dara.Validate(s)
}
