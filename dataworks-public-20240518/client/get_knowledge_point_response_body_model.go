// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetKnowledgePointResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *GetKnowledgePointResponseBody
	GetCode() *string
	SetData(v *GetKnowledgePointResponseBodyData) *GetKnowledgePointResponseBody
	GetData() *GetKnowledgePointResponseBodyData
	SetMessage(v string) *GetKnowledgePointResponseBody
	GetMessage() *string
	SetRequestId(v string) *GetKnowledgePointResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *GetKnowledgePointResponseBody
	GetSuccess() *bool
}

type GetKnowledgePointResponseBody struct {
	// The business error code. This parameter is omitted when the request is successful. For failures, the common error response protocol is followed.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The typed details of a single knowledge point. Only the item corresponding to the Type value is returned among Graph, Wiki, and KnowledgeContent.
	Data *GetKnowledgePointResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// The readable description of the business error. This parameter is omitted when the request is successful. Provide the RequestId when troubleshooting failures.
	//
	// example:
	//
	// The request parameters are invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request, which can be used to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s GetKnowledgePointResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBody) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBody) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgePointResponseBody) GetData() *GetKnowledgePointResponseBodyData {
	return s.Data
}

func (s *GetKnowledgePointResponseBody) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgePointResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetKnowledgePointResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *GetKnowledgePointResponseBody) SetCode(v string) *GetKnowledgePointResponseBody {
	s.Code = &v
	return s
}

func (s *GetKnowledgePointResponseBody) SetData(v *GetKnowledgePointResponseBodyData) *GetKnowledgePointResponseBody {
	s.Data = v
	return s
}

func (s *GetKnowledgePointResponseBody) SetMessage(v string) *GetKnowledgePointResponseBody {
	s.Message = &v
	return s
}

func (s *GetKnowledgePointResponseBody) SetRequestId(v string) *GetKnowledgePointResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetKnowledgePointResponseBody) SetSuccess(v bool) *GetKnowledgePointResponseBody {
	s.Success = &v
	return s
}

func (s *GetKnowledgePointResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgePointResponseBodyData struct {
	// The one-hop local graph with the requested Id as the entry point. This parameter is returned only when Type is set to GRAPH. It includes the entry point, the relationships on the current page, and their visible endpoints.
	Graph *GetKnowledgePointResponseBodyDataGraph `json:"Graph,omitempty" xml:"Graph,omitempty" type:"Struct"`
	// The location ID of the knowledge point. Copy this value exactly as it appears in the search results and use it together with the corresponding KnowledgeBaseId and Type. Do not parse or convert this value.
	//
	// example:
	//
	// node-sales
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The ID of the knowledge base to which the knowledge point belongs. This parameter is used together with Type and Id to locate the knowledge point.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The title, Markdown body, and optional sources of the knowledge content. This parameter is returned only when Type is set to KNOWLEDGE_CONTENT.
	KnowledgeContent *GetKnowledgePointResponseBodyDataKnowledgeContent `json:"KnowledgeContent,omitempty" xml:"KnowledgeContent,omitempty" type:"Struct"`
	// The pagination cursor. Do not specify this parameter when you request the first page. For subsequent pages, use the NextToken value returned in the previous response. An empty response indicates that no more pages are available. If the current page is empty but the cursor is not empty, continue pagination. This parameter is used only when Type is set to GRAPH. You must keep the same knowledge base, entry point, and MaxResults values.
	//
	// example:
	//
	// eyJvZmZzZXQiOjIwfQ
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The read category of the knowledge point. Valid values: GRAPH for local graphs, WIKI for topic pages, and KNOWLEDGE_CONTENT for knowledge content.
	//
	// example:
	//
	// GRAPH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
	// The truncation, scope limitation, or degradation prompts for the current call. Use Code for programmatic judgment. Message is for reading purposes only.
	Warnings []*GetKnowledgePointResponseBodyDataWarnings `json:"Warnings,omitempty" xml:"Warnings,omitempty" type:"Repeated"`
	// The title, Markdown body, and optional sources of the topic page. This parameter is returned only when Type is set to WIKI.
	Wiki *GetKnowledgePointResponseBodyDataWiki `json:"Wiki,omitempty" xml:"Wiki,omitempty" type:"Struct"`
}

func (s GetKnowledgePointResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyData) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyData) GetGraph() *GetKnowledgePointResponseBodyDataGraph {
	return s.Graph
}

