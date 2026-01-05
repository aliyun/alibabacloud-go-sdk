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
	SetSuccess(v bool) *DeleteDatasetVersionResponseBody
	GetSuccess() *bool
}

type DeleteDatasetVersionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 6AABBBD3-F2E4-5860-8CF7-2E9CEE3BDXXX
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the deletion succeeded.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
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

func (s *DeleteDatasetVersionResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDatasetVersionResponseBody) SetRequestId(v string) *DeleteDatasetVersionResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDatasetVersionResponseBody) SetSuccess(v bool) *DeleteDatasetVersionResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDatasetVersionResponseBody) Validate() error {
	return dara.Validate(s)
}
