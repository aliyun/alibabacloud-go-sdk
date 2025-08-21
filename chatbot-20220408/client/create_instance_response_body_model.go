// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceId(v string) *CreateInstanceResponseBody
	GetInstanceId() *string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
}

type CreateInstanceResponseBody struct {
	// example:
	//
	// chatbot-cn-mp90s2lrk00050
	InstanceId *string `json:"InstanceId,omitempty" xml:"InstanceId,omitempty"`
	// Id of the request
	//
	// example:
	//
	// F2E2C40D-AB09-45A1-B5C5-EB9F5C4E4E4A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) SetInstanceId(v string) *CreateInstanceResponseBody {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}
