// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPostInnerConvertRequest interface {
	dara.Model
	String() string
	GoString() string
	SetSqlConvertMap(v map[string]interface{}) *PostInnerConvertRequest
	GetSqlConvertMap() map[string]interface{}
	SetSrcDataSourceName(v string) *PostInnerConvertRequest
	GetSrcDataSourceName() *string
	SetTgtDataSourceName(v string) *PostInnerConvertRequest
	GetTgtDataSourceName() *string
}

type PostInnerConvertRequest struct {
	// SQL node type mapping, where the key is the source node type and the value is the target node type. If not empty, it will be assembled into `workflow.converter.sqlNodeTypeMapping` in `innerConvertConfig` and written via the task configuration update interface after creating the scheduling transformation task.
	SqlConvertMap map[string]interface{} `json:"sqlConvertMap,omitempty" xml:"sqlConvertMap,omitempty"`
	// Source data source name, i.e., the name of the scheduling data source at the source end of the transformation task.
	//
	// example:
	//
	// SourceDS1
	SrcDataSourceName *string `json:"srcDataSourceName,omitempty" xml:"srcDataSourceName,omitempty"`
	// Target data source name, i.e., the name of the scheduling data source at the target end of the transformation task.
	//
	// example:
	//
	// TargetDS1
	TgtDataSourceName *string `json:"tgtDataSourceName,omitempty" xml:"tgtDataSourceName,omitempty"`
}

func (s PostInnerConvertRequest) String() string {
	return dara.Prettify(s)
}

func (s PostInnerConvertRequest) GoString() string {
	return s.String()
}

func (s *PostInnerConvertRequest) GetSqlConvertMap() map[string]interface{} {
	return s.SqlConvertMap
}

func (s *PostInnerConvertRequest) GetSrcDataSourceName() *string {
	return s.SrcDataSourceName
}

func (s *PostInnerConvertRequest) GetTgtDataSourceName() *string {
	return s.TgtDataSourceName
}

func (s *PostInnerConvertRequest) SetSqlConvertMap(v map[string]interface{}) *PostInnerConvertRequest {
	s.SqlConvertMap = v
	return s
}

func (s *PostInnerConvertRequest) SetSrcDataSourceName(v string) *PostInnerConvertRequest {
	s.SrcDataSourceName = &v
	return s
}

func (s *PostInnerConvertRequest) SetTgtDataSourceName(v string) *PostInnerConvertRequest {
	s.TgtDataSourceName = &v
	return s
}

func (s *PostInnerConvertRequest) Validate() error {
	return dara.Validate(s)
}
