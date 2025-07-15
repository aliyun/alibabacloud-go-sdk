// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDocumentUploadParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFileName(v string) *GetDocumentUploadParametersRequest
	GetFileName() *string
	SetInstanceId(v string) *GetDocumentUploadParametersRequest
	GetInstanceId() *string
	SetRequestId(v string) *GetDocumentUploadParametersRequest
	GetRequestId() *string
}

type GetDocumentUploadParametersRequest struct {
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
	// 9cfad875-6260-4a53-ab6e-b13e3fb31f7d
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// example:
	//
	// 24BE19E8-BF7D-4992-A35E-15EBA874F2E5
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetDocumentUploadParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDocumentUploadParametersRequest) GoString() string {
	return s.String()
}

func (s *GetDocumentUploadParametersRequest) GetFileName() *string {
	return s.FileName
}

func (s *GetDocumentUploadParametersRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *GetDocumentUploadParametersRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDocumentUploadParametersRequest) SetFileName(v string) *GetDocumentUploadParametersRequest {
	s.FileName = &v
	return s
}

func (s *GetDocumentUploadParametersRequest) SetInstanceId(v string) *GetDocumentUploadParametersRequest {
	s.InstanceId = &v
	return s
}

func (s *GetDocumentUploadParametersRequest) SetRequestId(v string) *GetDocumentUploadParametersRequest {
	s.RequestId = &v
	return s
}

func (s *GetDocumentUploadParametersRequest) Validate() error {
	return dara.Validate(s)
}