func (s *GetKnowledgePointResponseBodyData) GetId() *string {
	return s.Id
}

func (s *GetKnowledgePointResponseBodyData) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgePointResponseBodyData) GetKnowledgeContent() *GetKnowledgePointResponseBodyDataKnowledgeContent {
	return s.KnowledgeContent
}

func (s *GetKnowledgePointResponseBodyData) GetNextToken() *string {
	return s.NextToken
}

func (s *GetKnowledgePointResponseBodyData) GetType() *string {
	return s.Type
}

func (s *GetKnowledgePointResponseBodyData) GetWarnings() []*GetKnowledgePointResponseBodyDataWarnings {
	return s.Warnings
}

func (s *GetKnowledgePointResponseBodyData) GetWiki() *GetKnowledgePointResponseBodyDataWiki {
	return s.Wiki
}

func (s *GetKnowledgePointResponseBodyData) SetGraph(v *GetKnowledgePointResponseBodyDataGraph) *GetKnowledgePointResponseBodyData {
	s.Graph = v
	return s
}

func (s *GetKnowledgePointResponseBodyData) SetId(v string) *GetKnowledgePointResponseBodyData {
	s.Id = &v
	return s
}

func (s *GetKnowledgePointResponseBodyData) SetKnowledgeBaseId(v string) *GetKnowledgePointResponseBodyData {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyData) SetKnowledgeContent(v *GetKnowledgePointResponseBodyDataKnowledgeContent) *GetKnowledgePointResponseBodyData {
	s.KnowledgeContent = v
	return s
}

func (s *GetKnowledgePointResponseBodyData) SetNextToken(v string) *GetKnowledgePointResponseBodyData {
	s.NextToken = &v
	return s
}

func (s *GetKnowledgePointResponseBodyData) SetType(v string) *GetKnowledgePointResponseBodyData {
	s.Type = &v
	return s
}

func (s *GetKnowledgePointResponseBodyData) SetWarnings(v []*GetKnowledgePointResponseBodyDataWarnings) *GetKnowledgePointResponseBodyData {
	s.Warnings = v
	return s
}

func (s *GetKnowledgePointResponseBodyData) SetWiki(v *GetKnowledgePointResponseBodyDataWiki) *GetKnowledgePointResponseBodyData {
	s.Wiki = v
	return s
}

func (s *GetKnowledgePointResponseBodyData) Validate() error {
	if s.Graph != nil {
		if err := s.Graph.Validate(); err != nil {
			return err
		}
	}
	if s.KnowledgeContent != nil {
		if err := s.KnowledgeContent.Validate(); err != nil {
			return err
		}
	}
	if s.Warnings != nil {
		for _, item := range s.Warnings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Wiki != nil {
		if err := s.Wiki.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgePointResponseBodyDataGraph struct {
	// The visible one-hop adjacency relationships on the current page. At least one end of each edge is the RootNodeId, and both ends are in the Nodes of the current page.
	Edges []*GetKnowledgePointResponseBodyDataGraphEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The ID of the knowledge base to which the knowledge point belongs. This parameter is used together with Type and Id to locate the knowledge point.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// All visible endpoints involved by the entry node and the relationships on the current page, deduplicated by NodeId. Subsequent pages still include the entry node.
	Nodes []*GetKnowledgePointResponseBodyDataGraphNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
	// The ID of the entry node for the current local graph, which is the same as the requested Id.
	//
	// example:
	//
	// node-sales
	RootNodeId *string `json:"RootNodeId,omitempty" xml:"RootNodeId,omitempty"`
	// The time when the current graph was last published. This parameter is returned when the time is reliable.
	//
	// example:
	//
	// 1794373200000
	UpdateTime *int64 `json:"UpdateTime,omitempty" xml:"UpdateTime,omitempty"`
}

func (s GetKnowledgePointResponseBodyDataGraph) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataGraph) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataGraph) GetEdges() []*GetKnowledgePointResponseBodyDataGraphEdges {
	return s.Edges
}

func (s *GetKnowledgePointResponseBodyDataGraph) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgePointResponseBodyDataGraph) GetNodes() []*GetKnowledgePointResponseBodyDataGraphNodes {
	return s.Nodes
}

