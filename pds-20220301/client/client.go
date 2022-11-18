// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	gatewayclient "github.com/alibabacloud-go/alibabacloud-gateway-pds/client"
	
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	openapiutil "github.com/alibabacloud-go/openapi-util/service"
	util "github.com/alibabacloud-go/tea-utils/v2/service"
	"github.com/alibabacloud-go/tea/tea"
)

type AccountLinkInfo struct {
	AuthenticationType *string `json:"authentication_type,omitempty" xml:"authentication_type,omitempty"`
	CreatedAt          *int64  `json:"created_at,omitempty" xml:"created_at,omitempty"`
	DisplayName        *string `json:"display_name,omitempty" xml:"display_name,omitempty"`
	DomainId           *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	Extra              *string `json:"extra,omitempty" xml:"extra,omitempty"`
	Identity           *string `json:"identity,omitempty" xml:"identity,omitempty"`
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

type Address struct {
	City     *string `json:"city,omitempty" xml:"city,omitempty"`
	Country  *string `json:"country,omitempty" xml:"country,omitempty"`
	District *string `json:"district,omitempty" xml:"district,omitempty"`
	Province *string `json:"province,omitempty" xml:"province,omitempty"`
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

type FaceGroup struct {
	CreatedAt              *string                          `json:"created_at,omitempty" xml:"created_at,omitempty"`
	GroupCoverFaceBoundary *FaceGroupGroupCoverFaceBoundary `json:"group_cover_face_boundary,omitempty" xml:"group_cover_face_boundary,omitempty" type:"Struct"`
	GroupCoverFileId       *string                          `json:"group_cover_file_id,omitempty" xml:"group_cover_file_id,omitempty"`
	GroupCoverHeight       *int64                           `json:"group_cover_height,omitempty" xml:"group_cover_height,omitempty"`
	GroupCoverUrl          *string                          `json:"group_cover_url,omitempty" xml:"group_cover_url,omitempty"`
	GroupCoverWidth        *int64                           `json:"group_cover_width,omitempty" xml:"group_cover_width,omitempty"`
	GroupId                *string                          `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName              *string                          `json:"group_name,omitempty" xml:"group_name,omitempty"`
	ImageCount             *int64                           `json:"image_count,omitempty" xml:"image_count,omitempty"`
	Remarks                *string                          `json:"remarks,omitempty" xml:"remarks,omitempty"`
	UpdatedAt              *string                          `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
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
	Height *int32 `json:"height,omitempty" xml:"height,omitempty"`
	Left   *int32 `json:"left,omitempty" xml:"left,omitempty"`
	Top    *int32 `json:"top,omitempty" xml:"top,omitempty"`
	Width  *int32 `json:"width,omitempty" xml:"width,omitempty"`
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

type File struct {
	Category        *string `json:"category,omitempty" xml:"category,omitempty"`
	ContentHash     *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	ContentType     *string `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Crc64Hash       *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	CreatedAt       *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Description     *string `json:"description,omitempty" xml:"description,omitempty"`
	DomainId        *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DownloadUrl     *string `json:"download_url,omitempty" xml:"download_url,omitempty"`
	DriveId         *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileExtension   *string `json:"file_extension,omitempty" xml:"file_extension,omitempty"`
	FileId          *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Hidden          *bool   `json:"hidden,omitempty" xml:"hidden,omitempty"`
	Labels          *string `json:"labels,omitempty" xml:"labels,omitempty"`
	LocalCreatedAt  *string `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	LocalModifiedAt *string `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	Name            *string `json:"name,omitempty" xml:"name,omitempty"`
	ParentFileId    *string `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	RevisionId      *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	Size            *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Starred         *bool   `json:"starred,omitempty" xml:"starred,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	Thumbnail       *string `json:"thumbnail,omitempty" xml:"thumbnail,omitempty"`
	TrashedAt       *string `json:"trashed_at,omitempty" xml:"trashed_at,omitempty"`
	Type            *string `json:"type,omitempty" xml:"type,omitempty"`
	UpdatedAt       *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	UploadId        *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
}

func (s File) String() string {
	return tea.Prettify(s)
}

func (s File) GoString() string {
	return s.String()
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

func (s *File) SetLabels(v string) *File {
	s.Labels = &v
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

type Identity struct {
	IdentityId   *string `json:"identity_id,omitempty" xml:"identity_id,omitempty"`
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

type ImageMediaMetadata struct {
	Height  *int64  `json:"height,omitempty" xml:"height,omitempty"`
	TakenAt *string `json:"taken_at,omitempty" xml:"taken_at,omitempty"`
	Width   *int64  `json:"width,omitempty" xml:"width,omitempty"`
}

func (s ImageMediaMetadata) String() string {
	return tea.Prettify(s)
}

func (s ImageMediaMetadata) GoString() string {
	return s.String()
}

func (s *ImageMediaMetadata) SetHeight(v int64) *ImageMediaMetadata {
	s.Height = &v
	return s
}

func (s *ImageMediaMetadata) SetTakenAt(v string) *ImageMediaMetadata {
	s.TakenAt = &v
	return s
}

func (s *ImageMediaMetadata) SetWidth(v int64) *ImageMediaMetadata {
	s.Width = &v
	return s
}

type ImageTag struct {
	Count              *int64   `json:"count,omitempty" xml:"count,omitempty"`
	CoverFileCategory  *string  `json:"cover_file_category,omitempty" xml:"cover_file_category,omitempty"`
	CoverFileId        *string  `json:"cover_file_id,omitempty" xml:"cover_file_id,omitempty"`
	CoverOverallScore  *float32 `json:"cover_overall_score,omitempty" xml:"cover_overall_score,omitempty"`
	CoverTagConfidence *float32 `json:"cover_tag_confidence,omitempty" xml:"cover_tag_confidence,omitempty"`
	CoverUrl           *string  `json:"cover_url,omitempty" xml:"cover_url,omitempty"`
	Name               *string  `json:"name,omitempty" xml:"name,omitempty"`
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

type Revision struct {
	ContentHash         *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName     *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	Crc64Hash           *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	CreatedAt           *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
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

type ShareLink struct {
	AccessCount       *int64  `json:"access_count,omitempty" xml:"access_count,omitempty"`
	CreatedAt         *string `json:"created_at,omitempty" xml:"created_at,omitempty"`
	Creator           *string `json:"creator,omitempty" xml:"creator,omitempty"`
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	DisableDownload   *bool   `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	DisablePreview    *bool   `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	DisableSave       *bool   `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	DownloadCount     *int64  `json:"download_count,omitempty" xml:"download_count,omitempty"`
	DownloadLimit     *int64  `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	DriveId           *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Expiration        *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	Expired           *bool   `json:"expired,omitempty" xml:"expired,omitempty"`
	FileIdList        *string `json:"file_id_list,omitempty" xml:"file_id_list,omitempty"`
	PreviewCount      *int64  `json:"preview_count,omitempty" xml:"preview_count,omitempty"`
	PreviewLimit      *int64  `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	ReportCount       *int64  `json:"report_count,omitempty" xml:"report_count,omitempty"`
	SaveCount         *int64  `json:"save_count,omitempty" xml:"save_count,omitempty"`
	SaveLimit         *int64  `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	ShareId           *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	ShareName         *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	SharePwd          *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
	Status            *string `json:"status,omitempty" xml:"status,omitempty"`
	UpdatedAt         *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	VideoPreviewCount *int64  `json:"video_preview_count,omitempty" xml:"video_preview_count,omitempty"`
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

func (s *ShareLink) SetFileIdList(v string) *ShareLink {
	s.FileIdList = &v
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

func (s *ShareLink) SetSaveLimit(v int64) *ShareLink {
	s.SaveLimit = &v
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

type SystemTag struct {
	Confidence *float32 `json:"confidence,omitempty" xml:"confidence,omitempty"`
	Name       *string  `json:"name,omitempty" xml:"name,omitempty"`
	ParentName *string  `json:"parent_name,omitempty" xml:"parent_name,omitempty"`
	Source     *string  `json:"source,omitempty" xml:"source,omitempty"`
	TagLevel   *int32   `json:"tag_level,omitempty" xml:"tag_level,omitempty"`
}

func (s SystemTag) String() string {
	return tea.Prettify(s)
}

func (s SystemTag) GoString() string {
	return s.String()
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

func (s *SystemTag) SetSource(v string) *SystemTag {
	s.Source = &v
	return s
}

func (s *SystemTag) SetTagLevel(v int32) *SystemTag {
	s.TagLevel = &v
	return s
}

type Token struct {
	AccessToken    *string `json:"access_token,omitempty" xml:"access_token,omitempty"`
	Avatar         *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	DefaultDriveId *string `json:"default_drive_id,omitempty" xml:"default_drive_id,omitempty"`
	DeviceId       *string `json:"device_id,omitempty" xml:"device_id,omitempty"`
	DeviceName     *string `json:"device_name,omitempty" xml:"device_name,omitempty"`
	DomainId       *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	ExpireTime     *string `json:"expire_time,omitempty" xml:"expire_time,omitempty"`
	ExpiresIn      *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
	IsFirstLogin   *bool   `json:"is_first_login,omitempty" xml:"is_first_login,omitempty"`
	NickName       *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	RefreshToken   *string `json:"refresh_token,omitempty" xml:"refresh_token,omitempty"`
	Role           *string `json:"role,omitempty" xml:"role,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	TokenType      *string `json:"token_type,omitempty" xml:"token_type,omitempty"`
	UserId         *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName       *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
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

func (s *Token) SetNickName(v string) *Token {
	s.NickName = &v
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

func (s *Token) SetStatus(v string) *Token {
	s.Status = &v
	return s
}

func (s *Token) SetTokenType(v string) *Token {
	s.TokenType = &v
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

type UploadPartInfo struct {
	Etag              *string                        `json:"etag,omitempty" xml:"etag,omitempty"`
	InternalUploadUrl *string                        `json:"internal_upload_url,omitempty" xml:"internal_upload_url,omitempty"`
	ParallelSha1Ctx   *UploadPartInfoParallelSha1Ctx `json:"parallel_sha1_ctx,omitempty" xml:"parallel_sha1_ctx,omitempty" type:"Struct"`
	PartNumber        *int32                         `json:"part_number,omitempty" xml:"part_number,omitempty"`
	PartSize          *int64                         `json:"part_size,omitempty" xml:"part_size,omitempty"`
	UploadUrl         *string                        `json:"upload_url,omitempty" xml:"upload_url,omitempty"`
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

type UserTag struct {
	Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UserTag) String() string {
	return tea.Prettify(s)
}

func (s UserTag) GoString() string {
	return s.String()
}

func (s *UserTag) SetValue(v string) *UserTag {
	s.Value = &v
	return s
}

func (s *UserTag) SetKey(v string) *UserTag {
	s.Key = &v
	return s
}

type VideoMediaMetadata struct {
	Duration *string `json:"duration,omitempty" xml:"duration,omitempty"`
	TakenAt  *string `json:"taken_at,omitempty" xml:"taken_at,omitempty"`
}

func (s VideoMediaMetadata) String() string {
	return tea.Prettify(s)
}

func (s VideoMediaMetadata) GoString() string {
	return s.String()
}

func (s *VideoMediaMetadata) SetDuration(v string) *VideoMediaMetadata {
	s.Duration = &v
	return s
}

func (s *VideoMediaMetadata) SetTakenAt(v string) *VideoMediaMetadata {
	s.TakenAt = &v
	return s
}

type VideoPreviewPlayInfo struct {
	Category                *string                                        `json:"category,omitempty" xml:"category,omitempty"`
	LiveTranscodingTaskList []*VideoPreviewPlayInfoLiveTranscodingTaskList `json:"live_transcoding_task_list,omitempty" xml:"live_transcoding_task_list,omitempty" type:"Repeated"`
	Meta                    *VideoPreviewPlayInfoMeta                      `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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

func (s *VideoPreviewPlayInfo) SetLiveTranscodingTaskList(v []*VideoPreviewPlayInfoLiveTranscodingTaskList) *VideoPreviewPlayInfo {
	s.LiveTranscodingTaskList = v
	return s
}

func (s *VideoPreviewPlayInfo) SetMeta(v *VideoPreviewPlayInfoMeta) *VideoPreviewPlayInfo {
	s.Meta = v
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

type VideoPreviewPlayMeta struct {
	Category                *string                                        `json:"category,omitempty" xml:"category,omitempty"`
	LiveTranscodingTaskList []*VideoPreviewPlayMetaLiveTranscodingTaskList `json:"live_transcoding_task_list,omitempty" xml:"live_transcoding_task_list,omitempty" type:"Repeated"`
	Meta                    *VideoPreviewPlayMetaMeta                      `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
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

type VideoPreviewPlayMetaLiveTranscodingTaskList struct {
	KeepOriginalResolution *bool   `json:"keep_original_resolution,omitempty" xml:"keep_original_resolution,omitempty"`
	Status                 *string `json:"status,omitempty" xml:"status,omitempty"`
	TemplateId             *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
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
	Duration *float64 `json:"duration,omitempty" xml:"duration,omitempty"`
	Height   *int64   `json:"height,omitempty" xml:"height,omitempty"`
	Width    *int64   `json:"width,omitempty" xml:"width,omitempty"`
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

type AuthorizeRequest struct {
	ClientId     *string   `json:"client_id,omitempty" xml:"client_id,omitempty"`
	HideConsent  *bool     `json:"hide_consent,omitempty" xml:"hide_consent,omitempty"`
	LoginType    *string   `json:"login_type,omitempty" xml:"login_type,omitempty"`
	RedirectUri  *string   `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	ResponseType *string   `json:"response_type,omitempty" xml:"response_type,omitempty"`
	Scope        []*string `json:"scope,omitempty" xml:"scope,omitempty" type:"Repeated"`
	State        *string   `json:"state,omitempty" xml:"state,omitempty"`
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
	ClientId     *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	HideConsent  *bool   `json:"hide_consent,omitempty" xml:"hide_consent,omitempty"`
	LoginType    *string `json:"login_type,omitempty" xml:"login_type,omitempty"`
	RedirectUri  *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
	ResponseType *string `json:"response_type,omitempty" xml:"response_type,omitempty"`
	ScopeShrink  *string `json:"scope,omitempty" xml:"scope,omitempty"`
	State        *string `json:"state,omitempty" xml:"state,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	Requests []*BatchRequestRequests `json:"requests,omitempty" xml:"requests,omitempty" type:"Repeated"`
	Resource *string                 `json:"resource,omitempty" xml:"resource,omitempty"`
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
	Body    map[string]*string `json:"body,omitempty" xml:"body,omitempty"`
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty"`
	Id      *string            `json:"id,omitempty" xml:"id,omitempty"`
	Method  *string            `json:"method,omitempty" xml:"method,omitempty"`
	Url     *string            `json:"url,omitempty" xml:"url,omitempty"`
}

func (s BatchRequestRequests) String() string {
	return tea.Prettify(s)
}

func (s BatchRequestRequests) GoString() string {
	return s.String()
}

func (s *BatchRequestRequests) SetBody(v map[string]*string) *BatchRequestRequests {
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
	Body   map[string]*string `json:"body,omitempty" xml:"body,omitempty"`
	Id     *string            `json:"id,omitempty" xml:"id,omitempty"`
	Status *int32             `json:"status,omitempty" xml:"status,omitempty"`
}

func (s BatchResponseBodyResponses) String() string {
	return tea.Prettify(s)
}

func (s BatchResponseBodyResponses) GoString() string {
	return s.String()
}

func (s *BatchResponseBodyResponses) SetBody(v map[string]*string) *BatchResponseBodyResponses {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *BatchResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CancelShareLinkRequest struct {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ClearRecyclebinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId  *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId   *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *File              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AutoRename     *bool   `json:"auto_rename,omitempty" xml:"auto_rename,omitempty"`
	DriveId        *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId         *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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

func (s *CopyFileRequest) SetToParentFileId(v string) *CopyFileRequest {
	s.ToParentFileId = &v
	return s
}

type CopyFileResponseBody struct {
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId      *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CopyFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateDriveRequest struct {
	Default     *bool   `json:"default,omitempty" xml:"default,omitempty"`
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DriveName   *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	DriveType   *string `json:"drive_type,omitempty" xml:"drive_type,omitempty"`
	Owner       *string `json:"owner,omitempty" xml:"owner,omitempty"`
	OwnerType   *string `json:"owner_type,omitempty" xml:"owner_type,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	TotalSize   *int64  `json:"total_size,omitempty" xml:"total_size,omitempty"`
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
	DomainId *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId  *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
}

func (s CreateDriveResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateDriveResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDriveResponseBody) SetDomainId(v string) *CreateDriveResponseBody {
	s.DomainId = &v
	return s
}

func (s *CreateDriveResponseBody) SetDriveId(v string) *CreateDriveResponseBody {
	s.DriveId = &v
	return s
}

type CreateDriveResponse struct {
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateDriveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CheckNameMode      *string                          `json:"check_name_mode,omitempty" xml:"check_name_mode,omitempty"`
	ContentHash        *string                          `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName    *string                          `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	ContentType        *string                          `json:"content_type,omitempty" xml:"content_type,omitempty"`
	Description        *string                          `json:"description,omitempty" xml:"description,omitempty"`
	DriveId            *string                          `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId             *string                          `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Hidden             *bool                            `json:"hidden,omitempty" xml:"hidden,omitempty"`
	ImageMediaMetadata *ImageMediaMetadata              `json:"image_media_metadata,omitempty" xml:"image_media_metadata,omitempty"`
	LocalCreatedAt     *string                          `json:"local_created_at,omitempty" xml:"local_created_at,omitempty"`
	LocalModifiedAt    *string                          `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	Name               *string                          `json:"name,omitempty" xml:"name,omitempty"`
	ParallelUpload     *bool                            `json:"parallel_upload,omitempty" xml:"parallel_upload,omitempty"`
	ParentFileId       *string                          `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	PartInfoList       []*CreateFileRequestPartInfoList `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	PreHash            *string                          `json:"pre_hash,omitempty" xml:"pre_hash,omitempty"`
	ShareId            *string                          `json:"share_id,omitempty" xml:"share_id,omitempty"`
	Size               *int64                           `json:"size,omitempty" xml:"size,omitempty"`
	Type               *string                          `json:"type,omitempty" xml:"type,omitempty"`
	UserTags           []*UserTag                       `json:"user_tags,omitempty" xml:"user_tags,omitempty" type:"Repeated"`
	VideoMediaMetadata *VideoMediaMetadata              `json:"video_media_metadata,omitempty" xml:"video_media_metadata,omitempty"`
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
	PartNumber *int32 `json:"part_number,omitempty" xml:"part_number,omitempty"`
}

func (s CreateFileRequestPartInfoList) String() string {
	return tea.Prettify(s)
}

func (s CreateFileRequestPartInfoList) GoString() string {
	return s.String()
}

func (s *CreateFileRequestPartInfoList) SetPartNumber(v int32) *CreateFileRequestPartInfoList {
	s.PartNumber = &v
	return s
}

type CreateFileResponseBody struct {
	DomainId     *string           `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId      *string           `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Exist        *bool             `json:"exist,omitempty" xml:"exist,omitempty"`
	FileId       *string           `json:"file_id,omitempty" xml:"file_id,omitempty"`
	FileName     *string           `json:"file_name,omitempty" xml:"file_name,omitempty"`
	ParentFileId *string           `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	PartInfoList []*UploadPartInfo `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	RapidUpload  *bool             `json:"rapid_upload,omitempty" xml:"rapid_upload,omitempty"`
	Status       *string           `json:"status,omitempty" xml:"status,omitempty"`
	Type         *string           `json:"type,omitempty" xml:"type,omitempty"`
	UploadId     *string           `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description   *string `json:"description,omitempty" xml:"description,omitempty"`
	GroupName     *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	IsRoot        *bool   `json:"is_root,omitempty" xml:"is_root,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Group             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateShareLinkRequest struct {
	Description     *string   `json:"description,omitempty" xml:"description,omitempty"`
	DisableDownload *bool     `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	DisablePreview  *bool     `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	DisableSave     *bool     `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	DownloadLimit   *int64    `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	DriveId         *string   `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Expiration      *string   `json:"expiration,omitempty" xml:"expiration,omitempty"`
	FileIdList      []*string `json:"file_id_list,omitempty" xml:"file_id_list,omitempty" type:"Repeated"`
	PreviewLimit    *int64    `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	SaveLimit       *int64    `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	ShareName       *string   `json:"share_name,omitempty" xml:"share_name,omitempty"`
	SharePwd        *string   `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
	UserId          *string   `json:"user_id,omitempty" xml:"user_id,omitempty"`
}

func (s CreateShareLinkRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateShareLinkRequest) GoString() string {
	return s.String()
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

func (s *CreateShareLinkRequest) SetPreviewLimit(v int64) *CreateShareLinkRequest {
	s.PreviewLimit = &v
	return s
}

func (s *CreateShareLinkRequest) SetSaveLimit(v int64) *CreateShareLinkRequest {
	s.SaveLimit = &v
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ShareLink         `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type CreateUserRequest struct {
	Avatar        *string                           `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Description   *string                           `json:"description,omitempty" xml:"description,omitempty"`
	Email         *string                           `json:"email,omitempty" xml:"email,omitempty"`
	GroupInfoList []*CreateUserRequestGroupInfoList `json:"group_info_list,omitempty" xml:"group_info_list,omitempty" type:"Repeated"`
	NickName      *string                           `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	Phone         *string                           `json:"phone,omitempty" xml:"phone,omitempty"`
	Role          *string                           `json:"role,omitempty" xml:"role,omitempty"`
	Status        *string                           `json:"status,omitempty" xml:"status,omitempty"`
	UserData      *string                           `json:"user_data,omitempty" xml:"user_data,omitempty"`
	UserId        *string                           `json:"user_id,omitempty" xml:"user_id,omitempty"`
	UserName      *string                           `json:"user_name,omitempty" xml:"user_name,omitempty"`
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

func (s *CreateUserRequest) SetUserData(v string) *CreateUserRequest {
	s.UserData = &v
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

func (s *CreateUserResponseBody) SetUserData(v map[string]*string) *CreateUserResponseBody {
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *CreateUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type DeleteDriveRequest struct {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId  *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId      *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeleteFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	DriveId    *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId     *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type DeleteUserRequest struct {
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	DriveId    *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
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
	Headers    map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                          `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *DeltaGetLastCursorResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId                *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId                 *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	ImageThumbnailProcess  *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	OfficeThumbnailProcess *string `json:"office_thumbnail_process,omitempty" xml:"office_thumbnail_process,omitempty"`
	VideoThumbnailProcess  *string `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
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

func (s *DownloadFileRequest) SetVideoThumbnailProcess(v string) *DownloadFileRequest {
	s.VideoThumbnailProcess = &v
	return s
}

type DownloadFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	DriveId    *string                 `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId     *string                 `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	DriveId *string   `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId  *string   `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId  *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       []*FilePermissionMember `json:"body,omitempty" xml:"body,omitempty" require:"true" type:"Repeated"`
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
	DriveId  *string                           `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId   *string                           `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Key   *string `json:"key,omitempty" xml:"key,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *FilePutUserTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId    *string                                  `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId     *string                                  `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Identity *Identity `json:"identity,omitempty" xml:"identity,omitempty"`
	RoleId   *string   `json:"role_id,omitempty" xml:"role_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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
	AsyncTaskId     *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	ConsumedProcess *int64  `json:"consumed_process,omitempty" xml:"consumed_process,omitempty"`
	ErrCode         *int64  `json:"err_code,omitempty" xml:"err_code,omitempty"`
	Message         *string `json:"message,omitempty" xml:"message,omitempty"`
	Status          *string `json:"status,omitempty" xml:"status,omitempty"`
	TotalProcess    *int64  `json:"total_process,omitempty" xml:"total_process,omitempty"`
	Url             *string `json:"url,omitempty" xml:"url,omitempty"`
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

func (s *GetAsyncTaskResponseBody) SetConsumedProcess(v int64) *GetAsyncTaskResponseBody {
	s.ConsumedProcess = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetErrCode(v int64) *GetAsyncTaskResponseBody {
	s.ErrCode = &v
	return s
}

func (s *GetAsyncTaskResponseBody) SetMessage(v string) *GetAsyncTaskResponseBody {
	s.Message = &v
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

func (s *GetAsyncTaskResponseBody) SetUrl(v string) *GetAsyncTaskResponseBody {
	s.Url = &v
	return s
}

type GetAsyncTaskResponse struct {
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetAsyncTaskResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Drive             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetDownloadUrlRequest struct {
	DriveId   *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	ExpireSec *int32  `json:"expire_sec,omitempty" xml:"expire_sec,omitempty"`
	FileId    *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	FileName  *string `json:"file_name,omitempty" xml:"file_name,omitempty"`
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

type GetDownloadUrlResponseBody struct {
	CdnUrl          *string `json:"cdn_url,omitempty" xml:"cdn_url,omitempty"`
	ContentHash     *string `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName *string `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	Crc64Hash       *string `json:"crc64_hash,omitempty" xml:"crc64_hash,omitempty"`
	Expiration      *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	InternalUrl     *string `json:"internal_url,omitempty" xml:"internal_url,omitempty"`
	Size            *int64  `json:"size,omitempty" xml:"size,omitempty"`
	Url             *string `json:"url,omitempty" xml:"url,omitempty"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Drive             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId      *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Fields       *string `json:"fields,omitempty" xml:"fields,omitempty"`
	FileId       *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	UrlExpireSec *int32  `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
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

func (s *GetFileRequest) SetUrlExpireSec(v int32) *GetFileRequest {
	s.UrlExpireSec = &v
	return s
}

type GetFileResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *File              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Group             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetLinkInfoRequest struct {
	Extra    *string `json:"extra,omitempty" xml:"extra,omitempty"`
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *AccountLinkInfo   `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetLinkInfoByUserIdResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId      *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Fields       *string `json:"fields,omitempty" xml:"fields,omitempty"`
	FileId       *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	RevisionId   *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
	UrlExpireSec *int64  `json:"url_expire_sec,omitempty" xml:"url_expire_sec,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Revision          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ShareLink         `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AccessCount       *int64  `json:"access_count,omitempty" xml:"access_count,omitempty"`
	Avatar            *string `json:"avatar,omitempty" xml:"avatar,omitempty"`
	CreatorId         *string `json:"creator_id,omitempty" xml:"creator_id,omitempty"`
	CreatorName       *string `json:"creator_name,omitempty" xml:"creator_name,omitempty"`
	CreatorPhone      *string `json:"creator_phone,omitempty" xml:"creator_phone,omitempty"`
	DisableDownload   *bool   `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	DisablePreview    *bool   `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	DisableSave       *bool   `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	DownloadCount     *int64  `json:"download_count,omitempty" xml:"download_count,omitempty"`
	DownloadLimit     *int64  `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	Expiration        *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	PreviewCount      *int64  `json:"preview_count,omitempty" xml:"preview_count,omitempty"`
	PreviewLimit      *int64  `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	ReportCount       *int64  `json:"report_count,omitempty" xml:"report_count,omitempty"`
	SaveCount         *int64  `json:"save_count,omitempty" xml:"save_count,omitempty"`
	SaveDownloadLimit *int64  `json:"save_download_limit,omitempty" xml:"save_download_limit,omitempty"`
	SaveLimit         *int64  `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	ShareName         *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	UpdatedAt         *string `json:"updated_at,omitempty" xml:"updated_at,omitempty"`
	VideoPreviewCount *int64  `json:"video_preview_count,omitempty" xml:"video_preview_count,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetShareLinkByAnonymousResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	ExpireSec *int32  `json:"expire_sec,omitempty" xml:"expire_sec,omitempty"`
	ShareId   *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	SharePwd  *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
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
	ExpiresIn  *int64  `json:"expires_in,omitempty" xml:"expires_in,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetShareLinkTokenResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type GetUploadUrlRequest struct {
	DriveId      *string                            `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId       *string                            `json:"file_id,omitempty" xml:"file_id,omitempty"`
	PartInfoList []*GetUploadUrlRequestPartInfoList `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	ShareId      *string                            `json:"share_id,omitempty" xml:"share_id,omitempty"`
	UploadId     *string                            `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
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
	ParallelSha1Ctx *GetUploadUrlRequestPartInfoListParallelSha1Ctx `json:"parallel_sha1_ctx,omitempty" xml:"parallel_sha1_ctx,omitempty" type:"Struct"`
	PartNumber      *int32                                          `json:"part_number,omitempty" xml:"part_number,omitempty"`
}

func (s GetUploadUrlRequestPartInfoList) String() string {
	return tea.Prettify(s)
}

func (s GetUploadUrlRequestPartInfoList) GoString() string {
	return s.String()
}

func (s *GetUploadUrlRequestPartInfoList) SetParallelSha1Ctx(v *GetUploadUrlRequestPartInfoListParallelSha1Ctx) *GetUploadUrlRequestPartInfoList {
	s.ParallelSha1Ctx = v
	return s
}

func (s *GetUploadUrlRequestPartInfoList) SetPartNumber(v int32) *GetUploadUrlRequestPartInfoList {
	s.PartNumber = &v
	return s
}

type GetUploadUrlRequestPartInfoListParallelSha1Ctx struct {
	H          []*int64 `json:"h,omitempty" xml:"h,omitempty" type:"Repeated"`
	PartOffset *int64   `json:"part_offset,omitempty" xml:"part_offset,omitempty"`
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

type GetUploadUrlResponseBody struct {
	CreateAt     *string           `json:"create_at,omitempty" xml:"create_at,omitempty"`
	DomainId     *string           `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId      *string           `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId       *string           `json:"file_id,omitempty" xml:"file_id,omitempty"`
	PartInfoList []*UploadPartInfo `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	UploadId     *string           `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetUploadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *User              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Category      *string `json:"category,omitempty" xml:"category,omitempty"`
	DriveId       *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId        *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	GetWithoutUrl *bool   `json:"get_without_url,omitempty" xml:"get_without_url,omitempty"`
	ShareId       *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	TemplateId    *string `json:"template_id,omitempty" xml:"template_id,omitempty"`
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

func (s *GetVideoPreviewPlayInfoRequest) SetGetWithoutUrl(v bool) *GetVideoPreviewPlayInfoRequest {
	s.GetWithoutUrl = &v
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

type GetVideoPreviewPlayInfoResponseBody struct {
	DomainId             *string               `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId              *string               `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId               *string               `json:"file_id,omitempty" xml:"file_id,omitempty"`
	ShareId              *string               `json:"share_id,omitempty" xml:"share_id,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVideoPreviewPlayInfoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Category *string `json:"category,omitempty" xml:"category,omitempty"`
	DriveId  *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId   *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	ShareId  *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
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
	DomainId             *string               `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId              *string               `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId               *string               `json:"file_id,omitempty" xml:"file_id,omitempty"`
	ShareId              *string               `json:"share_id,omitempty" xml:"share_id,omitempty"`
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
	Headers    map[string]*string                   `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                               `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *GetVideoPreviewPlayMetaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ImportUserRequest struct {
	AuthenticationDisplayName *string `json:"authentication_display_name,omitempty" xml:"authentication_display_name,omitempty"`
	AuthenticationType        *string `json:"authentication_type,omitempty" xml:"authentication_type,omitempty"`
	AutoCreateDrive           *bool   `json:"auto_create_drive,omitempty" xml:"auto_create_drive,omitempty"`
	DriveTotalSize            *int64  `json:"drive_total_size,omitempty" xml:"drive_total_size,omitempty"`
	Extra                     *string `json:"extra,omitempty" xml:"extra,omitempty"`
	Identity                  *string `json:"identity,omitempty" xml:"identity,omitempty"`
	NickName                  *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	ParentGroupId             *string `json:"parent_group_id,omitempty" xml:"parent_group_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *User              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type LinkAccountRequest struct {
	Extra    *string `json:"extra,omitempty" xml:"extra,omitempty"`
	Identity *string `json:"identity,omitempty" xml:"identity,omitempty"`
	Type     *string `json:"type,omitempty" xml:"type,omitempty"`
	UserId   *string `json:"user_id,omitempty" xml:"user_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Token             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId               *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	Limit                 *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker                *string `json:"marker,omitempty" xml:"marker,omitempty"`
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
	Items      []*AddressGroup `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string         `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListAddressGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListDeltaRequest struct {
	Cursor     *string `json:"cursor,omitempty" xml:"cursor,omitempty"`
	DriveId    *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Limit      *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
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
	Cursor  *string                       `json:"cursor,omitempty" xml:"cursor,omitempty"`
	HasMore *bool                         `json:"has_more,omitempty" xml:"has_more,omitempty"`
	Items   []*ListDeltaResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	File   *File   `json:"file,omitempty" xml:"file,omitempty"`
	FileId *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Op     *string `json:"op,omitempty" xml:"op,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDeltaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListDriveRequest struct {
	Limit     *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker    *string `json:"marker,omitempty" xml:"marker,omitempty"`
	Owner     *string `json:"owner,omitempty" xml:"owner,omitempty"`
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
	Items      []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string  `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListDriveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Limit   *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker  *string `json:"marker,omitempty" xml:"marker,omitempty"`
	Remarks *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
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

type ListFacegroupsResponseBody struct {
	Items      []*FaceGroup `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string      `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
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

type ListFacegroupsResponse struct {
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFacegroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Category       *string `json:"category,omitempty" xml:"category,omitempty"`
	DriveId        *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Fields         *string `json:"fields,omitempty" xml:"fields,omitempty"`
	Limit          *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker         *string `json:"marker,omitempty" xml:"marker,omitempty"`
	OrderBy        *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	OrderDirection *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
	ParentFileId   *string `json:"parent_file_id,omitempty" xml:"parent_file_id,omitempty"`
	Status         *string `json:"status,omitempty" xml:"status,omitempty"`
	Type           *string `json:"type,omitempty" xml:"type,omitempty"`
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

func (s *ListFileRequest) SetStatus(v string) *ListFileRequest {
	s.Status = &v
	return s
}

func (s *ListFileRequest) SetType(v string) *ListFileRequest {
	s.Type = &v
	return s
}

type ListFileResponseBody struct {
	Items      []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Limit  *string `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker *string `json:"marker,omitempty" xml:"marker,omitempty"`
}

func (s ListGroupRequest) String() string {
	return tea.Prettify(s)
}

func (s ListGroupRequest) GoString() string {
	return s.String()
}

func (s *ListGroupRequest) SetLimit(v string) *ListGroupRequest {
	s.Limit = &v
	return s
}

func (s *ListGroupRequest) SetMarker(v string) *ListGroupRequest {
	s.Marker = &v
	return s
}

type ListGroupResponseBody struct {
	Items      []*Group `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string  `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListGroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListMyDrivesRequest struct {
	Limit  *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
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
	Items      []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string  `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListMyDrivesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ListRecyclebinRequest struct {
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Fields  *string `json:"fields,omitempty" xml:"fields,omitempty"`
	Limit   *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker  *string `json:"marker,omitempty" xml:"marker,omitempty"`
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
	Items      []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	Headers    map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                      `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRecyclebinResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Fields  *string `json:"fields,omitempty" xml:"fields,omitempty"`
	FileId  *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Limit   *int64  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker  *string `json:"marker,omitempty" xml:"marker,omitempty"`
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
	Items      []*Revision `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string     `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
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
	Headers    map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                    `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListRevisionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Creator          *string `json:"creator,omitempty" xml:"creator,omitempty"`
	IncludeCancelled *bool   `json:"include_cancelled,omitempty" xml:"include_cancelled,omitempty"`
	Limit            *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker           *string `json:"marker,omitempty" xml:"marker,omitempty"`
	OrderBy          *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	OrderDirection   *string `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
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
	Items      []*ShareLink `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string      `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
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
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListShareLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId               *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	ImageThumbnailProcess *string `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId          *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId           *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Limit            *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	PartNumberMarker *int32  `json:"part_number_marker,omitempty" xml:"part_number_marker,omitempty"`
	ShareId          *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	UploadId         *string `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
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
	FileId               *string           `json:"file_id,omitempty" xml:"file_id,omitempty"`
	NextPartNumberMarker *string           `json:"next_part_number_marker,omitempty" xml:"next_part_number_marker,omitempty"`
	ParallelUpload       *bool             `json:"parallel_upload,omitempty" xml:"parallel_upload,omitempty"`
	UploadId             *string           `json:"upload_id,omitempty" xml:"upload_id,omitempty"`
	UploadedParts        []*UploadPartInfo `json:"uploaded_parts,omitempty" xml:"uploaded_parts,omitempty" type:"Repeated"`
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
	Headers    map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                         `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUploadedPartsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Limit  *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
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
	Items      []*User `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ListUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CheckNameMode  *string `json:"check_name_mode,omitempty" xml:"check_name_mode,omitempty"`
	DriveId        *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId         *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Exist       *bool   `json:"exist,omitempty" xml:"exist,omitempty"`
	FileId      *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *MoveFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type ParseKeywordsRequest struct {
	Keywords *string `json:"keywords,omitempty" xml:"keywords,omitempty"`
}

func (s ParseKeywordsRequest) String() string {
	return tea.Prettify(s)
}

func (s ParseKeywordsRequest) GoString() string {
	return s.String()
}

func (s *ParseKeywordsRequest) SetKeywords(v string) *ParseKeywordsRequest {
	s.Keywords = &v
	return s
}

type ParseKeywordsResponseBody struct {
	AddressLine *string                             `json:"address_line,omitempty" xml:"address_line,omitempty"`
	Tags        []*SystemTag                        `json:"tags,omitempty" xml:"tags,omitempty" type:"Repeated"`
	TimeRange   *ParseKeywordsResponseBodyTimeRange `json:"time_range,omitempty" xml:"time_range,omitempty" type:"Struct"`
}

func (s ParseKeywordsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ParseKeywordsResponseBody) GoString() string {
	return s.String()
}

func (s *ParseKeywordsResponseBody) SetAddressLine(v string) *ParseKeywordsResponseBody {
	s.AddressLine = &v
	return s
}

func (s *ParseKeywordsResponseBody) SetTags(v []*SystemTag) *ParseKeywordsResponseBody {
	s.Tags = v
	return s
}

func (s *ParseKeywordsResponseBody) SetTimeRange(v *ParseKeywordsResponseBodyTimeRange) *ParseKeywordsResponseBody {
	s.TimeRange = v
	return s
}

type ParseKeywordsResponseBodyTimeRange struct {
	End   *string `json:"end,omitempty" xml:"end,omitempty"`
	Start *string `json:"start,omitempty" xml:"start,omitempty"`
}

func (s ParseKeywordsResponseBodyTimeRange) String() string {
	return tea.Prettify(s)
}

func (s ParseKeywordsResponseBodyTimeRange) GoString() string {
	return s.String()
}

func (s *ParseKeywordsResponseBodyTimeRange) SetEnd(v string) *ParseKeywordsResponseBodyTimeRange {
	s.End = &v
	return s
}

func (s *ParseKeywordsResponseBodyTimeRange) SetStart(v string) *ParseKeywordsResponseBodyTimeRange {
	s.Start = &v
	return s
}

type ParseKeywordsResponse struct {
	Headers    map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                     `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ParseKeywordsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ParseKeywordsResponse) String() string {
	return tea.Prettify(s)
}

func (s ParseKeywordsResponse) GoString() string {
	return s.String()
}

func (s *ParseKeywordsResponse) SetHeaders(v map[string]*string) *ParseKeywordsResponse {
	s.Headers = v
	return s
}

func (s *ParseKeywordsResponse) SetStatusCode(v int32) *ParseKeywordsResponse {
	s.StatusCode = &v
	return s
}

func (s *ParseKeywordsResponse) SetBody(v *ParseKeywordsResponseBody) *ParseKeywordsResponse {
	s.Body = v
	return s
}

type RemoveFaceGroupFileRequest struct {
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FaceGroupId *string `json:"face_group_id,omitempty" xml:"face_group_id,omitempty"`
	FileId      *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
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

type RestoreFileRequest struct {
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId  *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId      *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *RestoreFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId    *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId     *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Revision          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Fields  *string `json:"fields,omitempty" xml:"fields,omitempty"`
	Limit   *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker  *string `json:"marker,omitempty" xml:"marker,omitempty"`
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
	Items      []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	Headers    map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ScanFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	AddressLevel          *string   `json:"address_level,omitempty" xml:"address_level,omitempty"`
	AddressNames          []*string `json:"address_names,omitempty" xml:"address_names,omitempty" type:"Repeated"`
	BrGeoPoint            *string   `json:"br_geo_point,omitempty" xml:"br_geo_point,omitempty"`
	DriveId               *string   `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	ImageThumbnailProcess *string   `json:"image_thumbnail_process,omitempty" xml:"image_thumbnail_process,omitempty"`
	TlGeoPoint            *string   `json:"tl_geo_point,omitempty" xml:"tl_geo_point,omitempty"`
	VideoThumbnailProcess *string   `json:"video_thumbnail_process,omitempty" xml:"video_thumbnail_process,omitempty"`
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
	Headers    map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                           `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchAddressGroupsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SearchDriveRequest struct {
	DriveName *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	Limit     *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker    *string `json:"marker,omitempty" xml:"marker,omitempty"`
	Owner     *string `json:"owner,omitempty" xml:"owner,omitempty"`
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
	Items      []*Drive `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string  `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
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
	Headers    map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                   `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchDriveResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId          *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	Limit            *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker           *string `json:"marker,omitempty" xml:"marker,omitempty"`
	OrderBy          *string `json:"order_by,omitempty" xml:"order_by,omitempty"`
	Query            *string `json:"query,omitempty" xml:"query,omitempty"`
	ReturnTotalCount *bool   `json:"return_total_count,omitempty" xml:"return_total_count,omitempty"`
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

func (s *SearchFileRequest) SetReturnTotalCount(v bool) *SearchFileRequest {
	s.ReturnTotalCount = &v
	return s
}

type SearchFileResponseBody struct {
	Items      []*File `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	TotalCount *int64  `json:"total_count,omitempty" xml:"total_count,omitempty"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Creators         []*string `json:"creators,omitempty" xml:"creators,omitempty" type:"Repeated"`
	Limit            *int32    `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker           *string   `json:"marker,omitempty" xml:"marker,omitempty"`
	OrderBy          *string   `json:"order_by,omitempty" xml:"order_by,omitempty"`
	OrderDirection   *string   `json:"order_direction,omitempty" xml:"order_direction,omitempty"`
	Query            *string   `json:"query,omitempty" xml:"query,omitempty"`
	ReturnTotalCount *bool     `json:"return_total_count,omitempty" xml:"return_total_count,omitempty"`
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
	Items      []*ShareLink `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	NextMarker *string      `json:"next_marker,omitempty" xml:"next_marker,omitempty"`
	TotalCount *int64       `json:"total_count,omitempty" xml:"total_count,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchShareLinkResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type SearchUserRequest struct {
	Email            *string `json:"email,omitempty" xml:"email,omitempty"`
	Limit            *int32  `json:"limit,omitempty" xml:"limit,omitempty"`
	Marker           *string `json:"marker,omitempty" xml:"marker,omitempty"`
	NickName         *string `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	NickNameForFuzzy *string `json:"nick_name_for_fuzzy,omitempty" xml:"nick_name_for_fuzzy,omitempty"`
	Phone            *string `json:"phone,omitempty" xml:"phone,omitempty"`
	Role             *string `json:"role,omitempty" xml:"role,omitempty"`
	Status           *string `json:"status,omitempty" xml:"status,omitempty"`
	UserName         *string `json:"user_name,omitempty" xml:"user_name,omitempty"`
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
	Items      []*User `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
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
	Headers    map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                  `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *SearchUserResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Assertion    *string `json:"assertion,omitempty" xml:"assertion,omitempty"`
	ClientId     *string `json:"client_id,omitempty" xml:"client_id,omitempty"`
	ClientSecret *string `json:"client_secret,omitempty" xml:"client_secret,omitempty"`
	Code         *string `json:"code,omitempty" xml:"code,omitempty"`
	GrantType    *string `json:"grant_type,omitempty" xml:"grant_type,omitempty"`
	RedirectUri  *string `json:"redirect_uri,omitempty" xml:"redirect_uri,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Token             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId  *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	AsyncTaskId *string `json:"async_task_id,omitempty" xml:"async_task_id,omitempty"`
	DomainId    *string `json:"domain_id,omitempty" xml:"domain_id,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId      *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
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
	Headers    map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                 `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *TrashFileResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateDriveRequest struct {
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	DriveId     *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	DriveName   *string `json:"drive_name,omitempty" xml:"drive_name,omitempty"`
	Status      *string `json:"status,omitempty" xml:"status,omitempty"`
	TotalSize   *int64  `json:"total_size,omitempty" xml:"total_size,omitempty"`
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

func (s *UpdateDriveRequest) SetStatus(v string) *UpdateDriveRequest {
	s.Status = &v
	return s
}

func (s *UpdateDriveRequest) SetTotalSize(v int64) *UpdateDriveRequest {
	s.TotalSize = &v
	return s
}

type UpdateDriveResponse struct {
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Drive             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	DriveId          *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	GroupCoverFaceId *string `json:"group_cover_face_id,omitempty" xml:"group_cover_face_id,omitempty"`
	GroupId          *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName        *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
	Remarks          *string `json:"remarks,omitempty" xml:"remarks,omitempty"`
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
	DriveId *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
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
	Headers    map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32                       `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *UpdateFacegroupResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	CheckNameMode   *string   `json:"check_name_mode,omitempty" xml:"check_name_mode,omitempty"`
	Description     *string   `json:"description,omitempty" xml:"description,omitempty"`
	DriveId         *string   `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId          *string   `json:"file_id,omitempty" xml:"file_id,omitempty"`
	Hidden          *bool     `json:"hidden,omitempty" xml:"hidden,omitempty"`
	Labels          []*string `json:"labels,omitempty" xml:"labels,omitempty" type:"Repeated"`
	LocalModifiedAt *string   `json:"local_modified_at,omitempty" xml:"local_modified_at,omitempty"`
	Name            *string   `json:"name,omitempty" xml:"name,omitempty"`
	Starred         *bool     `json:"starred,omitempty" xml:"starred,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *File              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
	GroupId     *string `json:"group_id,omitempty" xml:"group_id,omitempty"`
	GroupName   *string `json:"group_name,omitempty" xml:"group_name,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Group             `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateRevisionRequest struct {
	DriveId             *string `json:"drive_id,omitempty" xml:"drive_id,omitempty"`
	FileId              *string `json:"file_id,omitempty" xml:"file_id,omitempty"`
	KeepForever         *bool   `json:"keep_forever,omitempty" xml:"keep_forever,omitempty"`
	RevisionDescription *string `json:"revision_description,omitempty" xml:"revision_description,omitempty"`
	RevisionId          *string `json:"revision_id,omitempty" xml:"revision_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *Revision          `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	Description       *string `json:"description,omitempty" xml:"description,omitempty"`
	DisableDownload   *bool   `json:"disable_download,omitempty" xml:"disable_download,omitempty"`
	DisablePreview    *bool   `json:"disable_preview,omitempty" xml:"disable_preview,omitempty"`
	DisableSave       *bool   `json:"disable_save,omitempty" xml:"disable_save,omitempty"`
	DownloadCount     *int64  `json:"download_count,omitempty" xml:"download_count,omitempty"`
	DownloadLimit     *int64  `json:"download_limit,omitempty" xml:"download_limit,omitempty"`
	Expiration        *string `json:"expiration,omitempty" xml:"expiration,omitempty"`
	PreviewCount      *int64  `json:"preview_count,omitempty" xml:"preview_count,omitempty"`
	PreviewLimit      *int64  `json:"preview_limit,omitempty" xml:"preview_limit,omitempty"`
	ReportCount       *int64  `json:"report_count,omitempty" xml:"report_count,omitempty"`
	SaveCount         *int64  `json:"save_count,omitempty" xml:"save_count,omitempty"`
	SaveLimit         *int64  `json:"save_limit,omitempty" xml:"save_limit,omitempty"`
	ShareId           *string `json:"share_id,omitempty" xml:"share_id,omitempty"`
	ShareName         *string `json:"share_name,omitempty" xml:"share_name,omitempty"`
	SharePwd          *string `json:"share_pwd,omitempty" xml:"share_pwd,omitempty"`
	Status            *string `json:"status,omitempty" xml:"status,omitempty"`
	VideoPreviewCount *int64  `json:"video_preview_count,omitempty" xml:"video_preview_count,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *ShareLink         `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

type UpdateUserRequest struct {
	Avatar        *string                           `json:"avatar,omitempty" xml:"avatar,omitempty"`
	Description   *string                           `json:"description,omitempty" xml:"description,omitempty"`
	Email         *string                           `json:"email,omitempty" xml:"email,omitempty"`
	GroupInfoList []*UpdateUserRequestGroupInfoList `json:"group_info_list,omitempty" xml:"group_info_list,omitempty" type:"Repeated"`
	NickName      *string                           `json:"nick_name,omitempty" xml:"nick_name,omitempty"`
	Phone         *string                           `json:"phone,omitempty" xml:"phone,omitempty"`
	Role          *string                           `json:"role,omitempty" xml:"role,omitempty"`
	Status        *string                           `json:"status,omitempty" xml:"status,omitempty"`
	UserData      map[string]*string                `json:"user_data,omitempty" xml:"user_data,omitempty"`
	UserId        *string                           `json:"user_id,omitempty" xml:"user_id,omitempty"`
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
	Headers    map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	StatusCode *int32             `json:"statusCode,omitempty" xml:"statusCode,omitempty" require:"true"`
	Body       *User              `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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
	interfaceSPI, _err := gatewayclient.NewClient()
	if _err != nil {
		return _err
	}

	client.Spi = interfaceSPI
	client.SignatureAlgorithm = tea.String("v2")
	client.EndpointRule = tea.String("")
	return nil
}

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
		BodyType:    tea.String("json"),
	}
	_result = &AuthorizeResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.ImageMediaMetadata))) {
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

	if !tea.BoolValue(util.IsUnset(tea.ToMap(request.VideoMediaMetadata))) {
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

func (client *Client) CreateShareLinkWithOptions(request *CreateShareLinkRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *CreateShareLinkResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.PreviewLimit)) {
		body["preview_limit"] = request.PreviewLimit
	}

	if !tea.BoolValue(util.IsUnset(request.SaveLimit)) {
		body["save_limit"] = request.SaveLimit
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

func (client *Client) DownloadFileWithOptions(request *DownloadFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *DownloadFileResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.ImageThumbnailProcess)) {
		body["image_thumbnail_process"] = request.ImageThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.OfficeThumbnailProcess)) {
		body["office_thumbnail_process"] = request.OfficeThumbnailProcess
	}

	if !tea.BoolValue(util.IsUnset(request.VideoThumbnailProcess)) {
		body["video_thumbnail_process"] = request.VideoThumbnailProcess
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
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
		BodyType:    tea.String("json"),
	}
	_result = &DownloadFileResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

	if !tea.BoolValue(util.IsUnset(request.GetWithoutUrl)) {
		body["get_without_url"] = request.GetWithoutUrl
	}

	if !tea.BoolValue(util.IsUnset(request.ShareId)) {
		body["share_id"] = request.ShareId
	}

	if !tea.BoolValue(util.IsUnset(request.TemplateId)) {
		body["template_id"] = request.TemplateId
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

	if !tea.BoolValue(util.IsUnset(request.Status)) {
		body["status"] = request.Status
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

func (client *Client) ParseKeywords(request *ParseKeywordsRequest) (_result *ParseKeywordsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	headers := make(map[string]*string)
	_result = &ParseKeywordsResponse{}
	_body, _err := client.ParseKeywordsWithOptions(request, headers, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ParseKeywordsWithOptions(request *ParseKeywordsRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *ParseKeywordsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	body := map[string]interface{}{}
	if !tea.BoolValue(util.IsUnset(request.Keywords)) {
		body["keywords"] = request.Keywords
	}

	req := &openapi.OpenApiRequest{
		Headers: headers,
		Body:    openapiutil.ParseToMap(body),
	}
	params := &openapi.Params{
		Action:      tea.String("ParseKeywords"),
		Version:     tea.String("2022-03-01"),
		Protocol:    tea.String("HTTPS"),
		Pathname:    tea.String("/v2/image/parse_keywords"),
		Method:      tea.String("POST"),
		AuthType:    tea.String("AK"),
		Style:       tea.String("ROA"),
		ReqBodyType: tea.String("json"),
		BodyType:    tea.String("json"),
	}
	_result = &ParseKeywordsResponse{}
	_body, _err := client.Execute(params, req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

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

func (client *Client) SearchFileWithOptions(request *SearchFileRequest, headers map[string]*string, runtime *util.RuntimeOptions) (_result *SearchFileResponse, _err error) {
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

	if !tea.BoolValue(util.IsUnset(request.OrderBy)) {
		body["order_by"] = request.OrderBy
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
