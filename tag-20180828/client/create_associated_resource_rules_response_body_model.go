// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAssociatedResourceRulesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateAssociatedResourceRulesResponseBody
	GetRequestId() *string
}

type CreateAssociatedResourceRulesResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 33BD6957-D7B0-500C-ADA1-300414EDCE89
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateAssociatedResourceRulesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAssociatedResourceRulesResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAssociatedResourceRulesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAssociatedResourceRulesResponseBody) SetRequestId(v string) *CreateAssociatedResourceRulesResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAssociatedResourceRulesResponseBody) Validate() error {
	return dara.Validate(s)
}
