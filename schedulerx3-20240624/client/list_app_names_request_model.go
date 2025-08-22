// This file is auto-generated, don't edit it. Thanks.
package client

import (
	"github.com/alibabacloud-go/tea/dara"
)

type iListAppNamesRequest interface {
	dara.Model
	String() string
	GoString() string
	SetAppName(v string) *ListAppNamesRequest
	GetAppName() *string
	SetClusterId(v string) *ListAppNamesRequest
	GetClusterId() *string
}

type ListAppNamesRequest struct {
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

func (s ListAppNamesRequest) String() string {
	return dara.Prettify(s)
}

func (s ListAppNamesRequest) GoString() string {
	return s.String()
}

func (s *ListAppNamesRequest) GetAppName() *string {
	return s.AppName
}

func (s *ListAppNamesRequest) GetClusterId() *string {
	return s.ClusterId
}

func (s *ListAppNamesRequest) SetAppName(v string) *ListAppNamesRequest {
	s.AppName = &v
	return s
}

func (s *ListAppNamesRequest) SetClusterId(v string) *ListAppNamesRequest {
	s.ClusterId = &v
	return s
}

func (s *ListAppNamesRequest) Validate() error {
	return dara.Validate(s)
}
