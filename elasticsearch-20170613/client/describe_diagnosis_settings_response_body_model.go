// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeDiagnosisSettingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetRequestId(v string) *DescribeDiagnosisSettingsResponseBody
	GetRequestId() *string
	SetResult(v *DescribeDiagnosisSettingsResponseBodyResult) *DescribeDiagnosisSettingsResponseBody
	GetResult() *DescribeDiagnosisSettingsResponseBodyResult
}

type DescribeDiagnosisSettingsResponseBody struct {
	// The ID of the request.
	//
	// example:
	//
	// 5E82B8A8-EED7-4557-A6E9-D1AD3E58****
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
	// The return results.
	Result *DescribeDiagnosisSettingsResponseBodyResult `json:"Result,omitempty" xml:"Result,omitempty" type:"Struct"`
}

func (s DescribeDiagnosisSettingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSettingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeDiagnosisSettingsResponseBody) GetResult() *DescribeDiagnosisSettingsResponseBodyResult {
	return s.Result
}

func (s *DescribeDiagnosisSettingsResponseBody) SetRequestId(v string) *DescribeDiagnosisSettingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBody) SetResult(v *DescribeDiagnosisSettingsResponseBodyResult) *DescribeDiagnosisSettingsResponseBody {
	s.Result = v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBody) Validate() error {
	if s.Result != nil {
		if err := s.Result.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DescribeDiagnosisSettingsResponseBodyResult struct {
	// Scenarios of intelligent maintenance.
	//
	// example:
	//
	// Business Search
	Scene *string `json:"scene,omitempty" xml:"scene,omitempty"`
	// The timestamp of the last update for Intelligent Maintenance scenarios.
	//
	// example:
	//
	// 1588994035385
	UpdateTime *int64 `json:"updateTime,omitempty" xml:"updateTime,omitempty"`
}

func (s DescribeDiagnosisSettingsResponseBodyResult) String() string {
	return dara.Prettify(s)
}

func (s DescribeDiagnosisSettingsResponseBodyResult) GoString() string {
	return s.String()
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) GetScene() *string {
	return s.Scene
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) GetUpdateTime() *int64 {
	return s.UpdateTime
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetScene(v string) *DescribeDiagnosisSettingsResponseBodyResult {
	s.Scene = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) SetUpdateTime(v int64) *DescribeDiagnosisSettingsResponseBodyResult {
	s.UpdateTime = &v
	return s
}

func (s *DescribeDiagnosisSettingsResponseBodyResult) Validate() error {
	return dara.Validate(s)
}
