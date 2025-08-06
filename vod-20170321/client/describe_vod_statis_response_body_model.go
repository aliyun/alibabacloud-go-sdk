// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeVodStatisResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetCdnData(v *DescribeVodStatisResponseBodyCdnData) *DescribeVodStatisResponseBody
	GetCdnData() *DescribeVodStatisResponseBodyCdnData
	SetRequestId(v string) *DescribeVodStatisResponseBody
	GetRequestId() *string
	SetStorageData(v *DescribeVodStatisResponseBodyStorageData) *DescribeVodStatisResponseBody
	GetStorageData() *DescribeVodStatisResponseBodyStorageData
	SetTranscodeData(v *DescribeVodStatisResponseBodyTranscodeData) *DescribeVodStatisResponseBody
	GetTranscodeData() *DescribeVodStatisResponseBodyTranscodeData
}

type DescribeVodStatisResponseBody struct {
	CdnData       *DescribeVodStatisResponseBodyCdnData       `json:"CdnData,omitempty" xml:"CdnData,omitempty" type:"Struct"`
	RequestId     *string                                     `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	StorageData   *DescribeVodStatisResponseBodyStorageData   `json:"StorageData,omitempty" xml:"StorageData,omitempty" type:"Struct"`
	TranscodeData *DescribeVodStatisResponseBodyTranscodeData `json:"TranscodeData,omitempty" xml:"TranscodeData,omitempty" type:"Struct"`
}

func (s DescribeVodStatisResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStatisResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeVodStatisResponseBody) GetCdnData() *DescribeVodStatisResponseBodyCdnData {
	return s.CdnData
}

func (s *DescribeVodStatisResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeVodStatisResponseBody) GetStorageData() *DescribeVodStatisResponseBodyStorageData {
	return s.StorageData
}

func (s *DescribeVodStatisResponseBody) GetTranscodeData() *DescribeVodStatisResponseBodyTranscodeData {
	return s.TranscodeData
}

func (s *DescribeVodStatisResponseBody) SetCdnData(v *DescribeVodStatisResponseBodyCdnData) *DescribeVodStatisResponseBody {
	s.CdnData = v
	return s
}

func (s *DescribeVodStatisResponseBody) SetRequestId(v string) *DescribeVodStatisResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeVodStatisResponseBody) SetStorageData(v *DescribeVodStatisResponseBodyStorageData) *DescribeVodStatisResponseBody {
	s.StorageData = v
	return s
}

func (s *DescribeVodStatisResponseBody) SetTranscodeData(v *DescribeVodStatisResponseBodyTranscodeData) *DescribeVodStatisResponseBody {
	s.TranscodeData = v
	return s
}

func (s *DescribeVodStatisResponseBody) Validate() error {
	return dara.Validate(s)
}

type DescribeVodStatisResponseBodyCdnData struct {
	Bps     *float32 `json:"Bps,omitempty" xml:"Bps,omitempty"`
	Traffic *int64   `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeVodStatisResponseBodyCdnData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStatisResponseBodyCdnData) GoString() string {
	return s.String()
}

func (s *DescribeVodStatisResponseBodyCdnData) GetBps() *float32 {
	return s.Bps
}

func (s *DescribeVodStatisResponseBodyCdnData) GetTraffic() *int64 {
	return s.Traffic
}

func (s *DescribeVodStatisResponseBodyCdnData) SetBps(v float32) *DescribeVodStatisResponseBodyCdnData {
	s.Bps = &v
	return s
}

func (s *DescribeVodStatisResponseBodyCdnData) SetTraffic(v int64) *DescribeVodStatisResponseBodyCdnData {
	s.Traffic = &v
	return s
}

func (s *DescribeVodStatisResponseBodyCdnData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodStatisResponseBodyStorageData struct {
	Size    *int64 `json:"Size,omitempty" xml:"Size,omitempty"`
	Traffic *int64 `json:"Traffic,omitempty" xml:"Traffic,omitempty"`
}

func (s DescribeVodStatisResponseBodyStorageData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStatisResponseBodyStorageData) GoString() string {
	return s.String()
}

func (s *DescribeVodStatisResponseBodyStorageData) GetSize() *int64 {
	return s.Size
}

func (s *DescribeVodStatisResponseBodyStorageData) GetTraffic() *int64 {
	return s.Traffic
}

func (s *DescribeVodStatisResponseBodyStorageData) SetSize(v int64) *DescribeVodStatisResponseBodyStorageData {
	s.Size = &v
	return s
}

func (s *DescribeVodStatisResponseBodyStorageData) SetTraffic(v int64) *DescribeVodStatisResponseBodyStorageData {
	s.Traffic = &v
	return s
}

func (s *DescribeVodStatisResponseBodyStorageData) Validate() error {
	return dara.Validate(s)
}

type DescribeVodStatisResponseBodyTranscodeData struct {
	Duration *int64 `json:"Duration,omitempty" xml:"Duration,omitempty"`
}

func (s DescribeVodStatisResponseBodyTranscodeData) String() string {
	return dara.Prettify(s)
}

func (s DescribeVodStatisResponseBodyTranscodeData) GoString() string {
	return s.String()
}

func (s *DescribeVodStatisResponseBodyTranscodeData) GetDuration() *int64 {
	return s.Duration
}

func (s *DescribeVodStatisResponseBodyTranscodeData) SetDuration(v int64) *DescribeVodStatisResponseBodyTranscodeData {
	s.Duration = &v
	return s
}

func (s *DescribeVodStatisResponseBodyTranscodeData) Validate() error {
	return dara.Validate(s)
}
