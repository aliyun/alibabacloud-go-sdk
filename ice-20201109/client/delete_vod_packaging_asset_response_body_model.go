// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteVodPackagingAssetResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteVodPackagingAssetResponseBody
	GetRequestId() *string
}

type DeleteVodPackagingAssetResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// ******11-DB8D-4A9A-875B-275798******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteVodPackagingAssetResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteVodPackagingAssetResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteVodPackagingAssetResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteVodPackagingAssetResponseBody) SetRequestId(v string) *DeleteVodPackagingAssetResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteVodPackagingAssetResponseBody) Validate() error {
	return dara.Validate(s)
}
