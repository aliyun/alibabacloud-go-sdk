// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPreviewOutputFileResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *PreviewOutputFileResponseBody
	GetCode() *string
	SetContent(v string) *PreviewOutputFileResponseBody
	GetContent() *string
	SetDegraded(v bool) *PreviewOutputFileResponseBody
	GetDegraded() *bool
	SetDownloadExpiresAt(v string) *PreviewOutputFileResponseBody
	GetDownloadExpiresAt() *string
	SetDownloadUrl(v string) *PreviewOutputFileResponseBody
	GetDownloadUrl() *string
	SetFileExt(v string) *PreviewOutputFileResponseBody
	GetFileExt() *string
	SetGmtModified(v string) *PreviewOutputFileResponseBody
	GetGmtModified() *string
	SetItemName(v string) *PreviewOutputFileResponseBody
	GetItemName() *string
	SetItemType(v string) *PreviewOutputFileResponseBody
	GetItemType() *string
	SetMessage(v string) *PreviewOutputFileResponseBody
	GetMessage() *string
	SetOutputItemId(v string) *PreviewOutputFileResponseBody
	GetOutputItemId() *string
	SetPreviewExpiresAt(v string) *PreviewOutputFileResponseBody
	GetPreviewExpiresAt() *string
	SetPreviewType(v string) *PreviewOutputFileResponseBody
	GetPreviewType() *string
	SetPreviewUrl(v string) *PreviewOutputFileResponseBody
	GetPreviewUrl() *string
	SetReadState(v string) *PreviewOutputFileResponseBody
	GetReadState() *string
	SetRequestId(v string) *PreviewOutputFileResponseBody
	GetRequestId() *string
	SetSlides(v []*PreviewOutputFileResponseBodySlides) *PreviewOutputFileResponseBody
	GetSlides() []*PreviewOutputFileResponseBodySlides
	SetUnavailableReason(v string) *PreviewOutputFileResponseBody
	GetUnavailableReason() *string
	SetVersion(v *PreviewOutputFileResponseBodyVersion) *PreviewOutputFileResponseBody
	GetVersion() *PreviewOutputFileResponseBodyVersion
}

