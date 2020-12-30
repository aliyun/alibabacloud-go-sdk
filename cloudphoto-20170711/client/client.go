// This file is auto-generated, don't edit it. Thanks.
/**
 *
 */
package client

import (
	openapi "github.com/alibabacloud-go/darabonba-openapi/client"
	endpointutil "github.com/alibabacloud-go/endpoint-util/service"
	util "github.com/alibabacloud-go/tea-utils/service"
	"github.com/alibabacloud-go/tea/tea"
)

type ActivatePhotosRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s ActivatePhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s ActivatePhotosRequest) GoString() string {
	return s.String()
}

func (s *ActivatePhotosRequest) SetStoreName(v string) *ActivatePhotosRequest {
	s.StoreName = &v
	return s
}

func (s *ActivatePhotosRequest) SetLibraryId(v string) *ActivatePhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *ActivatePhotosRequest) SetPhotoId(v []*int) *ActivatePhotosRequest {
	s.PhotoId = v
	return s
}

type ActivatePhotosResponseBody struct {
	Action    *string                              `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*ActivatePhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ActivatePhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ActivatePhotosResponseBody) GoString() string {
	return s.String()
}

func (s *ActivatePhotosResponseBody) SetAction(v string) *ActivatePhotosResponseBody {
	s.Action = &v
	return s
}

func (s *ActivatePhotosResponseBody) SetMessage(v string) *ActivatePhotosResponseBody {
	s.Message = &v
	return s
}

func (s *ActivatePhotosResponseBody) SetRequestId(v string) *ActivatePhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ActivatePhotosResponseBody) SetResults(v []*ActivatePhotosResponseBodyResults) *ActivatePhotosResponseBody {
	s.Results = v
	return s
}

func (s *ActivatePhotosResponseBody) SetCode(v string) *ActivatePhotosResponseBody {
	s.Code = &v
	return s
}

type ActivatePhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ActivatePhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ActivatePhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ActivatePhotosResponseBodyResults) SetIdStr(v string) *ActivatePhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *ActivatePhotosResponseBodyResults) SetCode(v string) *ActivatePhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ActivatePhotosResponseBodyResults) SetMessage(v string) *ActivatePhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ActivatePhotosResponseBodyResults) SetId(v int64) *ActivatePhotosResponseBodyResults {
	s.Id = &v
	return s
}

type ActivatePhotosResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ActivatePhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ActivatePhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s ActivatePhotosResponse) GoString() string {
	return s.String()
}

func (s *ActivatePhotosResponse) SetHeaders(v map[string]*string) *ActivatePhotosResponse {
	s.Headers = v
	return s
}

func (s *ActivatePhotosResponse) SetBody(v *ActivatePhotosResponseBody) *ActivatePhotosResponse {
	s.Body = v
	return s
}

type AddAlbumPhotosRequest struct {
	AlbumId   *int64  `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s AddAlbumPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s AddAlbumPhotosRequest) GoString() string {
	return s.String()
}

func (s *AddAlbumPhotosRequest) SetAlbumId(v int64) *AddAlbumPhotosRequest {
	s.AlbumId = &v
	return s
}

func (s *AddAlbumPhotosRequest) SetStoreName(v string) *AddAlbumPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *AddAlbumPhotosRequest) SetLibraryId(v string) *AddAlbumPhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *AddAlbumPhotosRequest) SetPhotoId(v []*int) *AddAlbumPhotosRequest {
	s.PhotoId = v
	return s
}

type AddAlbumPhotosResponseBody struct {
	Action    *string                              `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*AddAlbumPhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s AddAlbumPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s AddAlbumPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *AddAlbumPhotosResponseBody) SetAction(v string) *AddAlbumPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *AddAlbumPhotosResponseBody) SetMessage(v string) *AddAlbumPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *AddAlbumPhotosResponseBody) SetRequestId(v string) *AddAlbumPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAlbumPhotosResponseBody) SetResults(v []*AddAlbumPhotosResponseBodyResults) *AddAlbumPhotosResponseBody {
	s.Results = v
	return s
}

func (s *AddAlbumPhotosResponseBody) SetCode(v string) *AddAlbumPhotosResponseBody {
	s.Code = &v
	return s
}

type AddAlbumPhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s AddAlbumPhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s AddAlbumPhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *AddAlbumPhotosResponseBodyResults) SetIdStr(v string) *AddAlbumPhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *AddAlbumPhotosResponseBodyResults) SetCode(v string) *AddAlbumPhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *AddAlbumPhotosResponseBodyResults) SetMessage(v string) *AddAlbumPhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *AddAlbumPhotosResponseBodyResults) SetId(v int64) *AddAlbumPhotosResponseBodyResults {
	s.Id = &v
	return s
}

type AddAlbumPhotosResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *AddAlbumPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s AddAlbumPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s AddAlbumPhotosResponse) GoString() string {
	return s.String()
}

func (s *AddAlbumPhotosResponse) SetHeaders(v map[string]*string) *AddAlbumPhotosResponse {
	s.Headers = v
	return s
}

func (s *AddAlbumPhotosResponse) SetBody(v *AddAlbumPhotosResponseBody) *AddAlbumPhotosResponse {
	s.Body = v
	return s
}

type CreateAlbumRequest struct {
	AlbumName *string `json:"AlbumName,omitempty" xml:"AlbumName,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	Remark    *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s CreateAlbumRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateAlbumRequest) GoString() string {
	return s.String()
}

func (s *CreateAlbumRequest) SetAlbumName(v string) *CreateAlbumRequest {
	s.AlbumName = &v
	return s
}

func (s *CreateAlbumRequest) SetStoreName(v string) *CreateAlbumRequest {
	s.StoreName = &v
	return s
}

func (s *CreateAlbumRequest) SetRemark(v string) *CreateAlbumRequest {
	s.Remark = &v
	return s
}

func (s *CreateAlbumRequest) SetLibraryId(v string) *CreateAlbumRequest {
	s.LibraryId = &v
	return s
}

type CreateAlbumResponseBody struct {
	Action    *string                       `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Album     *CreateAlbumResponseBodyAlbum `json:"Album,omitempty" xml:"Album,omitempty" type:"Struct"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateAlbumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateAlbumResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlbumResponseBody) SetAction(v string) *CreateAlbumResponseBody {
	s.Action = &v
	return s
}

func (s *CreateAlbumResponseBody) SetMessage(v string) *CreateAlbumResponseBody {
	s.Message = &v
	return s
}

func (s *CreateAlbumResponseBody) SetRequestId(v string) *CreateAlbumResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlbumResponseBody) SetAlbum(v *CreateAlbumResponseBodyAlbum) *CreateAlbumResponseBody {
	s.Album = v
	return s
}

func (s *CreateAlbumResponseBody) SetCode(v string) *CreateAlbumResponseBody {
	s.Code = &v
	return s
}

type CreateAlbumResponseBodyAlbum struct {
	IdStr       *string                            `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	PhotosCount *int64                             `json:"PhotosCount,omitempty" xml:"PhotosCount,omitempty"`
	Cover       *CreateAlbumResponseBodyAlbumCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Mtime       *int64                             `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime       *int64                             `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Remark      *string                            `json:"Remark,omitempty" xml:"Remark,omitempty"`
	State       *string                            `json:"State,omitempty" xml:"State,omitempty"`
	Name        *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *int64                             `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateAlbumResponseBodyAlbum) String() string {
	return tea.Prettify(s)
}

func (s CreateAlbumResponseBodyAlbum) GoString() string {
	return s.String()
}

func (s *CreateAlbumResponseBodyAlbum) SetIdStr(v string) *CreateAlbumResponseBodyAlbum {
	s.IdStr = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbum) SetPhotosCount(v int64) *CreateAlbumResponseBodyAlbum {
	s.PhotosCount = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbum) SetCover(v *CreateAlbumResponseBodyAlbumCover) *CreateAlbumResponseBodyAlbum {
	s.Cover = v
	return s
}

func (s *CreateAlbumResponseBodyAlbum) SetMtime(v int64) *CreateAlbumResponseBodyAlbum {
	s.Mtime = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbum) SetCtime(v int64) *CreateAlbumResponseBodyAlbum {
	s.Ctime = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbum) SetRemark(v string) *CreateAlbumResponseBodyAlbum {
	s.Remark = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbum) SetState(v string) *CreateAlbumResponseBodyAlbum {
	s.State = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbum) SetName(v string) *CreateAlbumResponseBodyAlbum {
	s.Name = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbum) SetId(v int64) *CreateAlbumResponseBodyAlbum {
	s.Id = &v
	return s
}

type CreateAlbumResponseBodyAlbumCover struct {
	Remark  *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	State   *string `json:"State,omitempty" xml:"State,omitempty"`
	Height  *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	FileId  *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Mtime   *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime   *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Width   *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Md5     *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s CreateAlbumResponseBodyAlbumCover) String() string {
	return tea.Prettify(s)
}

func (s CreateAlbumResponseBodyAlbumCover) GoString() string {
	return s.String()
}

func (s *CreateAlbumResponseBodyAlbumCover) SetRemark(v string) *CreateAlbumResponseBodyAlbumCover {
	s.Remark = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetState(v string) *CreateAlbumResponseBodyAlbumCover {
	s.State = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetHeight(v int64) *CreateAlbumResponseBodyAlbumCover {
	s.Height = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetFileId(v string) *CreateAlbumResponseBodyAlbumCover {
	s.FileId = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetIdStr(v string) *CreateAlbumResponseBodyAlbumCover {
	s.IdStr = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetMtime(v int64) *CreateAlbumResponseBodyAlbumCover {
	s.Mtime = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetCtime(v int64) *CreateAlbumResponseBodyAlbumCover {
	s.Ctime = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetWidth(v int64) *CreateAlbumResponseBodyAlbumCover {
	s.Width = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetMd5(v string) *CreateAlbumResponseBodyAlbumCover {
	s.Md5 = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetIsVideo(v bool) *CreateAlbumResponseBodyAlbumCover {
	s.IsVideo = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetTitle(v string) *CreateAlbumResponseBodyAlbumCover {
	s.Title = &v
	return s
}

func (s *CreateAlbumResponseBodyAlbumCover) SetId(v int64) *CreateAlbumResponseBodyAlbumCover {
	s.Id = &v
	return s
}

type CreateAlbumResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateAlbumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateAlbumResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateAlbumResponse) GoString() string {
	return s.String()
}

func (s *CreateAlbumResponse) SetHeaders(v map[string]*string) *CreateAlbumResponse {
	s.Headers = v
	return s
}

func (s *CreateAlbumResponse) SetBody(v *CreateAlbumResponseBody) *CreateAlbumResponse {
	s.Body = v
	return s
}

type CreatePhotoRequest struct {
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	SessionId       *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	UploadType      *string `json:"UploadType,omitempty" xml:"UploadType,omitempty"`
	PhotoTitle      *string `json:"PhotoTitle,omitempty" xml:"PhotoTitle,omitempty"`
	StoreName       *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	LibraryId       *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	Staging         *string `json:"Staging,omitempty" xml:"Staging,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
}

func (s CreatePhotoRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePhotoRequest) GoString() string {
	return s.String()
}

func (s *CreatePhotoRequest) SetFileId(v string) *CreatePhotoRequest {
	s.FileId = &v
	return s
}

func (s *CreatePhotoRequest) SetSessionId(v string) *CreatePhotoRequest {
	s.SessionId = &v
	return s
}

func (s *CreatePhotoRequest) SetUploadType(v string) *CreatePhotoRequest {
	s.UploadType = &v
	return s
}

func (s *CreatePhotoRequest) SetPhotoTitle(v string) *CreatePhotoRequest {
	s.PhotoTitle = &v
	return s
}

func (s *CreatePhotoRequest) SetStoreName(v string) *CreatePhotoRequest {
	s.StoreName = &v
	return s
}

func (s *CreatePhotoRequest) SetRemark(v string) *CreatePhotoRequest {
	s.Remark = &v
	return s
}

func (s *CreatePhotoRequest) SetLibraryId(v string) *CreatePhotoRequest {
	s.LibraryId = &v
	return s
}

func (s *CreatePhotoRequest) SetStaging(v string) *CreatePhotoRequest {
	s.Staging = &v
	return s
}

func (s *CreatePhotoRequest) SetShareExpireTime(v int64) *CreatePhotoRequest {
	s.ShareExpireTime = &v
	return s
}

func (s *CreatePhotoRequest) SetTakenAt(v int64) *CreatePhotoRequest {
	s.TakenAt = &v
	return s
}

type CreatePhotoResponseBody struct {
	Action    *string                       `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Photo     *CreatePhotoResponseBodyPhoto `json:"Photo,omitempty" xml:"Photo,omitempty" type:"Struct"`
	Code      *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreatePhotoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePhotoResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhotoResponseBody) SetAction(v string) *CreatePhotoResponseBody {
	s.Action = &v
	return s
}

func (s *CreatePhotoResponseBody) SetMessage(v string) *CreatePhotoResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePhotoResponseBody) SetRequestId(v string) *CreatePhotoResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhotoResponseBody) SetPhoto(v *CreatePhotoResponseBodyPhoto) *CreatePhotoResponseBody {
	s.Photo = v
	return s
}

func (s *CreatePhotoResponseBody) SetCode(v string) *CreatePhotoResponseBody {
	s.Code = &v
	return s
}

type CreatePhotoResponseBodyPhoto struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s CreatePhotoResponseBodyPhoto) String() string {
	return tea.Prettify(s)
}

func (s CreatePhotoResponseBodyPhoto) GoString() string {
	return s.String()
}

func (s *CreatePhotoResponseBodyPhoto) SetRemark(v string) *CreatePhotoResponseBodyPhoto {
	s.Remark = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetTakenAt(v int64) *CreatePhotoResponseBodyPhoto {
	s.TakenAt = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetState(v string) *CreatePhotoResponseBodyPhoto {
	s.State = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetHeight(v int64) *CreatePhotoResponseBodyPhoto {
	s.Height = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetShareExpireTime(v int64) *CreatePhotoResponseBodyPhoto {
	s.ShareExpireTime = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetFileId(v string) *CreatePhotoResponseBodyPhoto {
	s.FileId = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetIdStr(v string) *CreatePhotoResponseBodyPhoto {
	s.IdStr = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetCtime(v int64) *CreatePhotoResponseBodyPhoto {
	s.Ctime = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetMtime(v int64) *CreatePhotoResponseBodyPhoto {
	s.Mtime = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetWidth(v int64) *CreatePhotoResponseBodyPhoto {
	s.Width = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetSize(v int64) *CreatePhotoResponseBodyPhoto {
	s.Size = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetMd5(v string) *CreatePhotoResponseBodyPhoto {
	s.Md5 = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetTitle(v string) *CreatePhotoResponseBodyPhoto {
	s.Title = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetIsVideo(v bool) *CreatePhotoResponseBodyPhoto {
	s.IsVideo = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetId(v int64) *CreatePhotoResponseBodyPhoto {
	s.Id = &v
	return s
}

func (s *CreatePhotoResponseBodyPhoto) SetLocation(v string) *CreatePhotoResponseBodyPhoto {
	s.Location = &v
	return s
}

type CreatePhotoResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePhotoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePhotoResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePhotoResponse) GoString() string {
	return s.String()
}

func (s *CreatePhotoResponse) SetHeaders(v map[string]*string) *CreatePhotoResponse {
	s.Headers = v
	return s
}

func (s *CreatePhotoResponse) SetBody(v *CreatePhotoResponseBody) *CreatePhotoResponse {
	s.Body = v
	return s
}

type CreatePhotoStoreRequest struct {
	StoreName    *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	DefaultQuota *int64  `json:"DefaultQuota,omitempty" xml:"DefaultQuota,omitempty"`
	BucketName   *string `json:"BucketName,omitempty" xml:"BucketName,omitempty"`
	Remark       *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s CreatePhotoStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s CreatePhotoStoreRequest) GoString() string {
	return s.String()
}

func (s *CreatePhotoStoreRequest) SetStoreName(v string) *CreatePhotoStoreRequest {
	s.StoreName = &v
	return s
}

func (s *CreatePhotoStoreRequest) SetDefaultQuota(v int64) *CreatePhotoStoreRequest {
	s.DefaultQuota = &v
	return s
}

func (s *CreatePhotoStoreRequest) SetBucketName(v string) *CreatePhotoStoreRequest {
	s.BucketName = &v
	return s
}

func (s *CreatePhotoStoreRequest) SetRemark(v string) *CreatePhotoStoreRequest {
	s.Remark = &v
	return s
}

type CreatePhotoStoreResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreatePhotoStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreatePhotoStoreResponseBody) GoString() string {
	return s.String()
}

func (s *CreatePhotoStoreResponseBody) SetAction(v string) *CreatePhotoStoreResponseBody {
	s.Action = &v
	return s
}

func (s *CreatePhotoStoreResponseBody) SetMessage(v string) *CreatePhotoStoreResponseBody {
	s.Message = &v
	return s
}

func (s *CreatePhotoStoreResponseBody) SetRequestId(v string) *CreatePhotoStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreatePhotoStoreResponseBody) SetCode(v string) *CreatePhotoStoreResponseBody {
	s.Code = &v
	return s
}

type CreatePhotoStoreResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreatePhotoStoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreatePhotoStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s CreatePhotoStoreResponse) GoString() string {
	return s.String()
}

func (s *CreatePhotoStoreResponse) SetHeaders(v map[string]*string) *CreatePhotoStoreResponse {
	s.Headers = v
	return s
}

func (s *CreatePhotoStoreResponse) SetBody(v *CreatePhotoStoreResponseBody) *CreatePhotoStoreResponse {
	s.Body = v
	return s
}

type CreateTransactionRequest struct {
	Size      *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Ext       *string `json:"Ext,omitempty" xml:"Ext,omitempty"`
	Force     *string `json:"Force,omitempty" xml:"Force,omitempty"`
	Md5       *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s CreateTransactionRequest) String() string {
	return tea.Prettify(s)
}

func (s CreateTransactionRequest) GoString() string {
	return s.String()
}

func (s *CreateTransactionRequest) SetSize(v int64) *CreateTransactionRequest {
	s.Size = &v
	return s
}

func (s *CreateTransactionRequest) SetExt(v string) *CreateTransactionRequest {
	s.Ext = &v
	return s
}

func (s *CreateTransactionRequest) SetForce(v string) *CreateTransactionRequest {
	s.Force = &v
	return s
}

func (s *CreateTransactionRequest) SetMd5(v string) *CreateTransactionRequest {
	s.Md5 = &v
	return s
}

func (s *CreateTransactionRequest) SetStoreName(v string) *CreateTransactionRequest {
	s.StoreName = &v
	return s
}

func (s *CreateTransactionRequest) SetLibraryId(v string) *CreateTransactionRequest {
	s.LibraryId = &v
	return s
}

type CreateTransactionResponseBody struct {
	Action      *string                                   `json:"Action,omitempty" xml:"Action,omitempty"`
	Message     *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Transaction *CreateTransactionResponseBodyTransaction `json:"Transaction,omitempty" xml:"Transaction,omitempty" type:"Struct"`
	Code        *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s CreateTransactionResponseBody) String() string {
	return tea.Prettify(s)
}

func (s CreateTransactionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateTransactionResponseBody) SetAction(v string) *CreateTransactionResponseBody {
	s.Action = &v
	return s
}

func (s *CreateTransactionResponseBody) SetMessage(v string) *CreateTransactionResponseBody {
	s.Message = &v
	return s
}

func (s *CreateTransactionResponseBody) SetRequestId(v string) *CreateTransactionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateTransactionResponseBody) SetTransaction(v *CreateTransactionResponseBodyTransaction) *CreateTransactionResponseBody {
	s.Transaction = v
	return s
}

func (s *CreateTransactionResponseBody) SetCode(v string) *CreateTransactionResponseBody {
	s.Code = &v
	return s
}

type CreateTransactionResponseBodyTransaction struct {
	Upload *CreateTransactionResponseBodyTransactionUpload `json:"Upload,omitempty" xml:"Upload,omitempty" type:"Struct"`
}

func (s CreateTransactionResponseBodyTransaction) String() string {
	return tea.Prettify(s)
}

func (s CreateTransactionResponseBodyTransaction) GoString() string {
	return s.String()
}

func (s *CreateTransactionResponseBodyTransaction) SetUpload(v *CreateTransactionResponseBodyTransactionUpload) *CreateTransactionResponseBodyTransaction {
	s.Upload = v
	return s
}

type CreateTransactionResponseBodyTransactionUpload struct {
	ObjectKey       *string `json:"ObjectKey,omitempty" xml:"ObjectKey,omitempty"`
	AccessKeySecret *string `json:"AccessKeySecret,omitempty" xml:"AccessKeySecret,omitempty"`
	SessionId       *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	AccessKeyId     *string `json:"AccessKeyId,omitempty" xml:"AccessKeyId,omitempty"`
	StsToken        *string `json:"StsToken,omitempty" xml:"StsToken,omitempty"`
	OssEndpoint     *string `json:"OssEndpoint,omitempty" xml:"OssEndpoint,omitempty"`
	Bucket          *string `json:"Bucket,omitempty" xml:"Bucket,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
}

func (s CreateTransactionResponseBodyTransactionUpload) String() string {
	return tea.Prettify(s)
}

func (s CreateTransactionResponseBodyTransactionUpload) GoString() string {
	return s.String()
}

func (s *CreateTransactionResponseBodyTransactionUpload) SetObjectKey(v string) *CreateTransactionResponseBodyTransactionUpload {
	s.ObjectKey = &v
	return s
}

func (s *CreateTransactionResponseBodyTransactionUpload) SetAccessKeySecret(v string) *CreateTransactionResponseBodyTransactionUpload {
	s.AccessKeySecret = &v
	return s
}

func (s *CreateTransactionResponseBodyTransactionUpload) SetSessionId(v string) *CreateTransactionResponseBodyTransactionUpload {
	s.SessionId = &v
	return s
}

func (s *CreateTransactionResponseBodyTransactionUpload) SetAccessKeyId(v string) *CreateTransactionResponseBodyTransactionUpload {
	s.AccessKeyId = &v
	return s
}

func (s *CreateTransactionResponseBodyTransactionUpload) SetStsToken(v string) *CreateTransactionResponseBodyTransactionUpload {
	s.StsToken = &v
	return s
}

func (s *CreateTransactionResponseBodyTransactionUpload) SetOssEndpoint(v string) *CreateTransactionResponseBodyTransactionUpload {
	s.OssEndpoint = &v
	return s
}

func (s *CreateTransactionResponseBodyTransactionUpload) SetBucket(v string) *CreateTransactionResponseBodyTransactionUpload {
	s.Bucket = &v
	return s
}

func (s *CreateTransactionResponseBodyTransactionUpload) SetFileId(v string) *CreateTransactionResponseBodyTransactionUpload {
	s.FileId = &v
	return s
}

type CreateTransactionResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *CreateTransactionResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s CreateTransactionResponse) String() string {
	return tea.Prettify(s)
}

func (s CreateTransactionResponse) GoString() string {
	return s.String()
}

func (s *CreateTransactionResponse) SetHeaders(v map[string]*string) *CreateTransactionResponse {
	s.Headers = v
	return s
}

func (s *CreateTransactionResponse) SetBody(v *CreateTransactionResponseBody) *CreateTransactionResponse {
	s.Body = v
	return s
}

type DeleteAlbumsRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	AlbumId   []*int  `json:"AlbumId,omitempty" xml:"AlbumId,omitempty" type:"Repeated"`
}

func (s DeleteAlbumsRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlbumsRequest) GoString() string {
	return s.String()
}

func (s *DeleteAlbumsRequest) SetStoreName(v string) *DeleteAlbumsRequest {
	s.StoreName = &v
	return s
}

func (s *DeleteAlbumsRequest) SetLibraryId(v string) *DeleteAlbumsRequest {
	s.LibraryId = &v
	return s
}

func (s *DeleteAlbumsRequest) SetAlbumId(v []*int) *DeleteAlbumsRequest {
	s.AlbumId = v
	return s
}

type DeleteAlbumsResponseBody struct {
	Action    *string                            `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DeleteAlbumsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteAlbumsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlbumsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAlbumsResponseBody) SetAction(v string) *DeleteAlbumsResponseBody {
	s.Action = &v
	return s
}

func (s *DeleteAlbumsResponseBody) SetMessage(v string) *DeleteAlbumsResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteAlbumsResponseBody) SetRequestId(v string) *DeleteAlbumsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAlbumsResponseBody) SetResults(v []*DeleteAlbumsResponseBodyResults) *DeleteAlbumsResponseBody {
	s.Results = v
	return s
}

func (s *DeleteAlbumsResponseBody) SetCode(v string) *DeleteAlbumsResponseBody {
	s.Code = &v
	return s
}

type DeleteAlbumsResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteAlbumsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlbumsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DeleteAlbumsResponseBodyResults) SetIdStr(v string) *DeleteAlbumsResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *DeleteAlbumsResponseBodyResults) SetCode(v string) *DeleteAlbumsResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DeleteAlbumsResponseBodyResults) SetMessage(v string) *DeleteAlbumsResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DeleteAlbumsResponseBodyResults) SetId(v int64) *DeleteAlbumsResponseBodyResults {
	s.Id = &v
	return s
}

type DeleteAlbumsResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteAlbumsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteAlbumsResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteAlbumsResponse) GoString() string {
	return s.String()
}

func (s *DeleteAlbumsResponse) SetHeaders(v map[string]*string) *DeleteAlbumsResponse {
	s.Headers = v
	return s
}

func (s *DeleteAlbumsResponse) SetBody(v *DeleteAlbumsResponseBody) *DeleteAlbumsResponse {
	s.Body = v
	return s
}

type DeleteEventRequest struct {
	EventId   *int64  `json:"EventId,omitempty" xml:"EventId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s DeleteEventRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventRequest) GoString() string {
	return s.String()
}

func (s *DeleteEventRequest) SetEventId(v int64) *DeleteEventRequest {
	s.EventId = &v
	return s
}

func (s *DeleteEventRequest) SetStoreName(v string) *DeleteEventRequest {
	s.StoreName = &v
	return s
}

func (s *DeleteEventRequest) SetLibraryId(v string) *DeleteEventRequest {
	s.LibraryId = &v
	return s
}

type DeleteEventResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteEventResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteEventResponseBody) SetAction(v string) *DeleteEventResponseBody {
	s.Action = &v
	return s
}

func (s *DeleteEventResponseBody) SetMessage(v string) *DeleteEventResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteEventResponseBody) SetRequestId(v string) *DeleteEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteEventResponseBody) SetCode(v string) *DeleteEventResponseBody {
	s.Code = &v
	return s
}

type DeleteEventResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteEventResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteEventResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteEventResponse) GoString() string {
	return s.String()
}

func (s *DeleteEventResponse) SetHeaders(v map[string]*string) *DeleteEventResponse {
	s.Headers = v
	return s
}

func (s *DeleteEventResponse) SetBody(v *DeleteEventResponseBody) *DeleteEventResponse {
	s.Body = v
	return s
}

type DeleteFacesRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	FaceId    []*int  `json:"FaceId,omitempty" xml:"FaceId,omitempty" type:"Repeated"`
}

func (s DeleteFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacesRequest) GoString() string {
	return s.String()
}

func (s *DeleteFacesRequest) SetStoreName(v string) *DeleteFacesRequest {
	s.StoreName = &v
	return s
}

func (s *DeleteFacesRequest) SetLibraryId(v string) *DeleteFacesRequest {
	s.LibraryId = &v
	return s
}

func (s *DeleteFacesRequest) SetFaceId(v []*int) *DeleteFacesRequest {
	s.FaceId = v
	return s
}

type DeleteFacesResponseBody struct {
	Action    *string                           `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DeleteFacesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeleteFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacesResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteFacesResponseBody) SetAction(v string) *DeleteFacesResponseBody {
	s.Action = &v
	return s
}

func (s *DeleteFacesResponseBody) SetMessage(v string) *DeleteFacesResponseBody {
	s.Message = &v
	return s
}

func (s *DeleteFacesResponseBody) SetRequestId(v string) *DeleteFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteFacesResponseBody) SetResults(v []*DeleteFacesResponseBodyResults) *DeleteFacesResponseBody {
	s.Results = v
	return s
}

func (s *DeleteFacesResponseBody) SetCode(v string) *DeleteFacesResponseBody {
	s.Code = &v
	return s
}

type DeleteFacesResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeleteFacesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DeleteFacesResponseBodyResults) SetIdStr(v string) *DeleteFacesResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *DeleteFacesResponseBodyResults) SetCode(v string) *DeleteFacesResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DeleteFacesResponseBodyResults) SetMessage(v string) *DeleteFacesResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DeleteFacesResponseBodyResults) SetId(v int64) *DeleteFacesResponseBodyResults {
	s.Id = &v
	return s
}

type DeleteFacesResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeleteFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeleteFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s DeleteFacesResponse) GoString() string {
	return s.String()
}

func (s *DeleteFacesResponse) SetHeaders(v map[string]*string) *DeleteFacesResponse {
	s.Headers = v
	return s
}

func (s *DeleteFacesResponse) SetBody(v *DeleteFacesResponseBody) *DeleteFacesResponse {
	s.Body = v
	return s
}

type DeletePhotosRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s DeletePhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePhotosRequest) GoString() string {
	return s.String()
}

func (s *DeletePhotosRequest) SetStoreName(v string) *DeletePhotosRequest {
	s.StoreName = &v
	return s
}

func (s *DeletePhotosRequest) SetLibraryId(v string) *DeletePhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *DeletePhotosRequest) SetPhotoId(v []*int) *DeletePhotosRequest {
	s.PhotoId = v
	return s
}

