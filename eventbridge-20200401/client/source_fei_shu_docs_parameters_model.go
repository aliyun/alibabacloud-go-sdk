// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iSourceFeiShuDocsParameters interface {
	dara.Model
	String() string
	GoString() string
	SetAppId(v string) *SourceFeiShuDocsParameters
	GetAppId() *string
	SetAppSecret(v string) *SourceFeiShuDocsParameters
	GetAppSecret() *string
	SetKnowledgeSpaceName(v string) *SourceFeiShuDocsParameters
	GetKnowledgeSpaceName() *string
	SetLoadMode(v string) *SourceFeiShuDocsParameters
	GetLoadMode() *string
}

type SourceFeiShuDocsParameters struct {
	AppId              *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	AppSecret          *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	KnowledgeSpaceName *string `json:"KnowledgeSpaceName,omitempty" xml:"KnowledgeSpaceName,omitempty"`
	LoadMode           *string `json:"LoadMode,omitempty" xml:"LoadMode,omitempty"`
}

func (s SourceFeiShuDocsParameters) String() string {
	return dara.Prettify(s)
}

func (s SourceFeiShuDocsParameters) GoString() string {
	return s.String()
}

func (s *SourceFeiShuDocsParameters) GetAppId() *string {
	return s.AppId
}

func (s *SourceFeiShuDocsParameters) GetAppSecret() *string {
	return s.AppSecret
}

func (s *SourceFeiShuDocsParameters) GetKnowledgeSpaceName() *string {
	return s.KnowledgeSpaceName
}

func (s *SourceFeiShuDocsParameters) GetLoadMode() *string {
	return s.LoadMode
}

func (s *SourceFeiShuDocsParameters) SetAppId(v string) *SourceFeiShuDocsParameters {
	s.AppId = &v
	return s
}

func (s *SourceFeiShuDocsParameters) SetAppSecret(v string) *SourceFeiShuDocsParameters {
	s.AppSecret = &v
	return s
}

func (s *SourceFeiShuDocsParameters) SetKnowledgeSpaceName(v string) *SourceFeiShuDocsParameters {
	s.KnowledgeSpaceName = &v
	return s
}

func (s *SourceFeiShuDocsParameters) SetLoadMode(v string) *SourceFeiShuDocsParameters {
	s.LoadMode = &v
	return s
}

func (s *SourceFeiShuDocsParameters) Validate() error {
	return dara.Validate(s)
}
