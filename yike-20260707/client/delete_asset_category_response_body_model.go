// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAssetCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteAssetCategoryResponseBody
	GetRequestId() *string
}

type DeleteAssetCategoryResponseBody struct {
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteAssetCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteAssetCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteAssetCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteAssetCategoryResponseBody) SetRequestId(v string) *DeleteAssetCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteAssetCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