type DeletePhotosResponseBody struct {
	Action    *string                            `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*DeletePhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeletePhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePhotosResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePhotosResponseBody) SetAction(v string) *DeletePhotosResponseBody {
	s.Action = &v
	return s
}

func (s *DeletePhotosResponseBody) SetMessage(v string) *DeletePhotosResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePhotosResponseBody) SetRequestId(v string) *DeletePhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePhotosResponseBody) SetResults(v []*DeletePhotosResponseBodyResults) *DeletePhotosResponseBody {
	s.Results = v
	return s
}

func (s *DeletePhotosResponseBody) SetCode(v string) *DeletePhotosResponseBody {
	s.Code = &v
	return s
}

type DeletePhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s DeletePhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s DeletePhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *DeletePhotosResponseBodyResults) SetIdStr(v string) *DeletePhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *DeletePhotosResponseBodyResults) SetCode(v string) *DeletePhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *DeletePhotosResponseBodyResults) SetMessage(v string) *DeletePhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *DeletePhotosResponseBodyResults) SetId(v int64) *DeletePhotosResponseBodyResults {
	s.Id = &v
	return s
}

type DeletePhotosResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePhotosResponse) GoString() string {
	return s.String()
}

func (s *DeletePhotosResponse) SetHeaders(v map[string]*string) *DeletePhotosResponse {
	s.Headers = v
	return s
}

func (s *DeletePhotosResponse) SetBody(v *DeletePhotosResponseBody) *DeletePhotosResponse {
	s.Body = v
	return s
}

type DeletePhotoStoreRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
}

func (s DeletePhotoStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s DeletePhotoStoreRequest) GoString() string {
	return s.String()
}

func (s *DeletePhotoStoreRequest) SetStoreName(v string) *DeletePhotoStoreRequest {
	s.StoreName = &v
	return s
}

type DeletePhotoStoreResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s DeletePhotoStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s DeletePhotoStoreResponseBody) GoString() string {
	return s.String()
}

func (s *DeletePhotoStoreResponseBody) SetAction(v string) *DeletePhotoStoreResponseBody {
	s.Action = &v
	return s
}

func (s *DeletePhotoStoreResponseBody) SetMessage(v string) *DeletePhotoStoreResponseBody {
	s.Message = &v
	return s
}

func (s *DeletePhotoStoreResponseBody) SetRequestId(v string) *DeletePhotoStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeletePhotoStoreResponseBody) SetCode(v string) *DeletePhotoStoreResponseBody {
	s.Code = &v
	return s
}

type DeletePhotoStoreResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *DeletePhotoStoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s DeletePhotoStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s DeletePhotoStoreResponse) GoString() string {
	return s.String()
}

func (s *DeletePhotoStoreResponse) SetHeaders(v map[string]*string) *DeletePhotoStoreResponse {
	s.Headers = v
	return s
}

func (s *DeletePhotoStoreResponse) SetBody(v *DeletePhotoStoreResponseBody) *DeletePhotoStoreResponse {
	s.Body = v
	return s
}

type EditPhotosRequest struct {
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	StoreName       *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId       *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId         []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s EditPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s EditPhotosRequest) GoString() string {
	return s.String()
}

func (s *EditPhotosRequest) SetShareExpireTime(v int64) *EditPhotosRequest {
	s.ShareExpireTime = &v
	return s
}

func (s *EditPhotosRequest) SetTakenAt(v int64) *EditPhotosRequest {
	s.TakenAt = &v
	return s
}

func (s *EditPhotosRequest) SetTitle(v string) *EditPhotosRequest {
	s.Title = &v
	return s
}

func (s *EditPhotosRequest) SetRemark(v string) *EditPhotosRequest {
	s.Remark = &v
	return s
}

func (s *EditPhotosRequest) SetStoreName(v string) *EditPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *EditPhotosRequest) SetLibraryId(v string) *EditPhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *EditPhotosRequest) SetPhotoId(v []*int) *EditPhotosRequest {
	s.PhotoId = v
	return s
}

type EditPhotosResponseBody struct {
	Action    *string                          `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*EditPhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s EditPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *EditPhotosResponseBody) SetAction(v string) *EditPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *EditPhotosResponseBody) SetMessage(v string) *EditPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *EditPhotosResponseBody) SetRequestId(v string) *EditPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditPhotosResponseBody) SetResults(v []*EditPhotosResponseBodyResults) *EditPhotosResponseBody {
	s.Results = v
	return s
}

func (s *EditPhotosResponseBody) SetCode(v string) *EditPhotosResponseBody {
	s.Code = &v
	return s
}

type EditPhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s EditPhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s EditPhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *EditPhotosResponseBodyResults) SetIdStr(v string) *EditPhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *EditPhotosResponseBodyResults) SetCode(v string) *EditPhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *EditPhotosResponseBodyResults) SetMessage(v string) *EditPhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *EditPhotosResponseBodyResults) SetId(v int64) *EditPhotosResponseBodyResults {
	s.Id = &v
	return s
}

type EditPhotosResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s EditPhotosResponse) GoString() string {
	return s.String()
}

func (s *EditPhotosResponse) SetHeaders(v map[string]*string) *EditPhotosResponse {
	s.Headers = v
	return s
}

func (s *EditPhotosResponse) SetBody(v *EditPhotosResponseBody) *EditPhotosResponse {
	s.Body = v
	return s
}

type EditPhotoStoreRequest struct {
	AutoCleanEnabled  *string `json:"AutoCleanEnabled,omitempty" xml:"AutoCleanEnabled,omitempty"`
	AutoCleanDays     *int32  `json:"AutoCleanDays,omitempty" xml:"AutoCleanDays,omitempty"`
	DefaultQuota      *int64  `json:"DefaultQuota,omitempty" xml:"DefaultQuota,omitempty"`
	DefaultTrashQuota *int64  `json:"DefaultTrashQuota,omitempty" xml:"DefaultTrashQuota,omitempty"`
	Remark            *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	StoreName         *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
}

func (s EditPhotoStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s EditPhotoStoreRequest) GoString() string {
	return s.String()
}

func (s *EditPhotoStoreRequest) SetAutoCleanEnabled(v string) *EditPhotoStoreRequest {
	s.AutoCleanEnabled = &v
	return s
}

func (s *EditPhotoStoreRequest) SetAutoCleanDays(v int32) *EditPhotoStoreRequest {
	s.AutoCleanDays = &v
	return s
}

func (s *EditPhotoStoreRequest) SetDefaultQuota(v int64) *EditPhotoStoreRequest {
	s.DefaultQuota = &v
	return s
}

func (s *EditPhotoStoreRequest) SetDefaultTrashQuota(v int64) *EditPhotoStoreRequest {
	s.DefaultTrashQuota = &v
	return s
}

func (s *EditPhotoStoreRequest) SetRemark(v string) *EditPhotoStoreRequest {
	s.Remark = &v
	return s
}

func (s *EditPhotoStoreRequest) SetStoreName(v string) *EditPhotoStoreRequest {
	s.StoreName = &v
	return s
}

type EditPhotoStoreResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s EditPhotoStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s EditPhotoStoreResponseBody) GoString() string {
	return s.String()
}

func (s *EditPhotoStoreResponseBody) SetAction(v string) *EditPhotoStoreResponseBody {
	s.Action = &v
	return s
}

func (s *EditPhotoStoreResponseBody) SetMessage(v string) *EditPhotoStoreResponseBody {
	s.Message = &v
	return s
}

func (s *EditPhotoStoreResponseBody) SetRequestId(v string) *EditPhotoStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *EditPhotoStoreResponseBody) SetCode(v string) *EditPhotoStoreResponseBody {
	s.Code = &v
	return s
}

type EditPhotoStoreResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *EditPhotoStoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s EditPhotoStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s EditPhotoStoreResponse) GoString() string {
	return s.String()
}

func (s *EditPhotoStoreResponse) SetHeaders(v map[string]*string) *EditPhotoStoreResponse {
	s.Headers = v
	return s
}

func (s *EditPhotoStoreResponse) SetBody(v *EditPhotoStoreResponseBody) *EditPhotoStoreResponse {
	s.Body = v
	return s
}

type FetchAlbumTagPhotosRequest struct {
	AlbumId   *int64  `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	TagId     *int64  `json:"TagId,omitempty" xml:"TagId,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s FetchAlbumTagPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchAlbumTagPhotosRequest) GoString() string {
	return s.String()
}

func (s *FetchAlbumTagPhotosRequest) SetAlbumId(v int64) *FetchAlbumTagPhotosRequest {
	s.AlbumId = &v
	return s
}

func (s *FetchAlbumTagPhotosRequest) SetTagId(v int64) *FetchAlbumTagPhotosRequest {
	s.TagId = &v
	return s
}

func (s *FetchAlbumTagPhotosRequest) SetSize(v int32) *FetchAlbumTagPhotosRequest {
	s.Size = &v
	return s
}

func (s *FetchAlbumTagPhotosRequest) SetPage(v int32) *FetchAlbumTagPhotosRequest {
	s.Page = &v
	return s
}

func (s *FetchAlbumTagPhotosRequest) SetStoreName(v string) *FetchAlbumTagPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *FetchAlbumTagPhotosRequest) SetLibraryId(v string) *FetchAlbumTagPhotosRequest {
	s.LibraryId = &v
	return s
}

type FetchAlbumTagPhotosResponseBody struct {
	Action     *string                                   `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                                    `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results    []*FetchAlbumTagPhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code       *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s FetchAlbumTagPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchAlbumTagPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *FetchAlbumTagPhotosResponseBody) SetAction(v string) *FetchAlbumTagPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *FetchAlbumTagPhotosResponseBody) SetTotalCount(v int32) *FetchAlbumTagPhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *FetchAlbumTagPhotosResponseBody) SetMessage(v string) *FetchAlbumTagPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *FetchAlbumTagPhotosResponseBody) SetRequestId(v string) *FetchAlbumTagPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchAlbumTagPhotosResponseBody) SetResults(v []*FetchAlbumTagPhotosResponseBodyResults) *FetchAlbumTagPhotosResponseBody {
	s.Results = v
	return s
}

func (s *FetchAlbumTagPhotosResponseBody) SetCode(v string) *FetchAlbumTagPhotosResponseBody {
	s.Code = &v
	return s
}

type FetchAlbumTagPhotosResponseBodyResults struct {
	PhotoIdStr *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	Mtime      *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	PhotoId    *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
}

func (s FetchAlbumTagPhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s FetchAlbumTagPhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *FetchAlbumTagPhotosResponseBodyResults) SetPhotoIdStr(v string) *FetchAlbumTagPhotosResponseBodyResults {
	s.PhotoIdStr = &v
	return s
}

func (s *FetchAlbumTagPhotosResponseBodyResults) SetMtime(v int64) *FetchAlbumTagPhotosResponseBodyResults {
	s.Mtime = &v
	return s
}

func (s *FetchAlbumTagPhotosResponseBodyResults) SetState(v string) *FetchAlbumTagPhotosResponseBodyResults {
	s.State = &v
	return s
}

func (s *FetchAlbumTagPhotosResponseBodyResults) SetPhotoId(v int64) *FetchAlbumTagPhotosResponseBodyResults {
	s.PhotoId = &v
	return s
}

type FetchAlbumTagPhotosResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FetchAlbumTagPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FetchAlbumTagPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchAlbumTagPhotosResponse) GoString() string {
	return s.String()
}

func (s *FetchAlbumTagPhotosResponse) SetHeaders(v map[string]*string) *FetchAlbumTagPhotosResponse {
	s.Headers = v
	return s
}

func (s *FetchAlbumTagPhotosResponse) SetBody(v *FetchAlbumTagPhotosResponseBody) *FetchAlbumTagPhotosResponse {
	s.Body = v
	return s
}

type FetchLibrariesRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	NeedQuota *bool   `json:"NeedQuota,omitempty" xml:"NeedQuota,omitempty"`
}

func (s FetchLibrariesRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchLibrariesRequest) GoString() string {
	return s.String()
}

func (s *FetchLibrariesRequest) SetStoreName(v string) *FetchLibrariesRequest {
	s.StoreName = &v
	return s
}

func (s *FetchLibrariesRequest) SetPage(v int32) *FetchLibrariesRequest {
	s.Page = &v
	return s
}

func (s *FetchLibrariesRequest) SetSize(v int32) *FetchLibrariesRequest {
	s.Size = &v
	return s
}

func (s *FetchLibrariesRequest) SetNeedQuota(v bool) *FetchLibrariesRequest {
	s.NeedQuota = &v
	return s
}

type FetchLibrariesResponseBody struct {
	Action     *string                                `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Libraries  []*FetchLibrariesResponseBodyLibraries `json:"Libraries,omitempty" xml:"Libraries,omitempty" type:"Repeated"`
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s FetchLibrariesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchLibrariesResponseBody) GoString() string {
	return s.String()
}

func (s *FetchLibrariesResponseBody) SetAction(v string) *FetchLibrariesResponseBody {
	s.Action = &v
	return s
}

func (s *FetchLibrariesResponseBody) SetTotalCount(v int32) *FetchLibrariesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *FetchLibrariesResponseBody) SetMessage(v string) *FetchLibrariesResponseBody {
	s.Message = &v
	return s
}

func (s *FetchLibrariesResponseBody) SetRequestId(v string) *FetchLibrariesResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchLibrariesResponseBody) SetLibraries(v []*FetchLibrariesResponseBodyLibraries) *FetchLibrariesResponseBody {
	s.Libraries = v
	return s
}

func (s *FetchLibrariesResponseBody) SetCode(v string) *FetchLibrariesResponseBody {
	s.Code = &v
	return s
}

type FetchLibrariesResponseBodyLibraries struct {
	Ctime      *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	LibraryId  *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	TotalQuota *int64  `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
}

func (s FetchLibrariesResponseBodyLibraries) String() string {
	return tea.Prettify(s)
}

func (s FetchLibrariesResponseBodyLibraries) GoString() string {
	return s.String()
}

func (s *FetchLibrariesResponseBodyLibraries) SetCtime(v int64) *FetchLibrariesResponseBodyLibraries {
	s.Ctime = &v
	return s
}

func (s *FetchLibrariesResponseBodyLibraries) SetLibraryId(v string) *FetchLibrariesResponseBodyLibraries {
	s.LibraryId = &v
	return s
}

func (s *FetchLibrariesResponseBodyLibraries) SetTotalQuota(v int64) *FetchLibrariesResponseBodyLibraries {
	s.TotalQuota = &v
	return s
}

type FetchLibrariesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FetchLibrariesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FetchLibrariesResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchLibrariesResponse) GoString() string {
	return s.String()
}

func (s *FetchLibrariesResponse) SetHeaders(v map[string]*string) *FetchLibrariesResponse {
	s.Headers = v
	return s
}

func (s *FetchLibrariesResponse) SetBody(v *FetchLibrariesResponseBody) *FetchLibrariesResponse {
	s.Body = v
	return s
}

type FetchMomentPhotosRequest struct {
	MomentId  *int64  `json:"MomentId,omitempty" xml:"MomentId,omitempty"`
	OrderBy   *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Order     *string `json:"Order,omitempty" xml:"Order,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s FetchMomentPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchMomentPhotosRequest) GoString() string {
	return s.String()
}

func (s *FetchMomentPhotosRequest) SetMomentId(v int64) *FetchMomentPhotosRequest {
	s.MomentId = &v
	return s
}

func (s *FetchMomentPhotosRequest) SetOrderBy(v string) *FetchMomentPhotosRequest {
	s.OrderBy = &v
	return s
}

func (s *FetchMomentPhotosRequest) SetOrder(v string) *FetchMomentPhotosRequest {
	s.Order = &v
	return s
}

func (s *FetchMomentPhotosRequest) SetSize(v int32) *FetchMomentPhotosRequest {
	s.Size = &v
	return s
}

func (s *FetchMomentPhotosRequest) SetPage(v int32) *FetchMomentPhotosRequest {
	s.Page = &v
	return s
}

func (s *FetchMomentPhotosRequest) SetStoreName(v string) *FetchMomentPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *FetchMomentPhotosRequest) SetLibraryId(v string) *FetchMomentPhotosRequest {
	s.LibraryId = &v
	return s
}

type FetchMomentPhotosResponseBody struct {
	Photos     []*FetchMomentPhotosResponseBodyPhotos `json:"Photos,omitempty" xml:"Photos,omitempty" type:"Repeated"`
	Action     *string                                `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s FetchMomentPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchMomentPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *FetchMomentPhotosResponseBody) SetPhotos(v []*FetchMomentPhotosResponseBodyPhotos) *FetchMomentPhotosResponseBody {
	s.Photos = v
	return s
}

func (s *FetchMomentPhotosResponseBody) SetAction(v string) *FetchMomentPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *FetchMomentPhotosResponseBody) SetTotalCount(v int32) *FetchMomentPhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *FetchMomentPhotosResponseBody) SetMessage(v string) *FetchMomentPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *FetchMomentPhotosResponseBody) SetRequestId(v string) *FetchMomentPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchMomentPhotosResponseBody) SetCode(v string) *FetchMomentPhotosResponseBody {
	s.Code = &v
	return s
}

type FetchMomentPhotosResponseBodyPhotos struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	InactiveTime    *int64  `json:"InactiveTime,omitempty" xml:"InactiveTime,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FetchMomentPhotosResponseBodyPhotos) String() string {
	return tea.Prettify(s)
}

func (s FetchMomentPhotosResponseBodyPhotos) GoString() string {
	return s.String()
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetRemark(v string) *FetchMomentPhotosResponseBodyPhotos {
	s.Remark = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetTakenAt(v int64) *FetchMomentPhotosResponseBodyPhotos {
	s.TakenAt = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetState(v string) *FetchMomentPhotosResponseBodyPhotos {
	s.State = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetHeight(v int64) *FetchMomentPhotosResponseBodyPhotos {
	s.Height = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetShareExpireTime(v int64) *FetchMomentPhotosResponseBodyPhotos {
	s.ShareExpireTime = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetFileId(v string) *FetchMomentPhotosResponseBodyPhotos {
	s.FileId = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetIdStr(v string) *FetchMomentPhotosResponseBodyPhotos {
	s.IdStr = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetCtime(v int64) *FetchMomentPhotosResponseBodyPhotos {
	s.Ctime = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetMtime(v int64) *FetchMomentPhotosResponseBodyPhotos {
	s.Mtime = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetSize(v int64) *FetchMomentPhotosResponseBodyPhotos {
	s.Size = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetWidth(v int64) *FetchMomentPhotosResponseBodyPhotos {
	s.Width = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetInactiveTime(v int64) *FetchMomentPhotosResponseBodyPhotos {
	s.InactiveTime = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetMd5(v string) *FetchMomentPhotosResponseBodyPhotos {
	s.Md5 = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetIsVideo(v bool) *FetchMomentPhotosResponseBodyPhotos {
	s.IsVideo = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetTitle(v string) *FetchMomentPhotosResponseBodyPhotos {
	s.Title = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetLocation(v string) *FetchMomentPhotosResponseBodyPhotos {
	s.Location = &v
	return s
}

func (s *FetchMomentPhotosResponseBodyPhotos) SetId(v int64) *FetchMomentPhotosResponseBodyPhotos {
	s.Id = &v
	return s
}

type FetchMomentPhotosResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FetchMomentPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FetchMomentPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchMomentPhotosResponse) GoString() string {
	return s.String()
}

func (s *FetchMomentPhotosResponse) SetHeaders(v map[string]*string) *FetchMomentPhotosResponse {
	s.Headers = v
	return s
}

func (s *FetchMomentPhotosResponse) SetBody(v *FetchMomentPhotosResponseBody) *FetchMomentPhotosResponse {
	s.Body = v
	return s
}

type FetchPhotosRequest struct {
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	OrderBy   *string `json:"OrderBy,omitempty" xml:"OrderBy,omitempty"`
	Order     *string `json:"Order,omitempty" xml:"Order,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s FetchPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s FetchPhotosRequest) GoString() string {
	return s.String()
}

func (s *FetchPhotosRequest) SetState(v string) *FetchPhotosRequest {
	s.State = &v
	return s
}

func (s *FetchPhotosRequest) SetOrderBy(v string) *FetchPhotosRequest {
	s.OrderBy = &v
	return s
}

func (s *FetchPhotosRequest) SetOrder(v string) *FetchPhotosRequest {
	s.Order = &v
	return s
}

func (s *FetchPhotosRequest) SetSize(v int32) *FetchPhotosRequest {
	s.Size = &v
	return s
}

func (s *FetchPhotosRequest) SetPage(v int32) *FetchPhotosRequest {
	s.Page = &v
	return s
}

func (s *FetchPhotosRequest) SetStoreName(v string) *FetchPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *FetchPhotosRequest) SetLibraryId(v string) *FetchPhotosRequest {
	s.LibraryId = &v
	return s
}

type FetchPhotosResponseBody struct {
	Photos     []*FetchPhotosResponseBodyPhotos `json:"Photos,omitempty" xml:"Photos,omitempty" type:"Repeated"`
	Action     *string                          `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                           `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s FetchPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s FetchPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *FetchPhotosResponseBody) SetPhotos(v []*FetchPhotosResponseBodyPhotos) *FetchPhotosResponseBody {
	s.Photos = v
	return s
}

func (s *FetchPhotosResponseBody) SetAction(v string) *FetchPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *FetchPhotosResponseBody) SetTotalCount(v int32) *FetchPhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *FetchPhotosResponseBody) SetMessage(v string) *FetchPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *FetchPhotosResponseBody) SetRequestId(v string) *FetchPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *FetchPhotosResponseBody) SetCode(v string) *FetchPhotosResponseBody {
	s.Code = &v
	return s
}

type FetchPhotosResponseBodyPhotos struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	InactiveTime    *int64  `json:"InactiveTime,omitempty" xml:"InactiveTime,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s FetchPhotosResponseBodyPhotos) String() string {
	return tea.Prettify(s)
}

func (s FetchPhotosResponseBodyPhotos) GoString() string {
	return s.String()
}

func (s *FetchPhotosResponseBodyPhotos) SetRemark(v string) *FetchPhotosResponseBodyPhotos {
	s.Remark = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetTakenAt(v int64) *FetchPhotosResponseBodyPhotos {
	s.TakenAt = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetState(v string) *FetchPhotosResponseBodyPhotos {
	s.State = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetHeight(v int64) *FetchPhotosResponseBodyPhotos {
	s.Height = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetShareExpireTime(v int64) *FetchPhotosResponseBodyPhotos {
	s.ShareExpireTime = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetFileId(v string) *FetchPhotosResponseBodyPhotos {
	s.FileId = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetIdStr(v string) *FetchPhotosResponseBodyPhotos {
	s.IdStr = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetCtime(v int64) *FetchPhotosResponseBodyPhotos {
	s.Ctime = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetMtime(v int64) *FetchPhotosResponseBodyPhotos {
	s.Mtime = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetSize(v int64) *FetchPhotosResponseBodyPhotos {
	s.Size = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetWidth(v int64) *FetchPhotosResponseBodyPhotos {
	s.Width = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetInactiveTime(v int64) *FetchPhotosResponseBodyPhotos {
	s.InactiveTime = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetMd5(v string) *FetchPhotosResponseBodyPhotos {
	s.Md5 = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetIsVideo(v bool) *FetchPhotosResponseBodyPhotos {
	s.IsVideo = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetTitle(v string) *FetchPhotosResponseBodyPhotos {
	s.Title = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetLocation(v string) *FetchPhotosResponseBodyPhotos {
	s.Location = &v
	return s
}

func (s *FetchPhotosResponseBodyPhotos) SetId(v int64) *FetchPhotosResponseBodyPhotos {
	s.Id = &v
	return s
}

type FetchPhotosResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *FetchPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s FetchPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s FetchPhotosResponse) GoString() string {
	return s.String()
}

func (s *FetchPhotosResponse) SetHeaders(v map[string]*string) *FetchPhotosResponse {
	s.Headers = v
	return s
}

func (s *FetchPhotosResponse) SetBody(v *FetchPhotosResponseBody) *FetchPhotosResponse {
	s.Body = v
	return s
}

type GetAlbumsByNamesRequest struct {
	StoreName *string   `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string   `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	Name      []*string `json:"Name,omitempty" xml:"Name,omitempty" type:"Repeated"`
}

func (s GetAlbumsByNamesRequest) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumsByNamesRequest) GoString() string {
	return s.String()
}

func (s *GetAlbumsByNamesRequest) SetStoreName(v string) *GetAlbumsByNamesRequest {
	s.StoreName = &v
	return s
}

func (s *GetAlbumsByNamesRequest) SetLibraryId(v string) *GetAlbumsByNamesRequest {
	s.LibraryId = &v
	return s
}

func (s *GetAlbumsByNamesRequest) SetName(v []*string) *GetAlbumsByNamesRequest {
	s.Name = v
	return s
}

type GetAlbumsByNamesResponseBody struct {
	Action    *string                               `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
	Albums    []*GetAlbumsByNamesResponseBodyAlbums `json:"Albums,omitempty" xml:"Albums,omitempty" type:"Repeated"`
}

func (s GetAlbumsByNamesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumsByNamesResponseBody) GoString() string {
	return s.String()
}

func (s *GetAlbumsByNamesResponseBody) SetAction(v string) *GetAlbumsByNamesResponseBody {
	s.Action = &v
	return s
}

func (s *GetAlbumsByNamesResponseBody) SetMessage(v string) *GetAlbumsByNamesResponseBody {
	s.Message = &v
	return s
}

func (s *GetAlbumsByNamesResponseBody) SetRequestId(v string) *GetAlbumsByNamesResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAlbumsByNamesResponseBody) SetCode(v string) *GetAlbumsByNamesResponseBody {
	s.Code = &v
	return s
}

func (s *GetAlbumsByNamesResponseBody) SetAlbums(v []*GetAlbumsByNamesResponseBodyAlbums) *GetAlbumsByNamesResponseBody {
	s.Albums = v
	return s
}

type GetAlbumsByNamesResponseBodyAlbums struct {
	IdStr       *string                                  `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	PhotosCount *int64                                   `json:"PhotosCount,omitempty" xml:"PhotosCount,omitempty"`
	Cover       *GetAlbumsByNamesResponseBodyAlbumsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Mtime       *int64                                   `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime       *int64                                   `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	State       *string                                  `json:"State,omitempty" xml:"State,omitempty"`
	Name        *string                                  `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *int64                                   `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetAlbumsByNamesResponseBodyAlbums) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumsByNamesResponseBodyAlbums) GoString() string {
	return s.String()
}

func (s *GetAlbumsByNamesResponseBodyAlbums) SetIdStr(v string) *GetAlbumsByNamesResponseBodyAlbums {
	s.IdStr = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbums) SetPhotosCount(v int64) *GetAlbumsByNamesResponseBodyAlbums {
	s.PhotosCount = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbums) SetCover(v *GetAlbumsByNamesResponseBodyAlbumsCover) *GetAlbumsByNamesResponseBodyAlbums {
	s.Cover = v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbums) SetMtime(v int64) *GetAlbumsByNamesResponseBodyAlbums {
	s.Mtime = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbums) SetCtime(v int64) *GetAlbumsByNamesResponseBodyAlbums {
	s.Ctime = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbums) SetState(v string) *GetAlbumsByNamesResponseBodyAlbums {
	s.State = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbums) SetName(v string) *GetAlbumsByNamesResponseBodyAlbums {
	s.Name = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbums) SetId(v int64) *GetAlbumsByNamesResponseBodyAlbums {
	s.Id = &v
	return s
}

type GetAlbumsByNamesResponseBodyAlbumsCover struct {
	Remark  *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	State   *string `json:"State,omitempty" xml:"State,omitempty"`
	Height  *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	FileId  *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Mtime   *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime   *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Width   *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Md5     *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetAlbumsByNamesResponseBodyAlbumsCover) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumsByNamesResponseBodyAlbumsCover) GoString() string {
	return s.String()
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetRemark(v string) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.Remark = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetState(v string) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.State = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetHeight(v int64) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.Height = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetFileId(v string) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.FileId = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetIdStr(v string) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.IdStr = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetMtime(v int64) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.Mtime = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetCtime(v int64) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.Ctime = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetWidth(v int64) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.Width = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetMd5(v string) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.Md5 = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetIsVideo(v bool) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.IsVideo = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetTitle(v string) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.Title = &v
	return s
}

func (s *GetAlbumsByNamesResponseBodyAlbumsCover) SetId(v int64) *GetAlbumsByNamesResponseBodyAlbumsCover {
	s.Id = &v
	return s
}

type GetAlbumsByNamesResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetAlbumsByNamesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetAlbumsByNamesResponse) String() string {
	return tea.Prettify(s)
}

func (s GetAlbumsByNamesResponse) GoString() string {
	return s.String()
}

func (s *GetAlbumsByNamesResponse) SetHeaders(v map[string]*string) *GetAlbumsByNamesResponse {
	s.Headers = v
	return s
}

func (s *GetAlbumsByNamesResponse) SetBody(v *GetAlbumsByNamesResponseBody) *GetAlbumsByNamesResponse {
	s.Body = v
	return s
}

type GetDownloadUrlRequest struct {
	PhotoId   *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s GetDownloadUrlRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlRequest) SetPhotoId(v int64) *GetDownloadUrlRequest {
	s.PhotoId = &v
	return s
}

func (s *GetDownloadUrlRequest) SetStoreName(v string) *GetDownloadUrlRequest {
	s.StoreName = &v
	return s
}

func (s *GetDownloadUrlRequest) SetLibraryId(v string) *GetDownloadUrlRequest {
	s.LibraryId = &v
	return s
}

type GetDownloadUrlResponseBody struct {
	Action      *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
}

func (s GetDownloadUrlResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlResponseBody) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlResponseBody) SetAction(v string) *GetDownloadUrlResponseBody {
	s.Action = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetMessage(v string) *GetDownloadUrlResponseBody {
	s.Message = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetRequestId(v string) *GetDownloadUrlResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetCode(v string) *GetDownloadUrlResponseBody {
	s.Code = &v
	return s
}

func (s *GetDownloadUrlResponseBody) SetDownloadUrl(v string) *GetDownloadUrlResponseBody {
	s.DownloadUrl = &v
	return s
}

type GetDownloadUrlResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDownloadUrlResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *GetDownloadUrlResponse) SetBody(v *GetDownloadUrlResponseBody) *GetDownloadUrlResponse {
	s.Body = v
	return s
}

type GetDownloadUrlsRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s GetDownloadUrlsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlsRequest) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlsRequest) SetStoreName(v string) *GetDownloadUrlsRequest {
	s.StoreName = &v
	return s
}

func (s *GetDownloadUrlsRequest) SetLibraryId(v string) *GetDownloadUrlsRequest {
	s.LibraryId = &v
	return s
}

func (s *GetDownloadUrlsRequest) SetPhotoId(v []*int) *GetDownloadUrlsRequest {
	s.PhotoId = v
	return s
}

type GetDownloadUrlsResponseBody struct {
	Action    *string                             `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *GetDownloadUrlsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	Code      *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetDownloadUrlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlsResponseBody) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlsResponseBody) SetAction(v string) *GetDownloadUrlsResponseBody {
	s.Action = &v
	return s
}

func (s *GetDownloadUrlsResponseBody) SetMessage(v string) *GetDownloadUrlsResponseBody {
	s.Message = &v
	return s
}

func (s *GetDownloadUrlsResponseBody) SetRequestId(v string) *GetDownloadUrlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDownloadUrlsResponseBody) SetResults(v *GetDownloadUrlsResponseBodyResults) *GetDownloadUrlsResponseBody {
	s.Results = v
	return s
}

func (s *GetDownloadUrlsResponseBody) SetCode(v string) *GetDownloadUrlsResponseBody {
	s.Code = &v
	return s
}

type GetDownloadUrlsResponseBodyResults struct {
	Result []*GetDownloadUrlsResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetDownloadUrlsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlsResponseBodyResults) SetResult(v []*GetDownloadUrlsResponseBodyResultsResult) *GetDownloadUrlsResponseBodyResults {
	s.Result = v
	return s
}

type GetDownloadUrlsResponseBodyResultsResult struct {
	PhotoIdStr  *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	Code        *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message     *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PhotoId     *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
}

func (s GetDownloadUrlsResponseBodyResultsResult) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlsResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlsResponseBodyResultsResult) SetPhotoIdStr(v string) *GetDownloadUrlsResponseBodyResultsResult {
	s.PhotoIdStr = &v
	return s
}

func (s *GetDownloadUrlsResponseBodyResultsResult) SetDownloadUrl(v string) *GetDownloadUrlsResponseBodyResultsResult {
	s.DownloadUrl = &v
	return s
}

func (s *GetDownloadUrlsResponseBodyResultsResult) SetCode(v string) *GetDownloadUrlsResponseBodyResultsResult {
	s.Code = &v
	return s
}

func (s *GetDownloadUrlsResponseBodyResultsResult) SetMessage(v string) *GetDownloadUrlsResponseBodyResultsResult {
	s.Message = &v
	return s
}

