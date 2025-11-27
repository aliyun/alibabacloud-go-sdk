// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataset(v *Dataset) *CreateDatasetResponseBody
	GetDataset() *Dataset
	SetRequestId(v string) *CreateDatasetResponseBody
	GetRequestId() *string
}

type CreateDatasetResponseBody struct {
	// Dataset information.
	Dataset *Dataset `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// Request ID.
	//
	// example:
	//
	// 6D74B3A9-5AC0-49F9-914D-E01589D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetResponseBody) GetDataset() *Dataset {
	return s.Dataset
}

func (s *CreateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetResponseBody) SetDataset(v *Dataset) *CreateDatasetResponseBody {
	s.Dataset = v
	return s
}

func (s *CreateDatasetResponseBody) SetRequestId(v string) *CreateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetResponseBody) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}
