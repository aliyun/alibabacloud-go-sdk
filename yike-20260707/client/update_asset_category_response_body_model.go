// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateAssetCategoryResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *UpdateAssetCategoryResponseBody
	GetRequestId() *string
}

type UpdateAssetCategoryResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// ****63E8B7C7-4812-46AD-0FA56029AC86****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s UpdateAssetCategoryResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateAssetCategoryResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateAssetCategoryResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateAssetCategoryResponseBody) SetRequestId(v string) *UpdateAssetCategoryResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateAssetCategoryResponseBody) Validate() error {
	return dara.Validate(s)
}
