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
	// Cache preview flag:
	//
	// - true: When enabled, the document preview will no longer update collaborative editing content, suitable for scenarios where only preview is needed.
	//
	// - false: When disabled, it defaults to collaborative preview, allowing the preview to synchronously update collaborative editing content.
	//
	// 	Notice: The price for cache preview and non-cache preview differs. Please refer to the billing item description for more details.</notice> 	Notice: Search and print functions are not supported during cache preview.</notice> <notice>Updating cached content is currently not supported in cache preview mode.</notice>
	//
	// example:
	//
	// true、false
	CachePreview *bool `json:"CachePreview,omitempty" xml:"CachePreview,omitempty"`
	// **If there are no special requirements, leave this blank.**
	//
	// Chained authorization configuration, not required. For more information, see [Using Chained Authorization to Access Other Entity Resources](https://help.aliyun.com/document_detail/465340.html).
	CredentialConfigShrink *string `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
	// Indicates whether uploading a file with the same name to OSS is an expected behavior. Possible values are as follows:
	//
	// - true: Uploading a file with the same name to OSS is an expected behavior. The uploaded document will overwrite the original document and generate a new version. After setting it to true, you still need to close the currently editing document and wait for about 5 minutes before reopening it to load the new document. The upload is only effective when the document is closed; if the document is open, the new save will overwrite the uploaded file.
	//
	// - false (default): Uploading a file with the same name to OSS is not an expected behavior, and the interface will return an error.
	//
	// example:
	//
	// false
	ExternalUploaded *bool `json:"ExternalUploaded,omitempty" xml:"ExternalUploaded,omitempty"`
	// Filename, which must include the file extension. By default, it is the last segment of the **SourceURI*	- parameter.
	//
	// Supported file extensions (PDF is only supported for preview):
	//
	// - Text documents (Word): doc, docx, txt, dot, wps, wpt, dotx, docm, dotm, rtf
	//
	// - Presentation documents (PPT): ppt, pptx, pptm, ppsx, ppsm, pps, potx, potm, dpt, dps - Spreadsheet documents (Excel): et, xls, xlt, xlsx, xlsm, xltx, xltm, csv
	//
	// - PDF documents: pdf
	//
	// example:
	//
	// test-Object.pptx
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// Whether to hide the toolbar. This parameter can be set in document preview mode. Possible values are as follows:
	//
	// - false (default): Do not hide the toolbar.
	//
	// - true: Hide the toolbar.
	//
	// example:
	//
	// false
	Hidecmb *bool `json:"Hidecmb,omitempty" xml:"Hidecmb,omitempty"`
	// Notification message configuration, currently supporting only MNS. For the asynchronous notification message format, refer to [WebOffice Message Notification Format](https://help.aliyun.com/document_detail/2743999.html).
	//
	// > There will be message notifications when the file is saved or renamed.
	NotificationShrink *string `json:"Notification,omitempty" xml:"Notification,omitempty"`
	// Supports notifying some events to customers via MNS messages. This parameter is the topic for MNS asynchronous message notifications.
	//
	// example:
	//
	// test-topic
	NotifyTopicName *string `json:"NotifyTopicName,omitempty" xml:"NotifyTopicName,omitempty"`
	// The password to open the document.
	//
	// > If you need to preview or edit a password-protected document, set this parameter.
	//
	// example:
	//
	// 123456
	Password *string `json:"Password,omitempty" xml:"Password,omitempty"`
	// User permission information, represented in JSON format.
	//
	// User permissions include the following options:
	//
	// Each option is of type Boolean, with a default value of false, and can be set to true or false.
	//
	// - Readonly (optional): Preview mode.
	//
	// - Rename (optional): File renaming permission, which only provides message notification functionality. The renaming event will be sent to MNS.
	//
	// - History (optional): Permission to view historical versions.
	//
	// - Copy (optional): Copy permission.
	//
	// - Export (optional): PDF export permission.
	//
	// - Print (optional): Print permission.
	//
	// >PDF only supports preview functionality, so the "Readonly" parameter must be set to true.
	//
	// >
	//
	// >PDF files do not support exporting.
	//
	// >
	//
	// >To use the multi-version feature, you must first enable the multi-version feature in OSS and then set the "History" parameter to true.
	//
	// >
	//
	// 	Notice: Printing is not supported in cached preview.
	//
	// 	Notice: Historical versions can be viewed in edit mode but not in preview mode.
	PermissionShrink *string `json:"Permission,omitempty" xml:"Permission,omitempty"`
	// Limits the number of pages that can be previewed. By default, there is no limit. The maximum cannot exceed 5000.
	//
	// example:
	//
	// 5
	PreviewPages *int64 `json:"PreviewPages,omitempty" xml:"PreviewPages,omitempty"`
	// Project name, for how to obtain it, please refer to [Create Project](https://help.aliyun.com/document_detail/478153.html).
	//
	// This parameter is required.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// OSS anti-leeching. IMM needs to obtain the source file from OSS. If OSS has set up anti-leeching, IMM must pass the corresponding header to OSS to get the source file.
	//
	// > If the Bucket where the document is located has Referer set, please configure this parameter.
	//
	// example:
	//
	// *
	Referer *string `json:"Referer,omitempty" xml:"Referer,omitempty"`
	// OSS address of the document to be previewed or edited. The OSS address follows the rule `oss://${Bucket}/${Object}`, where `Bucket` is the name of the OSS Bucket in the same region as the current project, and `Object` is the full path of the file including the file extension.
	//
	// This parameter is required.
	//
	// example:
	//
	// oss://test-bucket/test-object.docx
	SourceURI *string `json:"SourceURI,omitempty" xml:"SourceURI,omitempty"`
	// User information. You can pass in user information from the business side, which will be displayed on the WebOffice page.
	//
	// The system distinguishes different users by User.Id, and User.Name is used only for front-end display. If User.Id is not provided, the backend will generate a random ID. Users with different IDs are considered different entities and cannot modify or delete each other\\"s comments.
	//
	// The default format is: Unknown_random string. If User.Id is not provided, the user information will default to "Unknown".
	UserShrink *string `json:"User,omitempty" xml:"User,omitempty"`
	// User-defined information. It only takes effect when Notification parameters are filled in for MNS configuration. It will be returned in asynchronous message notifications, which can help you correlate and process messages within your system. The maximum length is 2048 bytes.
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
	// Watermark information. The watermark is generated on the front end and is not written into the source document. The same document with different parameters will result in different watermarks.
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
