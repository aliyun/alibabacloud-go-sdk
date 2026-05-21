// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDisableOperationEventResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v bool) *DisableOperationEventResponseBody
	GetData() *bool
	SetRequestId(v string) *DisableOperationEventResponseBody
	GetRequestId() *string
}

type DisableOperationEventResponseBody struct {
	// example:
	//
	// true
	Data *bool `json:"Data,omitempty" xml:"Data,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D3AE84AB-0873-5FC7-A4C4-8CF869D2FA70
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DisableOperationEventResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DisableOperationEventResponseBody) GoString() string {
	return s.String()
}

func (s *DisableOperationEventResponseBody) GetData() *bool {
	return s.Data
}

func (s *DisableOperationEventResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DisableOperationEventResponseBody) SetData(v bool) *DisableOperationEventResponseBody {
	s.Data = &v
	return s
}

func (s *DisableOperationEventResponseBody) SetRequestId(v string) *DisableOperationEventResponseBody {
	s.RequestId = &v
	return s
}

func (s *DisableOperationEventResponseBody) Validate() error {
	return dara.Validate(s)
}
