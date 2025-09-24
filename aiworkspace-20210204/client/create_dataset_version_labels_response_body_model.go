// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDatasetVersionLabelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDatasetVersionLabelsResponseBody
	GetRequestId() *string
}

type CreateDatasetVersionLabelsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 8D7B2E70-F770-505B-A672-09F1D8F2EC1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateDatasetVersionLabelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDatasetVersionLabelsResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDatasetVersionLabelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDatasetVersionLabelsResponseBody) SetRequestId(v string) *CreateDatasetVersionLabelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDatasetVersionLabelsResponseBody) Validate() error {
	return dara.Validate(s)
}
