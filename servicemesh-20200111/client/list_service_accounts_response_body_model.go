// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListServiceAccountsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *ListServiceAccountsResponseBody
	GetRequestId() *string
	SetServiceAccounts(v []*ListServiceAccountsResponseBodyServiceAccounts) *ListServiceAccountsResponseBody
	GetServiceAccounts() []*ListServiceAccountsResponseBodyServiceAccounts
}

type ListServiceAccountsResponseBody struct {
	// The request ID.
	//
	// example:
	//
	// 8349374D-0F22-5CAB-9DE3-8CCE8EFA71FF
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of service accounts.
	ServiceAccounts []*ListServiceAccountsResponseBodyServiceAccounts `json:"ServiceAccounts,omitempty" xml:"ServiceAccounts,omitempty" type:"Repeated"`
}

func (s ListServiceAccountsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListServiceAccountsResponseBody) GoString() string {
	return s.String()
}

func (s *ListServiceAccountsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListServiceAccountsResponseBody) GetServiceAccounts() []*ListServiceAccountsResponseBodyServiceAccounts {
	return s.ServiceAccounts
}

func (s *ListServiceAccountsResponseBody) SetRequestId(v string) *ListServiceAccountsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListServiceAccountsResponseBody) SetServiceAccounts(v []*ListServiceAccountsResponseBodyServiceAccounts) *ListServiceAccountsResponseBody {
	s.ServiceAccounts = v
	return s
}

func (s *ListServiceAccountsResponseBody) Validate() error {
	if s.ServiceAccounts != nil {
		for _, item := range s.ServiceAccounts {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListServiceAccountsResponseBodyServiceAccounts struct {
	// The name of the service account.
	//
	// example:
	//
	// bookinfo-reviews
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The namespace where the service account resides.
	//
	// example:
	//
	// default
	Namespace *string `json:"Namespace,omitempty" xml:"Namespace,omitempty"`
}

func (s ListServiceAccountsResponseBodyServiceAccounts) String() string {
	return dara.Prettify(s)
}

func (s ListServiceAccountsResponseBodyServiceAccounts) GoString() string {
	return s.String()
}

func (s *ListServiceAccountsResponseBodyServiceAccounts) GetName() *string {
	return s.Name
}

func (s *ListServiceAccountsResponseBodyServiceAccounts) GetNamespace() *string {
	return s.Namespace
}

func (s *ListServiceAccountsResponseBodyServiceAccounts) SetName(v string) *ListServiceAccountsResponseBodyServiceAccounts {
	s.Name = &v
	return s
}

func (s *ListServiceAccountsResponseBodyServiceAccounts) SetNamespace(v string) *ListServiceAccountsResponseBodyServiceAccounts {
	s.Namespace = &v
	return s
}

func (s *ListServiceAccountsResponseBodyServiceAccounts) Validate() error {
	return dara.Validate(s)
}
