// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iModifyModelApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *ModifyModelApiResponseBody
	GetCategory() *string
	SetGmtCreated(v string) *ModifyModelApiResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *ModifyModelApiResponseBody
	GetGmtModified() *string
	SetModelApiId(v string) *ModifyModelApiResponseBody
	GetModelApiId() *string
	SetModelMappings(v string) *ModifyModelApiResponseBody
	GetModelMappings() *string
	SetModelRouting(v string) *ModifyModelApiResponseBody
	GetModelRouting() *string
	SetName(v string) *ModifyModelApiResponseBody
	GetName() *string
	SetPathPrefix(v string) *ModifyModelApiResponseBody
	GetPathPrefix() *string
	SetProtocol(v string) *ModifyModelApiResponseBody
	GetProtocol() *string
	SetRequestId(v string) *ModifyModelApiResponseBody
	GetRequestId() *string
	SetRouteRules(v string) *ModifyModelApiResponseBody
	GetRouteRules() *string
	SetStatus(v string) *ModifyModelApiResponseBody
	GetStatus() *string
}

type ModifyModelApiResponseBody struct {
	// The category of the model.
	//
	// example:
	//
	// text
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The time when the model API was created.
	//
	// example:
	//
	// 2026-11-09T06:43:33Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The time when the model API was last updated.
	//
	// example:
	//
	// 2026-11-09T06:43:33Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The ID of the model API.
	//
	// example:
	//
	// mi-xxxxx
	ModelApiId *string `json:"ModelApiId,omitempty" xml:"ModelApiId,omitempty"`
	// The current model name mappings, returned as a JSON array string. This parameter is empty if no mappings are configured.
	//
	// example:
	//
	// [{"Source":"client-model-a","Target":"upstream-model-a"}]
	ModelMappings *string `json:"ModelMappings,omitempty" xml:"ModelMappings,omitempty"`
	// The current intelligent routing candidate list, returned as a JSON object string.
	//
	// example:
	//
	// {"AllowedModels":["deepseek-v4-pro","glm-5.3","glm-5.3-flash"]}
	ModelRouting *string `json:"ModelRouting,omitempty" xml:"ModelRouting,omitempty"`
	// The updated name of the model API.
	//
	// example:
	//
	// jev-chat-api
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The API path prefix.
	//
	// example:
	//
	// /jev/v1
	PathPrefix *string `json:"PathPrefix,omitempty" xml:"PathPrefix,omitempty"`
	// The access protocol.
	//
	// example:
	//
	// openai-chat
	Protocol *string `json:"Protocol,omitempty" xml:"Protocol,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 925B84D9-CA72-432C-95CF-738C22******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The current routing rules, returned as a JSON string. A byte-for-byte exact match of the input is not guaranteed.
	//
	// example:
	//
	// [{"RuleName":"default","Providers":[{"ModelServiceName":"model-service-a","ModelList":"[\\\\"deepseek-v4-pro\\\\"]","Weight":100}]}]
	RouteRules *string `json:"RouteRules,omitempty" xml:"RouteRules,omitempty"`
	// The status of the model API.
	//
	// example:
	//
	// Running
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s ModifyModelApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ModifyModelApiResponseBody) GoString() string {
	return s.String()
}

func (s *ModifyModelApiResponseBody) GetCategory() *string {
	return s.Category
}

func (s *ModifyModelApiResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *ModifyModelApiResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *ModifyModelApiResponseBody) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *ModifyModelApiResponseBody) GetModelMappings() *string {
	return s.ModelMappings
}

func (s *ModifyModelApiResponseBody) GetModelRouting() *string {
	return s.ModelRouting
}

func (s *ModifyModelApiResponseBody) GetName() *string {
	return s.Name
}

func (s *ModifyModelApiResponseBody) GetPathPrefix() *string {
	return s.PathPrefix
}

func (s *ModifyModelApiResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *ModifyModelApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ModifyModelApiResponseBody) GetRouteRules() *string {
	return s.RouteRules
}

func (s *ModifyModelApiResponseBody) GetStatus() *string {
	return s.Status
}

func (s *ModifyModelApiResponseBody) SetCategory(v string) *ModifyModelApiResponseBody {
	s.Category = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetGmtCreated(v string) *ModifyModelApiResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetGmtModified(v string) *ModifyModelApiResponseBody {
	s.GmtModified = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetModelApiId(v string) *ModifyModelApiResponseBody {
	s.ModelApiId = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetModelMappings(v string) *ModifyModelApiResponseBody {
	s.ModelMappings = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetModelRouting(v string) *ModifyModelApiResponseBody {
	s.ModelRouting = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetName(v string) *ModifyModelApiResponseBody {
	s.Name = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetPathPrefix(v string) *ModifyModelApiResponseBody {
	s.PathPrefix = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetProtocol(v string) *ModifyModelApiResponseBody {
	s.Protocol = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetRequestId(v string) *ModifyModelApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetRouteRules(v string) *ModifyModelApiResponseBody {
	s.RouteRules = &v
	return s
}

func (s *ModifyModelApiResponseBody) SetStatus(v string) *ModifyModelApiResponseBody {
	s.Status = &v
	return s
}

func (s *ModifyModelApiResponseBody) Validate() error {
	return dara.Validate(s)
}
