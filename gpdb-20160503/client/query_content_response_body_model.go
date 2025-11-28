// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEmbeddingTokens(v string) *QueryContentResponseBody
	GetEmbeddingTokens() *string
	SetEntities(v *QueryContentResponseBodyEntities) *QueryContentResponseBody
	GetEntities() *QueryContentResponseBodyEntities
	SetMatches(v *QueryContentResponseBodyMatches) *QueryContentResponseBody
	GetMatches() *QueryContentResponseBodyMatches
	SetMessage(v string) *QueryContentResponseBody
	GetMessage() *string
	SetRelations(v *QueryContentResponseBodyRelations) *QueryContentResponseBody
	GetRelations() *QueryContentResponseBodyRelations
	SetRequestId(v string) *QueryContentResponseBody
	GetRequestId() *string
	SetStatus(v string) *QueryContentResponseBody
	GetStatus() *string
	SetUsage(v *QueryContentResponseBodyUsage) *QueryContentResponseBody
	GetUsage() *QueryContentResponseBodyUsage
	SetWindowMatches(v *QueryContentResponseBodyWindowMatches) *QueryContentResponseBody
	GetWindowMatches() *QueryContentResponseBodyWindowMatches
}

type QueryContentResponseBody struct {
	// Number of tokens used for vectorization.
	//
	// > A token refers to the smallest unit into which the input text is divided; a token can be a word, a phrase, a punctuation mark, or a character, etc.
	//
	// example:
	//
	// 100
	EmbeddingTokens *string `json:"EmbeddingTokens,omitempty" xml:"EmbeddingTokens,omitempty"`
	// The entities.
	Entities *QueryContentResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Struct"`
	// The retrieved data.
	Matches *QueryContentResponseBodyMatches `json:"Matches,omitempty" xml:"Matches,omitempty" type:"Struct"`
	// Return message.
	//
	// example:
	//
	// success
	Message *string `json:"Message,omitempty" xml:"Message,omitempty"`
	// The list of relationship edges.
	Relations *QueryContentResponseBodyRelations `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Struct"`
	// The request ID.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The execution state of the operation. Valid values:
	//
	// 	- **false**: The operation fails.
	//
	// 	- **true**: The operation is successful.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// Resource usage for this query.
	Usage *QueryContentResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
	// List of windowed matches.
	WindowMatches *QueryContentResponseBodyWindowMatches `json:"WindowMatches,omitempty" xml:"WindowMatches,omitempty" type:"Struct"`
}

func (s QueryContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBody) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBody) GetEmbeddingTokens() *string {
	return s.EmbeddingTokens
}

func (s *QueryContentResponseBody) GetEntities() *QueryContentResponseBodyEntities {
	return s.Entities
}

func (s *QueryContentResponseBody) GetMatches() *QueryContentResponseBodyMatches {
	return s.Matches
}

func (s *QueryContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryContentResponseBody) GetRelations() *QueryContentResponseBodyRelations {
	return s.Relations
}

func (s *QueryContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryContentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QueryContentResponseBody) GetUsage() *QueryContentResponseBodyUsage {
	return s.Usage
}

func (s *QueryContentResponseBody) GetWindowMatches() *QueryContentResponseBodyWindowMatches {
	return s.WindowMatches
}

func (s *QueryContentResponseBody) SetEmbeddingTokens(v string) *QueryContentResponseBody {
	s.EmbeddingTokens = &v
	return s
}

func (s *QueryContentResponseBody) SetEntities(v *QueryContentResponseBodyEntities) *QueryContentResponseBody {
	s.Entities = v
	return s
}

func (s *QueryContentResponseBody) SetMatches(v *QueryContentResponseBodyMatches) *QueryContentResponseBody {
	s.Matches = v
	return s
}

func (s *QueryContentResponseBody) SetMessage(v string) *QueryContentResponseBody {
	s.Message = &v
	return s
}

func (s *QueryContentResponseBody) SetRelations(v *QueryContentResponseBodyRelations) *QueryContentResponseBody {
	s.Relations = v
	return s
}

func (s *QueryContentResponseBody) SetRequestId(v string) *QueryContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryContentResponseBody) SetStatus(v string) *QueryContentResponseBody {
	s.Status = &v
	return s
}

func (s *QueryContentResponseBody) SetUsage(v *QueryContentResponseBodyUsage) *QueryContentResponseBody {
	s.Usage = v
	return s
}

