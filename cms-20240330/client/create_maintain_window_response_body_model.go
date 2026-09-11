// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateMaintainWindowResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaintainWindowId(v string) *CreateMaintainWindowResponseBody
	GetMaintainWindowId() *string
	SetRequestId(v string) *CreateMaintainWindowResponseBody
	GetRequestId() *string
}

type CreateMaintainWindowResponseBody struct {
	MaintainWindowId *string `json:"maintainWindowId,omitempty" xml:"maintainWindowId,omitempty"`
	RequestId        *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateMaintainWindowResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateMaintainWindowResponseBody) GoString() string {
	return s.String()
}

func (s *CreateMaintainWindowResponseBody) GetMaintainWindowId() *string {
	return s.MaintainWindowId
}

func (s *CreateMaintainWindowResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateMaintainWindowResponseBody) SetMaintainWindowId(v string) *CreateMaintainWindowResponseBody {
	s.MaintainWindowId = &v
	return s
}

func (s *CreateMaintainWindowResponseBody) SetRequestId(v string) *CreateMaintainWindowResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateMaintainWindowResponseBody) Validate() error {
	return dara.Validate(s)
}
