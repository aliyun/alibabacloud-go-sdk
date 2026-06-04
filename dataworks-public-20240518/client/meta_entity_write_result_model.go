// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iMetaEntityWriteResult interface {
	dara.Model
	String() string
	GoString() string
	SetEntityType(v string) *MetaEntityWriteResult
	GetEntityType() *string
	SetErrorMessage(v string) *MetaEntityWriteResult
	GetErrorMessage() *string
	SetId(v string) *MetaEntityWriteResult
	GetId() *string
	SetName(v string) *MetaEntityWriteResult
	GetName() *string
	SetSuccess(v bool) *MetaEntityWriteResult
	GetSuccess() *bool
}

type MetaEntityWriteResult struct {
	EntityType   *string `json:"EntityType,omitempty" xml:"EntityType,omitempty"`
	ErrorMessage *string `json:"ErrorMessage,omitempty" xml:"ErrorMessage,omitempty"`
	Id           *string `json:"Id,omitempty" xml:"Id,omitempty"`
	Name         *string `json:"Name,omitempty" xml:"Name,omitempty"`
	Success      *bool   `json:"Success,omitempty" xml:"Success,omitempty"`
}

func (s MetaEntityWriteResult) String() string {
	return dara.Prettify(s)
}

func (s MetaEntityWriteResult) GoString() string {
	return s.String()
}

func (s *MetaEntityWriteResult) GetEntityType() *string {
	return s.EntityType
}

func (s *MetaEntityWriteResult) GetErrorMessage() *string {
	return s.ErrorMessage
}

func (s *MetaEntityWriteResult) GetId() *string {
	return s.Id
}

func (s *MetaEntityWriteResult) GetName() *string {
	return s.Name
}

func (s *MetaEntityWriteResult) GetSuccess() *bool {
	return s.Success
}

func (s *MetaEntityWriteResult) SetEntityType(v string) *MetaEntityWriteResult {
	s.EntityType = &v
	return s
}

func (s *MetaEntityWriteResult) SetErrorMessage(v string) *MetaEntityWriteResult {
	s.ErrorMessage = &v
	return s
}

func (s *MetaEntityWriteResult) SetId(v string) *MetaEntityWriteResult {
	s.Id = &v
	return s
}

func (s *MetaEntityWriteResult) SetName(v string) *MetaEntityWriteResult {
	s.Name = &v
	return s
}

func (s *MetaEntityWriteResult) SetSuccess(v bool) *MetaEntityWriteResult {
	s.Success = &v
	return s
}

func (s *MetaEntityWriteResult) Validate() error {
	return dara.Validate(s)
}
