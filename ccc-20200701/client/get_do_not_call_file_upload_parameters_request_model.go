// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDoNotCallFileUploadParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetDoNotCallFileUploadParametersRequest
	GetFileName() *string
	SetInstanceId(v string) *GetDoNotCallFileUploadParametersRequest
	GetInstanceId() *string
}

type GetDoNotCallFileUploadParametersRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// blacklist.xlsx
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s GetDoNotCallFileUploadParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDoNotCallFileUploadParametersRequest) GoString() string {
	return s.String()
}

func (s *GetDoNotCallFileUploadParametersRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetDoNotCallFileUploadParametersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDoNotCallFileUploadParametersRequest) SetFileName(v string) *GetDoNotCallFileUploadParametersRequest {
	s.FileName = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersRequest) SetInstanceId(v string) *GetDoNotCallFileUploadParametersRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDoNotCallFileUploadParametersRequest) Validate() error {
	return dara.Validate(s)
}
