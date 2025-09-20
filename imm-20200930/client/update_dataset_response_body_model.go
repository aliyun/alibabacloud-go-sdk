// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataset(v *Dataset) *UpdateDatasetResponseBody
	GetDataset() *Dataset
	SetRequestId(v string) *UpdateDatasetResponseBody
	GetRequestId() *string
}

type UpdateDatasetResponseBody struct {
	// The dataset.
	Dataset *Dataset `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// example:
	//
	// 45234D4A-A3E3-4B23-AACA-8D897514****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDatasetResponseBody) GetDataset() *Dataset {
	return s.Dataset
}

func (s *UpdateDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDatasetResponseBody) SetDataset(v *Dataset) *UpdateDatasetResponseBody {
	s.Dataset = v
	return s
}

func (s *UpdateDatasetResponseBody) SetRequestId(v string) *UpdateDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
