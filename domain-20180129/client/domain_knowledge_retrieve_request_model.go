// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainKnowledgeRetrieveRequest interface {
	dara.Model
	String() string
	GoString() string
	SetGlobalTopN(v int32) *DomainKnowledgeRetrieveRequest
	GetGlobalTopN() *int32
	SetKeyword(v string) *DomainKnowledgeRetrieveRequest
	GetKeyword() *string
	SetSite(v string) *DomainKnowledgeRetrieveRequest
	GetSite() *string
}

type DomainKnowledgeRetrieveRequest struct {
	// Le nombre de résultats à renvoyer.
	//
	// example:
	//
	// 5
	GlobalTopN *int32 `json:"GlobalTopN,omitempty" xml:"GlobalTopN,omitempty"`
	// Les mots-clés à récupérer.
	//
	// This parameter is required.
	//
	// example:
	//
	// comment renouveler
	Keyword *string `json:"Keyword,omitempty" xml:"Keyword,omitempty"`
	// Les sites de la base de connaissances à interroger, y compris cn pour le national, intl pour l\\"international et all pour tous.
	//
	// example:
	//
	// all
	Site *string `json:"Site,omitempty" xml:"Site,omitempty"`
}

func (s DomainKnowledgeRetrieveRequest) String() string {
	return dara.Prettify(s)
}

func (s DomainKnowledgeRetrieveRequest) GoString() string {
	return s.String()
}

func (s *DomainKnowledgeRetrieveRequest) GetGlobalTopN() *int32 {
	return s.GlobalTopN
}

func (s *DomainKnowledgeRetrieveRequest) GetKeyword() *string {
	return s.Keyword
}

func (s *DomainKnowledgeRetrieveRequest) GetSite() *string {
	return s.Site
}

func (s *DomainKnowledgeRetrieveRequest) SetGlobalTopN(v int32) *DomainKnowledgeRetrieveRequest {
	s.GlobalTopN = &v
	return s
}

func (s *DomainKnowledgeRetrieveRequest) SetKeyword(v string) *DomainKnowledgeRetrieveRequest {
	s.Keyword = &v
	return s
}

func (s *DomainKnowledgeRetrieveRequest) SetSite(v string) *DomainKnowledgeRetrieveRequest {
	s.Site = &v
	return s
}

func (s *DomainKnowledgeRetrieveRequest) Validate() error {
	return dara.Validate(s)
}