func (s *GetKnowledgePointResponseBodyDataGraph) GetRootNodeId() *string {
	return s.RootNodeId
}

func (s *GetKnowledgePointResponseBodyDataGraph) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *GetKnowledgePointResponseBodyDataGraph) SetEdges(v []*GetKnowledgePointResponseBodyDataGraphEdges) *GetKnowledgePointResponseBodyDataGraph {
	s.Edges = v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraph) SetKnowledgeBaseId(v string) *GetKnowledgePointResponseBodyDataGraph {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraph) SetNodes(v []*GetKnowledgePointResponseBodyDataGraphNodes) *GetKnowledgePointResponseBodyDataGraph {
	s.Nodes = v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraph) SetRootNodeId(v string) *GetKnowledgePointResponseBodyDataGraph {
	s.RootNodeId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraph) SetUpdateTime(v int64) *GetKnowledgePointResponseBodyDataGraph {
	s.UpdateTime = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraph) Validate() error {
	if s.Edges != nil {
		for _, item := range s.Edges {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Nodes != nil {
		for _, item := range s.Nodes {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKnowledgePointResponseBodyDataGraphEdges struct {
	// The readable supplementary description of the current node or relationship. This parameter is omitted when no reliable content is available.
	//
	// example:
	//
	// Sales revenue is calculated based on the actual paid amount of orders
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The ID of the relationship in the current knowledge base graph.
	//
	// example:
	//
	// edge-sales-orders
	EdgeId *string `json:"EdgeId,omitempty" xml:"EdgeId,omitempty"`
	// The NodeId of the start node of the relationship. The start node must exist in the Graph.Nodes of the same page.
	//
	// example:
	//
	// node-sales
	FromNodeId *string `json:"FromNodeId,omitempty" xml:"FromNodeId,omitempty"`
	// The typed field association information for verified relationships. This parameter is omitted when no reliable and complete field pairing is available.
	Join *GetKnowledgePointResponseBodyDataGraphEdgesJoin `json:"Join,omitempty" xml:"Join,omitempty" type:"Struct"`
	// The ID of the knowledge base to which the knowledge point belongs. This parameter is used together with Type and Id to locate the knowledge point.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The public relationship type from the start node to the end node. For example, CALCULATED_FROM indicates that the start node is calculated based on the end node. You cannot infer the complete formula based solely on the type.
	//
	// example:
	//
	// CALCULATED_FROM
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The NodeId of the end node of the relationship. The end node must exist in the Graph.Nodes of the same page.
	//
	// example:
	//
	// node-orders
	ToNodeId *string `json:"ToNodeId,omitempty" xml:"ToNodeId,omitempty"`
}

func (s GetKnowledgePointResponseBodyDataGraphEdges) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataGraphEdges) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) GetEdgeId() *string {
	return s.EdgeId
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) GetFromNodeId() *string {
	return s.FromNodeId
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) GetJoin() *GetKnowledgePointResponseBodyDataGraphEdgesJoin {
	return s.Join
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) GetRelationType() *string {
	return s.RelationType
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) GetToNodeId() *string {
	return s.ToNodeId
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) SetDescription(v string) *GetKnowledgePointResponseBodyDataGraphEdges {
	s.Description = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) SetEdgeId(v string) *GetKnowledgePointResponseBodyDataGraphEdges {
	s.EdgeId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) SetFromNodeId(v string) *GetKnowledgePointResponseBodyDataGraphEdges {
	s.FromNodeId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) SetJoin(v *GetKnowledgePointResponseBodyDataGraphEdgesJoin) *GetKnowledgePointResponseBodyDataGraphEdges {
	s.Join = v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) SetKnowledgeBaseId(v string) *GetKnowledgePointResponseBodyDataGraphEdges {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) SetRelationType(v string) *GetKnowledgePointResponseBodyDataGraphEdges {
	s.RelationType = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) SetToNodeId(v string) *GetKnowledgePointResponseBodyDataGraphEdges {
	s.ToNodeId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdges) Validate() error {
	if s.Join != nil {
		if err := s.Join.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type GetKnowledgePointResponseBodyDataGraphEdgesJoin struct {
	// The association cardinality from the start node to the end node. Valid values: ONE_TO_ONE, ONE_TO_MANY, MANY_TO_ONE, and MANY_TO_MANY. This parameter is returned when the cardinality is reliable.
	//
	// example:
	//
	// MANY_TO_ONE
	Cardinality *string `json:"Cardinality,omitempty" xml:"Cardinality,omitempty"`
	// The equivalent association field pairing for the nodes at both ends. Composite keys are returned in full.
	JoinKeys []*GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys `json:"JoinKeys,omitempty" xml:"JoinKeys,omitempty" type:"Repeated"`
}

func (s GetKnowledgePointResponseBodyDataGraphEdgesJoin) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataGraphEdgesJoin) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoin) GetCardinality() *string {
	return s.Cardinality
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoin) GetJoinKeys() []*GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys {
	return s.JoinKeys
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoin) SetCardinality(v string) *GetKnowledgePointResponseBodyDataGraphEdgesJoin {
	s.Cardinality = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoin) SetJoinKeys(v []*GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys) *GetKnowledgePointResponseBodyDataGraphEdgesJoin {
	s.JoinKeys = v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoin) Validate() error {
	if s.JoinKeys != nil {
		for _, item := range s.JoinKeys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys struct {
	// The name of the association field for the node corresponding to FromNodeId.
	//
	// example:
	//
	// customer_id
	FromField *string `json:"FromField,omitempty" xml:"FromField,omitempty"`
	// The name of the association field for the node corresponding to ToNodeId.
	//
	// example:
	//
	// id
	ToField *string `json:"ToField,omitempty" xml:"ToField,omitempty"`
}

func (s GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys) GetFromField() *string {
	return s.FromField
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys) GetToField() *string {
	return s.ToField
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys) SetFromField(v string) *GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys {
	s.FromField = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys) SetToField(v string) *GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys {
	s.ToField = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphEdgesJoinJoinKeys) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgePointResponseBodyDataGraphNodes struct {
	// The list of reliable aliases for the node. This parameter is omitted when no aliases are available.
	Aliases []*string `json:"Aliases,omitempty" xml:"Aliases,omitempty" type:"Repeated"`
	// The reliable business definition of the graph node. This parameter is omitted when no public definition is available.
	//
	// example:
	//
	// The sum of the actual paid amounts of paid orders
	Definition *string `json:"Definition,omitempty" xml:"Definition,omitempty"`
	// The readable supplementary description of the current node or relationship. This parameter is omitted when no reliable content is available.
	//
	// example:
	//
	// Sales revenue is calculated based on the actual paid amount of orders
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The asset ID in DataWorks Data Map that can be reliably located. This is not the NodeId of the graph node. This parameter is omitted when no unique and reliable asset mapping is available.
	//
	// example:
	//
	// maxcompute-table:::sales:public:orders
	EntityId *string `json:"EntityId,omitempty" xml:"EntityId,omitempty"`
	// The ID of the knowledge base to which the knowledge point belongs. This parameter is used together with Type and Id to locate the knowledge point.
	//
	// example:
	//
	// 12345678-1234-4234-8234-123456789abc
	KnowledgeBaseId *string `json:"KnowledgeBaseId,omitempty" xml:"KnowledgeBaseId,omitempty"`
	// The display name of the graph node.
	//
	// example:
	//
	// Sales revenue
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The ID of the node in the current knowledge base graph. You can use this value as the Id parameter when Type is set to GRAPH to read the local graph of this node.
	//
	// example:
	//
	// node-sales
	NodeId *string `json:"NodeId,omitempty" xml:"NodeId,omitempty"`
	// The business category of the graph node, such as TABLE for a data table. This is not the read category of the knowledge point.
	//
	// example:
	//
	// TABLE
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s GetKnowledgePointResponseBodyDataGraphNodes) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataGraphNodes) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) GetAliases() []*string {
	return s.Aliases
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) GetDefinition() *string {
	return s.Definition
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) GetDescription() *string {
	return s.Description
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) GetEntityId() *string {
	return s.EntityId
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) GetName() *string {
	return s.Name
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) GetNodeId() *string {
	return s.NodeId
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) GetType() *string {
	return s.Type
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) SetAliases(v []*string) *GetKnowledgePointResponseBodyDataGraphNodes {
	s.Aliases = v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) SetDefinition(v string) *GetKnowledgePointResponseBodyDataGraphNodes {
	s.Definition = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) SetDescription(v string) *GetKnowledgePointResponseBodyDataGraphNodes {
	s.Description = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) SetEntityId(v string) *GetKnowledgePointResponseBodyDataGraphNodes {
	s.EntityId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) SetKnowledgeBaseId(v string) *GetKnowledgePointResponseBodyDataGraphNodes {
	s.KnowledgeBaseId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) SetName(v string) *GetKnowledgePointResponseBodyDataGraphNodes {
	s.Name = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) SetNodeId(v string) *GetKnowledgePointResponseBodyDataGraphNodes {
	s.NodeId = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) SetType(v string) *GetKnowledgePointResponseBodyDataGraphNodes {
	s.Type = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataGraphNodes) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgePointResponseBodyDataKnowledgeContent struct {
	// The currently public Markdown body. The knowledge content may be a document fragment or a readable projection of structured knowledge, and is not guaranteed to be the full text of the original file.
	//
	// example:
	//
	// Sales revenue is aggregated based on the actual paid amount of paid orders
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The download information of the uploaded original document explicitly associated with the current content. This parameter is omitted when no reliable source is available or when you do not have read permissions for the original file. The ability to retrieve a knowledge point does not imply the ability to download the complete original file.
	Sources []*GetKnowledgePointResponseBodyDataKnowledgeContentSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The display title of the current knowledge point.
	//
	// example:
	//
	// Sales revenue
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetKnowledgePointResponseBodyDataKnowledgeContent) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataKnowledgeContent) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContent) GetContent() *string {
	return s.Content
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContent) GetSources() []*GetKnowledgePointResponseBodyDataKnowledgeContentSources {
	return s.Sources
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContent) GetTitle() *string {
	return s.Title
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContent) SetContent(v string) *GetKnowledgePointResponseBodyDataKnowledgeContent {
	s.Content = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContent) SetSources(v []*GetKnowledgePointResponseBodyDataKnowledgeContentSources) *GetKnowledgePointResponseBodyDataKnowledgeContent {
	s.Sources = v
	return s
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContent) SetTitle(v string) *GetKnowledgePointResponseBodyDataKnowledgeContent {
	s.Title = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContent) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKnowledgePointResponseBodyDataKnowledgeContentSources struct {
	// The temporary HTTPS download URL of the original file, which is valid for 2 hours. This parameter is returned only when the source can be reliably located and the caller has read permissions for the original file. Do not forward this URL publicly or store it for a long time.
	//
	// example:
	//
	// https://example.com/sales-policy.pdf
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The expiration time of the download URL, in UNIX milliseconds.
	//
	// example:
	//
	// 1794376800000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The full file name of the original document used for building. This is not the name of the current knowledge point fragment.
	//
	// example:
	//
	// Sales_policy.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetKnowledgePointResponseBodyDataKnowledgeContentSources) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataKnowledgeContentSources) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContentSources) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContentSources) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContentSources) GetFileName() *string {
	return s.FileName
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContentSources) SetDownloadUrl(v string) *GetKnowledgePointResponseBodyDataKnowledgeContentSources {
	s.DownloadUrl = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContentSources) SetExpirationTime(v int64) *GetKnowledgePointResponseBodyDataKnowledgeContentSources {
	s.ExpirationTime = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContentSources) SetFileName(v string) *GetKnowledgePointResponseBodyDataKnowledgeContentSources {
	s.FileName = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataKnowledgeContentSources) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgePointResponseBodyDataWarnings struct {
	// The truncation or degradation prompt code. For example, CONTENT_TRUNCATED indicates that the content is truncated, and SOURCE_UNAVAILABLE indicates that the original file download is temporarily unavailable.
	//
	// example:
	//
	// CONTENT_TRUNCATED
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The sanitized and readable description of the current prompt. Use Code for programmatic judgment instead of relying on the text.
	//
	// example:
	//
	// Some content exceeds the length limit and has been truncated
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s GetKnowledgePointResponseBodyDataWarnings) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataWarnings) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataWarnings) GetCode() *string {
	return s.Code
}

