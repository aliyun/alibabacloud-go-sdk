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
}

type CreateSkillFileCheckRequest struct {
	Files []*CreateSkillFileCheckRequestFiles `json:"Files,omitempty" xml:"Files,omitempty" type:"Repeated"`
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

func (s *CreateSkillFileCheckRequest) SetFiles(v []*CreateSkillFileCheckRequestFiles) *CreateSkillFileCheckRequest {
	s.Files = v
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
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
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

func (s *CreateSkillFileCheckRequestFiles) SetDownloadUrl(v string) *CreateSkillFileCheckRequestFiles {
	s.DownloadUrl = &v
	return s
}

func (s *CreateSkillFileCheckRequestFiles) SetFileName(v string) *CreateSkillFileCheckRequestFiles {
	s.FileName = &v
	return s
}

func (s *CreateSkillFileCheckRequestFiles) Validate() error {
	return dara.Validate(s)
}
