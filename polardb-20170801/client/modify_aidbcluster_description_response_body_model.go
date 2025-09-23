// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyAIDBClusterDescriptionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ModifyAIDBClusterDescriptionResponseBody
	GetRequestId() *string
}

type ModifyAIDBClusterDescriptionResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ModifyAIDBClusterDescriptionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyAIDBClusterDescriptionResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyAIDBClusterDescriptionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyAIDBClusterDescriptionResponseBody) SetRequestId(v string) *ModifyAIDBClusterDescriptionResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyAIDBClusterDescriptionResponseBody) Validate() error {
	return dara.Validate(s)
}
