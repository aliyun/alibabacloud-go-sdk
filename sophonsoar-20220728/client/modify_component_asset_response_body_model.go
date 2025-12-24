// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyComponentAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyComponentAssetResponseBody
	GetRequestId() *string
}

type ModifyComponentAssetResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1C5F11E9-464E-51F0-9296-43BB312A0557
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyComponentAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyComponentAssetResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyComponentAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyComponentAssetResponseBody) SetRequestId(v string) *ModifyComponentAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyComponentAssetResponseBody) Validate() error {
	return dara.Validate(s)
}
