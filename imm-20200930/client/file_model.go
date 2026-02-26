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
	// The origin allowed in cross-origin requests.
	//
	// example:
	//
	// https://aliyundoc.com
	AccessControlAllowOrigin *string `json:"AccessControlAllowOrigin,omitempty" xml:"AccessControlAllowOrigin,omitempty"`
	// The method to be used in the actual cross-origin request.
	//
	// example:
	//
	// PUT
	AccessControlRequestMethod *string `json:"AccessControlRequestMethod,omitempty" xml:"AccessControlRequestMethod,omitempty"`
	// The addresses.
	Addresses []*Address `json:"Addresses,omitempty" xml:"Addresses,omitempty" type:"Repeated"`
	// The album.
	//
	// example:
	//
	// FirstAlbum
	Album *string `json:"Album,omitempty" xml:"Album,omitempty"`
	// The singer.
	//
	// example:
	//
	// Jane
	AlbumArtist *string `json:"AlbumArtist,omitempty" xml:"AlbumArtist,omitempty"`
	// The artist.
	//
	// example:
	//
	// Jane
	Artist *string `json:"Artist,omitempty" xml:"Artist,omitempty"`
	// The audio covers.
	AudioCovers []*Image `json:"AudioCovers,omitempty" xml:"AudioCovers,omitempty" type:"Repeated"`
	// The list of audio streams.
	AudioStreams []*AudioStream `json:"AudioStreams,omitempty" xml:"AudioStreams,omitempty" type:"Repeated"`
	// The bitrate. Unit: bit/s.
	//
	// example:
	//
	// 13091201
	Bitrate *int64 `json:"Bitrate,omitempty" xml:"Bitrate,omitempty"`
	// The caching behavior of the web page when the object is downloaded.
	//
	// This parameter corresponds to the Cache-Control HTTP header of the OSS object. For more information, see [Manage object metadata](https://help.aliyun.com/document_detail/31859.html).
	//
	// example:
	//
	// no-cache
	CacheControl *string `json:"CacheControl,omitempty" xml:"CacheControl,omitempty"`
	// The composer.
	//
	// example:
	//
	// Jane
	Composer *string `json:"Composer,omitempty" xml:"Composer,omitempty"`
	// The name of the object during the download.
	//
	// This parameter corresponds to the Content-Disposition HTTP header of the OSS object. For more information, see [Manage object metadata](https://help.aliyun.com/document_detail/31859.html).
	//
	// example:
	//
	// attachment; filename =test.jpg
	ContentDisposition *string `json:"ContentDisposition,omitempty" xml:"ContentDisposition,omitempty"`
	// The content encoding format of the object when the object is downloaded.
	//
	// This parameter corresponds to the Content-Encoding HTTP header of the OSS object. For more information, see [Manage object metadata](https://help.aliyun.com/document_detail/31859.html).
	//
	// example:
	//
	// UTF-8
	ContentEncoding *string `json:"ContentEncoding,omitempty" xml:"ContentEncoding,omitempty"`
	// The language of the object content.
	//
	// This parameter corresponds to the Content-Language HTTP header of the OSS object. For more information, see [Manage object metadata](https://help.aliyun.com/document_detail/31859.html).
	//
	// example:
	//
	// zh-CN
	ContentLanguage *string `json:"ContentLanguage,omitempty" xml:"ContentLanguage,omitempty"`
	// The MD5 value.
	//
	// example:
	//
	// HZwoCnxPZ/fvhz4oRJ2+Fw==
	ContentMd5 *string `json:"ContentMd5,omitempty" xml:"ContentMd5,omitempty"`
	// The Multipurpose Internet Mail Extensions (MIME) type of the file.
	//
	// example:
	//
	// image/jpeg
	ContentType *string `json:"ContentType,omitempty" xml:"ContentType,omitempty"`
	// The RFC3339Nano timestamp when the metadata was created.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	CreateTime *string `json:"CreateTime,omitempty" xml:"CreateTime,omitempty"`
	// The cropping suggestions for the image.
	//
	// > Not supported.
	CroppingSuggestions []*CroppingSuggestion `json:"CroppingSuggestions,omitempty" xml:"CroppingSuggestions,omitempty" type:"Repeated"`
	// The custom ID of the file. When the cluster is indexed into the dataset, the custom ID is stored as the data attribute. You can map the custom ID to other data in your business system. Configure this parameter based on your business requirements. For example, you can associate a URI with an ID in your system. We recommend that you set this parameter to a globally unique value.
	//
	// example:
	//
	// member-image-id-0001
	CustomId *string `json:"CustomId,omitempty" xml:"CustomId,omitempty"`
	// The custom labels of the file. This parameter is optional. The parameter stores custom key-value labels, which can be used to filter data.
	//
	// example:
	//
	// {
	//
	//       "MemberName": "Tim",
	//
	//       "Enabled": "True",
	//
	//       "ItemCount": "10"
	//
	// }
	CustomLabels map[string]interface{} `json:"CustomLabels,omitempty" xml:"CustomLabels,omitempty"`
	// The name of the dataset. You can obtain the name of the dataset from the response of the [CreateDataset](https://help.aliyun.com/document_detail/478160.html) operation.
	//
	// example:
	//
	// test-dataset
	DatasetName *string `json:"DatasetName,omitempty" xml:"DatasetName,omitempty"`
	// The total duration of the video. Unit: seconds.
	//
	// example:
	//
	// 15.263000
	Duration *float64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The ETag of the object. ETags are used to identify the content of objects.
	//
	// example:
	//
	// "1D9C280A7C4F67F7EF873E28449****"
	ETag *string `json:"ETag,omitempty" xml:"ETag,omitempty"`
	// The original EXIF information about the image. The EXIF information is stored in the serialized JSON format. For more information, see [Query image information](https://help.aliyun.com/document_detail/44975.html).
	//
	// example:
	//
	// {"Compression":{"value":"6"},"DateTime":{"value":"2020:08:19 17:11:11"}}
	EXIF *string `json:"EXIF,omitempty" xml:"EXIF,omitempty"`
	// The document elements that match the current query content when you call the SemanticQuery operation for semantic search.
	Elements []*Element `json:"Elements,omitempty" xml:"Elements,omitempty" type:"Repeated"`
	// The number of persons.
	//
	// example:
	//
	// 10
	FigureCount *int64 `json:"FigureCount,omitempty" xml:"FigureCount,omitempty"`
	// The list of persons. The persons are detected via AI models.
	Figures []*Figure `json:"Figures,omitempty" xml:"Figures,omitempty" type:"Repeated"`
	// The RFC3339Nano timestamp when the file was accessed.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	FileAccessTime *string `json:"FileAccessTime,omitempty" xml:"FileAccessTime,omitempty"`
	// The RFC3339Nano timestamp when the file was created.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	FileCreateTime *string `json:"FileCreateTime,omitempty" xml:"FileCreateTime,omitempty"`
	// The hash value of the file.
	//
	// example:
	//
	// 1d9c280a7c4f67f7ef873e28449dbe17
	FileHash *string `json:"FileHash,omitempty" xml:"FileHash,omitempty"`
	// The RFC3339Nano timestamp when the file was last modified.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	FileModifiedTime *string `json:"FileModifiedTime,omitempty" xml:"FileModifiedTime,omitempty"`
	// The name of the object. For an OSS object, the value of this parameter is the object name.
	//
	// example:
	//
	// sampleobject.jpg
	Filename *string `json:"Filename,omitempty" xml:"Filename,omitempty"`
	// The full name of the media format.
	//
	// example:
	//
	// QuickTime / MOV
	FormatLongName *string `json:"FormatLongName,omitempty" xml:"FormatLongName,omitempty"`
	// The name of the media format.
	//
	// example:
	//
	// mov
	FormatName *string `json:"FormatName,omitempty" xml:"FormatName,omitempty"`
	// The height of the image. Unit: pixels.
	//
	// example:
	//
	// 500
	ImageHeight *int64 `json:"ImageHeight,omitempty" xml:"ImageHeight,omitempty"`
	// The score of the image. The score is calculated by using AI models.
	ImageScore *ImageScore `json:"ImageScore,omitempty" xml:"ImageScore,omitempty"`
	// The width of the image. Unit: pixels.
	//
	// example:
	//
	// 270
	ImageWidth *int64 `json:"ImageWidth,omitempty" xml:"ImageWidth,omitempty"`
	// Summary and description of the file.
	//
	// > Not supported.
	//
	// if can be null:
	// true
	Insights *Insights `json:"Insights,omitempty" xml:"Insights,omitempty"`
	// The labels of the file. The labels are detected via AI models.
	Labels []*Label `json:"Labels,omitempty" xml:"Labels,omitempty" type:"Repeated"`
	// The language specified by using a BCP 47 language tag.
	//
	// example:
	//
	// eng
	Language *string `json:"Language,omitempty" xml:"Language,omitempty"`
	// The latitude and longitude.
	//
	// example:
	//
	// 30.134390,120.074997
	LatLong *string `json:"LatLong,omitempty" xml:"LatLong,omitempty"`
	// The media type of the file.
	//
	// Valid values:
	//
	// 	- image
	//
	// 	- other
	//
	// 	- document
	//
	// 	- archive
	//
	// 	- audio
	//
	// 	- video
	//
	// example:
	//
	// image
	MediaType *string `json:"MediaType,omitempty" xml:"MediaType,omitempty"`
	// The Optical Character Recognition (OCR) results.
	//
	// > Not supported.
	OCRContents []*OCRContents `json:"OCRContents,omitempty" xml:"OCRContents,omitempty" type:"Repeated"`
	// The text detected in the image.
	//
	// example:
	//
	// Alibaba Cloud IMM
	OCRTexts *string `json:"OCRTexts,omitempty" xml:"OCRTexts,omitempty"`
	// The CRC64 value.
	//
	// example:
	//
	// 559890638950338001
	OSSCRC64 *string `json:"OSSCRC64,omitempty" xml:"OSSCRC64,omitempty"`
	// The delete marker of the object.
	//
	// example:
	//
	// CAEQMhiBgIDXiaaB0BYiIGQzYmRkZGUxMTM1ZDRjOTZhNjk4YjRjMTAyZjhl****
	OSSDeleteMarker *string `json:"OSSDeleteMarker,omitempty" xml:"OSSDeleteMarker,omitempty"`
	// The expiration time of the OSS object.
	//
	// This parameter corresponds to the Expires HTTP header of the OSS object. For more information, see [Manage object metadata](https://help.aliyun.com/document_detail/31859.html).
	//
	// example:
	//
	// 2120-01-01T12:00:00.000Z
	OSSExpiration *string `json:"OSSExpiration,omitempty" xml:"OSSExpiration,omitempty"`
	// The type of the OSS object. Set the value to `Normal`.
	//
	// example:
	//
	// Normal
	OSSObjectType *string `json:"OSSObjectType,omitempty" xml:"OSSObjectType,omitempty"`
	// The storage class of the OSS bucket.
	//
	// example:
	//
	// Standard
	OSSStorageClass *string `json:"OSSStorageClass,omitempty" xml:"OSSStorageClass,omitempty"`
	// The tag of the object.
	//
	// For more information, see [Add tags to an object](https://help.aliyun.com/document_detail/106678.html).
	//
	// example:
	//
	// {"key": "val"}
	OSSTagging map[string]interface{} `json:"OSSTagging,omitempty" xml:"OSSTagging,omitempty"`
	// The number of OSS object tags.
	//
	// This parameter is available only if tags are added to the corresponding OSS object. For more information, see [Add tags to an object](https://help.aliyun.com/document_detail/106678.html).
	//
	// example:
	//
	// 2
	OSSTaggingCount *int64 `json:"OSSTaggingCount,omitempty" xml:"OSSTaggingCount,omitempty"`
	// The URI of the OSS object. This parameter is available only if the value of the URI parameter is the URI of a file in Photo and Drive Service.
	//
	// example:
	//
	// oss://examplebucket/sampleobject.jpg
	OSSURI *string `json:"OSSURI,omitempty" xml:"OSSURI,omitempty"`
	// The user metadata of the OSS object.
	//
	// This parameter is available only if user metadata is configured for the OSS object. For more information, see [Manage object metadata](https://help.aliyun.com/document_detail/31859.html).
	//
	// example:
	//
	// {"key": "val"}
	OSSUserMeta map[string]interface{} `json:"OSSUserMeta,omitempty" xml:"OSSUserMeta,omitempty"`
	// The version of the object.
	//
	// This parameter is available only if versioning is enabled for the bucket. For more information, see [Overview](https://help.aliyun.com/document_detail/109695.html).
	//
	// example:
	//
	// CAEQNhiBgMDJgZCA0BYiIDc4MGZjZGI2OTBjOTRmNTE5NmU5NmFhZjhjYmY0****
	OSSVersionId *string `json:"OSSVersionId,omitempty" xml:"OSSVersionId,omitempty"`
	// The access control list (ACL) of the OSS object.
	//
	// example:
	//
	// default
	ObjectACL *string `json:"ObjectACL,omitempty" xml:"ObjectACL,omitempty"`
	// The unique ID of the object.
	//
	// example:
	//
	// 75d5de2c50754e3dadd5c35dbca5f9949369e37eb342a73821f690c94c36c7f7
	ObjectId *string `json:"ObjectId,omitempty" xml:"ObjectId,omitempty"`
	// The status of the object.
	//
	// example:
	//
	// Indexed
	ObjectStatus *string `json:"ObjectStatus,omitempty" xml:"ObjectStatus,omitempty"`
	// The type of the object. Set the value to **file**.
	//
	// example:
	//
	// file
	ObjectType *string `json:"ObjectType,omitempty" xml:"ObjectType,omitempty"`
	// The image rotation angle. You can obtain the value from the exchangeable image file format (EXIF).
	//
	// If the EXIF metadata does not contain the image rotation angle, this parameter is not included in the response.
	//
	// example:
	//
	// 0
	Orientation *int64 `json:"Orientation,omitempty" xml:"Orientation,omitempty"`
	// The ID of the Alibaba Cloud account.
	//
	// example:
	//
	// 102321002467****
	OwnerId *string `json:"OwnerId,omitempty" xml:"OwnerId,omitempty"`
	// The number of pages.
	//
	// > Not supported.
	//
	// example:
	//
	// 5
	PageCount *int64 `json:"PageCount,omitempty" xml:"PageCount,omitempty"`
	// The performer.
	//
	// example:
	//
	// Jane
	Performer *string `json:"Performer,omitempty" xml:"Performer,omitempty"`
	// The time when the image was taken.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	ProduceTime *string `json:"ProduceTime,omitempty" xml:"ProduceTime,omitempty"`
	// The number of programs in the media container.
	//
	// example:
	//
	// 1
	ProgramCount *int64 `json:"ProgramCount,omitempty" xml:"ProgramCount,omitempty"`
	// The name of the project. You can obtain the name of the project from the response of the [CreateProject](https://help.aliyun.com/document_detail/478153.html) operation.
	//
	// example:
	//
	// test-project
	ProjectName *string `json:"ProjectName,omitempty" xml:"ProjectName,omitempty"`
	// The reason why the file failed to run the index.
	//
	// example:
	//
	// [InternalError] The request has been failed due to some unknown error. status: 500, requestId: CC5ACFBD-BB7A-496D-A9D6-****
	Reason *string `json:"Reason,omitempty" xml:"Reason,omitempty"`
	// The elements in the video segment, which are scene elements that you can extract from the video by using an AI model.
	SceneElements []*SceneElement `json:"SceneElements,omitempty" xml:"SceneElements,omitempty" type:"Repeated"`
	// The reasons for which the current file is included in the search results when you call the SemanticQuery operation for semantic search.
	SemanticTypes []*string `json:"SemanticTypes,omitempty" xml:"SemanticTypes,omitempty" type:"Repeated"`
	// The encryption method of the object.
	//
	// This parameter is available only if server encryption is configured for the OSS bucket. For more information, see [Server-side encryption](https://help.aliyun.com/document_detail/31871.html).
	//
	// example:
	//
	// SM4
	ServerSideDataEncryption *string `json:"ServerSideDataEncryption,omitempty" xml:"ServerSideDataEncryption,omitempty"`
	// The encryption method on the server side.
	//
	// This parameter is available only if server encryption is configured for the OSS bucket. For more information, see [Server-side encryption](https://help.aliyun.com/document_detail/31871.html).
	//
	// example:
	//
	// AES256
	ServerSideEncryption *string `json:"ServerSideEncryption,omitempty" xml:"ServerSideEncryption,omitempty"`
	// The algorithm that is used to encrypt the file on the server side.
	//
	// example:
	//
	// SM4
	ServerSideEncryptionCustomerAlgorithm *string `json:"ServerSideEncryptionCustomerAlgorithm,omitempty" xml:"ServerSideEncryptionCustomerAlgorithm,omitempty"`
	// The ID of the customer master key (CMK) managed by Key Management Service (KMS).
	//
	// This parameter is available only if server encryption is configured for the OSS bucket. For more information, see [Server-side encryption](https://help.aliyun.com/document_detail/31871.html).
	//
	// example:
	//
	// 9468da86-3509-4f8d-a61e-6eab1eac****
	ServerSideEncryptionKeyId *string `json:"ServerSideEncryptionKeyId,omitempty" xml:"ServerSideEncryptionKeyId,omitempty"`
	// The size of the object. Unit: bytes.
	//
	// example:
	//
	// 1000
	Size *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	// The time of the first frame. Unit: seconds.
	//
	// example:
	//
	// 0.000000
	StartTime *float64 `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	// The number of media streams in the media container.
	//
	// example:
	//
	// 1
	StreamCount *int64 `json:"StreamCount,omitempty" xml:"StreamCount,omitempty"`
	// The list of subtitle streams.
	Subtitles []*SubtitleStream `json:"Subtitles,omitempty" xml:"Subtitles,omitempty" type:"Repeated"`
	// The time zone.
	//
	// >  Not supported.
	//
	// example:
	//
	// ""
	Timezone *string `json:"Timezone,omitempty" xml:"Timezone,omitempty"`
	// The title of the file.
	//
	// example:
	//
	// test
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// A reserved parameter.
	//
	// > Not supported.
	//
	// example:
	//
	// -
	TravelClusterId *string `json:"TravelClusterId,omitempty" xml:"TravelClusterId,omitempty"`
	// The URI of the file.
	//
	// The URI of an OSS object follows the oss://${Bucket}/${Object} format, where `${Bucket}` is the name of the bucket in the same region as the current project and `${Object}` is the path of the object with the extension included.
	//
	// The URI of a file in Photo and Drive Service follows the pds://domains/${domain}/drives/${drive}/files/${file}/revisions/${revision} format.
	//
	// example:
	//
	// oss://examplebucket/sampleobject.jpg
	URI *string `json:"URI,omitempty" xml:"URI,omitempty"`
	// The RFC3339Nano timestamp when the metadata was modified.
	//
	// example:
	//
	// 2021-06-29T14:50:13.011643661+08:00
	UpdateTime *string `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
	// The height of the video. Unit: pixels.
	//
	// example:
	//
	// 1920
	VideoHeight *int64 `json:"VideoHeight,omitempty" xml:"VideoHeight,omitempty"`
	// The list of video streams.
	VideoStreams []*VideoStream `json:"VideoStreams,omitempty" xml:"VideoStreams,omitempty" type:"Repeated"`
	// The width of the video. Unit: pixels.
	//
	// example:
	//
	// 1080
	VideoWidth *int64 `json:"VideoWidth,omitempty" xml:"VideoWidth,omitempty"`
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
