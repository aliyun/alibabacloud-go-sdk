// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetVersionLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatasetVersionLabelsResponseBody
	GetRequestId() *string
}

type DeleteDatasetVersionLabelsResponseBody struct {
	// Id of the request
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetVersionLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetVersionLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetVersionLabelsResponseBody) SetRequestId(v string) *DeleteDatasetVersionLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetVersionLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
