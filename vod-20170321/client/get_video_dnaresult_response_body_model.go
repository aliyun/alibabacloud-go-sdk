// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetVideoDNAResultResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetAIVideoDNAResult(v *GetVideoDNAResultResponseBodyAIVideoDNAResult) *GetVideoDNAResultResponseBody
	GetAIVideoDNAResult() *GetVideoDNAResultResponseBodyAIVideoDNAResult
	SetRequestId(v string) *GetVideoDNAResultResponseBody
	GetRequestId() *string
}

type GetVideoDNAResultResponseBody struct {
	AIVideoDNAResult *GetVideoDNAResultResponseBodyAIVideoDNAResult `json:"AIVideoDNAResult,omitempty" xml:"AIVideoDNAResult,omitempty" type:"Struct"`
	RequestId        *string                                        `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s GetVideoDNAResultResponseBody) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDNAResultResponseBody) GoString() string {
	return s.String()
}

func (s *GetVideoDNAResultResponseBody) GetAIVideoDNAResult() *GetVideoDNAResultResponseBodyAIVideoDNAResult {
	return s.AIVideoDNAResult
}

func (s *GetVideoDNAResultResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *GetVideoDNAResultResponseBody) SetAIVideoDNAResult(v *GetVideoDNAResultResponseBodyAIVideoDNAResult) *GetVideoDNAResultResponseBody {
	s.AIVideoDNAResult = v
	return s
}

func (s *GetVideoDNAResultResponseBody) SetRequestId(v string) *GetVideoDNAResultResponseBody {
	s.RequestId = &v
	return s
}

func (s *GetVideoDNAResultResponseBody) Validate() error {
	return dara.Validate(s)
}

type GetVideoDNAResultResponseBodyAIVideoDNAResult struct {
	FpShots []*GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots `json:"FpShots,omitempty" xml:"FpShots,omitempty" type:"Repeated"`
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResult) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResult) GoString() string {
	return s.String()
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResult) GetFpShots() []*GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots {
	return s.FpShots
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResult) SetFpShots(v []*GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) *GetVideoDNAResultResponseBodyAIVideoDNAResult {
	s.FpShots = v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResult) Validate() error {
	return dara.Validate(s)
}

type GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots struct {
	FpShotSlices []*GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices `json:"FpShotSlices,omitempty" xml:"FpShotSlices,omitempty" type:"Repeated"`
	PrimaryKey   *string                                                             `json:"PrimaryKey,omitempty" xml:"PrimaryKey,omitempty"`
	Similarity   *string                                                             `json:"Similarity,omitempty" xml:"Similarity,omitempty"`
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) GoString() string {
	return s.String()
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) GetFpShotSlices() []*GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices {
	return s.FpShotSlices
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) GetPrimaryKey() *string {
	return s.PrimaryKey
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) GetSimilarity() *string {
	return s.Similarity
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) SetFpShotSlices(v []*GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices) *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots {
	s.FpShotSlices = v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) SetPrimaryKey(v string) *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots {
	s.PrimaryKey = &v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) SetSimilarity(v string) *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots {
	s.Similarity = &v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShots) Validate() error {
	return dara.Validate(s)
}

type GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices struct {
	Duplication *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication `json:"Duplication,omitempty" xml:"Duplication,omitempty" type:"Struct"`
	Input       *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput       `json:"Input,omitempty" xml:"Input,omitempty" type:"Struct"`
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices) GoString() string {
	return s.String()
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices) GetDuplication() *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication {
	return s.Duplication
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices) GetInput() *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput {
	return s.Input
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices) SetDuplication(v *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication) *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices {
	s.Duplication = v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices) SetInput(v *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput) *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices {
	s.Input = v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlices) Validate() error {
	return dara.Validate(s)
}

type GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication) GoString() string {
	return s.String()
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication) GetDuration() *string {
	return s.Duration
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication) GetStart() *string {
	return s.Start
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication) SetDuration(v string) *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication {
	s.Duration = &v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication) SetStart(v string) *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication {
	s.Start = &v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesDuplication) Validate() error {
	return dara.Validate(s)
}

type GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput struct {
	Duration *string `json:"Duration,omitempty" xml:"Duration,omitempty"`
	Start    *string `json:"Start,omitempty" xml:"Start,omitempty"`
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput) String() string {
	return dara.Prettify(s)
}

func (s GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput) GoString() string {
	return s.String()
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput) GetDuration() *string {
	return s.Duration
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput) GetStart() *string {
	return s.Start
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput) SetDuration(v string) *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput {
	s.Duration = &v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput) SetStart(v string) *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput {
	s.Start = &v
	return s
}

func (s *GetVideoDNAResultResponseBodyAIVideoDNAResultFpShotsFpShotSlicesInput) Validate() error {
	return dara.Validate(s)
}
