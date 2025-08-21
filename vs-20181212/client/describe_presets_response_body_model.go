// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribePresetsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetId(v string) *DescribePresetsResponseBody
	GetId() *string
	SetPresets(v []*DescribePresetsResponseBodyPresets) *DescribePresetsResponseBody
	GetPresets() []*DescribePresetsResponseBodyPresets
	SetRequestId(v string) *DescribePresetsResponseBody
	GetRequestId() *string
}

type DescribePresetsResponseBody struct {
	// example:
	//
	// 348*****380-cn-qingdao
	Id      *string                               `json:"Id,omitempty" xml:"Id,omitempty"`
	Presets []*DescribePresetsResponseBodyPresets `json:"Presets,omitempty" xml:"Presets,omitempty" type:"Repeated"`
	// example:
	//
	// 9FE0CA83-BFD3-4EBD-A429-FABB9B9AE772
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribePresetsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribePresetsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribePresetsResponseBody) GetId() *string {
	return s.Id
}

func (s *DescribePresetsResponseBody) GetPresets() []*DescribePresetsResponseBodyPresets {
	return s.Presets
}

func (s *DescribePresetsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribePresetsResponseBody) SetId(v string) *DescribePresetsResponseBody {
	s.Id = &v
	return s
}

func (s *DescribePresetsResponseBody) SetPresets(v []*DescribePresetsResponseBodyPresets) *DescribePresetsResponseBody {
	s.Presets = v
	return s
}

func (s *DescribePresetsResponseBody) SetRequestId(v string) *DescribePresetsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribePresetsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribePresetsResponseBodyPresets struct {
	// example:
	//
	// 2
	Id   *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
}

func (s DescribePresetsResponseBodyPresets) String() string {
	return dara.Prettify(s)
}

func (s DescribePresetsResponseBodyPresets) GoString() string {
	return s.String()
}

func (s *DescribePresetsResponseBodyPresets) GetId() *string {
	return s.Id
}

func (s *DescribePresetsResponseBodyPresets) GetName() *string {
	return s.Name
}

func (s *DescribePresetsResponseBodyPresets) SetId(v string) *DescribePresetsResponseBodyPresets {
	s.Id = &v
	return s
}

func (s *DescribePresetsResponseBodyPresets) SetName(v string) *DescribePresetsResponseBodyPresets {
	s.Name = &v
	return s
}

func (s *DescribePresetsResponseBodyPresets) Validate() error {
	return dara.Validate(s)
}
