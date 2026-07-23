// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListFeatureConsistencyCheckJobFeatureReportsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDataPath(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody
	GetDataPath() *string
	SetOssPath(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody
	GetOssPath() *string
	SetReportsOfFeatureDiff(v []*ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody
	GetReportsOfFeatureDiff() []*ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff
	SetRequestId(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody
	GetRequestId() *string
}

type ListFeatureConsistencyCheckJobFeatureReportsResponseBody struct {
	// The data path.
	//
	// example:
	//
	// https://********
	DataPath *string `json:"DataPath,omitempty" xml:"DataPath,omitempty"`
	// The OSS path.
	//
	// example:
	//
	// oss://********
	OssPath *string `json:"OssPath,omitempty" xml:"OssPath,omitempty"`
	// The feature comparison reports.
	ReportsOfFeatureDiff []*ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff `json:"ReportsOfFeatureDiff,omitempty" xml:"ReportsOfFeatureDiff,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// BBD41FBF-E75C-551A-92FA-CAD654AA006F
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponseBody) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) GetDataPath() *string {
	return s.DataPath
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) GetOssPath() *string {
	return s.OssPath
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) GetReportsOfFeatureDiff() []*ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	return s.ReportsOfFeatureDiff
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) SetDataPath(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody {
	s.DataPath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) SetOssPath(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody {
	s.OssPath = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) SetReportsOfFeatureDiff(v []*ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody {
	s.ReportsOfFeatureDiff = v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) SetRequestId(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBody {
	s.RequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBody) Validate() error {
	if s.ReportsOfFeatureDiff != nil {
		for _, item := range s.ReportsOfFeatureDiff {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff struct {
	// The feature name.
	//
	// example:
	//
	// gender
	FeatureName *string `json:"FeatureName,omitempty" xml:"FeatureName,omitempty"`
	// The item ID.
	//
	// example:
	//
	// 9010
	LogItemId *string `json:"LogItemId,omitempty" xml:"LogItemId,omitempty"`
	// The request ID.
	//
	// example:
	//
	// F7AC05FF-EDE7-5C2B-B9AE-33D6DF4178BA
	LogRequestId *string `json:"LogRequestId,omitempty" xml:"LogRequestId,omitempty"`
	// The user ID.
	//
	// example:
	//
	// 1010
	LogUserId *string `json:"LogUserId,omitempty" xml:"LogUserId,omitempty"`
	// The offline value.
	//
	// example:
	//
	// male
	OfflineValue *string `json:"OfflineValue,omitempty" xml:"OfflineValue,omitempty"`
	// The online value.
	//
	// example:
	//
	// male
	OnlineValue *string `json:"OnlineValue,omitempty" xml:"OnlineValue,omitempty"`
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) String() string {
	return dara.Prettify(s)
}

func (s ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) GoString() string {
	return s.String()
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) GetFeatureName() *string {
	return s.FeatureName
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) GetLogItemId() *string {
	return s.LogItemId
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) GetLogRequestId() *string {
	return s.LogRequestId
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) GetLogUserId() *string {
	return s.LogUserId
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) GetOfflineValue() *string {
	return s.OfflineValue
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) GetOnlineValue() *string {
	return s.OnlineValue
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetFeatureName(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.FeatureName = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetLogItemId(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.LogItemId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetLogRequestId(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.LogRequestId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetLogUserId(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.LogUserId = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetOfflineValue(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.OfflineValue = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) SetOnlineValue(v string) *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff {
	s.OnlineValue = &v
	return s
}

func (s *ListFeatureConsistencyCheckJobFeatureReportsResponseBodyReportsOfFeatureDiff) Validate() error {
	return dara.Validate(s)
}
