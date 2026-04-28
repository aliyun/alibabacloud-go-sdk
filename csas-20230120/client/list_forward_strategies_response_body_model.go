// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListForwardStrategiesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetForwardStrategies(v []*ListForwardStrategiesResponseBodyForwardStrategies) *ListForwardStrategiesResponseBody
	GetForwardStrategies() []*ListForwardStrategiesResponseBodyForwardStrategies
	SetRequestId(v string) *ListForwardStrategiesResponseBody
	GetRequestId() *string
	SetTotalNum(v int64) *ListForwardStrategiesResponseBody
	GetTotalNum() *int64
}

type ListForwardStrategiesResponseBody struct {
	ForwardStrategies []*ListForwardStrategiesResponseBodyForwardStrategies `json:"ForwardStrategies,omitempty" xml:"ForwardStrategies,omitempty" type:"Repeated"`
	// Id of the request
	//
	// example:
	//
	// D6707286-A50E-57B1-B2CF-EFAC59E850D8
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// example:
	//
	// 1
	TotalNum *int64 `json:"TotalNum,omitempty" xml:"TotalNum,omitempty"`
}

func (s ListForwardStrategiesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListForwardStrategiesResponseBody) GoString() string {
	return s.String()
}

func (s *ListForwardStrategiesResponseBody) GetForwardStrategies() []*ListForwardStrategiesResponseBodyForwardStrategies {
	return s.ForwardStrategies
}

func (s *ListForwardStrategiesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListForwardStrategiesResponseBody) GetTotalNum() *int64 {
	return s.TotalNum
}

func (s *ListForwardStrategiesResponseBody) SetForwardStrategies(v []*ListForwardStrategiesResponseBodyForwardStrategies) *ListForwardStrategiesResponseBody {
	s.ForwardStrategies = v
	return s
}

func (s *ListForwardStrategiesResponseBody) SetRequestId(v string) *ListForwardStrategiesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListForwardStrategiesResponseBody) SetTotalNum(v int64) *ListForwardStrategiesResponseBody {
	s.TotalNum = &v
	return s
}

func (s *ListForwardStrategiesResponseBody) Validate() error {
	if s.ForwardStrategies != nil {
		for _, item := range s.ForwardStrategies {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListForwardStrategiesResponseBodyForwardStrategies struct {
	// example:
	//
	// asdasdasd
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// example:
	//
	// connector-4178bc59bec56df1
	DestinationId *string `json:"DestinationId,omitempty" xml:"DestinationId,omitempty"`
	// example:
	//
	// Connector
	DestinationType *string `json:"DestinationType,omitempty" xml:"DestinationType,omitempty"`
	// example:
	//
	// fs-8b299ac5a93a0a3a
	ForwardId *string `json:"ForwardId,omitempty" xml:"ForwardId,omitempty"`
	Name      *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// example:
	//
	// 100
	Priority *string `json:"Priority,omitempty" xml:"Priority,omitempty"`
	// example:
	//
	// Enabled
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListForwardStrategiesResponseBodyForwardStrategies) String() string {
	return dara.Prettify(s)
}

func (s ListForwardStrategiesResponseBodyForwardStrategies) GoString() string {
	return s.String()
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) GetDescription() *string {
	return s.Description
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) GetDestinationId() *string {
	return s.DestinationId
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) GetDestinationType() *string {
	return s.DestinationType
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) GetForwardId() *string {
	return s.ForwardId
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) GetName() *string {
	return s.Name
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) GetPriority() *string {
	return s.Priority
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) GetStatus() *string {
	return s.Status
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) SetDescription(v string) *ListForwardStrategiesResponseBodyForwardStrategies {
	s.Description = &v
	return s
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) SetDestinationId(v string) *ListForwardStrategiesResponseBodyForwardStrategies {
	s.DestinationId = &v
	return s
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) SetDestinationType(v string) *ListForwardStrategiesResponseBodyForwardStrategies {
	s.DestinationType = &v
	return s
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) SetForwardId(v string) *ListForwardStrategiesResponseBodyForwardStrategies {
	s.ForwardId = &v
	return s
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) SetName(v string) *ListForwardStrategiesResponseBodyForwardStrategies {
	s.Name = &v
	return s
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) SetPriority(v string) *ListForwardStrategiesResponseBodyForwardStrategies {
	s.Priority = &v
	return s
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) SetStatus(v string) *ListForwardStrategiesResponseBodyForwardStrategies {
	s.Status = &v
	return s
}

func (s *ListForwardStrategiesResponseBodyForwardStrategies) Validate() error {
	return dara.Validate(s)
}
