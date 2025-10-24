// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceContainersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetContainers(v []*ContainerInfo) *ListServiceContainersResponseBody
	GetContainers() []*ContainerInfo
	SetRequestId(v string) *ListServiceContainersResponseBody
	GetRequestId() *string
	SetServiceName(v string) *ListServiceContainersResponseBody
	GetServiceName() *string
}

type ListServiceContainersResponseBody struct {
	// The containers of the service.
	Containers []*ContainerInfo `json:"Containers,omitempty" xml:"Containers,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 40325405-579C-4D82****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The service name.
	//
	// example:
	//
	// foo
	ServiceName *string `json:"ServiceName,omitempty" xml:"ServiceName,omitempty"`
}

func (s ListServiceContainersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceContainersResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceContainersResponseBody) GetContainers() []*ContainerInfo {
	return s.Containers
}

func (s *ListServiceContainersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceContainersResponseBody) GetServiceName() *string {
	return s.ServiceName
}

func (s *ListServiceContainersResponseBody) SetContainers(v []*ContainerInfo) *ListServiceContainersResponseBody {
	s.Containers = v
	return s
}

func (s *ListServiceContainersResponseBody) SetRequestId(v string) *ListServiceContainersResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceContainersResponseBody) SetServiceName(v string) *ListServiceContainersResponseBody {
	s.ServiceName = &v
	return s
}

func (s *ListServiceContainersResponseBody) Validate() error {
	if s.Containers != nil {
		for _, item := range s.Containers {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
