// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRunMaterializedViewRecommendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *RunMaterializedViewRecommendResponseBody
	GetRequestId() *string
}

type RunMaterializedViewRecommendResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A444291****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s RunMaterializedViewRecommendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RunMaterializedViewRecommendResponseBody) GoString() string {
	return s.String()
}

func (s *RunMaterializedViewRecommendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RunMaterializedViewRecommendResponseBody) SetRequestId(v string) *RunMaterializedViewRecommendResponseBody {
	s.RequestId = &v
	return s
}

func (s *RunMaterializedViewRecommendResponseBody) Validate() error {
	return dara.Validate(s)
}
