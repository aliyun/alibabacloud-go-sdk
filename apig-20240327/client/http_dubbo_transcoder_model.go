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
	SetMethodMapList(v []*HttpDubboTranscoderMethodMapList) *HttpDubboTranscoder
	GetMethodMapList() []*HttpDubboTranscoderMethodMapList
}

type HttpDubboTranscoder struct {
	// The Dubbo service group.
	DubboServiceGroup *string `json:"dubboServiceGroup,omitempty" xml:"dubboServiceGroup,omitempty"`
	// The Dubbo service name.
	DubboServiceName *string `json:"dubboServiceName,omitempty" xml:"dubboServiceName,omitempty"`
	// The Dubbo service version.
	DubboServiceVersion *string `json:"dubboServiceVersion,omitempty" xml:"dubboServiceVersion,omitempty"`
	// The method mapping list.
	MethodMapList []*HttpDubboTranscoderMethodMapList `json:"methodMapList,omitempty" xml:"methodMapList,omitempty" type:"Repeated"`
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

func (s *HttpDubboTranscoder) GetMethodMapList() []*HttpDubboTranscoderMethodMapList {
	return s.MethodMapList
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

func (s *HttpDubboTranscoder) SetMethodMapList(v []*HttpDubboTranscoderMethodMapList) *HttpDubboTranscoder {
	s.MethodMapList = v
	return s
}

func (s *HttpDubboTranscoder) Validate() error {
	if s.MethodMapList != nil {
		for _, item := range s.MethodMapList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type HttpDubboTranscoderMethodMapList struct {
	// The Dubbo method name.
	DubboMethodName *string `json:"dubboMethodName,omitempty" xml:"dubboMethodName,omitempty"`
	// The HTTP method. Valid values: ALL_GET. ALL_POST. ALL_PUT. ALL_DELETE. ALL_PATCH.
	//
	// example:
	//
	// ALL_GET
	HttpMethod *string `json:"httpMethod,omitempty" xml:"httpMethod,omitempty"`
	// The method matching path.
	//
	// example:
	//
	// /mytestzbk/sayhello
	MethodPath *string `json:"methodPath,omitempty" xml:"methodPath,omitempty"`
	// The parameter mapping list.
	ParamMapsList []*HttpDubboTranscoderMethodMapListParamMapsList `json:"paramMapsList,omitempty" xml:"paramMapsList,omitempty" type:"Repeated"`
	// The header pass-through type. Valid values: PASS_ALL: passes through all headers. PASS_NOT: does not pass through any headers. PASS_ASSIGN: passes through specified headers.
	//
	// example:
	//
	// PASS_NOT
	PassThroughAllHeaders *string `json:"passThroughAllHeaders,omitempty" xml:"passThroughAllHeaders,omitempty"`
	// The list of specified pass-through headers.
	PassThroughList []*string `json:"passThroughList,omitempty" xml:"passThroughList,omitempty" type:"Repeated"`
}

func (s HttpDubboTranscoderMethodMapList) String() string {
	return dara.Prettify(s)
}

func (s HttpDubboTranscoderMethodMapList) GoString() string {
	return s.String()
}

func (s *HttpDubboTranscoderMethodMapList) GetDubboMethodName() *string {
	return s.DubboMethodName
}

func (s *HttpDubboTranscoderMethodMapList) GetHttpMethod() *string {
	return s.HttpMethod
}

func (s *HttpDubboTranscoderMethodMapList) GetMethodPath() *string {
	return s.MethodPath
}

func (s *HttpDubboTranscoderMethodMapList) GetParamMapsList() []*HttpDubboTranscoderMethodMapListParamMapsList {
	return s.ParamMapsList
}

func (s *HttpDubboTranscoderMethodMapList) GetPassThroughAllHeaders() *string {
	return s.PassThroughAllHeaders
}

func (s *HttpDubboTranscoderMethodMapList) GetPassThroughList() []*string {
	return s.PassThroughList
}

func (s *HttpDubboTranscoderMethodMapList) SetDubboMethodName(v string) *HttpDubboTranscoderMethodMapList {
	s.DubboMethodName = &v
	return s
}

func (s *HttpDubboTranscoderMethodMapList) SetHttpMethod(v string) *HttpDubboTranscoderMethodMapList {
	s.HttpMethod = &v
	return s
}

func (s *HttpDubboTranscoderMethodMapList) SetMethodPath(v string) *HttpDubboTranscoderMethodMapList {
	s.MethodPath = &v
	return s
}

func (s *HttpDubboTranscoderMethodMapList) SetParamMapsList(v []*HttpDubboTranscoderMethodMapListParamMapsList) *HttpDubboTranscoderMethodMapList {
	s.ParamMapsList = v
	return s
}

func (s *HttpDubboTranscoderMethodMapList) SetPassThroughAllHeaders(v string) *HttpDubboTranscoderMethodMapList {
	s.PassThroughAllHeaders = &v
	return s
}

func (s *HttpDubboTranscoderMethodMapList) SetPassThroughList(v []*string) *HttpDubboTranscoderMethodMapList {
	s.PassThroughList = v
	return s
}

func (s *HttpDubboTranscoderMethodMapList) Validate() error {
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

type HttpDubboTranscoderMethodMapListParamMapsList struct {
	// The key used to extract the input parameter.
	//
	// example:
	//
	// name
	ExtractKey *string `json:"extractKey,omitempty" xml:"extractKey,omitempty"`
	// The input parameter location. Valid values: ALL_QUERY_PARAMETER: request parameter. ALL_HEADER: request header. ALL_PATH: URI of the request. ALL_BODY: request body.
	//
	// example:
	//
	// ALL_QUERY_PARAMETER
	ExtractKeySpec *string `json:"extractKeySpec,omitempty" xml:"extractKeySpec,omitempty"`
	// The backend parameter type.
	//
	// example:
	//
	// java.lang.String
	MappingType *string `json:"mappingType,omitempty" xml:"mappingType,omitempty"`
}

func (s HttpDubboTranscoderMethodMapListParamMapsList) String() string {
	return dara.Prettify(s)
}

func (s HttpDubboTranscoderMethodMapListParamMapsList) GoString() string {
	return s.String()
}

func (s *HttpDubboTranscoderMethodMapListParamMapsList) GetExtractKey() *string {
	return s.ExtractKey
}

func (s *HttpDubboTranscoderMethodMapListParamMapsList) GetExtractKeySpec() *string {
	return s.ExtractKeySpec
}

func (s *HttpDubboTranscoderMethodMapListParamMapsList) GetMappingType() *string {
	return s.MappingType
}

func (s *HttpDubboTranscoderMethodMapListParamMapsList) SetExtractKey(v string) *HttpDubboTranscoderMethodMapListParamMapsList {
	s.ExtractKey = &v
	return s
}

func (s *HttpDubboTranscoderMethodMapListParamMapsList) SetExtractKeySpec(v string) *HttpDubboTranscoderMethodMapListParamMapsList {
	s.ExtractKeySpec = &v
	return s
}

func (s *HttpDubboTranscoderMethodMapListParamMapsList) SetMappingType(v string) *HttpDubboTranscoderMethodMapListParamMapsList {
	s.MappingType = &v
	return s
}

func (s *HttpDubboTranscoderMethodMapListParamMapsList) Validate() error {
	return dara.Validate(s)
}
