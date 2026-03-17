// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeSagVbrRelationsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeSagVbrRelationsResponseBody
	GetRequestId() *string
	SetSagVbrRelations(v []*DescribeSagVbrRelationsResponseBodySagVbrRelations) *DescribeSagVbrRelationsResponseBody
	GetSagVbrRelations() []*DescribeSagVbrRelationsResponseBodySagVbrRelations
}

type DescribeSagVbrRelationsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 17D79124-104A-42DB-8FCA-CE2957CD1723
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the specified VBR is associated with an SAG instance.
	SagVbrRelations []*DescribeSagVbrRelationsResponseBodySagVbrRelations `json:"SagVbrRelations,omitempty" xml:"SagVbrRelations,omitempty" type:"Repeated"`
}

func (s DescribeSagVbrRelationsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagVbrRelationsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeSagVbrRelationsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeSagVbrRelationsResponseBody) GetSagVbrRelations() []*DescribeSagVbrRelationsResponseBodySagVbrRelations {
	return s.SagVbrRelations
}

func (s *DescribeSagVbrRelationsResponseBody) SetRequestId(v string) *DescribeSagVbrRelationsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeSagVbrRelationsResponseBody) SetSagVbrRelations(v []*DescribeSagVbrRelationsResponseBodySagVbrRelations) *DescribeSagVbrRelationsResponseBody {
	s.SagVbrRelations = v
	return s
}

func (s *DescribeSagVbrRelationsResponseBody) Validate() error {
	if s.SagVbrRelations != nil {
		for _, item := range s.SagVbrRelations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeSagVbrRelationsResponseBodySagVbrRelations struct {
	// The ID of the SAG instance.
	//
	// example:
	//
	// sag-0nnteglltw6z4b****
	SagInstanceId *string `json:"SagInstanceId,omitempty" xml:"SagInstanceId,omitempty"`
	// The ID of the Alibaba Cloud account to which the SAG instance belongs.
	//
	// example:
	//
	// 16884015121212****
	SagUid *string `json:"SagUid,omitempty" xml:"SagUid,omitempty"`
	// The ID of the VBR.
	//
	// example:
	//
	// vbr-bp15ihkk93ezxppk****
	VbrInstanceId *string `json:"VbrInstanceId,omitempty" xml:"VbrInstanceId,omitempty"`
}

func (s DescribeSagVbrRelationsResponseBodySagVbrRelations) String() string {
	return dara.Prettify(s)
}

func (s DescribeSagVbrRelationsResponseBodySagVbrRelations) GoString() string {
	return s.String()
}

func (s *DescribeSagVbrRelationsResponseBodySagVbrRelations) GetSagInstanceId() *string {
	return s.SagInstanceId
}

func (s *DescribeSagVbrRelationsResponseBodySagVbrRelations) GetSagUid() *string {
	return s.SagUid
}

func (s *DescribeSagVbrRelationsResponseBodySagVbrRelations) GetVbrInstanceId() *string {
	return s.VbrInstanceId
}

func (s *DescribeSagVbrRelationsResponseBodySagVbrRelations) SetSagInstanceId(v string) *DescribeSagVbrRelationsResponseBodySagVbrRelations {
	s.SagInstanceId = &v
	return s
}

func (s *DescribeSagVbrRelationsResponseBodySagVbrRelations) SetSagUid(v string) *DescribeSagVbrRelationsResponseBodySagVbrRelations {
	s.SagUid = &v
	return s
}

func (s *DescribeSagVbrRelationsResponseBodySagVbrRelations) SetVbrInstanceId(v string) *DescribeSagVbrRelationsResponseBodySagVbrRelations {
	s.VbrInstanceId = &v
	return s
}

func (s *DescribeSagVbrRelationsResponseBodySagVbrRelations) Validate() error {
	return dara.Validate(s)
}
