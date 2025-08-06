// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetMediaDNAResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDNAResult(v *GetMediaDNAResultResponseBodyDNAResult) *GetMediaDNAResultResponseBody
	GetDNAResult() *GetMediaDNAResultResponseBodyDNAResult
	SetRequestId(v string) *GetMediaDNAResultResponseBody
	GetRequestId() *string
}

type GetMediaDNAResultResponseBody struct {
	// The media fingerprinting results.
	DNAResult *GetMediaDNAResultResponseBodyDNAResult `json:"DNAResult,omitempty" xml:"DNAResult,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 63FC4896-E956-4B*****7D-134FF1BC597A
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetMediaDNAResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetMediaDNAResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBody) GetDNAResult() *GetMediaDNAResultResponseBodyDNAResult {
	return s.DNAResult
}

func (s *GetMediaDNAResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetMediaDNAResultResponseBody) SetDNAResult(v *GetMediaDNAResultResponseBodyDNAResult) *GetMediaDNAResultResponseBody {
	s.DNAResult = v
	return s
}

func (s *GetMediaDNAResultResponseBody) SetRequestId(v string) *GetMediaDNAResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetMediaDNAResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetMediaDNAResultResponseBodyDNAResult struct {
	// The video fingerprint recognition result.
	VideoDNA []*GetMediaDNAResultResponseBodyDNAResultVideoDNA `json:"VideoDNA,omitempty" xml:"VideoDNA,omitempty" type:"Repeated"`
}

func (s GetMediaDNAResultResponseBodyDNAResult) String() string {
	return dara.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResult) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResult) GetVideoDNA() []*GetMediaDNAResultResponseBodyDNAResultVideoDNA {
	return s.VideoDNA
}

func (s *GetMediaDNAResultResponseBodyDNAResult) SetVideoDNA(v []*GetMediaDNAResultResponseBodyDNAResultVideoDNA) *GetMediaDNAResultResponseBodyDNAResult {
	s.VideoDNA = v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResult) Validate() error {
	return dara.Validate(s)
}

type GetMediaDNAResultResponseBodyDNAResultVideoDNA struct {
	// The details of the matched video. Information such as the location and duration of the video is returned.
	Detail []*GetMediaDNAResultResponseBodyDNAResultVideoDNADetail `json:"Detail,omitempty" xml:"Detail,omitempty" type:"Repeated"`
	// The ID of the video that has a similar fingerprint.
	//
	// example:
	//
	// 6ad8987da46f4b*****490ce2873745
	PrimaryKey *string `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	// The similarity between the fingerprints of the input video and the matched video. 1 indicates that the fingerprints of the two videos are the same.
	//
	// example:
	//
	// 0.98
	Similarity *string `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNA) String() string {
	return dara.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNA) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) GetDetail() []*GetMediaDNAResultResponseBodyDNAResultVideoDNADetail {
	return s.Detail
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) GetSimilarity() *string {
	return s.Similarity
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) SetDetail(v []*GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) *GetMediaDNAResultResponseBodyDNAResultVideoDNA {
	s.Detail = v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) SetPrimaryKey(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNA {
	s.PrimaryKey = &v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) SetSimilarity(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNA {
	s.Similarity = &v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNA) Validate() error {
	return dara.Validate(s)
}

type GetMediaDNAResultResponseBodyDNAResultVideoDNADetail struct {
	// The start time and duration of the matched video.
	Duplication *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication `json:"Duplication,omitempty" xml:"Duplication,omitempty" type:"Struct"`
	// The start time and duration of the input video.
	Input *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) String() string {
	return dara.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) GetDuplication() *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication {
	return s.Duplication
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) GetInput() *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput {
	return s.Input
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) SetDuplication(v *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail {
	s.Duplication = v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) SetInput(v *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail {
	s.Input = v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetail) Validate() error {
	return dara.Validate(s)
}

type GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication struct {
	// The duration of the video. Unit: seconds.
	//
	// example:
	//
	// 12.0
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the video. Unit: seconds.
	//
	// example:
	//
	// 2.0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) String() string {
	return dara.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) GetStart() *string {
	return s.Start
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) SetDuration(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication {
	s.Duration = &v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) SetStart(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication {
	s.Start = &v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailDuplication) Validate() error {
	return dara.Validate(s)
}

type GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput struct {
	// The duration of the video. Unit: seconds.
	//
	// example:
	//
	// 12.0
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	// The start time of the video. Unit: seconds.
	//
	// example:
	//
	// 2.0
	Start *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) String() string {
	return dara.Prettify(s)
}

func (s GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) GoString() string {
	return s.String()
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) GetDuration() *string {
	return s.Duration
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) GetStart() *string {
	return s.Start
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) SetDuration(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput {
	s.Duration = &v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) SetStart(v string) *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput {
	s.Start = &v
	return s
}

func (s *GetMediaDNAResultResponseBodyDNAResultVideoDNADetailInput) Validate() error {
	return dara.Validate(s)
}
