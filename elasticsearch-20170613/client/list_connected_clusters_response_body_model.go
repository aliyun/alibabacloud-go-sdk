// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListConnectedClustersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListConnectedClustersResponseBody
	GetRequestId() *string
	SetResult(v *ListConnectedClustersResponseBodyResult) *ListConnectedClustersResponseBody
	GetResult() *ListConnectedClustersResponseBodyResult
}

type ListConnectedClustersResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5FFD9ED4-C2EC-4E89-B22B-1ACB6FE1D***
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result *ListConnectedClustersResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s ListConnectedClustersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListConnectedClustersResponseBody) GoString() string {
	return s.String()
}

func (s *ListConnectedClustersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListConnectedClustersResponseBody) GetResult() *ListConnectedClustersResponseBodyResult {
	return s.Result
}

func (s *ListConnectedClustersResponseBody) SetRequestId(v string) *ListConnectedClustersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListConnectedClustersResponseBody) SetResult(v *ListConnectedClustersResponseBodyResult) *ListConnectedClustersResponseBody {
	s.Result = v
	return s
}

func (s *ListConnectedClustersResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListConnectedClustersResponseBodyResult struct {
	Result []*ListConnectedClustersResponseBodyResultResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Repeated"`
}

func (s ListConnectedClustersResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s ListConnectedClustersResponseBodyResult) GoString() string {
	return s.String()
}

func (s *ListConnectedClustersResponseBodyResult) GetResult() []*ListConnectedClustersResponseBodyResultResult {
	return s.Result
}

func (s *ListConnectedClustersResponseBodyResult) SetResult(v []*ListConnectedClustersResponseBodyResultResult) *ListConnectedClustersResponseBodyResult {
	s.Result = v
	return s
}

func (s *ListConnectedClustersResponseBodyResult) Validate() error {
	if s.Result != nil {
		for _, item := range s.Result {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListConnectedClustersResponseBodyResultResult struct {
	// The ID of the remote instance that is connected to the network of the current instance.
	//
	// example:
	//
	// es-cn-09k1rocex0006****
	Instances *string `json:"instances,omitempty" xml:"instances,omitempty"`
	// The network type of the instance.
	//
	// example:
	//
	// vpc
	NetworkType *string `json:"networkType,omitempty" xml:"networkType,omitempty"`
}

func (s ListConnectedClustersResponseBodyResultResult) String() string {
	return dara.Prettify(s)
}

func (s ListConnectedClustersResponseBodyResultResult) GoString() string {
	return s.String()
}

func (s *ListConnectedClustersResponseBodyResultResult) GetInstances() *string {
	return s.Instances
}

func (s *ListConnectedClustersResponseBodyResultResult) GetNetworkType() *string {
	return s.NetworkType
}

func (s *ListConnectedClustersResponseBodyResultResult) SetInstances(v string) *ListConnectedClustersResponseBodyResultResult {
	s.Instances = &v
	return s
}

func (s *ListConnectedClustersResponseBodyResultResult) SetNetworkType(v string) *ListConnectedClustersResponseBodyResultResult {
	s.NetworkType = &v
	return s
}

func (s *ListConnectedClustersResponseBodyResultResult) Validate() error {
	return dara.Validate(s)
}
