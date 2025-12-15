// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAddonTemplatesResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAddons(v []*ListAddonTemplatesResponseBodyAddons) *ListAddonTemplatesResponseBody
	GetAddons() []*ListAddonTemplatesResponseBodyAddons
	SetPageNumber(v int64) *ListAddonTemplatesResponseBody
	GetPageNumber() *int64
	SetPageSize(v int64) *ListAddonTemplatesResponseBody
	GetPageSize() *int64
	SetRequestId(v string) *ListAddonTemplatesResponseBody
	GetRequestId() *string
	SetTotalCount(v int32) *ListAddonTemplatesResponseBody
	GetTotalCount() *int32
}

type ListAddonTemplatesResponseBody struct {
	// The information about the addon templates.
	Addons []*ListAddonTemplatesResponseBodyAddons `json:"Addons,omitempty" xml:"Addons,omitempty" type:"Repeated"`
	// The page number of the returned page.
	//
	// example:
	//
	// 1
	PageNumber *int64 `json:"PageNumber,omitempty" xml:"PageNumber,omitempty"`
	// The number of entries per page.
	//
	// example:
	//
	// 20
	PageSize *int64 `json:"PageSize,omitempty" xml:"PageSize,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 04F0F334-1335-436C-A1D7-6C044FE7****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of entries returned.
	//
	// example:
	//
	// 1
	TotalCount *int32 `json:"TotalCount,omitempty" xml:"TotalCount,omitempty"`
}

func (s ListAddonTemplatesResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListAddonTemplatesResponseBody) GoString() string {
	return s.String()
}

func (s *ListAddonTemplatesResponseBody) GetAddons() []*ListAddonTemplatesResponseBodyAddons {
	return s.Addons
}

func (s *ListAddonTemplatesResponseBody) GetPageNumber() *int64 {
	return s.PageNumber
}

func (s *ListAddonTemplatesResponseBody) GetPageSize() *int64 {
	return s.PageSize
}

func (s *ListAddonTemplatesResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListAddonTemplatesResponseBody) GetTotalCount() *int32 {
	return s.TotalCount
}

func (s *ListAddonTemplatesResponseBody) SetAddons(v []*ListAddonTemplatesResponseBodyAddons) *ListAddonTemplatesResponseBody {
	s.Addons = v
	return s
}

func (s *ListAddonTemplatesResponseBody) SetPageNumber(v int64) *ListAddonTemplatesResponseBody {
	s.PageNumber = &v
	return s
}

func (s *ListAddonTemplatesResponseBody) SetPageSize(v int64) *ListAddonTemplatesResponseBody {
	s.PageSize = &v
	return s
}

func (s *ListAddonTemplatesResponseBody) SetRequestId(v string) *ListAddonTemplatesResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListAddonTemplatesResponseBody) SetTotalCount(v int32) *ListAddonTemplatesResponseBody {
	s.TotalCount = &v
	return s
}

func (s *ListAddonTemplatesResponseBody) Validate() error {
	if s.Addons != nil {
		for _, item := range s.Addons {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListAddonTemplatesResponseBodyAddons struct {
	// The addon description.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The addon label
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
	// The addon name.
	//
	// This parameter is required.
	//
	// example:
	//
	// Login
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The addon version.
	//
	// This parameter is required.
	//
	// example:
	//
	// 1.0
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s ListAddonTemplatesResponseBodyAddons) String() string {
	return dara.Prettify(s)
}

func (s ListAddonTemplatesResponseBodyAddons) GoString() string {
	return s.String()
}

func (s *ListAddonTemplatesResponseBodyAddons) GetDescription() *string {
	return s.Description
}

func (s *ListAddonTemplatesResponseBodyAddons) GetLabel() *string {
	return s.Label
}

func (s *ListAddonTemplatesResponseBodyAddons) GetName() *string {
	return s.Name
}

func (s *ListAddonTemplatesResponseBodyAddons) GetVersion() *string {
	return s.Version
}

func (s *ListAddonTemplatesResponseBodyAddons) SetDescription(v string) *ListAddonTemplatesResponseBodyAddons {
	s.Description = &v
	return s
}

func (s *ListAddonTemplatesResponseBodyAddons) SetLabel(v string) *ListAddonTemplatesResponseBodyAddons {
	s.Label = &v
	return s
}

func (s *ListAddonTemplatesResponseBodyAddons) SetName(v string) *ListAddonTemplatesResponseBodyAddons {
	s.Name = &v
	return s
}

func (s *ListAddonTemplatesResponseBodyAddons) SetVersion(v string) *ListAddonTemplatesResponseBodyAddons {
	s.Version = &v
	return s
}

func (s *ListAddonTemplatesResponseBodyAddons) Validate() error {
	return dara.Validate(s)
}
