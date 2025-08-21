// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRenderingInstanceSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeRenderingInstanceSettingsResponseBody
	GetRequestId() *string
	SetSettings(v []*DescribeRenderingInstanceSettingsResponseBodySettings) *DescribeRenderingInstanceSettingsResponseBody
	GetSettings() []*DescribeRenderingInstanceSettingsResponseBodySettings
}

type DescribeRenderingInstanceSettingsResponseBody struct {
	// Id of the request
	//
	// example:
	//
	// BEA5625F-8FCF-48F4-851B-CA63946DA664
	RequestId *string                                                  `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	Settings  []*DescribeRenderingInstanceSettingsResponseBodySettings `json:"Settings,omitempty" xml:"Settings,omitempty" type:"Repeated"`
}

func (s DescribeRenderingInstanceSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRenderingInstanceSettingsResponseBody) GetSettings() []*DescribeRenderingInstanceSettingsResponseBodySettings {
	return s.Settings
}

func (s *DescribeRenderingInstanceSettingsResponseBody) SetRequestId(v string) *DescribeRenderingInstanceSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRenderingInstanceSettingsResponseBody) SetSettings(v []*DescribeRenderingInstanceSettingsResponseBodySettings) *DescribeRenderingInstanceSettingsResponseBody {
	s.Settings = v
	return s
}

func (s *DescribeRenderingInstanceSettingsResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeRenderingInstanceSettingsResponseBodySettings struct {
	// example:
	//
	// navbar.hide
	AttributeName *string `json:"AttributeName,omitempty" xml:"AttributeName,omitempty"`
	// example:
	//
	// 1
	AttributeValue *string `json:"AttributeValue,omitempty" xml:"AttributeValue,omitempty"`
}

func (s DescribeRenderingInstanceSettingsResponseBodySettings) String() string {
	return dara.Prettify(s)
}

func (s DescribeRenderingInstanceSettingsResponseBodySettings) GoString() string {
	return s.String()
}

func (s *DescribeRenderingInstanceSettingsResponseBodySettings) GetAttributeName() *string {
	return s.AttributeName
}

func (s *DescribeRenderingInstanceSettingsResponseBodySettings) GetAttributeValue() *string {
	return s.AttributeValue
}

func (s *DescribeRenderingInstanceSettingsResponseBodySettings) SetAttributeName(v string) *DescribeRenderingInstanceSettingsResponseBodySettings {
	s.AttributeName = &v
	return s
}

func (s *DescribeRenderingInstanceSettingsResponseBodySettings) SetAttributeValue(v string) *DescribeRenderingInstanceSettingsResponseBodySettings {
	s.AttributeValue = &v
	return s
}

func (s *DescribeRenderingInstanceSettingsResponseBodySettings) Validate() error {
	return dara.Validate(s)
}
