// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetTemplateRecommendParametersResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRecommendParameterValues(v []*GetTemplateRecommendParametersResponseBodyRecommendParameterValues) *GetTemplateRecommendParametersResponseBody
	GetRecommendParameterValues() []*GetTemplateRecommendParametersResponseBodyRecommendParameterValues
	SetRequestId(v string) *GetTemplateRecommendParametersResponseBody
	GetRequestId() *string
}

type GetTemplateRecommendParametersResponseBody struct {
	RecommendParameterValues []*GetTemplateRecommendParametersResponseBodyRecommendParameterValues `json:"RecommendParameterValues,omitempty" xml:"RecommendParameterValues,omitempty" type:"Repeated"`
	RequestId                *string                                                               `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetTemplateRecommendParametersResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRecommendParametersResponseBody) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersResponseBody) GetRecommendParameterValues() []*GetTemplateRecommendParametersResponseBodyRecommendParameterValues {
	return s.RecommendParameterValues
}

func (s *GetTemplateRecommendParametersResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetTemplateRecommendParametersResponseBody) SetRecommendParameterValues(v []*GetTemplateRecommendParametersResponseBodyRecommendParameterValues) *GetTemplateRecommendParametersResponseBody {
	s.RecommendParameterValues = v
	return s
}

func (s *GetTemplateRecommendParametersResponseBody) SetRequestId(v string) *GetTemplateRecommendParametersResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetTemplateRecommendParametersResponseBody) Validate() error {
	if s.RecommendParameterValues != nil {
		for _, item := range s.RecommendParameterValues {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetTemplateRecommendParametersResponseBodyRecommendParameterValues struct {
	ParameterKey   *string `json:"ParameterKey,omitempty" xml:"ParameterKey,omitempty"`
	RecommendValue *string `json:"RecommendValue,omitempty" xml:"RecommendValue,omitempty"`
}

func (s GetTemplateRecommendParametersResponseBodyRecommendParameterValues) String() string {
	return dara.Prettify(s)
}

func (s GetTemplateRecommendParametersResponseBodyRecommendParameterValues) GoString() string {
	return s.String()
}

func (s *GetTemplateRecommendParametersResponseBodyRecommendParameterValues) GetParameterKey() *string {
	return s.ParameterKey
}

func (s *GetTemplateRecommendParametersResponseBodyRecommendParameterValues) GetRecommendValue() *string {
	return s.RecommendValue
}

func (s *GetTemplateRecommendParametersResponseBodyRecommendParameterValues) SetParameterKey(v string) *GetTemplateRecommendParametersResponseBodyRecommendParameterValues {
	s.ParameterKey = &v
	return s
}

func (s *GetTemplateRecommendParametersResponseBodyRecommendParameterValues) SetRecommendValue(v string) *GetTemplateRecommendParametersResponseBodyRecommendParameterValues {
	s.RecommendValue = &v
	return s
}

func (s *GetTemplateRecommendParametersResponseBodyRecommendParameterValues) Validate() error {
	return dara.Validate(s)
}
