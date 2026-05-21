// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodeTypesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodeTypes(v []*NodeType) *ListNodeTypesResponseBody
	GetNodeTypes() []*NodeType
	SetRequestId(v string) *ListNodeTypesResponseBody
	GetRequestId() *string
	SetStatistics(v []*NodeTypeStatistic) *ListNodeTypesResponseBody
	GetStatistics() []*NodeTypeStatistic
}

type ListNodeTypesResponseBody struct {
	NodeTypes []*NodeType `json:"NodeTypes,omitempty" xml:"NodeTypes,omitempty" type:"Repeated"`
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId  *string              `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Statistics []*NodeTypeStatistic `json:"Statistics,omitempty" xml:"Statistics,omitempty" type:"Repeated"`
}

func (s ListNodeTypesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodeTypesResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodeTypesResponseBody) GetNodeTypes() []*NodeType {
	return s.NodeTypes
}

func (s *ListNodeTypesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodeTypesResponseBody) GetStatistics() []*NodeTypeStatistic {
	return s.Statistics
}

func (s *ListNodeTypesResponseBody) SetNodeTypes(v []*NodeType) *ListNodeTypesResponseBody {
	s.NodeTypes = v
	return s
}

func (s *ListNodeTypesResponseBody) SetRequestId(v string) *ListNodeTypesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodeTypesResponseBody) SetStatistics(v []*NodeTypeStatistic) *ListNodeTypesResponseBody {
	s.Statistics = v
	return s
}

func (s *ListNodeTypesResponseBody) Validate() error {
	if s.NodeTypes != nil {
		for _, item := range s.NodeTypes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Statistics != nil {
		for _, item := range s.Statistics {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
