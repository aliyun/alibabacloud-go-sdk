// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateAlertActionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAlertActionId(v string) *CreateAlertActionResponseBody
	GetAlertActionId() *string
	SetRequestId(v string) *CreateAlertActionResponseBody
	GetRequestId() *string
}

type CreateAlertActionResponseBody struct {
	// example:
	//
	// testId
	AlertActionId *string `json:"alertActionId,omitempty" xml:"alertActionId,omitempty"`
	// example:
	//
	// 8FDE2569-626B-5176-9844-28877A*****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateAlertActionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateAlertActionResponseBody) GoString() string {
	return s.String()
}

func (s *CreateAlertActionResponseBody) GetAlertActionId() *string {
	return s.AlertActionId
}

func (s *CreateAlertActionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateAlertActionResponseBody) SetAlertActionId(v string) *CreateAlertActionResponseBody {
	s.AlertActionId = &v
	return s
}

func (s *CreateAlertActionResponseBody) SetRequestId(v string) *CreateAlertActionResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateAlertActionResponseBody) Validate() error {
	return dara.Validate(s)
}
