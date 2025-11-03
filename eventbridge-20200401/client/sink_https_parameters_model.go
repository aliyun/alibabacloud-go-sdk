// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSinkHttpsParameters interface {
	dara.Model
	String() string
	GoString() string
	SetBody(v *SinkHttpsParametersBody) *SinkHttpsParameters
	GetBody() *SinkHttpsParametersBody
	SetMethod(v string) *SinkHttpsParameters
	GetMethod() *string
	SetNetworkType(v string) *SinkHttpsParameters
	GetNetworkType() *string
	SetSecurityGroupId(v string) *SinkHttpsParameters
	GetSecurityGroupId() *string
	SetToken(v string) *SinkHttpsParameters
	GetToken() *string
	SetURL(v *SinkHttpsParametersURL) *SinkHttpsParameters
	GetURL() *SinkHttpsParametersURL
	SetVSwitchIds(v string) *SinkHttpsParameters
	GetVSwitchIds() *string
	SetVpcId(v string) *SinkHttpsParameters
	GetVpcId() *string
}

type SinkHttpsParameters struct {
	Body            *SinkHttpsParametersBody `json:"Body,omitempty" xml:"Body,omitempty" type:"Struct"`
	Method          *string                  `json:"Method,omitempty" xml:"Method,omitempty"`
	NetworkType     *string                  `json:"NetworkType,omitempty" xml:"NetworkType,omitempty"`
	SecurityGroupId *string                  `json:"SecurityGroupId,omitempty" xml:"SecurityGroupId,omitempty"`
	Token           *string                  `json:"Token,omitempty" xml:"Token,omitempty"`
	URL             *SinkHttpsParametersURL  `json:"URL,omitempty" xml:"URL,omitempty" type:"Struct"`
	VSwitchIds      *string                  `json:"VSwitchIds,omitempty" xml:"VSwitchIds,omitempty"`
	VpcId           *string                  `json:"VpcId,omitempty" xml:"VpcId,omitempty"`
}

func (s SinkHttpsParameters) String() string {
	return dara.Prettify(s)
}

func (s SinkHttpsParameters) GoString() string {
	return s.String()
}

func (s *SinkHttpsParameters) GetBody() *SinkHttpsParametersBody {
	return s.Body
}

func (s *SinkHttpsParameters) GetMethod() *string {
	return s.Method
}

func (s *SinkHttpsParameters) GetNetworkType() *string {
	return s.NetworkType
}

func (s *SinkHttpsParameters) GetSecurityGroupId() *string {
	return s.SecurityGroupId
}

func (s *SinkHttpsParameters) GetToken() *string {
	return s.Token
}

func (s *SinkHttpsParameters) GetURL() *SinkHttpsParametersURL {
	return s.URL
}

func (s *SinkHttpsParameters) GetVSwitchIds() *string {
	return s.VSwitchIds
}

func (s *SinkHttpsParameters) GetVpcId() *string {
	return s.VpcId
}

func (s *SinkHttpsParameters) SetBody(v *SinkHttpsParametersBody) *SinkHttpsParameters {
	s.Body = v
	return s
}

func (s *SinkHttpsParameters) SetMethod(v string) *SinkHttpsParameters {
	s.Method = &v
	return s
}

func (s *SinkHttpsParameters) SetNetworkType(v string) *SinkHttpsParameters {
	s.NetworkType = &v
	return s
}

func (s *SinkHttpsParameters) SetSecurityGroupId(v string) *SinkHttpsParameters {
	s.SecurityGroupId = &v
	return s
}

func (s *SinkHttpsParameters) SetToken(v string) *SinkHttpsParameters {
	s.Token = &v
	return s
}

func (s *SinkHttpsParameters) SetURL(v *SinkHttpsParametersURL) *SinkHttpsParameters {
	s.URL = v
	return s
}

func (s *SinkHttpsParameters) SetVSwitchIds(v string) *SinkHttpsParameters {
	s.VSwitchIds = &v
	return s
}

func (s *SinkHttpsParameters) SetVpcId(v string) *SinkHttpsParameters {
	s.VpcId = &v
	return s
}

func (s *SinkHttpsParameters) Validate() error {
	if s.Body != nil {
		if err := s.Body.Validate(); err != nil {
			return err
		}
	}
	if s.URL != nil {
		if err := s.URL.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type SinkHttpsParametersBody struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkHttpsParametersBody) String() string {
	return dara.Prettify(s)
}

func (s SinkHttpsParametersBody) GoString() string {
	return s.String()
}

func (s *SinkHttpsParametersBody) GetForm() *string {
	return s.Form
}

func (s *SinkHttpsParametersBody) GetTemplate() *string {
	return s.Template
}

func (s *SinkHttpsParametersBody) GetValue() *string {
	return s.Value
}

func (s *SinkHttpsParametersBody) SetForm(v string) *SinkHttpsParametersBody {
	s.Form = &v
	return s
}

func (s *SinkHttpsParametersBody) SetTemplate(v string) *SinkHttpsParametersBody {
	s.Template = &v
	return s
}

func (s *SinkHttpsParametersBody) SetValue(v string) *SinkHttpsParametersBody {
	s.Value = &v
	return s
}

func (s *SinkHttpsParametersBody) Validate() error {
	return dara.Validate(s)
}

type SinkHttpsParametersURL struct {
	Form     *string `json:"Form,omitempty" xml:"Form,omitempty"`
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
	Value    *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s SinkHttpsParametersURL) String() string {
	return dara.Prettify(s)
}

func (s SinkHttpsParametersURL) GoString() string {
	return s.String()
}

func (s *SinkHttpsParametersURL) GetForm() *string {
	return s.Form
}

func (s *SinkHttpsParametersURL) GetTemplate() *string {
	return s.Template
}

func (s *SinkHttpsParametersURL) GetValue() *string {
	return s.Value
}

func (s *SinkHttpsParametersURL) SetForm(v string) *SinkHttpsParametersURL {
	s.Form = &v
	return s
}

func (s *SinkHttpsParametersURL) SetTemplate(v string) *SinkHttpsParametersURL {
	s.Template = &v
	return s
}

func (s *SinkHttpsParametersURL) SetValue(v string) *SinkHttpsParametersURL {
	s.Value = &v
	return s
}

func (s *SinkHttpsParametersURL) Validate() error {
	return dara.Validate(s)
}
