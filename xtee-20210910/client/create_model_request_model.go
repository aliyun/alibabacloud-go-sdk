// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBucId(v string) *CreateModelRequest
	GetBucId() *string
	SetCounts(v string) *CreateModelRequest
	GetCounts() *string
	SetFileMD5(v string) *CreateModelRequest
	GetFileMD5() *string
	SetFilePath(v string) *CreateModelRequest
	GetFilePath() *string
	SetModelName(v string) *CreateModelRequest
	GetModelName() *string
	SetModelScene(v string) *CreateModelRequest
	GetModelScene() *string
	SetParameterNum(v string) *CreateModelRequest
	GetParameterNum() *string
	SetRegId(v string) *CreateModelRequest
	GetRegId() *string
	SetUserLocalFileName(v string) *CreateModelRequest
	GetUserLocalFileName() *string
}

type CreateModelRequest struct {
	// Submitter ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// WB01160353
	BucId *string `json:"BucId,omitempty" xml:"BucId,omitempty"`
	// Number of file rows.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1000
	Counts *string `json:"Counts,omitempty" xml:"Counts,omitempty"`
	// File MD5 value.
	//
	// This parameter is required.
	//
	// example:
	//
	// VC6Sj3u138yfWHLxLl7dtA==
	FileMD5 *string `json:"FileMD5,omitempty" xml:"FileMD5,omitempty"`
	// File path.
	//
	// This parameter is required.
	//
	// example:
	//
	// tempCache/ef2387dfb357ffe87925fd51d2b5305b/1665717035328/mockData.csv
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// Model name.
	//
	// This parameter is required.
	//
	// example:
	//
	// qwen-max
	ModelName *string `json:"ModelName,omitempty" xml:"ModelName,omitempty"`
	// Model scenario.
	//
	// This parameter is required.
	//
	// example:
	//
	// saf_ai_account_abuse
	ModelScene *string `json:"ModelScene,omitempty" xml:"ModelScene,omitempty"`
	// Number of file columns.
	//
	// This parameter is required.
	//
	// example:
	//
	// 10
	ParameterNum *string `json:"ParameterNum,omitempty" xml:"ParameterNum,omitempty"`
	// Region code
	//
	// example:
	//
	// cn-hangzhou
	RegId *string `json:"RegId,omitempty" xml:"RegId,omitempty"`
	// Uploaded file name.
	//
	// This parameter is required.
	//
	// example:
	//
	// fileTest
	UserLocalFileName *string `json:"UserLocalFileName,omitempty" xml:"UserLocalFileName,omitempty"`
}

func (s CreateModelRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateModelRequest) GoString() string {
	return s.String()
}

func (s *CreateModelRequest) GetBucId() *string {
	return s.BucId
}

func (s *CreateModelRequest) GetCounts() *string {
	return s.Counts
}

func (s *CreateModelRequest) GetFileMD5() *string {
	return s.FileMD5
}

func (s *CreateModelRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *CreateModelRequest) GetModelName() *string {
	return s.ModelName
}

func (s *CreateModelRequest) GetModelScene() *string {
	return s.ModelScene
}

func (s *CreateModelRequest) GetParameterNum() *string {
	return s.ParameterNum
}

func (s *CreateModelRequest) GetRegId() *string {
	return s.RegId
}

func (s *CreateModelRequest) GetUserLocalFileName() *string {
	return s.UserLocalFileName
}

func (s *CreateModelRequest) SetBucId(v string) *CreateModelRequest {
	s.BucId = &v
	return s
}

func (s *CreateModelRequest) SetCounts(v string) *CreateModelRequest {
	s.Counts = &v
	return s
}

func (s *CreateModelRequest) SetFileMD5(v string) *CreateModelRequest {
	s.FileMD5 = &v
	return s
}

func (s *CreateModelRequest) SetFilePath(v string) *CreateModelRequest {
	s.FilePath = &v
	return s
}

func (s *CreateModelRequest) SetModelName(v string) *CreateModelRequest {
	s.ModelName = &v
	return s
}

func (s *CreateModelRequest) SetModelScene(v string) *CreateModelRequest {
	s.ModelScene = &v
	return s
}

func (s *CreateModelRequest) SetParameterNum(v string) *CreateModelRequest {
	s.ParameterNum = &v
	return s
}

func (s *CreateModelRequest) SetRegId(v string) *CreateModelRequest {
	s.RegId = &v
	return s
}

func (s *CreateModelRequest) SetUserLocalFileName(v string) *CreateModelRequest {
	s.UserLocalFileName = &v
	return s
}

func (s *CreateModelRequest) Validate() error {
	return dara.Validate(s)
}
