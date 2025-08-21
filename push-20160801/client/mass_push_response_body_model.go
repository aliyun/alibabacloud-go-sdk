// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMassPushResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMessageIds(v *MassPushResponseBodyMessageIds) *MassPushResponseBody
	GetMessageIds() *MassPushResponseBodyMessageIds
	SetRequestId(v string) *MassPushResponseBody
	GetRequestId() *string
}

type MassPushResponseBody struct {
	MessageIds *MassPushResponseBodyMessageIds `json:"MessageIds,omitempty" xml:"MessageIds,omitempty" type:"Struct"`
	// example:
	//
	// 9998B3CC-ED9E-4CB3-A8FB-DCC61296BFBC
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s MassPushResponseBody) String() string {
	return dara.Prettify(s)
}

func (s MassPushResponseBody) GoString() string {
	return s.String()
}

func (s *MassPushResponseBody) GetMessageIds() *MassPushResponseBodyMessageIds {
	return s.MessageIds
}

func (s *MassPushResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *MassPushResponseBody) SetMessageIds(v *MassPushResponseBodyMessageIds) *MassPushResponseBody {
	s.MessageIds = v
	return s
}

func (s *MassPushResponseBody) SetRequestId(v string) *MassPushResponseBody {
	s.RequestId = &v
	return s
}

func (s *MassPushResponseBody) Validate() error {
	return dara.Validate(s)
}

type MassPushResponseBodyMessageIds struct {
	MessageId []*string `json:"MessageId,omitempty" xml:"MessageId,omitempty" type:"Repeated"`
}

func (s MassPushResponseBodyMessageIds) String() string {
	return dara.Prettify(s)
}

func (s MassPushResponseBodyMessageIds) GoString() string {
	return s.String()
}

func (s *MassPushResponseBodyMessageIds) GetMessageId() []*string {
	return s.MessageId
}

func (s *MassPushResponseBodyMessageIds) SetMessageId(v []*string) *MassPushResponseBodyMessageIds {
	s.MessageId = v
	return s
}

func (s *MassPushResponseBodyMessageIds) Validate() error {
	return dara.Validate(s)
}
