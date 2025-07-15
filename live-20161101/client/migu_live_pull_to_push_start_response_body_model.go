// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMiguLivePullToPushStartResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MiguLivePullToPushStartResponseBody
	GetCode() *string
	SetMessage(v string) *MiguLivePullToPushStartResponseBody
	GetMessage() *string
	SetRequestId(v string) *MiguLivePullToPushStartResponseBody
	GetRequestId() *string
	SetTimestamp(v string) *MiguLivePullToPushStartResponseBody
	GetTimestamp() *string
}

type MiguLivePullToPushStartResponseBody struct {
	Code      *string `json:"code,omitempty" xml:"code,omitempty"`
	Message   *string `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Timestamp *string `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s MiguLivePullToPushStartResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MiguLivePullToPushStartResponseBody) GoString() string {
	return s.String()
}

func (s *MiguLivePullToPushStartResponseBody) GetCode() *string {
	return s.Code
}

func (s *MiguLivePullToPushStartResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MiguLivePullToPushStartResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MiguLivePullToPushStartResponseBody) GetTimestamp() *string {
	return s.Timestamp
}

func (s *MiguLivePullToPushStartResponseBody) SetCode(v string) *MiguLivePullToPushStartResponseBody {
	s.Code = &v
	return s
}

func (s *MiguLivePullToPushStartResponseBody) SetMessage(v string) *MiguLivePullToPushStartResponseBody {
	s.Message = &v
	return s
}

func (s *MiguLivePullToPushStartResponseBody) SetRequestId(v string) *MiguLivePullToPushStartResponseBody {
	s.RequestId = &v
	return s
}

func (s *MiguLivePullToPushStartResponseBody) SetTimestamp(v string) *MiguLivePullToPushStartResponseBody {
	s.Timestamp = &v
	return s
}

func (s *MiguLivePullToPushStartResponseBody) Validate() error {
	return dara.Validate(s)
}
