// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateServiceRolloutResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessage(v string) *CreateServiceRolloutResponseBody
	GetMessage() *string
	SetRequestId(v string) *CreateServiceRolloutResponseBody
	GetRequestId() *string
}

type CreateServiceRolloutResponseBody struct {
	// A message that indicates the result of the operation.
	//
	// example:
	//
	// Rollout created successfully
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique request ID. Use this ID for troubleshooting and traceability.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateServiceRolloutResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateServiceRolloutResponseBody) GoString() string {
	return s.String()
}

func (s *CreateServiceRolloutResponseBody) GetMessage() *string {
	return s.Message
}

func (s *CreateServiceRolloutResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateServiceRolloutResponseBody) SetMessage(v string) *CreateServiceRolloutResponseBody {
	s.Message = &v
	return s
}

func (s *CreateServiceRolloutResponseBody) SetRequestId(v string) *CreateServiceRolloutResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateServiceRolloutResponseBody) Validate() error {
	return dara.Validate(s)
}
