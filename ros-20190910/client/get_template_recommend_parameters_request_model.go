// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRecommendParametersRequest interface {
	dara.Model
	String() string
	GoString() string
	SetClientToken(v string) *GetTemplateRecommendParametersRequest
	GetClientToken() *string
	SetParameters(v []*GetTemplateRecommendParametersRequestParameters) *GetTemplateRecommendParametersRequest
	GetParameters() []*GetTemplateRecommendParametersRequestParameters
	SetRegionId(v string) *GetTemplateRecommendParametersRequest
	GetRegionId() *string
	SetTemplateBody(v string) *GetTemplateRecommendParametersRequest
	GetTemplateBody() *string
	SetTemplateId(v string) *GetTemplateRecommendParametersRequest
	GetTemplateId() *string
	SetTemplateURL(v string) *GetTemplateRecommendParametersRequest
	GetTemplateURL() *string
	SetTemplateVersion(v string) *GetTemplateRecommendParametersRequest
	GetTemplateVersion() *string
}

type GetTemplateRecommendParametersRequest struct {
	ClientToken *string                                            `json:"ClientToken,omitempty" xml:"ClientToken,omitempty"`
	Parameters  []*GetTemplateRecommendParametersRequestParameters `json:"Parameters,omitempty" xml:"Parameters,omitempty" type:"Repeated"`
	// This parameter is required.
	RegionId        *string `json:"RegionId,omitempty" xml:"RegionId,omitempty"`
	TemplateBody    *string `json:"TemplateBody,omitempty" xml:"TemplateBody,omitempty"`
	TemplateId      *string `json:"TemplateId,omitempty" xml:"TemplateId,omitempty"`
	TemplateURL     *string `json:"TemplateURL,omitempty" xml:"TemplateURL,omitempty"`
	TemplateVersion *string `json:"TemplateVersion,omitempty" xml:"TemplateVersion,omitempty"`
}

func (s GetTemplateRecommendParametersRequest) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRecommendParametersRequest) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersRequest) GetClientToken() *string {
	return s.ClientToken
}

func (s *GetTemplateRecommendParametersRequest) GetParameters() []*GetTemplateRecommendParametersRequestParameters {
	return s.Parameters
}

func (s *GetTemplateRecommendParametersRequest) GetRegionId() *string {
	return s.RegionId
}

func (s *GetTemplateRecommendParametersRequest) GetTemplateBody() *string {
	return s.TemplateBody
}

func (s *GetTemplateRecommendParametersRequest) GetTemplateId() *string {
	return s.TemplateId
}

func (s *GetTemplateRecommendParametersRequest) GetTemplateURL() *string {
	return s.TemplateURL
}

func (s *GetTemplateRecommendParametersRequest) GetTemplateVersion() *string {
	return s.TemplateVersion
}

func (s *GetTemplateRecommendParametersRequest) SetClientToken(v string) *GetTemplateRecommendParametersRequest {
	s.ClientToken = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetParameters(v []*GetTemplateRecommendParametersRequestParameters) *GetTemplateRecommendParametersRequest {
	s.Parameters = v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetRegionId(v string) *GetTemplateRecommendParametersRequest {
	s.RegionId = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetTemplateBody(v string) *GetTemplateRecommendParametersRequest {
	s.TemplateBody = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetTemplateId(v string) *GetTemplateRecommendParametersRequest {
	s.TemplateId = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetTemplateURL(v string) *GetTemplateRecommendParametersRequest {
	s.TemplateURL = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) SetTemplateVersion(v string) *GetTemplateRecommendParametersRequest {
	s.TemplateVersion = &v
	return s
}

func (s *GetTemplateRecommendParametersRequest) Validate() error {
	return dara.Validate(s)
}

type GetTemplateRecommendParametersRequestParameters struct {
	ParameterCandidateValues []*string `json:"ParameterCandidateValues,omitempty" xml:"ParameterCandidateValues,omitempty" type:"Repeated"`
	ParameterKey             *string   `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	ParameterValue           *string   `json:"ParameterValue,omitempty" xml:"ParameterValue,omitempty"`
}

func (s GetTemplateRecommendParametersRequestParameters) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRecommendParametersRequestParameters) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersRequestParameters) GetParameterCandidateValues() []*string {
	return s.ParameterCandidateValues
}

func (s *GetTemplateRecommendParametersRequestParameters) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetTemplateRecommendParametersRequestParameters) GetParameterValue() *string {
	return s.ParameterValue
}

func (s *GetTemplateRecommendParametersRequestParameters) SetParameterCandidateValues(v []*string) *GetTemplateRecommendParametersRequestParameters {
	s.ParameterCandidateValues = v
	return s
}

func (s *GetTemplateRecommendParametersRequestParameters) SetParameterKey(v string) *GetTemplateRecommendParametersRequestParameters {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateRecommendParametersRequestParameters) SetParameterValue(v string) *GetTemplateRecommendParametersRequestParameters {
	s.ParameterValue = &v
	return s
}

func (s *GetTemplateRecommendParametersRequestParameters) Validate() error {
	return dara.Validate(s)
}
