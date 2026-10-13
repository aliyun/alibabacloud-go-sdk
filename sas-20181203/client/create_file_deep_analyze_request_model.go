// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateFileDeepAnalyzeRequest interface {
	dara.Model
	String() string
	GoString() string
	SetContext(v string) *CreateFileDeepAnalyzeRequest
	GetContext() *string
	SetEventId(v int64) *CreateFileDeepAnalyzeRequest
	GetEventId() *int64
	SetHash(v string) *CreateFileDeepAnalyzeRequest
	GetHash() *string
	SetSource(v string) *CreateFileDeepAnalyzeRequest
	GetSource() *string
	SetSourceIp(v string) *CreateFileDeepAnalyzeRequest
	GetSourceIp() *string
}

type CreateFileDeepAnalyzeRequest struct {
	// The request context information.
	//
	// example:
	//
	// {}
	Context *string `json:"Context,omitempty" xml:"Context,omitempty"`
	// The event ID. If a file-related alert event already exists, you can use this parameter to trigger deep analysis.
	//
	// example:
	//
	// 100000001
	EventId *int64 `json:"EventId,omitempty" xml:"EventId,omitempty"`
	// The file hash. If the file already exists in the cloud, you can use this parameter to trigger deep analysis.
	//
	// example:
	//
	// d41d8cd98f00b204e9800998ecf8427e
	Hash *string `json:"Hash,omitempty" xml:"Hash,omitempty"`
	// The request source. The system verifies the purchased edition based on the request source. Valid values:
	//
	// - **edr**: AgenticEDR
	//
	// - **secOpsAgent**: Security O&M agent
	//
	// example:
	//
	// edr
	Source *string `json:"Source,omitempty" xml:"Source,omitempty"`
	// The IP address of the access source.
	//
	// example:
	//
	// 192.168.XX.XX
	SourceIp *string `json:"SourceIp,omitempty" xml:"SourceIp,omitempty"`
}

func (s CreateFileDeepAnalyzeRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateFileDeepAnalyzeRequest) GoString() string {
	return s.String()
}

func (s *CreateFileDeepAnalyzeRequest) GetContext() *string {
	return s.Context
}

func (s *CreateFileDeepAnalyzeRequest) GetEventId() *int64 {
	return s.EventId
}

func (s *CreateFileDeepAnalyzeRequest) GetHash() *string {
	return s.Hash
}

func (s *CreateFileDeepAnalyzeRequest) GetSource() *string {
	return s.Source
}

func (s *CreateFileDeepAnalyzeRequest) GetSourceIp() *string {
	return s.SourceIp
}

func (s *CreateFileDeepAnalyzeRequest) SetContext(v string) *CreateFileDeepAnalyzeRequest {
	s.Context = &v
	return s
}

func (s *CreateFileDeepAnalyzeRequest) SetEventId(v int64) *CreateFileDeepAnalyzeRequest {
	s.EventId = &v
	return s
}

func (s *CreateFileDeepAnalyzeRequest) SetHash(v string) *CreateFileDeepAnalyzeRequest {
	s.Hash = &v
	return s
}

func (s *CreateFileDeepAnalyzeRequest) SetSource(v string) *CreateFileDeepAnalyzeRequest {
	s.Source = &v
	return s
}

func (s *CreateFileDeepAnalyzeRequest) SetSourceIp(v string) *CreateFileDeepAnalyzeRequest {
	s.SourceIp = &v
	return s
}

func (s *CreateFileDeepAnalyzeRequest) Validate() error {
	return dara.Validate(s)
}