func (s *GetDownloadUrlsResponseBodyResultsResult) SetPhotoId(v int64) *GetDownloadUrlsResponseBodyResultsResult {
	s.PhotoId = &v
	return s
}

type GetDownloadUrlsResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetDownloadUrlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetDownloadUrlsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetDownloadUrlsResponse) GoString() string {
	return s.String()
}

func (s *GetDownloadUrlsResponse) SetHeaders(v map[string]*string) *GetDownloadUrlsResponse {
	s.Headers = v
	return s
}

func (s *GetDownloadUrlsResponse) SetBody(v *GetDownloadUrlsResponseBody) *GetDownloadUrlsResponse {
	s.Body = v
	return s
}

type GetFramedPhotoUrlsRequest struct {
	FrameId   *string `json:"FrameId,omitempty" xml:"FrameId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s GetFramedPhotoUrlsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetFramedPhotoUrlsRequest) GoString() string {
	return s.String()
}

func (s *GetFramedPhotoUrlsRequest) SetFrameId(v string) *GetFramedPhotoUrlsRequest {
	s.FrameId = &v
	return s
}

func (s *GetFramedPhotoUrlsRequest) SetStoreName(v string) *GetFramedPhotoUrlsRequest {
	s.StoreName = &v
	return s
}

func (s *GetFramedPhotoUrlsRequest) SetLibraryId(v string) *GetFramedPhotoUrlsRequest {
	s.LibraryId = &v
	return s
}

func (s *GetFramedPhotoUrlsRequest) SetPhotoId(v []*int) *GetFramedPhotoUrlsRequest {
	s.PhotoId = v
	return s
}

type GetFramedPhotoUrlsResponseBody struct {
	Action    *string                                `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *GetFramedPhotoUrlsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetFramedPhotoUrlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetFramedPhotoUrlsResponseBody) GoString() string {
	return s.String()
}

func (s *GetFramedPhotoUrlsResponseBody) SetAction(v string) *GetFramedPhotoUrlsResponseBody {
	s.Action = &v
	return s
}

func (s *GetFramedPhotoUrlsResponseBody) SetMessage(v string) *GetFramedPhotoUrlsResponseBody {
	s.Message = &v
	return s
}

func (s *GetFramedPhotoUrlsResponseBody) SetRequestId(v string) *GetFramedPhotoUrlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetFramedPhotoUrlsResponseBody) SetResults(v *GetFramedPhotoUrlsResponseBodyResults) *GetFramedPhotoUrlsResponseBody {
	s.Results = v
	return s
}

func (s *GetFramedPhotoUrlsResponseBody) SetCode(v string) *GetFramedPhotoUrlsResponseBody {
	s.Code = &v
	return s
}

type GetFramedPhotoUrlsResponseBodyResults struct {
	Result []*GetFramedPhotoUrlsResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetFramedPhotoUrlsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetFramedPhotoUrlsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetFramedPhotoUrlsResponseBodyResults) SetResult(v []*GetFramedPhotoUrlsResponseBodyResultsResult) *GetFramedPhotoUrlsResponseBodyResults {
	s.Result = v
	return s
}

type GetFramedPhotoUrlsResponseBodyResultsResult struct {
	PhotoIdStr     *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	FramedPhotoUrl *string `json:"FramedPhotoUrl,omitempty" xml:"FramedPhotoUrl,omitempty"`
	Code           *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message        *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PhotoId        *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
}

func (s GetFramedPhotoUrlsResponseBodyResultsResult) String() string {
	return tea.Prettify(s)
}

func (s GetFramedPhotoUrlsResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *GetFramedPhotoUrlsResponseBodyResultsResult) SetPhotoIdStr(v string) *GetFramedPhotoUrlsResponseBodyResultsResult {
	s.PhotoIdStr = &v
	return s
}

func (s *GetFramedPhotoUrlsResponseBodyResultsResult) SetFramedPhotoUrl(v string) *GetFramedPhotoUrlsResponseBodyResultsResult {
	s.FramedPhotoUrl = &v
	return s
}

func (s *GetFramedPhotoUrlsResponseBodyResultsResult) SetCode(v string) *GetFramedPhotoUrlsResponseBodyResultsResult {
	s.Code = &v
	return s
}

func (s *GetFramedPhotoUrlsResponseBodyResultsResult) SetMessage(v string) *GetFramedPhotoUrlsResponseBodyResultsResult {
	s.Message = &v
	return s
}

func (s *GetFramedPhotoUrlsResponseBodyResultsResult) SetPhotoId(v int64) *GetFramedPhotoUrlsResponseBodyResultsResult {
	s.PhotoId = &v
	return s
}

type GetFramedPhotoUrlsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetFramedPhotoUrlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetFramedPhotoUrlsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetFramedPhotoUrlsResponse) GoString() string {
	return s.String()
}

func (s *GetFramedPhotoUrlsResponse) SetHeaders(v map[string]*string) *GetFramedPhotoUrlsResponse {
	s.Headers = v
	return s
}

func (s *GetFramedPhotoUrlsResponse) SetBody(v *GetFramedPhotoUrlsResponseBody) *GetFramedPhotoUrlsResponse {
	s.Body = v
	return s
}

type GetLibraryRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s GetLibraryRequest) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryRequest) GoString() string {
	return s.String()
}

func (s *GetLibraryRequest) SetStoreName(v string) *GetLibraryRequest {
	s.StoreName = &v
	return s
}

func (s *GetLibraryRequest) SetLibraryId(v string) *GetLibraryRequest {
	s.LibraryId = &v
	return s
}

type GetLibraryResponseBody struct {
	Action    *string                        `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Library   *GetLibraryResponseBodyLibrary `json:"Library,omitempty" xml:"Library,omitempty" type:"Struct"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetLibraryResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBody) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBody) SetAction(v string) *GetLibraryResponseBody {
	s.Action = &v
	return s
}

func (s *GetLibraryResponseBody) SetMessage(v string) *GetLibraryResponseBody {
	s.Message = &v
	return s
}

func (s *GetLibraryResponseBody) SetRequestId(v string) *GetLibraryResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetLibraryResponseBody) SetLibrary(v *GetLibraryResponseBodyLibrary) *GetLibraryResponseBody {
	s.Library = v
	return s
}

func (s *GetLibraryResponseBody) SetCode(v string) *GetLibraryResponseBody {
	s.Code = &v
	return s
}

type GetLibraryResponseBodyLibrary struct {
	AutoCleanConfig *GetLibraryResponseBodyLibraryAutoCleanConfig `json:"AutoCleanConfig,omitempty" xml:"AutoCleanConfig,omitempty" type:"Struct"`
	Quota           *GetLibraryResponseBodyLibraryQuota           `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	Ctime           *int64                                        `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
}

func (s GetLibraryResponseBodyLibrary) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyLibrary) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyLibrary) SetAutoCleanConfig(v *GetLibraryResponseBodyLibraryAutoCleanConfig) *GetLibraryResponseBodyLibrary {
	s.AutoCleanConfig = v
	return s
}

func (s *GetLibraryResponseBodyLibrary) SetQuota(v *GetLibraryResponseBodyLibraryQuota) *GetLibraryResponseBodyLibrary {
	s.Quota = v
	return s
}

func (s *GetLibraryResponseBodyLibrary) SetCtime(v int64) *GetLibraryResponseBodyLibrary {
	s.Ctime = &v
	return s
}

type GetLibraryResponseBodyLibraryAutoCleanConfig struct {
	AutoCleanDays    *int32 `json:"AutoCleanDays,omitempty" xml:"AutoCleanDays,omitempty"`
	AutoCleanEnabled *bool  `json:"AutoCleanEnabled,omitempty" xml:"AutoCleanEnabled,omitempty"`
}

func (s GetLibraryResponseBodyLibraryAutoCleanConfig) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyLibraryAutoCleanConfig) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyLibraryAutoCleanConfig) SetAutoCleanDays(v int32) *GetLibraryResponseBodyLibraryAutoCleanConfig {
	s.AutoCleanDays = &v
	return s
}

func (s *GetLibraryResponseBodyLibraryAutoCleanConfig) SetAutoCleanEnabled(v bool) *GetLibraryResponseBodyLibraryAutoCleanConfig {
	s.AutoCleanEnabled = &v
	return s
}

type GetLibraryResponseBodyLibraryQuota struct {
	PhotosCount     *int32 `json:"PhotosCount,omitempty" xml:"PhotosCount,omitempty"`
	TotalTrashQuota *int64 `json:"TotalTrashQuota,omitempty" xml:"TotalTrashQuota,omitempty"`
	InactiveSize    *int64 `json:"InactiveSize,omitempty" xml:"InactiveSize,omitempty"`
	ActiveSize      *int64 `json:"ActiveSize,omitempty" xml:"ActiveSize,omitempty"`
	FacesCount      *int32 `json:"FacesCount,omitempty" xml:"FacesCount,omitempty"`
	VideosCount     *int32 `json:"VideosCount,omitempty" xml:"VideosCount,omitempty"`
	UsedQuota       *int64 `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
	TotalQuota      *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
}

func (s GetLibraryResponseBodyLibraryQuota) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponseBodyLibraryQuota) GoString() string {
	return s.String()
}

func (s *GetLibraryResponseBodyLibraryQuota) SetPhotosCount(v int32) *GetLibraryResponseBodyLibraryQuota {
	s.PhotosCount = &v
	return s
}

func (s *GetLibraryResponseBodyLibraryQuota) SetTotalTrashQuota(v int64) *GetLibraryResponseBodyLibraryQuota {
	s.TotalTrashQuota = &v
	return s
}

func (s *GetLibraryResponseBodyLibraryQuota) SetInactiveSize(v int64) *GetLibraryResponseBodyLibraryQuota {
	s.InactiveSize = &v
	return s
}

func (s *GetLibraryResponseBodyLibraryQuota) SetActiveSize(v int64) *GetLibraryResponseBodyLibraryQuota {
	s.ActiveSize = &v
	return s
}

func (s *GetLibraryResponseBodyLibraryQuota) SetFacesCount(v int32) *GetLibraryResponseBodyLibraryQuota {
	s.FacesCount = &v
	return s
}

func (s *GetLibraryResponseBodyLibraryQuota) SetVideosCount(v int32) *GetLibraryResponseBodyLibraryQuota {
	s.VideosCount = &v
	return s
}

func (s *GetLibraryResponseBodyLibraryQuota) SetUsedQuota(v int64) *GetLibraryResponseBodyLibraryQuota {
	s.UsedQuota = &v
	return s
}

func (s *GetLibraryResponseBodyLibraryQuota) SetTotalQuota(v int64) *GetLibraryResponseBodyLibraryQuota {
	s.TotalQuota = &v
	return s
}

type GetLibraryResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetLibraryResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetLibraryResponse) String() string {
	return tea.Prettify(s)
}

func (s GetLibraryResponse) GoString() string {
	return s.String()
}

func (s *GetLibraryResponse) SetHeaders(v map[string]*string) *GetLibraryResponse {
	s.Headers = v
	return s
}

func (s *GetLibraryResponse) SetBody(v *GetLibraryResponseBody) *GetLibraryResponse {
	s.Body = v
	return s
}

type GetPhotosRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s GetPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhotosRequest) GoString() string {
	return s.String()
}

func (s *GetPhotosRequest) SetStoreName(v string) *GetPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *GetPhotosRequest) SetLibraryId(v string) *GetPhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *GetPhotosRequest) SetPhotoId(v []*int) *GetPhotosRequest {
	s.PhotoId = v
	return s
}

type GetPhotosResponseBody struct {
	Photos    []*GetPhotosResponseBodyPhotos `json:"Photos,omitempty" xml:"Photos,omitempty" type:"Repeated"`
	Action    *string                        `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhotosResponseBody) SetPhotos(v []*GetPhotosResponseBodyPhotos) *GetPhotosResponseBody {
	s.Photos = v
	return s
}

func (s *GetPhotosResponseBody) SetAction(v string) *GetPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *GetPhotosResponseBody) SetMessage(v string) *GetPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhotosResponseBody) SetRequestId(v string) *GetPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhotosResponseBody) SetCode(v string) *GetPhotosResponseBody {
	s.Code = &v
	return s
}

type GetPhotosResponseBodyPhotos struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	InactiveTime    *int64  `json:"InactiveTime,omitempty" xml:"InactiveTime,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Like            *int64  `json:"Like,omitempty" xml:"Like,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetPhotosResponseBodyPhotos) String() string {
	return tea.Prettify(s)
}

func (s GetPhotosResponseBodyPhotos) GoString() string {
	return s.String()
}

