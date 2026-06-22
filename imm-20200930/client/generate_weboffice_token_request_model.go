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
	CredentialConfig *CredentialConfig `json:"CredentialConfig,omitempty" xml:"CredentialConfig,omitempty"`
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
	Notification *Notification `json:"Notification,omitempty" xml:"Notification,omitempty"`
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
	Permission *WebofficePermission `json:"Permission,omitempty" xml:"Permission,omitempty"`
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
	User *WebofficeUser `json:"User,omitempty" xml:"User,omitempty"`
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
	if s.CredentialConfig != nil {
		if err := s.CredentialConfig.Validate(); err != nil {
			return err
		}
	}
	if s.Notification != nil {
		if err := s.Notification.Validate(); err != nil {
			return err
		}
	}
	if s.Permission != nil {
		if err := s.Permission.Validate(); err != nil {
			return err
		}
	}
	if s.User != nil {
		if err := s.User.Validate(); err != nil {
			return err
		}
	}
	if s.Watermark != nil {
		if err := s.Watermark.Validate(); err != nil {
			return err
		}
	}
	return nil
}
