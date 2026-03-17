// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeUnbindFlowLogSagsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCount(v int32) *DescribeUnbindFlowLogSagsResponseBody
	GetCount() *int32
	SetRequestId(v string) *DescribeUnbindFlowLogSagsResponseBody
	GetRequestId() *string
	SetSags(v *DescribeUnbindFlowLogSagsResponseBodySags) *DescribeUnbindFlowLogSagsResponseBody
	GetSags() *DescribeUnbindFlowLogSagsResponseBodySags
}

type DescribeUnbindFlowLogSagsResponseBody struct {
	// The total number of the SAG instances.
	//
	// example:
	//
	// 20
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The ID of the request.
	//
	// example:
	//
	// C850C10E-9856-45FF-8640-80288BA467DF
	RequestId *string                                    `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Sags      *DescribeUnbindFlowLogSagsResponseBodySags `json:"Sags,omitempty" xml:"Sags,omitempty" type:"Struct"`
}

func (s DescribeUnbindFlowLogSagsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnbindFlowLogSagsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeUnbindFlowLogSagsResponseBody) GetCount() *int32 {
	return s.Count
}

func (s *DescribeUnbindFlowLogSagsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeUnbindFlowLogSagsResponseBody) GetSags() *DescribeUnbindFlowLogSagsResponseBodySags {
	return s.Sags
}

func (s *DescribeUnbindFlowLogSagsResponseBody) SetCount(v int32) *DescribeUnbindFlowLogSagsResponseBody {
	s.Count = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponseBody) SetRequestId(v string) *DescribeUnbindFlowLogSagsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponseBody) SetSags(v *DescribeUnbindFlowLogSagsResponseBodySags) *DescribeUnbindFlowLogSagsResponseBody {
	s.Sags = v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponseBody) Validate() error {
	if s.Sags != nil {
		if err := s.Sags.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeUnbindFlowLogSagsResponseBodySags struct {
	Sag []*DescribeUnbindFlowLogSagsResponseBodySagsSag `json:"Sag,omitempty" xml:"Sag,omitempty" type:"Repeated"`
}

func (s DescribeUnbindFlowLogSagsResponseBodySags) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnbindFlowLogSagsResponseBodySags) GoString() string {
	return s.String()
}

func (s *DescribeUnbindFlowLogSagsResponseBodySags) GetSag() []*DescribeUnbindFlowLogSagsResponseBodySagsSag {
	return s.Sag
}

func (s *DescribeUnbindFlowLogSagsResponseBodySags) SetSag(v []*DescribeUnbindFlowLogSagsResponseBodySagsSag) *DescribeUnbindFlowLogSagsResponseBodySags {
	s.Sag = v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponseBodySags) Validate() error {
	if s.Sag != nil {
		for _, item := range s.Sag {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeUnbindFlowLogSagsResponseBodySagsSag struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Name        *string `json:"Name,omitempty" xml:"Name,omitempty"`
	SmartAGId   *string `json:"SmartAGId,omitempty" xml:"SmartAGId,omitempty"`
}

func (s DescribeUnbindFlowLogSagsResponseBodySagsSag) String() string {
	return dara.Prettify(s)
}

func (s DescribeUnbindFlowLogSagsResponseBodySagsSag) GoString() string {
	return s.String()
}

func (s *DescribeUnbindFlowLogSagsResponseBodySagsSag) GetDescription() *string {
	return s.Description
}

func (s *DescribeUnbindFlowLogSagsResponseBodySagsSag) GetName() *string {
	return s.Name
}

func (s *DescribeUnbindFlowLogSagsResponseBodySagsSag) GetSmartAGId() *string {
	return s.SmartAGId
}

func (s *DescribeUnbindFlowLogSagsResponseBodySagsSag) SetDescription(v string) *DescribeUnbindFlowLogSagsResponseBodySagsSag {
	s.Description = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponseBodySagsSag) SetName(v string) *DescribeUnbindFlowLogSagsResponseBodySagsSag {
	s.Name = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponseBodySagsSag) SetSmartAGId(v string) *DescribeUnbindFlowLogSagsResponseBodySagsSag {
	s.SmartAGId = &v
	return s
}

func (s *DescribeUnbindFlowLogSagsResponseBodySagsSag) Validate() error {
	return dara.Validate(s)
}
