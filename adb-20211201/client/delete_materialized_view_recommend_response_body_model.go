// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteMaterializedViewRecommendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DeleteMaterializedViewRecommendResponseBody
	GetRequestId() *string
}

type DeleteMaterializedViewRecommendResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 7B081A85-0568-5E54-82EF-6958C4A3D4B4
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DeleteMaterializedViewRecommendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DeleteMaterializedViewRecommendResponseBody) GoString() string {
	return s.String()
}

func (s *DeleteMaterializedViewRecommendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DeleteMaterializedViewRecommendResponseBody) SetRequestId(v string) *DeleteMaterializedViewRecommendResponseBody {
	s.RequestId = &v
	return s
}

func (s *DeleteMaterializedViewRecommendResponseBody) Validate() error {
	return dara.Validate(s)
}
