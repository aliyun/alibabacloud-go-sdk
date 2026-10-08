// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDomainKnowledgeRetrieveResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetData(v []*DomainKnowledgeRetrieveResponseBodyData) *DomainKnowledgeRetrieveResponseBody
	GetData() []*DomainKnowledgeRetrieveResponseBodyData
	SetRequestId(v string) *DomainKnowledgeRetrieveResponseBody
	GetRequestId() *string
}

type DomainKnowledgeRetrieveResponseBody struct {
	// La liste des résultats récupérés.
	Data []*DomainKnowledgeRetrieveResponseBodyData `json:"Data,omitempty" xml:"Data,omitempty" type:"Repeated"`
	// L\\"identifiant de la requête.
	//
	// example:
	//
	// 019FABCB-6C7D-18FE-AA42-922BFC9555D9
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DomainKnowledgeRetrieveResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DomainKnowledgeRetrieveResponseBody) GoString() string {
	return s.String()
}

func (s *DomainKnowledgeRetrieveResponseBody) GetData() []*DomainKnowledgeRetrieveResponseBodyData {
	return s.Data
}

func (s *DomainKnowledgeRetrieveResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DomainKnowledgeRetrieveResponseBody) SetData(v []*DomainKnowledgeRetrieveResponseBodyData) *DomainKnowledgeRetrieveResponseBody {
	s.Data = v
	return s
}

func (s *DomainKnowledgeRetrieveResponseBody) SetRequestId(v string) *DomainKnowledgeRetrieveResponseBody {
	s.RequestId = &v
	return s
}

func (s *DomainKnowledgeRetrieveResponseBody) Validate() error {
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

type DomainKnowledgeRetrieveResponseBodyData struct {
	// Le score du texte récupéré ; plus le score est élevé, plus le résultat est pertinent.
	//
	// example:
	//
	// 0.6
	Score *float64 `json:"Score,omitempty" xml:"Score,omitempty"`
	// La source des résultats récupérés.
	//
	// example:
	//
	// Base de connaissances de l\\"activité nationale
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// Le texte récupéré.
	//
	// example:
	//
	// Sur le site national d\\"Alibaba Cloud, le renouvellement de nom de domaine peut être effectué via les méthodes suivantes
	Text *string `json:"Text,omitempty" xml:"Text,omitempty"`
}

func (s DomainKnowledgeRetrieveResponseBodyData) String() string {
	return dara.Prettify(s)
}

func (s DomainKnowledgeRetrieveResponseBodyData) GoString() string {
	return s.String()
}

func (s *DomainKnowledgeRetrieveResponseBodyData) GetScore() *float64 {
	return s.Score
}

func (s *DomainKnowledgeRetrieveResponseBodyData) GetSource() *string {
	return s.Source
}

func (s *DomainKnowledgeRetrieveResponseBodyData) GetText() *string {
	return s.Text
}

func (s *DomainKnowledgeRetrieveResponseBodyData) SetScore(v float64) *DomainKnowledgeRetrieveResponseBodyData {
	s.Score = &v
	return s
}

func (s *DomainKnowledgeRetrieveResponseBodyData) SetSource(v string) *DomainKnowledgeRetrieveResponseBodyData {
	s.Source = &v
	return s
}

func (s *DomainKnowledgeRetrieveResponseBodyData) SetText(v string) *DomainKnowledgeRetrieveResponseBodyData {
	s.Text = &v
	return s
}

func (s *DomainKnowledgeRetrieveResponseBodyData) Validate() error {
	return dara.Validate(s)
}
