// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateApplicationAgentRelationResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetApplicationId(v string) *CreateApplicationAgentRelationResponseBody
	GetApplicationId() *string
	SetRequestId(v string) *CreateApplicationAgentRelationResponseBody
	GetRequestId() *string
}

type CreateApplicationAgentRelationResponseBody struct {
	// The ID of the Squad application.
	//
	// example:
	//
	// pa-xxx
	ApplicationId *string `json:"ApplicationId,omitempty" xml:"ApplicationId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// CD3FA5F3-FAF3-44CA-AFFF-BAF869******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateApplicationAgentRelationResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateApplicationAgentRelationResponseBody) GoString() string {
	return s.String()
}

func (s *CreateApplicationAgentRelationResponseBody) GetApplicationId() *string {
	return s.ApplicationId
}

func (s *CreateApplicationAgentRelationResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateApplicationAgentRelationResponseBody) SetApplicationId(v string) *CreateApplicationAgentRelationResponseBody {
	s.ApplicationId = &v
	return s
}

func (s *CreateApplicationAgentRelationResponseBody) SetRequestId(v string) *CreateApplicationAgentRelationResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateApplicationAgentRelationResponseBody) Validate() error {
	return dara.Validate(s)
}
