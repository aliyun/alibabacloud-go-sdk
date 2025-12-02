// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyMaterializedViewRecommendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyMaterializedViewRecommendResponseBody
	GetRequestId() *string
}

type ModifyMaterializedViewRecommendResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// F0983B43-B2EC-536A-8791-142B5CF1E9B6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyMaterializedViewRecommendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyMaterializedViewRecommendResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyMaterializedViewRecommendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyMaterializedViewRecommendResponseBody) SetRequestId(v string) *ModifyMaterializedViewRecommendResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyMaterializedViewRecommendResponseBody) Validate() error {
	return dara.Validate(s)
}
