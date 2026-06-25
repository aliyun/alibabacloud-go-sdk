// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetLabelsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLabelKeys(v string) *DeleteDatasetLabelsRequest
	GetLabelKeys() *string
}

type DeleteDatasetLabelsRequest struct {
	// The keys of the labels. For more information about how to query the keys of labels, see [GetDataset](https://help.aliyun.com/document_detail/457218.html). Separate multiple keys with commas (,).
	//
	// example:
	//
	// key1,key2
	LabelKeys *string `json:"LabelKeys,omitempty" xml:"LabelKeys,omitempty"`
}

func (s DeleteDatasetLabelsRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetLabelsRequest) GoString() string {
	return s.String()
}

func (s *DeleteDatasetLabelsRequest) GetLabelKeys() *string {
	return s.LabelKeys
}

func (s *DeleteDatasetLabelsRequest) SetLabelKeys(v string) *DeleteDatasetLabelsRequest {
	s.LabelKeys = &v
	return s
}

func (s *DeleteDatasetLabelsRequest) Validate() error {
	return dara.Validate(s)
}
