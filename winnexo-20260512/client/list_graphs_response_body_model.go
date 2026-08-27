// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListGraphsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListGraphsResponseBody
	GetCode() *string
	SetItems(v []*ListGraphsResponseBodyItems) *ListGraphsResponseBody
	GetItems() []*ListGraphsResponseBodyItems
	SetMessage(v string) *ListGraphsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListGraphsResponseBody
	GetRequestId() *string
}

type ListGraphsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// 200
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The list of MCP cards.
	//
	// This parameter is required.
	Items []*ListGraphsResponseBodyItems `json:"items,omitempty" xml:"items,omitempty" type:"Repeated"`
	// The prompt message.
	//
	// example:
	//
	// successful
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 019FF406-1B10-0065-A97D-2D1920C2A03D
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListGraphsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListGraphsResponseBody) GoString() string {
	return s.String()
}

func (s *ListGraphsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListGraphsResponseBody) GetItems() []*ListGraphsResponseBodyItems {
	return s.Items
}

func (s *ListGraphsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListGraphsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListGraphsResponseBody) SetCode(v string) *ListGraphsResponseBody {
	s.Code = &v
	return s
}

func (s *ListGraphsResponseBody) SetItems(v []*ListGraphsResponseBodyItems) *ListGraphsResponseBody {
	s.Items = v
	return s
}

func (s *ListGraphsResponseBody) SetMessage(v string) *ListGraphsResponseBody {
	s.Message = &v
	return s
}

func (s *ListGraphsResponseBody) SetRequestId(v string) *ListGraphsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListGraphsResponseBody) Validate() error {
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

type ListGraphsResponseBodyItems struct {
	// The business description of the knowledge graph. An empty string is returned if not configured.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	BusinessProfile *string `json:"businessProfile,omitempty" xml:"businessProfile,omitempty"`
	// The tool display name.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	DisplayName *string `json:"displayName,omitempty" xml:"displayName,omitempty"`
	// The knowledge graph name.
	//
	// This parameter is required.
	//
	// example:
	//
	// string_value
	GraphName *string `json:"graphName,omitempty" xml:"graphName,omitempty"`
	// Indicates whether this is the default group.
	//
	// This parameter is required.
	//
	// example:
	//
	// true
	IsDefault *bool `json:"isDefault,omitempty" xml:"isDefault,omitempty"`
}

func (s ListGraphsResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s ListGraphsResponseBodyItems) GoString() string {
	return s.String()
}

func (s *ListGraphsResponseBodyItems) GetBusinessProfile() *string {
	return s.BusinessProfile
}

func (s *ListGraphsResponseBodyItems) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListGraphsResponseBodyItems) GetGraphName() *string {
	return s.GraphName
}

func (s *ListGraphsResponseBodyItems) GetIsDefault() *bool {
	return s.IsDefault
}

func (s *ListGraphsResponseBodyItems) SetBusinessProfile(v string) *ListGraphsResponseBodyItems {
	s.BusinessProfile = &v
	return s
}

func (s *ListGraphsResponseBodyItems) SetDisplayName(v string) *ListGraphsResponseBodyItems {
	s.DisplayName = &v
	return s
}

func (s *ListGraphsResponseBodyItems) SetGraphName(v string) *ListGraphsResponseBodyItems {
	s.GraphName = &v
	return s
}

func (s *ListGraphsResponseBodyItems) SetIsDefault(v bool) *ListGraphsResponseBodyItems {
	s.IsDefault = &v
	return s
}

func (s *ListGraphsResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
