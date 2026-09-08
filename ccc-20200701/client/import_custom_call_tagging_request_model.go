// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iImportCustomCallTaggingRequest interface {
	dara.Model
	String() string
	GoString() string
	SetFilePath(v string) *ImportCustomCallTaggingRequest
	GetFilePath() *string
	SetInstanceId(v string) *ImportCustomCallTaggingRequest
	GetInstanceId() *string
}

type ImportCustomCallTaggingRequest struct {
	// The ObjectKey of the OSS object that contains the inbound number mark file. OSS is configured with the public customer storage bucket for Cloud Call Center. You can upload the inbound number mark file to this public customer storage bucket through the inbound management page of Cloud Call Center. After the upload succeeds, invoking this API reads the file content from OSS and imports it in batch. We do not recommend directly invoking this API. Instead, you can perform this operation through the default public cloud CRM System provided by Cloud Call Center.
	//
	// This parameter is required.
	//
	// example:
	//
	// temp/ImportCustomCallTagging.xlsx
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The instance ID.
	//
	// This parameter is required.
	//
	// example:
	//
	// ccc-test
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
}

func (s ImportCustomCallTaggingRequest) String() string {
	return dara.Prettify(s)
}

func (s ImportCustomCallTaggingRequest) GoString() string {
	return s.String()
}

func (s *ImportCustomCallTaggingRequest) GetFilePath() *string {
	return s.FilePath
}

func (s *ImportCustomCallTaggingRequest) GetInstanceId() *string {
	return s.InstanceId
}

func (s *ImportCustomCallTaggingRequest) SetFilePath(v string) *ImportCustomCallTaggingRequest {
	s.FilePath = &v
	return s
}

func (s *ImportCustomCallTaggingRequest) SetInstanceId(v string) *ImportCustomCallTaggingRequest {
	s.InstanceId = &v
	return s
}

func (s *ImportCustomCallTaggingRequest) Validate() error {
	return dara.Validate(s)
}
