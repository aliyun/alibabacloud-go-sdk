// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIndexKey interface {
	dara.Model
	String() string
	GoString() string
	SetChn(v bool) *IndexKey
	GetChn() *bool
	SetEmbedding(v string) *IndexKey
	GetEmbedding() *string
	SetJsonKeys(v map[string]*IndexJsonKey) *IndexKey
	GetJsonKeys() map[string]*IndexJsonKey
	SetType(v string) *IndexKey
	GetType() *string
}

type IndexKey struct {
	Chn       *bool                    `json:"chn,omitempty" xml:"chn,omitempty"`
	Embedding *string                  `json:"embedding,omitempty" xml:"embedding,omitempty"`
	JsonKeys  map[string]*IndexJsonKey `json:"jsonKeys,omitempty" xml:"jsonKeys,omitempty"`
	Type      *string                  `json:"type,omitempty" xml:"type,omitempty"`
}

func (s IndexKey) String() string {
	return dara.Prettify(s)
}

func (s IndexKey) GoString() string {
	return s.String()
}

func (s *IndexKey) GetChn() *bool {
	return s.Chn
}

func (s *IndexKey) GetEmbedding() *string {
	return s.Embedding
}

func (s *IndexKey) GetJsonKeys() map[string]*IndexJsonKey {
	return s.JsonKeys
}

func (s *IndexKey) GetType() *string {
	return s.Type
}

func (s *IndexKey) SetChn(v bool) *IndexKey {
	s.Chn = &v
	return s
}

func (s *IndexKey) SetEmbedding(v string) *IndexKey {
	s.Embedding = &v
	return s
}

func (s *IndexKey) SetJsonKeys(v map[string]*IndexJsonKey) *IndexKey {
	s.JsonKeys = v
	return s
}

func (s *IndexKey) SetType(v string) *IndexKey {
	s.Type = &v
	return s
}

func (s *IndexKey) Validate() error {
	return dara.Validate(s)
}
