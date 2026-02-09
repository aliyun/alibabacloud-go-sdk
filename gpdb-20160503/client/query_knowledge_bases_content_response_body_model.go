// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iQueryKnowledgeBasesContentResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetEmbeddingTokens(v string) *QueryKnowledgeBasesContentResponseBody
	GetEmbeddingTokens() *string
	SetEntities(v *QueryKnowledgeBasesContentResponseBodyEntities) *QueryKnowledgeBasesContentResponseBody
	GetEntities() *QueryKnowledgeBasesContentResponseBodyEntities
	SetMatches(v *QueryKnowledgeBasesContentResponseBodyMatches) *QueryKnowledgeBasesContentResponseBody
	GetMatches() *QueryKnowledgeBasesContentResponseBodyMatches
	SetMessage(v string) *QueryKnowledgeBasesContentResponseBody
	GetMessage() *string
	SetRelations(v *QueryKnowledgeBasesContentResponseBodyRelations) *QueryKnowledgeBasesContentResponseBody
	GetRelations() *QueryKnowledgeBasesContentResponseBodyRelations
	SetRequestId(v string) *QueryKnowledgeBasesContentResponseBody
	GetRequestId() *string
	SetStatus(v string) *QueryKnowledgeBasesContentResponseBody
	GetStatus() *string
	SetUsage(v *QueryKnowledgeBasesContentResponseBodyUsage) *QueryKnowledgeBasesContentResponseBody
	GetUsage() *QueryKnowledgeBasesContentResponseBodyUsage
}

type QueryKnowledgeBasesContentResponseBody struct {
	// The number of tokens that are used during vectorization.
	//
	// >  A token is the minimum unit for segmenting text. A token can be a word, phrase, punctuation, or character.
	//
	// example:
	//
	// 100
	EmbeddingTokens *string                                         `json:"EmbeddingTokens,omitempty" xml:"EmbeddingTokens,omitempty"`
	Entities        *QueryKnowledgeBasesContentResponseBodyEntities `json:"Entities,omitempty" xml:"Entities,omitempty" type:"Struct"`
	Matches         *QueryKnowledgeBasesContentResponseBodyMatches  `json:"Matches,omitempty" xml:"Matches,omitempty" type:"Struct"`
	// The returned information.
	//
	// example:
	//
	// success
	Message   *string                                          `json:"Message,omitempty" xml:"Message,omitempty"`
	Relations *QueryKnowledgeBasesContentResponseBodyRelations `json:"Relations,omitempty" xml:"Relations,omitempty" type:"Struct"`
	// The unique ID of the request.
	//
	// example:
	//
	// ABB39CC3-4488-4857-905D-2E4A051D0521
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The status of the operation. Valid values:
	//
	// 	- **success**.
	//
	// 	- **fail**.
	//
	// example:
	//
	// success
	Status *string `json:"Status,omitempty" xml:"Status,omitempty"`
	// The total number of tokens that are consumed by this query.
	Usage *QueryKnowledgeBasesContentResponseBodyUsage `json:"Usage,omitempty" xml:"Usage,omitempty" type:"Struct"`
}

func (s QueryKnowledgeBasesContentResponseBody) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentResponseBody) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentResponseBody) GetEmbeddingTokens() *string {
	return s.EmbeddingTokens
}

func (s *QueryKnowledgeBasesContentResponseBody) GetEntities() *QueryKnowledgeBasesContentResponseBodyEntities {
	return s.Entities
}

func (s *QueryKnowledgeBasesContentResponseBody) GetMatches() *QueryKnowledgeBasesContentResponseBodyMatches {
	return s.Matches
}

func (s *QueryKnowledgeBasesContentResponseBody) GetMessage() *string {
	return s.Message
}

func (s *QueryKnowledgeBasesContentResponseBody) GetRelations() *QueryKnowledgeBasesContentResponseBodyRelations {
	return s.Relations
}

func (s *QueryKnowledgeBasesContentResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *QueryKnowledgeBasesContentResponseBody) GetStatus() *string {
	return s.Status
}

func (s *QueryKnowledgeBasesContentResponseBody) GetUsage() *QueryKnowledgeBasesContentResponseBodyUsage {
	return s.Usage
}

func (s *QueryKnowledgeBasesContentResponseBody) SetEmbeddingTokens(v string) *QueryKnowledgeBasesContentResponseBody {
	s.EmbeddingTokens = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBody) SetEntities(v *QueryKnowledgeBasesContentResponseBodyEntities) *QueryKnowledgeBasesContentResponseBody {
	s.Entities = v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBody) SetMatches(v *QueryKnowledgeBasesContentResponseBodyMatches) *QueryKnowledgeBasesContentResponseBody {
	s.Matches = v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBody) SetMessage(v string) *QueryKnowledgeBasesContentResponseBody {
	s.Message = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBody) SetRelations(v *QueryKnowledgeBasesContentResponseBodyRelations) *QueryKnowledgeBasesContentResponseBody {
	s.Relations = v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBody) SetRequestId(v string) *QueryKnowledgeBasesContentResponseBody {
	s.RequestId = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBody) SetStatus(v string) *QueryKnowledgeBasesContentResponseBody {
	s.Status = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBody) SetUsage(v *QueryKnowledgeBasesContentResponseBodyUsage) *QueryKnowledgeBasesContentResponseBody {
	s.Usage = v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBody) Validate() error {
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
	return nil
}

