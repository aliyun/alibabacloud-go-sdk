// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCdnFileDownloadCallbackInfo interface {
	dara.Model
	String() string
	GoString() string
	SetBucket(v string) *CdnFileDownloadCallbackInfo
	GetBucket() *string
	SetDomainId(v string) *CdnFileDownloadCallbackInfo
	GetDomainId() *string
	SetDriveId(v string) *CdnFileDownloadCallbackInfo
	GetDriveId() *string
	SetExpire(v int64) *CdnFileDownloadCallbackInfo
	GetExpire() *int64
	SetFileId(v string) *CdnFileDownloadCallbackInfo
	GetFileId() *string
	SetObject(v string) *CdnFileDownloadCallbackInfo
	GetObject() *string
	SetToken(v string) *CdnFileDownloadCallbackInfo
	GetToken() *string
	SetUserId(v string) *CdnFileDownloadCallbackInfo
	GetUserId() *string
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
	return dara.Prettify(s)
}

func (s CdnFileDownloadCallbackInfo) GoString() string {
	return s.String()
}

func (s *CdnFileDownloadCallbackInfo) GetBucket() *string {
	return s.Bucket
}

func (s *CdnFileDownloadCallbackInfo) GetDomainId() *string {
	return s.DomainId
}

func (s *CdnFileDownloadCallbackInfo) GetDriveId() *string {
	return s.DriveId
}

func (s *CdnFileDownloadCallbackInfo) GetExpire() *int64 {
	return s.Expire
}

func (s *CdnFileDownloadCallbackInfo) GetFileId() *string {
	return s.FileId
}

func (s *CdnFileDownloadCallbackInfo) GetObject() *string {
	return s.Object
}

func (s *CdnFileDownloadCallbackInfo) GetToken() *string {
	return s.Token
}

func (s *CdnFileDownloadCallbackInfo) GetUserId() *string {
	return s.UserId
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

func (s *CdnFileDownloadCallbackInfo) Validate() error {
	return dara.Validate(s)
}
