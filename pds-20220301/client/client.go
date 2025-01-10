// This file is auto-generated, don't edit it. Thanks.
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pds/client"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AccountAccessTokenResponse struct {
	AccessToken        *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Avatar             *string            `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DefaultDriveId     *string            `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	DefaultSboxDriveId *string            `json:"default_sbox_drive_id,omitempty" xml:"default_sbox_drive_id,omitempty"`
	DeviceId           *string            `json:"device_id,omitempty" xml:"device_id,omitempty"`
	DeviceName         *string            `json:"device_name,omitempty" xml:"device_name,omitempty"`
	DomainId           *string            `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	ExistLink          []*LinkInfo        `json:"exist_link,omitempty" xml:"exist_link,omitempty" type:"Repeated"`
	ExpireTime         *string            `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	ExpiresIn          *int64             `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	IsFirstLogin       *bool              `json:"is_first_login,omitempty" xml:"is_first_login,omitempty"`
	NeedLink           *bool              `json:"need_link,omitempty" xml:"need_link,omitempty"`
	NeedRpVerify       *bool              `json:"need_rp_verify,omitempty" xml:"need_rp_verify,omitempty"`
	NickName           *string            `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	PathStatus         *string            `json:"path_status,omitempty" xml:"path_status,omitempty"`
	PinSetup           *bool              `json:"pin_setup,omitempty" xml:"pin_setup,omitempty"`
	RefreshToken       *string            `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Role               *string            `json:"role,omitempty" xml:"role,omitempty"`
	State              *string            `json:"state,omitempty" xml:"state,omitempty"`
	Status             *string            `json:"status,omitempty" xml:"status,omitempty"`
	TokenType          *string            `json:"token_type,omitempty" xml:"token_type,omitempty"`
	UserData           map[string]*string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	UserId             *string            `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName           *string            `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s AccountAccessTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s AccountAccessTokenResponse) GoString() string {
	return s.String()
}

func (s *AccountAccessTokenResponse) SetAccessToken(v string) *AccountAccessTokenResponse {
	s.AccessToken = &v
	return s
}

func (s *AccountAccessTokenResponse) SetAvatar(v string) *AccountAccessTokenResponse {
	s.Avatar = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDefaultDriveId(v string) *AccountAccessTokenResponse {
	s.DefaultDriveId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDefaultSboxDriveId(v string) *AccountAccessTokenResponse {
	s.DefaultSboxDriveId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDeviceId(v string) *AccountAccessTokenResponse {
	s.DeviceId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDeviceName(v string) *AccountAccessTokenResponse {
	s.DeviceName = &v
	return s
}

func (s *AccountAccessTokenResponse) SetDomainId(v string) *AccountAccessTokenResponse {
	s.DomainId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetExistLink(v []*LinkInfo) *AccountAccessTokenResponse {
	s.ExistLink = v
	return s
}

func (s *AccountAccessTokenResponse) SetExpireTime(v string) *AccountAccessTokenResponse {
	s.ExpireTime = &v
	return s
}

func (s *AccountAccessTokenResponse) SetExpiresIn(v int64) *AccountAccessTokenResponse {
	s.ExpiresIn = &v
	return s
}

func (s *AccountAccessTokenResponse) SetIsFirstLogin(v bool) *AccountAccessTokenResponse {
	s.IsFirstLogin = &v
	return s
}

func (s *AccountAccessTokenResponse) SetNeedLink(v bool) *AccountAccessTokenResponse {
	s.NeedLink = &v
	return s
}

func (s *AccountAccessTokenResponse) SetNeedRpVerify(v bool) *AccountAccessTokenResponse {
	s.NeedRpVerify = &v
	return s
}

func (s *AccountAccessTokenResponse) SetNickName(v string) *AccountAccessTokenResponse {
	s.NickName = &v
	return s
}

func (s *AccountAccessTokenResponse) SetPathStatus(v string) *AccountAccessTokenResponse {
	s.PathStatus = &v
	return s
}

func (s *AccountAccessTokenResponse) SetPinSetup(v bool) *AccountAccessTokenResponse {
	s.PinSetup = &v
	return s
}

func (s *AccountAccessTokenResponse) SetRefreshToken(v string) *AccountAccessTokenResponse {
	s.RefreshToken = &v
	return s
}

func (s *AccountAccessTokenResponse) SetRole(v string) *AccountAccessTokenResponse {
	s.Role = &v
	return s
}

func (s *AccountAccessTokenResponse) SetState(v string) *AccountAccessTokenResponse {
	s.State = &v
	return s
}

func (s *AccountAccessTokenResponse) SetStatus(v string) *AccountAccessTokenResponse {
	s.Status = &v
	return s
}

func (s *AccountAccessTokenResponse) SetTokenType(v string) *AccountAccessTokenResponse {
	s.TokenType = &v
	return s
}

func (s *AccountAccessTokenResponse) SetUserData(v map[string]*string) *AccountAccessTokenResponse {
	s.UserData = v
	return s
}

func (s *AccountAccessTokenResponse) SetUserId(v string) *AccountAccessTokenResponse {
	s.UserId = &v
	return s
}

func (s *AccountAccessTokenResponse) SetUserName(v string) *AccountAccessTokenResponse {
	s.UserName = &v
	return s
}

type AccountLinkInfo struct {
	AuthenticationType *string `json:"authentication_type,omitempty" xml:"authentication_type,omitempty"`
	CreatedAt          *int64  `json:"created_at,omitempty" xml:"created_at,omitempty"`
	DisplayName        *string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	DomainId           *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	Extra              *string `json:"extra,omitempty" xml:"extra,omitempty"`
	Identity           *string `json:"identity,omitempty" xml:"identity,omitempty"`
	LastLoginTime      *int64  `json:"last_login_time,omitempty" xml:"last_login_time,omitempty"`
	Status             *string `json:"status,omitempty" xml:"status,omitempty"`
	UserId             *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s AccountLinkInfo) String() string {
	return tea.Prettify(s)
}

func (s AccountLinkInfo) GoString() string {
	return s.String()
}

func (s *AccountLinkInfo) SetAuthenticationType(v string) *AccountLinkInfo {
	s.AuthenticationType = &v
	return s
}

func (s *AccountLinkInfo) SetCreatedAt(v int64) *AccountLinkInfo {
	s.CreatedAt = &v
	return s
}

func (s *AccountLinkInfo) SetDisplayName(v string) *AccountLinkInfo {
	s.DisplayName = &v
	return s
}

func (s *AccountLinkInfo) SetDomainId(v string) *AccountLinkInfo {
	s.DomainId = &v
	return s
}

func (s *AccountLinkInfo) SetExtra(v string) *AccountLinkInfo {
	s.Extra = &v
	return s
}

func (s *AccountLinkInfo) SetIdentity(v string) *AccountLinkInfo {
	s.Identity = &v
	return s
}

func (s *AccountLinkInfo) SetLastLoginTime(v int64) *AccountLinkInfo {
	s.LastLoginTime = &v
	return s
}

func (s *AccountLinkInfo) SetStatus(v string) *AccountLinkInfo {
	s.Status = &v
	return s
}

func (s *AccountLinkInfo) SetUserId(v string) *AccountLinkInfo {
	s.UserId = &v
	return s
}

type Activity struct {
	ActivityId         *string                  `json:"activity_id,omitempty" xml:"activity_id,omitempty"`
	Device             *string                  `json:"device,omitempty" xml:"device,omitempty"`
	DriveId            *string                  `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	EventType          *int32                   `json:"event_type,omitempty" xml:"event_type,omitempty"`
	LatestEventTime    *string                  `json:"latest_event_time,omitempty" xml:"latest_event_time,omitempty"`
	ResourceCategory   *int32                   `json:"resource_category,omitempty" xml:"resource_category,omitempty"`
	ResourceList       []map[string]interface{} `json:"resource_list,omitempty" xml:"resource_list,omitempty" type:"Repeated"`
	TotalResourceCount *int64                   `json:"total_resource_count,omitempty" xml:"total_resource_count,omitempty"`
	UserId             *string                  `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s Activity) String() string {
	return tea.Prettify(s)
}

func (s Activity) GoString() string {
	return s.String()
}

func (s *Activity) SetActivityId(v string) *Activity {
	s.ActivityId = &v
	return s
}

func (s *Activity) SetDevice(v string) *Activity {
	s.Device = &v
	return s
}

func (s *Activity) SetDriveId(v string) *Activity {
	s.DriveId = &v
	return s
}

func (s *Activity) SetEventType(v int32) *Activity {
	s.EventType = &v
	return s
}

func (s *Activity) SetLatestEventTime(v string) *Activity {
	s.LatestEventTime = &v
	return s
}

func (s *Activity) SetResourceCategory(v int32) *Activity {
	s.ResourceCategory = &v
	return s
}

func (s *Activity) SetResourceList(v []map[string]interface{}) *Activity {
	s.ResourceList = v
	return s
}

func (s *Activity) SetTotalResourceCount(v int64) *Activity {
	s.TotalResourceCount = &v
	return s
}

func (s *Activity) SetUserId(v string) *Activity {
	s.UserId = &v
	return s
}

type AddStoryFile struct {
	ErrorCode    *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	ErrorMessage *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88dd06e49d9c0a14411ebae606f70edd9a59
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s AddStoryFile) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFile) GoString() string {
	return s.String()
}

func (s *AddStoryFile) SetErrorCode(v string) *AddStoryFile {
	s.ErrorCode = &v
	return s
}

func (s *AddStoryFile) SetErrorMessage(v string) *AddStoryFile {
	s.ErrorMessage = &v
	return s
}

func (s *AddStoryFile) SetFileId(v string) *AddStoryFile {
	s.FileId = &v
	return s
}

func (s *AddStoryFile) SetRevisionId(v string) *AddStoryFile {
	s.RevisionId = &v
	return s
}

type Address struct {
	// example:
	//
	// 杭州市
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// example:
	//
	// 中国
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// example:
	//
	// 余杭区
	District *string `json:"district,omitempty" xml:"district,omitempty"`
	// example:
	//
	// 浙江省
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	// example:
	//
	// 文一西路
	Township *string `json:"township,omitempty" xml:"township,omitempty"`
}

func (s Address) String() string {
	return tea.Prettify(s)
}

func (s Address) GoString() string {
	return s.String()
}

func (s *Address) SetCity(v string) *Address {
	s.City = &v
	return s
}

func (s *Address) SetCountry(v string) *Address {
	s.Country = &v
	return s
}

func (s *Address) SetDistrict(v string) *Address {
	s.District = &v
	return s
}

func (s *Address) SetProvince(v string) *Address {
	s.Province = &v
	return s
}

func (s *Address) SetTownship(v string) *Address {
	s.Township = &v
	return s
}

type AddressGroup struct {
	AddressDetail *Address `json:"address_detail,omitempty" xml:"address_detail,omitempty"`
	Count         *int64   `json:"count,omitempty" xml:"count,omitempty"`
	CoverFileId   *string  `json:"cover_file_id,omitempty" xml:"cover_file_id,omitempty"`
	CoverUrl      *string  `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	Location      *string  `json:"location,omitempty" xml:"location,omitempty"`
	Name          *string  `json:"name,omitempty" xml:"name,omitempty"`
}

func (s AddressGroup) String() string {
	return tea.Prettify(s)
}

func (s AddressGroup) GoString() string {
	return s.String()
}

func (s *AddressGroup) SetAddressDetail(v *Address) *AddressGroup {
	s.AddressDetail = v
	return s
}

func (s *AddressGroup) SetCount(v int64) *AddressGroup {
	s.Count = &v
	return s
}

func (s *AddressGroup) SetCoverFileId(v string) *AddressGroup {
	s.CoverFileId = &v
	return s
}

func (s *AddressGroup) SetCoverUrl(v string) *AddressGroup {
	s.CoverUrl = &v
	return s
}

func (s *AddressGroup) SetLocation(v string) *AddressGroup {
	s.Location = &v
	return s
}

func (s *AddressGroup) SetName(v string) *AddressGroup {
	s.Name = &v
	return s
}

type Aggregation struct {
	Field     []byte               `json:"field,omitempty" xml:"field,omitempty"`
	Groups    []*AggregationsGroup `json:"groups,omitempty" xml:"groups,omitempty" type:"Repeated"`
	Operation []byte               `json:"operation,omitempty" xml:"operation,omitempty"`
	Value     *float64             `json:"value,omitempty" xml:"value,omitempty"`
}

func (s Aggregation) String() string {
	return tea.Prettify(s)
}

func (s Aggregation) GoString() string {
	return s.String()
}

func (s *Aggregation) SetField(v []byte) *Aggregation {
	s.Field = v
	return s
}

func (s *Aggregation) SetGroups(v []*AggregationsGroup) *Aggregation {
	s.Groups = v
	return s
}

func (s *Aggregation) SetOperation(v []byte) *Aggregation {
	s.Operation = v
	return s
}

func (s *Aggregation) SetValue(v float64) *Aggregation {
	s.Value = &v
	return s
}

type AggregationsGroup struct {
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	Value []byte `json:"value,omitempty" xml:"value,omitempty"`
}

func (s AggregationsGroup) String() string {
	return tea.Prettify(s)
}

func (s AggregationsGroup) GoString() string {
	return s.String()
}

func (s *AggregationsGroup) SetCount(v int64) *AggregationsGroup {
	s.Count = &v
	return s
}

func (s *AggregationsGroup) SetValue(v []byte) *AggregationsGroup {
	s.Value = v
	return s
}

type Album struct {
	AlbumId         *string            `json:"album_id,omitempty" xml:"album_id,omitempty"`
	BaseFaceFile    *File              `json:"base_face_file,omitempty" xml:"base_face_file,omitempty"`
	BaseFaceGroupId *string            `json:"base_face_group_id,omitempty" xml:"base_face_group_id,omitempty"`
	CoverFile       *File              `json:"cover_file,omitempty" xml:"cover_file,omitempty"`
	CreatedAt       *string            `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description     *string            `json:"description,omitempty" xml:"description,omitempty"`
	FileCount       *int64             `json:"file_count,omitempty" xml:"file_count,omitempty"`
	Name            *string            `json:"name,omitempty" xml:"name,omitempty"`
	Owner           *string            `json:"owner,omitempty" xml:"owner,omitempty"`
	UpdatedAt       *string            `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UserTags        map[string]*string `json:"user_tags,omitempty" xml:"user_tags,omitempty"`
}

func (s Album) String() string {
	return tea.Prettify(s)
}

func (s Album) GoString() string {
	return s.String()
}

func (s *Album) SetAlbumId(v string) *Album {
	s.AlbumId = &v
	return s
}

func (s *Album) SetBaseFaceFile(v *File) *Album {
	s.BaseFaceFile = v
	return s
}

func (s *Album) SetBaseFaceGroupId(v string) *Album {
	s.BaseFaceGroupId = &v
	return s
}

func (s *Album) SetCoverFile(v *File) *Album {
	s.CoverFile = v
	return s
}

func (s *Album) SetCreatedAt(v string) *Album {
	s.CreatedAt = &v
	return s
}

func (s *Album) SetDescription(v string) *Album {
	s.Description = &v
	return s
}

func (s *Album) SetFileCount(v int64) *Album {
	s.FileCount = &v
	return s
}

func (s *Album) SetName(v string) *Album {
	s.Name = &v
	return s
}

func (s *Album) SetOwner(v string) *Album {
	s.Owner = &v
	return s
}

func (s *Album) SetUpdatedAt(v string) *Album {
	s.UpdatedAt = &v
	return s
}

func (s *Album) SetUserTags(v map[string]*string) *Album {
	s.UserTags = v
	return s
}

type AlbumFile struct {
	AlbumId            *string                `json:"album_id,omitempty" xml:"album_id,omitempty"`
	Category           *string                `json:"category,omitempty" xml:"category,omitempty"`
	ContentHash        *string                `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName    *string                `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	ContentType        *string                `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Crc64Hash          *string                `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	CreatedAt          *string                `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description        *string                `json:"description,omitempty" xml:"description,omitempty"`
	DomainId           *string                `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DownloadUrl        *string                `json:"download_url,omitempty" xml:"download_url,omitempty"`
	DriveId            *string                `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	ExFieldsInfo       map[string]interface{} `json:"ex_fields_info,omitempty" xml:"ex_fields_info,omitempty"`
	FileExtension      *string                `json:"file_extension,omitempty" xml:"file_extension,omitempty"`
	FileId             *string                `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Hidden             *bool                  `json:"hidden,omitempty" xml:"hidden,omitempty"`
	ImageMediaMetadata *ImageMediaMetadata    `json:"image_media_metadata,omitempty" xml:"image_media_metadata,omitempty"`
	InvestigationInfo  *InvestigationInfo     `json:"investigation_info,omitempty" xml:"investigation_info,omitempty"`
	JoinedAt           *int64                 `json:"joined_at,omitempty" xml:"joined_at,omitempty"`
	Labels             []*string              `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	LocalCreatedAt     *string                `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	LocalModifiedAt    *string                `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	MimeType           *string                `json:"mime_type,omitempty" xml:"mime_type,omitempty"`
	Name               *string                `json:"name,omitempty" xml:"name,omitempty"`
	ObjectUri          *string                `json:"object_uri,omitempty" xml:"object_uri,omitempty"`
	ParentFileId       *string                `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	RevisionId         *string                `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	Size               *int64                 `json:"size,omitempty" xml:"size,omitempty"`
	Starred            *bool                  `json:"starred,omitempty" xml:"starred,omitempty"`
	Status             *string                `json:"status,omitempty" xml:"status,omitempty"`
	Thumbnail          *string                `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	ThumbnailUrls      map[string]*string     `json:"thumbnail_urls,omitempty" xml:"thumbnail_urls,omitempty"`
	TranshedAt         *string                `json:"transhed_at,omitempty" xml:"transhed_at,omitempty"`
	Type               *string                `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt          *string                `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UploadId           *string                `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	UserMeta           *string                `json:"user_meta,omitempty" xml:"user_meta,omitempty"`
}

func (s AlbumFile) String() string {
	return tea.Prettify(s)
}

func (s AlbumFile) GoString() string {
	return s.String()
}

func (s *AlbumFile) SetAlbumId(v string) *AlbumFile {
	s.AlbumId = &v
	return s
}

func (s *AlbumFile) SetCategory(v string) *AlbumFile {
	s.Category = &v
	return s
}

func (s *AlbumFile) SetContentHash(v string) *AlbumFile {
	s.ContentHash = &v
	return s
}

func (s *AlbumFile) SetContentHashName(v string) *AlbumFile {
	s.ContentHashName = &v
	return s
}

func (s *AlbumFile) SetContentType(v string) *AlbumFile {
	s.ContentType = &v
	return s
}

func (s *AlbumFile) SetCrc64Hash(v string) *AlbumFile {
	s.Crc64Hash = &v
	return s
}

func (s *AlbumFile) SetCreatedAt(v string) *AlbumFile {
	s.CreatedAt = &v
	return s
}

func (s *AlbumFile) SetDescription(v string) *AlbumFile {
	s.Description = &v
	return s
}

func (s *AlbumFile) SetDomainId(v string) *AlbumFile {
	s.DomainId = &v
	return s
}

func (s *AlbumFile) SetDownloadUrl(v string) *AlbumFile {
	s.DownloadUrl = &v
	return s
}

func (s *AlbumFile) SetDriveId(v string) *AlbumFile {
	s.DriveId = &v
	return s
}

func (s *AlbumFile) SetExFieldsInfo(v map[string]interface{}) *AlbumFile {
	s.ExFieldsInfo = v
	return s
}

func (s *AlbumFile) SetFileExtension(v string) *AlbumFile {
	s.FileExtension = &v
	return s
}

func (s *AlbumFile) SetFileId(v string) *AlbumFile {
	s.FileId = &v
	return s
}

func (s *AlbumFile) SetHidden(v bool) *AlbumFile {
	s.Hidden = &v
	return s
}

func (s *AlbumFile) SetImageMediaMetadata(v *ImageMediaMetadata) *AlbumFile {
	s.ImageMediaMetadata = v
	return s
}

func (s *AlbumFile) SetInvestigationInfo(v *InvestigationInfo) *AlbumFile {
	s.InvestigationInfo = v
	return s
}

func (s *AlbumFile) SetJoinedAt(v int64) *AlbumFile {
	s.JoinedAt = &v
	return s
}

func (s *AlbumFile) SetLabels(v []*string) *AlbumFile {
	s.Labels = v
	return s
}

func (s *AlbumFile) SetLocalCreatedAt(v string) *AlbumFile {
	s.LocalCreatedAt = &v
	return s
}

func (s *AlbumFile) SetLocalModifiedAt(v string) *AlbumFile {
	s.LocalModifiedAt = &v
	return s
}

func (s *AlbumFile) SetMimeType(v string) *AlbumFile {
	s.MimeType = &v
	return s
}

func (s *AlbumFile) SetName(v string) *AlbumFile {
	s.Name = &v
	return s
}

func (s *AlbumFile) SetObjectUri(v string) *AlbumFile {
	s.ObjectUri = &v
	return s
}

func (s *AlbumFile) SetParentFileId(v string) *AlbumFile {
	s.ParentFileId = &v
	return s
}

func (s *AlbumFile) SetRevisionId(v string) *AlbumFile {
	s.RevisionId = &v
	return s
}

func (s *AlbumFile) SetSize(v int64) *AlbumFile {
	s.Size = &v
	return s
}

func (s *AlbumFile) SetStarred(v bool) *AlbumFile {
	s.Starred = &v
	return s
}

func (s *AlbumFile) SetStatus(v string) *AlbumFile {
	s.Status = &v
	return s
}

func (s *AlbumFile) SetThumbnail(v string) *AlbumFile {
	s.Thumbnail = &v
	return s
}

func (s *AlbumFile) SetThumbnailUrls(v map[string]*string) *AlbumFile {
	s.ThumbnailUrls = v
	return s
}

func (s *AlbumFile) SetTranshedAt(v string) *AlbumFile {
	s.TranshedAt = &v
	return s
}

func (s *AlbumFile) SetType(v string) *AlbumFile {
	s.Type = &v
	return s
}

func (s *AlbumFile) SetUpdatedAt(v string) *AlbumFile {
	s.UpdatedAt = &v
	return s
}

func (s *AlbumFile) SetUploadId(v string) *AlbumFile {
	s.UploadId = &v
	return s
}

func (s *AlbumFile) SetUserMeta(v string) *AlbumFile {
	s.UserMeta = &v
	return s
}

type App struct {
	AppId       *string   `json:"app_id,omitempty" xml:"app_id,omitempty"`
	AppName     *string   `json:"app_name,omitempty" xml:"app_name,omitempty"`
	AppSecret   *string   `json:"app_secret,omitempty" xml:"app_secret,omitempty"`
	CreatedAt   *string   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description *string   `json:"description,omitempty" xml:"description,omitempty"`
	Logo        *string   `json:"logo,omitempty" xml:"logo,omitempty"`
	Provider    *string   `json:"provider,omitempty" xml:"provider,omitempty"`
	RedirectUri *string   `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	Scope       []*string `json:"scope,omitempty" xml:"scope,omitempty" type:"Repeated"`
	Stage       *string   `json:"stage,omitempty" xml:"stage,omitempty"`
	Type        *string   `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt   *string   `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s App) String() string {
	return tea.Prettify(s)
}

func (s App) GoString() string {
	return s.String()
}

func (s *App) SetAppId(v string) *App {
	s.AppId = &v
	return s
}

func (s *App) SetAppName(v string) *App {
	s.AppName = &v
	return s
}

func (s *App) SetAppSecret(v string) *App {
	s.AppSecret = &v
	return s
}

func (s *App) SetCreatedAt(v string) *App {
	s.CreatedAt = &v
	return s
}

func (s *App) SetDescription(v string) *App {
	s.Description = &v
	return s
}

func (s *App) SetLogo(v string) *App {
	s.Logo = &v
	return s
}

func (s *App) SetProvider(v string) *App {
	s.Provider = &v
	return s
}

func (s *App) SetRedirectUri(v string) *App {
	s.RedirectUri = &v
	return s
}

func (s *App) SetScope(v []*string) *App {
	s.Scope = v
	return s
}

func (s *App) SetStage(v string) *App {
	s.Stage = &v
	return s
}

func (s *App) SetType(v string) *App {
	s.Type = &v
	return s
}

func (s *App) SetUpdatedAt(v string) *App {
	s.UpdatedAt = &v
	return s
}

type AppAccessStrategy struct {
	Effect          *string   `json:"effect,omitempty" xml:"effect,omitempty"`
	ExceptAppIdList []*string `json:"except_app_id_list,omitempty" xml:"except_app_id_list,omitempty" type:"Repeated"`
}

func (s AppAccessStrategy) String() string {
	return tea.Prettify(s)
}

func (s AppAccessStrategy) GoString() string {
	return s.String()
}

func (s *AppAccessStrategy) SetEffect(v string) *AppAccessStrategy {
	s.Effect = &v
	return s
}

func (s *AppAccessStrategy) SetExceptAppIdList(v []*string) *AppAccessStrategy {
	s.ExceptAppIdList = v
	return s
}

type ArchiveFilesConfigResponse struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s ArchiveFilesConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s ArchiveFilesConfigResponse) GoString() string {
	return s.String()
}

func (s *ArchiveFilesConfigResponse) SetEnabled(v bool) *ArchiveFilesConfigResponse {
	s.Enabled = &v
	return s
}

func (s *ArchiveFilesConfigResponse) SetVersion(v string) *ArchiveFilesConfigResponse {
	s.Version = &v
	return s
}

type AuditLog struct {
	ActedAt        *string         `json:"acted_at,omitempty" xml:"acted_at,omitempty"`
	ActionCategory *string         `json:"action_category,omitempty" xml:"action_category,omitempty"`
	ActionType     *string         `json:"action_type,omitempty" xml:"action_type,omitempty"`
	ActorId        *string         `json:"actor_id,omitempty" xml:"actor_id,omitempty"`
	ActorName      *string         `json:"actor_name,omitempty" xml:"actor_name,omitempty"`
	ActorType      *string         `json:"actor_type,omitempty" xml:"actor_type,omitempty"`
	ClientDevice   *string         `json:"client_device,omitempty" xml:"client_device,omitempty"`
	ClientIp       *string         `json:"client_ip,omitempty" xml:"client_ip,omitempty"`
	ClientType     *string         `json:"client_type,omitempty" xml:"client_type,omitempty"`
	ClientVersion  *string         `json:"client_version,omitempty" xml:"client_version,omitempty"`
	Detail         *AuditLogDetail `json:"detail,omitempty" xml:"detail,omitempty"`
	FilePathType   *string         `json:"file_path_type,omitempty" xml:"file_path_type,omitempty"`
	LogId          *string         `json:"log_id,omitempty" xml:"log_id,omitempty"`
	ObjectId       *string         `json:"object_id,omitempty" xml:"object_id,omitempty"`
	ObjectName     *string         `json:"object_name,omitempty" xml:"object_name,omitempty"`
}

func (s AuditLog) String() string {
	return tea.Prettify(s)
}

func (s AuditLog) GoString() string {
	return s.String()
}

func (s *AuditLog) SetActedAt(v string) *AuditLog {
	s.ActedAt = &v
	return s
}

func (s *AuditLog) SetActionCategory(v string) *AuditLog {
	s.ActionCategory = &v
	return s
}

func (s *AuditLog) SetActionType(v string) *AuditLog {
	s.ActionType = &v
	return s
}

func (s *AuditLog) SetActorId(v string) *AuditLog {
	s.ActorId = &v
	return s
}

func (s *AuditLog) SetActorName(v string) *AuditLog {
	s.ActorName = &v
	return s
}

func (s *AuditLog) SetActorType(v string) *AuditLog {
	s.ActorType = &v
	return s
}

func (s *AuditLog) SetClientDevice(v string) *AuditLog {
	s.ClientDevice = &v
	return s
}

func (s *AuditLog) SetClientIp(v string) *AuditLog {
	s.ClientIp = &v
	return s
}

func (s *AuditLog) SetClientType(v string) *AuditLog {
	s.ClientType = &v
	return s
}

func (s *AuditLog) SetClientVersion(v string) *AuditLog {
	s.ClientVersion = &v
	return s
}

func (s *AuditLog) SetDetail(v *AuditLogDetail) *AuditLog {
	s.Detail = v
	return s
}

func (s *AuditLog) SetFilePathType(v string) *AuditLog {
	s.FilePathType = &v
	return s
}

func (s *AuditLog) SetLogId(v string) *AuditLog {
	s.LogId = &v
	return s
}

func (s *AuditLog) SetObjectId(v string) *AuditLog {
	s.ObjectId = &v
	return s
}

func (s *AuditLog) SetObjectName(v string) *AuditLog {
	s.ObjectName = &v
	return s
}

type AuditLogDetail struct {
	DriveLogDetail *DriveLogDetail `json:"drive_log_detail,omitempty" xml:"drive_log_detail,omitempty"`
	FileLogDetail  *FileLogDetail  `json:"file_log_detail,omitempty" xml:"file_log_detail,omitempty"`
	UserLogDetail  *UserLogDetail  `json:"user_log_detail,omitempty" xml:"user_log_detail,omitempty"`
}

func (s AuditLogDetail) String() string {
	return tea.Prettify(s)
}

func (s AuditLogDetail) GoString() string {
	return s.String()
}

func (s *AuditLogDetail) SetDriveLogDetail(v *DriveLogDetail) *AuditLogDetail {
	s.DriveLogDetail = v
	return s
}

func (s *AuditLogDetail) SetFileLogDetail(v *FileLogDetail) *AuditLogDetail {
	s.FileLogDetail = v
	return s
}

func (s *AuditLogDetail) SetUserLogDetail(v *UserLogDetail) *AuditLogDetail {
	s.UserLogDetail = v
	return s
}

type BaseAlbumFileOperationResult struct {
	ErrorCode    *string         `json:"error_code,omitempty" xml:"error_code,omitempty"`
	ErrorMessage *string         `json:"error_message,omitempty" xml:"error_message,omitempty"`
	File         *CommonFileItem `json:"file,omitempty" xml:"file,omitempty"`
	IsSucceed    *bool           `json:"is_succeed,omitempty" xml:"is_succeed,omitempty"`
}

func (s BaseAlbumFileOperationResult) String() string {
	return tea.Prettify(s)
}

func (s BaseAlbumFileOperationResult) GoString() string {
	return s.String()
}

func (s *BaseAlbumFileOperationResult) SetErrorCode(v string) *BaseAlbumFileOperationResult {
	s.ErrorCode = &v
	return s
}

func (s *BaseAlbumFileOperationResult) SetErrorMessage(v string) *BaseAlbumFileOperationResult {
	s.ErrorMessage = &v
	return s
}

func (s *BaseAlbumFileOperationResult) SetFile(v *CommonFileItem) *BaseAlbumFileOperationResult {
	s.File = v
	return s
}

func (s *BaseAlbumFileOperationResult) SetIsSucceed(v bool) *BaseAlbumFileOperationResult {
	s.IsSucceed = &v
	return s
}

type BaseAssignmentResponse struct {
	AssociatedRoleTagId *string   `json:"associated_role_tag_id,omitempty" xml:"associated_role_tag_id,omitempty"`
	CreatedAt           *string   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator             *string   `json:"creator,omitempty" xml:"creator,omitempty"`
	DisinheritSubGroup  *bool     `json:"disinherit_sub_group,omitempty" xml:"disinherit_sub_group,omitempty"`
	DomainId            *string   `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	Identity            *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	ManageResourceId    *string   `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	ManageResourceType  *string   `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	RoleId              *string   `json:"role_id,omitempty" xml:"role_id,omitempty"`
	UpdatedAt           *string   `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s BaseAssignmentResponse) String() string {
	return tea.Prettify(s)
}

func (s BaseAssignmentResponse) GoString() string {
	return s.String()
}

func (s *BaseAssignmentResponse) SetAssociatedRoleTagId(v string) *BaseAssignmentResponse {
	s.AssociatedRoleTagId = &v
	return s
}

func (s *BaseAssignmentResponse) SetCreatedAt(v string) *BaseAssignmentResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseAssignmentResponse) SetCreator(v string) *BaseAssignmentResponse {
	s.Creator = &v
	return s
}

func (s *BaseAssignmentResponse) SetDisinheritSubGroup(v bool) *BaseAssignmentResponse {
	s.DisinheritSubGroup = &v
	return s
}

func (s *BaseAssignmentResponse) SetDomainId(v string) *BaseAssignmentResponse {
	s.DomainId = &v
	return s
}

func (s *BaseAssignmentResponse) SetIdentity(v *Identity) *BaseAssignmentResponse {
	s.Identity = v
	return s
}

func (s *BaseAssignmentResponse) SetManageResourceId(v string) *BaseAssignmentResponse {
	s.ManageResourceId = &v
	return s
}

func (s *BaseAssignmentResponse) SetManageResourceType(v string) *BaseAssignmentResponse {
	s.ManageResourceType = &v
	return s
}

func (s *BaseAssignmentResponse) SetRoleId(v string) *BaseAssignmentResponse {
	s.RoleId = &v
	return s
}

func (s *BaseAssignmentResponse) SetUpdatedAt(v string) *BaseAssignmentResponse {
	s.UpdatedAt = &v
	return s
}

type BaseDomainResponse struct {
	CreatedAt                  *string            `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description                *string            `json:"description,omitempty" xml:"description,omitempty"`
	DomainId                   *string            `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DomainName                 *string            `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	InitDriveEnable            *bool              `json:"init_drive_enable,omitempty" xml:"init_drive_enable,omitempty"`
	InitDriveSize              *int64             `json:"init_drive_size,omitempty" xml:"init_drive_size,omitempty"`
	ParentDomainId             *string            `json:"parent_domain_id,omitempty" xml:"parent_domain_id,omitempty"`
	PublishedAppAccessStrategy *AppAccessStrategy `json:"published_app_access_strategy,omitempty" xml:"published_app_access_strategy,omitempty"`
	ShareLinkEnabled           *bool              `json:"share_link_enabled,omitempty" xml:"share_link_enabled,omitempty"`
	SizeQuota                  *int64             `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	SizeQuotaUsed              *int64             `json:"size_quota_used,omitempty" xml:"size_quota_used,omitempty"`
	Status                     *int64             `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt                  *string            `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UsedSize                   *int64             `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s BaseDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s BaseDomainResponse) GoString() string {
	return s.String()
}

func (s *BaseDomainResponse) SetCreatedAt(v string) *BaseDomainResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseDomainResponse) SetDescription(v string) *BaseDomainResponse {
	s.Description = &v
	return s
}

func (s *BaseDomainResponse) SetDomainId(v string) *BaseDomainResponse {
	s.DomainId = &v
	return s
}

func (s *BaseDomainResponse) SetDomainName(v string) *BaseDomainResponse {
	s.DomainName = &v
	return s
}

func (s *BaseDomainResponse) SetInitDriveEnable(v bool) *BaseDomainResponse {
	s.InitDriveEnable = &v
	return s
}

func (s *BaseDomainResponse) SetInitDriveSize(v int64) *BaseDomainResponse {
	s.InitDriveSize = &v
	return s
}

func (s *BaseDomainResponse) SetParentDomainId(v string) *BaseDomainResponse {
	s.ParentDomainId = &v
	return s
}

func (s *BaseDomainResponse) SetPublishedAppAccessStrategy(v *AppAccessStrategy) *BaseDomainResponse {
	s.PublishedAppAccessStrategy = v
	return s
}

func (s *BaseDomainResponse) SetShareLinkEnabled(v bool) *BaseDomainResponse {
	s.ShareLinkEnabled = &v
	return s
}

func (s *BaseDomainResponse) SetSizeQuota(v int64) *BaseDomainResponse {
	s.SizeQuota = &v
	return s
}

func (s *BaseDomainResponse) SetSizeQuotaUsed(v int64) *BaseDomainResponse {
	s.SizeQuotaUsed = &v
	return s
}

func (s *BaseDomainResponse) SetStatus(v int64) *BaseDomainResponse {
	s.Status = &v
	return s
}

func (s *BaseDomainResponse) SetUpdatedAt(v string) *BaseDomainResponse {
	s.UpdatedAt = &v
	return s
}

func (s *BaseDomainResponse) SetUsedSize(v int64) *BaseDomainResponse {
	s.UsedSize = &v
	return s
}

type BaseDriveResponse struct {
	ActionList []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	Category   *string   `json:"category,omitempty" xml:"category,omitempty"`
	CreatedAt  *string   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// ccp
	Creator      *string `json:"creator,omitempty" xml:"creator,omitempty"`
	DeltaEnabled *bool   `json:"delta_enabled,omitempty" xml:"delta_enabled,omitempty"`
	// example:
	//
	// ccp team drive
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// hz999
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// example:
	//
	// 123
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// ccpdrive
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// example:
	//
	// normal
	DriveType         *string `json:"drive_type,omitempty" xml:"drive_type,omitempty"`
	EncryptDataAccess *bool   `json:"encrypt_data_access,omitempty" xml:"encrypt_data_access,omitempty"`
	EncryptMode       *string `json:"encrypt_mode,omitempty" xml:"encrypt_mode,omitempty"`
	IsHandover        *bool   `json:"is_handover,omitempty" xml:"is_handover,omitempty"`
	// example:
	//
	// ccp
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// example:
	//
	// user
	OwnerType    *string                  `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	PathStatus   *string                  `json:"path_status,omitempty" xml:"path_status,omitempty"`
	Permission   map[string]*IDPermission `json:"permission,omitempty" xml:"permission,omitempty"`
	RelativePath *string                  `json:"relative_path,omitempty" xml:"relative_path,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 123
	StoreId *string `json:"store_id,omitempty" xml:"store_id,omitempty"`
	// example:
	//
	// 102400
	TotalSize *int64  `json:"total_size,omitempty" xml:"total_size,omitempty"`
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// example:
	//
	// 1024
	UsedSize *int64 `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s BaseDriveResponse) String() string {
	return tea.Prettify(s)
}

func (s BaseDriveResponse) GoString() string {
	return s.String()
}

func (s *BaseDriveResponse) SetActionList(v []*string) *BaseDriveResponse {
	s.ActionList = v
	return s
}

func (s *BaseDriveResponse) SetCategory(v string) *BaseDriveResponse {
	s.Category = &v
	return s
}

func (s *BaseDriveResponse) SetCreatedAt(v string) *BaseDriveResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseDriveResponse) SetCreator(v string) *BaseDriveResponse {
	s.Creator = &v
	return s
}

func (s *BaseDriveResponse) SetDeltaEnabled(v bool) *BaseDriveResponse {
	s.DeltaEnabled = &v
	return s
}

func (s *BaseDriveResponse) SetDescription(v string) *BaseDriveResponse {
	s.Description = &v
	return s
}

func (s *BaseDriveResponse) SetDomainId(v string) *BaseDriveResponse {
	s.DomainId = &v
	return s
}

func (s *BaseDriveResponse) SetDriveId(v string) *BaseDriveResponse {
	s.DriveId = &v
	return s
}

func (s *BaseDriveResponse) SetDriveName(v string) *BaseDriveResponse {
	s.DriveName = &v
	return s
}

func (s *BaseDriveResponse) SetDriveType(v string) *BaseDriveResponse {
	s.DriveType = &v
	return s
}

func (s *BaseDriveResponse) SetEncryptDataAccess(v bool) *BaseDriveResponse {
	s.EncryptDataAccess = &v
	return s
}

func (s *BaseDriveResponse) SetEncryptMode(v string) *BaseDriveResponse {
	s.EncryptMode = &v
	return s
}

func (s *BaseDriveResponse) SetIsHandover(v bool) *BaseDriveResponse {
	s.IsHandover = &v
	return s
}

func (s *BaseDriveResponse) SetOwner(v string) *BaseDriveResponse {
	s.Owner = &v
	return s
}

func (s *BaseDriveResponse) SetOwnerType(v string) *BaseDriveResponse {
	s.OwnerType = &v
	return s
}

func (s *BaseDriveResponse) SetPathStatus(v string) *BaseDriveResponse {
	s.PathStatus = &v
	return s
}

func (s *BaseDriveResponse) SetPermission(v map[string]*IDPermission) *BaseDriveResponse {
	s.Permission = v
	return s
}

func (s *BaseDriveResponse) SetRelativePath(v string) *BaseDriveResponse {
	s.RelativePath = &v
	return s
}

func (s *BaseDriveResponse) SetStatus(v string) *BaseDriveResponse {
	s.Status = &v
	return s
}

func (s *BaseDriveResponse) SetStoreId(v string) *BaseDriveResponse {
	s.StoreId = &v
	return s
}

func (s *BaseDriveResponse) SetTotalSize(v int64) *BaseDriveResponse {
	s.TotalSize = &v
	return s
}

func (s *BaseDriveResponse) SetUpdatedAt(v string) *BaseDriveResponse {
	s.UpdatedAt = &v
	return s
}

func (s *BaseDriveResponse) SetUsedSize(v int64) *BaseDriveResponse {
	s.UsedSize = &v
	return s
}

type BaseFileListInheritPermissionResponse struct {
	FileId *string               `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Member *FilePermissionMember `json:"member,omitempty" xml:"member,omitempty"`
}

func (s BaseFileListInheritPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s BaseFileListInheritPermissionResponse) GoString() string {
	return s.String()
}

func (s *BaseFileListInheritPermissionResponse) SetFileId(v string) *BaseFileListInheritPermissionResponse {
	s.FileId = &v
	return s
}

func (s *BaseFileListInheritPermissionResponse) SetMember(v *FilePermissionMember) *BaseFileListInheritPermissionResponse {
	s.Member = v
	return s
}

type BaseFileUserPermissionResponse struct {
	CanAccess          *bool   `json:"can_access,omitempty" xml:"can_access,omitempty"`
	CreatedAt          *int64  `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator            *string `json:"creator,omitempty" xml:"creator,omitempty"`
	DisinheritSubGroup *bool   `json:"disinherit_sub_group,omitempty" xml:"disinherit_sub_group,omitempty"`
	// example:
	//
	// bj23
	DomainId     *string   `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId      *string   `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	ExpireTime   *int64    `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	FileFullPath *string   `json:"file_full_path,omitempty" xml:"file_full_path,omitempty"`
	FileId       *string   `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Identity     *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	RoleId       *string   `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s BaseFileUserPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s BaseFileUserPermissionResponse) GoString() string {
	return s.String()
}

func (s *BaseFileUserPermissionResponse) SetCanAccess(v bool) *BaseFileUserPermissionResponse {
	s.CanAccess = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetCreatedAt(v int64) *BaseFileUserPermissionResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetCreator(v string) *BaseFileUserPermissionResponse {
	s.Creator = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetDisinheritSubGroup(v bool) *BaseFileUserPermissionResponse {
	s.DisinheritSubGroup = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetDomainId(v string) *BaseFileUserPermissionResponse {
	s.DomainId = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetDriveId(v string) *BaseFileUserPermissionResponse {
	s.DriveId = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetExpireTime(v int64) *BaseFileUserPermissionResponse {
	s.ExpireTime = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetFileFullPath(v string) *BaseFileUserPermissionResponse {
	s.FileFullPath = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetFileId(v string) *BaseFileUserPermissionResponse {
	s.FileId = &v
	return s
}

func (s *BaseFileUserPermissionResponse) SetIdentity(v *Identity) *BaseFileUserPermissionResponse {
	s.Identity = v
	return s
}

func (s *BaseFileUserPermissionResponse) SetRoleId(v string) *BaseFileUserPermissionResponse {
	s.RoleId = &v
	return s
}

type BaseGroupResponse struct {
	// example:
	//
	// 111111
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// system
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// desc-111
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// bj123
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// example:
	//
	// b38b5681bd964950ad8bc0f8ea504793
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// example:
	//
	// name-111
	GroupName  *string                  `json:"group_name,omitempty" xml:"group_name,omitempty"`
	IsSync     *bool                    `json:"is_sync,omitempty" xml:"is_sync,omitempty"`
	Permission map[string]*IDPermission `json:"permission,omitempty" xml:"permission,omitempty"`
	// example:
	//
	// 111111
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s BaseGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s BaseGroupResponse) GoString() string {
	return s.String()
}

func (s *BaseGroupResponse) SetCreatedAt(v int64) *BaseGroupResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseGroupResponse) SetCreator(v string) *BaseGroupResponse {
	s.Creator = &v
	return s
}

func (s *BaseGroupResponse) SetDescription(v string) *BaseGroupResponse {
	s.Description = &v
	return s
}

func (s *BaseGroupResponse) SetDomainId(v string) *BaseGroupResponse {
	s.DomainId = &v
	return s
}

func (s *BaseGroupResponse) SetGroupId(v string) *BaseGroupResponse {
	s.GroupId = &v
	return s
}

func (s *BaseGroupResponse) SetGroupName(v string) *BaseGroupResponse {
	s.GroupName = &v
	return s
}

func (s *BaseGroupResponse) SetIsSync(v bool) *BaseGroupResponse {
	s.IsSync = &v
	return s
}

func (s *BaseGroupResponse) SetPermission(v map[string]*IDPermission) *BaseGroupResponse {
	s.Permission = v
	return s
}

func (s *BaseGroupResponse) SetUpdatedAt(v string) *BaseGroupResponse {
	s.UpdatedAt = &v
	return s
}

type BaseRoleMemberResponse struct {
	AssignmentList []*BaseAssignmentResponse `json:"assignment_list,omitempty" xml:"assignment_list,omitempty" type:"Repeated"`
	CreatedAt      *string                   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator        *string                   `json:"creator,omitempty" xml:"creator,omitempty"`
	DomainId       *string                   `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	Identity       *Identity                 `json:"identity,omitempty" xml:"identity,omitempty"`
	IdentityName   *string                   `json:"identity_name,omitempty" xml:"identity_name,omitempty"`
	IsAdmin        *bool                     `json:"is_admin,omitempty" xml:"is_admin,omitempty"`
	SubdomainId    *string                   `json:"subdomain_id,omitempty" xml:"subdomain_id,omitempty"`
}

func (s BaseRoleMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s BaseRoleMemberResponse) GoString() string {
	return s.String()
}

func (s *BaseRoleMemberResponse) SetAssignmentList(v []*BaseAssignmentResponse) *BaseRoleMemberResponse {
	s.AssignmentList = v
	return s
}

func (s *BaseRoleMemberResponse) SetCreatedAt(v string) *BaseRoleMemberResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseRoleMemberResponse) SetCreator(v string) *BaseRoleMemberResponse {
	s.Creator = &v
	return s
}

func (s *BaseRoleMemberResponse) SetDomainId(v string) *BaseRoleMemberResponse {
	s.DomainId = &v
	return s
}

func (s *BaseRoleMemberResponse) SetIdentity(v *Identity) *BaseRoleMemberResponse {
	s.Identity = v
	return s
}

func (s *BaseRoleMemberResponse) SetIdentityName(v string) *BaseRoleMemberResponse {
	s.IdentityName = &v
	return s
}

func (s *BaseRoleMemberResponse) SetIsAdmin(v bool) *BaseRoleMemberResponse {
	s.IsAdmin = &v
	return s
}

func (s *BaseRoleMemberResponse) SetSubdomainId(v string) *BaseRoleMemberResponse {
	s.SubdomainId = &v
	return s
}

type BaseUserResponse struct {
	// example:
	//
	// http://a.b.c/ccp.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// 1567407718386
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// system
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// example:
	//
	// 123
	DefaultDriveId           *string `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	DefaultLocation          *string `json:"default_location,omitempty" xml:"default_location,omitempty"`
	DenyChangePasswordBySelf *bool   `json:"deny_change_password_by_self,omitempty" xml:"deny_change_password_by_self,omitempty"`
	// example:
	//
	// ccp team user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// hz999
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// example:
	//
	// 123@ccp.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 0
	ExpiredAt                   *int64 `json:"expired_at,omitempty" xml:"expired_at,omitempty"`
	IsSync                      *bool  `json:"is_sync,omitempty" xml:"is_sync,omitempty"`
	LastLoginTime               *int64 `json:"last_login_time,omitempty" xml:"last_login_time,omitempty"`
	NeedChangePasswordNextLogin *bool  `json:"need_change_password_next_login,omitempty" xml:"need_change_password_next_login,omitempty"`
	// example:
	//
	// abc
	NickName   *string                  `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	PathStatus *string                  `json:"path_status,omitempty" xml:"path_status,omitempty"`
	Permission map[string]*IDPermission `json:"permission,omitempty" xml:"permission,omitempty"`
	// example:
	//
	// 13700000000
	Phone       *string `json:"phone,omitempty" xml:"phone,omitempty"`
	PhoneRegion *string `json:"phone_region,omitempty" xml:"phone_region,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1567407718386
	UpdatedAt *string                `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UserData  map[string]interface{} `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// example:
	//
	// ccpuserid
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// name
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s BaseUserResponse) String() string {
	return tea.Prettify(s)
}

func (s BaseUserResponse) GoString() string {
	return s.String()
}

func (s *BaseUserResponse) SetAvatar(v string) *BaseUserResponse {
	s.Avatar = &v
	return s
}

func (s *BaseUserResponse) SetCreatedAt(v string) *BaseUserResponse {
	s.CreatedAt = &v
	return s
}

func (s *BaseUserResponse) SetCreator(v string) *BaseUserResponse {
	s.Creator = &v
	return s
}

func (s *BaseUserResponse) SetDefaultDriveId(v string) *BaseUserResponse {
	s.DefaultDriveId = &v
	return s
}

func (s *BaseUserResponse) SetDefaultLocation(v string) *BaseUserResponse {
	s.DefaultLocation = &v
	return s
}

func (s *BaseUserResponse) SetDenyChangePasswordBySelf(v bool) *BaseUserResponse {
	s.DenyChangePasswordBySelf = &v
	return s
}

func (s *BaseUserResponse) SetDescription(v string) *BaseUserResponse {
	s.Description = &v
	return s
}

func (s *BaseUserResponse) SetDomainId(v string) *BaseUserResponse {
	s.DomainId = &v
	return s
}

func (s *BaseUserResponse) SetEmail(v string) *BaseUserResponse {
	s.Email = &v
	return s
}

func (s *BaseUserResponse) SetExpiredAt(v int64) *BaseUserResponse {
	s.ExpiredAt = &v
	return s
}

func (s *BaseUserResponse) SetIsSync(v bool) *BaseUserResponse {
	s.IsSync = &v
	return s
}

func (s *BaseUserResponse) SetLastLoginTime(v int64) *BaseUserResponse {
	s.LastLoginTime = &v
	return s
}

func (s *BaseUserResponse) SetNeedChangePasswordNextLogin(v bool) *BaseUserResponse {
	s.NeedChangePasswordNextLogin = &v
	return s
}

func (s *BaseUserResponse) SetNickName(v string) *BaseUserResponse {
	s.NickName = &v
	return s
}

func (s *BaseUserResponse) SetPathStatus(v string) *BaseUserResponse {
	s.PathStatus = &v
	return s
}

func (s *BaseUserResponse) SetPermission(v map[string]*IDPermission) *BaseUserResponse {
	s.Permission = v
	return s
}

func (s *BaseUserResponse) SetPhone(v string) *BaseUserResponse {
	s.Phone = &v
	return s
}

func (s *BaseUserResponse) SetPhoneRegion(v string) *BaseUserResponse {
	s.PhoneRegion = &v
	return s
}

func (s *BaseUserResponse) SetRole(v string) *BaseUserResponse {
	s.Role = &v
	return s
}

func (s *BaseUserResponse) SetStatus(v string) *BaseUserResponse {
	s.Status = &v
	return s
}

func (s *BaseUserResponse) SetUpdatedAt(v string) *BaseUserResponse {
	s.UpdatedAt = &v
	return s
}

func (s *BaseUserResponse) SetUserData(v map[string]interface{}) *BaseUserResponse {
	s.UserData = v
	return s
}

func (s *BaseUserResponse) SetUserId(v string) *BaseUserResponse {
	s.UserId = &v
	return s
}

func (s *BaseUserResponse) SetUserName(v string) *BaseUserResponse {
	s.UserName = &v
	return s
}

type BenefitPkgDeliveryInfo struct {
	Amount      *int64  `json:"amount,omitempty" xml:"amount,omitempty"`
	CreatedAt   *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	ExpireTime  *string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	IsPermanent *bool   `json:"is_permanent,omitempty" xml:"is_permanent,omitempty"`
}

func (s BenefitPkgDeliveryInfo) String() string {
	return tea.Prettify(s)
}

func (s BenefitPkgDeliveryInfo) GoString() string {
	return s.String()
}

func (s *BenefitPkgDeliveryInfo) SetAmount(v int64) *BenefitPkgDeliveryInfo {
	s.Amount = &v
	return s
}

func (s *BenefitPkgDeliveryInfo) SetCreatedAt(v string) *BenefitPkgDeliveryInfo {
	s.CreatedAt = &v
	return s
}

func (s *BenefitPkgDeliveryInfo) SetExpireTime(v string) *BenefitPkgDeliveryInfo {
	s.ExpireTime = &v
	return s
}

func (s *BenefitPkgDeliveryInfo) SetIsPermanent(v bool) *BenefitPkgDeliveryInfo {
	s.IsPermanent = &v
	return s
}

type CNameStatus struct {
	// example:
	//
	// BINDING/BOUND
	BingdingState *string `json:"bingding_state,omitempty" xml:"bingding_state,omitempty"`
	// example:
	//
	// NORMAL/ABNORMAL
	LegalState *string `json:"legal_state,omitempty" xml:"legal_state,omitempty"`
	// example:
	//
	// beian
	Remark *string `json:"remark,omitempty" xml:"remark,omitempty"`
}

func (s CNameStatus) String() string {
	return tea.Prettify(s)
}

func (s CNameStatus) GoString() string {
	return s.String()
}

func (s *CNameStatus) SetBingdingState(v string) *CNameStatus {
	s.BingdingState = &v
	return s
}

func (s *CNameStatus) SetLegalState(v string) *CNameStatus {
	s.LegalState = &v
	return s
}

func (s *CNameStatus) SetRemark(v string) *CNameStatus {
	s.Remark = &v
	return s
}

type CdnFileDownloadCallbackInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// ccp-bj1-bj-1234
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// This parameter is required.
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// This parameter is required.
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	Expire *int64 `json:"expire,omitempty" xml:"expire,omitempty"`
	// This parameter is required.
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// This parameter is required.
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// md5.Sum([]byte(fmt.Sprintf("%v%v%v%v%v%v...%v", 		req.Object, req.Range, req.DomainID, req.DriveID, req.UserID, req.FileID, req.Expire)))
	Token *string `json:"token,omitempty" xml:"token,omitempty"`
	// This parameter is required.
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CdnFileDownloadCallbackInfo) String() string {
	return tea.Prettify(s)
}

func (s CdnFileDownloadCallbackInfo) GoString() string {
	return s.String()
}

func (s *CdnFileDownloadCallbackInfo) SetBucket(v string) *CdnFileDownloadCallbackInfo {
	s.Bucket = &v
	return s
}

func (s *CdnFileDownloadCallbackInfo) SetDomainId(v string) *CdnFileDownloadCallbackInfo {
	s.DomainId = &v
	return s
}

func (s *CdnFileDownloadCallbackInfo) SetDriveId(v string) *CdnFileDownloadCallbackInfo {
	s.DriveId = &v
	return s
}

func (s *CdnFileDownloadCallbackInfo) SetExpire(v int64) *CdnFileDownloadCallbackInfo {
	s.Expire = &v
	return s
}

func (s *CdnFileDownloadCallbackInfo) SetFileId(v string) *CdnFileDownloadCallbackInfo {
	s.FileId = &v
	return s
}

func (s *CdnFileDownloadCallbackInfo) SetObject(v string) *CdnFileDownloadCallbackInfo {
	s.Object = &v
	return s
}

func (s *CdnFileDownloadCallbackInfo) SetToken(v string) *CdnFileDownloadCallbackInfo {
	s.Token = &v
	return s
}

func (s *CdnFileDownloadCallbackInfo) SetUserId(v string) *CdnFileDownloadCallbackInfo {
	s.UserId = &v
	return s
}

type CertInfo struct {
	// example:
	//
	// xxx
	CertBody *string `json:"cert_body,omitempty" xml:"cert_body,omitempty"`
	// example:
	//
	// xxx
	CertName *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	// example:
	//
	// xxx
	CertPrivatekey *string `json:"cert_privatekey,omitempty" xml:"cert_privatekey,omitempty"`
}

func (s CertInfo) String() string {
	return tea.Prettify(s)
}

func (s CertInfo) GoString() string {
	return s.String()
}

func (s *CertInfo) SetCertBody(v string) *CertInfo {
	s.CertBody = &v
	return s
}

func (s *CertInfo) SetCertName(v string) *CertInfo {
	s.CertName = &v
	return s
}

func (s *CertInfo) SetCertPrivatekey(v string) *CertInfo {
	s.CertPrivatekey = &v
	return s
}

type ClearRecycleBinItem struct {
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	TaskId      *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s ClearRecycleBinItem) String() string {
	return tea.Prettify(s)
}

func (s ClearRecycleBinItem) GoString() string {
	return s.String()
}

func (s *ClearRecycleBinItem) SetAsyncTaskId(v string) *ClearRecycleBinItem {
	s.AsyncTaskId = &v
	return s
}

func (s *ClearRecycleBinItem) SetDomainId(v string) *ClearRecycleBinItem {
	s.DomainId = &v
	return s
}

func (s *ClearRecycleBinItem) SetDriveId(v string) *ClearRecycleBinItem {
	s.DriveId = &v
	return s
}

func (s *ClearRecycleBinItem) SetTaskId(v string) *ClearRecycleBinItem {
	s.TaskId = &v
	return s
}

type CommonFileItem struct {
	DriveId    *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId     *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s CommonFileItem) String() string {
	return tea.Prettify(s)
}

func (s CommonFileItem) GoString() string {
	return s.String()
}

func (s *CommonFileItem) SetDriveId(v string) *CommonFileItem {
	s.DriveId = &v
	return s
}

func (s *CommonFileItem) SetFileId(v string) *CommonFileItem {
	s.FileId = &v
	return s
}

func (s *CommonFileItem) SetRevisionId(v string) *CommonFileItem {
	s.RevisionId = &v
	return s
}

type CssCreateOrderParam struct {
	AgentId             *string                 `json:"agentId,omitempty" xml:"agentId,omitempty"`
	AutoPay             *bool                   `json:"autoPay,omitempty" xml:"autoPay,omitempty"`
	AutoUseCoupon       *bool                   `json:"autoUseCoupon,omitempty" xml:"autoUseCoupon,omitempty"`
	Bid                 *string                 `json:"bid,omitempty" xml:"bid,omitempty"`
	BuyerId             *int64                  `json:"buyerId,omitempty" xml:"buyerId,omitempty"`
	Certificate         *string                 `json:"certificate,omitempty" xml:"certificate,omitempty"`
	ChildId             *int64                  `json:"childId,omitempty" xml:"childId,omitempty"`
	CilentIp            *string                 `json:"cilentIp,omitempty" xml:"cilentIp,omitempty"`
	Commodities         []*CssInstanceCommodity `json:"commodities,omitempty" xml:"commodities,omitempty" type:"Repeated"`
	CreaterNick         *string                 `json:"createrNick,omitempty" xml:"createrNick,omitempty"`
	CssAuthRequestParam interface{}             `json:"cssAuthRequestParam,omitempty" xml:"cssAuthRequestParam,omitempty"`
	FromApp             *string                 `json:"fromApp,omitempty" xml:"fromApp,omitempty"`
	Language            *string                 `json:"language,omitempty" xml:"language,omitempty"`
	MarketType          *int64                  `json:"marketType,omitempty" xml:"marketType,omitempty"`
	Memo                *string                 `json:"memo,omitempty" xml:"memo,omitempty"`
	OrderOrigin         *string                 `json:"orderOrigin,omitempty" xml:"orderOrigin,omitempty"`
	OrderParams         map[string]*string      `json:"orderParams,omitempty" xml:"orderParams,omitempty"`
	PayerId             *int64                  `json:"payerId,omitempty" xml:"payerId,omitempty"`
	PlanGroupId         *int64                  `json:"planGroupId,omitempty" xml:"planGroupId,omitempty"`
	PlanId              *int64                  `json:"planId,omitempty" xml:"planId,omitempty"`
	PlanInstanceId      *string                 `json:"planInstanceId,omitempty" xml:"planInstanceId,omitempty"`
	PromotionCode       *string                 `json:"promotionCode,omitempty" xml:"promotionCode,omitempty"`
	PromotionInputParam interface{}             `json:"promotionInputParam,omitempty" xml:"promotionInputParam,omitempty"`
	RequestId           *string                 `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SkipChannel         *bool                   `json:"skipChannel,omitempty" xml:"skipChannel,omitempty"`
	Token               *string                 `json:"token,omitempty" xml:"token,omitempty"`
	TransientAccess     interface{}             `json:"transientAccess,omitempty" xml:"transientAccess,omitempty"`
	UmidToken           *string                 `json:"umidToken,omitempty" xml:"umidToken,omitempty"`
	UserId              *int64                  `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CssCreateOrderParam) String() string {
	return tea.Prettify(s)
}

func (s CssCreateOrderParam) GoString() string {
	return s.String()
}

func (s *CssCreateOrderParam) SetAgentId(v string) *CssCreateOrderParam {
	s.AgentId = &v
	return s
}

func (s *CssCreateOrderParam) SetAutoPay(v bool) *CssCreateOrderParam {
	s.AutoPay = &v
	return s
}

func (s *CssCreateOrderParam) SetAutoUseCoupon(v bool) *CssCreateOrderParam {
	s.AutoUseCoupon = &v
	return s
}

func (s *CssCreateOrderParam) SetBid(v string) *CssCreateOrderParam {
	s.Bid = &v
	return s
}

func (s *CssCreateOrderParam) SetBuyerId(v int64) *CssCreateOrderParam {
	s.BuyerId = &v
	return s
}

func (s *CssCreateOrderParam) SetCertificate(v string) *CssCreateOrderParam {
	s.Certificate = &v
	return s
}

func (s *CssCreateOrderParam) SetChildId(v int64) *CssCreateOrderParam {
	s.ChildId = &v
	return s
}

func (s *CssCreateOrderParam) SetCilentIp(v string) *CssCreateOrderParam {
	s.CilentIp = &v
	return s
}

func (s *CssCreateOrderParam) SetCommodities(v []*CssInstanceCommodity) *CssCreateOrderParam {
	s.Commodities = v
	return s
}

func (s *CssCreateOrderParam) SetCreaterNick(v string) *CssCreateOrderParam {
	s.CreaterNick = &v
	return s
}

func (s *CssCreateOrderParam) SetCssAuthRequestParam(v interface{}) *CssCreateOrderParam {
	s.CssAuthRequestParam = v
	return s
}

func (s *CssCreateOrderParam) SetFromApp(v string) *CssCreateOrderParam {
	s.FromApp = &v
	return s
}

func (s *CssCreateOrderParam) SetLanguage(v string) *CssCreateOrderParam {
	s.Language = &v
	return s
}

func (s *CssCreateOrderParam) SetMarketType(v int64) *CssCreateOrderParam {
	s.MarketType = &v
	return s
}

func (s *CssCreateOrderParam) SetMemo(v string) *CssCreateOrderParam {
	s.Memo = &v
	return s
}

func (s *CssCreateOrderParam) SetOrderOrigin(v string) *CssCreateOrderParam {
	s.OrderOrigin = &v
	return s
}

func (s *CssCreateOrderParam) SetOrderParams(v map[string]*string) *CssCreateOrderParam {
	s.OrderParams = v
	return s
}

func (s *CssCreateOrderParam) SetPayerId(v int64) *CssCreateOrderParam {
	s.PayerId = &v
	return s
}

func (s *CssCreateOrderParam) SetPlanGroupId(v int64) *CssCreateOrderParam {
	s.PlanGroupId = &v
	return s
}

func (s *CssCreateOrderParam) SetPlanId(v int64) *CssCreateOrderParam {
	s.PlanId = &v
	return s
}

func (s *CssCreateOrderParam) SetPlanInstanceId(v string) *CssCreateOrderParam {
	s.PlanInstanceId = &v
	return s
}

func (s *CssCreateOrderParam) SetPromotionCode(v string) *CssCreateOrderParam {
	s.PromotionCode = &v
	return s
}

func (s *CssCreateOrderParam) SetPromotionInputParam(v interface{}) *CssCreateOrderParam {
	s.PromotionInputParam = v
	return s
}

func (s *CssCreateOrderParam) SetRequestId(v string) *CssCreateOrderParam {
	s.RequestId = &v
	return s
}

func (s *CssCreateOrderParam) SetSkipChannel(v bool) *CssCreateOrderParam {
	s.SkipChannel = &v
	return s
}

func (s *CssCreateOrderParam) SetToken(v string) *CssCreateOrderParam {
	s.Token = &v
	return s
}

func (s *CssCreateOrderParam) SetTransientAccess(v interface{}) *CssCreateOrderParam {
	s.TransientAccess = v
	return s
}

func (s *CssCreateOrderParam) SetUmidToken(v string) *CssCreateOrderParam {
	s.UmidToken = &v
	return s
}

func (s *CssCreateOrderParam) SetUserId(v int64) *CssCreateOrderParam {
	s.UserId = &v
	return s
}

type CssInstanceCommodity struct {
	ActivityId                   *int64                  `json:"activityId,omitempty" xml:"activityId,omitempty"`
	AliyunProduceCode            *string                 `json:"aliyunProduceCode,omitempty" xml:"aliyunProduceCode,omitempty"`
	ChargeType                   *string                 `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	CommodityCode                *string                 `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	Components                   []*CssInstanceComponent `json:"components,omitempty" xml:"components,omitempty" type:"Repeated"`
	Duration                     *int64                  `json:"duration,omitempty" xml:"duration,omitempty"`
	InstanceId                   *string                 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	IsFree                       *bool                   `json:"isFree,omitempty" xml:"isFree,omitempty"`
	IsPrePayPostCharge           *bool                   `json:"isPrePayPostCharge,omitempty" xml:"isPrePayPostCharge,omitempty"`
	IsRenewChange                *bool                   `json:"isRenewChange,omitempty" xml:"isRenewChange,omitempty"`
	IsSyncToSubscription         *bool                   `json:"isSyncToSubscription,omitempty" xml:"isSyncToSubscription,omitempty"`
	OrderParams                  map[string]*string      `json:"orderParams,omitempty" xml:"orderParams,omitempty"`
	OrderType                    *string                 `json:"orderType,omitempty" xml:"orderType,omitempty"`
	PlanItemId                   *int64                  `json:"planItemId,omitempty" xml:"planItemId,omitempty"`
	PricingCycle                 *string                 `json:"pricingCycle,omitempty" xml:"pricingCycle,omitempty"`
	Quantity                     *int64                  `json:"quantity,omitempty" xml:"quantity,omitempty"`
	RedeemNoList                 []*string               `json:"redeemNoList,omitempty" xml:"redeemNoList,omitempty" type:"Repeated"`
	RedeemOrderType              *string                 `json:"redeemOrderType,omitempty" xml:"redeemOrderType,omitempty"`
	RefundSpecCode               *string                 `json:"refundSpecCode,omitempty" xml:"refundSpecCode,omitempty"`
	SpecCode                     *string                 `json:"specCode,omitempty" xml:"specCode,omitempty"`
	SpecUpgradeOriginSpecCodes   []*string               `json:"specUpgradeOriginSpecCodes,omitempty" xml:"specUpgradeOriginSpecCodes,omitempty" type:"Repeated"`
	SpecifyStartDate             *int64                  `json:"specifyStartDate,omitempty" xml:"specifyStartDate,omitempty"`
	UpgradeInquireFinancialValue *bool                   `json:"upgradeInquireFinancialValue,omitempty" xml:"upgradeInquireFinancialValue,omitempty"`
}

func (s CssInstanceCommodity) String() string {
	return tea.Prettify(s)
}

func (s CssInstanceCommodity) GoString() string {
	return s.String()
}

func (s *CssInstanceCommodity) SetActivityId(v int64) *CssInstanceCommodity {
	s.ActivityId = &v
	return s
}

func (s *CssInstanceCommodity) SetAliyunProduceCode(v string) *CssInstanceCommodity {
	s.AliyunProduceCode = &v
	return s
}

func (s *CssInstanceCommodity) SetChargeType(v string) *CssInstanceCommodity {
	s.ChargeType = &v
	return s
}

func (s *CssInstanceCommodity) SetCommodityCode(v string) *CssInstanceCommodity {
	s.CommodityCode = &v
	return s
}

func (s *CssInstanceCommodity) SetComponents(v []*CssInstanceComponent) *CssInstanceCommodity {
	s.Components = v
	return s
}

func (s *CssInstanceCommodity) SetDuration(v int64) *CssInstanceCommodity {
	s.Duration = &v
	return s
}

func (s *CssInstanceCommodity) SetInstanceId(v string) *CssInstanceCommodity {
	s.InstanceId = &v
	return s
}

func (s *CssInstanceCommodity) SetIsFree(v bool) *CssInstanceCommodity {
	s.IsFree = &v
	return s
}

func (s *CssInstanceCommodity) SetIsPrePayPostCharge(v bool) *CssInstanceCommodity {
	s.IsPrePayPostCharge = &v
	return s
}

func (s *CssInstanceCommodity) SetIsRenewChange(v bool) *CssInstanceCommodity {
	s.IsRenewChange = &v
	return s
}

func (s *CssInstanceCommodity) SetIsSyncToSubscription(v bool) *CssInstanceCommodity {
	s.IsSyncToSubscription = &v
	return s
}

func (s *CssInstanceCommodity) SetOrderParams(v map[string]*string) *CssInstanceCommodity {
	s.OrderParams = v
	return s
}

func (s *CssInstanceCommodity) SetOrderType(v string) *CssInstanceCommodity {
	s.OrderType = &v
	return s
}

func (s *CssInstanceCommodity) SetPlanItemId(v int64) *CssInstanceCommodity {
	s.PlanItemId = &v
	return s
}

func (s *CssInstanceCommodity) SetPricingCycle(v string) *CssInstanceCommodity {
	s.PricingCycle = &v
	return s
}

func (s *CssInstanceCommodity) SetQuantity(v int64) *CssInstanceCommodity {
	s.Quantity = &v
	return s
}

func (s *CssInstanceCommodity) SetRedeemNoList(v []*string) *CssInstanceCommodity {
	s.RedeemNoList = v
	return s
}

func (s *CssInstanceCommodity) SetRedeemOrderType(v string) *CssInstanceCommodity {
	s.RedeemOrderType = &v
	return s
}

func (s *CssInstanceCommodity) SetRefundSpecCode(v string) *CssInstanceCommodity {
	s.RefundSpecCode = &v
	return s
}

func (s *CssInstanceCommodity) SetSpecCode(v string) *CssInstanceCommodity {
	s.SpecCode = &v
	return s
}

func (s *CssInstanceCommodity) SetSpecUpgradeOriginSpecCodes(v []*string) *CssInstanceCommodity {
	s.SpecUpgradeOriginSpecCodes = v
	return s
}

func (s *CssInstanceCommodity) SetSpecifyStartDate(v int64) *CssInstanceCommodity {
	s.SpecifyStartDate = &v
	return s
}

func (s *CssInstanceCommodity) SetUpgradeInquireFinancialValue(v bool) *CssInstanceCommodity {
	s.UpgradeInquireFinancialValue = &v
	return s
}

type CssInstanceComponent struct {
	ComponentCode    *string                `json:"componentCode,omitempty" xml:"componentCode,omitempty"`
	ComponentName    *string                `json:"componentName,omitempty" xml:"componentName,omitempty"`
	GlobalKey        *string                `json:"globalKey,omitempty" xml:"globalKey,omitempty"`
	InstanceProperty []*CssInstanceProperty `json:"instanceProperty,omitempty" xml:"instanceProperty,omitempty" type:"Repeated"`
	ModuleAttrStatus *int64                 `json:"moduleAttrStatus,omitempty" xml:"moduleAttrStatus,omitempty"`
	Tag              *string                `json:"tag,omitempty" xml:"tag,omitempty"`
}

func (s CssInstanceComponent) String() string {
	return tea.Prettify(s)
}

func (s CssInstanceComponent) GoString() string {
	return s.String()
}

func (s *CssInstanceComponent) SetComponentCode(v string) *CssInstanceComponent {
	s.ComponentCode = &v
	return s
}

func (s *CssInstanceComponent) SetComponentName(v string) *CssInstanceComponent {
	s.ComponentName = &v
	return s
}

func (s *CssInstanceComponent) SetGlobalKey(v string) *CssInstanceComponent {
	s.GlobalKey = &v
	return s
}

func (s *CssInstanceComponent) SetInstanceProperty(v []*CssInstanceProperty) *CssInstanceComponent {
	s.InstanceProperty = v
	return s
}

func (s *CssInstanceComponent) SetModuleAttrStatus(v int64) *CssInstanceComponent {
	s.ModuleAttrStatus = &v
	return s
}

func (s *CssInstanceComponent) SetTag(v string) *CssInstanceComponent {
	s.Tag = &v
	return s
}

type CssInstanceProperty struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	GlobalKey *string `json:"globalKey,omitempty" xml:"globalKey,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Unit      *string `json:"unit,omitempty" xml:"unit,omitempty"`
	Value     *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s CssInstanceProperty) String() string {
	return tea.Prettify(s)
}

func (s CssInstanceProperty) GoString() string {
	return s.String()
}

func (s *CssInstanceProperty) SetCode(v string) *CssInstanceProperty {
	s.Code = &v
	return s
}

func (s *CssInstanceProperty) SetGlobalKey(v string) *CssInstanceProperty {
	s.GlobalKey = &v
	return s
}

func (s *CssInstanceProperty) SetName(v string) *CssInstanceProperty {
	s.Name = &v
	return s
}

func (s *CssInstanceProperty) SetUnit(v string) *CssInstanceProperty {
	s.Unit = &v
	return s
}

func (s *CssInstanceProperty) SetValue(v string) *CssInstanceProperty {
	s.Value = &v
	return s
}

type CssProduce struct {
	Bid         *string        `json:"bid,omitempty" xml:"bid,omitempty"`
	BuyerId     *int64         `json:"buyerId,omitempty" xml:"buyerId,omitempty"`
	ChildId     *int64         `json:"childId,omitempty" xml:"childId,omitempty"`
	FromApp     *string        `json:"fromApp,omitempty" xml:"fromApp,omitempty"`
	OrderId     *int64         `json:"orderId,omitempty" xml:"orderId,omitempty"`
	PayerId     *int64         `json:"payerId,omitempty" xml:"payerId,omitempty"`
	Purchases   []*CssPurchase `json:"purchases,omitempty" xml:"purchases,omitempty" type:"Repeated"`
	RequestId   *string        `json:"requestId,omitempty" xml:"requestId,omitempty"`
	SkipChannel *bool          `json:"skipChannel,omitempty" xml:"skipChannel,omitempty"`
	Token       *string        `json:"token,omitempty" xml:"token,omitempty"`
	UserId      *int64         `json:"userId,omitempty" xml:"userId,omitempty"`
}

func (s CssProduce) String() string {
	return tea.Prettify(s)
}

func (s CssProduce) GoString() string {
	return s.String()
}

func (s *CssProduce) SetBid(v string) *CssProduce {
	s.Bid = &v
	return s
}

func (s *CssProduce) SetBuyerId(v int64) *CssProduce {
	s.BuyerId = &v
	return s
}

func (s *CssProduce) SetChildId(v int64) *CssProduce {
	s.ChildId = &v
	return s
}

func (s *CssProduce) SetFromApp(v string) *CssProduce {
	s.FromApp = &v
	return s
}

func (s *CssProduce) SetOrderId(v int64) *CssProduce {
	s.OrderId = &v
	return s
}

func (s *CssProduce) SetPayerId(v int64) *CssProduce {
	s.PayerId = &v
	return s
}

func (s *CssProduce) SetPurchases(v []*CssPurchase) *CssProduce {
	s.Purchases = v
	return s
}

func (s *CssProduce) SetRequestId(v string) *CssProduce {
	s.RequestId = &v
	return s
}

func (s *CssProduce) SetSkipChannel(v bool) *CssProduce {
	s.SkipChannel = &v
	return s
}

func (s *CssProduce) SetToken(v string) *CssProduce {
	s.Token = &v
	return s
}

func (s *CssProduce) SetUserId(v int64) *CssProduce {
	s.UserId = &v
	return s
}

type CssPurchase struct {
	ChargeType         *string                 `json:"chargeType,omitempty" xml:"chargeType,omitempty"`
	CommodityCode      *string                 `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	EndDate            *int64                  `json:"endDate,omitempty" xml:"endDate,omitempty"`
	GmtCreate          *int64                  `json:"gmtCreate,omitempty" xml:"gmtCreate,omitempty"`
	InstanceComponents []*CssInstanceComponent `json:"instanceComponents,omitempty" xml:"instanceComponents,omitempty" type:"Repeated"`
	InstanceId         *string                 `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	OrderType          *string                 `json:"orderType,omitempty" xml:"orderType,omitempty"`
	PurchaseParams     map[string]*string      `json:"purchaseParams,omitempty" xml:"purchaseParams,omitempty"`
	StartDate          *int64                  `json:"startDate,omitempty" xml:"startDate,omitempty"`
}

func (s CssPurchase) String() string {
	return tea.Prettify(s)
}

func (s CssPurchase) GoString() string {
	return s.String()
}

func (s *CssPurchase) SetChargeType(v string) *CssPurchase {
	s.ChargeType = &v
	return s
}

func (s *CssPurchase) SetCommodityCode(v string) *CssPurchase {
	s.CommodityCode = &v
	return s
}

func (s *CssPurchase) SetEndDate(v int64) *CssPurchase {
	s.EndDate = &v
	return s
}

func (s *CssPurchase) SetGmtCreate(v int64) *CssPurchase {
	s.GmtCreate = &v
	return s
}

func (s *CssPurchase) SetInstanceComponents(v []*CssInstanceComponent) *CssPurchase {
	s.InstanceComponents = v
	return s
}

func (s *CssPurchase) SetInstanceId(v string) *CssPurchase {
	s.InstanceId = &v
	return s
}

func (s *CssPurchase) SetOrderType(v string) *CssPurchase {
	s.OrderType = &v
	return s
}

func (s *CssPurchase) SetPurchaseParams(v map[string]*string) *CssPurchase {
	s.PurchaseParams = v
	return s
}

func (s *CssPurchase) SetStartDate(v int64) *CssPurchase {
	s.StartDate = &v
	return s
}

type CustomSideLinkConfig struct {
	Icon *string `json:"icon,omitempty" xml:"icon,omitempty"`
	Link *string `json:"link,omitempty" xml:"link,omitempty"`
	Text *string `json:"text,omitempty" xml:"text,omitempty"`
}

func (s CustomSideLinkConfig) String() string {
	return tea.Prettify(s)
}

func (s CustomSideLinkConfig) GoString() string {
	return s.String()
}

func (s *CustomSideLinkConfig) SetIcon(v string) *CustomSideLinkConfig {
	s.Icon = &v
	return s
}

func (s *CustomSideLinkConfig) SetLink(v string) *CustomSideLinkConfig {
	s.Link = &v
	return s
}

func (s *CustomSideLinkConfig) SetText(v string) *CustomSideLinkConfig {
	s.Text = &v
	return s
}

type DataBoxPrivileges struct {
	FeatureAttrId *string `json:"feature_attr_id,omitempty" xml:"feature_attr_id,omitempty"`
	FeatureId     *string `json:"feature_id,omitempty" xml:"feature_id,omitempty"`
	Quota         *int64  `json:"quota,omitempty" xml:"quota,omitempty"`
}

func (s DataBoxPrivileges) String() string {
	return tea.Prettify(s)
}

func (s DataBoxPrivileges) GoString() string {
	return s.String()
}

func (s *DataBoxPrivileges) SetFeatureAttrId(v string) *DataBoxPrivileges {
	s.FeatureAttrId = &v
	return s
}

func (s *DataBoxPrivileges) SetFeatureId(v string) *DataBoxPrivileges {
	s.FeatureId = &v
	return s
}

func (s *DataBoxPrivileges) SetQuota(v int64) *DataBoxPrivileges {
	s.Quota = &v
	return s
}

type DataCName struct {
	CertExpireTime *int64  `json:"cert_expire_time,omitempty" xml:"cert_expire_time,omitempty"`
	CertName       *string `json:"cert_name,omitempty" xml:"cert_name,omitempty"`
	Cname          *string `json:"cname,omitempty" xml:"cname,omitempty"`
	CnameType      *string `json:"cname_type,omitempty" xml:"cname_type,omitempty"`
	Location       *string `json:"location,omitempty" xml:"location,omitempty"`
	StoreId        *string `json:"store_id,omitempty" xml:"store_id,omitempty"`
}

func (s DataCName) String() string {
	return tea.Prettify(s)
}

func (s DataCName) GoString() string {
	return s.String()
}

func (s *DataCName) SetCertExpireTime(v int64) *DataCName {
	s.CertExpireTime = &v
	return s
}

func (s *DataCName) SetCertName(v string) *DataCName {
	s.CertName = &v
	return s
}

func (s *DataCName) SetCname(v string) *DataCName {
	s.Cname = &v
	return s
}

func (s *DataCName) SetCnameType(v string) *DataCName {
	s.CnameType = &v
	return s
}

func (s *DataCName) SetLocation(v string) *DataCName {
	s.Location = &v
	return s
}

func (s *DataCName) SetStoreId(v string) *DataCName {
	s.StoreId = &v
	return s
}

type Domain struct {
	CreatedAt                  *string            `json:"created_at,omitempty" xml:"created_at,omitempty"`
	DataHashName               *string            `json:"data_hash_name,omitempty" xml:"data_hash_name,omitempty"`
	Description                *string            `json:"description,omitempty" xml:"description,omitempty"`
	DomainId                   *string            `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DomainName                 *string            `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	InitDriveEnable            *bool              `json:"init_drive_enable,omitempty" xml:"init_drive_enable,omitempty"`
	InitDriveSize              *int64             `json:"init_drive_size,omitempty" xml:"init_drive_size,omitempty"`
	ParentDomainId             *string            `json:"parent_domain_id,omitempty" xml:"parent_domain_id,omitempty"`
	PublishedAppAccessStrategy *AppAccessStrategy `json:"published_app_access_strategy,omitempty" xml:"published_app_access_strategy,omitempty"`
	Sharable                   *bool              `json:"sharable,omitempty" xml:"sharable,omitempty"`
	SizeQuota                  *int64             `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	SizeQuotaUsed              *int64             `json:"size_quota_used,omitempty" xml:"size_quota_used,omitempty"`
	Status                     *int64             `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt                  *string            `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UsedSize                   *int64             `json:"used_size,omitempty" xml:"used_size,omitempty"`
	UserCountQuota             *int64             `json:"user_count_quota,omitempty" xml:"user_count_quota,omitempty"`
}

func (s Domain) String() string {
	return tea.Prettify(s)
}

func (s Domain) GoString() string {
	return s.String()
}

func (s *Domain) SetCreatedAt(v string) *Domain {
	s.CreatedAt = &v
	return s
}

func (s *Domain) SetDataHashName(v string) *Domain {
	s.DataHashName = &v
	return s
}

func (s *Domain) SetDescription(v string) *Domain {
	s.Description = &v
	return s
}

func (s *Domain) SetDomainId(v string) *Domain {
	s.DomainId = &v
	return s
}

func (s *Domain) SetDomainName(v string) *Domain {
	s.DomainName = &v
	return s
}

func (s *Domain) SetInitDriveEnable(v bool) *Domain {
	s.InitDriveEnable = &v
	return s
}

func (s *Domain) SetInitDriveSize(v int64) *Domain {
	s.InitDriveSize = &v
	return s
}

func (s *Domain) SetParentDomainId(v string) *Domain {
	s.ParentDomainId = &v
	return s
}

func (s *Domain) SetPublishedAppAccessStrategy(v *AppAccessStrategy) *Domain {
	s.PublishedAppAccessStrategy = v
	return s
}

func (s *Domain) SetSharable(v bool) *Domain {
	s.Sharable = &v
	return s
}

func (s *Domain) SetSizeQuota(v int64) *Domain {
	s.SizeQuota = &v
	return s
}

func (s *Domain) SetSizeQuotaUsed(v int64) *Domain {
	s.SizeQuotaUsed = &v
	return s
}

func (s *Domain) SetStatus(v int64) *Domain {
	s.Status = &v
	return s
}

func (s *Domain) SetUpdatedAt(v string) *Domain {
	s.UpdatedAt = &v
	return s
}

func (s *Domain) SetUsedSize(v int64) *Domain {
	s.UsedSize = &v
	return s
}

func (s *Domain) SetUserCountQuota(v int64) *Domain {
	s.UserCountQuota = &v
	return s
}

type DomainAppConfig struct {
	AllowUploadCustomFileExtList []*string `json:"allow_upload_custom_file_ext_list,omitempty" xml:"allow_upload_custom_file_ext_list,omitempty" type:"Repeated"`
	AllowUploadFileCategoryList  []*string `json:"allow_upload_file_category_list,omitempty" xml:"allow_upload_file_category_list,omitempty" type:"Repeated"`
	SameNameFileUploadMode       *string   `json:"same_name_file_upload_mode,omitempty" xml:"same_name_file_upload_mode,omitempty"`
	SingleFileUploadSizeLimit    *int64    `json:"single_file_upload_size_limit,omitempty" xml:"single_file_upload_size_limit,omitempty"`
	WebClientDownloadMode        *string   `json:"web_client_download_mode,omitempty" xml:"web_client_download_mode,omitempty"`
}

func (s DomainAppConfig) String() string {
	return tea.Prettify(s)
}

func (s DomainAppConfig) GoString() string {
	return s.String()
}

func (s *DomainAppConfig) SetAllowUploadCustomFileExtList(v []*string) *DomainAppConfig {
	s.AllowUploadCustomFileExtList = v
	return s
}

func (s *DomainAppConfig) SetAllowUploadFileCategoryList(v []*string) *DomainAppConfig {
	s.AllowUploadFileCategoryList = v
	return s
}

func (s *DomainAppConfig) SetSameNameFileUploadMode(v string) *DomainAppConfig {
	s.SameNameFileUploadMode = &v
	return s
}

func (s *DomainAppConfig) SetSingleFileUploadSizeLimit(v int64) *DomainAppConfig {
	s.SingleFileUploadSizeLimit = &v
	return s
}

func (s *DomainAppConfig) SetWebClientDownloadMode(v string) *DomainAppConfig {
	s.WebClientDownloadMode = &v
	return s
}

type DomainBuildClientConfig struct {
	Copyright *string `json:"copyright,omitempty" xml:"copyright,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s DomainBuildClientConfig) String() string {
	return tea.Prettify(s)
}

func (s DomainBuildClientConfig) GoString() string {
	return s.String()
}

func (s *DomainBuildClientConfig) SetCopyright(v string) *DomainBuildClientConfig {
	s.Copyright = &v
	return s
}

func (s *DomainBuildClientConfig) SetName(v string) *DomainBuildClientConfig {
	s.Name = &v
	return s
}

type DomainEndpoints struct {
}

func (s DomainEndpoints) String() string {
	return tea.Prettify(s)
}

func (s DomainEndpoints) GoString() string {
	return s.String()
}

type DomainSeniorConfig struct {
	ClientDownloadEnable *bool                   `json:"client_download_enable,omitempty" xml:"client_download_enable,omitempty"`
	CspFrameAncestors    *string                 `json:"csp_frame_ancestors,omitempty" xml:"csp_frame_ancestors,omitempty"`
	CustomLoginAppid     *string                 `json:"custom_login_appid,omitempty" xml:"custom_login_appid,omitempty"`
	CustomLoginUrl       *string                 `json:"custom_login_url,omitempty" xml:"custom_login_url,omitempty"`
	CustomLogoutUrl      *string                 `json:"custom_logout_url,omitempty" xml:"custom_logout_url,omitempty"`
	CustomSideLinkList   []*CustomSideLinkConfig `json:"custom_side_link_list,omitempty" xml:"custom_side_link_list,omitempty" type:"Repeated"`
	HomePageBgImageUrl   *string                 `json:"home_page_bg_image_url,omitempty" xml:"home_page_bg_image_url,omitempty"`
	HomePageFooter       *string                 `json:"home_page_footer,omitempty" xml:"home_page_footer,omitempty"`
	HomePageFooter2      *string                 `json:"home_page_footer2,omitempty" xml:"home_page_footer2,omitempty"`
	HomePageSlogan       *string                 `json:"home_page_slogan,omitempty" xml:"home_page_slogan,omitempty"`
	RefererEnable        *bool                   `json:"referer_enable,omitempty" xml:"referer_enable,omitempty"`
	WxTxtList            *WxTrustedDomainConfig  `json:"wx_txt_list,omitempty" xml:"wx_txt_list,omitempty"`
}

func (s DomainSeniorConfig) String() string {
	return tea.Prettify(s)
}

func (s DomainSeniorConfig) GoString() string {
	return s.String()
}

func (s *DomainSeniorConfig) SetClientDownloadEnable(v bool) *DomainSeniorConfig {
	s.ClientDownloadEnable = &v
	return s
}

func (s *DomainSeniorConfig) SetCspFrameAncestors(v string) *DomainSeniorConfig {
	s.CspFrameAncestors = &v
	return s
}

func (s *DomainSeniorConfig) SetCustomLoginAppid(v string) *DomainSeniorConfig {
	s.CustomLoginAppid = &v
	return s
}

func (s *DomainSeniorConfig) SetCustomLoginUrl(v string) *DomainSeniorConfig {
	s.CustomLoginUrl = &v
	return s
}

func (s *DomainSeniorConfig) SetCustomLogoutUrl(v string) *DomainSeniorConfig {
	s.CustomLogoutUrl = &v
	return s
}

func (s *DomainSeniorConfig) SetCustomSideLinkList(v []*CustomSideLinkConfig) *DomainSeniorConfig {
	s.CustomSideLinkList = v
	return s
}

func (s *DomainSeniorConfig) SetHomePageBgImageUrl(v string) *DomainSeniorConfig {
	s.HomePageBgImageUrl = &v
	return s
}

func (s *DomainSeniorConfig) SetHomePageFooter(v string) *DomainSeniorConfig {
	s.HomePageFooter = &v
	return s
}

func (s *DomainSeniorConfig) SetHomePageFooter2(v string) *DomainSeniorConfig {
	s.HomePageFooter2 = &v
	return s
}

func (s *DomainSeniorConfig) SetHomePageSlogan(v string) *DomainSeniorConfig {
	s.HomePageSlogan = &v
	return s
}

func (s *DomainSeniorConfig) SetRefererEnable(v bool) *DomainSeniorConfig {
	s.RefererEnable = &v
	return s
}

func (s *DomainSeniorConfig) SetWxTxtList(v *WxTrustedDomainConfig) *DomainSeniorConfig {
	s.WxTxtList = v
	return s
}

type Drive struct {
	CreatedAt   *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator     *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	DriveName   *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	DriveType   *string `json:"drive_type,omitempty" xml:"drive_type,omitempty"`
	Owner       *string `json:"owner,omitempty" xml:"owner,omitempty"`
	OwnerType   *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	TotalSize   *int64  `json:"total_size,omitempty" xml:"total_size,omitempty"`
	UsedSize    *int64  `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s Drive) String() string {
	return tea.Prettify(s)
}

func (s Drive) GoString() string {
	return s.String()
}

func (s *Drive) SetCreatedAt(v string) *Drive {
	s.CreatedAt = &v
	return s
}

func (s *Drive) SetCreator(v string) *Drive {
	s.Creator = &v
	return s
}

func (s *Drive) SetDescription(v string) *Drive {
	s.Description = &v
	return s
}

func (s *Drive) SetDomainId(v string) *Drive {
	s.DomainId = &v
	return s
}

func (s *Drive) SetDriveId(v string) *Drive {
	s.DriveId = &v
	return s
}

func (s *Drive) SetDriveName(v string) *Drive {
	s.DriveName = &v
	return s
}

func (s *Drive) SetDriveType(v string) *Drive {
	s.DriveType = &v
	return s
}

func (s *Drive) SetOwner(v string) *Drive {
	s.Owner = &v
	return s
}

func (s *Drive) SetOwnerType(v string) *Drive {
	s.OwnerType = &v
	return s
}

func (s *Drive) SetStatus(v string) *Drive {
	s.Status = &v
	return s
}

func (s *Drive) SetTotalSize(v int64) *Drive {
	s.TotalSize = &v
	return s
}

func (s *Drive) SetUsedSize(v int64) *Drive {
	s.UsedSize = &v
	return s
}

type DriveLogDetail struct {
	ForceDelete       *bool                   `json:"force_delete,omitempty" xml:"force_delete,omitempty"`
	HandoverOwnerName *string                 `json:"handover_owner_name,omitempty" xml:"handover_owner_name,omitempty"`
	Name              *string                 `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId           *string                 `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	OwnerName         *string                 `json:"owner_name,omitempty" xml:"owner_name,omitempty"`
	OwnerType         *string                 `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	TotalSize         *int64                  `json:"total_size,omitempty" xml:"total_size,omitempty"`
	UpdateTo          *DriveLogDetailUpdateTo `json:"update_to,omitempty" xml:"update_to,omitempty" type:"Struct"`
}

func (s DriveLogDetail) String() string {
	return tea.Prettify(s)
}

func (s DriveLogDetail) GoString() string {
	return s.String()
}

func (s *DriveLogDetail) SetForceDelete(v bool) *DriveLogDetail {
	s.ForceDelete = &v
	return s
}

func (s *DriveLogDetail) SetHandoverOwnerName(v string) *DriveLogDetail {
	s.HandoverOwnerName = &v
	return s
}

func (s *DriveLogDetail) SetName(v string) *DriveLogDetail {
	s.Name = &v
	return s
}

func (s *DriveLogDetail) SetOwnerId(v string) *DriveLogDetail {
	s.OwnerId = &v
	return s
}

func (s *DriveLogDetail) SetOwnerName(v string) *DriveLogDetail {
	s.OwnerName = &v
	return s
}

func (s *DriveLogDetail) SetOwnerType(v string) *DriveLogDetail {
	s.OwnerType = &v
	return s
}

func (s *DriveLogDetail) SetTotalSize(v int64) *DriveLogDetail {
	s.TotalSize = &v
	return s
}

func (s *DriveLogDetail) SetUpdateTo(v *DriveLogDetailUpdateTo) *DriveLogDetail {
	s.UpdateTo = v
	return s
}

type DriveLogDetailUpdateTo struct {
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId   *string `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	OwnerName *string `json:"owner_name,omitempty" xml:"owner_name,omitempty"`
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	TotalSize *int64  `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s DriveLogDetailUpdateTo) String() string {
	return tea.Prettify(s)
}

func (s DriveLogDetailUpdateTo) GoString() string {
	return s.String()
}

func (s *DriveLogDetailUpdateTo) SetName(v string) *DriveLogDetailUpdateTo {
	s.Name = &v
	return s
}

func (s *DriveLogDetailUpdateTo) SetOwnerId(v string) *DriveLogDetailUpdateTo {
	s.OwnerId = &v
	return s
}

func (s *DriveLogDetailUpdateTo) SetOwnerName(v string) *DriveLogDetailUpdateTo {
	s.OwnerName = &v
	return s
}

func (s *DriveLogDetailUpdateTo) SetOwnerType(v string) *DriveLogDetailUpdateTo {
	s.OwnerType = &v
	return s
}

func (s *DriveLogDetailUpdateTo) SetTotalSize(v int64) *DriveLogDetailUpdateTo {
	s.TotalSize = &v
	return s
}

type ExternalMultiFileRevisionConfig struct {
	RevisionCount         *int64 `json:"revision_count,omitempty" xml:"revision_count,omitempty"`
	RevisionMergeEnabled  *bool  `json:"revision_merge_enabled,omitempty" xml:"revision_merge_enabled,omitempty"`
	RevisionRecyclePeriod *int64 `json:"revision_recycle_period,omitempty" xml:"revision_recycle_period,omitempty"`
}

func (s ExternalMultiFileRevisionConfig) String() string {
	return tea.Prettify(s)
}

func (s ExternalMultiFileRevisionConfig) GoString() string {
	return s.String()
}

func (s *ExternalMultiFileRevisionConfig) SetRevisionCount(v int64) *ExternalMultiFileRevisionConfig {
	s.RevisionCount = &v
	return s
}

func (s *ExternalMultiFileRevisionConfig) SetRevisionMergeEnabled(v bool) *ExternalMultiFileRevisionConfig {
	s.RevisionMergeEnabled = &v
	return s
}

func (s *ExternalMultiFileRevisionConfig) SetRevisionRecyclePeriod(v int64) *ExternalMultiFileRevisionConfig {
	s.RevisionRecyclePeriod = &v
	return s
}

type FaceGroup struct {
	// example:
	//
	// 2022-01-14T10:10:52.83948013+08:00
	CreatedAt              *string                          `json:"created_at,omitempty" xml:"created_at,omitempty"`
	GroupCoverFaceBoundary *FaceGroupGroupCoverFaceBoundary `json:"group_cover_face_boundary,omitempty" xml:"group_cover_face_boundary,omitempty" type:"Struct"`
	// example:
	//
	// 6549c959640fbd517c9b4d93b3b36aecc45xxxxx
	GroupCoverFileId *string `json:"group_cover_file_id,omitempty" xml:"group_cover_file_id,omitempty"`
	// example:
	//
	// 1080
	GroupCoverHeight *int64 `json:"group_cover_height,omitempty" xml:"group_cover_height,omitempty"`
	// example:
	//
	// https://xxx
	GroupCoverUrl *string `json:"group_cover_url,omitempty" xml:"group_cover_url,omitempty"`
	// example:
	//
	// 1920
	GroupCoverWidth *int64 `json:"group_cover_width,omitempty" xml:"group_cover_width,omitempty"`
	// example:
	//
	// Cluster-ae6e3472-999e-410b-b54e-cd5dba****
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// example:
	//
	// name
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// example:
	//
	// 10
	ImageCount *int64  `json:"image_count,omitempty" xml:"image_count,omitempty"`
	Remarks    *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
	// example:
	//
	// 2022-01-14T10:10:52.83948013+08:00
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s FaceGroup) String() string {
	return tea.Prettify(s)
}

func (s FaceGroup) GoString() string {
	return s.String()
}

func (s *FaceGroup) SetCreatedAt(v string) *FaceGroup {
	s.CreatedAt = &v
	return s
}

func (s *FaceGroup) SetGroupCoverFaceBoundary(v *FaceGroupGroupCoverFaceBoundary) *FaceGroup {
	s.GroupCoverFaceBoundary = v
	return s
}

func (s *FaceGroup) SetGroupCoverFileId(v string) *FaceGroup {
	s.GroupCoverFileId = &v
	return s
}

func (s *FaceGroup) SetGroupCoverHeight(v int64) *FaceGroup {
	s.GroupCoverHeight = &v
	return s
}

func (s *FaceGroup) SetGroupCoverUrl(v string) *FaceGroup {
	s.GroupCoverUrl = &v
	return s
}

func (s *FaceGroup) SetGroupCoverWidth(v int64) *FaceGroup {
	s.GroupCoverWidth = &v
	return s
}

func (s *FaceGroup) SetGroupId(v string) *FaceGroup {
	s.GroupId = &v
	return s
}

func (s *FaceGroup) SetGroupName(v string) *FaceGroup {
	s.GroupName = &v
	return s
}

func (s *FaceGroup) SetImageCount(v int64) *FaceGroup {
	s.ImageCount = &v
	return s
}

func (s *FaceGroup) SetRemarks(v string) *FaceGroup {
	s.Remarks = &v
	return s
}

func (s *FaceGroup) SetUpdatedAt(v string) *FaceGroup {
	s.UpdatedAt = &v
	return s
}

type FaceGroupGroupCoverFaceBoundary struct {
	// example:
	//
	// 300
	Height *int32 `json:"Height,omitempty" xml:"Height,omitempty"`
	// example:
	//
	// 10
	Left *int32 `json:"Left,omitempty" xml:"Left,omitempty"`
	// example:
	//
	// 30
	Top *int32 `json:"Top,omitempty" xml:"Top,omitempty"`
	// example:
	//
	// 200
	Width *int32 `json:"Width,omitempty" xml:"Width,omitempty"`
}

func (s FaceGroupGroupCoverFaceBoundary) String() string {
	return tea.Prettify(s)
}

func (s FaceGroupGroupCoverFaceBoundary) GoString() string {
	return s.String()
}

func (s *FaceGroupGroupCoverFaceBoundary) SetHeight(v int32) *FaceGroupGroupCoverFaceBoundary {
	s.Height = &v
	return s
}

func (s *FaceGroupGroupCoverFaceBoundary) SetLeft(v int32) *FaceGroupGroupCoverFaceBoundary {
	s.Left = &v
	return s
}

func (s *FaceGroupGroupCoverFaceBoundary) SetTop(v int32) *FaceGroupGroupCoverFaceBoundary {
	s.Top = &v
	return s
}

func (s *FaceGroupGroupCoverFaceBoundary) SetWidth(v int32) *FaceGroupGroupCoverFaceBoundary {
	s.Width = &v
	return s
}

type FaceThumbnail struct {
	// example:
	//
	// Cluster-e3b7fb52-22b3-44f2-9746-8c1804bd6af0
	FaceGroupId *string `json:"face_group_id,omitempty" xml:"face_group_id,omitempty"`
	// example:
	//
	// a9a66a86-73dd-4c95-8b79-1d8a49db5226
	FaceId *string `json:"face_id,omitempty" xml:"face_id,omitempty"`
	// Deprecated
	//
	// example:
	//
	// https://pds-domain.region.aliyuncs.com/QieGeH98%2F1001%2F63e5e551ee621482ab934a0687c6cda75fc07864%2F642a8a40c00f1ad379df421694713ee65170f09b?security-token=CAIS%2BgF1q6Ft5B2yfSjIr5bjHPCNnrdR8aSaSW7woVlmVd1Bt5HorDz2IHpPfHdoBe0btvU%2BlWxX6fwZlq5rR4QAXlDfNSyFeX20qFHPWZHInuDox55m4cTXNAr%2BIhr%2F29CoEIedZdjBe%2FCrRknZnytou9XTfimjWFrXWv%2Fgy%2BQQDLItUxK%2FcCBNCfpPOwJms7V6D3bKMuu3OROY6Qi5TmgQ41En1DIlt%2FXuk5DCtkqB12eXkLFF%2B97DRbG%2FdNRpMZtFVNO44fd7bKKp0lQLsUMSqv8q0fEcqGaW4o7CWQJLnzyCMvvJ9OVDFyN0aKEnH7J%2Bq%2FzxhTPrMnpkSlacGoABPMvZ8rSESUEP96Vbf%2Bk0JRg9Qb1MnaIJqWAgo8K6K0UP1CtqL2zrUtugpKKDHOYiKbq2O0S5yLUPVX5vBHqEi%2FFc7i6ZnHCMcXLJs4rKDKwRBEhovUXXlklq2q43OSVtLrXkBy9Xs1ers%2FhJhcxpNA0Vl3EWfJxa2BTylEdnLOQ%3D&x-oss-access-key-id=STS.NUVWJ9shpFfqKHAEY3YRmXTCN&x-oss-expires=1686455451&x-oss-process=image%2Fcrop%2Cx_1128%2Cy_1211%2Cw_914%2Ch_914%2Fformat%2Cjpg&x-oss-signature=jmhOz91Tww1ciMEwadDiioU7d93FDiBNr8s8mHyMqW0%3D&x-oss-signature-version=OSS2
	FaceThumbnail *string `json:"face_thumbnail,omitempty" xml:"face_thumbnail,omitempty"`
}

func (s FaceThumbnail) String() string {
	return tea.Prettify(s)
}

func (s FaceThumbnail) GoString() string {
	return s.String()
}

func (s *FaceThumbnail) SetFaceGroupId(v string) *FaceThumbnail {
	s.FaceGroupId = &v
	return s
}

func (s *FaceThumbnail) SetFaceId(v string) *FaceThumbnail {
	s.FaceId = &v
	return s
}

func (s *FaceThumbnail) SetFaceThumbnail(v string) *FaceThumbnail {
	s.FaceThumbnail = &v
	return s
}

type File struct {
	ActionList         []*string           `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	Category           *string             `json:"category,omitempty" xml:"category,omitempty"`
	ContentHash        *string             `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName    *string             `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	ContentType        *string             `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Crc64Hash          *string             `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	CreatedAt          *string             `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description        *string             `json:"description,omitempty" xml:"description,omitempty"`
	DomainId           *string             `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DownloadUrl        *string             `json:"download_url,omitempty" xml:"download_url,omitempty"`
	DriveId            *string             `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileExtension      *string             `json:"file_extension,omitempty" xml:"file_extension,omitempty"`
	FileId             *string             `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Hidden             *bool               `json:"hidden,omitempty" xml:"hidden,omitempty"`
	IdPath             *string             `json:"id_path,omitempty" xml:"id_path,omitempty"`
	ImageMediaMetadata *ImageMediaMetadata `json:"image_media_metadata,omitempty" xml:"image_media_metadata,omitempty"`
	Labels             []*string           `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	LocalCreatedAt     *string             `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	LocalModifiedAt    *string             `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	Name               *string             `json:"name,omitempty" xml:"name,omitempty"`
	NamePath           *string             `json:"name_path,omitempty" xml:"name_path,omitempty"`
	ParentFileId       *string             `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	RevisionId         *string             `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	Size               *int64              `json:"size,omitempty" xml:"size,omitempty"`
	Starred            *bool               `json:"starred,omitempty" xml:"starred,omitempty"`
	Status             *string             `json:"status,omitempty" xml:"status,omitempty"`
	Thumbnail          *string             `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	ThumbnailUrls      map[string]*string  `json:"thumbnail_urls,omitempty" xml:"thumbnail_urls,omitempty"`
	TrashedAt          *string             `json:"trashed_at,omitempty" xml:"trashed_at,omitempty"`
	Type               *string             `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt          *string             `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UploadId           *string             `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	UserTags           map[string]*string  `json:"user_tags,omitempty" xml:"user_tags,omitempty"`
	VideoMediaMetadata *VideoMediaMetadata `json:"video_media_metadata,omitempty" xml:"video_media_metadata,omitempty"`
}

func (s File) String() string {
	return tea.Prettify(s)
}

func (s File) GoString() string {
	return s.String()
}

func (s *File) SetActionList(v []*string) *File {
	s.ActionList = v
	return s
}

func (s *File) SetCategory(v string) *File {
	s.Category = &v
	return s
}

func (s *File) SetContentHash(v string) *File {
	s.ContentHash = &v
	return s
}

func (s *File) SetContentHashName(v string) *File {
	s.ContentHashName = &v
	return s
}

func (s *File) SetContentType(v string) *File {
	s.ContentType = &v
	return s
}

func (s *File) SetCrc64Hash(v string) *File {
	s.Crc64Hash = &v
	return s
}

func (s *File) SetCreatedAt(v string) *File {
	s.CreatedAt = &v
	return s
}

func (s *File) SetDescription(v string) *File {
	s.Description = &v
	return s
}

func (s *File) SetDomainId(v string) *File {
	s.DomainId = &v
	return s
}

func (s *File) SetDownloadUrl(v string) *File {
	s.DownloadUrl = &v
	return s
}

func (s *File) SetDriveId(v string) *File {
	s.DriveId = &v
	return s
}

func (s *File) SetFileExtension(v string) *File {
	s.FileExtension = &v
	return s
}

func (s *File) SetFileId(v string) *File {
	s.FileId = &v
	return s
}

func (s *File) SetHidden(v bool) *File {
	s.Hidden = &v
	return s
}

func (s *File) SetIdPath(v string) *File {
	s.IdPath = &v
	return s
}

func (s *File) SetImageMediaMetadata(v *ImageMediaMetadata) *File {
	s.ImageMediaMetadata = v
	return s
}

func (s *File) SetLabels(v []*string) *File {
	s.Labels = v
	return s
}

func (s *File) SetLocalCreatedAt(v string) *File {
	s.LocalCreatedAt = &v
	return s
}

func (s *File) SetLocalModifiedAt(v string) *File {
	s.LocalModifiedAt = &v
	return s
}

func (s *File) SetName(v string) *File {
	s.Name = &v
	return s
}

func (s *File) SetNamePath(v string) *File {
	s.NamePath = &v
	return s
}

func (s *File) SetParentFileId(v string) *File {
	s.ParentFileId = &v
	return s
}

func (s *File) SetRevisionId(v string) *File {
	s.RevisionId = &v
	return s
}

func (s *File) SetSize(v int64) *File {
	s.Size = &v
	return s
}

func (s *File) SetStarred(v bool) *File {
	s.Starred = &v
	return s
}

func (s *File) SetStatus(v string) *File {
	s.Status = &v
	return s
}

func (s *File) SetThumbnail(v string) *File {
	s.Thumbnail = &v
	return s
}

func (s *File) SetThumbnailUrls(v map[string]*string) *File {
	s.ThumbnailUrls = v
	return s
}

func (s *File) SetTrashedAt(v string) *File {
	s.TrashedAt = &v
	return s
}

func (s *File) SetType(v string) *File {
	s.Type = &v
	return s
}

func (s *File) SetUpdatedAt(v string) *File {
	s.UpdatedAt = &v
	return s
}

func (s *File) SetUploadId(v string) *File {
	s.UploadId = &v
	return s
}

func (s *File) SetUserTags(v map[string]*string) *File {
	s.UserTags = v
	return s
}

func (s *File) SetVideoMediaMetadata(v *VideoMediaMetadata) *File {
	s.VideoMediaMetadata = v
	return s
}

type FileDownloadCallbackInfo struct {
	// This parameter is required.
	//
	// example:
	//
	// ccp-bj1-bj-1234
	Bucket *string `json:"bucket,omitempty" xml:"bucket,omitempty"`
	// This parameter is required.
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// This parameter is required.
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// This parameter is required.
	Object *string `json:"object,omitempty" xml:"object,omitempty"`
	// This parameter is required.
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s FileDownloadCallbackInfo) String() string {
	return tea.Prettify(s)
}

func (s FileDownloadCallbackInfo) GoString() string {
	return s.String()
}

func (s *FileDownloadCallbackInfo) SetBucket(v string) *FileDownloadCallbackInfo {
	s.Bucket = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetDomainId(v string) *FileDownloadCallbackInfo {
	s.DomainId = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetDriveId(v string) *FileDownloadCallbackInfo {
	s.DriveId = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetFileId(v string) *FileDownloadCallbackInfo {
	s.FileId = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetObject(v string) *FileDownloadCallbackInfo {
	s.Object = &v
	return s
}

func (s *FileDownloadCallbackInfo) SetUserId(v string) *FileDownloadCallbackInfo {
	s.UserId = &v
	return s
}

type FileLogDetail struct {
	DecompressFileList []*string `json:"decompress_file_list,omitempty" xml:"decompress_file_list,omitempty" type:"Repeated"`
	NewName            *string   `json:"new_name,omitempty" xml:"new_name,omitempty"`
	ParentPath         *string   `json:"parent_path,omitempty" xml:"parent_path,omitempty"`
	RevVersion         *int64    `json:"rev_version,omitempty" xml:"rev_version,omitempty"`
	RevisionId         *string   `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	Size               *int64    `json:"size,omitempty" xml:"size,omitempty"`
	ToParentPath       *string   `json:"to_parent_path,omitempty" xml:"to_parent_path,omitempty"`
	ToParentPathType   *string   `json:"to_parent_path_type,omitempty" xml:"to_parent_path_type,omitempty"`
	Type               *string   `json:"type,omitempty" xml:"type,omitempty"`
}

func (s FileLogDetail) String() string {
	return tea.Prettify(s)
}

func (s FileLogDetail) GoString() string {
	return s.String()
}

func (s *FileLogDetail) SetDecompressFileList(v []*string) *FileLogDetail {
	s.DecompressFileList = v
	return s
}

func (s *FileLogDetail) SetNewName(v string) *FileLogDetail {
	s.NewName = &v
	return s
}

func (s *FileLogDetail) SetParentPath(v string) *FileLogDetail {
	s.ParentPath = &v
	return s
}

func (s *FileLogDetail) SetRevVersion(v int64) *FileLogDetail {
	s.RevVersion = &v
	return s
}

func (s *FileLogDetail) SetRevisionId(v string) *FileLogDetail {
	s.RevisionId = &v
	return s
}

func (s *FileLogDetail) SetSize(v int64) *FileLogDetail {
	s.Size = &v
	return s
}

func (s *FileLogDetail) SetToParentPath(v string) *FileLogDetail {
	s.ToParentPath = &v
	return s
}

func (s *FileLogDetail) SetToParentPathType(v string) *FileLogDetail {
	s.ToParentPathType = &v
	return s
}

func (s *FileLogDetail) SetType(v string) *FileLogDetail {
	s.Type = &v
	return s
}

type FilePermissionMember struct {
	ActionList         []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	DisinheritSubGroup *bool     `json:"disinherit_sub_group,omitempty" xml:"disinherit_sub_group,omitempty"`
	ExpireTime         *int64    `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	Identity           *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	RoleId             *string   `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s FilePermissionMember) String() string {
	return tea.Prettify(s)
}

func (s FilePermissionMember) GoString() string {
	return s.String()
}

func (s *FilePermissionMember) SetActionList(v []*string) *FilePermissionMember {
	s.ActionList = v
	return s
}

func (s *FilePermissionMember) SetDisinheritSubGroup(v bool) *FilePermissionMember {
	s.DisinheritSubGroup = &v
	return s
}

func (s *FilePermissionMember) SetExpireTime(v int64) *FilePermissionMember {
	s.ExpireTime = &v
	return s
}

func (s *FilePermissionMember) SetIdentity(v *Identity) *FilePermissionMember {
	s.Identity = v
	return s
}

func (s *FilePermissionMember) SetRoleId(v string) *FilePermissionMember {
	s.RoleId = &v
	return s
}

type FileStreamInfo struct {
	ContentHash     *string         `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName *string         `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	ContentMd5      *string         `json:"content_md5,omitempty" xml:"content_md5,omitempty"`
	PartInfoList    *UploadPartInfo `json:"part_info_list,omitempty" xml:"part_info_list,omitempty"`
	PreHash         *string         `json:"pre_hash,omitempty" xml:"pre_hash,omitempty"`
	ProofCode       *string         `json:"proof_code,omitempty" xml:"proof_code,omitempty"`
	ProofVersion    *string         `json:"proof_version,omitempty" xml:"proof_version,omitempty"`
	Size            *int64          `json:"size,omitempty" xml:"size,omitempty"`
}

func (s FileStreamInfo) String() string {
	return tea.Prettify(s)
}

func (s FileStreamInfo) GoString() string {
	return s.String()
}

func (s *FileStreamInfo) SetContentHash(v string) *FileStreamInfo {
	s.ContentHash = &v
	return s
}

func (s *FileStreamInfo) SetContentHashName(v string) *FileStreamInfo {
	s.ContentHashName = &v
	return s
}

func (s *FileStreamInfo) SetContentMd5(v string) *FileStreamInfo {
	s.ContentMd5 = &v
	return s
}

func (s *FileStreamInfo) SetPartInfoList(v *UploadPartInfo) *FileStreamInfo {
	s.PartInfoList = v
	return s
}

func (s *FileStreamInfo) SetPreHash(v string) *FileStreamInfo {
	s.PreHash = &v
	return s
}

func (s *FileStreamInfo) SetProofCode(v string) *FileStreamInfo {
	s.ProofCode = &v
	return s
}

func (s *FileStreamInfo) SetProofVersion(v string) *FileStreamInfo {
	s.ProofVersion = &v
	return s
}

func (s *FileStreamInfo) SetSize(v int64) *FileStreamInfo {
	s.Size = &v
	return s
}

type GetOfficeEditUrlOption struct {
	Copy     *bool `json:"copy,omitempty" xml:"copy,omitempty"`
	Print    *bool `json:"print,omitempty" xml:"print,omitempty"`
	Readonly *bool `json:"readonly,omitempty" xml:"readonly,omitempty"`
}

func (s GetOfficeEditUrlOption) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeEditUrlOption) GoString() string {
	return s.String()
}

func (s *GetOfficeEditUrlOption) SetCopy(v bool) *GetOfficeEditUrlOption {
	s.Copy = &v
	return s
}

func (s *GetOfficeEditUrlOption) SetPrint(v bool) *GetOfficeEditUrlOption {
	s.Print = &v
	return s
}

func (s *GetOfficeEditUrlOption) SetReadonly(v bool) *GetOfficeEditUrlOption {
	s.Readonly = &v
	return s
}

type GetOfficeEditUrlWatermark struct {
	Fillstyle  *string  `json:"fillstyle,omitempty" xml:"fillstyle,omitempty"`
	Font       *string  `json:"font,omitempty" xml:"font,omitempty"`
	Horizontal *int64   `json:"horizontal,omitempty" xml:"horizontal,omitempty"`
	Rotate     *float64 `json:"rotate,omitempty" xml:"rotate,omitempty"`
	Type       *int32   `json:"type,omitempty" xml:"type,omitempty"`
	Value      *string  `json:"value,omitempty" xml:"value,omitempty"`
	Vertical   *int64   `json:"vertical,omitempty" xml:"vertical,omitempty"`
}

func (s GetOfficeEditUrlWatermark) String() string {
	return tea.Prettify(s)
}

func (s GetOfficeEditUrlWatermark) GoString() string {
	return s.String()
}

func (s *GetOfficeEditUrlWatermark) SetFillstyle(v string) *GetOfficeEditUrlWatermark {
	s.Fillstyle = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetFont(v string) *GetOfficeEditUrlWatermark {
	s.Font = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetHorizontal(v int64) *GetOfficeEditUrlWatermark {
	s.Horizontal = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetRotate(v float64) *GetOfficeEditUrlWatermark {
	s.Rotate = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetType(v int32) *GetOfficeEditUrlWatermark {
	s.Type = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetValue(v string) *GetOfficeEditUrlWatermark {
	s.Value = &v
	return s
}

func (s *GetOfficeEditUrlWatermark) SetVertical(v int64) *GetOfficeEditUrlWatermark {
	s.Vertical = &v
	return s
}

type GetOfficePreviewUrlOption struct {
	Copy  *bool `json:"copy,omitempty" xml:"copy,omitempty"`
	Print *bool `json:"print,omitempty" xml:"print,omitempty"`
}

func (s GetOfficePreviewUrlOption) String() string {
	return tea.Prettify(s)
}

func (s GetOfficePreviewUrlOption) GoString() string {
	return s.String()
}

func (s *GetOfficePreviewUrlOption) SetCopy(v bool) *GetOfficePreviewUrlOption {
	s.Copy = &v
	return s
}

func (s *GetOfficePreviewUrlOption) SetPrint(v bool) *GetOfficePreviewUrlOption {
	s.Print = &v
	return s
}

type Group struct {
	CreatedAt   *int64  `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator     *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	GroupId     *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName   *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	UpdatedAt   *int64  `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s Group) String() string {
	return tea.Prettify(s)
}

func (s Group) GoString() string {
	return s.String()
}

func (s *Group) SetCreatedAt(v int64) *Group {
	s.CreatedAt = &v
	return s
}

func (s *Group) SetCreator(v string) *Group {
	s.Creator = &v
	return s
}

func (s *Group) SetDescription(v string) *Group {
	s.Description = &v
	return s
}

func (s *Group) SetDomainId(v string) *Group {
	s.DomainId = &v
	return s
}

func (s *Group) SetGroupId(v string) *Group {
	s.GroupId = &v
	return s
}

func (s *Group) SetGroupName(v string) *Group {
	s.GroupName = &v
	return s
}

func (s *Group) SetUpdatedAt(v int64) *Group {
	s.UpdatedAt = &v
	return s
}

type HotDriveFile struct {
	// example:
	//
	// 2
	ActionCount *int64    `json:"action_count,omitempty" xml:"action_count,omitempty"`
	ActionList  []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	// example:
	//
	// doc
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 1727059860000
	CountAt *int64 `json:"count_at,omitempty" xml:"count_at,omitempty"`
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// 666ff36c22278f023ec
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// a.jpg
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 666ff36c22278f023ec
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s HotDriveFile) String() string {
	return tea.Prettify(s)
}

func (s HotDriveFile) GoString() string {
	return s.String()
}

func (s *HotDriveFile) SetActionCount(v int64) *HotDriveFile {
	s.ActionCount = &v
	return s
}

func (s *HotDriveFile) SetActionList(v []*string) *HotDriveFile {
	s.ActionList = v
	return s
}

func (s *HotDriveFile) SetCategory(v string) *HotDriveFile {
	s.Category = &v
	return s
}

func (s *HotDriveFile) SetCountAt(v int64) *HotDriveFile {
	s.CountAt = &v
	return s
}

func (s *HotDriveFile) SetDriveId(v string) *HotDriveFile {
	s.DriveId = &v
	return s
}

func (s *HotDriveFile) SetFileId(v string) *HotDriveFile {
	s.FileId = &v
	return s
}

func (s *HotDriveFile) SetName(v string) *HotDriveFile {
	s.Name = &v
	return s
}

func (s *HotDriveFile) SetRevisionId(v string) *HotDriveFile {
	s.RevisionId = &v
	return s
}

type HotKnowledgeBaseFile struct {
	// example:
	//
	// 1
	ActionCount *int64    `json:"action_count,omitempty" xml:"action_count,omitempty"`
	ActionList  []*string `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	// example:
	//
	// image
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// example:
	//
	// 1727578860000
	CountAt *int64 `json:"count_at,omitempty" xml:"count_at,omitempty"`
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// 666ff36c22278f023ec
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 4jTsp3AgW
	KnowledgeBaseId *string `json:"knowledge_base_id,omitempty" xml:"knowledge_base_id,omitempty"`
	// example:
	//
	// a.jpg
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 666ff36c22278f023ec
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s HotKnowledgeBaseFile) String() string {
	return tea.Prettify(s)
}

func (s HotKnowledgeBaseFile) GoString() string {
	return s.String()
}

func (s *HotKnowledgeBaseFile) SetActionCount(v int64) *HotKnowledgeBaseFile {
	s.ActionCount = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetActionList(v []*string) *HotKnowledgeBaseFile {
	s.ActionList = v
	return s
}

func (s *HotKnowledgeBaseFile) SetCategory(v string) *HotKnowledgeBaseFile {
	s.Category = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetCountAt(v int64) *HotKnowledgeBaseFile {
	s.CountAt = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetDriveId(v string) *HotKnowledgeBaseFile {
	s.DriveId = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetFileId(v string) *HotKnowledgeBaseFile {
	s.FileId = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetKnowledgeBaseId(v string) *HotKnowledgeBaseFile {
	s.KnowledgeBaseId = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetName(v string) *HotKnowledgeBaseFile {
	s.Name = &v
	return s
}

func (s *HotKnowledgeBaseFile) SetRevisionId(v string) *HotKnowledgeBaseFile {
	s.RevisionId = &v
	return s
}

type IDPermission struct {
	DisinheritSubGroup *bool  `json:"disinherit_sub_group,omitempty" xml:"disinherit_sub_group,omitempty"`
	ExpireTime         *int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// if can be null:
	// false
	Permission *Permission `json:"permission,omitempty" xml:"permission,omitempty"`
	Roles      []*string   `json:"roles,omitempty" xml:"roles,omitempty" type:"Repeated"`
}

func (s IDPermission) String() string {
	return tea.Prettify(s)
}

func (s IDPermission) GoString() string {
	return s.String()
}

func (s *IDPermission) SetDisinheritSubGroup(v bool) *IDPermission {
	s.DisinheritSubGroup = &v
	return s
}

func (s *IDPermission) SetExpireTime(v int64) *IDPermission {
	s.ExpireTime = &v
	return s
}

func (s *IDPermission) SetPermission(v *Permission) *IDPermission {
	s.Permission = v
	return s
}

func (s *IDPermission) SetRoles(v []*string) *IDPermission {
	s.Roles = v
	return s
}

type Identity struct {
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// example:
	//
	// IT_User
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
}

func (s Identity) String() string {
	return tea.Prettify(s)
}

func (s Identity) GoString() string {
	return s.String()
}

func (s *Identity) SetIdentityId(v string) *Identity {
	s.IdentityId = &v
	return s
}

func (s *Identity) SetIdentityType(v string) *Identity {
	s.IdentityType = &v
	return s
}

type IdentityToBenefitPkgMapping struct {
	BenefitPkgComputationRule *string                   `json:"benefit_pkg_computation_rule,omitempty" xml:"benefit_pkg_computation_rule,omitempty"`
	BenefitPkgId              *string                   `json:"benefit_pkg_id,omitempty" xml:"benefit_pkg_id,omitempty"`
	BenefitPkgName            *string                   `json:"benefit_pkg_name,omitempty" xml:"benefit_pkg_name,omitempty"`
	BenefitPkgOwnerId         *string                   `json:"benefit_pkg_owner_id,omitempty" xml:"benefit_pkg_owner_id,omitempty"`
	BenefitPkgPriority        *int64                    `json:"benefit_pkg_priority,omitempty" xml:"benefit_pkg_priority,omitempty"`
	BenefitPkgType            *string                   `json:"benefit_pkg_type,omitempty" xml:"benefit_pkg_type,omitempty"`
	CreatedAt                 *string                   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	DeliveryInfoList          []*BenefitPkgDeliveryInfo `json:"delivery_info_list,omitempty" xml:"delivery_info_list,omitempty" type:"Repeated"`
	IdentityId                *string                   `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	IdentityType              *string                   `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
	UpdatedAt                 *string                   `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s IdentityToBenefitPkgMapping) String() string {
	return tea.Prettify(s)
}

func (s IdentityToBenefitPkgMapping) GoString() string {
	return s.String()
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgComputationRule(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgComputationRule = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgId(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgId = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgName(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgName = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgOwnerId(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgOwnerId = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgPriority(v int64) *IdentityToBenefitPkgMapping {
	s.BenefitPkgPriority = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetBenefitPkgType(v string) *IdentityToBenefitPkgMapping {
	s.BenefitPkgType = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetCreatedAt(v string) *IdentityToBenefitPkgMapping {
	s.CreatedAt = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetDeliveryInfoList(v []*BenefitPkgDeliveryInfo) *IdentityToBenefitPkgMapping {
	s.DeliveryInfoList = v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetIdentityId(v string) *IdentityToBenefitPkgMapping {
	s.IdentityId = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetIdentityType(v string) *IdentityToBenefitPkgMapping {
	s.IdentityType = &v
	return s
}

func (s *IdentityToBenefitPkgMapping) SetUpdatedAt(v string) *IdentityToBenefitPkgMapping {
	s.UpdatedAt = &v
	return s
}

type ImageMediaMetadata struct {
	// example:
	//
	// 浙江省杭州市滨江区西兴街道江陵路
	AddressLine *string `json:"address_line,omitempty" xml:"address_line,omitempty"`
	// example:
	//
	// 杭州市
	City *string `json:"city,omitempty" xml:"city,omitempty"`
	// example:
	//
	// 中国
	Country *string `json:"country,omitempty" xml:"country,omitempty"`
	// example:
	//
	// 滨江区
	District *string `json:"district,omitempty" xml:"district,omitempty"`
	// example:
	//
	// {"Compression":{"value":"6"},"DateTime":{"value":"2020:08:19 17:11:11"}}
	Exif           *string          `json:"exif,omitempty" xml:"exif,omitempty"`
	FacesThumbnail []*FaceThumbnail `json:"faces_thumbnail,omitempty" xml:"faces_thumbnail,omitempty" type:"Repeated"`
	// example:
	//
	// 1080
	Height       *int64        `json:"height,omitempty" xml:"height,omitempty"`
	ImageQuality *ImageQuality `json:"image_quality,omitempty" xml:"image_quality,omitempty"`
	ImageTags    []*SystemTag  `json:"image_tags,omitempty" xml:"image_tags,omitempty" type:"Repeated"`
	// example:
	//
	// 30.185453,120.218522
	Location *string `json:"location,omitempty" xml:"location,omitempty"`
	// example:
	//
	// 浙江省
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
	// example:
	//
	// 2006-01-02T15:04:05.000Z07:00
	Time *string `json:"time,omitempty" xml:"time,omitempty"`
	// example:
	//
	// 西兴街道
	Township *string `json:"township,omitempty" xml:"township,omitempty"`
	// example:
	//
	// 1920
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ImageMediaMetadata) String() string {
	return tea.Prettify(s)
}

func (s ImageMediaMetadata) GoString() string {
	return s.String()
}

func (s *ImageMediaMetadata) SetAddressLine(v string) *ImageMediaMetadata {
	s.AddressLine = &v
	return s
}

func (s *ImageMediaMetadata) SetCity(v string) *ImageMediaMetadata {
	s.City = &v
	return s
}

func (s *ImageMediaMetadata) SetCountry(v string) *ImageMediaMetadata {
	s.Country = &v
	return s
}

func (s *ImageMediaMetadata) SetDistrict(v string) *ImageMediaMetadata {
	s.District = &v
	return s
}

func (s *ImageMediaMetadata) SetExif(v string) *ImageMediaMetadata {
	s.Exif = &v
	return s
}

func (s *ImageMediaMetadata) SetFacesThumbnail(v []*FaceThumbnail) *ImageMediaMetadata {
	s.FacesThumbnail = v
	return s
}

func (s *ImageMediaMetadata) SetHeight(v int64) *ImageMediaMetadata {
	s.Height = &v
	return s
}

func (s *ImageMediaMetadata) SetImageQuality(v *ImageQuality) *ImageMediaMetadata {
	s.ImageQuality = v
	return s
}

func (s *ImageMediaMetadata) SetImageTags(v []*SystemTag) *ImageMediaMetadata {
	s.ImageTags = v
	return s
}

func (s *ImageMediaMetadata) SetLocation(v string) *ImageMediaMetadata {
	s.Location = &v
	return s
}

func (s *ImageMediaMetadata) SetProvince(v string) *ImageMediaMetadata {
	s.Province = &v
	return s
}

func (s *ImageMediaMetadata) SetTime(v string) *ImageMediaMetadata {
	s.Time = &v
	return s
}

func (s *ImageMediaMetadata) SetTownship(v string) *ImageMediaMetadata {
	s.Township = &v
	return s
}

func (s *ImageMediaMetadata) SetWidth(v int64) *ImageMediaMetadata {
	s.Width = &v
	return s
}

type ImageProcess struct {
	ImageThumbnailProcess  *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	OfficeThumbnailProcess *string `json:"office_thumbnail_process,omitempty" xml:"office_thumbnail_process,omitempty"`
	VideoThumbnailProcess  *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s ImageProcess) String() string {
	return tea.Prettify(s)
}

func (s ImageProcess) GoString() string {
	return s.String()
}

func (s *ImageProcess) SetImageThumbnailProcess(v string) *ImageProcess {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *ImageProcess) SetOfficeThumbnailProcess(v string) *ImageProcess {
	s.OfficeThumbnailProcess = &v
	return s
}

func (s *ImageProcess) SetVideoThumbnailProcess(v string) *ImageProcess {
	s.VideoThumbnailProcess = &v
	return s
}

type ImageQuality struct {
	// example:
	//
	// 0.736
	OverallScore *float64 `json:"overall_score,omitempty" xml:"overall_score,omitempty"`
}

func (s ImageQuality) String() string {
	return tea.Prettify(s)
}

func (s ImageQuality) GoString() string {
	return s.String()
}

func (s *ImageQuality) SetOverallScore(v float64) *ImageQuality {
	s.OverallScore = &v
	return s
}

type ImageTag struct {
	// example:
	//
	// 10
	Count *int64 `json:"count,omitempty" xml:"count,omitempty"`
	// example:
	//
	// image
	CoverFileCategory *string `json:"cover_file_category,omitempty" xml:"cover_file_category,omitempty"`
	// example:
	//
	// 5d79206586bb5dd69fb34c349282718146c55da7
	CoverFileId *string `json:"cover_file_id,omitempty" xml:"cover_file_id,omitempty"`
	// example:
	//
	// 0.736
	CoverOverallScore *float32 `json:"cover_overall_score,omitempty" xml:"cover_overall_score,omitempty"`
	// example:
	//
	// 1
	CoverTagConfidence *float32 `json:"cover_tag_confidence,omitempty" xml:"cover_tag_confidence,omitempty"`
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	CoverUrl *string `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	// example:
	//
	// 动物
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s ImageTag) String() string {
	return tea.Prettify(s)
}

func (s ImageTag) GoString() string {
	return s.String()
}

func (s *ImageTag) SetCount(v int64) *ImageTag {
	s.Count = &v
	return s
}

func (s *ImageTag) SetCoverFileCategory(v string) *ImageTag {
	s.CoverFileCategory = &v
	return s
}

func (s *ImageTag) SetCoverFileId(v string) *ImageTag {
	s.CoverFileId = &v
	return s
}

func (s *ImageTag) SetCoverOverallScore(v float32) *ImageTag {
	s.CoverOverallScore = &v
	return s
}

func (s *ImageTag) SetCoverTagConfidence(v float32) *ImageTag {
	s.CoverTagConfidence = &v
	return s
}

func (s *ImageTag) SetCoverUrl(v string) *ImageTag {
	s.CoverUrl = &v
	return s
}

func (s *ImageTag) SetName(v string) *ImageTag {
	s.Name = &v
	return s
}

type Int64Range struct {
	From *int64 `json:"from,omitempty" xml:"from,omitempty"`
	To   *int64 `json:"to,omitempty" xml:"to,omitempty"`
}

func (s Int64Range) String() string {
	return tea.Prettify(s)
}

func (s Int64Range) GoString() string {
	return s.String()
}

func (s *Int64Range) SetFrom(v int64) *Int64Range {
	s.From = &v
	return s
}

func (s *Int64Range) SetTo(v int64) *Int64Range {
	s.To = &v
	return s
}

type InvestigationInfo struct {
	Status      *int64                        `json:"status,omitempty" xml:"status,omitempty"`
	Suggestion  *string                       `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
	VideoDetail *InvestigationInfoVideoDetail `json:"video_detail,omitempty" xml:"video_detail,omitempty" type:"Struct"`
}

func (s InvestigationInfo) String() string {
	return tea.Prettify(s)
}

func (s InvestigationInfo) GoString() string {
	return s.String()
}

func (s *InvestigationInfo) SetStatus(v int64) *InvestigationInfo {
	s.Status = &v
	return s
}

func (s *InvestigationInfo) SetSuggestion(v string) *InvestigationInfo {
	s.Suggestion = &v
	return s
}

func (s *InvestigationInfo) SetVideoDetail(v *InvestigationInfoVideoDetail) *InvestigationInfo {
	s.VideoDetail = v
	return s
}

type InvestigationInfoVideoDetail struct {
	BlockFrames []*InvestigationInfoVideoDetailBlockFrames `json:"block_frames,omitempty" xml:"block_frames,omitempty" type:"Repeated"`
}

func (s InvestigationInfoVideoDetail) String() string {
	return tea.Prettify(s)
}

func (s InvestigationInfoVideoDetail) GoString() string {
	return s.String()
}

func (s *InvestigationInfoVideoDetail) SetBlockFrames(v []*InvestigationInfoVideoDetailBlockFrames) *InvestigationInfoVideoDetail {
	s.BlockFrames = v
	return s
}

type InvestigationInfoVideoDetailBlockFrames struct {
	Label  *string  `json:"label,omitempty" xml:"label,omitempty"`
	Offset *int64   `json:"offset,omitempty" xml:"offset,omitempty"`
	Rate   *float64 `json:"rate,omitempty" xml:"rate,omitempty"`
}

func (s InvestigationInfoVideoDetailBlockFrames) String() string {
	return tea.Prettify(s)
}

func (s InvestigationInfoVideoDetailBlockFrames) GoString() string {
	return s.String()
}

func (s *InvestigationInfoVideoDetailBlockFrames) SetLabel(v string) *InvestigationInfoVideoDetailBlockFrames {
	s.Label = &v
	return s
}

func (s *InvestigationInfoVideoDetailBlockFrames) SetOffset(v int64) *InvestigationInfoVideoDetailBlockFrames {
	s.Offset = &v
	return s
}

func (s *InvestigationInfoVideoDetailBlockFrames) SetRate(v float64) *InvestigationInfoVideoDetailBlockFrames {
	s.Rate = &v
	return s
}

type JWTPayload struct {
	Aud        *string `json:"aud,omitempty" xml:"aud,omitempty"`
	AutoCreate *bool   `json:"auto_create,omitempty" xml:"auto_create,omitempty"`
	Exp        *int64  `json:"exp,omitempty" xml:"exp,omitempty"`
	Iat        *int64  `json:"iat,omitempty" xml:"iat,omitempty"`
	Iss        *string `json:"iss,omitempty" xml:"iss,omitempty"`
	Jti        *string `json:"jti,omitempty" xml:"jti,omitempty"`
	Nbf        *int64  `json:"nbf,omitempty" xml:"nbf,omitempty"`
	Sub        *string `json:"sub,omitempty" xml:"sub,omitempty"`
	SubType    *string `json:"sub_type,omitempty" xml:"sub_type,omitempty"`
}

func (s JWTPayload) String() string {
	return tea.Prettify(s)
}

func (s JWTPayload) GoString() string {
	return s.String()
}

func (s *JWTPayload) SetAud(v string) *JWTPayload {
	s.Aud = &v
	return s
}

func (s *JWTPayload) SetAutoCreate(v bool) *JWTPayload {
	s.AutoCreate = &v
	return s
}

func (s *JWTPayload) SetExp(v int64) *JWTPayload {
	s.Exp = &v
	return s
}

func (s *JWTPayload) SetIat(v int64) *JWTPayload {
	s.Iat = &v
	return s
}

func (s *JWTPayload) SetIss(v string) *JWTPayload {
	s.Iss = &v
	return s
}

func (s *JWTPayload) SetJti(v string) *JWTPayload {
	s.Jti = &v
	return s
}

func (s *JWTPayload) SetNbf(v int64) *JWTPayload {
	s.Nbf = &v
	return s
}

func (s *JWTPayload) SetSub(v string) *JWTPayload {
	s.Sub = &v
	return s
}

func (s *JWTPayload) SetSubType(v string) *JWTPayload {
	s.SubType = &v
	return s
}

type KnowledgeBase struct {
	CoverUri        *string     `json:"cover_uri,omitempty" xml:"cover_uri,omitempty"`
	CreatedAt       *int64      `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description     *string     `json:"description,omitempty" xml:"description,omitempty"`
	FileFilter      *string     `json:"file_filter,omitempty" xml:"file_filter,omitempty"`
	KnowledgeBaseId *string     `json:"knowledge_base_id,omitempty" xml:"knowledge_base_id,omitempty"`
	LinkRuleList    []*LinkRule `json:"link_rule_list,omitempty" xml:"link_rule_list,omitempty" type:"Repeated"`
	Name            *string     `json:"name,omitempty" xml:"name,omitempty"`
	OwnerId         *string     `json:"owner_id,omitempty" xml:"owner_id,omitempty"`
	OwnerName       *string     `json:"owner_name,omitempty" xml:"owner_name,omitempty"`
	OwnerType       *string     `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	UpdatedAt       *int64      `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s KnowledgeBase) String() string {
	return tea.Prettify(s)
}

func (s KnowledgeBase) GoString() string {
	return s.String()
}

func (s *KnowledgeBase) SetCoverUri(v string) *KnowledgeBase {
	s.CoverUri = &v
	return s
}

func (s *KnowledgeBase) SetCreatedAt(v int64) *KnowledgeBase {
	s.CreatedAt = &v
	return s
}

func (s *KnowledgeBase) SetDescription(v string) *KnowledgeBase {
	s.Description = &v
	return s
}

func (s *KnowledgeBase) SetFileFilter(v string) *KnowledgeBase {
	s.FileFilter = &v
	return s
}

func (s *KnowledgeBase) SetKnowledgeBaseId(v string) *KnowledgeBase {
	s.KnowledgeBaseId = &v
	return s
}

func (s *KnowledgeBase) SetLinkRuleList(v []*LinkRule) *KnowledgeBase {
	s.LinkRuleList = v
	return s
}

func (s *KnowledgeBase) SetName(v string) *KnowledgeBase {
	s.Name = &v
	return s
}

func (s *KnowledgeBase) SetOwnerId(v string) *KnowledgeBase {
	s.OwnerId = &v
	return s
}

func (s *KnowledgeBase) SetOwnerName(v string) *KnowledgeBase {
	s.OwnerName = &v
	return s
}

func (s *KnowledgeBase) SetOwnerType(v string) *KnowledgeBase {
	s.OwnerType = &v
	return s
}

func (s *KnowledgeBase) SetUpdatedAt(v int64) *KnowledgeBase {
	s.UpdatedAt = &v
	return s
}

type KnowledgeCategory struct {
	CreatedAt                 *int64    `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description               *string   `json:"description,omitempty" xml:"description,omitempty"`
	Keywords                  []*string `json:"keywords,omitempty" xml:"keywords,omitempty" type:"Repeated"`
	KnowledgeBaseId           *string   `json:"knowledge_base_id,omitempty" xml:"knowledge_base_id,omitempty"`
	KnowledgeBaseName         *string   `json:"knowledge_base_name,omitempty" xml:"knowledge_base_name,omitempty"`
	KnowledgeCategoryId       *string   `json:"knowledge_category_id,omitempty" xml:"knowledge_category_id,omitempty"`
	Name                      *string   `json:"name,omitempty" xml:"name,omitempty"`
	Owner                     *string   `json:"owner,omitempty" xml:"owner,omitempty"`
	OwnerType                 *string   `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	ParentKnowledgeCategoryId *string   `json:"parent_knowledge_category_id,omitempty" xml:"parent_knowledge_category_id,omitempty"`
	Status                    *string   `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt                 *int64    `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s KnowledgeCategory) String() string {
	return tea.Prettify(s)
}

func (s KnowledgeCategory) GoString() string {
	return s.String()
}

func (s *KnowledgeCategory) SetCreatedAt(v int64) *KnowledgeCategory {
	s.CreatedAt = &v
	return s
}

func (s *KnowledgeCategory) SetDescription(v string) *KnowledgeCategory {
	s.Description = &v
	return s
}

func (s *KnowledgeCategory) SetKeywords(v []*string) *KnowledgeCategory {
	s.Keywords = v
	return s
}

func (s *KnowledgeCategory) SetKnowledgeBaseId(v string) *KnowledgeCategory {
	s.KnowledgeBaseId = &v
	return s
}

func (s *KnowledgeCategory) SetKnowledgeBaseName(v string) *KnowledgeCategory {
	s.KnowledgeBaseName = &v
	return s
}

func (s *KnowledgeCategory) SetKnowledgeCategoryId(v string) *KnowledgeCategory {
	s.KnowledgeCategoryId = &v
	return s
}

func (s *KnowledgeCategory) SetName(v string) *KnowledgeCategory {
	s.Name = &v
	return s
}

func (s *KnowledgeCategory) SetOwner(v string) *KnowledgeCategory {
	s.Owner = &v
	return s
}

func (s *KnowledgeCategory) SetOwnerType(v string) *KnowledgeCategory {
	s.OwnerType = &v
	return s
}

func (s *KnowledgeCategory) SetParentKnowledgeCategoryId(v string) *KnowledgeCategory {
	s.ParentKnowledgeCategoryId = &v
	return s
}

func (s *KnowledgeCategory) SetStatus(v string) *KnowledgeCategory {
	s.Status = &v
	return s
}

func (s *KnowledgeCategory) SetUpdatedAt(v int64) *KnowledgeCategory {
	s.UpdatedAt = &v
	return s
}

type KnowledgeFile struct {
	CreatorId            *string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	DriveId              *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	DriveName            *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	FileCategory         *string `json:"file_category,omitempty" xml:"file_category,omitempty"`
	FileCreatedAt        *int64  `json:"file_created_at,omitempty" xml:"file_created_at,omitempty"`
	FileCreatorId        *string `json:"file_creator_id,omitempty" xml:"file_creator_id,omitempty"`
	FileId               *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	FileImageTime        *int64  `json:"file_image_time,omitempty" xml:"file_image_time,omitempty"`
	FileLastModifierId   *string `json:"file_last_modifier_id,omitempty" xml:"file_last_modifier_id,omitempty"`
	FileLastModifierType *string `json:"file_last_modifier_type,omitempty" xml:"file_last_modifier_type,omitempty"`
	FileName             *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	FileNamePath         *string `json:"file_name_path,omitempty" xml:"file_name_path,omitempty"`
	FileSize             *int64  `json:"file_size,omitempty" xml:"file_size,omitempty"`
	FileUpdatedAt        *int64  `json:"file_updated_at,omitempty" xml:"file_updated_at,omitempty"`
	JoinedAt             *int64  `json:"joined_at,omitempty" xml:"joined_at,omitempty"`
	KnowledgeBaseId      *string `json:"knowledge_base_id,omitempty" xml:"knowledge_base_id,omitempty"`
	KnowledgeCategoryId  *string `json:"knowledge_category_id,omitempty" xml:"knowledge_category_id,omitempty"`
	RevisionId           *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s KnowledgeFile) String() string {
	return tea.Prettify(s)
}

func (s KnowledgeFile) GoString() string {
	return s.String()
}

func (s *KnowledgeFile) SetCreatorId(v string) *KnowledgeFile {
	s.CreatorId = &v
	return s
}

func (s *KnowledgeFile) SetDriveId(v string) *KnowledgeFile {
	s.DriveId = &v
	return s
}

func (s *KnowledgeFile) SetDriveName(v string) *KnowledgeFile {
	s.DriveName = &v
	return s
}

func (s *KnowledgeFile) SetFileCategory(v string) *KnowledgeFile {
	s.FileCategory = &v
	return s
}

func (s *KnowledgeFile) SetFileCreatedAt(v int64) *KnowledgeFile {
	s.FileCreatedAt = &v
	return s
}

func (s *KnowledgeFile) SetFileCreatorId(v string) *KnowledgeFile {
	s.FileCreatorId = &v
	return s
}

func (s *KnowledgeFile) SetFileId(v string) *KnowledgeFile {
	s.FileId = &v
	return s
}

func (s *KnowledgeFile) SetFileImageTime(v int64) *KnowledgeFile {
	s.FileImageTime = &v
	return s
}

func (s *KnowledgeFile) SetFileLastModifierId(v string) *KnowledgeFile {
	s.FileLastModifierId = &v
	return s
}

func (s *KnowledgeFile) SetFileLastModifierType(v string) *KnowledgeFile {
	s.FileLastModifierType = &v
	return s
}

func (s *KnowledgeFile) SetFileName(v string) *KnowledgeFile {
	s.FileName = &v
	return s
}

func (s *KnowledgeFile) SetFileNamePath(v string) *KnowledgeFile {
	s.FileNamePath = &v
	return s
}

func (s *KnowledgeFile) SetFileSize(v int64) *KnowledgeFile {
	s.FileSize = &v
	return s
}

func (s *KnowledgeFile) SetFileUpdatedAt(v int64) *KnowledgeFile {
	s.FileUpdatedAt = &v
	return s
}

func (s *KnowledgeFile) SetJoinedAt(v int64) *KnowledgeFile {
	s.JoinedAt = &v
	return s
}

func (s *KnowledgeFile) SetKnowledgeBaseId(v string) *KnowledgeFile {
	s.KnowledgeBaseId = &v
	return s
}

func (s *KnowledgeFile) SetKnowledgeCategoryId(v string) *KnowledgeFile {
	s.KnowledgeCategoryId = &v
	return s
}

func (s *KnowledgeFile) SetRevisionId(v string) *KnowledgeFile {
	s.RevisionId = &v
	return s
}

type KnowledgeFileItem struct {
	// This parameter is required.
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s KnowledgeFileItem) String() string {
	return tea.Prettify(s)
}

func (s KnowledgeFileItem) GoString() string {
	return s.String()
}

func (s *KnowledgeFileItem) SetDriveId(v string) *KnowledgeFileItem {
	s.DriveId = &v
	return s
}

func (s *KnowledgeFileItem) SetFileId(v string) *KnowledgeFileItem {
	s.FileId = &v
	return s
}

type LinkInfo struct {
	Extra    *string `json:"extra,omitempty" xml:"extra,omitempty"`
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s LinkInfo) String() string {
	return tea.Prettify(s)
}

func (s LinkInfo) GoString() string {
	return s.String()
}

func (s *LinkInfo) SetExtra(v string) *LinkInfo {
	s.Extra = &v
	return s
}

func (s *LinkInfo) SetIdentity(v string) *LinkInfo {
	s.Identity = &v
	return s
}

func (s *LinkInfo) SetType(v string) *LinkInfo {
	s.Type = &v
	return s
}

type LinkRule struct {
	LinkType     *string `json:"link_type,omitempty" xml:"link_type,omitempty"`
	SrcDriveId   *string `json:"src_drive_id,omitempty" xml:"src_drive_id,omitempty"`
	SrcDriveName *string `json:"src_drive_name,omitempty" xml:"src_drive_name,omitempty"`
	SrcFileId    *string `json:"src_file_id,omitempty" xml:"src_file_id,omitempty"`
	SrcFileName  *string `json:"src_file_name,omitempty" xml:"src_file_name,omitempty"`
	SrcValid     *bool   `json:"src_valid,omitempty" xml:"src_valid,omitempty"`
}

func (s LinkRule) String() string {
	return tea.Prettify(s)
}

func (s LinkRule) GoString() string {
	return s.String()
}

func (s *LinkRule) SetLinkType(v string) *LinkRule {
	s.LinkType = &v
	return s
}

func (s *LinkRule) SetSrcDriveId(v string) *LinkRule {
	s.SrcDriveId = &v
	return s
}

func (s *LinkRule) SetSrcDriveName(v string) *LinkRule {
	s.SrcDriveName = &v
	return s
}

func (s *LinkRule) SetSrcFileId(v string) *LinkRule {
	s.SrcFileId = &v
	return s
}

func (s *LinkRule) SetSrcFileName(v string) *LinkRule {
	s.SrcFileName = &v
	return s
}

func (s *LinkRule) SetSrcValid(v bool) *LinkRule {
	s.SrcValid = &v
	return s
}

type LocationDateCluster struct {
	Address      *Address           `json:"address,omitempty" xml:"address,omitempty"`
	ClusterId    *string            `json:"cluster_id,omitempty" xml:"cluster_id,omitempty"`
	CreatedAt    *string            `json:"created_at,omitempty" xml:"created_at,omitempty"`
	CustomLabels map[string]*string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	DriveId      *string            `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	EndTime      *string            `json:"end_time,omitempty" xml:"end_time,omitempty"`
	Level        *string            `json:"level,omitempty" xml:"level,omitempty"`
	StartTime    *string            `json:"start_time,omitempty" xml:"start_time,omitempty"`
	Title        *string            `json:"title,omitempty" xml:"title,omitempty"`
	UpdatedAt    *string            `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s LocationDateCluster) String() string {
	return tea.Prettify(s)
}

func (s LocationDateCluster) GoString() string {
	return s.String()
}

func (s *LocationDateCluster) SetAddress(v *Address) *LocationDateCluster {
	s.Address = v
	return s
}

func (s *LocationDateCluster) SetClusterId(v string) *LocationDateCluster {
	s.ClusterId = &v
	return s
}

func (s *LocationDateCluster) SetCreatedAt(v string) *LocationDateCluster {
	s.CreatedAt = &v
	return s
}

func (s *LocationDateCluster) SetCustomLabels(v map[string]*string) *LocationDateCluster {
	s.CustomLabels = v
	return s
}

func (s *LocationDateCluster) SetDriveId(v string) *LocationDateCluster {
	s.DriveId = &v
	return s
}

func (s *LocationDateCluster) SetEndTime(v string) *LocationDateCluster {
	s.EndTime = &v
	return s
}

func (s *LocationDateCluster) SetLevel(v string) *LocationDateCluster {
	s.Level = &v
	return s
}

func (s *LocationDateCluster) SetStartTime(v string) *LocationDateCluster {
	s.StartTime = &v
	return s
}

func (s *LocationDateCluster) SetTitle(v string) *LocationDateCluster {
	s.Title = &v
	return s
}

func (s *LocationDateCluster) SetUpdatedAt(v string) *LocationDateCluster {
	s.UpdatedAt = &v
	return s
}

type Membership struct {
	CreatedAt   *int64  `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator     *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	GroupId     *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	MemberRole  *string `json:"member_role,omitempty" xml:"member_role,omitempty"`
	MemberType  *string `json:"member_type,omitempty" xml:"member_type,omitempty"`
	SubGroupId  *string `json:"sub_group_id,omitempty" xml:"sub_group_id,omitempty"`
	UpdatedAt   *int64  `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UserId      *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s Membership) String() string {
	return tea.Prettify(s)
}

func (s Membership) GoString() string {
	return s.String()
}

func (s *Membership) SetCreatedAt(v int64) *Membership {
	s.CreatedAt = &v
	return s
}

func (s *Membership) SetCreator(v string) *Membership {
	s.Creator = &v
	return s
}

func (s *Membership) SetDescription(v string) *Membership {
	s.Description = &v
	return s
}

func (s *Membership) SetDomainId(v string) *Membership {
	s.DomainId = &v
	return s
}

func (s *Membership) SetGroupId(v string) *Membership {
	s.GroupId = &v
	return s
}

func (s *Membership) SetMemberRole(v string) *Membership {
	s.MemberRole = &v
	return s
}

func (s *Membership) SetMemberType(v string) *Membership {
	s.MemberType = &v
	return s
}

func (s *Membership) SetSubGroupId(v string) *Membership {
	s.SubGroupId = &v
	return s
}

func (s *Membership) SetUpdatedAt(v int64) *Membership {
	s.UpdatedAt = &v
	return s
}

func (s *Membership) SetUserId(v string) *Membership {
	s.UserId = &v
	return s
}

type NameCheckResult struct {
	ExistFileId   *string `json:"exist_file_id,omitempty" xml:"exist_file_id,omitempty"`
	ExistFileType *string `json:"exist_file_type,omitempty" xml:"exist_file_type,omitempty"`
}

func (s NameCheckResult) String() string {
	return tea.Prettify(s)
}

func (s NameCheckResult) GoString() string {
	return s.String()
}

func (s *NameCheckResult) SetExistFileId(v string) *NameCheckResult {
	s.ExistFileId = &v
	return s
}

func (s *NameCheckResult) SetExistFileType(v string) *NameCheckResult {
	s.ExistFileType = &v
	return s
}

type OfficeEditConfig struct {
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s OfficeEditConfig) String() string {
	return tea.Prettify(s)
}

func (s OfficeEditConfig) GoString() string {
	return s.String()
}

func (s *OfficeEditConfig) SetEnabled(v bool) *OfficeEditConfig {
	s.Enabled = &v
	return s
}

type OfficePreviewConfig struct {
	Enabled *bool `json:"enabled,omitempty" xml:"enabled,omitempty"`
}

func (s OfficePreviewConfig) String() string {
	return tea.Prettify(s)
}

func (s OfficePreviewConfig) GoString() string {
	return s.String()
}

func (s *OfficePreviewConfig) SetEnabled(v bool) *OfficePreviewConfig {
	s.Enabled = &v
	return s
}

type Permission struct {
	ActionList   []*PermissionActionList `json:"action_list,omitempty" xml:"action_list,omitempty" type:"Repeated"`
	Collection   *string                 `json:"collection,omitempty" xml:"collection,omitempty"`
	Condition    *PermissionCondition    `json:"condition,omitempty" xml:"condition,omitempty"`
	CreatedAt    *int64                  `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Effect       *string                 `json:"effect,omitempty" xml:"effect,omitempty"`
	IdentityId   *string                 `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	IdentityType *string                 `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
	Resource     *string                 `json:"resource,omitempty" xml:"resource,omitempty"`
	ResourceType *string                 `json:"resource_type,omitempty" xml:"resource_type,omitempty"`
	UpdatedAt    *int64                  `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UserTags     []*string               `json:"user_tags,omitempty" xml:"user_tags,omitempty" type:"Repeated"`
}

func (s Permission) String() string {
	return tea.Prettify(s)
}

func (s Permission) GoString() string {
	return s.String()
}

func (s *Permission) SetActionList(v []*PermissionActionList) *Permission {
	s.ActionList = v
	return s
}

func (s *Permission) SetCollection(v string) *Permission {
	s.Collection = &v
	return s
}

func (s *Permission) SetCondition(v *PermissionCondition) *Permission {
	s.Condition = v
	return s
}

func (s *Permission) SetCreatedAt(v int64) *Permission {
	s.CreatedAt = &v
	return s
}

func (s *Permission) SetEffect(v string) *Permission {
	s.Effect = &v
	return s
}

func (s *Permission) SetIdentityId(v string) *Permission {
	s.IdentityId = &v
	return s
}

func (s *Permission) SetIdentityType(v string) *Permission {
	s.IdentityType = &v
	return s
}

func (s *Permission) SetResource(v string) *Permission {
	s.Resource = &v
	return s
}

func (s *Permission) SetResourceType(v string) *Permission {
	s.ResourceType = &v
	return s
}

func (s *Permission) SetUpdatedAt(v int64) *Permission {
	s.UpdatedAt = &v
	return s
}

func (s *Permission) SetUserTags(v []*string) *Permission {
	s.UserTags = v
	return s
}

type PermissionActionList struct {
	Action *string `json:"action,omitempty" xml:"action,omitempty"`
}

func (s PermissionActionList) String() string {
	return tea.Prettify(s)
}

func (s PermissionActionList) GoString() string {
	return s.String()
}

func (s *PermissionActionList) SetAction(v string) *PermissionActionList {
	s.Action = &v
	return s
}

type PermissionCondition struct {
	IpEquals      *PermissionConditionIpEquals      `json:"ip_equals,omitempty" xml:"ip_equals,omitempty" type:"Struct"`
	IpNotEquals   *PermissionConditionIpNotEquals   `json:"ip_not_equals,omitempty" xml:"ip_not_equals,omitempty" type:"Struct"`
	StringLike    *PermissionConditionStringLike    `json:"string_like,omitempty" xml:"string_like,omitempty" type:"Struct"`
	StringNotLike *PermissionConditionStringNotLike `json:"string_not_like,omitempty" xml:"string_not_like,omitempty" type:"Struct"`
}

func (s PermissionCondition) String() string {
	return tea.Prettify(s)
}

func (s PermissionCondition) GoString() string {
	return s.String()
}

func (s *PermissionCondition) SetIpEquals(v *PermissionConditionIpEquals) *PermissionCondition {
	s.IpEquals = v
	return s
}

func (s *PermissionCondition) SetIpNotEquals(v *PermissionConditionIpNotEquals) *PermissionCondition {
	s.IpNotEquals = v
	return s
}

func (s *PermissionCondition) SetStringLike(v *PermissionConditionStringLike) *PermissionCondition {
	s.StringLike = v
	return s
}

func (s *PermissionCondition) SetStringNotLike(v *PermissionConditionStringNotLike) *PermissionCondition {
	s.StringNotLike = v
	return s
}

type PermissionConditionIpEquals struct {
	ClientIp []*string `json:"client_ip,omitempty" xml:"client_ip,omitempty" type:"Repeated"`
}

func (s PermissionConditionIpEquals) String() string {
	return tea.Prettify(s)
}

func (s PermissionConditionIpEquals) GoString() string {
	return s.String()
}

func (s *PermissionConditionIpEquals) SetClientIp(v []*string) *PermissionConditionIpEquals {
	s.ClientIp = v
	return s
}

type PermissionConditionIpNotEquals struct {
	ClientIp []*string `json:"client_ip,omitempty" xml:"client_ip,omitempty" type:"Repeated"`
}

func (s PermissionConditionIpNotEquals) String() string {
	return tea.Prettify(s)
}

func (s PermissionConditionIpNotEquals) GoString() string {
	return s.String()
}

func (s *PermissionConditionIpNotEquals) SetClientIp(v []*string) *PermissionConditionIpNotEquals {
	s.ClientIp = v
	return s
}

type PermissionConditionStringLike struct {
	VpcId []*string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty" type:"Repeated"`
}

func (s PermissionConditionStringLike) String() string {
	return tea.Prettify(s)
}

func (s PermissionConditionStringLike) GoString() string {
	return s.String()
}

func (s *PermissionConditionStringLike) SetVpcId(v []*string) *PermissionConditionStringLike {
	s.VpcId = v
	return s
}

type PermissionConditionStringNotLike struct {
	VpcId []*string `json:"vpc_id,omitempty" xml:"vpc_id,omitempty" type:"Repeated"`
}

func (s PermissionConditionStringNotLike) String() string {
	return tea.Prettify(s)
}

func (s PermissionConditionStringNotLike) GoString() string {
	return s.String()
}

func (s *PermissionConditionStringNotLike) SetVpcId(v []*string) *PermissionConditionStringNotLike {
	s.VpcId = v
	return s
}

type PersonalRightsInfoResponse struct {
	ExpiresTime         *string                     `json:"expires_time,omitempty" xml:"expires_time,omitempty"`
	HistoryLatestRights *PersonalRightsInfoResponse `json:"history_latest_rights,omitempty" xml:"history_latest_rights,omitempty"`
	Icon                *string                     `json:"icon,omitempty" xml:"icon,omitempty"`
	IsExpires           *bool                       `json:"is_expires,omitempty" xml:"is_expires,omitempty"`
	Name                *string                     `json:"name,omitempty" xml:"name,omitempty"`
	OtherRights         *PersonalRightsInfoResponse `json:"other_rights,omitempty" xml:"other_rights,omitempty"`
	Privileges          []*DataBoxPrivileges        `json:"privileges,omitempty" xml:"privileges,omitempty" type:"Repeated"`
	SpuId               *string                     `json:"spu_id,omitempty" xml:"spu_id,omitempty"`
	Title               *string                     `json:"title,omitempty" xml:"title,omitempty"`
}

func (s PersonalRightsInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s PersonalRightsInfoResponse) GoString() string {
	return s.String()
}

func (s *PersonalRightsInfoResponse) SetExpiresTime(v string) *PersonalRightsInfoResponse {
	s.ExpiresTime = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetHistoryLatestRights(v *PersonalRightsInfoResponse) *PersonalRightsInfoResponse {
	s.HistoryLatestRights = v
	return s
}

func (s *PersonalRightsInfoResponse) SetIcon(v string) *PersonalRightsInfoResponse {
	s.Icon = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetIsExpires(v bool) *PersonalRightsInfoResponse {
	s.IsExpires = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetName(v string) *PersonalRightsInfoResponse {
	s.Name = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetOtherRights(v *PersonalRightsInfoResponse) *PersonalRightsInfoResponse {
	s.OtherRights = v
	return s
}

func (s *PersonalRightsInfoResponse) SetPrivileges(v []*DataBoxPrivileges) *PersonalRightsInfoResponse {
	s.Privileges = v
	return s
}

func (s *PersonalRightsInfoResponse) SetSpuId(v string) *PersonalRightsInfoResponse {
	s.SpuId = &v
	return s
}

func (s *PersonalRightsInfoResponse) SetTitle(v string) *PersonalRightsInfoResponse {
	s.Title = &v
	return s
}

type PersonalSpaceInfo struct {
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	UsedSize  *int64 `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s PersonalSpaceInfo) String() string {
	return tea.Prettify(s)
}

func (s PersonalSpaceInfo) GoString() string {
	return s.String()
}

func (s *PersonalSpaceInfo) SetTotalSize(v int64) *PersonalSpaceInfo {
	s.TotalSize = &v
	return s
}

func (s *PersonalSpaceInfo) SetUsedSize(v int64) *PersonalSpaceInfo {
	s.UsedSize = &v
	return s
}

type RecycleBinConfig struct {
	AutoDeleteEnabled             *bool  `json:"auto_delete_enabled,omitempty" xml:"auto_delete_enabled,omitempty"`
	AutoDeleteKeepSecond          *int32 `json:"auto_delete_keep_second,omitempty" xml:"auto_delete_keep_second,omitempty"`
	DeleteTrashNormalFileDisabled *bool  `json:"delete_trash_normal_file_disabled,omitempty" xml:"delete_trash_normal_file_disabled,omitempty"`
}

func (s RecycleBinConfig) String() string {
	return tea.Prettify(s)
}

func (s RecycleBinConfig) GoString() string {
	return s.String()
}

func (s *RecycleBinConfig) SetAutoDeleteEnabled(v bool) *RecycleBinConfig {
	s.AutoDeleteEnabled = &v
	return s
}

func (s *RecycleBinConfig) SetAutoDeleteKeepSecond(v int32) *RecycleBinConfig {
	s.AutoDeleteKeepSecond = &v
	return s
}

func (s *RecycleBinConfig) SetDeleteTrashNormalFileDisabled(v bool) *RecycleBinConfig {
	s.DeleteTrashNormalFileDisabled = &v
	return s
}

type RefundNoticeParam struct {
	Aliuid            *int64             `json:"aliuid,omitempty" xml:"aliuid,omitempty"`
	AliyunProduceCode *string            `json:"aliyunProduceCode,omitempty" xml:"aliyunProduceCode,omitempty"`
	CommodityCode     *string            `json:"commodityCode,omitempty" xml:"commodityCode,omitempty"`
	InstanceId        *string            `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
	NewExpireTime     interface{}        `json:"newExpireTime,omitempty" xml:"newExpireTime,omitempty"`
	OrderIds          []*int64           `json:"orderIds,omitempty" xml:"orderIds,omitempty" type:"Repeated"`
	RefundParamMap    map[string]*string `json:"refundParamMap,omitempty" xml:"refundParamMap,omitempty"`
	RefundType        *string            `json:"refundType,omitempty" xml:"refundType,omitempty"`
}

func (s RefundNoticeParam) String() string {
	return tea.Prettify(s)
}

func (s RefundNoticeParam) GoString() string {
	return s.String()
}

func (s *RefundNoticeParam) SetAliuid(v int64) *RefundNoticeParam {
	s.Aliuid = &v
	return s
}

func (s *RefundNoticeParam) SetAliyunProduceCode(v string) *RefundNoticeParam {
	s.AliyunProduceCode = &v
	return s
}

func (s *RefundNoticeParam) SetCommodityCode(v string) *RefundNoticeParam {
	s.CommodityCode = &v
	return s
}

func (s *RefundNoticeParam) SetInstanceId(v string) *RefundNoticeParam {
	s.InstanceId = &v
	return s
}

func (s *RefundNoticeParam) SetNewExpireTime(v interface{}) *RefundNoticeParam {
	s.NewExpireTime = v
	return s
}

func (s *RefundNoticeParam) SetOrderIds(v []*int64) *RefundNoticeParam {
	s.OrderIds = v
	return s
}

func (s *RefundNoticeParam) SetRefundParamMap(v map[string]*string) *RefundNoticeParam {
	s.RefundParamMap = v
	return s
}

func (s *RefundNoticeParam) SetRefundType(v string) *RefundNoticeParam {
	s.RefundType = &v
	return s
}

type Revision struct {
	ContentHash         *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName     *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	Crc64Hash           *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	CreatedAt           *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	CreatorId           *string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	CreatorName         *string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	DomainId            *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DownloadUrl         *string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	DriveId             *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileExtension       *string `json:"file_extension,omitempty" xml:"file_extension,omitempty"`
	FileId              *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	IsLatestVersion     *bool   `json:"is_latest_version,omitempty" xml:"is_latest_version,omitempty"`
	KeepForever         *bool   `json:"keep_forever,omitempty" xml:"keep_forever,omitempty"`
	RevisionDescription *string `json:"revision_description,omitempty" xml:"revision_description,omitempty"`
	RevisionId          *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	RevisionName        *string `json:"revision_name,omitempty" xml:"revision_name,omitempty"`
	RevisionVersion     *int64  `json:"revision_version,omitempty" xml:"revision_version,omitempty"`
	Size                *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Thumbnail           *string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	UpdatedAt           *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	Url                 *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s Revision) String() string {
	return tea.Prettify(s)
}

func (s Revision) GoString() string {
	return s.String()
}

func (s *Revision) SetContentHash(v string) *Revision {
	s.ContentHash = &v
	return s
}

func (s *Revision) SetContentHashName(v string) *Revision {
	s.ContentHashName = &v
	return s
}

func (s *Revision) SetCrc64Hash(v string) *Revision {
	s.Crc64Hash = &v
	return s
}

func (s *Revision) SetCreatedAt(v string) *Revision {
	s.CreatedAt = &v
	return s
}

func (s *Revision) SetCreatorId(v string) *Revision {
	s.CreatorId = &v
	return s
}

func (s *Revision) SetCreatorName(v string) *Revision {
	s.CreatorName = &v
	return s
}

func (s *Revision) SetDomainId(v string) *Revision {
	s.DomainId = &v
	return s
}

func (s *Revision) SetDownloadUrl(v string) *Revision {
	s.DownloadUrl = &v
	return s
}

func (s *Revision) SetDriveId(v string) *Revision {
	s.DriveId = &v
	return s
}

func (s *Revision) SetFileExtension(v string) *Revision {
	s.FileExtension = &v
	return s
}

func (s *Revision) SetFileId(v string) *Revision {
	s.FileId = &v
	return s
}

func (s *Revision) SetIsLatestVersion(v bool) *Revision {
	s.IsLatestVersion = &v
	return s
}

func (s *Revision) SetKeepForever(v bool) *Revision {
	s.KeepForever = &v
	return s
}

func (s *Revision) SetRevisionDescription(v string) *Revision {
	s.RevisionDescription = &v
	return s
}

func (s *Revision) SetRevisionId(v string) *Revision {
	s.RevisionId = &v
	return s
}

func (s *Revision) SetRevisionName(v string) *Revision {
	s.RevisionName = &v
	return s
}

func (s *Revision) SetRevisionVersion(v int64) *Revision {
	s.RevisionVersion = &v
	return s
}

func (s *Revision) SetSize(v int64) *Revision {
	s.Size = &v
	return s
}

func (s *Revision) SetThumbnail(v string) *Revision {
	s.Thumbnail = &v
	return s
}

func (s *Revision) SetUpdatedAt(v string) *Revision {
	s.UpdatedAt = &v
	return s
}

func (s *Revision) SetUrl(v string) *Revision {
	s.Url = &v
	return s
}

type Role struct {
	CreatedAt          *int64        `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator            *string       `json:"creator,omitempty" xml:"creator,omitempty"`
	Description        *string       `json:"description,omitempty" xml:"description,omitempty"`
	ManageResourceType *string       `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	Name               *string       `json:"name,omitempty" xml:"name,omitempty"`
	Permissions        []*Permission `json:"permissions,omitempty" xml:"permissions,omitempty" type:"Repeated"`
	RoleId             *string       `json:"role_id,omitempty" xml:"role_id,omitempty"`
	Status             *string       `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt          *int64        `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s Role) String() string {
	return tea.Prettify(s)
}

func (s Role) GoString() string {
	return s.String()
}

func (s *Role) SetCreatedAt(v int64) *Role {
	s.CreatedAt = &v
	return s
}

func (s *Role) SetCreator(v string) *Role {
	s.Creator = &v
	return s
}

func (s *Role) SetDescription(v string) *Role {
	s.Description = &v
	return s
}

func (s *Role) SetManageResourceType(v string) *Role {
	s.ManageResourceType = &v
	return s
}

func (s *Role) SetName(v string) *Role {
	s.Name = &v
	return s
}

func (s *Role) SetPermissions(v []*Permission) *Role {
	s.Permissions = v
	return s
}

func (s *Role) SetRoleId(v string) *Role {
	s.RoleId = &v
	return s
}

func (s *Role) SetStatus(v string) *Role {
	s.Status = &v
	return s
}

func (s *Role) SetUpdatedAt(v int64) *Role {
	s.UpdatedAt = &v
	return s
}

type SearchFromThirdPartyItem struct {
	AuthenticationType *string                `json:"authentication_type,omitempty" xml:"authentication_type,omitempty"`
	Extra              *string                `json:"extra,omitempty" xml:"extra,omitempty"`
	Identity           *string                `json:"identity,omitempty" xml:"identity,omitempty"`
	Others             map[string]interface{} `json:"others,omitempty" xml:"others,omitempty"`
}

func (s SearchFromThirdPartyItem) String() string {
	return tea.Prettify(s)
}

func (s SearchFromThirdPartyItem) GoString() string {
	return s.String()
}

func (s *SearchFromThirdPartyItem) SetAuthenticationType(v string) *SearchFromThirdPartyItem {
	s.AuthenticationType = &v
	return s
}

func (s *SearchFromThirdPartyItem) SetExtra(v string) *SearchFromThirdPartyItem {
	s.Extra = &v
	return s
}

func (s *SearchFromThirdPartyItem) SetIdentity(v string) *SearchFromThirdPartyItem {
	s.Identity = &v
	return s
}

func (s *SearchFromThirdPartyItem) SetOthers(v map[string]interface{}) *SearchFromThirdPartyItem {
	s.Others = v
	return s
}

type ShareLink struct {
	AccessCount       *int64    `json:"access_count,omitempty" xml:"access_count,omitempty"`
	CreatedAt         *string   `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator           *string   `json:"creator,omitempty" xml:"creator,omitempty"`
	Description       *string   `json:"description,omitempty" xml:"description,omitempty"`
	DisableDownload   *bool     `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	DisablePreview    *bool     `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	DisableSave       *bool     `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	DownloadCount     *int64    `json:"download_count,omitempty" xml:"download_count,omitempty"`
	DownloadLimit     *int64    `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	DriveId           *string   `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Expiration        *string   `json:"expiration,omitempty" xml:"expiration,omitempty"`
	Expired           *bool     `json:"expired,omitempty" xml:"expired,omitempty"`
	FileIdList        []*string `json:"file_id_list,omitempty" xml:"file_id_list,omitempty" type:"Repeated"`
	OfficeEditable    *bool     `json:"office_editable,omitempty" xml:"office_editable,omitempty"`
	PreviewCount      *int64    `json:"preview_count,omitempty" xml:"preview_count,omitempty"`
	PreviewLimit      *int64    `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	ReportCount       *int64    `json:"report_count,omitempty" xml:"report_count,omitempty"`
	SaveCount         *int64    `json:"save_count,omitempty" xml:"save_count,omitempty"`
	SaveDownloadLimit *int64    `json:"save_download_limit,omitempty" xml:"save_download_limit,omitempty"`
	SaveLimit         *int64    `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	ShareAllFiles     *bool     `json:"share_all_files,omitempty" xml:"share_all_files,omitempty"`
	ShareId           *string   `json:"share_id,omitempty" xml:"share_id,omitempty"`
	ShareName         *string   `json:"share_name,omitempty" xml:"share_name,omitempty"`
	SharePwd          *string   `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
	Status            *string   `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt         *string   `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	VideoPreviewCount *int64    `json:"video_preview_count,omitempty" xml:"video_preview_count,omitempty"`
}

func (s ShareLink) String() string {
	return tea.Prettify(s)
}

func (s ShareLink) GoString() string {
	return s.String()
}

func (s *ShareLink) SetAccessCount(v int64) *ShareLink {
	s.AccessCount = &v
	return s
}

func (s *ShareLink) SetCreatedAt(v string) *ShareLink {
	s.CreatedAt = &v
	return s
}

func (s *ShareLink) SetCreator(v string) *ShareLink {
	s.Creator = &v
	return s
}

func (s *ShareLink) SetDescription(v string) *ShareLink {
	s.Description = &v
	return s
}

func (s *ShareLink) SetDisableDownload(v bool) *ShareLink {
	s.DisableDownload = &v
	return s
}

func (s *ShareLink) SetDisablePreview(v bool) *ShareLink {
	s.DisablePreview = &v
	return s
}

func (s *ShareLink) SetDisableSave(v bool) *ShareLink {
	s.DisableSave = &v
	return s
}

func (s *ShareLink) SetDownloadCount(v int64) *ShareLink {
	s.DownloadCount = &v
	return s
}

func (s *ShareLink) SetDownloadLimit(v int64) *ShareLink {
	s.DownloadLimit = &v
	return s
}

func (s *ShareLink) SetDriveId(v string) *ShareLink {
	s.DriveId = &v
	return s
}

func (s *ShareLink) SetExpiration(v string) *ShareLink {
	s.Expiration = &v
	return s
}

func (s *ShareLink) SetExpired(v bool) *ShareLink {
	s.Expired = &v
	return s
}

func (s *ShareLink) SetFileIdList(v []*string) *ShareLink {
	s.FileIdList = v
	return s
}

func (s *ShareLink) SetOfficeEditable(v bool) *ShareLink {
	s.OfficeEditable = &v
	return s
}

func (s *ShareLink) SetPreviewCount(v int64) *ShareLink {
	s.PreviewCount = &v
	return s
}

func (s *ShareLink) SetPreviewLimit(v int64) *ShareLink {
	s.PreviewLimit = &v
	return s
}

func (s *ShareLink) SetReportCount(v int64) *ShareLink {
	s.ReportCount = &v
	return s
}

func (s *ShareLink) SetSaveCount(v int64) *ShareLink {
	s.SaveCount = &v
	return s
}

func (s *ShareLink) SetSaveDownloadLimit(v int64) *ShareLink {
	s.SaveDownloadLimit = &v
	return s
}

func (s *ShareLink) SetSaveLimit(v int64) *ShareLink {
	s.SaveLimit = &v
	return s
}

func (s *ShareLink) SetShareAllFiles(v bool) *ShareLink {
	s.ShareAllFiles = &v
	return s
}

func (s *ShareLink) SetShareId(v string) *ShareLink {
	s.ShareId = &v
	return s
}

func (s *ShareLink) SetShareName(v string) *ShareLink {
	s.ShareName = &v
	return s
}

func (s *ShareLink) SetSharePwd(v string) *ShareLink {
	s.SharePwd = &v
	return s
}

func (s *ShareLink) SetStatus(v string) *ShareLink {
	s.Status = &v
	return s
}

func (s *ShareLink) SetUpdatedAt(v string) *ShareLink {
	s.UpdatedAt = &v
	return s
}

func (s *ShareLink) SetVideoPreviewCount(v int64) *ShareLink {
	s.VideoPreviewCount = &v
	return s
}

type ShareLinkConfig struct {
	EnableShareLinkOfficeEdit *bool `json:"enable_share_link_office_edit,omitempty" xml:"enable_share_link_office_edit,omitempty"`
}

func (s ShareLinkConfig) String() string {
	return tea.Prettify(s)
}

func (s ShareLinkConfig) GoString() string {
	return s.String()
}

func (s *ShareLinkConfig) SetEnableShareLinkOfficeEdit(v bool) *ShareLinkConfig {
	s.EnableShareLinkOfficeEdit = &v
	return s
}

type ShareLinkDetail struct {
	EnableOfficeEditable *bool `json:"enable_office_editable,omitempty" xml:"enable_office_editable,omitempty"`
}

func (s ShareLinkDetail) String() string {
	return tea.Prettify(s)
}

func (s ShareLinkDetail) GoString() string {
	return s.String()
}

func (s *ShareLinkDetail) SetEnableOfficeEditable(v bool) *ShareLinkDetail {
	s.EnableOfficeEditable = &v
	return s
}

type SimpleQuery struct {
	Field      []byte         `json:"field,omitempty" xml:"field,omitempty"`
	Operation  []byte         `json:"operation,omitempty" xml:"operation,omitempty"`
	SubQueries []*SimpleQuery `json:"sub_queries,omitempty" xml:"sub_queries,omitempty" type:"Repeated"`
	Value      []byte         `json:"value,omitempty" xml:"value,omitempty"`
}

func (s SimpleQuery) String() string {
	return tea.Prettify(s)
}

func (s SimpleQuery) GoString() string {
	return s.String()
}

func (s *SimpleQuery) SetField(v []byte) *SimpleQuery {
	s.Field = v
	return s
}

func (s *SimpleQuery) SetOperation(v []byte) *SimpleQuery {
	s.Operation = v
	return s
}

func (s *SimpleQuery) SetSubQueries(v []*SimpleQuery) *SimpleQuery {
	s.SubQueries = v
	return s
}

func (s *SimpleQuery) SetValue(v []byte) *SimpleQuery {
	s.Value = v
	return s
}

type SimpleStreamInfo struct {
	ContentHash     *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	Crc64Hash       *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	DownloadUrl     *string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	Size            *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Thumbnail       *string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	Url             *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s SimpleStreamInfo) String() string {
	return tea.Prettify(s)
}

func (s SimpleStreamInfo) GoString() string {
	return s.String()
}

func (s *SimpleStreamInfo) SetContentHash(v string) *SimpleStreamInfo {
	s.ContentHash = &v
	return s
}

func (s *SimpleStreamInfo) SetContentHashName(v string) *SimpleStreamInfo {
	s.ContentHashName = &v
	return s
}

func (s *SimpleStreamInfo) SetCrc64Hash(v string) *SimpleStreamInfo {
	s.Crc64Hash = &v
	return s
}

func (s *SimpleStreamInfo) SetDownloadUrl(v string) *SimpleStreamInfo {
	s.DownloadUrl = &v
	return s
}

func (s *SimpleStreamInfo) SetSize(v int64) *SimpleStreamInfo {
	s.Size = &v
	return s
}

func (s *SimpleStreamInfo) SetThumbnail(v string) *SimpleStreamInfo {
	s.Thumbnail = &v
	return s
}

func (s *SimpleStreamInfo) SetUrl(v string) *SimpleStreamInfo {
	s.Url = &v
	return s
}

type Story struct {
	CoverFileId           *string                `json:"cover_file_id,omitempty" xml:"cover_file_id,omitempty"`
	CoverFileThumbnailUrl *string                `json:"cover_file_thumbnail_url,omitempty" xml:"cover_file_thumbnail_url,omitempty"`
	CreatedAt             *string                `json:"created_at,omitempty" xml:"created_at,omitempty"`
	CustomLabels          map[string]interface{} `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	FaceGroupIds          []*string              `json:"face_group_ids,omitempty" xml:"face_group_ids,omitempty" type:"Repeated"`
	StoryEndTime          *string                `json:"story_end_time,omitempty" xml:"story_end_time,omitempty"`
	StoryFileList         []*File                `json:"story_file_list,omitempty" xml:"story_file_list,omitempty" type:"Repeated"`
	StoryId               *string                `json:"story_id,omitempty" xml:"story_id,omitempty"`
	StoryName             *string                `json:"story_name,omitempty" xml:"story_name,omitempty"`
	StoryStartTime        *string                `json:"story_start_time,omitempty" xml:"story_start_time,omitempty"`
	StorySubType          *string                `json:"story_sub_type,omitempty" xml:"story_sub_type,omitempty"`
	StoryType             *string                `json:"story_type,omitempty" xml:"story_type,omitempty"`
	UpdatedAt             *string                `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s Story) String() string {
	return tea.Prettify(s)
}

func (s Story) GoString() string {
	return s.String()
}

func (s *Story) SetCoverFileId(v string) *Story {
	s.CoverFileId = &v
	return s
}

func (s *Story) SetCoverFileThumbnailUrl(v string) *Story {
	s.CoverFileThumbnailUrl = &v
	return s
}

func (s *Story) SetCreatedAt(v string) *Story {
	s.CreatedAt = &v
	return s
}

func (s *Story) SetCustomLabels(v map[string]interface{}) *Story {
	s.CustomLabels = v
	return s
}

func (s *Story) SetFaceGroupIds(v []*string) *Story {
	s.FaceGroupIds = v
	return s
}

func (s *Story) SetStoryEndTime(v string) *Story {
	s.StoryEndTime = &v
	return s
}

func (s *Story) SetStoryFileList(v []*File) *Story {
	s.StoryFileList = v
	return s
}

func (s *Story) SetStoryId(v string) *Story {
	s.StoryId = &v
	return s
}

func (s *Story) SetStoryName(v string) *Story {
	s.StoryName = &v
	return s
}

func (s *Story) SetStoryStartTime(v string) *Story {
	s.StoryStartTime = &v
	return s
}

func (s *Story) SetStorySubType(v string) *Story {
	s.StorySubType = &v
	return s
}

func (s *Story) SetStoryType(v string) *Story {
	s.StoryType = &v
	return s
}

func (s *Story) SetUpdatedAt(v string) *Story {
	s.UpdatedAt = &v
	return s
}

type SystemTag struct {
	// example:
	//
	// 0.877
	CentricScore *float32 `json:"centric_score,omitempty" xml:"centric_score,omitempty"`
	// example:
	//
	// 0.98
	Confidence *float32 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	// example:
	//
	// 篮球
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// example:
	//
	// 运动
	ParentName *string `json:"parent_name,omitempty" xml:"parent_name,omitempty"`
	// example:
	//
	// 3
	TagLevel *int32 `json:"tag_level,omitempty" xml:"tag_level,omitempty"`
}

func (s SystemTag) String() string {
	return tea.Prettify(s)
}

func (s SystemTag) GoString() string {
	return s.String()
}

func (s *SystemTag) SetCentricScore(v float32) *SystemTag {
	s.CentricScore = &v
	return s
}

func (s *SystemTag) SetConfidence(v float32) *SystemTag {
	s.Confidence = &v
	return s
}

func (s *SystemTag) SetName(v string) *SystemTag {
	s.Name = &v
	return s
}

func (s *SystemTag) SetParentName(v string) *SystemTag {
	s.ParentName = &v
	return s
}

func (s *SystemTag) SetTagLevel(v int32) *SystemTag {
	s.TagLevel = &v
	return s
}

type TimeRange struct {
	End   *string `json:"end,omitempty" xml:"end,omitempty"`
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s TimeRange) String() string {
	return tea.Prettify(s)
}

func (s TimeRange) GoString() string {
	return s.String()
}

func (s *TimeRange) SetEnd(v string) *TimeRange {
	s.End = &v
	return s
}

func (s *TimeRange) SetStart(v string) *TimeRange {
	s.Start = &v
	return s
}

type Token struct {
	AccessToken        *string            `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Avatar             *string            `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DefaultDriveId     *string            `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	DefaultSboxDriveId *string            `json:"default_sbox_drive_id,omitempty" xml:"default_sbox_drive_id,omitempty"`
	DeviceId           *string            `json:"device_id,omitempty" xml:"device_id,omitempty"`
	DeviceName         *string            `json:"device_name,omitempty" xml:"device_name,omitempty"`
	DomainId           *string            `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	ExistLink          []*LinkInfo        `json:"exist_link,omitempty" xml:"exist_link,omitempty" type:"Repeated"`
	ExpireTime         *string            `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	ExpiresIn          *int64             `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	IsFirstLogin       *bool              `json:"is_first_login,omitempty" xml:"is_first_login,omitempty"`
	NeedLink           *bool              `json:"need_link,omitempty" xml:"need_link,omitempty"`
	NeedRpVerify       *bool              `json:"need_rp_verify,omitempty" xml:"need_rp_verify,omitempty"`
	NickName           *string            `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	PinSetup           *bool              `json:"pin_setup,omitempty" xml:"pin_setup,omitempty"`
	RefreshToken       *string            `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Role               *string            `json:"role,omitempty" xml:"role,omitempty"`
	State              *string            `json:"state,omitempty" xml:"state,omitempty"`
	Status             *string            `json:"status,omitempty" xml:"status,omitempty"`
	TokenType          *string            `json:"token_type,omitempty" xml:"token_type,omitempty"`
	UserData           map[string]*string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	UserId             *string            `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName           *string            `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s Token) String() string {
	return tea.Prettify(s)
}

func (s Token) GoString() string {
	return s.String()
}

func (s *Token) SetAccessToken(v string) *Token {
	s.AccessToken = &v
	return s
}

func (s *Token) SetAvatar(v string) *Token {
	s.Avatar = &v
	return s
}

func (s *Token) SetDefaultDriveId(v string) *Token {
	s.DefaultDriveId = &v
	return s
}

func (s *Token) SetDefaultSboxDriveId(v string) *Token {
	s.DefaultSboxDriveId = &v
	return s
}

func (s *Token) SetDeviceId(v string) *Token {
	s.DeviceId = &v
	return s
}

func (s *Token) SetDeviceName(v string) *Token {
	s.DeviceName = &v
	return s
}

func (s *Token) SetDomainId(v string) *Token {
	s.DomainId = &v
	return s
}

func (s *Token) SetExistLink(v []*LinkInfo) *Token {
	s.ExistLink = v
	return s
}

func (s *Token) SetExpireTime(v string) *Token {
	s.ExpireTime = &v
	return s
}

func (s *Token) SetExpiresIn(v int64) *Token {
	s.ExpiresIn = &v
	return s
}

func (s *Token) SetIsFirstLogin(v bool) *Token {
	s.IsFirstLogin = &v
	return s
}

func (s *Token) SetNeedLink(v bool) *Token {
	s.NeedLink = &v
	return s
}

func (s *Token) SetNeedRpVerify(v bool) *Token {
	s.NeedRpVerify = &v
	return s
}

func (s *Token) SetNickName(v string) *Token {
	s.NickName = &v
	return s
}

func (s *Token) SetPinSetup(v bool) *Token {
	s.PinSetup = &v
	return s
}

func (s *Token) SetRefreshToken(v string) *Token {
	s.RefreshToken = &v
	return s
}

func (s *Token) SetRole(v string) *Token {
	s.Role = &v
	return s
}

func (s *Token) SetState(v string) *Token {
	s.State = &v
	return s
}

func (s *Token) SetStatus(v string) *Token {
	s.Status = &v
	return s
}

func (s *Token) SetTokenType(v string) *Token {
	s.TokenType = &v
	return s
}

func (s *Token) SetUserData(v map[string]*string) *Token {
	s.UserData = v
	return s
}

func (s *Token) SetUserId(v string) *Token {
	s.UserId = &v
	return s
}

func (s *Token) SetUserName(v string) *Token {
	s.UserName = &v
	return s
}

type UncompressConfigResponse struct {
	Enabled *bool   `json:"enabled,omitempty" xml:"enabled,omitempty"`
	Version *string `json:"version,omitempty" xml:"version,omitempty"`
}

func (s UncompressConfigResponse) String() string {
	return tea.Prettify(s)
}

func (s UncompressConfigResponse) GoString() string {
	return s.String()
}

func (s *UncompressConfigResponse) SetEnabled(v bool) *UncompressConfigResponse {
	s.Enabled = &v
	return s
}

func (s *UncompressConfigResponse) SetVersion(v string) *UncompressConfigResponse {
	s.Version = &v
	return s
}

type UncompressedFileInfo struct {
	DriveId   *string                 `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId    *string                 `json:"file_id,omitempty" xml:"file_id,omitempty"`
	IsFolder  *bool                   `json:"is_folder,omitempty" xml:"is_folder,omitempty"`
	Items     []*UncompressedFileInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	Name      *string                 `json:"name,omitempty" xml:"name,omitempty"`
	Size      *int64                  `json:"size,omitempty" xml:"size,omitempty"`
	UpdatedAt *int64                  `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
}

func (s UncompressedFileInfo) String() string {
	return tea.Prettify(s)
}

func (s UncompressedFileInfo) GoString() string {
	return s.String()
}

func (s *UncompressedFileInfo) SetDriveId(v string) *UncompressedFileInfo {
	s.DriveId = &v
	return s
}

func (s *UncompressedFileInfo) SetFileId(v string) *UncompressedFileInfo {
	s.FileId = &v
	return s
}

func (s *UncompressedFileInfo) SetIsFolder(v bool) *UncompressedFileInfo {
	s.IsFolder = &v
	return s
}

func (s *UncompressedFileInfo) SetItems(v []*UncompressedFileInfo) *UncompressedFileInfo {
	s.Items = v
	return s
}

func (s *UncompressedFileInfo) SetName(v string) *UncompressedFileInfo {
	s.Name = &v
	return s
}

func (s *UncompressedFileInfo) SetSize(v int64) *UncompressedFileInfo {
	s.Size = &v
	return s
}

func (s *UncompressedFileInfo) SetUpdatedAt(v int64) *UncompressedFileInfo {
	s.UpdatedAt = &v
	return s
}

type UploadFormInfo struct {
	BucketName       *string            `json:"bucket_name,omitempty" xml:"bucket_name,omitempty"`
	Endpoint         *string            `json:"endpoint,omitempty" xml:"endpoint,omitempty"`
	FormData         map[string]*string `json:"form_data,omitempty" xml:"form_data,omitempty"`
	ObjectKey        *string            `json:"object_key,omitempty" xml:"object_key,omitempty"`
	OssAccessKeyId   *string            `json:"oss_access_key_id,omitempty" xml:"oss_access_key_id,omitempty"`
	OssEndPoint      *string            `json:"oss_end_point,omitempty" xml:"oss_end_point,omitempty"`
	OssSecurityToken *string            `json:"oss_security_token,omitempty" xml:"oss_security_token,omitempty"`
	Policy           *string            `json:"policy,omitempty" xml:"policy,omitempty"`
	Signature        *string            `json:"signature,omitempty" xml:"signature,omitempty"`
}

func (s UploadFormInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadFormInfo) GoString() string {
	return s.String()
}

func (s *UploadFormInfo) SetBucketName(v string) *UploadFormInfo {
	s.BucketName = &v
	return s
}

func (s *UploadFormInfo) SetEndpoint(v string) *UploadFormInfo {
	s.Endpoint = &v
	return s
}

func (s *UploadFormInfo) SetFormData(v map[string]*string) *UploadFormInfo {
	s.FormData = v
	return s
}

func (s *UploadFormInfo) SetObjectKey(v string) *UploadFormInfo {
	s.ObjectKey = &v
	return s
}

func (s *UploadFormInfo) SetOssAccessKeyId(v string) *UploadFormInfo {
	s.OssAccessKeyId = &v
	return s
}

func (s *UploadFormInfo) SetOssEndPoint(v string) *UploadFormInfo {
	s.OssEndPoint = &v
	return s
}

func (s *UploadFormInfo) SetOssSecurityToken(v string) *UploadFormInfo {
	s.OssSecurityToken = &v
	return s
}

func (s *UploadFormInfo) SetPolicy(v string) *UploadFormInfo {
	s.Policy = &v
	return s
}

func (s *UploadFormInfo) SetSignature(v string) *UploadFormInfo {
	s.Signature = &v
	return s
}

type UploadPartInfo struct {
	// example:
	//
	// 0CC175B9C0F1B6A831C399E269772661
	Etag              *string                          `json:"etag,omitempty" xml:"etag,omitempty"`
	InternalUploadUrl *string                          `json:"internal_upload_url,omitempty" xml:"internal_upload_url,omitempty"`
	ParallelSha1Ctx   *UploadPartInfoParallelSha1Ctx   `json:"parallel_sha1_ctx,omitempty" xml:"parallel_sha1_ctx,omitempty" type:"Struct"`
	ParallelSha256Ctx *UploadPartInfoParallelSha256Ctx `json:"parallel_sha256_ctx,omitempty" xml:"parallel_sha256_ctx,omitempty" type:"Struct"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	PartNumber *int32 `json:"part_number,omitempty" xml:"part_number,omitempty"`
	// example:
	//
	// 1024
	PartSize *int64 `json:"part_size,omitempty" xml:"part_size,omitempty"`
	// This parameter is required.
	UploadUrl *string `json:"upload_url,omitempty" xml:"upload_url,omitempty"`
}

func (s UploadPartInfo) String() string {
	return tea.Prettify(s)
}

func (s UploadPartInfo) GoString() string {
	return s.String()
}

func (s *UploadPartInfo) SetEtag(v string) *UploadPartInfo {
	s.Etag = &v
	return s
}

func (s *UploadPartInfo) SetInternalUploadUrl(v string) *UploadPartInfo {
	s.InternalUploadUrl = &v
	return s
}

func (s *UploadPartInfo) SetParallelSha1Ctx(v *UploadPartInfoParallelSha1Ctx) *UploadPartInfo {
	s.ParallelSha1Ctx = v
	return s
}

func (s *UploadPartInfo) SetParallelSha256Ctx(v *UploadPartInfoParallelSha256Ctx) *UploadPartInfo {
	s.ParallelSha256Ctx = v
	return s
}

func (s *UploadPartInfo) SetPartNumber(v int32) *UploadPartInfo {
	s.PartNumber = &v
	return s
}

func (s *UploadPartInfo) SetPartSize(v int64) *UploadPartInfo {
	s.PartSize = &v
	return s
}

func (s *UploadPartInfo) SetUploadUrl(v string) *UploadPartInfo {
	s.UploadUrl = &v
	return s
}

type UploadPartInfoParallelSha1Ctx struct {
	H          []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	PartOffset *int64   `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s UploadPartInfoParallelSha1Ctx) String() string {
	return tea.Prettify(s)
}

func (s UploadPartInfoParallelSha1Ctx) GoString() string {
	return s.String()
}

func (s *UploadPartInfoParallelSha1Ctx) SetH(v []*int64) *UploadPartInfoParallelSha1Ctx {
	s.H = v
	return s
}

func (s *UploadPartInfoParallelSha1Ctx) SetPartOffset(v int64) *UploadPartInfoParallelSha1Ctx {
	s.PartOffset = &v
	return s
}

type UploadPartInfoParallelSha256Ctx struct {
	H          []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	PartOffset *int64   `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s UploadPartInfoParallelSha256Ctx) String() string {
	return tea.Prettify(s)
}

func (s UploadPartInfoParallelSha256Ctx) GoString() string {
	return s.String()
}

func (s *UploadPartInfoParallelSha256Ctx) SetH(v []*int64) *UploadPartInfoParallelSha256Ctx {
	s.H = v
	return s
}

func (s *UploadPartInfoParallelSha256Ctx) SetPartOffset(v int64) *UploadPartInfoParallelSha256Ctx {
	s.PartOffset = &v
	return s
}

type User struct {
	Avatar         *string            `json:"avatar,omitempty" xml:"avatar,omitempty"`
	CreatedAt      *int64             `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator        *string            `json:"creator,omitempty" xml:"creator,omitempty"`
	DefaultDriveId *string            `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	Description    *string            `json:"description,omitempty" xml:"description,omitempty"`
	DomainId       *string            `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	Email          *string            `json:"email,omitempty" xml:"email,omitempty"`
	NickName       *string            `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	Phone          *string            `json:"phone,omitempty" xml:"phone,omitempty"`
	Role           *string            `json:"role,omitempty" xml:"role,omitempty"`
	Status         *string            `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt      *int64             `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UserData       map[string]*string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	UserId         *string            `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName       *string            `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s User) String() string {
	return tea.Prettify(s)
}

func (s User) GoString() string {
	return s.String()
}

func (s *User) SetAvatar(v string) *User {
	s.Avatar = &v
	return s
}

func (s *User) SetCreatedAt(v int64) *User {
	s.CreatedAt = &v
	return s
}

func (s *User) SetCreator(v string) *User {
	s.Creator = &v
	return s
}

func (s *User) SetDefaultDriveId(v string) *User {
	s.DefaultDriveId = &v
	return s
}

func (s *User) SetDescription(v string) *User {
	s.Description = &v
	return s
}

func (s *User) SetDomainId(v string) *User {
	s.DomainId = &v
	return s
}

func (s *User) SetEmail(v string) *User {
	s.Email = &v
	return s
}

func (s *User) SetNickName(v string) *User {
	s.NickName = &v
	return s
}

func (s *User) SetPhone(v string) *User {
	s.Phone = &v
	return s
}

func (s *User) SetRole(v string) *User {
	s.Role = &v
	return s
}

func (s *User) SetStatus(v string) *User {
	s.Status = &v
	return s
}

func (s *User) SetUpdatedAt(v int64) *User {
	s.UpdatedAt = &v
	return s
}

func (s *User) SetUserData(v map[string]*string) *User {
	s.UserData = v
	return s
}

func (s *User) SetUserId(v string) *User {
	s.UserId = &v
	return s
}

func (s *User) SetUserName(v string) *User {
	s.UserName = &v
	return s
}

type UserExtraItem struct {
	Account []*AccountLinkInfo `json:"account,omitempty" xml:"account,omitempty" type:"Repeated"`
	// example:
	//
	// http://a.b.c/ccp.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// example:
	//
	// 1567407718386
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// example:
	//
	// system
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// if can be null:
	// true
	DefaultDrive *BaseDriveResponse `json:"default_drive,omitempty" xml:"default_drive,omitempty"`
	// example:
	//
	// 123
	DefaultDriveId           *string `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	DefaultLocation          *string `json:"default_location,omitempty" xml:"default_location,omitempty"`
	DenyChangePasswordBySelf *bool   `json:"deny_change_password_by_self,omitempty" xml:"deny_change_password_by_self,omitempty"`
	// example:
	//
	// ccp team user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// example:
	//
	// hz999
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// example:
	//
	// 123@ccp.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// example:
	//
	// 0
	ExpiredAt                   *int64 `json:"expired_at,omitempty" xml:"expired_at,omitempty"`
	IsSync                      *bool  `json:"is_sync,omitempty" xml:"is_sync,omitempty"`
	LastLoginTime               *int64 `json:"last_login_time,omitempty" xml:"last_login_time,omitempty"`
	NeedChangePasswordNextLogin *bool  `json:"need_change_password_next_login,omitempty" xml:"need_change_password_next_login,omitempty"`
	// example:
	//
	// abc
	NickName    *string                  `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	ParentGroup []*BaseDriveResponse     `json:"parent_group,omitempty" xml:"parent_group,omitempty" type:"Repeated"`
	PathStatus  *string                  `json:"path_status,omitempty" xml:"path_status,omitempty"`
	Permission  map[string]*IDPermission `json:"permission,omitempty" xml:"permission,omitempty"`
	// example:
	//
	// 13700000000
	Phone       *string `json:"phone,omitempty" xml:"phone,omitempty"`
	PhoneRegion *string `json:"phone_region,omitempty" xml:"phone_region,omitempty"`
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 1567407718386
	UpdatedAt *string                `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UserData  map[string]interface{} `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// example:
	//
	// ccpuserid
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// example:
	//
	// name
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s UserExtraItem) String() string {
	return tea.Prettify(s)
}

func (s UserExtraItem) GoString() string {
	return s.String()
}

func (s *UserExtraItem) SetAccount(v []*AccountLinkInfo) *UserExtraItem {
	s.Account = v
	return s
}

func (s *UserExtraItem) SetAvatar(v string) *UserExtraItem {
	s.Avatar = &v
	return s
}

func (s *UserExtraItem) SetCreatedAt(v string) *UserExtraItem {
	s.CreatedAt = &v
	return s
}

func (s *UserExtraItem) SetCreator(v string) *UserExtraItem {
	s.Creator = &v
	return s
}

func (s *UserExtraItem) SetDefaultDrive(v *BaseDriveResponse) *UserExtraItem {
	s.DefaultDrive = v
	return s
}

func (s *UserExtraItem) SetDefaultDriveId(v string) *UserExtraItem {
	s.DefaultDriveId = &v
	return s
}

func (s *UserExtraItem) SetDefaultLocation(v string) *UserExtraItem {
	s.DefaultLocation = &v
	return s
}

func (s *UserExtraItem) SetDenyChangePasswordBySelf(v bool) *UserExtraItem {
	s.DenyChangePasswordBySelf = &v
	return s
}

func (s *UserExtraItem) SetDescription(v string) *UserExtraItem {
	s.Description = &v
	return s
}

func (s *UserExtraItem) SetDomainId(v string) *UserExtraItem {
	s.DomainId = &v
	return s
}

func (s *UserExtraItem) SetEmail(v string) *UserExtraItem {
	s.Email = &v
	return s
}

func (s *UserExtraItem) SetExpiredAt(v int64) *UserExtraItem {
	s.ExpiredAt = &v
	return s
}

func (s *UserExtraItem) SetIsSync(v bool) *UserExtraItem {
	s.IsSync = &v
	return s
}

func (s *UserExtraItem) SetLastLoginTime(v int64) *UserExtraItem {
	s.LastLoginTime = &v
	return s
}

func (s *UserExtraItem) SetNeedChangePasswordNextLogin(v bool) *UserExtraItem {
	s.NeedChangePasswordNextLogin = &v
	return s
}

func (s *UserExtraItem) SetNickName(v string) *UserExtraItem {
	s.NickName = &v
	return s
}

func (s *UserExtraItem) SetParentGroup(v []*BaseDriveResponse) *UserExtraItem {
	s.ParentGroup = v
	return s
}

func (s *UserExtraItem) SetPathStatus(v string) *UserExtraItem {
	s.PathStatus = &v
	return s
}

func (s *UserExtraItem) SetPermission(v map[string]*IDPermission) *UserExtraItem {
	s.Permission = v
	return s
}

func (s *UserExtraItem) SetPhone(v string) *UserExtraItem {
	s.Phone = &v
	return s
}

func (s *UserExtraItem) SetPhoneRegion(v string) *UserExtraItem {
	s.PhoneRegion = &v
	return s
}

func (s *UserExtraItem) SetRole(v string) *UserExtraItem {
	s.Role = &v
	return s
}

func (s *UserExtraItem) SetStatus(v string) *UserExtraItem {
	s.Status = &v
	return s
}

func (s *UserExtraItem) SetUpdatedAt(v string) *UserExtraItem {
	s.UpdatedAt = &v
	return s
}

func (s *UserExtraItem) SetUserData(v map[string]interface{}) *UserExtraItem {
	s.UserData = v
	return s
}

func (s *UserExtraItem) SetUserId(v string) *UserExtraItem {
	s.UserId = &v
	return s
}

func (s *UserExtraItem) SetUserName(v string) *UserExtraItem {
	s.UserName = &v
	return s
}

type UserLogDetail struct {
	Email     *string                `json:"email,omitempty" xml:"email,omitempty"`
	ExpiredAt *int64                 `json:"expired_at,omitempty" xml:"expired_at,omitempty"`
	Name      *string                `json:"name,omitempty" xml:"name,omitempty"`
	Phone     *string                `json:"phone,omitempty" xml:"phone,omitempty"`
	RoleId    *string                `json:"role_id,omitempty" xml:"role_id,omitempty"`
	UpdateTo  *UserLogDetailUpdateTo `json:"update_to,omitempty" xml:"update_to,omitempty" type:"Struct"`
}

func (s UserLogDetail) String() string {
	return tea.Prettify(s)
}

func (s UserLogDetail) GoString() string {
	return s.String()
}

func (s *UserLogDetail) SetEmail(v string) *UserLogDetail {
	s.Email = &v
	return s
}

func (s *UserLogDetail) SetExpiredAt(v int64) *UserLogDetail {
	s.ExpiredAt = &v
	return s
}

func (s *UserLogDetail) SetName(v string) *UserLogDetail {
	s.Name = &v
	return s
}

func (s *UserLogDetail) SetPhone(v string) *UserLogDetail {
	s.Phone = &v
	return s
}

func (s *UserLogDetail) SetRoleId(v string) *UserLogDetail {
	s.RoleId = &v
	return s
}

func (s *UserLogDetail) SetUpdateTo(v *UserLogDetailUpdateTo) *UserLogDetail {
	s.UpdateTo = v
	return s
}

type UserLogDetailUpdateTo struct {
	Email     *string `json:"email,omitempty" xml:"email,omitempty"`
	ExpiredAt *int64  `json:"expired_at,omitempty" xml:"expired_at,omitempty"`
	Name      *string `json:"name,omitempty" xml:"name,omitempty"`
	Phone     *string `json:"phone,omitempty" xml:"phone,omitempty"`
	RoleId    *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s UserLogDetailUpdateTo) String() string {
	return tea.Prettify(s)
}

func (s UserLogDetailUpdateTo) GoString() string {
	return s.String()
}

func (s *UserLogDetailUpdateTo) SetEmail(v string) *UserLogDetailUpdateTo {
	s.Email = &v
	return s
}

func (s *UserLogDetailUpdateTo) SetExpiredAt(v int64) *UserLogDetailUpdateTo {
	s.ExpiredAt = &v
	return s
}

func (s *UserLogDetailUpdateTo) SetName(v string) *UserLogDetailUpdateTo {
	s.Name = &v
	return s
}

func (s *UserLogDetailUpdateTo) SetPhone(v string) *UserLogDetailUpdateTo {
	s.Phone = &v
	return s
}

func (s *UserLogDetailUpdateTo) SetRoleId(v string) *UserLogDetailUpdateTo {
	s.RoleId = &v
	return s
}

type UserTag struct {
	// This parameter is required.
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// This parameter is required.
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s UserTag) String() string {
	return tea.Prettify(s)
}

func (s UserTag) GoString() string {
	return s.String()
}

func (s *UserTag) SetKey(v string) *UserTag {
	s.Key = &v
	return s
}

func (s *UserTag) SetValue(v string) *UserTag {
	s.Value = &v
	return s
}

type VideoMediaAudioStream struct {
	// example:
	//
	// 129280
	BitRate *string `json:"bit_rate,omitempty" xml:"bit_rate,omitempty"`
	// example:
	//
	// aac
	CodeName *string `json:"code_name,omitempty" xml:"code_name,omitempty"`
	// example:
	//
	// 7704.573000
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
}

func (s VideoMediaAudioStream) String() string {
	return tea.Prettify(s)
}

func (s VideoMediaAudioStream) GoString() string {
	return s.String()
}

func (s *VideoMediaAudioStream) SetBitRate(v string) *VideoMediaAudioStream {
	s.BitRate = &v
	return s
}

func (s *VideoMediaAudioStream) SetCodeName(v string) *VideoMediaAudioStream {
	s.CodeName = &v
	return s
}

func (s *VideoMediaAudioStream) SetDuration(v string) *VideoMediaAudioStream {
	s.Duration = &v
	return s
}

type VideoMediaMetadata struct {
	// example:
	//
	// 1080
	Height                *int64                   `json:"height,omitempty" xml:"height,omitempty"`
	VideoMediaAudioStream []*VideoMediaAudioStream `json:"video_media_audio_stream,omitempty" xml:"video_media_audio_stream,omitempty" type:"Repeated"`
	VideoMediaVideoStream []*VideoMediaVideoStream `json:"video_media_video_stream,omitempty" xml:"video_media_video_stream,omitempty" type:"Repeated"`
	// example:
	//
	// 1920
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s VideoMediaMetadata) String() string {
	return tea.Prettify(s)
}

func (s VideoMediaMetadata) GoString() string {
	return s.String()
}

func (s *VideoMediaMetadata) SetHeight(v int64) *VideoMediaMetadata {
	s.Height = &v
	return s
}

func (s *VideoMediaMetadata) SetVideoMediaAudioStream(v []*VideoMediaAudioStream) *VideoMediaMetadata {
	s.VideoMediaAudioStream = v
	return s
}

func (s *VideoMediaMetadata) SetVideoMediaVideoStream(v []*VideoMediaVideoStream) *VideoMediaMetadata {
	s.VideoMediaVideoStream = v
	return s
}

func (s *VideoMediaMetadata) SetWidth(v int64) *VideoMediaMetadata {
	s.Width = &v
	return s
}

type VideoMediaVideoStream struct {
	// example:
	//
	// 108420
	Bitrate *string `json:"bitrate,omitempty" xml:"bitrate,omitempty"`
	// example:
	//
	// h264
	CodeName *string `json:"code_name,omitempty" xml:"code_name,omitempty"`
	// example:
	//
	// 22.88
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 90
	FrameCount *string `json:"frame_count,omitempty" xml:"frame_count,omitempty"`
}

func (s VideoMediaVideoStream) String() string {
	return tea.Prettify(s)
}

func (s VideoMediaVideoStream) GoString() string {
	return s.String()
}

func (s *VideoMediaVideoStream) SetBitrate(v string) *VideoMediaVideoStream {
	s.Bitrate = &v
	return s
}

func (s *VideoMediaVideoStream) SetCodeName(v string) *VideoMediaVideoStream {
	s.CodeName = &v
	return s
}

func (s *VideoMediaVideoStream) SetDuration(v string) *VideoMediaVideoStream {
	s.Duration = &v
	return s
}

func (s *VideoMediaVideoStream) SetFrameCount(v string) *VideoMediaVideoStream {
	s.FrameCount = &v
	return s
}

type VideoPreviewPlayInfo struct {
	// example:
	//
	// live_transcoding
	Category                            *string                                            `json:"category,omitempty" xml:"category,omitempty"`
	LiveTranscodingSubtitleTaskList     []*VideoPreviewSubtitleInfo                        `json:"live_transcoding_subtitle_task_list,omitempty" xml:"live_transcoding_subtitle_task_list,omitempty" type:"Repeated"`
	LiveTranscodingTaskList             []*VideoPreviewPlayInfoLiveTranscodingTaskList     `json:"live_transcoding_task_list,omitempty" xml:"live_transcoding_task_list,omitempty" type:"Repeated"`
	MasterUrl                           *string                                            `json:"master_url,omitempty" xml:"master_url,omitempty"`
	Meta                                *VideoPreviewPlayInfoMeta                          `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	OfflineVideoTranscodingList         []*VideoPreviewPlayInfoOfflineVideoTranscodingList `json:"offline_video_transcoding_list,omitempty" xml:"offline_video_transcoding_list,omitempty" type:"Repeated"`
	OfflineVideoTranscodingSubtitleList []*VideoPreviewSubtitleInfo                        `json:"offline_video_transcoding_subtitle_list,omitempty" xml:"offline_video_transcoding_subtitle_list,omitempty" type:"Repeated"`
	QuickVideoList                      []*VideoPreviewPlayInfoQuickVideoList              `json:"quick_video_list,omitempty" xml:"quick_video_list,omitempty" type:"Repeated"`
	QuickVideoSubtitleList              []*VideoPreviewSubtitleInfo                        `json:"quick_video_subtitle_list,omitempty" xml:"quick_video_subtitle_list,omitempty" type:"Repeated"`
}

func (s VideoPreviewPlayInfo) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayInfo) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfo) SetCategory(v string) *VideoPreviewPlayInfo {
	s.Category = &v
	return s
}

func (s *VideoPreviewPlayInfo) SetLiveTranscodingSubtitleTaskList(v []*VideoPreviewSubtitleInfo) *VideoPreviewPlayInfo {
	s.LiveTranscodingSubtitleTaskList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetLiveTranscodingTaskList(v []*VideoPreviewPlayInfoLiveTranscodingTaskList) *VideoPreviewPlayInfo {
	s.LiveTranscodingTaskList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetMasterUrl(v string) *VideoPreviewPlayInfo {
	s.MasterUrl = &v
	return s
}

func (s *VideoPreviewPlayInfo) SetMeta(v *VideoPreviewPlayInfoMeta) *VideoPreviewPlayInfo {
	s.Meta = v
	return s
}

func (s *VideoPreviewPlayInfo) SetOfflineVideoTranscodingList(v []*VideoPreviewPlayInfoOfflineVideoTranscodingList) *VideoPreviewPlayInfo {
	s.OfflineVideoTranscodingList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetOfflineVideoTranscodingSubtitleList(v []*VideoPreviewSubtitleInfo) *VideoPreviewPlayInfo {
	s.OfflineVideoTranscodingSubtitleList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetQuickVideoList(v []*VideoPreviewPlayInfoQuickVideoList) *VideoPreviewPlayInfo {
	s.QuickVideoList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetQuickVideoSubtitleList(v []*VideoPreviewSubtitleInfo) *VideoPreviewPlayInfo {
	s.QuickVideoSubtitleList = v
	return s
}

type VideoPreviewPlayInfoLiveTranscodingTaskList struct {
	KeepOriginalResolution *bool   `json:"keep_original_resolution,omitempty" xml:"keep_original_resolution,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
	TemplateId             *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	Url                    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s VideoPreviewPlayInfoLiveTranscodingTaskList) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayInfoLiveTranscodingTaskList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) SetKeepOriginalResolution(v bool) *VideoPreviewPlayInfoLiveTranscodingTaskList {
	s.KeepOriginalResolution = &v
	return s
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) SetStatus(v string) *VideoPreviewPlayInfoLiveTranscodingTaskList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) SetTemplateId(v string) *VideoPreviewPlayInfoLiveTranscodingTaskList {
	s.TemplateId = &v
	return s
}

func (s *VideoPreviewPlayInfoLiveTranscodingTaskList) SetUrl(v string) *VideoPreviewPlayInfoLiveTranscodingTaskList {
	s.Url = &v
	return s
}

type VideoPreviewPlayInfoMeta struct {
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	Height   *int64   `json:"height,omitempty" xml:"height,omitempty"`
	Width    *int64   `json:"width,omitempty" xml:"width,omitempty"`
}

func (s VideoPreviewPlayInfoMeta) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayInfoMeta) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfoMeta) SetDuration(v float64) *VideoPreviewPlayInfoMeta {
	s.Duration = &v
	return s
}

func (s *VideoPreviewPlayInfoMeta) SetHeight(v int64) *VideoPreviewPlayInfoMeta {
	s.Height = &v
	return s
}

func (s *VideoPreviewPlayInfoMeta) SetWidth(v int64) *VideoPreviewPlayInfoMeta {
	s.Width = &v
	return s
}

type VideoPreviewPlayInfoOfflineVideoTranscodingList struct {
	KeepOriginalResolution *bool   `json:"keep_original_resolution,omitempty" xml:"keep_original_resolution,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
	TemplateId             *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	Url                    *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s VideoPreviewPlayInfoOfflineVideoTranscodingList) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayInfoOfflineVideoTranscodingList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) SetKeepOriginalResolution(v bool) *VideoPreviewPlayInfoOfflineVideoTranscodingList {
	s.KeepOriginalResolution = &v
	return s
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) SetStatus(v string) *VideoPreviewPlayInfoOfflineVideoTranscodingList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) SetTemplateId(v string) *VideoPreviewPlayInfoOfflineVideoTranscodingList {
	s.TemplateId = &v
	return s
}

func (s *VideoPreviewPlayInfoOfflineVideoTranscodingList) SetUrl(v string) *VideoPreviewPlayInfoOfflineVideoTranscodingList {
	s.Url = &v
	return s
}

type VideoPreviewPlayInfoQuickVideoList struct {
	Status     *string `json:"status,omitempty" xml:"status,omitempty"`
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	Url        *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s VideoPreviewPlayInfoQuickVideoList) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayInfoQuickVideoList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayInfoQuickVideoList) SetStatus(v string) *VideoPreviewPlayInfoQuickVideoList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayInfoQuickVideoList) SetTemplateId(v string) *VideoPreviewPlayInfoQuickVideoList {
	s.TemplateId = &v
	return s
}

func (s *VideoPreviewPlayInfoQuickVideoList) SetUrl(v string) *VideoPreviewPlayInfoQuickVideoList {
	s.Url = &v
	return s
}

type VideoPreviewPlayMeta struct {
	// example:
	//
	// live_transcoding
	Category                    *string                                            `json:"category,omitempty" xml:"category,omitempty"`
	LiveTranscodingTaskList     []*VideoPreviewPlayMetaLiveTranscodingTaskList     `json:"live_transcoding_task_list,omitempty" xml:"live_transcoding_task_list,omitempty" type:"Repeated"`
	Meta                        *VideoPreviewPlayMetaMeta                          `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	OfflineVideoTranscodingList []*VideoPreviewPlayMetaOfflineVideoTranscodingList `json:"offline_video_transcoding_list,omitempty" xml:"offline_video_transcoding_list,omitempty" type:"Repeated"`
	QuickVideoList              []*VideoPreviewPlayMetaQuickVideoList              `json:"quick_video_list,omitempty" xml:"quick_video_list,omitempty" type:"Repeated"`
}

func (s VideoPreviewPlayMeta) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayMeta) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMeta) SetCategory(v string) *VideoPreviewPlayMeta {
	s.Category = &v
	return s
}

func (s *VideoPreviewPlayMeta) SetLiveTranscodingTaskList(v []*VideoPreviewPlayMetaLiveTranscodingTaskList) *VideoPreviewPlayMeta {
	s.LiveTranscodingTaskList = v
	return s
}

func (s *VideoPreviewPlayMeta) SetMeta(v *VideoPreviewPlayMetaMeta) *VideoPreviewPlayMeta {
	s.Meta = v
	return s
}

func (s *VideoPreviewPlayMeta) SetOfflineVideoTranscodingList(v []*VideoPreviewPlayMetaOfflineVideoTranscodingList) *VideoPreviewPlayMeta {
	s.OfflineVideoTranscodingList = v
	return s
}

func (s *VideoPreviewPlayMeta) SetQuickVideoList(v []*VideoPreviewPlayMetaQuickVideoList) *VideoPreviewPlayMeta {
	s.QuickVideoList = v
	return s
}

type VideoPreviewPlayMetaLiveTranscodingTaskList struct {
	// example:
	//
	// true
	KeepOriginalResolution *bool `json:"keep_original_resolution,omitempty" xml:"keep_original_resolution,omitempty"`
	// example:
	//
	// finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 264_720p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

func (s VideoPreviewPlayMetaLiveTranscodingTaskList) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayMetaLiveTranscodingTaskList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) SetKeepOriginalResolution(v bool) *VideoPreviewPlayMetaLiveTranscodingTaskList {
	s.KeepOriginalResolution = &v
	return s
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) SetStatus(v string) *VideoPreviewPlayMetaLiveTranscodingTaskList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayMetaLiveTranscodingTaskList) SetTemplateId(v string) *VideoPreviewPlayMetaLiveTranscodingTaskList {
	s.TemplateId = &v
	return s
}

type VideoPreviewPlayMetaMeta struct {
	// example:
	//
	// 10
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	// example:
	//
	// 720
	Height *int64 `json:"height,omitempty" xml:"height,omitempty"`
	// example:
	//
	// 1280
	Width *int64 `json:"width,omitempty" xml:"width,omitempty"`
}

func (s VideoPreviewPlayMetaMeta) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayMetaMeta) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMetaMeta) SetDuration(v float64) *VideoPreviewPlayMetaMeta {
	s.Duration = &v
	return s
}

func (s *VideoPreviewPlayMetaMeta) SetHeight(v int64) *VideoPreviewPlayMetaMeta {
	s.Height = &v
	return s
}

func (s *VideoPreviewPlayMetaMeta) SetWidth(v int64) *VideoPreviewPlayMetaMeta {
	s.Width = &v
	return s
}

type VideoPreviewPlayMetaOfflineVideoTranscodingList struct {
	// example:
	//
	// true
	KeepOriginalResolution *string `json:"keep_original_resolution,omitempty" xml:"keep_original_resolution,omitempty"`
	// example:
	//
	// finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 264_720p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

func (s VideoPreviewPlayMetaOfflineVideoTranscodingList) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayMetaOfflineVideoTranscodingList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) SetKeepOriginalResolution(v string) *VideoPreviewPlayMetaOfflineVideoTranscodingList {
	s.KeepOriginalResolution = &v
	return s
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) SetStatus(v string) *VideoPreviewPlayMetaOfflineVideoTranscodingList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayMetaOfflineVideoTranscodingList) SetTemplateId(v string) *VideoPreviewPlayMetaOfflineVideoTranscodingList {
	s.TemplateId = &v
	return s
}

type VideoPreviewPlayMetaQuickVideoList struct {
	// example:
	//
	// finished
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// example:
	//
	// 264_720p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
}

func (s VideoPreviewPlayMetaQuickVideoList) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewPlayMetaQuickVideoList) GoString() string {
	return s.String()
}

func (s *VideoPreviewPlayMetaQuickVideoList) SetStatus(v string) *VideoPreviewPlayMetaQuickVideoList {
	s.Status = &v
	return s
}

func (s *VideoPreviewPlayMetaQuickVideoList) SetTemplateId(v string) *VideoPreviewPlayMetaQuickVideoList {
	s.TemplateId = &v
	return s
}

type VideoPreviewSubtitleInfo struct {
	Language *string `json:"language,omitempty" xml:"language,omitempty"`
	Status   *string `json:"status,omitempty" xml:"status,omitempty"`
	Url      *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s VideoPreviewSubtitleInfo) String() string {
	return tea.Prettify(s)
}

func (s VideoPreviewSubtitleInfo) GoString() string {
	return s.String()
}

func (s *VideoPreviewSubtitleInfo) SetLanguage(v string) *VideoPreviewSubtitleInfo {
	s.Language = &v
	return s
}

func (s *VideoPreviewSubtitleInfo) SetStatus(v string) *VideoPreviewSubtitleInfo {
	s.Status = &v
	return s
}

func (s *VideoPreviewSubtitleInfo) SetUrl(v string) *VideoPreviewSubtitleInfo {
	s.Url = &v
	return s
}

type View struct {
	Category     *string                `json:"category,omitempty" xml:"category,omitempty"`
	CreatedAt    *string                `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description  *string                `json:"description,omitempty" xml:"description,omitempty"`
	ExFieldsInfo map[string]interface{} `json:"ex_fields_info,omitempty" xml:"ex_fields_info,omitempty"`
	FileCount    *int64                 `json:"file_count,omitempty" xml:"file_count,omitempty"`
	Name         *string                `json:"name,omitempty" xml:"name,omitempty"`
	Owner        *string                `json:"owner,omitempty" xml:"owner,omitempty"`
	UpdatedAt    *string                `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	ViewId       *string                `json:"view_id,omitempty" xml:"view_id,omitempty"`
}

func (s View) String() string {
	return tea.Prettify(s)
}

func (s View) GoString() string {
	return s.String()
}

func (s *View) SetCategory(v string) *View {
	s.Category = &v
	return s
}

func (s *View) SetCreatedAt(v string) *View {
	s.CreatedAt = &v
	return s
}

func (s *View) SetDescription(v string) *View {
	s.Description = &v
	return s
}

func (s *View) SetExFieldsInfo(v map[string]interface{}) *View {
	s.ExFieldsInfo = v
	return s
}

func (s *View) SetFileCount(v int64) *View {
	s.FileCount = &v
	return s
}

func (s *View) SetName(v string) *View {
	s.Name = &v
	return s
}

func (s *View) SetOwner(v string) *View {
	s.Owner = &v
	return s
}

func (s *View) SetUpdatedAt(v string) *View {
	s.UpdatedAt = &v
	return s
}

func (s *View) SetViewId(v string) *View {
	s.ViewId = &v
	return s
}

type ViewFile struct {
	Category          *string                    `json:"category,omitempty" xml:"category,omitempty"`
	ContentHash       *string                    `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName   *string                    `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	ContentType       *string                    `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Crc64Hash         *string                    `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	CreatedAt         *string                    `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description       *string                    `json:"description,omitempty" xml:"description,omitempty"`
	DomainId          *string                    `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DownloadUrl       *string                    `json:"download_url,omitempty" xml:"download_url,omitempty"`
	DriveId           *string                    `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Fields            map[string]interface{}     `json:"fields,omitempty" xml:"fields,omitempty"`
	FileExtension     *string                    `json:"file_extension,omitempty" xml:"file_extension,omitempty"`
	FileId            *string                    `json:"file_id,omitempty" xml:"file_id,omitempty"`
	FileRevisionId    *string                    `json:"file_revision_id,omitempty" xml:"file_revision_id,omitempty"`
	Hidden            *bool                      `json:"hidden,omitempty" xml:"hidden,omitempty"`
	InvestigationInfo *ViewFileInvestigationInfo `json:"investigation_info,omitempty" xml:"investigation_info,omitempty" type:"Struct"`
	JoinedAt          *int64                     `json:"joined_at,omitempty" xml:"joined_at,omitempty"`
	Labels            []*string                  `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	LocalCreatedAt    *string                    `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	LocalModifiedAt   *string                    `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	Name              *string                    `json:"name,omitempty" xml:"name,omitempty"`
	ParentFileId      *string                    `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	RevisionId        *string                    `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	Size              *int64                     `json:"size,omitempty" xml:"size,omitempty"`
	Starred           *bool                      `json:"starred,omitempty" xml:"starred,omitempty"`
	Status            *string                    `json:"status,omitempty" xml:"status,omitempty"`
	Thumbnail         *string                    `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	ThumbnailUrls     map[string]*string         `json:"thumbnail_urls,omitempty" xml:"thumbnail_urls,omitempty"`
	TrashedAt         *string                    `json:"trashed_at,omitempty" xml:"trashed_at,omitempty"`
	Type              *string                    `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt         *string                    `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UploadId          *string                    `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	ViewId            *string                    `json:"view_id,omitempty" xml:"view_id,omitempty"`
}

func (s ViewFile) String() string {
	return tea.Prettify(s)
}

func (s ViewFile) GoString() string {
	return s.String()
}

func (s *ViewFile) SetCategory(v string) *ViewFile {
	s.Category = &v
	return s
}

func (s *ViewFile) SetContentHash(v string) *ViewFile {
	s.ContentHash = &v
	return s
}

func (s *ViewFile) SetContentHashName(v string) *ViewFile {
	s.ContentHashName = &v
	return s
}

func (s *ViewFile) SetContentType(v string) *ViewFile {
	s.ContentType = &v
	return s
}

func (s *ViewFile) SetCrc64Hash(v string) *ViewFile {
	s.Crc64Hash = &v
	return s
}

func (s *ViewFile) SetCreatedAt(v string) *ViewFile {
	s.CreatedAt = &v
	return s
}

func (s *ViewFile) SetDescription(v string) *ViewFile {
	s.Description = &v
	return s
}

func (s *ViewFile) SetDomainId(v string) *ViewFile {
	s.DomainId = &v
	return s
}

func (s *ViewFile) SetDownloadUrl(v string) *ViewFile {
	s.DownloadUrl = &v
	return s
}

func (s *ViewFile) SetDriveId(v string) *ViewFile {
	s.DriveId = &v
	return s
}

func (s *ViewFile) SetFields(v map[string]interface{}) *ViewFile {
	s.Fields = v
	return s
}

func (s *ViewFile) SetFileExtension(v string) *ViewFile {
	s.FileExtension = &v
	return s
}

func (s *ViewFile) SetFileId(v string) *ViewFile {
	s.FileId = &v
	return s
}

func (s *ViewFile) SetFileRevisionId(v string) *ViewFile {
	s.FileRevisionId = &v
	return s
}

func (s *ViewFile) SetHidden(v bool) *ViewFile {
	s.Hidden = &v
	return s
}

func (s *ViewFile) SetInvestigationInfo(v *ViewFileInvestigationInfo) *ViewFile {
	s.InvestigationInfo = v
	return s
}

func (s *ViewFile) SetJoinedAt(v int64) *ViewFile {
	s.JoinedAt = &v
	return s
}

func (s *ViewFile) SetLabels(v []*string) *ViewFile {
	s.Labels = v
	return s
}

func (s *ViewFile) SetLocalCreatedAt(v string) *ViewFile {
	s.LocalCreatedAt = &v
	return s
}

func (s *ViewFile) SetLocalModifiedAt(v string) *ViewFile {
	s.LocalModifiedAt = &v
	return s
}

func (s *ViewFile) SetName(v string) *ViewFile {
	s.Name = &v
	return s
}

func (s *ViewFile) SetParentFileId(v string) *ViewFile {
	s.ParentFileId = &v
	return s
}

func (s *ViewFile) SetRevisionId(v string) *ViewFile {
	s.RevisionId = &v
	return s
}

func (s *ViewFile) SetSize(v int64) *ViewFile {
	s.Size = &v
	return s
}

func (s *ViewFile) SetStarred(v bool) *ViewFile {
	s.Starred = &v
	return s
}

func (s *ViewFile) SetStatus(v string) *ViewFile {
	s.Status = &v
	return s
}

func (s *ViewFile) SetThumbnail(v string) *ViewFile {
	s.Thumbnail = &v
	return s
}

func (s *ViewFile) SetThumbnailUrls(v map[string]*string) *ViewFile {
	s.ThumbnailUrls = v
	return s
}

func (s *ViewFile) SetTrashedAt(v string) *ViewFile {
	s.TrashedAt = &v
	return s
}

func (s *ViewFile) SetType(v string) *ViewFile {
	s.Type = &v
	return s
}

func (s *ViewFile) SetUpdatedAt(v string) *ViewFile {
	s.UpdatedAt = &v
	return s
}

func (s *ViewFile) SetUploadId(v string) *ViewFile {
	s.UploadId = &v
	return s
}

func (s *ViewFile) SetViewId(v string) *ViewFile {
	s.ViewId = &v
	return s
}

type ViewFileInvestigationInfo struct {
	Status     *int64  `json:"status,omitempty" xml:"status,omitempty"`
	Suggestion *string `json:"suggestion,omitempty" xml:"suggestion,omitempty"`
}

func (s ViewFileInvestigationInfo) String() string {
	return tea.Prettify(s)
}

func (s ViewFileInvestigationInfo) GoString() string {
	return s.String()
}

func (s *ViewFileInvestigationInfo) SetStatus(v int64) *ViewFileInvestigationInfo {
	s.Status = &v
	return s
}

func (s *ViewFileInvestigationInfo) SetSuggestion(v string) *ViewFileInvestigationInfo {
	s.Suggestion = &v
	return s
}

type WatermarkEnableConfig struct {
	DisplayAccessUserName       *bool   `json:"display_access_user_name,omitempty" xml:"display_access_user_name,omitempty"`
	DisplayCustomText           *string `json:"display_custom_text,omitempty" xml:"display_custom_text,omitempty"`
	DisplayShareLinkCreatorName *bool   `json:"display_shareLink_creator_name,omitempty" xml:"display_shareLink_creator_name,omitempty"`
	EnableDocPreview            *bool   `json:"enable_doc_preview,omitempty" xml:"enable_doc_preview,omitempty"`
}

func (s WatermarkEnableConfig) String() string {
	return tea.Prettify(s)
}

func (s WatermarkEnableConfig) GoString() string {
	return s.String()
}

func (s *WatermarkEnableConfig) SetDisplayAccessUserName(v bool) *WatermarkEnableConfig {
	s.DisplayAccessUserName = &v
	return s
}

func (s *WatermarkEnableConfig) SetDisplayCustomText(v string) *WatermarkEnableConfig {
	s.DisplayCustomText = &v
	return s
}

func (s *WatermarkEnableConfig) SetDisplayShareLinkCreatorName(v bool) *WatermarkEnableConfig {
	s.DisplayShareLinkCreatorName = &v
	return s
}

func (s *WatermarkEnableConfig) SetEnableDocPreview(v bool) *WatermarkEnableConfig {
	s.EnableDocPreview = &v
	return s
}

type WxTrustedDomainConfig struct {
	Content *string `json:"content,omitempty" xml:"content,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
	Show    *bool   `json:"show,omitempty" xml:"show,omitempty"`
}

func (s WxTrustedDomainConfig) String() string {
	return tea.Prettify(s)
}

func (s WxTrustedDomainConfig) GoString() string {
	return s.String()
}

func (s *WxTrustedDomainConfig) SetContent(v string) *WxTrustedDomainConfig {
	s.Content = &v
	return s
}

func (s *WxTrustedDomainConfig) SetName(v string) *WxTrustedDomainConfig {
	s.Name = &v
	return s
}

func (s *WxTrustedDomainConfig) SetShow(v bool) *WxTrustedDomainConfig {
	s.Show = &v
	return s
}

type AddGroupMemberRequest struct {
	// The ID of the destination group to which the member is added.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e5***2c2
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The member ID. If member_type is set to user, set this parameter to a user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e4***1b1
	MemberId *string `json:"member_id,omitempty" xml:"member_id,omitempty"`
	// The type of the member. Set the value to user. When you create a group, you can directly add the group to a parent group.
	//
	// 	- user
	//
	// Note: A group can be added to only one group. A user can be added to multiple groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	MemberType *string `json:"member_type,omitempty" xml:"member_type,omitempty"`
}

func (s AddGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *AddGroupMemberRequest) SetGroupId(v string) *AddGroupMemberRequest {
	s.GroupId = &v
	return s
}

func (s *AddGroupMemberRequest) SetMemberId(v string) *AddGroupMemberRequest {
	s.MemberId = &v
	return s
}

func (s *AddGroupMemberRequest) SetMemberType(v string) *AddGroupMemberRequest {
	s.MemberType = &v
	return s
}

type AddGroupMemberResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AddGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s AddGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *AddGroupMemberResponse) SetHeaders(v map[string]*string) *AddGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *AddGroupMemberResponse) SetStatusCode(v int32) *AddGroupMemberResponse {
	s.StatusCode = &v
	return s
}

type AddStoryFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string                      `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Files   []*AddStoryFilesRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s AddStoryFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesRequest) GoString() string {
	return s.String()
}

func (s *AddStoryFilesRequest) SetDriveId(v string) *AddStoryFilesRequest {
	s.DriveId = &v
	return s
}

func (s *AddStoryFilesRequest) SetFiles(v []*AddStoryFilesRequestFiles) *AddStoryFilesRequest {
	s.Files = v
	return s
}

func (s *AddStoryFilesRequest) SetStoryId(v string) *AddStoryFilesRequest {
	s.StoryId = &v
	return s
}

type AddStoryFilesRequestFiles struct {
	// This parameter is required.
	//
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88dd06e49d9c0a14411ebae606f70edd9a59
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s AddStoryFilesRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesRequestFiles) GoString() string {
	return s.String()
}

func (s *AddStoryFilesRequestFiles) SetFileId(v string) *AddStoryFilesRequestFiles {
	s.FileId = &v
	return s
}

func (s *AddStoryFilesRequestFiles) SetRevisionId(v string) *AddStoryFilesRequestFiles {
	s.RevisionId = &v
	return s
}

type AddStoryFilesResponseBody struct {
	// example:
	//
	// 1
	DriveId   *string         `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Files     []*AddStoryFile `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	RequestId *string         `json:"request_id,omitempty" xml:"request_id,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s AddStoryFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesResponseBody) GoString() string {
	return s.String()
}

func (s *AddStoryFilesResponseBody) SetDriveId(v string) *AddStoryFilesResponseBody {
	s.DriveId = &v
	return s
}

func (s *AddStoryFilesResponseBody) SetFiles(v []*AddStoryFile) *AddStoryFilesResponseBody {
	s.Files = v
	return s
}

func (s *AddStoryFilesResponseBody) SetRequestId(v string) *AddStoryFilesResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddStoryFilesResponseBody) SetStoryId(v string) *AddStoryFilesResponseBody {
	s.StoryId = &v
	return s
}

type AddStoryFilesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AddStoryFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s AddStoryFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s AddStoryFilesResponse) GoString() string {
	return s.String()
}

func (s *AddStoryFilesResponse) SetHeaders(v map[string]*string) *AddStoryFilesResponse {
	s.Headers = v
	return s
}

func (s *AddStoryFilesResponse) SetStatusCode(v int32) *AddStoryFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *AddStoryFilesResponse) SetBody(v *AddStoryFilesResponseBody) *AddStoryFilesResponse {
	s.Body = v
	return s
}

type AssignRoleRequest struct {
	// The unique identifier of a user. The group administrator role can only be assigned to a user.
	//
	// This parameter is required.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	// The ID of the resource that the role can manage. You can only set this parameter to the ID of a group.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105***b82
	ManageResourceId *string `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	// The type of the resource that the role can manage. Valid value: RT_Group.
	//
	// This parameter is required.
	//
	// example:
	//
	// RT_Group
	ManageResourceType *string `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	// The ID of the role that is assigned to a user. Valid value: SystemGroupAdmin.
	//
	// This parameter is required.
	//
	// example:
	//
	// SystemGroupAdmin
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s AssignRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s AssignRoleRequest) GoString() string {
	return s.String()
}

func (s *AssignRoleRequest) SetIdentity(v *Identity) *AssignRoleRequest {
	s.Identity = v
	return s
}

func (s *AssignRoleRequest) SetManageResourceId(v string) *AssignRoleRequest {
	s.ManageResourceId = &v
	return s
}

func (s *AssignRoleRequest) SetManageResourceType(v string) *AssignRoleRequest {
	s.ManageResourceType = &v
	return s
}

func (s *AssignRoleRequest) SetRoleId(v string) *AssignRoleRequest {
	s.RoleId = &v
	return s
}

type AssignRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AssignRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s AssignRoleResponse) GoString() string {
	return s.String()
}

func (s *AssignRoleResponse) SetHeaders(v map[string]*string) *AssignRoleResponse {
	s.Headers = v
	return s
}

func (s *AssignRoleResponse) SetStatusCode(v int32) *AssignRoleResponse {
	s.StatusCode = &v
	return s
}

type AuthorizeRequest struct {
	// The application ID returned when the application was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47eUHhrzgWBvlLWj
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// Specifies whether to hide the consent page.
	//
	// example:
	//
	// true
	HideConsent *bool `json:"hide_consent,omitempty" xml:"hide_consent,omitempty"`
	// The authentication method. Valid values:
	//
	// 	- default: all logon methods that are integrated on the default logon page provided by Drive and Photo Service.
	//
	// 	- ding: logs on by scanning a DingTalk QR code.
	//
	// 	- ding_sns: logs on by entering a DingTalk account and its password.
	//
	// 	- ram: logs on as an Alibaba Cloud Resource Access Management (RAM) user.
	//
	// 	- wechat: logs on by scanning a WeCom QR code.
	//
	// 	- wechat_app: logs on without authentication in WeCom.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	LoginType *string `json:"login_type,omitempty" xml:"login_type,omitempty"`
	// The callback URL specified when the application was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyunpds.com/sign/callback
	RedirectUri *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	// The format in which to return the response. Set the value to code.
	//
	// This parameter is required.
	//
	// example:
	//
	// code
	ResponseType *string `json:"response_type,omitempty" xml:"response_type,omitempty"`
	// The requested permissions. By default, all permissions are requested.
	Scope []*string `json:"scope,omitempty" xml:"scope,omitempty" type:"Repeated"`
	// The user-defined parameter to return in the callback URL after the requested permissions are granted.
	//
	// example:
	//
	// customdata
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s AuthorizeRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeRequest) SetClientId(v string) *AuthorizeRequest {
	s.ClientId = &v
	return s
}

func (s *AuthorizeRequest) SetHideConsent(v bool) *AuthorizeRequest {
	s.HideConsent = &v
	return s
}

func (s *AuthorizeRequest) SetLoginType(v string) *AuthorizeRequest {
	s.LoginType = &v
	return s
}

func (s *AuthorizeRequest) SetRedirectUri(v string) *AuthorizeRequest {
	s.RedirectUri = &v
	return s
}

func (s *AuthorizeRequest) SetResponseType(v string) *AuthorizeRequest {
	s.ResponseType = &v
	return s
}

func (s *AuthorizeRequest) SetScope(v []*string) *AuthorizeRequest {
	s.Scope = v
	return s
}

func (s *AuthorizeRequest) SetState(v string) *AuthorizeRequest {
	s.State = &v
	return s
}

type AuthorizeShrinkRequest struct {
	// The application ID returned when the application was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// 47eUHhrzgWBvlLWj
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// Specifies whether to hide the consent page.
	//
	// example:
	//
	// true
	HideConsent *bool `json:"hide_consent,omitempty" xml:"hide_consent,omitempty"`
	// The authentication method. Valid values:
	//
	// 	- default: all logon methods that are integrated on the default logon page provided by Drive and Photo Service.
	//
	// 	- ding: logs on by scanning a DingTalk QR code.
	//
	// 	- ding_sns: logs on by entering a DingTalk account and its password.
	//
	// 	- ram: logs on as an Alibaba Cloud Resource Access Management (RAM) user.
	//
	// 	- wechat: logs on by scanning a WeCom QR code.
	//
	// 	- wechat_app: logs on without authentication in WeCom.
	//
	// This parameter is required.
	//
	// example:
	//
	// default
	LoginType *string `json:"login_type,omitempty" xml:"login_type,omitempty"`
	// The callback URL specified when the application was created.
	//
	// This parameter is required.
	//
	// example:
	//
	// https://www.aliyunpds.com/sign/callback
	RedirectUri *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	// The format in which to return the response. Set the value to code.
	//
	// This parameter is required.
	//
	// example:
	//
	// code
	ResponseType *string `json:"response_type,omitempty" xml:"response_type,omitempty"`
	// The requested permissions. By default, all permissions are requested.
	ScopeShrink *string `json:"scope,omitempty" xml:"scope,omitempty"`
	// The user-defined parameter to return in the callback URL after the requested permissions are granted.
	//
	// example:
	//
	// customdata
	State *string `json:"state,omitempty" xml:"state,omitempty"`
}

func (s AuthorizeShrinkRequest) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeShrinkRequest) GoString() string {
	return s.String()
}

func (s *AuthorizeShrinkRequest) SetClientId(v string) *AuthorizeShrinkRequest {
	s.ClientId = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetHideConsent(v bool) *AuthorizeShrinkRequest {
	s.HideConsent = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetLoginType(v string) *AuthorizeShrinkRequest {
	s.LoginType = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetRedirectUri(v string) *AuthorizeShrinkRequest {
	s.RedirectUri = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetResponseType(v string) *AuthorizeShrinkRequest {
	s.ResponseType = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetScopeShrink(v string) *AuthorizeShrinkRequest {
	s.ScopeShrink = &v
	return s
}

func (s *AuthorizeShrinkRequest) SetState(v string) *AuthorizeShrinkRequest {
	s.State = &v
	return s
}

type AuthorizeResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s AuthorizeResponse) String() string {
	return tea.Prettify(s)
}

func (s AuthorizeResponse) GoString() string {
	return s.String()
}

func (s *AuthorizeResponse) SetHeaders(v map[string]*string) *AuthorizeResponse {
	s.Headers = v
	return s
}

func (s *AuthorizeResponse) SetStatusCode(v int32) *AuthorizeResponse {
	s.StatusCode = &v
	return s
}

type BatchRequest struct {
	// The child requests.
	//
	// The number of child requests. Valid value: 1 to 100.
	//
	// This parameter is required.
	Requests []*BatchRequestRequests `json:"requests,omitempty" xml:"requests,omitempty" type:"Repeated"`
	// The type of the resource that you want to manage. Valid values:
	//
	// 	- file: a file.
	//
	// 	- drive: an individual drive or a team drive.
	//
	// 	- user: a user.
	//
	// 	- group: a group.
	//
	// 	- membership: a group member.
	//
	// 	- share_link: a share.
	//
	// 	- async_task: an asynchronous task.
	//
	// This parameter is required.
	//
	// example:
	//
	// file
	Resource *string `json:"resource,omitempty" xml:"resource,omitempty"`
}

func (s BatchRequest) String() string {
	return tea.Prettify(s)
}

func (s BatchRequest) GoString() string {
	return s.String()
}

func (s *BatchRequest) SetRequests(v []*BatchRequestRequests) *BatchRequest {
	s.Requests = v
	return s
}

func (s *BatchRequest) SetResource(v string) *BatchRequest {
	s.Resource = &v
	return s
}

type BatchRequestRequests struct {
	// The request parameters of a child request. The parameter value must be a JSON string. For more information, see the topic of the corresponding child request.
	//
	// Before you specify the request body, you must specify a header by using Content-Type. Content-Type can only be set to application/json.
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// The header of a child request, which indicates the type of the data specified in the request body.
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	// The ID of the child request. The ID is used to associate a child request with a response. The ID of a child request must be unique.
	//
	// This parameter is required.
	//
	// example:
	//
	// 93433894994ad2e1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The method of a child request. Valid values:
	//
	// 	- POST
	//
	// 	- GET
	//
	// 	- PUT
	//
	// 	- DELETE
	//
	// 	- HEAD
	//
	// This parameter is required.
	//
	// example:
	//
	// POST
	Method *string `json:"method,omitempty" xml:"method,omitempty"`
	// The API path of a child request. Valid values:
	//
	// 	- /file/get: queries the information about a file.
	//
	// 	- /file/update: modifies the information about a file.
	//
	// 	- /file/search: searches for a file.
	//
	// 	- /file/copy: copies a file or folder.
	//
	// 	- /file/move: moves a file or folder.
	//
	// 	- /file/delete: deletes a file or folder.
	//
	// 	- /file/get_download_url: queries the download URL of a file.
	//
	// 	- /file/get_share_link_download_url: queries the download URL of a file in a share.
	//
	// 	- /recyclebin/trash: moves a file or folder to the recycle bin.
	//
	// 	- /recyclebin/restore: restores a file or folder.
	//
	// 	- /file/put_usertags: adds tags to a user.
	//
	// 	- /file/delete_usertags: removes tags from a user.
	//
	// 	- /drive/get: queries the information about a drive.
	//
	// 	- /user/get: queries the information about a user.
	//
	// 	- /group/get: queries the information about a group.
	//
	// 	- /share_link/create: creates a share.
	//
	// 	- /share_link/update: modifies a share.
	//
	// 	- /share_link/cancel: cancels a share.
	//
	// 	- /share_link/list: queries shares.
	//
	// 	- /share_link/get: queries the information about a share.
	//
	// 	- /share_link/get_share_token: queries an access token of a share.
	//
	// 	- /async_task/get: queries the information about an asynchronous task.
	//
	// This parameter is required.
	//
	// example:
	//
	// /file/get
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s BatchRequestRequests) String() string {
	return tea.Prettify(s)
}

func (s BatchRequestRequests) GoString() string {
	return s.String()
}

func (s *BatchRequestRequests) SetBody(v map[string]interface{}) *BatchRequestRequests {
	s.Body = v
	return s
}

func (s *BatchRequestRequests) SetHeaders(v map[string]*string) *BatchRequestRequests {
	s.Headers = v
	return s
}

func (s *BatchRequestRequests) SetId(v string) *BatchRequestRequests {
	s.Id = &v
	return s
}

func (s *BatchRequestRequests) SetMethod(v string) *BatchRequestRequests {
	s.Method = &v
	return s
}

func (s *BatchRequestRequests) SetUrl(v string) *BatchRequestRequests {
	s.Url = &v
	return s
}

type BatchResponseBody struct {
	// All responses of the child requests.
	Responses []*BatchResponseBodyResponses `json:"responses,omitempty" xml:"responses,omitempty" type:"Repeated"`
}

func (s BatchResponseBody) String() string {
	return tea.Prettify(s)
}

func (s BatchResponseBody) GoString() string {
	return s.String()
}

func (s *BatchResponseBody) SetResponses(v []*BatchResponseBodyResponses) *BatchResponseBody {
	s.Responses = v
	return s
}

type BatchResponseBodyResponses struct {
	// The response parameters of a child request. For more information, see the topic of the corresponding child request.
	Body map[string]interface{} `json:"body,omitempty" xml:"body,omitempty"`
	// The ID of the child request. The ID is used to associate a child request with a response.
	//
	// example:
	//
	// 93433894994ad2e1
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
	// The returned HTTP status code of a child request. For more information, see the topic of the corresponding child request.
	//
	// example:
	//
	// 200
	Status *int32 `json:"status,omitempty" xml:"status,omitempty"`
}

func (s BatchResponseBodyResponses) String() string {
	return tea.Prettify(s)
}

func (s BatchResponseBodyResponses) GoString() string {
	return s.String()
}

func (s *BatchResponseBodyResponses) SetBody(v map[string]interface{}) *BatchResponseBodyResponses {
	s.Body = v
	return s
}

func (s *BatchResponseBodyResponses) SetId(v string) *BatchResponseBodyResponses {
	s.Id = &v
	return s
}

func (s *BatchResponseBodyResponses) SetStatus(v int32) *BatchResponseBodyResponses {
	s.Status = &v
	return s
}

type BatchResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BatchResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s BatchResponse) String() string {
	return tea.Prettify(s)
}

func (s BatchResponse) GoString() string {
	return s.String()
}

func (s *BatchResponse) SetHeaders(v map[string]*string) *BatchResponse {
	s.Headers = v
	return s
}

func (s *BatchResponse) SetStatusCode(v int32) *BatchResponse {
	s.StatusCode = &v
	return s
}

func (s *BatchResponse) SetBody(v *BatchResponseBody) *BatchResponse {
	s.Body = v
	return s
}

type CancelAssignRoleRequest struct {
	// The unique identifier. You can cancel only the role assigned to a user.
	//
	// This parameter is required.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	// The ID of the resource that the role manages. Set the value to a group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105***b82
	ManageResourceId *string `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	// The type of the resource that the role manages. Set the value to RT_Group, which specifies group.
	//
	// This parameter is required.
	//
	// example:
	//
	// RT_Group
	ManageResourceType *string `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	// The ID of the role to be canceled. Set the value to SystemGroupAdmin, which is the ID of the group administrator role.
	//
	// This parameter is required.
	//
	// example:
	//
	// SystemGroupAdmin
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s CancelAssignRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelAssignRoleRequest) GoString() string {
	return s.String()
}

func (s *CancelAssignRoleRequest) SetIdentity(v *Identity) *CancelAssignRoleRequest {
	s.Identity = v
	return s
}

func (s *CancelAssignRoleRequest) SetManageResourceId(v string) *CancelAssignRoleRequest {
	s.ManageResourceId = &v
	return s
}

func (s *CancelAssignRoleRequest) SetManageResourceType(v string) *CancelAssignRoleRequest {
	s.ManageResourceType = &v
	return s
}

func (s *CancelAssignRoleRequest) SetRoleId(v string) *CancelAssignRoleRequest {
	s.RoleId = &v
	return s
}

type CancelAssignRoleResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CancelAssignRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelAssignRoleResponse) GoString() string {
	return s.String()
}

func (s *CancelAssignRoleResponse) SetHeaders(v map[string]*string) *CancelAssignRoleResponse {
	s.Headers = v
	return s
}

func (s *CancelAssignRoleResponse) SetStatusCode(v int32) *CancelAssignRoleResponse {
	s.StatusCode = &v
	return s
}

type CancelShareLinkRequest struct {
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s CancelShareLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CancelShareLinkRequest) GoString() string {
	return s.String()
}

func (s *CancelShareLinkRequest) SetShareId(v string) *CancelShareLinkRequest {
	s.ShareId = &v
	return s
}

type CancelShareLinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CancelShareLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s CancelShareLinkResponse) GoString() string {
	return s.String()
}

func (s *CancelShareLinkResponse) SetHeaders(v map[string]*string) *CancelShareLinkResponse {
	s.Headers = v
	return s
}

func (s *CancelShareLinkResponse) SetStatusCode(v int32) *CancelShareLinkResponse {
	s.StatusCode = &v
	return s
}

type ClearRecyclebinRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s ClearRecyclebinRequest) String() string {
	return tea.Prettify(s)
}

func (s ClearRecyclebinRequest) GoString() string {
	return s.String()
}

func (s *ClearRecyclebinRequest) SetDriveId(v string) *ClearRecyclebinRequest {
	s.DriveId = &v
	return s
}

type ClearRecyclebinResponseBody struct {
	// The ID of the asynchronous task.
	//
	// You can call the GetAsyncTask operation to query the information about the asynchronous task based on the task ID.
	//
	// example:
	//
	// 13ebd3a24dba4166b1527add676ef2866051b4d5dele16
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s ClearRecyclebinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ClearRecyclebinResponseBody) GoString() string {
	return s.String()
}

func (s *ClearRecyclebinResponseBody) SetAsyncTaskId(v string) *ClearRecyclebinResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *ClearRecyclebinResponseBody) SetDomainId(v string) *ClearRecyclebinResponseBody {
	s.DomainId = &v
	return s
}

func (s *ClearRecyclebinResponseBody) SetDriveId(v string) *ClearRecyclebinResponseBody {
	s.DriveId = &v
	return s
}

type ClearRecyclebinResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ClearRecyclebinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ClearRecyclebinResponse) String() string {
	return tea.Prettify(s)
}

func (s ClearRecyclebinResponse) GoString() string {
	return s.String()
}

func (s *ClearRecyclebinResponse) SetHeaders(v map[string]*string) *ClearRecyclebinResponse {
	s.Headers = v
	return s
}

func (s *ClearRecyclebinResponse) SetStatusCode(v int32) *ClearRecyclebinResponse {
	s.StatusCode = &v
	return s
}

func (s *ClearRecyclebinResponse) SetBody(v *ClearRecyclebinResponseBody) *ClearRecyclebinResponse {
	s.Body = v
	return s
}

type CompleteFileRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The upload ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// C9DCFE5A82644AC7A02DB74C30C934A6
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s CompleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CompleteFileRequest) GoString() string {
	return s.String()
}

func (s *CompleteFileRequest) SetDriveId(v string) *CompleteFileRequest {
	s.DriveId = &v
	return s
}

func (s *CompleteFileRequest) SetFileId(v string) *CompleteFileRequest {
	s.FileId = &v
	return s
}

func (s *CompleteFileRequest) SetUploadId(v string) *CompleteFileRequest {
	s.UploadId = &v
	return s
}

type CompleteFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *File              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CompleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CompleteFileResponse) GoString() string {
	return s.String()
}

func (s *CompleteFileResponse) SetHeaders(v map[string]*string) *CompleteFileResponse {
	s.Headers = v
	return s
}

func (s *CompleteFileResponse) SetStatusCode(v int32) *CompleteFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CompleteFileResponse) SetBody(v *File) *CompleteFileResponse {
	s.Body = v
	return s
}

type CopyFileRequest struct {
	// Specifies whether to automatically rename the file if the file name already exists in the destination folder. Default value: false.
	//
	// example:
	//
	// true
	AutoRename *bool `json:"auto_rename,omitempty" xml:"auto_rename,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID or folder ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The share ID. If you want to manage a file by using a share link, carry the `x-share-token` header for authentication in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify one of `share_id` and `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The ID of the drive to which you want to copy the file or folder. Default value: the value of drive_id.
	//
	// example:
	//
	// 1
	ToDriveId *string `json:"to_drive_id,omitempty" xml:"to_drive_id,omitempty"`
	// The ID of the destination parent folder. If you want to copy the file or folder to a root directory, set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6520943DC261
	ToParentFileId *string `json:"to_parent_file_id,omitempty" xml:"to_parent_file_id,omitempty"`
}

func (s CopyFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CopyFileRequest) GoString() string {
	return s.String()
}

func (s *CopyFileRequest) SetAutoRename(v bool) *CopyFileRequest {
	s.AutoRename = &v
	return s
}

func (s *CopyFileRequest) SetDriveId(v string) *CopyFileRequest {
	s.DriveId = &v
	return s
}

func (s *CopyFileRequest) SetFileId(v string) *CopyFileRequest {
	s.FileId = &v
	return s
}

func (s *CopyFileRequest) SetShareId(v string) *CopyFileRequest {
	s.ShareId = &v
	return s
}

func (s *CopyFileRequest) SetToDriveId(v string) *CopyFileRequest {
	s.ToDriveId = &v
	return s
}

func (s *CopyFileRequest) SetToParentFileId(v string) *CopyFileRequest {
	s.ToParentFileId = &v
	return s
}

type CopyFileResponseBody struct {
	// The ID of the asynchronous task.
	//
	// If a file is copied, this parameter is not returned. If a folder is copied, the folder is asynchronously copied in the background and this parameter is returned. You can call the GetAsyncTask operation to query the information about the asynchronous task based on the task ID.
	//
	// example:
	//
	// 000e89fb-cf8f-11e9-8ab4-b6e980803a3b
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the copied file or folder.
	//
	// example:
	//
	// 4221bf6e6ab43a255edc4463bffa6f5f5d317401
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s CopyFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CopyFileResponseBody) GoString() string {
	return s.String()
}

func (s *CopyFileResponseBody) SetAsyncTaskId(v string) *CopyFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *CopyFileResponseBody) SetDomainId(v string) *CopyFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *CopyFileResponseBody) SetDriveId(v string) *CopyFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *CopyFileResponseBody) SetFileId(v string) *CopyFileResponseBody {
	s.FileId = &v
	return s
}

type CopyFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CopyFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CopyFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CopyFileResponse) GoString() string {
	return s.String()
}

func (s *CopyFileResponse) SetHeaders(v map[string]*string) *CopyFileResponse {
	s.Headers = v
	return s
}

func (s *CopyFileResponse) SetStatusCode(v int32) *CopyFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CopyFileResponse) SetBody(v *CopyFileResponseBody) *CopyFileResponse {
	s.Body = v
	return s
}

type CreateCustomizedStoryRequest struct {
	// Deprecated
	CustomLabels map[string]*string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	StoryCover *CreateCustomizedStoryRequestStoryCover `json:"story_cover,omitempty" xml:"story_cover,omitempty" type:"Struct"`
	// This parameter is required.
	StoryFiles []*CreateCustomizedStoryRequestStoryFiles `json:"story_files,omitempty" xml:"story_files,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// test_name
	StoryName *string `json:"story_name,omitempty" xml:"story_name,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_created
	StorySubType *string `json:"story_sub_type,omitempty" xml:"story_sub_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// user_created
	StoryType *string `json:"story_type,omitempty" xml:"story_type,omitempty"`
}

func (s CreateCustomizedStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryRequest) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequest) SetCustomLabels(v map[string]*string) *CreateCustomizedStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetDriveId(v string) *CreateCustomizedStoryRequest {
	s.DriveId = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryCover(v *CreateCustomizedStoryRequestStoryCover) *CreateCustomizedStoryRequest {
	s.StoryCover = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryFiles(v []*CreateCustomizedStoryRequestStoryFiles) *CreateCustomizedStoryRequest {
	s.StoryFiles = v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryName(v string) *CreateCustomizedStoryRequest {
	s.StoryName = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStorySubType(v string) *CreateCustomizedStoryRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateCustomizedStoryRequest) SetStoryType(v string) *CreateCustomizedStoryRequest {
	s.StoryType = &v
	return s
}

type CreateCustomizedStoryRequestStoryCover struct {
	// This parameter is required.
	//
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88d4aff041ee68fd4fc89beb80e1119da343
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s CreateCustomizedStoryRequestStoryCover) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryRequestStoryCover) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequestStoryCover) SetFileId(v string) *CreateCustomizedStoryRequestStoryCover {
	s.FileId = &v
	return s
}

func (s *CreateCustomizedStoryRequestStoryCover) SetRevisionId(v string) *CreateCustomizedStoryRequestStoryCover {
	s.RevisionId = &v
	return s
}

type CreateCustomizedStoryRequestStoryFiles struct {
	// This parameter is required.
	//
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88d4aff041ee68fd4fc89beb80e1119da343
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s CreateCustomizedStoryRequestStoryFiles) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryRequestStoryFiles) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryRequestStoryFiles) SetFileId(v string) *CreateCustomizedStoryRequestStoryFiles {
	s.FileId = &v
	return s
}

func (s *CreateCustomizedStoryRequestStoryFiles) SetRevisionId(v string) *CreateCustomizedStoryRequestStoryFiles {
	s.RevisionId = &v
	return s
}

type CreateCustomizedStoryResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s CreateCustomizedStoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryResponseBody) SetDriveId(v string) *CreateCustomizedStoryResponseBody {
	s.DriveId = &v
	return s
}

func (s *CreateCustomizedStoryResponseBody) SetStoryId(v string) *CreateCustomizedStoryResponseBody {
	s.StoryId = &v
	return s
}

type CreateCustomizedStoryResponse struct {
	Headers    map[string]*string                 `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateCustomizedStoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateCustomizedStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateCustomizedStoryResponse) GoString() string {
	return s.String()
}

func (s *CreateCustomizedStoryResponse) SetHeaders(v map[string]*string) *CreateCustomizedStoryResponse {
	s.Headers = v
	return s
}

func (s *CreateCustomizedStoryResponse) SetStatusCode(v int32) *CreateCustomizedStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateCustomizedStoryResponse) SetBody(v *CreateCustomizedStoryResponseBody) *CreateCustomizedStoryResponse {
	s.Body = v
	return s
}

type CreateDomainRequest struct {
	// domain 描述
	//
	// example:
	//
	// 你好企业网盘开发环境
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// If you want to perform secondary operations based on Drive and Photo Service and perform fine-grained control on your tenants, you can use the parent-child domain feature of Drive and Photo Service. For more information, join the DingTalk group whose ID is 23146118.
	//
	// This parameter is required.
	//
	// example:
	//
	// 你好企业网盘
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	// https
	//
	// example:
	//
	// true
	InitDriveEnable *bool `json:"init_drive_enable,omitempty" xml:"init_drive_enable,omitempty"`
	// http
	//
	// example:
	//
	// 1073741824
	InitDriveSize *int64 `json:"init_drive_size,omitempty" xml:"init_drive_size,omitempty"`
	// Create domain.
	//
	// example:
	//
	// bj1
	ParentDomainId *string `json:"parent_domain_id,omitempty" xml:"parent_domain_id,omitempty"`
	// The ID of the parent domain. If you want to create a child domain, specify parent_domain_id. In most cases, you do not need to create a child domain. If you want to perform secondary operations based on Drive and Photo Service, contact the customer service.
	//
	// example:
	//
	// 1099511627776
	SizeQuota *int64 `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	// The information about the domain.
	//
	// example:
	//
	// 50
	UserCountQuota *int64 `json:"user_count_quota,omitempty" xml:"user_count_quota,omitempty"`
}

func (s CreateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainRequest) GoString() string {
	return s.String()
}

func (s *CreateDomainRequest) SetDescription(v string) *CreateDomainRequest {
	s.Description = &v
	return s
}

func (s *CreateDomainRequest) SetDomainName(v string) *CreateDomainRequest {
	s.DomainName = &v
	return s
}

func (s *CreateDomainRequest) SetInitDriveEnable(v bool) *CreateDomainRequest {
	s.InitDriveEnable = &v
	return s
}

func (s *CreateDomainRequest) SetInitDriveSize(v int64) *CreateDomainRequest {
	s.InitDriveSize = &v
	return s
}

func (s *CreateDomainRequest) SetParentDomainId(v string) *CreateDomainRequest {
	s.ParentDomainId = &v
	return s
}

func (s *CreateDomainRequest) SetSizeQuota(v int64) *CreateDomainRequest {
	s.SizeQuota = &v
	return s
}

func (s *CreateDomainRequest) SetUserCountQuota(v int64) *CreateDomainRequest {
	s.UserCountQuota = &v
	return s
}

type CreateDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Domain            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDomainResponse) GoString() string {
	return s.String()
}

func (s *CreateDomainResponse) SetHeaders(v map[string]*string) *CreateDomainResponse {
	s.Headers = v
	return s
}

func (s *CreateDomainResponse) SetStatusCode(v int32) *CreateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDomainResponse) SetBody(v *Domain) *CreateDomainResponse {
	s.Body = v
	return s
}

type CreateDriveRequest struct {
	// Specifies whether the drive is the default drive. Default value: false.
	//
	// example:
	//
	// true
	Default *bool `json:"default,omitempty" xml:"default,omitempty"`
	// The description of the drive. The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// drive for test
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the drive. The name can be up to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test_drive
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// The type of the drive. Set the value to normal.
	//
	// example:
	//
	// normal
	DriveType *string `json:"drive_type,omitempty" xml:"drive_type,omitempty"`
	// The owner of the drive.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3b3d7245c159488da17d081ad6c64687
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The type of the owner. Valid values:
	//
	// user and group.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	// The state of the drive. Valid values:
	//
	// enabled and disabled.
	//
	// Default value: enabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total size of the drive. Unit: bytes. By default, the size is unlimited.
	//
	// example:
	//
	// 1024
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s CreateDriveRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateDriveRequest) GoString() string {
	return s.String()
}

func (s *CreateDriveRequest) SetDefault(v bool) *CreateDriveRequest {
	s.Default = &v
	return s
}

func (s *CreateDriveRequest) SetDescription(v string) *CreateDriveRequest {
	s.Description = &v
	return s
}

func (s *CreateDriveRequest) SetDriveName(v string) *CreateDriveRequest {
	s.DriveName = &v
	return s
}

func (s *CreateDriveRequest) SetDriveType(v string) *CreateDriveRequest {
	s.DriveType = &v
	return s
}

func (s *CreateDriveRequest) SetOwner(v string) *CreateDriveRequest {
	s.Owner = &v
	return s
}

func (s *CreateDriveRequest) SetOwnerType(v string) *CreateDriveRequest {
	s.OwnerType = &v
	return s
}

func (s *CreateDriveRequest) SetStatus(v string) *CreateDriveRequest {
	s.Status = &v
	return s
}

func (s *CreateDriveRequest) SetTotalSize(v int64) *CreateDriveRequest {
	s.TotalSize = &v
	return s
}

type CreateDriveResponseBody struct {
	CreatedAt   *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator     *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId   *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	DriveType *string `json:"drive_type,omitempty" xml:"drive_type,omitempty"`
	Owner     *string `json:"owner,omitempty" xml:"owner,omitempty"`
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
	TotalSize *int64  `json:"total_size,omitempty" xml:"total_size,omitempty"`
	UsedSize  *int64  `json:"used_size,omitempty" xml:"used_size,omitempty"`
}

func (s CreateDriveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDriveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDriveResponseBody) SetCreatedAt(v string) *CreateDriveResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateDriveResponseBody) SetCreator(v string) *CreateDriveResponseBody {
	s.Creator = &v
	return s
}

func (s *CreateDriveResponseBody) SetDescription(v string) *CreateDriveResponseBody {
	s.Description = &v
	return s
}

func (s *CreateDriveResponseBody) SetDomainId(v string) *CreateDriveResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateDriveResponseBody) SetDriveId(v string) *CreateDriveResponseBody {
	s.DriveId = &v
	return s
}

func (s *CreateDriveResponseBody) SetDriveName(v string) *CreateDriveResponseBody {
	s.DriveName = &v
	return s
}

func (s *CreateDriveResponseBody) SetDriveType(v string) *CreateDriveResponseBody {
	s.DriveType = &v
	return s
}

func (s *CreateDriveResponseBody) SetOwner(v string) *CreateDriveResponseBody {
	s.Owner = &v
	return s
}

func (s *CreateDriveResponseBody) SetOwnerType(v string) *CreateDriveResponseBody {
	s.OwnerType = &v
	return s
}

func (s *CreateDriveResponseBody) SetStatus(v string) *CreateDriveResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDriveResponseBody) SetTotalSize(v int64) *CreateDriveResponseBody {
	s.TotalSize = &v
	return s
}

func (s *CreateDriveResponseBody) SetUsedSize(v int64) *CreateDriveResponseBody {
	s.UsedSize = &v
	return s
}

type CreateDriveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateDriveResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateDriveResponse) GoString() string {
	return s.String()
}

func (s *CreateDriveResponse) SetHeaders(v map[string]*string) *CreateDriveResponse {
	s.Headers = v
	return s
}

func (s *CreateDriveResponse) SetStatusCode(v int32) *CreateDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateDriveResponse) SetBody(v *CreateDriveResponseBody) *CreateDriveResponse {
	s.Body = v
	return s
}

type CreateFileRequest struct {
	// The processing method that is used if the file that you want to create has the same name as an existing file in the cloud. Valid values:
	//
	// ignore: allows you to create the file by using the same name as an existing file in the cloud.
	//
	// auto_rename: automatically renames the file that you want to create. By default, the current point in time is added to the end of the file name. Example: xxx_20060102_150405.
	//
	// refuse: does not create the file that you want to create but returns the information about the file that has the same name in the cloud.
	//
	// Default value: ignore.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// ignore
	CheckNameMode *string `json:"check_name_mode,omitempty" xml:"check_name_mode,omitempty"`
	// The hash value of the file content. The value is calculated based on the algorithm specified by content_hash_name.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 7C4A8D09CA3762AF61E59520943DC26494F8941B
	ContentHash *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	// The name of the algorithm that is used to calculate the hash value of the file content. Only SHA1 is supported.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// sha1
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	// The type of the file content. Default value: application/oct-stream.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// application/json
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// The description of the file. The description can be up to 1,024 characters in length. By default, this parameter is left empty.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 重要文件
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The drive ID. This parameter is required if the file is not uploaded by using the share URL of the file.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID. This parameter is required if check_name_mode is set to ignore.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// Specifies whether to hide the file or folder. By default, the file or folder is not hidden.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// false
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// The information about the image specified by the client.
	ImageMediaMetadata *ImageMediaMetadata `json:"image_media_metadata,omitempty" xml:"image_media_metadata,omitempty"`
	// The time when the local file was created. By default, this parameter is left empty. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format based on the UTC+0 time zone.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	LocalCreatedAt *string `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	// The time when the local file was modified. By default, this parameter is left empty. Specify the time in the yyyy-MM-ddTHH:mm:ssZ format based on the UTC+0 time zone.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	LocalModifiedAt *string `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	// The name of the file. The name can be up to 1,024 bytes in length based on the UTF-8 encoding rule and cannot end with a forward slash (/).
	//
	// This parameter is required.
	//
	// if can be null:
	// false
	//
	// example:
	//
	// a.txt
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to enable the parallel upload feature.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// true
	ParallelUpload *bool `json:"parallel_upload,omitempty" xml:"parallel_upload,omitempty"`
	// The ID of the parent directory. If you want to create a file or folder in the root directory, set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// fileid1
	ParentFileId *string `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	// The information about the file parts. You can specify up to 10,000 parts. By default, if you do not specify this parameter, only one part is returned.
	PartInfoList []*CreateFileRequestPartInfoList `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	// The SHA-1 hash value of the first 1 KB data of the file. This parameter is required if you perform instant file upload by using the pre-hashing feature. If the SHA-1 hash value is not matched on the cloud, the client does not need to calculate the SHA-1 hash value of the entire file.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 7C4A8D09CA3762AF61E59520943DC26494F89411
	PreHash *string `json:"pre_hash,omitempty" xml:"pre_hash,omitempty"`
	// The share ID. This parameter is required if the file is uploaded by using the share URL of the file.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// 1024
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The type of the file. Valid values:
	//
	// file folder
	//
	// This parameter is required.
	//
	// if can be null:
	// true
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The custom tags. You can specify up to 1,000 tags.
	UserTags []*UserTag `json:"user_tags,omitempty" xml:"user_tags,omitempty" type:"Repeated"`
	// The information about the video specified by the client.
	VideoMediaMetadata *VideoMediaMetadata `json:"video_media_metadata,omitempty" xml:"video_media_metadata,omitempty"`
}

func (s CreateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateFileRequest) GoString() string {
	return s.String()
}

func (s *CreateFileRequest) SetCheckNameMode(v string) *CreateFileRequest {
	s.CheckNameMode = &v
	return s
}

func (s *CreateFileRequest) SetContentHash(v string) *CreateFileRequest {
	s.ContentHash = &v
	return s
}

func (s *CreateFileRequest) SetContentHashName(v string) *CreateFileRequest {
	s.ContentHashName = &v
	return s
}

func (s *CreateFileRequest) SetContentType(v string) *CreateFileRequest {
	s.ContentType = &v
	return s
}

func (s *CreateFileRequest) SetDescription(v string) *CreateFileRequest {
	s.Description = &v
	return s
}

func (s *CreateFileRequest) SetDriveId(v string) *CreateFileRequest {
	s.DriveId = &v
	return s
}

func (s *CreateFileRequest) SetFileId(v string) *CreateFileRequest {
	s.FileId = &v
	return s
}

func (s *CreateFileRequest) SetHidden(v bool) *CreateFileRequest {
	s.Hidden = &v
	return s
}

func (s *CreateFileRequest) SetImageMediaMetadata(v *ImageMediaMetadata) *CreateFileRequest {
	s.ImageMediaMetadata = v
	return s
}

func (s *CreateFileRequest) SetLocalCreatedAt(v string) *CreateFileRequest {
	s.LocalCreatedAt = &v
	return s
}

func (s *CreateFileRequest) SetLocalModifiedAt(v string) *CreateFileRequest {
	s.LocalModifiedAt = &v
	return s
}

func (s *CreateFileRequest) SetName(v string) *CreateFileRequest {
	s.Name = &v
	return s
}

func (s *CreateFileRequest) SetParallelUpload(v bool) *CreateFileRequest {
	s.ParallelUpload = &v
	return s
}

func (s *CreateFileRequest) SetParentFileId(v string) *CreateFileRequest {
	s.ParentFileId = &v
	return s
}

func (s *CreateFileRequest) SetPartInfoList(v []*CreateFileRequestPartInfoList) *CreateFileRequest {
	s.PartInfoList = v
	return s
}

func (s *CreateFileRequest) SetPreHash(v string) *CreateFileRequest {
	s.PreHash = &v
	return s
}

func (s *CreateFileRequest) SetShareId(v string) *CreateFileRequest {
	s.ShareId = &v
	return s
}

func (s *CreateFileRequest) SetSize(v int64) *CreateFileRequest {
	s.Size = &v
	return s
}

func (s *CreateFileRequest) SetType(v string) *CreateFileRequest {
	s.Type = &v
	return s
}

func (s *CreateFileRequest) SetUserTags(v []*UserTag) *CreateFileRequest {
	s.UserTags = v
	return s
}

func (s *CreateFileRequest) SetVideoMediaMetadata(v *VideoMediaMetadata) *CreateFileRequest {
	s.VideoMediaMetadata = v
	return s
}

type CreateFileRequestPartInfoList struct {
	// The MD5 hash value of the file part. This parameter is required when the MD5 hash value of the file part needs to be verified during part upload.
	//
	// example:
	//
	// ASKJDJSKDJJSJDJS
	ContentMd5 *string `json:"content_md5,omitempty" xml:"content_md5,omitempty"`
	// The SHA-1 hash value of the file content before the file part. This parameter takes effect only if the parallel upload feature is enabled.
	ParallelSha1Ctx *CreateFileRequestPartInfoListParallelSha1Ctx `json:"parallel_sha1_ctx,omitempty" xml:"parallel_sha1_ctx,omitempty" type:"Struct"`
	// The serial number of a file part. The number starts from 1.
	//
	// example:
	//
	// 1
	PartNumber *int32 `json:"part_number,omitempty" xml:"part_number,omitempty"`
}

func (s CreateFileRequestPartInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreateFileRequestPartInfoList) GoString() string {
	return s.String()
}

func (s *CreateFileRequestPartInfoList) SetContentMd5(v string) *CreateFileRequestPartInfoList {
	s.ContentMd5 = &v
	return s
}

func (s *CreateFileRequestPartInfoList) SetParallelSha1Ctx(v *CreateFileRequestPartInfoListParallelSha1Ctx) *CreateFileRequestPartInfoList {
	s.ParallelSha1Ctx = v
	return s
}

func (s *CreateFileRequestPartInfoList) SetPartNumber(v int32) *CreateFileRequestPartInfoList {
	s.PartNumber = &v
	return s
}

type CreateFileRequestPartInfoListParallelSha1Ctx struct {
	// The first to fifth 32-bit variables of the SHA-1 hash value of the file content before the file part. This parameter takes effect only if the parallel upload feature is enabled.
	H []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	// The size of the file content before the file part. Unit: bytes. The value must be a multiple of 64. This parameter takes effect only if the parallel upload feature is enabled.
	//
	// example:
	//
	// 10240
	PartOffset *int64 `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s CreateFileRequestPartInfoListParallelSha1Ctx) String() string {
	return tea.Prettify(s)
}

func (s CreateFileRequestPartInfoListParallelSha1Ctx) GoString() string {
	return s.String()
}

func (s *CreateFileRequestPartInfoListParallelSha1Ctx) SetH(v []*int64) *CreateFileRequestPartInfoListParallelSha1Ctx {
	s.H = v
	return s
}

func (s *CreateFileRequestPartInfoListParallelSha1Ctx) SetPartOffset(v int64) *CreateFileRequestPartInfoListParallelSha1Ctx {
	s.PartOffset = &v
	return s
}

type CreateFileResponseBody struct {
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Indicates whether the file exists.
	//
	// example:
	//
	// false
	Exist *bool `json:"exist,omitempty" xml:"exist,omitempty"`
	// The file ID.
	//
	// example:
	//
	// fileid1
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The file name.
	//
	// example:
	//
	// a.txt
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// The ID of the parent directory.
	//
	// example:
	//
	// fileid5
	ParentFileId *string `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	// The information about the file parts.
	PartInfoList []*UploadPartInfo `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	// Indicates whether the file is instantly uploaded.
	//
	// example:
	//
	// true
	RapidUpload *bool `json:"rapid_upload,omitempty" xml:"rapid_upload,omitempty"`
	// The state of the file.
	//
	// example:
	//
	// uploading
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The type of the file.
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The ID of the upload task.
	//
	// example:
	//
	// uploadid1
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s CreateFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponseBody) GoString() string {
	return s.String()
}

func (s *CreateFileResponseBody) SetDomainId(v string) *CreateFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateFileResponseBody) SetDriveId(v string) *CreateFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *CreateFileResponseBody) SetExist(v bool) *CreateFileResponseBody {
	s.Exist = &v
	return s
}

func (s *CreateFileResponseBody) SetFileId(v string) *CreateFileResponseBody {
	s.FileId = &v
	return s
}

func (s *CreateFileResponseBody) SetFileName(v string) *CreateFileResponseBody {
	s.FileName = &v
	return s
}

func (s *CreateFileResponseBody) SetParentFileId(v string) *CreateFileResponseBody {
	s.ParentFileId = &v
	return s
}

func (s *CreateFileResponseBody) SetPartInfoList(v []*UploadPartInfo) *CreateFileResponseBody {
	s.PartInfoList = v
	return s
}

func (s *CreateFileResponseBody) SetRapidUpload(v bool) *CreateFileResponseBody {
	s.RapidUpload = &v
	return s
}

func (s *CreateFileResponseBody) SetStatus(v string) *CreateFileResponseBody {
	s.Status = &v
	return s
}

func (s *CreateFileResponseBody) SetType(v string) *CreateFileResponseBody {
	s.Type = &v
	return s
}

func (s *CreateFileResponseBody) SetUploadId(v string) *CreateFileResponseBody {
	s.UploadId = &v
	return s
}

type CreateFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateFileResponse) GoString() string {
	return s.String()
}

func (s *CreateFileResponse) SetHeaders(v map[string]*string) *CreateFileResponse {
	s.Headers = v
	return s
}

func (s *CreateFileResponse) SetStatusCode(v int32) *CreateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateFileResponse) SetBody(v *CreateFileResponseBody) *CreateFileResponse {
	s.Body = v
	return s
}

type CreateGroupRequest struct {
	// The description of the group. The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// test group description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The name of the group. The name must be 1 to 128 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// test group
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// Specifies whether the group is a root group. A root group cannot be added to any other group. In most cases, a root group is the top-level organization in the organizational structure.
	//
	// example:
	//
	// false
	IsRoot *bool `json:"is_root,omitempty" xml:"is_root,omitempty"`
	// The ID of the parent group to which the group is added. If this parameter is specified, the system automatically adds the group to the specified parent group after the group is created.
	//
	// example:
	//
	// 2e43ec8427dd45f19431b7504649a1b3
	ParentGroupId *string `json:"parent_group_id,omitempty" xml:"parent_group_id,omitempty"`
}

func (s CreateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupRequest) GoString() string {
	return s.String()
}

func (s *CreateGroupRequest) SetDescription(v string) *CreateGroupRequest {
	s.Description = &v
	return s
}

func (s *CreateGroupRequest) SetGroupName(v string) *CreateGroupRequest {
	s.GroupName = &v
	return s
}

func (s *CreateGroupRequest) SetIsRoot(v bool) *CreateGroupRequest {
	s.IsRoot = &v
	return s
}

func (s *CreateGroupRequest) SetParentGroupId(v string) *CreateGroupRequest {
	s.ParentGroupId = &v
	return s
}

type CreateGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Group             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateGroupResponse) GoString() string {
	return s.String()
}

func (s *CreateGroupResponse) SetHeaders(v map[string]*string) *CreateGroupResponse {
	s.Headers = v
	return s
}

func (s *CreateGroupResponse) SetStatusCode(v int32) *CreateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateGroupResponse) SetBody(v *Group) *CreateGroupResponse {
	s.Body = v
	return s
}

type CreateIdentityToBenefitPkgMappingRequest struct {
	// The number of benefit packages.
	//
	// This parameter takes effect only for the benefit packages of the resource type. Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The unique identifier of the benefit package.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40cb7794c9294
	BenefitPkgId *string `json:"benefit_pkg_id,omitempty" xml:"benefit_pkg_id,omitempty"`
	// The time when the benefit package expires. Set the value to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// By default, the benefit package never expires.
	//
	// example:
	//
	// 1633167071000
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// The unique identifier of the entity.
	//
	// If you want to manage the benefits of a user, set this parameter to a user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user123
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// The type of the entity.
	//
	// If you want to manage the benefits of a user, set this parameter to user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
}

func (s CreateIdentityToBenefitPkgMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateIdentityToBenefitPkgMappingRequest) GoString() string {
	return s.String()
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetAmount(v int64) *CreateIdentityToBenefitPkgMappingRequest {
	s.Amount = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetBenefitPkgId(v string) *CreateIdentityToBenefitPkgMappingRequest {
	s.BenefitPkgId = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetExpireTime(v int64) *CreateIdentityToBenefitPkgMappingRequest {
	s.ExpireTime = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetIdentityId(v string) *CreateIdentityToBenefitPkgMappingRequest {
	s.IdentityId = &v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingRequest) SetIdentityType(v string) *CreateIdentityToBenefitPkgMappingRequest {
	s.IdentityType = &v
	return s
}

type CreateIdentityToBenefitPkgMappingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s CreateIdentityToBenefitPkgMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateIdentityToBenefitPkgMappingResponse) GoString() string {
	return s.String()
}

func (s *CreateIdentityToBenefitPkgMappingResponse) SetHeaders(v map[string]*string) *CreateIdentityToBenefitPkgMappingResponse {
	s.Headers = v
	return s
}

func (s *CreateIdentityToBenefitPkgMappingResponse) SetStatusCode(v int32) *CreateIdentityToBenefitPkgMappingResponse {
	s.StatusCode = &v
	return s
}

type CreateOrderRequest struct {
	AutoPay   *bool `json:"auto_pay,omitempty" xml:"auto_pay,omitempty"`
	AutoRenew *bool `json:"auto_renew,omitempty" xml:"auto_renew,omitempty"`
	// This parameter is required.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// This parameter is required.
	OrderType *string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// This parameter is required.
	Package *string `json:"package,omitempty" xml:"package,omitempty"`
	// This parameter is required.
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// This parameter is required.
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	// This parameter is required.
	UserCount *int64 `json:"user_count,omitempty" xml:"user_count,omitempty"`
}

func (s CreateOrderRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderRequest) GoString() string {
	return s.String()
}

func (s *CreateOrderRequest) SetAutoPay(v bool) *CreateOrderRequest {
	s.AutoPay = &v
	return s
}

func (s *CreateOrderRequest) SetAutoRenew(v bool) *CreateOrderRequest {
	s.AutoRenew = &v
	return s
}

func (s *CreateOrderRequest) SetCode(v string) *CreateOrderRequest {
	s.Code = &v
	return s
}

func (s *CreateOrderRequest) SetInstanceId(v string) *CreateOrderRequest {
	s.InstanceId = &v
	return s
}

func (s *CreateOrderRequest) SetOrderType(v string) *CreateOrderRequest {
	s.OrderType = &v
	return s
}

func (s *CreateOrderRequest) SetPackage(v string) *CreateOrderRequest {
	s.Package = &v
	return s
}

func (s *CreateOrderRequest) SetPeriod(v int64) *CreateOrderRequest {
	s.Period = &v
	return s
}

func (s *CreateOrderRequest) SetPeriodUnit(v string) *CreateOrderRequest {
	s.PeriodUnit = &v
	return s
}

func (s *CreateOrderRequest) SetTotalSize(v int64) *CreateOrderRequest {
	s.TotalSize = &v
	return s
}

func (s *CreateOrderRequest) SetUserCount(v int64) *CreateOrderRequest {
	s.UserCount = &v
	return s
}

type CreateOrderResponseBody struct {
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	OrderId    *string `json:"order_id,omitempty" xml:"order_id,omitempty"`
}

func (s CreateOrderResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOrderResponseBody) SetInstanceId(v string) *CreateOrderResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateOrderResponseBody) SetOrderId(v string) *CreateOrderResponseBody {
	s.OrderId = &v
	return s
}

type CreateOrderResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateOrderResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateOrderResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateOrderResponse) GoString() string {
	return s.String()
}

func (s *CreateOrderResponse) SetHeaders(v map[string]*string) *CreateOrderResponse {
	s.Headers = v
	return s
}

func (s *CreateOrderResponse) SetStatusCode(v int32) *CreateOrderResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateOrderResponse) SetBody(v *CreateOrderResponseBody) *CreateOrderResponse {
	s.Body = v
	return s
}

type CreateShareLinkRequest struct {
	Creatable           *bool     `json:"creatable,omitempty" xml:"creatable,omitempty"`
	CreatableFileIdList []*string `json:"creatable_file_id_list,omitempty" xml:"creatable_file_id_list,omitempty" type:"Repeated"`
	// The description of the share. The description must be 0 to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to disable the download feature.
	//
	// example:
	//
	// false
	DisableDownload *bool `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	// Specifies whether to disable the preview feature.
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	// Specifies whether to disable the dump feature.
	//
	// example:
	//
	// false
	DisableSave *bool `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	// The limit on the number of times that the shared files can be downloaded. The value of this parameter must be equal to or greater than 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The time when the share URL expires. The value of this parameter follows the RFC 3339 standard. Example: "2020-06-28T11:33:00.000+08:00". If expiration is set to "", the share URL never expires.
	//
	// example:
	//
	// 2020-06-28T11:33:00.000+08:00
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// The IDs of the files to share in the parent path. The number of files in the parent path ranges from 1 to 100. If share_all_files is set to true, this parameter does not take effect. Otherwise, you must specify this parameter.``
	//
	// example:
	//
	// ["520b217f13adf4fc24f2191991b1664ce045b393"]
	FileIdList     []*string `json:"file_id_list,omitempty" xml:"file_id_list,omitempty" type:"Repeated"`
	OfficeEditable *bool     `json:"office_editable,omitempty" xml:"office_editable,omitempty"`
	// The limit on the number of times that the shared files can be previewed. The value of this parameter must be equal to or greater than 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	RequireLogin *bool  `json:"require_login,omitempty" xml:"require_login,omitempty"`
	// The limit on the number of times that the shared files can be dumped. The value of this parameter must be equal to or greater than 0. A value of 0 indicates no limit.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	// Specifies whether to share all files in the drive.
	//
	// example:
	//
	// true
	ShareAllFiles *bool `json:"share_all_files,omitempty" xml:"share_all_files,omitempty"`
	// The name of the share. If you leave this parameter empty, the file name that corresponds to the first ID in the file ID list is used. The name must be 0 to 128 characters in length.
	ShareName *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	// The access code. An access code must be 0 to 64 bytes in length. If you do not specify this parameter or leave this parameter empty, the files can be accessed without an access code. In this case, you do not need to specify the share_pwd parameter when you call an operation to query the share URL. The access code can contain only visible ASCII characters.
	//
	// example:
	//
	// abcF123x
	SharePwd *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
	// The user ID.
	//
	// example:
	//
	// u123
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CreateShareLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateShareLinkRequest) GoString() string {
	return s.String()
}

func (s *CreateShareLinkRequest) SetCreatable(v bool) *CreateShareLinkRequest {
	s.Creatable = &v
	return s
}

func (s *CreateShareLinkRequest) SetCreatableFileIdList(v []*string) *CreateShareLinkRequest {
	s.CreatableFileIdList = v
	return s
}

func (s *CreateShareLinkRequest) SetDescription(v string) *CreateShareLinkRequest {
	s.Description = &v
	return s
}

func (s *CreateShareLinkRequest) SetDisableDownload(v bool) *CreateShareLinkRequest {
	s.DisableDownload = &v
	return s
}

func (s *CreateShareLinkRequest) SetDisablePreview(v bool) *CreateShareLinkRequest {
	s.DisablePreview = &v
	return s
}

func (s *CreateShareLinkRequest) SetDisableSave(v bool) *CreateShareLinkRequest {
	s.DisableSave = &v
	return s
}

func (s *CreateShareLinkRequest) SetDownloadLimit(v int64) *CreateShareLinkRequest {
	s.DownloadLimit = &v
	return s
}

func (s *CreateShareLinkRequest) SetDriveId(v string) *CreateShareLinkRequest {
	s.DriveId = &v
	return s
}

func (s *CreateShareLinkRequest) SetExpiration(v string) *CreateShareLinkRequest {
	s.Expiration = &v
	return s
}

func (s *CreateShareLinkRequest) SetFileIdList(v []*string) *CreateShareLinkRequest {
	s.FileIdList = v
	return s
}

func (s *CreateShareLinkRequest) SetOfficeEditable(v bool) *CreateShareLinkRequest {
	s.OfficeEditable = &v
	return s
}

func (s *CreateShareLinkRequest) SetPreviewLimit(v int64) *CreateShareLinkRequest {
	s.PreviewLimit = &v
	return s
}

func (s *CreateShareLinkRequest) SetRequireLogin(v bool) *CreateShareLinkRequest {
	s.RequireLogin = &v
	return s
}

func (s *CreateShareLinkRequest) SetSaveLimit(v int64) *CreateShareLinkRequest {
	s.SaveLimit = &v
	return s
}

func (s *CreateShareLinkRequest) SetShareAllFiles(v bool) *CreateShareLinkRequest {
	s.ShareAllFiles = &v
	return s
}

func (s *CreateShareLinkRequest) SetShareName(v string) *CreateShareLinkRequest {
	s.ShareName = &v
	return s
}

func (s *CreateShareLinkRequest) SetSharePwd(v string) *CreateShareLinkRequest {
	s.SharePwd = &v
	return s
}

func (s *CreateShareLinkRequest) SetUserId(v string) *CreateShareLinkRequest {
	s.UserId = &v
	return s
}

type CreateShareLinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareLink         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateShareLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateShareLinkResponse) GoString() string {
	return s.String()
}

func (s *CreateShareLinkResponse) SetHeaders(v map[string]*string) *CreateShareLinkResponse {
	s.Headers = v
	return s
}

func (s *CreateShareLinkResponse) SetStatusCode(v int32) *CreateShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateShareLinkResponse) SetBody(v *ShareLink) *CreateShareLinkResponse {
	s.Body = v
	return s
}

type CreateSimilarImageClusterTaskRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s CreateSimilarImageClusterTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarImageClusterTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusterTaskRequest) SetDriveId(v string) *CreateSimilarImageClusterTaskRequest {
	s.DriveId = &v
	return s
}

type CreateSimilarImageClusterTaskResponseBody struct {
	// example:
	//
	// i:SimilarImageClustering-b67d53e7-2fe8-460f-9b95-1e93636923eb
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s CreateSimilarImageClusterTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarImageClusterTaskResponseBody) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusterTaskResponseBody) SetTaskId(v string) *CreateSimilarImageClusterTaskResponseBody {
	s.TaskId = &v
	return s
}

type CreateSimilarImageClusterTaskResponse struct {
	Headers    map[string]*string                         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateSimilarImageClusterTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateSimilarImageClusterTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateSimilarImageClusterTaskResponse) GoString() string {
	return s.String()
}

func (s *CreateSimilarImageClusterTaskResponse) SetHeaders(v map[string]*string) *CreateSimilarImageClusterTaskResponse {
	s.Headers = v
	return s
}

func (s *CreateSimilarImageClusterTaskResponse) SetStatusCode(v int32) *CreateSimilarImageClusterTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateSimilarImageClusterTaskResponse) SetBody(v *CreateSimilarImageClusterTaskResponseBody) *CreateSimilarImageClusterTaskResponse {
	s.Body = v
	return s
}

type CreateStoryRequest struct {
	Address *Address `json:"address,omitempty" xml:"address,omitempty"`
	// Deprecated
	CustomLabels map[string]*string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 30
	MaxImageCount *int64 `json:"max_image_count,omitempty" xml:"max_image_count,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// 1
	MinImageCount *int64 `json:"min_image_count,omitempty" xml:"min_image_count,omitempty"`
	// example:
	//
	// 2022-12-30T16:00:00Z
	StoryEndTime *string `json:"story_end_time,omitempty" xml:"story_end_time,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId   *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
	StoryName *string `json:"story_name,omitempty" xml:"story_name,omitempty"`
	// example:
	//
	// 2016-12-30T16:00:00Z
	StoryStartTime *string `json:"story_start_time,omitempty" xml:"story_start_time,omitempty"`
	// example:
	//
	// Food
	StorySubType *string `json:"story_sub_type,omitempty" xml:"story_sub_type,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// TagMemory
	StoryType *string `json:"story_type,omitempty" xml:"story_type,omitempty"`
}

func (s CreateStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateStoryRequest) GoString() string {
	return s.String()
}

func (s *CreateStoryRequest) SetAddress(v *Address) *CreateStoryRequest {
	s.Address = v
	return s
}

func (s *CreateStoryRequest) SetCustomLabels(v map[string]*string) *CreateStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *CreateStoryRequest) SetDriveId(v string) *CreateStoryRequest {
	s.DriveId = &v
	return s
}

func (s *CreateStoryRequest) SetMaxImageCount(v int64) *CreateStoryRequest {
	s.MaxImageCount = &v
	return s
}

func (s *CreateStoryRequest) SetMinImageCount(v int64) *CreateStoryRequest {
	s.MinImageCount = &v
	return s
}

func (s *CreateStoryRequest) SetStoryEndTime(v string) *CreateStoryRequest {
	s.StoryEndTime = &v
	return s
}

func (s *CreateStoryRequest) SetStoryId(v string) *CreateStoryRequest {
	s.StoryId = &v
	return s
}

func (s *CreateStoryRequest) SetStoryName(v string) *CreateStoryRequest {
	s.StoryName = &v
	return s
}

func (s *CreateStoryRequest) SetStoryStartTime(v string) *CreateStoryRequest {
	s.StoryStartTime = &v
	return s
}

func (s *CreateStoryRequest) SetStorySubType(v string) *CreateStoryRequest {
	s.StorySubType = &v
	return s
}

func (s *CreateStoryRequest) SetStoryType(v string) *CreateStoryRequest {
	s.StoryType = &v
	return s
}

type CreateStoryResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s CreateStoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateStoryResponseBody) GoString() string {
	return s.String()
}

func (s *CreateStoryResponseBody) SetDriveId(v string) *CreateStoryResponseBody {
	s.DriveId = &v
	return s
}

type CreateStoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateStoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateStoryResponse) GoString() string {
	return s.String()
}

func (s *CreateStoryResponse) SetHeaders(v map[string]*string) *CreateStoryResponse {
	s.Headers = v
	return s
}

func (s *CreateStoryResponse) SetStatusCode(v int32) *CreateStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateStoryResponse) SetBody(v *CreateStoryResponseBody) *CreateStoryResponse {
	s.Body = v
	return s
}

type CreateUserRequest struct {
	// The URL of the profile picture.
	//
	// If you specify the parameter in the HTTP URL format, the URL must start with http:// or https:// and can be up to 4 KB in size.
	//
	// If you specify the parameter in the data URL format, the URL must start with data:// and be encoded in Base64. The URL can be up to 300 KB in size.
	//
	// example:
	//
	// http://a.b.c/pds.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The description of the user. The description can be up to 1,024 characters in length.
	//
	// example:
	//
	// The VIP user
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The email address.
	//
	// example:
	//
	// 123@pds.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The information about the group.
	GroupInfoList []*CreateUserRequestGroupInfoList `json:"group_info_list,omitempty" xml:"group_info_list,omitempty" type:"Repeated"`
	// The nickname of the user. The nickname can be up to 128 characters in length.
	//
	// example:
	//
	// pdsuer
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 13900001111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The role of the user. Default value: user. Valid values:
	//
	// 	- superadmin
	//
	// 	- admin
	//
	// 	- user
	//
	// If the domain can be divided into subdomains, the subdomain_super_admin and subdomain_admin roles are also supported.
	//
	// Valid values:
	//
	// 	- subdomain_super_admin
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- subdomain_admin
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- superadmin
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- admin
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- user
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The state of the user. Default value: enabled. Valid values:
	//
	// 	- enabled: The user is in a normal state.
	//
	// 	- disabled: The user is prohibited from logon.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The custom data. The data can be up to 1,024 characters in length.
	//
	// example:
	//
	// md
	UserData map[string]interface{} `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The user ID. The ID can be up to 64 characters in length and cannot contain number signs (#).
	//
	// This parameter is required.
	//
	// example:
	//
	// pdsuserid1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// The username. The username can be up to 128 characters in length.
	//
	// example:
	//
	// pdsusername
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CreateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequest) GoString() string {
	return s.String()
}

func (s *CreateUserRequest) SetAvatar(v string) *CreateUserRequest {
	s.Avatar = &v
	return s
}

func (s *CreateUserRequest) SetDescription(v string) *CreateUserRequest {
	s.Description = &v
	return s
}

func (s *CreateUserRequest) SetEmail(v string) *CreateUserRequest {
	s.Email = &v
	return s
}

func (s *CreateUserRequest) SetGroupInfoList(v []*CreateUserRequestGroupInfoList) *CreateUserRequest {
	s.GroupInfoList = v
	return s
}

func (s *CreateUserRequest) SetNickName(v string) *CreateUserRequest {
	s.NickName = &v
	return s
}

func (s *CreateUserRequest) SetPhone(v string) *CreateUserRequest {
	s.Phone = &v
	return s
}

func (s *CreateUserRequest) SetRole(v string) *CreateUserRequest {
	s.Role = &v
	return s
}

func (s *CreateUserRequest) SetStatus(v string) *CreateUserRequest {
	s.Status = &v
	return s
}

func (s *CreateUserRequest) SetUserData(v map[string]interface{}) *CreateUserRequest {
	s.UserData = v
	return s
}

func (s *CreateUserRequest) SetUserId(v string) *CreateUserRequest {
	s.UserId = &v
	return s
}

func (s *CreateUserRequest) SetUserName(v string) *CreateUserRequest {
	s.UserName = &v
	return s
}

type CreateUserRequestGroupInfoList struct {
	// The group ID.
	//
	// example:
	//
	// g123
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s CreateUserRequestGroupInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreateUserRequestGroupInfoList) GoString() string {
	return s.String()
}

func (s *CreateUserRequestGroupInfoList) SetGroupId(v string) *CreateUserRequestGroupInfoList {
	s.GroupId = &v
	return s
}

type CreateUserResponseBody struct {
	// The URL of the profile picture.
	//
	// example:
	//
	// http://aa.com/1.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The time when the user was created. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1639762579768
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The user who created the user.
	//
	// example:
	//
	// user1
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The ID of the default drive.
	//
	// example:
	//
	// 1
	DefaultDriveId *string `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	// The description of the user.
	//
	// example:
	//
	// vipuser
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The email address.
	//
	// example:
	//
	// a@a.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The nickname of the user.
	//
	// example:
	//
	// 001
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The phone number.
	//
	// example:
	//
	// 13900001111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The role of the user. Valid values:
	//
	// 	- superadmin
	//
	// 	- admin
	//
	// 	- user
	//
	// example:
	//
	// admin
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The state of the user. Valid values:
	//
	// 	- disabled: The user is prohibited from logon.
	//
	// 	- enabled: The user is in a normal state.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The time when the user was modified. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1639762579768
	UpdatedAt *int64 `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The custom data.
	UserData map[string]interface{} `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The user ID.
	//
	// example:
	//
	// dingding_abc001
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	// The username.
	//
	// example:
	//
	// pds
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s CreateUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponseBody) GoString() string {
	return s.String()
}

func (s *CreateUserResponseBody) SetAvatar(v string) *CreateUserResponseBody {
	s.Avatar = &v
	return s
}

func (s *CreateUserResponseBody) SetCreatedAt(v int64) *CreateUserResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *CreateUserResponseBody) SetCreator(v string) *CreateUserResponseBody {
	s.Creator = &v
	return s
}

func (s *CreateUserResponseBody) SetDefaultDriveId(v string) *CreateUserResponseBody {
	s.DefaultDriveId = &v
	return s
}

func (s *CreateUserResponseBody) SetDescription(v string) *CreateUserResponseBody {
	s.Description = &v
	return s
}

func (s *CreateUserResponseBody) SetDomainId(v string) *CreateUserResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateUserResponseBody) SetEmail(v string) *CreateUserResponseBody {
	s.Email = &v
	return s
}

func (s *CreateUserResponseBody) SetNickName(v string) *CreateUserResponseBody {
	s.NickName = &v
	return s
}

func (s *CreateUserResponseBody) SetPhone(v string) *CreateUserResponseBody {
	s.Phone = &v
	return s
}

func (s *CreateUserResponseBody) SetRole(v string) *CreateUserResponseBody {
	s.Role = &v
	return s
}

func (s *CreateUserResponseBody) SetStatus(v string) *CreateUserResponseBody {
	s.Status = &v
	return s
}

func (s *CreateUserResponseBody) SetUpdatedAt(v int64) *CreateUserResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *CreateUserResponseBody) SetUserData(v map[string]interface{}) *CreateUserResponseBody {
	s.UserData = v
	return s
}

func (s *CreateUserResponseBody) SetUserId(v string) *CreateUserResponseBody {
	s.UserId = &v
	return s
}

func (s *CreateUserResponseBody) SetUserName(v string) *CreateUserResponseBody {
	s.UserName = &v
	return s
}

type CreateUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CreateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateUserResponse) GoString() string {
	return s.String()
}

func (s *CreateUserResponse) SetHeaders(v map[string]*string) *CreateUserResponse {
	s.Headers = v
	return s
}

func (s *CreateUserResponse) SetStatusCode(v int32) *CreateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *CreateUserResponse) SetBody(v *CreateUserResponseBody) *CreateUserResponse {
	s.Body = v
	return s
}

type CsiGetFileInfoRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 100
	UrlExpireSec *int32 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
}

func (s CsiGetFileInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s CsiGetFileInfoRequest) GoString() string {
	return s.String()
}

func (s *CsiGetFileInfoRequest) SetDriveId(v string) *CsiGetFileInfoRequest {
	s.DriveId = &v
	return s
}

func (s *CsiGetFileInfoRequest) SetFileId(v string) *CsiGetFileInfoRequest {
	s.FileId = &v
	return s
}

func (s *CsiGetFileInfoRequest) SetUrlExpireSec(v int32) *CsiGetFileInfoRequest {
	s.UrlExpireSec = &v
	return s
}

type CsiGetFileInfoResponseBody struct {
	InvestigationInfo *InvestigationInfo `json:"investigation_info,omitempty" xml:"investigation_info,omitempty"`
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s CsiGetFileInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CsiGetFileInfoResponseBody) GoString() string {
	return s.String()
}

func (s *CsiGetFileInfoResponseBody) SetInvestigationInfo(v *InvestigationInfo) *CsiGetFileInfoResponseBody {
	s.InvestigationInfo = v
	return s
}

func (s *CsiGetFileInfoResponseBody) SetUrl(v string) *CsiGetFileInfoResponseBody {
	s.Url = &v
	return s
}

type CsiGetFileInfoResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *CsiGetFileInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s CsiGetFileInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s CsiGetFileInfoResponse) GoString() string {
	return s.String()
}

func (s *CsiGetFileInfoResponse) SetHeaders(v map[string]*string) *CsiGetFileInfoResponse {
	s.Headers = v
	return s
}

func (s *CsiGetFileInfoResponse) SetStatusCode(v int32) *CsiGetFileInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *CsiGetFileInfoResponse) SetBody(v *CsiGetFileInfoResponseBody) *CsiGetFileInfoResponse {
	s.Body = v
	return s
}

type DeleteDomainRequest struct {
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
}

func (s DeleteDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainRequest) GoString() string {
	return s.String()
}

func (s *DeleteDomainRequest) SetDomainId(v string) *DeleteDomainRequest {
	s.DomainId = &v
	return s
}

type DeleteDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDomainResponse) GoString() string {
	return s.String()
}

func (s *DeleteDomainResponse) SetHeaders(v map[string]*string) *DeleteDomainResponse {
	s.Headers = v
	return s
}

func (s *DeleteDomainResponse) SetStatusCode(v int32) *DeleteDomainResponse {
	s.StatusCode = &v
	return s
}

type DeleteDriveRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s DeleteDriveRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteDriveRequest) GoString() string {
	return s.String()
}

func (s *DeleteDriveRequest) SetDriveId(v string) *DeleteDriveRequest {
	s.DriveId = &v
	return s
}

type DeleteDriveResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteDriveResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteDriveResponse) GoString() string {
	return s.String()
}

func (s *DeleteDriveResponse) SetHeaders(v map[string]*string) *DeleteDriveResponse {
	s.Headers = v
	return s
}

func (s *DeleteDriveResponse) SetStatusCode(v int32) *DeleteDriveResponse {
	s.StatusCode = &v
	return s
}

type DeleteFileRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID or folder ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s DeleteFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileRequest) GoString() string {
	return s.String()
}

func (s *DeleteFileRequest) SetDriveId(v string) *DeleteFileRequest {
	s.DriveId = &v
	return s
}

func (s *DeleteFileRequest) SetFileId(v string) *DeleteFileRequest {
	s.FileId = &v
	return s
}

type DeleteFileResponseBody struct {
	// The ID of the asynchronous task. This parameter is returned only in asynchronous processing scenarios. You can call the [GetAsyncTask](https://help.aliyun.com/document_detail/440456.html) operation to query the information about the asynchronous task based on the task ID.
	//
	// example:
	//
	// 000e89fb-cf8f-11e9-8ab4-b6e980803a3b
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s DeleteFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFileResponseBody) SetAsyncTaskId(v string) *DeleteFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *DeleteFileResponseBody) SetDomainId(v string) *DeleteFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *DeleteFileResponseBody) SetDriveId(v string) *DeleteFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *DeleteFileResponseBody) SetFileId(v string) *DeleteFileResponseBody {
	s.FileId = &v
	return s
}

type DeleteFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFileResponse) GoString() string {
	return s.String()
}

func (s *DeleteFileResponse) SetHeaders(v map[string]*string) *DeleteFileResponse {
	s.Headers = v
	return s
}

func (s *DeleteFileResponse) SetStatusCode(v int32) *DeleteFileResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteFileResponse) SetBody(v *DeleteFileResponseBody) *DeleteFileResponse {
	s.Body = v
	return s
}

type DeleteGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// g123
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s DeleteGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupRequest) GoString() string {
	return s.String()
}

func (s *DeleteGroupRequest) SetGroupId(v string) *DeleteGroupRequest {
	s.GroupId = &v
	return s
}

type DeleteGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteGroupResponse) GoString() string {
	return s.String()
}

func (s *DeleteGroupResponse) SetHeaders(v map[string]*string) *DeleteGroupResponse {
	s.Headers = v
	return s
}

func (s *DeleteGroupResponse) SetStatusCode(v int32) *DeleteGroupResponse {
	s.StatusCode = &v
	return s
}

type DeleteRevisionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40CB7794C929
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s DeleteRevisionRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteRevisionRequest) GoString() string {
	return s.String()
}

func (s *DeleteRevisionRequest) SetDriveId(v string) *DeleteRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *DeleteRevisionRequest) SetFileId(v string) *DeleteRevisionRequest {
	s.FileId = &v
	return s
}

func (s *DeleteRevisionRequest) SetRevisionId(v string) *DeleteRevisionRequest {
	s.RevisionId = &v
	return s
}

type DeleteRevisionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteRevisionResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteRevisionResponse) GoString() string {
	return s.String()
}

func (s *DeleteRevisionResponse) SetHeaders(v map[string]*string) *DeleteRevisionResponse {
	s.Headers = v
	return s
}

func (s *DeleteRevisionResponse) SetStatusCode(v int32) *DeleteRevisionResponse {
	s.StatusCode = &v
	return s
}

type DeleteStoryRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s DeleteStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoryRequest) GoString() string {
	return s.String()
}

func (s *DeleteStoryRequest) SetDriveId(v string) *DeleteStoryRequest {
	s.DriveId = &v
	return s
}

func (s *DeleteStoryRequest) SetStoryId(v string) *DeleteStoryRequest {
	s.StoryId = &v
	return s
}

type DeleteStoryResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s DeleteStoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteStoryResponseBody) SetDriveId(v string) *DeleteStoryResponseBody {
	s.DriveId = &v
	return s
}

type DeleteStoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeleteStoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeleteStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteStoryResponse) GoString() string {
	return s.String()
}

func (s *DeleteStoryResponse) SetHeaders(v map[string]*string) *DeleteStoryResponse {
	s.Headers = v
	return s
}

func (s *DeleteStoryResponse) SetStatusCode(v int32) *DeleteStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *DeleteStoryResponse) SetBody(v *DeleteStoryResponseBody) *DeleteStoryResponse {
	s.Body = v
	return s
}

type DeleteUserRequest struct {
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s DeleteUserRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserRequest) GoString() string {
	return s.String()
}

func (s *DeleteUserRequest) SetUserId(v string) *DeleteUserRequest {
	s.UserId = &v
	return s
}

type DeleteUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DeleteUserResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteUserResponse) GoString() string {
	return s.String()
}

func (s *DeleteUserResponse) SetHeaders(v map[string]*string) *DeleteUserResponse {
	s.Headers = v
	return s
}

func (s *DeleteUserResponse) SetStatusCode(v int32) *DeleteUserResponse {
	s.StatusCode = &v
	return s
}

type DeltaGetLastCursorRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the root file of the synced folder.
	//
	// example:
	//
	// 622fb09598ae66777c7040109a16f49381f6abe1
	SyncRootId *string `json:"sync_root_id,omitempty" xml:"sync_root_id,omitempty"`
}

func (s DeltaGetLastCursorRequest) String() string {
	return tea.Prettify(s)
}

func (s DeltaGetLastCursorRequest) GoString() string {
	return s.String()
}

func (s *DeltaGetLastCursorRequest) SetDriveId(v string) *DeltaGetLastCursorRequest {
	s.DriveId = &v
	return s
}

func (s *DeltaGetLastCursorRequest) SetSyncRootId(v string) *DeltaGetLastCursorRequest {
	s.SyncRootId = &v
	return s
}

type DeltaGetLastCursorResponseBody struct {
	// The latest cursor of incremental information in the specified drive or synced folder.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
}

func (s DeltaGetLastCursorResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeltaGetLastCursorResponseBody) GoString() string {
	return s.String()
}

func (s *DeltaGetLastCursorResponseBody) SetCursor(v string) *DeltaGetLastCursorResponseBody {
	s.Cursor = &v
	return s
}

type DeltaGetLastCursorResponse struct {
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *DeltaGetLastCursorResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s DeltaGetLastCursorResponse) String() string {
	return tea.Prettify(s)
}

func (s DeltaGetLastCursorResponse) GoString() string {
	return s.String()
}

func (s *DeltaGetLastCursorResponse) SetHeaders(v map[string]*string) *DeltaGetLastCursorResponse {
	s.Headers = v
	return s
}

func (s *DeltaGetLastCursorResponse) SetStatusCode(v int32) *DeltaGetLastCursorResponse {
	s.StatusCode = &v
	return s
}

func (s *DeltaGetLastCursorResponse) SetBody(v *DeltaGetLastCursorResponseBody) *DeltaGetLastCursorResponse {
	s.Body = v
	return s
}

type DownloadFileRequest struct {
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The method used to generate the thumbnail of an image. If this parameter is specified, you are redirected to the URL of the generated thumbnail.
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// The method used to generate the thumbnail of a document. If this parameter is specified, you are redirected to the URL of the generated thumbnail.
	//
	// example:
	//
	// image/resize,w_200
	OfficeThumbnailProcess *string `json:"office_thumbnail_process,omitempty" xml:"office_thumbnail_process,omitempty"`
	// The share ID. If you want to manage a file by using a share link, carry the `x-share-token` header for authentication in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify one of `share_id` and `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The method used to generate the thumbnail of a video. If this parameter is specified, you are redirected to the URL of the generated thumbnail.
	//
	// example:
	//
	// video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s DownloadFileRequest) String() string {
	return tea.Prettify(s)
}

func (s DownloadFileRequest) GoString() string {
	return s.String()
}

func (s *DownloadFileRequest) SetDriveId(v string) *DownloadFileRequest {
	s.DriveId = &v
	return s
}

func (s *DownloadFileRequest) SetFileId(v string) *DownloadFileRequest {
	s.FileId = &v
	return s
}

func (s *DownloadFileRequest) SetImageThumbnailProcess(v string) *DownloadFileRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *DownloadFileRequest) SetOfficeThumbnailProcess(v string) *DownloadFileRequest {
	s.OfficeThumbnailProcess = &v
	return s
}

func (s *DownloadFileRequest) SetShareId(v string) *DownloadFileRequest {
	s.ShareId = &v
	return s
}

func (s *DownloadFileRequest) SetVideoThumbnailProcess(v string) *DownloadFileRequest {
	s.VideoThumbnailProcess = &v
	return s
}

type DownloadFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s DownloadFileResponse) String() string {
	return tea.Prettify(s)
}

func (s DownloadFileResponse) GoString() string {
	return s.String()
}

func (s *DownloadFileResponse) SetHeaders(v map[string]*string) *DownloadFileResponse {
	s.Headers = v
	return s
}

func (s *DownloadFileResponse) SetStatusCode(v int32) *DownloadFileResponse {
	s.StatusCode = &v
	return s
}

type FileAddPermissionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the folder. If you want to authorize a user or group to access a team drive, set this parameter to root. If you want to authorize a user or group to access an individual drive, you cannot set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The members that are authorized to access files.
	//
	// This parameter is required.
	MemberList []*FilePermissionMember `json:"member_list,omitempty" xml:"member_list,omitempty" type:"Repeated"`
}

func (s FileAddPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s FileAddPermissionRequest) GoString() string {
	return s.String()
}

func (s *FileAddPermissionRequest) SetDriveId(v string) *FileAddPermissionRequest {
	s.DriveId = &v
	return s
}

func (s *FileAddPermissionRequest) SetFileId(v string) *FileAddPermissionRequest {
	s.FileId = &v
	return s
}

func (s *FileAddPermissionRequest) SetMemberList(v []*FilePermissionMember) *FileAddPermissionRequest {
	s.MemberList = v
	return s
}

type FileAddPermissionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s FileAddPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s FileAddPermissionResponse) GoString() string {
	return s.String()
}

func (s *FileAddPermissionResponse) SetHeaders(v map[string]*string) *FileAddPermissionResponse {
	s.Headers = v
	return s
}

func (s *FileAddPermissionResponse) SetStatusCode(v int32) *FileAddPermissionResponse {
	s.StatusCode = &v
	return s
}

type FileDeleteUserTagsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The tags that you want to remove from a file. You cannot leave this parameter empty. You can specify up to 1,000 tags.
	//
	// This parameter is required.
	KeyList []*string `json:"key_list,omitempty" xml:"key_list,omitempty" type:"Repeated"`
}

func (s FileDeleteUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s FileDeleteUserTagsRequest) GoString() string {
	return s.String()
}

func (s *FileDeleteUserTagsRequest) SetDriveId(v string) *FileDeleteUserTagsRequest {
	s.DriveId = &v
	return s
}

func (s *FileDeleteUserTagsRequest) SetFileId(v string) *FileDeleteUserTagsRequest {
	s.FileId = &v
	return s
}

func (s *FileDeleteUserTagsRequest) SetKeyList(v []*string) *FileDeleteUserTagsRequest {
	s.KeyList = v
	return s
}

type FileDeleteUserTagsResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s FileDeleteUserTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s FileDeleteUserTagsResponse) GoString() string {
	return s.String()
}

func (s *FileDeleteUserTagsResponse) SetHeaders(v map[string]*string) *FileDeleteUserTagsResponse {
	s.Headers = v
	return s
}

func (s *FileDeleteUserTagsResponse) SetStatusCode(v int32) *FileDeleteUserTagsResponse {
	s.StatusCode = &v
	return s
}

type FileListPermissionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43a255edc4463bffa6f5f5d317401
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s FileListPermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s FileListPermissionRequest) GoString() string {
	return s.String()
}

func (s *FileListPermissionRequest) SetDriveId(v string) *FileListPermissionRequest {
	s.DriveId = &v
	return s
}

func (s *FileListPermissionRequest) SetFileId(v string) *FileListPermissionRequest {
	s.FileId = &v
	return s
}

type FileListPermissionResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       []*FilePermissionMember `json:"body,omitempty" xml:"body,omitempty" type:"Repeated"`
}

func (s FileListPermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s FileListPermissionResponse) GoString() string {
	return s.String()
}

func (s *FileListPermissionResponse) SetHeaders(v map[string]*string) *FileListPermissionResponse {
	s.Headers = v
	return s
}

func (s *FileListPermissionResponse) SetStatusCode(v int32) *FileListPermissionResponse {
	s.StatusCode = &v
	return s
}

func (s *FileListPermissionResponse) SetBody(v []*FilePermissionMember) *FileListPermissionResponse {
	s.Body = v
	return s
}

type FilePutUserTagsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The tags to be added to the file. You cannot leave this parameter empty. You can specify up to 1,000 tags. You cannot specify tags that have the same name.
	//
	// This parameter is required.
	UserTags []*FilePutUserTagsRequestUserTags `json:"user_tags,omitempty" xml:"user_tags,omitempty" type:"Repeated"`
}

func (s FilePutUserTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s FilePutUserTagsRequest) GoString() string {
	return s.String()
}

func (s *FilePutUserTagsRequest) SetDriveId(v string) *FilePutUserTagsRequest {
	s.DriveId = &v
	return s
}

func (s *FilePutUserTagsRequest) SetFileId(v string) *FilePutUserTagsRequest {
	s.FileId = &v
	return s
}

func (s *FilePutUserTagsRequest) SetUserTags(v []*FilePutUserTagsRequestUserTags) *FilePutUserTagsRequest {
	s.UserTags = v
	return s
}

type FilePutUserTagsRequestUserTags struct {
	// The name of the tag. The tag name cannot be empty and cannot contain number signs (#).
	//
	// This parameter is required.
	//
	// example:
	//
	// tag
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// The value of the tag. The tag value cannot contain number signs (#).
	//
	// example:
	//
	// value
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s FilePutUserTagsRequestUserTags) String() string {
	return tea.Prettify(s)
}

func (s FilePutUserTagsRequestUserTags) GoString() string {
	return s.String()
}

func (s *FilePutUserTagsRequestUserTags) SetKey(v string) *FilePutUserTagsRequestUserTags {
	s.Key = &v
	return s
}

func (s *FilePutUserTagsRequestUserTags) SetValue(v string) *FilePutUserTagsRequestUserTags {
	s.Value = &v
	return s
}

type FilePutUserTagsResponseBody struct {
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s FilePutUserTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FilePutUserTagsResponseBody) GoString() string {
	return s.String()
}

func (s *FilePutUserTagsResponseBody) SetFileId(v string) *FilePutUserTagsResponseBody {
	s.FileId = &v
	return s
}

type FilePutUserTagsResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *FilePutUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s FilePutUserTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s FilePutUserTagsResponse) GoString() string {
	return s.String()
}

func (s *FilePutUserTagsResponse) SetHeaders(v map[string]*string) *FilePutUserTagsResponse {
	s.Headers = v
	return s
}

func (s *FilePutUserTagsResponse) SetStatusCode(v int32) *FilePutUserTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *FilePutUserTagsResponse) SetBody(v *FilePutUserTagsResponseBody) *FilePutUserTagsResponse {
	s.Body = v
	return s
}

type FileRemovePermissionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The identities with whom the file is shared.
	//
	// This parameter is required.
	MemberList []*FileRemovePermissionRequestMemberList `json:"member_list,omitempty" xml:"member_list,omitempty" type:"Repeated"`
}

func (s FileRemovePermissionRequest) String() string {
	return tea.Prettify(s)
}

func (s FileRemovePermissionRequest) GoString() string {
	return s.String()
}

func (s *FileRemovePermissionRequest) SetDriveId(v string) *FileRemovePermissionRequest {
	s.DriveId = &v
	return s
}

func (s *FileRemovePermissionRequest) SetFileId(v string) *FileRemovePermissionRequest {
	s.FileId = &v
	return s
}

func (s *FileRemovePermissionRequest) SetMemberList(v []*FileRemovePermissionRequestMemberList) *FileRemovePermissionRequest {
	s.MemberList = v
	return s
}

type FileRemovePermissionRequestMemberList struct {
	// The identity to whom the permissions are granted, which is a user or a group.
	//
	// This parameter is required.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	// The role ID. You can grant permissions by assigning roles to identities, or you can customize the permissions. To grant permissions by assigning roles to identities, specify role_id. role_id and action_list are mutually exclusive. If both parameters are specified, role_id has a higher priority.
	//
	// Valid values:
	//
	// SystemFileOwner: collaborator.
	//
	// SystemFileDownloader: downloader.
	//
	// SystemFileEditor: editor.
	//
	// SystemFileEditorWithoutDelete: editor without permissions to delete the file.
	//
	// SystemFileEditorWithoutShareLink: editor without permissions to share the file.
	//
	// SystemFileMetaViewer: viewer of lists.
	//
	// SystemFileUploader: uploader. SystemFileUploaderAndDownloader: uploader and downloader.
	//
	// SystemFileDownloaderWithShareLink: downloader and sharer.
	//
	// SystemFileUploaderAndDownloaderWithShareLink: uploader, downloader, and sharer.
	//
	// SystemFileUploaderAndViewer: viewer and uploader.
	//
	// SystemFileUploaderWithShareLink: uploader and sharer.
	//
	// SystemFileViewer: viewer.
	//
	// This parameter is required.
	//
	// example:
	//
	// SystemFileDownloader
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s FileRemovePermissionRequestMemberList) String() string {
	return tea.Prettify(s)
}

func (s FileRemovePermissionRequestMemberList) GoString() string {
	return s.String()
}

func (s *FileRemovePermissionRequestMemberList) SetIdentity(v *Identity) *FileRemovePermissionRequestMemberList {
	s.Identity = v
	return s
}

func (s *FileRemovePermissionRequestMemberList) SetRoleId(v string) *FileRemovePermissionRequestMemberList {
	s.RoleId = &v
	return s
}

type FileRemovePermissionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s FileRemovePermissionResponse) String() string {
	return tea.Prettify(s)
}

func (s FileRemovePermissionResponse) GoString() string {
	return s.String()
}

func (s *FileRemovePermissionResponse) SetHeaders(v map[string]*string) *FileRemovePermissionResponse {
	s.Headers = v
	return s
}

func (s *FileRemovePermissionResponse) SetStatusCode(v int32) *FileRemovePermissionResponse {
	s.StatusCode = &v
	return s
}

type GetAsyncTaskRequest struct {
	// The ID of the asynchronous task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 000e89fb-cf8f-11e9-8ab4-b6e980803a3b
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
}

func (s GetAsyncTaskRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskRequest) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskRequest) SetAsyncTaskId(v string) *GetAsyncTaskRequest {
	s.AsyncTaskId = &v
	return s
}

type GetAsyncTaskResponseBody struct {
	// The ID of the asynchronous task.
	//
	// example:
	//
	// 000e89fb-cf8f-11e9-8ab4-b6e980803a3b
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The custom category of the task.
	//
	// example:
	//
	// album
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The total amount of work that is done in the asynchronous task, such as the number of files that are packaged for package download on the server.
	//
	// example:
	//
	// 100
	ConsumedProcess *int64 `json:"consumed_process,omitempty" xml:"consumed_process,omitempty"`
	// The time when the task was created. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. Example: 2019-03-28T13:03:29.298Z.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	CreatedAt *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// <warning>This parameter is no longer used. We recommend that you use error_code instead.</warning>
	//
	// The error code returned if the asynchronous task failed.
	//
	// example:
	//
	// InternalError
	ErrCode *int64 `json:"err_code,omitempty" xml:"err_code,omitempty"`
	// The error code returned if the asynchronous task failed.
	//
	// example:
	//
	// InternalError
	ErrorCode *string `json:"error_code,omitempty" xml:"error_code,omitempty"`
	// The error message returned if the asynchronous task failed.
	//
	// example:
	//
	// The request has been failed due to some unknown error. Please try again later.
	ErrorMessage  *string `json:"error_message,omitempty" xml:"error_message,omitempty"`
	FailedProcess *int64  `json:"failed_process,omitempty" xml:"failed_process,omitempty"`
	// The time when the task was complete. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. Example: 2019-03-28T13:03:29.298Z.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	FinishedAt *string `json:"finished_at,omitempty" xml:"finished_at,omitempty"`
	// <warning>This parameter is no longer used. We recommend that you use error_message instead.</warning>
	//
	// The error message returned if the asynchronous task failed.
	//
	// example:
	//
	// The request has been failed due to some unknown error. Please try again later.
	Message        *string `json:"message,omitempty" xml:"message,omitempty"`
	SkippedProcess *int64  `json:"skipped_process,omitempty" xml:"skipped_process,omitempty"`
	// The time when the task was started. The time follows the ISO 8601 standard in the yyyy-MM-ddTHH:mm:ssZ format. Example: 2019-03-28T13:03:29.298Z.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	StartedAt *string `json:"started_at,omitempty" xml:"started_at,omitempty"`
	// The state of the task. Valid values:
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- PartialSucceed
	//
	// 	- Succeed
	//
	// example:
	//
	// Succeed
	State *string `json:"state,omitempty" xml:"state,omitempty"`
	// <warning>This parameter is no longer used. We recommend that you use state instead.</warning>
	//
	// The state of the task. Valid values:
	//
	// 	- Failed
	//
	// 	- Running
	//
	// 	- PartialSucceed
	//
	// 	- Succeed
	//
	// example:
	//
	// Succeed
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total amount of work to be done in the asynchronous task, such as the number of files to be packaged for package download on the server.
	//
	// example:
	//
	// 1000
	TotalProcess *int64 `json:"total_process,omitempty" xml:"total_process,omitempty"`
	// The extracted files.
	UncompressFileList []*UncompressedFileInfo `json:"uncompress_file_list,omitempty" xml:"uncompress_file_list,omitempty" type:"Repeated"`
	// The download URL of the data generated by the asynchronous task, such as the download URL of the packaged files generated by the task of package download on the server.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d5b986facbec311ef844c25954f96821497b383%2F5d5b986f955410dd991646bb87c6b4e899eff525?Expires=xxx&OSSAccessKeyId=xxx&Signature=xxx
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetAsyncTaskResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskResponseBody) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskResponseBody) SetAsyncTaskId(v string) *GetAsyncTaskResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetCategory(v string) *GetAsyncTaskResponseBody {
	s.Category = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetConsumedProcess(v int64) *GetAsyncTaskResponseBody {
	s.ConsumedProcess = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetCreatedAt(v string) *GetAsyncTaskResponseBody {
	s.CreatedAt = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetErrCode(v int64) *GetAsyncTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetErrorCode(v string) *GetAsyncTaskResponseBody {
	s.ErrorCode = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetErrorMessage(v string) *GetAsyncTaskResponseBody {
	s.ErrorMessage = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetFailedProcess(v int64) *GetAsyncTaskResponseBody {
	s.FailedProcess = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetFinishedAt(v string) *GetAsyncTaskResponseBody {
	s.FinishedAt = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetMessage(v string) *GetAsyncTaskResponseBody {
	s.Message = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetSkippedProcess(v int64) *GetAsyncTaskResponseBody {
	s.SkippedProcess = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetStartedAt(v string) *GetAsyncTaskResponseBody {
	s.StartedAt = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetState(v string) *GetAsyncTaskResponseBody {
	s.State = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetStatus(v string) *GetAsyncTaskResponseBody {
	s.Status = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetTotalProcess(v int64) *GetAsyncTaskResponseBody {
	s.TotalProcess = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetUncompressFileList(v []*UncompressedFileInfo) *GetAsyncTaskResponseBody {
	s.UncompressFileList = v
	return s
}

func (s *GetAsyncTaskResponseBody) SetUrl(v string) *GetAsyncTaskResponseBody {
	s.Url = &v
	return s
}

type GetAsyncTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetAsyncTaskResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAsyncTaskResponse) GoString() string {
	return s.String()
}

func (s *GetAsyncTaskResponse) SetHeaders(v map[string]*string) *GetAsyncTaskResponse {
	s.Headers = v
	return s
}

func (s *GetAsyncTaskResponse) SetStatusCode(v int32) *GetAsyncTaskResponse {
	s.StatusCode = &v
	return s
}

func (s *GetAsyncTaskResponse) SetBody(v *GetAsyncTaskResponseBody) *GetAsyncTaskResponse {
	s.Body = v
	return s
}

type GetDefaultDriveRequest struct {
	// The user ID. If you use an AccessKey pair for authentication, you must specify this parameter. If you use an access token for authentication, this parameter is optional. By default, the user ID associated with the access token is used.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s GetDefaultDriveRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultDriveRequest) GoString() string {
	return s.String()
}

func (s *GetDefaultDriveRequest) SetUserId(v string) *GetDefaultDriveRequest {
	s.UserId = &v
	return s
}

type GetDefaultDriveResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Drive             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDefaultDriveResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDefaultDriveResponse) GoString() string {
	return s.String()
}

func (s *GetDefaultDriveResponse) SetHeaders(v map[string]*string) *GetDefaultDriveResponse {
	s.Headers = v
	return s
}

func (s *GetDefaultDriveResponse) SetStatusCode(v int32) *GetDefaultDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDefaultDriveResponse) SetBody(v *Drive) *GetDefaultDriveResponse {
	s.Body = v
	return s
}

type GetDomainRequest struct {
	// The ID of the domain.
	//
	// This parameter is required.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	Fields   *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// Specifies whether to return the used quota of the domain. Default value: false. If the quota of the domain is greater than 0 and you set this parameter to true, the used quota of the domain is returned.
	//
	// example:
	//
	// true
	GetQuotaUsed *bool `json:"get_quota_used,omitempty" xml:"get_quota_used,omitempty"`
}

func (s GetDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDomainRequest) GoString() string {
	return s.String()
}

func (s *GetDomainRequest) SetDomainId(v string) *GetDomainRequest {
	s.DomainId = &v
	return s
}

func (s *GetDomainRequest) SetFields(v string) *GetDomainRequest {
	s.Fields = &v
	return s
}

func (s *GetDomainRequest) SetGetQuotaUsed(v bool) *GetDomainRequest {
	s.GetQuotaUsed = &v
	return s
}

type GetDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Domain            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainResponse) GoString() string {
	return s.String()
}

func (s *GetDomainResponse) SetHeaders(v map[string]*string) *GetDomainResponse {
	s.Headers = v
	return s
}

func (s *GetDomainResponse) SetStatusCode(v int32) *GetDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainResponse) SetBody(v *Domain) *GetDomainResponse {
	s.Body = v
	return s
}

type GetDomainQuotaResponseBody struct {
	SizeQuota      *int64 `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	SizeUsed       *int64 `json:"size_used,omitempty" xml:"size_used,omitempty"`
	UserCountQuota *int64 `json:"user_count_quota,omitempty" xml:"user_count_quota,omitempty"`
	UserCountUsed  *int64 `json:"user_count_used,omitempty" xml:"user_count_used,omitempty"`
}

func (s GetDomainQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDomainQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetDomainQuotaResponseBody) SetSizeQuota(v int64) *GetDomainQuotaResponseBody {
	s.SizeQuota = &v
	return s
}

func (s *GetDomainQuotaResponseBody) SetSizeUsed(v int64) *GetDomainQuotaResponseBody {
	s.SizeUsed = &v
	return s
}

func (s *GetDomainQuotaResponseBody) SetUserCountQuota(v int64) *GetDomainQuotaResponseBody {
	s.UserCountQuota = &v
	return s
}

func (s *GetDomainQuotaResponseBody) SetUserCountUsed(v int64) *GetDomainQuotaResponseBody {
	s.UserCountUsed = &v
	return s
}

type GetDomainQuotaResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDomainQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDomainQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDomainQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetDomainQuotaResponse) SetHeaders(v map[string]*string) *GetDomainQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetDomainQuotaResponse) SetStatusCode(v int32) *GetDomainQuotaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDomainQuotaResponse) SetBody(v *GetDomainQuotaResponseBody) *GetDomainQuotaResponse {
	s.Body = v
	return s
}

type GetDownloadUrlRequest struct {
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The validity period of the download URL. Maximum value: 115200. Default value: 900. Unit: seconds.
	//
	// example:
	//
	// 100
	ExpireSec *int32 `json:"expire_sec,omitempty" xml:"expire_sec,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The name of the file. The name can be up to 1,024 characters in length.
	//
	// example:
	//
	// 1.txt
	FileName *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
	// example:
	//
	// video/mp4
	ResponseContentType *string `json:"response_content_type,omitempty" xml:"response_content_type,omitempty"`
	// The share ID. If you want to manage a file by using a sharing link, carry the `x-share-token` header in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify at least either `share_id` or `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s GetDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlRequest) SetDriveId(v string) *GetDownloadUrlRequest {
	s.DriveId = &v
	return s
}

func (s *GetDownloadUrlRequest) SetExpireSec(v int32) *GetDownloadUrlRequest {
	s.ExpireSec = &v
	return s
}

func (s *GetDownloadUrlRequest) SetFileId(v string) *GetDownloadUrlRequest {
	s.FileId = &v
	return s
}

func (s *GetDownloadUrlRequest) SetFileName(v string) *GetDownloadUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetDownloadUrlRequest) SetResponseContentType(v string) *GetDownloadUrlRequest {
	s.ResponseContentType = &v
	return s
}

func (s *GetDownloadUrlRequest) SetShareId(v string) *GetDownloadUrlRequest {
	s.ShareId = &v
	return s
}

type GetDownloadUrlResponseBody struct {
	// The download URL of a file that is downloaded by using Alibaba Cloud CDN.
	//
	// example:
	//
	// https://data-cdn.aliyunpds.com/hz22%2F5d79219b0aa9a7c995a94a96993ba3205cd91c5a%2F5d79219bf3261a5d38744da0834ed489b677a27a?Expires=xxxOSSAccessKeyId=xxx&Signature=xxx&response-content-disposition=attachment%3Bfilename%3DtBiZAoJPC2c8b13450eda4292b7f5f8010618e078.txt
	CdnUrl *string `json:"cdn_url,omitempty" xml:"cdn_url,omitempty"`
	// The hash value of the file content.
	//
	// example:
	//
	// EA4942AA8761213890A5C386F88E6464D2C31CA1
	ContentHash *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	// The name of the algorithm that is used to calculate the hash value of the file content.
	//
	// example:
	//
	// sha1
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	// The hash value calculated by using 64-bit cyclic redundancy check (CRC-64).
	//
	// example:
	//
	// 5498595269368962671
	Crc64Hash *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	// The time when the download URL expires.
	//
	// example:
	//
	// 2022-01-02T15:04:05.999Z07:00
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// The download URL of a file that is downloaded over a virtual private cloud (VPC).
	//
	// example:
	//
	// https://data-vpc.aliyunpds.com/hz22%2F5d79219b0aa9a7c995a94a96993ba3205cd91c5a%2F5d79219bf3261a5d38744da0834ed489b677a27a?Expires=xxxOSSAccessKeyId=xxx&Signature=xxx&response-content-disposition=attachment%3Bfilename%3DtBiZAoJPC2c8b13450eda4292b7f5f8010618e078.txt
	InternalUrl *string `json:"internal_url,omitempty" xml:"internal_url,omitempty"`
	// The size of the file. Unit: bytes.
	//
	// example:
	//
	// 10
	Size *int64 `json:"size,omitempty" xml:"size,omitempty"`
	// The download URL of a file that is downloaded over the Internet.
	//
	// example:
	//
	// https://data.aliyunpds.com/hz22%2F5d79219b0aa9a7c995a94a96993ba3205cd91c5a%2F5d79219bf3261a5d38744da0834ed489b677a27a?Expires=xxxOSSAccessKeyId=xxx&Signature=xxx&response-content-disposition=attachment%3Bfilename%3DtBiZAoJPC2c8b13450eda4292b7f5f8010618e078.txt
	Url *string `json:"url,omitempty" xml:"url,omitempty"`
}

func (s GetDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlResponseBody) SetCdnUrl(v string) *GetDownloadUrlResponseBody {
	s.CdnUrl = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetContentHash(v string) *GetDownloadUrlResponseBody {
	s.ContentHash = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetContentHashName(v string) *GetDownloadUrlResponseBody {
	s.ContentHashName = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetCrc64Hash(v string) *GetDownloadUrlResponseBody {
	s.Crc64Hash = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetExpiration(v string) *GetDownloadUrlResponseBody {
	s.Expiration = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetInternalUrl(v string) *GetDownloadUrlResponseBody {
	s.InternalUrl = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetSize(v int64) *GetDownloadUrlResponseBody {
	s.Size = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetUrl(v string) *GetDownloadUrlResponseBody {
	s.Url = &v
	return s
}

type GetDownloadUrlResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDownloadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlResponse) SetHeaders(v map[string]*string) *GetDownloadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetDownloadUrlResponse) SetStatusCode(v int32) *GetDownloadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDownloadUrlResponse) SetBody(v *GetDownloadUrlResponseBody) *GetDownloadUrlResponse {
	s.Body = v
	return s
}

type GetDriveRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s GetDriveRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDriveRequest) GoString() string {
	return s.String()
}

func (s *GetDriveRequest) SetDriveId(v string) *GetDriveRequest {
	s.DriveId = &v
	return s
}

type GetDriveResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Drive             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetDriveResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDriveResponse) GoString() string {
	return s.String()
}

func (s *GetDriveResponse) SetHeaders(v map[string]*string) *GetDriveResponse {
	s.Headers = v
	return s
}

func (s *GetDriveResponse) SetStatusCode(v int32) *GetDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *GetDriveResponse) SetBody(v *Drive) *GetDriveResponse {
	s.Body = v
	return s
}

type GetFileRequest struct {
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The fields to return.
	//
	// 1.  If this parameter is set to \\*, all fields of the file except the fields that must be specified are returned.
	//
	// 2.  If only specific fields are required, you can specify the following fields: url, thumbnail, exif, cropping_suggestion, characteristic_hash, video_metadata, and video_preview_metadata. If multiple fields are required, separate them with commas (,). Example: url,thumbnail.
	//
	// 3.  The investigation_info field is returned only if you specify this field.
	//
	// By default, all fields except the fields that must be specified are returned.
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The share ID. If you want to manage a file by using a share link, carry the `x-share-token` header for authentication in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify one of `share_id` and `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// 缩略图配置，可一次性返回最多5个缩略图，map的key可以自定义，返回时按key返回对应的缩略图链接
	ThumbnailProcesses map[string]*ImageProcess `json:"thumbnail_processes,omitempty" xml:"thumbnail_processes,omitempty"`
	// The time when the file expires. Unit: seconds. Valid values: 10 to 14400.
	//
	// example:
	//
	// 100
	UrlExpireSec *int32 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
}

func (s GetFileRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFileRequest) GoString() string {
	return s.String()
}

func (s *GetFileRequest) SetDriveId(v string) *GetFileRequest {
	s.DriveId = &v
	return s
}

func (s *GetFileRequest) SetFields(v string) *GetFileRequest {
	s.Fields = &v
	return s
}

func (s *GetFileRequest) SetFileId(v string) *GetFileRequest {
	s.FileId = &v
	return s
}

func (s *GetFileRequest) SetShareId(v string) *GetFileRequest {
	s.ShareId = &v
	return s
}

func (s *GetFileRequest) SetThumbnailProcesses(v map[string]*ImageProcess) *GetFileRequest {
	s.ThumbnailProcesses = v
	return s
}

func (s *GetFileRequest) SetUrlExpireSec(v int32) *GetFileRequest {
	s.UrlExpireSec = &v
	return s
}

type GetFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *File              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetFileResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFileResponse) GoString() string {
	return s.String()
}

func (s *GetFileResponse) SetHeaders(v map[string]*string) *GetFileResponse {
	s.Headers = v
	return s
}

func (s *GetFileResponse) SetStatusCode(v int32) *GetFileResponse {
	s.StatusCode = &v
	return s
}

func (s *GetFileResponse) SetBody(v *File) *GetFileResponse {
	s.Body = v
	return s
}

type GetGroupRequest struct {
	// The group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e43ec8427dd45f19431b7504649a1b1
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s GetGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s GetGroupRequest) GoString() string {
	return s.String()
}

func (s *GetGroupRequest) SetGroupId(v string) *GetGroupRequest {
	s.GroupId = &v
	return s
}

type GetGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Group             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s GetGroupResponse) GoString() string {
	return s.String()
}

func (s *GetGroupResponse) SetHeaders(v map[string]*string) *GetGroupResponse {
	s.Headers = v
	return s
}

func (s *GetGroupResponse) SetStatusCode(v int32) *GetGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *GetGroupResponse) SetBody(v *Group) *GetGroupResponse {
	s.Body = v
	return s
}

type GetIdentityToBenefitPkgMappingRequest struct {
	// The unique identifier of the benefit package.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40cb7794c9294
	BenefitPkgId *string `json:"benefit_pkg_id,omitempty" xml:"benefit_pkg_id,omitempty"`
	// The unique identifier of the entity.
	//
	// If you want to manage the benefits of a user, set this parameter to a user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// user123
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// The type of the entity. If you want to manage the benefits of a user, set this parameter to user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
}

func (s GetIdentityToBenefitPkgMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s GetIdentityToBenefitPkgMappingRequest) GoString() string {
	return s.String()
}

func (s *GetIdentityToBenefitPkgMappingRequest) SetBenefitPkgId(v string) *GetIdentityToBenefitPkgMappingRequest {
	s.BenefitPkgId = &v
	return s
}

func (s *GetIdentityToBenefitPkgMappingRequest) SetIdentityId(v string) *GetIdentityToBenefitPkgMappingRequest {
	s.IdentityId = &v
	return s
}

func (s *GetIdentityToBenefitPkgMappingRequest) SetIdentityType(v string) *GetIdentityToBenefitPkgMappingRequest {
	s.IdentityType = &v
	return s
}

type GetIdentityToBenefitPkgMappingResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *IdentityToBenefitPkgMapping `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetIdentityToBenefitPkgMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s GetIdentityToBenefitPkgMappingResponse) GoString() string {
	return s.String()
}

func (s *GetIdentityToBenefitPkgMappingResponse) SetHeaders(v map[string]*string) *GetIdentityToBenefitPkgMappingResponse {
	s.Headers = v
	return s
}

func (s *GetIdentityToBenefitPkgMappingResponse) SetStatusCode(v int32) *GetIdentityToBenefitPkgMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *GetIdentityToBenefitPkgMappingResponse) SetBody(v *IdentityToBenefitPkgMapping) *GetIdentityToBenefitPkgMappingResponse {
	s.Body = v
	return s
}

type GetLinkInfoRequest struct {
	// The additional information about the unique identifier of the account. For example, if type is set to mobile, set the value of extra to a country code.
	//
	// example:
	//
	// 1
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The unique identifier of the account, such as a mobile number.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130***
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	// The account type. Valid values:
	//
	// 	- mobile: a mobile number.
	//
	// 	- email: an email address.
	//
	// 	- ding: a DingTalk account.
	//
	// 	- ram: an Alibaba Cloud Resource Access Management (RAM) user.
	//
	// 	- wechat: a WeCom account.
	//
	// 	- ldap: a Lightweight Directory Access Protocol (LDAP) account.
	//
	// 	- custom: a custom account.
	//
	// This parameter is required.
	//
	// example:
	//
	// mobile
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s GetLinkInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLinkInfoRequest) GoString() string {
	return s.String()
}

func (s *GetLinkInfoRequest) SetExtra(v string) *GetLinkInfoRequest {
	s.Extra = &v
	return s
}

func (s *GetLinkInfoRequest) SetIdentity(v string) *GetLinkInfoRequest {
	s.Identity = &v
	return s
}

func (s *GetLinkInfoRequest) SetType(v string) *GetLinkInfoRequest {
	s.Type = &v
	return s
}

type GetLinkInfoResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *AccountLinkInfo   `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLinkInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLinkInfoResponse) GoString() string {
	return s.String()
}

func (s *GetLinkInfoResponse) SetHeaders(v map[string]*string) *GetLinkInfoResponse {
	s.Headers = v
	return s
}

func (s *GetLinkInfoResponse) SetStatusCode(v int32) *GetLinkInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLinkInfoResponse) SetBody(v *AccountLinkInfo) *GetLinkInfoResponse {
	s.Body = v
	return s
}

type GetLinkInfoByUserIdRequest struct {
	// The user ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s GetLinkInfoByUserIdRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLinkInfoByUserIdRequest) GoString() string {
	return s.String()
}

func (s *GetLinkInfoByUserIdRequest) SetUserId(v string) *GetLinkInfoByUserIdRequest {
	s.UserId = &v
	return s
}

type GetLinkInfoByUserIdResponseBody struct {
	// The information about the users.
	Items []*AccountLinkInfo `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s GetLinkInfoByUserIdResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLinkInfoByUserIdResponseBody) GoString() string {
	return s.String()
}

func (s *GetLinkInfoByUserIdResponseBody) SetItems(v []*AccountLinkInfo) *GetLinkInfoByUserIdResponseBody {
	s.Items = v
	return s
}

type GetLinkInfoByUserIdResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetLinkInfoByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetLinkInfoByUserIdResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLinkInfoByUserIdResponse) GoString() string {
	return s.String()
}

func (s *GetLinkInfoByUserIdResponse) SetHeaders(v map[string]*string) *GetLinkInfoByUserIdResponse {
	s.Headers = v
	return s
}

func (s *GetLinkInfoByUserIdResponse) SetStatusCode(v int32) *GetLinkInfoByUserIdResponse {
	s.StatusCode = &v
	return s
}

func (s *GetLinkInfoByUserIdResponse) SetBody(v *GetLinkInfoByUserIdResponseBody) *GetLinkInfoByUserIdResponse {
	s.Body = v
	return s
}

type GetRevisionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Specifies the returned fields.
	//
	// By default, this parameter is left empty. If you set this parameter to \\*, all fields are returned. If you leave this parameter empty, the creator of the file is not returned.
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40CB7794C929
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	// The validity period of the file download or preview. Valid values: 10 to 86400.
	//
	// Default value: 900. Unit: seconds.
	//
	// example:
	//
	// 900
	UrlExpireSec *int64 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
}

func (s GetRevisionRequest) String() string {
	return tea.Prettify(s)
}

func (s GetRevisionRequest) GoString() string {
	return s.String()
}

func (s *GetRevisionRequest) SetDriveId(v string) *GetRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *GetRevisionRequest) SetFields(v string) *GetRevisionRequest {
	s.Fields = &v
	return s
}

func (s *GetRevisionRequest) SetFileId(v string) *GetRevisionRequest {
	s.FileId = &v
	return s
}

func (s *GetRevisionRequest) SetRevisionId(v string) *GetRevisionRequest {
	s.RevisionId = &v
	return s
}

func (s *GetRevisionRequest) SetUrlExpireSec(v int64) *GetRevisionRequest {
	s.UrlExpireSec = &v
	return s
}

type GetRevisionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Revision          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetRevisionResponse) String() string {
	return tea.Prettify(s)
}

func (s GetRevisionResponse) GoString() string {
	return s.String()
}

func (s *GetRevisionResponse) SetHeaders(v map[string]*string) *GetRevisionResponse {
	s.Headers = v
	return s
}

func (s *GetRevisionResponse) SetStatusCode(v int32) *GetRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *GetRevisionResponse) SetBody(v *Revision) *GetRevisionResponse {
	s.Body = v
	return s
}

type GetShareLinkRequest struct {
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s GetShareLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s GetShareLinkRequest) GoString() string {
	return s.String()
}

func (s *GetShareLinkRequest) SetShareId(v string) *GetShareLinkRequest {
	s.ShareId = &v
	return s
}

type GetShareLinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareLink         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetShareLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShareLinkResponse) GoString() string {
	return s.String()
}

func (s *GetShareLinkResponse) SetHeaders(v map[string]*string) *GetShareLinkResponse {
	s.Headers = v
	return s
}

func (s *GetShareLinkResponse) SetStatusCode(v int32) *GetShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShareLinkResponse) SetBody(v *ShareLink) *GetShareLinkResponse {
	s.Body = v
	return s
}

type GetShareLinkByAnonymousRequest struct {
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s GetShareLinkByAnonymousRequest) String() string {
	return tea.Prettify(s)
}

func (s GetShareLinkByAnonymousRequest) GoString() string {
	return s.String()
}

func (s *GetShareLinkByAnonymousRequest) SetShareId(v string) *GetShareLinkByAnonymousRequest {
	s.ShareId = &v
	return s
}

type GetShareLinkByAnonymousResponseBody struct {
	// The number of times that the shared files are visited.
	//
	// example:
	//
	// 30
	AccessCount *int64 `json:"access_count,omitempty" xml:"access_count,omitempty"`
	// The profile picture of the user who created the share link.
	//
	// example:
	//
	// https://aliyunpds.com/a.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The ID of the user who created the share link.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	CreatorId *string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	// The name of the user who created the share link. The value is masked.
	//
	// example:
	//
	// AB***CD
	CreatorName *string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	// The mobile number of the user who created the share link. The value is masked.
	//
	// example:
	//
	// 136****00
	CreatorPhone *string `json:"creator_phone,omitempty" xml:"creator_phone,omitempty"`
	// Indicates whether the downloads of the shared files are prohibited.
	//
	// example:
	//
	// false
	DisableDownload *bool `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	// Indicates whether the previews of the shared files are prohibited.
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	// Indicates whether the saves of the shared files are prohibited.
	//
	// example:
	//
	// false
	DisableSave *bool `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	// The number of times that the shared files are downloaded.
	//
	// example:
	//
	// 50
	DownloadCount *int64 `json:"download_count,omitempty" xml:"download_count,omitempty"`
	// The maximum number of times that the shared files can be downloaded.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	// The time when the share link expires.
	//
	// example:
	//
	// 2020-08-20T06:51:27.292Z
	Expiration *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	// The number of times that the shared files are previewed.
	//
	// example:
	//
	// 80
	PreviewCount *int64 `json:"preview_count,omitempty" xml:"preview_count,omitempty"`
	// The maximum number of times that the shared files can be previewed.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	// The number of times that the shared files are reported.
	//
	// example:
	//
	// 0
	ReportCount *int64 `json:"report_count,omitempty" xml:"report_count,omitempty"`
	// The number of times that the shared files are saved.
	//
	// example:
	//
	// 2
	SaveCount *int64 `json:"save_count,omitempty" xml:"save_count,omitempty"`
	// The maximum number of times that the shared files can be saved and downloaded.
	//
	// example:
	//
	// 200
	SaveDownloadLimit *int64 `json:"save_download_limit,omitempty" xml:"save_download_limit,omitempty"`
	// The maximum number of times that the shared files can be saved.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	// The name of the share link.
	ShareName *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	// The time when the share link was last modified.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	UpdatedAt *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	// The number of times that the videos are previewed in the shared files.
	//
	// example:
	//
	// 5
	VideoPreviewCount *int64 `json:"video_preview_count,omitempty" xml:"video_preview_count,omitempty"`
}

func (s GetShareLinkByAnonymousResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShareLinkByAnonymousResponseBody) GoString() string {
	return s.String()
}

func (s *GetShareLinkByAnonymousResponseBody) SetAccessCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.AccessCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetAvatar(v string) *GetShareLinkByAnonymousResponseBody {
	s.Avatar = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetCreatorId(v string) *GetShareLinkByAnonymousResponseBody {
	s.CreatorId = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetCreatorName(v string) *GetShareLinkByAnonymousResponseBody {
	s.CreatorName = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetCreatorPhone(v string) *GetShareLinkByAnonymousResponseBody {
	s.CreatorPhone = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDisableDownload(v bool) *GetShareLinkByAnonymousResponseBody {
	s.DisableDownload = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDisablePreview(v bool) *GetShareLinkByAnonymousResponseBody {
	s.DisablePreview = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDisableSave(v bool) *GetShareLinkByAnonymousResponseBody {
	s.DisableSave = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDownloadCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.DownloadCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetDownloadLimit(v int64) *GetShareLinkByAnonymousResponseBody {
	s.DownloadLimit = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetExpiration(v string) *GetShareLinkByAnonymousResponseBody {
	s.Expiration = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetPreviewCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.PreviewCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetPreviewLimit(v int64) *GetShareLinkByAnonymousResponseBody {
	s.PreviewLimit = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetReportCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.ReportCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetSaveCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.SaveCount = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetSaveDownloadLimit(v int64) *GetShareLinkByAnonymousResponseBody {
	s.SaveDownloadLimit = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetSaveLimit(v int64) *GetShareLinkByAnonymousResponseBody {
	s.SaveLimit = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetShareName(v string) *GetShareLinkByAnonymousResponseBody {
	s.ShareName = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetUpdatedAt(v string) *GetShareLinkByAnonymousResponseBody {
	s.UpdatedAt = &v
	return s
}

func (s *GetShareLinkByAnonymousResponseBody) SetVideoPreviewCount(v int64) *GetShareLinkByAnonymousResponseBody {
	s.VideoPreviewCount = &v
	return s
}

type GetShareLinkByAnonymousResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetShareLinkByAnonymousResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetShareLinkByAnonymousResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShareLinkByAnonymousResponse) GoString() string {
	return s.String()
}

func (s *GetShareLinkByAnonymousResponse) SetHeaders(v map[string]*string) *GetShareLinkByAnonymousResponse {
	s.Headers = v
	return s
}

func (s *GetShareLinkByAnonymousResponse) SetStatusCode(v int32) *GetShareLinkByAnonymousResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShareLinkByAnonymousResponse) SetBody(v *GetShareLinkByAnonymousResponseBody) *GetShareLinkByAnonymousResponse {
	s.Body = v
	return s
}

type GetShareLinkTokenRequest struct {
	// The validity period of the token. Valid values: (0,7200]. Default value: 7200. Unit: seconds.
	//
	// example:
	//
	// 7200
	ExpireSec *int32 `json:"expire_sec,omitempty" xml:"expire_sec,omitempty"`
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The access code.
	//
	// example:
	//
	// abcF123x
	SharePwd *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
}

func (s GetShareLinkTokenRequest) String() string {
	return tea.Prettify(s)
}

func (s GetShareLinkTokenRequest) GoString() string {
	return s.String()
}

func (s *GetShareLinkTokenRequest) SetExpireSec(v int32) *GetShareLinkTokenRequest {
	s.ExpireSec = &v
	return s
}

func (s *GetShareLinkTokenRequest) SetShareId(v string) *GetShareLinkTokenRequest {
	s.ShareId = &v
	return s
}

func (s *GetShareLinkTokenRequest) SetSharePwd(v string) *GetShareLinkTokenRequest {
	s.SharePwd = &v
	return s
}

type GetShareLinkTokenResponseBody struct {
	// The validity period of the token. Unit: seconds. For example, a value of 7200 indicates two hours.
	//
	// example:
	//
	// 7200
	ExpiresIn *int64 `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	// The JSON Web Token (JWT).
	//
	// example:
	//
	// eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VySWQiOiJjOWI3YTVhYTA0ZDE0YWUzODY3ZmRjODg2ZmEwMWRhNCIsImN1c3RvbUpzb24iOiJ7XCJjbGllbnRJZFwiOlwiMjVkelgzdmJZcWt0Vnh5WFwiLFwiZG9tYWluSWRcIjpcImJqMjlcIixcInNjb3BlXCI6W1wiRFJJVkUuQUxMXCIsXCJTSEFSRS5BTExcIixcIkZJTEUuQUxMXCIsXCJVU0VSLkFMTFwiLFwiVklFVy5BTExcIixcIlNUT1JBR0UuQUxMXCIsXCJTVE9SQUdFRklMRS5MSVNUXCIsXCJCQVRDSFwiLFwiT0FVVEguQUxMXCIsXCJJTUFHRS5BTExcIixcIklOVklURS5BTExcIixcIkFDQ09VTlQuQUxMXCJdLFwicm9sZVwiOlwidXNlclwiLFwicmVmXCI6XCJodHRwczovL3d3dy5hbGl5dW5kcml2ZS5jb20vXCIsXCJkZXZpY2VfaWRcIjpcImIyODIwNWU1YzU5NzRjY2JiODI3MDNiNjhkYjhjNDUxXCJ9IiwiZXhwIjoxNjQ4NjE0NDkzLCJpYXQiOjE2NDg2MDcyMzN9.d3HVLvv_LFw2QhPrhvjH_kICWQJX9sKKt7NjQEqI_xE2JO_b7D8rPsFTZz93PLvZ7MhCmudTjGImUpd-ehFnI4Go-1S7BGaKaHFILvP-sWy18Wpikowjxx9mSbzBM_cO6D1LI-kyYhXKWHgVdADfVIPniTDA7-ffhUpi7cAebEs
	ShareToken *string `json:"share_token,omitempty" xml:"share_token,omitempty"`
}

func (s GetShareLinkTokenResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetShareLinkTokenResponseBody) GoString() string {
	return s.String()
}

func (s *GetShareLinkTokenResponseBody) SetExpiresIn(v int64) *GetShareLinkTokenResponseBody {
	s.ExpiresIn = &v
	return s
}

func (s *GetShareLinkTokenResponseBody) SetShareToken(v string) *GetShareLinkTokenResponseBody {
	s.ShareToken = &v
	return s
}

type GetShareLinkTokenResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetShareLinkTokenResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetShareLinkTokenResponse) String() string {
	return tea.Prettify(s)
}

func (s GetShareLinkTokenResponse) GoString() string {
	return s.String()
}

func (s *GetShareLinkTokenResponse) SetHeaders(v map[string]*string) *GetShareLinkTokenResponse {
	s.Headers = v
	return s
}

func (s *GetShareLinkTokenResponse) SetStatusCode(v int32) *GetShareLinkTokenResponse {
	s.StatusCode = &v
	return s
}

func (s *GetShareLinkTokenResponse) SetBody(v *GetShareLinkTokenResponseBody) *GetShareLinkTokenResponse {
	s.Body = v
	return s
}

type GetStoryRequest struct {
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	CoverImageThumbnailProcess *string `json:"cover_image_thumbnail_process,omitempty" xml:"cover_image_thumbnail_process,omitempty"`
	// Deprecated
	//
	// example:
	//
	// video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast,ar_auto
	CoverVideoThumbnailProcess *string `json:"cover_video_thumbnail_process,omitempty" xml:"cover_video_thumbnail_process,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	ImageUrlProcess *string `json:"image_url_process,omitempty" xml:"image_url_process,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 900
	UrlExpireSec *int64 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
	// Deprecated
	//
	// example:
	//
	// video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast,ar_auto
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s GetStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetStoryRequest) GoString() string {
	return s.String()
}

func (s *GetStoryRequest) SetCoverImageThumbnailProcess(v string) *GetStoryRequest {
	s.CoverImageThumbnailProcess = &v
	return s
}

func (s *GetStoryRequest) SetCoverVideoThumbnailProcess(v string) *GetStoryRequest {
	s.CoverVideoThumbnailProcess = &v
	return s
}

func (s *GetStoryRequest) SetDriveId(v string) *GetStoryRequest {
	s.DriveId = &v
	return s
}

func (s *GetStoryRequest) SetImageThumbnailProcess(v string) *GetStoryRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *GetStoryRequest) SetImageUrlProcess(v string) *GetStoryRequest {
	s.ImageUrlProcess = &v
	return s
}

func (s *GetStoryRequest) SetStoryId(v string) *GetStoryRequest {
	s.StoryId = &v
	return s
}

func (s *GetStoryRequest) SetUrlExpireSec(v int64) *GetStoryRequest {
	s.UrlExpireSec = &v
	return s
}

func (s *GetStoryRequest) SetVideoThumbnailProcess(v string) *GetStoryRequest {
	s.VideoThumbnailProcess = &v
	return s
}

type GetStoryResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Story             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetStoryResponse) GoString() string {
	return s.String()
}

func (s *GetStoryResponse) SetHeaders(v map[string]*string) *GetStoryResponse {
	s.Headers = v
	return s
}

func (s *GetStoryResponse) SetStatusCode(v int32) *GetStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *GetStoryResponse) SetBody(v *Story) *GetStoryResponse {
	s.Body = v
	return s
}

type GetTaskStatusRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the task.
	//
	// This parameter is required.
	//
	// example:
	//
	// i:SimilarImageClustering-b67d53e7-2fe8-460f-9b95-1e93636923eb
	TaskId *string `json:"task_id,omitempty" xml:"task_id,omitempty"`
}

func (s GetTaskStatusRequest) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusRequest) GoString() string {
	return s.String()
}

func (s *GetTaskStatusRequest) SetDriveId(v string) *GetTaskStatusRequest {
	s.DriveId = &v
	return s
}

func (s *GetTaskStatusRequest) SetTaskId(v string) *GetTaskStatusRequest {
	s.TaskId = &v
	return s
}

type GetTaskStatusResponseBody struct {
	// The state of the task.
	//
	// Valid values:
	//
	// 	- running
	//
	//     <!-- -->
	//
	//     : The task is
	//
	//     <!-- -->
	//
	//     running
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- failed
	//
	//     <!-- -->
	//
	//     : The task
	//
	//     <!-- -->
	//
	//     fails
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- succeeded
	//
	//     <!-- -->
	//
	//     : The task is
	//
	//     <!-- -->
	//
	//     successful
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// running
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s GetTaskStatusResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponseBody) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponseBody) SetStatus(v string) *GetTaskStatusResponseBody {
	s.Status = &v
	return s
}

type GetTaskStatusResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetTaskStatusResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetTaskStatusResponse) String() string {
	return tea.Prettify(s)
}

func (s GetTaskStatusResponse) GoString() string {
	return s.String()
}

func (s *GetTaskStatusResponse) SetHeaders(v map[string]*string) *GetTaskStatusResponse {
	s.Headers = v
	return s
}

func (s *GetTaskStatusResponse) SetStatusCode(v int32) *GetTaskStatusResponse {
	s.StatusCode = &v
	return s
}

func (s *GetTaskStatusResponse) SetBody(v *GetTaskStatusResponseBody) *GetTaskStatusResponse {
	s.Body = v
	return s
}

type GetUploadUrlRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 5d5b846942cf94fa72324c14a4bda34e81da635d
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The information about the file parts.
	//
	// This parameter is required.
	PartInfoList []*GetUploadUrlRequestPartInfoList `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	// The share ID.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The ID of the upload task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10166D06127B413BA1EC8ABB1144D111
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s GetUploadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequest) SetDriveId(v string) *GetUploadUrlRequest {
	s.DriveId = &v
	return s
}

func (s *GetUploadUrlRequest) SetFileId(v string) *GetUploadUrlRequest {
	s.FileId = &v
	return s
}

func (s *GetUploadUrlRequest) SetPartInfoList(v []*GetUploadUrlRequestPartInfoList) *GetUploadUrlRequest {
	s.PartInfoList = v
	return s
}

func (s *GetUploadUrlRequest) SetShareId(v string) *GetUploadUrlRequest {
	s.ShareId = &v
	return s
}

func (s *GetUploadUrlRequest) SetUploadId(v string) *GetUploadUrlRequest {
	s.UploadId = &v
	return s
}

type GetUploadUrlRequestPartInfoList struct {
	ContentMd5  *string `json:"content_md5,omitempty" xml:"content_md5,omitempty"`
	ContentType *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	// The SHA-1 hash value of the file content before the file part. This parameter takes effect only if the parallel upload feature is enabled.
	ParallelSha1Ctx   *GetUploadUrlRequestPartInfoListParallelSha1Ctx   `json:"parallel_sha1_ctx,omitempty" xml:"parallel_sha1_ctx,omitempty" type:"Struct"`
	ParallelSha256Ctx *GetUploadUrlRequestPartInfoListParallelSha256Ctx `json:"parallel_sha256_ctx,omitempty" xml:"parallel_sha256_ctx,omitempty" type:"Struct"`
	// The serial number of a part.
	//
	// example:
	//
	// 1
	PartNumber *int32 `json:"part_number,omitempty" xml:"part_number,omitempty"`
}

func (s GetUploadUrlRequestPartInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlRequestPartInfoList) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequestPartInfoList) SetContentMd5(v string) *GetUploadUrlRequestPartInfoList {
	s.ContentMd5 = &v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) SetContentType(v string) *GetUploadUrlRequestPartInfoList {
	s.ContentType = &v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) SetParallelSha1Ctx(v *GetUploadUrlRequestPartInfoListParallelSha1Ctx) *GetUploadUrlRequestPartInfoList {
	s.ParallelSha1Ctx = v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) SetParallelSha256Ctx(v *GetUploadUrlRequestPartInfoListParallelSha256Ctx) *GetUploadUrlRequestPartInfoList {
	s.ParallelSha256Ctx = v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) SetPartNumber(v int32) *GetUploadUrlRequestPartInfoList {
	s.PartNumber = &v
	return s
}

type GetUploadUrlRequestPartInfoListParallelSha1Ctx struct {
	// The first to fifth 32-bit variables of the SHA-1 hash value of the file content before the file part. This parameter takes effect only if the parallel upload feature is enabled.
	H []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	// The size of the file part. Unit: bytes. The value must be a multiple of 64. This parameter takes effect only if the parallel upload feature is enabled.
	//
	// example:
	//
	// 10240
	PartOffset *int64 `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s GetUploadUrlRequestPartInfoListParallelSha1Ctx) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlRequestPartInfoListParallelSha1Ctx) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequestPartInfoListParallelSha1Ctx) SetH(v []*int64) *GetUploadUrlRequestPartInfoListParallelSha1Ctx {
	s.H = v
	return s
}

func (s *GetUploadUrlRequestPartInfoListParallelSha1Ctx) SetPartOffset(v int64) *GetUploadUrlRequestPartInfoListParallelSha1Ctx {
	s.PartOffset = &v
	return s
}

type GetUploadUrlRequestPartInfoListParallelSha256Ctx struct {
	H          []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	PartOffset *int64   `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
}

func (s GetUploadUrlRequestPartInfoListParallelSha256Ctx) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlRequestPartInfoListParallelSha256Ctx) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequestPartInfoListParallelSha256Ctx) SetH(v []*int64) *GetUploadUrlRequestPartInfoListParallelSha256Ctx {
	s.H = v
	return s
}

func (s *GetUploadUrlRequestPartInfoListParallelSha256Ctx) SetPartOffset(v int64) *GetUploadUrlRequestPartInfoListParallelSha256Ctx {
	s.PartOffset = &v
	return s
}

type GetUploadUrlResponseBody struct {
	// The time when the upload task was created.
	//
	// example:
	//
	// 2019-09-11T16:34:36.977Z
	CreateAt *string `json:"create_at,omitempty" xml:"create_at,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 5d5b846942cf94fa72324c14a4bda34e81da635d
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The information about the file parts.
	PartInfoList []*UploadPartInfo `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	// The ID of the upload task.
	//
	// example:
	//
	// 10166D06127B413BA1EC8ABB1144D111
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s GetUploadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponseBody) SetCreateAt(v string) *GetUploadUrlResponseBody {
	s.CreateAt = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetDomainId(v string) *GetUploadUrlResponseBody {
	s.DomainId = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetDriveId(v string) *GetUploadUrlResponseBody {
	s.DriveId = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetFileId(v string) *GetUploadUrlResponseBody {
	s.FileId = &v
	return s
}

func (s *GetUploadUrlResponseBody) SetPartInfoList(v []*UploadPartInfo) *GetUploadUrlResponseBody {
	s.PartInfoList = v
	return s
}

func (s *GetUploadUrlResponseBody) SetUploadId(v string) *GetUploadUrlResponseBody {
	s.UploadId = &v
	return s
}

type GetUploadUrlResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUploadUrlResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlResponse) GoString() string {
	return s.String()
}

func (s *GetUploadUrlResponse) SetHeaders(v map[string]*string) *GetUploadUrlResponse {
	s.Headers = v
	return s
}

func (s *GetUploadUrlResponse) SetStatusCode(v int32) *GetUploadUrlResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUploadUrlResponse) SetBody(v *GetUploadUrlResponseBody) *GetUploadUrlResponse {
	s.Body = v
	return s
}

type GetUserRequest struct {
	// The user ID. If you use an AccessKey pair to access Drive and Photo Service, you must specify this parameter. If you use an access token to access Drive and Photo Service, you do not need to specify this parameter, and Drive and Photo Service automatically finds the user ID contained in the access token.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s GetUserRequest) String() string {
	return tea.Prettify(s)
}

func (s GetUserRequest) GoString() string {
	return s.String()
}

func (s *GetUserRequest) SetUserId(v string) *GetUserRequest {
	s.UserId = &v
	return s
}

type GetUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *User              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetUserResponse) String() string {
	return tea.Prettify(s)
}

func (s GetUserResponse) GoString() string {
	return s.String()
}

func (s *GetUserResponse) SetHeaders(v map[string]*string) *GetUserResponse {
	s.Headers = v
	return s
}

func (s *GetUserResponse) SetStatusCode(v int32) *GetUserResponse {
	s.StatusCode = &v
	return s
}

func (s *GetUserResponse) SetBody(v *User) *GetUserResponse {
	s.Body = v
	return s
}

type GetVideoPreviewPlayInfoRequest struct {
	// The preview type. You must enable the corresponding video transcoding feature. Valid values:
	//
	// 	- live_transcoding: previews a live stream while transcoding is in progress.
	//
	// 	- quick_video: previews a video while transcoding is in progress.
	//
	// 	- offline_audio: previews a piece of audio after the audio is transcoded offline.
	//
	// 	- offline_video: previews a video after the video is transcoded offline.
	//
	// This parameter is required.
	//
	// example:
	//
	// live_transcoding
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId       *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	GetMasterUrl *bool   `json:"get_master_url,omitempty" xml:"get_master_url,omitempty"`
	// Specifies whether not to query the playback URL. If you set this parameter to true, only transcoding metadata is returned. The video is not transcoded in the TS format, and the playback URL is not returned. If you set this parameter to false, the playback URL is returned. If the video has not been transcoded by using the template specified by template_id, the transcoding process is triggered. You are charged for the value-added service fees generated for transcoding.
	//
	// example:
	//
	// true
	GetWithoutUrl *bool `json:"get_without_url,omitempty" xml:"get_without_url,omitempty"`
	ReTranscode   *bool `json:"re_transcode,omitempty" xml:"re_transcode,omitempty"`
	// The share ID. If you want to manage a file by using a sharing link, carry the `x-share-token` header in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify at least either `share_id` or `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The ID of the definition template. If you leave this parameter empty, all definition templates are available.
	//
	// example:
	//
	// 264_480p
	TemplateId *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
	// The validity period of the video preview. Unit: seconds. Default value: 900. Maximum value: 14400.
	//
	// example:
	//
	// 3600
	UrlExpireSec *int64 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
}

func (s GetVideoPreviewPlayInfoRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPreviewPlayInfoRequest) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayInfoRequest) SetCategory(v string) *GetVideoPreviewPlayInfoRequest {
	s.Category = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetDriveId(v string) *GetVideoPreviewPlayInfoRequest {
	s.DriveId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetFileId(v string) *GetVideoPreviewPlayInfoRequest {
	s.FileId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetGetMasterUrl(v bool) *GetVideoPreviewPlayInfoRequest {
	s.GetMasterUrl = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetGetWithoutUrl(v bool) *GetVideoPreviewPlayInfoRequest {
	s.GetWithoutUrl = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetReTranscode(v bool) *GetVideoPreviewPlayInfoRequest {
	s.ReTranscode = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetShareId(v string) *GetVideoPreviewPlayInfoRequest {
	s.ShareId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetTemplateId(v string) *GetVideoPreviewPlayInfoRequest {
	s.TemplateId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoRequest) SetUrlExpireSec(v int64) *GetVideoPreviewPlayInfoRequest {
	s.UrlExpireSec = &v
	return s
}

type GetVideoPreviewPlayInfoResponseBody struct {
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// example:
	//
	// fileid1
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The share ID.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The information about video playback.
	VideoPreviewPlayInfo *VideoPreviewPlayInfo `json:"video_preview_play_info,omitempty" xml:"video_preview_play_info,omitempty"`
}

func (s GetVideoPreviewPlayInfoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPreviewPlayInfoResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetDomainId(v string) *GetVideoPreviewPlayInfoResponseBody {
	s.DomainId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetDriveId(v string) *GetVideoPreviewPlayInfoResponseBody {
	s.DriveId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetFileId(v string) *GetVideoPreviewPlayInfoResponseBody {
	s.FileId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetShareId(v string) *GetVideoPreviewPlayInfoResponseBody {
	s.ShareId = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponseBody) SetVideoPreviewPlayInfo(v *VideoPreviewPlayInfo) *GetVideoPreviewPlayInfoResponseBody {
	s.VideoPreviewPlayInfo = v
	return s
}

type GetVideoPreviewPlayInfoResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoPreviewPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoPreviewPlayInfoResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPreviewPlayInfoResponse) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayInfoResponse) SetHeaders(v map[string]*string) *GetVideoPreviewPlayInfoResponse {
	s.Headers = v
	return s
}

func (s *GetVideoPreviewPlayInfoResponse) SetStatusCode(v int32) *GetVideoPreviewPlayInfoResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoPreviewPlayInfoResponse) SetBody(v *GetVideoPreviewPlayInfoResponseBody) *GetVideoPreviewPlayInfoResponse {
	s.Body = v
	return s
}

type GetVideoPreviewPlayMetaRequest struct {
	// The preview type. You must enable the corresponding video transcoding feature. Valid values:
	//
	// 	- live_transcoding: previews a live stream while transcoding is in progress.
	//
	// 	- quick_video: previews a video while transcoding is in progress.
	//
	// 	- offline_audio: previews a piece of audio after the audio is transcoded offline.
	//
	// 	- offline_video: previews a video after the video is transcoded offline.
	//
	// This parameter is required.
	//
	// example:
	//
	// live_transcoding
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The share ID. If you want to manage a file by using a sharing link, carry the `x-share-token` header in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify at least either `share_id` or `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
}

func (s GetVideoPreviewPlayMetaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPreviewPlayMetaRequest) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayMetaRequest) SetCategory(v string) *GetVideoPreviewPlayMetaRequest {
	s.Category = &v
	return s
}

func (s *GetVideoPreviewPlayMetaRequest) SetDriveId(v string) *GetVideoPreviewPlayMetaRequest {
	s.DriveId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaRequest) SetFileId(v string) *GetVideoPreviewPlayMetaRequest {
	s.FileId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaRequest) SetShareId(v string) *GetVideoPreviewPlayMetaRequest {
	s.ShareId = &v
	return s
}

type GetVideoPreviewPlayMetaResponseBody struct {
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// example:
	//
	// fileid1
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The share ID.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The preview metadata of the video.
	VideoPreviewPlayMeta *VideoPreviewPlayMeta `json:"video_preview_play_meta,omitempty" xml:"video_preview_play_meta,omitempty"`
}

func (s GetVideoPreviewPlayMetaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPreviewPlayMetaResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetDomainId(v string) *GetVideoPreviewPlayMetaResponseBody {
	s.DomainId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetDriveId(v string) *GetVideoPreviewPlayMetaResponseBody {
	s.DriveId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetFileId(v string) *GetVideoPreviewPlayMetaResponseBody {
	s.FileId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetShareId(v string) *GetVideoPreviewPlayMetaResponseBody {
	s.ShareId = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponseBody) SetVideoPreviewPlayMeta(v *VideoPreviewPlayMeta) *GetVideoPreviewPlayMetaResponseBody {
	s.VideoPreviewPlayMeta = v
	return s
}

type GetVideoPreviewPlayMetaResponse struct {
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *GetVideoPreviewPlayMetaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s GetVideoPreviewPlayMetaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoPreviewPlayMetaResponse) GoString() string {
	return s.String()
}

func (s *GetVideoPreviewPlayMetaResponse) SetHeaders(v map[string]*string) *GetVideoPreviewPlayMetaResponse {
	s.Headers = v
	return s
}

func (s *GetVideoPreviewPlayMetaResponse) SetStatusCode(v int32) *GetVideoPreviewPlayMetaResponse {
	s.StatusCode = &v
	return s
}

func (s *GetVideoPreviewPlayMetaResponse) SetBody(v *GetVideoPreviewPlayMetaResponseBody) *GetVideoPreviewPlayMetaResponse {
	s.Body = v
	return s
}

type GroupUpdateNameRequest struct {
	// This parameter is required.
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	Name    *string `json:"name,omitempty" xml:"name,omitempty"`
}

func (s GroupUpdateNameRequest) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateNameRequest) GoString() string {
	return s.String()
}

func (s *GroupUpdateNameRequest) SetGroupId(v string) *GroupUpdateNameRequest {
	s.GroupId = &v
	return s
}

func (s *GroupUpdateNameRequest) SetName(v string) *GroupUpdateNameRequest {
	s.Name = &v
	return s
}

type GroupUpdateNameResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s GroupUpdateNameResponse) String() string {
	return tea.Prettify(s)
}

func (s GroupUpdateNameResponse) GoString() string {
	return s.String()
}

func (s *GroupUpdateNameResponse) SetHeaders(v map[string]*string) *GroupUpdateNameResponse {
	s.Headers = v
	return s
}

func (s *GroupUpdateNameResponse) SetStatusCode(v int32) *GroupUpdateNameResponse {
	s.StatusCode = &v
	return s
}

type ImportUserRequest struct {
	// The display name of the authentication type.
	//
	// example:
	//
	// 10000
	AuthenticationDisplayName *string `json:"authentication_display_name,omitempty" xml:"authentication_display_name,omitempty"`
	// The authentication type. Valid values:
	//
	// 	- mobile: mobile number.
	//
	// 	- email: email address.
	//
	// 	- ding: DingTalk account.
	//
	// 	- ram: Alibaba Cloud Resource Access Management (RAM) user.
	//
	// 	- wechat: WeCom account.
	//
	// 	- ldap: Lightweight Directory Access Protocol (LDAP) account.
	//
	// 	- custom: custom account.
	//
	// This parameter is required.
	//
	// example:
	//
	// mobile
	AuthenticationType *string `json:"authentication_type,omitempty" xml:"authentication_type,omitempty"`
	// Specifies whether to automatically create a drive.
	//
	// example:
	//
	// false
	AutoCreateDrive *bool `json:"auto_create_drive,omitempty" xml:"auto_create_drive,omitempty"`
	// The size of the drive. The value cannot be smaller than -1. A value of -1 specifies that the size is unlimited.
	//
	// example:
	//
	// 10240
	DriveTotalSize *int64 `json:"drive_total_size,omitempty" xml:"drive_total_size,omitempty"`
	// The additional information.
	//
	// If authentication_type is set to mobile, set this parameter to a country code. If you do not specify this parameter, 86 is used by default.
	//
	// example:
	//
	// 1
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The unique identifier.
	//
	// This parameter is required.
	//
	// example:
	//
	// 130****
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	// The nickname of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// pdsuer
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The ID of the group to which the user is added.
	//
	// example:
	//
	// g12
	ParentGroupId *string `json:"parent_group_id,omitempty" xml:"parent_group_id,omitempty"`
}

func (s ImportUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ImportUserRequest) GoString() string {
	return s.String()
}

func (s *ImportUserRequest) SetAuthenticationDisplayName(v string) *ImportUserRequest {
	s.AuthenticationDisplayName = &v
	return s
}

func (s *ImportUserRequest) SetAuthenticationType(v string) *ImportUserRequest {
	s.AuthenticationType = &v
	return s
}

func (s *ImportUserRequest) SetAutoCreateDrive(v bool) *ImportUserRequest {
	s.AutoCreateDrive = &v
	return s
}

func (s *ImportUserRequest) SetDriveTotalSize(v int64) *ImportUserRequest {
	s.DriveTotalSize = &v
	return s
}

func (s *ImportUserRequest) SetExtra(v string) *ImportUserRequest {
	s.Extra = &v
	return s
}

func (s *ImportUserRequest) SetIdentity(v string) *ImportUserRequest {
	s.Identity = &v
	return s
}

func (s *ImportUserRequest) SetNickName(v string) *ImportUserRequest {
	s.NickName = &v
	return s
}

func (s *ImportUserRequest) SetParentGroupId(v string) *ImportUserRequest {
	s.ParentGroupId = &v
	return s
}

type ImportUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *User              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ImportUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ImportUserResponse) GoString() string {
	return s.String()
}

func (s *ImportUserResponse) SetHeaders(v map[string]*string) *ImportUserResponse {
	s.Headers = v
	return s
}

func (s *ImportUserResponse) SetStatusCode(v int32) *ImportUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ImportUserResponse) SetBody(v *User) *ImportUserResponse {
	s.Body = v
	return s
}

type InvestigateFileRequest struct {
	// This parameter is required.
	DriveFileIds []*InvestigateFileRequestDriveFileIds `json:"drive_file_ids,omitempty" xml:"drive_file_ids,omitempty" type:"Repeated"`
}

func (s InvestigateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s InvestigateFileRequest) GoString() string {
	return s.String()
}

func (s *InvestigateFileRequest) SetDriveFileIds(v []*InvestigateFileRequestDriveFileIds) *InvestigateFileRequest {
	s.DriveFileIds = v
	return s
}

type InvestigateFileRequestDriveFileIds struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s InvestigateFileRequestDriveFileIds) String() string {
	return tea.Prettify(s)
}

func (s InvestigateFileRequestDriveFileIds) GoString() string {
	return s.String()
}

func (s *InvestigateFileRequestDriveFileIds) SetDriveId(v string) *InvestigateFileRequestDriveFileIds {
	s.DriveId = &v
	return s
}

func (s *InvestigateFileRequestDriveFileIds) SetFileId(v string) *InvestigateFileRequestDriveFileIds {
	s.FileId = &v
	return s
}

type InvestigateFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s InvestigateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s InvestigateFileResponse) GoString() string {
	return s.String()
}

func (s *InvestigateFileResponse) SetHeaders(v map[string]*string) *InvestigateFileResponse {
	s.Headers = v
	return s
}

func (s *InvestigateFileResponse) SetStatusCode(v int32) *InvestigateFileResponse {
	s.StatusCode = &v
	return s
}

type LinkAccountRequest struct {
	// The additional information about the unique identifier of the account. For example, if type is set to mobile, set the value of extra to a country code. For example, a value of 86 specifies a mobile number in the Chinese mainland. If you do not specify this parameter, 86 is used by default.
	//
	// example:
	//
	// 86
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// The unique identifier of the account, such as a mobile number.
	//
	// This parameter is required.
	//
	// example:
	//
	// eyy***
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	// The account type. Valid values:
	//
	// 	- mobile: a mobile number.
	//
	// 	- email: an email address.
	//
	// 	- ding: a DingTalk account.
	//
	// 	- ram: an Alibaba Cloud Resource Access Management (RAM) user.
	//
	// 	- wechat: a WeCom account.
	//
	// 	- ldap: a Lightweight Directory Access Protocol (LDAP) account.
	//
	// 	- custom: a custom account.
	//
	// This parameter is required.
	//
	// example:
	//
	// ding
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// The ID of the user with which you want to associate an account.
	//
	// This parameter is required.
	//
	// example:
	//
	// xxx
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s LinkAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s LinkAccountRequest) GoString() string {
	return s.String()
}

func (s *LinkAccountRequest) SetExtra(v string) *LinkAccountRequest {
	s.Extra = &v
	return s
}

func (s *LinkAccountRequest) SetIdentity(v string) *LinkAccountRequest {
	s.Identity = &v
	return s
}

func (s *LinkAccountRequest) SetType(v string) *LinkAccountRequest {
	s.Type = &v
	return s
}

func (s *LinkAccountRequest) SetUserId(v string) *LinkAccountRequest {
	s.UserId = &v
	return s
}

type LinkAccountResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Token             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s LinkAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s LinkAccountResponse) GoString() string {
	return s.String()
}

func (s *LinkAccountResponse) SetHeaders(v map[string]*string) *LinkAccountResponse {
	s.Headers = v
	return s
}

func (s *LinkAccountResponse) SetStatusCode(v int32) *LinkAccountResponse {
	s.StatusCode = &v
	return s
}

func (s *LinkAccountResponse) SetBody(v *Token) *LinkAccountResponse {
	s.Body = v
	return s
}

type ListAddressGroupsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The method that is used to generate a thumbnail of an image.
	//
	// example:
	//
	// image/resize,w_200
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The method that is used to generate a thumbnail of a video.
	//
	// example:
	//
	// video_thumbnail_process
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s ListAddressGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAddressGroupsRequest) GoString() string {
	return s.String()
}

func (s *ListAddressGroupsRequest) SetDriveId(v string) *ListAddressGroupsRequest {
	s.DriveId = &v
	return s
}

func (s *ListAddressGroupsRequest) SetImageThumbnailProcess(v string) *ListAddressGroupsRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *ListAddressGroupsRequest) SetLimit(v int32) *ListAddressGroupsRequest {
	s.Limit = &v
	return s
}

func (s *ListAddressGroupsRequest) SetMarker(v string) *ListAddressGroupsRequest {
	s.Marker = &v
	return s
}

func (s *ListAddressGroupsRequest) SetVideoThumbnailProcess(v string) *ListAddressGroupsRequest {
	s.VideoThumbnailProcess = &v
	return s
}

type ListAddressGroupsResponseBody struct {
	// The information about the location-based groups.
	Items []*AddressGroup `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListAddressGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAddressGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddressGroupsResponseBody) SetItems(v []*AddressGroup) *ListAddressGroupsResponseBody {
	s.Items = v
	return s
}

func (s *ListAddressGroupsResponseBody) SetNextMarker(v string) *ListAddressGroupsResponseBody {
	s.NextMarker = &v
	return s
}

type ListAddressGroupsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAddressGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAddressGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAddressGroupsResponse) GoString() string {
	return s.String()
}

func (s *ListAddressGroupsResponse) SetHeaders(v map[string]*string) *ListAddressGroupsResponse {
	s.Headers = v
	return s
}

func (s *ListAddressGroupsResponse) SetStatusCode(v int32) *ListAddressGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAddressGroupsResponse) SetBody(v *ListAddressGroupsResponseBody) *ListAddressGroupsResponse {
	s.Body = v
	return s
}

type ListAssignmentRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The ID of the managed resource, such as a group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 105***b82
	ManageResourceId *string `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	// The type of the managed resource. Set the value to RT_Group, which specifies that the administrators of a group are queried.
	//
	// This parameter is required.
	//
	// example:
	//
	// RT_Group
	ManageResourceType *string `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListAssignmentRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAssignmentRequest) GoString() string {
	return s.String()
}

func (s *ListAssignmentRequest) SetLimit(v int32) *ListAssignmentRequest {
	s.Limit = &v
	return s
}

func (s *ListAssignmentRequest) SetManageResourceId(v string) *ListAssignmentRequest {
	s.ManageResourceId = &v
	return s
}

func (s *ListAssignmentRequest) SetManageResourceType(v string) *ListAssignmentRequest {
	s.ManageResourceType = &v
	return s
}

func (s *ListAssignmentRequest) SetMarker(v string) *ListAssignmentRequest {
	s.Marker = &v
	return s
}

type ListAssignmentResponseBody struct {
	// The assigned roles.
	AssignmentList []*ListAssignmentResponseBodyAssignmentList `json:"assignment_list,omitempty" xml:"assignment_list,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListAssignmentResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAssignmentResponseBody) GoString() string {
	return s.String()
}

func (s *ListAssignmentResponseBody) SetAssignmentList(v []*ListAssignmentResponseBodyAssignmentList) *ListAssignmentResponseBody {
	s.AssignmentList = v
	return s
}

func (s *ListAssignmentResponseBody) SetNextMarker(v string) *ListAssignmentResponseBody {
	s.NextMarker = &v
	return s
}

type ListAssignmentResponseBodyAssignmentList struct {
	// The time when the role was assigned. This value is a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// example:
	//
	// 1622682267564
	CreatedAt *int64 `json:"created_at,omitempty" xml:"created_at,omitempty"`
	// The ID of the user who assigned the role.
	//
	// example:
	//
	// 216***c83
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// hz1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The identity to whom the role is assigned, which is a user or a group.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	// The ID of the managed resource, such as a group ID.
	//
	// example:
	//
	// 105***b82
	ManageResourceId *string `json:"manage_resource_id,omitempty" xml:"manage_resource_id,omitempty"`
	// The type of the managed resource. For example, a value of RT_Group indicates group.
	//
	// example:
	//
	// RT_Group
	ManageResourceType *string `json:"manage_resource_type,omitempty" xml:"manage_resource_type,omitempty"`
	// The ID of the role assigned to the identity.
	//
	// example:
	//
	// SystemGroupAdmin
	RoleId *string `json:"role_id,omitempty" xml:"role_id,omitempty"`
}

func (s ListAssignmentResponseBodyAssignmentList) String() string {
	return tea.Prettify(s)
}

func (s ListAssignmentResponseBodyAssignmentList) GoString() string {
	return s.String()
}

func (s *ListAssignmentResponseBodyAssignmentList) SetCreatedAt(v int64) *ListAssignmentResponseBodyAssignmentList {
	s.CreatedAt = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetCreator(v string) *ListAssignmentResponseBodyAssignmentList {
	s.Creator = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetDomainId(v string) *ListAssignmentResponseBodyAssignmentList {
	s.DomainId = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetIdentity(v *Identity) *ListAssignmentResponseBodyAssignmentList {
	s.Identity = v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetManageResourceId(v string) *ListAssignmentResponseBodyAssignmentList {
	s.ManageResourceId = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetManageResourceType(v string) *ListAssignmentResponseBodyAssignmentList {
	s.ManageResourceType = &v
	return s
}

func (s *ListAssignmentResponseBodyAssignmentList) SetRoleId(v string) *ListAssignmentResponseBodyAssignmentList {
	s.RoleId = &v
	return s
}

type ListAssignmentResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListAssignmentResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListAssignmentResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAssignmentResponse) GoString() string {
	return s.String()
}

func (s *ListAssignmentResponse) SetHeaders(v map[string]*string) *ListAssignmentResponse {
	s.Headers = v
	return s
}

func (s *ListAssignmentResponse) SetStatusCode(v int32) *ListAssignmentResponse {
	s.StatusCode = &v
	return s
}

func (s *ListAssignmentResponse) SetBody(v *ListAssignmentResponseBody) *ListAssignmentResponse {
	s.Body = v
	return s
}

type ListDeltaRequest struct {
	// The cursor of the incremental information.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The maximum number of results to return. Valid values: 0 to 100. Default value: 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The ID of the root file of the synced folder.
	//
	// example:
	//
	// 622fb09598ae66777c7040109a16f49381f6abe1
	SyncRootId *string `json:"sync_root_id,omitempty" xml:"sync_root_id,omitempty"`
}

func (s ListDeltaRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDeltaRequest) GoString() string {
	return s.String()
}

func (s *ListDeltaRequest) SetCursor(v string) *ListDeltaRequest {
	s.Cursor = &v
	return s
}

func (s *ListDeltaRequest) SetDriveId(v string) *ListDeltaRequest {
	s.DriveId = &v
	return s
}

func (s *ListDeltaRequest) SetLimit(v int32) *ListDeltaRequest {
	s.Limit = &v
	return s
}

func (s *ListDeltaRequest) SetSyncRootId(v string) *ListDeltaRequest {
	s.SyncRootId = &v
	return s
}

type ListDeltaResponseBody struct {
	// The cursor of the incremental information.
	//
	// example:
	//
	// 1WQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Cursor *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	// Indicates whether more information is returned.
	//
	// example:
	//
	// true
	HasMore *bool `json:"has_more,omitempty" xml:"has_more,omitempty"`
	// The incremental information returned.
	Items []*ListDeltaResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListDeltaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDeltaResponseBody) GoString() string {
	return s.String()
}

func (s *ListDeltaResponseBody) SetCursor(v string) *ListDeltaResponseBody {
	s.Cursor = &v
	return s
}

func (s *ListDeltaResponseBody) SetHasMore(v bool) *ListDeltaResponseBody {
	s.HasMore = &v
	return s
}

func (s *ListDeltaResponseBody) SetItems(v []*ListDeltaResponseBodyItems) *ListDeltaResponseBody {
	s.Items = v
	return s
}

type ListDeltaResponseBodyItems struct {
	// The information about the file.
	File *File `json:"file,omitempty" xml:"file,omitempty"`
	// The file ID.
	//
	// example:
	//
	// 122fb09598ae66777c7040109a16f49381f6abe2
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The operation that is performed. Valid values: Valid values:
	//
	// 	- create
	//
	// 	- overwrite
	//
	// 	- delete
	//
	// 	- update
	//
	// 	- move
	//
	// 	- trash
	//
	// 	- restore
	//
	// 	- rename
	//
	// example:
	//
	// create
	Op *string `json:"op,omitempty" xml:"op,omitempty"`
}

func (s ListDeltaResponseBodyItems) String() string {
	return tea.Prettify(s)
}

func (s ListDeltaResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListDeltaResponseBodyItems) SetFile(v *File) *ListDeltaResponseBodyItems {
	s.File = v
	return s
}

func (s *ListDeltaResponseBodyItems) SetFileId(v string) *ListDeltaResponseBodyItems {
	s.FileId = &v
	return s
}

func (s *ListDeltaResponseBodyItems) SetOp(v string) *ListDeltaResponseBodyItems {
	s.Op = &v
	return s
}

type ListDeltaResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDeltaResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDeltaResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDeltaResponse) GoString() string {
	return s.String()
}

func (s *ListDeltaResponse) SetHeaders(v map[string]*string) *ListDeltaResponse {
	s.Headers = v
	return s
}

func (s *ListDeltaResponse) SetStatusCode(v int32) *ListDeltaResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDeltaResponse) SetBody(v *ListDeltaResponseBody) *ListDeltaResponse {
	s.Body = v
	return s
}

type ListDomainsRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 50.
	//
	// example:
	//
	// 60
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The ID of the parent domain.
	//
	// example:
	//
	// bj1
	ParentDomainId *string `json:"parent_domain_id,omitempty" xml:"parent_domain_id,omitempty"`
	ServiceCode    *string `json:"service_code,omitempty" xml:"service_code,omitempty"`
}

func (s ListDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsRequest) GoString() string {
	return s.String()
}

func (s *ListDomainsRequest) SetLimit(v int64) *ListDomainsRequest {
	s.Limit = &v
	return s
}

func (s *ListDomainsRequest) SetMarker(v string) *ListDomainsRequest {
	s.Marker = &v
	return s
}

func (s *ListDomainsRequest) SetParentDomainId(v string) *ListDomainsRequest {
	s.ParentDomainId = &v
	return s
}

func (s *ListDomainsRequest) SetServiceCode(v string) *ListDomainsRequest {
	s.ServiceCode = &v
	return s
}

type ListDomainsResponseBody struct {
	// The information about the domains.
	Items []*Domain `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDomainsResponseBody) SetItems(v []*Domain) *ListDomainsResponseBody {
	s.Items = v
	return s
}

func (s *ListDomainsResponseBody) SetNextMarker(v string) *ListDomainsResponseBody {
	s.NextMarker = &v
	return s
}

type ListDomainsResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDomainsResponse) GoString() string {
	return s.String()
}

func (s *ListDomainsResponse) SetHeaders(v map[string]*string) *ListDomainsResponse {
	s.Headers = v
	return s
}

func (s *ListDomainsResponse) SetStatusCode(v int32) *ListDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDomainsResponse) SetBody(v *ListDomainsResponseBody) *ListDomainsResponse {
	s.Body = v
	return s
}

type ListDriveRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The owner of the drive. If this parameter is not specified, all drives are returned.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The type of the owner. Valid values:
	//
	// user and group.
	//
	// By default, drives of all owner types are returned.
	//
	// example:
	//
	// user
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
}

func (s ListDriveRequest) String() string {
	return tea.Prettify(s)
}

func (s ListDriveRequest) GoString() string {
	return s.String()
}

func (s *ListDriveRequest) SetLimit(v int32) *ListDriveRequest {
	s.Limit = &v
	return s
}

func (s *ListDriveRequest) SetMarker(v string) *ListDriveRequest {
	s.Marker = &v
	return s
}

func (s *ListDriveRequest) SetOwner(v string) *ListDriveRequest {
	s.Owner = &v
	return s
}

func (s *ListDriveRequest) SetOwnerType(v string) *ListDriveRequest {
	s.OwnerType = &v
	return s
}

type ListDriveResponseBody struct {
	// The queried drives.
	Items []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListDriveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListDriveResponseBody) GoString() string {
	return s.String()
}

func (s *ListDriveResponseBody) SetItems(v []*Drive) *ListDriveResponseBody {
	s.Items = v
	return s
}

func (s *ListDriveResponseBody) SetNextMarker(v string) *ListDriveResponseBody {
	s.NextMarker = &v
	return s
}

type ListDriveResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListDriveResponse) String() string {
	return tea.Prettify(s)
}

func (s ListDriveResponse) GoString() string {
	return s.String()
}

func (s *ListDriveResponse) SetHeaders(v map[string]*string) *ListDriveResponse {
	s.Headers = v
	return s
}

func (s *ListDriveResponse) SetStatusCode(v int32) *ListDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *ListDriveResponse) SetBody(v *ListDriveResponseBody) *ListDriveResponse {
	s.Body = v
	return s
}

type ListFacegroupsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The filter condition that is used to query groups. The value can be up to 128 characters in length. An exact match is used.
	Remarks          *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
	ReturnTotalCount *bool   `json:"return_total_count,omitempty" xml:"return_total_count,omitempty"`
}

func (s ListFacegroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFacegroupsRequest) GoString() string {
	return s.String()
}

func (s *ListFacegroupsRequest) SetDriveId(v string) *ListFacegroupsRequest {
	s.DriveId = &v
	return s
}

func (s *ListFacegroupsRequest) SetLimit(v int32) *ListFacegroupsRequest {
	s.Limit = &v
	return s
}

func (s *ListFacegroupsRequest) SetMarker(v string) *ListFacegroupsRequest {
	s.Marker = &v
	return s
}

func (s *ListFacegroupsRequest) SetRemarks(v string) *ListFacegroupsRequest {
	s.Remarks = &v
	return s
}

func (s *ListFacegroupsRequest) SetReturnTotalCount(v bool) *ListFacegroupsRequest {
	s.ReturnTotalCount = &v
	return s
}

type ListFacegroupsResponseBody struct {
	// The information about the face-based groups.
	Items []*FaceGroup `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	TotalCount *int64  `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s ListFacegroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFacegroupsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFacegroupsResponseBody) SetItems(v []*FaceGroup) *ListFacegroupsResponseBody {
	s.Items = v
	return s
}

func (s *ListFacegroupsResponseBody) SetNextMarker(v string) *ListFacegroupsResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListFacegroupsResponseBody) SetTotalCount(v int64) *ListFacegroupsResponseBody {
	s.TotalCount = &v
	return s
}

type ListFacegroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFacegroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFacegroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFacegroupsResponse) GoString() string {
	return s.String()
}

func (s *ListFacegroupsResponse) SetHeaders(v map[string]*string) *ListFacegroupsResponse {
	s.Headers = v
	return s
}

func (s *ListFacegroupsResponse) SetStatusCode(v int32) *ListFacegroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFacegroupsResponse) SetBody(v *ListFacegroupsResponseBody) *ListFacegroupsResponse {
	s.Body = v
	return s
}

type ListFileRequest struct {
	// The category of the file. Valid values:
	//
	// app: installation package. zip: compressed package. image: image. doc: document. video: video. audio: audio. others: other files.
	//
	// By default, files of all categories are returned.
	//
	// example:
	//
	// image
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The fields to return.
	//
	// 1.  If this parameter is set to \\*, all fields of the file except the fields that must be specified are returned.
	//
	// 2.  If only specific fields are required, you can specify the following fields: url, exif, cropping_suggestion, characteristic_hash, video_metadata, and video_preview_metadata. If multiple fields are required, separate them with commas (,). Example: url,exif.
	//
	// 3.  The investigation_info field is returned only if you specify this field.
	//
	// By default, all fields except the fields that must be specified are returned.
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The sorting field.
	//
	// Default value: created_at.
	//
	// Valid values:
	//
	// 	- updated_at
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts the results based on the time when the file was last modified
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- size
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts the results based on the size of the file
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- name
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts the results based on the name of the file
	//
	//     <!-- -->
	//
	//     .
	//
	// 	- created_at
	//
	//     <!-- -->
	//
	//     :
	//
	//     <!-- -->
	//
	//     sorts the results based on the time when the file was created
	//
	//     <!-- -->
	//
	//     .
	//
	// example:
	//
	// updated_at
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// The sorting direction. Valid values:
	//
	// ASC: ascending order. DESC: descending order.
	//
	// Default value: ASC.
	//
	// example:
	//
	// ASC
	OrderDirection *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
	// The ID of the parent folder. If the parent folder is a root directory, set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// root
	ParentFileId *string `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	// The share ID. If you want to manage a file by using a share link, carry the `x-share-token` header for authentication in the request and specify share_id. In this case, `drive_id` is invalid. Otherwise, use an `AccessKey pair` or `access token` for authentication and specify `drive_id`. You must specify one of `share_id` and `drive_id`.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The state of the file. Valid values:
	//
	// available: Only normal files are returned. uploading: Only files that are being uploaded are returned.
	//
	// By default, only files in the available state are returned.
	//
	// example:
	//
	// available
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The thumbnail configurations. Up to five thumbnails can be returned at a time. The value contains key-value pairs. You can customize the keys. The URL of a thumbnail is returned based on the key.
	ThumbnailProcesses map[string]*ImageProcess `json:"thumbnail_processes,omitempty" xml:"thumbnail_processes,omitempty"`
	// The type of the file. Valid values:
	//
	// file: Only files are returned. folder: Only folders are returned.
	//
	// By default, files of all types are returned.
	//
	// example:
	//
	// file
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFileRequest) GoString() string {
	return s.String()
}

func (s *ListFileRequest) SetCategory(v string) *ListFileRequest {
	s.Category = &v
	return s
}

func (s *ListFileRequest) SetDriveId(v string) *ListFileRequest {
	s.DriveId = &v
	return s
}

func (s *ListFileRequest) SetFields(v string) *ListFileRequest {
	s.Fields = &v
	return s
}

func (s *ListFileRequest) SetLimit(v int32) *ListFileRequest {
	s.Limit = &v
	return s
}

func (s *ListFileRequest) SetMarker(v string) *ListFileRequest {
	s.Marker = &v
	return s
}

func (s *ListFileRequest) SetOrderBy(v string) *ListFileRequest {
	s.OrderBy = &v
	return s
}

func (s *ListFileRequest) SetOrderDirection(v string) *ListFileRequest {
	s.OrderDirection = &v
	return s
}

func (s *ListFileRequest) SetParentFileId(v string) *ListFileRequest {
	s.ParentFileId = &v
	return s
}

func (s *ListFileRequest) SetShareId(v string) *ListFileRequest {
	s.ShareId = &v
	return s
}

func (s *ListFileRequest) SetStatus(v string) *ListFileRequest {
	s.Status = &v
	return s
}

func (s *ListFileRequest) SetThumbnailProcesses(v map[string]*ImageProcess) *ListFileRequest {
	s.ThumbnailProcesses = v
	return s
}

func (s *ListFileRequest) SetType(v string) *ListFileRequest {
	s.Type = &v
	return s
}

type ListFileResponseBody struct {
	// The queried files.
	Items []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListFileResponseBody) SetItems(v []*File) *ListFileResponseBody {
	s.Items = v
	return s
}

func (s *ListFileResponseBody) SetNextMarker(v string) *ListFileResponseBody {
	s.NextMarker = &v
	return s
}

type ListFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFileResponse) GoString() string {
	return s.String()
}

func (s *ListFileResponse) SetHeaders(v map[string]*string) *ListFileResponse {
	s.Headers = v
	return s
}

func (s *ListFileResponse) SetStatusCode(v int32) *ListFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListFileResponse) SetBody(v *ListFileResponseBody) *ListFileResponse {
	s.Body = v
	return s
}

type ListGroupRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRequest) GoString() string {
	return s.String()
}

func (s *ListGroupRequest) SetLimit(v int32) *ListGroupRequest {
	s.Limit = &v
	return s
}

func (s *ListGroupRequest) SetMarker(v string) *ListGroupRequest {
	s.Marker = &v
	return s
}

type ListGroupResponseBody struct {
	// The information about the groups.
	Items []*Group `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListGroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupResponseBody) SetItems(v []*Group) *ListGroupResponseBody {
	s.Items = v
	return s
}

func (s *ListGroupResponseBody) SetNextMarker(v string) *ListGroupResponseBody {
	s.NextMarker = &v
	return s
}

type ListGroupResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupResponse) GoString() string {
	return s.String()
}

func (s *ListGroupResponse) SetHeaders(v map[string]*string) *ListGroupResponse {
	s.Headers = v
	return s
}

func (s *ListGroupResponse) SetStatusCode(v int32) *ListGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupResponse) SetBody(v *ListGroupResponseBody) *ListGroupResponse {
	s.Body = v
	return s
}

type ListGroupMemberRequest struct {
	// The ID of the group of which you want to query members.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e5***2c2
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The member type. If you do not specify this parameter, both types of members are returned. Valid values:
	//
	// 	- user
	//
	// 	- group
	//
	// Note: A group can be a member of only one group. It cannot be a member of multiple groups. A user can be a member of multiple groups.
	//
	// example:
	//
	// user
	MemberType *string `json:"member_type,omitempty" xml:"member_type,omitempty"`
}

func (s ListGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *ListGroupMemberRequest) SetGroupId(v string) *ListGroupMemberRequest {
	s.GroupId = &v
	return s
}

func (s *ListGroupMemberRequest) SetLimit(v int32) *ListGroupMemberRequest {
	s.Limit = &v
	return s
}

func (s *ListGroupMemberRequest) SetMarker(v string) *ListGroupMemberRequest {
	s.Marker = &v
	return s
}

func (s *ListGroupMemberRequest) SetMemberType(v string) *ListGroupMemberRequest {
	s.MemberType = &v
	return s
}

type ListGroupMemberResponseBody struct {
	// The information about the groups.
	GroupItems []*Group `json:"group_items,omitempty" xml:"group_items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhM1
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	// The information about the users.
	UserItems []*User `json:"user_items,omitempty" xml:"user_items,omitempty" type:"Repeated"`
}

func (s ListGroupMemberResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMemberResponseBody) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponseBody) SetGroupItems(v []*Group) *ListGroupMemberResponseBody {
	s.GroupItems = v
	return s
}

func (s *ListGroupMemberResponseBody) SetNextMarker(v string) *ListGroupMemberResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListGroupMemberResponseBody) SetUserItems(v []*User) *ListGroupMemberResponseBody {
	s.UserItems = v
	return s
}

type ListGroupMemberResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListGroupMemberResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s ListGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *ListGroupMemberResponse) SetHeaders(v map[string]*string) *ListGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *ListGroupMemberResponse) SetStatusCode(v int32) *ListGroupMemberResponse {
	s.StatusCode = &v
	return s
}

func (s *ListGroupMemberResponse) SetBody(v *ListGroupMemberResponseBody) *ListGroupMemberResponse {
	s.Body = v
	return s
}

type ListIdentityRoleRequest struct {
	// This parameter is required.
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
}

func (s ListIdentityRoleRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIdentityRoleRequest) GoString() string {
	return s.String()
}

func (s *ListIdentityRoleRequest) SetIdentity(v *Identity) *ListIdentityRoleRequest {
	s.Identity = v
	return s
}

type ListIdentityRoleResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *BaseRoleMemberResponse `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdentityRoleResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIdentityRoleResponse) GoString() string {
	return s.String()
}

func (s *ListIdentityRoleResponse) SetHeaders(v map[string]*string) *ListIdentityRoleResponse {
	s.Headers = v
	return s
}

func (s *ListIdentityRoleResponse) SetStatusCode(v int32) *ListIdentityRoleResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdentityRoleResponse) SetBody(v *BaseRoleMemberResponse) *ListIdentityRoleResponse {
	s.Body = v
	return s
}

type ListIdentityToBenefitPkgMappingRequest struct {
	// The unique identifier of the entity.
	//
	// If you call this operation to manage the benefits of a user, set this parameter to the ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user123
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// The type of the entity. If you call this operation to manage the benefits of a user, set this parameter to user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
	// Specifies whether to return the benefit packages that expire. Default value: false.
	//
	// example:
	//
	// false
	IncludeExpired *bool `json:"include_expired,omitempty" xml:"include_expired,omitempty"`
}

func (s ListIdentityToBenefitPkgMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s ListIdentityToBenefitPkgMappingRequest) GoString() string {
	return s.String()
}

func (s *ListIdentityToBenefitPkgMappingRequest) SetIdentityId(v string) *ListIdentityToBenefitPkgMappingRequest {
	s.IdentityId = &v
	return s
}

func (s *ListIdentityToBenefitPkgMappingRequest) SetIdentityType(v string) *ListIdentityToBenefitPkgMappingRequest {
	s.IdentityType = &v
	return s
}

func (s *ListIdentityToBenefitPkgMappingRequest) SetIncludeExpired(v bool) *ListIdentityToBenefitPkgMappingRequest {
	s.IncludeExpired = &v
	return s
}

type ListIdentityToBenefitPkgMappingResponseBody struct {
	// The information about the benefit packages that are associated with an entity.
	Items []*IdentityToBenefitPkgMapping `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListIdentityToBenefitPkgMappingResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListIdentityToBenefitPkgMappingResponseBody) GoString() string {
	return s.String()
}

func (s *ListIdentityToBenefitPkgMappingResponseBody) SetItems(v []*IdentityToBenefitPkgMapping) *ListIdentityToBenefitPkgMappingResponseBody {
	s.Items = v
	return s
}

type ListIdentityToBenefitPkgMappingResponse struct {
	Headers    map[string]*string                           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListIdentityToBenefitPkgMappingResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListIdentityToBenefitPkgMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s ListIdentityToBenefitPkgMappingResponse) GoString() string {
	return s.String()
}

func (s *ListIdentityToBenefitPkgMappingResponse) SetHeaders(v map[string]*string) *ListIdentityToBenefitPkgMappingResponse {
	s.Headers = v
	return s
}

func (s *ListIdentityToBenefitPkgMappingResponse) SetStatusCode(v int32) *ListIdentityToBenefitPkgMappingResponse {
	s.StatusCode = &v
	return s
}

func (s *ListIdentityToBenefitPkgMappingResponse) SetBody(v *ListIdentityToBenefitPkgMappingResponseBody) *ListIdentityToBenefitPkgMappingResponse {
	s.Body = v
	return s
}

type ListMyDrivesRequest struct {
	// The maximum number of results to return. Default value: 100. Valid values: 1 to 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListMyDrivesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMyDrivesRequest) GoString() string {
	return s.String()
}

func (s *ListMyDrivesRequest) SetLimit(v int32) *ListMyDrivesRequest {
	s.Limit = &v
	return s
}

func (s *ListMyDrivesRequest) SetMarker(v string) *ListMyDrivesRequest {
	s.Marker = &v
	return s
}

type ListMyDrivesResponseBody struct {
	// The queried drives.
	Items []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListMyDrivesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMyDrivesResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyDrivesResponseBody) SetItems(v []*Drive) *ListMyDrivesResponseBody {
	s.Items = v
	return s
}

func (s *ListMyDrivesResponseBody) SetNextMarker(v string) *ListMyDrivesResponseBody {
	s.NextMarker = &v
	return s
}

type ListMyDrivesResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyDrivesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyDrivesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMyDrivesResponse) GoString() string {
	return s.String()
}

func (s *ListMyDrivesResponse) SetHeaders(v map[string]*string) *ListMyDrivesResponse {
	s.Headers = v
	return s
}

func (s *ListMyDrivesResponse) SetStatusCode(v int32) *ListMyDrivesResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyDrivesResponse) SetBody(v *ListMyDrivesResponseBody) *ListMyDrivesResponse {
	s.Body = v
	return s
}

type ListMyGroupDriveRequest struct {
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListMyGroupDriveRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupDriveRequest) GoString() string {
	return s.String()
}

func (s *ListMyGroupDriveRequest) SetDriveName(v string) *ListMyGroupDriveRequest {
	s.DriveName = &v
	return s
}

func (s *ListMyGroupDriveRequest) SetLimit(v int32) *ListMyGroupDriveRequest {
	s.Limit = &v
	return s
}

func (s *ListMyGroupDriveRequest) SetMarker(v string) *ListMyGroupDriveRequest {
	s.Marker = &v
	return s
}

type ListMyGroupDriveResponseBody struct {
	// The information about the team drives.
	Items []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker     *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	RootGroupDrive *Drive  `json:"root_group_drive,omitempty" xml:"root_group_drive,omitempty"`
}

func (s ListMyGroupDriveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupDriveResponseBody) GoString() string {
	return s.String()
}

func (s *ListMyGroupDriveResponseBody) SetItems(v []*Drive) *ListMyGroupDriveResponseBody {
	s.Items = v
	return s
}

func (s *ListMyGroupDriveResponseBody) SetNextMarker(v string) *ListMyGroupDriveResponseBody {
	s.NextMarker = &v
	return s
}

func (s *ListMyGroupDriveResponseBody) SetRootGroupDrive(v *Drive) *ListMyGroupDriveResponseBody {
	s.RootGroupDrive = v
	return s
}

type ListMyGroupDriveResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListMyGroupDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListMyGroupDriveResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMyGroupDriveResponse) GoString() string {
	return s.String()
}

func (s *ListMyGroupDriveResponse) SetHeaders(v map[string]*string) *ListMyGroupDriveResponse {
	s.Headers = v
	return s
}

func (s *ListMyGroupDriveResponse) SetStatusCode(v int32) *ListMyGroupDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *ListMyGroupDriveResponse) SetBody(v *ListMyGroupDriveResponseBody) *ListMyGroupDriveResponse {
	s.Body = v
	return s
}

type ListReceivedFileRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is empty.
	//
	// example:
	//
	// eym***
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListReceivedFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ListReceivedFileRequest) GoString() string {
	return s.String()
}

func (s *ListReceivedFileRequest) SetLimit(v int32) *ListReceivedFileRequest {
	s.Limit = &v
	return s
}

func (s *ListReceivedFileRequest) SetMarker(v string) *ListReceivedFileRequest {
	s.Marker = &v
	return s
}

type ListReceivedFileResponseBody struct {
	// The queried files.
	Items []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// eym***
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListReceivedFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListReceivedFileResponseBody) GoString() string {
	return s.String()
}

func (s *ListReceivedFileResponseBody) SetItems(v []*File) *ListReceivedFileResponseBody {
	s.Items = v
	return s
}

func (s *ListReceivedFileResponseBody) SetNextMarker(v string) *ListReceivedFileResponseBody {
	s.NextMarker = &v
	return s
}

type ListReceivedFileResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListReceivedFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListReceivedFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ListReceivedFileResponse) GoString() string {
	return s.String()
}

func (s *ListReceivedFileResponse) SetHeaders(v map[string]*string) *ListReceivedFileResponse {
	s.Headers = v
	return s
}

func (s *ListReceivedFileResponse) SetStatusCode(v int32) *ListReceivedFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ListReceivedFileResponse) SetBody(v *ListReceivedFileResponseBody) *ListReceivedFileResponse {
	s.Body = v
	return s
}

type ListRecyclebinRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Specifies the returned fields.
	//
	// 1\\. If you set this parameter to \\*, all fields of the file are returned.
	//
	// 2\\. If you set this parameter to a null value or leave this parameter empty, the fields, such as file creator, file modifier, and custom tags, are not returned.
	//
	// The default value is a null value, which indicates that only some fields are returned.
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 200. Default value: 50.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListRecyclebinRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRecyclebinRequest) GoString() string {
	return s.String()
}

func (s *ListRecyclebinRequest) SetDriveId(v string) *ListRecyclebinRequest {
	s.DriveId = &v
	return s
}

func (s *ListRecyclebinRequest) SetFields(v string) *ListRecyclebinRequest {
	s.Fields = &v
	return s
}

func (s *ListRecyclebinRequest) SetLimit(v int32) *ListRecyclebinRequest {
	s.Limit = &v
	return s
}

func (s *ListRecyclebinRequest) SetMarker(v string) *ListRecyclebinRequest {
	s.Marker = &v
	return s
}

type ListRecyclebinResponseBody struct {
	// The information about the files and folders in the recycle bin.
	Items []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhM1
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListRecyclebinResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRecyclebinResponseBody) GoString() string {
	return s.String()
}

func (s *ListRecyclebinResponseBody) SetItems(v []*File) *ListRecyclebinResponseBody {
	s.Items = v
	return s
}

func (s *ListRecyclebinResponseBody) SetNextMarker(v string) *ListRecyclebinResponseBody {
	s.NextMarker = &v
	return s
}

type ListRecyclebinResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRecyclebinResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRecyclebinResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRecyclebinResponse) GoString() string {
	return s.String()
}

func (s *ListRecyclebinResponse) SetHeaders(v map[string]*string) *ListRecyclebinResponse {
	s.Headers = v
	return s
}

func (s *ListRecyclebinResponse) SetStatusCode(v int32) *ListRecyclebinResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRecyclebinResponse) SetBody(v *ListRecyclebinResponseBody) *ListRecyclebinResponse {
	s.Body = v
	return s
}

type ListRevisionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Specifies the returned fields.
	//
	// By default, this parameter is left empty. If you set this parameter to \\*, all fields are returned. If you leave this parameter empty, the creator of the file is not returned.
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// Default value: 50.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 100
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// 40CB7794C929
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListRevisionRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRevisionRequest) GoString() string {
	return s.String()
}

func (s *ListRevisionRequest) SetDriveId(v string) *ListRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *ListRevisionRequest) SetFields(v string) *ListRevisionRequest {
	s.Fields = &v
	return s
}

func (s *ListRevisionRequest) SetFileId(v string) *ListRevisionRequest {
	s.FileId = &v
	return s
}

func (s *ListRevisionRequest) SetLimit(v int64) *ListRevisionRequest {
	s.Limit = &v
	return s
}

func (s *ListRevisionRequest) SetMarker(v string) *ListRevisionRequest {
	s.Marker = &v
	return s
}

type ListRevisionResponseBody struct {
	// The information about the versions.
	Items []*Revision `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListRevisionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRevisionResponseBody) GoString() string {
	return s.String()
}

func (s *ListRevisionResponseBody) SetItems(v []*Revision) *ListRevisionResponseBody {
	s.Items = v
	return s
}

func (s *ListRevisionResponseBody) SetNextMarker(v string) *ListRevisionResponseBody {
	s.NextMarker = &v
	return s
}

type ListRevisionResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListRevisionResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRevisionResponse) GoString() string {
	return s.String()
}

func (s *ListRevisionResponse) SetHeaders(v map[string]*string) *ListRevisionResponse {
	s.Headers = v
	return s
}

func (s *ListRevisionResponse) SetStatusCode(v int32) *ListRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *ListRevisionResponse) SetBody(v *ListRevisionResponseBody) *ListRevisionResponse {
	s.Body = v
	return s
}

type ListShareLinkRequest struct {
	// The creator of the share.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Creator *string `json:"creator,omitempty" xml:"creator,omitempty"`
	// Specifies whether to return the shares that are canceled.
	//
	// example:
	//
	// true
	IncludeCancelled *bool `json:"include_cancelled,omitempty" xml:"include_cancelled,omitempty"`
	// The maximum number of results to return. Valid values: 0 to 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The field by which to sort the returned results. Default value: created_at. Valid values:
	//
	// 	- share_name: sorts the results by the name of the share.
	//
	// 	- updated_at: sorts the results by the time when the share was modified.
	//
	// 	- description: sorts the results by the description of the share.
	//
	// 	- created_at: sorts the results by the time when the share was created.
	//
	// example:
	//
	// share_name
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// The order in which you want to sort the returned results. By default, order_direction is set to DESC if order_by is set to created_at or updated_at, and is set to ASC if order_by is set to other values. Valid values:
	//
	// 	- ASC: sorts the results in ascending order.
	//
	// 	- DESC: sorts the results in descending order.
	//
	// example:
	//
	// ASC
	OrderDirection *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
}

func (s ListShareLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s ListShareLinkRequest) GoString() string {
	return s.String()
}

func (s *ListShareLinkRequest) SetCreator(v string) *ListShareLinkRequest {
	s.Creator = &v
	return s
}

func (s *ListShareLinkRequest) SetIncludeCancelled(v bool) *ListShareLinkRequest {
	s.IncludeCancelled = &v
	return s
}

func (s *ListShareLinkRequest) SetLimit(v int32) *ListShareLinkRequest {
	s.Limit = &v
	return s
}

func (s *ListShareLinkRequest) SetMarker(v string) *ListShareLinkRequest {
	s.Marker = &v
	return s
}

func (s *ListShareLinkRequest) SetOrderBy(v string) *ListShareLinkRequest {
	s.OrderBy = &v
	return s
}

func (s *ListShareLinkRequest) SetOrderDirection(v string) *ListShareLinkRequest {
	s.OrderDirection = &v
	return s
}

type ListShareLinkResponseBody struct {
	// The information about the shares.
	Items []*ShareLink `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListShareLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListShareLinkResponseBody) GoString() string {
	return s.String()
}

func (s *ListShareLinkResponseBody) SetItems(v []*ShareLink) *ListShareLinkResponseBody {
	s.Items = v
	return s
}

func (s *ListShareLinkResponseBody) SetNextMarker(v string) *ListShareLinkResponseBody {
	s.NextMarker = &v
	return s
}

type ListShareLinkResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListShareLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListShareLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s ListShareLinkResponse) GoString() string {
	return s.String()
}

func (s *ListShareLinkResponse) SetHeaders(v map[string]*string) *ListShareLinkResponse {
	s.Headers = v
	return s
}

func (s *ListShareLinkResponse) SetStatusCode(v int32) *ListShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *ListShareLinkResponse) SetBody(v *ListShareLinkResponseBody) *ListShareLinkResponse {
	s.Body = v
	return s
}

type ListTagsRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The method that is used to generate the thumbnail of an image.
	//
	// example:
	//
	// image/resize,w_200
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// The method that is used to generate the thumbnail of a video.
	//
	// example:
	//
	// video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s ListTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) SetDriveId(v string) *ListTagsRequest {
	s.DriveId = &v
	return s
}

func (s *ListTagsRequest) SetImageThumbnailProcess(v string) *ListTagsRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *ListTagsRequest) SetVideoThumbnailProcess(v string) *ListTagsRequest {
	s.VideoThumbnailProcess = &v
	return s
}

type ListTagsResponseBody struct {
	// The information about the tags.
	Tags []*ImageTag `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) SetTags(v []*ImageTag) *ListTagsResponseBody {
	s.Tags = v
	return s
}

type ListTagsResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponse) GoString() string {
	return s.String()
}

func (s *ListTagsResponse) SetHeaders(v map[string]*string) *ListTagsResponse {
	s.Headers = v
	return s
}

func (s *ListTagsResponse) SetStatusCode(v int32) *ListTagsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

type ListUploadedPartsRequest struct {
	// The drive ID. This parameter is required if the file is not uploaded by using the share URL of the file.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 322fb07b975f4b0ae1b543fe8475eee4c19eb2b2
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	PartNumberMarker *int32 `json:"part_number_marker,omitempty" xml:"part_number_marker,omitempty"`
	// The share ID. This parameter is required if the file is uploaded by using the share URL of the file.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The ID of the upload task.
	//
	// This parameter is required.
	//
	// example:
	//
	// 00166D06127B413BA1EC8ABB1144D101
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s ListUploadedPartsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUploadedPartsRequest) GoString() string {
	return s.String()
}

func (s *ListUploadedPartsRequest) SetDriveId(v string) *ListUploadedPartsRequest {
	s.DriveId = &v
	return s
}

func (s *ListUploadedPartsRequest) SetFileId(v string) *ListUploadedPartsRequest {
	s.FileId = &v
	return s
}

func (s *ListUploadedPartsRequest) SetLimit(v int32) *ListUploadedPartsRequest {
	s.Limit = &v
	return s
}

func (s *ListUploadedPartsRequest) SetPartNumberMarker(v int32) *ListUploadedPartsRequest {
	s.PartNumberMarker = &v
	return s
}

func (s *ListUploadedPartsRequest) SetShareId(v string) *ListUploadedPartsRequest {
	s.ShareId = &v
	return s
}

func (s *ListUploadedPartsRequest) SetUploadId(v string) *ListUploadedPartsRequest {
	s.UploadId = &v
	return s
}

type ListUploadedPartsResponseBody struct {
	// The file ID.
	//
	// example:
	//
	// 322fb07b975f4b0ae1b543fe8475eee4c19eb2b2
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextPartNumberMarker *string `json:"next_part_number_marker,omitempty" xml:"next_part_number_marker,omitempty"`
	// Indicates whether the parallel upload feature is enabled.
	//
	// example:
	//
	// false
	ParallelUpload *bool `json:"parallel_upload,omitempty" xml:"parallel_upload,omitempty"`
	// The ID of the upload task.
	//
	// example:
	//
	// 00166D06127B413BA1EC8ABB1144D101
	UploadId *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	// The information about the file parts.
	UploadedParts []*UploadPartInfo `json:"uploaded_parts,omitempty" xml:"uploaded_parts,omitempty" type:"Repeated"`
}

func (s ListUploadedPartsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUploadedPartsResponseBody) GoString() string {
	return s.String()
}

func (s *ListUploadedPartsResponseBody) SetFileId(v string) *ListUploadedPartsResponseBody {
	s.FileId = &v
	return s
}

func (s *ListUploadedPartsResponseBody) SetNextPartNumberMarker(v string) *ListUploadedPartsResponseBody {
	s.NextPartNumberMarker = &v
	return s
}

func (s *ListUploadedPartsResponseBody) SetParallelUpload(v bool) *ListUploadedPartsResponseBody {
	s.ParallelUpload = &v
	return s
}

func (s *ListUploadedPartsResponseBody) SetUploadId(v string) *ListUploadedPartsResponseBody {
	s.UploadId = &v
	return s
}

func (s *ListUploadedPartsResponseBody) SetUploadedParts(v []*UploadPartInfo) *ListUploadedPartsResponseBody {
	s.UploadedParts = v
	return s
}

type ListUploadedPartsResponse struct {
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUploadedPartsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUploadedPartsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUploadedPartsResponse) GoString() string {
	return s.String()
}

func (s *ListUploadedPartsResponse) SetHeaders(v map[string]*string) *ListUploadedPartsResponse {
	s.Headers = v
	return s
}

func (s *ListUploadedPartsResponse) SetStatusCode(v int32) *ListUploadedPartsResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUploadedPartsResponse) SetBody(v *ListUploadedPartsResponseBody) *ListUploadedPartsResponse {
	s.Body = v
	return s
}

type ListUserRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker. By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListUserRequest) String() string {
	return tea.Prettify(s)
}

func (s ListUserRequest) GoString() string {
	return s.String()
}

func (s *ListUserRequest) SetLimit(v int32) *ListUserRequest {
	s.Limit = &v
	return s
}

func (s *ListUserRequest) SetMarker(v string) *ListUserRequest {
	s.Marker = &v
	return s
}

type ListUserResponseBody struct {
	// The information about the users.
	Items []*User `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ListUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListUserResponseBody) GoString() string {
	return s.String()
}

func (s *ListUserResponseBody) SetItems(v []*User) *ListUserResponseBody {
	s.Items = v
	return s
}

func (s *ListUserResponseBody) SetNextMarker(v string) *ListUserResponseBody {
	s.NextMarker = &v
	return s
}

type ListUserResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ListUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ListUserResponse) String() string {
	return tea.Prettify(s)
}

func (s ListUserResponse) GoString() string {
	return s.String()
}

func (s *ListUserResponse) SetHeaders(v map[string]*string) *ListUserResponse {
	s.Headers = v
	return s
}

func (s *ListUserResponse) SetStatusCode(v int32) *ListUserResponse {
	s.StatusCode = &v
	return s
}

func (s *ListUserResponse) SetBody(v *ListUserResponseBody) *ListUserResponse {
	s.Body = v
	return s
}

type MoveFileRequest struct {
	// The processing method that is used if the file that you want to move has the same name as an existing file in the destination directory. Valid values:
	//
	// ignore: allows you to move the file by using the same name as an existing file in the destination directory.
	//
	// auto_rename: automatically renames the file that has the same name exists in the destination directory. By default, the current point in time is added to the end of the file name. Example: xxx_20060102_150405.
	//
	// refuse: does not move the file that you want to move but returns the information about the file that has the same name in the destination directory.
	//
	// Default value: ignore.
	//
	// example:
	//
	// ignore
	CheckNameMode *string `json:"check_name_mode,omitempty" xml:"check_name_mode,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The ID of the destination parent directory to which you want to move a file or folder. If you want to move a file or folder to the root directory, set this parameter to root.
	//
	// This parameter is required.
	//
	// example:
	//
	// 6520943DC261
	ToParentFileId *string `json:"to_parent_file_id,omitempty" xml:"to_parent_file_id,omitempty"`
}

func (s MoveFileRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveFileRequest) GoString() string {
	return s.String()
}

func (s *MoveFileRequest) SetCheckNameMode(v string) *MoveFileRequest {
	s.CheckNameMode = &v
	return s
}

func (s *MoveFileRequest) SetDriveId(v string) *MoveFileRequest {
	s.DriveId = &v
	return s
}

func (s *MoveFileRequest) SetFileId(v string) *MoveFileRequest {
	s.FileId = &v
	return s
}

func (s *MoveFileRequest) SetToParentFileId(v string) *MoveFileRequest {
	s.ToParentFileId = &v
	return s
}

type MoveFileResponseBody struct {
	// The ID of the asynchronous task.
	//
	// If an empty string is returned, the file is moved.
	//
	// If a non-empty string is returned, an asynchronous task is required. You can call the GetAsyncTask operation to obtain the information about an asynchronous task based on the task ID.
	//
	// example:
	//
	// 23ebd1a24dba4166b1527add476ef2866051b4d5del106
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Indicates whether the file already exists in the destination directory.
	//
	// example:
	//
	// false
	Exist *bool `json:"exist,omitempty" xml:"exist,omitempty"`
	// The file ID.
	//
	// example:
	//
	// fileid1
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s MoveFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveFileResponseBody) GoString() string {
	return s.String()
}

func (s *MoveFileResponseBody) SetAsyncTaskId(v string) *MoveFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *MoveFileResponseBody) SetDomainId(v string) *MoveFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *MoveFileResponseBody) SetDriveId(v string) *MoveFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *MoveFileResponseBody) SetExist(v bool) *MoveFileResponseBody {
	s.Exist = &v
	return s
}

func (s *MoveFileResponseBody) SetFileId(v string) *MoveFileResponseBody {
	s.FileId = &v
	return s
}

type MoveFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *MoveFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s MoveFileResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveFileResponse) GoString() string {
	return s.String()
}

func (s *MoveFileResponse) SetHeaders(v map[string]*string) *MoveFileResponse {
	s.Headers = v
	return s
}

func (s *MoveFileResponse) SetStatusCode(v int32) *MoveFileResponse {
	s.StatusCode = &v
	return s
}

func (s *MoveFileResponse) SetBody(v *MoveFileResponseBody) *MoveFileResponse {
	s.Body = v
	return s
}

type QueryOrderPriceRequest struct {
	// This parameter is required.
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// This parameter is required.
	InstanceId *string `json:"instance_id,omitempty" xml:"instance_id,omitempty"`
	// This parameter is required.
	OrderType *string `json:"order_type,omitempty" xml:"order_type,omitempty"`
	// This parameter is required.
	Package *string `json:"package,omitempty" xml:"package,omitempty"`
	// This parameter is required.
	Period *int64 `json:"period,omitempty" xml:"period,omitempty"`
	// This parameter is required.
	PeriodUnit *string `json:"period_unit,omitempty" xml:"period_unit,omitempty"`
	// This parameter is required.
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
	// This parameter is required.
	UserCount *int64 `json:"user_count,omitempty" xml:"user_count,omitempty"`
}

func (s QueryOrderPriceRequest) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderPriceRequest) GoString() string {
	return s.String()
}

func (s *QueryOrderPriceRequest) SetCode(v string) *QueryOrderPriceRequest {
	s.Code = &v
	return s
}

func (s *QueryOrderPriceRequest) SetInstanceId(v string) *QueryOrderPriceRequest {
	s.InstanceId = &v
	return s
}

func (s *QueryOrderPriceRequest) SetOrderType(v string) *QueryOrderPriceRequest {
	s.OrderType = &v
	return s
}

func (s *QueryOrderPriceRequest) SetPackage(v string) *QueryOrderPriceRequest {
	s.Package = &v
	return s
}

func (s *QueryOrderPriceRequest) SetPeriod(v int64) *QueryOrderPriceRequest {
	s.Period = &v
	return s
}

func (s *QueryOrderPriceRequest) SetPeriodUnit(v string) *QueryOrderPriceRequest {
	s.PeriodUnit = &v
	return s
}

func (s *QueryOrderPriceRequest) SetTotalSize(v int64) *QueryOrderPriceRequest {
	s.TotalSize = &v
	return s
}

func (s *QueryOrderPriceRequest) SetUserCount(v int64) *QueryOrderPriceRequest {
	s.UserCount = &v
	return s
}

type QueryOrderPriceResponseBody struct {
	DiscountPrice *float64 `json:"discount_price,omitempty" xml:"discount_price,omitempty"`
	OriginalPrice *float64 `json:"original_price,omitempty" xml:"original_price,omitempty"`
	TradePrice    *float64 `json:"trade_price,omitempty" xml:"trade_price,omitempty"`
}

func (s QueryOrderPriceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderPriceResponseBody) GoString() string {
	return s.String()
}

func (s *QueryOrderPriceResponseBody) SetDiscountPrice(v float64) *QueryOrderPriceResponseBody {
	s.DiscountPrice = &v
	return s
}

func (s *QueryOrderPriceResponseBody) SetOriginalPrice(v float64) *QueryOrderPriceResponseBody {
	s.OriginalPrice = &v
	return s
}

func (s *QueryOrderPriceResponseBody) SetTradePrice(v float64) *QueryOrderPriceResponseBody {
	s.TradePrice = &v
	return s
}

type QueryOrderPriceResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *QueryOrderPriceResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s QueryOrderPriceResponse) String() string {
	return tea.Prettify(s)
}

func (s QueryOrderPriceResponse) GoString() string {
	return s.String()
}

func (s *QueryOrderPriceResponse) SetHeaders(v map[string]*string) *QueryOrderPriceResponse {
	s.Headers = v
	return s
}

func (s *QueryOrderPriceResponse) SetStatusCode(v int32) *QueryOrderPriceResponse {
	s.StatusCode = &v
	return s
}

func (s *QueryOrderPriceResponse) SetBody(v *QueryOrderPriceResponseBody) *QueryOrderPriceResponse {
	s.Body = v
	return s
}

type RemoveFaceGroupFileRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 123
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// Cluster-abc
	FaceGroupId *string `json:"face_group_id,omitempty" xml:"face_group_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// abcd
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s RemoveFaceGroupFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceGroupFileRequest) GoString() string {
	return s.String()
}

func (s *RemoveFaceGroupFileRequest) SetDriveId(v string) *RemoveFaceGroupFileRequest {
	s.DriveId = &v
	return s
}

func (s *RemoveFaceGroupFileRequest) SetFaceGroupId(v string) *RemoveFaceGroupFileRequest {
	s.FaceGroupId = &v
	return s
}

func (s *RemoveFaceGroupFileRequest) SetFileId(v string) *RemoveFaceGroupFileRequest {
	s.FileId = &v
	return s
}

type RemoveFaceGroupFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RemoveFaceGroupFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveFaceGroupFileResponse) GoString() string {
	return s.String()
}

func (s *RemoveFaceGroupFileResponse) SetHeaders(v map[string]*string) *RemoveFaceGroupFileResponse {
	s.Headers = v
	return s
}

func (s *RemoveFaceGroupFileResponse) SetStatusCode(v int32) *RemoveFaceGroupFileResponse {
	s.StatusCode = &v
	return s
}

type RemoveGroupMemberRequest struct {
	// The ID of the group from which you want to remove a member.
	//
	// This parameter is required.
	//
	// example:
	//
	// 3e5***2c2
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The ID of the member. If member_type is set to user, set this parameter to the ID of the corresponding user.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e4***1b1
	MemberId *string `json:"member_id,omitempty" xml:"member_id,omitempty"`
	// The type of the member that you want to remove from the group. Only common users can be removed. If you want to remove all members from a group, you can directly delete the group. Valid value:
	//
	// 	- user
	//
	// Note: A group can be a member of only one group. It cannot be a member of multiple groups. A user can be a member of multiple groups.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	MemberType *string `json:"member_type,omitempty" xml:"member_type,omitempty"`
}

func (s RemoveGroupMemberRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberRequest) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberRequest) SetGroupId(v string) *RemoveGroupMemberRequest {
	s.GroupId = &v
	return s
}

func (s *RemoveGroupMemberRequest) SetMemberId(v string) *RemoveGroupMemberRequest {
	s.MemberId = &v
	return s
}

func (s *RemoveGroupMemberRequest) SetMemberType(v string) *RemoveGroupMemberRequest {
	s.MemberType = &v
	return s
}

type RemoveGroupMemberResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s RemoveGroupMemberResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveGroupMemberResponse) GoString() string {
	return s.String()
}

func (s *RemoveGroupMemberResponse) SetHeaders(v map[string]*string) *RemoveGroupMemberResponse {
	s.Headers = v
	return s
}

func (s *RemoveGroupMemberResponse) SetStatusCode(v int32) *RemoveGroupMemberResponse {
	s.StatusCode = &v
	return s
}

type RemoveStoryFilesRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string                         `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Files   []*RemoveStoryFilesRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// This parameter is required.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s RemoveStoryFilesRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveStoryFilesRequest) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesRequest) SetDriveId(v string) *RemoveStoryFilesRequest {
	s.DriveId = &v
	return s
}

func (s *RemoveStoryFilesRequest) SetFiles(v []*RemoveStoryFilesRequestFiles) *RemoveStoryFilesRequest {
	s.Files = v
	return s
}

func (s *RemoveStoryFilesRequest) SetStoryId(v string) *RemoveStoryFilesRequest {
	s.StoryId = &v
	return s
}

type RemoveStoryFilesRequestFiles struct {
	// This parameter is required.
	//
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88dd06e49d9c0a14411ebae606f70edd9a59
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s RemoveStoryFilesRequestFiles) String() string {
	return tea.Prettify(s)
}

func (s RemoveStoryFilesRequestFiles) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesRequestFiles) SetFileId(v string) *RemoveStoryFilesRequestFiles {
	s.FileId = &v
	return s
}

func (s *RemoveStoryFilesRequestFiles) SetRevisionId(v string) *RemoveStoryFilesRequestFiles {
	s.RevisionId = &v
	return s
}

type RemoveStoryFilesResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s RemoveStoryFilesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveStoryFilesResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesResponseBody) SetDriveId(v string) *RemoveStoryFilesResponseBody {
	s.DriveId = &v
	return s
}

func (s *RemoveStoryFilesResponseBody) SetStoryId(v string) *RemoveStoryFilesResponseBody {
	s.StoryId = &v
	return s
}

type RemoveStoryFilesResponse struct {
	Headers    map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                        `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RemoveStoryFilesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RemoveStoryFilesResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveStoryFilesResponse) GoString() string {
	return s.String()
}

func (s *RemoveStoryFilesResponse) SetHeaders(v map[string]*string) *RemoveStoryFilesResponse {
	s.Headers = v
	return s
}

func (s *RemoveStoryFilesResponse) SetStatusCode(v int32) *RemoveStoryFilesResponse {
	s.StatusCode = &v
	return s
}

func (s *RemoveStoryFilesResponse) SetBody(v *RemoveStoryFilesResponseBody) *RemoveStoryFilesResponse {
	s.Body = v
	return s
}

type RestoreFileRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the file or folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43a255edc4463bffa6f5f5d317401
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s RestoreFileRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreFileRequest) GoString() string {
	return s.String()
}

func (s *RestoreFileRequest) SetDriveId(v string) *RestoreFileRequest {
	s.DriveId = &v
	return s
}

func (s *RestoreFileRequest) SetFileId(v string) *RestoreFileRequest {
	s.FileId = &v
	return s
}

type RestoreFileResponseBody struct {
	// The ID of the asynchronous task.
	//
	// If an empty string is returned, the file or folder is restored.
	//
	// If a non-empty string is returned, an asynchronous task is required. You can call the GetAsyncTask operation to obtain the information about an asynchronous task based on the task ID.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the file or folder.
	//
	// example:
	//
	// 4221bf6e6ab43a255edc4463bffa6f5f5d317401
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s RestoreFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RestoreFileResponseBody) GoString() string {
	return s.String()
}

func (s *RestoreFileResponseBody) SetAsyncTaskId(v string) *RestoreFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *RestoreFileResponseBody) SetDomainId(v string) *RestoreFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *RestoreFileResponseBody) SetDriveId(v string) *RestoreFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *RestoreFileResponseBody) SetFileId(v string) *RestoreFileResponseBody {
	s.FileId = &v
	return s
}

type RestoreFileResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *RestoreFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreFileResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreFileResponse) GoString() string {
	return s.String()
}

func (s *RestoreFileResponse) SetHeaders(v map[string]*string) *RestoreFileResponse {
	s.Headers = v
	return s
}

func (s *RestoreFileResponse) SetStatusCode(v int32) *RestoreFileResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreFileResponse) SetBody(v *RestoreFileResponseBody) *RestoreFileResponse {
	s.Body = v
	return s
}

type RestoreRevisionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// The version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40CB7794C929
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s RestoreRevisionRequest) String() string {
	return tea.Prettify(s)
}

func (s RestoreRevisionRequest) GoString() string {
	return s.String()
}

func (s *RestoreRevisionRequest) SetDriveId(v string) *RestoreRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *RestoreRevisionRequest) SetFileId(v string) *RestoreRevisionRequest {
	s.FileId = &v
	return s
}

func (s *RestoreRevisionRequest) SetRevisionId(v string) *RestoreRevisionRequest {
	s.RevisionId = &v
	return s
}

type RestoreRevisionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Revision          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s RestoreRevisionResponse) String() string {
	return tea.Prettify(s)
}

func (s RestoreRevisionResponse) GoString() string {
	return s.String()
}

func (s *RestoreRevisionResponse) SetHeaders(v map[string]*string) *RestoreRevisionResponse {
	s.Headers = v
	return s
}

func (s *RestoreRevisionResponse) SetStatusCode(v int32) *RestoreRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *RestoreRevisionResponse) SetBody(v *Revision) *RestoreRevisionResponse {
	s.Body = v
	return s
}

type ScanFileRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file properties to return.
	//
	// 	- If you want to return all file properties, set this parameter to \\*.
	//
	// 	- By default, if you do not specify this parameter, the following properties of a file are returned: - file_id, - drive_id, - parent_file_id, - type, - created_at, - updated_at, - file_extention, - size, - starred, - status, - category, and - permissions.
	//
	// 	- You can also specify properties to return. Separate multiple properties with commas (,).
	//
	// example:
	//
	// *
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ScanFileRequest) String() string {
	return tea.Prettify(s)
}

func (s ScanFileRequest) GoString() string {
	return s.String()
}

func (s *ScanFileRequest) SetDriveId(v string) *ScanFileRequest {
	s.DriveId = &v
	return s
}

func (s *ScanFileRequest) SetFields(v string) *ScanFileRequest {
	s.Fields = &v
	return s
}

func (s *ScanFileRequest) SetLimit(v int32) *ScanFileRequest {
	s.Limit = &v
	return s
}

func (s *ScanFileRequest) SetMarker(v string) *ScanFileRequest {
	s.Marker = &v
	return s
}

type ScanFileResponseBody struct {
	// The information about the files.
	Items []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s ScanFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ScanFileResponseBody) GoString() string {
	return s.String()
}

func (s *ScanFileResponseBody) SetItems(v []*File) *ScanFileResponseBody {
	s.Items = v
	return s
}

func (s *ScanFileResponseBody) SetNextMarker(v string) *ScanFileResponseBody {
	s.NextMarker = &v
	return s
}

type ScanFileResponse struct {
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ScanFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s ScanFileResponse) String() string {
	return tea.Prettify(s)
}

func (s ScanFileResponse) GoString() string {
	return s.String()
}

func (s *ScanFileResponse) SetHeaders(v map[string]*string) *ScanFileResponse {
	s.Headers = v
	return s
}

func (s *ScanFileResponse) SetStatusCode(v int32) *ScanFileResponse {
	s.StatusCode = &v
	return s
}

func (s *ScanFileResponse) SetBody(v *ScanFileResponseBody) *ScanFileResponse {
	s.Body = v
	return s
}

type SearchAddressGroupsRequest struct {
	// The level of the location.
	//
	// Valid values:
	//
	// 	- country
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- province
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- city
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- district
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	// 	- township
	//
	//     <!-- -->
	//
	//     <!-- -->
	//
	//     <!-- -->
	AddressLevel *string `json:"address_level,omitempty" xml:"address_level,omitempty"`
	// The locations.
	AddressNames []*string `json:"address_names,omitempty" xml:"address_names,omitempty" type:"Repeated"`
	// The coordinates of the bottom right vertex of the rectangle. Set the value in the format of latitude,longitude.
	//
	// example:
	//
	// 40.121,105.2121
	BrGeoPoint *string `json:"br_geo_point,omitempty" xml:"br_geo_point,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The method used to generate the thumbnail of an image.
	//
	// example:
	//
	// image/resize,w_200
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// The coordinates of the top left vertex of the rectangle. Set the value in the format of latitude,longitude.
	//
	// example:
	//
	// 39.121,101.2121
	TlGeoPoint *string `json:"tl_geo_point,omitempty" xml:"tl_geo_point,omitempty"`
	// The method used to generate the thumbnail of a video.
	//
	// example:
	//
	// video/snapshot,t_7000,f_jpg,w_800,h_600,m_fast
	VideoThumbnailProcess *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
}

func (s SearchAddressGroupsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchAddressGroupsRequest) GoString() string {
	return s.String()
}

func (s *SearchAddressGroupsRequest) SetAddressLevel(v string) *SearchAddressGroupsRequest {
	s.AddressLevel = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetAddressNames(v []*string) *SearchAddressGroupsRequest {
	s.AddressNames = v
	return s
}

func (s *SearchAddressGroupsRequest) SetBrGeoPoint(v string) *SearchAddressGroupsRequest {
	s.BrGeoPoint = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetDriveId(v string) *SearchAddressGroupsRequest {
	s.DriveId = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetImageThumbnailProcess(v string) *SearchAddressGroupsRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetTlGeoPoint(v string) *SearchAddressGroupsRequest {
	s.TlGeoPoint = &v
	return s
}

func (s *SearchAddressGroupsRequest) SetVideoThumbnailProcess(v string) *SearchAddressGroupsRequest {
	s.VideoThumbnailProcess = &v
	return s
}

type SearchAddressGroupsResponseBody struct {
	// The location-based groups.
	Items []*AddressGroup `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s SearchAddressGroupsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchAddressGroupsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchAddressGroupsResponseBody) SetItems(v []*AddressGroup) *SearchAddressGroupsResponseBody {
	s.Items = v
	return s
}

type SearchAddressGroupsResponse struct {
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchAddressGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchAddressGroupsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchAddressGroupsResponse) GoString() string {
	return s.String()
}

func (s *SearchAddressGroupsResponse) SetHeaders(v map[string]*string) *SearchAddressGroupsResponse {
	s.Headers = v
	return s
}

func (s *SearchAddressGroupsResponse) SetStatusCode(v int32) *SearchAddressGroupsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchAddressGroupsResponse) SetBody(v *SearchAddressGroupsResponseBody) *SearchAddressGroupsResponse {
	s.Body = v
	return s
}

type SearchDomainsRequest struct {
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The name of the domain. Fuzzy search is supported.
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The sorting rule. Set the value to created_at, which specifies that the results are sorted based on the time when the domain was created.
	//
	// example:
	//
	// created_at
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
}

func (s SearchDomainsRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDomainsRequest) GoString() string {
	return s.String()
}

func (s *SearchDomainsRequest) SetLimit(v int64) *SearchDomainsRequest {
	s.Limit = &v
	return s
}

func (s *SearchDomainsRequest) SetMarker(v string) *SearchDomainsRequest {
	s.Marker = &v
	return s
}

func (s *SearchDomainsRequest) SetName(v string) *SearchDomainsRequest {
	s.Name = &v
	return s
}

func (s *SearchDomainsRequest) SetOrderBy(v string) *SearchDomainsRequest {
	s.OrderBy = &v
	return s
}

type SearchDomainsResponseBody struct {
	// The queried domains.
	Items []*Domain `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s SearchDomainsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchDomainsResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDomainsResponseBody) SetItems(v []*Domain) *SearchDomainsResponseBody {
	s.Items = v
	return s
}

func (s *SearchDomainsResponseBody) SetNextMarker(v string) *SearchDomainsResponseBody {
	s.NextMarker = &v
	return s
}

type SearchDomainsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDomainsResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDomainsResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchDomainsResponse) GoString() string {
	return s.String()
}

func (s *SearchDomainsResponse) SetHeaders(v map[string]*string) *SearchDomainsResponse {
	s.Headers = v
	return s
}

func (s *SearchDomainsResponse) SetStatusCode(v int32) *SearchDomainsResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDomainsResponse) SetBody(v *SearchDomainsResponseBody) *SearchDomainsResponse {
	s.Body = v
	return s
}

type SearchDriveRequest struct {
	// The drive name.
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// The maximum number of asynchronous tasks to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The owner of the drive.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The type of the owner. Valid values:
	//
	// user group
	//
	// example:
	//
	// user
	OwnerType *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
}

func (s SearchDriveRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchDriveRequest) GoString() string {
	return s.String()
}

func (s *SearchDriveRequest) SetDriveName(v string) *SearchDriveRequest {
	s.DriveName = &v
	return s
}

func (s *SearchDriveRequest) SetLimit(v int32) *SearchDriveRequest {
	s.Limit = &v
	return s
}

func (s *SearchDriveRequest) SetMarker(v string) *SearchDriveRequest {
	s.Marker = &v
	return s
}

func (s *SearchDriveRequest) SetOwner(v string) *SearchDriveRequest {
	s.Owner = &v
	return s
}

func (s *SearchDriveRequest) SetOwnerType(v string) *SearchDriveRequest {
	s.OwnerType = &v
	return s
}

type SearchDriveResponseBody struct {
	// The information about the drives.
	Items []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s SearchDriveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchDriveResponseBody) GoString() string {
	return s.String()
}

func (s *SearchDriveResponseBody) SetItems(v []*Drive) *SearchDriveResponseBody {
	s.Items = v
	return s
}

func (s *SearchDriveResponseBody) SetNextMarker(v string) *SearchDriveResponseBody {
	s.NextMarker = &v
	return s
}

type SearchDriveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchDriveResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchDriveResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchDriveResponse) GoString() string {
	return s.String()
}

func (s *SearchDriveResponse) SetHeaders(v map[string]*string) *SearchDriveResponse {
	s.Headers = v
	return s
}

func (s *SearchDriveResponse) SetStatusCode(v int32) *SearchDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchDriveResponse) SetBody(v *SearchDriveResponseBody) *SearchDriveResponse {
	s.Body = v
	return s
}

type SearchFileRequest struct {
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Deprecated
	//
	// example:
	//
	// url,thumbnail
	Fields *string `json:"fields,omitempty" xml:"fields,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The field by which to sort the returned results. Default value: created_at. Valid values:
	//
	// 	- created_at: sorts the results by the time when the file was created.
	//
	// 	- updated_at: sorts the results by the time when the file was modified.
	//
	// 	- size: sorts the results by the size of the file.
	//
	// 	- name: sorts the results by the name of the file.
	//
	// The order in which you want to sort the returned results. Valid values:
	//
	// 	- ASC: sorts the results in ascending order.
	//
	// 	- DESC: sorts the results in descending order.
	//
	// You must specify this parameter in the \\<field name> \\<ASC or DESC> format. Separate multiple field names with commas (,). A preceding field has a higher priority than a following field. Examples:
	//
	// 	- If you want to sort the results based on the file name in ascending order, set this parameter to "name ASC".
	//
	// 	- If you want to sort the results based on the creation time in descending order, set this parameter to "created_at DESC".
	//
	// 	- If you want to sort the results based on the creation time in descending order first, and then sort the results based on the file name in ascending order if the creation time is the same, set this parameter to "created_at DESC,name ASC".
	//
	// example:
	//
	// name
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// The search condition. Fuzzy searches based on the file name or directory name are supported. The search condition can be up to 4,096 characters in length.
	//
	// This parameter is required.
	//
	// example:
	//
	// not name=123
	Query     *string `json:"query,omitempty" xml:"query,omitempty"`
	Recursive *bool   `json:"recursive,omitempty" xml:"recursive,omitempty"`
	// Specifies whether to return the total number of retrieved files.
	//
	// example:
	//
	// true
	ReturnTotalCount *bool `json:"return_total_count,omitempty" xml:"return_total_count,omitempty"`
}

func (s SearchFileRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchFileRequest) GoString() string {
	return s.String()
}

func (s *SearchFileRequest) SetDriveId(v string) *SearchFileRequest {
	s.DriveId = &v
	return s
}

func (s *SearchFileRequest) SetFields(v string) *SearchFileRequest {
	s.Fields = &v
	return s
}

func (s *SearchFileRequest) SetLimit(v int32) *SearchFileRequest {
	s.Limit = &v
	return s
}

func (s *SearchFileRequest) SetMarker(v string) *SearchFileRequest {
	s.Marker = &v
	return s
}

func (s *SearchFileRequest) SetOrderBy(v string) *SearchFileRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchFileRequest) SetQuery(v string) *SearchFileRequest {
	s.Query = &v
	return s
}

func (s *SearchFileRequest) SetRecursive(v bool) *SearchFileRequest {
	s.Recursive = &v
	return s
}

func (s *SearchFileRequest) SetReturnTotalCount(v bool) *SearchFileRequest {
	s.ReturnTotalCount = &v
	return s
}

type SearchFileResponseBody struct {
	// The information about the files.
	Items []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	// The total number of retrieved files.
	//
	// example:
	//
	// 1022
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s SearchFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchFileResponseBody) GoString() string {
	return s.String()
}

func (s *SearchFileResponseBody) SetItems(v []*File) *SearchFileResponseBody {
	s.Items = v
	return s
}

func (s *SearchFileResponseBody) SetNextMarker(v string) *SearchFileResponseBody {
	s.NextMarker = &v
	return s
}

func (s *SearchFileResponseBody) SetTotalCount(v int64) *SearchFileResponseBody {
	s.TotalCount = &v
	return s
}

type SearchFileResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchFileResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchFileResponse) GoString() string {
	return s.String()
}

func (s *SearchFileResponse) SetHeaders(v map[string]*string) *SearchFileResponse {
	s.Headers = v
	return s
}

func (s *SearchFileResponse) SetStatusCode(v int32) *SearchFileResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchFileResponse) SetBody(v *SearchFileResponseBody) *SearchFileResponse {
	s.Body = v
	return s
}

type SearchShareLinkRequest struct {
	// The creators of shares. Set this parameter to a user ID. If you do not specify this parameter, Drive and Photo Service automatically queries shares based on your permissions. If you are a drive administrator or the super administrator, the shares created by all members are queried. If you are a team administrator, the shares created by all team members are queried. If you are a common user, only the shares created by yourself are queried.
	//
	// If you specify this parameter, this parameter must be set to the ID of the super administrator, a drive administrator, or a team administrator.
	Creators []*string `json:"creators,omitempty" xml:"creators,omitempty" type:"Repeated"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// The number of returned results must be less than or equal to the specified number.
	//
	// example:
	//
	// 50
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The field by which to sort the returned results. Default value: created_at. Valid values:
	//
	// 	- share_name: sorts the results by the name of the share.
	//
	// 	- updated_at: sorts the results by the time when the share was modified.
	//
	// 	- description: sorts the results by the description of the share.
	//
	// 	- created_at: sorts the results by the time when the share was created.
	//
	// example:
	//
	// created_at
	OrderBy *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	// The order in which you want to sort the returned results. By default, order_direction is set to DESC if order_by is set to created_at or updated_at, and is set to ASC if order_by is set to other values. Valid values:
	//
	// 	- ASC: sorts the results in ascending order.
	//
	// 	- DESC: sorts the results in descending order.
	//
	// example:
	//
	// ASC
	OrderDirection *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
	// The query condition that is used to search for share URLs. You can use the following fields to specify query conditions: created_at: queries a share URL based on the time when the share URL was created. updated_at: queries a share URL based on the time when the share URL was modified. share_name_for_fuzzy: queries a share URL based on the name of the share. A fuzzy match is supported. status: queries a share URL based on the status of the share. Valid values: enabled and disabled. A value of enabled indicates that the share is valid. A value of disabled indicates that the share is canceled. expired_time: queries a share URL based on the expiration time of the share. If the share never expires, set this field to 1970-01-01T00:00:00. Otherwise, set this field to 1970-01-02T00:00:00.
	//
	// example:
	//
	// created_at>=\\"2022-01-18T02:50:00\\" and created_at<\\"2022-01-19T09:45:28\\" and share_name_for_fuzzy match \\"HD.mp4\\" and status in [\\"enabled\\", \\"disabled\\"] and expired_time=\\"1970-01-01T00:00:00\\"
	Query *string `json:"query,omitempty" xml:"query,omitempty"`
	// Specifies whether to return the total number of returned results.
	//
	// example:
	//
	// false
	ReturnTotalCount *bool `json:"return_total_count,omitempty" xml:"return_total_count,omitempty"`
}

func (s SearchShareLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchShareLinkRequest) GoString() string {
	return s.String()
}

func (s *SearchShareLinkRequest) SetCreators(v []*string) *SearchShareLinkRequest {
	s.Creators = v
	return s
}

func (s *SearchShareLinkRequest) SetLimit(v int32) *SearchShareLinkRequest {
	s.Limit = &v
	return s
}

func (s *SearchShareLinkRequest) SetMarker(v string) *SearchShareLinkRequest {
	s.Marker = &v
	return s
}

func (s *SearchShareLinkRequest) SetOrderBy(v string) *SearchShareLinkRequest {
	s.OrderBy = &v
	return s
}

func (s *SearchShareLinkRequest) SetOrderDirection(v string) *SearchShareLinkRequest {
	s.OrderDirection = &v
	return s
}

func (s *SearchShareLinkRequest) SetQuery(v string) *SearchShareLinkRequest {
	s.Query = &v
	return s
}

func (s *SearchShareLinkRequest) SetReturnTotalCount(v bool) *SearchShareLinkRequest {
	s.ReturnTotalCount = &v
	return s
}

type SearchShareLinkResponseBody struct {
	// The share URLs.
	Items []*ShareLink `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 101
	TotalCount *int64 `json:"total_count,omitempty" xml:"total_count,omitempty"`
}

func (s SearchShareLinkResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchShareLinkResponseBody) GoString() string {
	return s.String()
}

func (s *SearchShareLinkResponseBody) SetItems(v []*ShareLink) *SearchShareLinkResponseBody {
	s.Items = v
	return s
}

func (s *SearchShareLinkResponseBody) SetNextMarker(v string) *SearchShareLinkResponseBody {
	s.NextMarker = &v
	return s
}

func (s *SearchShareLinkResponseBody) SetTotalCount(v int64) *SearchShareLinkResponseBody {
	s.TotalCount = &v
	return s
}

type SearchShareLinkResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchShareLinkResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchShareLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchShareLinkResponse) GoString() string {
	return s.String()
}

func (s *SearchShareLinkResponse) SetHeaders(v map[string]*string) *SearchShareLinkResponse {
	s.Headers = v
	return s
}

func (s *SearchShareLinkResponse) SetStatusCode(v int32) *SearchShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchShareLinkResponse) SetBody(v *SearchShareLinkResponseBody) *SearchShareLinkResponse {
	s.Body = v
	return s
}

type SearchSimilarImageClustersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	// example:
	//
	// 50
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// YWRzX3VzZXJfcHJvZmlsZV9je1bnQh***
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
}

func (s SearchSimilarImageClustersRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarImageClustersRequest) GoString() string {
	return s.String()
}

func (s *SearchSimilarImageClustersRequest) SetDriveId(v string) *SearchSimilarImageClustersRequest {
	s.DriveId = &v
	return s
}

func (s *SearchSimilarImageClustersRequest) SetImageThumbnailProcess(v string) *SearchSimilarImageClustersRequest {
	s.ImageThumbnailProcess = &v
	return s
}

func (s *SearchSimilarImageClustersRequest) SetLimit(v int64) *SearchSimilarImageClustersRequest {
	s.Limit = &v
	return s
}

func (s *SearchSimilarImageClustersRequest) SetMarker(v string) *SearchSimilarImageClustersRequest {
	s.Marker = &v
	return s
}

func (s *SearchSimilarImageClustersRequest) SetOrder(v string) *SearchSimilarImageClustersRequest {
	s.Order = &v
	return s
}

type SearchSimilarImageClustersResponseBody struct {
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0***
	NextMarker           *string                                                       `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	SimilarImageClusters []*SearchSimilarImageClustersResponseBodySimilarImageClusters `json:"similar_image_clusters,omitempty" xml:"similar_image_clusters,omitempty" type:"Repeated"`
}

func (s SearchSimilarImageClustersResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarImageClustersResponseBody) GoString() string {
	return s.String()
}

func (s *SearchSimilarImageClustersResponseBody) SetNextMarker(v string) *SearchSimilarImageClustersResponseBody {
	s.NextMarker = &v
	return s
}

func (s *SearchSimilarImageClustersResponseBody) SetSimilarImageClusters(v []*SearchSimilarImageClustersResponseBodySimilarImageClusters) *SearchSimilarImageClustersResponseBody {
	s.SimilarImageClusters = v
	return s
}

type SearchSimilarImageClustersResponseBodySimilarImageClusters struct {
	Files []*File `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
}

func (s SearchSimilarImageClustersResponseBodySimilarImageClusters) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarImageClustersResponseBodySimilarImageClusters) GoString() string {
	return s.String()
}

func (s *SearchSimilarImageClustersResponseBodySimilarImageClusters) SetFiles(v []*File) *SearchSimilarImageClustersResponseBodySimilarImageClusters {
	s.Files = v
	return s
}

type SearchSimilarImageClustersResponse struct {
	Headers    map[string]*string                      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchSimilarImageClustersResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchSimilarImageClustersResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchSimilarImageClustersResponse) GoString() string {
	return s.String()
}

func (s *SearchSimilarImageClustersResponse) SetHeaders(v map[string]*string) *SearchSimilarImageClustersResponse {
	s.Headers = v
	return s
}

func (s *SearchSimilarImageClustersResponse) SetStatusCode(v int32) *SearchSimilarImageClustersResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchSimilarImageClustersResponse) SetBody(v *SearchSimilarImageClustersResponseBody) *SearchSimilarImageClustersResponse {
	s.Body = v
	return s
}

type SearchStoriesRequest struct {
	// Deprecated
	//
	// example:
	//
	// image/resize,m_fill,h_128,w_128,limit_0/format,jpg
	CoverImageThumbnailProcess *string `json:"cover_image_thumbnail_process,omitempty" xml:"cover_image_thumbnail_process,omitempty"`
	// Deprecated
	//
	// example:
	//
	// video/snapshot,t_1000,f_jpg,w_0,h_0,m_fast,ar_auto
	CoverVideoThumbnailProcess *string `json:"cover_video_thumbnail_process,omitempty" xml:"cover_video_thumbnail_process,omitempty"`
	// if can be null:
	// true
	CreateTimeRange *SearchStoriesRequestCreateTimeRange `json:"create_time_range,omitempty" xml:"create_time_range,omitempty" type:"Struct"`
	// Deprecated
	//
	// example:
	//
	// key1=value1,key2!=value2
	CustomLabels *string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId      *string   `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FaceGroupIds []*string `json:"face_group_ids,omitempty" xml:"face_group_ids,omitempty" type:"Repeated"`
	// example:
	//
	// 100
	Limit *int64 `json:"limit,omitempty" xml:"limit,omitempty"`
	// example:
	//
	// NWQ1Yjk4YmI1ZDODBhNDQ2Nzhl***
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// example:
	//
	// desc
	Order *string `json:"order,omitempty" xml:"order,omitempty"`
	// example:
	//
	// CreateTime
	Sort *string `json:"sort,omitempty" xml:"sort,omitempty"`
	// if can be null:
	// true
	StoryEndTimeRange *SearchStoriesRequestStoryEndTimeRange `json:"story_end_time_range,omitempty" xml:"story_end_time_range,omitempty" type:"Struct"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId   *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
	StoryName *string `json:"story_name,omitempty" xml:"story_name,omitempty"`
	// if can be null:
	// true
	StoryStartTimeRange *SearchStoriesRequestStoryStartTimeRange `json:"story_start_time_range,omitempty" xml:"story_start_time_range,omitempty" type:"Struct"`
	// example:
	//
	// PeopleMemory
	StoryType *string `json:"story_type,omitempty" xml:"story_type,omitempty"`
	// Deprecated
	//
	// example:
	//
	// 900
	UrlExpireSec *int64 `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// false
	WithEmptyStories *bool `json:"with_empty_stories,omitempty" xml:"with_empty_stories,omitempty"`
}

func (s SearchStoriesRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchStoriesRequest) GoString() string {
	return s.String()
}

func (s *SearchStoriesRequest) SetCoverImageThumbnailProcess(v string) *SearchStoriesRequest {
	s.CoverImageThumbnailProcess = &v
	return s
}

func (s *SearchStoriesRequest) SetCoverVideoThumbnailProcess(v string) *SearchStoriesRequest {
	s.CoverVideoThumbnailProcess = &v
	return s
}

func (s *SearchStoriesRequest) SetCreateTimeRange(v *SearchStoriesRequestCreateTimeRange) *SearchStoriesRequest {
	s.CreateTimeRange = v
	return s
}

func (s *SearchStoriesRequest) SetCustomLabels(v string) *SearchStoriesRequest {
	s.CustomLabels = &v
	return s
}

func (s *SearchStoriesRequest) SetDriveId(v string) *SearchStoriesRequest {
	s.DriveId = &v
	return s
}

func (s *SearchStoriesRequest) SetFaceGroupIds(v []*string) *SearchStoriesRequest {
	s.FaceGroupIds = v
	return s
}

func (s *SearchStoriesRequest) SetLimit(v int64) *SearchStoriesRequest {
	s.Limit = &v
	return s
}

func (s *SearchStoriesRequest) SetMarker(v string) *SearchStoriesRequest {
	s.Marker = &v
	return s
}

func (s *SearchStoriesRequest) SetOrder(v string) *SearchStoriesRequest {
	s.Order = &v
	return s
}

func (s *SearchStoriesRequest) SetSort(v string) *SearchStoriesRequest {
	s.Sort = &v
	return s
}

func (s *SearchStoriesRequest) SetStoryEndTimeRange(v *SearchStoriesRequestStoryEndTimeRange) *SearchStoriesRequest {
	s.StoryEndTimeRange = v
	return s
}

func (s *SearchStoriesRequest) SetStoryId(v string) *SearchStoriesRequest {
	s.StoryId = &v
	return s
}

func (s *SearchStoriesRequest) SetStoryName(v string) *SearchStoriesRequest {
	s.StoryName = &v
	return s
}

func (s *SearchStoriesRequest) SetStoryStartTimeRange(v *SearchStoriesRequestStoryStartTimeRange) *SearchStoriesRequest {
	s.StoryStartTimeRange = v
	return s
}

func (s *SearchStoriesRequest) SetStoryType(v string) *SearchStoriesRequest {
	s.StoryType = &v
	return s
}

func (s *SearchStoriesRequest) SetUrlExpireSec(v int64) *SearchStoriesRequest {
	s.UrlExpireSec = &v
	return s
}

func (s *SearchStoriesRequest) SetWithEmptyStories(v bool) *SearchStoriesRequest {
	s.WithEmptyStories = &v
	return s
}

type SearchStoriesRequestCreateTimeRange struct {
	// example:
	//
	// 2022-12-31T00:00:00+08:00
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 2016-12-31T00:00:00+08:00
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s SearchStoriesRequestCreateTimeRange) String() string {
	return tea.Prettify(s)
}

func (s SearchStoriesRequestCreateTimeRange) GoString() string {
	return s.String()
}

func (s *SearchStoriesRequestCreateTimeRange) SetEnd(v string) *SearchStoriesRequestCreateTimeRange {
	s.End = &v
	return s
}

func (s *SearchStoriesRequestCreateTimeRange) SetStart(v string) *SearchStoriesRequestCreateTimeRange {
	s.Start = &v
	return s
}

type SearchStoriesRequestStoryEndTimeRange struct {
	// example:
	//
	// 2022-12-31T00:00:00+08:00
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 2016-12-31T00:00:00+08:00
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s SearchStoriesRequestStoryEndTimeRange) String() string {
	return tea.Prettify(s)
}

func (s SearchStoriesRequestStoryEndTimeRange) GoString() string {
	return s.String()
}

func (s *SearchStoriesRequestStoryEndTimeRange) SetEnd(v string) *SearchStoriesRequestStoryEndTimeRange {
	s.End = &v
	return s
}

func (s *SearchStoriesRequestStoryEndTimeRange) SetStart(v string) *SearchStoriesRequestStoryEndTimeRange {
	s.Start = &v
	return s
}

type SearchStoriesRequestStoryStartTimeRange struct {
	// example:
	//
	// 2022-12-31T00:00:00+08:00
	End *string `json:"end,omitempty" xml:"end,omitempty"`
	// example:
	//
	// 2016-12-31T00:00:00+08:00
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s SearchStoriesRequestStoryStartTimeRange) String() string {
	return tea.Prettify(s)
}

func (s SearchStoriesRequestStoryStartTimeRange) GoString() string {
	return s.String()
}

func (s *SearchStoriesRequestStoryStartTimeRange) SetEnd(v string) *SearchStoriesRequestStoryStartTimeRange {
	s.End = &v
	return s
}

func (s *SearchStoriesRequestStoryStartTimeRange) SetStart(v string) *SearchStoriesRequestStoryStartTimeRange {
	s.Start = &v
	return s
}

type SearchStoriesResponseBody struct {
	Items []*Story `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// example:
	//
	// Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJ***
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s SearchStoriesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchStoriesResponseBody) GoString() string {
	return s.String()
}

func (s *SearchStoriesResponseBody) SetItems(v []*Story) *SearchStoriesResponseBody {
	s.Items = v
	return s
}

func (s *SearchStoriesResponseBody) SetNextMarker(v string) *SearchStoriesResponseBody {
	s.NextMarker = &v
	return s
}

type SearchStoriesResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchStoriesResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchStoriesResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchStoriesResponse) GoString() string {
	return s.String()
}

func (s *SearchStoriesResponse) SetHeaders(v map[string]*string) *SearchStoriesResponse {
	s.Headers = v
	return s
}

func (s *SearchStoriesResponse) SetStatusCode(v int32) *SearchStoriesResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchStoriesResponse) SetBody(v *SearchStoriesResponseBody) *SearchStoriesResponse {
	s.Body = v
	return s
}

type SearchUserRequest struct {
	// The email address of the user.
	//
	// example:
	//
	// 123@pds.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The maximum number of results to return. Valid values: 1 to 100. Default value: 100.
	//
	// example:
	//
	// 100
	Limit *int32 `json:"limit,omitempty" xml:"limit,omitempty"`
	// The pagination token that is used in the next request to retrieve a new page of results. You do not need to specify this parameter for the first request. You must specify the token that is obtained from the previous query as the value of marker.\\
	//
	// By default, this parameter is left empty.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
	// The nickname of the user. The nickname can be up to 128 characters in length.
	//
	// example:
	//
	// pdsuer
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The nickname used for fuzzy searches. The nickname can be up to 128 characters in length.
	//
	// example:
	//
	// la
	NickNameForFuzzy *string `json:"nick_name_for_fuzzy,omitempty" xml:"nick_name_for_fuzzy,omitempty"`
	// The mobile number of the user.
	//
	// example:
	//
	// 13900001111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The role of the user. Valid values:
	//
	// 	- superadmin
	//
	// 	- admin
	//
	// 	- user
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The state of the user. Valid values:
	//
	// 	- disabled: The user is prohibited from logon.
	//
	// 	- enabled: The user is in a normal state.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The name of the user. The name can be up to 128 characters in length.
	//
	// example:
	//
	// pds
	UserName *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
}

func (s SearchUserRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchUserRequest) GoString() string {
	return s.String()
}

func (s *SearchUserRequest) SetEmail(v string) *SearchUserRequest {
	s.Email = &v
	return s
}

func (s *SearchUserRequest) SetLimit(v int32) *SearchUserRequest {
	s.Limit = &v
	return s
}

func (s *SearchUserRequest) SetMarker(v string) *SearchUserRequest {
	s.Marker = &v
	return s
}

func (s *SearchUserRequest) SetNickName(v string) *SearchUserRequest {
	s.NickName = &v
	return s
}

func (s *SearchUserRequest) SetNickNameForFuzzy(v string) *SearchUserRequest {
	s.NickNameForFuzzy = &v
	return s
}

func (s *SearchUserRequest) SetPhone(v string) *SearchUserRequest {
	s.Phone = &v
	return s
}

func (s *SearchUserRequest) SetRole(v string) *SearchUserRequest {
	s.Role = &v
	return s
}

func (s *SearchUserRequest) SetStatus(v string) *SearchUserRequest {
	s.Status = &v
	return s
}

func (s *SearchUserRequest) SetUserName(v string) *SearchUserRequest {
	s.UserName = &v
	return s
}

type SearchUserResponseBody struct {
	// The information about the users.
	Items []*User `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// A pagination token. It can be used in the next request to retrieve a new page of results. If next_marker is empty, no next page exists.
	//
	// example:
	//
	// NWQ1Yjk4YmI1ZDRlYmU1Y2E0YWE0NmJhYWJmODBhNDQ2NzhlMTRhMg
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
}

func (s SearchUserResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchUserResponseBody) GoString() string {
	return s.String()
}

func (s *SearchUserResponseBody) SetItems(v []*User) *SearchUserResponseBody {
	s.Items = v
	return s
}

func (s *SearchUserResponseBody) SetNextMarker(v string) *SearchUserResponseBody {
	s.NextMarker = &v
	return s
}

type SearchUserResponse struct {
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *SearchUserResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s SearchUserResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchUserResponse) GoString() string {
	return s.String()
}

func (s *SearchUserResponse) SetHeaders(v map[string]*string) *SearchUserResponse {
	s.Headers = v
	return s
}

func (s *SearchUserResponse) SetStatusCode(v int32) *SearchUserResponse {
	s.StatusCode = &v
	return s
}

func (s *SearchUserResponse) SetBody(v *SearchUserResponseBody) *SearchUserResponse {
	s.Body = v
	return s
}

type TokenRequest struct {
	// The JWT assertion that is signed by using the JWT private key. The JWT assertion contains the information about the user to be authorized and the authorization parameters. For more information about the structure of the JWT assertion, see JWTPayload. This parameter is required if grant_type is set to urn:ietf:params:oauth:grant-type:jwt-bearer.
	//
	// example:
	//
	// ey***asd
	Assertion *string `json:"assertion,omitempty" xml:"assertion,omitempty"`
	// The AppId of the application that is created in the Drive and Photo Service console.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1Zu***flH
	ClientId *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	// The AppSecret of the application that is created in the Drive and Photo Service console. This parameter is required if the application is of the WebServer type.
	//
	// example:
	//
	// 80D***3i5
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	// The authorization code in the redirect URI that is specified after the authorization process is complete. This parameter is required if grant_type is set to authorization_code.
	//
	// example:
	//
	// 0045157fa8e24f4f9a0d9e3ff158c1e0
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The method that is used to generate an access token. Valid values:
	//
	// authorization_code: generates an access token by using the authorization code that is returned after the authorization process is complete.
	//
	// refresh_token: generates an access token by using the refresh token that is returned after the authorization process is complete.
	//
	// urn:ietf:params:oauth:grant-type:jwt-bearer: generates an access token by using a JWT assertion.
	//
	// This parameter is required.
	//
	// example:
	//
	// refresh_token
	GrantType *string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	// The redirect URI that is specified when you initiate the authorization request. This parameter is required if grant_type is set to authorization_code.
	//
	// example:
	//
	// https://aliyun.com/pds
	RedirectUri *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	// The refresh token that is used to refresh the access token. This parameter is required if grant_type is set to refresh_token.
	//
	// example:
	//
	// 399623e13353490391266c7d48a13ed1
	RefreshToken *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
}

func (s TokenRequest) String() string {
	return tea.Prettify(s)
}

func (s TokenRequest) GoString() string {
	return s.String()
}

func (s *TokenRequest) SetAssertion(v string) *TokenRequest {
	s.Assertion = &v
	return s
}

func (s *TokenRequest) SetClientId(v string) *TokenRequest {
	s.ClientId = &v
	return s
}

func (s *TokenRequest) SetClientSecret(v string) *TokenRequest {
	s.ClientSecret = &v
	return s
}

func (s *TokenRequest) SetCode(v string) *TokenRequest {
	s.Code = &v
	return s
}

func (s *TokenRequest) SetGrantType(v string) *TokenRequest {
	s.GrantType = &v
	return s
}

func (s *TokenRequest) SetRedirectUri(v string) *TokenRequest {
	s.RedirectUri = &v
	return s
}

func (s *TokenRequest) SetRefreshToken(v string) *TokenRequest {
	s.RefreshToken = &v
	return s
}

type TokenResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Token             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TokenResponse) String() string {
	return tea.Prettify(s)
}

func (s TokenResponse) GoString() string {
	return s.String()
}

func (s *TokenResponse) SetHeaders(v map[string]*string) *TokenResponse {
	s.Headers = v
	return s
}

func (s *TokenResponse) SetStatusCode(v int32) *TokenResponse {
	s.StatusCode = &v
	return s
}

func (s *TokenResponse) SetBody(v *Token) *TokenResponse {
	s.Body = v
	return s
}

type TrashFileRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the file or folder.
	//
	// This parameter is required.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s TrashFileRequest) String() string {
	return tea.Prettify(s)
}

func (s TrashFileRequest) GoString() string {
	return s.String()
}

func (s *TrashFileRequest) SetDriveId(v string) *TrashFileRequest {
	s.DriveId = &v
	return s
}

func (s *TrashFileRequest) SetFileId(v string) *TrashFileRequest {
	s.FileId = &v
	return s
}

type TrashFileResponseBody struct {
	// The ID of the asynchronous task.
	//
	// If an empty string is returned, the file or folder is moved to the recycle bin.
	//
	// If a non-empty string is returned, an asynchronous task is required. You can call the GetAsyncTask operation to obtain the information about an asynchronous task based on the task ID.
	//
	// example:
	//
	// 13ebd3a24dba4166b1527add676ef2866051b4d5dele16
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	// The domain ID.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The ID of the file or folder.
	//
	// example:
	//
	// 4221bf6e6ab43c255edc4463bf3a6f5f5d317406
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
}

func (s TrashFileResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TrashFileResponseBody) GoString() string {
	return s.String()
}

func (s *TrashFileResponseBody) SetAsyncTaskId(v string) *TrashFileResponseBody {
	s.AsyncTaskId = &v
	return s
}

func (s *TrashFileResponseBody) SetDomainId(v string) *TrashFileResponseBody {
	s.DomainId = &v
	return s
}

func (s *TrashFileResponseBody) SetDriveId(v string) *TrashFileResponseBody {
	s.DriveId = &v
	return s
}

func (s *TrashFileResponseBody) SetFileId(v string) *TrashFileResponseBody {
	s.FileId = &v
	return s
}

type TrashFileResponse struct {
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *TrashFileResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s TrashFileResponse) String() string {
	return tea.Prettify(s)
}

func (s TrashFileResponse) GoString() string {
	return s.String()
}

func (s *TrashFileResponse) SetHeaders(v map[string]*string) *TrashFileResponse {
	s.Headers = v
	return s
}

func (s *TrashFileResponse) SetStatusCode(v int32) *TrashFileResponse {
	s.StatusCode = &v
	return s
}

func (s *TrashFileResponse) SetBody(v *TrashFileResponseBody) *TrashFileResponse {
	s.Body = v
	return s
}

type UnLinkAccountRequest struct {
	// Additional information for the unique account identifier. For example, when the account is a phone number, this field should be filled with the area code of the phone, such as 86 for Mainland China. If not provided, it defaults to 86.
	//
	// example:
	//
	// 1
	Extra *string `json:"extra,omitempty" xml:"extra,omitempty"`
	// Unique identifier of the account, such as a phone number
	//
	// This parameter is required.
	//
	// example:
	//
	// 139****
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	// Account type
	//
	// mobile: Phone number
	//
	// email: Email address
	//
	// ding: DingTalk
	//
	// ram: Alibaba Cloud RAM User
	//
	// wechat: WeCom
	//
	// ldap: LDAP account
	//
	// custom: Custom account
	//
	// This parameter is required.
	//
	// example:
	//
	// mobile
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// User identifier
	//
	// This parameter is required.
	//
	// example:
	//
	// uid1
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s UnLinkAccountRequest) String() string {
	return tea.Prettify(s)
}

func (s UnLinkAccountRequest) GoString() string {
	return s.String()
}

func (s *UnLinkAccountRequest) SetExtra(v string) *UnLinkAccountRequest {
	s.Extra = &v
	return s
}

func (s *UnLinkAccountRequest) SetIdentity(v string) *UnLinkAccountRequest {
	s.Identity = &v
	return s
}

func (s *UnLinkAccountRequest) SetType(v string) *UnLinkAccountRequest {
	s.Type = &v
	return s
}

func (s *UnLinkAccountRequest) SetUserId(v string) *UnLinkAccountRequest {
	s.UserId = &v
	return s
}

type UnLinkAccountResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UnLinkAccountResponse) String() string {
	return tea.Prettify(s)
}

func (s UnLinkAccountResponse) GoString() string {
	return s.String()
}

func (s *UnLinkAccountResponse) SetHeaders(v map[string]*string) *UnLinkAccountResponse {
	s.Headers = v
	return s
}

func (s *UnLinkAccountResponse) SetStatusCode(v int32) *UnLinkAccountResponse {
	s.StatusCode = &v
	return s
}

type UpdateDomainRequest struct {
	// The description of the domain.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The domain ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// bj1
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	// The name of the domain.
	DomainName *string `json:"domain_name,omitempty" xml:"domain_name,omitempty"`
	// Specifies whether to enable the default drive feature. A value of true specifies that all users are assigned a drive by default on the first logon. Default value: false.
	//
	// example:
	//
	// true
	InitDriveEnable *bool `json:"init_drive_enable,omitempty" xml:"init_drive_enable,omitempty"`
	// The size of the default drive. Unit: bytes. You must specify init_drive_size if you set init_drive_enable to true. Default value: 0. A value of 0 specifies that the size of the default drive is 0 bytes and you cannot upload files to the drive. To initialize the default drive, set init_drive_size to 0. A value of -1 specifies that the size is unlimited.
	//
	// example:
	//
	// 1073741824
	InitDriveSize *int64 `json:"init_drive_size,omitempty" xml:"init_drive_size,omitempty"`
	// The access policy of the application.
	PublishedAppAccessStrategy *AppAccessStrategy `json:"published_app_access_strategy,omitempty" xml:"published_app_access_strategy,omitempty"`
	// The total storage quota for all drives in the domain. A value of 0 specifies that the quota is unlimited.
	//
	// example:
	//
	// 1099511627776
	SizeQuota *int64 `json:"size_quota,omitempty" xml:"size_quota,omitempty"`
	// The maximum number of users that can be created in the domain.
	//
	// example:
	//
	// 50
	UserCountQuota *int64 `json:"user_count_quota,omitempty" xml:"user_count_quota,omitempty"`
}

func (s UpdateDomainRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainRequest) GoString() string {
	return s.String()
}

func (s *UpdateDomainRequest) SetDescription(v string) *UpdateDomainRequest {
	s.Description = &v
	return s
}

func (s *UpdateDomainRequest) SetDomainId(v string) *UpdateDomainRequest {
	s.DomainId = &v
	return s
}

func (s *UpdateDomainRequest) SetDomainName(v string) *UpdateDomainRequest {
	s.DomainName = &v
	return s
}

func (s *UpdateDomainRequest) SetInitDriveEnable(v bool) *UpdateDomainRequest {
	s.InitDriveEnable = &v
	return s
}

func (s *UpdateDomainRequest) SetInitDriveSize(v int64) *UpdateDomainRequest {
	s.InitDriveSize = &v
	return s
}

func (s *UpdateDomainRequest) SetPublishedAppAccessStrategy(v *AppAccessStrategy) *UpdateDomainRequest {
	s.PublishedAppAccessStrategy = v
	return s
}

func (s *UpdateDomainRequest) SetSizeQuota(v int64) *UpdateDomainRequest {
	s.SizeQuota = &v
	return s
}

func (s *UpdateDomainRequest) SetUserCountQuota(v int64) *UpdateDomainRequest {
	s.UserCountQuota = &v
	return s
}

type UpdateDomainResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Domain            `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDomainResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDomainResponse) GoString() string {
	return s.String()
}

func (s *UpdateDomainResponse) SetHeaders(v map[string]*string) *UpdateDomainResponse {
	s.Headers = v
	return s
}

func (s *UpdateDomainResponse) SetStatusCode(v int32) *UpdateDomainResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDomainResponse) SetBody(v *Domain) *UpdateDomainResponse {
	s.Body = v
	return s
}

type UpdateDriveRequest struct {
	// The description of the drive. The description can be up to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The name of the drive. The name can be up to 128 characters in length.
	//
	// example:
	//
	// my_drive
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	// The owner of the drive. Note: You can modify the owner of a personal drive only by using an AccessKey pair.
	//
	// example:
	//
	// user1
	Owner *string `json:"owner,omitempty" xml:"owner,omitempty"`
	// The state of the drive. Valid values:
	//
	// enabled and disabled.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The total size of the drive. Unit: bytes. A value of -1 specifies that the size is unlimited.
	//
	// example:
	//
	// 10240
	TotalSize *int64 `json:"total_size,omitempty" xml:"total_size,omitempty"`
}

func (s UpdateDriveRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateDriveRequest) GoString() string {
	return s.String()
}

func (s *UpdateDriveRequest) SetDescription(v string) *UpdateDriveRequest {
	s.Description = &v
	return s
}

func (s *UpdateDriveRequest) SetDriveId(v string) *UpdateDriveRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateDriveRequest) SetDriveName(v string) *UpdateDriveRequest {
	s.DriveName = &v
	return s
}

func (s *UpdateDriveRequest) SetOwner(v string) *UpdateDriveRequest {
	s.Owner = &v
	return s
}

func (s *UpdateDriveRequest) SetStatus(v string) *UpdateDriveRequest {
	s.Status = &v
	return s
}

func (s *UpdateDriveRequest) SetTotalSize(v int64) *UpdateDriveRequest {
	s.TotalSize = &v
	return s
}

type UpdateDriveResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Drive             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateDriveResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateDriveResponse) GoString() string {
	return s.String()
}

func (s *UpdateDriveResponse) SetHeaders(v map[string]*string) *UpdateDriveResponse {
	s.Headers = v
	return s
}

func (s *UpdateDriveResponse) SetStatusCode(v int32) *UpdateDriveResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateDriveResponse) SetBody(v *Drive) *UpdateDriveResponse {
	s.Body = v
	return s
}

type UpdateFacegroupRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The face ID of the thumbnail of the face-based group. You can obtain the face ID from the **image_media_metadata*	- parameter in the returned results of the GetFile, ListFile, or SearchFile operation.
	//
	// example:
	//
	// face1
	GroupCoverFaceId *string `json:"group_cover_face_id,omitempty" xml:"group_cover_face_id,omitempty"`
	// The ID of the face-based group. You can call the ListFacegroups operation to query the group ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// group-abc
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The name of the face-based group. The name can be up to 128 characters in length.
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	// The remarks. The remarks can be up to 128 characters in length.
	Remarks *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
}

func (s UpdateFacegroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFacegroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateFacegroupRequest) SetDriveId(v string) *UpdateFacegroupRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateFacegroupRequest) SetGroupCoverFaceId(v string) *UpdateFacegroupRequest {
	s.GroupCoverFaceId = &v
	return s
}

func (s *UpdateFacegroupRequest) SetGroupId(v string) *UpdateFacegroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateFacegroupRequest) SetGroupName(v string) *UpdateFacegroupRequest {
	s.GroupName = &v
	return s
}

func (s *UpdateFacegroupRequest) SetRemarks(v string) *UpdateFacegroupRequest {
	s.Remarks = &v
	return s
}

type UpdateFacegroupResponseBody struct {
	// The drive ID.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The group ID.
	//
	// example:
	//
	// group-abc
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s UpdateFacegroupResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateFacegroupResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateFacegroupResponseBody) SetDriveId(v string) *UpdateFacegroupResponseBody {
	s.DriveId = &v
	return s
}

func (s *UpdateFacegroupResponseBody) SetGroupId(v string) *UpdateFacegroupResponseBody {
	s.GroupId = &v
	return s
}

type UpdateFacegroupResponse struct {
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateFacegroupResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFacegroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFacegroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateFacegroupResponse) SetHeaders(v map[string]*string) *UpdateFacegroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateFacegroupResponse) SetStatusCode(v int32) *UpdateFacegroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFacegroupResponse) SetBody(v *UpdateFacegroupResponseBody) *UpdateFacegroupResponse {
	s.Body = v
	return s
}

type UpdateFileRequest struct {
	// The processing method that is used if the file that you want to modify has the same name as an existing file on the cloud. Valid values:
	//
	// ignore: allows you to modify the file by using the same name as an existing file on the cloud.
	//
	// auto_rename: automatically renames the file that has the same name on the cloud. By default, the current point in time is added to the end of the file name. Example: xxx_20060102_150405.
	//
	// refuse: does not modify the file that you want to modify but returns the information about the file that has the same name on the cloud.
	//
	// Default value: ignore.
	//
	// example:
	//
	// ignore
	CheckNameMode *string `json:"check_name_mode,omitempty" xml:"check_name_mode,omitempty"`
	// The description of the file. The description can be up to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// Specifies whether to hide the file.
	//
	// example:
	//
	// true
	Hidden *bool `json:"hidden,omitempty" xml:"hidden,omitempty"`
	// The tags of the file. You can specify up to 100 tags.
	Labels []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	// The local time when the file was modified. The time is in the yyyy-MM-ddTHH:mm:ssZ format based on the UTC+0 time zone.
	//
	// example:
	//
	// 2019-08-20T06:51:27.292Z
	LocalModifiedAt *string `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	// The name of the file. The name can be up to 1,024 bytes in length based on the UTF-8 encoding rule.
	//
	// example:
	//
	// a.jpg
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// Specifies whether to add the file to favorites.
	//
	// example:
	//
	// true
	Starred *bool `json:"starred,omitempty" xml:"starred,omitempty"`
}

func (s UpdateFileRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileRequest) GoString() string {
	return s.String()
}

func (s *UpdateFileRequest) SetCheckNameMode(v string) *UpdateFileRequest {
	s.CheckNameMode = &v
	return s
}

func (s *UpdateFileRequest) SetDescription(v string) *UpdateFileRequest {
	s.Description = &v
	return s
}

func (s *UpdateFileRequest) SetDriveId(v string) *UpdateFileRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateFileRequest) SetFileId(v string) *UpdateFileRequest {
	s.FileId = &v
	return s
}

func (s *UpdateFileRequest) SetHidden(v bool) *UpdateFileRequest {
	s.Hidden = &v
	return s
}

func (s *UpdateFileRequest) SetLabels(v []*string) *UpdateFileRequest {
	s.Labels = v
	return s
}

func (s *UpdateFileRequest) SetLocalModifiedAt(v string) *UpdateFileRequest {
	s.LocalModifiedAt = &v
	return s
}

func (s *UpdateFileRequest) SetName(v string) *UpdateFileRequest {
	s.Name = &v
	return s
}

func (s *UpdateFileRequest) SetStarred(v bool) *UpdateFileRequest {
	s.Starred = &v
	return s
}

type UpdateFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *File              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateFileResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateFileResponse) GoString() string {
	return s.String()
}

func (s *UpdateFileResponse) SetHeaders(v map[string]*string) *UpdateFileResponse {
	s.Headers = v
	return s
}

func (s *UpdateFileResponse) SetStatusCode(v int32) *UpdateFileResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateFileResponse) SetBody(v *File) *UpdateFileResponse {
	s.Body = v
	return s
}

type UpdateGroupRequest struct {
	// The description of the group after modification.
	//
	// example:
	//
	// test group description
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The ID of the group that you want to modify.
	//
	// This parameter is required.
	//
	// example:
	//
	// 2e43ec8427dd45f19431b7504649a1b4
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	// The name of the group after modification.
	//
	// example:
	//
	// test group
	GroupName *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
}

func (s UpdateGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupRequest) GoString() string {
	return s.String()
}

func (s *UpdateGroupRequest) SetDescription(v string) *UpdateGroupRequest {
	s.Description = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupId(v string) *UpdateGroupRequest {
	s.GroupId = &v
	return s
}

func (s *UpdateGroupRequest) SetGroupName(v string) *UpdateGroupRequest {
	s.GroupName = &v
	return s
}

type UpdateGroupResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Group             `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateGroupResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateGroupResponse) GoString() string {
	return s.String()
}

func (s *UpdateGroupResponse) SetHeaders(v map[string]*string) *UpdateGroupResponse {
	s.Headers = v
	return s
}

func (s *UpdateGroupResponse) SetStatusCode(v int32) *UpdateGroupResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateGroupResponse) SetBody(v *Group) *UpdateGroupResponse {
	s.Body = v
	return s
}

type UpdateIdentityToBenefitPkgMappingRequest struct {
	// The number of benefit packages.
	//
	// This parameter specifies the number of benefit packages of the resource type. Default value: 1.
	//
	// example:
	//
	// 1
	Amount *int64 `json:"amount,omitempty" xml:"amount,omitempty"`
	// The unique identifier of the benefit package.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40cb7794c9294
	BenefitPkgId *string `json:"benefit_pkg_id,omitempty" xml:"benefit_pkg_id,omitempty"`
	// The expiration time of the benefit package. Set this parameter to a UNIX timestamp representing the number of milliseconds that have elapsed since January 1, 1970, 00:00:00 UTC.
	//
	// By default, the benefit package never expires.
	//
	// example:
	//
	// 1633167071000
	ExpireTime *int64 `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	// The unique identifier of the entity.
	//
	// If you call this operation to manage the benefits of a user, set this parameter to the ID of the user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user123
	IdentityId *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
	// The type of the entity. If you call this operation to manage the benefits of a user, set this parameter to user.
	//
	// This parameter is required.
	//
	// example:
	//
	// user
	IdentityType *string `json:"identity_type,omitempty" xml:"identity_type,omitempty"`
}

func (s UpdateIdentityToBenefitPkgMappingRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateIdentityToBenefitPkgMappingRequest) GoString() string {
	return s.String()
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetAmount(v int64) *UpdateIdentityToBenefitPkgMappingRequest {
	s.Amount = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetBenefitPkgId(v string) *UpdateIdentityToBenefitPkgMappingRequest {
	s.BenefitPkgId = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetExpireTime(v int64) *UpdateIdentityToBenefitPkgMappingRequest {
	s.ExpireTime = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetIdentityId(v string) *UpdateIdentityToBenefitPkgMappingRequest {
	s.IdentityId = &v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingRequest) SetIdentityType(v string) *UpdateIdentityToBenefitPkgMappingRequest {
	s.IdentityType = &v
	return s
}

type UpdateIdentityToBenefitPkgMappingResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
}

func (s UpdateIdentityToBenefitPkgMappingResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateIdentityToBenefitPkgMappingResponse) GoString() string {
	return s.String()
}

func (s *UpdateIdentityToBenefitPkgMappingResponse) SetHeaders(v map[string]*string) *UpdateIdentityToBenefitPkgMappingResponse {
	s.Headers = v
	return s
}

func (s *UpdateIdentityToBenefitPkgMappingResponse) SetStatusCode(v int32) *UpdateIdentityToBenefitPkgMappingResponse {
	s.StatusCode = &v
	return s
}

type UpdateRevisionRequest struct {
	// The drive ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// The file ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 9520943DC264
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// Specifies whether to permanently retain a version.
	//
	// By default, this parameter is not specified, which indicates that you do not modify the permanent retention configuration of the version.
	//
	// example:
	//
	// true
	KeepForever *bool `json:"keep_forever,omitempty" xml:"keep_forever,omitempty"`
	// The description of the version. The description can be up to 1,024 characters in length.
	//
	// By default, this parameter is not specified, which indicates that you do not modify the description of the version.
	//
	// example:
	//
	// aaa
	RevisionDescription *string `json:"revision_description,omitempty" xml:"revision_description,omitempty"`
	// The version ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 40CB7794C929
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s UpdateRevisionRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateRevisionRequest) GoString() string {
	return s.String()
}

func (s *UpdateRevisionRequest) SetDriveId(v string) *UpdateRevisionRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateRevisionRequest) SetFileId(v string) *UpdateRevisionRequest {
	s.FileId = &v
	return s
}

func (s *UpdateRevisionRequest) SetKeepForever(v bool) *UpdateRevisionRequest {
	s.KeepForever = &v
	return s
}

func (s *UpdateRevisionRequest) SetRevisionDescription(v string) *UpdateRevisionRequest {
	s.RevisionDescription = &v
	return s
}

func (s *UpdateRevisionRequest) SetRevisionId(v string) *UpdateRevisionRequest {
	s.RevisionId = &v
	return s
}

type UpdateRevisionResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *Revision          `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateRevisionResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateRevisionResponse) GoString() string {
	return s.String()
}

func (s *UpdateRevisionResponse) SetHeaders(v map[string]*string) *UpdateRevisionResponse {
	s.Headers = v
	return s
}

func (s *UpdateRevisionResponse) SetStatusCode(v int32) *UpdateRevisionResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateRevisionResponse) SetBody(v *Revision) *UpdateRevisionResponse {
	s.Body = v
	return s
}

type UpdateShareLinkRequest struct {
	// The description of the share link. The description can be up to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// Specifies whether to prohibit the downloads of the shared files.
	//
	// example:
	//
	// false
	DisableDownload *bool `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	// Specifies whether to prohibit the previews of the shared files.
	//
	// example:
	//
	// false
	DisablePreview *bool `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	// Specifies whether to prohibit the saves of the shared files.
	//
	// example:
	//
	// false
	DisableSave *bool `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	// The number of times that the shared files are downloaded. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 30
	DownloadCount *int64 `json:"download_count,omitempty" xml:"download_count,omitempty"`
	// The maximum number of times that the shared files can be downloaded. The value must be greater than or equal to 0. A value of 0 specifies that the number is unlimited.
	//
	// example:
	//
	// 100
	DownloadLimit *int64 `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	// The time when the share link expires. The time follows the RFC 3339 standard. Example: 2020-06-28T11:33:00.000+08:00. If you leave this parameter empty, the share link never expires.
	//
	// example:
	//
	// 2020-06-28T11:33:00.000+08:00
	Expiration     *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	OfficeEditable *bool   `json:"office_editable,omitempty" xml:"office_editable,omitempty"`
	// The number of times that the shared files are previewed. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 3
	PreviewCount *int64 `json:"preview_count,omitempty" xml:"preview_count,omitempty"`
	// The maximum number of times that the shared files can be previewed. The value must be greater than or equal to 0. A value of 0 specifies that the number is unlimited.
	//
	// example:
	//
	// 100
	PreviewLimit *int64 `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	// The number of times that the shared files are reported. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 1
	ReportCount *int64 `json:"report_count,omitempty" xml:"report_count,omitempty"`
	// The number of times that the shared files are saved. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 5
	SaveCount *int64 `json:"save_count,omitempty" xml:"save_count,omitempty"`
	// The maximum number of times that the shared files can be saved. The value must be greater than or equal to 0. A value of 0 specifies that the number is unlimited.
	//
	// example:
	//
	// 100
	SaveLimit *int64 `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	// The share ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// 7JQX1FswpQ8
	ShareId *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	// The name of the share link. By default, the name of the first file is used. The name can be up to 128 characters in length.
	ShareName *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	// The access code. The access code can be up to 64 characters in length. A value of 0 specifies an empty string.
	//
	// example:
	//
	// abcF123x
	SharePwd *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
	// The state of the share link. Valid values:
	//
	// 	- disabled: The share link is canceled.
	//
	// 	- enabled: The share link is effective.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The number of times that the videos are previewed in the shared files. The value must be greater than or equal to 0.
	//
	// example:
	//
	// 100
	VideoPreviewCount *int64 `json:"video_preview_count,omitempty" xml:"video_preview_count,omitempty"`
}

func (s UpdateShareLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateShareLinkRequest) GoString() string {
	return s.String()
}

func (s *UpdateShareLinkRequest) SetDescription(v string) *UpdateShareLinkRequest {
	s.Description = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDisableDownload(v bool) *UpdateShareLinkRequest {
	s.DisableDownload = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDisablePreview(v bool) *UpdateShareLinkRequest {
	s.DisablePreview = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDisableSave(v bool) *UpdateShareLinkRequest {
	s.DisableSave = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDownloadCount(v int64) *UpdateShareLinkRequest {
	s.DownloadCount = &v
	return s
}

func (s *UpdateShareLinkRequest) SetDownloadLimit(v int64) *UpdateShareLinkRequest {
	s.DownloadLimit = &v
	return s
}

func (s *UpdateShareLinkRequest) SetExpiration(v string) *UpdateShareLinkRequest {
	s.Expiration = &v
	return s
}

func (s *UpdateShareLinkRequest) SetOfficeEditable(v bool) *UpdateShareLinkRequest {
	s.OfficeEditable = &v
	return s
}

func (s *UpdateShareLinkRequest) SetPreviewCount(v int64) *UpdateShareLinkRequest {
	s.PreviewCount = &v
	return s
}

func (s *UpdateShareLinkRequest) SetPreviewLimit(v int64) *UpdateShareLinkRequest {
	s.PreviewLimit = &v
	return s
}

func (s *UpdateShareLinkRequest) SetReportCount(v int64) *UpdateShareLinkRequest {
	s.ReportCount = &v
	return s
}

func (s *UpdateShareLinkRequest) SetSaveCount(v int64) *UpdateShareLinkRequest {
	s.SaveCount = &v
	return s
}

func (s *UpdateShareLinkRequest) SetSaveLimit(v int64) *UpdateShareLinkRequest {
	s.SaveLimit = &v
	return s
}

func (s *UpdateShareLinkRequest) SetShareId(v string) *UpdateShareLinkRequest {
	s.ShareId = &v
	return s
}

func (s *UpdateShareLinkRequest) SetShareName(v string) *UpdateShareLinkRequest {
	s.ShareName = &v
	return s
}

func (s *UpdateShareLinkRequest) SetSharePwd(v string) *UpdateShareLinkRequest {
	s.SharePwd = &v
	return s
}

func (s *UpdateShareLinkRequest) SetStatus(v string) *UpdateShareLinkRequest {
	s.Status = &v
	return s
}

func (s *UpdateShareLinkRequest) SetVideoPreviewCount(v int64) *UpdateShareLinkRequest {
	s.VideoPreviewCount = &v
	return s
}

type UpdateShareLinkResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *ShareLink         `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateShareLinkResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateShareLinkResponse) GoString() string {
	return s.String()
}

func (s *UpdateShareLinkResponse) SetHeaders(v map[string]*string) *UpdateShareLinkResponse {
	s.Headers = v
	return s
}

func (s *UpdateShareLinkResponse) SetStatusCode(v int32) *UpdateShareLinkResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateShareLinkResponse) SetBody(v *ShareLink) *UpdateShareLinkResponse {
	s.Body = v
	return s
}

type UpdateStoryRequest struct {
	// if can be null:
	// true
	Cover *UpdateStoryRequestCover `json:"cover,omitempty" xml:"cover,omitempty" type:"Struct"`
	// Deprecated
	CustomLabels map[string]*string `json:"custom_labels,omitempty" xml:"custom_labels,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
	// if can be null:
	// true
	//
	// example:
	//
	// name1
	StoryName *string `json:"story_name,omitempty" xml:"story_name,omitempty"`
}

func (s UpdateStoryRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoryRequest) GoString() string {
	return s.String()
}

func (s *UpdateStoryRequest) SetCover(v *UpdateStoryRequestCover) *UpdateStoryRequest {
	s.Cover = v
	return s
}

func (s *UpdateStoryRequest) SetCustomLabels(v map[string]*string) *UpdateStoryRequest {
	s.CustomLabels = v
	return s
}

func (s *UpdateStoryRequest) SetDriveId(v string) *UpdateStoryRequest {
	s.DriveId = &v
	return s
}

func (s *UpdateStoryRequest) SetStoryId(v string) *UpdateStoryRequest {
	s.StoryId = &v
	return s
}

func (s *UpdateStoryRequest) SetStoryName(v string) *UpdateStoryRequest {
	s.StoryName = &v
	return s
}

type UpdateStoryRequestCover struct {
	// example:
	//
	// 63e5e4340f76cb3ead5f40f68163f0f967c1a7bf
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	// example:
	//
	// 642a88dd06e49d9c0a14411ebae606f70edd9a59
	RevisionId *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
}

func (s UpdateStoryRequestCover) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoryRequestCover) GoString() string {
	return s.String()
}

func (s *UpdateStoryRequestCover) SetFileId(v string) *UpdateStoryRequestCover {
	s.FileId = &v
	return s
}

func (s *UpdateStoryRequestCover) SetRevisionId(v string) *UpdateStoryRequestCover {
	s.RevisionId = &v
	return s
}

type UpdateStoryResponseBody struct {
	// example:
	//
	// 1
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	// example:
	//
	// 9132e0d8-fe92-4e56-86c3-f5f112308003
	StoryId *string `json:"story_id,omitempty" xml:"story_id,omitempty"`
}

func (s UpdateStoryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateStoryResponseBody) SetDriveId(v string) *UpdateStoryResponseBody {
	s.DriveId = &v
	return s
}

func (s *UpdateStoryResponseBody) SetStoryId(v string) *UpdateStoryResponseBody {
	s.StoryId = &v
	return s
}

type UpdateStoryResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *UpdateStoryResponseBody `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateStoryResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateStoryResponse) GoString() string {
	return s.String()
}

func (s *UpdateStoryResponse) SetHeaders(v map[string]*string) *UpdateStoryResponse {
	s.Headers = v
	return s
}

func (s *UpdateStoryResponse) SetStatusCode(v int32) *UpdateStoryResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateStoryResponse) SetBody(v *UpdateStoryResponseBody) *UpdateStoryResponse {
	s.Body = v
	return s
}

type UpdateUserRequest struct {
	// The URL of the profile picture.
	//
	// If you specify the parameter in the HTTP URL format, the URL must start with http:// or https:// and can be up to 4 KB in size.
	//
	// If you specify the parameter in the DATA URL format, the URL must start with data:// and be encoded in Base64. The URL can be up to 300 KB in size.
	//
	// example:
	//
	// http://a.b.c/pds.jpg
	Avatar *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	// The description of the user. The description can be up to 1,024 characters in length.
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	// The email address of the user.
	//
	// example:
	//
	// a@aliyunpds.com
	Email *string `json:"email,omitempty" xml:"email,omitempty"`
	// The information about the group.
	GroupInfoList []*UpdateUserRequestGroupInfoList `json:"group_info_list,omitempty" xml:"group_info_list,omitempty" type:"Repeated"`
	// The nickname of the user. The nickname can be up to 128 characters in length.
	//
	// example:
	//
	// pdsuer
	NickName *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	// The mobile number of the user.
	//
	// example:
	//
	// 13900001111
	Phone *string `json:"phone,omitempty" xml:"phone,omitempty"`
	// The role of the user. Valid values:
	//
	// 	- superadmin
	//
	// 	- admin
	//
	// 	- user
	//
	// example:
	//
	// user
	Role *string `json:"role,omitempty" xml:"role,omitempty"`
	// The state of the user. Valid values:
	//
	// 	- disabled: The user is prohibited from logon.
	//
	// 	- enabled: The user is in a normal state.
	//
	// example:
	//
	// enabled
	Status *string `json:"status,omitempty" xml:"status,omitempty"`
	// The custom data. The data can be up to 1,024 characters in length.
	UserData map[string]*string `json:"user_data,omitempty" xml:"user_data,omitempty"`
	// The user ID. The ID can be up to 64 characters in length and cannot contain a number sign (#).
	//
	// This parameter is required.
	//
	// example:
	//
	// c9b7a5aa04d14ae3867fdc886fa01da4
	UserId *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s UpdateUserRequest) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequest) GoString() string {
	return s.String()
}

func (s *UpdateUserRequest) SetAvatar(v string) *UpdateUserRequest {
	s.Avatar = &v
	return s
}

func (s *UpdateUserRequest) SetDescription(v string) *UpdateUserRequest {
	s.Description = &v
	return s
}

func (s *UpdateUserRequest) SetEmail(v string) *UpdateUserRequest {
	s.Email = &v
	return s
}

func (s *UpdateUserRequest) SetGroupInfoList(v []*UpdateUserRequestGroupInfoList) *UpdateUserRequest {
	s.GroupInfoList = v
	return s
}

func (s *UpdateUserRequest) SetNickName(v string) *UpdateUserRequest {
	s.NickName = &v
	return s
}

func (s *UpdateUserRequest) SetPhone(v string) *UpdateUserRequest {
	s.Phone = &v
	return s
}

func (s *UpdateUserRequest) SetRole(v string) *UpdateUserRequest {
	s.Role = &v
	return s
}

func (s *UpdateUserRequest) SetStatus(v string) *UpdateUserRequest {
	s.Status = &v
	return s
}

func (s *UpdateUserRequest) SetUserData(v map[string]*string) *UpdateUserRequest {
	s.UserData = v
	return s
}

func (s *UpdateUserRequest) SetUserId(v string) *UpdateUserRequest {
	s.UserId = &v
	return s
}

type UpdateUserRequestGroupInfoList struct {
	// The group ID.
	//
	// example:
	//
	// g123
	GroupId *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
}

func (s UpdateUserRequestGroupInfoList) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserRequestGroupInfoList) GoString() string {
	return s.String()
}

func (s *UpdateUserRequestGroupInfoList) SetGroupId(v string) *UpdateUserRequestGroupInfoList {
	s.GroupId = &v
	return s
}

type UpdateUserResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty"`
	Body       *User              `json:"body,omitempty" xml:"body,omitempty"`
}

func (s UpdateUserResponse) String() string {
	return tea.Prettify(s)
}

func (s UpdateUserResponse) GoString() string {
	return s.String()
}

func (s *UpdateUserResponse) SetHeaders(v map[string]*string) *UpdateUserResponse {
	s.Headers = v
	return s
}

func (s *UpdateUserResponse) SetStatusCode(v int32) *UpdateUserResponse {
	s.StatusCode = &v
	return s
}

func (s *UpdateUserResponse) SetBody(v *User) *UpdateUserResponse {
	s.Body = v
	return s
}

type Client struct {
	openapi.Client
}

func NewClient(config *openapi.Config) (*Client, error) {
	client := new(Client)
	err := client.Init(config)
	return client, err
}

func (client *Client) Init(config *openapi.Config) (_err error) {
	_err = client.Client.Init(config)
	if _err != nil {
		return _err
	}
	gatewayClient, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = gatewayClient
	
	client.EndpointRule = tea.String("")
	return nil
}

// Summary:
//
// Adds a member to a group.
//
// @param request - AddGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddGroupMemberResponse
func (client *Client) AddGroupMemberWithOptions(request *AddGroupMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["group_id"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		body["member_id"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberType)) {
		body["member_type"] = request.MemberType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddGroupMember"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/group/add_member"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds a member to a group.
//
// @param request - AddGroupMemberRequest
//
// @return AddGroupMemberResponse
func (client *Client) AddGroupMember(request *AddGroupMemberRequest) (_result *AddGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddGroupMemberResponse{}
	_body, _err := client.AddGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 故事添加文件
//
// @param request - AddStoryFilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AddStoryFilesResponse
func (client *Client) AddStoryFilesWithOptions(request *AddStoryFilesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AddStoryFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Files)) {
		body["files"] = request.Files
	}

	if !tea.BoolValue(util.IsUnset(request.StoryId)) {
		body["story_id"] = request.StoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AddStoryFiles"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/add_story_files"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AddStoryFilesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 故事添加文件
//
// @param request - AddStoryFilesRequest
//
// @return AddStoryFilesResponse
func (client *Client) AddStoryFiles(request *AddStoryFilesRequest) (_result *AddStoryFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AddStoryFilesResponse{}
	_body, _err := client.AddStoryFilesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Assigns a group administrator role to a user.
//
// Description:
//
// You can call this operation to assign a group administrator role to a user.
//
// @param request - AssignRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AssignRoleResponse
func (client *Client) AssignRoleWithOptions(request *AssignRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AssignRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		body["identity"] = request.Identity
	}

	if !tea.BoolValue(util.IsUnset(request.ManageResourceId)) {
		body["manage_resource_id"] = request.ManageResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManageResourceType)) {
		body["manage_resource_type"] = request.ManageResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		body["role_id"] = request.RoleId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("AssignRole"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/role/assign"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &AssignRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Assigns a group administrator role to a user.
//
// Description:
//
// You can call this operation to assign a group administrator role to a user.
//
// @param request - AssignRoleRequest
//
// @return AssignRoleResponse
func (client *Client) AssignRole(request *AssignRoleRequest) (_result *AssignRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AssignRoleResponse{}
	_body, _err := client.AssignRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Requests permissions by using OAuth 2.0.
//
// Description:
//
// For more information, see "OAuth 2.0 For Web Server Applications" at [OAuth 2.0 For Web Server Applications](https://www.alibabacloud.com/help/en/pds/drive-and-photo-service-dev/user-guide/oauth-2-0-access-process-for-web-server-applications) in User Guide.
//
// @param tmpReq - AuthorizeRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return AuthorizeResponse
func (client *Client) AuthorizeWithOptions(tmpReq *AuthorizeRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *AuthorizeResponse, _err error) {
	_err = util.ValidateModel(tmpReq)
	if _err != nil {
		return _result, _err
	}
	request := &AuthorizeShrinkRequest{}
	openapiutil.Convert(tmpReq, request)
	if !tea.BoolValue(util.IsUnset(tmpReq.Scope)) {
		request.ScopeShrink = openapiutil.ArrayToStringWithSpecifiedStyle(tmpReq.Scope, tea.String("scope"), tea.String("simple"))
	}

	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		query["client_id"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.HideConsent)) {
		query["hide_consent"] = request.HideConsent
	}

	if !tea.BoolValue(util.IsUnset(request.LoginType)) {
		query["login_type"] = request.LoginType
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectUri)) {
		query["redirect_uri"] = request.RedirectUri
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseType)) {
		query["response_type"] = request.ResponseType
	}

	if !tea.BoolValue(util.IsUnset(request.ScopeShrink)) {
		query["scope"] = request.ScopeShrink
	}

	if !tea.BoolValue(util.IsUnset(request.State)) {
		query["state"] = request.State
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("Authorize"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/oauth/authorize"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("binary"),
	}
	_result = &AuthorizeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Requests permissions by using OAuth 2.0.
//
// Description:
//
// For more information, see "OAuth 2.0 For Web Server Applications" at [OAuth 2.0 For Web Server Applications](https://www.alibabacloud.com/help/en/pds/drive-and-photo-service-dev/user-guide/oauth-2-0-access-process-for-web-server-applications) in User Guide.
//
// @param request - AuthorizeRequest
//
// @return AuthorizeResponse
func (client *Client) Authorize(request *AuthorizeRequest) (_result *AuthorizeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &AuthorizeResponse{}
	_body, _err := client.AuthorizeWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Calls multiple operations at a time to improve call efficiency.
//
// @param request - BatchRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return BatchResponse
func (client *Client) BatchWithOptions(request *BatchRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *BatchResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Requests)) {
		body["requests"] = request.Requests
	}

	if !tea.BoolValue(util.IsUnset(request.Resource)) {
		body["resource"] = request.Resource
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Batch"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/batch"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &BatchResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Calls multiple operations at a time to improve call efficiency.
//
// @param request - BatchRequest
//
// @return BatchResponse
func (client *Client) Batch(request *BatchRequest) (_result *BatchResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &BatchResponse{}
	_body, _err := client.BatchWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels a role.
//
// Description:
//
// You can cancel only the group administrator role.
//
// @param request - CancelAssignRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelAssignRoleResponse
func (client *Client) CancelAssignRoleWithOptions(request *CancelAssignRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelAssignRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		body["identity"] = request.Identity
	}

	if !tea.BoolValue(util.IsUnset(request.ManageResourceId)) {
		body["manage_resource_id"] = request.ManageResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManageResourceType)) {
		body["manage_resource_type"] = request.ManageResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.RoleId)) {
		body["role_id"] = request.RoleId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelAssignRole"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/role/cancel_assign"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelAssignRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels a role.
//
// Description:
//
// You can cancel only the group administrator role.
//
// @param request - CancelAssignRoleRequest
//
// @return CancelAssignRoleResponse
func (client *Client) CancelAssignRole(request *CancelAssignRoleRequest) (_result *CancelAssignRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelAssignRoleResponse{}
	_body, _err := client.CancelAssignRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a share link.
//
// @param request - CancelShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CancelShareLinkResponse
func (client *Client) CancelShareLinkWithOptions(request *CancelShareLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CancelShareLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CancelShareLink"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/share_link/cancel"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CancelShareLinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a share link.
//
// @param request - CancelShareLinkRequest
//
// @return CancelShareLinkResponse
func (client *Client) CancelShareLink(request *CancelShareLinkRequest) (_result *CancelShareLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CancelShareLinkResponse{}
	_body, _err := client.CancelShareLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Empties the recycle bin.
//
// @param request - ClearRecyclebinRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ClearRecyclebinResponse
func (client *Client) ClearRecyclebinWithOptions(request *ClearRecyclebinRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ClearRecyclebinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ClearRecyclebin"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/recyclebin/clear"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ClearRecyclebinResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Empties the recycle bin.
//
// @param request - ClearRecyclebinRequest
//
// @return ClearRecyclebinResponse
func (client *Client) ClearRecyclebin(request *ClearRecyclebinRequest) (_result *ClearRecyclebinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ClearRecyclebinResponse{}
	_body, _err := client.ClearRecyclebinWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Completes the upload of a file.
//
// @param request - CompleteFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CompleteFileResponse
func (client *Client) CompleteFileWithOptions(request *CompleteFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CompleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.UploadId)) {
		body["upload_id"] = request.UploadId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CompleteFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/complete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CompleteFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Completes the upload of a file.
//
// @param request - CompleteFileRequest
//
// @return CompleteFileResponse
func (client *Client) CompleteFile(request *CompleteFileRequest) (_result *CompleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CompleteFileResponse{}
	_body, _err := client.CompleteFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Copies a file or folder.
//
// @param request - CopyFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CopyFileResponse
func (client *Client) CopyFileWithOptions(request *CopyFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CopyFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoRename)) {
		body["auto_rename"] = request.AutoRename
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.ToDriveId)) {
		body["to_drive_id"] = request.ToDriveId
	}

	if !tea.BoolValue(util.IsUnset(request.ToParentFileId)) {
		body["to_parent_file_id"] = request.ToParentFileId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CopyFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/copy"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CopyFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Copies a file or folder.
//
// @param request - CopyFileRequest
//
// @return CopyFileResponse
func (client *Client) CopyFile(request *CopyFileRequest) (_result *CopyFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CopyFileResponse{}
	_body, _err := client.CopyFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建自定义故事
//
// @param request - CreateCustomizedStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateCustomizedStoryResponse
func (client *Client) CreateCustomizedStoryWithOptions(request *CreateCustomizedStoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateCustomizedStoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CustomLabels)) {
		body["custom_labels"] = request.CustomLabels
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.StoryCover)) {
		body["story_cover"] = request.StoryCover
	}

	if !tea.BoolValue(util.IsUnset(request.StoryFiles)) {
		body["story_files"] = request.StoryFiles
	}

	if !tea.BoolValue(util.IsUnset(request.StoryName)) {
		body["story_name"] = request.StoryName
	}

	if !tea.BoolValue(util.IsUnset(request.StorySubType)) {
		body["story_sub_type"] = request.StorySubType
	}

	if !tea.BoolValue(util.IsUnset(request.StoryType)) {
		body["story_type"] = request.StoryType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateCustomizedStory"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/create_customized_story"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateCustomizedStoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建自定义故事
//
// @param request - CreateCustomizedStoryRequest
//
// @return CreateCustomizedStoryResponse
func (client *Client) CreateCustomizedStory(request *CreateCustomizedStoryRequest) (_result *CreateCustomizedStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateCustomizedStoryResponse{}
	_body, _err := client.CreateCustomizedStoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// test_domain
//
// Description:
//
// The description of the domain.
//
// @param request - CreateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDomainResponse
func (client *Client) CreateDomainWithOptions(request *CreateDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["domain_name"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.InitDriveEnable)) {
		body["init_drive_enable"] = request.InitDriveEnable
	}

	if !tea.BoolValue(util.IsUnset(request.InitDriveSize)) {
		body["init_drive_size"] = request.InitDriveSize
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDomainId)) {
		body["parent_domain_id"] = request.ParentDomainId
	}

	if !tea.BoolValue(util.IsUnset(request.SizeQuota)) {
		body["size_quota"] = request.SizeQuota
	}

	if !tea.BoolValue(util.IsUnset(request.UserCountQuota)) {
		body["user_count_quota"] = request.UserCountQuota
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDomain"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/domain/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDomainResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// test_domain
//
// Description:
//
// The description of the domain.
//
// @param request - CreateDomainRequest
//
// @return CreateDomainResponse
func (client *Client) CreateDomain(request *CreateDomainRequest) (_result *CreateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDomainResponse{}
	_body, _err := client.CreateDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a drive.
//
// @param request - CreateDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateDriveResponse
func (client *Client) CreateDriveWithOptions(request *CreateDriveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateDriveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Default)) {
		body["default"] = request.Default
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DriveName)) {
		body["drive_name"] = request.DriveName
	}

	if !tea.BoolValue(util.IsUnset(request.DriveType)) {
		body["drive_type"] = request.DriveType
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		body["owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerType)) {
		body["owner_type"] = request.OwnerType
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TotalSize)) {
		body["total_size"] = request.TotalSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateDrive"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/drive/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateDriveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a drive.
//
// @param request - CreateDriveRequest
//
// @return CreateDriveResponse
func (client *Client) CreateDrive(request *CreateDriveRequest) (_result *CreateDriveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateDriveResponse{}
	_body, _err := client.CreateDriveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a file or folder.
//
// @param request - CreateFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateFileResponse
func (client *Client) CreateFileWithOptions(request *CreateFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckNameMode)) {
		body["check_name_mode"] = request.CheckNameMode
	}

	if !tea.BoolValue(util.IsUnset(request.ContentHash)) {
		body["content_hash"] = request.ContentHash
	}

	if !tea.BoolValue(util.IsUnset(request.ContentHashName)) {
		body["content_hash_name"] = request.ContentHashName
	}

	if !tea.BoolValue(util.IsUnset(request.ContentType)) {
		body["content_type"] = request.ContentType
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.Hidden)) {
		body["hidden"] = request.Hidden
	}

	if !tea.BoolValue(util.IsUnset(request.ImageMediaMetadata)) {
		body["image_media_metadata"] = request.ImageMediaMetadata
	}

	if !tea.BoolValue(util.IsUnset(request.LocalCreatedAt)) {
		body["local_created_at"] = request.LocalCreatedAt
	}

	if !tea.BoolValue(util.IsUnset(request.LocalModifiedAt)) {
		body["local_modified_at"] = request.LocalModifiedAt
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.ParallelUpload)) {
		body["parallel_upload"] = request.ParallelUpload
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFileId)) {
		body["parent_file_id"] = request.ParentFileId
	}

	if !tea.BoolValue(util.IsUnset(request.PartInfoList)) {
		body["part_info_list"] = request.PartInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.PreHash)) {
		body["pre_hash"] = request.PreHash
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.Size)) {
		body["size"] = request.Size
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserTags)) {
		body["user_tags"] = request.UserTags
	}

	if !tea.BoolValue(util.IsUnset(request.VideoMediaMetadata)) {
		body["video_media_metadata"] = request.VideoMediaMetadata
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a file or folder.
//
// @param request - CreateFileRequest
//
// @return CreateFileResponse
func (client *Client) CreateFile(request *CreateFileRequest) (_result *CreateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateFileResponse{}
	_body, _err := client.CreateFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a group.
//
// @param request - CreateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateGroupResponse
func (client *Client) CreateGroupWithOptions(request *CreateGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["group_name"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.IsRoot)) {
		body["is_root"] = request.IsRoot
	}

	if !tea.BoolValue(util.IsUnset(request.ParentGroupId)) {
		body["parent_group_id"] = request.ParentGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateGroup"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/group/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a group.
//
// @param request - CreateGroupRequest
//
// @return CreateGroupResponse
func (client *Client) CreateGroup(request *CreateGroupRequest) (_result *CreateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateGroupResponse{}
	_body, _err := client.CreateGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a mapping between an entity and a benefit package. You can call this operation to associate a benefit package with a user.
//
// Description:
//
// If you need to manage a large number of users based on Drive and Photo Service, you can control the features and quotas that users can use based on the benefits to which they are entitled. For more information, join the DingTalk group (ID 23146118).
//
// @param request - CreateIdentityToBenefitPkgMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateIdentityToBenefitPkgMappingResponse
func (client *Client) CreateIdentityToBenefitPkgMappingWithOptions(request *CreateIdentityToBenefitPkgMappingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateIdentityToBenefitPkgMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.BenefitPkgId)) {
		body["benefit_pkg_id"] = request.BenefitPkgId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["expire_time"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityId)) {
		body["identity_id"] = request.IdentityId
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityType)) {
		body["identity_type"] = request.IdentityType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateIdentityToBenefitPkgMapping"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/benefit/identity_to_benefit_pkg_mapping/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a mapping between an entity and a benefit package. You can call this operation to associate a benefit package with a user.
//
// Description:
//
// If you need to manage a large number of users based on Drive and Photo Service, you can control the features and quotas that users can use based on the benefits to which they are entitled. For more information, join the DingTalk group (ID 23146118).
//
// @param request - CreateIdentityToBenefitPkgMappingRequest
//
// @return CreateIdentityToBenefitPkgMappingResponse
func (client *Client) CreateIdentityToBenefitPkgMapping(request *CreateIdentityToBenefitPkgMappingRequest) (_result *CreateIdentityToBenefitPkgMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.CreateIdentityToBenefitPkgMappingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建凌霄订单
//
// @param request - CreateOrderRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateOrderResponse
func (client *Client) CreateOrderWithOptions(request *CreateOrderRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateOrderResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AutoPay)) {
		body["auto_pay"] = request.AutoPay
	}

	if !tea.BoolValue(util.IsUnset(request.AutoRenew)) {
		body["auto_renew"] = request.AutoRenew
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instance_id"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		body["order_type"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.Package)) {
		body["package"] = request.Package
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		body["period_unit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.TotalSize)) {
		body["total_size"] = request.TotalSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserCount)) {
		body["user_count"] = request.UserCount
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateOrder"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/domain/create_order"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateOrderResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建凌霄订单
//
// @param request - CreateOrderRequest
//
// @return CreateOrderResponse
func (client *Client) CreateOrder(request *CreateOrderRequest) (_result *CreateOrderResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateOrderResponse{}
	_body, _err := client.CreateOrderWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a share URL.
//
// Description:
//
// A share is a file view container. You can grant anonymous users the permissions to access files in the user drive by using the share URL. Anonymous users can access the files based on the granted permissions.
//
// @param request - CreateShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateShareLinkResponse
func (client *Client) CreateShareLinkWithOptions(request *CreateShareLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateShareLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creatable)) {
		body["creatable"] = request.Creatable
	}

	if !tea.BoolValue(util.IsUnset(request.CreatableFileIdList)) {
		body["creatable_file_id_list"] = request.CreatableFileIdList
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisableDownload)) {
		body["disable_download"] = request.DisableDownload
	}

	if !tea.BoolValue(util.IsUnset(request.DisablePreview)) {
		body["disable_preview"] = request.DisablePreview
	}

	if !tea.BoolValue(util.IsUnset(request.DisableSave)) {
		body["disable_save"] = request.DisableSave
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadLimit)) {
		body["download_limit"] = request.DownloadLimit
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		body["expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.FileIdList)) {
		body["file_id_list"] = request.FileIdList
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeEditable)) {
		body["office_editable"] = request.OfficeEditable
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewLimit)) {
		body["preview_limit"] = request.PreviewLimit
	}

	if !tea.BoolValue(util.IsUnset(request.RequireLogin)) {
		body["require_login"] = request.RequireLogin
	}

	if !tea.BoolValue(util.IsUnset(request.SaveLimit)) {
		body["save_limit"] = request.SaveLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ShareAllFiles)) {
		body["share_all_files"] = request.ShareAllFiles
	}

	if !tea.BoolValue(util.IsUnset(request.ShareName)) {
		body["share_name"] = request.ShareName
	}

	if !tea.BoolValue(util.IsUnset(request.SharePwd)) {
		body["share_pwd"] = request.SharePwd
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateShareLink"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/share_link/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateShareLinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a share URL.
//
// Description:
//
// A share is a file view container. You can grant anonymous users the permissions to access files in the user drive by using the share URL. Anonymous users can access the files based on the granted permissions.
//
// @param request - CreateShareLinkRequest
//
// @return CreateShareLinkResponse
func (client *Client) CreateShareLink(request *CreateShareLinkRequest) (_result *CreateShareLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateShareLinkResponse{}
	_body, _err := client.CreateShareLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建相似图片聚类任务
//
// @param request - CreateSimilarImageClusterTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateSimilarImageClusterTaskResponse
func (client *Client) CreateSimilarImageClusterTaskWithOptions(request *CreateSimilarImageClusterTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateSimilarImageClusterTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateSimilarImageClusterTask"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/create_similar_image_cluster_task"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateSimilarImageClusterTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建相似图片聚类任务
//
// @param request - CreateSimilarImageClusterTaskRequest
//
// @return CreateSimilarImageClusterTaskResponse
func (client *Client) CreateSimilarImageClusterTask(request *CreateSimilarImageClusterTaskRequest) (_result *CreateSimilarImageClusterTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateSimilarImageClusterTaskResponse{}
	_body, _err := client.CreateSimilarImageClusterTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 创建推荐故事
//
// @param request - CreateStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateStoryResponse
func (client *Client) CreateStoryWithOptions(request *CreateStoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateStoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Address)) {
		body["address"] = request.Address
	}

	if !tea.BoolValue(util.IsUnset(request.CustomLabels)) {
		body["custom_labels"] = request.CustomLabels
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.MaxImageCount)) {
		body["max_image_count"] = request.MaxImageCount
	}

	if !tea.BoolValue(util.IsUnset(request.MinImageCount)) {
		body["min_image_count"] = request.MinImageCount
	}

	if !tea.BoolValue(util.IsUnset(request.StoryEndTime)) {
		body["story_end_time"] = request.StoryEndTime
	}

	if !tea.BoolValue(util.IsUnset(request.StoryId)) {
		body["story_id"] = request.StoryId
	}

	if !tea.BoolValue(util.IsUnset(request.StoryName)) {
		body["story_name"] = request.StoryName
	}

	if !tea.BoolValue(util.IsUnset(request.StoryStartTime)) {
		body["story_start_time"] = request.StoryStartTime
	}

	if !tea.BoolValue(util.IsUnset(request.StorySubType)) {
		body["story_sub_type"] = request.StorySubType
	}

	if !tea.BoolValue(util.IsUnset(request.StoryType)) {
		body["story_type"] = request.StoryType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateStory"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/create_story"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateStoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 创建推荐故事
//
// @param request - CreateStoryRequest
//
// @return CreateStoryResponse
func (client *Client) CreateStory(request *CreateStoryRequest) (_result *CreateStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateStoryResponse{}
	_body, _err := client.CreateStoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Creates a user.
//
// @param request - CreateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CreateUserResponse
func (client *Client) CreateUserWithOptions(request *CreateUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Avatar)) {
		body["avatar"] = request.Avatar
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.GroupInfoList)) {
		body["group_info_list"] = request.GroupInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		body["nick_name"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["user_data"] = request.UserData
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["user_name"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CreateUser"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/user/create"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CreateUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Creates a user.
//
// @param request - CreateUserRequest
//
// @return CreateUserResponse
func (client *Client) CreateUser(request *CreateUserRequest) (_result *CreateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CreateUserResponse{}
	_body, _err := client.CreateUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取文件内容安全信息
//
// @param request - CsiGetFileInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return CsiGetFileInfoResponse
func (client *Client) CsiGetFileInfoWithOptions(request *CsiGetFileInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CsiGetFileInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.UrlExpireSec)) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("CsiGetFileInfo"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/csi/get_file_info"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &CsiGetFileInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取文件内容安全信息
//
// @param request - CsiGetFileInfoRequest
//
// @return CsiGetFileInfoResponse
func (client *Client) CsiGetFileInfo(request *CsiGetFileInfoRequest) (_result *CsiGetFileInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &CsiGetFileInfoResponse{}
	_body, _err := client.CsiGetFileInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Delete the domain
//
// @param request - DeleteDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomainWithOptions(request *DeleteDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		body["domain_id"] = request.DomainId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDomain"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/domain/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDomainResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Delete the domain
//
// @param request - DeleteDomainRequest
//
// @return DeleteDomainResponse
func (client *Client) DeleteDomain(request *DeleteDomainRequest) (_result *DeleteDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDomainResponse{}
	_body, _err := client.DeleteDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a drive.
//
// @param request - DeleteDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteDriveResponse
func (client *Client) DeleteDriveWithOptions(request *DeleteDriveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteDriveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteDrive"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/drive/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteDriveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a drive.
//
// @param request - DeleteDriveRequest
//
// @return DeleteDriveResponse
func (client *Client) DeleteDrive(request *DeleteDriveRequest) (_result *DeleteDriveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteDriveResponse{}
	_body, _err := client.DeleteDriveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a file or folder.
//
// @param request - DeleteFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteFileResponse
func (client *Client) DeleteFileWithOptions(request *DeleteFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a file or folder.
//
// @param request - DeleteFileRequest
//
// @return DeleteFileResponse
func (client *Client) DeleteFile(request *DeleteFileRequest) (_result *DeleteFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteFileResponse{}
	_body, _err := client.DeleteFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes groups. Before you delete a group, make sure that no other groups or users exist in the group. Otherwise, the group fails to be deleted.
//
// @param request - DeleteGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroupWithOptions(request *DeleteGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["group_id"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteGroup"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/group/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes groups. Before you delete a group, make sure that no other groups or users exist in the group. Otherwise, the group fails to be deleted.
//
// @param request - DeleteGroupRequest
//
// @return DeleteGroupResponse
func (client *Client) DeleteGroup(request *DeleteGroupRequest) (_result *DeleteGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteGroupResponse{}
	_body, _err := client.DeleteGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a historical version of a file. You cannot delete the latest version of a file.
//
// @param request - DeleteRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteRevisionResponse
func (client *Client) DeleteRevisionWithOptions(request *DeleteRevisionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteRevisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionId)) {
		body["revision_id"] = request.RevisionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteRevision"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/revision/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteRevisionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a historical version of a file. You cannot delete the latest version of a file.
//
// @param request - DeleteRevisionRequest
//
// @return DeleteRevisionResponse
func (client *Client) DeleteRevision(request *DeleteRevisionRequest) (_result *DeleteRevisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteRevisionResponse{}
	_body, _err := client.DeleteRevisionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 删除故事
//
// @param request - DeleteStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteStoryResponse
func (client *Client) DeleteStoryWithOptions(request *DeleteStoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteStoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.StoryId)) {
		body["story_id"] = request.StoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteStory"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/delete_story"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteStoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 删除故事
//
// @param request - DeleteStoryRequest
//
// @return DeleteStoryResponse
func (client *Client) DeleteStory(request *DeleteStoryRequest) (_result *DeleteStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteStoryResponse{}
	_body, _err := client.DeleteStoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Deletes a user.
//
// @param request - DeleteUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeleteUserResponse
func (client *Client) DeleteUserWithOptions(request *DeleteUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeleteUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeleteUser"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/user/delete"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeleteUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Deletes a user.
//
// @param request - DeleteUserRequest
//
// @return DeleteUserResponse
func (client *Client) DeleteUser(request *DeleteUserRequest) (_result *DeleteUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeleteUserResponse{}
	_body, _err := client.DeleteUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the cursor of incremental information.
//
// @param request - DeltaGetLastCursorRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DeltaGetLastCursorResponse
func (client *Client) DeltaGetLastCursorWithOptions(request *DeltaGetLastCursorRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DeltaGetLastCursorResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.SyncRootId)) {
		body["sync_root_id"] = request.SyncRootId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("DeltaGetLastCursor"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/get_last_cursor"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &DeltaGetLastCursorResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the cursor of incremental information.
//
// @param request - DeltaGetLastCursorRequest
//
// @return DeltaGetLastCursorResponse
func (client *Client) DeltaGetLastCursor(request *DeltaGetLastCursorRequest) (_result *DeltaGetLastCursorResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DeltaGetLastCursorResponse{}
	_body, _err := client.DeltaGetLastCursorWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Downloads a file.
//
// Description:
//
// For information about best practices for downloading a file.
//
// @param request - DownloadFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return DownloadFileResponse
func (client *Client) DownloadFileWithOptions(request *DownloadFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DownloadFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	query := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		query["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		query["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageThumbnailProcess)) {
		query["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeThumbnailProcess)) {
		query["office_thumbnail_process"] = request.OfficeThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		query["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.VideoThumbnailProcess)) {
		query["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Query:   openapiutil.Query(query),
	}
	params := &openapi.Params{
		Action:      tea.String("DownloadFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/download"),
		Method:      tea.String("GET"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("binary"),
	}
	_result = &DownloadFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Downloads a file.
//
// Description:
//
// For information about best practices for downloading a file.
//
// @param request - DownloadFileRequest
//
// @return DownloadFileResponse
func (client *Client) DownloadFile(request *DownloadFileRequest) (_result *DownloadFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &DownloadFileResponse{}
	_body, _err := client.DownloadFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Grants permissions to access files to a user or group.
//
// @param request - FileAddPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileAddPermissionResponse
func (client *Client) FileAddPermissionWithOptions(request *FileAddPermissionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FileAddPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberList)) {
		body["member_list"] = request.MemberList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FileAddPermission"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/add_permission"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FileAddPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Grants permissions to access files to a user or group.
//
// @param request - FileAddPermissionRequest
//
// @return FileAddPermissionResponse
func (client *Client) FileAddPermission(request *FileAddPermissionRequest) (_result *FileAddPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FileAddPermissionResponse{}
	_body, _err := client.FileAddPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes custom tags from a file.
//
// @param request - FileDeleteUserTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileDeleteUserTagsResponse
func (client *Client) FileDeleteUserTagsWithOptions(request *FileDeleteUserTagsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FileDeleteUserTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.KeyList)) {
		body["key_list"] = request.KeyList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FileDeleteUserTags"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/delete_usertags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FileDeleteUserTagsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes custom tags from a file.
//
// @param request - FileDeleteUserTagsRequest
//
// @return FileDeleteUserTagsResponse
func (client *Client) FileDeleteUserTags(request *FileDeleteUserTagsRequest) (_result *FileDeleteUserTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FileDeleteUserTagsResponse{}
	_body, _err := client.FileDeleteUserTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the sharing authorization records of a file.
//
// @param request - FileListPermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileListPermissionResponse
func (client *Client) FileListPermissionWithOptions(request *FileListPermissionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FileListPermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FileListPermission"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/list_permission"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("array"),
	}
	_result = &FileListPermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the sharing authorization records of a file.
//
// @param request - FileListPermissionRequest
//
// @return FileListPermissionResponse
func (client *Client) FileListPermission(request *FileListPermissionRequest) (_result *FileListPermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FileListPermissionResponse{}
	_body, _err := client.FileListPermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Adds custom tags to a file.
//
// Description:
//
// This operation is an incremental update operation. Take note of the following items:
//
// 	- If a tag name specified in the request is the same as an existing tag name, the existing tag is overwritten.
//
// 	- If a tag name specified in the request is different from the existing tag names, the specified tag is added.
//
// 	- The existing tags with unique names are not affected.
//
// @param request - FilePutUserTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FilePutUserTagsResponse
func (client *Client) FilePutUserTagsWithOptions(request *FilePutUserTagsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FilePutUserTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.UserTags)) {
		body["user_tags"] = request.UserTags
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FilePutUserTags"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/put_usertags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FilePutUserTagsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Adds custom tags to a file.
//
// Description:
//
// This operation is an incremental update operation. Take note of the following items:
//
// 	- If a tag name specified in the request is the same as an existing tag name, the existing tag is overwritten.
//
// 	- If a tag name specified in the request is different from the existing tag names, the specified tag is added.
//
// 	- The existing tags with unique names are not affected.
//
// @param request - FilePutUserTagsRequest
//
// @return FilePutUserTagsResponse
func (client *Client) FilePutUserTags(request *FilePutUserTagsRequest) (_result *FilePutUserTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FilePutUserTagsResponse{}
	_body, _err := client.FilePutUserTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Cancels the permissions on a shared file.
//
// @param request - FileRemovePermissionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return FileRemovePermissionResponse
func (client *Client) FileRemovePermissionWithOptions(request *FileRemovePermissionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *FileRemovePermissionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberList)) {
		body["member_list"] = request.MemberList
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("FileRemovePermission"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/remove_permission"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &FileRemovePermissionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Cancels the permissions on a shared file.
//
// @param request - FileRemovePermissionRequest
//
// @return FileRemovePermissionResponse
func (client *Client) FileRemovePermission(request *FileRemovePermissionRequest) (_result *FileRemovePermissionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &FileRemovePermissionResponse{}
	_body, _err := client.FileRemovePermissionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an asynchronous task.
//
// @param request - GetAsyncTaskRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTaskWithOptions(request *GetAsyncTaskRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetAsyncTaskResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AsyncTaskId)) {
		body["async_task_id"] = request.AsyncTaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetAsyncTask"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/async_task/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an asynchronous task.
//
// @param request - GetAsyncTaskRequest
//
// @return GetAsyncTaskResponse
func (client *Client) GetAsyncTask(request *GetAsyncTaskRequest) (_result *GetAsyncTaskResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetAsyncTaskResponse{}
	_body, _err := client.GetAsyncTaskWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the default drive of a user.
//
// @param request - GetDefaultDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDefaultDriveResponse
func (client *Client) GetDefaultDriveWithOptions(request *GetDefaultDriveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDefaultDriveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDefaultDrive"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/drive/get_default_drive"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDefaultDriveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the default drive of a user.
//
// @param request - GetDefaultDriveRequest
//
// @return GetDefaultDriveResponse
func (client *Client) GetDefaultDrive(request *GetDefaultDriveRequest) (_result *GetDefaultDriveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDefaultDriveResponse{}
	_body, _err := client.GetDefaultDriveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Get domain information.
//
// @param request - GetDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainResponse
func (client *Client) GetDomainWithOptions(request *GetDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		body["domain_id"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.GetQuotaUsed)) {
		body["get_quota_used"] = request.GetQuotaUsed
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomain"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/domain/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Get domain information.
//
// @param request - GetDomainRequest
//
// @return GetDomainResponse
func (client *Client) GetDomain(request *GetDomainRequest) (_result *GetDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDomainResponse{}
	_body, _err := client.GetDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取domain限额
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDomainQuotaResponse
func (client *Client) GetDomainQuotaWithOptions(headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDomainQuotaResponse, _err error) {
	req := &openapi.OpenApiRequest{
		Headers: headers,
	}
	params := &openapi.Params{
		Action:      tea.String("GetDomainQuota"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/domain/get_quota"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDomainQuotaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取domain限额
//
// @return GetDomainQuotaResponse
func (client *Client) GetDomainQuota() (_result *GetDomainQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDomainQuotaResponse{}
	_body, _err := client.GetDomainQuotaWithOptions(headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the download URL of a file. For more information about best practices, visit https://help.aliyun.com/document_detail/175889.html.
//
// @param request - GetDownloadUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDownloadUrlResponse
func (client *Client) GetDownloadUrlWithOptions(request *GetDownloadUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireSec)) {
		body["expire_sec"] = request.ExpireSec
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.FileName)) {
		body["file_name"] = request.FileName
	}

	if !tea.BoolValue(util.IsUnset(request.ResponseContentType)) {
		body["response_content_type"] = request.ResponseContentType
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDownloadUrl"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/get_download_url"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDownloadUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the download URL of a file. For more information about best practices, visit https://help.aliyun.com/document_detail/175889.html.
//
// @param request - GetDownloadUrlRequest
//
// @return GetDownloadUrlResponse
func (client *Client) GetDownloadUrl(request *GetDownloadUrlRequest) (_result *GetDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDownloadUrlResponse{}
	_body, _err := client.GetDownloadUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a drive.
//
// @param request - GetDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetDriveResponse
func (client *Client) GetDriveWithOptions(request *GetDriveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetDriveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetDrive"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/drive/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetDriveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a drive.
//
// @param request - GetDriveRequest
//
// @return GetDriveResponse
func (client *Client) GetDrive(request *GetDriveRequest) (_result *GetDriveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetDriveResponse{}
	_body, _err := client.GetDriveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a file.
//
// @param request - GetFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetFileResponse
func (client *Client) GetFileWithOptions(request *GetFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.ThumbnailProcesses)) {
		body["thumbnail_processes"] = request.ThumbnailProcesses
	}

	if !tea.BoolValue(util.IsUnset(request.UrlExpireSec)) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a file.
//
// @param request - GetFileRequest
//
// @return GetFileResponse
func (client *Client) GetFile(request *GetFileRequest) (_result *GetFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetFileResponse{}
	_body, _err := client.GetFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a group.
//
// @param request - GetGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetGroupResponse
func (client *Client) GetGroupWithOptions(request *GetGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["group_id"] = request.GroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetGroup"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/group/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a group.
//
// @param request - GetGroupRequest
//
// @return GetGroupResponse
func (client *Client) GetGroup(request *GetGroupRequest) (_result *GetGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetGroupResponse{}
	_body, _err := client.GetGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the mapping between an entity and a benefit package. You can call this operation to query the benefit package that is associated with a user.
//
// @param request - GetIdentityToBenefitPkgMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetIdentityToBenefitPkgMappingResponse
func (client *Client) GetIdentityToBenefitPkgMappingWithOptions(request *GetIdentityToBenefitPkgMappingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetIdentityToBenefitPkgMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.BenefitPkgId)) {
		body["benefit_pkg_id"] = request.BenefitPkgId
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityId)) {
		body["identity_id"] = request.IdentityId
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityType)) {
		body["identity_type"] = request.IdentityType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetIdentityToBenefitPkgMapping"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/benefit/identity_to_benefit_pkg_mapping/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mapping between an entity and a benefit package. You can call this operation to query the benefit package that is associated with a user.
//
// @param request - GetIdentityToBenefitPkgMappingRequest
//
// @return GetIdentityToBenefitPkgMappingResponse
func (client *Client) GetIdentityToBenefitPkgMapping(request *GetIdentityToBenefitPkgMappingRequest) (_result *GetIdentityToBenefitPkgMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.GetIdentityToBenefitPkgMappingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about an account.
//
// @param request - GetLinkInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLinkInfoResponse
func (client *Client) GetLinkInfoWithOptions(request *GetLinkInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLinkInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		body["identity"] = request.Identity
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLinkInfo"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/account/get_link_info"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLinkInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about an account.
//
// @param request - GetLinkInfoRequest
//
// @return GetLinkInfoResponse
func (client *Client) GetLinkInfo(request *GetLinkInfoRequest) (_result *GetLinkInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLinkInfoResponse{}
	_body, _err := client.GetLinkInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a user based on the user ID.
//
// @param request - GetLinkInfoByUserIdRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetLinkInfoByUserIdResponse
func (client *Client) GetLinkInfoByUserIdWithOptions(request *GetLinkInfoByUserIdRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetLinkInfoByUserIdResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetLinkInfoByUserId"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/account/get_link_info_by_user_id"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetLinkInfoByUserIdResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a user based on the user ID.
//
// @param request - GetLinkInfoByUserIdRequest
//
// @return GetLinkInfoByUserIdResponse
func (client *Client) GetLinkInfoByUserId(request *GetLinkInfoByUserIdRequest) (_result *GetLinkInfoByUserIdResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetLinkInfoByUserIdResponse{}
	_body, _err := client.GetLinkInfoByUserIdWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a version.
//
// @param request - GetRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetRevisionResponse
func (client *Client) GetRevisionWithOptions(request *GetRevisionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetRevisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionId)) {
		body["revision_id"] = request.RevisionId
	}

	if !tea.BoolValue(util.IsUnset(request.UrlExpireSec)) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetRevision"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/revision/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetRevisionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a version.
//
// @param request - GetRevisionRequest
//
// @return GetRevisionResponse
func (client *Client) GetRevision(request *GetRevisionRequest) (_result *GetRevisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetRevisionResponse{}
	_body, _err := client.GetRevisionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the share URL of a file.
//
// @param request - GetShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShareLinkResponse
func (client *Client) GetShareLinkWithOptions(request *GetShareLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetShareLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetShareLink"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/share_link/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetShareLinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the share URL of a file.
//
// @param request - GetShareLinkRequest
//
// @return GetShareLinkResponse
func (client *Client) GetShareLink(request *GetShareLinkRequest) (_result *GetShareLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetShareLinkResponse{}
	_body, _err := client.GetShareLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a share link anonymously.
//
// @param request - GetShareLinkByAnonymousRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShareLinkByAnonymousResponse
func (client *Client) GetShareLinkByAnonymousWithOptions(request *GetShareLinkByAnonymousRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetShareLinkByAnonymousResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetShareLinkByAnonymous"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/share_link/get_by_anonymous"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetShareLinkByAnonymousResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a share link anonymously.
//
// @param request - GetShareLinkByAnonymousRequest
//
// @return GetShareLinkByAnonymousResponse
func (client *Client) GetShareLinkByAnonymous(request *GetShareLinkByAnonymousRequest) (_result *GetShareLinkByAnonymousResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetShareLinkByAnonymousResponse{}
	_body, _err := client.GetShareLinkByAnonymousWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a share token anonymously.
//
// Description:
//
// To access a file by using a share link, you must first obtain a share token, even if the value of share_pwd of this share is an empty string, which specifies that the share is not private.
//
// @param request - GetShareLinkTokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetShareLinkTokenResponse
func (client *Client) GetShareLinkTokenWithOptions(request *GetShareLinkTokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetShareLinkTokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.ExpireSec)) {
		body["expire_sec"] = request.ExpireSec
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.SharePwd)) {
		body["share_pwd"] = request.SharePwd
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetShareLinkToken"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/share_link/get_share_token"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetShareLinkTokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a share token anonymously.
//
// Description:
//
// To access a file by using a share link, you must first obtain a share token, even if the value of share_pwd of this share is an empty string, which specifies that the share is not private.
//
// @param request - GetShareLinkTokenRequest
//
// @return GetShareLinkTokenResponse
func (client *Client) GetShareLinkToken(request *GetShareLinkTokenRequest) (_result *GetShareLinkTokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetShareLinkTokenResponse{}
	_body, _err := client.GetShareLinkTokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取故事详情
//
// @param request - GetStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetStoryResponse
func (client *Client) GetStoryWithOptions(request *GetStoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetStoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoverImageThumbnailProcess)) {
		body["cover_image_thumbnail_process"] = request.CoverImageThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.CoverVideoThumbnailProcess)) {
		body["cover_video_thumbnail_process"] = request.CoverVideoThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageThumbnailProcess)) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.ImageUrlProcess)) {
		body["image_url_process"] = request.ImageUrlProcess
	}

	if !tea.BoolValue(util.IsUnset(request.StoryId)) {
		body["story_id"] = request.StoryId
	}

	if !tea.BoolValue(util.IsUnset(request.UrlExpireSec)) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	if !tea.BoolValue(util.IsUnset(request.VideoThumbnailProcess)) {
		body["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetStory"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/get_story"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetStoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取故事详情
//
// @param request - GetStoryRequest
//
// @return GetStoryResponse
func (client *Client) GetStory(request *GetStoryRequest) (_result *GetStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetStoryResponse{}
	_body, _err := client.GetStoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the execution status of a value-added asynchronous task. You can call this operation to query the execution status of an asynchronous task that is created by calling the CreateSimilarImageClusterTask operation.
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/425220.html) of Drive and Photo Service**.
//
// To call this operation, make sure that the value-added image processing feature is enabled.
//
// Before you call this operation, a value-added asynchronous task must be created. For example, you can call the CreateSimilarImageClusterTask operation to create an asynchronous task. Then, you can call this operation to query the execution status of the asynchronous task based on the task ID.
//
// @param request - GetTaskStatusRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatusWithOptions(request *GetTaskStatusRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetTaskStatusResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.TaskId)) {
		body["task_id"] = request.TaskId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetTaskStatus"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/get_task_status"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetTaskStatusResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the execution status of a value-added asynchronous task. You can call this operation to query the execution status of an asynchronous task that is created by calling the CreateSimilarImageClusterTask operation.
//
// Description:
//
// *Before you call this operation, make sure that you are familiar with the [billing](https://help.aliyun.com/document_detail/425220.html) of Drive and Photo Service**.
//
// To call this operation, make sure that the value-added image processing feature is enabled.
//
// Before you call this operation, a value-added asynchronous task must be created. For example, you can call the CreateSimilarImageClusterTask operation to create an asynchronous task. Then, you can call this operation to query the execution status of the asynchronous task based on the task ID.
//
// @param request - GetTaskStatusRequest
//
// @return GetTaskStatusResponse
func (client *Client) GetTaskStatus(request *GetTaskStatusRequest) (_result *GetTaskStatusResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetTaskStatusResponse{}
	_body, _err := client.GetTaskStatusWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the upload URL of a file.
//
// @param request - GetUploadUrlRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUploadUrlResponse
func (client *Client) GetUploadUrlWithOptions(request *GetUploadUrlRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUploadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.PartInfoList)) {
		body["part_info_list"] = request.PartInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.UploadId)) {
		body["upload_id"] = request.UploadId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUploadUrl"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/get_upload_url"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUploadUrlResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the upload URL of a file.
//
// @param request - GetUploadUrlRequest
//
// @return GetUploadUrlResponse
func (client *Client) GetUploadUrl(request *GetUploadUrlRequest) (_result *GetUploadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUploadUrlResponse{}
	_body, _err := client.GetUploadUrlWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about a user.
//
// @param request - GetUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetUserResponse
func (client *Client) GetUserWithOptions(request *GetUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetUser"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/user/get"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about a user.
//
// @param request - GetUserRequest
//
// @return GetUserResponse
func (client *Client) GetUser(request *GetUserRequest) (_result *GetUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetUserResponse{}
	_body, _err := client.GetUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about video playback.
//
// Description:
//
// For more information about best practices, see [Preview videos online](https://help.aliyun.com/document_detail/427477.html).
//
// @param request - GetVideoPreviewPlayInfoRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoPreviewPlayInfoResponse
func (client *Client) GetVideoPreviewPlayInfoWithOptions(request *GetVideoPreviewPlayInfoRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVideoPreviewPlayInfoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.GetMasterUrl)) {
		body["get_master_url"] = request.GetMasterUrl
	}

	if !tea.BoolValue(util.IsUnset(request.GetWithoutUrl)) {
		body["get_without_url"] = request.GetWithoutUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ReTranscode)) {
		body["re_transcode"] = request.ReTranscode
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["template_id"] = request.TemplateId
	}

	if !tea.BoolValue(util.IsUnset(request.UrlExpireSec)) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVideoPreviewPlayInfo"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/get_video_preview_play_info"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVideoPreviewPlayInfoResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about video playback.
//
// Description:
//
// For more information about best practices, see [Preview videos online](https://help.aliyun.com/document_detail/427477.html).
//
// @param request - GetVideoPreviewPlayInfoRequest
//
// @return GetVideoPreviewPlayInfoResponse
func (client *Client) GetVideoPreviewPlayInfo(request *GetVideoPreviewPlayInfoRequest) (_result *GetVideoPreviewPlayInfoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVideoPreviewPlayInfoResponse{}
	_body, _err := client.GetVideoPreviewPlayInfoWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the preview metadata of a video.
//
// Description:
//
// For more information about best practices, see [Preview videos online](https://help.aliyun.com/document_detail/427477.html).
//
// @param request - GetVideoPreviewPlayMetaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GetVideoPreviewPlayMetaResponse
func (client *Client) GetVideoPreviewPlayMetaWithOptions(request *GetVideoPreviewPlayMetaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GetVideoPreviewPlayMetaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GetVideoPreviewPlayMeta"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/get_video_preview_play_meta"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GetVideoPreviewPlayMetaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the preview metadata of a video.
//
// Description:
//
// For more information about best practices, see [Preview videos online](https://help.aliyun.com/document_detail/427477.html).
//
// @param request - GetVideoPreviewPlayMetaRequest
//
// @return GetVideoPreviewPlayMetaResponse
func (client *Client) GetVideoPreviewPlayMeta(request *GetVideoPreviewPlayMetaRequest) (_result *GetVideoPreviewPlayMetaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GetVideoPreviewPlayMetaResponse{}
	_body, _err := client.GetVideoPreviewPlayMetaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新用户组名字
//
// @param request - GroupUpdateNameRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return GroupUpdateNameResponse
func (client *Client) GroupUpdateNameWithOptions(request *GroupUpdateNameRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *GroupUpdateNameResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["group_id"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("GroupUpdateName"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/group/update_name"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &GroupUpdateNameResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新用户组名字
//
// @param request - GroupUpdateNameRequest
//
// @return GroupUpdateNameResponse
func (client *Client) GroupUpdateName(request *GroupUpdateNameRequest) (_result *GroupUpdateNameResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &GroupUpdateNameResponse{}
	_body, _err := client.GroupUpdateNameWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Imports a user.
//
// @param request - ImportUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ImportUserResponse
func (client *Client) ImportUserWithOptions(request *ImportUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ImportUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AuthenticationDisplayName)) {
		body["authentication_display_name"] = request.AuthenticationDisplayName
	}

	if !tea.BoolValue(util.IsUnset(request.AuthenticationType)) {
		body["authentication_type"] = request.AuthenticationType
	}

	if !tea.BoolValue(util.IsUnset(request.AutoCreateDrive)) {
		body["auto_create_drive"] = request.AutoCreateDrive
	}

	if !tea.BoolValue(util.IsUnset(request.DriveTotalSize)) {
		body["drive_total_size"] = request.DriveTotalSize
	}

	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		body["identity"] = request.Identity
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		body["nick_name"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.ParentGroupId)) {
		body["parent_group_id"] = request.ParentGroupId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ImportUser"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/user/import"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ImportUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Imports a user.
//
// @param request - ImportUserRequest
//
// @return ImportUserResponse
func (client *Client) ImportUser(request *ImportUserRequest) (_result *ImportUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ImportUserResponse{}
	_body, _err := client.ImportUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 送审文件
//
// @param request - InvestigateFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return InvestigateFileResponse
func (client *Client) InvestigateFileWithOptions(request *InvestigateFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *InvestigateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveFileIds)) {
		body["drive_file_ids"] = request.DriveFileIds
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("InvestigateFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/csi/investigate_file"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &InvestigateFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 送审文件
//
// @param request - InvestigateFileRequest
//
// @return InvestigateFileResponse
func (client *Client) InvestigateFile(request *InvestigateFileRequest) (_result *InvestigateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &InvestigateFileResponse{}
	_body, _err := client.InvestigateFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Associates an account with a user.
//
// @param request - LinkAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return LinkAccountResponse
func (client *Client) LinkAccountWithOptions(request *LinkAccountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *LinkAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		body["identity"] = request.Identity
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("LinkAccount"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/account/link"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &LinkAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Associates an account with a user.
//
// @param request - LinkAccountRequest
//
// @return LinkAccountResponse
func (client *Client) LinkAccount(request *LinkAccountRequest) (_result *LinkAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &LinkAccountResponse{}
	_body, _err := client.LinkAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries location-based groups.
//
// @param request - ListAddressGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAddressGroupsResponse
func (client *Client) ListAddressGroupsWithOptions(request *ListAddressGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAddressGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageThumbnailProcess)) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.VideoThumbnailProcess)) {
		body["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAddressGroups"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/list_address_groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAddressGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries location-based groups.
//
// @param request - ListAddressGroupsRequest
//
// @return ListAddressGroupsResponse
func (client *Client) ListAddressGroups(request *ListAddressGroupsRequest) (_result *ListAddressGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAddressGroupsResponse{}
	_body, _err := client.ListAddressGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of assigned roles. For example, you can query the administrators of a group by group ID.
//
// @param request - ListAssignmentRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListAssignmentResponse
func (client *Client) ListAssignmentWithOptions(request *ListAssignmentRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListAssignmentResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.ManageResourceId)) {
		body["manage_resource_id"] = request.ManageResourceId
	}

	if !tea.BoolValue(util.IsUnset(request.ManageResourceType)) {
		body["manage_resource_type"] = request.ManageResourceType
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListAssignment"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/role/list_assignment"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListAssignmentResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of assigned roles. For example, you can query the administrators of a group by group ID.
//
// @param request - ListAssignmentRequest
//
// @return ListAssignmentResponse
func (client *Client) ListAssignment(request *ListAssignmentRequest) (_result *ListAssignmentResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListAssignmentResponse{}
	_body, _err := client.ListAssignmentWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries incremental information.
//
// @param request - ListDeltaRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDeltaResponse
func (client *Client) ListDeltaWithOptions(request *ListDeltaRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDeltaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cursor)) {
		body["cursor"] = request.Cursor
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.SyncRootId)) {
		body["sync_root_id"] = request.SyncRootId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDelta"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/list_delta"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDeltaResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries incremental information.
//
// @param request - ListDeltaRequest
//
// @return ListDeltaResponse
func (client *Client) ListDelta(request *ListDeltaRequest) (_result *ListDeltaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDeltaResponse{}
	_body, _err := client.ListDeltaWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举 domain
//
// @param request - ListDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDomainsResponse
func (client *Client) ListDomainsWithOptions(request *ListDomainsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.ParentDomainId)) {
		body["parent_domain_id"] = request.ParentDomainId
	}

	if !tea.BoolValue(util.IsUnset(request.ServiceCode)) {
		body["service_code"] = request.ServiceCode
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDomains"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/domain/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDomainsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举 domain
//
// @param request - ListDomainsRequest
//
// @return ListDomainsResponse
func (client *Client) ListDomains(request *ListDomainsRequest) (_result *ListDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDomainsResponse{}
	_body, _err := client.ListDomainsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of drives.
//
// @param request - ListDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListDriveResponse
func (client *Client) ListDriveWithOptions(request *ListDriveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListDriveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		body["owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerType)) {
		body["owner_type"] = request.OwnerType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListDrive"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/drive/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListDriveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of drives.
//
// @param request - ListDriveRequest
//
// @return ListDriveResponse
func (client *Client) ListDrive(request *ListDriveRequest) (_result *ListDriveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListDriveResponse{}
	_body, _err := client.ListDriveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries face-based groups.
//
// @param request - ListFacegroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFacegroupsResponse
func (client *Client) ListFacegroupsWithOptions(request *ListFacegroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFacegroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Remarks)) {
		body["remarks"] = request.Remarks
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnTotalCount)) {
		body["return_total_count"] = request.ReturnTotalCount
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFacegroups"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/list_facegroups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFacegroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries face-based groups.
//
// @param request - ListFacegroupsRequest
//
// @return ListFacegroupsResponse
func (client *Client) ListFacegroups(request *ListFacegroupsRequest) (_result *ListFacegroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFacegroupsResponse{}
	_body, _err := client.ListFacegroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of files and folders.
//
// @param request - ListFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListFileResponse
func (client *Client) ListFileWithOptions(request *ListFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Category)) {
		body["category"] = request.Category
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["order_by"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDirection)) {
		body["order_direction"] = request.OrderDirection
	}

	if !tea.BoolValue(util.IsUnset(request.ParentFileId)) {
		body["parent_file_id"] = request.ParentFileId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.ThumbnailProcesses)) {
		body["thumbnail_processes"] = request.ThumbnailProcesses
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of files and folders.
//
// @param request - ListFileRequest
//
// @return ListFileResponse
func (client *Client) ListFile(request *ListFileRequest) (_result *ListFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListFileResponse{}
	_body, _err := client.ListFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries groups.
//
// @param request - ListGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupResponse
func (client *Client) ListGroupWithOptions(request *ListGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroup"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/group/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries groups.
//
// @param request - ListGroupRequest
//
// @return ListGroupResponse
func (client *Client) ListGroup(request *ListGroupRequest) (_result *ListGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGroupResponse{}
	_body, _err := client.ListGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the members of a group.
//
// @param request - ListGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListGroupMemberResponse
func (client *Client) ListGroupMemberWithOptions(request *ListGroupMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["group_id"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.MemberType)) {
		body["member_type"] = request.MemberType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListGroupMember"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/group/list_member"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the members of a group.
//
// @param request - ListGroupMemberRequest
//
// @return ListGroupMemberResponse
func (client *Client) ListGroupMember(request *ListGroupMemberRequest) (_result *ListGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListGroupMemberResponse{}
	_body, _err := client.ListGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 列举用户或团队已分配的角色列表
//
// @param request - ListIdentityRoleRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityRoleResponse
func (client *Client) ListIdentityRoleWithOptions(request *ListIdentityRoleRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIdentityRoleResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		body["identity"] = request.Identity
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIdentityRole"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/role/list_identity_role"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIdentityRoleResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 列举用户或团队已分配的角色列表
//
// @param request - ListIdentityRoleRequest
//
// @return ListIdentityRoleResponse
func (client *Client) ListIdentityRole(request *ListIdentityRoleRequest) (_result *ListIdentityRoleResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIdentityRoleResponse{}
	_body, _err := client.ListIdentityRoleWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the mappings between entities and benefit packages. You can call this operation to query the benefit packages that are associated with a user.
//
// @param request - ListIdentityToBenefitPkgMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListIdentityToBenefitPkgMappingResponse
func (client *Client) ListIdentityToBenefitPkgMappingWithOptions(request *ListIdentityToBenefitPkgMappingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListIdentityToBenefitPkgMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.IdentityId)) {
		body["identity_id"] = request.IdentityId
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityType)) {
		body["identity_type"] = request.IdentityType
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeExpired)) {
		body["include_expired"] = request.IncludeExpired
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListIdentityToBenefitPkgMapping"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/benefit/identity_to_benefit_pkg_mapping/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the mappings between entities and benefit packages. You can call this operation to query the benefit packages that are associated with a user.
//
// @param request - ListIdentityToBenefitPkgMappingRequest
//
// @return ListIdentityToBenefitPkgMappingResponse
func (client *Client) ListIdentityToBenefitPkgMapping(request *ListIdentityToBenefitPkgMappingRequest) (_result *ListIdentityToBenefitPkgMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.ListIdentityToBenefitPkgMappingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the drives of the current user.
//
// @param request - ListMyDrivesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMyDrivesResponse
func (client *Client) ListMyDrivesWithOptions(request *ListMyDrivesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMyDrivesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMyDrives"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/drive/list_my_drives"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMyDrivesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the drives of the current user.
//
// @param request - ListMyDrivesRequest
//
// @return ListMyDrivesResponse
func (client *Client) ListMyDrives(request *ListMyDrivesRequest) (_result *ListMyDrivesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMyDrivesResponse{}
	_body, _err := client.ListMyDrivesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the team drives that can be accessed by the authorized users.
//
// @param request - ListMyGroupDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListMyGroupDriveResponse
func (client *Client) ListMyGroupDriveWithOptions(request *ListMyGroupDriveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListMyGroupDriveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveName)) {
		body["drive_name"] = request.DriveName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListMyGroupDrive"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/drive/list_my_group_drive"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListMyGroupDriveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the team drives that can be accessed by the authorized users.
//
// @param request - ListMyGroupDriveRequest
//
// @return ListMyGroupDriveResponse
func (client *Client) ListMyGroupDrive(request *ListMyGroupDriveRequest) (_result *ListMyGroupDriveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListMyGroupDriveResponse{}
	_body, _err := client.ListMyGroupDriveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries a list of files that are shared with a user. You can call this operation to query a list of files in a personal drive on which a user is granted permissions.
//
// @param request - ListReceivedFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListReceivedFileResponse
func (client *Client) ListReceivedFileWithOptions(request *ListReceivedFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListReceivedFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListReceivedFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/list_received_file"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListReceivedFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries a list of files that are shared with a user. You can call this operation to query a list of files in a personal drive on which a user is granted permissions.
//
// @param request - ListReceivedFileRequest
//
// @return ListReceivedFileResponse
func (client *Client) ListReceivedFile(request *ListReceivedFileRequest) (_result *ListReceivedFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListReceivedFileResponse{}
	_body, _err := client.ListReceivedFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the information about files and folders in the recycle bin.
//
// @param request - ListRecyclebinRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRecyclebinResponse
func (client *Client) ListRecyclebinWithOptions(request *ListRecyclebinRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRecyclebinResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRecyclebin"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/recyclebin/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRecyclebinResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the information about files and folders in the recycle bin.
//
// @param request - ListRecyclebinRequest
//
// @return ListRecyclebinResponse
func (client *Client) ListRecyclebin(request *ListRecyclebinRequest) (_result *ListRecyclebinResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRecyclebinResponse{}
	_body, _err := client.ListRecyclebinWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the versions of a file.
//
// @param request - ListRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListRevisionResponse
func (client *Client) ListRevisionWithOptions(request *ListRevisionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListRevisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListRevision"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/revision/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListRevisionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the versions of a file.
//
// @param request - ListRevisionRequest
//
// @return ListRevisionResponse
func (client *Client) ListRevision(request *ListRevisionRequest) (_result *ListRevisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListRevisionResponse{}
	_body, _err := client.ListRevisionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries shares.
//
// Description:
//
// This operation is discontinued. To query shares, you can call the SearchShareLink operation.
//
// @param request - ListShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListShareLinkResponse
func (client *Client) ListShareLinkWithOptions(request *ListShareLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListShareLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creator)) {
		body["creator"] = request.Creator
	}

	if !tea.BoolValue(util.IsUnset(request.IncludeCancelled)) {
		body["include_cancelled"] = request.IncludeCancelled
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["order_by"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDirection)) {
		body["order_direction"] = request.OrderDirection
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListShareLink"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/share_link/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListShareLinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries shares.
//
// Description:
//
// This operation is discontinued. To query shares, you can call the SearchShareLink operation.
//
// @param request - ListShareLinkRequest
//
// @return ListShareLinkResponse
func (client *Client) ListShareLink(request *ListShareLinkRequest) (_result *ListShareLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListShareLinkResponse{}
	_body, _err := client.ListShareLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries tags.
//
// Description:
//
// You can call this operation to query the tags within the specified drive at a time. The top 2,000 tags of the images are returned.
//
// @param request - ListTagsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListTagsResponse
func (client *Client) ListTagsWithOptions(request *ListTagsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageThumbnailProcess)) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.VideoThumbnailProcess)) {
		body["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListTags"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/list_tags"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries tags.
//
// Description:
//
// You can call this operation to query the tags within the specified drive at a time. The top 2,000 tags of the images are returned.
//
// @param request - ListTagsRequest
//
// @return ListTagsResponse
func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries the file parts that are uploaded.
//
// @param request - ListUploadedPartsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUploadedPartsResponse
func (client *Client) ListUploadedPartsWithOptions(request *ListUploadedPartsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUploadedPartsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.PartNumberMarker)) {
		body["part_number_marker"] = request.PartNumberMarker
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.UploadId)) {
		body["upload_id"] = request.UploadId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUploadedParts"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/list_uploaded_parts"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUploadedPartsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries the file parts that are uploaded.
//
// @param request - ListUploadedPartsRequest
//
// @return ListUploadedPartsResponse
func (client *Client) ListUploadedParts(request *ListUploadedPartsRequest) (_result *ListUploadedPartsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUploadedPartsResponse{}
	_body, _err := client.ListUploadedPartsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries users.
//
// @param request - ListUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ListUserResponse
func (client *Client) ListUserWithOptions(request *ListUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ListUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ListUser"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/user/list"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ListUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries users.
//
// @param request - ListUserRequest
//
// @return ListUserResponse
func (client *Client) ListUser(request *ListUserRequest) (_result *ListUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ListUserResponse{}
	_body, _err := client.ListUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves files or folders.
//
// @param request - MoveFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return MoveFileResponse
func (client *Client) MoveFileWithOptions(request *MoveFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *MoveFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckNameMode)) {
		body["check_name_mode"] = request.CheckNameMode
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.ToParentFileId)) {
		body["to_parent_file_id"] = request.ToParentFileId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("MoveFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/move"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &MoveFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves files or folders.
//
// @param request - MoveFileRequest
//
// @return MoveFileResponse
func (client *Client) MoveFile(request *MoveFileRequest) (_result *MoveFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &MoveFileResponse{}
	_body, _err := client.MoveFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询凌霄订单价格
//
// @param request - QueryOrderPriceRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return QueryOrderPriceResponse
func (client *Client) QueryOrderPriceWithOptions(request *QueryOrderPriceRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *QueryOrderPriceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.InstanceId)) {
		body["instance_id"] = request.InstanceId
	}

	if !tea.BoolValue(util.IsUnset(request.OrderType)) {
		body["order_type"] = request.OrderType
	}

	if !tea.BoolValue(util.IsUnset(request.Package)) {
		body["package"] = request.Package
	}

	if !tea.BoolValue(util.IsUnset(request.Period)) {
		body["period"] = request.Period
	}

	if !tea.BoolValue(util.IsUnset(request.PeriodUnit)) {
		body["period_unit"] = request.PeriodUnit
	}

	if !tea.BoolValue(util.IsUnset(request.TotalSize)) {
		body["total_size"] = request.TotalSize
	}

	if !tea.BoolValue(util.IsUnset(request.UserCount)) {
		body["user_count"] = request.UserCount
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("QueryOrderPrice"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/domain/query_order_price"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &QueryOrderPriceResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询凌霄订单价格
//
// @param request - QueryOrderPriceRequest
//
// @return QueryOrderPriceResponse
func (client *Client) QueryOrderPrice(request *QueryOrderPriceRequest) (_result *QueryOrderPriceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &QueryOrderPriceResponse{}
	_body, _err := client.QueryOrderPriceWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 从人脸分组中的移除指定的文件
//
// @param request - RemoveFaceGroupFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveFaceGroupFileResponse
func (client *Client) RemoveFaceGroupFileWithOptions(request *RemoveFaceGroupFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveFaceGroupFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FaceGroupId)) {
		body["face_group_id"] = request.FaceGroupId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveFaceGroupFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/albums/unassign_facegroup_item"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveFaceGroupFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 从人脸分组中的移除指定的文件
//
// @param request - RemoveFaceGroupFileRequest
//
// @return RemoveFaceGroupFileResponse
func (client *Client) RemoveFaceGroupFile(request *RemoveFaceGroupFileRequest) (_result *RemoveFaceGroupFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveFaceGroupFileResponse{}
	_body, _err := client.RemoveFaceGroupFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Removes a member from a group.
//
// @param request - RemoveGroupMemberRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveGroupMemberResponse
func (client *Client) RemoveGroupMemberWithOptions(request *RemoveGroupMemberRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveGroupMemberResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["group_id"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberId)) {
		body["member_id"] = request.MemberId
	}

	if !tea.BoolValue(util.IsUnset(request.MemberType)) {
		body["member_type"] = request.MemberType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveGroupMember"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/group/remove_member"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveGroupMemberResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Removes a member from a group.
//
// @param request - RemoveGroupMemberRequest
//
// @return RemoveGroupMemberResponse
func (client *Client) RemoveGroupMember(request *RemoveGroupMemberRequest) (_result *RemoveGroupMemberResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveGroupMemberResponse{}
	_body, _err := client.RemoveGroupMemberWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 故事移除文件
//
// @param request - RemoveStoryFilesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RemoveStoryFilesResponse
func (client *Client) RemoveStoryFilesWithOptions(request *RemoveStoryFilesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RemoveStoryFilesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Files)) {
		body["files"] = request.Files
	}

	if !tea.BoolValue(util.IsUnset(request.StoryId)) {
		body["story_id"] = request.StoryId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RemoveStoryFiles"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/remove_story_files"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RemoveStoryFilesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 故事移除文件
//
// @param request - RemoveStoryFilesRequest
//
// @return RemoveStoryFilesResponse
func (client *Client) RemoveStoryFiles(request *RemoveStoryFilesRequest) (_result *RemoveStoryFilesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RemoveStoryFilesResponse{}
	_body, _err := client.RemoveStoryFilesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a file or folder from the recycle bin.
//
// @param request - RestoreFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreFileResponse
func (client *Client) RestoreFileWithOptions(request *RestoreFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestoreFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/recyclebin/restore"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a file or folder from the recycle bin.
//
// @param request - RestoreFileRequest
//
// @return RestoreFileResponse
func (client *Client) RestoreFile(request *RestoreFileRequest) (_result *RestoreFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestoreFileResponse{}
	_body, _err := client.RestoreFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Restores a historical version of a file. You cannot restore the latest version of a file.
//
// @param request - RestoreRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return RestoreRevisionResponse
func (client *Client) RestoreRevisionWithOptions(request *RestoreRevisionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *RestoreRevisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionId)) {
		body["revision_id"] = request.RevisionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("RestoreRevision"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/revision/restore"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &RestoreRevisionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Restores a historical version of a file. You cannot restore the latest version of a file.
//
// @param request - RestoreRevisionRequest
//
// @return RestoreRevisionResponse
func (client *Client) RestoreRevision(request *RestoreRevisionRequest) (_result *RestoreRevisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &RestoreRevisionResponse{}
	_body, _err := client.RestoreRevisionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Scans files.
//
// @param request - ScanFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return ScanFileResponse
func (client *Client) ScanFileWithOptions(request *ScanFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ScanFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ScanFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/scan"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ScanFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Scans files.
//
// @param request - ScanFileRequest
//
// @return ScanFileResponse
func (client *Client) ScanFile(request *ScanFileRequest) (_result *ScanFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ScanFileResponse{}
	_body, _err := client.ScanFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries location-based groups based on specific locations.
//
// @param request - SearchAddressGroupsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchAddressGroupsResponse
func (client *Client) SearchAddressGroupsWithOptions(request *SearchAddressGroupsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchAddressGroupsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.AddressLevel)) {
		body["address_level"] = request.AddressLevel
	}

	if !tea.BoolValue(util.IsUnset(request.AddressNames)) {
		body["address_names"] = request.AddressNames
	}

	if !tea.BoolValue(util.IsUnset(request.BrGeoPoint)) {
		body["br_geo_point"] = request.BrGeoPoint
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageThumbnailProcess)) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.TlGeoPoint)) {
		body["tl_geo_point"] = request.TlGeoPoint
	}

	if !tea.BoolValue(util.IsUnset(request.VideoThumbnailProcess)) {
		body["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchAddressGroups"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/search_address_groups"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchAddressGroupsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries location-based groups based on specific locations.
//
// @param request - SearchAddressGroupsRequest
//
// @return SearchAddressGroupsResponse
func (client *Client) SearchAddressGroups(request *SearchAddressGroupsRequest) (_result *SearchAddressGroupsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchAddressGroupsResponse{}
	_body, _err := client.SearchAddressGroupsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Search domain with specified attributes
//
// @param request - SearchDomainsRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDomainsResponse
func (client *Client) SearchDomainsWithOptions(request *SearchDomainsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchDomainsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["order_by"] = request.OrderBy
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchDomains"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/domain/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchDomainsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Search domain with specified attributes
//
// @param request - SearchDomainsRequest
//
// @return SearchDomainsResponse
func (client *Client) SearchDomains(request *SearchDomainsRequest) (_result *SearchDomainsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchDomainsResponse{}
	_body, _err := client.SearchDomainsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries drives.
//
// @param request - SearchDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchDriveResponse
func (client *Client) SearchDriveWithOptions(request *SearchDriveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchDriveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveName)) {
		body["drive_name"] = request.DriveName
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		body["owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.OwnerType)) {
		body["owner_type"] = request.OwnerType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchDrive"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/drive/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchDriveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries drives.
//
// @param request - SearchDriveRequest
//
// @return SearchDriveResponse
func (client *Client) SearchDrive(request *SearchDriveRequest) (_result *SearchDriveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchDriveResponse{}
	_body, _err := client.SearchDriveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries files. For more information about best practices, visit https://help.aliyun.com/document_detail/175890.html.
//
// @param request - SearchFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchFileResponse
func (client *Client) SearchFileWithOptions(request *SearchFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.Fields)) {
		body["fields"] = request.Fields
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["order_by"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.Recursive)) {
		body["recursive"] = request.Recursive
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnTotalCount)) {
		body["return_total_count"] = request.ReturnTotalCount
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries files. For more information about best practices, visit https://help.aliyun.com/document_detail/175890.html.
//
// @param request - SearchFileRequest
//
// @return SearchFileResponse
func (client *Client) SearchFile(request *SearchFileRequest) (_result *SearchFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchFileResponse{}
	_body, _err := client.SearchFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Queries share URLs.
//
// @param request - SearchShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchShareLinkResponse
func (client *Client) SearchShareLinkWithOptions(request *SearchShareLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchShareLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Creators)) {
		body["creators"] = request.Creators
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["order_by"] = request.OrderBy
	}

	if !tea.BoolValue(util.IsUnset(request.OrderDirection)) {
		body["order_direction"] = request.OrderDirection
	}

	if !tea.BoolValue(util.IsUnset(request.Query)) {
		body["query"] = request.Query
	}

	if !tea.BoolValue(util.IsUnset(request.ReturnTotalCount)) {
		body["return_total_count"] = request.ReturnTotalCount
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchShareLink"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/share_link/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchShareLinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Queries share URLs.
//
// @param request - SearchShareLinkRequest
//
// @return SearchShareLinkResponse
func (client *Client) SearchShareLink(request *SearchShareLinkRequest) (_result *SearchShareLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchShareLinkResponse{}
	_body, _err := client.SearchShareLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 获取相似图片聚类结果
//
// @param request - SearchSimilarImageClustersRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchSimilarImageClustersResponse
func (client *Client) SearchSimilarImageClustersWithOptions(request *SearchSimilarImageClustersRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchSimilarImageClustersResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.ImageThumbnailProcess)) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchSimilarImageClusters"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/query_similar_image_clusters"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchSimilarImageClustersResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 获取相似图片聚类结果
//
// @param request - SearchSimilarImageClustersRequest
//
// @return SearchSimilarImageClustersResponse
func (client *Client) SearchSimilarImageClusters(request *SearchSimilarImageClustersRequest) (_result *SearchSimilarImageClustersResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchSimilarImageClustersResponse{}
	_body, _err := client.SearchSimilarImageClustersWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 查询故事列表
//
// @param request - SearchStoriesRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchStoriesResponse
func (client *Client) SearchStoriesWithOptions(request *SearchStoriesRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchStoriesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CoverImageThumbnailProcess)) {
		body["cover_image_thumbnail_process"] = request.CoverImageThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.CoverVideoThumbnailProcess)) {
		body["cover_video_thumbnail_process"] = request.CoverVideoThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.CreateTimeRange)) {
		body["create_time_range"] = request.CreateTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.CustomLabels)) {
		body["custom_labels"] = request.CustomLabels
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FaceGroupIds)) {
		body["face_group_ids"] = request.FaceGroupIds
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.Order)) {
		body["order"] = request.Order
	}

	if !tea.BoolValue(util.IsUnset(request.Sort)) {
		body["sort"] = request.Sort
	}

	if !tea.BoolValue(util.IsUnset(request.StoryEndTimeRange)) {
		body["story_end_time_range"] = request.StoryEndTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.StoryId)) {
		body["story_id"] = request.StoryId
	}

	if !tea.BoolValue(util.IsUnset(request.StoryName)) {
		body["story_name"] = request.StoryName
	}

	if !tea.BoolValue(util.IsUnset(request.StoryStartTimeRange)) {
		body["story_start_time_range"] = request.StoryStartTimeRange
	}

	if !tea.BoolValue(util.IsUnset(request.StoryType)) {
		body["story_type"] = request.StoryType
	}

	if !tea.BoolValue(util.IsUnset(request.UrlExpireSec)) {
		body["url_expire_sec"] = request.UrlExpireSec
	}

	if !tea.BoolValue(util.IsUnset(request.WithEmptyStories)) {
		body["with_empty_stories"] = request.WithEmptyStories
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchStories"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/find_stories"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchStoriesResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 查询故事列表
//
// @param request - SearchStoriesRequest
//
// @return SearchStoriesResponse
func (client *Client) SearchStories(request *SearchStoriesRequest) (_result *SearchStoriesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchStoriesResponse{}
	_body, _err := client.SearchStoriesWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Searches for users.
//
// @param request - SearchUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return SearchUserResponse
func (client *Client) SearchUserWithOptions(request *SearchUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.Limit)) {
		body["limit"] = request.Limit
	}

	if !tea.BoolValue(util.IsUnset(request.Marker)) {
		body["marker"] = request.Marker
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		body["nick_name"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.NickNameForFuzzy)) {
		body["nick_name_for_fuzzy"] = request.NickNameForFuzzy
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserName)) {
		body["user_name"] = request.UserName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("SearchUser"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/user/search"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &SearchUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Searches for users.
//
// @param request - SearchUserRequest
//
// @return SearchUserResponse
func (client *Client) SearchUser(request *SearchUserRequest) (_result *SearchUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &SearchUserResponse{}
	_body, _err := client.SearchUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Generates an access token based on Open Authorization (OAuth) 2.0.
//
// Description:
//
// For more information about how to access Drive and Photo Service from a web server application by using OAuth 2.0, visit [OAuth 2.0 For Web Server Applications](https://www.alibabacloud.com/help/zh/pds/drive-and-photo-service-dev/user-guide/oauth-2-0-access-process-for-web-server-applications).
//
// For more information about how to access Drive and Photo Service by using a JSON Web Token (JWT) application, visit [Access process for JWT applications](https://www.alibabacloud.com/help/zh/pds/drive-and-photo-service-dev/user-guide/access-process-for-jwt-applications).
//
// @param request - TokenRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TokenResponse
func (client *Client) TokenWithOptions(request *TokenRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TokenResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Assertion)) {
		body["assertion"] = request.Assertion
	}

	if !tea.BoolValue(util.IsUnset(request.ClientId)) {
		body["client_id"] = request.ClientId
	}

	if !tea.BoolValue(util.IsUnset(request.ClientSecret)) {
		body["client_secret"] = request.ClientSecret
	}

	if !tea.BoolValue(util.IsUnset(request.Code)) {
		body["code"] = request.Code
	}

	if !tea.BoolValue(util.IsUnset(request.GrantType)) {
		body["grant_type"] = request.GrantType
	}

	if !tea.BoolValue(util.IsUnset(request.RedirectUri)) {
		body["redirect_uri"] = request.RedirectUri
	}

	if !tea.BoolValue(util.IsUnset(request.RefreshToken)) {
		body["refresh_token"] = request.RefreshToken
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("Token"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/oauth/token"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("Anonymous"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("formData"),
		BodyType:    tea.String("json"),
	}
	_result = &TokenResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Generates an access token based on Open Authorization (OAuth) 2.0.
//
// Description:
//
// For more information about how to access Drive and Photo Service from a web server application by using OAuth 2.0, visit [OAuth 2.0 For Web Server Applications](https://www.alibabacloud.com/help/zh/pds/drive-and-photo-service-dev/user-guide/oauth-2-0-access-process-for-web-server-applications).
//
// For more information about how to access Drive and Photo Service by using a JSON Web Token (JWT) application, visit [Access process for JWT applications](https://www.alibabacloud.com/help/zh/pds/drive-and-photo-service-dev/user-guide/access-process-for-jwt-applications).
//
// @param request - TokenRequest
//
// @return TokenResponse
func (client *Client) Token(request *TokenRequest) (_result *TokenResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TokenResponse{}
	_body, _err := client.TokenWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Moves a file or folder to the recycle bin.
//
// @param request - TrashFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return TrashFileResponse
func (client *Client) TrashFileWithOptions(request *TrashFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *TrashFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("TrashFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/recyclebin/trash"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &TrashFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Moves a file or folder to the recycle bin.
//
// @param request - TrashFileRequest
//
// @return TrashFileResponse
func (client *Client) TrashFile(request *TrashFileRequest) (_result *TrashFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &TrashFileResponse{}
	_body, _err := client.TrashFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Unlink Account Binding
//
// @param request - UnLinkAccountRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UnLinkAccountResponse
func (client *Client) UnLinkAccountWithOptions(request *UnLinkAccountRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UnLinkAccountResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Extra)) {
		body["extra"] = request.Extra
	}

	if !tea.BoolValue(util.IsUnset(request.Identity)) {
		body["identity"] = request.Identity
	}

	if !tea.BoolValue(util.IsUnset(request.Type)) {
		body["type"] = request.Type
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UnLinkAccount"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/account/unlink"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UnLinkAccountResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Unlink Account Binding
//
// @param request - UnLinkAccountRequest
//
// @return UnLinkAccountResponse
func (client *Client) UnLinkAccount(request *UnLinkAccountRequest) (_result *UnLinkAccountResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UnLinkAccountResponse{}
	_body, _err := client.UnLinkAccountWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Update domain information.
//
// @param request - UpdateDomainRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomainWithOptions(request *UpdateDomainRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDomainResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DomainId)) {
		body["domain_id"] = request.DomainId
	}

	if !tea.BoolValue(util.IsUnset(request.DomainName)) {
		body["domain_name"] = request.DomainName
	}

	if !tea.BoolValue(util.IsUnset(request.InitDriveEnable)) {
		body["init_drive_enable"] = request.InitDriveEnable
	}

	if !tea.BoolValue(util.IsUnset(request.InitDriveSize)) {
		body["init_drive_size"] = request.InitDriveSize
	}

	if !tea.BoolValue(util.IsUnset(request.PublishedAppAccessStrategy)) {
		body["published_app_access_strategy"] = request.PublishedAppAccessStrategy
	}

	if !tea.BoolValue(util.IsUnset(request.SizeQuota)) {
		body["size_quota"] = request.SizeQuota
	}

	if !tea.BoolValue(util.IsUnset(request.UserCountQuota)) {
		body["user_count_quota"] = request.UserCountQuota
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDomain"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/domain/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDomainResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Update domain information.
//
// @param request - UpdateDomainRequest
//
// @return UpdateDomainResponse
func (client *Client) UpdateDomain(request *UpdateDomainRequest) (_result *UpdateDomainResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDomainResponse{}
	_body, _err := client.UpdateDomainWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a drive.
//
// @param request - UpdateDriveRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateDriveResponse
func (client *Client) UpdateDriveWithOptions(request *UpdateDriveRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateDriveResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.DriveName)) {
		body["drive_name"] = request.DriveName
	}

	if !tea.BoolValue(util.IsUnset(request.Owner)) {
		body["owner"] = request.Owner
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.TotalSize)) {
		body["total_size"] = request.TotalSize
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateDrive"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/drive/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateDriveResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a drive.
//
// @param request - UpdateDriveRequest
//
// @return UpdateDriveResponse
func (client *Client) UpdateDrive(request *UpdateDriveRequest) (_result *UpdateDriveResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateDriveResponse{}
	_body, _err := client.UpdateDriveWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates a face-based group.
//
// @param request - UpdateFacegroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFacegroupResponse
func (client *Client) UpdateFacegroupWithOptions(request *UpdateFacegroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFacegroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupCoverFaceId)) {
		body["group_cover_face_id"] = request.GroupCoverFaceId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["group_id"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["group_name"] = request.GroupName
	}

	if !tea.BoolValue(util.IsUnset(request.Remarks)) {
		body["remarks"] = request.Remarks
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFacegroup"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/update_facegroup_info"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFacegroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates a face-based group.
//
// @param request - UpdateFacegroupRequest
//
// @return UpdateFacegroupResponse
func (client *Client) UpdateFacegroup(request *UpdateFacegroupRequest) (_result *UpdateFacegroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFacegroupResponse{}
	_body, _err := client.UpdateFacegroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a file instead of the file data.
//
// @param request - UpdateFileRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateFileResponse
func (client *Client) UpdateFileWithOptions(request *UpdateFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateFileResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.CheckNameMode)) {
		body["check_name_mode"] = request.CheckNameMode
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.Hidden)) {
		body["hidden"] = request.Hidden
	}

	if !tea.BoolValue(util.IsUnset(request.Labels)) {
		body["labels"] = request.Labels
	}

	if !tea.BoolValue(util.IsUnset(request.LocalModifiedAt)) {
		body["local_modified_at"] = request.LocalModifiedAt
	}

	if !tea.BoolValue(util.IsUnset(request.Name)) {
		body["name"] = request.Name
	}

	if !tea.BoolValue(util.IsUnset(request.Starred)) {
		body["starred"] = request.Starred
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateFile"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a file instead of the file data.
//
// @param request - UpdateFileRequest
//
// @return UpdateFileResponse
func (client *Client) UpdateFile(request *UpdateFileRequest) (_result *UpdateFileResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateFileResponse{}
	_body, _err := client.UpdateFileWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a group.
//
// @param request - UpdateGroupRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroupWithOptions(request *UpdateGroupRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateGroupResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.GroupId)) {
		body["group_id"] = request.GroupId
	}

	if !tea.BoolValue(util.IsUnset(request.GroupName)) {
		body["group_name"] = request.GroupName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateGroup"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/group/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateGroupResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a group.
//
// @param request - UpdateGroupRequest
//
// @return UpdateGroupResponse
func (client *Client) UpdateGroup(request *UpdateGroupRequest) (_result *UpdateGroupResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateGroupResponse{}
	_body, _err := client.UpdateGroupWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the mapping between an entity and a benefit package. You can call this operation to associate a benefit package with a user.
//
// @param request - UpdateIdentityToBenefitPkgMappingRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateIdentityToBenefitPkgMappingResponse
func (client *Client) UpdateIdentityToBenefitPkgMappingWithOptions(request *UpdateIdentityToBenefitPkgMappingRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateIdentityToBenefitPkgMappingResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Amount)) {
		body["amount"] = request.Amount
	}

	if !tea.BoolValue(util.IsUnset(request.BenefitPkgId)) {
		body["benefit_pkg_id"] = request.BenefitPkgId
	}

	if !tea.BoolValue(util.IsUnset(request.ExpireTime)) {
		body["expire_time"] = request.ExpireTime
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityId)) {
		body["identity_id"] = request.IdentityId
	}

	if !tea.BoolValue(util.IsUnset(request.IdentityType)) {
		body["identity_type"] = request.IdentityType
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateIdentityToBenefitPkgMapping"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/benefit/identity_to_benefit_pkg_mapping/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the mapping between an entity and a benefit package. You can call this operation to associate a benefit package with a user.
//
// @param request - UpdateIdentityToBenefitPkgMappingRequest
//
// @return UpdateIdentityToBenefitPkgMappingResponse
func (client *Client) UpdateIdentityToBenefitPkgMapping(request *UpdateIdentityToBenefitPkgMappingRequest) (_result *UpdateIdentityToBenefitPkgMappingResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateIdentityToBenefitPkgMappingResponse{}
	_body, _err := client.UpdateIdentityToBenefitPkgMappingWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Updates the version information. You can call this operation to permanently retain a version or modify the description of a version. You can permanently retain up to 50 versions of a file.
//
// @param request - UpdateRevisionRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateRevisionResponse
func (client *Client) UpdateRevisionWithOptions(request *UpdateRevisionRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateRevisionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.FileId)) {
		body["file_id"] = request.FileId
	}

	if !tea.BoolValue(util.IsUnset(request.KeepForever)) {
		body["keep_forever"] = request.KeepForever
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionDescription)) {
		body["revision_description"] = request.RevisionDescription
	}

	if !tea.BoolValue(util.IsUnset(request.RevisionId)) {
		body["revision_id"] = request.RevisionId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateRevision"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/file/revision/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateRevisionResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Updates the version information. You can call this operation to permanently retain a version or modify the description of a version. You can permanently retain up to 50 versions of a file.
//
// @param request - UpdateRevisionRequest
//
// @return UpdateRevisionResponse
func (client *Client) UpdateRevision(request *UpdateRevisionRequest) (_result *UpdateRevisionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateRevisionResponse{}
	_body, _err := client.UpdateRevisionWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies a share link.
//
// @param request - UpdateShareLinkRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateShareLinkResponse
func (client *Client) UpdateShareLinkWithOptions(request *UpdateShareLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateShareLinkResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.DisableDownload)) {
		body["disable_download"] = request.DisableDownload
	}

	if !tea.BoolValue(util.IsUnset(request.DisablePreview)) {
		body["disable_preview"] = request.DisablePreview
	}

	if !tea.BoolValue(util.IsUnset(request.DisableSave)) {
		body["disable_save"] = request.DisableSave
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadCount)) {
		body["download_count"] = request.DownloadCount
	}

	if !tea.BoolValue(util.IsUnset(request.DownloadLimit)) {
		body["download_limit"] = request.DownloadLimit
	}

	if !tea.BoolValue(util.IsUnset(request.Expiration)) {
		body["expiration"] = request.Expiration
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeEditable)) {
		body["office_editable"] = request.OfficeEditable
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewCount)) {
		body["preview_count"] = request.PreviewCount
	}

	if !tea.BoolValue(util.IsUnset(request.PreviewLimit)) {
		body["preview_limit"] = request.PreviewLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ReportCount)) {
		body["report_count"] = request.ReportCount
	}

	if !tea.BoolValue(util.IsUnset(request.SaveCount)) {
		body["save_count"] = request.SaveCount
	}

	if !tea.BoolValue(util.IsUnset(request.SaveLimit)) {
		body["save_limit"] = request.SaveLimit
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.ShareName)) {
		body["share_name"] = request.ShareName
	}

	if !tea.BoolValue(util.IsUnset(request.SharePwd)) {
		body["share_pwd"] = request.SharePwd
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.VideoPreviewCount)) {
		body["video_preview_count"] = request.VideoPreviewCount
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateShareLink"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/share_link/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateShareLinkResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies a share link.
//
// @param request - UpdateShareLinkRequest
//
// @return UpdateShareLinkResponse
func (client *Client) UpdateShareLink(request *UpdateShareLinkRequest) (_result *UpdateShareLinkResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateShareLinkResponse{}
	_body, _err := client.UpdateShareLinkWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// 更新故事
//
// @param request - UpdateStoryRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateStoryResponse
func (client *Client) UpdateStoryWithOptions(request *UpdateStoryRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateStoryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Cover)) {
		body["cover"] = request.Cover
	}

	if !tea.BoolValue(util.IsUnset(request.CustomLabels)) {
		body["custom_labels"] = request.CustomLabels
	}

	if !tea.BoolValue(util.IsUnset(request.DriveId)) {
		body["drive_id"] = request.DriveId
	}

	if !tea.BoolValue(util.IsUnset(request.StoryId)) {
		body["story_id"] = request.StoryId
	}

	if !tea.BoolValue(util.IsUnset(request.StoryName)) {
		body["story_name"] = request.StoryName
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateStory"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/update_story"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateStoryResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// 更新故事
//
// @param request - UpdateStoryRequest
//
// @return UpdateStoryResponse
func (client *Client) UpdateStory(request *UpdateStoryRequest) (_result *UpdateStoryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateStoryResponse{}
	_body, _err := client.UpdateStoryWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

// Summary:
//
// Modifies the information about a user.
//
// @param request - UpdateUserRequest
//
// @param headers - map
//
// @param runtime - runtime options for this request RuntimeOptions
//
// @return UpdateUserResponse
func (client *Client) UpdateUserWithOptions(request *UpdateUserRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *UpdateUserResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Avatar)) {
		body["avatar"] = request.Avatar
	}

	if !tea.BoolValue(util.IsUnset(request.Description)) {
		body["description"] = request.Description
	}

	if !tea.BoolValue(util.IsUnset(request.Email)) {
		body["email"] = request.Email
	}

	if !tea.BoolValue(util.IsUnset(request.GroupInfoList)) {
		body["group_info_list"] = request.GroupInfoList
	}

	if !tea.BoolValue(util.IsUnset(request.NickName)) {
		body["nick_name"] = request.NickName
	}

	if !tea.BoolValue(util.IsUnset(request.Phone)) {
		body["phone"] = request.Phone
	}

	if !tea.BoolValue(util.IsUnset(request.Role)) {
		body["role"] = request.Role
	}

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
	}

	if !tea.BoolValue(util.IsUnset(request.UserData)) {
		body["user_data"] = request.UserData
	}

	if !tea.BoolValue(util.IsUnset(request.UserId)) {
		body["user_id"] = request.UserId
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("UpdateUser"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/user/update"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &UpdateUserResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

// Summary:
//
// Modifies the information about a user.
//
// @param request - UpdateUserRequest
//
// @return UpdateUserResponse
func (client *Client) UpdateUser(request *UpdateUserRequest) (_result *UpdateUserResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &UpdateUserResponse{}
	_body, _err := client.UpdateUserWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
