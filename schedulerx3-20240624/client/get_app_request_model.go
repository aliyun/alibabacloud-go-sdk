// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iGetAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *GetAppRequest
	GetAppName() *string
	SetClusterId(v string) *GetAppRequest
	GetClusterId() *string
}

type GetAppRequest struct {
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
	// xxljob-d6a5243b6fa
	ClusterId *string `json:"ClusterId,omitempty" xml:"ClusterId,omitempty"`
}

func (s GetAppRequest) String() string {
	return dara.Prettify(s)
}

func (s GetAppRequest) GoString() string {
	return s.String()
}

func (s *GetAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *GetAppRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *GetAppRequest) SetAppName(v string) *GetAppRequest {
	s.AppName = &v
	return s
}

func (s *GetAppRequest) SetClusterId(v string) *GetAppRequest {
	s.ClusterId = &v
	return s
}

func (s *GetAppRequest) Validate() error {
	return dara.Validate(s)
}
