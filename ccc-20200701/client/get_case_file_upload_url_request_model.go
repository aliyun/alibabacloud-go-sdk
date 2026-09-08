// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetCaseFileUploadUrlRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetCaseFileUploadUrlRequest
	GetFileName() *string
	SetInstanceId(v string) *GetCaseFileUploadUrlRequest
	GetInstanceId() *string
}

type GetCaseFileUploadUrlRequest struct {
	// The file name of the predictive outbound calling Activity list. The name must consist of uppercase and lowercase English letters, and the file format must be CSV.
	//
	// This parameter is required.
	//
	// example:
	//
	// case.csv
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetCaseFileUploadUrlRequest) String() string {
	return dara.Prettify(s)
}

func (s GetCaseFileUploadUrlRequest) GoString() string {
	return s.String()
}

func (s *GetCaseFileUploadUrlRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetCaseFileUploadUrlRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetCaseFileUploadUrlRequest) SetFileName(v string) *GetCaseFileUploadUrlRequest {
	s.FileName = &v
	return s
}

func (s *GetCaseFileUploadUrlRequest) SetInstanceId(v string) *GetCaseFileUploadUrlRequest {
	s.InstanceId = &v
	return s
}

func (s *GetCaseFileUploadUrlRequest) Validate() error {
	return dara.Validate(s)
}
