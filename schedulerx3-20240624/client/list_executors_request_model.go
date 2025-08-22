// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListExecutorsRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListExecutorsRequest
	GetAppName() *string
	SetClusterId(v string) *ListExecutorsRequest
	GetClusterId() *string
	SetJobId(v int64) *ListExecutorsRequest
	GetJobId() *int64
	SetLabel(v string) *ListExecutorsRequest
	GetLabel() *string
}

type ListExecutorsRequest struct {
	// This parameter is required.
	//
	// example:
	//
	// test-app
	AppName *string `json:"AppName,omitempty" xml:"AppName,omitempty"`
	// This parameter is required.
	//
	// example:
	//
	// xxljob-b6ec1xxxx
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
	// example:
	//
	// gray
	Label *string `json:"Label,omitempty" xml:"Label,omitempty"`
}

func (s ListExecutorsRequest) String() string {
	return dara.Prettify(s)
}

func (s ListExecutorsRequest) GoString() string {
	return s.String()
}

func (s *ListExecutorsRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListExecutorsRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListExecutorsRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListExecutorsRequest) GetLabel() *string {
	return s.Label
}

func (s *ListExecutorsRequest) SetAppName(v string) *ListExecutorsRequest {
	s.AppName = &v
	return s
}

func (s *ListExecutorsRequest) SetClusterId(v string) *ListExecutorsRequest {
	s.ClusterId = &v
	return s
}

func (s *ListExecutorsRequest) SetJobId(v int64) *ListExecutorsRequest {
	s.JobId = &v
	return s
}

func (s *ListExecutorsRequest) SetLabel(v string) *ListExecutorsRequest {
	s.Label = &v
	return s
}

func (s *ListExecutorsRequest) Validate() error {
	return dara.Validate(s)
}
