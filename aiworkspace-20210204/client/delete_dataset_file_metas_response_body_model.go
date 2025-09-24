// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetFileMetasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedDetails(v []*DatasetFileMetaResponse) *DeleteDatasetFileMetasResponseBody
	GetFailedDetails() []*DatasetFileMetaResponse
	SetRequestId(v string) *DeleteDatasetFileMetasResponseBody
	GetRequestId() *string
	SetStatus(v bool) *DeleteDatasetFileMetasResponseBody
	GetStatus() *bool
}

type DeleteDatasetFileMetasResponseBody struct {
	// The metadata records that fail to be deleted for the dataset files.
	FailedDetails []*DatasetFileMetaResponse `json:"FailedDetails,omitempty" xml:"FailedDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the metadata records of all dataset files were deleted. The value true indicates that the metadata records of all dataset files are deleted. If the value is false, view the failure details specified by FailedDetails.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s DeleteDatasetFileMetasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetFileMetasResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetFileMetasResponseBody) GetFailedDetails() []*DatasetFileMetaResponse {
	return s.FailedDetails
}

func (s *DeleteDatasetFileMetasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetFileMetasResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *DeleteDatasetFileMetasResponseBody) SetFailedDetails(v []*DatasetFileMetaResponse) *DeleteDatasetFileMetasResponseBody {
	s.FailedDetails = v
	return s
}

func (s *DeleteDatasetFileMetasResponseBody) SetRequestId(v string) *DeleteDatasetFileMetasResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetFileMetasResponseBody) SetStatus(v bool) *DeleteDatasetFileMetasResponseBody {
	s.Status = &v
	return s
}

func (s *DeleteDatasetFileMetasResponseBody) Validate() error {
	return dara.Validate(s)
}
