// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iRetrieveKnowledgePointsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCode(v string) *RetrieveKnowledgePointsResponseBody
	GetCode() *string
	SetData(v *RetrieveKnowledgePointsResponseBodyData) *RetrieveKnowledgePointsResponseBody
	GetData() *RetrieveKnowledgePointsResponseBodyData
	SetMessage(v string) *RetrieveKnowledgePointsResponseBody
	GetMessage() *string
	SetRequestId(v string) *RetrieveKnowledgePointsResponseBody
	GetRequestId() *string
	SetSuccess(v bool) *RetrieveKnowledgePointsResponseBody
	GetSuccess() *bool
}

type RetrieveKnowledgePointsResponseBody struct {
	// The business error code. This field may be omitted if the request succeeds. If the request fails, the common error response protocol applies.
	//
	// example:
	//
	// InvalidParameter
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The knowledge point summaries, relationship paths, ambiguity candidates, and warnings for the current search.
	Data *RetrieveKnowledgePointsResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Struct"`
	// A human-readable description of the business error. This field may be omitted if the request succeeds. Provide the RequestId when troubleshooting failures.
	//
	// example:
	//
	// The request parameters are invalid.
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The unique identifier of the request. You can use this ID to troubleshoot issues.
	//
	// example:
	//
	// 00000000-0000-4000-8000-000000000001
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// Indicates whether the request is successful.
	Success *bool `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s RetrieveKnowledgePointsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBody) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBody) GetCode() *string {
	return s.Code
}

func (s *RetrieveKnowledgePointsResponseBody) GetData() *RetrieveKnowledgePointsResponseBodyData {
	return s.Data
}

func (s *RetrieveKnowledgePointsResponseBody) GetMessage() *string {
	return s.Message
}

func (s *RetrieveKnowledgePointsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *RetrieveKnowledgePointsResponseBody) GetSuccess() *bool {
	return s.Success
}

func (s *RetrieveKnowledgePointsResponseBody) SetCode(v string) *RetrieveKnowledgePointsResponseBody {
	s.Code = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBody) SetData(v *RetrieveKnowledgePointsResponseBodyData) *RetrieveKnowledgePointsResponseBody {
	s.Data = v
	return s
}

func (s *RetrieveKnowledgePointsResponseBody) SetMessage(v string) *RetrieveKnowledgePointsResponseBody {
	s.Message = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBody) SetRequestId(v string) *RetrieveKnowledgePointsResponseBody {
	s.RequestId = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBody) SetSuccess(v bool) *RetrieveKnowledgePointsResponseBody {
	s.Success = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBody) Validate() error {
	if s.Data != nil {
		if err := s.Data.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type RetrieveKnowledgePointsResponseBodyData struct {
	// The ambiguity descriptions and candidate objects returned when the target is not yet confirmed.
	Ambiguities []*RetrieveKnowledgePointsResponseBodyDataAmbiguities `json:"Ambiguities,omitempty" xml:"Ambiguities,omitempty" type:"Repeated"`
	// The relationship paths matched by the search. Path indexes do not correspond to indexes in the Results array. The paths may cover only a subset of relationships.
	Paths []*RetrieveKnowledgePointsResponseBodyDataPaths `json:"Paths,omitempty" xml:"Paths,omitempty" type:"Repeated"`
	// Knowledge point summaries ranked by relevance across all knowledge bases. The total number of results does not exceed MaxResults. Results can include GRAPH, WIKI, and KNOWLEDGE_CONTENT.
	Results []*RetrieveKnowledgePointsResponseBodyDataResults `json:"Results,omitempty" xml:"Results,omitempty" type:"Repeated"`
	// The truncation, scope limitation, or degradation warnings for the current call. Use Code for programmatic judgment. Message is for reading purposes only.
	Warnings []*RetrieveKnowledgePointsResponseBodyDataWarnings `json:"Warnings,omitempty" xml:"Warnings,omitempty" type:"Repeated"`
}

func (s RetrieveKnowledgePointsResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBodyData) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBodyData) GetAmbiguities() []*RetrieveKnowledgePointsResponseBodyDataAmbiguities {
	return s.Ambiguities
}

func (s *RetrieveKnowledgePointsResponseBodyData) GetPaths() []*RetrieveKnowledgePointsResponseBodyDataPaths {
	return s.Paths
}

func (s *RetrieveKnowledgePointsResponseBodyData) GetResults() []*RetrieveKnowledgePointsResponseBodyDataResults {
	return s.Results
}

func (s *RetrieveKnowledgePointsResponseBodyData) GetWarnings() []*RetrieveKnowledgePointsResponseBodyDataWarnings {
	return s.Warnings
}

func (s *RetrieveKnowledgePointsResponseBodyData) SetAmbiguities(v []*RetrieveKnowledgePointsResponseBodyDataAmbiguities) *RetrieveKnowledgePointsResponseBodyData {
	s.Ambiguities = v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyData) SetPaths(v []*RetrieveKnowledgePointsResponseBodyDataPaths) *RetrieveKnowledgePointsResponseBodyData {
	s.Paths = v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyData) SetResults(v []*RetrieveKnowledgePointsResponseBodyDataResults) *RetrieveKnowledgePointsResponseBodyData {
	s.Results = v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyData) SetWarnings(v []*RetrieveKnowledgePointsResponseBodyDataWarnings) *RetrieveKnowledgePointsResponseBodyData {
	s.Warnings = v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyData) Validate() error {
	if s.Ambiguities != nil {
		for _, item := range s.Ambiguities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Paths != nil {
		for _, item := range s.Paths {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	if s.Results != nil {
		for _, item := range s.Results {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
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
	return nil
}

type RetrieveKnowledgePointsResponseBodyDataAmbiguities struct {
	// The candidate graph nodes when ambiguity exists. Each candidate carries the KnowledgeBaseId, Type, and Id for reading details, which does not indicate that the user has confirmed the selection.
	Candidates []*RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates `json:"Candidates,omitempty" xml:"Candidates,omitempty" type:"Repeated"`
	// The name or description of the object whose referent is not yet determined in the query.
	//
	// example:
	//
	// Sales revenue
	Mention *string `json:"Mention,omitempty" xml:"Mention,omitempty"`
}

func (s RetrieveKnowledgePointsResponseBodyDataAmbiguities) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBodyDataAmbiguities) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguities) GetCandidates() []*RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates {
	return s.Candidates
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguities) GetMention() *string {
	return s.Mention
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguities) SetCandidates(v []*RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) *RetrieveKnowledgePointsResponseBodyDataAmbiguities {
	s.Candidates = v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguities) SetMention(v string) *RetrieveKnowledgePointsResponseBodyDataAmbiguities {
	s.Mention = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguities) Validate() error {
	if s.Candidates != nil {
		for _, item := range s.Candidates {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates struct {
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
	// The display title of the current knowledge point.
	//
	// example:
	//
	// Sales revenue
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The read category of the knowledge point. Use GRAPH for the current focus objects and graph nodes.
	//
	// example:
	//
	// GRAPH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) GetId() *string {
	return s.Id
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) GetTitle() *string {
	return s.Title
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) GetType() *string {
	return s.Type
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) SetId(v string) *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates {
	s.Id = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) SetKnowledgeBaseId(v string) *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) SetTitle(v string) *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates {
	s.Title = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) SetType(v string) *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates {
	s.Type = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataAmbiguitiesCandidates) Validate() error {
	return dara.Validate(s)
}

type RetrieveKnowledgePointsResponseBodyDataPaths struct {
	// The directed relationships that connect the nodes in the Nodes array of the current path.
	Edges []*RetrieveKnowledgePointsResponseBodyDataPathsEdges `json:"Edges,omitempty" xml:"Edges,omitempty" type:"Repeated"`
	// The graph nodes in the relationship path. You can read the details of each node by using its KnowledgeBaseId, Type, and Id.
	Nodes []*RetrieveKnowledgePointsResponseBodyDataPathsNodes `json:"Nodes,omitempty" xml:"Nodes,omitempty" type:"Repeated"`
}

func (s RetrieveKnowledgePointsResponseBodyDataPaths) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBodyDataPaths) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBodyDataPaths) GetEdges() []*RetrieveKnowledgePointsResponseBodyDataPathsEdges {
	return s.Edges
}

func (s *RetrieveKnowledgePointsResponseBodyDataPaths) GetNodes() []*RetrieveKnowledgePointsResponseBodyDataPathsNodes {
	return s.Nodes
}

func (s *RetrieveKnowledgePointsResponseBodyDataPaths) SetEdges(v []*RetrieveKnowledgePointsResponseBodyDataPathsEdges) *RetrieveKnowledgePointsResponseBodyDataPaths {
	s.Edges = v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataPaths) SetNodes(v []*RetrieveKnowledgePointsResponseBodyDataPathsNodes) *RetrieveKnowledgePointsResponseBodyDataPaths {
	s.Nodes = v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataPaths) Validate() error {
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

type RetrieveKnowledgePointsResponseBodyDataPathsEdges struct {
	// The Id of the start node in the Nodes array of the same path.
	//
	// example:
	//
	// node-sales
	FromId *string `json:"FromId,omitempty" xml:"FromId,omitempty"`
	// The public relationship type directed from the start node to the end node. For example, CALCULATED_FROM indicates that the start node is calculated based on the end node. You cannot infer the complete formula based solely on the relationship type.
	//
	// example:
	//
	// CALCULATED_FROM
	RelationType *string `json:"RelationType,omitempty" xml:"RelationType,omitempty"`
	// The Id of the end node in the Nodes array of the same path.
	//
	// example:
	//
	// node-orders
	ToId *string `json:"ToId,omitempty" xml:"ToId,omitempty"`
}

func (s RetrieveKnowledgePointsResponseBodyDataPathsEdges) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBodyDataPathsEdges) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsEdges) GetFromId() *string {
	return s.FromId
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsEdges) GetRelationType() *string {
	return s.RelationType
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsEdges) GetToId() *string {
	return s.ToId
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsEdges) SetFromId(v string) *RetrieveKnowledgePointsResponseBodyDataPathsEdges {
	s.FromId = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsEdges) SetRelationType(v string) *RetrieveKnowledgePointsResponseBodyDataPathsEdges {
	s.RelationType = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsEdges) SetToId(v string) *RetrieveKnowledgePointsResponseBodyDataPathsEdges {
	s.ToId = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsEdges) Validate() error {
	return dara.Validate(s)
}

type RetrieveKnowledgePointsResponseBodyDataPathsNodes struct {
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
	// The display title of the current knowledge point.
	//
	// example:
	//
	// Sales revenue
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The read category of the knowledge point. Use GRAPH for the current focus objects and graph nodes.
	//
	// example:
	//
	// GRAPH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RetrieveKnowledgePointsResponseBodyDataPathsNodes) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBodyDataPathsNodes) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsNodes) GetId() *string {
	return s.Id
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsNodes) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsNodes) GetTitle() *string {
	return s.Title
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsNodes) GetType() *string {
	return s.Type
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsNodes) SetId(v string) *RetrieveKnowledgePointsResponseBodyDataPathsNodes {
	s.Id = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsNodes) SetKnowledgeBaseId(v string) *RetrieveKnowledgePointsResponseBodyDataPathsNodes {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsNodes) SetTitle(v string) *RetrieveKnowledgePointsResponseBodyDataPathsNodes {
	s.Title = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsNodes) SetType(v string) *RetrieveKnowledgePointsResponseBodyDataPathsNodes {
	s.Type = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataPathsNodes) Validate() error {
	return dara.Validate(s)
}

type RetrieveKnowledgePointsResponseBodyDataResults struct {
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
	// The download information of the uploaded original documents explicitly associated with the current content. This parameter is omitted if no reliable source is available or if the caller lacks read permissions for the original files. The ability to retrieve knowledge points does not imply the permission to download the complete original files.
	Sources []*RetrieveKnowledgePointsResponseBodyDataResultsSources `json:"Sources,omitempty" xml:"Sources,omitempty" type:"Repeated"`
	// The readable summary of the knowledge point. The summary is not guaranteed to contain the complete content. To obtain details, call GetKnowledgePoint.
	//
	// example:
	//
	// The sales revenue metric and its associated order table
	Summary *string `json:"Summary,omitempty" xml:"Summary,omitempty"`
	// The display title of the current knowledge point.
	//
	// example:
	//
	// Sales revenue
	Title *string `json:"Title,omitempty" xml:"Title,omitempty"`
	// The read category of the knowledge point. GRAPH indicates a local graph, WIKI indicates a topic page, and KNOWLEDGE_CONTENT indicates knowledge content.
	//
	// example:
	//
	// GRAPH
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s RetrieveKnowledgePointsResponseBodyDataResults) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBodyDataResults) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) GetId() *string {
	return s.Id
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) GetKnowledgeBaseId() *string {
	return s.KnowledgeBaseId
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) GetSources() []*RetrieveKnowledgePointsResponseBodyDataResultsSources {
	return s.Sources
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) GetSummary() *string {
	return s.Summary
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) GetTitle() *string {
	return s.Title
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) GetType() *string {
	return s.Type
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) SetId(v string) *RetrieveKnowledgePointsResponseBodyDataResults {
	s.Id = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) SetKnowledgeBaseId(v string) *RetrieveKnowledgePointsResponseBodyDataResults {
	s.KnowledgeBaseId = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) SetSources(v []*RetrieveKnowledgePointsResponseBodyDataResultsSources) *RetrieveKnowledgePointsResponseBodyDataResults {
	s.Sources = v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) SetSummary(v string) *RetrieveKnowledgePointsResponseBodyDataResults {
	s.Summary = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) SetTitle(v string) *RetrieveKnowledgePointsResponseBodyDataResults {
	s.Title = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) SetType(v string) *RetrieveKnowledgePointsResponseBodyDataResults {
	s.Type = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataResults) Validate() error {
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

type RetrieveKnowledgePointsResponseBodyDataResultsSources struct {
	// The temporary HTTPS download URL of the original file. The URL is valid for 2 hours. This parameter is returned only when the source can be reliably located and the caller has read permissions for the original file. Do not forward the URL publicly or store it for a long time.
	//
	// example:
	//
	// https://example.com/sales-policy.pdf
	DownloadUrl *string `json:"DownloadUrl,omitempty" xml:"DownloadUrl,omitempty"`
	// The expiration time of the download URL. The value is a UNIX timestamp in milliseconds.
	//
	// example:
	//
	// 1794376800000
	ExpirationTime *int64 `json:"ExpirationTime,omitempty" xml:"ExpirationTime,omitempty"`
	// The full file name of the original document used for building. This is not the name of the current knowledge point segment.
	//
	// example:
	//
	// Sales_policy.pdf
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
}

func (s RetrieveKnowledgePointsResponseBodyDataResultsSources) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBodyDataResultsSources) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBodyDataResultsSources) GetDownloadUrl() *string {
	return s.DownloadUrl
}

func (s *RetrieveKnowledgePointsResponseBodyDataResultsSources) GetExpirationTime() *int64 {
	return s.ExpirationTime
}

func (s *RetrieveKnowledgePointsResponseBodyDataResultsSources) GetFileName() *string {
	return s.FileName
}

func (s *RetrieveKnowledgePointsResponseBodyDataResultsSources) SetDownloadUrl(v string) *RetrieveKnowledgePointsResponseBodyDataResultsSources {
	s.DownloadUrl = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataResultsSources) SetExpirationTime(v int64) *RetrieveKnowledgePointsResponseBodyDataResultsSources {
	s.ExpirationTime = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataResultsSources) SetFileName(v string) *RetrieveKnowledgePointsResponseBodyDataResultsSources {
	s.FileName = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataResultsSources) Validate() error {
	return dara.Validate(s)
}

type RetrieveKnowledgePointsResponseBodyDataWarnings struct {
	// The truncation or degradation warning code. For example, CONTENT_TRUNCATED indicates that the content is truncated, and SOURCE_UNAVAILABLE indicates that the original file download is temporarily unavailable.
	//
	// example:
	//
	// CONTENT_TRUNCATED
	Code *string `json:"Code,omitempty" xml:"Code,omitempty"`
	// The sanitized readable description of the current warning. Use Code for programmatic judgment instead of relying on the text.
	//
	// example:
	//
	// Some content exceeds the length limit and has been truncated
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
}

func (s RetrieveKnowledgePointsResponseBodyDataWarnings) String() string {
	return dara.Prettify(s)
}

func (s RetrieveKnowledgePointsResponseBodyDataWarnings) GoString() string {
	return s.String()
}

func (s *RetrieveKnowledgePointsResponseBodyDataWarnings) GetCode() *string {
	return s.Code
}

func (s *RetrieveKnowledgePointsResponseBodyDataWarnings) GetMessage() *string {
	return s.Message
}

func (s *RetrieveKnowledgePointsResponseBodyDataWarnings) SetCode(v string) *RetrieveKnowledgePointsResponseBodyDataWarnings {
	s.Code = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataWarnings) SetMessage(v string) *RetrieveKnowledgePointsResponseBodyDataWarnings {
	s.Message = &v
	return s
}

func (s *RetrieveKnowledgePointsResponseBodyDataWarnings) Validate() error {
	return dara.Validate(s)
}
