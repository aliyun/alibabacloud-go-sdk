// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListEdgeContainerAppImageSecretsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetImageSecretList(v []*ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) *ListEdgeContainerAppImageSecretsResponseBody
	GetImageSecretList() []*ListEdgeContainerAppImageSecretsResponseBodyImageSecretList
	SetRequestId(v string) *ListEdgeContainerAppImageSecretsResponseBody
	GetRequestId() *string
}

type ListEdgeContainerAppImageSecretsResponseBody struct {
	// List of image secrets.
	ImageSecretList []*ListEdgeContainerAppImageSecretsResponseBodyImageSecretList `json:"ImageSecretList,omitempty" xml:"ImageSecretList,omitempty" type:"Repeated"`
	// Request ID.
	//
	// example:
	//
	// 3558df77-8a7a-4060-a900-2d7949403836
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListEdgeContainerAppImageSecretsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppImageSecretsResponseBody) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppImageSecretsResponseBody) GetImageSecretList() []*ListEdgeContainerAppImageSecretsResponseBodyImageSecretList {
	return s.ImageSecretList
}

func (s *ListEdgeContainerAppImageSecretsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListEdgeContainerAppImageSecretsResponseBody) SetImageSecretList(v []*ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) *ListEdgeContainerAppImageSecretsResponseBody {
	s.ImageSecretList = v
	return s
}

func (s *ListEdgeContainerAppImageSecretsResponseBody) SetRequestId(v string) *ListEdgeContainerAppImageSecretsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListEdgeContainerAppImageSecretsResponseBody) Validate() error {
	if s.ImageSecretList != nil {
		for _, item := range s.ImageSecretList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListEdgeContainerAppImageSecretsResponseBodyImageSecretList struct {
	// Name of the image secret.
	//
	// example:
	//
	// reg-123****
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// Registry address.
	//
	// example:
	//
	// registry-vpc.cn-hangzhou.aliyuncs.com
	Registry *string `json:"Registry,omitempty" xml:"Registry,omitempty"`
	// Username for the image repository
	//
	// example:
	//
	// admin
	Username *string `json:"Username,omitempty" xml:"Username,omitempty"`
}

func (s ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) String() string {
	return dara.Prettify(s)
}

func (s ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) GoString() string {
	return s.String()
}

func (s *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) GetName() *string {
	return s.Name
}

func (s *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) GetRegistry() *string {
	return s.Registry
}

func (s *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) GetUsername() *string {
	return s.Username
}

func (s *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) SetName(v string) *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList {
	s.Name = &v
	return s
}

func (s *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) SetRegistry(v string) *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList {
	s.Registry = &v
	return s
}

func (s *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) SetUsername(v string) *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList {
	s.Username = &v
	return s
}

func (s *ListEdgeContainerAppImageSecretsResponseBodyImageSecretList) Validate() error {
	return dara.Validate(s)
}
