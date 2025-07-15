// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportDocumentsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *ImportDocumentsRequest
	GetInstanceId() *string
	SetOssFileKey(v string) *ImportDocumentsRequest
	GetOssFileKey() *string
	SetRequestId(v string) *ImportDocumentsRequest
	GetRequestId() *string
	SetSchemaId(v string) *ImportDocumentsRequest
	GetSchemaId() *string
}

type ImportDocumentsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// b0eb2742-f37e-4c67-82d4-25c651c1c450
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// ccc-test/test-file.wav
	OssFileKey *string `json:"OssFileKey,omitempty" xml:"OssFileKey,omitempty"`
	// example:
	//
	// 7BEEA660-A45A-45E3-98CC-AFC65E715C23
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// schema id
	//
	// This parameter is required.
	//
	// example:
	//
	// profile
	SchemaId *string `json:"SchemaId,omitempty" xml:"SchemaId,omitempty"`
}

func (s ImportDocumentsRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportDocumentsRequest) GoString() string {
	return s.String()
}

func (s *ImportDocumentsRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportDocumentsRequest) GetOssFileKey() *string {
	return s.OssFileKey
}

func (s *ImportDocumentsRequest) GetRequestId() *string {
	return s.RequestId
}

func (s *ImportDocumentsRequest) GetSchemaId() *string {
	return s.SchemaId
}

func (s *ImportDocumentsRequest) SetInstanceId(v string) *ImportDocumentsRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportDocumentsRequest) SetOssFileKey(v string) *ImportDocumentsRequest {
	s.OssFileKey = &v
	return s
}

func (s *ImportDocumentsRequest) SetRequestId(v string) *ImportDocumentsRequest {
	s.RequestId = &v
	return s
}

func (s *ImportDocumentsRequest) SetSchemaId(v string) *ImportDocumentsRequest {
	s.SchemaId = &v
	return s
}

func (s *ImportDocumentsRequest) Validate() error {
	return dara.Validate(s)
}
