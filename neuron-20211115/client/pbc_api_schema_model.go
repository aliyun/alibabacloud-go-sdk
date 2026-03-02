// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPbcApiSchema interface {
	dara.Model
	String() string
	GoString() string
	SetApiSchema(v string) *PbcApiSchema
	GetApiSchema() *string
	SetId(v int64) *PbcApiSchema
	GetId() *int64
	SetPbcVersionId(v int64) *PbcApiSchema
	GetPbcVersionId() *int64
	SetRequestId(v string) *PbcApiSchema
	GetRequestId() *string
}

type PbcApiSchema struct {
	// This parameter is required.
	ApiSchema *string `json:"apiSchema,omitempty" xml:"apiSchema,omitempty"`
	// example:
	//
	// 1
	Id           *int64  `json:"id,omitempty" xml:"id,omitempty"`
	PbcVersionId *int64  `json:"pbcVersionId,omitempty" xml:"pbcVersionId,omitempty"`
	RequestId    *string `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s PbcApiSchema) String() string {
	return dara.Prettify(s)
}

func (s PbcApiSchema) GoString() string {
	return s.String()
}

func (s *PbcApiSchema) GetApiSchema() *string {
	return s.ApiSchema
}

func (s *PbcApiSchema) GetId() *int64 {
	return s.Id
}

func (s *PbcApiSchema) GetPbcVersionId() *int64 {
	return s.PbcVersionId
}

func (s *PbcApiSchema) GetRequestId() *string {
	return s.RequestId
}

func (s *PbcApiSchema) SetApiSchema(v string) *PbcApiSchema {
	s.ApiSchema = &v
	return s
}

func (s *PbcApiSchema) SetId(v int64) *PbcApiSchema {
	s.Id = &v
	return s
}

func (s *PbcApiSchema) SetPbcVersionId(v int64) *PbcApiSchema {
	s.PbcVersionId = &v
	return s
}

func (s *PbcApiSchema) SetRequestId(v string) *PbcApiSchema {
	s.RequestId = &v
	return s
}

func (s *PbcApiSchema) Validate() error {
	return dara.Validate(s)
}
