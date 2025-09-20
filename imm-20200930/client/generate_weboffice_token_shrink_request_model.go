// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGenerateWebofficeTokenShrinkRequest interface {
	dara.Model
	String() string
	GoString() string
	SetCachePreview(v bool) *GenerateWebofficeTokenShrinkRequest
	GetCachePreview() *bool
	SetCredentialConfigShrink(v string) *GenerateWebofficeTokenShrinkRequest
	GetCredentialConfigShrink() *string
	SetExternalUploaded(v bool) *GenerateWebofficeTokenShrinkRequest
	GetExternalUploaded() *bool
	SetFilename(v string) *GenerateWebofficeTokenShrinkRequest
	GetFilename() *string
	SetHidecmb(v bool) *GenerateWebofficeTokenShrinkRequest
	GetHidecmb() *bool
	SetNotificationShrink(v string) *GenerateWebofficeTokenShrinkRequest
	GetNotificationShrink() *string
	SetNotifyTopicName(v string) *GenerateWebofficeTokenShrinkRequest
	GetNotifyTopicName() *string
	SetPassword(v string) *GenerateWebofficeTokenShrinkRequest
	GetPassword() *string
	SetPermissionShrink(v string) *GenerateWebofficeTokenShrinkRequest
	GetPermissionShrink() *string
	SetPreviewPages(v int64) *GenerateWebofficeTokenShrinkRequest
	GetPreviewPages() *int64
	SetProjectName(v string) *GenerateWebofficeTokenShrinkRequest
	GetProjectName() *string
	SetReferer(v string) *GenerateWebofficeTokenShrinkRequest
	GetReferer() *string
	SetSourceURI(v string) *GenerateWebofficeTokenShrinkRequest
	GetSourceURI() *string
	SetUserShrink(v string) *GenerateWebofficeTokenShrinkRequest
	GetUserShrink() *string
	SetUserData(v string) *GenerateWebofficeTokenShrinkRequest
	GetUserData() *string
	SetWatermarkShrink(v string) *GenerateWebofficeTokenShrinkRequest
	GetWatermarkShrink() *string
}

type GenerateWebofficeTokenShrinkRequest struct {
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
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	PermissionShrink *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
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
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
	// The user-defined data that you want to return in asynchronous messages. This parameter takes effect only when you specify the MNS settings in the Notification parameter. The maximum length of the value is 2,048 bytes.
	//
	// example:
	//
	// {"file_id": "abc"}
	UserData        *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	WatermarkShrink *string `json:"Watermark,omitempty" xml:"Watermark,omitempty"`
}

func (s GenerateWebofficeTokenShrinkRequest) String() string {
	return dara.Prettify(s)
}

func (s GenerateWebofficeTokenShrinkRequest) GoString() string {
	return s.String()
}

func (s *GenerateWebofficeTokenShrinkRequest) GetCachePreview() *bool {
	return s.CachePreview
}

func (s *GenerateWebofficeTokenShrinkRequest) GetCredentialConfigShrink() *string {
	return s.CredentialConfigShrink
}

func (s *GenerateWebofficeTokenShrinkRequest) GetExternalUploaded() *bool {
	return s.ExternalUploaded
}

func (s *GenerateWebofficeTokenShrinkRequest) GetFilename() *string {
	return s.Filename
}

func (s *GenerateWebofficeTokenShrinkRequest) GetHidecmb() *bool {
	return s.Hidecmb
}

func (s *GenerateWebofficeTokenShrinkRequest) GetNotificationShrink() *string {
	return s.NotificationShrink
}

func (s *GenerateWebofficeTokenShrinkRequest) GetNotifyTopicName() *string {
	return s.NotifyTopicName
}

func (s *GenerateWebofficeTokenShrinkRequest) GetPassword() *string {
	return s.Password
}

func (s *GenerateWebofficeTokenShrinkRequest) GetPermissionShrink() *string {
	return s.PermissionShrink
}

func (s *GenerateWebofficeTokenShrinkRequest) GetPreviewPages() *int64 {
	return s.PreviewPages
}

func (s *GenerateWebofficeTokenShrinkRequest) GetProjectName() *string {
	return s.ProjectName
}

func (s *GenerateWebofficeTokenShrinkRequest) GetReferer() *string {
	return s.Referer
}

func (s *GenerateWebofficeTokenShrinkRequest) GetSourceURI() *string {
	return s.SourceURI
}

func (s *GenerateWebofficeTokenShrinkRequest) GetUserShrink() *string {
	return s.UserShrink
}

func (s *GenerateWebofficeTokenShrinkRequest) GetUserData() *string {
	return s.UserData
}

func (s *GenerateWebofficeTokenShrinkRequest) GetWatermarkShrink() *string {
	return s.WatermarkShrink
}

func (s *GenerateWebofficeTokenShrinkRequest) SetCachePreview(v bool) *GenerateWebofficeTokenShrinkRequest {
	s.CachePreview = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetCredentialConfigShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.CredentialConfigShrink = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetExternalUploaded(v bool) *GenerateWebofficeTokenShrinkRequest {
	s.ExternalUploaded = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetFilename(v string) *GenerateWebofficeTokenShrinkRequest {
	s.Filename = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetHidecmb(v bool) *GenerateWebofficeTokenShrinkRequest {
	s.Hidecmb = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetNotificationShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.NotificationShrink = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetNotifyTopicName(v string) *GenerateWebofficeTokenShrinkRequest {
	s.NotifyTopicName = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetPassword(v string) *GenerateWebofficeTokenShrinkRequest {
	s.Password = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetPermissionShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.PermissionShrink = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetPreviewPages(v int64) *GenerateWebofficeTokenShrinkRequest {
	s.PreviewPages = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetProjectName(v string) *GenerateWebofficeTokenShrinkRequest {
	s.ProjectName = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetReferer(v string) *GenerateWebofficeTokenShrinkRequest {
	s.Referer = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetSourceURI(v string) *GenerateWebofficeTokenShrinkRequest {
	s.SourceURI = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetUserShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.UserShrink = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetUserData(v string) *GenerateWebofficeTokenShrinkRequest {
	s.UserData = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) SetWatermarkShrink(v string) *GenerateWebofficeTokenShrinkRequest {
	s.WatermarkShrink = &v
	return s
}

func (s *GenerateWebofficeTokenShrinkRequest) Validate() error {
	return dara.Validate(s)
}
