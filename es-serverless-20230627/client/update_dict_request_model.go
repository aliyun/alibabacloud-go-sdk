// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDictRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAllowCover(v bool) *UpdateDictRequest
	GetAllowCover() *bool
	SetFiles(v []*UpdateDictRequestFiles) *UpdateDictRequest
	GetFiles() []*UpdateDictRequestFiles
	SetSourceType(v string) *UpdateDictRequest
	GetSourceType() *string
	SetType(v string) *UpdateDictRequest
	GetType() *string
	SetDryRun(v bool) *UpdateDictRequest
	GetDryRun() *bool
}

type UpdateDictRequest struct {
	// example:
	//
	// true
	AllowCover *bool `json:"allowCover,omitempty" xml:"allowCover,omitempty"`
	// This parameter is required.
	Files []*UpdateDictRequestFiles `json:"files,omitempty" xml:"files,omitempty" type:"Repeated"`
	// example:
	//
	// OSS
	SourceType *string `json:"sourceType,omitempty" xml:"sourceType,omitempty"`
	// example:
	//
	// MAIN
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s UpdateDictRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateDictRequest) GoString() string {
	return s.String()
}

func (s *UpdateDictRequest) GetAllowCover() *bool {
	return s.AllowCover
}

func (s *UpdateDictRequest) GetFiles() []*UpdateDictRequestFiles {
	return s.Files
}

func (s *UpdateDictRequest) GetSourceType() *string {
	return s.SourceType
}

func (s *UpdateDictRequest) GetType() *string {
	return s.Type
}

func (s *UpdateDictRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *UpdateDictRequest) SetAllowCover(v bool) *UpdateDictRequest {
	s.AllowCover = &v
	return s
}

func (s *UpdateDictRequest) SetFiles(v []*UpdateDictRequestFiles) *UpdateDictRequest {
	s.Files = v
	return s
}

func (s *UpdateDictRequest) SetSourceType(v string) *UpdateDictRequest {
	s.SourceType = &v
	return s
}

func (s *UpdateDictRequest) SetType(v string) *UpdateDictRequest {
	s.Type = &v
	return s
}

func (s *UpdateDictRequest) SetDryRun(v bool) *UpdateDictRequest {
	s.DryRun = &v
	return s
}

func (s *UpdateDictRequest) Validate() error {
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

type UpdateDictRequestFiles struct {
	// example:
	//
	// dic_0.dic
	Name      *string                          `json:"name,omitempty" xml:"name,omitempty"`
	OssObject *UpdateDictRequestFilesOssObject `json:"ossObject,omitempty" xml:"ossObject,omitempty" type:"Struct"`
}

func (s UpdateDictRequestFiles) String() string {
	return dara.Prettify(s)
}

func (s UpdateDictRequestFiles) GoString() string {
	return s.String()
}

func (s *UpdateDictRequestFiles) GetName() *string {
	return s.Name
}

func (s *UpdateDictRequestFiles) GetOssObject() *UpdateDictRequestFilesOssObject {
	return s.OssObject
}

func (s *UpdateDictRequestFiles) SetName(v string) *UpdateDictRequestFiles {
	s.Name = &v
	return s
}

func (s *UpdateDictRequestFiles) SetOssObject(v *UpdateDictRequestFilesOssObject) *UpdateDictRequestFiles {
	s.OssObject = v
	return s
}

func (s *UpdateDictRequestFiles) Validate() error {
	if s.OssObject != nil {
		if err := s.OssObject.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type UpdateDictRequestFilesOssObject struct {
	// example:
	//
	// bucket1
	BucketName *string `json:"bucketName,omitempty" xml:"bucketName,omitempty"`
	// example:
	//
	// oss/dic_0.dic
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
}

func (s UpdateDictRequestFilesOssObject) String() string {
	return dara.Prettify(s)
}

func (s UpdateDictRequestFilesOssObject) GoString() string {
	return s.String()
}

func (s *UpdateDictRequestFilesOssObject) GetBucketName() *string {
	return s.BucketName
}

func (s *UpdateDictRequestFilesOssObject) GetKey() *string {
	return s.Key
}

func (s *UpdateDictRequestFilesOssObject) SetBucketName(v string) *UpdateDictRequestFilesOssObject {
	s.BucketName = &v
	return s
}

func (s *UpdateDictRequestFilesOssObject) SetKey(v string) *UpdateDictRequestFilesOssObject {
	s.Key = &v
	return s
}

func (s *UpdateDictRequestFilesOssObject) Validate() error {
	return dara.Validate(s)
}
