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
	// Specifies whether to enable cached preview.
	//
	// -  true: When enabled, the document preview no longer updates collaborative editing content. This is suitable for preview-only scenarios.
	//
	// -  false: When disabled, collaborative preview is used by default, which synchronizes collaborative editing content during preview.
	//
	// 	Notice: Cached preview and non-cached preview have different unit prices. For more information, see the billing item description.
	//
	// </notice>	Notice: Cached preview does not support document content search or printing.</notice>
	//
	// <notice>Cached preview does not support updating cached content.</notice>.
	//
	// example:
	//
	// true
	CachePreview *bool `json:"CachePreview,omitempty" xml:"CachePreview,omitempty"`
	// **Leave this parameter empty unless you have specific requirements.**
	//
	// The China authorization configuration. This parameter is optional. For more information, see [Use chained authorization to access resources of other entities](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// Specifies whether uploading a file with the same name to OSS is expected behavior. Valid values:
	//
	// - true: Uploading a file with the same name to OSS is expected behavior. The uploaded document overwrites the original document and generates a new version. After this parameter is set to true, you must first close the document that is being edited, wait about 5 minutes, and then reopen it to load the new document. The upload takes effect only when the document is closed. If the document is open, new saves overwrite the uploaded file.
	//
	// - false (default): Uploading a file with the same name to OSS is not expected behavior. The operation returns an error.
	//
	// example:
	//
	// false
	ExternalUploaded *bool `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	// The file name, which must include the file name extension. The default value is the last segment of the **SourceURI*	- parameter.
	//
	// Supported file name extensions (PDF supports preview only):
	//
	// - Word documents: doc, docx, txt, dot, wps, wpt, dotx, docm, dotm, and rtf
	//
	// - PowerPoint documents: ppt, pptx, pptm, ppsx, ppsm, pps, potx, potm, dpt, and dps
	//
	// - Excel documents: et, xls, xlt, xlsx, xlsm, xltx, xltm, and csv
	//
	// - PDF documents: pdf.
	//
	// example:
	//
	// test-Object.pptx
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Specifies whether to hide the toolbar. This parameter is supported in document preview mode. Valid values:
	//
	// - false (default): The toolbar is not hidden.
	//
	// - true: The toolbar is hidden.
	//
	// example:
	//
	// false
	Hidecmb *bool `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	// The notification configuration. Currently, only MNS is supported. For the format of asynchronous notification messages, see [WebOffice message notification format](https://help.aliyun.com/document_detail/2743999.html).
	//
	// > Message notifications are sent when a file is saved or renamed.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// Sends event notifications to you as MNS messages. This parameter specifies the MNS topic for asynchronous message notifications.
	//
	// example:
	//
	// test-topic
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// The password to open the document.
	//
	// > Set this parameter if you want to preview or edit a password-protected document.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// The user permission information in JSON format.
	//
	// User permissions include the following options:
	//
	// Each option is of the Boolean type. The default value is false. Valid values: true and false.
	//
	// - Readonly (optional): Preview mode.
	//
	// - Rename (optional): The permission to rename a file. Only message notification is provided. The rename event is sent to MNS.
	//
	// - History (optional): The permission to view historical versions.
	//
	// - Copy (optional): The copy permission.
	//
	// - Export (optional): The permission to export to PDF.
	//
	// - Print (optional): The print permission.
	//
	//
	// > PDF supports only the preview feature. You must set the Readonly parameter to true.
	//
	// >
	//
	// > PDF files do not support export.
	//
	// >
	//
	// > To use the versioning feature, you must first enable versioning in OSS and then set the History parameter to true.
	//
	// >
	//
	// 	Notice: Printing is not supported in cached preview.
	//
	// 	Notice: Historical versions can be viewed in edit mode but not in preview mode..
	PermissionShrink *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// The maximum number of pages that can be previewed. By default, no limit is imposed. The maximum value is 5,000.
	//
	// example:
	//
	// 5
	PreviewPages *int64 `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	// The project name. For information about how to obtain the project name, see [Create a project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The OSS hotlink protection referer. Intelligent Media Management (IMM) needs to retrieve the source file from OSS. If hotlink protection is configured for OSS, IMM must pass the corresponding header to OSS to retrieve the source file.
	//
	// > Set this parameter if the bucket that stores the document has a referer configured.
	//
	// example:
	//
	// *
	Referer *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	// The OSS URI of the document to preview or edit.
	//
	// The OSS URI follows the format `oss://${Bucket}/${Object}`, where `Bucket` is the name of an OSS bucket in the same region as the current project, and `Object` is the full path of the file including the file name extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object.docx
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// The user information. You can pass in user information from the business side, and the WebOffice page displays this information.
	//
	// The system distinguishes different users by User.Id. User.Name is used only for frontend display. If User.Id is not specified, the backend automatically generates a random ID. Users with different IDs are treated as different principals and cannot modify or delete each other\\"s comments.
	//
	// The default format is: Unknown_RandomString. If User.Id is not specified, the user information is displayed as "Unknown" by default.
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
	// The custom user data. This parameter takes effect only when the Notification parameter is specified with MNS configurations. The data is returned in asynchronous message notifications for you to associate and process message notifications within your system. Maximum length: 2,048 bytes.
	//
	// example:
	//
	// {
	//
	//       "id": "test-id",
	//
	//       "name": "test-name"
	//
	// }
	UserData *string `json:"UserData,omitempty" xml:"UserData,omitempty"`
	// The watermark information. The watermark is generated on the frontend and is not written to the source document. Different parameters passed for the same document produce different watermarks.
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
