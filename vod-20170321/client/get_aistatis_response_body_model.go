// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIStatisList(v []*GetAIStatisResponseBodyAIStatisList) *GetAIStatisResponseBody
	GetAIStatisList() []*GetAIStatisResponseBodyAIStatisList
	SetRequestId(v string) *GetAIStatisResponseBody
	GetRequestId() *string
}

type GetAIStatisResponseBody struct {
	AIStatisList []*GetAIStatisResponseBodyAIStatisList `json:"AIStatisList,omitempty" xml:"AIStatisList,omitempty" type:"Repeated"`
	RequestId    *string                                `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAIStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIStatisResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIStatisResponseBody) GetAIStatisList() []*GetAIStatisResponseBodyAIStatisList {
	return s.AIStatisList
}

func (s *GetAIStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIStatisResponseBody) SetAIStatisList(v []*GetAIStatisResponseBodyAIStatisList) *GetAIStatisResponseBody {
	s.AIStatisList = v
	return s
}

func (s *GetAIStatisResponseBody) SetRequestId(v string) *GetAIStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIStatisResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAIStatisResponseBodyAIStatisList struct {
	Duration    *int64  `json:"Duration,omitempty" xml:"Duration,omitempty"`
	StatTime    *string `json:"StatTime,omitempty" xml:"StatTime,omitempty"`
	StatTimeUTC *string `json:"StatTimeUTC,omitempty" xml:"StatTimeUTC,omitempty"`
}

func (s GetAIStatisResponseBodyAIStatisList) String() string {
	return dara.Prettify(s)
}

func (s GetAIStatisResponseBodyAIStatisList) GoString() string {
	return s.String()
}

func (s *GetAIStatisResponseBodyAIStatisList) GetDuration() *int64 {
	return s.Duration
}

func (s *GetAIStatisResponseBodyAIStatisList) GetStatTime() *string {
	return s.StatTime
}

func (s *GetAIStatisResponseBodyAIStatisList) GetStatTimeUTC() *string {
	return s.StatTimeUTC
}

func (s *GetAIStatisResponseBodyAIStatisList) SetDuration(v int64) *GetAIStatisResponseBodyAIStatisList {
	s.Duration = &v
	return s
}

func (s *GetAIStatisResponseBodyAIStatisList) SetStatTime(v string) *GetAIStatisResponseBodyAIStatisList {
	s.StatTime = &v
	return s
}

func (s *GetAIStatisResponseBodyAIStatisList) SetStatTimeUTC(v string) *GetAIStatisResponseBodyAIStatisList {
	s.StatTimeUTC = &v
	return s
}

func (s *GetAIStatisResponseBodyAIStatisList) Validate() error {
	return dara.Validate(s)
}
