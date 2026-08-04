// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListNodePodsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNodePodInfos(v []*NodePodInfo) *ListNodePodsResponseBody
	GetNodePodInfos() []*NodePodInfo
	SetRequestId(v string) *ListNodePodsResponseBody
	GetRequestId() *string
}

type ListNodePodsResponseBody struct {
	// The node pod information.
	NodePodInfos []*NodePodInfo `json:"NodePodInfos,omitempty" xml:"NodePodInfos,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListNodePodsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListNodePodsResponseBody) GoString() string {
	return s.String()
}

func (s *ListNodePodsResponseBody) GetNodePodInfos() []*NodePodInfo {
	return s.NodePodInfos
}

func (s *ListNodePodsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListNodePodsResponseBody) SetNodePodInfos(v []*NodePodInfo) *ListNodePodsResponseBody {
	s.NodePodInfos = v
	return s
}

func (s *ListNodePodsResponseBody) SetRequestId(v string) *ListNodePodsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListNodePodsResponseBody) Validate() error {
	if s.NodePodInfos != nil {
		for _, item := range s.NodePodInfos {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
