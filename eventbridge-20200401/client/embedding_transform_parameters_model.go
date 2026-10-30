// This file is auto-generated, don't edit it. Thanks.
package client

import (
  "github.com/alibabacloud-go/tea/dara"
)

type iEmbeddingTransformParameters interface {
  dara.Model
  String() string
  GoString() string
  SetApiKey(v string) *EmbeddingTransformParameters
  GetApiKey() *string 
  SetEmbeddingData(v *EmbeddingTransformParametersEmbeddingData) *EmbeddingTransformParameters
  GetEmbeddingData() *EmbeddingTransformParametersEmbeddingData 
  SetEmbeddingModel(v string) *EmbeddingTransformParameters
  GetEmbeddingModel() *string 
}

type EmbeddingTransformParameters struct {
  // The API key used to call the Model Studio embedding model service.
  // 
  // example:
  // 
  // Q37sQ****
  ApiKey *string `json:"ApiKey,omitempty" xml:"ApiKey,omitempty"`
  // The input data to be vectorized and its transformation method.
  EmbeddingData *EmbeddingTransformParametersEmbeddingData `json:"EmbeddingData,omitempty" xml:"EmbeddingData,omitempty" type:"Struct"`
  // The name of the model used to generate embeddings.
  // 
  // example:
  // 
  // text-embedding-v3
  EmbeddingModel *string `json:"EmbeddingModel,omitempty" xml:"EmbeddingModel,omitempty"`
}

func (s EmbeddingTransformParameters) String() string {
  return dara.Prettify(s)
}

func (s EmbeddingTransformParameters) GoString() string {
  return s.String()
}

func (s *EmbeddingTransformParameters) GetApiKey() *string  {
  return s.ApiKey
}

func (s *EmbeddingTransformParameters) GetEmbeddingData() *EmbeddingTransformParametersEmbeddingData  {
  return s.EmbeddingData
}

func (s *EmbeddingTransformParameters) GetEmbeddingModel() *string  {
  return s.EmbeddingModel
}

func (s *EmbeddingTransformParameters) SetApiKey(v string) *EmbeddingTransformParameters {
  s.ApiKey = &v
  return s
}

func (s *EmbeddingTransformParameters) SetEmbeddingData(v *EmbeddingTransformParametersEmbeddingData) *EmbeddingTransformParameters {
  s.EmbeddingData = v
  return s
}

func (s *EmbeddingTransformParameters) SetEmbeddingModel(v string) *EmbeddingTransformParameters {
  s.EmbeddingModel = &v
  return s
}

func (s *EmbeddingTransformParameters) Validate() error {
  if s.EmbeddingData != nil {
    if err := s.EmbeddingData.Validate(); err != nil {
      return err
    }
  }
  return nil
}

type EmbeddingTransformParametersEmbeddingData struct {
  // The method used to generate the input data. JSONPATH indicates that data is extracted from the event by using JSONPath.
  // 
  // example:
  // 
  // JSONPATH
  Form *string `json:"Form,omitempty" xml:"Form,omitempty"`
  // The template used to build the input data. This parameter is used only when Form is set to TEMPLATE.
  // 
  // example:
  // 
  // The value of ${key} is ${value}!
  Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
  // The source content of the input data. If Form is set to JSONPATH, specify a JSONPath expression.
  // 
  // example:
  // 
  // $.data.message
  Value *string `json:"Value,omitempty" xml:"Value,omitempty"`
}

func (s EmbeddingTransformParametersEmbeddingData) String() string {
  return dara.Prettify(s)
}

func (s EmbeddingTransformParametersEmbeddingData) GoString() string {
  return s.String()
}

func (s *EmbeddingTransformParametersEmbeddingData) GetForm() *string  {
  return s.Form
}

func (s *EmbeddingTransformParametersEmbeddingData) GetTemplate() *string  {
  return s.Template
}

func (s *EmbeddingTransformParametersEmbeddingData) GetValue() *string  {
  return s.Value
}

func (s *EmbeddingTransformParametersEmbeddingData) SetForm(v string) *EmbeddingTransformParametersEmbeddingData {
  s.Form = &v
  return s
}

func (s *EmbeddingTransformParametersEmbeddingData) SetTemplate(v string) *EmbeddingTransformParametersEmbeddingData {
  s.Template = &v
  return s
}

func (s *EmbeddingTransformParametersEmbeddingData) SetValue(v string) *EmbeddingTransformParametersEmbeddingData {
  s.Value = &v
  return s
}

func (s *EmbeddingTransformParametersEmbeddingData) Validate() error {
  return dara.Validate(s)
}

