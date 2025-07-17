// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateDataAssetTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateDataAssetTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *UpdateDataAssetTagResponseBody
	GetSuccess() *bool
}

type UpdateDataAssetTagResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0bc1ec92159376
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request was successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s UpdateDataAssetTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateDataAssetTagResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateDataAssetTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateDataAssetTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *UpdateDataAssetTagResponseBody) SetRequestId(v string) *UpdateDataAssetTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateDataAssetTagResponseBody) SetSuccess(v bool) *UpdateDataAssetTagResponseBody {
	s.Success = &v
	return s
}

func (s *UpdateDataAssetTagResponseBody) Validate() error {
	return dara.Validate(s)
}
