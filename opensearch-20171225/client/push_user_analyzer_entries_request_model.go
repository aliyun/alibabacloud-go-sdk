// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushUserAnalyzerEntriesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEntries(v []*PushUserAnalyzerEntriesRequestEntries) *PushUserAnalyzerEntriesRequest
	GetEntries() []*PushUserAnalyzerEntriesRequestEntries
	SetDryRun(v bool) *PushUserAnalyzerEntriesRequest
	GetDryRun() *bool
}

type PushUserAnalyzerEntriesRequest struct {
	// The entries of the custom analyzer.
	Entries []*PushUserAnalyzerEntriesRequestEntries `json:"entries,omitempty" xml:"entries,omitempty" type:"Repeated"`
	// Specifies whether to perform a dry run. This parameter is only used to check whether the data source is valid. Valid values: true and false.
	//
	// example:
	//
	// true
	DryRun *bool `json:"dryRun,omitempty" xml:"dryRun,omitempty"`
}

func (s PushUserAnalyzerEntriesRequest) String() string {
	return dara.Prettify(s)
}

func (s PushUserAnalyzerEntriesRequest) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesRequest) GetEntries() []*PushUserAnalyzerEntriesRequestEntries {
	return s.Entries
}

func (s *PushUserAnalyzerEntriesRequest) GetDryRun() *bool {
	return s.DryRun
}

func (s *PushUserAnalyzerEntriesRequest) SetEntries(v []*PushUserAnalyzerEntriesRequestEntries) *PushUserAnalyzerEntriesRequest {
	s.Entries = v
	return s
}

func (s *PushUserAnalyzerEntriesRequest) SetDryRun(v bool) *PushUserAnalyzerEntriesRequest {
	s.DryRun = &v
	return s
}

func (s *PushUserAnalyzerEntriesRequest) Validate() error {
	if s.Entries != nil {
		for _, item := range s.Entries {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PushUserAnalyzerEntriesRequestEntries struct {
	// The operation to be performed on the entries.
	//
	// Valid values:
	//
	// 	- add
	//
	// 	- delete
	//
	// example:
	//
	// "add"
	Cmd *string `json:"cmd,omitempty" xml:"cmd,omitempty"`
	// The key to be used to query entries.
	//
	// example:
	//
	// "testvalue"
	Key *string `json:"key,omitempty" xml:"key,omitempty"`
	// Specifies whether to further analyze the terms that are generated after the search query is analyzed.
	//
	// Default value: true.
	//
	// example:
	//
	// true
	SplitEnabled *bool `json:"splitEnabled,omitempty" xml:"splitEnabled,omitempty"`
	// The analysis result.
	//
	// example:
	//
	// "test value"
	Value *string `json:"value,omitempty" xml:"value,omitempty"`
}

func (s PushUserAnalyzerEntriesRequestEntries) String() string {
	return dara.Prettify(s)
}

func (s PushUserAnalyzerEntriesRequestEntries) GoString() string {
	return s.String()
}

func (s *PushUserAnalyzerEntriesRequestEntries) GetCmd() *string {
	return s.Cmd
}

func (s *PushUserAnalyzerEntriesRequestEntries) GetKey() *string {
	return s.Key
}

func (s *PushUserAnalyzerEntriesRequestEntries) GetSplitEnabled() *bool {
	return s.SplitEnabled
}

func (s *PushUserAnalyzerEntriesRequestEntries) GetValue() *string {
	return s.Value
}

func (s *PushUserAnalyzerEntriesRequestEntries) SetCmd(v string) *PushUserAnalyzerEntriesRequestEntries {
	s.Cmd = &v
	return s
}

func (s *PushUserAnalyzerEntriesRequestEntries) SetKey(v string) *PushUserAnalyzerEntriesRequestEntries {
	s.Key = &v
	return s
}

func (s *PushUserAnalyzerEntriesRequestEntries) SetSplitEnabled(v bool) *PushUserAnalyzerEntriesRequestEntries {
	s.SplitEnabled = &v
	return s
}

func (s *PushUserAnalyzerEntriesRequestEntries) SetValue(v string) *PushUserAnalyzerEntriesRequestEntries {
	s.Value = &v
	return s
}

func (s *PushUserAnalyzerEntriesRequestEntries) Validate() error {
	return dara.Validate(s)
}
