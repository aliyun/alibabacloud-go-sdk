// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateModelApiResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCategory(v string) *CreateModelApiResponseBody
	GetCategory() *string
	SetGmtCreated(v string) *CreateModelApiResponseBody
	GetGmtCreated() *string
	SetGmtModified(v string) *CreateModelApiResponseBody
	GetGmtModified() *string
	SetInvokeEndpoint(v string) *CreateModelApiResponseBody
	GetInvokeEndpoint() *string
	SetModelApiId(v string) *CreateModelApiResponseBody
	GetModelApiId() *string
	SetModelMappings(v string) *CreateModelApiResponseBody
	GetModelMappings() *string
	SetName(v string) *CreateModelApiResponseBody
	GetName() *string
	SetPathPrefix(v string) *CreateModelApiResponseBody
	GetPathPrefix() *string
	SetProtocol(v string) *CreateModelApiResponseBody
	GetProtocol() *string
	SetRequestId(v string) *CreateModelApiResponseBody
	GetRequestId() *string
	SetRouteRules(v string) *CreateModelApiResponseBody
	GetRouteRules() *string
	SetStatus(v string) *CreateModelApiResponseBody
	GetStatus() *string
}

type CreateModelApiResponseBody struct {
	// The model category.
	//
	// example:
	//
	// text
	Category *string `json:"Category,omitempty" xml:"Category,omitempty"`
	// The creation time.
	//
	// example:
	//
	// 2026-11-09T06:43:33Z
	GmtCreated *string `json:"GmtCreated,omitempty" xml:"GmtCreated,omitempty"`
	// The modification time.
	//
	// example:
	//
	// 2026-11-09T06:43:33Z
	GmtModified *string `json:"GmtModified,omitempty" xml:"GmtModified,omitempty"`
	// The invocation endpoint.
	//
	// example:
	//
	// https://model-api.example.com/v1/chat/completions
	InvokeEndpoint *string `json:"InvokeEndpoint,omitempty" xml:"InvokeEndpoint,omitempty"`
	// The model API ID.
	//
	// example:
	//
	// mi-xxxxxx
	ModelApiId *string `json:"ModelApiId,omitempty" xml:"ModelApiId,omitempty"`
	// The saved model name mappings, returned as a JSON array string. This parameter is empty if no mappings are configured.
	//
	// example:
	//
	// [{"Source":"client-model-a","Target":"upstream-model-a"}]
	ModelMappings *string `json:"ModelMappings,omitempty" xml:"ModelMappings,omitempty"`
	// The name of the model API.
	//
	// example:
	//
	// jev-chat-api
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The invocation path prefix.
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
	// 6BD9CDE4-5E7B-4BF3-9BB8-83C73E******
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The saved routing rules, returned as a JSON string. A byte-by-byte exact echo is not guaranteed.
	//
	// example:
	//
	// [{"RuleName":"default","Providers":[{"ModelServiceName":"model-service-a","ModelList":"[\\\\"deepseek-v4-pro\\\\"]","Weight":100}]}]
	RouteRules *string `json:"RouteRules,omitempty" xml:"RouteRules,omitempty"`
	// The status of the model API.
	//
	// example:
	//
	// Enable
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
}

func (s CreateModelApiResponseBody) String() string {
	return dara.Prettify(s)
}

func (s CreateModelApiResponseBody) GoString() string {
	return s.String()
}

func (s *CreateModelApiResponseBody) GetCategory() *string {
	return s.Category
}

func (s *CreateModelApiResponseBody) GetGmtCreated() *string {
	return s.GmtCreated
}

func (s *CreateModelApiResponseBody) GetGmtModified() *string {
	return s.GmtModified
}

func (s *CreateModelApiResponseBody) GetInvokeEndpoint() *string {
	return s.InvokeEndpoint
}

func (s *CreateModelApiResponseBody) GetModelApiId() *string {
	return s.ModelApiId
}

func (s *CreateModelApiResponseBody) GetModelMappings() *string {
	return s.ModelMappings
}

func (s *CreateModelApiResponseBody) GetName() *string {
	return s.Name
}

func (s *CreateModelApiResponseBody) GetPathPrefix() *string {
	return s.PathPrefix
}

func (s *CreateModelApiResponseBody) GetProtocol() *string {
	return s.Protocol
}

func (s *CreateModelApiResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *CreateModelApiResponseBody) GetRouteRules() *string {
	return s.RouteRules
}

func (s *CreateModelApiResponseBody) GetStatus() *string {
	return s.Status
}

func (s *CreateModelApiResponseBody) SetCategory(v string) *CreateModelApiResponseBody {
	s.Category = &v
	return s
}

func (s *CreateModelApiResponseBody) SetGmtCreated(v string) *CreateModelApiResponseBody {
	s.GmtCreated = &v
	return s
}

func (s *CreateModelApiResponseBody) SetGmtModified(v string) *CreateModelApiResponseBody {
	s.GmtModified = &v
	return s
}

func (s *CreateModelApiResponseBody) SetInvokeEndpoint(v string) *CreateModelApiResponseBody {
	s.InvokeEndpoint = &v
	return s
}

func (s *CreateModelApiResponseBody) SetModelApiId(v string) *CreateModelApiResponseBody {
	s.ModelApiId = &v
	return s
}

func (s *CreateModelApiResponseBody) SetModelMappings(v string) *CreateModelApiResponseBody {
	s.ModelMappings = &v
	return s
}

func (s *CreateModelApiResponseBody) SetName(v string) *CreateModelApiResponseBody {
	s.Name = &v
	return s
}

func (s *CreateModelApiResponseBody) SetPathPrefix(v string) *CreateModelApiResponseBody {
	s.PathPrefix = &v
	return s
}

func (s *CreateModelApiResponseBody) SetProtocol(v string) *CreateModelApiResponseBody {
	s.Protocol = &v
	return s
}

func (s *CreateModelApiResponseBody) SetRequestId(v string) *CreateModelApiResponseBody {
	s.RequestId = &v
	return s
}

func (s *CreateModelApiResponseBody) SetRouteRules(v string) *CreateModelApiResponseBody {
	s.RouteRules = &v
	return s
}

func (s *CreateModelApiResponseBody) SetStatus(v string) *CreateModelApiResponseBody {
	s.Status = &v
	return s
}

func (s *CreateModelApiResponseBody) Validate() error {
	return dara.Validate(s)
}
