// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDistApplicationDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetDistInstanceIds(v *DistApplicationDataResponseBodyDistInstanceIds) *DistApplicationDataResponseBody
	GetDistInstanceIds() *DistApplicationDataResponseBodyDistInstanceIds
	SetDistInstanceTotalCount(v int32) *DistApplicationDataResponseBody
	GetDistInstanceTotalCount() *int32
	SetDistResults(v *DistApplicationDataResponseBodyDistResults) *DistApplicationDataResponseBody
	GetDistResults() *DistApplicationDataResponseBodyDistResults
	SetRequestId(v string) *DistApplicationDataResponseBody
	GetRequestId() *string
}

type DistApplicationDataResponseBody struct {
	// The list of ENS instance IDs.
	DistInstanceIds *DistApplicationDataResponseBodyDistInstanceIds `json:"DistInstanceIds,omitempty" xml:"DistInstanceIds,omitempty" type:"Struct"`
	// The total number of ENS instance IDs.
	//
	// example:
	//
	// 2
	DistInstanceTotalCount *int32 `json:"DistInstanceTotalCount,omitempty" xml:"DistInstanceTotalCount,omitempty"`
	// The distribution result of the data file.
	DistResults *DistApplicationDataResponseBodyDistResults `json:"DistResults,omitempty" xml:"DistResults,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s DistApplicationDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s DistApplicationDataResponseBody) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseBody) GetDistInstanceIds() *DistApplicationDataResponseBodyDistInstanceIds {
	return s.DistInstanceIds
}

func (s *DistApplicationDataResponseBody) GetDistInstanceTotalCount() *int32 {
	return s.DistInstanceTotalCount
}

func (s *DistApplicationDataResponseBody) GetDistResults() *DistApplicationDataResponseBodyDistResults {
	return s.DistResults
}

func (s *DistApplicationDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *DistApplicationDataResponseBody) SetDistInstanceIds(v *DistApplicationDataResponseBodyDistInstanceIds) *DistApplicationDataResponseBody {
	s.DistInstanceIds = v
	return s
}

func (s *DistApplicationDataResponseBody) SetDistInstanceTotalCount(v int32) *DistApplicationDataResponseBody {
	s.DistInstanceTotalCount = &v
	return s
}

func (s *DistApplicationDataResponseBody) SetDistResults(v *DistApplicationDataResponseBodyDistResults) *DistApplicationDataResponseBody {
	s.DistResults = v
	return s
}

func (s *DistApplicationDataResponseBody) SetRequestId(v string) *DistApplicationDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *DistApplicationDataResponseBody) Validate() error {
	if s.DistInstanceIds != nil {
		if err := s.DistInstanceIds.Validate(); err != nil {
			return err
		}
	}
	if s.DistResults != nil {
		if err := s.DistResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type DistApplicationDataResponseBodyDistInstanceIds struct {
	DistInstanceId []*string `json:"DistInstanceId,omitempty" xml:"DistInstanceId,omitempty" type:"Repeated"`
}

func (s DistApplicationDataResponseBodyDistInstanceIds) String() string {
	return dara.Prettify(s)
}

func (s DistApplicationDataResponseBodyDistInstanceIds) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseBodyDistInstanceIds) GetDistInstanceId() []*string {
	return s.DistInstanceId
}

func (s *DistApplicationDataResponseBodyDistInstanceIds) SetDistInstanceId(v []*string) *DistApplicationDataResponseBodyDistInstanceIds {
	s.DistInstanceId = v
	return s
}

func (s *DistApplicationDataResponseBodyDistInstanceIds) Validate() error {
	return dara.Validate(s)
}

type DistApplicationDataResponseBodyDistResults struct {
	DistResult []*DistApplicationDataResponseBodyDistResultsDistResult `json:"DistResult,omitempty" xml:"DistResult,omitempty" type:"Repeated"`
}

func (s DistApplicationDataResponseBodyDistResults) String() string {
	return dara.Prettify(s)
}

func (s DistApplicationDataResponseBodyDistResults) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseBodyDistResults) GetDistResult() []*DistApplicationDataResponseBodyDistResultsDistResult {
	return s.DistResult
}

func (s *DistApplicationDataResponseBodyDistResults) SetDistResult(v []*DistApplicationDataResponseBodyDistResultsDistResult) *DistApplicationDataResponseBodyDistResults {
	s.DistResult = v
	return s
}

func (s *DistApplicationDataResponseBodyDistResults) Validate() error {
	if s.DistResult != nil {
		for _, item := range s.DistResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type DistApplicationDataResponseBodyDistResultsDistResult struct {
	// The name of the data file.
	//
	// example:
	//
	// gcs-prod-websocket-eip-unicom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The error code. The value is of the enumerated data type.
	//
	// example:
	//
	// 400
	ResultCode *int32 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The description of the distribution result.
	//
	// example:
	//
	// Success
	ResultDescrip *string `json:"ResultDescrip,omitempty" xml:"ResultDescrip,omitempty"`
	// The version number of the data file.
	//
	// example:
	//
	// standard
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s DistApplicationDataResponseBodyDistResultsDistResult) String() string {
	return dara.Prettify(s)
}

func (s DistApplicationDataResponseBodyDistResultsDistResult) GoString() string {
	return s.String()
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) GetName() *string {
	return s.Name
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) GetResultDescrip() *string {
	return s.ResultDescrip
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) GetVersion() *string {
	return s.Version
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) SetName(v string) *DistApplicationDataResponseBodyDistResultsDistResult {
	s.Name = &v
	return s
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) SetResultCode(v int32) *DistApplicationDataResponseBodyDistResultsDistResult {
	s.ResultCode = &v
	return s
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) SetResultDescrip(v string) *DistApplicationDataResponseBodyDistResultsDistResult {
	s.ResultDescrip = &v
	return s
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) SetVersion(v string) *DistApplicationDataResponseBodyDistResultsDistResult {
	s.Version = &v
	return s
}

func (s *DistApplicationDataResponseBodyDistResultsDistResult) Validate() error {
	return dara.Validate(s)
}