func (s *GetKnowledgePointResponseBodyDataWarnings) GetMessage() *string {
	return s.Message
}

func (s *GetKnowledgePointResponseBodyDataWarnings) SetCode(v string) *GetKnowledgePointResponseBodyDataWarnings {
	s.Code = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataWarnings) SetMessage(v string) *GetKnowledgePointResponseBodyDataWarnings {
	s.Message = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataWarnings) Validate() error {
	return dara.Validate(s)
}

type GetKnowledgePointResponseBodyDataWiki struct {
	// The currently public Markdown body. The knowledge content may be a document fragment or a readable projection of structured knowledge, and is not guaranteed to be the full text of the original file.
	//
	// example:
	//
	// Sales revenue is aggregated based on the actual paid amount of paid orders
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The download information of the uploaded original document explicitly associated with the current content. This parameter is omitted when no reliable source is available or when you do not have read permissions for the original file. The ability to retrieve a knowledge point does not imply the ability to download the complete original file.
	Sources []*GetKnowledgePointResponseBodyDataWikiSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The display title of the current knowledge point.
	//
	// example:
	//
	// Sales revenue
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
}

func (s GetKnowledgePointResponseBodyDataWiki) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataWiki) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataWiki) GetContent() *string {
	return s.Content
}

