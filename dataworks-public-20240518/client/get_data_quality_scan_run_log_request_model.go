// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDataQualityScanRunLogRequest interface {
	dara.Model
	String() string
	GoString() string
	SetId(v int64) *GetDataQualityScanRunLogRequest
	GetId() *int64
	SetOffset(v int64) *GetDataQualityScanRunLogRequest
	GetOffset() *int64
}

type GetDataQualityScanRunLogRequest struct {
	// The ID of the data quality monitoring run record.
	//
	// example:
	//
	// 10001
	Id *int64 `json:"Id,omitempty" xml:"Id,omitempty"`
	// The start position of the log, which is the offset relative to the beginning of the file. Unit: bytes. Each query returns a maximum of 512 KB of content.
	//
	// example:
	//
	// 200
	Offset *int64 `json:"Offset,omitempty" xml:"Offset,omitempty"`
}

func (s GetDataQualityScanRunLogRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDataQualityScanRunLogRequest) GoString() string {
	return s.String()
}

func (s *GetDataQualityScanRunLogRequest) GetId() *int64 {
	return s.Id
}

func (s *GetDataQualityScanRunLogRequest) GetOffset() *int64 {
	return s.Offset
}

func (s *GetDataQualityScanRunLogRequest) SetId(v int64) *GetDataQualityScanRunLogRequest {
	s.Id = &v
	return s
}

func (s *GetDataQualityScanRunLogRequest) SetOffset(v int64) *GetDataQualityScanRunLogRequest {
	s.Offset = &v
	return s
}

func (s *GetDataQualityScanRunLogRequest) Validate() error {
	return dara.Validate(s)
}
