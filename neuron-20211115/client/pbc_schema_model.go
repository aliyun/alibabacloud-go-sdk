// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcSchema interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *PbcSchema
	GetId() *int64
	SetPbcVersionId(v int64) *PbcSchema
	GetPbcVersionId() *int64
	SetRequestId(v string) *PbcSchema
	GetRequestId() *string
	SetSchema(v string) *PbcSchema
	GetSchema() *string
}

type PbcSchema struct {
	// example:
	//
	// 1
	Id *int64 `json:"id,omitempty" xml:"id,omitempty"`
	// example:
	//
	// 2
	PbcVersionId *int64  `json:"pbcVersionId,omitempty" xml:"pbcVersionId,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
	// This parameter is required.
	Schema *string `json:"schema,omitempty" xml:"schema,omitempty"`
}

func (s PbcSchema) String() string {
	return dara.Prettify(s)
}

func (s PbcSchema) GoString() string {
	return s.String()
}

func (s *PbcSchema) GetId() *int64 {
	return s.Id
}

func (s *PbcSchema) GetPbcVersionId() *int64 {
	return s.PbcVersionId
}

func (s *PbcSchema) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcSchema) GetSchema() *string {
	return s.Schema
}

func (s *PbcSchema) SetId(v int64) *PbcSchema {
	s.Id = &v
	return s
}

func (s *PbcSchema) SetPbcVersionId(v int64) *PbcSchema {
	s.PbcVersionId = &v
	return s
}

func (s *PbcSchema) SetRequestId(v string) *PbcSchema {
	s.RequestId = &v
	return s
}

func (s *PbcSchema) SetSchema(v string) *PbcSchema {
	s.Schema = &v
	return s
}

func (s *PbcSchema) Validate() error {
	return dara.Validate(s)
}
