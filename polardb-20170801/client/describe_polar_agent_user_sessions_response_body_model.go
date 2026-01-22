// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePolarAgentUserSessionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DescribePolarAgentUserSessionsResponseBodyData) *DescribePolarAgentUserSessionsResponseBody
	GetData() []*DescribePolarAgentUserSessionsResponseBodyData
	SetRequestId(v string) *DescribePolarAgentUserSessionsResponseBody
	GetRequestId() *string
}

type DescribePolarAgentUserSessionsResponseBody struct {
	Data []*DescribePolarAgentUserSessionsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// example:
	//
	// E2FDB684-751D-424D-98B9-704BEA******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePolarAgentUserSessionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentUserSessionsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentUserSessionsResponseBody) GetData() []*DescribePolarAgentUserSessionsResponseBodyData {
	return s.Data
}

func (s *DescribePolarAgentUserSessionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePolarAgentUserSessionsResponseBody) SetData(v []*DescribePolarAgentUserSessionsResponseBodyData) *DescribePolarAgentUserSessionsResponseBody {
	s.Data = v
	return s
}

func (s *DescribePolarAgentUserSessionsResponseBody) SetRequestId(v string) *DescribePolarAgentUserSessionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePolarAgentUserSessionsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribePolarAgentUserSessionsResponseBodyData struct {
	// example:
	//
	// 01IC17MLISBB98SL345H7B5AES1E8VB1
	SessionId *string `json:"SessionId,omitempty" xml:"SessionId,omitempty"`
	// example:
	//
	// 1
	SessionStatus *int64 `json:"SessionStatus,omitempty" xml:"SessionStatus,omitempty"`
	// example:
	//
	// content
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s DescribePolarAgentUserSessionsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DescribePolarAgentUserSessionsResponseBodyData) GoString() string {
	return s.String()
}

func (s *DescribePolarAgentUserSessionsResponseBodyData) GetSessionId() *string {
	return s.SessionId
}

func (s *DescribePolarAgentUserSessionsResponseBodyData) GetSessionStatus() *int64 {
	return s.SessionStatus
}

func (s *DescribePolarAgentUserSessionsResponseBodyData) GetTitle() *string {
	return s.Title
}

func (s *DescribePolarAgentUserSessionsResponseBodyData) SetSessionId(v string) *DescribePolarAgentUserSessionsResponseBodyData {
	s.SessionId = &v
	return s
}

func (s *DescribePolarAgentUserSessionsResponseBodyData) SetSessionStatus(v int64) *DescribePolarAgentUserSessionsResponseBodyData {
	s.SessionStatus = &v
	return s
}

func (s *DescribePolarAgentUserSessionsResponseBodyData) SetTitle(v string) *DescribePolarAgentUserSessionsResponseBodyData {
	s.Title = &v
	return s
}

func (s *DescribePolarAgentUserSessionsResponseBodyData) Validate() error {
	return dara.Validate(s)
}
