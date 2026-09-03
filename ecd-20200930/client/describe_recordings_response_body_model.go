// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDescribeRecordingsResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetNextToken(v string) *DescribeRecordingsResponseBody
	GetNextToken() *string
	SetRecordings(v []*DescribeRecordingsResponseBodyRecordings) *DescribeRecordingsResponseBody
	GetRecordings() []*DescribeRecordingsResponseBodyRecordings
	SetRequestId(v string) *DescribeRecordingsResponseBody
	GetRequestId() *string
}

type DescribeRecordingsResponseBody struct {
	// The pagination token for the next query. If NextToken is empty, no more results exist.
	//
	// example:
	//
	// AAAAAV3MpHK1AP0pfERHZN5pu6nbCQ7ar+fECeh1IuWQXi39R5eoJ68zWp99mTAKRRNRhw==
	NextToken *string `json:"NextToken,omitempty" xml:"NextToken,omitempty"`
	// The collection of cloud computer screen recording information.
	Recordings []*DescribeRecordingsResponseBodyRecordings `json:"Recordings,omitempty" xml:"Recordings,omitempty" type:"Repeated"`
	// The request ID.
	//
	// example:
	//
	// 327CFE78-1C0D-51AC-A9C6-BCEDF0DD44D6
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DescribeRecordingsResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordingsResponseBody) GoString() string {
	return s.String()
}

func (s *DescribeRecordingsResponseBody) GetNextToken() *string {
	return s.NextToken
}

func (s *DescribeRecordingsResponseBody) GetRecordings() []*DescribeRecordingsResponseBodyRecordings {
	return s.Recordings
}

func (s *DescribeRecordingsResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DescribeRecordingsResponseBody) SetNextToken(v string) *DescribeRecordingsResponseBody {
	s.NextToken = &v
	return s
}

func (s *DescribeRecordingsResponseBody) SetRecordings(v []*DescribeRecordingsResponseBodyRecordings) *DescribeRecordingsResponseBody {
	s.Recordings = v
	return s
}

func (s *DescribeRecordingsResponseBody) SetRequestId(v string) *DescribeRecordingsResponseBody {
	s.RequestId = &v
	return s
}

func (s *DescribeRecordingsResponseBody) Validate() error {
	if s.Recordings != nil {
		for _, item := range s.Recordings {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DescribeRecordingsResponseBodyRecordings struct {
	// The cloud computer ID.
	//
	// example:
	//
	// ecd-10v0vuvm616sk****
	DesktopId *string `json:"DesktopId,omitempty" xml:"DesktopId,omitempty"`
	// The cloud computer name.
	//
	// example:
	//
	// DemoComputer
	DesktopName *string `json:"DesktopName,omitempty" xml:"DesktopName,omitempty"`
	// The end time of the recording. The time is in the ISO 8601 standard in UTC: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-04-10T07:26:06Z
	EndTime *string `json:"EndTime,omitempty" xml:"EndTime,omitempty"`
	// The list of end user IDs.
	EndUserIds []*string `json:"EndUserIds,omitempty" xml:"EndUserIds,omitempty" type:"Repeated"`
	// The file path.
	FilePath *string `json:"FilePath,omitempty" xml:"FilePath,omitempty"`
	// The policy ID.
	//
	// example:
	//
	// pg-6dn811rzrwh9w****
	PolicyGroupId *string `json:"PolicyGroupId,omitempty" xml:"PolicyGroupId,omitempty"`
	// The size of the recording file. Unit: bytes.
	//
	// example:
	//
	// 1742845
	RecordingSize *int32 `json:"RecordingSize,omitempty" xml:"RecordingSize,omitempty"`
	// The recording type.
	//
	// example:
	//
	// alltime
	RecordingType *string `json:"RecordingType,omitempty" xml:"RecordingType,omitempty"`
	// The download URL of the screen recording file.
	SignedUrl *string `json:"SignedUrl,omitempty" xml:"SignedUrl,omitempty"`
	// The start time of the recording. The time is in the ISO 8601 standard in UTC: yyyy-MM-ddTHH:mm:ssZ.
	//
	// example:
	//
	// 2023-04-10T07:26:06Z
	StartTime *string `json:"StartTime,omitempty" xml:"StartTime,omitempty"`
}

func (s DescribeRecordingsResponseBodyRecordings) String() string {
	return dara.Prettify(s)
}

func (s DescribeRecordingsResponseBodyRecordings) GoString() string {
	return s.String()
}

func (s *DescribeRecordingsResponseBodyRecordings) GetDesktopId() *string {
	return s.DesktopId
}

func (s *DescribeRecordingsResponseBodyRecordings) GetDesktopName() *string {
	return s.DesktopName
}

func (s *DescribeRecordingsResponseBodyRecordings) GetEndTime() *string {
	return s.EndTime
}

func (s *DescribeRecordingsResponseBodyRecordings) GetEndUserIds() []*string {
	return s.EndUserIds
}

func (s *DescribeRecordingsResponseBodyRecordings) GetFilePath() *string {
	return s.FilePath
}

func (s *DescribeRecordingsResponseBodyRecordings) GetPolicyGroupId() *string {
	return s.PolicyGroupId
}

func (s *DescribeRecordingsResponseBodyRecordings) GetRecordingSize() *int32 {
	return s.RecordingSize
}

func (s *DescribeRecordingsResponseBodyRecordings) GetRecordingType() *string {
	return s.RecordingType
}

func (s *DescribeRecordingsResponseBodyRecordings) GetSignedUrl() *string {
	return s.SignedUrl
}

func (s *DescribeRecordingsResponseBodyRecordings) GetStartTime() *string {
	return s.StartTime
}

func (s *DescribeRecordingsResponseBodyRecordings) SetDesktopId(v string) *DescribeRecordingsResponseBodyRecordings {
	s.DesktopId = &v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) SetDesktopName(v string) *DescribeRecordingsResponseBodyRecordings {
	s.DesktopName = &v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) SetEndTime(v string) *DescribeRecordingsResponseBodyRecordings {
	s.EndTime = &v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) SetEndUserIds(v []*string) *DescribeRecordingsResponseBodyRecordings {
	s.EndUserIds = v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) SetFilePath(v string) *DescribeRecordingsResponseBodyRecordings {
	s.FilePath = &v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) SetPolicyGroupId(v string) *DescribeRecordingsResponseBodyRecordings {
	s.PolicyGroupId = &v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) SetRecordingSize(v int32) *DescribeRecordingsResponseBodyRecordings {
	s.RecordingSize = &v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) SetRecordingType(v string) *DescribeRecordingsResponseBodyRecordings {
	s.RecordingType = &v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) SetSignedUrl(v string) *DescribeRecordingsResponseBodyRecordings {
	s.SignedUrl = &v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) SetStartTime(v string) *DescribeRecordingsResponseBodyRecordings {
	s.StartTime = &v
	return s
}

func (s *DescribeRecordingsResponseBodyRecordings) Validate() error {
	return dara.Validate(s)
}
