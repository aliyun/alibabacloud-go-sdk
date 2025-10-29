// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeLiveStreamTranscodeStreamNumResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetLazyTranscodedNumber(v int64) *DescribeLiveStreamTranscodeStreamNumResponseBody
	GetLazyTranscodedNumber() *int64
	SetRequestId(v string) *DescribeLiveStreamTranscodeStreamNumResponseBody
	GetRequestId() *string
	SetTotal(v int64) *DescribeLiveStreamTranscodeStreamNumResponseBody
	GetTotal() *int64
	SetTranscodeStreamCountDetails(v []*DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails) *DescribeLiveStreamTranscodeStreamNumResponseBody
	GetTranscodeStreamCountDetails() []*DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails
	SetTranscodedNumber(v int64) *DescribeLiveStreamTranscodeStreamNumResponseBody
	GetTranscodedNumber() *int64
	SetUntranscodeNumber(v int64) *DescribeLiveStreamTranscodeStreamNumResponseBody
	GetUntranscodeNumber() *int64
}

type DescribeLiveStreamTranscodeStreamNumResponseBody struct {
	// The number of streams for which transcoding is triggered by stream pulling.
	//
	// example:
	//
	// 10
	LazyTranscodedNumber *int64 `json:"LazyTranscodedNumber,omitempty" xml:"LazyTranscodedNumber,omitempty"`
	// The request ID.
	//
	// example:
	//
	// 150191A4-DD88-5941-B48C-9DF59E0A8B1F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The total number of streams.
	//
	// example:
	//
	// 57
	Total *int64 `json:"Total,omitempty" xml:"Total,omitempty"`
	// The details about the transcoding templates.
	TranscodeStreamCountDetails []*DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails `json:"TranscodeStreamCountDetails,omitempty" xml:"TranscodeStreamCountDetails,omitempty" type:"Repeated"`
	// The number of streams that are transcoded.
	//
	// example:
	//
	// 30
	TranscodedNumber *int64 `json:"TranscodedNumber,omitempty" xml:"TranscodedNumber,omitempty"`
	// The number of streams that are not transcoded.
	//
	// example:
	//
	// 27
	UntranscodeNumber *int64 `json:"UntranscodeNumber,omitempty" xml:"UntranscodeNumber,omitempty"`
}

func (s DescribeLiveStreamTranscodeStreamNumResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeStreamNumResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) GetLazyTranscodedNumber() *int64 {
	return s.LazyTranscodedNumber
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) GetTotal() *int64 {
	return s.Total
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) GetTranscodeStreamCountDetails() []*DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails {
	return s.TranscodeStreamCountDetails
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) GetTranscodedNumber() *int64 {
	return s.TranscodedNumber
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) GetUntranscodeNumber() *int64 {
	return s.UntranscodeNumber
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) SetLazyTranscodedNumber(v int64) *DescribeLiveStreamTranscodeStreamNumResponseBody {
	s.LazyTranscodedNumber = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) SetRequestId(v string) *DescribeLiveStreamTranscodeStreamNumResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) SetTotal(v int64) *DescribeLiveStreamTranscodeStreamNumResponseBody {
	s.Total = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) SetTranscodeStreamCountDetails(v []*DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails) *DescribeLiveStreamTranscodeStreamNumResponseBody {
	s.TranscodeStreamCountDetails = v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) SetTranscodedNumber(v int64) *DescribeLiveStreamTranscodeStreamNumResponseBody {
	s.TranscodedNumber = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) SetUntranscodeNumber(v int64) *DescribeLiveStreamTranscodeStreamNumResponseBody {
	s.UntranscodeNumber = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBody) Validate() error {
	if s.TranscodeStreamCountDetails != nil {
		for _, item := range s.TranscodeStreamCountDetails {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails struct {
	// The number of streams that use the transcoding template.
	//
	// example:
	//
	// 30
	Count *int32 `json:"Count,omitempty" xml:"Count,omitempty"`
	// The name of the transcoding template.
	//
	// example:
	//
	// template_name
	Template *string `json:"Template,omitempty" xml:"Template,omitempty"`
}

func (s DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails) String() string {
	return dara.Prettify(s)
}

func (s DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails) GoString() string {
	return s.String()
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails) GetCount() *int32 {
	return s.Count
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails) GetTemplate() *string {
	return s.Template
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails) SetCount(v int32) *DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails {
	s.Count = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails) SetTemplate(v string) *DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails {
	s.Template = &v
	return s
}

func (s *DescribeLiveStreamTranscodeStreamNumResponseBodyTranscodeStreamCountDetails) Validate() error {
	return dara.Validate(s)
}
