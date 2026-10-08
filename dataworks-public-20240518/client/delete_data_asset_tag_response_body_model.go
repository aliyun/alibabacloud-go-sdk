// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteDataAssetTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteDataAssetTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *DeleteDataAssetTagResponseBody
	GetSuccess() *bool
}

type DeleteDataAssetTagResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bc1411515937635973****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s DeleteDataAssetTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteDataAssetTagResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteDataAssetTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteDataAssetTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *DeleteDataAssetTagResponseBody) SetRequestId(v string) *DeleteDataAssetTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteDataAssetTagResponseBody) SetSuccess(v bool) *DeleteDataAssetTagResponseBody {
	s.Success = &v
	return s
}

func (s *DeleteDataAssetTagResponseBody) Validate() error {
	return dara.Validate(s)
}
