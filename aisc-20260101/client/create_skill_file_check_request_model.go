// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateSkillFileCheckRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFiles(v []*CreateSkillFileCheckRequestFiles) *CreateSkillFileCheckRequest
	GetFiles() []*CreateSkillFileCheckRequestFiles
	SetSource(v string) *CreateSkillFileCheckRequest
	GetSource() *string
}

type CreateSkillFileCheckRequest struct {
	// The file information.
	Files []*CreateSkillFileCheckRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
	// The upload source. If left empty, user_upload is used by default. Security operations agents use sec_ops_agent.
	//
	// example:
	//
	// sec_ops_agent
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
}

func (s CreateSkillFileCheckRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillFileCheckRequest) GoString() string {
	return s.String()
}

func (s *CreateSkillFileCheckRequest) GetFiles() []*CreateSkillFileCheckRequestFiles {
	return s.Files
}

func (s *CreateSkillFileCheckRequest) GetSource() *string {
	return s.Source
}

func (s *CreateSkillFileCheckRequest) SetFiles(v []*CreateSkillFileCheckRequestFiles) *CreateSkillFileCheckRequest {
	s.Files = v
	return s
}

func (s *CreateSkillFileCheckRequest) SetSource(v string) *CreateSkillFileCheckRequest {
	s.Source = &v
	return s
}

func (s *CreateSkillFileCheckRequest) Validate() error {
	if s.Files != nil {
		for _, item := range s.Files {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type CreateSkillFileCheckRequestFiles struct {
	// The public URL for downloading the file. The downloaded file must be a compressed package in tar.gz or zip format.
	//
	// example:
	//
	// https://test.oss-cn-hangzhou.aliyuncs.com/xxxx
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The file name. If not specified, the file name is parsed from DownloadUrl.
	//
	// example:
	//
	// test-file
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The tenant-isolated OSS temporary object key returned by GenerateSkillOssUploadCredential. Specify either this parameter or DownloadUrl.
	//
	// example:
	//
	// user-upload/staging/123456789/550e8400-e29b-41d4-a716-446655440000
	UploadKey *string `json:"UploadKey,omitempty" xml:"UploadKey,omitempty"`
}

func (s CreateSkillFileCheckRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s CreateSkillFileCheckRequestFiles) GoString() string {
	return s.String()
}

func (s *CreateSkillFileCheckRequestFiles) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *CreateSkillFileCheckRequestFiles) GetFileName() *string {
	return s.FileName
}

func (s *CreateSkillFileCheckRequestFiles) GetUploadKey() *string {
	return s.UploadKey
}

func (s *CreateSkillFileCheckRequestFiles) SetDownloadUrl(v string) *CreateSkillFileCheckRequestFiles {
	s.DownloadUrl = &v
	return s
}

func (s *CreateSkillFileCheckRequestFiles) SetFileName(v string) *CreateSkillFileCheckRequestFiles {
	s.FileName = &v
	return s
}

func (s *CreateSkillFileCheckRequestFiles) SetUploadKey(v string) *CreateSkillFileCheckRequestFiles {
	s.UploadKey = &v
	return s
}

func (s *CreateSkillFileCheckRequestFiles) Validate() error {
	return dara.Validate(s)
}
