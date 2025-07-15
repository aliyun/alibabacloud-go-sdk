// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMiguLivePullToPushStatusResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *MiguLivePullToPushStatusResponseBody
	GetCode() *string
	SetData(v *MiguLivePullToPushStatusResponseBodyData) *MiguLivePullToPushStatusResponseBody
	GetData() *MiguLivePullToPushStatusResponseBodyData
	SetMessage(v string) *MiguLivePullToPushStatusResponseBody
	GetMessage() *string
	SetRequestId(v string) *MiguLivePullToPushStatusResponseBody
	GetRequestId() *string
	SetTimestamp(v string) *MiguLivePullToPushStatusResponseBody
	GetTimestamp() *string
}

type MiguLivePullToPushStatusResponseBody struct {
	Code      *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data      *MiguLivePullToPushStatusResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	Message   *string                                   `json:"message,omitempty" xml:"message,omitempty"`
	RequestId *string                                   `json:"requestId,omitempty" xml:"requestId,omitempty"`
	Timestamp *string                                   `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
}

func (s MiguLivePullToPushStatusResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MiguLivePullToPushStatusResponseBody) GoString() string {
	return s.String()
}

func (s *MiguLivePullToPushStatusResponseBody) GetCode() *string {
	return s.Code
}

func (s *MiguLivePullToPushStatusResponseBody) GetData() *MiguLivePullToPushStatusResponseBodyData {
	return s.Data
}

func (s *MiguLivePullToPushStatusResponseBody) GetMessage() *string {
	return s.Message
}

func (s *MiguLivePullToPushStatusResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MiguLivePullToPushStatusResponseBody) GetTimestamp() *string {
	return s.Timestamp
}

func (s *MiguLivePullToPushStatusResponseBody) SetCode(v string) *MiguLivePullToPushStatusResponseBody {
	s.Code = &v
	return s
}

func (s *MiguLivePullToPushStatusResponseBody) SetData(v *MiguLivePullToPushStatusResponseBodyData) *MiguLivePullToPushStatusResponseBody {
	s.Data = v
	return s
}

func (s *MiguLivePullToPushStatusResponseBody) SetMessage(v string) *MiguLivePullToPushStatusResponseBody {
	s.Message = &v
	return s
}

func (s *MiguLivePullToPushStatusResponseBody) SetRequestId(v string) *MiguLivePullToPushStatusResponseBody {
	s.RequestId = &v
	return s
}

func (s *MiguLivePullToPushStatusResponseBody) SetTimestamp(v string) *MiguLivePullToPushStatusResponseBody {
	s.Timestamp = &v
	return s
}

func (s *MiguLivePullToPushStatusResponseBody) Validate() error {
	return dara.Validate(s)
}

type MiguLivePullToPushStatusResponseBodyData struct {
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	Status  *string `json:"status,omitempty" xml:"status,omitempty"`
}

func (s MiguLivePullToPushStatusResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s MiguLivePullToPushStatusResponseBodyData) GoString() string {
	return s.String()
}

func (s *MiguLivePullToPushStatusResponseBodyData) GetMessage() *string {
	return s.Message
}

func (s *MiguLivePullToPushStatusResponseBodyData) GetStatus() *string {
	return s.Status
}

func (s *MiguLivePullToPushStatusResponseBodyData) SetMessage(v string) *MiguLivePullToPushStatusResponseBodyData {
	s.Message = &v
	return s
}

func (s *MiguLivePullToPushStatusResponseBodyData) SetStatus(v string) *MiguLivePullToPushStatusResponseBodyData {
	s.Status = &v
	return s
}

func (s *MiguLivePullToPushStatusResponseBodyData) Validate() error {
	return dara.Validate(s)
}