func (s *QueryContentResponseBody) SetWindowMatches(v *QueryContentResponseBodyWindowMatches) *QueryContentResponseBody {
	s.WindowMatches = v
	return s
}

func (s *QueryContentResponseBody) Validate() error {
	if s.Entities != nil {
		if err := s.Entities.Validate(); err != nil {
			return err
		}
	}
	if s.Matches != nil {
		if err := s.Matches.Validate(); err != nil {
			return err
		}
	}
	if s.Relations != nil {
		if err := s.Relations.Validate(); err != nil {
			return err
		}
	}
	if s.Usage != nil {
		if err := s.Usage.Validate(); err != nil {
			return err
		}
	}
	if s.WindowMatches != nil {
		if err := s.WindowMatches.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryContentResponseBodyEntities struct {
	Entities []*QueryContentResponseBodyEntitiesEntities `json:"entities,omitempty" xml:"entities,omitempty" type:"Repeated"`
}

func (s QueryContentResponseBodyEntities) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyEntities) GetEntities() []*QueryContentResponseBodyEntitiesEntities {
	return s.Entities
}

func (s *QueryContentResponseBodyEntities) SetEntities(v []*QueryContentResponseBodyEntitiesEntities) *QueryContentResponseBodyEntities {
	s.Entities = v
	return s
}

func (s *QueryContentResponseBodyEntities) Validate() error {
	if s.Entities != nil {
		for _, item := range s.Entities {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryContentResponseBodyEntitiesEntities struct {
	// Entity description.
	//
	// example:
	//
	// A former advisor at DeepMind.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The entity name.
	//
	// example:
	//
	// Dr. Wang
	Entity *string `json:"Entity,omitempty" xml:"Entity,omitempty"`
	// The file name.
	//
	// example:
	//
	// my_doc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The entity ID.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The entity type.
	//
	// example:
	//
	// Figure
	Type *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryContentResponseBodyEntitiesEntities) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyEntitiesEntities) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyEntitiesEntities) GetDescription() *string {
	return s.Description
}

func (s *QueryContentResponseBodyEntitiesEntities) GetEntity() *string {
	return s.Entity
}

func (s *QueryContentResponseBodyEntitiesEntities) GetFileName() *string {
	return s.FileName
}

func (s *QueryContentResponseBodyEntitiesEntities) GetId() *string {
	return s.Id
}

func (s *QueryContentResponseBodyEntitiesEntities) GetType() *string {
	return s.Type
}

func (s *QueryContentResponseBodyEntitiesEntities) SetDescription(v string) *QueryContentResponseBodyEntitiesEntities {
	s.Description = &v
	return s
}

func (s *QueryContentResponseBodyEntitiesEntities) SetEntity(v string) *QueryContentResponseBodyEntitiesEntities {
	s.Entity = &v
	return s
}

func (s *QueryContentResponseBodyEntitiesEntities) SetFileName(v string) *QueryContentResponseBodyEntitiesEntities {
	s.FileName = &v
	return s
}

func (s *QueryContentResponseBodyEntitiesEntities) SetId(v string) *QueryContentResponseBodyEntitiesEntities {
	s.Id = &v
	return s
}

func (s *QueryContentResponseBodyEntitiesEntities) SetType(v string) *QueryContentResponseBodyEntitiesEntities {
	s.Type = &v
	return s
}

func (s *QueryContentResponseBodyEntitiesEntities) Validate() error {
	return dara.Validate(s)
}

type QueryContentResponseBodyMatches struct {
	MatchList []*QueryContentResponseBodyMatchesMatchList `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
}

func (s QueryContentResponseBodyMatches) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyMatches) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyMatches) GetMatchList() []*QueryContentResponseBodyMatchesMatchList {
	return s.MatchList
}

func (s *QueryContentResponseBodyMatches) SetMatchList(v []*QueryContentResponseBodyMatchesMatchList) *QueryContentResponseBodyMatches {
	s.MatchList = v
	return s
}

func (s *QueryContentResponseBodyMatches) Validate() error {
	if s.MatchList != nil {
		for _, item := range s.MatchList {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryContentResponseBodyMatchesMatchList struct {
	// The content that is used for full-text search. If you leave this parameter empty, only vector search is used. If you do not leave this parameter empty, two-way retrieval based on vector search and full-text search is used.
	//
	// >  You must specify at least one of the Content and Vector parameters.
	//
	// example:
	//
	// Cloud-native data warehouse AnalyticDB PostgreSQL Edition provides a simple, fast, and cost-effective PB-level cloud data warehouse solution.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The name of the document.
	//
	// >  You can call the [ListDocuments](https://help.aliyun.com/document_detail/2618453.html) operation to query a list of documents.
	//
	// example:
	//
	// my_doc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The public URL of the query result image, valid for 2 hours
	//
	// example:
	//
	// https://xxx-cn-beijing.aliyuncs.com/image/test.png
	FileURL *string `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	// The unique ID of the vector data.
	//
	// example:
	//
	// doca-1234
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Metadata during document loader loading.
	//
	// example:
	//
	// {"page_pos": 1}
	LoaderMetadata *string `json:"LoaderMetadata,omitempty" xml:"LoaderMetadata,omitempty"`
	// The metadata.
	Metadata map[string]*string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	// Re-ranking score.
	//
	// example:
	//
	// 6.2345
	RerankScore *float64 `json:"RerankScore,omitempty" xml:"RerankScore,omitempty"`
	// Source of the retrieval results:
	//
	// - 1 indicates vector retrieval
	//
	// - 2 indicates full-text retrieval
	//
	// - 3 indicates dual-path recall
	//
	// example:
	//
	// 1
	RetrievalSource *int32 `json:"RetrievalSource,omitempty" xml:"RetrievalSource,omitempty"`
	// The similarity score of the data. It is related to the `l2, ip, or cosine` algorithm that is specified when you create an index.
	//
	// example:
	//
	// 0.12345
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// The vector data. The length of the value must be the same as that of the Dimension parameter in the [CreateCollection](https://help.aliyun.com/document_detail/2401497.html) operation.
	//
	// >  If you leave this parameter empty, only full-text search results are returned.
	Vector *QueryContentResponseBodyMatchesMatchListVector `json:"Vector,omitempty" xml:"Vector,omitempty" type:"Struct"`
}

func (s QueryContentResponseBodyMatchesMatchList) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyMatchesMatchList) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyMatchesMatchList) GetContent() *string {
	return s.Content
}

