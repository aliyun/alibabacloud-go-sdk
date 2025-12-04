// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataset(v *Dataset) *GetDatasetResponseBody
	GetDataset() *Dataset
	SetRequestId(v string) *GetDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetDatasetResponseBody
	GetSuccess() *bool
}

type GetDatasetResponseBody struct {
	Dataset *Dataset `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 204EAF68-CCE3-5112-8DA0-E7A60F02XXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *GetDatasetResponseBody) GetDataset() *Dataset {
	return s.Dataset
}

func (s *GetDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetDatasetResponseBody) SetDataset(v *Dataset) *GetDatasetResponseBody {
	s.Dataset = v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetResponseBody) SetSuccess(v bool) *GetDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *GetDatasetResponseBody) Validate() error {
	if s.Dataset != nil {
		if err := s.Dataset.Validate(); err != nil {
			return err
		}
	}
	return nil
}
