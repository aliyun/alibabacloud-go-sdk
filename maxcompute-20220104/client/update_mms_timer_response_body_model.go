// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMmsTimerResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v int64) *UpdateMmsTimerResponseBody
	GetData() *int64
	SetRequestId(v string) *UpdateMmsTimerResponseBody
	GetRequestId() *string
}

type UpdateMmsTimerResponseBody struct {
	// example:
	//
	// success
	Data *int64 `json:"data,omitempty" xml:"data,omitempty"`
	// example:
	//
	// EA1320AB-7766-5EC7-B0F6-8B20E2298567
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s UpdateMmsTimerResponseBody) String() string {
	return dara.Prettify(s)
}

func (s UpdateMmsTimerResponseBody) GoString() string {
	return s.String()
}

func (s *UpdateMmsTimerResponseBody) GetData() *int64 {
	return s.Data
}

func (s *UpdateMmsTimerResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *UpdateMmsTimerResponseBody) SetData(v int64) *UpdateMmsTimerResponseBody {
	s.Data = &v
	return s
}

func (s *UpdateMmsTimerResponseBody) SetRequestId(v string) *UpdateMmsTimerResponseBody {
	s.RequestId = &v
	return s
}

func (s *UpdateMmsTimerResponseBody) Validate() error {
	return dara.Validate(s)
}
