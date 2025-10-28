// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetFileMetasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedDetails(v []*DatasetFileMetaResponse) *UpdateDatasetFileMetasResponseBody
	GetFailedDetails() []*DatasetFileMetaResponse
	SetRequestId(v string) *UpdateDatasetFileMetasResponseBody
	GetRequestId() *string
	SetStatus(v bool) *UpdateDatasetFileMetasResponseBody
	GetStatus() *bool
}

type UpdateDatasetFileMetasResponseBody struct {
	// The metadata records that fail to be updated for the dataset files.
	FailedDetails []*DatasetFileMetaResponse `json:"FailedDetails,omitempty" xml:"FailedDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the metadata records of all dataset files were updated. Valid values: true and false. If the value is false, view the failure details specified by FailedDetails.
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s UpdateDatasetFileMetasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetFileMetasResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetFileMetasResponseBody) GetFailedDetails() []*DatasetFileMetaResponse {
	return s.FailedDetails
}

func (s *UpdateDatasetFileMetasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasetFileMetasResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *UpdateDatasetFileMetasResponseBody) SetFailedDetails(v []*DatasetFileMetaResponse) *UpdateDatasetFileMetasResponseBody {
	s.FailedDetails = v
	return s
}

func (s *UpdateDatasetFileMetasResponseBody) SetRequestId(v string) *UpdateDatasetFileMetasResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetFileMetasResponseBody) SetStatus(v bool) *UpdateDatasetFileMetasResponseBody {
	s.Status = &v
	return s
}

func (s *UpdateDatasetFileMetasResponseBody) Validate() error {
	if s.FailedDetails != nil {
		for _, item := range s.FailedDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
