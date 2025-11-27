// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFile interface {
	dara.Model
	String() string
	GoString() string
	SetAccessControlAllowOrigin(v string) *File
	GetAccessControlAllowOrigin() *string
	SetAccessControlRequestMethod(v string) *File
	GetAccessControlRequestMethod() *string
	SetAddresses(v []*Address) *File
	GetAddresses() []*Address
	SetAlbum(v string) *File
	GetAlbum() *string
	SetAlbumArtist(v string) *File
	GetAlbumArtist() *string
	SetArtist(v string) *File
	GetArtist() *string
	SetAudioCovers(v []*Image) *File
	GetAudioCovers() []*Image
	SetAudioStreams(v []*AudioStream) *File
	GetAudioStreams() []*AudioStream
	SetBitrate(v int64) *File
	GetBitrate() *int64
	SetCacheControl(v string) *File
	GetCacheControl() *string
	SetComposer(v string) *File
	GetComposer() *string
	SetContentDisposition(v string) *File
	GetContentDisposition() *string
	SetContentEncoding(v string) *File
	GetContentEncoding() *string
	SetContentLanguage(v string) *File
	GetContentLanguage() *string
	SetContentMd5(v string) *File
	GetContentMd5() *string
	SetContentType(v string) *File
	GetContentType() *string
	SetCreateTime(v string) *File
	GetCreateTime() *string
	SetCroppingSuggestions(v []*CroppingSuggestion) *File
	GetCroppingSuggestions() []*CroppingSuggestion
	SetCustomId(v string) *File
	GetCustomId() *string
	SetCustomLabels(v map[string]interface{}) *File
	GetCustomLabels() map[string]interface{}
	SetDatasetName(v string) *File
	GetDatasetName() *string
	SetDuration(v float64) *File
	GetDuration() *float64
	SetETag(v string) *File
	GetETag() *string
	SetEXIF(v string) *File
	GetEXIF() *string
	SetElements(v []*Element) *File
	GetElements() []*Element
	SetFigureCount(v int64) *File
	GetFigureCount() *int64
	SetFigures(v []*Figure) *File
	GetFigures() []*Figure
	SetFileAccessTime(v string) *File
	GetFileAccessTime() *string
	SetFileCreateTime(v string) *File
	GetFileCreateTime() *string
	SetFileHash(v string) *File
	GetFileHash() *string
	SetFileModifiedTime(v string) *File
	GetFileModifiedTime() *string
	SetFilename(v string) *File
	GetFilename() *string
	SetFormatLongName(v string) *File
	GetFormatLongName() *string
	SetFormatName(v string) *File
	GetFormatName() *string
	SetImageHeight(v int64) *File
	GetImageHeight() *int64
	SetImageScore(v *ImageScore) *File
	GetImageScore() *ImageScore
	SetImageWidth(v int64) *File
	GetImageWidth() *int64
	SetInsights(v *Insights) *File
	GetInsights() *Insights
	SetLabels(v []*Label) *File
	GetLabels() []*Label
	SetLanguage(v string) *File
	GetLanguage() *string
	SetLatLong(v string) *File
	GetLatLong() *string
	SetMediaType(v string) *File
	GetMediaType() *string
	SetOCRContents(v []*OCRContents) *File
	GetOCRContents() []*OCRContents
	SetOCRTexts(v string) *File
	GetOCRTexts() *string
	SetOSSCRC64(v string) *File
	GetOSSCRC64() *string
	SetOSSDeleteMarker(v string) *File
	GetOSSDeleteMarker() *string
	SetOSSExpiration(v string) *File
	GetOSSExpiration() *string
	SetOSSObjectType(v string) *File
	GetOSSObjectType() *string
	SetOSSStorageClass(v string) *File
	GetOSSStorageClass() *string
	SetOSSTagging(v map[string]interface{}) *File
	GetOSSTagging() map[string]interface{}
	SetOSSTaggingCount(v int64) *File
	GetOSSTaggingCount() *int64
	SetOSSURI(v string) *File
	GetOSSURI() *string
	SetOSSUserMeta(v map[string]interface{}) *File
	GetOSSUserMeta() map[string]interface{}
	SetOSSVersionId(v string) *File
	GetOSSVersionId() *string
	SetObjectACL(v string) *File
	GetObjectACL() *string
	SetObjectId(v string) *File
	GetObjectId() *string
	SetObjectStatus(v string) *File
	GetObjectStatus() *string
	SetObjectType(v string) *File
	GetObjectType() *string
	SetOrientation(v int64) *File
	GetOrientation() *int64
	SetOwnerId(v string) *File
	GetOwnerId() *string
	SetPageCount(v int64) *File
	GetPageCount() *int64
	SetPerformer(v string) *File
	GetPerformer() *string
	SetProduceTime(v string) *File
	GetProduceTime() *string
	SetProgramCount(v int64) *File
	GetProgramCount() *int64
	SetProjectName(v string) *File
	GetProjectName() *string
	SetReason(v string) *File
	GetReason() *string
	SetSceneElements(v []*SceneElement) *File
	GetSceneElements() []*SceneElement
	SetSemanticTypes(v []*string) *File
	GetSemanticTypes() []*string
	SetServerSideDataEncryption(v string) *File
	GetServerSideDataEncryption() *string
	SetServerSideEncryption(v string) *File
	GetServerSideEncryption() *string
	SetServerSideEncryptionCustomerAlgorithm(v string) *File
	GetServerSideEncryptionCustomerAlgorithm() *string
	SetServerSideEncryptionKeyId(v string) *File
	GetServerSideEncryptionKeyId() *string
	SetSize(v int64) *File
	GetSize() *int64
	SetStartTime(v float64) *File
	GetStartTime() *float64
	SetStreamCount(v int64) *File
	GetStreamCount() *int64
	SetSubtitles(v []*SubtitleStream) *File
	GetSubtitles() []*SubtitleStream
	SetTimezone(v string) *File
	GetTimezone() *string
	SetTitle(v string) *File
	GetTitle() *string
	SetTravelClusterId(v string) *File
	GetTravelClusterId() *string
	SetURI(v string) *File
	GetURI() *string
	SetUpdateTime(v string) *File
	GetUpdateTime() *string
	SetVideoHeight(v int64) *File
	GetVideoHeight() *int64
	SetVideoStreams(v []*VideoStream) *File
	GetVideoStreams() []*VideoStream
	SetVideoWidth(v int64) *File
	GetVideoWidth() *int64
}

