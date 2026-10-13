// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelCardsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *ListModelCardsResponseBody
	GetCode() *string
	SetData(v []*ListModelCardsResponseBodyData) *ListModelCardsResponseBody
	GetData() []*ListModelCardsResponseBodyData
	SetMessage(v string) *ListModelCardsResponseBody
	GetMessage() *string
	SetRequestId(v string) *ListModelCardsResponseBody
	GetRequestId() *string
}

type ListModelCardsResponseBody struct {
	// The response status code.
	//
	// example:
	//
	// Ok
	Code *string `json:"code,omitempty" xml:"code,omitempty"`
	// The response data.
	Data []*ListModelCardsResponseBodyData `json:"data,omitempty" xml:"data,omitempty" type:"Repeated"`
	// The response message.
	//
	// example:
	//
	// success
	Message *string `json:"message,omitempty" xml:"message,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 393E2630-DBE7-5221-AB35-9E740675491A
	RequestId *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s ListModelCardsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelCardsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelCardsResponseBody) GetCode() *string {
	return s.Code
}

func (s *ListModelCardsResponseBody) GetData() []*ListModelCardsResponseBodyData {
	return s.Data
}

func (s *ListModelCardsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *ListModelCardsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelCardsResponseBody) SetCode(v string) *ListModelCardsResponseBody {
	s.Code = &v
	return s
}

func (s *ListModelCardsResponseBody) SetData(v []*ListModelCardsResponseBodyData) *ListModelCardsResponseBody {
	s.Data = v
	return s
}

func (s *ListModelCardsResponseBody) SetMessage(v string) *ListModelCardsResponseBody {
	s.Message = &v
	return s
}

func (s *ListModelCardsResponseBody) SetRequestId(v string) *ListModelCardsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelCardsResponseBody) Validate() error {
	if s.Data != nil {
		for _, item := range s.Data {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelCardsResponseBodyData struct {
	// The list of available paths for the model card. Each element consists of a path and a type member. The API definition does not specify the criteria for determining availability.
	AvailablePaths []*ListModelCardsResponseBodyDataAvailablePaths `json:"availablePaths,omitempty" xml:"availablePaths,omitempty" type:"Repeated"`
	// The feature configuration of the model card in a key-value pair map structure. The set of keys is not specified in the API definition.
	Features map[string]interface{} `json:"features,omitempty" xml:"features,omitempty"`
	// The metadata container of the model card. Only the description member is declared in the API definition.
	Meta *ListModelCardsResponseBodyDataMeta `json:"meta,omitempty" xml:"meta,omitempty" type:"Struct"`
	// The name of the model.
	//
	// example:
	//
	// qwen-plus
	ModelName *string `json:"modelName,omitempty" xml:"modelName,omitempty"`
	// The pricing information of the model card in a key-value pair map structure. The set of keys and billing criteria are not specified in the API definition.
	Pricing map[string]interface{} `json:"pricing,omitempty" xml:"pricing,omitempty"`
	// The identifier of the model provider for the model card.
	//
	// example:
	//
	// provider-xxxxxxxx
	Provider *string `json:"provider,omitempty" xml:"provider,omitempty"`
	// The list of snapshots for the model card. Array items are map structures with unspecified members. The API definition does not specify the snapshot content or generation rules.
	Snapshots []map[string]interface{} `json:"snapshots,omitempty" xml:"snapshots,omitempty" type:"Repeated"`
	// The source identifier of the model card. The valid values are not specified in the API definition.
	//
	// example:
	//
	// source-xxxxxxxx
	Source *string `json:"source,omitempty" xml:"source,omitempty"`
}

func (s ListModelCardsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s ListModelCardsResponseBodyData) GoString() string {
	return s.String()
}

func (s *ListModelCardsResponseBodyData) GetAvailablePaths() []*ListModelCardsResponseBodyDataAvailablePaths {
	return s.AvailablePaths
}

func (s *ListModelCardsResponseBodyData) GetFeatures() map[string]interface{} {
	return s.Features
}

func (s *ListModelCardsResponseBodyData) GetMeta() *ListModelCardsResponseBodyDataMeta {
	return s.Meta
}

func (s *ListModelCardsResponseBodyData) GetModelName() *string {
	return s.ModelName
}

func (s *ListModelCardsResponseBodyData) GetPricing() map[string]interface{} {
	return s.Pricing
}

func (s *ListModelCardsResponseBodyData) GetProvider() *string {
	return s.Provider
}

func (s *ListModelCardsResponseBodyData) GetSnapshots() []map[string]interface{} {
	return s.Snapshots
}

func (s *ListModelCardsResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *ListModelCardsResponseBodyData) SetAvailablePaths(v []*ListModelCardsResponseBodyDataAvailablePaths) *ListModelCardsResponseBodyData {
	s.AvailablePaths = v
	return s
}

func (s *ListModelCardsResponseBodyData) SetFeatures(v map[string]interface{}) *ListModelCardsResponseBodyData {
	s.Features = v
	return s
}

func (s *ListModelCardsResponseBodyData) SetMeta(v *ListModelCardsResponseBodyDataMeta) *ListModelCardsResponseBodyData {
	s.Meta = v
	return s
}

func (s *ListModelCardsResponseBodyData) SetModelName(v string) *ListModelCardsResponseBodyData {
	s.ModelName = &v
	return s
}

func (s *ListModelCardsResponseBodyData) SetPricing(v map[string]interface{}) *ListModelCardsResponseBodyData {
	s.Pricing = v
	return s
}

func (s *ListModelCardsResponseBodyData) SetProvider(v string) *ListModelCardsResponseBodyData {
	s.Provider = &v
	return s
}

func (s *ListModelCardsResponseBodyData) SetSnapshots(v []map[string]interface{}) *ListModelCardsResponseBodyData {
	s.Snapshots = v
	return s
}

func (s *ListModelCardsResponseBodyData) SetSource(v string) *ListModelCardsResponseBodyData {
	s.Source = &v
	return s
}

func (s *ListModelCardsResponseBodyData) Validate() error {
	if s.AvailablePaths != nil {
		for _, item := range s.AvailablePaths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Meta != nil {
		if err := s.Meta.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListModelCardsResponseBodyDataAvailablePaths struct {
	// The request path of the API.
	//
	// example:
	//
	// /v1/chat/completions
	Path *string `json:"path,omitempty" xml:"path,omitempty"`
	// The type identifier of the available path record. The valid values are not specified in the API definition.
	//
	// example:
	//
	// type-xxxxxxxx
	Type *string `json:"type,omitempty" xml:"type,omitempty"`
}

func (s ListModelCardsResponseBodyDataAvailablePaths) String() string {
	return dara.Prettify(s)
}

func (s ListModelCardsResponseBodyDataAvailablePaths) GoString() string {
	return s.String()
}

func (s *ListModelCardsResponseBodyDataAvailablePaths) GetPath() *string {
	return s.Path
}

func (s *ListModelCardsResponseBodyDataAvailablePaths) GetType() *string {
	return s.Type
}

func (s *ListModelCardsResponseBodyDataAvailablePaths) SetPath(v string) *ListModelCardsResponseBodyDataAvailablePaths {
	s.Path = &v
	return s
}

func (s *ListModelCardsResponseBodyDataAvailablePaths) SetType(v string) *ListModelCardsResponseBodyDataAvailablePaths {
	s.Type = &v
	return s
}

func (s *ListModelCardsResponseBodyDataAvailablePaths) Validate() error {
	return dara.Validate(s)
}

type ListModelCardsResponseBodyDataMeta struct {
	// The description text of the model card.
	//
	// example:
	//
	// Sample description text for the model card
	Description *string `json:"description,omitempty" xml:"description,omitempty"`
}

func (s ListModelCardsResponseBodyDataMeta) String() string {
	return dara.Prettify(s)
}

func (s ListModelCardsResponseBodyDataMeta) GoString() string {
	return s.String()
}

func (s *ListModelCardsResponseBodyDataMeta) GetDescription() *string {
	return s.Description
}

func (s *ListModelCardsResponseBodyDataMeta) SetDescription(v string) *ListModelCardsResponseBodyDataMeta {
	s.Description = &v
	return s
}

func (s *ListModelCardsResponseBodyDataMeta) Validate() error {
	return dara.Validate(s)
}
