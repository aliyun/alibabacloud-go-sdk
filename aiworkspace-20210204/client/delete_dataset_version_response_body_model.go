// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetVersionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatasetVersionResponseBody
	GetRequestId() *string
}

type DeleteDatasetVersionResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 8D7B2E70-F770-505B-A672-09F1D8F2EC1E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteDatasetVersionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetVersionResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetVersionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetVersionResponseBody) SetRequestId(v string) *DeleteDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
