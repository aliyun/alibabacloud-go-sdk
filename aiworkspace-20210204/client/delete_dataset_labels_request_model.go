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
	// The tag key. You can call [GetDataset](https://help.aliyun.com/document_detail/457218.html) to obtain the tag key. Multiple tag keys are separated by commas (,).
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