func (s *QueryContentResponseBodyMatchesMatchList) GetFileName() *string {
	return s.FileName
}

func (s *QueryContentResponseBodyMatchesMatchList) GetFileURL() *string {
	return s.FileURL
}

func (s *QueryContentResponseBodyMatchesMatchList) GetId() *string {
	return s.Id
}

func (s *QueryContentResponseBodyMatchesMatchList) GetLoaderMetadata() *string {
	return s.LoaderMetadata
}

func (s *QueryContentResponseBodyMatchesMatchList) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *QueryContentResponseBodyMatchesMatchList) GetRerankScore() *float64 {
	return s.RerankScore
}

func (s *QueryContentResponseBodyMatchesMatchList) GetRetrievalSource() *int32 {
	return s.RetrievalSource
}

func (s *QueryContentResponseBodyMatchesMatchList) GetScore() *float64 {
	return s.Score
}

func (s *QueryContentResponseBodyMatchesMatchList) GetVector() *QueryContentResponseBodyMatchesMatchListVector {
	return s.Vector
}

func (s *QueryContentResponseBodyMatchesMatchList) SetContent(v string) *QueryContentResponseBodyMatchesMatchList {
	s.Content = &v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) SetFileName(v string) *QueryContentResponseBodyMatchesMatchList {
	s.FileName = &v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) SetFileURL(v string) *QueryContentResponseBodyMatchesMatchList {
	s.FileURL = &v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) SetId(v string) *QueryContentResponseBodyMatchesMatchList {
	s.Id = &v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) SetLoaderMetadata(v string) *QueryContentResponseBodyMatchesMatchList {
	s.LoaderMetadata = &v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) SetMetadata(v map[string]*string) *QueryContentResponseBodyMatchesMatchList {
	s.Metadata = v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) SetRerankScore(v float64) *QueryContentResponseBodyMatchesMatchList {
	s.RerankScore = &v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) SetRetrievalSource(v int32) *QueryContentResponseBodyMatchesMatchList {
	s.RetrievalSource = &v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) SetScore(v float64) *QueryContentResponseBodyMatchesMatchList {
	s.Score = &v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) SetVector(v *QueryContentResponseBodyMatchesMatchListVector) *QueryContentResponseBodyMatchesMatchList {
	s.Vector = v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchList) Validate() error {
	if s.Vector != nil {
		if err := s.Vector.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryContentResponseBodyMatchesMatchListVector struct {
	VectorList []*float64 `json:"VectorList,omitempty" xml:"VectorList,omitempty" type:"Repeated"`
}

func (s QueryContentResponseBodyMatchesMatchListVector) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyMatchesMatchListVector) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyMatchesMatchListVector) GetVectorList() []*float64 {
	return s.VectorList
}

