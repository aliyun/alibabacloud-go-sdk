// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAgentsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListAgentsResponseBody
	GetCode() *string
	SetData(v *ListAgentsResponseBodyData) *ListAgentsResponseBody
	GetData() *ListAgentsResponseBodyData
	SetMessage(v string) *ListAgentsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListAgentsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *ListAgentsResponseBody
	GetSuccess() *bool
}

type ListAgentsResponseBody struct {
	// example:
	//
	// Success
	Code *string                     `json:"Code,omitempty" xml:"Code,omitempty"`
	Data *ListAgentsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// Failed to list agents
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// example:
	//
	// 7DA60DED-CD36-5837-B848-C01A23D2****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// true
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s ListAgentsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBody) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListAgentsResponseBody) GetData() *ListAgentsResponseBodyData {
	return s.Data
}

func (s *ListAgentsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListAgentsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAgentsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *ListAgentsResponseBody) SetCode(v string) *ListAgentsResponseBody {
	s.Code = &v
	return s
}

func (s *ListAgentsResponseBody) SetData(v *ListAgentsResponseBodyData) *ListAgentsResponseBody {
	s.Data = v
	return s
}

func (s *ListAgentsResponseBody) SetMessage(v string) *ListAgentsResponseBody {
	s.Message = &v
	return s
}

func (s *ListAgentsResponseBody) SetRequestId(v string) *ListAgentsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAgentsResponseBody) SetSuccess(v bool) *ListAgentsResponseBody {
	s.Success = &v
	return s
}

func (s *ListAgentsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListAgentsResponseBodyData struct {
	Agents []*Agent `json:"Agents,omitempty" xml:"Agents,omitempty" type:"Repeated"`
	// example:
	//
	// uat-agent
	FirstId *string `json:"FirstId,omitempty" xml:"FirstId,omitempty"`
	// example:
	//
	// false
	HasMore *string `json:"HasMore,omitempty" xml:"HasMore,omitempty"`
	// example:
	//
	// last-agent
	LastId *string `json:"LastId,omitempty" xml:"LastId,omitempty"`
}

func (s ListAgentsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListAgentsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListAgentsResponseBodyData) GetAgents() []*Agent {
	return s.Agents
}

func (s *ListAgentsResponseBodyData) GetFirstId() *string {
	return s.FirstId
}

func (s *ListAgentsResponseBodyData) GetHasMore() *string {
	return s.HasMore
}

func (s *ListAgentsResponseBodyData) GetLastId() *string {
	return s.LastId
}

func (s *ListAgentsResponseBodyData) SetAgents(v []*Agent) *ListAgentsResponseBodyData {
	s.Agents = v
	return s
}

func (s *ListAgentsResponseBodyData) SetFirstId(v string) *ListAgentsResponseBodyData {
	s.FirstId = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetHasMore(v string) *ListAgentsResponseBodyData {
	s.HasMore = &v
	return s
}

func (s *ListAgentsResponseBodyData) SetLastId(v string) *ListAgentsResponseBodyData {
	s.LastId = &v
	return s
}

func (s *ListAgentsResponseBodyData) Validate() error {
	if s.Agents != nil {
		for _, item := range s.Agents {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
