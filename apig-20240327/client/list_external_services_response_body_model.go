// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExternalServicesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListExternalServicesResponseBody
	GetCode() *string
	SetData(v *ListExternalServicesResponseBodyData) *ListExternalServicesResponseBody
	GetData() *ListExternalServicesResponseBodyData
	SetMessage(v string) *ListExternalServicesResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListExternalServicesResponseBody
	GetRequestId() *string
}

type ListExternalServicesResponseBody struct {
	// The response code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of APIs.
	Data *ListExternalServicesResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Struct"`
	// The response message.
	//
	// example:
	//
	// Failed to upload data. Please try again
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// Id of the request
	//
	// example:
	//
	// D9B03B59-B2D6-51B0-A7E9-AF8466E320CA
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListExternalServicesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListExternalServicesResponseBody) GoString() string {
	return s.String()
}

func (s *ListExternalServicesResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListExternalServicesResponseBody) GetData() *ListExternalServicesResponseBodyData {
	return s.Data
}

func (s *ListExternalServicesResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListExternalServicesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListExternalServicesResponseBody) SetCode(v string) *ListExternalServicesResponseBody {
	s.Code = &v
	return s
}

func (s *ListExternalServicesResponseBody) SetData(v *ListExternalServicesResponseBodyData) *ListExternalServicesResponseBody {
	s.Data = v
	return s
}

func (s *ListExternalServicesResponseBody) SetMessage(v string) *ListExternalServicesResponseBody {
	s.Message = &v
	return s
}

func (s *ListExternalServicesResponseBody) SetRequestId(v string) *ListExternalServicesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListExternalServicesResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListExternalServicesResponseBodyData struct {
	// The service information.
	Items []*ListExternalServicesResponseBodyDataItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
}

func (s ListExternalServicesResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListExternalServicesResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListExternalServicesResponseBodyData) GetItems() []*ListExternalServicesResponseBodyDataItems {
	return s.Items
}

func (s *ListExternalServicesResponseBodyData) SetItems(v []*ListExternalServicesResponseBodyDataItems) *ListExternalServicesResponseBodyData {
	s.Items = v
	return s
}

func (s *ListExternalServicesResponseBodyData) Validate() error {
	if s.Items != nil {
		for _, item := range s.Items {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExternalServicesResponseBodyDataItems struct {
	// The namespace.
	//
	// example:
	//
	// PUBLIC
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
	// The display name of the namespace.
	//
	// example:
	//
	// PUBLIC
	NamespaceShowName *string `json:"namespaceShowName,omitempty" xml:"namespaceShowName,omitempty"`
	// The backend services.
	Services []*ListExternalServicesResponseBodyDataItemsServices `json:"services,omitempty" xml:"services,omitempty" type:"Repeated"`
}

func (s ListExternalServicesResponseBodyDataItems) String() string {
	return dara.Prettify(s)
}

func (s ListExternalServicesResponseBodyDataItems) GoString() string {
	return s.String()
}

func (s *ListExternalServicesResponseBodyDataItems) GetNamespace() *string {
	return s.Namespace
}

func (s *ListExternalServicesResponseBodyDataItems) GetNamespaceShowName() *string {
	return s.NamespaceShowName
}

func (s *ListExternalServicesResponseBodyDataItems) GetServices() []*ListExternalServicesResponseBodyDataItemsServices {
	return s.Services
}

func (s *ListExternalServicesResponseBodyDataItems) SetNamespace(v string) *ListExternalServicesResponseBodyDataItems {
	s.Namespace = &v
	return s
}

func (s *ListExternalServicesResponseBodyDataItems) SetNamespaceShowName(v string) *ListExternalServicesResponseBodyDataItems {
	s.NamespaceShowName = &v
	return s
}

func (s *ListExternalServicesResponseBodyDataItems) SetServices(v []*ListExternalServicesResponseBodyDataItemsServices) *ListExternalServicesResponseBodyDataItems {
	s.Services = v
	return s
}

func (s *ListExternalServicesResponseBodyDataItems) Validate() error {
	if s.Services != nil {
		for _, item := range s.Services {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListExternalServicesResponseBodyDataItemsServices struct {
	// The name.
	//
	// example:
	//
	// enable-slash-merge
	Name *string `json:"name,omitempty" xml:"name,omitempty"`
	// The namespace.
	//
	// example:
	//
	// public
	Namespace *string `json:"namespace,omitempty" xml:"namespace,omitempty"`
}

func (s ListExternalServicesResponseBodyDataItemsServices) String() string {
	return dara.Prettify(s)
}

func (s ListExternalServicesResponseBodyDataItemsServices) GoString() string {
	return s.String()
}

func (s *ListExternalServicesResponseBodyDataItemsServices) GetName() *string {
	return s.Name
}

func (s *ListExternalServicesResponseBodyDataItemsServices) GetNamespace() *string {
	return s.Namespace
}

func (s *ListExternalServicesResponseBodyDataItemsServices) SetName(v string) *ListExternalServicesResponseBodyDataItemsServices {
	s.Name = &v
	return s
}

func (s *ListExternalServicesResponseBodyDataItemsServices) SetNamespace(v string) *ListExternalServicesResponseBodyDataItemsServices {
	s.Namespace = &v
	return s
}

func (s *ListExternalServicesResponseBodyDataItemsServices) Validate() error {
	return dara.Validate(s)
}
