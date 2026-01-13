// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobScoreReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataPath(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBody
	GetDataPath() *string
	SetOssPath(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBody
	GetOssPath() *string
	SetReportsOfScoreDiff(v []*ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) *ListFeatureConsistencyCheckJobScoreReportsResponseBody
	GetReportsOfScoreDiff() []*ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff
	SetRequestId(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBody
	GetRequestId() *string
}

type ListFeatureConsistencyCheckJobScoreReportsResponseBody struct {
	// example:
	//
	// http://*******
	DataPath *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	// example:
	//
	// oss://********
	OssPath            *string                                                                     `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	ReportsOfScoreDiff []*ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff `json:"ReportsOfScoreDiff,omitempty" xml:"ReportsOfScoreDiff,omitempty" type:"Repeated"`
	// example:
	//
	// F0AB6527-093F-5C44-B3BD-42C8C210C619
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) GetDataPath() *string {
	return s.DataPath
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) GetOssPath() *string {
	return s.OssPath
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) GetReportsOfScoreDiff() []*ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	return s.ReportsOfScoreDiff
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) SetDataPath(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBody {
	s.DataPath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) SetOssPath(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBody {
	s.OssPath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) SetReportsOfScoreDiff(v []*ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) *ListFeatureConsistencyCheckJobScoreReportsResponseBody {
	s.ReportsOfScoreDiff = v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) SetRequestId(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBody) Validate() error {
	if s.ReportsOfScoreDiff != nil {
		for _, item := range s.ReportsOfScoreDiff {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff struct {
	// example:
	//
	// 4
	LogItemId *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	// example:
	//
	// 323
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// example:
	//
	// 3
	LogUserId *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// example:
	//
	// 0.00000234
	ScoreDiff *string `json:"ScoreDiff,omitempty" xml:"ScoreDiff,omitempty"`
	// example:
	//
	// {}
	ScoreDiffDetail *string `json:"ScoreDiffDetail,omitempty" xml:"ScoreDiffDetail,omitempty"`
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) GetLogItemId() *string {
	return s.LogItemId
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) GetLogUserId() *string {
	return s.LogUserId
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) GetScoreDiff() *string {
	return s.ScoreDiff
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) GetScoreDiffDetail() *string {
	return s.ScoreDiffDetail
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetLogItemId(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.LogItemId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetLogRequestId(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.LogRequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetLogUserId(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.LogUserId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetScoreDiff(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.ScoreDiff = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) SetScoreDiffDetail(v string) *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff {
	s.ScoreDiffDetail = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobScoreReportsResponseBodyReportsOfScoreDiff) Validate() error {
	return dara.Validate(s)
}
