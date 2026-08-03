// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iStopAvatarSessionResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *StopAvatarSessionResponseBody
	GetRequestId() *string
	SetStatus(v string) *StopAvatarSessionResponseBody
	GetStatus() *string
}

type StopAvatarSessionResponseBody struct {
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Status    *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s StopAvatarSessionResponseBody) String() string {
	return dara.Prettify(s)
}

func (s StopAvatarSessionResponseBody) GoString() string {
	return s.String()
}

func (s *StopAvatarSessionResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *StopAvatarSessionResponseBody) GetStatus() *string {
	return s.Status
}

func (s *StopAvatarSessionResponseBody) SetRequestId(v string) *StopAvatarSessionResponseBody {
	s.RequestId = &v
	return s
}

func (s *StopAvatarSessionResponseBody) SetStatus(v string) *StopAvatarSessionResponseBody {
	s.Status = &v
	return s
}

func (s *StopAvatarSessionResponseBody) Validate() error {
	return dara.Validate(s)
}