func (s *QueryContentResponseBodyMatchesMatchListVector) SetVectorList(v []*float64) *QueryContentResponseBodyMatchesMatchListVector {
	s.VectorList = v
	return s
}

func (s *QueryContentResponseBodyMatchesMatchListVector) Validate() error {
	return dara.Validate(s)
}

type QueryContentResponseBodyRelations struct {
	Relations []*QueryContentResponseBodyRelationsRelations `json:"relations,omitempty" xml:"relations,omitempty" type:"Repeated"`
}

func (s QueryContentResponseBodyRelations) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyRelations) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyRelations) GetRelations() []*QueryContentResponseBodyRelationsRelations {
	return s.Relations
}

func (s *QueryContentResponseBodyRelations) SetRelations(v []*QueryContentResponseBodyRelationsRelations) *QueryContentResponseBodyRelations {
	s.Relations = v
	return s
}

func (s *QueryContentResponseBodyRelations) Validate() error {
	if s.Relations != nil {
		for _, item := range s.Relations {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryContentResponseBodyRelationsRelations struct {
	// The description of the relationship edge.
	//
	// example:
	//
	// Dr. Wang previously served as an advisor at DeepMind.
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	// The file name.
	//
	// example:
	//
	// my_doc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// The ID of the relationship edge.
	//
	// example:
	//
	// 1
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// The source entity.
	//
	// example:
	//
	// Former DeepMind consultant
	SourceEntity *string `json:"SourceEntity,omitempty" xml:"SourceEntity,omitempty"`
	// The destination entity.
	//
	// example:
	//
	// Dr. Wang
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
}

func (s QueryContentResponseBodyRelationsRelations) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyRelationsRelations) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyRelationsRelations) GetDescription() *string {
	return s.Description
}

func (s *QueryContentResponseBodyRelationsRelations) GetFileName() *string {
	return s.FileName
}

func (s *QueryContentResponseBodyRelationsRelations) GetId() *string {
	return s.Id
}

func (s *QueryContentResponseBodyRelationsRelations) GetSourceEntity() *string {
	return s.SourceEntity
}

func (s *QueryContentResponseBodyRelationsRelations) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *QueryContentResponseBodyRelationsRelations) SetDescription(v string) *QueryContentResponseBodyRelationsRelations {
	s.Description = &v
	return s
}

func (s *QueryContentResponseBodyRelationsRelations) SetFileName(v string) *QueryContentResponseBodyRelationsRelations {
	s.FileName = &v
	return s
}

func (s *QueryContentResponseBodyRelationsRelations) SetId(v string) *QueryContentResponseBodyRelationsRelations {
	s.Id = &v
	return s
}

func (s *QueryContentResponseBodyRelationsRelations) SetSourceEntity(v string) *QueryContentResponseBodyRelationsRelations {
	s.SourceEntity = &v
	return s
}

func (s *QueryContentResponseBodyRelationsRelations) SetTargetEntity(v string) *QueryContentResponseBodyRelationsRelations {
	s.TargetEntity = &v
	return s
}

func (s *QueryContentResponseBodyRelationsRelations) Validate() error {
	return dara.Validate(s)
}

type QueryContentResponseBodyUsage struct {
	// The number of entries used for vectorization.
	//
	// > An entry refers to the number of processing items when performing vectorization on text or images. For example, processing one piece of text counts as 1 entry, while processing one image counts as 2 entries.
	//
	// example:
	//
	// 10
	EmbeddingEntries *string `json:"EmbeddingEntries,omitempty" xml:"EmbeddingEntries,omitempty"`
	// Number of tokens used for vectorization.
	//
	// > A token refers to the smallest unit into which the input text is divided; a token can be a word, a phrase, a punctuation mark, or a character, etc.
	//
	// example:
	//
	// 100
	EmbeddingTokens *string `json:"EmbeddingTokens,omitempty" xml:"EmbeddingTokens,omitempty"`
}

