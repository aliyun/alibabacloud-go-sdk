// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDatasetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDatasetResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDatasetResponseBody
	GetSuccess() *bool
}

type DeleteDatasetResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// AAC30B35-820D-5F3E-A42C-E96BB****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the deletion succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDatasetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDatasetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDatasetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDatasetResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDatasetResponseBody) SetRequestId(v string) *DeleteDatasetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetResponseBody) SetSuccess(v bool) *DeleteDatasetResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDatasetResponseBody) Validate() error {
	return dara.Validate(s)
}
