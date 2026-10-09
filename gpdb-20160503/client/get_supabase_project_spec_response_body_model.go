// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSupabaseProjectSpecResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetItems(v []*GetSupabaseProjectSpecResponseBodyItems) *GetSupabaseProjectSpecResponseBody
	GetItems() []*GetSupabaseProjectSpecResponseBodyItems
	SetRequestId(v string) *GetSupabaseProjectSpecResponseBody
	GetRequestId() *string
	SetZoneIds(v []*string) *GetSupabaseProjectSpecResponseBody
	GetZoneIds() []*string
}

type GetSupabaseProjectSpecResponseBody struct {
	// The list of Supabase project specifications.
	Items []*GetSupabaseProjectSpecResponseBodyItems `json:"Items,omitempty" xml:"Items,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// B4CAF581-2AC7-41AD-8940-D56DF7AADF5B
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The list of zone IDs that support creating Supabase projects.
	ZoneIds []*string `json:"ZoneIds,omitempty" xml:"ZoneIds,omitempty" type:"Repeated"`
}

func (s GetSupabaseProjectSpecResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectSpecResponseBody) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectSpecResponseBody) GetItems() []*GetSupabaseProjectSpecResponseBodyItems {
	return s.Items
}

func (s *GetSupabaseProjectSpecResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetSupabaseProjectSpecResponseBody) GetZoneIds() []*string {
	return s.ZoneIds
}

func (s *GetSupabaseProjectSpecResponseBody) SetItems(v []*GetSupabaseProjectSpecResponseBodyItems) *GetSupabaseProjectSpecResponseBody {
	s.Items = v
	return s
}

func (s *GetSupabaseProjectSpecResponseBody) SetRequestId(v string) *GetSupabaseProjectSpecResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetSupabaseProjectSpecResponseBody) SetZoneIds(v []*string) *GetSupabaseProjectSpecResponseBody {
	s.ZoneIds = v
	return s
}

func (s *GetSupabaseProjectSpecResponseBody) Validate() error {
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

type GetSupabaseProjectSpecResponseBodyItems struct {
	// Indicates whether the specification is free.
	//
	// example:
	//
	// false
	Free *bool `json:"Free,omitempty" xml:"Free,omitempty"`
	// The specification code.
	//
	// example:
	//
	// 2C4G
	Spec *string `json:"Spec,omitempty" xml:"Spec,omitempty"`
	// Indicates whether the specification is visible.
	//
	// example:
	//
	// true
	Visible *bool `json:"Visible,omitempty" xml:"Visible,omitempty"`
}

func (s GetSupabaseProjectSpecResponseBodyItems) String() string {
	return dara.Prettify(s)
}

func (s GetSupabaseProjectSpecResponseBodyItems) GoString() string {
	return s.String()
}

func (s *GetSupabaseProjectSpecResponseBodyItems) GetFree() *bool {
	return s.Free
}

func (s *GetSupabaseProjectSpecResponseBodyItems) GetSpec() *string {
	return s.Spec
}

func (s *GetSupabaseProjectSpecResponseBodyItems) GetVisible() *bool {
	return s.Visible
}

func (s *GetSupabaseProjectSpecResponseBodyItems) SetFree(v bool) *GetSupabaseProjectSpecResponseBodyItems {
	s.Free = &v
	return s
}

func (s *GetSupabaseProjectSpecResponseBodyItems) SetSpec(v string) *GetSupabaseProjectSpecResponseBodyItems {
	s.Spec = &v
	return s
}

func (s *GetSupabaseProjectSpecResponseBodyItems) SetVisible(v bool) *GetSupabaseProjectSpecResponseBodyItems {
	s.Visible = &v
	return s
}

func (s *GetSupabaseProjectSpecResponseBodyItems) Validate() error {
	return dara.Validate(s)
}
