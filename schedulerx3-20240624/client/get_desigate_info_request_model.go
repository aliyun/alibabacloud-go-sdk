// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetDesigateInfoRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetDesigateInfoRequest
	GetAppName() *string
	SetClusterId(v string) *GetDesigateInfoRequest
	GetClusterId() *string
	SetJobId(v int64) *GetDesigateInfoRequest
	GetJobId() *int64
}

type GetDesigateInfoRequest struct {
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
	// This parameter is required.
	//
	// example:
	//
	// 74
	JobId *int64 `json:"JobId,omitempty" xml:"JobId,omitempty"`
}

func (s GetDesigateInfoRequest) String() string {
	return dara.Prettify(s)
}

func (s GetDesigateInfoRequest) GoString() string {
	return s.String()
}

func (s *GetDesigateInfoRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetDesigateInfoRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetDesigateInfoRequest) GetJobId() *int64 {
	return s.JobId
}

func (s *GetDesigateInfoRequest) SetAppName(v string) *GetDesigateInfoRequest {
	s.AppName = &v
	return s
}

func (s *GetDesigateInfoRequest) SetClusterId(v string) *GetDesigateInfoRequest {
	s.ClusterId = &v
	return s
}

func (s *GetDesigateInfoRequest) SetJobId(v int64) *GetDesigateInfoRequest {
	s.JobId = &v
	return s
}

func (s *GetDesigateInfoRequest) Validate() error {
	return dara.Validate(s)
}
