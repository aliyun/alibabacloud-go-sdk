// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListModelsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetMaxResults(v int32) *ListModelsResponseBody
	GetMaxResults() *int32
	SetModels(v []*ListModelsResponseBodyModels) *ListModelsResponseBody
	GetModels() []*ListModelsResponseBodyModels
	SetNextToken(v string) *ListModelsResponseBody
	GetNextToken() *string
	SetRequestId(v string) *ListModelsResponseBody
	GetRequestId() *string
	SetTotal(v int64) *ListModelsResponseBody
	GetTotal() *int64
}

type ListModelsResponseBody struct {
	// The maximum number of results used for this pagination. If not paginated, it is the total number of results. It is not the actual number of entries returned on this page.
	//
	// example:
	//
	// 20
	MaxResults *int32 `json:"MaxResults,omitempty" xml:"MaxResults,omitempty"`
	// The array of model structures. For nested fields, refer to the supplementary description of response parameters.
	//
	// example:
	//
	// [{"Id":"text-model-example","Object":"model","Created":1793491200,"OwnedBy":"provider-example","DisplayName":"Sample text model","Capabilities":{"InputModalities":["text"]},"CreditsCoefficient":"2.5","ThinkingLevels":["off","low","medium","high"],"ThinkingLevel":"off"}]
	Models []*ListModelsResponseBodyModels `json:"Models,omitempty" xml:"Models,omitempty" type:"Repeated"`
	// The token for the next page. An empty value indicates that there is no next page.
	//
	// example:
	//
	// next-token-example
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The ID of the request. It is used for troubleshooting.
	//
	// example:
	//
	// 0A1B2C3D-4E5F-6789-ABCD-EF0123456789
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of models that meet the filter conditions.
	//
	// example:
	//
	// 1
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
}

func (s ListModelsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBody) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBody) GetMaxResults() *int32 {
	return s.MaxResults
}

func (s *ListModelsResponseBody) GetModels() []*ListModelsResponseBodyModels {
	return s.Models
}

func (s *ListModelsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *ListModelsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListModelsResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *ListModelsResponseBody) SetMaxResults(v int32) *ListModelsResponseBody {
	s.MaxResults = &v
	return s
}

func (s *ListModelsResponseBody) SetModels(v []*ListModelsResponseBodyModels) *ListModelsResponseBody {
	s.Models = v
	return s
}

func (s *ListModelsResponseBody) SetNextToken(v string) *ListModelsResponseBody {
	s.NextToken = &v
	return s
}

func (s *ListModelsResponseBody) SetRequestId(v string) *ListModelsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListModelsResponseBody) SetTotal(v int64) *ListModelsResponseBody {
	s.Total = &v
	return s
}

