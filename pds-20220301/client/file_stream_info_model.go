// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iFileStreamInfo interface {
	dara.Model
	String() string
	GoString() string
	SetContentHash(v string) *FileStreamInfo
	GetContentHash() *string
	SetContentHashName(v string) *FileStreamInfo
	GetContentHashName() *string
	SetContentMd5(v string) *FileStreamInfo
	GetContentMd5() *string
	SetPartInfoList(v []*UploadPartInfo) *FileStreamInfo
	GetPartInfoList() []*UploadPartInfo
	SetPreHash(v string) *FileStreamInfo
	GetPreHash() *string
	SetProofCode(v string) *FileStreamInfo
	GetProofCode() *string
	SetProofVersion(v string) *FileStreamInfo
	GetProofVersion() *string
	SetSize(v int64) *FileStreamInfo
	GetSize() *int64
}

type FileStreamInfo struct {
	ContentHash     *string           `json:"content_hash,omitempty" xml:"content_hash,omitempty"`
	ContentHashName *string           `json:"content_hash_name,omitempty" xml:"content_hash_name,omitempty"`
	ContentMd5      *string           `json:"content_md5,omitempty" xml:"content_md5,omitempty"`
	PartInfoList    []*UploadPartInfo `json:"part_info_list,omitempty" xml:"part_info_list,omitempty" type:"Repeated"`
	PreHash         *string           `json:"pre_hash,omitempty" xml:"pre_hash,omitempty"`
	ProofCode       *string           `json:"proof_code,omitempty" xml:"proof_code,omitempty"`
	ProofVersion    *string           `json:"proof_version,omitempty" xml:"proof_version,omitempty"`
	Size            *int64            `json:"size,omitempty" xml:"size,omitempty"`
}

func (s FileStreamInfo) String() string {
	return dara.Prettify(s)
}

func (s FileStreamInfo) GoString() string {
	return s.String()
}

func (s *FileStreamInfo) GetContentHash() *string {
	return s.ContentHash
}

func (s *FileStreamInfo) GetContentHashName() *string {
	return s.ContentHashName
}

func (s *FileStreamInfo) GetContentMd5() *string {
	return s.ContentMd5
}

func (s *FileStreamInfo) GetPartInfoList() []*UploadPartInfo {
	return s.PartInfoList
}

func (s *FileStreamInfo) GetPreHash() *string {
	return s.PreHash
}

func (s *FileStreamInfo) GetProofCode() *string {
	return s.ProofCode
}

func (s *FileStreamInfo) GetProofVersion() *string {
	return s.ProofVersion
}

func (s *FileStreamInfo) GetSize() *int64 {
	return s.Size
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

func (s *FileStreamInfo) SetPartInfoList(v []*UploadPartInfo) *FileStreamInfo {
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

func (s *FileStreamInfo) Validate() error {
	if s.PartInfoList != nil {
		for _, item := range s.PartInfoList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
