// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDataSourcesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataSources(v []*ListDataSourcesResponseBodyDataSources) *ListDataSourcesResponseBody
	GetDataSources() []*ListDataSourcesResponseBodyDataSources
	SetRequestId(v string) *ListDataSourcesResponseBody
	GetRequestId() *string
}

type ListDataSourcesResponseBody struct {
	// The list of data.
	DataSources []*ListDataSourcesResponseBodyDataSources `json:"dataSources,omitempty" xml:"dataSources,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3****
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListDataSourcesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBody) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBody) GetDataSources() []*ListDataSourcesResponseBodyDataSources {
	return s.DataSources
}

func (s *ListDataSourcesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDataSourcesResponseBody) SetDataSources(v []*ListDataSourcesResponseBodyDataSources) *ListDataSourcesResponseBody {
	s.DataSources = v
	return s
}

func (s *ListDataSourcesResponseBody) SetRequestId(v string) *ListDataSourcesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDataSourcesResponseBody) Validate() error {
	if s.DataSources != nil {
		for _, item := range s.DataSources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDataSourcesResponseBodyDataSources struct {
	// The data ID.
	//
	// example:
	//
	// cn-beijing
	Id *string `json:"id,omitempty" xml:"id,omitempty"`
}

func (s ListDataSourcesResponseBodyDataSources) String() string {
	return dara.Prettify(s)
}

func (s ListDataSourcesResponseBodyDataSources) GoString() string {
	return s.String()
}

func (s *ListDataSourcesResponseBodyDataSources) GetId() *string {
	return s.Id
}

func (s *ListDataSourcesResponseBodyDataSources) SetId(v string) *ListDataSourcesResponseBodyDataSources {
	s.Id = &v
	return s
}

func (s *ListDataSourcesResponseBodyDataSources) Validate() error {
	return dara.Validate(s)
}
