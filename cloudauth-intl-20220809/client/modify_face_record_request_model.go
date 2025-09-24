// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyFaceRecordRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFaceGroupCode(v string) *ModifyFaceRecordRequest
	GetFaceGroupCode() *string
	SetImgOssInfos(v string) *ModifyFaceRecordRequest
	GetImgOssInfos() *string
}

type ModifyFaceRecordRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 4EB35****87
	FaceGroupCode *string `json:"FaceGroupCode,omitempty" xml:"FaceGroupCode,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// JSON
	ImgOssInfos *string `json:"ImgOssInfos,omitempty" xml:"ImgOssInfos,omitempty"`
}

func (s ModifyFaceRecordRequest) String() string {
	return dara.Prettify(s)
}

func (s ModifyFaceRecordRequest) GoString() string {
	return s.String()
}

func (s *ModifyFaceRecordRequest) GetFaceGroupCode() *string {
	return s.FaceGroupCode
}

func (s *ModifyFaceRecordRequest) GetImgOssInfos() *string {
	return s.ImgOssInfos
}

func (s *ModifyFaceRecordRequest) SetFaceGroupCode(v string) *ModifyFaceRecordRequest {
	s.FaceGroupCode = &v
	return s
}

func (s *ModifyFaceRecordRequest) SetImgOssInfos(v string) *ModifyFaceRecordRequest {
	s.ImgOssInfos = &v
	return s
}

func (s *ModifyFaceRecordRequest) Validate() error {
	return dara.Validate(s)
}
