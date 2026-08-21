// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSetZoneLbaStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *SetZoneLbaStatusResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *SetZoneLbaStatusResponseBody
	GetSuccess() *bool
}

type SetZoneLbaStatusResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 0B7AD377-7E86-44A8-B9A8-53E8666E72FE
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the operation is successful.
	//
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s SetZoneLbaStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s SetZoneLbaStatusResponseBody) GoString() string {
	return s.String()
}

func (s *SetZoneLbaStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *SetZoneLbaStatusResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *SetZoneLbaStatusResponseBody) SetRequestId(v string) *SetZoneLbaStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *SetZoneLbaStatusResponseBody) SetSuccess(v bool) *SetZoneLbaStatusResponseBody {
	s.Success = &v
	return s
}

func (s *SetZoneLbaStatusResponseBody) Validate() error {
	return dara.Validate(s)
}
