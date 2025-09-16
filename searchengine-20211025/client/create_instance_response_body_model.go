// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateInstanceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *CreateInstanceResponseBody
	GetRequestId() *string
	SetResult(v *CreateInstanceResponseBodyResult) *CreateInstanceResponseBody
	GetResult() *CreateInstanceResponseBodyResult
}

type CreateInstanceResponseBody struct {
	// The ID of the request
	//
	// example:
	//
	// D39EE0F1-D7EF-5F46-B781-6BF4185308B0
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// The results returned.
	Result *CreateInstanceResponseBodyResult `json:"result,omitempty" xml:"result,omitempty" type:"Struct"`
}

func (s CreateInstanceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBody) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateInstanceResponseBody) GetResult() *CreateInstanceResponseBodyResult {
	return s.Result
}

func (s *CreateInstanceResponseBody) SetRequestId(v string) *CreateInstanceResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateInstanceResponseBody) SetResult(v *CreateInstanceResponseBodyResult) *CreateInstanceResponseBody {
	s.Result = v
	return s
}

func (s *CreateInstanceResponseBody) Validate() error {
	return dara.Validate(s)
}

type CreateInstanceResponseBodyResult struct {
	// The instance ID.
	//
	// example:
	//
	// ha-cn-2r42ppr7901
	InstanceId *string `json:"instanceId,omitempty" xml:"instanceId,omitempty"`
}

func (s CreateInstanceResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s CreateInstanceResponseBodyResult) GoString() string {
	return s.String()
}

func (s *CreateInstanceResponseBodyResult) GetInstanceId() *string {
	return s.InstanceId
}

func (s *CreateInstanceResponseBodyResult) SetInstanceId(v string) *CreateInstanceResponseBodyResult {
	s.InstanceId = &v
	return s
}

func (s *CreateInstanceResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
