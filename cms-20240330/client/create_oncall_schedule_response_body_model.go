// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateOncallScheduleResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetOncallScheduleId(v string) *CreateOncallScheduleResponseBody
	GetOncallScheduleId() *string
	SetRequestId(v string) *CreateOncallScheduleResponseBody
	GetRequestId() *string
}

type CreateOncallScheduleResponseBody struct {
	// example:
	//
	// testOncallScheduleId
	OncallScheduleId *string `json:"oncallScheduleId,omitempty" xml:"oncallScheduleId,omitempty"`
	// example:
	//
	// 3B311FD9-A60B-55E0-A896-A0C73*********
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s CreateOncallScheduleResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateOncallScheduleResponseBody) GoString() string {
	return s.String()
}

func (s *CreateOncallScheduleResponseBody) GetOncallScheduleId() *string {
	return s.OncallScheduleId
}

func (s *CreateOncallScheduleResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateOncallScheduleResponseBody) SetOncallScheduleId(v string) *CreateOncallScheduleResponseBody {
	s.OncallScheduleId = &v
	return s
}

func (s *CreateOncallScheduleResponseBody) SetRequestId(v string) *CreateOncallScheduleResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateOncallScheduleResponseBody) Validate() error {
	return dara.Validate(s)
}