type QueryKnowledgeBasesContentResponseBodyEntities struct {
	Entities []*QueryKnowledgeBasesContentResponseBodyEntitiesEntities `json:"entities,omitempty" xml:"entities,omitempty" type:"Repeated"`
}

func (s QueryKnowledgeBasesContentResponseBodyEntities) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentResponseBodyEntities) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentResponseBodyEntities) GetEntities() []*QueryKnowledgeBasesContentResponseBodyEntitiesEntities {
	return s.Entities
}

func (s *QueryKnowledgeBasesContentResponseBodyEntities) SetEntities(v []*QueryKnowledgeBasesContentResponseBodyEntitiesEntities) *QueryKnowledgeBasesContentResponseBodyEntities {
	s.Entities = v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyEntities) Validate() error {
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

type QueryKnowledgeBasesContentResponseBodyEntitiesEntities struct {
	Description *string `json:"Description,omitempty" xml:"Description,omitempty"`
	Entity      *string `json:"Entity,omitempty" xml:"Entity,omitempty"`
	FileName    *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id          *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Type        *string `json:"Type,omitempty" xml:"Type,omitempty"`
}

func (s QueryKnowledgeBasesContentResponseBodyEntitiesEntities) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentResponseBodyEntitiesEntities) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) GetDescription() *string {
	return s.Description
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) GetEntity() *string {
	return s.Entity
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) GetFileName() *string {
	return s.FileName
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) GetId() *string {
	return s.Id
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) GetType() *string {
	return s.Type
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) SetDescription(v string) *QueryKnowledgeBasesContentResponseBodyEntitiesEntities {
	s.Description = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) SetEntity(v string) *QueryKnowledgeBasesContentResponseBodyEntitiesEntities {
	s.Entity = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) SetFileName(v string) *QueryKnowledgeBasesContentResponseBodyEntitiesEntities {
	s.FileName = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) SetId(v string) *QueryKnowledgeBasesContentResponseBodyEntitiesEntities {
	s.Id = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) SetType(v string) *QueryKnowledgeBasesContentResponseBodyEntitiesEntities {
	s.Type = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyEntitiesEntities) Validate() error {
	return dara.Validate(s)
}

type QueryKnowledgeBasesContentResponseBodyMatches struct {
	MatchList []*QueryKnowledgeBasesContentResponseBodyMatchesMatchList `json:"MatchList,omitempty" xml:"MatchList,omitempty" type:"Repeated"`
}

func (s QueryKnowledgeBasesContentResponseBodyMatches) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentResponseBodyMatches) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentResponseBodyMatches) GetMatchList() []*QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	return s.MatchList
}

func (s *QueryKnowledgeBasesContentResponseBodyMatches) SetMatchList(v []*QueryKnowledgeBasesContentResponseBodyMatchesMatchList) *QueryKnowledgeBasesContentResponseBodyMatches {
	s.MatchList = v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatches) Validate() error {
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

type QueryKnowledgeBasesContentResponseBodyMatchesMatchList struct {
	Content         *string            `json:"Content,omitempty" xml:"Content,omitempty"`
	FileName        *string            `json:"FileName,omitempty" xml:"FileName,omitempty"`
	FileURL         *string            `json:"FileURL,omitempty" xml:"FileURL,omitempty"`
	Id              *string            `json:"Id,omitempty" xml:"Id,omitempty"`
	LoaderMetadata  *string            `json:"LoaderMetadata,omitempty" xml:"LoaderMetadata,omitempty"`
	Metadata        map[string]*string `json:"Metadata,omitempty" xml:"Metadata,omitempty"`
	RerankScore     *float64           `json:"RerankScore,omitempty" xml:"RerankScore,omitempty"`
	RetrievalSource *int32             `json:"RetrievalSource,omitempty" xml:"RetrievalSource,omitempty"`
	Score           *float64           `json:"Score,omitempty" xml:"Score,omitempty"`
}

func (s QueryKnowledgeBasesContentResponseBodyMatchesMatchList) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GetContent() *string {
	return s.Content
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GetFileName() *string {
	return s.FileName
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GetFileURL() *string {
	return s.FileURL
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GetId() *string {
	return s.Id
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GetLoaderMetadata() *string {
	return s.LoaderMetadata
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GetMetadata() map[string]*string {
	return s.Metadata
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GetRerankScore() *float64 {
	return s.RerankScore
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GetRetrievalSource() *int32 {
	return s.RetrievalSource
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) GetScore() *float64 {
	return s.Score
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) SetContent(v string) *QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	s.Content = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) SetFileName(v string) *QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	s.FileName = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) SetFileURL(v string) *QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	s.FileURL = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) SetId(v string) *QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	s.Id = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) SetLoaderMetadata(v string) *QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	s.LoaderMetadata = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) SetMetadata(v map[string]*string) *QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	s.Metadata = v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) SetRerankScore(v float64) *QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	s.RerankScore = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) SetRetrievalSource(v int32) *QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	s.RetrievalSource = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) SetScore(v float64) *QueryKnowledgeBasesContentResponseBodyMatchesMatchList {
	s.Score = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyMatchesMatchList) Validate() error {
	return dara.Validate(s)
}

