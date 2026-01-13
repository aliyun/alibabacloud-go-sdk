// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListDatabaseExtensionsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetExtensions(v []*ListDatabaseExtensionsResponseBodyExtensions) *ListDatabaseExtensionsResponseBody
	GetExtensions() []*ListDatabaseExtensionsResponseBodyExtensions
	SetRequestId(v string) *ListDatabaseExtensionsResponseBody
	GetRequestId() *string
}

type ListDatabaseExtensionsResponseBody struct {
	// Extension list.
	Extensions []*ListDatabaseExtensionsResponseBodyExtensions `json:"Extensions,omitempty" xml:"Extensions,omitempty" type:"Repeated"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListDatabaseExtensionsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseExtensionsResponseBody) GoString() string {
	return s.String()
}

func (s *ListDatabaseExtensionsResponseBody) GetExtensions() []*ListDatabaseExtensionsResponseBodyExtensions {
	return s.Extensions
}

func (s *ListDatabaseExtensionsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListDatabaseExtensionsResponseBody) SetExtensions(v []*ListDatabaseExtensionsResponseBodyExtensions) *ListDatabaseExtensionsResponseBody {
	s.Extensions = v
	return s
}

func (s *ListDatabaseExtensionsResponseBody) SetRequestId(v string) *ListDatabaseExtensionsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListDatabaseExtensionsResponseBody) Validate() error {
	if s.Extensions != nil {
		for _, item := range s.Extensions {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListDatabaseExtensionsResponseBodyExtensions struct {
	// The description of the extension.
	//
	// example:
	//
	// zhparser
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The extension name.
	//
	// example:
	//
	// zhparser
	ExtensionName *string `json:"ExtensionName,omitempty" xml:"ExtensionName,omitempty"`
	// The status of the extension.
	//
	// example:
	//
	// installed
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ListDatabaseExtensionsResponseBodyExtensions) String() string {
	return dara.Prettify(s)
}

func (s ListDatabaseExtensionsResponseBodyExtensions) GoString() string {
	return s.String()
}

func (s *ListDatabaseExtensionsResponseBodyExtensions) GetDescription() *string {
	return s.Description
}

func (s *ListDatabaseExtensionsResponseBodyExtensions) GetExtensionName() *string {
	return s.ExtensionName
}

func (s *ListDatabaseExtensionsResponseBodyExtensions) GetStatus() *string {
	return s.Status
}

func (s *ListDatabaseExtensionsResponseBodyExtensions) SetDescription(v string) *ListDatabaseExtensionsResponseBodyExtensions {
	s.Description = &v
	return s
}

func (s *ListDatabaseExtensionsResponseBodyExtensions) SetExtensionName(v string) *ListDatabaseExtensionsResponseBodyExtensions {
	s.ExtensionName = &v
	return s
}

func (s *ListDatabaseExtensionsResponseBodyExtensions) SetStatus(v string) *ListDatabaseExtensionsResponseBodyExtensions {
	s.Status = &v
	return s
}

func (s *ListDatabaseExtensionsResponseBodyExtensions) Validate() error {
	return dara.Validate(s)
}
