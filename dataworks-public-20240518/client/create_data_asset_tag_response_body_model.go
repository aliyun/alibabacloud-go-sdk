// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateDataAssetTagResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateDataAssetTagResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *CreateDataAssetTagResponseBody
	GetSuccess() *bool
}

type CreateDataAssetTagResponseBody struct {
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

func (s CreateDataAssetTagResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateDataAssetTagResponseBody) GoString() string {
	return s.String()
}

func (s *CreateDataAssetTagResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateDataAssetTagResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *CreateDataAssetTagResponseBody) SetRequestId(v string) *CreateDataAssetTagResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateDataAssetTagResponseBody) SetSuccess(v bool) *CreateDataAssetTagResponseBody {
	s.Success = &v
	return s
}

func (s *CreateDataAssetTagResponseBody) Validate() error {
	return dara.Validate(s)
}