func (s *GetKnowledgePointResponseBodyDataWiki) GetSources() []*GetKnowledgePointResponseBodyDataWikiSources {
	return s.Sources
}

func (s *GetKnowledgePointResponseBodyDataWiki) GetTitle() *string {
	return s.Title
}

func (s *GetKnowledgePointResponseBodyDataWiki) SetContent(v string) *GetKnowledgePointResponseBodyDataWiki {
	s.Content = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataWiki) SetSources(v []*GetKnowledgePointResponseBodyDataWikiSources) *GetKnowledgePointResponseBodyDataWiki {
	s.Sources = v
	return s
}

func (s *GetKnowledgePointResponseBodyDataWiki) SetTitle(v string) *GetKnowledgePointResponseBodyDataWiki {
	s.Title = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataWiki) Validate() error {
	if s.Sources != nil {
		for _, item := range s.Sources {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type GetKnowledgePointResponseBodyDataWikiSources struct {
	// The temporary HTTPS download URL of the original file, which is valid for 2 hours. This parameter is returned only when the source can be reliably located and the caller has read permissions for the original file. Do not forward this URL publicly or store it for a long time.
	//
	// example:
	//
	// https://example.com/sales-policy.pdf
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The expiration time of the download URL, in UNIX milliseconds.
	//
	// example:
	//
	// 1794376800000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The full file name of the original document used for building. This is not the name of the current knowledge point fragment.
	//
	// example:
	//
	// Sales_policy.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s GetKnowledgePointResponseBodyDataWikiSources) String() string {
	return dara.Prettify(s)
}

func (s GetKnowledgePointResponseBodyDataWikiSources) GoString() string {
	return s.String()
}

func (s *GetKnowledgePointResponseBodyDataWikiSources) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *GetKnowledgePointResponseBodyDataWikiSources) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *GetKnowledgePointResponseBodyDataWikiSources) GetFileName() *string {
	return s.FileName
}

func (s *GetKnowledgePointResponseBodyDataWikiSources) SetDownloadUrl(v string) *GetKnowledgePointResponseBodyDataWikiSources {
	s.DownloadUrl = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataWikiSources) SetExpirationTime(v int64) *GetKnowledgePointResponseBodyDataWikiSources {
	s.ExpirationTime = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataWikiSources) SetFileName(v string) *GetKnowledgePointResponseBodyDataWikiSources {
	s.FileName = &v
	return s
}

func (s *GetKnowledgePointResponseBodyDataWikiSources) Validate() error {
	return dara.Validate(s)
}
