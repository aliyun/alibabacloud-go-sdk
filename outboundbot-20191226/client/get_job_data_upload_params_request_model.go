// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetJobDataUploadParamsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetBusiType(v string) *GetJobDataUploadParamsRequest
	GetBusiType() *string
	SetFileName(v string) *GetJobDataUploadParamsRequest
	GetFileName() *string
	SetInstanceId(v string) *GetJobDataUploadParamsRequest
	GetInstanceId() *string
	SetPath(v string) *GetJobDataUploadParamsRequest
	GetPath() *string
	SetUniqueId(v string) *GetJobDataUploadParamsRequest
	GetUniqueId() *string
}

type GetJobDataUploadParamsRequest struct {
	// Business type
	//
	// example:
	//
	// SCRIPT_RECORDING
	BusiType *string `json:"BusiType,omitempty" xml:"BusiType,omitempty"`
	// File name
	//
	// This parameter is required.
	//
	// example:
	//
	// SHUKE_ZA_ol_2_开场白.wav
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Instance ID
	//
	// This parameter is required.
	//
	// example:
	//
	// 4eee9bf8-1319-468f-ac82-83c50ae389f8
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Relative path of file storage
	//
	// example:
	//
	// //airwaybill/1237185904146124802
	Path *string `json:"Path,omitempty" xml:"Path,omitempty"`
	// Unique business ID (prevents file name duplication; if left blank, the system automatically adds one)
	//
	// example:
	//
	// sas_siema_1477832102462645_siem_f07e90c2c147cf8cf1549ccb974e1956
	UniqueId *string `json:"UniqueId,omitempty" xml:"UniqueId,omitempty"`
}

func (s GetJobDataUploadParamsRequest) String() string {
	return dara.Prettify(s)
}

func (s GetJobDataUploadParamsRequest) GoString() string {
	return s.String()
}

func (s *GetJobDataUploadParamsRequest) GetBusiType() *string {
	return s.BusiType
}

func (s *GetJobDataUploadParamsRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetJobDataUploadParamsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetJobDataUploadParamsRequest) GetPath() *string {
	return s.Path
}

func (s *GetJobDataUploadParamsRequest) GetUniqueId() *string {
	return s.UniqueId
}

func (s *GetJobDataUploadParamsRequest) SetBusiType(v string) *GetJobDataUploadParamsRequest {
	s.BusiType = &v
	return s
}

func (s *GetJobDataUploadParamsRequest) SetFileName(v string) *GetJobDataUploadParamsRequest {
	s.FileName = &v
	return s
}

func (s *GetJobDataUploadParamsRequest) SetInstanceId(v string) *GetJobDataUploadParamsRequest {
	s.InstanceId = &v
	return s
}

func (s *GetJobDataUploadParamsRequest) SetPath(v string) *GetJobDataUploadParamsRequest {
	s.Path = &v
	return s
}

func (s *GetJobDataUploadParamsRequest) SetUniqueId(v string) *GetJobDataUploadParamsRequest {
	s.UniqueId = &v
	return s
}

func (s *GetJobDataUploadParamsRequest) Validate() error {
	return dara.Validate(s)
}
