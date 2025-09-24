// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatasetLabelsResponseBody
	GetRequestId() *string
}

type DeleteDatasetLabelsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 64B50C1D-D4C2-560C-86A3-A6ED****16D
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetLabelsResponseBody) SetRequestId(v string) *DeleteDatasetLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