func (s *ListModelsResponseBody) Validate() error {
	if s.Models != nil {
		for _, item := range s.Models {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListModelsResponseBodyModels struct {
	// The model capabilities structure.
	//
	// example:
	//
	// {"InputModalities":["text"]}
	Capabilities *ListModelsResponseBodyModelsCapabilities `json:"Capabilities,omitempty" xml:"Capabilities,omitempty" type:"Struct"`
	// Deprecated
	//
	// **[Deprecated]*	- The legacy compatibility field. It is no longer returned. Do not use it as a dependency for integration.
	//
	// example:
	//
	// 32768
	ContextWindow *int64 `json:"ContextWindow,omitempty" xml:"ContextWindow,omitempty"`
	// The time when the model was created. It is a UNIX timestamp in seconds.
	//
	// example:
	//
	// 1793491200
	Created *int64 `json:"Created,omitempty" xml:"Created,omitempty"`
	// The credit coefficient of the model, returned as a decimal string. Not returned if not configured.
	//
	// example:
	//
	// 2.5
	CreditsCoefficient *string `json:"CreditsCoefficient,omitempty" xml:"CreditsCoefficient,omitempty"`
	// The display name of the model. Do not pass it as the model identifier.
	//
	// example:
	//
	// Sample text model
	DisplayName *string `json:"DisplayName,omitempty" xml:"DisplayName,omitempty"`
	// The model identifier. Used as Params.modelId for SetAgentSessionModel. Use the actual returned value.
	//
	// example:
	//
	// text-model-example
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The legacy compatibility field. It is no longer returned. Do not use it as a dependency for integration.
	//
	// example:
	//
	// 131072
	MaxContextWindow *int64 `json:"MaxContextWindow,omitempty" xml:"MaxContextWindow,omitempty"`
	// The object type. Fixed to model.
	//
	// example:
	//
	// model
	Object *string `json:"Object,omitempty" xml:"Object,omitempty"`
	// The identifier of the model provider.
	//
	// example:
	//
	// provider-example
	OwnedBy *string `json:"OwnedBy,omitempty" xml:"OwnedBy,omitempty"`
	// Deprecated
	//
	// **[Deprecated]*	- The legacy compatibility field. It is no longer returned. Use ThinkingLevels instead.
	//
	// example:
	//
	// 1
	SupportsThinking *int32 `json:"SupportsThinking,omitempty" xml:"SupportsThinking,omitempty"`
	// The default thinking level in the model catalog. It does not represent the current value of the session. Query the current value by using GetAgentSessionModel.
	//
	// example:
	//
	// off
	ThinkingLevel *string `json:"ThinkingLevel,omitempty" xml:"ThinkingLevel,omitempty"`
	// The thinking levels supported by the model. An empty array does not mean all levels are supported. The target session still needs to be verified.
	//
	// example:
	//
	// ["off","low","medium","high"]
	ThinkingLevels []*string `json:"ThinkingLevels,omitempty" xml:"ThinkingLevels,omitempty" type:"Repeated"`
	// Deprecated
	//
	// **[Deprecated]*	- The legacy compatibility field. It is no longer returned. Use ThinkingLevel instead.
	//
	// example:
	//
	// 0
	ThinkingMode *int32 `json:"ThinkingMode,omitempty" xml:"ThinkingMode,omitempty"`
}

func (s ListModelsResponseBodyModels) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyModels) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyModels) GetCapabilities() *ListModelsResponseBodyModelsCapabilities {
	return s.Capabilities
}

func (s *ListModelsResponseBodyModels) GetContextWindow() *int64 {
	return s.ContextWindow
}

func (s *ListModelsResponseBodyModels) GetCreated() *int64 {
	return s.Created
}

func (s *ListModelsResponseBodyModels) GetCreditsCoefficient() *string {
	return s.CreditsCoefficient
}

func (s *ListModelsResponseBodyModels) GetDisplayName() *string {
	return s.DisplayName
}

func (s *ListModelsResponseBodyModels) GetId() *string {
	return s.Id
}

func (s *ListModelsResponseBodyModels) GetMaxContextWindow() *int64 {
	return s.MaxContextWindow
}

func (s *ListModelsResponseBodyModels) GetObject() *string {
	return s.Object
}

func (s *ListModelsResponseBodyModels) GetOwnedBy() *string {
	return s.OwnedBy
}

func (s *ListModelsResponseBodyModels) GetSupportsThinking() *int32 {
	return s.SupportsThinking
}

func (s *ListModelsResponseBodyModels) GetThinkingLevel() *string {
	return s.ThinkingLevel
}

func (s *ListModelsResponseBodyModels) GetThinkingLevels() []*string {
	return s.ThinkingLevels
}

func (s *ListModelsResponseBodyModels) GetThinkingMode() *int32 {
	return s.ThinkingMode
}

func (s *ListModelsResponseBodyModels) SetCapabilities(v *ListModelsResponseBodyModelsCapabilities) *ListModelsResponseBodyModels {
	s.Capabilities = v
	return s
}

func (s *ListModelsResponseBodyModels) SetContextWindow(v int64) *ListModelsResponseBodyModels {
	s.ContextWindow = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetCreated(v int64) *ListModelsResponseBodyModels {
	s.Created = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetCreditsCoefficient(v string) *ListModelsResponseBodyModels {
	s.CreditsCoefficient = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetDisplayName(v string) *ListModelsResponseBodyModels {
	s.DisplayName = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetId(v string) *ListModelsResponseBodyModels {
	s.Id = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetMaxContextWindow(v int64) *ListModelsResponseBodyModels {
	s.MaxContextWindow = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetObject(v string) *ListModelsResponseBodyModels {
	s.Object = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetOwnedBy(v string) *ListModelsResponseBodyModels {
	s.OwnedBy = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetSupportsThinking(v int32) *ListModelsResponseBodyModels {
	s.SupportsThinking = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetThinkingLevel(v string) *ListModelsResponseBodyModels {
	s.ThinkingLevel = &v
	return s
}

func (s *ListModelsResponseBodyModels) SetThinkingLevels(v []*string) *ListModelsResponseBodyModels {
	s.ThinkingLevels = v
	return s
}

func (s *ListModelsResponseBodyModels) SetThinkingMode(v int32) *ListModelsResponseBodyModels {
	s.ThinkingMode = &v
	return s
}

func (s *ListModelsResponseBodyModels) Validate() error {
	if s.Capabilities != nil {
		if err := s.Capabilities.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type ListModelsResponseBodyModelsCapabilities struct {
	// The supported input modalities. For text sessions, select a model that includes text.
	//
	// example:
	//
	// ["text"]
	InputModalities []*string `json:"InputModalities,omitempty" xml:"InputModalities,omitempty" type:"Repeated"`
}

func (s ListModelsResponseBodyModelsCapabilities) String() string {
	return dara.Prettify(s)
}

func (s ListModelsResponseBodyModelsCapabilities) GoString() string {
	return s.String()
}

func (s *ListModelsResponseBodyModelsCapabilities) GetInputModalities() []*string {
	return s.InputModalities
}

func (s *ListModelsResponseBodyModelsCapabilities) SetInputModalities(v []*string) *ListModelsResponseBodyModelsCapabilities {
	s.InputModalities = v
	return s
}

func (s *ListModelsResponseBodyModelsCapabilities) Validate() error {
	return dara.Validate(s)
}