type QueryKnowledgeBasesContentResponseBodyRelations struct {
	Relations []*QueryKnowledgeBasesContentResponseBodyRelationsRelations `json:"relations,omitempty" xml:"relations,omitempty" type:"Repeated"`
}

func (s QueryKnowledgeBasesContentResponseBodyRelations) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentResponseBodyRelations) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentResponseBodyRelations) GetRelations() []*QueryKnowledgeBasesContentResponseBodyRelationsRelations {
	return s.Relations
}

func (s *QueryKnowledgeBasesContentResponseBodyRelations) SetRelations(v []*QueryKnowledgeBasesContentResponseBodyRelationsRelations) *QueryKnowledgeBasesContentResponseBodyRelations {
	s.Relations = v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyRelations) Validate() error {
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

type QueryKnowledgeBasesContentResponseBodyRelationsRelations struct {
	Description  *string `json:"Description,omitempty" xml:"Description,omitempty"`
	FileName     *string `json:"FileName,omitempty" xml:"FileName,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	SourceEntity *string `json:"SourceEntity,omitempty" xml:"SourceEntity,omitempty"`
	TargetEntity *string `json:"TargetEntity,omitempty" xml:"TargetEntity,omitempty"`
}

func (s QueryKnowledgeBasesContentResponseBodyRelationsRelations) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentResponseBodyRelationsRelations) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) GetDescription() *string {
	return s.Description
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) GetFileName() *string {
	return s.FileName
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) GetId() *string {
	return s.Id
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) GetSourceEntity() *string {
	return s.SourceEntity
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) GetTargetEntity() *string {
	return s.TargetEntity
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) SetDescription(v string) *QueryKnowledgeBasesContentResponseBodyRelationsRelations {
	s.Description = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) SetFileName(v string) *QueryKnowledgeBasesContentResponseBodyRelationsRelations {
	s.FileName = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) SetId(v string) *QueryKnowledgeBasesContentResponseBodyRelationsRelations {
	s.Id = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) SetSourceEntity(v string) *QueryKnowledgeBasesContentResponseBodyRelationsRelations {
	s.SourceEntity = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) SetTargetEntity(v string) *QueryKnowledgeBasesContentResponseBodyRelationsRelations {
	s.TargetEntity = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyRelationsRelations) Validate() error {
	return dara.Validate(s)
}

type QueryKnowledgeBasesContentResponseBodyUsage struct {
	// The number of entries that are used during vectorization.
	//
	// >  An entry refers to a single unit of vectorization processing. Processing one text input counts as 1 entry, while processing one image counts as 2 entries.
	//
	// example:
	//
	// 10
	EmbeddingEntries *string `json:"EmbeddingEntries,omitempty" xml:"EmbeddingEntries,omitempty"`
	// The number of tokens that are used for vectorization.
	//
	// >  A token is the minimum unit for splitting text. A token can be a word, phrase, punctuation, or character.
	//
	// example:
	//
	// 475
	EmbeddingTokens *string `json:"EmbeddingTokens,omitempty" xml:"EmbeddingTokens,omitempty"`
}

func (s QueryKnowledgeBasesContentResponseBodyUsage) String() string {
	return dara.Prettify(s)
}

func (s QueryKnowledgeBasesContentResponseBodyUsage) GoString() string {
	return s.String()
}

func (s *QueryKnowledgeBasesContentResponseBodyUsage) GetEmbeddingEntries() *string {
	return s.EmbeddingEntries
}

func (s *QueryKnowledgeBasesContentResponseBodyUsage) GetEmbeddingTokens() *string {
	return s.EmbeddingTokens
}

func (s *QueryKnowledgeBasesContentResponseBodyUsage) SetEmbeddingEntries(v string) *QueryKnowledgeBasesContentResponseBodyUsage {
	s.EmbeddingEntries = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyUsage) SetEmbeddingTokens(v string) *QueryKnowledgeBasesContentResponseBodyUsage {
	s.EmbeddingTokens = &v
	return s
}

func (s *QueryKnowledgeBasesContentResponseBodyUsage) Validate() error {
	return dara.Validate(s)
}
