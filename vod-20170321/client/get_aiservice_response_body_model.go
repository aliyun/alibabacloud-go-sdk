// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAIServiceResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIList(v []*GetAIServiceResponseBodyAIList) *GetAIServiceResponseBody
	GetAIList() []*GetAIServiceResponseBodyAIList
	SetRequestId(v string) *GetAIServiceResponseBody
	GetRequestId() *string
}

type GetAIServiceResponseBody struct {
	AIList    []*GetAIServiceResponseBodyAIList `json:"AIList,omitempty" xml:"AIList,omitempty" type:"Repeated"`
	RequestId *string                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetAIServiceResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetAIServiceResponseBody) GoString() string {
	return s.String()
}

func (s *GetAIServiceResponseBody) GetAIList() []*GetAIServiceResponseBodyAIList {
	return s.AIList
}

func (s *GetAIServiceResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetAIServiceResponseBody) SetAIList(v []*GetAIServiceResponseBodyAIList) *GetAIServiceResponseBody {
	s.AIList = v
	return s
}

func (s *GetAIServiceResponseBody) SetRequestId(v string) *GetAIServiceResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetAIServiceResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetAIServiceResponseBodyAIList struct {
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	Type   *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetAIServiceResponseBodyAIList) String() string {
	return dara.Prettify(s)
}

func (s GetAIServiceResponseBodyAIList) GoString() string {
	return s.String()
}

func (s *GetAIServiceResponseBodyAIList) GetStatus() *string {
	return s.Status
}

func (s *GetAIServiceResponseBodyAIList) GetType() *string {
	return s.Type
}

func (s *GetAIServiceResponseBodyAIList) SetStatus(v string) *GetAIServiceResponseBodyAIList {
	s.Status = &v
	return s
}

func (s *GetAIServiceResponseBodyAIList) SetType(v string) *GetAIServiceResponseBodyAIList {
	s.Type = &v
	return s
}

func (s *GetAIServiceResponseBodyAIList) Validate() error {
	return dara.Validate(s)
}
