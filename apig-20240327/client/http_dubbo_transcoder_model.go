// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iHttpDubboTranscoder interface {
	dara.Model
	String() string
	GoString() string
	SetDubboServiceGroup(v string) *HttpDubboTranscoder
	GetDubboServiceGroup() *string
	SetDubboServiceName(v string) *HttpDubboTranscoder
	GetDubboServiceName() *string
	SetDubboServiceVersion(v string) *HttpDubboTranscoder
	GetDubboServiceVersion() *string
	SetMothedMapList(v []*HttpDubboTranscoderMothedMapList) *HttpDubboTranscoder
	GetMothedMapList() []*HttpDubboTranscoderMothedMapList
}

type HttpDubboTranscoder struct {
	DubboServiceGroup   *string                             `json:"dubboServiceGroup,omitempty" xml:"dubboServiceGroup,omitempty"`
	DubboServiceName    *string                             `json:"dubboServiceName,omitempty" xml:"dubboServiceName,omitempty"`
	DubboServiceVersion *string                             `json:"dubboServiceVersion,omitempty" xml:"dubboServiceVersion,omitempty"`
	MothedMapList       []*HttpDubboTranscoderMothedMapList `json:"mothedMapList,omitempty" xml:"mothedMapList,omitempty" type:"Repeated"`
}

func (s HttpDubboTranscoder) String() string {
	return dara.Prettify(s)
}

func (s HttpDubboTranscoder) GoString() string {
	return s.String()
}

func (s *HttpDubboTranscoder) GetDubboServiceGroup() *string {
	return s.DubboServiceGroup
}

func (s *HttpDubboTranscoder) GetDubboServiceName() *string {
	return s.DubboServiceName
}

func (s *HttpDubboTranscoder) GetDubboServiceVersion() *string {
	return s.DubboServiceVersion
}

func (s *HttpDubboTranscoder) GetMothedMapList() []*HttpDubboTranscoderMothedMapList {
	return s.MothedMapList
}

func (s *HttpDubboTranscoder) SetDubboServiceGroup(v string) *HttpDubboTranscoder {
	s.DubboServiceGroup = &v
	return s
}

func (s *HttpDubboTranscoder) SetDubboServiceName(v string) *HttpDubboTranscoder {
	s.DubboServiceName = &v
	return s
}

func (s *HttpDubboTranscoder) SetDubboServiceVersion(v string) *HttpDubboTranscoder {
	s.DubboServiceVersion = &v
	return s
}

func (s *HttpDubboTranscoder) SetMothedMapList(v []*HttpDubboTranscoderMothedMapList) *HttpDubboTranscoder {
	s.MothedMapList = v
	return s
}

func (s *HttpDubboTranscoder) Validate() error {
	if s.MothedMapList != nil {
		for _, item := range s.MothedMapList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpDubboTranscoderMothedMapList struct {
	DubboMothedName *string `json:"dubboMothedName,omitempty" xml:"dubboMothedName,omitempty"`
	// example:
	//
	// ALL_GET
	HttpMothed *string `json:"httpMothed,omitempty" xml:"httpMothed,omitempty"`
	// example:
	//
	// /mytestzbk/sayhello
	Mothedpath    *string                                          `json:"mothedpath,omitempty" xml:"mothedpath,omitempty"`
	ParamMapsList []*HttpDubboTranscoderMothedMapListParamMapsList `json:"paramMapsList,omitempty" xml:"paramMapsList,omitempty" type:"Repeated"`
	// example:
	//
	// PASS_NOT
	PassThroughAllHeaders *string   `json:"passThroughAllHeaders,omitempty" xml:"passThroughAllHeaders,omitempty"`
	PassThroughList       []*string `json:"passThroughList,omitempty" xml:"passThroughList,omitempty" type:"Repeated"`
}

func (s HttpDubboTranscoderMothedMapList) String() string {
	return dara.Prettify(s)
}

func (s HttpDubboTranscoderMothedMapList) GoString() string {
	return s.String()
}

func (s *HttpDubboTranscoderMothedMapList) GetDubboMothedName() *string {
	return s.DubboMothedName
}

func (s *HttpDubboTranscoderMothedMapList) GetHttpMothed() *string {
	return s.HttpMothed
}

func (s *HttpDubboTranscoderMothedMapList) GetMothedpath() *string {
	return s.Mothedpath
}

func (s *HttpDubboTranscoderMothedMapList) GetParamMapsList() []*HttpDubboTranscoderMothedMapListParamMapsList {
	return s.ParamMapsList
}

func (s *HttpDubboTranscoderMothedMapList) GetPassThroughAllHeaders() *string {
	return s.PassThroughAllHeaders
}

func (s *HttpDubboTranscoderMothedMapList) GetPassThroughList() []*string {
	return s.PassThroughList
}

func (s *HttpDubboTranscoderMothedMapList) SetDubboMothedName(v string) *HttpDubboTranscoderMothedMapList {
	s.DubboMothedName = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetHttpMothed(v string) *HttpDubboTranscoderMothedMapList {
	s.HttpMothed = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetMothedpath(v string) *HttpDubboTranscoderMothedMapList {
	s.Mothedpath = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetParamMapsList(v []*HttpDubboTranscoderMothedMapListParamMapsList) *HttpDubboTranscoderMothedMapList {
	s.ParamMapsList = v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetPassThroughAllHeaders(v string) *HttpDubboTranscoderMothedMapList {
	s.PassThroughAllHeaders = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) SetPassThroughList(v []*string) *HttpDubboTranscoderMothedMapList {
	s.PassThroughList = v
	return s
}

func (s *HttpDubboTranscoderMothedMapList) Validate() error {
	if s.ParamMapsList != nil {
		for _, item := range s.ParamMapsList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpDubboTranscoderMothedMapListParamMapsList struct {
	// example:
	//
	// name
	ExtractKey *string `json:"extractKey,omitempty" xml:"extractKey,omitempty"`
	// example:
	//
	// ALL_QUERY_PARAMETER
	ExtractKeySpec *string `json:"extractKeySpec,omitempty" xml:"extractKeySpec,omitempty"`
	// example:
	//
	// java.lang.String
	MappingType *string `json:"mappingType,omitempty" xml:"mappingType,omitempty"`
}

func (s HttpDubboTranscoderMothedMapListParamMapsList) String() string {
	return dara.Prettify(s)
}

func (s HttpDubboTranscoderMothedMapListParamMapsList) GoString() string {
	return s.String()
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) GetExtractKey() *string {
	return s.ExtractKey
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) GetExtractKeySpec() *string {
	return s.ExtractKeySpec
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) GetMappingType() *string {
	return s.MappingType
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) SetExtractKey(v string) *HttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKey = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) SetExtractKeySpec(v string) *HttpDubboTranscoderMothedMapListParamMapsList {
	s.ExtractKeySpec = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) SetMappingType(v string) *HttpDubboTranscoderMothedMapListParamMapsList {
	s.MappingType = &v
	return s
}

func (s *HttpDubboTranscoderMothedMapListParamMapsList) Validate() error {
	return dara.Validate(s)
}