func (s QueryContentResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyUsage) GetEmbeddingEntries() *string {
	return s.EmbeddingEntries
}

func (s *QueryContentResponseBodyUsage) GetEmbeddingTokens() *string {
	return s.EmbeddingTokens
}

func (s *QueryContentResponseBodyUsage) SetEmbeddingEntries(v string) *QueryContentResponseBodyUsage {
	s.EmbeddingEntries = &v
	return s
}

func (s *QueryContentResponseBodyUsage) SetEmbeddingTokens(v string) *QueryContentResponseBodyUsage {
	s.EmbeddingTokens = &v
	return s
}

func (s *QueryContentResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}

type QueryContentResponseBodyWindowMatches struct {
	WindowMatches []*QueryContentResponseBodyWindowMatchesWindowMatches `json:"windowMatches,omitempty" xml:"windowMatches,omitempty" type:"Repeated"`
}

func (s QueryContentResponseBodyWindowMatches) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyWindowMatches) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyWindowMatches) GetWindowMatches() []*QueryContentResponseBodyWindowMatchesWindowMatches {
	return s.WindowMatches
}

func (s *QueryContentResponseBodyWindowMatches) SetWindowMatches(v []*QueryContentResponseBodyWindowMatchesWindowMatches) *QueryContentResponseBodyWindowMatches {
	s.WindowMatches = v
	return s
}

func (s *QueryContentResponseBodyWindowMatches) Validate() error {
	if s.WindowMatches != nil {
		for _, item := range s.WindowMatches {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryContentResponseBodyWindowMatchesWindowMatches struct {
	// List of individual top windowed matches.
	WindowMatch *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch `json:"WindowMatch,omitempty" xml:"WindowMatch,omitempty" type:"Struct"`
}

func (s QueryContentResponseBodyWindowMatchesWindowMatches) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyWindowMatchesWindowMatches) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatches) GetWindowMatch() *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch {
	return s.WindowMatch
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatches) SetWindowMatch(v *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch) *QueryContentResponseBodyWindowMatchesWindowMatches {
	s.WindowMatch = v
	return s
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatches) Validate() error {
	if s.WindowMatch != nil {
		if err := s.WindowMatch.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch struct {
	WindowMatch []*QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch `json:"windowMatch,omitempty" xml:"windowMatch,omitempty" type:"Repeated"`
}

func (s QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch) GetWindowMatch() []*QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch {
	return s.WindowMatch
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch) SetWindowMatch(v []*QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch {
	s.WindowMatch = v
	return s
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatch) Validate() error {
	if s.WindowMatch != nil {
		for _, item := range s.WindowMatch {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch struct {
	// Text content.
	//
	// example:
	//
	// AnalyticDB for PostgreSQL is a cloud-native data warehouse service that provides large-scale parallel processing (MPP) capabilities for massive online data analysis.
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// File name.
	//
	// example:
	//
	// my_doc.txt
	FileName *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	// Unique ID of the vector data.
	//
	// example:
	//
	// doca-2345
	Id *string `json:"Id,omitempty" xml:"Id,omitempty"`
	// Metadata information when the document loader was loaded.
	//
	// example:
	//
	// {"page_pos": 2}
	LoaderMetadata *string `json:"LoaderMetadata,omitempty" xml:"LoaderMetadata,omitempty"`
	// Metadata map.
	Metadata map[string]*string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
}

func (s QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) String() string {
	return dara.Prettify(s)
}

func (s QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) GoString() string {
	return s.String()
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) GetContent() *string {
	return s.Content
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) GetFileName() *string {
	return s.FileName
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) GetId() *string {
	return s.Id
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) GetLoaderMetadata() *string {
	return s.LoaderMetadata
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) SetContent(v string) *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch {
	s.Content = &v
	return s
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) SetFileName(v string) *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch {
	s.FileName = &v
	return s
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) SetId(v string) *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch {
	s.Id = &v
	return s
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) SetLoaderMetadata(v string) *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch {
	s.LoaderMetadata = &v
	return s
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) SetMetadata(v map[string]*string) *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch {
	s.Metadata = v
	return s
}

func (s *QueryContentResponseBodyWindowMatchesWindowMatchesWindowMatchWindowMatch) Validate() error {
	return dara.Validate(s)
}
