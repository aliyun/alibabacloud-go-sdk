// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetIndexMonitorRequest interface {
	dara.Model
	String() string
	GoString() string
	SetEndTimestamp(v int64) *GetIndexMonitorRequest
	GetEndTimestamp() *int64
	SetIndexId(v string) *GetIndexMonitorRequest
	GetIndexId() *string
	SetStartTimestamp(v int64) *GetIndexMonitorRequest
	GetStartTimestamp() *int64
}

type GetIndexMonitorRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// 1767604500
	EndTimestamp *int64 `json:"EndTimestamp,omitempty" xml:"EndTimestamp,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// kb-123456xxxx
	IndexId *string `json:"IndexId,omitempty" xml:"IndexId,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// 1767604500
	StartTimestamp *int64 `json:"StartTimestamp,omitempty" xml:"StartTimestamp,omitempty"`
}

func (s GetIndexMonitorRequest) String() string {
	return dara.Prettify(s)
}

func (s GetIndexMonitorRequest) GoString() string {
	return s.String()
}

func (s *GetIndexMonitorRequest) GetEndTimestamp() *int64 {
	return s.EndTimestamp
}

func (s *GetIndexMonitorRequest) GetIndexId() *string {
	return s.IndexId
}

func (s *GetIndexMonitorRequest) GetStartTimestamp() *int64 {
	return s.StartTimestamp
}

func (s *GetIndexMonitorRequest) SetEndTimestamp(v int64) *GetIndexMonitorRequest {
	s.EndTimestamp = &v
	return s
}

func (s *GetIndexMonitorRequest) SetIndexId(v string) *GetIndexMonitorRequest {
	s.IndexId = &v
	return s
}

func (s *GetIndexMonitorRequest) SetStartTimestamp(v int64) *GetIndexMonitorRequest {
	s.StartTimestamp = &v
	return s
}

func (s *GetIndexMonitorRequest) Validate() error {
	return dara.Validate(s)
}
