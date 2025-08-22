// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListLablesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListLablesRequest
	GetAppName() *string
	SetClusterId(v string) *ListLablesRequest
	GetClusterId() *string
	SetJobId(v int64) *ListLablesRequest
	GetJobId() *int64
}

type ListLablesRequest struct {
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
	// 15
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s ListLablesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListLablesRequest) GoString() string {
	return s.String()
}

func (s *ListLablesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListLablesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListLablesRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *ListLablesRequest) SetAppName(v string) *ListLablesRequest {
	s.AppName = &v
	return s
}

func (s *ListLablesRequest) SetClusterId(v string) *ListLablesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListLablesRequest) SetJobId(v int64) *ListLablesRequest {
	s.JobId = &v
	return s
}

func (s *ListLablesRequest) Validate() error {
	return dara.Validate(s)
}
