// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iIcebergSnapshot interface {
	dara.Model
	String() string
	GoString() string
	SetAddedRows(v int64) *IcebergSnapshot
	GetAddedRows() *int64
	SetId(v int64) *IcebergSnapshot
	GetId() *int64
	SetOperation(v string) *IcebergSnapshot
	GetOperation() *string
	SetParentId(v int64) *IcebergSnapshot
	GetParentId() *int64
	SetSchemaId(v int64) *IcebergSnapshot
	GetSchemaId() *int64
	SetSequenceNumber(v int64) *IcebergSnapshot
	GetSequenceNumber() *int64
	SetSummary(v map[string]*string) *IcebergSnapshot
	GetSummary() map[string]*string
	SetTimestampMillis(v int64) *IcebergSnapshot
	GetTimestampMillis() *int64
}

type IcebergSnapshot struct {
	AddedRows       *int64             `json:"addedRows,omitempty" xml:"addedRows,omitempty"`
	Id              *int64             `json:"id,omitempty" xml:"id,omitempty"`
	Operation       *string            `json:"operation,omitempty" xml:"operation,omitempty"`
	ParentId        *int64             `json:"parentId,omitempty" xml:"parentId,omitempty"`
	SchemaId        *int64             `json:"schemaId,omitempty" xml:"schemaId,omitempty"`
	SequenceNumber  *int64             `json:"sequenceNumber,omitempty" xml:"sequenceNumber,omitempty"`
	Summary         map[string]*string `json:"summary,omitempty" xml:"summary,omitempty"`
	TimestampMillis *int64             `json:"timestampMillis,omitempty" xml:"timestampMillis,omitempty"`
}

func (s IcebergSnapshot) String() string {
	return dara.Prettify(s)
}

func (s IcebergSnapshot) GoString() string {
	return s.String()
}

func (s *IcebergSnapshot) GetAddedRows() *int64 {
	return s.AddedRows
}

func (s *IcebergSnapshot) GetId() *int64 {
	return s.Id
}

func (s *IcebergSnapshot) GetOperation() *string {
	return s.Operation
}

func (s *IcebergSnapshot) GetParentId() *int64 {
	return s.ParentId
}

func (s *IcebergSnapshot) GetSchemaId() *int64 {
	return s.SchemaId
}

func (s *IcebergSnapshot) GetSequenceNumber() *int64 {
	return s.SequenceNumber
}

func (s *IcebergSnapshot) GetSummary() map[string]*string {
	return s.Summary
}

func (s *IcebergSnapshot) GetTimestampMillis() *int64 {
	return s.TimestampMillis
}

func (s *IcebergSnapshot) SetAddedRows(v int64) *IcebergSnapshot {
	s.AddedRows = &v
	return s
}

func (s *IcebergSnapshot) SetId(v int64) *IcebergSnapshot {
	s.Id = &v
	return s
}

func (s *IcebergSnapshot) SetOperation(v string) *IcebergSnapshot {
	s.Operation = &v
	return s
}

func (s *IcebergSnapshot) SetParentId(v int64) *IcebergSnapshot {
	s.ParentId = &v
	return s
}

func (s *IcebergSnapshot) SetSchemaId(v int64) *IcebergSnapshot {
	s.SchemaId = &v
	return s
}

func (s *IcebergSnapshot) SetSequenceNumber(v int64) *IcebergSnapshot {
	s.SequenceNumber = &v
	return s
}

func (s *IcebergSnapshot) SetSummary(v map[string]*string) *IcebergSnapshot {
	s.Summary = v
	return s
}

func (s *IcebergSnapshot) SetTimestampMillis(v int64) *IcebergSnapshot {
	s.TimestampMillis = &v
	return s
}

func (s *IcebergSnapshot) Validate() error {
	return dara.Validate(s)
}
