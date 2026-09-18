// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iCreateTargetScanTaskRequest interface {
	dara.Model
	String() string
	GoString() string
	SetLang(v string) *CreateTargetScanTaskRequest
	GetLang() *string
	SetSampleLevel(v int64) *CreateTargetScanTaskRequest
	GetSampleLevel() *int64
	SetTargetId(v string) *CreateTargetScanTaskRequest
	GetTargetId() *string
}

type CreateTargetScanTaskRequest struct {
	// The language filter for samples, in locale format (such as zh_CN or en_US, which is internally normalized to zh or en). If this parameter is not specified, samples are selected based on the default language policy combined with general-purpose samples.
	//
	// example:
	//
	// zh_CN
	Lang *string `json:"Lang,omitempty" xml:"Lang,omitempty"`
	// The detection intensity. The backend derives the corresponding technique level based on this value. If this parameter is not specified, the system falls back to the scan task configuration saved for the target, and then to the system default value.
	//
	// example:
	//
	// 3
	SampleLevel *int64 `json:"SampleLevel,omitempty" xml:"SampleLevel,omitempty"`
	// The unique identifier of the scan target. The target must have passed connectivity verification (verifyStatus=verified). Otherwise, a 400 error is returned. You can call TestConnectivity to complete the verification first.
	//
	// This parameter is required.
	//
	// example:
	//
	// target-abc123def4567
	TargetId *string `json:"TargetId,omitempty" xml:"TargetId,omitempty"`
}

func (s CreateTargetScanTaskRequest) String() string {
	return dara.Prettify(s)
}

func (s CreateTargetScanTaskRequest) GoString() string {
	return s.String()
}

func (s *CreateTargetScanTaskRequest) GetLang() *string {
	return s.Lang
}

func (s *CreateTargetScanTaskRequest) GetSampleLevel() *int64 {
	return s.SampleLevel
}

func (s *CreateTargetScanTaskRequest) GetTargetId() *string {
	return s.TargetId
}

func (s *CreateTargetScanTaskRequest) SetLang(v string) *CreateTargetScanTaskRequest {
	s.Lang = &v
	return s
}

func (s *CreateTargetScanTaskRequest) SetSampleLevel(v int64) *CreateTargetScanTaskRequest {
	s.SampleLevel = &v
	return s
}

func (s *CreateTargetScanTaskRequest) SetTargetId(v string) *CreateTargetScanTaskRequest {
	s.TargetId = &v
	return s
}

func (s *CreateTargetScanTaskRequest) Validate() error {
	return dara.Validate(s)
}
