// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateRayHistoryServerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRayHistoryServerId(v string) *CreateRayHistoryServerResponseBody
	GetRayHistoryServerId() *string
	SetRequestId(v string) *CreateRayHistoryServerResponseBody
	GetRequestId() *string
}

type CreateRayHistoryServerResponseBody struct {
	// example:
	//
	// rhsxxxx
	RayHistoryServerId *string `json:"RayHistoryServerId,omitempty" xml:"RayHistoryServerId,omitempty"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-xxxxxxx
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s CreateRayHistoryServerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateRayHistoryServerResponseBody) GoString() string {
	return s.String()
}

func (s *CreateRayHistoryServerResponseBody) GetRayHistoryServerId() *string {
	return s.RayHistoryServerId
}

func (s *CreateRayHistoryServerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateRayHistoryServerResponseBody) SetRayHistoryServerId(v string) *CreateRayHistoryServerResponseBody {
	s.RayHistoryServerId = &v
	return s
}

func (s *CreateRayHistoryServerResponseBody) SetRequestId(v string) *CreateRayHistoryServerResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateRayHistoryServerResponseBody) Validate() error {
	return dara.Validate(s)
}
