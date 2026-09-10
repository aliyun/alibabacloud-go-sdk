// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetSqlTransTableMetaInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSourceDialect(v string) *GetSqlTransTableMetaInfoRequest
	GetSourceDialect() *string
	SetSourceSqlScript(v string) *GetSqlTransTableMetaInfoRequest
	GetSourceSqlScript() *string
	SetTargetDialect(v string) *GetSqlTransTableMetaInfoRequest
	GetTargetDialect() *string
}

type GetSqlTransTableMetaInfoRequest struct {
	// The source SQL dialect type.
	//
	// example:
	//
	// hive
	SourceDialect *string `json:"sourceDialect,omitempty" xml:"sourceDialect,omitempty"`
	// The source script content. You must Base64-encode the script before passing it in. The server decodes the content before parsing.
	//
	// example:
	//
	// SELECT 	- FROM t;
	SourceSqlScript *string `json:"sourceSqlScript,omitempty" xml:"sourceSqlScript,omitempty"`
	// The target SQL dialect type.
	//
	// example:
	//
	// hive
	TargetDialect *string `json:"targetDialect,omitempty" xml:"targetDialect,omitempty"`
}

func (s GetSqlTransTableMetaInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetSqlTransTableMetaInfoRequest) GoString() string {
	return s.String()
}

func (s *GetSqlTransTableMetaInfoRequest) GetSourceDialect() *string {
	return s.SourceDialect
}

func (s *GetSqlTransTableMetaInfoRequest) GetSourceSqlScript() *string {
	return s.SourceSqlScript
}

func (s *GetSqlTransTableMetaInfoRequest) GetTargetDialect() *string {
	return s.TargetDialect
}

func (s *GetSqlTransTableMetaInfoRequest) SetSourceDialect(v string) *GetSqlTransTableMetaInfoRequest {
	s.SourceDialect = &v
	return s
}

func (s *GetSqlTransTableMetaInfoRequest) SetSourceSqlScript(v string) *GetSqlTransTableMetaInfoRequest {
	s.SourceSqlScript = &v
	return s
}

func (s *GetSqlTransTableMetaInfoRequest) SetTargetDialect(v string) *GetSqlTransTableMetaInfoRequest {
	s.TargetDialect = &v
	return s
}

func (s *GetSqlTransTableMetaInfoRequest) Validate() error {
	return dara.Validate(s)
}
