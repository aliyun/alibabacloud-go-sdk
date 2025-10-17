// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaterializedViewRecommendResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateMaterializedViewRecommendResponseBody
	GetRequestId() *string
}

type CreateMaterializedViewRecommendResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// 1AD222E9-E606-4A42-BF6D-8A4442913CEF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateMaterializedViewRecommendResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMaterializedViewRecommendResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMaterializedViewRecommendResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMaterializedViewRecommendResponseBody) SetRequestId(v string) *CreateMaterializedViewRecommendResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMaterializedViewRecommendResponseBody) Validate() error {
	return dara.Validate(s)
}
