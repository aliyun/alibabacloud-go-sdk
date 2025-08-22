// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iDeleteAppRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *DeleteAppRequest
	GetAppName() *string
	SetClusterId(v string) *DeleteAppRequest
	GetClusterId() *string
}

type DeleteAppRequest struct {
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
}

func (s DeleteAppRequest) String() string {
	return dara.Prettify(s)
}

func (s DeleteAppRequest) GoString() string {
	return s.String()
}

func (s *DeleteAppRequest) GetAppName() *string {
	return s.AppName
}

func (s *DeleteAppRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *DeleteAppRequest) SetAppName(v string) *DeleteAppRequest {
	s.AppName = &v
	return s
}

func (s *DeleteAppRequest) SetClusterId(v string) *DeleteAppRequest {
	s.ClusterId = &v
	return s
}

func (s *DeleteAppRequest) Validate() error {
	return dara.Validate(s)
}
