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
}

type GetDatasetResponseBody struct {
	// The dataset.
	Dataset *Dataset `json:"Dataset,omitempty" xml:"Dataset,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 6D74B3A9-5AC0-49F9-914D-E01589D3****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
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

func (s *GetDatasetResponseBody) SetDataset(v *Dataset) *GetDatasetResponseBody {
	s.Dataset = v
	return s
}

func (s *GetDatasetResponseBody) SetRequestId(v string) *GetDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
