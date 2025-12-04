// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataEventServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v *ListDataEventServicesResponseBodyData) *ListDataEventServicesResponseBody
	GetData() *ListDataEventServicesResponseBodyData
	SetRequestId(v string) *ListDataEventServicesResponseBody
	GetRequestId() *string
}

type ListDataEventServicesResponseBody struct {
	Data *ListDataEventServicesResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// example:
	//
	// 851038F3-33AB-4C49-97D7-6AB37D35****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDataEventServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataEventServicesResponseBody) GetData() *ListDataEventServicesResponseBodyData {
	return s.Data
}

func (s *ListDataEventServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataEventServicesResponseBody) SetData(v *ListDataEventServicesResponseBodyData) *ListDataEventServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListDataEventServicesResponseBody) SetRequestId(v string) *ListDataEventServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataEventServicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListDataEventServicesResponseBodyData struct {
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// example:
	//
	// VjE6bHJlTGoxdm1M****
	NextToken    *string                                              `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	ServiceInfos []*ListDataEventServicesResponseBodyDataServiceInfos `json:"ServiceInfos,omitempty" xml:"ServiceInfos,omitempty" type:"Repeated"`
}

func (s ListDataEventServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListDataEventServicesResponseBodyData) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListDataEventServicesResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *ListDataEventServicesResponseBodyData) GetServiceInfos() []*ListDataEventServicesResponseBodyDataServiceInfos {
	return s.ServiceInfos
}

func (s *ListDataEventServicesResponseBodyData) SetMaxResults(v int32) *ListDataEventServicesResponseBodyData {
	s.MaxResults = &v
	return s
}

func (s *ListDataEventServicesResponseBodyData) SetNextToken(v string) *ListDataEventServicesResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *ListDataEventServicesResponseBodyData) SetServiceInfos(v []*ListDataEventServicesResponseBodyDataServiceInfos) *ListDataEventServicesResponseBodyData {
	s.ServiceInfos = v
	return s
}

func (s *ListDataEventServicesResponseBodyData) Validate() error {
	if s.ServiceInfos != nil {
		for _, item := range s.ServiceInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataEventServicesResponseBodyDataServiceInfos struct {
	EventNames []*string `json:"EventNames,omitempty" xml:"EventNames,omitempty" type:"Repeated"`
	// example:
	//
	// Cms
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListDataEventServicesResponseBodyDataServiceInfos) String() string {
	return dara.Prettify(s)
}

func (s ListDataEventServicesResponseBodyDataServiceInfos) GoString() string {
	return s.String()
}

func (s *ListDataEventServicesResponseBodyDataServiceInfos) GetEventNames() []*string {
	return s.EventNames
}

func (s *ListDataEventServicesResponseBodyDataServiceInfos) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListDataEventServicesResponseBodyDataServiceInfos) SetEventNames(v []*string) *ListDataEventServicesResponseBodyDataServiceInfos {
	s.EventNames = v
	return s
}

func (s *ListDataEventServicesResponseBodyDataServiceInfos) SetServiceName(v string) *ListDataEventServicesResponseBodyDataServiceInfos {
	s.ServiceName = &v
	return s
}

func (s *ListDataEventServicesResponseBodyDataServiceInfos) Validate() error {
	return dara.Validate(s)
}
