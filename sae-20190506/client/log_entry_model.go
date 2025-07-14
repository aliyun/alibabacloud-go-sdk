// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iLogEntry interface {
	dara.Model
	String() string
	GoString() string
	SetInstanceID(v string) *LogEntry
	GetInstanceID() *string
	SetMessage(v string) *LogEntry
	GetMessage() *string
	SetOffset(v int64) *LogEntry
	GetOffset() *int64
	SetPackID(v string) *LogEntry
	GetPackID() *string
	SetPackMeta(v string) *LogEntry
	GetPackMeta() *string
	SetQualifier(v string) *LogEntry
	GetQualifier() *string
	SetTimestamp(v int32) *LogEntry
	GetTimestamp() *int32
	SetVersionID(v string) *LogEntry
	GetVersionID() *string
}

type LogEntry struct {
	InstanceID *string `json:"instanceID,omitempty" xml:"instanceID,omitempty"`
	Message    *string `json:"message,omitempty" xml:"message,omitempty"`
	Offset     *int64  `json:"offset,omitempty" xml:"offset,omitempty"`
	PackID     *string `json:"packID,omitempty" xml:"packID,omitempty"`
	PackMeta   *string `json:"packMeta,omitempty" xml:"packMeta,omitempty"`
	Qualifier  *string `json:"qualifier,omitempty" xml:"qualifier,omitempty"`
	Timestamp  *int32  `json:"timestamp,omitempty" xml:"timestamp,omitempty"`
	VersionID  *string `json:"versionID,omitempty" xml:"versionID,omitempty"`
}

func (s LogEntry) String() string {
	return dara.Prettify(s)
}

func (s LogEntry) GoString() string {
	return s.String()
}

func (s *LogEntry) GetInstanceID() *string {
	return s.InstanceID
}

func (s *LogEntry) GetMessage() *string {
	return s.Message
}

func (s *LogEntry) GetOffset() *int64 {
	return s.Offset
}

func (s *LogEntry) GetPackID() *string {
	return s.PackID
}

func (s *LogEntry) GetPackMeta() *string {
	return s.PackMeta
}

func (s *LogEntry) GetQualifier() *string {
	return s.Qualifier
}

func (s *LogEntry) GetTimestamp() *int32 {
	return s.Timestamp
}

func (s *LogEntry) GetVersionID() *string {
	return s.VersionID
}

func (s *LogEntry) SetInstanceID(v string) *LogEntry {
	s.InstanceID = &v
	return s
}

func (s *LogEntry) SetMessage(v string) *LogEntry {
	s.Message = &v
	return s
}

func (s *LogEntry) SetOffset(v int64) *LogEntry {
	s.Offset = &v
	return s
}

func (s *LogEntry) SetPackID(v string) *LogEntry {
	s.PackID = &v
	return s
}

func (s *LogEntry) SetPackMeta(v string) *LogEntry {
	s.PackMeta = &v
	return s
}

func (s *LogEntry) SetQualifier(v string) *LogEntry {
	s.Qualifier = &v
	return s
}

func (s *LogEntry) SetTimestamp(v int32) *LogEntry {
	s.Timestamp = &v
	return s
}

func (s *LogEntry) SetVersionID(v string) *LogEntry {
	s.VersionID = &v
	return s
}

func (s *LogEntry) Validate() error {
	return dara.Validate(s)
}
