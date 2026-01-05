// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDatasetVersion(v *DatasetVersion) *GetDatasetVersionResponseBody
	GetDatasetVersion() *DatasetVersion
	SetRequestId(v string) *GetDatasetVersionResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatasetVersionResponseBody
	GetSuccess() *bool
}

type GetDatasetVersionResponseBody struct {
	// The dataset version.
	DatasetVersion *DatasetVersion `json:"DatasetVersion,omitempty" xml:"DatasetVersion,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 4CDF7B72-020B-542A-8465-21CFFA8XXXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDatasetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetVersionResponseBody) GetDatasetVersion() *DatasetVersion {
	return s.DatasetVersion
}

func (s *GetDatasetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatasetVersionResponseBody) SetDatasetVersion(v *DatasetVersion) *GetDatasetVersionResponseBody {
	s.DatasetVersion = v
	return s
}

func (s *GetDatasetVersionResponseBody) SetRequestId(v string) *GetDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetVersionResponseBody) SetSuccess(v bool) *GetDatasetVersionResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatasetVersionResponseBody) Validate() error {
	if s.DatasetVersion != nil {
		if err := s.DatasetVersion.Validate(); err != nil {
			return err
		}
	}
	return nil
}
