// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetFileMetasResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFailedDetails(v []*DatasetFileMetaResponse) *CreateDatasetFileMetasResponseBody
	GetFailedDetails() []*DatasetFileMetaResponse
	SetRequestId(v string) *CreateDatasetFileMetasResponseBody
	GetRequestId() *string
	SetStatus(v bool) *CreateDatasetFileMetasResponseBody
	GetStatus() *bool
	SetSucceedDetails(v []*DatasetFileMetaResponse) *CreateDatasetFileMetasResponseBody
	GetSucceedDetails() []*DatasetFileMetaResponse
}

type CreateDatasetFileMetasResponseBody struct {
	// The metadata that failed to be created.
	FailedDetails []*DatasetFileMetaResponse `json:"FailedDetails,omitempty" xml:"FailedDetails,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 5A14FA81-DD4E-******-6343FE44B941
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the metadata records of all dataset files were created. The value true indicates that the metadata records of all dataset files are created. If the value is false, view the failure details specified by FailedDetails.
	//
	// Valid values:
	//
	// 	- true
	//
	// 	- false
	//
	// example:
	//
	// true
	Status *bool `json:"Status,omitempty" xml:"Status,omitempty"`
	// The metadata that is created.
	SucceedDetails []*DatasetFileMetaResponse `json:"SucceedDetails,omitempty" xml:"SucceedDetails,omitempty" type:"Repeated"`
}

func (s CreateDatasetFileMetasResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetFileMetasResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetFileMetasResponseBody) GetFailedDetails() []*DatasetFileMetaResponse {
	return s.FailedDetails
}

func (s *CreateDatasetFileMetasResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetFileMetasResponseBody) GetStatus() *bool {
	return s.Status
}

func (s *CreateDatasetFileMetasResponseBody) GetSucceedDetails() []*DatasetFileMetaResponse {
	return s.SucceedDetails
}

func (s *CreateDatasetFileMetasResponseBody) SetFailedDetails(v []*DatasetFileMetaResponse) *CreateDatasetFileMetasResponseBody {
	s.FailedDetails = v
	return s
}

func (s *CreateDatasetFileMetasResponseBody) SetRequestId(v string) *CreateDatasetFileMetasResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetFileMetasResponseBody) SetStatus(v bool) *CreateDatasetFileMetasResponseBody {
	s.Status = &v
	return s
}

func (s *CreateDatasetFileMetasResponseBody) SetSucceedDetails(v []*DatasetFileMetaResponse) *CreateDatasetFileMetasResponseBody {
	s.SucceedDetails = v
	return s
}

func (s *CreateDatasetFileMetasResponseBody) Validate() error {
	if s.FailedDetails != nil {
		for _, item := range s.FailedDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.SucceedDetails != nil {
		for _, item := range s.SucceedDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
