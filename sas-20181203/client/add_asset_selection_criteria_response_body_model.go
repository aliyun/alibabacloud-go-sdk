// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iAddAssetSelectionCriteriaResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *AddAssetSelectionCriteriaResponseBody
	GetRequestId() *string
}

type AddAssetSelectionCriteriaResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 7BC55C8F-226E-5AF5-9A2C-2EC43864****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s AddAssetSelectionCriteriaResponseBody) String() string {
	return dara.Prettify(s)
}

func (s AddAssetSelectionCriteriaResponseBody) GoString() string {
	return s.String()
}

func (s *AddAssetSelectionCriteriaResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *AddAssetSelectionCriteriaResponseBody) SetRequestId(v string) *AddAssetSelectionCriteriaResponseBody {
	s.RequestId = &v
	return s
}

func (s *AddAssetSelectionCriteriaResponseBody) Validate() error {
	return dara.Validate(s)
}
