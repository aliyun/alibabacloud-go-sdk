// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iUpdateMetaTableIntroWikiRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContent(v string) *UpdateMetaTableIntroWikiRequest
	GetContent() *string
	SetTableGuid(v string) *UpdateMetaTableIntroWikiRequest
	GetTableGuid() *string
}

type UpdateMetaTableIntroWikiRequest struct {
	// The details of the instructions on how to use the metatable.
	//
	// This parameter is required.
	//
	// example:
	//
	// abc
	Content *string `json:"Content,omitempty" xml:"Content,omitempty"`
	// The GUID of the table. Specify the GUID in the odps.{projectName}.{tableName} format.
	//
	// This parameter is required.
	//
	// example:
	//
	// odps.test.table1
	TableGuid *string `json:"TableGuid,omitempty" xml:"TableGuid,omitempty"`
}

func (s UpdateMetaTableIntroWikiRequest) String() string {
	return dara.Prettify(s)
}

func (s UpdateMetaTableIntroWikiRequest) GoString() string {
	return s.String()
}

func (s *UpdateMetaTableIntroWikiRequest) GetContent() *string {
	return s.Content
}

func (s *UpdateMetaTableIntroWikiRequest) GetTableGuid() *string {
	return s.TableGuid
}

func (s *UpdateMetaTableIntroWikiRequest) SetContent(v string) *UpdateMetaTableIntroWikiRequest {
	s.Content = &v
	return s
}

func (s *UpdateMetaTableIntroWikiRequest) SetTableGuid(v string) *UpdateMetaTableIntroWikiRequest {
	s.TableGuid = &v
	return s
}

func (s *UpdateMetaTableIntroWikiRequest) Validate() error {
	return dara.Validate(s)
}
