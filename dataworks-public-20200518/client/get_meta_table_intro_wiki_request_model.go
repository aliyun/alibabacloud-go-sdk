// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMetaTableIntroWikiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetTableGuid(v string) *GetMetaTableIntroWikiRequest
	GetTableGuid() *string
	SetWikiVersion(v int64) *GetMetaTableIntroWikiRequest
	GetWikiVersion() *int64
}

type GetMetaTableIntroWikiRequest struct {
	// The GUID of the metatable.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.engine_name.table_name
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
	// The version of the instructions.
	//
	// example:
	//
	// 1
	WikiVersion *int64 `json:"WikiVersion,omitempty" xml:"WikiVersion,omitempty"`
}

func (s GetMetaTableIntroWikiRequest) String() string {
	return dara.Prettify(s)
}

func (s GetMetaTableIntroWikiRequest) GoString() string {
	return s.String()
}

func (s *GetMetaTableIntroWikiRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *GetMetaTableIntroWikiRequest) GetWikiVersion() *int64 {
	return s.WikiVersion
}

func (s *GetMetaTableIntroWikiRequest) SetTableGuid(v string) *GetMetaTableIntroWikiRequest {
	s.TableGuid = &v
	return s
}

func (s *GetMetaTableIntroWikiRequest) SetWikiVersion(v int64) *GetMetaTableIntroWikiRequest {
	s.WikiVersion = &v
	return s
}

func (s *GetMetaTableIntroWikiRequest) Validate() error {
	return dara.Validate(s)
}
