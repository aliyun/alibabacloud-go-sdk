// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iPushApplicationDataResponseBody interface {
	dara.Model
	String() string
	GoString() string
	SetPushResults(v *PushApplicationDataResponseBodyPushResults) *PushApplicationDataResponseBody
	GetPushResults() *PushApplicationDataResponseBodyPushResults
	SetRequestId(v string) *PushApplicationDataResponseBody
	GetRequestId() *string
}

type PushApplicationDataResponseBody struct {
	// The push results of data files.
	PushResults *PushApplicationDataResponseBodyPushResults `json:"PushResults,omitempty" xml:"PushResults,omitempty" type:"Struct"`
	// The ID of the request.
	//
	// example:
	//
	// 473469C7-AA6F-4DC5-B3DB-A3DC0DE3C83E
	RequestId *string `json:"RequestId,omitempty" xml:"RequestId,omitempty"`
}

func (s PushApplicationDataResponseBody) String() string {
	return dara.Prettify(s)
}

func (s PushApplicationDataResponseBody) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponseBody) GetPushResults() *PushApplicationDataResponseBodyPushResults {
	return s.PushResults
}

func (s *PushApplicationDataResponseBody) GetRequestId() *string {
	return s.RequestId
}

func (s *PushApplicationDataResponseBody) SetPushResults(v *PushApplicationDataResponseBodyPushResults) *PushApplicationDataResponseBody {
	s.PushResults = v
	return s
}

func (s *PushApplicationDataResponseBody) SetRequestId(v string) *PushApplicationDataResponseBody {
	s.RequestId = &v
	return s
}

func (s *PushApplicationDataResponseBody) Validate() error {
	if s.PushResults != nil {
		if err := s.PushResults.Validate(); err != nil {
			return err
		}
	}
	return nil
}

type PushApplicationDataResponseBodyPushResults struct {
	PushResult []*PushApplicationDataResponseBodyPushResultsPushResult `json:"PushResult,omitempty" xml:"PushResult,omitempty" type:"Repeated"`
}

func (s PushApplicationDataResponseBodyPushResults) String() string {
	return dara.Prettify(s)
}

func (s PushApplicationDataResponseBodyPushResults) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponseBodyPushResults) GetPushResult() []*PushApplicationDataResponseBodyPushResultsPushResult {
	return s.PushResult
}

func (s *PushApplicationDataResponseBodyPushResults) SetPushResult(v []*PushApplicationDataResponseBodyPushResultsPushResult) *PushApplicationDataResponseBodyPushResults {
	s.PushResult = v
	return s
}

func (s *PushApplicationDataResponseBodyPushResults) Validate() error {
	if s.PushResult != nil {
		for _, item := range s.PushResult {
			if item != nil {
				if err := item.Validate(); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

type PushApplicationDataResponseBodyPushResultsPushResult struct {
	// The name of the data file.
	//
	// example:
	//
	// gcs-pre-websocket-eip-telecom
	Name *string `json:"Name,omitempty" xml:"Name,omitempty"`
	// The push result. The value is of the enumeration type. Valid values:
	//
	// 	- 0: The push operation is successful.
	//
	// 	- 100: The push operation has been performed and the file is pushed.
	//
	// 	- 200: The push operation has been performed and the file is being pushed to specific file servers.
	//
	// 	- 300: The push operation failed. You must trigger the push operation again. The ResultDescrip parameter indicates the error description.
	//
	// example:
	//
	// 100
	ResultCode *int32 `json:"ResultCode,omitempty" xml:"ResultCode,omitempty"`
	// The description of the push status.
	//
	// example:
	//
	// success
	ResultDescrip *string `json:"ResultDescrip,omitempty" xml:"ResultDescrip,omitempty"`
	// The version number of the data file.
	//
	// example:
	//
	// standard
	Version *string `json:"Version,omitempty" xml:"Version,omitempty"`
}

func (s PushApplicationDataResponseBodyPushResultsPushResult) String() string {
	return dara.Prettify(s)
}

func (s PushApplicationDataResponseBodyPushResultsPushResult) GoString() string {
	return s.String()
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) GetName() *string {
	return s.Name
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) GetResultCode() *int32 {
	return s.ResultCode
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) GetResultDescrip() *string {
	return s.ResultDescrip
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) GetVersion() *string {
	return s.Version
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) SetName(v string) *PushApplicationDataResponseBodyPushResultsPushResult {
	s.Name = &v
	return s
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) SetResultCode(v int32) *PushApplicationDataResponseBodyPushResultsPushResult {
	s.ResultCode = &v
	return s
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) SetResultDescrip(v string) *PushApplicationDataResponseBodyPushResultsPushResult {
	s.ResultDescrip = &v
	return s
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) SetVersion(v string) *PushApplicationDataResponseBodyPushResultsPushResult {
	s.Version = &v
	return s
}

func (s *PushApplicationDataResponseBodyPushResultsPushResult) Validate() error {
	return dara.Validate(s)
}