func (s *GetPhotosResponseBodyPhotos) SetRemark(v string) *GetPhotosResponseBodyPhotos {
	s.Remark = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetTakenAt(v int64) *GetPhotosResponseBodyPhotos {
	s.TakenAt = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetState(v string) *GetPhotosResponseBodyPhotos {
	s.State = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetHeight(v int64) *GetPhotosResponseBodyPhotos {
	s.Height = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetShareExpireTime(v int64) *GetPhotosResponseBodyPhotos {
	s.ShareExpireTime = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetFileId(v string) *GetPhotosResponseBodyPhotos {
	s.FileId = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetIdStr(v string) *GetPhotosResponseBodyPhotos {
	s.IdStr = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetCtime(v int64) *GetPhotosResponseBodyPhotos {
	s.Ctime = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetMtime(v int64) *GetPhotosResponseBodyPhotos {
	s.Mtime = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetWidth(v int64) *GetPhotosResponseBodyPhotos {
	s.Width = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetSize(v int64) *GetPhotosResponseBodyPhotos {
	s.Size = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetInactiveTime(v int64) *GetPhotosResponseBodyPhotos {
	s.InactiveTime = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetMd5(v string) *GetPhotosResponseBodyPhotos {
	s.Md5 = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetIsVideo(v bool) *GetPhotosResponseBodyPhotos {
	s.IsVideo = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetTitle(v string) *GetPhotosResponseBodyPhotos {
	s.Title = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetLike(v int64) *GetPhotosResponseBodyPhotos {
	s.Like = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetLocation(v string) *GetPhotosResponseBodyPhotos {
	s.Location = &v
	return s
}

func (s *GetPhotosResponseBodyPhotos) SetId(v int64) *GetPhotosResponseBodyPhotos {
	s.Id = &v
	return s
}

type GetPhotosResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhotosResponse) GoString() string {
	return s.String()
}

func (s *GetPhotosResponse) SetHeaders(v map[string]*string) *GetPhotosResponse {
	s.Headers = v
	return s
}

func (s *GetPhotosResponse) SetBody(v *GetPhotosResponseBody) *GetPhotosResponse {
	s.Body = v
	return s
}

type GetPhotosByMd5sRequest struct {
	State     *string   `json:"State,omitempty" xml:"State,omitempty"`
	StoreName *string   `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string   `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	Md5       []*string `json:"Md5,omitempty" xml:"Md5,omitempty" type:"Repeated"`
}

func (s GetPhotosByMd5sRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhotosByMd5sRequest) GoString() string {
	return s.String()
}

func (s *GetPhotosByMd5sRequest) SetState(v string) *GetPhotosByMd5sRequest {
	s.State = &v
	return s
}

func (s *GetPhotosByMd5sRequest) SetStoreName(v string) *GetPhotosByMd5sRequest {
	s.StoreName = &v
	return s
}

func (s *GetPhotosByMd5sRequest) SetLibraryId(v string) *GetPhotosByMd5sRequest {
	s.LibraryId = &v
	return s
}

func (s *GetPhotosByMd5sRequest) SetMd5(v []*string) *GetPhotosByMd5sRequest {
	s.Md5 = v
	return s
}

type GetPhotosByMd5sResponseBody struct {
	Photos    []*GetPhotosByMd5sResponseBodyPhotos `json:"Photos,omitempty" xml:"Photos,omitempty" type:"Repeated"`
	Action    *string                              `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPhotosByMd5sResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhotosByMd5sResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhotosByMd5sResponseBody) SetPhotos(v []*GetPhotosByMd5sResponseBodyPhotos) *GetPhotosByMd5sResponseBody {
	s.Photos = v
	return s
}

func (s *GetPhotosByMd5sResponseBody) SetAction(v string) *GetPhotosByMd5sResponseBody {
	s.Action = &v
	return s
}

func (s *GetPhotosByMd5sResponseBody) SetMessage(v string) *GetPhotosByMd5sResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhotosByMd5sResponseBody) SetRequestId(v string) *GetPhotosByMd5sResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhotosByMd5sResponseBody) SetCode(v string) *GetPhotosByMd5sResponseBody {
	s.Code = &v
	return s
}

type GetPhotosByMd5sResponseBodyPhotos struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s GetPhotosByMd5sResponseBodyPhotos) String() string {
	return tea.Prettify(s)
}

func (s GetPhotosByMd5sResponseBodyPhotos) GoString() string {
	return s.String()
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetRemark(v string) *GetPhotosByMd5sResponseBodyPhotos {
	s.Remark = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetTakenAt(v int64) *GetPhotosByMd5sResponseBodyPhotos {
	s.TakenAt = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetState(v string) *GetPhotosByMd5sResponseBodyPhotos {
	s.State = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetHeight(v int64) *GetPhotosByMd5sResponseBodyPhotos {
	s.Height = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetShareExpireTime(v int64) *GetPhotosByMd5sResponseBodyPhotos {
	s.ShareExpireTime = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetFileId(v string) *GetPhotosByMd5sResponseBodyPhotos {
	s.FileId = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetIdStr(v string) *GetPhotosByMd5sResponseBodyPhotos {
	s.IdStr = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetCtime(v int64) *GetPhotosByMd5sResponseBodyPhotos {
	s.Ctime = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetMtime(v int64) *GetPhotosByMd5sResponseBodyPhotos {
	s.Mtime = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetWidth(v int64) *GetPhotosByMd5sResponseBodyPhotos {
	s.Width = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetSize(v int64) *GetPhotosByMd5sResponseBodyPhotos {
	s.Size = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetMd5(v string) *GetPhotosByMd5sResponseBodyPhotos {
	s.Md5 = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetTitle(v string) *GetPhotosByMd5sResponseBodyPhotos {
	s.Title = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetIsVideo(v bool) *GetPhotosByMd5sResponseBodyPhotos {
	s.IsVideo = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetId(v int64) *GetPhotosByMd5sResponseBodyPhotos {
	s.Id = &v
	return s
}

func (s *GetPhotosByMd5sResponseBodyPhotos) SetLocation(v string) *GetPhotosByMd5sResponseBodyPhotos {
	s.Location = &v
	return s
}

type GetPhotosByMd5sResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPhotosByMd5sResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPhotosByMd5sResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhotosByMd5sResponse) GoString() string {
	return s.String()
}

func (s *GetPhotosByMd5sResponse) SetHeaders(v map[string]*string) *GetPhotosByMd5sResponse {
	s.Headers = v
	return s
}

func (s *GetPhotosByMd5sResponse) SetBody(v *GetPhotosByMd5sResponseBody) *GetPhotosByMd5sResponse {
	s.Body = v
	return s
}

type GetPhotoStoreRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
}

func (s GetPhotoStoreRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPhotoStoreRequest) GoString() string {
	return s.String()
}

func (s *GetPhotoStoreRequest) SetStoreName(v string) *GetPhotoStoreRequest {
	s.StoreName = &v
	return s
}

type GetPhotoStoreResponseBody struct {
	Action     *string                              `json:"Action,omitempty" xml:"Action,omitempty"`
	Message    *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PhotoStore *GetPhotoStoreResponseBodyPhotoStore `json:"PhotoStore,omitempty" xml:"PhotoStore,omitempty" type:"Struct"`
	Code       *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPhotoStoreResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPhotoStoreResponseBody) GoString() string {
	return s.String()
}

func (s *GetPhotoStoreResponseBody) SetAction(v string) *GetPhotoStoreResponseBody {
	s.Action = &v
	return s
}

func (s *GetPhotoStoreResponseBody) SetMessage(v string) *GetPhotoStoreResponseBody {
	s.Message = &v
	return s
}

func (s *GetPhotoStoreResponseBody) SetRequestId(v string) *GetPhotoStoreResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPhotoStoreResponseBody) SetPhotoStore(v *GetPhotoStoreResponseBodyPhotoStore) *GetPhotoStoreResponseBody {
	s.PhotoStore = v
	return s
}

func (s *GetPhotoStoreResponseBody) SetCode(v string) *GetPhotoStoreResponseBody {
	s.Code = &v
	return s
}

type GetPhotoStoreResponseBodyPhotoStore struct {
	AutoCleanDays     *int32                                        `json:"AutoCleanDays,omitempty" xml:"AutoCleanDays,omitempty"`
	IdStr             *string                                       `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Mtime             *int64                                        `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime             *int64                                        `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	DefaultTrashQuota *int64                                        `json:"DefaultTrashQuota,omitempty" xml:"DefaultTrashQuota,omitempty"`
	Remark            *string                                       `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Buckets           []*GetPhotoStoreResponseBodyPhotoStoreBuckets `json:"Buckets,omitempty" xml:"Buckets,omitempty" type:"Repeated"`
	DefaultQuota      *int64                                        `json:"DefaultQuota,omitempty" xml:"DefaultQuota,omitempty"`
	Name              *string                                       `json:"Name,omitempty" xml:"Name,omitempty"`
	AutoCleanEnabled  *bool                                         `json:"AutoCleanEnabled,omitempty" xml:"AutoCleanEnabled,omitempty"`
	Id                *int64                                        `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetPhotoStoreResponseBodyPhotoStore) String() string {
	return tea.Prettify(s)
}

func (s GetPhotoStoreResponseBodyPhotoStore) GoString() string {
	return s.String()
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetAutoCleanDays(v int32) *GetPhotoStoreResponseBodyPhotoStore {
	s.AutoCleanDays = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetIdStr(v string) *GetPhotoStoreResponseBodyPhotoStore {
	s.IdStr = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetMtime(v int64) *GetPhotoStoreResponseBodyPhotoStore {
	s.Mtime = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetCtime(v int64) *GetPhotoStoreResponseBodyPhotoStore {
	s.Ctime = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetDefaultTrashQuota(v int64) *GetPhotoStoreResponseBodyPhotoStore {
	s.DefaultTrashQuota = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetRemark(v string) *GetPhotoStoreResponseBodyPhotoStore {
	s.Remark = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetBuckets(v []*GetPhotoStoreResponseBodyPhotoStoreBuckets) *GetPhotoStoreResponseBodyPhotoStore {
	s.Buckets = v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetDefaultQuota(v int64) *GetPhotoStoreResponseBodyPhotoStore {
	s.DefaultQuota = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetName(v string) *GetPhotoStoreResponseBodyPhotoStore {
	s.Name = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetAutoCleanEnabled(v bool) *GetPhotoStoreResponseBodyPhotoStore {
	s.AutoCleanEnabled = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStore) SetId(v int64) *GetPhotoStoreResponseBodyPhotoStore {
	s.Id = &v
	return s
}

type GetPhotoStoreResponseBodyPhotoStoreBuckets struct {
	Acl    *string `json:"Acl,omitempty" xml:"Acl,omitempty"`
	State  *string `json:"State,omitempty" xml:"State,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s GetPhotoStoreResponseBodyPhotoStoreBuckets) String() string {
	return tea.Prettify(s)
}

func (s GetPhotoStoreResponseBodyPhotoStoreBuckets) GoString() string {
	return s.String()
}

func (s *GetPhotoStoreResponseBodyPhotoStoreBuckets) SetAcl(v string) *GetPhotoStoreResponseBodyPhotoStoreBuckets {
	s.Acl = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStoreBuckets) SetState(v string) *GetPhotoStoreResponseBodyPhotoStoreBuckets {
	s.State = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStoreBuckets) SetRegion(v string) *GetPhotoStoreResponseBodyPhotoStoreBuckets {
	s.Region = &v
	return s
}

func (s *GetPhotoStoreResponseBodyPhotoStoreBuckets) SetName(v string) *GetPhotoStoreResponseBodyPhotoStoreBuckets {
	s.Name = &v
	return s
}

type GetPhotoStoreResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPhotoStoreResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPhotoStoreResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPhotoStoreResponse) GoString() string {
	return s.String()
}

func (s *GetPhotoStoreResponse) SetHeaders(v map[string]*string) *GetPhotoStoreResponse {
	s.Headers = v
	return s
}

func (s *GetPhotoStoreResponse) SetBody(v *GetPhotoStoreResponseBody) *GetPhotoStoreResponse {
	s.Body = v
	return s
}

type GetPrivateAccessUrlsRequest struct {
	ZoomType  *string `json:"ZoomType,omitempty" xml:"ZoomType,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s GetPrivateAccessUrlsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessUrlsRequest) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessUrlsRequest) SetZoomType(v string) *GetPrivateAccessUrlsRequest {
	s.ZoomType = &v
	return s
}

func (s *GetPrivateAccessUrlsRequest) SetStoreName(v string) *GetPrivateAccessUrlsRequest {
	s.StoreName = &v
	return s
}

func (s *GetPrivateAccessUrlsRequest) SetLibraryId(v string) *GetPrivateAccessUrlsRequest {
	s.LibraryId = &v
	return s
}

func (s *GetPrivateAccessUrlsRequest) SetPhotoId(v []*int) *GetPrivateAccessUrlsRequest {
	s.PhotoId = v
	return s
}

type GetPrivateAccessUrlsResponseBody struct {
	Action    *string                                    `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*GetPrivateAccessUrlsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPrivateAccessUrlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessUrlsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessUrlsResponseBody) SetAction(v string) *GetPrivateAccessUrlsResponseBody {
	s.Action = &v
	return s
}

func (s *GetPrivateAccessUrlsResponseBody) SetMessage(v string) *GetPrivateAccessUrlsResponseBody {
	s.Message = &v
	return s
}

func (s *GetPrivateAccessUrlsResponseBody) SetRequestId(v string) *GetPrivateAccessUrlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPrivateAccessUrlsResponseBody) SetResults(v []*GetPrivateAccessUrlsResponseBodyResults) *GetPrivateAccessUrlsResponseBody {
	s.Results = v
	return s
}

func (s *GetPrivateAccessUrlsResponseBody) SetCode(v string) *GetPrivateAccessUrlsResponseBody {
	s.Code = &v
	return s
}

type GetPrivateAccessUrlsResponseBodyResults struct {
	PhotoIdStr *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PhotoId    *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	AccessUrl  *string `json:"AccessUrl,omitempty" xml:"AccessUrl,omitempty"`
}

func (s GetPrivateAccessUrlsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessUrlsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessUrlsResponseBodyResults) SetPhotoIdStr(v string) *GetPrivateAccessUrlsResponseBodyResults {
	s.PhotoIdStr = &v
	return s
}

func (s *GetPrivateAccessUrlsResponseBodyResults) SetCode(v string) *GetPrivateAccessUrlsResponseBodyResults {
	s.Code = &v
	return s
}

func (s *GetPrivateAccessUrlsResponseBodyResults) SetMessage(v string) *GetPrivateAccessUrlsResponseBodyResults {
	s.Message = &v
	return s
}

func (s *GetPrivateAccessUrlsResponseBodyResults) SetPhotoId(v int64) *GetPrivateAccessUrlsResponseBodyResults {
	s.PhotoId = &v
	return s
}

func (s *GetPrivateAccessUrlsResponseBodyResults) SetAccessUrl(v string) *GetPrivateAccessUrlsResponseBodyResults {
	s.AccessUrl = &v
	return s
}

type GetPrivateAccessUrlsResponse struct {
	Headers map[string]*string                `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPrivateAccessUrlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPrivateAccessUrlsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPrivateAccessUrlsResponse) GoString() string {
	return s.String()
}

func (s *GetPrivateAccessUrlsResponse) SetHeaders(v map[string]*string) *GetPrivateAccessUrlsResponse {
	s.Headers = v
	return s
}

func (s *GetPrivateAccessUrlsResponse) SetBody(v *GetPrivateAccessUrlsResponseBody) *GetPrivateAccessUrlsResponse {
	s.Body = v
	return s
}

type GetPublicAccessUrlsRequest struct {
	ZoomType   *string `json:"ZoomType,omitempty" xml:"ZoomType,omitempty"`
	DomainType *string `json:"DomainType,omitempty" xml:"DomainType,omitempty"`
	StoreName  *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId  *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId    []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s GetPublicAccessUrlsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetPublicAccessUrlsRequest) GoString() string {
	return s.String()
}

func (s *GetPublicAccessUrlsRequest) SetZoomType(v string) *GetPublicAccessUrlsRequest {
	s.ZoomType = &v
	return s
}

func (s *GetPublicAccessUrlsRequest) SetDomainType(v string) *GetPublicAccessUrlsRequest {
	s.DomainType = &v
	return s
}

func (s *GetPublicAccessUrlsRequest) SetStoreName(v string) *GetPublicAccessUrlsRequest {
	s.StoreName = &v
	return s
}

func (s *GetPublicAccessUrlsRequest) SetLibraryId(v string) *GetPublicAccessUrlsRequest {
	s.LibraryId = &v
	return s
}

func (s *GetPublicAccessUrlsRequest) SetPhotoId(v []*int) *GetPublicAccessUrlsRequest {
	s.PhotoId = v
	return s
}

type GetPublicAccessUrlsResponseBody struct {
	Action    *string                                   `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*GetPublicAccessUrlsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetPublicAccessUrlsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetPublicAccessUrlsResponseBody) GoString() string {
	return s.String()
}

func (s *GetPublicAccessUrlsResponseBody) SetAction(v string) *GetPublicAccessUrlsResponseBody {
	s.Action = &v
	return s
}

func (s *GetPublicAccessUrlsResponseBody) SetMessage(v string) *GetPublicAccessUrlsResponseBody {
	s.Message = &v
	return s
}

func (s *GetPublicAccessUrlsResponseBody) SetRequestId(v string) *GetPublicAccessUrlsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetPublicAccessUrlsResponseBody) SetResults(v []*GetPublicAccessUrlsResponseBodyResults) *GetPublicAccessUrlsResponseBody {
	s.Results = v
	return s
}

func (s *GetPublicAccessUrlsResponseBody) SetCode(v string) *GetPublicAccessUrlsResponseBody {
	s.Code = &v
	return s
}

type GetPublicAccessUrlsResponseBodyResults struct {
	PhotoIdStr *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	Code       *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message    *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PhotoId    *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	AccessUrl  *string `json:"AccessUrl,omitempty" xml:"AccessUrl,omitempty"`
}

func (s GetPublicAccessUrlsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetPublicAccessUrlsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetPublicAccessUrlsResponseBodyResults) SetPhotoIdStr(v string) *GetPublicAccessUrlsResponseBodyResults {
	s.PhotoIdStr = &v
	return s
}

func (s *GetPublicAccessUrlsResponseBodyResults) SetCode(v string) *GetPublicAccessUrlsResponseBodyResults {
	s.Code = &v
	return s
}

func (s *GetPublicAccessUrlsResponseBodyResults) SetMessage(v string) *GetPublicAccessUrlsResponseBodyResults {
	s.Message = &v
	return s
}

func (s *GetPublicAccessUrlsResponseBodyResults) SetPhotoId(v int64) *GetPublicAccessUrlsResponseBodyResults {
	s.PhotoId = &v
	return s
}

func (s *GetPublicAccessUrlsResponseBodyResults) SetAccessUrl(v string) *GetPublicAccessUrlsResponseBodyResults {
	s.AccessUrl = &v
	return s
}

type GetPublicAccessUrlsResponse struct {
	Headers map[string]*string               `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetPublicAccessUrlsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetPublicAccessUrlsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetPublicAccessUrlsResponse) GoString() string {
	return s.String()
}

func (s *GetPublicAccessUrlsResponse) SetHeaders(v map[string]*string) *GetPublicAccessUrlsResponse {
	s.Headers = v
	return s
}

func (s *GetPublicAccessUrlsResponse) SetBody(v *GetPublicAccessUrlsResponseBody) *GetPublicAccessUrlsResponse {
	s.Body = v
	return s
}

type GetQuotaRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s GetQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaRequest) GoString() string {
	return s.String()
}

func (s *GetQuotaRequest) SetStoreName(v string) *GetQuotaRequest {
	s.StoreName = &v
	return s
}

func (s *GetQuotaRequest) SetLibraryId(v string) *GetQuotaRequest {
	s.LibraryId = &v
	return s
}

type GetQuotaResponseBody struct {
	Action    *string                    `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                    `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Quota     *GetQuotaResponseBodyQuota `json:"Quota,omitempty" xml:"Quota,omitempty" type:"Struct"`
	Code      *string                    `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBody) SetAction(v string) *GetQuotaResponseBody {
	s.Action = &v
	return s
}

func (s *GetQuotaResponseBody) SetMessage(v string) *GetQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *GetQuotaResponseBody) SetRequestId(v string) *GetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetQuotaResponseBody) SetQuota(v *GetQuotaResponseBodyQuota) *GetQuotaResponseBody {
	s.Quota = v
	return s
}

func (s *GetQuotaResponseBody) SetCode(v string) *GetQuotaResponseBody {
	s.Code = &v
	return s
}

type GetQuotaResponseBodyQuota struct {
	PhotosCount *int32 `json:"PhotosCount,omitempty" xml:"PhotosCount,omitempty"`
	VideosCount *int32 `json:"VideosCount,omitempty" xml:"VideosCount,omitempty"`
	FacesCount  *int32 `json:"FacesCount,omitempty" xml:"FacesCount,omitempty"`
	UsedQuota   *int64 `json:"UsedQuota,omitempty" xml:"UsedQuota,omitempty"`
	TotalQuota  *int64 `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
}

func (s GetQuotaResponseBodyQuota) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponseBodyQuota) GoString() string {
	return s.String()
}

func (s *GetQuotaResponseBodyQuota) SetPhotosCount(v int32) *GetQuotaResponseBodyQuota {
	s.PhotosCount = &v
	return s
}

func (s *GetQuotaResponseBodyQuota) SetVideosCount(v int32) *GetQuotaResponseBodyQuota {
	s.VideosCount = &v
	return s
}

func (s *GetQuotaResponseBodyQuota) SetFacesCount(v int32) *GetQuotaResponseBodyQuota {
	s.FacesCount = &v
	return s
}

func (s *GetQuotaResponseBodyQuota) SetUsedQuota(v int64) *GetQuotaResponseBodyQuota {
	s.UsedQuota = &v
	return s
}

func (s *GetQuotaResponseBodyQuota) SetTotalQuota(v int64) *GetQuotaResponseBodyQuota {
	s.TotalQuota = &v
	return s
}

type GetQuotaResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s GetQuotaResponse) GoString() string {
	return s.String()
}

func (s *GetQuotaResponse) SetHeaders(v map[string]*string) *GetQuotaResponse {
	s.Headers = v
	return s
}

func (s *GetQuotaResponse) SetBody(v *GetQuotaResponseBody) *GetQuotaResponse {
	s.Body = v
	return s
}

type GetSimilarPhotosRequest struct {
	PhotoId   *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s GetSimilarPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarPhotosRequest) GoString() string {
	return s.String()
}

func (s *GetSimilarPhotosRequest) SetPhotoId(v int64) *GetSimilarPhotosRequest {
	s.PhotoId = &v
	return s
}

func (s *GetSimilarPhotosRequest) SetStoreName(v string) *GetSimilarPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *GetSimilarPhotosRequest) SetLibraryId(v string) *GetSimilarPhotosRequest {
	s.LibraryId = &v
	return s
}

type GetSimilarPhotosResponseBody struct {
	Photos    []*GetSimilarPhotosResponseBodyPhotos `json:"Photos,omitempty" xml:"Photos,omitempty" type:"Repeated"`
	Action    *string                               `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetSimilarPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *GetSimilarPhotosResponseBody) SetPhotos(v []*GetSimilarPhotosResponseBodyPhotos) *GetSimilarPhotosResponseBody {
	s.Photos = v
	return s
}

func (s *GetSimilarPhotosResponseBody) SetAction(v string) *GetSimilarPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *GetSimilarPhotosResponseBody) SetMessage(v string) *GetSimilarPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *GetSimilarPhotosResponseBody) SetRequestId(v string) *GetSimilarPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSimilarPhotosResponseBody) SetCode(v string) *GetSimilarPhotosResponseBody {
	s.Code = &v
	return s
}

type GetSimilarPhotosResponseBodyPhotos struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	InactiveTime    *int64  `json:"InactiveTime,omitempty" xml:"InactiveTime,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Like            *int64  `json:"Like,omitempty" xml:"Like,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s GetSimilarPhotosResponseBodyPhotos) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarPhotosResponseBodyPhotos) GoString() string {
	return s.String()
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetRemark(v string) *GetSimilarPhotosResponseBodyPhotos {
	s.Remark = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetTakenAt(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.TakenAt = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetState(v string) *GetSimilarPhotosResponseBodyPhotos {
	s.State = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetHeight(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.Height = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetShareExpireTime(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.ShareExpireTime = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetFileId(v string) *GetSimilarPhotosResponseBodyPhotos {
	s.FileId = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetIdStr(v string) *GetSimilarPhotosResponseBodyPhotos {
	s.IdStr = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetCtime(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.Ctime = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetMtime(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.Mtime = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetSize(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.Size = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetWidth(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.Width = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetInactiveTime(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.InactiveTime = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetMd5(v string) *GetSimilarPhotosResponseBodyPhotos {
	s.Md5 = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetIsVideo(v bool) *GetSimilarPhotosResponseBodyPhotos {
	s.IsVideo = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetTitle(v string) *GetSimilarPhotosResponseBodyPhotos {
	s.Title = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetLike(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.Like = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetLocation(v string) *GetSimilarPhotosResponseBodyPhotos {
	s.Location = &v
	return s
}

func (s *GetSimilarPhotosResponseBodyPhotos) SetId(v int64) *GetSimilarPhotosResponseBodyPhotos {
	s.Id = &v
	return s
}

type GetSimilarPhotosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetSimilarPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetSimilarPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s GetSimilarPhotosResponse) GoString() string {
	return s.String()
}

func (s *GetSimilarPhotosResponse) SetHeaders(v map[string]*string) *GetSimilarPhotosResponse {
	s.Headers = v
	return s
}

func (s *GetSimilarPhotosResponse) SetBody(v *GetSimilarPhotosResponseBody) *GetSimilarPhotosResponse {
	s.Body = v
	return s
}

type GetThumbnailRequest struct {
	PhotoId   *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	ZoomType  *string `json:"ZoomType,omitempty" xml:"ZoomType,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s GetThumbnailRequest) String() string {
	return tea.Prettify(s)
}

func (s GetThumbnailRequest) GoString() string {
	return s.String()
}

func (s *GetThumbnailRequest) SetPhotoId(v int64) *GetThumbnailRequest {
	s.PhotoId = &v
	return s
}

func (s *GetThumbnailRequest) SetZoomType(v string) *GetThumbnailRequest {
	s.ZoomType = &v
	return s
}

func (s *GetThumbnailRequest) SetStoreName(v string) *GetThumbnailRequest {
	s.StoreName = &v
	return s
}

func (s *GetThumbnailRequest) SetLibraryId(v string) *GetThumbnailRequest {
	s.LibraryId = &v
	return s
}

type GetThumbnailResponseBody struct {
	Action       *string `json:"Action,omitempty" xml:"Action,omitempty"`
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId    *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetThumbnailResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetThumbnailResponseBody) GoString() string {
	return s.String()
}

func (s *GetThumbnailResponseBody) SetAction(v string) *GetThumbnailResponseBody {
	s.Action = &v
	return s
}

func (s *GetThumbnailResponseBody) SetThumbnailUrl(v string) *GetThumbnailResponseBody {
	s.ThumbnailUrl = &v
	return s
}

func (s *GetThumbnailResponseBody) SetMessage(v string) *GetThumbnailResponseBody {
	s.Message = &v
	return s
}

func (s *GetThumbnailResponseBody) SetRequestId(v string) *GetThumbnailResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetThumbnailResponseBody) SetCode(v string) *GetThumbnailResponseBody {
	s.Code = &v
	return s
}

type GetThumbnailResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetThumbnailResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetThumbnailResponse) String() string {
	return tea.Prettify(s)
}

func (s GetThumbnailResponse) GoString() string {
	return s.String()
}

func (s *GetThumbnailResponse) SetHeaders(v map[string]*string) *GetThumbnailResponse {
	s.Headers = v
	return s
}

func (s *GetThumbnailResponse) SetBody(v *GetThumbnailResponseBody) *GetThumbnailResponse {
	s.Body = v
	return s
}

type GetThumbnailsRequest struct {
	ZoomType  *string `json:"ZoomType,omitempty" xml:"ZoomType,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s GetThumbnailsRequest) String() string {
	return tea.Prettify(s)
}

func (s GetThumbnailsRequest) GoString() string {
	return s.String()
}

func (s *GetThumbnailsRequest) SetZoomType(v string) *GetThumbnailsRequest {
	s.ZoomType = &v
	return s
}

func (s *GetThumbnailsRequest) SetStoreName(v string) *GetThumbnailsRequest {
	s.StoreName = &v
	return s
}

func (s *GetThumbnailsRequest) SetLibraryId(v string) *GetThumbnailsRequest {
	s.LibraryId = &v
	return s
}

func (s *GetThumbnailsRequest) SetPhotoId(v []*int) *GetThumbnailsRequest {
	s.PhotoId = v
	return s
}

type GetThumbnailsResponseBody struct {
	Action    *string                           `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *GetThumbnailsResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	Code      *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetThumbnailsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetThumbnailsResponseBody) GoString() string {
	return s.String()
}

func (s *GetThumbnailsResponseBody) SetAction(v string) *GetThumbnailsResponseBody {
	s.Action = &v
	return s
}

func (s *GetThumbnailsResponseBody) SetMessage(v string) *GetThumbnailsResponseBody {
	s.Message = &v
	return s
}

func (s *GetThumbnailsResponseBody) SetRequestId(v string) *GetThumbnailsResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetThumbnailsResponseBody) SetResults(v *GetThumbnailsResponseBodyResults) *GetThumbnailsResponseBody {
	s.Results = v
	return s
}

func (s *GetThumbnailsResponseBody) SetCode(v string) *GetThumbnailsResponseBody {
	s.Code = &v
	return s
}

type GetThumbnailsResponseBodyResults struct {
	Result []*GetThumbnailsResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s GetThumbnailsResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s GetThumbnailsResponseBodyResults) GoString() string {
	return s.String()
}

func (s *GetThumbnailsResponseBodyResults) SetResult(v []*GetThumbnailsResponseBodyResultsResult) *GetThumbnailsResponseBodyResults {
	s.Result = v
	return s
}

type GetThumbnailsResponseBodyResultsResult struct {
	ThumbnailUrl *string `json:"ThumbnailUrl,omitempty" xml:"ThumbnailUrl,omitempty"`
	PhotoIdStr   *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	Code         *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message      *string `json:"Message,omitempty" xml:"Message,omitempty"`
	PhotoId      *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
}

func (s GetThumbnailsResponseBodyResultsResult) String() string {
	return tea.Prettify(s)
}

func (s GetThumbnailsResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *GetThumbnailsResponseBodyResultsResult) SetThumbnailUrl(v string) *GetThumbnailsResponseBodyResultsResult {
	s.ThumbnailUrl = &v
	return s
}

func (s *GetThumbnailsResponseBodyResultsResult) SetPhotoIdStr(v string) *GetThumbnailsResponseBodyResultsResult {
	s.PhotoIdStr = &v
	return s
}

func (s *GetThumbnailsResponseBodyResultsResult) SetCode(v string) *GetThumbnailsResponseBodyResultsResult {
	s.Code = &v
	return s
}

func (s *GetThumbnailsResponseBodyResultsResult) SetMessage(v string) *GetThumbnailsResponseBodyResultsResult {
	s.Message = &v
	return s
}

func (s *GetThumbnailsResponseBodyResultsResult) SetPhotoId(v int64) *GetThumbnailsResponseBodyResultsResult {
	s.PhotoId = &v
	return s
}

type GetThumbnailsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetThumbnailsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetThumbnailsResponse) String() string {
	return tea.Prettify(s)
}

func (s GetThumbnailsResponse) GoString() string {
	return s.String()
}

func (s *GetThumbnailsResponse) SetHeaders(v map[string]*string) *GetThumbnailsResponse {
	s.Headers = v
	return s
}

func (s *GetThumbnailsResponse) SetBody(v *GetThumbnailsResponseBody) *GetThumbnailsResponse {
	s.Body = v
	return s
}

type GetVideoCoverRequest struct {
	PhotoId   *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	ZoomType  *string `json:"ZoomType,omitempty" xml:"ZoomType,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s GetVideoCoverRequest) String() string {
	return tea.Prettify(s)
}

func (s GetVideoCoverRequest) GoString() string {
	return s.String()
}

func (s *GetVideoCoverRequest) SetPhotoId(v int64) *GetVideoCoverRequest {
	s.PhotoId = &v
	return s
}

func (s *GetVideoCoverRequest) SetZoomType(v string) *GetVideoCoverRequest {
	s.ZoomType = &v
	return s
}

func (s *GetVideoCoverRequest) SetStoreName(v string) *GetVideoCoverRequest {
	s.StoreName = &v
	return s
}

func (s *GetVideoCoverRequest) SetLibraryId(v string) *GetVideoCoverRequest {
	s.LibraryId = &v
	return s
}

type GetVideoCoverResponseBody struct {
	Action        *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message       *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId     *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	VideoCoverUrl *string `json:"VideoCoverUrl,omitempty" xml:"VideoCoverUrl,omitempty"`
	Code          *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s GetVideoCoverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s GetVideoCoverResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoCoverResponseBody) SetAction(v string) *GetVideoCoverResponseBody {
	s.Action = &v
	return s
}

func (s *GetVideoCoverResponseBody) SetMessage(v string) *GetVideoCoverResponseBody {
	s.Message = &v
	return s
}

func (s *GetVideoCoverResponseBody) SetRequestId(v string) *GetVideoCoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoCoverResponseBody) SetVideoCoverUrl(v string) *GetVideoCoverResponseBody {
	s.VideoCoverUrl = &v
	return s
}

func (s *GetVideoCoverResponseBody) SetCode(v string) *GetVideoCoverResponseBody {
	s.Code = &v
	return s
}

type GetVideoCoverResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *GetVideoCoverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s GetVideoCoverResponse) String() string {
	return tea.Prettify(s)
}

func (s GetVideoCoverResponse) GoString() string {
	return s.String()
}

func (s *GetVideoCoverResponse) SetHeaders(v map[string]*string) *GetVideoCoverResponse {
	s.Headers = v
	return s
}

func (s *GetVideoCoverResponse) SetBody(v *GetVideoCoverResponseBody) *GetVideoCoverResponse {
	s.Body = v
	return s
}

type InactivatePhotosRequest struct {
	StoreName    *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId    *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	InactiveTime *int64  `json:"InactiveTime,omitempty" xml:"InactiveTime,omitempty"`
	PhotoId      []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s InactivatePhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s InactivatePhotosRequest) GoString() string {
	return s.String()
}

func (s *InactivatePhotosRequest) SetStoreName(v string) *InactivatePhotosRequest {
	s.StoreName = &v
	return s
}

func (s *InactivatePhotosRequest) SetLibraryId(v string) *InactivatePhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *InactivatePhotosRequest) SetInactiveTime(v int64) *InactivatePhotosRequest {
	s.InactiveTime = &v
	return s
}

func (s *InactivatePhotosRequest) SetPhotoId(v []*int) *InactivatePhotosRequest {
	s.PhotoId = v
	return s
}

type InactivatePhotosResponseBody struct {
	Action    *string                                `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*InactivatePhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s InactivatePhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s InactivatePhotosResponseBody) GoString() string {
	return s.String()
}

func (s *InactivatePhotosResponseBody) SetAction(v string) *InactivatePhotosResponseBody {
	s.Action = &v
	return s
}

func (s *InactivatePhotosResponseBody) SetMessage(v string) *InactivatePhotosResponseBody {
	s.Message = &v
	return s
}

func (s *InactivatePhotosResponseBody) SetRequestId(v string) *InactivatePhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *InactivatePhotosResponseBody) SetResults(v []*InactivatePhotosResponseBodyResults) *InactivatePhotosResponseBody {
	s.Results = v
	return s
}

func (s *InactivatePhotosResponseBody) SetCode(v string) *InactivatePhotosResponseBody {
	s.Code = &v
	return s
}

type InactivatePhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s InactivatePhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s InactivatePhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *InactivatePhotosResponseBodyResults) SetIdStr(v string) *InactivatePhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *InactivatePhotosResponseBodyResults) SetCode(v string) *InactivatePhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *InactivatePhotosResponseBodyResults) SetMessage(v string) *InactivatePhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *InactivatePhotosResponseBodyResults) SetId(v int64) *InactivatePhotosResponseBodyResults {
	s.Id = &v
	return s
}

type InactivatePhotosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *InactivatePhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s InactivatePhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s InactivatePhotosResponse) GoString() string {
	return s.String()
}

func (s *InactivatePhotosResponse) SetHeaders(v map[string]*string) *InactivatePhotosResponse {
	s.Headers = v
	return s
}

func (s *InactivatePhotosResponse) SetBody(v *InactivatePhotosResponseBody) *InactivatePhotosResponse {
	s.Body = v
	return s
}

type LikePhotoRequest struct {
	PhotoId   *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s LikePhotoRequest) String() string {
	return tea.Prettify(s)
}

func (s LikePhotoRequest) GoString() string {
	return s.String()
}

func (s *LikePhotoRequest) SetPhotoId(v int64) *LikePhotoRequest {
	s.PhotoId = &v
	return s
}

func (s *LikePhotoRequest) SetStoreName(v string) *LikePhotoRequest {
	s.StoreName = &v
	return s
}

func (s *LikePhotoRequest) SetLibraryId(v string) *LikePhotoRequest {
	s.LibraryId = &v
	return s
}

type LikePhotoResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s LikePhotoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s LikePhotoResponseBody) GoString() string {
	return s.String()
}

func (s *LikePhotoResponseBody) SetAction(v string) *LikePhotoResponseBody {
	s.Action = &v
	return s
}

func (s *LikePhotoResponseBody) SetMessage(v string) *LikePhotoResponseBody {
	s.Message = &v
	return s
}

func (s *LikePhotoResponseBody) SetRequestId(v string) *LikePhotoResponseBody {
	s.RequestId = &v
	return s
}

func (s *LikePhotoResponseBody) SetCode(v string) *LikePhotoResponseBody {
	s.Code = &v
	return s
}

type LikePhotoResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *LikePhotoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s LikePhotoResponse) String() string {
	return tea.Prettify(s)
}

func (s LikePhotoResponse) GoString() string {
	return s.String()
}

func (s *LikePhotoResponse) SetHeaders(v map[string]*string) *LikePhotoResponse {
	s.Headers = v
	return s
}

func (s *LikePhotoResponse) SetBody(v *LikePhotoResponseBody) *LikePhotoResponse {
	s.Body = v
	return s
}

type ListAlbumPhotosRequest struct {
	AlbumId   *int64  `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Cursor    *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListAlbumPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumPhotosRequest) GoString() string {
	return s.String()
}

func (s *ListAlbumPhotosRequest) SetAlbumId(v int64) *ListAlbumPhotosRequest {
	s.AlbumId = &v
	return s
}

func (s *ListAlbumPhotosRequest) SetDirection(v string) *ListAlbumPhotosRequest {
	s.Direction = &v
	return s
}

func (s *ListAlbumPhotosRequest) SetSize(v int32) *ListAlbumPhotosRequest {
	s.Size = &v
	return s
}

func (s *ListAlbumPhotosRequest) SetCursor(v string) *ListAlbumPhotosRequest {
	s.Cursor = &v
	return s
}

func (s *ListAlbumPhotosRequest) SetState(v string) *ListAlbumPhotosRequest {
	s.State = &v
	return s
}

func (s *ListAlbumPhotosRequest) SetStoreName(v string) *ListAlbumPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *ListAlbumPhotosRequest) SetLibraryId(v string) *ListAlbumPhotosRequest {
	s.LibraryId = &v
	return s
}

type ListAlbumPhotosResponseBody struct {
	Action     *string                               `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                                `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextCursor *string                               `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	Results    []*ListAlbumPhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListAlbumPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlbumPhotosResponseBody) SetAction(v string) *ListAlbumPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *ListAlbumPhotosResponseBody) SetTotalCount(v int32) *ListAlbumPhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlbumPhotosResponseBody) SetNextCursor(v string) *ListAlbumPhotosResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListAlbumPhotosResponseBody) SetRequestId(v string) *ListAlbumPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlbumPhotosResponseBody) SetMessage(v string) *ListAlbumPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlbumPhotosResponseBody) SetResults(v []*ListAlbumPhotosResponseBodyResults) *ListAlbumPhotosResponseBody {
	s.Results = v
	return s
}

func (s *ListAlbumPhotosResponseBody) SetCode(v string) *ListAlbumPhotosResponseBody {
	s.Code = &v
	return s
}

type ListAlbumPhotosResponseBodyResults struct {
	PhotoIdStr *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	Mtime      *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	PhotoId    *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
}

func (s ListAlbumPhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumPhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListAlbumPhotosResponseBodyResults) SetPhotoIdStr(v string) *ListAlbumPhotosResponseBodyResults {
	s.PhotoIdStr = &v
	return s
}

func (s *ListAlbumPhotosResponseBodyResults) SetMtime(v int64) *ListAlbumPhotosResponseBodyResults {
	s.Mtime = &v
	return s
}

func (s *ListAlbumPhotosResponseBodyResults) SetState(v string) *ListAlbumPhotosResponseBodyResults {
	s.State = &v
	return s
}

func (s *ListAlbumPhotosResponseBodyResults) SetPhotoId(v int64) *ListAlbumPhotosResponseBodyResults {
	s.PhotoId = &v
	return s
}

type ListAlbumPhotosResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlbumPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlbumPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumPhotosResponse) GoString() string {
	return s.String()
}

func (s *ListAlbumPhotosResponse) SetHeaders(v map[string]*string) *ListAlbumPhotosResponse {
	s.Headers = v
	return s
}

func (s *ListAlbumPhotosResponse) SetBody(v *ListAlbumPhotosResponseBody) *ListAlbumPhotosResponse {
	s.Body = v
	return s
}

type ListAlbumsRequest struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Cursor    *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListAlbumsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumsRequest) GoString() string {
	return s.String()
}

func (s *ListAlbumsRequest) SetDirection(v string) *ListAlbumsRequest {
	s.Direction = &v
	return s
}

func (s *ListAlbumsRequest) SetSize(v int32) *ListAlbumsRequest {
	s.Size = &v
	return s
}

func (s *ListAlbumsRequest) SetCursor(v string) *ListAlbumsRequest {
	s.Cursor = &v
	return s
}

func (s *ListAlbumsRequest) SetState(v string) *ListAlbumsRequest {
	s.State = &v
	return s
}

func (s *ListAlbumsRequest) SetStoreName(v string) *ListAlbumsRequest {
	s.StoreName = &v
	return s
}

func (s *ListAlbumsRequest) SetLibraryId(v string) *ListAlbumsRequest {
	s.LibraryId = &v
	return s
}

type ListAlbumsResponseBody struct {
	Action     *string                         `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextCursor *string                         `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Code       *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
	Albums     []*ListAlbumsResponseBodyAlbums `json:"Albums,omitempty" xml:"Albums,omitempty" type:"Repeated"`
}

func (s ListAlbumsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAlbumsResponseBody) SetAction(v string) *ListAlbumsResponseBody {
	s.Action = &v
	return s
}

func (s *ListAlbumsResponseBody) SetTotalCount(v int32) *ListAlbumsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAlbumsResponseBody) SetNextCursor(v string) *ListAlbumsResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListAlbumsResponseBody) SetRequestId(v string) *ListAlbumsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAlbumsResponseBody) SetMessage(v string) *ListAlbumsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAlbumsResponseBody) SetCode(v string) *ListAlbumsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAlbumsResponseBody) SetAlbums(v []*ListAlbumsResponseBodyAlbums) *ListAlbumsResponseBody {
	s.Albums = v
	return s
}

type ListAlbumsResponseBodyAlbums struct {
	IdStr       *string                            `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	PhotosCount *int64                             `json:"PhotosCount,omitempty" xml:"PhotosCount,omitempty"`
	Cover       *ListAlbumsResponseBodyAlbumsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Mtime       *int64                             `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime       *int64                             `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Remark      *string                            `json:"Remark,omitempty" xml:"Remark,omitempty"`
	State       *string                            `json:"State,omitempty" xml:"State,omitempty"`
	Name        *string                            `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *int64                             `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListAlbumsResponseBodyAlbums) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumsResponseBodyAlbums) GoString() string {
	return s.String()
}

func (s *ListAlbumsResponseBodyAlbums) SetIdStr(v string) *ListAlbumsResponseBodyAlbums {
	s.IdStr = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbums) SetPhotosCount(v int64) *ListAlbumsResponseBodyAlbums {
	s.PhotosCount = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbums) SetCover(v *ListAlbumsResponseBodyAlbumsCover) *ListAlbumsResponseBodyAlbums {
	s.Cover = v
	return s
}

func (s *ListAlbumsResponseBodyAlbums) SetMtime(v int64) *ListAlbumsResponseBodyAlbums {
	s.Mtime = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbums) SetCtime(v int64) *ListAlbumsResponseBodyAlbums {
	s.Ctime = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbums) SetRemark(v string) *ListAlbumsResponseBodyAlbums {
	s.Remark = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbums) SetState(v string) *ListAlbumsResponseBodyAlbums {
	s.State = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbums) SetName(v string) *ListAlbumsResponseBodyAlbums {
	s.Name = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbums) SetId(v int64) *ListAlbumsResponseBodyAlbums {
	s.Id = &v
	return s
}

type ListAlbumsResponseBodyAlbumsCover struct {
	Remark  *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	State   *string `json:"State,omitempty" xml:"State,omitempty"`
	Height  *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	FileId  *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Mtime   *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime   *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Width   *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Md5     *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListAlbumsResponseBodyAlbumsCover) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumsResponseBodyAlbumsCover) GoString() string {
	return s.String()
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetRemark(v string) *ListAlbumsResponseBodyAlbumsCover {
	s.Remark = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetState(v string) *ListAlbumsResponseBodyAlbumsCover {
	s.State = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetHeight(v int64) *ListAlbumsResponseBodyAlbumsCover {
	s.Height = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetFileId(v string) *ListAlbumsResponseBodyAlbumsCover {
	s.FileId = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetIdStr(v string) *ListAlbumsResponseBodyAlbumsCover {
	s.IdStr = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetMtime(v int64) *ListAlbumsResponseBodyAlbumsCover {
	s.Mtime = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetCtime(v int64) *ListAlbumsResponseBodyAlbumsCover {
	s.Ctime = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetWidth(v int64) *ListAlbumsResponseBodyAlbumsCover {
	s.Width = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetMd5(v string) *ListAlbumsResponseBodyAlbumsCover {
	s.Md5 = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetIsVideo(v bool) *ListAlbumsResponseBodyAlbumsCover {
	s.IsVideo = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetTitle(v string) *ListAlbumsResponseBodyAlbumsCover {
	s.Title = &v
	return s
}

func (s *ListAlbumsResponseBodyAlbumsCover) SetId(v int64) *ListAlbumsResponseBodyAlbumsCover {
	s.Id = &v
	return s
}

type ListAlbumsResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListAlbumsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListAlbumsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListAlbumsResponse) GoString() string {
	return s.String()
}

func (s *ListAlbumsResponse) SetHeaders(v map[string]*string) *ListAlbumsResponse {
	s.Headers = v
	return s
}

func (s *ListAlbumsResponse) SetBody(v *ListAlbumsResponseBody) *ListAlbumsResponse {
	s.Body = v
	return s
}

type ListFacePhotosRequest struct {
	FaceId    *int64  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Cursor    *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListFacePhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFacePhotosRequest) GoString() string {
	return s.String()
}

func (s *ListFacePhotosRequest) SetFaceId(v int64) *ListFacePhotosRequest {
	s.FaceId = &v
	return s
}

func (s *ListFacePhotosRequest) SetDirection(v string) *ListFacePhotosRequest {
	s.Direction = &v
	return s
}

func (s *ListFacePhotosRequest) SetSize(v int32) *ListFacePhotosRequest {
	s.Size = &v
	return s
}

func (s *ListFacePhotosRequest) SetCursor(v string) *ListFacePhotosRequest {
	s.Cursor = &v
	return s
}

func (s *ListFacePhotosRequest) SetState(v string) *ListFacePhotosRequest {
	s.State = &v
	return s
}

func (s *ListFacePhotosRequest) SetStoreName(v string) *ListFacePhotosRequest {
	s.StoreName = &v
	return s
}

func (s *ListFacePhotosRequest) SetLibraryId(v string) *ListFacePhotosRequest {
	s.LibraryId = &v
	return s
}

type ListFacePhotosResponseBody struct {
	Action     *string                              `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                               `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextCursor *string                              `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	RequestId  *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	Results    []*ListFacePhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code       *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListFacePhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFacePhotosResponseBody) GoString() string {
	return s.String()
}

func (s *ListFacePhotosResponseBody) SetAction(v string) *ListFacePhotosResponseBody {
	s.Action = &v
	return s
}

func (s *ListFacePhotosResponseBody) SetTotalCount(v int32) *ListFacePhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFacePhotosResponseBody) SetNextCursor(v string) *ListFacePhotosResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListFacePhotosResponseBody) SetRequestId(v string) *ListFacePhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFacePhotosResponseBody) SetMessage(v string) *ListFacePhotosResponseBody {
	s.Message = &v
	return s
}

func (s *ListFacePhotosResponseBody) SetResults(v []*ListFacePhotosResponseBodyResults) *ListFacePhotosResponseBody {
	s.Results = v
	return s
}

func (s *ListFacePhotosResponseBody) SetCode(v string) *ListFacePhotosResponseBody {
	s.Code = &v
	return s
}

type ListFacePhotosResponseBodyResults struct {
	PhotoIdStr *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	Mtime      *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	PhotoId    *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
}

func (s ListFacePhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListFacePhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListFacePhotosResponseBodyResults) SetPhotoIdStr(v string) *ListFacePhotosResponseBodyResults {
	s.PhotoIdStr = &v
	return s
}

func (s *ListFacePhotosResponseBodyResults) SetMtime(v int64) *ListFacePhotosResponseBodyResults {
	s.Mtime = &v
	return s
}

func (s *ListFacePhotosResponseBodyResults) SetState(v string) *ListFacePhotosResponseBodyResults {
	s.State = &v
	return s
}

func (s *ListFacePhotosResponseBodyResults) SetPhotoId(v int64) *ListFacePhotosResponseBodyResults {
	s.PhotoId = &v
	return s
}

type ListFacePhotosResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFacePhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFacePhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFacePhotosResponse) GoString() string {
	return s.String()
}

func (s *ListFacePhotosResponse) SetHeaders(v map[string]*string) *ListFacePhotosResponse {
	s.Headers = v
	return s
}

func (s *ListFacePhotosResponse) SetBody(v *ListFacePhotosResponseBody) *ListFacePhotosResponse {
	s.Body = v
	return s
}

type ListFacesRequest struct {
	Direction   *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Size        *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Cursor      *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	State       *string `json:"State,omitempty" xml:"State,omitempty"`
	StoreName   *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId   *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	HasFaceName *string `json:"HasFaceName,omitempty" xml:"HasFaceName,omitempty"`
}

func (s ListFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListFacesRequest) GoString() string {
	return s.String()
}

func (s *ListFacesRequest) SetDirection(v string) *ListFacesRequest {
	s.Direction = &v
	return s
}

func (s *ListFacesRequest) SetSize(v int32) *ListFacesRequest {
	s.Size = &v
	return s
}

func (s *ListFacesRequest) SetCursor(v string) *ListFacesRequest {
	s.Cursor = &v
	return s
}

func (s *ListFacesRequest) SetState(v string) *ListFacesRequest {
	s.State = &v
	return s
}

func (s *ListFacesRequest) SetStoreName(v string) *ListFacesRequest {
	s.StoreName = &v
	return s
}

func (s *ListFacesRequest) SetLibraryId(v string) *ListFacesRequest {
	s.LibraryId = &v
	return s
}

func (s *ListFacesRequest) SetHasFaceName(v string) *ListFacesRequest {
	s.HasFaceName = &v
	return s
}

type ListFacesResponseBody struct {
	Action     *string                       `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                        `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextCursor *string                       `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	RequestId  *string                       `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                       `json:"Message,omitempty" xml:"Message,omitempty"`
	Code       *string                       `json:"Code,omitempty" xml:"Code,omitempty"`
	Faces      []*ListFacesResponseBodyFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
}

func (s ListFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListFacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListFacesResponseBody) SetAction(v string) *ListFacesResponseBody {
	s.Action = &v
	return s
}

func (s *ListFacesResponseBody) SetTotalCount(v int32) *ListFacesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListFacesResponseBody) SetNextCursor(v string) *ListFacesResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListFacesResponseBody) SetRequestId(v string) *ListFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFacesResponseBody) SetMessage(v string) *ListFacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListFacesResponseBody) SetCode(v string) *ListFacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListFacesResponseBody) SetFaces(v []*ListFacesResponseBodyFaces) *ListFacesResponseBody {
	s.Faces = v
	return s
}

type ListFacesResponseBodyFaces struct {
	IdStr       *string                          `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	PhotosCount *int32                           `json:"PhotosCount,omitempty" xml:"PhotosCount,omitempty"`
	IsMe        *bool                            `json:"IsMe,omitempty" xml:"IsMe,omitempty"`
	Cover       *ListFacesResponseBodyFacesCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	Mtime       *int64                           `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime       *int64                           `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	State       *string                          `json:"State,omitempty" xml:"State,omitempty"`
	Axis        []*string                        `json:"Axis,omitempty" xml:"Axis,omitempty" type:"Repeated"`
	Name        *string                          `json:"Name,omitempty" xml:"Name,omitempty"`
	Id          *int64                           `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListFacesResponseBodyFaces) String() string {
	return tea.Prettify(s)
}

func (s ListFacesResponseBodyFaces) GoString() string {
	return s.String()
}

func (s *ListFacesResponseBodyFaces) SetIdStr(v string) *ListFacesResponseBodyFaces {
	s.IdStr = &v
	return s
}

func (s *ListFacesResponseBodyFaces) SetPhotosCount(v int32) *ListFacesResponseBodyFaces {
	s.PhotosCount = &v
	return s
}

func (s *ListFacesResponseBodyFaces) SetIsMe(v bool) *ListFacesResponseBodyFaces {
	s.IsMe = &v
	return s
}

func (s *ListFacesResponseBodyFaces) SetCover(v *ListFacesResponseBodyFacesCover) *ListFacesResponseBodyFaces {
	s.Cover = v
	return s
}

func (s *ListFacesResponseBodyFaces) SetMtime(v int64) *ListFacesResponseBodyFaces {
	s.Mtime = &v
	return s
}

func (s *ListFacesResponseBodyFaces) SetCtime(v int64) *ListFacesResponseBodyFaces {
	s.Ctime = &v
	return s
}

func (s *ListFacesResponseBodyFaces) SetState(v string) *ListFacesResponseBodyFaces {
	s.State = &v
	return s
}

func (s *ListFacesResponseBodyFaces) SetAxis(v []*string) *ListFacesResponseBodyFaces {
	s.Axis = v
	return s
}

func (s *ListFacesResponseBodyFaces) SetName(v string) *ListFacesResponseBodyFaces {
	s.Name = &v
	return s
}

func (s *ListFacesResponseBodyFaces) SetId(v int64) *ListFacesResponseBodyFaces {
	s.Id = &v
	return s
}

type ListFacesResponseBodyFacesCover struct {
	Remark  *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	State   *string `json:"State,omitempty" xml:"State,omitempty"`
	Height  *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	FileId  *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Mtime   *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime   *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Width   *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Md5     *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListFacesResponseBodyFacesCover) String() string {
	return tea.Prettify(s)
}

func (s ListFacesResponseBodyFacesCover) GoString() string {
	return s.String()
}

func (s *ListFacesResponseBodyFacesCover) SetRemark(v string) *ListFacesResponseBodyFacesCover {
	s.Remark = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetState(v string) *ListFacesResponseBodyFacesCover {
	s.State = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetHeight(v int64) *ListFacesResponseBodyFacesCover {
	s.Height = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetFileId(v string) *ListFacesResponseBodyFacesCover {
	s.FileId = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetIdStr(v string) *ListFacesResponseBodyFacesCover {
	s.IdStr = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetMtime(v int64) *ListFacesResponseBodyFacesCover {
	s.Mtime = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetCtime(v int64) *ListFacesResponseBodyFacesCover {
	s.Ctime = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetWidth(v int64) *ListFacesResponseBodyFacesCover {
	s.Width = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetMd5(v string) *ListFacesResponseBodyFacesCover {
	s.Md5 = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetIsVideo(v bool) *ListFacesResponseBodyFacesCover {
	s.IsVideo = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetTitle(v string) *ListFacesResponseBodyFacesCover {
	s.Title = &v
	return s
}

func (s *ListFacesResponseBodyFacesCover) SetId(v int64) *ListFacesResponseBodyFacesCover {
	s.Id = &v
	return s
}

type ListFacesResponse struct {
	Headers map[string]*string     `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListFacesResponse) GoString() string {
	return s.String()
}

func (s *ListFacesResponse) SetHeaders(v map[string]*string) *ListFacesResponse {
	s.Headers = v
	return s
}

func (s *ListFacesResponse) SetBody(v *ListFacesResponseBody) *ListFacesResponse {
	s.Body = v
	return s
}

type ListMomentPhotosRequest struct {
	MomentId  *int64  `json:"MomentId,omitempty" xml:"MomentId,omitempty"`
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Cursor    *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListMomentPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMomentPhotosRequest) GoString() string {
	return s.String()
}

func (s *ListMomentPhotosRequest) SetMomentId(v int64) *ListMomentPhotosRequest {
	s.MomentId = &v
	return s
}

func (s *ListMomentPhotosRequest) SetDirection(v string) *ListMomentPhotosRequest {
	s.Direction = &v
	return s
}

func (s *ListMomentPhotosRequest) SetSize(v int32) *ListMomentPhotosRequest {
	s.Size = &v
	return s
}

func (s *ListMomentPhotosRequest) SetCursor(v string) *ListMomentPhotosRequest {
	s.Cursor = &v
	return s
}

func (s *ListMomentPhotosRequest) SetState(v string) *ListMomentPhotosRequest {
	s.State = &v
	return s
}

func (s *ListMomentPhotosRequest) SetStoreName(v string) *ListMomentPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *ListMomentPhotosRequest) SetLibraryId(v string) *ListMomentPhotosRequest {
	s.LibraryId = &v
	return s
}

type ListMomentPhotosResponseBody struct {
	Action     *string                                `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                                 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextCursor *string                                `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	RequestId  *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	Results    []*ListMomentPhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code       *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListMomentPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMomentPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *ListMomentPhotosResponseBody) SetAction(v string) *ListMomentPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *ListMomentPhotosResponseBody) SetTotalCount(v int32) *ListMomentPhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMomentPhotosResponseBody) SetNextCursor(v string) *ListMomentPhotosResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListMomentPhotosResponseBody) SetRequestId(v string) *ListMomentPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMomentPhotosResponseBody) SetMessage(v string) *ListMomentPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *ListMomentPhotosResponseBody) SetResults(v []*ListMomentPhotosResponseBodyResults) *ListMomentPhotosResponseBody {
	s.Results = v
	return s
}

func (s *ListMomentPhotosResponseBody) SetCode(v string) *ListMomentPhotosResponseBody {
	s.Code = &v
	return s
}

type ListMomentPhotosResponseBodyResults struct {
	PhotoIdStr *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	PhotoId    *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
}

func (s ListMomentPhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListMomentPhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListMomentPhotosResponseBodyResults) SetPhotoIdStr(v string) *ListMomentPhotosResponseBodyResults {
	s.PhotoIdStr = &v
	return s
}

func (s *ListMomentPhotosResponseBodyResults) SetState(v string) *ListMomentPhotosResponseBodyResults {
	s.State = &v
	return s
}

func (s *ListMomentPhotosResponseBodyResults) SetPhotoId(v int64) *ListMomentPhotosResponseBodyResults {
	s.PhotoId = &v
	return s
}

type ListMomentPhotosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMomentPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMomentPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMomentPhotosResponse) GoString() string {
	return s.String()
}

func (s *ListMomentPhotosResponse) SetHeaders(v map[string]*string) *ListMomentPhotosResponse {
	s.Headers = v
	return s
}

func (s *ListMomentPhotosResponse) SetBody(v *ListMomentPhotosResponseBody) *ListMomentPhotosResponse {
	s.Body = v
	return s
}

type ListMomentsRequest struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Cursor    *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListMomentsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListMomentsRequest) GoString() string {
	return s.String()
}

func (s *ListMomentsRequest) SetDirection(v string) *ListMomentsRequest {
	s.Direction = &v
	return s
}

func (s *ListMomentsRequest) SetSize(v int32) *ListMomentsRequest {
	s.Size = &v
	return s
}

func (s *ListMomentsRequest) SetCursor(v string) *ListMomentsRequest {
	s.Cursor = &v
	return s
}

func (s *ListMomentsRequest) SetState(v string) *ListMomentsRequest {
	s.State = &v
	return s
}

func (s *ListMomentsRequest) SetStoreName(v string) *ListMomentsRequest {
	s.StoreName = &v
	return s
}

func (s *ListMomentsRequest) SetLibraryId(v string) *ListMomentsRequest {
	s.LibraryId = &v
	return s
}

type ListMomentsResponseBody struct {
	Action     *string                           `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextCursor *string                           `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	Moments    []*ListMomentsResponseBodyMoments `json:"Moments,omitempty" xml:"Moments,omitempty" type:"Repeated"`
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListMomentsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListMomentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListMomentsResponseBody) SetAction(v string) *ListMomentsResponseBody {
	s.Action = &v
	return s
}

func (s *ListMomentsResponseBody) SetTotalCount(v int32) *ListMomentsResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListMomentsResponseBody) SetNextCursor(v string) *ListMomentsResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListMomentsResponseBody) SetRequestId(v string) *ListMomentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListMomentsResponseBody) SetMessage(v string) *ListMomentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListMomentsResponseBody) SetMoments(v []*ListMomentsResponseBodyMoments) *ListMomentsResponseBody {
	s.Moments = v
	return s
}

func (s *ListMomentsResponseBody) SetCode(v string) *ListMomentsResponseBody {
	s.Code = &v
	return s
}

type ListMomentsResponseBodyMoments struct {
	IdStr        *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	PhotosCount  *int32  `json:"PhotosCount,omitempty" xml:"PhotosCount,omitempty"`
	Mtime        *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime        *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	TakenAt      *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State        *string `json:"State,omitempty" xml:"State,omitempty"`
	LocationName *string `json:"LocationName,omitempty" xml:"LocationName,omitempty"`
	Id           *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListMomentsResponseBodyMoments) String() string {
	return tea.Prettify(s)
}

func (s ListMomentsResponseBodyMoments) GoString() string {
	return s.String()
}

func (s *ListMomentsResponseBodyMoments) SetIdStr(v string) *ListMomentsResponseBodyMoments {
	s.IdStr = &v
	return s
}

func (s *ListMomentsResponseBodyMoments) SetPhotosCount(v int32) *ListMomentsResponseBodyMoments {
	s.PhotosCount = &v
	return s
}

func (s *ListMomentsResponseBodyMoments) SetMtime(v int64) *ListMomentsResponseBodyMoments {
	s.Mtime = &v
	return s
}

func (s *ListMomentsResponseBodyMoments) SetCtime(v int64) *ListMomentsResponseBodyMoments {
	s.Ctime = &v
	return s
}

func (s *ListMomentsResponseBodyMoments) SetTakenAt(v int64) *ListMomentsResponseBodyMoments {
	s.TakenAt = &v
	return s
}

func (s *ListMomentsResponseBodyMoments) SetState(v string) *ListMomentsResponseBodyMoments {
	s.State = &v
	return s
}

func (s *ListMomentsResponseBodyMoments) SetLocationName(v string) *ListMomentsResponseBodyMoments {
	s.LocationName = &v
	return s
}

func (s *ListMomentsResponseBodyMoments) SetId(v int64) *ListMomentsResponseBodyMoments {
	s.Id = &v
	return s
}

type ListMomentsResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListMomentsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListMomentsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListMomentsResponse) GoString() string {
	return s.String()
}

func (s *ListMomentsResponse) SetHeaders(v map[string]*string) *ListMomentsResponse {
	s.Headers = v
	return s
}

func (s *ListMomentsResponse) SetBody(v *ListMomentsResponseBody) *ListMomentsResponse {
	s.Body = v
	return s
}

type ListPhotoFacesRequest struct {
	PhotoId   *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListPhotoFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoFacesRequest) GoString() string {
	return s.String()
}

func (s *ListPhotoFacesRequest) SetPhotoId(v int64) *ListPhotoFacesRequest {
	s.PhotoId = &v
	return s
}

func (s *ListPhotoFacesRequest) SetStoreName(v string) *ListPhotoFacesRequest {
	s.StoreName = &v
	return s
}

func (s *ListPhotoFacesRequest) SetLibraryId(v string) *ListPhotoFacesRequest {
	s.LibraryId = &v
	return s
}

type ListPhotoFacesResponseBody struct {
	Action    *string                            `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                            `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                            `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                            `json:"Code,omitempty" xml:"Code,omitempty"`
	Faces     []*ListPhotoFacesResponseBodyFaces `json:"Faces,omitempty" xml:"Faces,omitempty" type:"Repeated"`
}

func (s ListPhotoFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoFacesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhotoFacesResponseBody) SetAction(v string) *ListPhotoFacesResponseBody {
	s.Action = &v
	return s
}

func (s *ListPhotoFacesResponseBody) SetMessage(v string) *ListPhotoFacesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhotoFacesResponseBody) SetRequestId(v string) *ListPhotoFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhotoFacesResponseBody) SetCode(v string) *ListPhotoFacesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPhotoFacesResponseBody) SetFaces(v []*ListPhotoFacesResponseBodyFaces) *ListPhotoFacesResponseBody {
	s.Faces = v
	return s
}

type ListPhotoFacesResponseBodyFaces struct {
	FaceIdStr *string   `json:"FaceIdStr,omitempty" xml:"FaceIdStr,omitempty"`
	FaceName  *string   `json:"FaceName,omitempty" xml:"FaceName,omitempty"`
	FaceId    *int64    `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	Axis      []*string `json:"Axis,omitempty" xml:"Axis,omitempty" type:"Repeated"`
}

func (s ListPhotoFacesResponseBodyFaces) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoFacesResponseBodyFaces) GoString() string {
	return s.String()
}

func (s *ListPhotoFacesResponseBodyFaces) SetFaceIdStr(v string) *ListPhotoFacesResponseBodyFaces {
	s.FaceIdStr = &v
	return s
}

func (s *ListPhotoFacesResponseBodyFaces) SetFaceName(v string) *ListPhotoFacesResponseBodyFaces {
	s.FaceName = &v
	return s
}

func (s *ListPhotoFacesResponseBodyFaces) SetFaceId(v int64) *ListPhotoFacesResponseBodyFaces {
	s.FaceId = &v
	return s
}

func (s *ListPhotoFacesResponseBodyFaces) SetAxis(v []*string) *ListPhotoFacesResponseBodyFaces {
	s.Axis = v
	return s
}

type ListPhotoFacesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPhotoFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPhotoFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoFacesResponse) GoString() string {
	return s.String()
}

func (s *ListPhotoFacesResponse) SetHeaders(v map[string]*string) *ListPhotoFacesResponse {
	s.Headers = v
	return s
}

func (s *ListPhotoFacesResponse) SetBody(v *ListPhotoFacesResponseBody) *ListPhotoFacesResponse {
	s.Body = v
	return s
}

type ListPhotosRequest struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Cursor    *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPhotosRequest) GoString() string {
	return s.String()
}

func (s *ListPhotosRequest) SetDirection(v string) *ListPhotosRequest {
	s.Direction = &v
	return s
}

func (s *ListPhotosRequest) SetSize(v int32) *ListPhotosRequest {
	s.Size = &v
	return s
}

func (s *ListPhotosRequest) SetCursor(v string) *ListPhotosRequest {
	s.Cursor = &v
	return s
}

func (s *ListPhotosRequest) SetState(v string) *ListPhotosRequest {
	s.State = &v
	return s
}

func (s *ListPhotosRequest) SetStoreName(v string) *ListPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *ListPhotosRequest) SetLibraryId(v string) *ListPhotosRequest {
	s.LibraryId = &v
	return s
}

type ListPhotosResponseBody struct {
	Photos     []*ListPhotosResponseBodyPhotos `json:"Photos,omitempty" xml:"Photos,omitempty" type:"Repeated"`
	Action     *string                         `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                          `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextCursor *string                         `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	RequestId  *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	Code       *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhotosResponseBody) SetPhotos(v []*ListPhotosResponseBodyPhotos) *ListPhotosResponseBody {
	s.Photos = v
	return s
}

func (s *ListPhotosResponseBody) SetAction(v string) *ListPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *ListPhotosResponseBody) SetTotalCount(v int32) *ListPhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListPhotosResponseBody) SetNextCursor(v string) *ListPhotosResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListPhotosResponseBody) SetRequestId(v string) *ListPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhotosResponseBody) SetMessage(v string) *ListPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhotosResponseBody) SetCode(v string) *ListPhotosResponseBody {
	s.Code = &v
	return s
}

type ListPhotosResponseBodyPhotos struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	InactiveTime    *int64  `json:"InactiveTime,omitempty" xml:"InactiveTime,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListPhotosResponseBodyPhotos) String() string {
	return tea.Prettify(s)
}

func (s ListPhotosResponseBodyPhotos) GoString() string {
	return s.String()
}

func (s *ListPhotosResponseBodyPhotos) SetRemark(v string) *ListPhotosResponseBodyPhotos {
	s.Remark = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetTakenAt(v int64) *ListPhotosResponseBodyPhotos {
	s.TakenAt = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetState(v string) *ListPhotosResponseBodyPhotos {
	s.State = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetHeight(v int64) *ListPhotosResponseBodyPhotos {
	s.Height = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetShareExpireTime(v int64) *ListPhotosResponseBodyPhotos {
	s.ShareExpireTime = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetFileId(v string) *ListPhotosResponseBodyPhotos {
	s.FileId = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetIdStr(v string) *ListPhotosResponseBodyPhotos {
	s.IdStr = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetCtime(v int64) *ListPhotosResponseBodyPhotos {
	s.Ctime = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetMtime(v int64) *ListPhotosResponseBodyPhotos {
	s.Mtime = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetSize(v int64) *ListPhotosResponseBodyPhotos {
	s.Size = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetWidth(v int64) *ListPhotosResponseBodyPhotos {
	s.Width = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetInactiveTime(v int64) *ListPhotosResponseBodyPhotos {
	s.InactiveTime = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetMd5(v string) *ListPhotosResponseBodyPhotos {
	s.Md5 = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetIsVideo(v bool) *ListPhotosResponseBodyPhotos {
	s.IsVideo = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetTitle(v string) *ListPhotosResponseBodyPhotos {
	s.Title = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetLocation(v string) *ListPhotosResponseBodyPhotos {
	s.Location = &v
	return s
}

func (s *ListPhotosResponseBodyPhotos) SetId(v int64) *ListPhotosResponseBodyPhotos {
	s.Id = &v
	return s
}

type ListPhotosResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPhotosResponse) GoString() string {
	return s.String()
}

func (s *ListPhotosResponse) SetHeaders(v map[string]*string) *ListPhotosResponse {
	s.Headers = v
	return s
}

func (s *ListPhotosResponse) SetBody(v *ListPhotosResponseBody) *ListPhotosResponse {
	s.Body = v
	return s
}

type ListPhotoStoresResponseBody struct {
	Action      *string                                   `json:"Action,omitempty" xml:"Action,omitempty"`
	Message     *string                                   `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId   *string                                   `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	PhotoStores []*ListPhotoStoresResponseBodyPhotoStores `json:"PhotoStores,omitempty" xml:"PhotoStores,omitempty" type:"Repeated"`
	Code        *string                                   `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListPhotoStoresResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoStoresResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhotoStoresResponseBody) SetAction(v string) *ListPhotoStoresResponseBody {
	s.Action = &v
	return s
}

func (s *ListPhotoStoresResponseBody) SetMessage(v string) *ListPhotoStoresResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhotoStoresResponseBody) SetRequestId(v string) *ListPhotoStoresResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhotoStoresResponseBody) SetPhotoStores(v []*ListPhotoStoresResponseBodyPhotoStores) *ListPhotoStoresResponseBody {
	s.PhotoStores = v
	return s
}

func (s *ListPhotoStoresResponseBody) SetCode(v string) *ListPhotoStoresResponseBody {
	s.Code = &v
	return s
}

type ListPhotoStoresResponseBodyPhotoStores struct {
	AutoCleanDays    *int32                                           `json:"AutoCleanDays,omitempty" xml:"AutoCleanDays,omitempty"`
	IdStr            *string                                          `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Mtime            *int64                                           `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime            *int64                                           `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Remark           *string                                          `json:"Remark,omitempty" xml:"Remark,omitempty"`
	Buckets          []*ListPhotoStoresResponseBodyPhotoStoresBuckets `json:"Buckets,omitempty" xml:"Buckets,omitempty" type:"Repeated"`
	DefaultQuota     *int64                                           `json:"DefaultQuota,omitempty" xml:"DefaultQuota,omitempty"`
	Name             *string                                          `json:"Name,omitempty" xml:"Name,omitempty"`
	AutoCleanEnabled *bool                                            `json:"AutoCleanEnabled,omitempty" xml:"AutoCleanEnabled,omitempty"`
	Id               *int64                                           `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListPhotoStoresResponseBodyPhotoStores) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoStoresResponseBodyPhotoStores) GoString() string {
	return s.String()
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetAutoCleanDays(v int32) *ListPhotoStoresResponseBodyPhotoStores {
	s.AutoCleanDays = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetIdStr(v string) *ListPhotoStoresResponseBodyPhotoStores {
	s.IdStr = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetMtime(v int64) *ListPhotoStoresResponseBodyPhotoStores {
	s.Mtime = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetCtime(v int64) *ListPhotoStoresResponseBodyPhotoStores {
	s.Ctime = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetRemark(v string) *ListPhotoStoresResponseBodyPhotoStores {
	s.Remark = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetBuckets(v []*ListPhotoStoresResponseBodyPhotoStoresBuckets) *ListPhotoStoresResponseBodyPhotoStores {
	s.Buckets = v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetDefaultQuota(v int64) *ListPhotoStoresResponseBodyPhotoStores {
	s.DefaultQuota = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetName(v string) *ListPhotoStoresResponseBodyPhotoStores {
	s.Name = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetAutoCleanEnabled(v bool) *ListPhotoStoresResponseBodyPhotoStores {
	s.AutoCleanEnabled = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStores) SetId(v int64) *ListPhotoStoresResponseBodyPhotoStores {
	s.Id = &v
	return s
}

type ListPhotoStoresResponseBodyPhotoStoresBuckets struct {
	State  *string `json:"State,omitempty" xml:"State,omitempty"`
	Region *string `json:"Region,omitempty" xml:"Region,omitempty"`
	Name   *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s ListPhotoStoresResponseBodyPhotoStoresBuckets) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoStoresResponseBodyPhotoStoresBuckets) GoString() string {
	return s.String()
}

func (s *ListPhotoStoresResponseBodyPhotoStoresBuckets) SetState(v string) *ListPhotoStoresResponseBodyPhotoStoresBuckets {
	s.State = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStoresBuckets) SetRegion(v string) *ListPhotoStoresResponseBodyPhotoStoresBuckets {
	s.Region = &v
	return s
}

func (s *ListPhotoStoresResponseBodyPhotoStoresBuckets) SetName(v string) *ListPhotoStoresResponseBodyPhotoStoresBuckets {
	s.Name = &v
	return s
}

type ListPhotoStoresResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPhotoStoresResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPhotoStoresResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoStoresResponse) GoString() string {
	return s.String()
}

func (s *ListPhotoStoresResponse) SetHeaders(v map[string]*string) *ListPhotoStoresResponse {
	s.Headers = v
	return s
}

func (s *ListPhotoStoresResponse) SetBody(v *ListPhotoStoresResponseBody) *ListPhotoStoresResponse {
	s.Body = v
	return s
}

type ListPhotoTagsRequest struct {
	PhotoId   *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
}

func (s ListPhotoTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoTagsRequest) GoString() string {
	return s.String()
}

func (s *ListPhotoTagsRequest) SetPhotoId(v int64) *ListPhotoTagsRequest {
	s.PhotoId = &v
	return s
}

func (s *ListPhotoTagsRequest) SetStoreName(v string) *ListPhotoTagsRequest {
	s.StoreName = &v
	return s
}

func (s *ListPhotoTagsRequest) SetLibraryId(v string) *ListPhotoTagsRequest {
	s.LibraryId = &v
	return s
}

func (s *ListPhotoTagsRequest) SetLang(v string) *ListPhotoTagsRequest {
	s.Lang = &v
	return s
}

type ListPhotoTagsResponseBody struct {
	Action    *string                          `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                          `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                          `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                          `json:"Code,omitempty" xml:"Code,omitempty"`
	Tags      []*ListPhotoTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListPhotoTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListPhotoTagsResponseBody) SetAction(v string) *ListPhotoTagsResponseBody {
	s.Action = &v
	return s
}

func (s *ListPhotoTagsResponseBody) SetMessage(v string) *ListPhotoTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListPhotoTagsResponseBody) SetRequestId(v string) *ListPhotoTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPhotoTagsResponseBody) SetCode(v string) *ListPhotoTagsResponseBody {
	s.Code = &v
	return s
}

func (s *ListPhotoTagsResponseBody) SetTags(v []*ListPhotoTagsResponseBodyTags) *ListPhotoTagsResponseBody {
	s.Tags = v
	return s
}

type ListPhotoTagsResponseBodyTags struct {
	IdStr     *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	IsSubTag  *bool   `json:"IsSubTag,omitempty" xml:"IsSubTag,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentTag *string `json:"ParentTag,omitempty" xml:"ParentTag,omitempty"`
	Id        *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListPhotoTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListPhotoTagsResponseBodyTags) SetIdStr(v string) *ListPhotoTagsResponseBodyTags {
	s.IdStr = &v
	return s
}

func (s *ListPhotoTagsResponseBodyTags) SetIsSubTag(v bool) *ListPhotoTagsResponseBodyTags {
	s.IsSubTag = &v
	return s
}

func (s *ListPhotoTagsResponseBodyTags) SetName(v string) *ListPhotoTagsResponseBodyTags {
	s.Name = &v
	return s
}

func (s *ListPhotoTagsResponseBodyTags) SetParentTag(v string) *ListPhotoTagsResponseBodyTags {
	s.ParentTag = &v
	return s
}

func (s *ListPhotoTagsResponseBodyTags) SetId(v int64) *ListPhotoTagsResponseBodyTags {
	s.Id = &v
	return s
}

type ListPhotoTagsResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListPhotoTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListPhotoTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListPhotoTagsResponse) GoString() string {
	return s.String()
}

func (s *ListPhotoTagsResponse) SetHeaders(v map[string]*string) *ListPhotoTagsResponse {
	s.Headers = v
	return s
}

func (s *ListPhotoTagsResponse) SetBody(v *ListPhotoTagsResponseBody) *ListPhotoTagsResponse {
	s.Body = v
	return s
}

type ListRegisteredTagsRequest struct {
	StoreName *string   `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	Lang      []*string `json:"Lang,omitempty" xml:"Lang,omitempty" type:"Repeated"`
}

func (s ListRegisteredTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListRegisteredTagsRequest) GoString() string {
	return s.String()
}

func (s *ListRegisteredTagsRequest) SetStoreName(v string) *ListRegisteredTagsRequest {
	s.StoreName = &v
	return s
}

func (s *ListRegisteredTagsRequest) SetLang(v []*string) *ListRegisteredTagsRequest {
	s.Lang = v
	return s
}

type ListRegisteredTagsResponseBody struct {
	Action         *string                                         `json:"Action,omitempty" xml:"Action,omitempty"`
	Message        *string                                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId      *string                                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	RegisteredTags []*ListRegisteredTagsResponseBodyRegisteredTags `json:"RegisteredTags,omitempty" xml:"RegisteredTags,omitempty" type:"Repeated"`
	Code           *string                                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListRegisteredTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListRegisteredTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListRegisteredTagsResponseBody) SetAction(v string) *ListRegisteredTagsResponseBody {
	s.Action = &v
	return s
}

func (s *ListRegisteredTagsResponseBody) SetMessage(v string) *ListRegisteredTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListRegisteredTagsResponseBody) SetRequestId(v string) *ListRegisteredTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListRegisteredTagsResponseBody) SetRegisteredTags(v []*ListRegisteredTagsResponseBodyRegisteredTags) *ListRegisteredTagsResponseBody {
	s.RegisteredTags = v
	return s
}

func (s *ListRegisteredTagsResponseBody) SetCode(v string) *ListRegisteredTagsResponseBody {
	s.Code = &v
	return s
}

type ListRegisteredTagsResponseBodyRegisteredTags struct {
	TagValues []*ListRegisteredTagsResponseBodyRegisteredTagsTagValues `json:"TagValues,omitempty" xml:"TagValues,omitempty" type:"Repeated"`
	TagKey    *string                                                  `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
}

func (s ListRegisteredTagsResponseBodyRegisteredTags) String() string {
	return tea.Prettify(s)
}

func (s ListRegisteredTagsResponseBodyRegisteredTags) GoString() string {
	return s.String()
}

func (s *ListRegisteredTagsResponseBodyRegisteredTags) SetTagValues(v []*ListRegisteredTagsResponseBodyRegisteredTagsTagValues) *ListRegisteredTagsResponseBodyRegisteredTags {
	s.TagValues = v
	return s
}

func (s *ListRegisteredTagsResponseBodyRegisteredTags) SetTagKey(v string) *ListRegisteredTagsResponseBodyRegisteredTags {
	s.TagKey = &v
	return s
}

type ListRegisteredTagsResponseBodyRegisteredTagsTagValues struct {
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s ListRegisteredTagsResponseBodyRegisteredTagsTagValues) String() string {
	return tea.Prettify(s)
}

func (s ListRegisteredTagsResponseBodyRegisteredTagsTagValues) GoString() string {
	return s.String()
}

func (s *ListRegisteredTagsResponseBodyRegisteredTagsTagValues) SetLang(v string) *ListRegisteredTagsResponseBodyRegisteredTagsTagValues {
	s.Lang = &v
	return s
}

func (s *ListRegisteredTagsResponseBodyRegisteredTagsTagValues) SetText(v string) *ListRegisteredTagsResponseBodyRegisteredTagsTagValues {
	s.Text = &v
	return s
}

type ListRegisteredTagsResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListRegisteredTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListRegisteredTagsResponse) String() string {
	return tea.Prettify(s)
}

func (s ListRegisteredTagsResponse) GoString() string {
	return s.String()
}

func (s *ListRegisteredTagsResponse) SetHeaders(v map[string]*string) *ListRegisteredTagsResponse {
	s.Headers = v
	return s
}

func (s *ListRegisteredTagsResponse) SetBody(v *ListRegisteredTagsResponseBody) *ListRegisteredTagsResponse {
	s.Body = v
	return s
}

type ListTagPhotosRequest struct {
	TagId     *int64  `json:"TagId,omitempty" xml:"TagId,omitempty"`
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Cursor    *string `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	State     *string `json:"State,omitempty" xml:"State,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListTagPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagPhotosRequest) GoString() string {
	return s.String()
}

func (s *ListTagPhotosRequest) SetTagId(v int64) *ListTagPhotosRequest {
	s.TagId = &v
	return s
}

func (s *ListTagPhotosRequest) SetDirection(v string) *ListTagPhotosRequest {
	s.Direction = &v
	return s
}

func (s *ListTagPhotosRequest) SetSize(v int32) *ListTagPhotosRequest {
	s.Size = &v
	return s
}

func (s *ListTagPhotosRequest) SetCursor(v string) *ListTagPhotosRequest {
	s.Cursor = &v
	return s
}

func (s *ListTagPhotosRequest) SetState(v string) *ListTagPhotosRequest {
	s.State = &v
	return s
}

func (s *ListTagPhotosRequest) SetStoreName(v string) *ListTagPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *ListTagPhotosRequest) SetLibraryId(v string) *ListTagPhotosRequest {
	s.LibraryId = &v
	return s
}

type ListTagPhotosResponseBody struct {
	Action     *string                             `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                              `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	NextCursor *string                             `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	RequestId  *string                             `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Message    *string                             `json:"Message,omitempty" xml:"Message,omitempty"`
	Results    []*ListTagPhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code       *string                             `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListTagPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagPhotosResponseBody) SetAction(v string) *ListTagPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *ListTagPhotosResponseBody) SetTotalCount(v int32) *ListTagPhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTagPhotosResponseBody) SetNextCursor(v string) *ListTagPhotosResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListTagPhotosResponseBody) SetRequestId(v string) *ListTagPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagPhotosResponseBody) SetMessage(v string) *ListTagPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagPhotosResponseBody) SetResults(v []*ListTagPhotosResponseBodyResults) *ListTagPhotosResponseBody {
	s.Results = v
	return s
}

func (s *ListTagPhotosResponseBody) SetCode(v string) *ListTagPhotosResponseBody {
	s.Code = &v
	return s
}

type ListTagPhotosResponseBodyResults struct {
	PhotoIdStr *string `json:"PhotoIdStr,omitempty" xml:"PhotoIdStr,omitempty"`
	State      *string `json:"State,omitempty" xml:"State,omitempty"`
	PhotoId    *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
}

func (s ListTagPhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ListTagPhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ListTagPhotosResponseBodyResults) SetPhotoIdStr(v string) *ListTagPhotosResponseBodyResults {
	s.PhotoIdStr = &v
	return s
}

func (s *ListTagPhotosResponseBodyResults) SetState(v string) *ListTagPhotosResponseBodyResults {
	s.State = &v
	return s
}

func (s *ListTagPhotosResponseBodyResults) SetPhotoId(v int64) *ListTagPhotosResponseBodyResults {
	s.PhotoId = &v
	return s
}

type ListTagPhotosResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTagPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTagPhotosResponse) GoString() string {
	return s.String()
}

func (s *ListTagPhotosResponse) SetHeaders(v map[string]*string) *ListTagPhotosResponse {
	s.Headers = v
	return s
}

func (s *ListTagPhotosResponse) SetBody(v *ListTagPhotosResponseBody) *ListTagPhotosResponse {
	s.Body = v
	return s
}

type ListTagsRequest struct {
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListTagsRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTagsRequest) GoString() string {
	return s.String()
}

func (s *ListTagsRequest) SetLang(v string) *ListTagsRequest {
	s.Lang = &v
	return s
}

func (s *ListTagsRequest) SetStoreName(v string) *ListTagsRequest {
	s.StoreName = &v
	return s
}

func (s *ListTagsRequest) SetLibraryId(v string) *ListTagsRequest {
	s.LibraryId = &v
	return s
}

type ListTagsResponseBody struct {
	Action    *string                     `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                     `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Tags      []*ListTagsResponseBodyTags `json:"Tags,omitempty" xml:"Tags,omitempty" type:"Repeated"`
}

func (s ListTagsResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBody) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBody) SetAction(v string) *ListTagsResponseBody {
	s.Action = &v
	return s
}

func (s *ListTagsResponseBody) SetMessage(v string) *ListTagsResponseBody {
	s.Message = &v
	return s
}

func (s *ListTagsResponseBody) SetRequestId(v string) *ListTagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTagsResponseBody) SetCode(v string) *ListTagsResponseBody {
	s.Code = &v
	return s
}

func (s *ListTagsResponseBody) SetTags(v []*ListTagsResponseBodyTags) *ListTagsResponseBody {
	s.Tags = v
	return s
}

type ListTagsResponseBodyTags struct {
	IdStr     *string                        `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Cover     *ListTagsResponseBodyTagsCover `json:"Cover,omitempty" xml:"Cover,omitempty" type:"Struct"`
	IsSubTag  *bool                          `json:"IsSubTag,omitempty" xml:"IsSubTag,omitempty"`
	Name      *string                        `json:"Name,omitempty" xml:"Name,omitempty"`
	ParentTag *string                        `json:"ParentTag,omitempty" xml:"ParentTag,omitempty"`
	Id        *int64                         `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListTagsResponseBodyTags) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTags) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTags) SetIdStr(v string) *ListTagsResponseBodyTags {
	s.IdStr = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetCover(v *ListTagsResponseBodyTagsCover) *ListTagsResponseBodyTags {
	s.Cover = v
	return s
}

func (s *ListTagsResponseBodyTags) SetIsSubTag(v bool) *ListTagsResponseBodyTags {
	s.IsSubTag = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetName(v string) *ListTagsResponseBodyTags {
	s.Name = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetParentTag(v string) *ListTagsResponseBodyTags {
	s.ParentTag = &v
	return s
}

func (s *ListTagsResponseBodyTags) SetId(v int64) *ListTagsResponseBodyTags {
	s.Id = &v
	return s
}

type ListTagsResponseBodyTagsCover struct {
	Remark  *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	State   *string `json:"State,omitempty" xml:"State,omitempty"`
	Height  *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	FileId  *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Mtime   *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Ctime   *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Width   *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Md5     *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title   *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListTagsResponseBodyTagsCover) String() string {
	return tea.Prettify(s)
}

func (s ListTagsResponseBodyTagsCover) GoString() string {
	return s.String()
}

func (s *ListTagsResponseBodyTagsCover) SetRemark(v string) *ListTagsResponseBodyTagsCover {
	s.Remark = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetState(v string) *ListTagsResponseBodyTagsCover {
	s.State = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetHeight(v int64) *ListTagsResponseBodyTagsCover {
	s.Height = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetFileId(v string) *ListTagsResponseBodyTagsCover {
	s.FileId = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetIdStr(v string) *ListTagsResponseBodyTagsCover {
	s.IdStr = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetMtime(v int64) *ListTagsResponseBodyTagsCover {
	s.Mtime = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetCtime(v int64) *ListTagsResponseBodyTagsCover {
	s.Ctime = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetWidth(v int64) *ListTagsResponseBodyTagsCover {
	s.Width = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetMd5(v string) *ListTagsResponseBodyTagsCover {
	s.Md5 = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetIsVideo(v bool) *ListTagsResponseBodyTagsCover {
	s.IsVideo = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetTitle(v string) *ListTagsResponseBodyTagsCover {
	s.Title = &v
	return s
}

func (s *ListTagsResponseBodyTagsCover) SetId(v int64) *ListTagsResponseBodyTagsCover {
	s.Id = &v
	return s
}

type ListTagsResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTagsResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
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

func (s *ListTagsResponse) SetBody(v *ListTagsResponseBody) *ListTagsResponse {
	s.Body = v
	return s
}

type ListTimeLinePhotosRequest struct {
	Direction *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	StartTime *int64  `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	EndTime   *int64  `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	FilterBy  *string `json:"FilterBy,omitempty" xml:"FilterBy,omitempty"`
	Order     *string `json:"Order,omitempty" xml:"Order,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListTimeLinePhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTimeLinePhotosRequest) GoString() string {
	return s.String()
}

func (s *ListTimeLinePhotosRequest) SetDirection(v string) *ListTimeLinePhotosRequest {
	s.Direction = &v
	return s
}

func (s *ListTimeLinePhotosRequest) SetPage(v int32) *ListTimeLinePhotosRequest {
	s.Page = &v
	return s
}

func (s *ListTimeLinePhotosRequest) SetSize(v int32) *ListTimeLinePhotosRequest {
	s.Size = &v
	return s
}

func (s *ListTimeLinePhotosRequest) SetStartTime(v int64) *ListTimeLinePhotosRequest {
	s.StartTime = &v
	return s
}

func (s *ListTimeLinePhotosRequest) SetEndTime(v int64) *ListTimeLinePhotosRequest {
	s.EndTime = &v
	return s
}

func (s *ListTimeLinePhotosRequest) SetFilterBy(v string) *ListTimeLinePhotosRequest {
	s.FilterBy = &v
	return s
}

func (s *ListTimeLinePhotosRequest) SetOrder(v string) *ListTimeLinePhotosRequest {
	s.Order = &v
	return s
}

func (s *ListTimeLinePhotosRequest) SetStoreName(v string) *ListTimeLinePhotosRequest {
	s.StoreName = &v
	return s
}

func (s *ListTimeLinePhotosRequest) SetLibraryId(v string) *ListTimeLinePhotosRequest {
	s.LibraryId = &v
	return s
}

type ListTimeLinePhotosResponseBody struct {
	Photos     []*ListTimeLinePhotosResponseBodyPhotos `json:"Photos,omitempty" xml:"Photos,omitempty" type:"Repeated"`
	Action     *string                                 `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                                  `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListTimeLinePhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTimeLinePhotosResponseBody) GoString() string {
	return s.String()
}

func (s *ListTimeLinePhotosResponseBody) SetPhotos(v []*ListTimeLinePhotosResponseBodyPhotos) *ListTimeLinePhotosResponseBody {
	s.Photos = v
	return s
}

func (s *ListTimeLinePhotosResponseBody) SetAction(v string) *ListTimeLinePhotosResponseBody {
	s.Action = &v
	return s
}

func (s *ListTimeLinePhotosResponseBody) SetTotalCount(v int32) *ListTimeLinePhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListTimeLinePhotosResponseBody) SetMessage(v string) *ListTimeLinePhotosResponseBody {
	s.Message = &v
	return s
}

func (s *ListTimeLinePhotosResponseBody) SetRequestId(v string) *ListTimeLinePhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTimeLinePhotosResponseBody) SetCode(v string) *ListTimeLinePhotosResponseBody {
	s.Code = &v
	return s
}

type ListTimeLinePhotosResponseBodyPhotos struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Like            *int64  `json:"Like,omitempty" xml:"Like,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListTimeLinePhotosResponseBodyPhotos) String() string {
	return tea.Prettify(s)
}

func (s ListTimeLinePhotosResponseBodyPhotos) GoString() string {
	return s.String()
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetRemark(v string) *ListTimeLinePhotosResponseBodyPhotos {
	s.Remark = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetTakenAt(v int64) *ListTimeLinePhotosResponseBodyPhotos {
	s.TakenAt = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetState(v string) *ListTimeLinePhotosResponseBodyPhotos {
	s.State = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetHeight(v int64) *ListTimeLinePhotosResponseBodyPhotos {
	s.Height = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetShareExpireTime(v int64) *ListTimeLinePhotosResponseBodyPhotos {
	s.ShareExpireTime = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetFileId(v string) *ListTimeLinePhotosResponseBodyPhotos {
	s.FileId = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetIdStr(v string) *ListTimeLinePhotosResponseBodyPhotos {
	s.IdStr = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetCtime(v int64) *ListTimeLinePhotosResponseBodyPhotos {
	s.Ctime = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetMtime(v int64) *ListTimeLinePhotosResponseBodyPhotos {
	s.Mtime = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetSize(v int64) *ListTimeLinePhotosResponseBodyPhotos {
	s.Size = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetWidth(v int64) *ListTimeLinePhotosResponseBodyPhotos {
	s.Width = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetMd5(v string) *ListTimeLinePhotosResponseBodyPhotos {
	s.Md5 = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetIsVideo(v bool) *ListTimeLinePhotosResponseBodyPhotos {
	s.IsVideo = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetTitle(v string) *ListTimeLinePhotosResponseBodyPhotos {
	s.Title = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetLike(v int64) *ListTimeLinePhotosResponseBodyPhotos {
	s.Like = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetLocation(v string) *ListTimeLinePhotosResponseBodyPhotos {
	s.Location = &v
	return s
}

func (s *ListTimeLinePhotosResponseBodyPhotos) SetId(v int64) *ListTimeLinePhotosResponseBodyPhotos {
	s.Id = &v
	return s
}

type ListTimeLinePhotosResponse struct {
	Headers map[string]*string              `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTimeLinePhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTimeLinePhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTimeLinePhotosResponse) GoString() string {
	return s.String()
}

func (s *ListTimeLinePhotosResponse) SetHeaders(v map[string]*string) *ListTimeLinePhotosResponse {
	s.Headers = v
	return s
}

func (s *ListTimeLinePhotosResponse) SetBody(v *ListTimeLinePhotosResponseBody) *ListTimeLinePhotosResponse {
	s.Body = v
	return s
}

type ListTimeLinesRequest struct {
	Direction     *string `json:"Direction,omitempty" xml:"Direction,omitempty"`
	PhotoSize     *int32  `json:"PhotoSize,omitempty" xml:"PhotoSize,omitempty"`
	Cursor        *int64  `json:"Cursor,omitempty" xml:"Cursor,omitempty"`
	TimeLineCount *int32  `json:"TimeLineCount,omitempty" xml:"TimeLineCount,omitempty"`
	TimeLineUnit  *string `json:"TimeLineUnit,omitempty" xml:"TimeLineUnit,omitempty"`
	FilterBy      *string `json:"FilterBy,omitempty" xml:"FilterBy,omitempty"`
	Order         *string `json:"Order,omitempty" xml:"Order,omitempty"`
	StoreName     *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId     *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s ListTimeLinesRequest) String() string {
	return tea.Prettify(s)
}

func (s ListTimeLinesRequest) GoString() string {
	return s.String()
}

func (s *ListTimeLinesRequest) SetDirection(v string) *ListTimeLinesRequest {
	s.Direction = &v
	return s
}

func (s *ListTimeLinesRequest) SetPhotoSize(v int32) *ListTimeLinesRequest {
	s.PhotoSize = &v
	return s
}

func (s *ListTimeLinesRequest) SetCursor(v int64) *ListTimeLinesRequest {
	s.Cursor = &v
	return s
}

func (s *ListTimeLinesRequest) SetTimeLineCount(v int32) *ListTimeLinesRequest {
	s.TimeLineCount = &v
	return s
}

func (s *ListTimeLinesRequest) SetTimeLineUnit(v string) *ListTimeLinesRequest {
	s.TimeLineUnit = &v
	return s
}

func (s *ListTimeLinesRequest) SetFilterBy(v string) *ListTimeLinesRequest {
	s.FilterBy = &v
	return s
}

func (s *ListTimeLinesRequest) SetOrder(v string) *ListTimeLinesRequest {
	s.Order = &v
	return s
}

func (s *ListTimeLinesRequest) SetStoreName(v string) *ListTimeLinesRequest {
	s.StoreName = &v
	return s
}

func (s *ListTimeLinesRequest) SetLibraryId(v string) *ListTimeLinesRequest {
	s.LibraryId = &v
	return s
}

type ListTimeLinesResponseBody struct {
	Action     *string                               `json:"Action,omitempty" xml:"Action,omitempty"`
	NextCursor *int32                                `json:"NextCursor,omitempty" xml:"NextCursor,omitempty"`
	TimeLines  []*ListTimeLinesResponseBodyTimeLines `json:"TimeLines,omitempty" xml:"TimeLines,omitempty" type:"Repeated"`
	Message    *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ListTimeLinesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ListTimeLinesResponseBody) GoString() string {
	return s.String()
}

func (s *ListTimeLinesResponseBody) SetAction(v string) *ListTimeLinesResponseBody {
	s.Action = &v
	return s
}

func (s *ListTimeLinesResponseBody) SetNextCursor(v int32) *ListTimeLinesResponseBody {
	s.NextCursor = &v
	return s
}

func (s *ListTimeLinesResponseBody) SetTimeLines(v []*ListTimeLinesResponseBodyTimeLines) *ListTimeLinesResponseBody {
	s.TimeLines = v
	return s
}

func (s *ListTimeLinesResponseBody) SetMessage(v string) *ListTimeLinesResponseBody {
	s.Message = &v
	return s
}

func (s *ListTimeLinesResponseBody) SetRequestId(v string) *ListTimeLinesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListTimeLinesResponseBody) SetCode(v string) *ListTimeLinesResponseBody {
	s.Code = &v
	return s
}

type ListTimeLinesResponseBodyTimeLines struct {
	EndTime     *int64                                      `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	PhotosCount *int32                                      `json:"PhotosCount,omitempty" xml:"PhotosCount,omitempty"`
	StartTime   *int64                                      `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
	Photos      []*ListTimeLinesResponseBodyTimeLinesPhotos `json:"Photos,omitempty" xml:"Photos,omitempty" type:"Repeated"`
	TotalCount  *int32                                      `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListTimeLinesResponseBodyTimeLines) String() string {
	return tea.Prettify(s)
}

func (s ListTimeLinesResponseBodyTimeLines) GoString() string {
	return s.String()
}

func (s *ListTimeLinesResponseBodyTimeLines) SetEndTime(v int64) *ListTimeLinesResponseBodyTimeLines {
	s.EndTime = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLines) SetPhotosCount(v int32) *ListTimeLinesResponseBodyTimeLines {
	s.PhotosCount = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLines) SetStartTime(v int64) *ListTimeLinesResponseBodyTimeLines {
	s.StartTime = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLines) SetPhotos(v []*ListTimeLinesResponseBodyTimeLinesPhotos) *ListTimeLinesResponseBodyTimeLines {
	s.Photos = v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLines) SetTotalCount(v int32) *ListTimeLinesResponseBodyTimeLines {
	s.TotalCount = &v
	return s
}

type ListTimeLinesResponseBodyTimeLinesPhotos struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	Like            *int64  `json:"Like,omitempty" xml:"Like,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ListTimeLinesResponseBodyTimeLinesPhotos) String() string {
	return tea.Prettify(s)
}

func (s ListTimeLinesResponseBodyTimeLinesPhotos) GoString() string {
	return s.String()
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetRemark(v string) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Remark = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetTakenAt(v int64) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.TakenAt = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetState(v string) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.State = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetHeight(v int64) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Height = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetShareExpireTime(v int64) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.ShareExpireTime = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetFileId(v string) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.FileId = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetIdStr(v string) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.IdStr = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetCtime(v int64) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Ctime = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetMtime(v int64) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Mtime = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetSize(v int64) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Size = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetWidth(v int64) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Width = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetMd5(v string) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Md5 = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetIsVideo(v bool) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.IsVideo = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetTitle(v string) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Title = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetLike(v int64) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Like = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetLocation(v string) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Location = &v
	return s
}

func (s *ListTimeLinesResponseBodyTimeLinesPhotos) SetId(v int64) *ListTimeLinesResponseBodyTimeLinesPhotos {
	s.Id = &v
	return s
}

type ListTimeLinesResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ListTimeLinesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ListTimeLinesResponse) String() string {
	return tea.Prettify(s)
}

func (s ListTimeLinesResponse) GoString() string {
	return s.String()
}

func (s *ListTimeLinesResponse) SetHeaders(v map[string]*string) *ListTimeLinesResponse {
	s.Headers = v
	return s
}

func (s *ListTimeLinesResponse) SetBody(v *ListTimeLinesResponseBody) *ListTimeLinesResponse {
	s.Body = v
	return s
}

type MergeFacesRequest struct {
	TargetFaceId *int64  `json:"TargetFaceId,omitempty" xml:"TargetFaceId,omitempty"`
	StoreName    *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId    *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	FaceId       []*int  `json:"FaceId,omitempty" xml:"FaceId,omitempty" type:"Repeated"`
}

func (s MergeFacesRequest) String() string {
	return tea.Prettify(s)
}

func (s MergeFacesRequest) GoString() string {
	return s.String()
}

func (s *MergeFacesRequest) SetTargetFaceId(v int64) *MergeFacesRequest {
	s.TargetFaceId = &v
	return s
}

func (s *MergeFacesRequest) SetStoreName(v string) *MergeFacesRequest {
	s.StoreName = &v
	return s
}

func (s *MergeFacesRequest) SetLibraryId(v string) *MergeFacesRequest {
	s.LibraryId = &v
	return s
}

func (s *MergeFacesRequest) SetFaceId(v []*int) *MergeFacesRequest {
	s.FaceId = v
	return s
}

type MergeFacesResponseBody struct {
	Action    *string                        `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                        `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   *MergeFacesResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Struct"`
	Code      *string                        `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s MergeFacesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MergeFacesResponseBody) GoString() string {
	return s.String()
}

func (s *MergeFacesResponseBody) SetAction(v string) *MergeFacesResponseBody {
	s.Action = &v
	return s
}

func (s *MergeFacesResponseBody) SetMessage(v string) *MergeFacesResponseBody {
	s.Message = &v
	return s
}

func (s *MergeFacesResponseBody) SetRequestId(v string) *MergeFacesResponseBody {
	s.RequestId = &v
	return s
}

func (s *MergeFacesResponseBody) SetResults(v *MergeFacesResponseBodyResults) *MergeFacesResponseBody {
	s.Results = v
	return s
}

func (s *MergeFacesResponseBody) SetCode(v string) *MergeFacesResponseBody {
	s.Code = &v
	return s
}

type MergeFacesResponseBodyResults struct {
	Result []*MergeFacesResponseBodyResultsResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s MergeFacesResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s MergeFacesResponseBodyResults) GoString() string {
	return s.String()
}

func (s *MergeFacesResponseBodyResults) SetResult(v []*MergeFacesResponseBodyResultsResult) *MergeFacesResponseBodyResults {
	s.Result = v
	return s
}

type MergeFacesResponseBodyResultsResult struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s MergeFacesResponseBodyResultsResult) String() string {
	return tea.Prettify(s)
}

func (s MergeFacesResponseBodyResultsResult) GoString() string {
	return s.String()
}

func (s *MergeFacesResponseBodyResultsResult) SetIdStr(v string) *MergeFacesResponseBodyResultsResult {
	s.IdStr = &v
	return s
}

func (s *MergeFacesResponseBodyResultsResult) SetCode(v string) *MergeFacesResponseBodyResultsResult {
	s.Code = &v
	return s
}

func (s *MergeFacesResponseBodyResultsResult) SetMessage(v string) *MergeFacesResponseBodyResultsResult {
	s.Message = &v
	return s
}

func (s *MergeFacesResponseBodyResultsResult) SetId(v int64) *MergeFacesResponseBodyResultsResult {
	s.Id = &v
	return s
}

type MergeFacesResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MergeFacesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MergeFacesResponse) String() string {
	return tea.Prettify(s)
}

func (s MergeFacesResponse) GoString() string {
	return s.String()
}

func (s *MergeFacesResponse) SetHeaders(v map[string]*string) *MergeFacesResponse {
	s.Headers = v
	return s
}

func (s *MergeFacesResponse) SetBody(v *MergeFacesResponseBody) *MergeFacesResponse {
	s.Body = v
	return s
}

type MoveAlbumPhotosRequest struct {
	SourceAlbumId *int64  `json:"SourceAlbumId,omitempty" xml:"SourceAlbumId,omitempty"`
	TargetAlbumId *int64  `json:"TargetAlbumId,omitempty" xml:"TargetAlbumId,omitempty"`
	StoreName     *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId     *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId       []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s MoveAlbumPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveAlbumPhotosRequest) GoString() string {
	return s.String()
}

func (s *MoveAlbumPhotosRequest) SetSourceAlbumId(v int64) *MoveAlbumPhotosRequest {
	s.SourceAlbumId = &v
	return s
}

func (s *MoveAlbumPhotosRequest) SetTargetAlbumId(v int64) *MoveAlbumPhotosRequest {
	s.TargetAlbumId = &v
	return s
}

func (s *MoveAlbumPhotosRequest) SetStoreName(v string) *MoveAlbumPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *MoveAlbumPhotosRequest) SetLibraryId(v string) *MoveAlbumPhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *MoveAlbumPhotosRequest) SetPhotoId(v []*int) *MoveAlbumPhotosRequest {
	s.PhotoId = v
	return s
}

type MoveAlbumPhotosResponseBody struct {
	Action    *string                               `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                               `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*MoveAlbumPhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                               `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s MoveAlbumPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveAlbumPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *MoveAlbumPhotosResponseBody) SetAction(v string) *MoveAlbumPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *MoveAlbumPhotosResponseBody) SetMessage(v string) *MoveAlbumPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *MoveAlbumPhotosResponseBody) SetRequestId(v string) *MoveAlbumPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveAlbumPhotosResponseBody) SetResults(v []*MoveAlbumPhotosResponseBodyResults) *MoveAlbumPhotosResponseBody {
	s.Results = v
	return s
}

func (s *MoveAlbumPhotosResponseBody) SetCode(v string) *MoveAlbumPhotosResponseBody {
	s.Code = &v
	return s
}

type MoveAlbumPhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s MoveAlbumPhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s MoveAlbumPhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *MoveAlbumPhotosResponseBodyResults) SetIdStr(v string) *MoveAlbumPhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *MoveAlbumPhotosResponseBodyResults) SetCode(v string) *MoveAlbumPhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *MoveAlbumPhotosResponseBodyResults) SetMessage(v string) *MoveAlbumPhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *MoveAlbumPhotosResponseBodyResults) SetId(v int64) *MoveAlbumPhotosResponseBodyResults {
	s.Id = &v
	return s
}

type MoveAlbumPhotosResponse struct {
	Headers map[string]*string           `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveAlbumPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveAlbumPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveAlbumPhotosResponse) GoString() string {
	return s.String()
}

func (s *MoveAlbumPhotosResponse) SetHeaders(v map[string]*string) *MoveAlbumPhotosResponse {
	s.Headers = v
	return s
}

func (s *MoveAlbumPhotosResponse) SetBody(v *MoveAlbumPhotosResponseBody) *MoveAlbumPhotosResponse {
	s.Body = v
	return s
}

type MoveFacePhotosRequest struct {
	SourceFaceId *int64  `json:"SourceFaceId,omitempty" xml:"SourceFaceId,omitempty"`
	TargetFaceId *int64  `json:"TargetFaceId,omitempty" xml:"TargetFaceId,omitempty"`
	StoreName    *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId    *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId      []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s MoveFacePhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s MoveFacePhotosRequest) GoString() string {
	return s.String()
}

func (s *MoveFacePhotosRequest) SetSourceFaceId(v int64) *MoveFacePhotosRequest {
	s.SourceFaceId = &v
	return s
}

func (s *MoveFacePhotosRequest) SetTargetFaceId(v int64) *MoveFacePhotosRequest {
	s.TargetFaceId = &v
	return s
}

func (s *MoveFacePhotosRequest) SetStoreName(v string) *MoveFacePhotosRequest {
	s.StoreName = &v
	return s
}

func (s *MoveFacePhotosRequest) SetLibraryId(v string) *MoveFacePhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *MoveFacePhotosRequest) SetPhotoId(v []*int) *MoveFacePhotosRequest {
	s.PhotoId = v
	return s
}

type MoveFacePhotosResponseBody struct {
	Action    *string                              `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                              `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*MoveFacePhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                              `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s MoveFacePhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s MoveFacePhotosResponseBody) GoString() string {
	return s.String()
}

func (s *MoveFacePhotosResponseBody) SetAction(v string) *MoveFacePhotosResponseBody {
	s.Action = &v
	return s
}

func (s *MoveFacePhotosResponseBody) SetMessage(v string) *MoveFacePhotosResponseBody {
	s.Message = &v
	return s
}

func (s *MoveFacePhotosResponseBody) SetRequestId(v string) *MoveFacePhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *MoveFacePhotosResponseBody) SetResults(v []*MoveFacePhotosResponseBodyResults) *MoveFacePhotosResponseBody {
	s.Results = v
	return s
}

func (s *MoveFacePhotosResponseBody) SetCode(v string) *MoveFacePhotosResponseBody {
	s.Code = &v
	return s
}

type MoveFacePhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s MoveFacePhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s MoveFacePhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *MoveFacePhotosResponseBodyResults) SetIdStr(v string) *MoveFacePhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *MoveFacePhotosResponseBodyResults) SetCode(v string) *MoveFacePhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *MoveFacePhotosResponseBodyResults) SetMessage(v string) *MoveFacePhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *MoveFacePhotosResponseBodyResults) SetId(v int64) *MoveFacePhotosResponseBodyResults {
	s.Id = &v
	return s
}

type MoveFacePhotosResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *MoveFacePhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s MoveFacePhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s MoveFacePhotosResponse) GoString() string {
	return s.String()
}

func (s *MoveFacePhotosResponse) SetHeaders(v map[string]*string) *MoveFacePhotosResponse {
	s.Headers = v
	return s
}

func (s *MoveFacePhotosResponse) SetBody(v *MoveFacePhotosResponseBody) *MoveFacePhotosResponse {
	s.Body = v
	return s
}

type ReactivatePhotosRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s ReactivatePhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s ReactivatePhotosRequest) GoString() string {
	return s.String()
}

func (s *ReactivatePhotosRequest) SetStoreName(v string) *ReactivatePhotosRequest {
	s.StoreName = &v
	return s
}

func (s *ReactivatePhotosRequest) SetLibraryId(v string) *ReactivatePhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *ReactivatePhotosRequest) SetPhotoId(v []*int) *ReactivatePhotosRequest {
	s.PhotoId = v
	return s
}

type ReactivatePhotosResponseBody struct {
	Action    *string                                `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*ReactivatePhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ReactivatePhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ReactivatePhotosResponseBody) GoString() string {
	return s.String()
}

func (s *ReactivatePhotosResponseBody) SetAction(v string) *ReactivatePhotosResponseBody {
	s.Action = &v
	return s
}

func (s *ReactivatePhotosResponseBody) SetMessage(v string) *ReactivatePhotosResponseBody {
	s.Message = &v
	return s
}

func (s *ReactivatePhotosResponseBody) SetRequestId(v string) *ReactivatePhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *ReactivatePhotosResponseBody) SetResults(v []*ReactivatePhotosResponseBodyResults) *ReactivatePhotosResponseBody {
	s.Results = v
	return s
}

func (s *ReactivatePhotosResponseBody) SetCode(v string) *ReactivatePhotosResponseBody {
	s.Code = &v
	return s
}

type ReactivatePhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s ReactivatePhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s ReactivatePhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *ReactivatePhotosResponseBodyResults) SetIdStr(v string) *ReactivatePhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *ReactivatePhotosResponseBodyResults) SetCode(v string) *ReactivatePhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *ReactivatePhotosResponseBodyResults) SetMessage(v string) *ReactivatePhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *ReactivatePhotosResponseBodyResults) SetId(v int64) *ReactivatePhotosResponseBodyResults {
	s.Id = &v
	return s
}

type ReactivatePhotosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ReactivatePhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ReactivatePhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s ReactivatePhotosResponse) GoString() string {
	return s.String()
}

func (s *ReactivatePhotosResponse) SetHeaders(v map[string]*string) *ReactivatePhotosResponse {
	s.Headers = v
	return s
}

func (s *ReactivatePhotosResponse) SetBody(v *ReactivatePhotosResponseBody) *ReactivatePhotosResponse {
	s.Body = v
	return s
}

type RegisterPhotoRequest struct {
	TakenAt    *int64   `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	Location   *string  `json:"Location,omitempty" xml:"Location,omitempty"`
	StoreName  *string  `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId  *string  `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	Latitude   *float32 `json:"Latitude,omitempty" xml:"Latitude,omitempty"`
	Longitude  *float32 `json:"Longitude,omitempty" xml:"Longitude,omitempty"`
	Width      *int32   `json:"Width,omitempty" xml:"Width,omitempty"`
	Height     *int32   `json:"Height,omitempty" xml:"Height,omitempty"`
	IsVideo    *string  `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Md5        *string  `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Size       *int64   `json:"Size,omitempty" xml:"Size,omitempty"`
	PhotoTitle *string  `json:"PhotoTitle,omitempty" xml:"PhotoTitle,omitempty"`
	Remark     *string  `json:"Remark,omitempty" xml:"Remark,omitempty"`
}

func (s RegisterPhotoRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterPhotoRequest) GoString() string {
	return s.String()
}

func (s *RegisterPhotoRequest) SetTakenAt(v int64) *RegisterPhotoRequest {
	s.TakenAt = &v
	return s
}

func (s *RegisterPhotoRequest) SetLocation(v string) *RegisterPhotoRequest {
	s.Location = &v
	return s
}

func (s *RegisterPhotoRequest) SetStoreName(v string) *RegisterPhotoRequest {
	s.StoreName = &v
	return s
}

func (s *RegisterPhotoRequest) SetLibraryId(v string) *RegisterPhotoRequest {
	s.LibraryId = &v
	return s
}

func (s *RegisterPhotoRequest) SetLatitude(v float32) *RegisterPhotoRequest {
	s.Latitude = &v
	return s
}

func (s *RegisterPhotoRequest) SetLongitude(v float32) *RegisterPhotoRequest {
	s.Longitude = &v
	return s
}

func (s *RegisterPhotoRequest) SetWidth(v int32) *RegisterPhotoRequest {
	s.Width = &v
	return s
}

func (s *RegisterPhotoRequest) SetHeight(v int32) *RegisterPhotoRequest {
	s.Height = &v
	return s
}

func (s *RegisterPhotoRequest) SetIsVideo(v string) *RegisterPhotoRequest {
	s.IsVideo = &v
	return s
}

func (s *RegisterPhotoRequest) SetMd5(v string) *RegisterPhotoRequest {
	s.Md5 = &v
	return s
}

func (s *RegisterPhotoRequest) SetSize(v int64) *RegisterPhotoRequest {
	s.Size = &v
	return s
}

func (s *RegisterPhotoRequest) SetPhotoTitle(v string) *RegisterPhotoRequest {
	s.PhotoTitle = &v
	return s
}

func (s *RegisterPhotoRequest) SetRemark(v string) *RegisterPhotoRequest {
	s.Remark = &v
	return s
}

type RegisterPhotoResponseBody struct {
	Action    *string                         `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                         `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                         `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Photo     *RegisterPhotoResponseBodyPhoto `json:"Photo,omitempty" xml:"Photo,omitempty" type:"Struct"`
	Code      *string                         `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RegisterPhotoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterPhotoResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterPhotoResponseBody) SetAction(v string) *RegisterPhotoResponseBody {
	s.Action = &v
	return s
}

func (s *RegisterPhotoResponseBody) SetMessage(v string) *RegisterPhotoResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterPhotoResponseBody) SetRequestId(v string) *RegisterPhotoResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterPhotoResponseBody) SetPhoto(v *RegisterPhotoResponseBodyPhoto) *RegisterPhotoResponseBody {
	s.Photo = v
	return s
}

func (s *RegisterPhotoResponseBody) SetCode(v string) *RegisterPhotoResponseBody {
	s.Code = &v
	return s
}

type RegisterPhotoResponseBodyPhoto struct {
	Remark          *string `json:"Remark,omitempty" xml:"Remark,omitempty"`
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s RegisterPhotoResponseBodyPhoto) String() string {
	return tea.Prettify(s)
}

func (s RegisterPhotoResponseBodyPhoto) GoString() string {
	return s.String()
}

func (s *RegisterPhotoResponseBodyPhoto) SetRemark(v string) *RegisterPhotoResponseBodyPhoto {
	s.Remark = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetTakenAt(v int64) *RegisterPhotoResponseBodyPhoto {
	s.TakenAt = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetState(v string) *RegisterPhotoResponseBodyPhoto {
	s.State = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetHeight(v int64) *RegisterPhotoResponseBodyPhoto {
	s.Height = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetShareExpireTime(v int64) *RegisterPhotoResponseBodyPhoto {
	s.ShareExpireTime = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetFileId(v string) *RegisterPhotoResponseBodyPhoto {
	s.FileId = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetIdStr(v string) *RegisterPhotoResponseBodyPhoto {
	s.IdStr = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetCtime(v int64) *RegisterPhotoResponseBodyPhoto {
	s.Ctime = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetMtime(v int64) *RegisterPhotoResponseBodyPhoto {
	s.Mtime = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetWidth(v int64) *RegisterPhotoResponseBodyPhoto {
	s.Width = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetSize(v int64) *RegisterPhotoResponseBodyPhoto {
	s.Size = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetMd5(v string) *RegisterPhotoResponseBodyPhoto {
	s.Md5 = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetTitle(v string) *RegisterPhotoResponseBodyPhoto {
	s.Title = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetIsVideo(v bool) *RegisterPhotoResponseBodyPhoto {
	s.IsVideo = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetId(v int64) *RegisterPhotoResponseBodyPhoto {
	s.Id = &v
	return s
}

func (s *RegisterPhotoResponseBodyPhoto) SetLocation(v string) *RegisterPhotoResponseBodyPhoto {
	s.Location = &v
	return s
}

type RegisterPhotoResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterPhotoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterPhotoResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterPhotoResponse) GoString() string {
	return s.String()
}

func (s *RegisterPhotoResponse) SetHeaders(v map[string]*string) *RegisterPhotoResponse {
	s.Headers = v
	return s
}

func (s *RegisterPhotoResponse) SetBody(v *RegisterPhotoResponseBody) *RegisterPhotoResponse {
	s.Body = v
	return s
}

type RegisterTagRequest struct {
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	TagKey    *string `json:"TagKey,omitempty" xml:"TagKey,omitempty"`
	Lang      *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	Text      *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s RegisterTagRequest) String() string {
	return tea.Prettify(s)
}

func (s RegisterTagRequest) GoString() string {
	return s.String()
}

func (s *RegisterTagRequest) SetStoreName(v string) *RegisterTagRequest {
	s.StoreName = &v
	return s
}

func (s *RegisterTagRequest) SetTagKey(v string) *RegisterTagRequest {
	s.TagKey = &v
	return s
}

func (s *RegisterTagRequest) SetLang(v string) *RegisterTagRequest {
	s.Lang = &v
	return s
}

func (s *RegisterTagRequest) SetText(v string) *RegisterTagRequest {
	s.Text = &v
	return s
}

type RegisterTagResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RegisterTagResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RegisterTagResponseBody) GoString() string {
	return s.String()
}

func (s *RegisterTagResponseBody) SetAction(v string) *RegisterTagResponseBody {
	s.Action = &v
	return s
}

func (s *RegisterTagResponseBody) SetMessage(v string) *RegisterTagResponseBody {
	s.Message = &v
	return s
}

func (s *RegisterTagResponseBody) SetRequestId(v string) *RegisterTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *RegisterTagResponseBody) SetCode(v string) *RegisterTagResponseBody {
	s.Code = &v
	return s
}

type RegisterTagResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RegisterTagResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RegisterTagResponse) String() string {
	return tea.Prettify(s)
}

func (s RegisterTagResponse) GoString() string {
	return s.String()
}

func (s *RegisterTagResponse) SetHeaders(v map[string]*string) *RegisterTagResponse {
	s.Headers = v
	return s
}

func (s *RegisterTagResponse) SetBody(v *RegisterTagResponseBody) *RegisterTagResponse {
	s.Body = v
	return s
}

type RemoveAlbumPhotosRequest struct {
	AlbumId   *int64  `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s RemoveAlbumPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveAlbumPhotosRequest) GoString() string {
	return s.String()
}

func (s *RemoveAlbumPhotosRequest) SetAlbumId(v int64) *RemoveAlbumPhotosRequest {
	s.AlbumId = &v
	return s
}

func (s *RemoveAlbumPhotosRequest) SetStoreName(v string) *RemoveAlbumPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *RemoveAlbumPhotosRequest) SetLibraryId(v string) *RemoveAlbumPhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *RemoveAlbumPhotosRequest) SetPhotoId(v []*int) *RemoveAlbumPhotosRequest {
	s.PhotoId = v
	return s
}

type RemoveAlbumPhotosResponseBody struct {
	Action    *string                                 `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                                 `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                 `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*RemoveAlbumPhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                                 `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RemoveAlbumPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveAlbumPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveAlbumPhotosResponseBody) SetAction(v string) *RemoveAlbumPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *RemoveAlbumPhotosResponseBody) SetMessage(v string) *RemoveAlbumPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveAlbumPhotosResponseBody) SetRequestId(v string) *RemoveAlbumPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveAlbumPhotosResponseBody) SetResults(v []*RemoveAlbumPhotosResponseBodyResults) *RemoveAlbumPhotosResponseBody {
	s.Results = v
	return s
}

func (s *RemoveAlbumPhotosResponseBody) SetCode(v string) *RemoveAlbumPhotosResponseBody {
	s.Code = &v
	return s
}

type RemoveAlbumPhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s RemoveAlbumPhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s RemoveAlbumPhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RemoveAlbumPhotosResponseBodyResults) SetIdStr(v string) *RemoveAlbumPhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *RemoveAlbumPhotosResponseBodyResults) SetCode(v string) *RemoveAlbumPhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *RemoveAlbumPhotosResponseBodyResults) SetMessage(v string) *RemoveAlbumPhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *RemoveAlbumPhotosResponseBodyResults) SetId(v int64) *RemoveAlbumPhotosResponseBodyResults {
	s.Id = &v
	return s
}

type RemoveAlbumPhotosResponse struct {
	Headers map[string]*string             `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveAlbumPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveAlbumPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveAlbumPhotosResponse) GoString() string {
	return s.String()
}

func (s *RemoveAlbumPhotosResponse) SetHeaders(v map[string]*string) *RemoveAlbumPhotosResponse {
	s.Headers = v
	return s
}

func (s *RemoveAlbumPhotosResponse) SetBody(v *RemoveAlbumPhotosResponseBody) *RemoveAlbumPhotosResponse {
	s.Body = v
	return s
}

type RemoveFacePhotosRequest struct {
	FaceId    *int64  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId   []*int  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty" type:"Repeated"`
}

func (s RemoveFacePhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s RemoveFacePhotosRequest) GoString() string {
	return s.String()
}

func (s *RemoveFacePhotosRequest) SetFaceId(v int64) *RemoveFacePhotosRequest {
	s.FaceId = &v
	return s
}

func (s *RemoveFacePhotosRequest) SetStoreName(v string) *RemoveFacePhotosRequest {
	s.StoreName = &v
	return s
}

func (s *RemoveFacePhotosRequest) SetLibraryId(v string) *RemoveFacePhotosRequest {
	s.LibraryId = &v
	return s
}

func (s *RemoveFacePhotosRequest) SetPhotoId(v []*int) *RemoveFacePhotosRequest {
	s.PhotoId = v
	return s
}

type RemoveFacePhotosResponseBody struct {
	Action    *string                                `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string                                `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Results   []*RemoveFacePhotosResponseBodyResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	Code      *string                                `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RemoveFacePhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RemoveFacePhotosResponseBody) GoString() string {
	return s.String()
}

func (s *RemoveFacePhotosResponseBody) SetAction(v string) *RemoveFacePhotosResponseBody {
	s.Action = &v
	return s
}

func (s *RemoveFacePhotosResponseBody) SetMessage(v string) *RemoveFacePhotosResponseBody {
	s.Message = &v
	return s
}

func (s *RemoveFacePhotosResponseBody) SetRequestId(v string) *RemoveFacePhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *RemoveFacePhotosResponseBody) SetResults(v []*RemoveFacePhotosResponseBodyResults) *RemoveFacePhotosResponseBody {
	s.Results = v
	return s
}

func (s *RemoveFacePhotosResponseBody) SetCode(v string) *RemoveFacePhotosResponseBody {
	s.Code = &v
	return s
}

type RemoveFacePhotosResponseBodyResults struct {
	IdStr   *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Code    *string `json:"Code,omitempty" xml:"Code,omitempty"`
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	Id      *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
}

func (s RemoveFacePhotosResponseBodyResults) String() string {
	return tea.Prettify(s)
}

func (s RemoveFacePhotosResponseBodyResults) GoString() string {
	return s.String()
}

func (s *RemoveFacePhotosResponseBodyResults) SetIdStr(v string) *RemoveFacePhotosResponseBodyResults {
	s.IdStr = &v
	return s
}

func (s *RemoveFacePhotosResponseBodyResults) SetCode(v string) *RemoveFacePhotosResponseBodyResults {
	s.Code = &v
	return s
}

func (s *RemoveFacePhotosResponseBodyResults) SetMessage(v string) *RemoveFacePhotosResponseBodyResults {
	s.Message = &v
	return s
}

func (s *RemoveFacePhotosResponseBodyResults) SetId(v int64) *RemoveFacePhotosResponseBodyResults {
	s.Id = &v
	return s
}

type RemoveFacePhotosResponse struct {
	Headers map[string]*string            `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RemoveFacePhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RemoveFacePhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s RemoveFacePhotosResponse) GoString() string {
	return s.String()
}

func (s *RemoveFacePhotosResponse) SetHeaders(v map[string]*string) *RemoveFacePhotosResponse {
	s.Headers = v
	return s
}

func (s *RemoveFacePhotosResponse) SetBody(v *RemoveFacePhotosResponseBody) *RemoveFacePhotosResponse {
	s.Body = v
	return s
}

type RenameAlbumRequest struct {
	AlbumId   *int64  `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	AlbumName *string `json:"AlbumName,omitempty" xml:"AlbumName,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s RenameAlbumRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameAlbumRequest) GoString() string {
	return s.String()
}

func (s *RenameAlbumRequest) SetAlbumId(v int64) *RenameAlbumRequest {
	s.AlbumId = &v
	return s
}

func (s *RenameAlbumRequest) SetAlbumName(v string) *RenameAlbumRequest {
	s.AlbumName = &v
	return s
}

func (s *RenameAlbumRequest) SetStoreName(v string) *RenameAlbumRequest {
	s.StoreName = &v
	return s
}

func (s *RenameAlbumRequest) SetLibraryId(v string) *RenameAlbumRequest {
	s.LibraryId = &v
	return s
}

type RenameAlbumResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RenameAlbumResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameAlbumResponseBody) GoString() string {
	return s.String()
}

func (s *RenameAlbumResponseBody) SetAction(v string) *RenameAlbumResponseBody {
	s.Action = &v
	return s
}

func (s *RenameAlbumResponseBody) SetMessage(v string) *RenameAlbumResponseBody {
	s.Message = &v
	return s
}

func (s *RenameAlbumResponseBody) SetRequestId(v string) *RenameAlbumResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameAlbumResponseBody) SetCode(v string) *RenameAlbumResponseBody {
	s.Code = &v
	return s
}

type RenameAlbumResponse struct {
	Headers map[string]*string       `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenameAlbumResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameAlbumResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameAlbumResponse) GoString() string {
	return s.String()
}

func (s *RenameAlbumResponse) SetHeaders(v map[string]*string) *RenameAlbumResponse {
	s.Headers = v
	return s
}

func (s *RenameAlbumResponse) SetBody(v *RenameAlbumResponseBody) *RenameAlbumResponse {
	s.Body = v
	return s
}

type RenameFaceRequest struct {
	FaceId    *int64  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	FaceName  *string `json:"FaceName,omitempty" xml:"FaceName,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s RenameFaceRequest) String() string {
	return tea.Prettify(s)
}

func (s RenameFaceRequest) GoString() string {
	return s.String()
}

func (s *RenameFaceRequest) SetFaceId(v int64) *RenameFaceRequest {
	s.FaceId = &v
	return s
}

func (s *RenameFaceRequest) SetFaceName(v string) *RenameFaceRequest {
	s.FaceName = &v
	return s
}

func (s *RenameFaceRequest) SetStoreName(v string) *RenameFaceRequest {
	s.StoreName = &v
	return s
}

func (s *RenameFaceRequest) SetLibraryId(v string) *RenameFaceRequest {
	s.LibraryId = &v
	return s
}

type RenameFaceResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s RenameFaceResponseBody) String() string {
	return tea.Prettify(s)
}

func (s RenameFaceResponseBody) GoString() string {
	return s.String()
}

func (s *RenameFaceResponseBody) SetAction(v string) *RenameFaceResponseBody {
	s.Action = &v
	return s
}

func (s *RenameFaceResponseBody) SetMessage(v string) *RenameFaceResponseBody {
	s.Message = &v
	return s
}

func (s *RenameFaceResponseBody) SetRequestId(v string) *RenameFaceResponseBody {
	s.RequestId = &v
	return s
}

func (s *RenameFaceResponseBody) SetCode(v string) *RenameFaceResponseBody {
	s.Code = &v
	return s
}

type RenameFaceResponse struct {
	Headers map[string]*string      `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *RenameFaceResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s RenameFaceResponse) String() string {
	return tea.Prettify(s)
}

func (s RenameFaceResponse) GoString() string {
	return s.String()
}

func (s *RenameFaceResponse) SetHeaders(v map[string]*string) *RenameFaceResponse {
	s.Headers = v
	return s
}

func (s *RenameFaceResponse) SetBody(v *RenameFaceResponseBody) *RenameFaceResponse {
	s.Body = v
	return s
}

type SearchPhotosRequest struct {
	Page      *int32  `json:"Page,omitempty" xml:"Page,omitempty"`
	Size      *int32  `json:"Size,omitempty" xml:"Size,omitempty"`
	Keyword   *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s SearchPhotosRequest) String() string {
	return tea.Prettify(s)
}

func (s SearchPhotosRequest) GoString() string {
	return s.String()
}

func (s *SearchPhotosRequest) SetPage(v int32) *SearchPhotosRequest {
	s.Page = &v
	return s
}

func (s *SearchPhotosRequest) SetSize(v int32) *SearchPhotosRequest {
	s.Size = &v
	return s
}

func (s *SearchPhotosRequest) SetKeyword(v string) *SearchPhotosRequest {
	s.Keyword = &v
	return s
}

func (s *SearchPhotosRequest) SetStoreName(v string) *SearchPhotosRequest {
	s.StoreName = &v
	return s
}

func (s *SearchPhotosRequest) SetLibraryId(v string) *SearchPhotosRequest {
	s.LibraryId = &v
	return s
}

type SearchPhotosResponseBody struct {
	Photos     []*SearchPhotosResponseBodyPhotos `json:"Photos,omitempty" xml:"Photos,omitempty" type:"Repeated"`
	Action     *string                           `json:"Action,omitempty" xml:"Action,omitempty"`
	TotalCount *int32                            `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
	Message    *string                           `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId  *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code       *string                           `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SearchPhotosResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SearchPhotosResponseBody) GoString() string {
	return s.String()
}

func (s *SearchPhotosResponseBody) SetPhotos(v []*SearchPhotosResponseBodyPhotos) *SearchPhotosResponseBody {
	s.Photos = v
	return s
}

func (s *SearchPhotosResponseBody) SetAction(v string) *SearchPhotosResponseBody {
	s.Action = &v
	return s
}

func (s *SearchPhotosResponseBody) SetTotalCount(v int32) *SearchPhotosResponseBody {
	s.TotalCount = &v
	return s
}

func (s *SearchPhotosResponseBody) SetMessage(v string) *SearchPhotosResponseBody {
	s.Message = &v
	return s
}

func (s *SearchPhotosResponseBody) SetRequestId(v string) *SearchPhotosResponseBody {
	s.RequestId = &v
	return s
}

func (s *SearchPhotosResponseBody) SetCode(v string) *SearchPhotosResponseBody {
	s.Code = &v
	return s
}

type SearchPhotosResponseBodyPhotos struct {
	TakenAt         *int64  `json:"TakenAt,omitempty" xml:"TakenAt,omitempty"`
	State           *string `json:"State,omitempty" xml:"State,omitempty"`
	Height          *int64  `json:"Height,omitempty" xml:"Height,omitempty"`
	ShareExpireTime *int64  `json:"ShareExpireTime,omitempty" xml:"ShareExpireTime,omitempty"`
	FileId          *string `json:"FileId,omitempty" xml:"FileId,omitempty"`
	IdStr           *string `json:"IdStr,omitempty" xml:"IdStr,omitempty"`
	Ctime           *int64  `json:"Ctime,omitempty" xml:"Ctime,omitempty"`
	Mtime           *int64  `json:"Mtime,omitempty" xml:"Mtime,omitempty"`
	Width           *int64  `json:"Width,omitempty" xml:"Width,omitempty"`
	Size            *int64  `json:"Size,omitempty" xml:"Size,omitempty"`
	Md5             *string `json:"Md5,omitempty" xml:"Md5,omitempty"`
	Title           *string `json:"Title,omitempty" xml:"Title,omitempty"`
	IsVideo         *bool   `json:"IsVideo,omitempty" xml:"IsVideo,omitempty"`
	Id              *int64  `json:"Id,omitempty" xml:"Id,omitempty"`
	Location        *string `json:"Location,omitempty" xml:"Location,omitempty"`
}

func (s SearchPhotosResponseBodyPhotos) String() string {
	return tea.Prettify(s)
}

func (s SearchPhotosResponseBodyPhotos) GoString() string {
	return s.String()
}

func (s *SearchPhotosResponseBodyPhotos) SetTakenAt(v int64) *SearchPhotosResponseBodyPhotos {
	s.TakenAt = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetState(v string) *SearchPhotosResponseBodyPhotos {
	s.State = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetHeight(v int64) *SearchPhotosResponseBodyPhotos {
	s.Height = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetShareExpireTime(v int64) *SearchPhotosResponseBodyPhotos {
	s.ShareExpireTime = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetFileId(v string) *SearchPhotosResponseBodyPhotos {
	s.FileId = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetIdStr(v string) *SearchPhotosResponseBodyPhotos {
	s.IdStr = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetCtime(v int64) *SearchPhotosResponseBodyPhotos {
	s.Ctime = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetMtime(v int64) *SearchPhotosResponseBodyPhotos {
	s.Mtime = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetWidth(v int64) *SearchPhotosResponseBodyPhotos {
	s.Width = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetSize(v int64) *SearchPhotosResponseBodyPhotos {
	s.Size = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetMd5(v string) *SearchPhotosResponseBodyPhotos {
	s.Md5 = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetTitle(v string) *SearchPhotosResponseBodyPhotos {
	s.Title = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetIsVideo(v bool) *SearchPhotosResponseBodyPhotos {
	s.IsVideo = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetId(v int64) *SearchPhotosResponseBodyPhotos {
	s.Id = &v
	return s
}

func (s *SearchPhotosResponseBodyPhotos) SetLocation(v string) *SearchPhotosResponseBodyPhotos {
	s.Location = &v
	return s
}

type SearchPhotosResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SearchPhotosResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SearchPhotosResponse) String() string {
	return tea.Prettify(s)
}

func (s SearchPhotosResponse) GoString() string {
	return s.String()
}

func (s *SearchPhotosResponse) SetHeaders(v map[string]*string) *SearchPhotosResponse {
	s.Headers = v
	return s
}

func (s *SearchPhotosResponse) SetBody(v *SearchPhotosResponseBody) *SearchPhotosResponse {
	s.Body = v
	return s
}

type SetAlbumCoverRequest struct {
	AlbumId   *int64  `json:"AlbumId,omitempty" xml:"AlbumId,omitempty"`
	PhotoId   *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s SetAlbumCoverRequest) String() string {
	return tea.Prettify(s)
}

func (s SetAlbumCoverRequest) GoString() string {
	return s.String()
}

func (s *SetAlbumCoverRequest) SetAlbumId(v int64) *SetAlbumCoverRequest {
	s.AlbumId = &v
	return s
}

func (s *SetAlbumCoverRequest) SetPhotoId(v int64) *SetAlbumCoverRequest {
	s.PhotoId = &v
	return s
}

func (s *SetAlbumCoverRequest) SetStoreName(v string) *SetAlbumCoverRequest {
	s.StoreName = &v
	return s
}

func (s *SetAlbumCoverRequest) SetLibraryId(v string) *SetAlbumCoverRequest {
	s.LibraryId = &v
	return s
}

type SetAlbumCoverResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SetAlbumCoverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetAlbumCoverResponseBody) GoString() string {
	return s.String()
}

func (s *SetAlbumCoverResponseBody) SetAction(v string) *SetAlbumCoverResponseBody {
	s.Action = &v
	return s
}

func (s *SetAlbumCoverResponseBody) SetMessage(v string) *SetAlbumCoverResponseBody {
	s.Message = &v
	return s
}

func (s *SetAlbumCoverResponseBody) SetRequestId(v string) *SetAlbumCoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetAlbumCoverResponseBody) SetCode(v string) *SetAlbumCoverResponseBody {
	s.Code = &v
	return s
}

type SetAlbumCoverResponse struct {
	Headers map[string]*string         `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetAlbumCoverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetAlbumCoverResponse) String() string {
	return tea.Prettify(s)
}

func (s SetAlbumCoverResponse) GoString() string {
	return s.String()
}

func (s *SetAlbumCoverResponse) SetHeaders(v map[string]*string) *SetAlbumCoverResponse {
	s.Headers = v
	return s
}

func (s *SetAlbumCoverResponse) SetBody(v *SetAlbumCoverResponseBody) *SetAlbumCoverResponse {
	s.Body = v
	return s
}

type SetFaceCoverRequest struct {
	FaceId    *int64  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	PhotoId   *int64  `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s SetFaceCoverRequest) String() string {
	return tea.Prettify(s)
}

func (s SetFaceCoverRequest) GoString() string {
	return s.String()
}

func (s *SetFaceCoverRequest) SetFaceId(v int64) *SetFaceCoverRequest {
	s.FaceId = &v
	return s
}

func (s *SetFaceCoverRequest) SetPhotoId(v int64) *SetFaceCoverRequest {
	s.PhotoId = &v
	return s
}

func (s *SetFaceCoverRequest) SetStoreName(v string) *SetFaceCoverRequest {
	s.StoreName = &v
	return s
}

func (s *SetFaceCoverRequest) SetLibraryId(v string) *SetFaceCoverRequest {
	s.LibraryId = &v
	return s
}

type SetFaceCoverResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SetFaceCoverResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetFaceCoverResponseBody) GoString() string {
	return s.String()
}

func (s *SetFaceCoverResponseBody) SetAction(v string) *SetFaceCoverResponseBody {
	s.Action = &v
	return s
}

func (s *SetFaceCoverResponseBody) SetMessage(v string) *SetFaceCoverResponseBody {
	s.Message = &v
	return s
}

func (s *SetFaceCoverResponseBody) SetRequestId(v string) *SetFaceCoverResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetFaceCoverResponseBody) SetCode(v string) *SetFaceCoverResponseBody {
	s.Code = &v
	return s
}

type SetFaceCoverResponse struct {
	Headers map[string]*string        `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetFaceCoverResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetFaceCoverResponse) String() string {
	return tea.Prettify(s)
}

func (s SetFaceCoverResponse) GoString() string {
	return s.String()
}

func (s *SetFaceCoverResponse) SetHeaders(v map[string]*string) *SetFaceCoverResponse {
	s.Headers = v
	return s
}

func (s *SetFaceCoverResponse) SetBody(v *SetFaceCoverResponseBody) *SetFaceCoverResponse {
	s.Body = v
	return s
}

type SetMeRequest struct {
	FaceId    *int64  `json:"FaceId,omitempty" xml:"FaceId,omitempty"`
	StoreName *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s SetMeRequest) String() string {
	return tea.Prettify(s)
}

func (s SetMeRequest) GoString() string {
	return s.String()
}

func (s *SetMeRequest) SetFaceId(v int64) *SetMeRequest {
	s.FaceId = &v
	return s
}

func (s *SetMeRequest) SetStoreName(v string) *SetMeRequest {
	s.StoreName = &v
	return s
}

func (s *SetMeRequest) SetLibraryId(v string) *SetMeRequest {
	s.LibraryId = &v
	return s
}

type SetMeResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SetMeResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetMeResponseBody) GoString() string {
	return s.String()
}

func (s *SetMeResponseBody) SetAction(v string) *SetMeResponseBody {
	s.Action = &v
	return s
}

func (s *SetMeResponseBody) SetMessage(v string) *SetMeResponseBody {
	s.Message = &v
	return s
}

func (s *SetMeResponseBody) SetRequestId(v string) *SetMeResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetMeResponseBody) SetCode(v string) *SetMeResponseBody {
	s.Code = &v
	return s
}

type SetMeResponse struct {
	Headers map[string]*string `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetMeResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetMeResponse) String() string {
	return tea.Prettify(s)
}

func (s SetMeResponse) GoString() string {
	return s.String()
}

func (s *SetMeResponse) SetHeaders(v map[string]*string) *SetMeResponse {
	s.Headers = v
	return s
}

func (s *SetMeResponse) SetBody(v *SetMeResponseBody) *SetMeResponse {
	s.Body = v
	return s
}

type SetQuotaRequest struct {
	TotalQuota *int64  `json:"TotalQuota,omitempty" xml:"TotalQuota,omitempty"`
	StoreName  *string `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId  *string `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
}

func (s SetQuotaRequest) String() string {
	return tea.Prettify(s)
}

func (s SetQuotaRequest) GoString() string {
	return s.String()
}

func (s *SetQuotaRequest) SetTotalQuota(v int64) *SetQuotaRequest {
	s.TotalQuota = &v
	return s
}

func (s *SetQuotaRequest) SetStoreName(v string) *SetQuotaRequest {
	s.StoreName = &v
	return s
}

func (s *SetQuotaRequest) SetLibraryId(v string) *SetQuotaRequest {
	s.LibraryId = &v
	return s
}

type SetQuotaResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s SetQuotaResponseBody) String() string {
	return tea.Prettify(s)
}

func (s SetQuotaResponseBody) GoString() string {
	return s.String()
}

func (s *SetQuotaResponseBody) SetAction(v string) *SetQuotaResponseBody {
	s.Action = &v
	return s
}

func (s *SetQuotaResponseBody) SetMessage(v string) *SetQuotaResponseBody {
	s.Message = &v
	return s
}

func (s *SetQuotaResponseBody) SetRequestId(v string) *SetQuotaResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetQuotaResponseBody) SetCode(v string) *SetQuotaResponseBody {
	s.Code = &v
	return s
}

type SetQuotaResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *SetQuotaResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s SetQuotaResponse) String() string {
	return tea.Prettify(s)
}

func (s SetQuotaResponse) GoString() string {
	return s.String()
}

func (s *SetQuotaResponse) SetHeaders(v map[string]*string) *SetQuotaResponse {
	s.Headers = v
	return s
}

func (s *SetQuotaResponse) SetBody(v *SetQuotaResponseBody) *SetQuotaResponse {
	s.Body = v
	return s
}

type TagPhotoRequest struct {
	StoreName  *string   `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	LibraryId  *string   `json:"LibraryId,omitempty" xml:"LibraryId,omitempty"`
	PhotoId    *int64    `json:"PhotoId,omitempty" xml:"PhotoId,omitempty"`
	TagKey     []*string `json:"TagKey,omitempty" xml:"TagKey,omitempty" type:"Repeated"`
	Confidence []*int    `json:"Confidence,omitempty" xml:"Confidence,omitempty" type:"Repeated"`
}

func (s TagPhotoRequest) String() string {
	return tea.Prettify(s)
}

func (s TagPhotoRequest) GoString() string {
	return s.String()
}

func (s *TagPhotoRequest) SetStoreName(v string) *TagPhotoRequest {
	s.StoreName = &v
	return s
}

func (s *TagPhotoRequest) SetLibraryId(v string) *TagPhotoRequest {
	s.LibraryId = &v
	return s
}

func (s *TagPhotoRequest) SetPhotoId(v int64) *TagPhotoRequest {
	s.PhotoId = &v
	return s
}

func (s *TagPhotoRequest) SetTagKey(v []*string) *TagPhotoRequest {
	s.TagKey = v
	return s
}

func (s *TagPhotoRequest) SetConfidence(v []*int) *TagPhotoRequest {
	s.Confidence = v
	return s
}

type TagPhotoResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s TagPhotoResponseBody) String() string {
	return tea.Prettify(s)
}

func (s TagPhotoResponseBody) GoString() string {
	return s.String()
}

func (s *TagPhotoResponseBody) SetAction(v string) *TagPhotoResponseBody {
	s.Action = &v
	return s
}

func (s *TagPhotoResponseBody) SetMessage(v string) *TagPhotoResponseBody {
	s.Message = &v
	return s
}

func (s *TagPhotoResponseBody) SetRequestId(v string) *TagPhotoResponseBody {
	s.RequestId = &v
	return s
}

func (s *TagPhotoResponseBody) SetCode(v string) *TagPhotoResponseBody {
	s.Code = &v
	return s
}

type TagPhotoResponse struct {
	Headers map[string]*string    `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *TagPhotoResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s TagPhotoResponse) String() string {
	return tea.Prettify(s)
}

func (s TagPhotoResponse) GoString() string {
	return s.String()
}

func (s *TagPhotoResponse) SetHeaders(v map[string]*string) *TagPhotoResponse {
	s.Headers = v
	return s
}

func (s *TagPhotoResponse) SetBody(v *TagPhotoResponseBody) *TagPhotoResponse {
	s.Body = v
	return s
}

type ToggleFeaturesRequest struct {
	StoreName        *string   `json:"StoreName,omitempty" xml:"StoreName,omitempty"`
	EnabledFeatures  []*string `json:"EnabledFeatures,omitempty" xml:"EnabledFeatures,omitempty" type:"Repeated"`
	DisabledFeatures []*string `json:"DisabledFeatures,omitempty" xml:"DisabledFeatures,omitempty" type:"Repeated"`
}

func (s ToggleFeaturesRequest) String() string {
	return tea.Prettify(s)
}

func (s ToggleFeaturesRequest) GoString() string {
	return s.String()
}

func (s *ToggleFeaturesRequest) SetStoreName(v string) *ToggleFeaturesRequest {
	s.StoreName = &v
	return s
}

func (s *ToggleFeaturesRequest) SetEnabledFeatures(v []*string) *ToggleFeaturesRequest {
	s.EnabledFeatures = v
	return s
}

func (s *ToggleFeaturesRequest) SetDisabledFeatures(v []*string) *ToggleFeaturesRequest {
	s.DisabledFeatures = v
	return s
}

type ToggleFeaturesResponseBody struct {
	Action    *string `json:"Action,omitempty" xml:"Action,omitempty"`
	Message   *string `json:"Message,omitempty" xml:"Message,omitempty"`
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Code      *string `json:"Code,omitempty" xml:"Code,omitempty"`
}

func (s ToggleFeaturesResponseBody) String() string {
	return tea.Prettify(s)
}

func (s ToggleFeaturesResponseBody) GoString() string {
	return s.String()
}

func (s *ToggleFeaturesResponseBody) SetAction(v string) *ToggleFeaturesResponseBody {
	s.Action = &v
	return s
}

func (s *ToggleFeaturesResponseBody) SetMessage(v string) *ToggleFeaturesResponseBody {
	s.Message = &v
	return s
}

func (s *ToggleFeaturesResponseBody) SetRequestId(v string) *ToggleFeaturesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ToggleFeaturesResponseBody) SetCode(v string) *ToggleFeaturesResponseBody {
	s.Code = &v
	return s
}

type ToggleFeaturesResponse struct {
	Headers map[string]*string          `json:"headers,omitempty" xml:"headers,omitempty" require:"true"`
	Body    *ToggleFeaturesResponseBody `json:"body,omitempty" xml:"body,omitempty" require:"true"`
}

func (s ToggleFeaturesResponse) String() string {
	return tea.Prettify(s)
}

func (s ToggleFeaturesResponse) GoString() string {
	return s.String()
}

func (s *ToggleFeaturesResponse) SetHeaders(v map[string]*string) *ToggleFeaturesResponse {
	s.Headers = v
	return s
}

func (s *ToggleFeaturesResponse) SetBody(v *ToggleFeaturesResponseBody) *ToggleFeaturesResponse {
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
	client.EndpointRule = tea.String("")
	_err = client.CheckConfig(config)
	if _err != nil {
		return _err
	}
	client.Endpoint, _err = client.GetEndpoint(tea.String("cloudphoto"), client.RegionId, client.EndpointRule, client.Network, client.Suffix, client.EndpointMap, client.Endpoint)
	if _err != nil {
		return _err
	}

	return nil
}

func (client *Client) GetEndpoint(productId *string, regionId *string, endpointRule *string, network *string, suffix *string, endpointMap map[string]*string, endpoint *string) (_result *string, _err error) {
	if !tea.BoolValue(util.Empty(endpoint)) {
		_result = endpoint
		return _result, _err
	}

	if !tea.BoolValue(util.IsUnset(endpointMap)) && !tea.BoolValue(util.Empty(endpointMap[tea.StringValue(regionId)])) {
		_result = endpointMap[tea.StringValue(regionId)]
		return _result, _err
	}

	_body, _err := endpointutil.GetEndpointRules(productId, regionId, endpointRule, network, suffix)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ActivatePhotosWithOptions(request *ActivatePhotosRequest, runtime *util.RuntimeOptions) (_result *ActivatePhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ActivatePhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ActivatePhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ActivatePhotos(request *ActivatePhotosRequest) (_result *ActivatePhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ActivatePhotosResponse{}
	_body, _err := client.ActivatePhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) AddAlbumPhotosWithOptions(request *AddAlbumPhotosRequest, runtime *util.RuntimeOptions) (_result *AddAlbumPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &AddAlbumPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("AddAlbumPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) AddAlbumPhotos(request *AddAlbumPhotosRequest) (_result *AddAlbumPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &AddAlbumPhotosResponse{}
	_body, _err := client.AddAlbumPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateAlbumWithOptions(request *CreateAlbumRequest, runtime *util.RuntimeOptions) (_result *CreateAlbumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateAlbumResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateAlbum"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateAlbum(request *CreateAlbumRequest) (_result *CreateAlbumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateAlbumResponse{}
	_body, _err := client.CreateAlbumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePhotoWithOptions(request *CreatePhotoRequest, runtime *util.RuntimeOptions) (_result *CreatePhotoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePhotoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePhoto"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePhoto(request *CreatePhotoRequest) (_result *CreatePhotoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePhotoResponse{}
	_body, _err := client.CreatePhotoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreatePhotoStoreWithOptions(request *CreatePhotoStoreRequest, runtime *util.RuntimeOptions) (_result *CreatePhotoStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreatePhotoStoreResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreatePhotoStore"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreatePhotoStore(request *CreatePhotoStoreRequest) (_result *CreatePhotoStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreatePhotoStoreResponse{}
	_body, _err := client.CreatePhotoStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) CreateTransactionWithOptions(request *CreateTransactionRequest, runtime *util.RuntimeOptions) (_result *CreateTransactionResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &CreateTransactionResponse{}
	_body, _err := client.DoRPCRequest(tea.String("CreateTransaction"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) CreateTransaction(request *CreateTransactionRequest) (_result *CreateTransactionResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &CreateTransactionResponse{}
	_body, _err := client.CreateTransactionWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteAlbumsWithOptions(request *DeleteAlbumsRequest, runtime *util.RuntimeOptions) (_result *DeleteAlbumsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteAlbumsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteAlbums"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteAlbums(request *DeleteAlbumsRequest) (_result *DeleteAlbumsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteAlbumsResponse{}
	_body, _err := client.DeleteAlbumsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteEventWithOptions(request *DeleteEventRequest, runtime *util.RuntimeOptions) (_result *DeleteEventResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteEventResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteEvent"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteEvent(request *DeleteEventRequest) (_result *DeleteEventResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteEventResponse{}
	_body, _err := client.DeleteEventWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeleteFacesWithOptions(request *DeleteFacesRequest, runtime *util.RuntimeOptions) (_result *DeleteFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeleteFacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeleteFaces"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeleteFaces(request *DeleteFacesRequest) (_result *DeleteFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeleteFacesResponse{}
	_body, _err := client.DeleteFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePhotosWithOptions(request *DeletePhotosRequest, runtime *util.RuntimeOptions) (_result *DeletePhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePhotos(request *DeletePhotosRequest) (_result *DeletePhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePhotosResponse{}
	_body, _err := client.DeletePhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) DeletePhotoStoreWithOptions(request *DeletePhotoStoreRequest, runtime *util.RuntimeOptions) (_result *DeletePhotoStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &DeletePhotoStoreResponse{}
	_body, _err := client.DoRPCRequest(tea.String("DeletePhotoStore"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) DeletePhotoStore(request *DeletePhotoStoreRequest) (_result *DeletePhotoStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &DeletePhotoStoreResponse{}
	_body, _err := client.DeletePhotoStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditPhotosWithOptions(request *EditPhotosRequest, runtime *util.RuntimeOptions) (_result *EditPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EditPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditPhotos(request *EditPhotosRequest) (_result *EditPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditPhotosResponse{}
	_body, _err := client.EditPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) EditPhotoStoreWithOptions(request *EditPhotoStoreRequest, runtime *util.RuntimeOptions) (_result *EditPhotoStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &EditPhotoStoreResponse{}
	_body, _err := client.DoRPCRequest(tea.String("EditPhotoStore"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) EditPhotoStore(request *EditPhotoStoreRequest) (_result *EditPhotoStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &EditPhotoStoreResponse{}
	_body, _err := client.EditPhotoStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FetchAlbumTagPhotosWithOptions(request *FetchAlbumTagPhotosRequest, runtime *util.RuntimeOptions) (_result *FetchAlbumTagPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FetchAlbumTagPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FetchAlbumTagPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FetchAlbumTagPhotos(request *FetchAlbumTagPhotosRequest) (_result *FetchAlbumTagPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchAlbumTagPhotosResponse{}
	_body, _err := client.FetchAlbumTagPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FetchLibrariesWithOptions(request *FetchLibrariesRequest, runtime *util.RuntimeOptions) (_result *FetchLibrariesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FetchLibrariesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FetchLibraries"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FetchLibraries(request *FetchLibrariesRequest) (_result *FetchLibrariesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchLibrariesResponse{}
	_body, _err := client.FetchLibrariesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FetchMomentPhotosWithOptions(request *FetchMomentPhotosRequest, runtime *util.RuntimeOptions) (_result *FetchMomentPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FetchMomentPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FetchMomentPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FetchMomentPhotos(request *FetchMomentPhotosRequest) (_result *FetchMomentPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchMomentPhotosResponse{}
	_body, _err := client.FetchMomentPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) FetchPhotosWithOptions(request *FetchPhotosRequest, runtime *util.RuntimeOptions) (_result *FetchPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &FetchPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("FetchPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) FetchPhotos(request *FetchPhotosRequest) (_result *FetchPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &FetchPhotosResponse{}
	_body, _err := client.FetchPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetAlbumsByNamesWithOptions(request *GetAlbumsByNamesRequest, runtime *util.RuntimeOptions) (_result *GetAlbumsByNamesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetAlbumsByNamesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetAlbumsByNames"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetAlbumsByNames(request *GetAlbumsByNamesRequest) (_result *GetAlbumsByNamesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetAlbumsByNamesResponse{}
	_body, _err := client.GetAlbumsByNamesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDownloadUrlWithOptions(request *GetDownloadUrlRequest, runtime *util.RuntimeOptions) (_result *GetDownloadUrlResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDownloadUrlResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDownloadUrl"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDownloadUrl(request *GetDownloadUrlRequest) (_result *GetDownloadUrlResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDownloadUrlResponse{}
	_body, _err := client.GetDownloadUrlWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetDownloadUrlsWithOptions(request *GetDownloadUrlsRequest, runtime *util.RuntimeOptions) (_result *GetDownloadUrlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetDownloadUrlsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetDownloadUrls"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetDownloadUrls(request *GetDownloadUrlsRequest) (_result *GetDownloadUrlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetDownloadUrlsResponse{}
	_body, _err := client.GetDownloadUrlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetFramedPhotoUrlsWithOptions(request *GetFramedPhotoUrlsRequest, runtime *util.RuntimeOptions) (_result *GetFramedPhotoUrlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetFramedPhotoUrlsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetFramedPhotoUrls"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetFramedPhotoUrls(request *GetFramedPhotoUrlsRequest) (_result *GetFramedPhotoUrlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetFramedPhotoUrlsResponse{}
	_body, _err := client.GetFramedPhotoUrlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetLibraryWithOptions(request *GetLibraryRequest, runtime *util.RuntimeOptions) (_result *GetLibraryResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetLibraryResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetLibrary"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetLibrary(request *GetLibraryRequest) (_result *GetLibraryResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetLibraryResponse{}
	_body, _err := client.GetLibraryWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPhotosWithOptions(request *GetPhotosRequest, runtime *util.RuntimeOptions) (_result *GetPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPhotos(request *GetPhotosRequest) (_result *GetPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhotosResponse{}
	_body, _err := client.GetPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPhotosByMd5sWithOptions(request *GetPhotosByMd5sRequest, runtime *util.RuntimeOptions) (_result *GetPhotosByMd5sResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPhotosByMd5sResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPhotosByMd5s"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPhotosByMd5s(request *GetPhotosByMd5sRequest) (_result *GetPhotosByMd5sResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhotosByMd5sResponse{}
	_body, _err := client.GetPhotosByMd5sWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPhotoStoreWithOptions(request *GetPhotoStoreRequest, runtime *util.RuntimeOptions) (_result *GetPhotoStoreResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPhotoStoreResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPhotoStore"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPhotoStore(request *GetPhotoStoreRequest) (_result *GetPhotoStoreResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPhotoStoreResponse{}
	_body, _err := client.GetPhotoStoreWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPrivateAccessUrlsWithOptions(request *GetPrivateAccessUrlsRequest, runtime *util.RuntimeOptions) (_result *GetPrivateAccessUrlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPrivateAccessUrlsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPrivateAccessUrls"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPrivateAccessUrls(request *GetPrivateAccessUrlsRequest) (_result *GetPrivateAccessUrlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPrivateAccessUrlsResponse{}
	_body, _err := client.GetPrivateAccessUrlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetPublicAccessUrlsWithOptions(request *GetPublicAccessUrlsRequest, runtime *util.RuntimeOptions) (_result *GetPublicAccessUrlsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetPublicAccessUrlsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetPublicAccessUrls"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetPublicAccessUrls(request *GetPublicAccessUrlsRequest) (_result *GetPublicAccessUrlsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetPublicAccessUrlsResponse{}
	_body, _err := client.GetPublicAccessUrlsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetQuotaWithOptions(request *GetQuotaRequest, runtime *util.RuntimeOptions) (_result *GetQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetQuota"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetQuota(request *GetQuotaRequest) (_result *GetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetQuotaResponse{}
	_body, _err := client.GetQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetSimilarPhotosWithOptions(request *GetSimilarPhotosRequest, runtime *util.RuntimeOptions) (_result *GetSimilarPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetSimilarPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetSimilarPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetSimilarPhotos(request *GetSimilarPhotosRequest) (_result *GetSimilarPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetSimilarPhotosResponse{}
	_body, _err := client.GetSimilarPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetThumbnailWithOptions(request *GetThumbnailRequest, runtime *util.RuntimeOptions) (_result *GetThumbnailResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetThumbnailResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetThumbnail"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetThumbnail(request *GetThumbnailRequest) (_result *GetThumbnailResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetThumbnailResponse{}
	_body, _err := client.GetThumbnailWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetThumbnailsWithOptions(request *GetThumbnailsRequest, runtime *util.RuntimeOptions) (_result *GetThumbnailsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetThumbnailsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetThumbnails"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetThumbnails(request *GetThumbnailsRequest) (_result *GetThumbnailsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetThumbnailsResponse{}
	_body, _err := client.GetThumbnailsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) GetVideoCoverWithOptions(request *GetVideoCoverRequest, runtime *util.RuntimeOptions) (_result *GetVideoCoverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &GetVideoCoverResponse{}
	_body, _err := client.DoRPCRequest(tea.String("GetVideoCover"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) GetVideoCover(request *GetVideoCoverRequest) (_result *GetVideoCoverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &GetVideoCoverResponse{}
	_body, _err := client.GetVideoCoverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) InactivatePhotosWithOptions(request *InactivatePhotosRequest, runtime *util.RuntimeOptions) (_result *InactivatePhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &InactivatePhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("InactivatePhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) InactivatePhotos(request *InactivatePhotosRequest) (_result *InactivatePhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &InactivatePhotosResponse{}
	_body, _err := client.InactivatePhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) LikePhotoWithOptions(request *LikePhotoRequest, runtime *util.RuntimeOptions) (_result *LikePhotoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &LikePhotoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("LikePhoto"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) LikePhoto(request *LikePhotoRequest) (_result *LikePhotoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &LikePhotoResponse{}
	_body, _err := client.LikePhotoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlbumPhotosWithOptions(request *ListAlbumPhotosRequest, runtime *util.RuntimeOptions) (_result *ListAlbumPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAlbumPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAlbumPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlbumPhotos(request *ListAlbumPhotosRequest) (_result *ListAlbumPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlbumPhotosResponse{}
	_body, _err := client.ListAlbumPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListAlbumsWithOptions(request *ListAlbumsRequest, runtime *util.RuntimeOptions) (_result *ListAlbumsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListAlbumsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListAlbums"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListAlbums(request *ListAlbumsRequest) (_result *ListAlbumsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListAlbumsResponse{}
	_body, _err := client.ListAlbumsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFacePhotosWithOptions(request *ListFacePhotosRequest, runtime *util.RuntimeOptions) (_result *ListFacePhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFacePhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFacePhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFacePhotos(request *ListFacePhotosRequest) (_result *ListFacePhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFacePhotosResponse{}
	_body, _err := client.ListFacePhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListFacesWithOptions(request *ListFacesRequest, runtime *util.RuntimeOptions) (_result *ListFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListFacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListFaces"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListFaces(request *ListFacesRequest) (_result *ListFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListFacesResponse{}
	_body, _err := client.ListFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMomentPhotosWithOptions(request *ListMomentPhotosRequest, runtime *util.RuntimeOptions) (_result *ListMomentPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMomentPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMomentPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMomentPhotos(request *ListMomentPhotosRequest) (_result *ListMomentPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMomentPhotosResponse{}
	_body, _err := client.ListMomentPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListMomentsWithOptions(request *ListMomentsRequest, runtime *util.RuntimeOptions) (_result *ListMomentsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListMomentsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListMoments"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListMoments(request *ListMomentsRequest) (_result *ListMomentsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListMomentsResponse{}
	_body, _err := client.ListMomentsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPhotoFacesWithOptions(request *ListPhotoFacesRequest, runtime *util.RuntimeOptions) (_result *ListPhotoFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPhotoFacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPhotoFaces"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPhotoFaces(request *ListPhotoFacesRequest) (_result *ListPhotoFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPhotoFacesResponse{}
	_body, _err := client.ListPhotoFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPhotosWithOptions(request *ListPhotosRequest, runtime *util.RuntimeOptions) (_result *ListPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPhotos(request *ListPhotosRequest) (_result *ListPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPhotosResponse{}
	_body, _err := client.ListPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPhotoStoresWithOptions(runtime *util.RuntimeOptions) (_result *ListPhotoStoresResponse, _err error) {
	req := &openapi.OpenApiRequest{}
	_result = &ListPhotoStoresResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPhotoStores"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPhotoStores() (_result *ListPhotoStoresResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPhotoStoresResponse{}
	_body, _err := client.ListPhotoStoresWithOptions(runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListPhotoTagsWithOptions(request *ListPhotoTagsRequest, runtime *util.RuntimeOptions) (_result *ListPhotoTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListPhotoTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListPhotoTags"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListPhotoTags(request *ListPhotoTagsRequest) (_result *ListPhotoTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListPhotoTagsResponse{}
	_body, _err := client.ListPhotoTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListRegisteredTagsWithOptions(request *ListRegisteredTagsRequest, runtime *util.RuntimeOptions) (_result *ListRegisteredTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListRegisteredTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListRegisteredTags"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListRegisteredTags(request *ListRegisteredTagsRequest) (_result *ListRegisteredTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListRegisteredTagsResponse{}
	_body, _err := client.ListRegisteredTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagPhotosWithOptions(request *ListTagPhotosRequest, runtime *util.RuntimeOptions) (_result *ListTagPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTagPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTagPhotos(request *ListTagPhotosRequest) (_result *ListTagPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagPhotosResponse{}
	_body, _err := client.ListTagPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTagsWithOptions(request *ListTagsRequest, runtime *util.RuntimeOptions) (_result *ListTagsResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTagsResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTags"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTags(request *ListTagsRequest) (_result *ListTagsResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTagsResponse{}
	_body, _err := client.ListTagsWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTimeLinePhotosWithOptions(request *ListTimeLinePhotosRequest, runtime *util.RuntimeOptions) (_result *ListTimeLinePhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTimeLinePhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTimeLinePhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTimeLinePhotos(request *ListTimeLinePhotosRequest) (_result *ListTimeLinePhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTimeLinePhotosResponse{}
	_body, _err := client.ListTimeLinePhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ListTimeLinesWithOptions(request *ListTimeLinesRequest, runtime *util.RuntimeOptions) (_result *ListTimeLinesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ListTimeLinesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ListTimeLines"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ListTimeLines(request *ListTimeLinesRequest) (_result *ListTimeLinesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ListTimeLinesResponse{}
	_body, _err := client.ListTimeLinesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MergeFacesWithOptions(request *MergeFacesRequest, runtime *util.RuntimeOptions) (_result *MergeFacesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MergeFacesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MergeFaces"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MergeFaces(request *MergeFacesRequest) (_result *MergeFacesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MergeFacesResponse{}
	_body, _err := client.MergeFacesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveAlbumPhotosWithOptions(request *MoveAlbumPhotosRequest, runtime *util.RuntimeOptions) (_result *MoveAlbumPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveAlbumPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveAlbumPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveAlbumPhotos(request *MoveAlbumPhotosRequest) (_result *MoveAlbumPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveAlbumPhotosResponse{}
	_body, _err := client.MoveAlbumPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) MoveFacePhotosWithOptions(request *MoveFacePhotosRequest, runtime *util.RuntimeOptions) (_result *MoveFacePhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &MoveFacePhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("MoveFacePhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) MoveFacePhotos(request *MoveFacePhotosRequest) (_result *MoveFacePhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &MoveFacePhotosResponse{}
	_body, _err := client.MoveFacePhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ReactivatePhotosWithOptions(request *ReactivatePhotosRequest, runtime *util.RuntimeOptions) (_result *ReactivatePhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ReactivatePhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ReactivatePhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ReactivatePhotos(request *ReactivatePhotosRequest) (_result *ReactivatePhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ReactivatePhotosResponse{}
	_body, _err := client.ReactivatePhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterPhotoWithOptions(request *RegisterPhotoRequest, runtime *util.RuntimeOptions) (_result *RegisterPhotoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterPhotoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterPhoto"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterPhoto(request *RegisterPhotoRequest) (_result *RegisterPhotoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterPhotoResponse{}
	_body, _err := client.RegisterPhotoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RegisterTagWithOptions(request *RegisterTagRequest, runtime *util.RuntimeOptions) (_result *RegisterTagResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RegisterTagResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RegisterTag"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RegisterTag(request *RegisterTagRequest) (_result *RegisterTagResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RegisterTagResponse{}
	_body, _err := client.RegisterTagWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveAlbumPhotosWithOptions(request *RemoveAlbumPhotosRequest, runtime *util.RuntimeOptions) (_result *RemoveAlbumPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveAlbumPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveAlbumPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveAlbumPhotos(request *RemoveAlbumPhotosRequest) (_result *RemoveAlbumPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveAlbumPhotosResponse{}
	_body, _err := client.RemoveAlbumPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RemoveFacePhotosWithOptions(request *RemoveFacePhotosRequest, runtime *util.RuntimeOptions) (_result *RemoveFacePhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RemoveFacePhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RemoveFacePhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RemoveFacePhotos(request *RemoveFacePhotosRequest) (_result *RemoveFacePhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RemoveFacePhotosResponse{}
	_body, _err := client.RemoveFacePhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameAlbumWithOptions(request *RenameAlbumRequest, runtime *util.RuntimeOptions) (_result *RenameAlbumResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenameAlbumResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenameAlbum"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameAlbum(request *RenameAlbumRequest) (_result *RenameAlbumResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameAlbumResponse{}
	_body, _err := client.RenameAlbumWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) RenameFaceWithOptions(request *RenameFaceRequest, runtime *util.RuntimeOptions) (_result *RenameFaceResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &RenameFaceResponse{}
	_body, _err := client.DoRPCRequest(tea.String("RenameFace"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) RenameFace(request *RenameFaceRequest) (_result *RenameFaceResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &RenameFaceResponse{}
	_body, _err := client.RenameFaceWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SearchPhotosWithOptions(request *SearchPhotosRequest, runtime *util.RuntimeOptions) (_result *SearchPhotosResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SearchPhotosResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SearchPhotos"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SearchPhotos(request *SearchPhotosRequest) (_result *SearchPhotosResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SearchPhotosResponse{}
	_body, _err := client.SearchPhotosWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetAlbumCoverWithOptions(request *SetAlbumCoverRequest, runtime *util.RuntimeOptions) (_result *SetAlbumCoverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetAlbumCoverResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetAlbumCover"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetAlbumCover(request *SetAlbumCoverRequest) (_result *SetAlbumCoverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetAlbumCoverResponse{}
	_body, _err := client.SetAlbumCoverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetFaceCoverWithOptions(request *SetFaceCoverRequest, runtime *util.RuntimeOptions) (_result *SetFaceCoverResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetFaceCoverResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetFaceCover"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetFaceCover(request *SetFaceCoverRequest) (_result *SetFaceCoverResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetFaceCoverResponse{}
	_body, _err := client.SetFaceCoverWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetMeWithOptions(request *SetMeRequest, runtime *util.RuntimeOptions) (_result *SetMeResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetMeResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetMe"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetMe(request *SetMeRequest) (_result *SetMeResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetMeResponse{}
	_body, _err := client.SetMeWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) SetQuotaWithOptions(request *SetQuotaRequest, runtime *util.RuntimeOptions) (_result *SetQuotaResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &SetQuotaResponse{}
	_body, _err := client.DoRPCRequest(tea.String("SetQuota"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) SetQuota(request *SetQuotaRequest) (_result *SetQuotaResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &SetQuotaResponse{}
	_body, _err := client.SetQuotaWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) TagPhotoWithOptions(request *TagPhotoRequest, runtime *util.RuntimeOptions) (_result *TagPhotoResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &TagPhotoResponse{}
	_body, _err := client.DoRPCRequest(tea.String("TagPhoto"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) TagPhoto(request *TagPhotoRequest) (_result *TagPhotoResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &TagPhotoResponse{}
	_body, _err := client.TagPhotoWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}

func (client *Client) ToggleFeaturesWithOptions(request *ToggleFeaturesRequest, runtime *util.RuntimeOptions) (_result *ToggleFeaturesResponse, _err error) {
	_err = util.ValidateModel(request)
	if _err != nil {
		return _result, _err
	}
	req := &openapi.OpenApiRequest{
		Body: util.ToMap(request),
	}
	_result = &ToggleFeaturesResponse{}
	_body, _err := client.DoRPCRequest(tea.String("ToggleFeatures"), tea.String("2017-07-11"), tea.String("HTTPS"), tea.String("POST"), tea.String("AK"), tea.String("json"), req, runtime)
	if _err != nil {
		return _result, _err
	}
	_err = tea.Convert(_body, &_result)
	return _result, _err
}

func (client *Client) ToggleFeatures(request *ToggleFeaturesRequest) (_result *ToggleFeaturesResponse, _err error) {
	runtime := &util.RuntimeOptions{}
	_result = &ToggleFeaturesResponse{}
	_body, _err := client.ToggleFeaturesWithOptions(request, runtime)
	if _err != nil {
		return _result, _err
	}
	_result = _body
	return _result, _err
}