type PreviewOutputFileResponseBody struct {
	// The error code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The text content. This parameter is used for the CONTENT type.
	//
	// example:
	//
	// # Quarterly Review
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	// Indicates whether the preview is degraded to a download. A value of true indicates that the preview is degraded to a download and cannot be opened in an online previewer.
	//
	// example:
	//
	// false
	Degraded *bool `json:"degraded,omitempty" xml:"degraded,omitempty"`
	// The expiration time of the download ticket.
	//
	// example:
	//
	// 2026-11-16T14:00:00+00:00
	DownloadExpiresAt *string `json:"downloadExpiresAt,omitempty" xml:"downloadExpiresAt,omitempty"`
	// The short-term signed download URL for the output. You can send a GET request to this URL before it expires.
	//
	// example:
	//
	// https://example.com/download/file
	DownloadUrl *string `json:"downloadUrl,omitempty" xml:"downloadUrl,omitempty"`
	// The file name extension. This parameter is used for the OSS_IMM type.
	//
	// example:
	//
	// pptx
	FileExt *string `json:"fileExt,omitempty" xml:"fileExt,omitempty"`
	// The modification time in ISO8601 format.
	//
	// example:
	//
	// 2026-11-16T13:00:00+00:00
	GmtModified *string `json:"gmtModified,omitempty" xml:"gmtModified,omitempty"`
	// The output name.
	//
	// example:
	//
	// Quarterly_Review.pptx
	ItemName *string `json:"itemName,omitempty" xml:"itemName,omitempty"`
	// The output item type. Valid values: ppt, html, document, picture, slides, video, audio, email, and others.
	//
	// example:
	//
	// ppt
	ItemType *string `json:"itemType,omitempty" xml:"itemType,omitempty"`
	// The message.
	//
	// example:
	//
	// Invalid request parameters
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The output item ID.
	//
	// example:
	//
	// exampleOutputItemId
	OutputItemId *string `json:"outputItemId,omitempty" xml:"outputItemId,omitempty"`
	// The expiration time of the preview ticket.
	//
	// example:
	//
	// 2026-11-16T14:00:00+00:00
	PreviewExpiresAt *string `json:"previewExpiresAt,omitempty" xml:"previewExpiresAt,omitempty"`
	// The preview type. Valid values: OSS_IMM, IMAGE, AUDIO, VIDEO, HTML, DING_TALK, VOICE_MEETING, and CONTENT.
	//
	// example:
	//
	// OSS_IMM
	PreviewType *string `json:"previewType,omitempty" xml:"previewType,omitempty"`
	// The preview URL. This parameter is used for OSS_IMM, DING_TALK, and VOICE_MEETING.
	//
	// example:
	//
	// https://example.com/preview/document
	PreviewUrl *string `json:"previewUrl,omitempty" xml:"previewUrl,omitempty"`
	// The read state. Valid values: AVAILABLE, PRUNED, PENDING, and UNVERSIONED.
	//
	// example:
	//
	// AVAILABLE
	ReadState *string `json:"readState,omitempty" xml:"readState,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The list of slides for the slide deck output.
	//
	// example:
	//
	// []
	Slides []*PreviewOutputFileResponseBodySlides `json:"slides,omitempty" xml:"slides,omitempty" type:"Repeated"`
	// The reason why the content is unavailable.
	//
	// example:
	//
	// The output content is still being generated
	UnavailableReason *string `json:"unavailableReason,omitempty" xml:"unavailableReason,omitempty"`
	// The optimistic version number.
	Version *PreviewOutputFileResponseBodyVersion `json:"version,omitempty" xml:"version,omitempty" type:"Struct"`
}

func (s PreviewOutputFileResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PreviewOutputFileResponseBody) GoString() string {
	return s.String()
}

func (s *PreviewOutputFileResponseBody) GetCode() *string {
	return s.Code
}

func (s *PreviewOutputFileResponseBody) GetContent() *string {
	return s.Content
}

func (s *PreviewOutputFileResponseBody) GetDegraded() *bool {
	return s.Degraded
}

func (s *PreviewOutputFileResponseBody) GetDownloadExpiresAt() *string {
	return s.DownloadExpiresAt
}

func (s *PreviewOutputFileResponseBody) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *PreviewOutputFileResponseBody) GetFileExt() *string {
	return s.FileExt
}

func (s *PreviewOutputFileResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *PreviewOutputFileResponseBody) GetItemName() *string {
	return s.ItemName
}

func (s *PreviewOutputFileResponseBody) GetItemType() *string {
	return s.ItemType
}

func (s *PreviewOutputFileResponseBody) GetMessage() *string {
	return s.Message
}

func (s *PreviewOutputFileResponseBody) GetOutputItemId() *string {
	return s.OutputItemId
}

func (s *PreviewOutputFileResponseBody) GetPreviewExpiresAt() *string {
	return s.PreviewExpiresAt
}

func (s *PreviewOutputFileResponseBody) GetPreviewType() *string {
	return s.PreviewType
}

func (s *PreviewOutputFileResponseBody) GetPreviewUrl() *string {
	return s.PreviewUrl
}

func (s *PreviewOutputFileResponseBody) GetReadState() *string {
	return s.ReadState
}

func (s *PreviewOutputFileResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PreviewOutputFileResponseBody) GetSlides() []*PreviewOutputFileResponseBodySlides {
	return s.Slides
}

func (s *PreviewOutputFileResponseBody) GetUnavailableReason() *string {
	return s.UnavailableReason
}

func (s *PreviewOutputFileResponseBody) GetVersion() *PreviewOutputFileResponseBodyVersion {
	return s.Version
}

func (s *PreviewOutputFileResponseBody) SetCode(v string) *PreviewOutputFileResponseBody {
	s.Code = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetContent(v string) *PreviewOutputFileResponseBody {
	s.Content = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetDegraded(v bool) *PreviewOutputFileResponseBody {
	s.Degraded = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetDownloadExpiresAt(v string) *PreviewOutputFileResponseBody {
	s.DownloadExpiresAt = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetDownloadUrl(v string) *PreviewOutputFileResponseBody {
	s.DownloadUrl = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetFileExt(v string) *PreviewOutputFileResponseBody {
	s.FileExt = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetGmtModified(v string) *PreviewOutputFileResponseBody {
	s.GmtModified = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetItemName(v string) *PreviewOutputFileResponseBody {
	s.ItemName = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetItemType(v string) *PreviewOutputFileResponseBody {
	s.ItemType = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetMessage(v string) *PreviewOutputFileResponseBody {
	s.Message = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetOutputItemId(v string) *PreviewOutputFileResponseBody {
	s.OutputItemId = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetPreviewExpiresAt(v string) *PreviewOutputFileResponseBody {
	s.PreviewExpiresAt = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetPreviewType(v string) *PreviewOutputFileResponseBody {
	s.PreviewType = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetPreviewUrl(v string) *PreviewOutputFileResponseBody {
	s.PreviewUrl = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetReadState(v string) *PreviewOutputFileResponseBody {
	s.ReadState = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetRequestId(v string) *PreviewOutputFileResponseBody {
	s.RequestId = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetSlides(v []*PreviewOutputFileResponseBodySlides) *PreviewOutputFileResponseBody {
	s.Slides = v
	return s
}

func (s *PreviewOutputFileResponseBody) SetUnavailableReason(v string) *PreviewOutputFileResponseBody {
	s.UnavailableReason = &v
	return s
}

func (s *PreviewOutputFileResponseBody) SetVersion(v *PreviewOutputFileResponseBodyVersion) *PreviewOutputFileResponseBody {
	s.Version = v
	return s
}

func (s *PreviewOutputFileResponseBody) Validate() error {
	if s.Slides != nil {
		for _, item := range s.Slides {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Version != nil {
		if err := s.Version.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewOutputFileResponseBodySlides struct {
	// The display text of the slide.
	//
	// example:
	//
	// Market Overview
	DisplayText *string `json:"displayText,omitempty" xml:"displayText,omitempty"`
	// Specifies the detail level of the image. Valid values: low, high, and auto. Default value: auto.
	//
	// example:
	//
	// https://example.com/slide/1
	ImageUrl *string `json:"imageUrl,omitempty" xml:"imageUrl,omitempty"`
	// The slide ID.
	//
	// example:
	//
	// exampleSlideId
	SlideId *string `json:"slideId,omitempty" xml:"slideId,omitempty"`
	// The slide order.
	//
	// example:
	//
	// 1
	SlideIndex *int64 `json:"slideIndex,omitempty" xml:"slideIndex,omitempty"`
	// The signing status. Valid values:
	//
	// - CREATED: created but not signed
	//
	// - SUCCESS: signed successfully
	//
	// - STOP: terminated
	//
	// example:
	//
	// SUCCESS
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s PreviewOutputFileResponseBodySlides) String() string {
	return dara.Prettify(s)
}

func (s PreviewOutputFileResponseBodySlides) GoString() string {
	return s.String()
}

func (s *PreviewOutputFileResponseBodySlides) GetDisplayText() *string {
	return s.DisplayText
}

func (s *PreviewOutputFileResponseBodySlides) GetImageUrl() *string {
	return s.ImageUrl
}

func (s *PreviewOutputFileResponseBodySlides) GetSlideId() *string {
	return s.SlideId
}

func (s *PreviewOutputFileResponseBodySlides) GetSlideIndex() *int64 {
	return s.SlideIndex
}

func (s *PreviewOutputFileResponseBodySlides) GetStatus() *string {
	return s.Status
}

func (s *PreviewOutputFileResponseBodySlides) SetDisplayText(v string) *PreviewOutputFileResponseBodySlides {
	s.DisplayText = &v
	return s
}

func (s *PreviewOutputFileResponseBodySlides) SetImageUrl(v string) *PreviewOutputFileResponseBodySlides {
	s.ImageUrl = &v
	return s
}

func (s *PreviewOutputFileResponseBodySlides) SetSlideId(v string) *PreviewOutputFileResponseBodySlides {
	s.SlideId = &v
	return s
}

func (s *PreviewOutputFileResponseBodySlides) SetSlideIndex(v int64) *PreviewOutputFileResponseBodySlides {
	s.SlideIndex = &v
	return s
}

func (s *PreviewOutputFileResponseBodySlides) SetStatus(v string) *PreviewOutputFileResponseBodySlides {
	s.Status = &v
	return s
}

func (s *PreviewOutputFileResponseBodySlides) Validate() error {
	return dara.Validate(s)
}

type PreviewOutputFileResponseBodyVersion struct {
	// The number of remaining seats.
	//
	// example:
	//
	// AVAILABLE
	Availability *string `json:"availability,omitempty" xml:"availability,omitempty"`
	// Specifies whether the model supports multiple tool calls in a single response.
	Capabilities *PreviewOutputFileResponseBodyVersionCapabilities `json:"capabilities,omitempty" xml:"capabilities,omitempty" type:"Struct"`
	// The content revision number.
	//
	// example:
	//
	// 1
	ContentRevision *int64 `json:"contentRevision,omitempty" xml:"contentRevision,omitempty"`
	// Indicates whether the version is the latest version.
	//
	// example:
	//
	// true
	IsLatest *bool `json:"isLatest,omitempty" xml:"isLatest,omitempty"`
	// The ID of the current latest version.
	//
	// example:
	//
	// exampleLatestVersionId
	LatestVersionId *string `json:"latestVersionId,omitempty" xml:"latestVersionId,omitempty"`
	// The read state of the version.
	//
	// example:
	//
	// AVAILABLE
	ReadState *string `json:"readState,omitempty" xml:"readState,omitempty"`
	// The ID of the version that is actually read.
	//
	// example:
	//
	// exampleVersionId
	SelectedVersionId *string `json:"selectedVersionId,omitempty" xml:"selectedVersionId,omitempty"`
	// The ID of the message that generated this version.
	//
	// example:
	//
	// exampleMessageId
	SourceMessageId *string `json:"sourceMessageId,omitempty" xml:"sourceMessageId,omitempty"`
	// The update time.
	//
	// example:
	//
	// 2026-11-16T13:00:00+00:00
	UpdatedAt *string `json:"updatedAt,omitempty" xml:"updatedAt,omitempty"`
	// The version serial number.
	//
	// example:
	//
	// 2
	VersionNo *int64 `json:"versionNo,omitempty" xml:"versionNo,omitempty"`
}

func (s PreviewOutputFileResponseBodyVersion) String() string {
	return dara.Prettify(s)
}

func (s PreviewOutputFileResponseBodyVersion) GoString() string {
	return s.String()
}

func (s *PreviewOutputFileResponseBodyVersion) GetAvailability() *string {
	return s.Availability
}

func (s *PreviewOutputFileResponseBodyVersion) GetCapabilities() *PreviewOutputFileResponseBodyVersionCapabilities {
	return s.Capabilities
}

func (s *PreviewOutputFileResponseBodyVersion) GetContentRevision() *int64 {
	return s.ContentRevision
}

func (s *PreviewOutputFileResponseBodyVersion) GetIsLatest() *bool {
	return s.IsLatest
}

func (s *PreviewOutputFileResponseBodyVersion) GetLatestVersionId() *string {
	return s.LatestVersionId
}

func (s *PreviewOutputFileResponseBodyVersion) GetReadState() *string {
	return s.ReadState
}

func (s *PreviewOutputFileResponseBodyVersion) GetSelectedVersionId() *string {
	return s.SelectedVersionId
}

func (s *PreviewOutputFileResponseBodyVersion) GetSourceMessageId() *string {
	return s.SourceMessageId
}

func (s *PreviewOutputFileResponseBodyVersion) GetUpdatedAt() *string {
	return s.UpdatedAt
}

func (s *PreviewOutputFileResponseBodyVersion) GetVersionNo() *int64 {
	return s.VersionNo
}

func (s *PreviewOutputFileResponseBodyVersion) SetAvailability(v string) *PreviewOutputFileResponseBodyVersion {
	s.Availability = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) SetCapabilities(v *PreviewOutputFileResponseBodyVersionCapabilities) *PreviewOutputFileResponseBodyVersion {
	s.Capabilities = v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) SetContentRevision(v int64) *PreviewOutputFileResponseBodyVersion {
	s.ContentRevision = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) SetIsLatest(v bool) *PreviewOutputFileResponseBodyVersion {
	s.IsLatest = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) SetLatestVersionId(v string) *PreviewOutputFileResponseBodyVersion {
	s.LatestVersionId = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) SetReadState(v string) *PreviewOutputFileResponseBodyVersion {
	s.ReadState = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) SetSelectedVersionId(v string) *PreviewOutputFileResponseBodyVersion {
	s.SelectedVersionId = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) SetSourceMessageId(v string) *PreviewOutputFileResponseBodyVersion {
	s.SourceMessageId = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) SetUpdatedAt(v string) *PreviewOutputFileResponseBodyVersion {
	s.UpdatedAt = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) SetVersionNo(v int64) *PreviewOutputFileResponseBodyVersion {
	s.VersionNo = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersion) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PreviewOutputFileResponseBodyVersionCapabilities struct {
	// Indicates whether the version can be opened in a browser.
	//
	// example:
	//
	// true
	CanBrowserOpen *bool `json:"canBrowserOpen,omitempty" xml:"canBrowserOpen,omitempty"`
	// Indicates whether the version can be downloaded.
	//
	// example:
	//
	// true
	CanDownload *bool `json:"canDownload,omitempty" xml:"canDownload,omitempty"`
	// Indicates whether the course can be edited or deleted.
	//
	// example:
	//
	// false
	CanEdit *bool `json:"canEdit,omitempty" xml:"canEdit,omitempty"`
	// Indicates whether the version can be restored to the latest version.
	//
	// example:
	//
	// false
	CanRestore *bool `json:"canRestore,omitempty" xml:"canRestore,omitempty"`
	// Indicates whether the version can be saved as knowledge.
	//
	// example:
	//
	// true
	CanSaveAsKnowledge *bool `json:"canSaveAsKnowledge,omitempty" xml:"canSaveAsKnowledge,omitempty"`
	// Indicates whether the version can be shared.
	//
	// example:
	//
	// true
	CanShare *bool `json:"canShare,omitempty" xml:"canShare,omitempty"`
}

func (s PreviewOutputFileResponseBodyVersionCapabilities) String() string {
	return dara.Prettify(s)
}

func (s PreviewOutputFileResponseBodyVersionCapabilities) GoString() string {
	return s.String()
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) GetCanBrowserOpen() *bool {
	return s.CanBrowserOpen
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) GetCanDownload() *bool {
	return s.CanDownload
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) GetCanEdit() *bool {
	return s.CanEdit
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) GetCanRestore() *bool {
	return s.CanRestore
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) GetCanSaveAsKnowledge() *bool {
	return s.CanSaveAsKnowledge
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) GetCanShare() *bool {
	return s.CanShare
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) SetCanBrowserOpen(v bool) *PreviewOutputFileResponseBodyVersionCapabilities {
	s.CanBrowserOpen = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) SetCanDownload(v bool) *PreviewOutputFileResponseBodyVersionCapabilities {
	s.CanDownload = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) SetCanEdit(v bool) *PreviewOutputFileResponseBodyVersionCapabilities {
	s.CanEdit = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) SetCanRestore(v bool) *PreviewOutputFileResponseBodyVersionCapabilities {
	s.CanRestore = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) SetCanSaveAsKnowledge(v bool) *PreviewOutputFileResponseBodyVersionCapabilities {
	s.CanSaveAsKnowledge = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) SetCanShare(v bool) *PreviewOutputFileResponseBodyVersionCapabilities {
	s.CanShare = &v
	return s
}

func (s *PreviewOutputFileResponseBodyVersionCapabilities) Validate() error {
	return dara.Validate(s)
}
