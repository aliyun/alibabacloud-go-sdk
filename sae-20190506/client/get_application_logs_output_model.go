// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetApplicationLogsOutput interface {
	dara.Model
	String() string
	GoString() string
	SetLogEntrys(v []*LogEntry) *GetApplicationLogsOutput
	GetLogEntrys() []*LogEntry
	SetNextOffset(v int64) *GetApplicationLogsOutput
	GetNextOffset() *int64
	SetRequestId(v string) *GetApplicationLogsOutput
	GetRequestId() *string
}

type GetApplicationLogsOutput struct {
	LogEntrys  []*LogEntry `json:"logEntrys,omitempty" xml:"logEntrys,omitempty" type:"Repeated"`
	NextOffset *int64      `json:"nextOffset,omitempty" xml:"nextOffset,omitempty"`
	RequestId  *string     `json:"requestId,omitempty" xml:"requestId,omitempty"`
}

func (s GetApplicationLogsOutput) String() string {
	return dara.Prettify(s)
}

func (s GetApplicationLogsOutput) GoString() string {
	return s.String()
}

func (s *GetApplicationLogsOutput) GetLogEntrys() []*LogEntry {
	return s.LogEntrys
}

func (s *GetApplicationLogsOutput) GetNextOffset() *int64 {
	return s.NextOffset
}

func (s *GetApplicationLogsOutput) GetRequestId() *string {
	return s.RequestId
}

func (s *GetApplicationLogsOutput) SetLogEntrys(v []*LogEntry) *GetApplicationLogsOutput {
	s.LogEntrys = v
	return s
}

func (s *GetApplicationLogsOutput) SetNextOffset(v int64) *GetApplicationLogsOutput {
	s.NextOffset = &v
	return s
}

func (s *GetApplicationLogsOutput) SetRequestId(v string) *GetApplicationLogsOutput {
	s.RequestId = &v
	return s
}

func (s *GetApplicationLogsOutput) Validate() error {
	if s.LogEntrys != nil {
		for _, item := range s.LogEntrys {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
