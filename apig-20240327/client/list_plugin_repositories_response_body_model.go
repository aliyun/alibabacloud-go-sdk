// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListPluginRepositoriesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListPluginRepositoriesResponseBody
	GetCode() *string
	SetData(v []*ListPluginRepositoriesResponseBodyData) *ListPluginRepositoriesResponseBody
	GetData() []*ListPluginRepositoriesResponseBodyData
	SetMessage(v string) *ListPluginRepositoriesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListPluginRepositoriesResponseBody
	GetRequestId() *string
}

type ListPluginRepositoriesResponseBody struct {
	// example:
	//
	// 200
	Code *string                                   `json:"code,omitempty" xml:"code,omitempty"`
	Data []*ListPluginRepositoriesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// 019FCA83-0416-588D-9763-2474980495F1
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListPluginRepositoriesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListPluginRepositoriesResponseBody) GoString() string {
	return s.String()
}

func (s *ListPluginRepositoriesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListPluginRepositoriesResponseBody) GetData() []*ListPluginRepositoriesResponseBodyData {
	return s.Data
}

func (s *ListPluginRepositoriesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListPluginRepositoriesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListPluginRepositoriesResponseBody) SetCode(v string) *ListPluginRepositoriesResponseBody {
	s.Code = &v
	return s
}

func (s *ListPluginRepositoriesResponseBody) SetData(v []*ListPluginRepositoriesResponseBodyData) *ListPluginRepositoriesResponseBody {
	s.Data = v
	return s
}

func (s *ListPluginRepositoriesResponseBody) SetMessage(v string) *ListPluginRepositoriesResponseBody {
	s.Message = &v
	return s
}

func (s *ListPluginRepositoriesResponseBody) SetRequestId(v string) *ListPluginRepositoriesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListPluginRepositoriesResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPluginRepositoriesResponseBodyData struct {
	// example:
	//
	// 664f1e2xxxx
	OrganizationId *string `json:"organizationId,omitempty" xml:"organizationId,omitempty"`
	// example:
	//
	// my-org
	OrganizationName *string                                               `json:"organizationName,omitempty" xml:"organizationName,omitempty"`
	Repositories     []*ListPluginRepositoriesResponseBodyDataRepositories `json:"repositories,omitempty" xml:"repositories,omitempty" type:"Repeated"`
}

func (s ListPluginRepositoriesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListPluginRepositoriesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListPluginRepositoriesResponseBodyData) GetOrganizationId() *string {
	return s.OrganizationId
}

func (s *ListPluginRepositoriesResponseBodyData) GetOrganizationName() *string {
	return s.OrganizationName
}

func (s *ListPluginRepositoriesResponseBodyData) GetRepositories() []*ListPluginRepositoriesResponseBodyDataRepositories {
	return s.Repositories
}

func (s *ListPluginRepositoriesResponseBodyData) SetOrganizationId(v string) *ListPluginRepositoriesResponseBodyData {
	s.OrganizationId = &v
	return s
}

func (s *ListPluginRepositoriesResponseBodyData) SetOrganizationName(v string) *ListPluginRepositoriesResponseBodyData {
	s.OrganizationName = &v
	return s
}

func (s *ListPluginRepositoriesResponseBodyData) SetRepositories(v []*ListPluginRepositoriesResponseBodyDataRepositories) *ListPluginRepositoriesResponseBodyData {
	s.Repositories = v
	return s
}

func (s *ListPluginRepositoriesResponseBodyData) Validate() error {
	if s.Repositories != nil {
		for _, item := range s.Repositories {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListPluginRepositoriesResponseBodyDataRepositories struct {
	// example:
	//
	// 12345678
	RepositoryId *string `json:"repositoryId,omitempty" xml:"repositoryId,omitempty"`
	// example:
	//
	// my-custom-plugin
	RepositoryName *string `json:"repositoryName,omitempty" xml:"repositoryName,omitempty"`
}

func (s ListPluginRepositoriesResponseBodyDataRepositories) String() string {
	return dara.Prettify(s)
}

func (s ListPluginRepositoriesResponseBodyDataRepositories) GoString() string {
	return s.String()
}

func (s *ListPluginRepositoriesResponseBodyDataRepositories) GetRepositoryId() *string {
	return s.RepositoryId
}

func (s *ListPluginRepositoriesResponseBodyDataRepositories) GetRepositoryName() *string {
	return s.RepositoryName
}

func (s *ListPluginRepositoriesResponseBodyDataRepositories) SetRepositoryId(v string) *ListPluginRepositoriesResponseBodyDataRepositories {
	s.RepositoryId = &v
	return s
}

func (s *ListPluginRepositoriesResponseBodyDataRepositories) SetRepositoryName(v string) *ListPluginRepositoriesResponseBodyDataRepositories {
	s.RepositoryName = &v
	return s
}

func (s *ListPluginRepositoriesResponseBodyDataRepositories) Validate() error {
	return dara.Validate(s)
}
