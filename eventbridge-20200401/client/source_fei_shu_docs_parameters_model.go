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
	// The App ID of the custom Lark application. Obtain this value after creating an application on the Lark Open Platform.
	//
	// example:
	//
	// cli_a946c046xxxx
	AppId *string `json:"AppId,omitempty" xml:"AppId,omitempty"`
	// The App Secret of the custom Lark application. Obtain this value after creating an application on the Lark Open Platform.
	//
	// example:
	//
	// yO3hEYiSjkBVxxxx
	AppSecret *string `json:"AppSecret,omitempty" xml:"AppSecret,omitempty"`
	// The name of the Lark knowledge space. An exact match is required.
	//
	// example:
	//
	// Product Documentation Center
	KnowledgeSpaceName *string `json:"KnowledgeSpaceName,omitempty" xml:"KnowledgeSpaceName,omitempty"`
	// The document loading mode. Valid values: block: splits the document by blocks, with each block as a separate event. single (default): treats the entire document as a single event, with metadata extension keys such as file name, modification time, and original link. Use this mode when importing into an EventHouse knowledge base.
	//
	// example:
	//
	// single
	LoadMode *string `json:"LoadMode,omitempty" xml:"LoadMode,omitempty"`
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
