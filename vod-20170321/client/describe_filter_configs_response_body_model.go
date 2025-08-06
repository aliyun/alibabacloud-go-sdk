// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeFilterConfigsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetFilterConfigs(v []*DescribeFilterConfigsResponseBodyFilterConfigs) *DescribeFilterConfigsResponseBody
	GetFilterConfigs() []*DescribeFilterConfigsResponseBodyFilterConfigs
	SetRequestId(v string) *DescribeFilterConfigsResponseBody
	GetRequestId() *string
}

type DescribeFilterConfigsResponseBody struct {
	FilterConfigs []*DescribeFilterConfigsResponseBodyFilterConfigs `json:"FilterConfigs,omitempty" xml:"FilterConfigs,omitempty" type:"Repeated"`
	RequestId     *string                                           `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeFilterConfigsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilterConfigsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeFilterConfigsResponseBody) GetFilterConfigs() []*DescribeFilterConfigsResponseBodyFilterConfigs {
	return s.FilterConfigs
}

func (s *DescribeFilterConfigsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeFilterConfigsResponseBody) SetFilterConfigs(v []*DescribeFilterConfigsResponseBodyFilterConfigs) *DescribeFilterConfigsResponseBody {
	s.FilterConfigs = v
	return s
}

func (s *DescribeFilterConfigsResponseBody) SetRequestId(v string) *DescribeFilterConfigsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeFilterConfigsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeFilterConfigsResponseBodyFilterConfigs struct {
	FilterName  *string `json:"FilterName,omitempty" xml:"FilterName,omitempty"`
	ItemConfigs *string `json:"ItemConfigs,omitempty" xml:"ItemConfigs,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
	UuId        *string `json:"UuId,omitempty" xml:"UuId,omitempty"`
}

func (s DescribeFilterConfigsResponseBodyFilterConfigs) String() string {
	return dara.Prettify(s)
}

func (s DescribeFilterConfigsResponseBodyFilterConfigs) GoString() string {
	return s.String()
}

func (s *DescribeFilterConfigsResponseBodyFilterConfigs) GetFilterName() *string {
	return s.FilterName
}

func (s *DescribeFilterConfigsResponseBodyFilterConfigs) GetItemConfigs() *string {
	return s.ItemConfigs
}

func (s *DescribeFilterConfigsResponseBodyFilterConfigs) GetType() *string {
	return s.Type
}

func (s *DescribeFilterConfigsResponseBodyFilterConfigs) GetUuId() *string {
	return s.UuId
}

func (s *DescribeFilterConfigsResponseBodyFilterConfigs) SetFilterName(v string) *DescribeFilterConfigsResponseBodyFilterConfigs {
	s.FilterName = &v
	return s
}

func (s *DescribeFilterConfigsResponseBodyFilterConfigs) SetItemConfigs(v string) *DescribeFilterConfigsResponseBodyFilterConfigs {
	s.ItemConfigs = &v
	return s
}

func (s *DescribeFilterConfigsResponseBodyFilterConfigs) SetType(v string) *DescribeFilterConfigsResponseBodyFilterConfigs {
	s.Type = &v
	return s
}

func (s *DescribeFilterConfigsResponseBodyFilterConfigs) SetUuId(v string) *DescribeFilterConfigsResponseBodyFilterConfigs {
	s.UuId = &v
	return s
}

func (s *DescribeFilterConfigsResponseBodyFilterConfigs) Validate() error {
	return dara.Validate(s)
}