type File struct {
	AccessControlAllowOrigin   *string                `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
	AccessControlRequestMethod *string                `json:"AccessControlRequestMethod,omitempty" xml:"AccessControlRequestMethod,omitempty"`
	Addresses                  []*Address             `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	Album                      *string                `json:"Album,omitempty" xml:"Album,omitempty"`
	AlbumArtist                *string                `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	Artist                     *string                `json:"Artist,omitempty" xml:"Artist,omitempty"`
	AudioCovers                []*Image               `json:"AudioCovers,omitempty" xml:"AudioCovers,omitempty" type:"Repeated"`
	AudioStreams               []*AudioStream         `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Repeated"`
	Bitrate                    *int64                 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	CacheControl               *string                `json:"CacheControl,omitempty" xml:"CacheControl,omitempty"`
	Composer                   *string                `json:"Composer,omitempty" xml:"Composer,omitempty"`
	ContentDisposition         *string                `json:"ContentDisposition,omitempty" xml:"ContentDisposition,omitempty"`
	ContentEncoding            *string                `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	ContentLanguage            *string                `json:"ContentLanguage,omitempty" xml:"ContentLanguage,omitempty"`
	ContentMd5                 *string                `json:"ContentMd5,omitempty" xml:"ContentMd5,omitempty"`
	ContentType                *string                `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	CreateTime                 *string                `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	CroppingSuggestions        []*CroppingSuggestion  `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	CustomId                   *string                `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	CustomLabels               map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	DatasetName                *string                `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	Duration                   *float64               `json:"Duration,omitempty" xml:"Duration,omitempty"`
	ETag                       *string                `json:"ETag,omitempty" xml:"ETag,omitempty"`
	EXIF                       *string                `json:"EXIF,omitempty" xml:"EXIF,omitempty"`
	Elements                   []*Element             `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	FigureCount                *int64                 `json:"FigureCount,omitempty" xml:"FigureCount,omitempty"`
	Figures                    []*Figure              `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
	FileAccessTime             *string                `json:"FileAccessTime,omitempty" xml:"FileAccessTime,omitempty"`
	FileCreateTime             *string                `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	FileHash                   *string                `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	FileModifiedTime           *string                `json:"FileModifiedTime,omitempty" xml:"FileModifiedTime,omitempty"`
	Filename                   *string                `json:"Filename,omitempty" xml:"Filename,omitempty"`
	FormatLongName             *string                `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	FormatName                 *string                `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	ImageHeight                *int64                 `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	ImageScore                 *ImageScore            `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	ImageWidth                 *int64                 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	// if can be null:
	// true
	Insights                              *Insights              `json:"Insights,omitempty" xml:"Insights,omitempty"`
	Labels                                []*Label               `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	Language                              *string                `json:"Language,omitempty" xml:"Language,omitempty"`
	LatLong                               *string                `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	MediaType                             *string                `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	OCRContents                           []*OCRContents         `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
	OCRTexts                              *string                `json:"OCRTexts,omitempty" xml:"OCRTexts,omitempty"`
	OSSCRC64                              *string                `json:"OSSCRC64,omitempty" xml:"OSSCRC64,omitempty"`
	OSSDeleteMarker                       *string                `json:"OSSDeleteMarker,omitempty" xml:"OSSDeleteMarker,omitempty"`
	OSSExpiration                         *string                `json:"OSSExpiration,omitempty" xml:"OSSExpiration,omitempty"`
	OSSObjectType                         *string                `json:"OSSObjectType,omitempty" xml:"OSSObjectType,omitempty"`
	OSSStorageClass                       *string                `json:"OSSStorageClass,omitempty" xml:"OSSStorageClass,omitempty"`
	OSSTagging                            map[string]interface{} `json:"OSSTagging,omitempty" xml:"OSSTagging,omitempty"`
	OSSTaggingCount                       *int64                 `json:"OSSTaggingCount,omitempty" xml:"OSSTaggingCount,omitempty"`
	OSSURI                                *string                `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	OSSUserMeta                           map[string]interface{} `json:"OSSUserMeta,omitempty" xml:"OSSUserMeta,omitempty"`
	OSSVersionId                          *string                `json:"OSSVersionId,omitempty" xml:"OSSVersionId,omitempty"`
	ObjectACL                             *string                `json:"ObjectACL,omitempty" xml:"ObjectACL,omitempty"`
	ObjectId                              *string                `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	ObjectStatus                          *string                `json:"ObjectStatus,omitempty" xml:"ObjectStatus,omitempty"`
	ObjectType                            *string                `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	Orientation                           *int64                 `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	OwnerId                               *string                `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	PageCount                             *int64                 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	Performer                             *string                `json:"Performer,omitempty" xml:"Performer,omitempty"`
	ProduceTime                           *string                `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	ProgramCount                          *int64                 `json:"ProgramCount,omitempty" xml:"ProgramCount,omitempty"`
	ProjectName                           *string                `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	Reason                                *string                `json:"Reason,omitempty" xml:"Reason,omitempty"`
	SceneElements                         []*SceneElement        `json:"SceneElements,omitempty" xml:"SceneElements,omitempty" type:"Repeated"`
	SemanticTypes                         []*string              `json:"SemanticTypes,omitempty" xml:"SemanticTypes,omitempty" type:"Repeated"`
	ServerSideDataEncryption              *string                `json:"ServerSideDataEncryption,omitempty" xml:"ServerSideDataEncryption,omitempty"`
	ServerSideEncryption                  *string                `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	ServerSideEncryptionCustomerAlgorithm *string                `json:"ServerSideEncryptionCustomerAlgorithm,omitempty" xml:"ServerSideEncryptionCustomerAlgorithm,omitempty"`
	ServerSideEncryptionKeyId             *string                `json:"ServerSideEncryptionKeyId,omitempty" xml:"ServerSideEncryptionKeyId,omitempty"`
	Size                                  *int64                 `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime                             *float64               `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	StreamCount                           *int64                 `json:"StreamCount,omitempty" xml:"StreamCount,omitempty"`
	Subtitles                             []*SubtitleStream      `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	Timezone                              *string                `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	Title                                 *string                `json:"Title,omitempty" xml:"Title,omitempty"`
	TravelClusterId                       *string                `json:"TravelClusterId,omitempty" xml:"TravelClusterId,omitempty"`
	URI                                   *string                `json:"URI,omitempty" xml:"URI,omitempty"`
	UpdateTime                            *string                `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	VideoHeight                           *int64                 `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	VideoStreams                          []*VideoStream         `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Repeated"`
	VideoWidth                            *int64                 `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
}

func (s File) String() string {
	return dara.Prettify(s)
}

func (s File) GoString() string {
	return s.String()
}

func (s *File) GetAccessControlAllowOrigin() *string {
	return s.AccessControlAllowOrigin
}

func (s *File) GetAccessControlRequestMethod() *string {
	return s.AccessControlRequestMethod
}

func (s *File) GetAddresses() []*Address {
	return s.Addresses
}

func (s *File) GetAlbum() *string {
	return s.Album
}

func (s *File) GetAlbumArtist() *string {
	return s.AlbumArtist
}

func (s *File) GetArtist() *string {
	return s.Artist
}

func (s *File) GetAudioCovers() []*Image {
	return s.AudioCovers
}

func (s *File) GetAudioStreams() []*AudioStream {
	return s.AudioStreams
}

func (s *File) GetBitrate() *int64 {
	return s.Bitrate
}

func (s *File) GetCacheControl() *string {
	return s.CacheControl
}

func (s *File) GetComposer() *string {
	return s.Composer
}

func (s *File) GetContentDisposition() *string {
	return s.ContentDisposition
}

func (s *File) GetContentEncoding() *string {
	return s.ContentEncoding
}

func (s *File) GetContentLanguage() *string {
	return s.ContentLanguage
}

func (s *File) GetContentMd5() *string {
	return s.ContentMd5
}

func (s *File) GetContentType() *string {
	return s.ContentType
}

func (s *File) GetCreateTime() *string {
	return s.CreateTime
}

func (s *File) GetCroppingSuggestions() []*CroppingSuggestion {
	return s.CroppingSuggestions
}

func (s *File) GetCustomId() *string {
	return s.CustomId
}

func (s *File) GetCustomLabels() map[string]interface{} {
	return s.CustomLabels
}

func (s *File) GetDatasetName() *string {
	return s.DatasetName
}

func (s *File) GetDuration() *float64 {
	return s.Duration
}

func (s *File) GetETag() *string {
	return s.ETag
}

func (s *File) GetEXIF() *string {
	return s.EXIF
}

func (s *File) GetElements() []*Element {
	return s.Elements
}

func (s *File) GetFigureCount() *int64 {
	return s.FigureCount
}

func (s *File) GetFigures() []*Figure {
	return s.Figures
}

func (s *File) GetFileAccessTime() *string {
	return s.FileAccessTime
}

func (s *File) GetFileCreateTime() *string {
	return s.FileCreateTime
}

func (s *File) GetFileHash() *string {
	return s.FileHash
}

func (s *File) GetFileModifiedTime() *string {
	return s.FileModifiedTime
}

func (s *File) GetFilename() *string {
	return s.Filename
}

func (s *File) GetFormatLongName() *string {
	return s.FormatLongName
}

func (s *File) GetFormatName() *string {
	return s.FormatName
}

func (s *File) GetImageHeight() *int64 {
	return s.ImageHeight
}

func (s *File) GetImageScore() *ImageScore {
	return s.ImageScore
}

func (s *File) GetImageWidth() *int64 {
	return s.ImageWidth
}

func (s *File) GetInsights() *Insights {
	return s.Insights
}

func (s *File) GetLabels() []*Label {
	return s.Labels
}

func (s *File) GetLanguage() *string {
	return s.Language
}

func (s *File) GetLatLong() *string {
	return s.LatLong
}

func (s *File) GetMediaType() *string {
	return s.MediaType
}

func (s *File) GetOCRContents() []*OCRContents {
	return s.OCRContents
}

func (s *File) GetOCRTexts() *string {
	return s.OCRTexts
}

func (s *File) GetOSSCRC64() *string {
	return s.OSSCRC64
}

func (s *File) GetOSSDeleteMarker() *string {
	return s.OSSDeleteMarker
}

func (s *File) GetOSSExpiration() *string {
	return s.OSSExpiration
}

func (s *File) GetOSSObjectType() *string {
	return s.OSSObjectType
}

func (s *File) GetOSSStorageClass() *string {
	return s.OSSStorageClass
}

func (s *File) GetOSSTagging() map[string]interface{} {
	return s.OSSTagging
}

func (s *File) GetOSSTaggingCount() *int64 {
	return s.OSSTaggingCount
}

func (s *File) GetOSSURI() *string {
	return s.OSSURI
}

func (s *File) GetOSSUserMeta() map[string]interface{} {
	return s.OSSUserMeta
}

func (s *File) GetOSSVersionId() *string {
	return s.OSSVersionId
}

func (s *File) GetObjectACL() *string {
	return s.ObjectACL
}

func (s *File) GetObjectId() *string {
	return s.ObjectId
}

func (s *File) GetObjectStatus() *string {
	return s.ObjectStatus
}

func (s *File) GetObjectType() *string {
	return s.ObjectType
}

func (s *File) GetOrientation() *int64 {
	return s.Orientation
}

func (s *File) GetOwnerId() *string {
	return s.OwnerId
}

func (s *File) GetPageCount() *int64 {
	return s.PageCount
}

func (s *File) GetPerformer() *string {
	return s.Performer
}

func (s *File) GetProduceTime() *string {
	return s.ProduceTime
}

func (s *File) GetProgramCount() *int64 {
	return s.ProgramCount
}

func (s *File) GetProjectName() *string {
	return s.ProjectName
}

func (s *File) GetReason() *string {
	return s.Reason
}

func (s *File) GetSceneElements() []*SceneElement {
	return s.SceneElements
}

func (s *File) GetSemanticTypes() []*string {
	return s.SemanticTypes
}

func (s *File) GetServerSideDataEncryption() *string {
	return s.ServerSideDataEncryption
}

func (s *File) GetServerSideEncryption() *string {
	return s.ServerSideEncryption
}

func (s *File) GetServerSideEncryptionCustomerAlgorithm() *string {
	return s.ServerSideEncryptionCustomerAlgorithm
}

func (s *File) GetServerSideEncryptionKeyId() *string {
	return s.ServerSideEncryptionKeyId
}

func (s *File) GetSize() *int64 {
	return s.Size
}

func (s *File) GetStartTime() *float64 {
	return s.StartTime
}

func (s *File) GetStreamCount() *int64 {
	return s.StreamCount
}

func (s *File) GetSubtitles() []*SubtitleStream {
	return s.Subtitles
}

func (s *File) GetTimezone() *string {
	return s.Timezone
}

func (s *File) GetTitle() *string {
	return s.Title
}

func (s *File) GetTravelClusterId() *string {
	return s.TravelClusterId
}

func (s *File) GetURI() *string {
	return s.URI
}

func (s *File) GetUpdateTime() *string {
	return s.UpdateTime
}

func (s *File) GetVideoHeight() *int64 {
	return s.VideoHeight
}

func (s *File) GetVideoStreams() []*VideoStream {
	return s.VideoStreams
}

func (s *File) GetVideoWidth() *int64 {
	return s.VideoWidth
}

func (s *File) SetAccessControlAllowOrigin(v string) *File {
	s.AccessControlAllowOrigin = &v
	return s
}

func (s *File) SetAccessControlRequestMethod(v string) *File {
	s.AccessControlRequestMethod = &v
	return s
}

func (s *File) SetAddresses(v []*Address) *File {
	s.Addresses = v
	return s
}

func (s *File) SetAlbum(v string) *File {
	s.Album = &v
	return s
}

func (s *File) SetAlbumArtist(v string) *File {
	s.AlbumArtist = &v
	return s
}

func (s *File) SetArtist(v string) *File {
	s.Artist = &v
	return s
}

func (s *File) SetAudioCovers(v []*Image) *File {
	s.AudioCovers = v
	return s
}

func (s *File) SetAudioStreams(v []*AudioStream) *File {
	s.AudioStreams = v
	return s
}

func (s *File) SetBitrate(v int64) *File {
	s.Bitrate = &v
	return s
}

func (s *File) SetCacheControl(v string) *File {
	s.CacheControl = &v
	return s
}

func (s *File) SetComposer(v string) *File {
	s.Composer = &v
	return s
}

func (s *File) SetContentDisposition(v string) *File {
	s.ContentDisposition = &v
	return s
}

func (s *File) SetContentEncoding(v string) *File {
	s.ContentEncoding = &v
	return s
}

func (s *File) SetContentLanguage(v string) *File {
	s.ContentLanguage = &v
	return s
}

func (s *File) SetContentMd5(v string) *File {
	s.ContentMd5 = &v
	return s
}

func (s *File) SetContentType(v string) *File {
	s.ContentType = &v
	return s
}

func (s *File) SetCreateTime(v string) *File {
	s.CreateTime = &v
	return s
}

func (s *File) SetCroppingSuggestions(v []*CroppingSuggestion) *File {
	s.CroppingSuggestions = v
	return s
}

func (s *File) SetCustomId(v string) *File {
	s.CustomId = &v
	return s
}

func (s *File) SetCustomLabels(v map[string]interface{}) *File {
	s.CustomLabels = v
	return s
}

func (s *File) SetDatasetName(v string) *File {
	s.DatasetName = &v
	return s
}

func (s *File) SetDuration(v float64) *File {
	s.Duration = &v
	return s
}

func (s *File) SetETag(v string) *File {
	s.ETag = &v
	return s
}

func (s *File) SetEXIF(v string) *File {
	s.EXIF = &v
	return s
}

func (s *File) SetElements(v []*Element) *File {
	s.Elements = v
	return s
}

func (s *File) SetFigureCount(v int64) *File {
	s.FigureCount = &v
	return s
}

func (s *File) SetFigures(v []*Figure) *File {
	s.Figures = v
	return s
}

func (s *File) SetFileAccessTime(v string) *File {
	s.FileAccessTime = &v
	return s
}

func (s *File) SetFileCreateTime(v string) *File {
	s.FileCreateTime = &v
	return s
}

func (s *File) SetFileHash(v string) *File {
	s.FileHash = &v
	return s
}

func (s *File) SetFileModifiedTime(v string) *File {
	s.FileModifiedTime = &v
	return s
}

func (s *File) SetFilename(v string) *File {
	s.Filename = &v
	return s
}

func (s *File) SetFormatLongName(v string) *File {
	s.FormatLongName = &v
	return s
}

func (s *File) SetFormatName(v string) *File {
	s.FormatName = &v
	return s
}

func (s *File) SetImageHeight(v int64) *File {
	s.ImageHeight = &v
	return s
}

func (s *File) SetImageScore(v *ImageScore) *File {
	s.ImageScore = v
	return s
}

func (s *File) SetImageWidth(v int64) *File {
	s.ImageWidth = &v
	return s
}

func (s *File) SetInsights(v *Insights) *File {
	s.Insights = v
	return s
}

func (s *File) SetLabels(v []*Label) *File {
	s.Labels = v
	return s
}

func (s *File) SetLanguage(v string) *File {
	s.Language = &v
	return s
}

func (s *File) SetLatLong(v string) *File {
	s.LatLong = &v
	return s
}

func (s *File) SetMediaType(v string) *File {
	s.MediaType = &v
	return s
}

func (s *File) SetOCRContents(v []*OCRContents) *File {
	s.OCRContents = v
	return s
}

func (s *File) SetOCRTexts(v string) *File {
	s.OCRTexts = &v
	return s
}

func (s *File) SetOSSCRC64(v string) *File {
	s.OSSCRC64 = &v
	return s
}

func (s *File) SetOSSDeleteMarker(v string) *File {
	s.OSSDeleteMarker = &v
	return s
}

func (s *File) SetOSSExpiration(v string) *File {
	s.OSSExpiration = &v
	return s
}

func (s *File) SetOSSObjectType(v string) *File {
	s.OSSObjectType = &v
	return s
}

func (s *File) SetOSSStorageClass(v string) *File {
	s.OSSStorageClass = &v
	return s
}

func (s *File) SetOSSTagging(v map[string]interface{}) *File {
	s.OSSTagging = v
	return s
}

func (s *File) SetOSSTaggingCount(v int64) *File {
	s.OSSTaggingCount = &v
	return s
}

func (s *File) SetOSSURI(v string) *File {
	s.OSSURI = &v
	return s
}

func (s *File) SetOSSUserMeta(v map[string]interface{}) *File {
	s.OSSUserMeta = v
	return s
}

func (s *File) SetOSSVersionId(v string) *File {
	s.OSSVersionId = &v
	return s
}

func (s *File) SetObjectACL(v string) *File {
	s.ObjectACL = &v
	return s
}

func (s *File) SetObjectId(v string) *File {
	s.ObjectId = &v
	return s
}

func (s *File) SetObjectStatus(v string) *File {
	s.ObjectStatus = &v
	return s
}

func (s *File) SetObjectType(v string) *File {
	s.ObjectType = &v
	return s
}

func (s *File) SetOrientation(v int64) *File {
	s.Orientation = &v
	return s
}

func (s *File) SetOwnerId(v string) *File {
	s.OwnerId = &v
	return s
}

func (s *File) SetPageCount(v int64) *File {
	s.PageCount = &v
	return s
}

func (s *File) SetPerformer(v string) *File {
	s.Performer = &v
	return s
}

func (s *File) SetProduceTime(v string) *File {
	s.ProduceTime = &v
	return s
}

func (s *File) SetProgramCount(v int64) *File {
	s.ProgramCount = &v
	return s
}

func (s *File) SetProjectName(v string) *File {
	s.ProjectName = &v
	return s
}

func (s *File) SetReason(v string) *File {
	s.Reason = &v
	return s
}

func (s *File) SetSceneElements(v []*SceneElement) *File {
	s.SceneElements = v
	return s
}

func (s *File) SetSemanticTypes(v []*string) *File {
	s.SemanticTypes = v
	return s
}

func (s *File) SetServerSideDataEncryption(v string) *File {
	s.ServerSideDataEncryption = &v
	return s
}

func (s *File) SetServerSideEncryption(v string) *File {
	s.ServerSideEncryption = &v
	return s
}

func (s *File) SetServerSideEncryptionCustomerAlgorithm(v string) *File {
	s.ServerSideEncryptionCustomerAlgorithm = &v
	return s
}

func (s *File) SetServerSideEncryptionKeyId(v string) *File {
	s.ServerSideEncryptionKeyId = &v
	return s
}

func (s *File) SetSize(v int64) *File {
	s.Size = &v
	return s
}

func (s *File) SetStartTime(v float64) *File {
	s.StartTime = &v
	return s
}

func (s *File) SetStreamCount(v int64) *File {
	s.StreamCount = &v
	return s
}

func (s *File) SetSubtitles(v []*SubtitleStream) *File {
	s.Subtitles = v
	return s
}

func (s *File) SetTimezone(v string) *File {
	s.Timezone = &v
	return s
}

func (s *File) SetTitle(v string) *File {
	s.Title = &v
	return s
}

func (s *File) SetTravelClusterId(v string) *File {
	s.TravelClusterId = &v
	return s
}

func (s *File) SetURI(v string) *File {
	s.URI = &v
	return s
}

func (s *File) SetUpdateTime(v string) *File {
	s.UpdateTime = &v
	return s
}

func (s *File) SetVideoHeight(v int64) *File {
	s.VideoHeight = &v
	return s
}

func (s *File) SetVideoStreams(v []*VideoStream) *File {
	s.VideoStreams = v
	return s
}

func (s *File) SetVideoWidth(v int64) *File {
	s.VideoWidth = &v
	return s
}

func (s *File) Validate() error {
	if s.Addresses != nil {
		for _, item := range s.Addresses {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AudioCovers != nil {
		for _, item := range s.AudioCovers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.AudioStreams != nil {
		for _, item := range s.AudioStreams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.CroppingSuggestions != nil {
		for _, item := range s.CroppingSuggestions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Elements != nil {
		for _, item := range s.Elements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Figures != nil {
		for _, item := range s.Figures {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.ImageScore != nil {
		if err := s.ImageScore.Validate(); err != nil {
			return err
		}
	}
	if s.Insights != nil {
		if err := s.Insights.Validate(); err != nil {
			return err
		}
	}
	if s.Labels != nil {
		for _, item := range s.Labels {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.OCRContents != nil {
		for _, item := range s.OCRContents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SceneElements != nil {
		for _, item := range s.SceneElements {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Subtitles != nil {
		for _, item := range s.Subtitles {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.VideoStreams != nil {
		for _, item := range s.VideoStreams {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
