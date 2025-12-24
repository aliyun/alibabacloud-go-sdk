// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetExecutorConfigRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetExecutorConfigRequest
	GetAppName() *string
	SetClusterId(v string) *GetExecutorConfigRequest
	GetClusterId() *string
}

type GetExecutorConfigRequest struct {
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
	// xxljob-a1804a3226d
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetExecutorConfigRequest) String() string {
	return dara.Prettify(s)
}

func (s GetExecutorConfigRequest) GoString() string {
	return s.String()
}

func (s *GetExecutorConfigRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetExecutorConfigRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetExecutorConfigRequest) SetAppName(v string) *GetExecutorConfigRequest {
	s.AppName = &v
	return s
}

func (s *GetExecutorConfigRequest) SetClusterId(v string) *GetExecutorConfigRequest {
	s.ClusterId = &v
	return s
}

func (s *GetExecutorConfigRequest) Validate() error {
	return dara.Validate(s)
}
